// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*deviceInterfaceResource)(nil)
	_ resource.ResourceWithConfigure   = (*deviceInterfaceResource)(nil)
	_ resource.ResourceWithImportState = (*deviceInterfaceResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*deviceInterfaceResource)(nil)
)

// NewDeviceInterfaceResource returns a new device_interface resource.
func NewDeviceInterfaceResource() resource.Resource {
	return &deviceInterfaceResource{}
}

type deviceInterfaceResource struct {
	client *netboxapi.Client
}

func (r *deviceInterfaceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_interface"
}

func (r *deviceInterfaceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A device interface (dcim.interface).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/interface/):\n\n> Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"device_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the device.",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Interface type as its NetBox slug, e.g. virtual, lag, bridge, 1000base-t, 10gbase-x-sfpp (any of NetBox's interface type choices).",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the interface is enabled.",
				Default:     booldefault.StaticBool(true),
			},
			"mgmt_only": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the interface is used for out-of-band management only.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
				Default: booldefault.StaticBool(false),
			},
			"mgmtonly": schema.BoolAttribute{
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use mgmt_only instead.",
				Description:        "Deprecated alias of mgmt_only.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"mark_connected": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat as if a cable is connected.",
				Default:     booldefault.StaticBool(false),
			},
			"mtu": schema.Int64Attribute{
				Optional: true,
				Validators: []validator.Int64{
					int64validator.Between(1, 65536),
				},
			},
			"speed": schema.Int64Attribute{
				Optional:    true,
				Description: "Speed in kbps.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"duplex": schema.StringAttribute{
				Optional:    true,
				Description: "One of: half, full, auto.",
				Validators: []validator.String{
					stringvalidator.OneOf("half", "full", "auto"),
				},
			},
			"wwn": schema.StringAttribute{
				Optional:    true,
				Description: "64-bit World Wide Name.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Description: "802.1Q tagging mode. One of: access, tagged, tagged-all, q-in-q.",
				Validators: []validator.String{
					stringvalidator.OneOf("access", "tagged", "tagged-all", "q-in-q"),
				},
			},
			"untagged_vlan_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the untagged VLAN.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"untagged_vlan": schema.Int64Attribute{
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use untagged_vlan_id instead.",
				Description:        "Deprecated alias of untagged_vlan_id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"tagged_vlan_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Computed:    true,
				Description: "Ids of the tagged VLANs.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"tagged_vlans": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use tagged_vlan_ids instead.",
				Description:        "Deprecated alias of tagged_vlan_ids.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"module_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the installed module this interface belongs to.",
			},
			"lag_device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the parent LAG interface.",
			},
			"parent_device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the parent interface.",
			},
			"bridge_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the bridge interface.",
			},
			"vrf_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the VRF.",
			},
			"vdc_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Optional:    true,
				Description: "Ids of the virtual device contexts.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Slugs of the tags assigned to the object (the provider's default_tags are added on top, see tags_all).",
			},
			"tags_all": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of all tags on the object, including the provider's default_tags.",
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *deviceInterfaceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

