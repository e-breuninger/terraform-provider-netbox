// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*deviceResource)(nil)
	_ resource.ResourceWithConfigure   = (*deviceResource)(nil)
	_ resource.ResourceWithImportState = (*deviceResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*deviceResource)(nil)
)

// NewDeviceResource returns a new device resource.
func NewDeviceResource() resource.Resource {
	return &deviceResource{}
}

type deviceResource struct {
	client *netboxapi.Client
}

func (r *deviceResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (r *deviceResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A physical device (dcim.device). primary_ip4_id/primary_ip6_id are set through netbox_primary_ip.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/device/):\n\n> Every piece of hardware which is installed within a site or rack exists in NetBox as a device. Devices are measured in rack units (U) and can be half depth or full depth. A device may have a height of 0U: These devices do not consume vertical rack space and cannot be assigned to a particular rack unit. A common example of a 0U device is a vertically-mounted PDU.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Description: "Device name; NetBox allows unnamed devices.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"device_type_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the device type.",
			},
			"role_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the device role.",
			},
			"site_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the site.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the location within the site.",
			},
			"rack_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the rack.",
			},
			"rack_position": schema.Float64Attribute{
				Optional:    true,
				Description: "Lowest rack unit occupied by the device; half units (e.g. 1.5) are allowed.",
				Validators: []validator.Float64{
					float64validator.Between(0.5, 1000),
				},
			},
			"rack_face": schema.StringAttribute{
				Optional:    true,
				Description: "Rack face the device is mounted on. One of: front, rear.",
				Validators: []validator.String{
					stringvalidator.OneOf("front", "rear"),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: offline, active, planned, staged, failed, inventory, decommissioning.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("offline", "active", "planned", "staged", "failed", "inventory", "decommissioning"),
				},
			},
			"airflow": schema.StringAttribute{
				Optional:    true,
				Description: "One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
				Validators: []validator.String{
					stringvalidator.OneOf("front-to-rear", "rear-to-front", "left-to-right", "right-to-left", "side-to-rear", "passive", "mixed", "rear-to-side", "bottom-to-top", "top-to-bottom"),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant.",
			},
			"platform_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the platform.",
			},
			"cluster_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtualization cluster the device hosts.",
			},
			"virtual_chassis_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual chassis.",
			},
			"virtual_chassis_position": schema.Int64Attribute{
				Optional:    true,
				Description: "Position in the virtual chassis (0-255).",
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"virtual_chassis_priority": schema.Int64Attribute{
				Optional:    true,
				Description: "Master election priority in the virtual chassis (0-255).",
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
			},
			"config_template_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the config template.",
			},
			"serial": schema.StringAttribute{
				Optional:    true,
				Description: "Chassis serial number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"asset_tag": schema.StringAttribute{
				Optional:    true,
				Description: "Unique asset tag.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"primary_ip4_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the primary IPv4 address (managed by netbox_primary_ip).",
			},
			"primary_ip6_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the primary IPv6 address (managed by netbox_primary_ip).",
			},
			"oob_ip_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the out-of-band management IP address (managed by netbox_device_oob_ip).",
			},
			"local_context_data": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Optional:    true,
				Description: "Local config context data as JSON text (use jsonencode()).",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Optional: true,
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
			"console_port_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of console ports on the device.",
			},
			"console_server_port_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of console server ports on the device.",
			},
			"power_port_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of power ports on the device.",
			},
			"power_outlet_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of power outlets on the device.",
			},
			"interface_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of interfaces on the device.",
			},
			"front_port_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of front ports on the device.",
			},
			"rear_port_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of rear ports on the device.",
			},
			"device_bay_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of device bays on the device.",
			},
			"module_bay_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of module bays on the device.",
			},
			"inventory_item_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of inventory items on the device.",
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

func (r *deviceResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *deviceResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var plan deviceResourceModel
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
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state deviceResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		if config.ConsolePortCount.IsNull() {
			plan.ConsolePortCount = state.ConsolePortCount
		}
		if config.ConsoleServerPortCount.IsNull() {
			plan.ConsoleServerPortCount = state.ConsoleServerPortCount
		}
		if config.PowerPortCount.IsNull() {
			plan.PowerPortCount = state.PowerPortCount
		}
		if config.PowerOutletCount.IsNull() {
			plan.PowerOutletCount = state.PowerOutletCount
		}
		if config.InterfaceCount.IsNull() {
			plan.InterfaceCount = state.InterfaceCount
		}
		if config.FrontPortCount.IsNull() {
			plan.FrontPortCount = state.FrontPortCount
		}
		if config.RearPortCount.IsNull() {
			plan.RearPortCount = state.RearPortCount
		}
		if config.DeviceBayCount.IsNull() {
			plan.DeviceBayCount = state.DeviceBayCount
		}
		if config.ModuleBayCount.IsNull() {
			plan.ModuleBayCount = state.ModuleBayCount
		}
		if config.InventoryItemCount.IsNull() {
			plan.InventoryItemCount = state.InventoryItemCount
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		if config.ConsolePortCount.IsNull() {
			plan.ConsolePortCount = types.Int64Unknown()
		}
		if config.ConsoleServerPortCount.IsNull() {
			plan.ConsoleServerPortCount = types.Int64Unknown()
		}
		if config.PowerPortCount.IsNull() {
			plan.PowerPortCount = types.Int64Unknown()
		}
		if config.PowerOutletCount.IsNull() {
			plan.PowerOutletCount = types.Int64Unknown()
		}
		if config.InterfaceCount.IsNull() {
			plan.InterfaceCount = types.Int64Unknown()
		}
		if config.FrontPortCount.IsNull() {
			plan.FrontPortCount = types.Int64Unknown()
		}
		if config.RearPortCount.IsNull() {
			plan.RearPortCount = types.Int64Unknown()
		}
		if config.DeviceBayCount.IsNull() {
			plan.DeviceBayCount = types.Int64Unknown()
		}
		if config.ModuleBayCount.IsNull() {
			plan.ModuleBayCount = types.Int64Unknown()
		}
		if config.InventoryItemCount.IsNull() {
			plan.InventoryItemCount = types.Int64Unknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *deviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data deviceResourceModel

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
	requestDTO, diags := expandDevice(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Dcim.DcimDevicesCreateContext(ctx, dcim.NewDcimDevicesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_device", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDevice(ctx, netboxapi.DeviceResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *deviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data deviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Dcim.DcimDevicesRetrieveContext(ctx, dcim.NewDcimDevicesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDevice(ctx, netboxapi.DeviceResponseDTOFromGoNetbox(res.Payload), state)...)
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

func (r *deviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data deviceResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandDevice(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior deviceResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Dcim.DcimDevicesUpdateContext(ctx, dcim.NewDcimDevicesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_device", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDevice(ctx, netboxapi.DeviceResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *deviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data deviceResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimDevicesDestroyContext(ctx, dcim.NewDcimDevicesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_device", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
