package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// netbox_device_oob_ip is the out-of-band counterpart of netbox_primary_ip: NetBox only
// accepts an oob_ip assigned to one of the device's own interfaces, so the address is set by a
// separate resource after it exists (device -> interface -> ip_address -> device would cycle).
// Virtual machines have no oob_ip, so this one takes devices only.

func init() {
	companionResources = append(companionResources, NewDeviceOOBIPResource)
}

var (
	_ resource.Resource                = (*deviceOOBIPResource)(nil)
	_ resource.ResourceWithConfigure   = (*deviceOOBIPResource)(nil)
	_ resource.ResourceWithImportState = (*deviceOOBIPResource)(nil)
)

func NewDeviceOOBIPResource() resource.Resource {
	return &deviceOOBIPResource{}
}

type deviceOOBIPResource struct {
	client *netboxapi.Client
}

type deviceOOBIPResourceModel struct {
	ID          types.String `tfsdk:"id"`
	DeviceID    types.Int64  `tfsdk:"device_id"`
	IPAddressID types.Int64  `tfsdk:"ip_address_id"`
}

func (resource *deviceOOBIPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_oob_ip"
}

func (resource *deviceOOBIPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Sets the out-of-band management IP address of a device. The address must be assigned to one of the device's interfaces. Import with the device id.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The device id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"device_id": schema.Int64Attribute{
				Required:      true,
				Description:   "Id of the device.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"ip_address_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the IP address to use for out-of-band management.",
			},
		},
	}
}

func (resource *deviceOOBIPResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (resource *deviceOOBIPResource) set(ctx context.Context, plan *deviceOOBIPResourceModel) error {
	ipID := plan.IPAddressID.ValueInt64()
	if err := patchDevice(ctx, resource.client, plan.DeviceID.ValueInt64(), map[string]any{"oob_ip": ipID}); err != nil {
		return err
	}
	plan.ID = types.StringValue(strconv.FormatInt(plan.DeviceID.ValueInt64(), 10))
	return nil
}

func (resource *deviceOOBIPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan deviceOOBIPResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := resource.set(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error setting the out-of-band IP of the device", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *deviceOOBIPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state deviceOOBIPResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	id, err := strconv.ParseInt(state.ID.ValueString(), 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid netbox_device_oob_ip id", fmt.Sprintf("Expected the device id, got %q.", state.ID.ValueString()))
		return
	}
	res, err := resource.client.Dcim.DcimDevicesRetrieveContext(ctx, dcim.NewDcimDevicesRetrieveParams().WithID(id), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading the device", err.Error())
		return
	}
	if res.Payload.OobIP == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.DeviceID = types.Int64Value(id)
	state.IPAddressID = types.Int64Value(res.Payload.OobIP.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (resource *deviceOOBIPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan deviceOOBIPResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := resource.set(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error setting the out-of-band IP of the device", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *deviceOOBIPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state deviceOOBIPResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := patchDevice(ctx, resource.client, state.DeviceID.ValueInt64(), map[string]any{"oob_ip": nil})
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error clearing the out-of-band IP of the device", err.Error())
	}
}

func (resource *deviceOOBIPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("Expected the device id, got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &deviceOOBIPResourceModel{
		ID: types.StringValue(req.ID), DeviceID: types.Int64Value(id), IPAddressID: types.Int64Null(),
	})...)
}