// ModifyPlan runs on every plan, after Terraform builds the proposed new state. It is the
// resource-level hook for changing the plan itself: filling computed attributes already known at
// plan time (avoiding "known after apply" noise), or vetoing the plan with diagnostics. It is
// generated when something needs it, each part only where the spec asks for it: a backend plan
// fragment, the companion file's modify_plan hook, the sync that keeps attributes marked as
// deprecated equal to their canonical ones and the stage that decides the volatile attributes.
func (r *deviceInterfaceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var plan deviceInterfaceResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if r.client != nil && !plan.Tags.IsUnknown() {
			configured := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
			plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, netboxapi.MergeTags(configured, r.client.DefaultTags), false, &resp.Diagnostics)
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}()
	if resp.Diagnostics.HasError() {
		return
	}
	// Alias sync, after the backend fragment and the companion hook (both rewrite resp.Plan
	// wholesale from req.Plan, so an earlier sync would be clobbered): fold each configured
	// deprecated alias into its canonical attribute and keep the two equal in the plan.
	if !resp.Diagnostics.HasError() && !resp.Plan.Raw.IsNull() {
		var config, plan deviceInterfaceResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		switch {
		case !config.Mgmtonly.IsNull() && !config.MgmtOnly.IsNull():
			if !config.Mgmtonly.IsUnknown() && !config.MgmtOnly.IsUnknown() && !config.Mgmtonly.Equal(config.MgmtOnly) {
				resp.Diagnostics.AddError("Conflicting attribute values", "mgmtonly is a deprecated alias of mgmt_only; both are set and the values differ.")
			}
		case !config.Mgmtonly.IsNull():
			plan.MgmtOnly = plan.Mgmtonly
		case !config.MgmtOnly.IsNull():
			plan.Mgmtonly = plan.MgmtOnly
		default:
			plan.Mgmtonly = plan.MgmtOnly
		}
		switch {
		case !config.UntaggedVlan.IsNull() && !config.UntaggedVlanID.IsNull():
			if !config.UntaggedVlan.IsUnknown() && !config.UntaggedVlanID.IsUnknown() && !config.UntaggedVlan.Equal(config.UntaggedVlanID) {
				resp.Diagnostics.AddError("Conflicting attribute values", "untagged_vlan is a deprecated alias of untagged_vlan_id; both are set and the values differ.")
			}
		case !config.UntaggedVlan.IsNull():
			plan.UntaggedVlanID = plan.UntaggedVlan
		case !config.UntaggedVlanID.IsNull():
			plan.UntaggedVlan = plan.UntaggedVlanID
		default:
			// untagged_vlan_id was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.UntaggedVlanID = types.Int64Null()
			plan.UntaggedVlan = plan.UntaggedVlanID
		}
		switch {
		case !config.TaggedVlans.IsNull() && !config.TaggedVlanIds.IsNull():
			if !config.TaggedVlans.IsUnknown() && !config.TaggedVlanIds.IsUnknown() && !config.TaggedVlans.Equal(config.TaggedVlanIds) {
				resp.Diagnostics.AddError("Conflicting attribute values", "tagged_vlans is a deprecated alias of tagged_vlan_ids; both are set and the values differ.")
			}
		case !config.TaggedVlans.IsNull():
			plan.TaggedVlanIds = plan.TaggedVlans
		case !config.TaggedVlanIds.IsNull():
			plan.TaggedVlans = plan.TaggedVlanIds
		default:
			// tagged_vlan_ids was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.TaggedVlanIds = types.SetNull(types.Int64Type)
			plan.TaggedVlans = plan.TaggedVlanIds
		}
		if resp.Diagnostics.HasError() {
			return
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state deviceInterfaceResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *deviceInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data deviceInterfaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandDeviceInterface(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device_interface custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Dcim.DcimInterfacesCreateContext(ctx, dcim.NewDcimInterfacesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_device_interface", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceInterface(ctx, netboxapi.DeviceInterfaceResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *deviceInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data deviceInterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Dcim.DcimInterfacesRetrieveContext(ctx, dcim.NewDcimInterfacesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device_interface", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceInterface(ctx, netboxapi.DeviceInterfaceResponseDTOFromGoNetbox(res.Payload), state)...)
	allTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	state.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	state.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	state.CustomFields = keepPriorCustomFields(priorCustomFields, state.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *deviceInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data deviceInterfaceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandDeviceInterface(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior deviceInterfaceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device_interface custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Dcim.DcimInterfacesUpdateContext(ctx, dcim.NewDcimInterfacesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_device_interface", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceInterface(ctx, netboxapi.DeviceInterfaceResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *deviceInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data deviceInterfaceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimInterfacesDestroyContext(ctx, dcim.NewDcimInterfacesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_device_interface", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Expected an integer id, got %q.", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
