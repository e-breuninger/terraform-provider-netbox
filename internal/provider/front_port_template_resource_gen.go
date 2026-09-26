// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/setplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                     = (*frontPortTemplateResource)(nil)
	_ resource.ResourceWithConfigure        = (*frontPortTemplateResource)(nil)
	_ resource.ResourceWithImportState      = (*frontPortTemplateResource)(nil)
	_ resource.ResourceWithModifyPlan       = (*frontPortTemplateResource)(nil)
	_ resource.ResourceWithConfigValidators = (*frontPortTemplateResource)(nil)
)

// NewFrontPortTemplateResource returns a new front_port_template resource.
func NewFrontPortTemplateResource() resource.Resource {
	return &frontPortTemplateResource{}
}

type frontPortTemplateResource struct {
	client *netboxapi.Client
}

func (r *frontPortTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_front_port_template"
}

func (r *frontPortTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A front port template of a device type or module type (dcim.frontporttemplate), mapped position by position onto rear port templates of the same type.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/frontporttemplate/):\n\n> A template for a front-facing pass-through port that will be created on all instantiations of the parent device type. See the [front port](https://netboxlabs.com/docs/netbox/models/dcim/frontport/) documentation for more detail.",
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
							Description: "Id of the rear port template the position maps to.",
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
				Description: "Mappings of the front port's positions onto rear port templates of the same device type or module type, one object per position. Derived from rear_port_id and rear_port_position when those are set; an empty set, or neither form, leaves the port unmapped. Changing the set recreates every mapping, since NetBox 4.6 rejects an update that resends an existing one.",
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
func (r *frontPortTemplateResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	out := []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(path.MatchRoot("device_type_id"), path.MatchRoot("module_type_id")),
		resourcevalidator.Conflicting(path.MatchRoot("rear_port_id"), path.MatchRoot("rear_ports")),
	}
	return out
}

func (r *frontPortTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *frontPortTemplateResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The companion file's modify_plan hook (spec hooks), after the backend fragment.
	r.modifyPlan(ctx, req, resp)
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state frontPortTemplateResourceModel
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

func (r *frontPortTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data frontPortTemplateResourceModel

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
	requestDTO, diags := expandFrontPortTemplate(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Dcim.DcimFrontPortTemplatesCreateContext(ctx, dcim.NewDcimFrontPortTemplatesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_front_port_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenFrontPortTemplate(ctx, netboxapi.FrontPortTemplateResponseDTOFromGoNetbox(res.Payload), plan)...)

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

func (r *frontPortTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data frontPortTemplateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	res, err := r.client.Dcim.DcimFrontPortTemplatesRetrieveContext(ctx, dcim.NewDcimFrontPortTemplatesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_front_port_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenFrontPortTemplate(ctx, netboxapi.FrontPortTemplateResponseDTOFromGoNetbox(res.Payload), state)...)

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

func (r *frontPortTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data frontPortTemplateResourceModel

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

	requestDTO, diags := expandFrontPortTemplate(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	res, err := r.client.Dcim.DcimFrontPortTemplatesUpdateContext(ctx, dcim.NewDcimFrontPortTemplatesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_front_port_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenFrontPortTemplate(ctx, netboxapi.FrontPortTemplateResponseDTOFromGoNetbox(res.Payload), plan)...)

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

func (r *frontPortTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data frontPortTemplateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimFrontPortTemplatesDestroyContext(ctx, dcim.NewDcimFrontPortTemplatesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_front_port_template", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *frontPortTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
