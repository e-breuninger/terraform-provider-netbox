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
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/listplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*cableResource)(nil)
	_ resource.ResourceWithConfigure   = (*cableResource)(nil)
	_ resource.ResourceWithImportState = (*cableResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*cableResource)(nil)
)

// NewCableResource returns a new cable resource.
func NewCableResource() resource.Resource {
	return &cableResource{}
}

type cableResource struct {
	client *netboxapi.Client
}

func (r *cableResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cable"
}

func (r *cableResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A cable between two sets of terminations (dcim.cable): interfaces, front and rear ports, console and power ports, power feeds or circuit terminations.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/cable/):\n\n> All connections between device components in NetBox are represented using cables. A cable represents a direct physical connection between two sets of endpoints (A and B), such as a console port and a patch panel port, or between two network interfaces.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"a_side": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"object_type": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Content type of the terminating objects. Derived from the *_ids list that is set; set it together with ids otherwise. One of: dcim.interface, dcim.frontport, dcim.rearport, dcim.consoleport, dcim.consoleserverport, dcim.powerport, dcim.poweroutlet, dcim.powerfeed, circuits.circuittermination.",
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf("dcim.interface", "dcim.frontport", "dcim.rearport", "dcim.consoleport", "dcim.consoleserverport", "dcim.powerport", "dcim.poweroutlet", "dcim.powerfeed", "circuits.circuittermination"),
						},
					},
					"ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the terminating objects, in order (see object_type).",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"device_interface_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the device interfaces this side terminates on (object_type dcim.interface). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"front_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the front ports this side terminates on (object_type dcim.frontport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"rear_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the rear ports this side terminates on (object_type dcim.rearport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"console_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the console ports this side terminates on (object_type dcim.consoleport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"console_server_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the console server ports this side terminates on (object_type dcim.consoleserverport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"power_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the power ports this side terminates on (object_type dcim.powerport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"power_outlet_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the power outlets this side terminates on (object_type dcim.poweroutlet). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"power_feed_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the power feeds this side terminates on (object_type dcim.powerfeed). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"circuit_termination_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the circuit terminations this side terminates on (object_type circuits.circuittermination). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
				},
				Required:    true,
				Description: "What the A side terminates on: objects of one type (NetBox's rule), one or several for a breakout, in order (with a profile the position is the connector). Give object_type and ids, or exactly one of the typed *_ids lists, which is an alias of the same thing.",
			},
			"b_side": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"object_type": schema.StringAttribute{
						Optional:    true,
						Computed:    true,
						Description: "Content type of the terminating objects. Derived from the *_ids list that is set; set it together with ids otherwise. One of: dcim.interface, dcim.frontport, dcim.rearport, dcim.consoleport, dcim.consoleserverport, dcim.powerport, dcim.poweroutlet, dcim.powerfeed, circuits.circuittermination.",
						PlanModifiers: []planmodifier.String{
							stringplanmodifier.UseStateForUnknown(),
						},
						Validators: []validator.String{
							stringvalidator.OneOf("dcim.interface", "dcim.frontport", "dcim.rearport", "dcim.consoleport", "dcim.consoleserverport", "dcim.powerport", "dcim.poweroutlet", "dcim.powerfeed", "circuits.circuittermination"),
						},
					},
					"ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the terminating objects, in order (see object_type).",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"device_interface_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the device interfaces this side terminates on (object_type dcim.interface). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"front_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the front ports this side terminates on (object_type dcim.frontport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"rear_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the rear ports this side terminates on (object_type dcim.rearport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"console_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the console ports this side terminates on (object_type dcim.consoleport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"console_server_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the console server ports this side terminates on (object_type dcim.consoleserverport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"power_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the power ports this side terminates on (object_type dcim.powerport). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"power_outlet_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the power outlets this side terminates on (object_type dcim.poweroutlet). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"power_feed_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the power feeds this side terminates on (object_type dcim.powerfeed). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
					"circuit_termination_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Optional:    true,
						Computed:    true,
						Description: "Ids of the circuit terminations this side terminates on (object_type circuits.circuittermination). Conflicts with the other *_ids lists and with object_type/ids.",
						PlanModifiers: []planmodifier.List{
							listplanmodifier.UseStateForUnknown(),
						},
					},
				},
				Required:    true,
				Description: "What the B side terminates on: objects of one type (NetBox's rule), one or several for a breakout, in order (with a profile the position is the connector). Give object_type and ids, or exactly one of the typed *_ids lists, which is an alias of the same thing.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Description: "Cable type as its NetBox slug, e.g. cat6, smf-os2, power (any of NetBox's cable type choices).",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cable status. One of: connected, planned, decommissioning.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("connected", "planned", "decommissioning"),
				},
			},
			"profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cable profile (connectors and positions per side); with one set, each termination's position in the list is its connector. NetBox rejects clearing it, so unset keeps the current value. One of: single-1c1p, single-1c2p, single-1c4p, single-1c6p, single-1c8p, single-1c12p, single-1c16p, trunk-2c1p, trunk-2c2p, trunk-2c4p, trunk-2c4p-shuffle, trunk-2c6p, trunk-2c8p, trunk-2c12p, trunk-4c1p, trunk-4c2p, trunk-4c4p, trunk-4c4p-shuffle, trunk-4c6p, trunk-4c8p, trunk-8c4p, breakout-1c2p-2c1p, breakout-1c4p-4c1p, breakout-1c6p-6c1p, breakout-1c8p-8c1p, breakout-2c4p-8c1p-shuffle.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("single-1c1p", "single-1c2p", "single-1c4p", "single-1c6p", "single-1c8p", "single-1c12p", "single-1c16p", "trunk-2c1p", "trunk-2c2p", "trunk-2c4p", "trunk-2c4p-shuffle", "trunk-2c6p", "trunk-2c8p", "trunk-2c12p", "trunk-4c1p", "trunk-4c2p", "trunk-4c4p", "trunk-4c4p-shuffle", "trunk-4c6p", "trunk-4c8p", "trunk-8c4p", "breakout-1c2p-2c1p", "breakout-1c4p-4c1p", "breakout-1c6p-6c1p", "breakout-1c8p-8c1p", "breakout-2c4p-8c1p-shuffle"),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant.",
			},
			"bundle_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cable bundle.",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(100),
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
			"length": schema.Float64Attribute{
				Optional:    true,
				Description: "Cable length; requires length_unit.",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"length_unit": schema.StringAttribute{
				Optional:    true,
				Description: "Unit of length. One of: km, m, cm, mi, ft, in.",
				Validators: []validator.String{
					stringvalidator.OneOf("km", "m", "cm", "mi", "ft", "in"),
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

func (r *cableResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *cableResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var plan cableResourceModel
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
	r.modifyPlanHook(ctx, req, resp)
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state cableResourceModel
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

func (r *cableResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data cableResourceModel

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
	priorASide := plan.ASide
	priorBSide := plan.BSide
	requestDTO, diags := expandCable(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_cable custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Dcim.DcimCablesCreateContext(ctx, dcim.NewDcimCablesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_cable", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCable(ctx, netboxapi.CableResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)
	plan.ASide = keepPriorSide(priorASide, plan.ASide)
	plan.BSide = keepPriorSide(priorBSide, plan.BSide)

	if resp.Diagnostics.HasError() {
		return
	}
	// The companion file's post_create hook (spec hooks).
	r.postCreateHook(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *cableResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data cableResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	priorASide := state.ASide
	priorBSide := state.BSide
	res, err := r.client.Dcim.DcimCablesRetrieveContext(ctx, dcim.NewDcimCablesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_cable", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCable(ctx, netboxapi.CableResponseDTOFromGoNetbox(res.Payload), state)...)
	allTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	state.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	state.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	state.CustomFields = keepPriorCustomFields(priorCustomFields, state.CustomFields)
	state.ASide = keepPriorSide(priorASide, state.ASide)
	state.BSide = keepPriorSide(priorBSide, state.BSide)

	if resp.Diagnostics.HasError() {
		return
	}
	// The companion file's post_read hook (spec hooks).
	r.postReadHook(ctx, state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *cableResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data cableResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	priorASide := plan.ASide
	priorBSide := plan.BSide
	requestDTO, diags := expandCable(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior cableResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_cable custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Dcim.DcimCablesUpdateContext(ctx, dcim.NewDcimCablesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_cable", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCable(ctx, netboxapi.CableResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)
	plan.ASide = keepPriorSide(priorASide, plan.ASide)
	plan.BSide = keepPriorSide(priorBSide, plan.BSide)

	if resp.Diagnostics.HasError() {
		return
	}
	// The companion file's post_update hook (spec hooks).
	r.postUpdateHook(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *cableResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data cableResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimCablesDestroyContext(ctx, dcim.NewDcimCablesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_cable", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *cableResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
