// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/float64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*deviceTypeResource)(nil)
	_ resource.ResourceWithConfigure   = (*deviceTypeResource)(nil)
	_ resource.ResourceWithImportState = (*deviceTypeResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*deviceTypeResource)(nil)
)

// NewDeviceTypeResource returns a new device_type resource.
func NewDeviceTypeResource() resource.Resource {
	return &deviceTypeResource{}
}

type deviceTypeResource struct {
	client *netboxapi.Client
}

func (r *deviceTypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_type"
}

func (r *deviceTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A device type (dcim.devicetype): a hardware model made by a manufacturer.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/devicetype/):\n\n> A device type represents a particular make and model of hardware that exists in the real world. Device types define the physical attributes of a device (rack height and depth) and its individual components (console, power, network interfaces, and so on).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"manufacturer_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the manufacturer.",
			},
			"model": schema.StringAttribute{
				Required:    true,
				Description: "Model name (unique per manufacturer).",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from model when not set.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"part_number": schema.StringAttribute{
				Optional:    true,
				Description: "Discrete part number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"default_platform_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the platform assigned to new devices of this type.",
			},
			"u_height": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Height in rack units (0.5 steps).",
				PlanModifiers: []planmodifier.Float64{
					float64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Float64{
					float64validator.Between(0, 1000),
				},
			},
			"is_full_depth": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Device consumes both the front and rear rack faces.",
				Default:     booldefault.StaticBool(true),
			},
			"exclude_from_utilization": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether devices of this type are left out of the rack utilization.",
				Default:     booldefault.StaticBool(false),
			},
			"subdevice_role": schema.StringAttribute{
				Optional:    true,
				Description: "Parent devices house child devices in device bays. One of: parent, child.",
				Validators: []validator.String{
					stringvalidator.OneOf("parent", "child"),
				},
			},
			"airflow": schema.StringAttribute{
				Optional:    true,
				Description: "Airflow direction. One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
				Validators: []validator.String{
					stringvalidator.OneOf("front-to-rear", "rear-to-front", "left-to-right", "right-to-left", "side-to-rear", "passive", "mixed", "rear-to-side", "bottom-to-top", "top-to-bottom"),
				},
			},
			"weight": schema.Float64Attribute{
				Optional:    true,
				Description: "Weight of the device type (requires weight_unit).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"weight_unit": schema.StringAttribute{
				Optional:    true,
				Description: "Unit of weight. One of: kg, g, lb, oz.",
				Validators: []validator.String{
					stringvalidator.OneOf("kg", "g", "lb", "oz"),
				},
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
			"device_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of devices of the device type.",
			},
			"console_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of console port templates of the device type.",
			},
			"console_server_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of console server port templates of the device type.",
			},
			"power_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of power port templates of the device type.",
			},
			"power_outlet_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of power outlet templates of the device type.",
			},
			"interface_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of interface templates of the device type.",
			},
			"front_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of front port templates of the device type.",
			},
			"rear_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of rear port templates of the device type.",
			},
			"device_bay_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of device bay templates of the device type.",
			},
			"module_bay_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of module bay templates of the device type.",
			},
			"inventory_item_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of inventory item templates of the device type.",
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

func (r *deviceTypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *deviceTypeResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var config, plan deviceTypeResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.Slug.IsNull() && !plan.Model.IsUnknown() && !plan.Model.IsNull() {
			plan.Slug = types.StringValue(netboxapi.Slugify(plan.Model.ValueString()))
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
		var config, plan, state deviceTypeResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		if config.DeviceCount.IsNull() {
			plan.DeviceCount = state.DeviceCount
		}
		if config.ConsolePortTemplateCount.IsNull() {
			plan.ConsolePortTemplateCount = state.ConsolePortTemplateCount
		}
		if config.ConsoleServerPortTemplateCount.IsNull() {
			plan.ConsoleServerPortTemplateCount = state.ConsoleServerPortTemplateCount
		}
		if config.PowerPortTemplateCount.IsNull() {
			plan.PowerPortTemplateCount = state.PowerPortTemplateCount
		}
		if config.PowerOutletTemplateCount.IsNull() {
			plan.PowerOutletTemplateCount = state.PowerOutletTemplateCount
		}
		if config.InterfaceTemplateCount.IsNull() {
			plan.InterfaceTemplateCount = state.InterfaceTemplateCount
		}
		if config.FrontPortTemplateCount.IsNull() {
			plan.FrontPortTemplateCount = state.FrontPortTemplateCount
		}
		if config.RearPortTemplateCount.IsNull() {
			plan.RearPortTemplateCount = state.RearPortTemplateCount
		}
		if config.DeviceBayTemplateCount.IsNull() {
			plan.DeviceBayTemplateCount = state.DeviceBayTemplateCount
		}
		if config.ModuleBayTemplateCount.IsNull() {
			plan.ModuleBayTemplateCount = state.ModuleBayTemplateCount
		}
		if config.InventoryItemTemplateCount.IsNull() {
			plan.InventoryItemTemplateCount = state.InventoryItemTemplateCount
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		if config.DeviceCount.IsNull() {
			plan.DeviceCount = types.Int64Unknown()
		}
		if config.ConsolePortTemplateCount.IsNull() {
			plan.ConsolePortTemplateCount = types.Int64Unknown()
		}
		if config.ConsoleServerPortTemplateCount.IsNull() {
			plan.ConsoleServerPortTemplateCount = types.Int64Unknown()
		}
		if config.PowerPortTemplateCount.IsNull() {
			plan.PowerPortTemplateCount = types.Int64Unknown()
		}
		if config.PowerOutletTemplateCount.IsNull() {
			plan.PowerOutletTemplateCount = types.Int64Unknown()
		}
		if config.InterfaceTemplateCount.IsNull() {
			plan.InterfaceTemplateCount = types.Int64Unknown()
		}
		if config.FrontPortTemplateCount.IsNull() {
			plan.FrontPortTemplateCount = types.Int64Unknown()
		}
		if config.RearPortTemplateCount.IsNull() {
			plan.RearPortTemplateCount = types.Int64Unknown()
		}
		if config.DeviceBayTemplateCount.IsNull() {
			plan.DeviceBayTemplateCount = types.Int64Unknown()
		}
		if config.ModuleBayTemplateCount.IsNull() {
			plan.ModuleBayTemplateCount = types.Int64Unknown()
		}
		if config.InventoryItemTemplateCount.IsNull() {
			plan.InventoryItemTemplateCount = types.Int64Unknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *deviceTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data deviceTypeResourceModel

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
	requestDTO, diags := expandDeviceType(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device_type custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Dcim.DcimDeviceTypesCreateContext(ctx, dcim.NewDcimDeviceTypesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_device_type", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceType(ctx, netboxapi.DeviceTypeResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *deviceTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data deviceTypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Dcim.DcimDeviceTypesRetrieveContext(ctx, dcim.NewDcimDeviceTypesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device_type", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceType(ctx, netboxapi.DeviceTypeResponseDTOFromGoNetbox(res.Payload), state)...)
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

func (r *deviceTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data deviceTypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandDeviceType(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior deviceTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device_type custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Dcim.DcimDeviceTypesUpdateContext(ctx, dcim.NewDcimDeviceTypesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_device_type", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceType(ctx, netboxapi.DeviceTypeResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *deviceTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data deviceTypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimDeviceTypesDestroyContext(ctx, dcim.NewDcimDeviceTypesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_device_type", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
