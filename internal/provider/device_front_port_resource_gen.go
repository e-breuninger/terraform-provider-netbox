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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
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
	_ resource.Resource                     = (*deviceFrontPortResource)(nil)
	_ resource.ResourceWithConfigure        = (*deviceFrontPortResource)(nil)
	_ resource.ResourceWithImportState      = (*deviceFrontPortResource)(nil)
	_ resource.ResourceWithModifyPlan       = (*deviceFrontPortResource)(nil)
	_ resource.ResourceWithConfigValidators = (*deviceFrontPortResource)(nil)
)

// NewDeviceFrontPortResource returns a new device_front_port resource.
func NewDeviceFrontPortResource() resource.Resource {
	return &deviceFrontPortResource{}
}

type deviceFrontPortResource struct {
	client *netboxapi.Client
}

func (r *deviceFrontPortResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_front_port"
}

func (r *deviceFrontPortResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A device front port (dcim.frontport): a pass-through port on the front of a patch panel or cassette, mapped position by position onto rear ports of the same device.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/frontport/):\n\n> Front ports are pass-through ports which represent physical cable connections that comprise part of a longer path. For example, the ports on the front face of a UTP patch panel would be modeled in NetBox as front ports. Each port is assigned a physical type, and must be mapped to a specific rear port on the same device. A single rear port may be mapped to multiple front ports, using numeric positions to annotate the specific alignment of each.",
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
			"module_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the installed module this component belongs to.",
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Port type as its NetBox slug, e.g. 8p8c, lc, mpo (any of NetBox's port type choices).",
			},
			"positions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of positions on the front port, at least the number of rear_ports mappings. NetBox checks a lower value against the mappings it already holds, so remove mappings in one apply and lower positions in the next.",
				Default:     int64default.StaticInt64(1),
				Validators: []validator.Int64{
					int64validator.Between(1, 1024),
				},
			},
			"rear_ports": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"position": schema.Int64Attribute{
							Required:    true,
							Description: "Position on the front port, 1-based and unique per port.",
						},
						"rear_port_id": schema.Int64Attribute{
							Required:    true,
							Description: "Id of the rear port the position maps to.",
						},
						"rear_port_position": schema.Int64Attribute{
							Required:    true,
							Description: "Position on that rear port, 1-based; each rear port position takes one mapping.",
							Validators: []validator.Int64{
								int64validator.Between(1, 1024),
							},
						},
					},
				},
				Optional:    true,
				Computed:    true,
				Description: "Mappings of the front port's positions onto rear ports of the same device, one object per position. Derived from rear_port_id and rear_port_position when those are set; an empty set, or neither form, leaves the port unmapped. Changing the set recreates every mapping, since NetBox 4.6 rejects an update that resends an existing one.",
				PlanModifiers: []planmodifier.Set{
					setplanmodifier.UseStateForUnknown(),
				},
			},
			"rear_port_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Rear port of a single-position front port: the same as rear_ports = [{position = 1, rear_port_id = ..., rear_port_position = ...}]. Conflicts with rear_ports and with positions > 1.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"rear_port_position": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Position on rear_port_id. Requires rear_port_id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"color_hex": schema.StringAttribute{
				Optional:    true,
				Description: "RGB color in hex (e.g. ff0000).",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(6),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[0-9a-f]{6}$`), ""),
				},
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"mark_connected": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat as if a cable is connected.",
				Default:     booldefault.StaticBool(false),
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

// ConfigValidators enforces the spec's cross-attribute rules (exactly_one_of / conflicts_with /
// required_with) on the configuration, and the ones the companion file's config_validators hook
// (spec hooks) contributes.
func (r *deviceFrontPortResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	out := []resource.ConfigValidator{
		resourcevalidator.Conflicting(path.MatchRoot("rear_port_id"), path.MatchRoot("rear_ports")),
	}
	return out
}

func (r *deviceFrontPortResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *deviceFrontPortResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var plan deviceFrontPortResourceModel
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
	// The companion file's modify_plan hook (spec hooks), after the backend fragment.
	r.modifyPlan(ctx, req, resp)
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state deviceFrontPortResourceModel
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

func (r *deviceFrontPortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data deviceFrontPortResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// The companion file's pre_create hook (spec hooks).
	r.preCreate(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// convert to DTO
	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandDeviceFrontPort(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device_front_port custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Dcim.DcimFrontPortsCreateContext(ctx, dcim.NewDcimFrontPortsCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_device_front_port", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceFrontPort(ctx, netboxapi.DeviceFrontPortResponseDTOFromGoNetbox(res.Payload), plan)...)
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
	// The companion file's post_create hook (spec hooks).
	r.postCreate(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *deviceFrontPortResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data deviceFrontPortResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Dcim.DcimFrontPortsRetrieveContext(ctx, dcim.NewDcimFrontPortsRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device_front_port", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceFrontPort(ctx, netboxapi.DeviceFrontPortResponseDTOFromGoNetbox(res.Payload), state)...)
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
	// The companion file's post_read hook (spec hooks).
	r.postRead(ctx, state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *deviceFrontPortResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data deviceFrontPortResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data
	// The companion file's pre_update hook (spec hooks).
	r.preUpdate(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandDeviceFrontPort(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior deviceFrontPortResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_device_front_port custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Dcim.DcimFrontPortsUpdateContext(ctx, dcim.NewDcimFrontPortsUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_device_front_port", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenDeviceFrontPort(ctx, netboxapi.DeviceFrontPortResponseDTOFromGoNetbox(res.Payload), plan)...)
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
	// The companion file's post_update hook (spec hooks).
	r.postUpdate(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *deviceFrontPortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data deviceFrontPortResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimFrontPortsDestroyContext(ctx, dcim.NewDcimFrontPortsDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_device_front_port", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *deviceFrontPortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
