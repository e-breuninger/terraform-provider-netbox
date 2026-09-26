// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                     = (*powerOutletTemplateResource)(nil)
	_ resource.ResourceWithConfigure        = (*powerOutletTemplateResource)(nil)
	_ resource.ResourceWithImportState      = (*powerOutletTemplateResource)(nil)
	_ resource.ResourceWithModifyPlan       = (*powerOutletTemplateResource)(nil)
	_ resource.ResourceWithConfigValidators = (*powerOutletTemplateResource)(nil)
)

// NewPowerOutletTemplateResource returns a new power_outlet_template resource.
func NewPowerOutletTemplateResource() resource.Resource {
	return &powerOutletTemplateResource{}
}

type powerOutletTemplateResource struct {
	client *netboxapi.Client
}

func (r *powerOutletTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_power_outlet_template"
}

func (r *powerOutletTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A power outlet template of a device type or module type (dcim.poweroutlettemplate).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/poweroutlettemplate/):\n\n> A template for a power outlet that will be created on all instantiations of the parent device type. See the power outlet documentation for more detail.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"device_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device type. Set exactly one of device_type_id and module_type_id; NetBox rejects both and neither.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"module_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the module type. Set exactly one of device_type_id and module_type_id; NetBox rejects both and neither.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Description: "Outlet type as its NetBox slug, e.g. iec-60320-c13, nema-5-15r, CS8364C (any of NetBox's power outlet type choices).",
			},
			"power_port_template_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the power port template this outlet draws from.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"power_port_id": schema.Int64Attribute{
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use power_port_template_id instead.",
				Description:        "Deprecated alias of power_port_template_id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"feed_leg": schema.StringAttribute{
				Optional:    true,
				Description: "Phase of the upstream feed this outlet is mapped to. One of: A, B, C.",
				Validators: []validator.String{
					stringvalidator.OneOf("A", "B", "C"),
				},
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
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
		},
	}
}

// ConfigValidators enforces the spec's cross-attribute rules (exactly_one_of / conflicts_with /
// required_with) on the configuration, and the ones the companion file's config_validators hook
// (spec hooks) contributes.
func (r *powerOutletTemplateResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	out := []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(path.MatchRoot("device_type_id"), path.MatchRoot("module_type_id")),
	}
	return out
}

func (r *powerOutletTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *powerOutletTemplateResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Alias sync, after the backend fragment and the companion hook (both rewrite resp.Plan
	// wholesale from req.Plan, so an earlier sync would be clobbered): fold each configured
	// deprecated alias into its canonical attribute and keep the two equal in the plan.
	if !resp.Diagnostics.HasError() && !resp.Plan.Raw.IsNull() {
		var config, plan powerOutletTemplateResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		switch {
		case !config.PowerPortID.IsNull() && !config.PowerPortTemplateID.IsNull():
			if !config.PowerPortID.IsUnknown() && !config.PowerPortTemplateID.IsUnknown() && !config.PowerPortID.Equal(config.PowerPortTemplateID) {
				resp.Diagnostics.AddError("Conflicting attribute values", "power_port_id is a deprecated alias of power_port_template_id; both are set and the values differ.")
			}
		case !config.PowerPortID.IsNull():
			plan.PowerPortTemplateID = plan.PowerPortID
		case !config.PowerPortTemplateID.IsNull():
			plan.PowerPortID = plan.PowerPortTemplateID
		default:
			// power_port_template_id was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.PowerPortTemplateID = types.Int64Null()
			plan.PowerPortID = plan.PowerPortTemplateID
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
		var config, plan, state powerOutletTemplateResourceModel
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

func (r *powerOutletTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data powerOutletTemplateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	requestDTO, diags := expandPowerOutletTemplate(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Dcim.DcimPowerOutletTemplatesCreateContext(ctx, dcim.NewDcimPowerOutletTemplatesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_power_outlet_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenPowerOutletTemplate(ctx, netboxapi.PowerOutletTemplateResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *powerOutletTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data powerOutletTemplateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	res, err := r.client.Dcim.DcimPowerOutletTemplatesRetrieveContext(ctx, dcim.NewDcimPowerOutletTemplatesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_power_outlet_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenPowerOutletTemplate(ctx, netboxapi.PowerOutletTemplateResponseDTOFromGoNetbox(res.Payload), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *powerOutletTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data powerOutletTemplateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	requestDTO, diags := expandPowerOutletTemplate(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	res, err := r.client.Dcim.DcimPowerOutletTemplatesUpdateContext(ctx, dcim.NewDcimPowerOutletTemplatesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_power_outlet_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenPowerOutletTemplate(ctx, netboxapi.PowerOutletTemplateResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *powerOutletTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data powerOutletTemplateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimPowerOutletTemplatesDestroyContext(ctx, dcim.NewDcimPowerOutletTemplatesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_power_outlet_template", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *powerOutletTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
