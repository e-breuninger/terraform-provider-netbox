// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*circuitTerminationResource)(nil)
	_ resource.ResourceWithConfigure   = (*circuitTerminationResource)(nil)
	_ resource.ResourceWithImportState = (*circuitTerminationResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*circuitTerminationResource)(nil)
)

// NewCircuitTerminationResource returns a new circuit_termination resource.
func NewCircuitTerminationResource() resource.Resource {
	return &circuitTerminationResource{}
}

type circuitTerminationResource struct {
	client *netboxapi.Client
}

func (r *circuitTerminationResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_circuit_termination"
}

func (r *circuitTerminationResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:A termination attaching one side of a circuit to a site or provider network (circuits.circuittermination).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuit-terminations):\n\n> The association of a circuit with a particular site and/or device is modeled separately as a circuit termination. A circuit may have up to two terminations, labeled A and Z. A single-termination circuit can be used when you don't know (or care) about the far end of a circuit (for example, an Internet access circuit which connects to a transit provider). A dual-termination circuit is useful for tracking circuits which connect two sites.\n>\n> Each circuit termination is attached to either a site or to a provider network. Site terminations may optionally be connected via a cable to a specific device interface or port within that site. Each termination must be assigned a port speed, and can optionally be assigned an upstream speed if it differs from the downstream speed (a common scenario with e.g. DOCSIS cable modems). Fields are also available to track cross-connect and patch panel details.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"circuit_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the circuit.",
			},
			"term_side": schema.StringAttribute{
				Required:    true,
				Description: "Side of the circuit this terminates. One of: A, Z.",
				Validators: []validator.String{
					stringvalidator.OneOf("A", "Z"),
				},
			},
			"termination_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the termination. Derived from site_id, location_id, region_id, site_group_id or provider_network_id when one of those is set; set it together with termination_id otherwise. One termination is required. One of: dcim.site, dcim.location, dcim.region, dcim.sitegroup, circuits.providernetwork.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("dcim.site", "dcim.location", "dcim.region", "dcim.sitegroup", "circuits.providernetwork"),
				},
			},
			"termination_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating object (see termination_type).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating site (termination_type dcim.site). Conflicts with the other aliases and with setting the termination_* pair directly.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating location (termination_type dcim.location). Conflicts with the other aliases and with setting the termination_* pair directly.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"region_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating region (termination_type dcim.region). Conflicts with the other aliases and with setting the termination_* pair directly.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"site_group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating site group (termination_type dcim.sitegroup). Conflicts with the other aliases and with setting the termination_* pair directly.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"provider_network_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating provider network (termination_type circuits.providernetwork). Conflicts with the other aliases and with setting the termination_* pair directly.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"port_speed_kbps": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical circuit speed in kbps.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"port_speed": schema.Int64Attribute{
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use port_speed_kbps instead.",
				Description:        "Deprecated alias of port_speed_kbps.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"upstream_speed_kbps": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Upstream speed in kbps if different from the port speed.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"upstream_speed": schema.Int64Attribute{
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use upstream_speed_kbps instead.",
				Description:        "Deprecated alias of upstream_speed_kbps.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"xconnect_id": schema.StringAttribute{
				Optional:    true,
				Description: "Cross-connect id assigned by the provider.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"pp_info": schema.StringAttribute{
				Optional:    true,
				Description: "Patch panel and port information.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(100),
				},
			},
			"mark_connected": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat the termination as connected even without a cable.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
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

func (r *circuitTerminationResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *circuitTerminationResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var plan circuitTerminationResourceModel
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
	// Alias sync, after the backend fragment and the companion hook (both rewrite resp.Plan
	// wholesale from req.Plan, so an earlier sync would be clobbered): fold each configured
	// deprecated alias into its canonical attribute and keep the two equal in the plan.
	if !resp.Diagnostics.HasError() && !resp.Plan.Raw.IsNull() {
		var config, plan circuitTerminationResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		switch {
		case !config.PortSpeed.IsNull() && !config.PortSpeedKbps.IsNull():
			if !config.PortSpeed.IsUnknown() && !config.PortSpeedKbps.IsUnknown() && !config.PortSpeed.Equal(config.PortSpeedKbps) {
				resp.Diagnostics.AddError("Conflicting attribute values", "port_speed is a deprecated alias of port_speed_kbps; both are set and the values differ.")
			}
		case !config.PortSpeed.IsNull():
			plan.PortSpeedKbps = plan.PortSpeed
		case !config.PortSpeedKbps.IsNull():
			plan.PortSpeed = plan.PortSpeedKbps
		default:
			// port_speed_kbps was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.PortSpeedKbps = types.Int64Null()
			plan.PortSpeed = plan.PortSpeedKbps
		}
		switch {
		case !config.UpstreamSpeed.IsNull() && !config.UpstreamSpeedKbps.IsNull():
			if !config.UpstreamSpeed.IsUnknown() && !config.UpstreamSpeedKbps.IsUnknown() && !config.UpstreamSpeed.Equal(config.UpstreamSpeedKbps) {
				resp.Diagnostics.AddError("Conflicting attribute values", "upstream_speed is a deprecated alias of upstream_speed_kbps; both are set and the values differ.")
			}
		case !config.UpstreamSpeed.IsNull():
			plan.UpstreamSpeedKbps = plan.UpstreamSpeed
		case !config.UpstreamSpeedKbps.IsNull():
			plan.UpstreamSpeed = plan.UpstreamSpeedKbps
		default:
			// upstream_speed_kbps was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.UpstreamSpeedKbps = types.Int64Null()
			plan.UpstreamSpeed = plan.UpstreamSpeedKbps
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
		var config, plan, state circuitTerminationResourceModel
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

func (r *circuitTerminationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data circuitTerminationResourceModel

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
	requestDTO, diags := expandCircuitTermination(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_circuit_termination custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Circuits.CircuitsCircuitTerminationsCreateContext(ctx, circuits.NewCircuitsCircuitTerminationsCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_circuit_termination", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCircuitTermination(ctx, netboxapi.CircuitTerminationResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *circuitTerminationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data circuitTerminationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Circuits.CircuitsCircuitTerminationsRetrieveContext(ctx, circuits.NewCircuitsCircuitTerminationsRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_circuit_termination", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCircuitTermination(ctx, netboxapi.CircuitTerminationResponseDTOFromGoNetbox(res.Payload), state)...)
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

func (r *circuitTerminationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data circuitTerminationResourceModel

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
	requestDTO, diags := expandCircuitTermination(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior circuitTerminationResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_circuit_termination custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Circuits.CircuitsCircuitTerminationsUpdateContext(ctx, circuits.NewCircuitsCircuitTerminationsUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_circuit_termination", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCircuitTermination(ctx, netboxapi.CircuitTerminationResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *circuitTerminationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data circuitTerminationResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Circuits.CircuitsCircuitTerminationsDestroyContext(ctx, circuits.NewCircuitsCircuitTerminationsDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_circuit_termination", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *circuitTerminationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
