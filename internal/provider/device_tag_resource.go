package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// netbox_device_tag is a hand-written companion resource: it owns the presence of one tag on a
// device without owning the device, the way aws_ec2_tag does for AWS resources.

func init() {
	companionResources = append(companionResources, NewDeviceTagResource)
}

var (
	_ resource.Resource                = (*deviceTagResource)(nil)
	_ resource.ResourceWithConfigure   = (*deviceTagResource)(nil)
	_ resource.ResourceWithImportState = (*deviceTagResource)(nil)
)

func NewDeviceTagResource() resource.Resource {
	return &deviceTagResource{}
}

type deviceTagResource struct {
	client *netboxapi.Client
}

type deviceTagResourceModel struct {
	ID       types.String `tfsdk:"id"`
	DeviceID types.Int64  `tfsdk:"device_id"`
	TagSlug  types.String `tfsdk:"tag_slug"`
}

func (resource *deviceTagResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_tag"
}

func (resource *deviceTagResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Attaches a single tag to a device, like aws_ec2_tag in the AWS provider: use it to tag a device that is created and managed outside of Terraform (onboarding automation, the NetBox UI) without taking ownership of the rest of the device. A netbox_device resource for the same device fights this resource — its updates strip every tag its own configuration does not declare — unless it sets lifecycle { ignore_changes = [tags] }. Import as <device id>/<tag slug>.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "<device id>/<tag slug>.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"device_id": schema.Int64Attribute{
				Required:      true,
				Description:   "Id of the device.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"tag_slug": schema.StringAttribute{
				Required:      true,
				Description:   "Slug of the tag. The tag must already exist.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (resource *deviceTagResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected resource configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	resource.client = client
}

func deviceTagID(deviceID int64, slug string) string {
	return strconv.FormatInt(deviceID, 10) + "/" + slug
}

func parseDeviceTagID(s string) (int64, string, error) {
	rest, slug, ok := strings.Cut(s, "/")
	id, err := strconv.ParseInt(rest, 10, 64)
	if !ok || slug == "" || err != nil {
		return 0, "", fmt.Errorf("expected <device id>/<tag slug>, got %q", s)
	}
	return id, slug, nil
}

// modifyDeviceTags rewrites the device's tag list as mutate(current). It holds the device's patch
// lock (deviceLocks, shared with patchDevice) across the whole read-modify-write, so concurrent
// netbox_device_tag resources on the same device cannot interleave and resurrect a stale list.
// A mutate that reports no change skips the write.
func modifyDeviceTags(ctx context.Context, client *netboxapi.Client, deviceID int64, mutate func(current []string) (desired []string, changed bool)) error {
	mutex, _ := deviceLocks.LoadOrStore(deviceID, &sync.Mutex{})
	mutex.(*sync.Mutex).Lock()
	defer mutex.(*sync.Mutex).Unlock()
	res, err := client.Dcim.DcimDevicesRetrieveContext(ctx, dcim.NewDcimDevicesRetrieveParams().WithID(deviceID), nil)
	if err != nil {
		return err
	}
	current := make([]string, 0, len(res.Payload.Tags))
	for _, tag := range res.Payload.Tags {
		if tag.Slug != nil {
			current = append(current, *tag.Slug)
		}
	}
	desired, changed := mutate(current)
	if !changed {
		return nil
	}
	refs := make([]map[string]string, 0, len(desired))
	for _, slug := range desired {
		refs = append(refs, map[string]string{"slug": slug})
	}
	params := dcim.NewDcimDevicesPartialUpdateParams().WithID(deviceID)
	_, err = client.Dcim.DcimDevicesPartialUpdateContext(ctx, params, nil, netboxapi.WithBody(map[string]any{"tags": refs}))
	return err
}

func (resource *deviceTagResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan deviceTagResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deviceID := plan.DeviceID.ValueInt64()
	slug := plan.TagSlug.ValueString()
	if err := netboxapi.CheckTags(ctx, resource.client.NetBoxAPI, []string{slug}); err != nil {
		resp.Diagnostics.AddError("Error creating netbox_device_tag", err.Error())
		return
	}
	err := modifyDeviceTags(ctx, resource.client, deviceID, func(current []string) ([]string, bool) {
		for _, currentSlug := range current {
			if currentSlug == slug {
				return nil, false // already tagged, adopt it
			}
		}
		return append(current, slug), true
	})
	if err != nil {
		resp.Diagnostics.AddError("Error tagging the device", err.Error())
		return
	}
	plan.ID = types.StringValue(deviceTagID(deviceID, slug))
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *deviceTagResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceTagResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deviceID, slug, err := parseDeviceTagID(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid netbox_device_tag id", err.Error())
		return
	}
	res, err := resource.client.Dcim.DcimDevicesRetrieveContext(ctx, dcim.NewDcimDevicesRetrieveParams().WithID(deviceID), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading the device", err.Error())
		return
	}
	for _, tag := range res.Payload.Tags {
		if tag.Slug != nil && *tag.Slug == slug {
			state.DeviceID = types.Int64Value(deviceID)
			state.TagSlug = types.StringValue(slug)
			resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
			return
		}
	}
	// the tag was removed from the device outside of Terraform
	resp.State.RemoveResource(ctx)
}

// Update is unreachable: both attributes require replacement.
func (resource *deviceTagResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan deviceTagResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *deviceTagResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state deviceTagResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deviceID, slug, err := parseDeviceTagID(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid netbox_device_tag id", err.Error())
		return
	}
	err = modifyDeviceTags(ctx, resource.client, deviceID, func(current []string) ([]string, bool) {
		desired := make([]string, 0, len(current))
		for _, currentSlug := range current {
			if currentSlug != slug {
				desired = append(desired, currentSlug)
			}
		}
		return desired, len(desired) != len(current)
	})
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error untagging the device", err.Error())
	}
}

func (resource *deviceTagResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	deviceID, slug, err := parseDeviceTagID(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &deviceTagResourceModel{
		ID:       types.StringValue(req.ID),
		DeviceID: types.Int64Value(deviceID),
		TagSlug:  types.StringValue(slug),
	})...)
}
