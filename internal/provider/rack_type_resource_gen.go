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
	_ resource.Resource                = (*rackTypeResource)(nil)
	_ resource.ResourceWithConfigure   = (*rackTypeResource)(nil)
	_ resource.ResourceWithImportState = (*rackTypeResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*rackTypeResource)(nil)
)

// NewRackTypeResource returns a new rack_type resource.
func NewRackTypeResource() resource.Resource {
	return &rackTypeResource{}
}

type rackTypeResource struct {
	client *netboxapi.Client
}

func (r *rackTypeResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rack_type"
}

func (r *rackTypeResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A NetBox rack type (dcim.racktype): a rack model made by a manufacturer, whose physical attributes racks inherit.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/en/stable/models/dcim/racktype/):\n\n> A rack type defines the physical characteristics of a particular model of rack.",
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
				Required: true,
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
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"form_factor": schema.StringAttribute{
				Required:    true,
				Description: "Rack form factor. One of: 2-post-frame, 4-post-frame, 4-post-cabinet, wall-frame, wall-frame-vertical, wall-cabinet, wall-cabinet-vertical.",
				Validators: []validator.String{
					stringvalidator.OneOf("2-post-frame", "4-post-frame", "4-post-cabinet", "wall-frame", "wall-frame-vertical", "wall-cabinet", "wall-cabinet-vertical"),
				},
			},
			"width": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Rail-to-rail width in inches. One of: 10, 19, 21, 23.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.OneOf(10, 19, 21, 23),
				},
			},
			"u_height": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Height in rack units.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"starting_unit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Lowest unit number.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(1, 32767),
				},
			},
			"desc_units": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units are numbered top-to-bottom.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"outer_width": schema.Int64Attribute{
				Optional:    true,
				Description: "Outer width (requires outer_unit).",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"outer_height": schema.Int64Attribute{
				Optional:    true,
				Description: "Outer height (requires outer_unit).",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"outer_depth": schema.Int64Attribute{
				Optional:    true,
				Description: "Outer depth (requires outer_unit).",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"outer_unit": schema.StringAttribute{
				Optional:    true,
				Description: "Unit of the outer dimensions. One of: mm, in.",
				Validators: []validator.String{
					stringvalidator.OneOf("mm", "in"),
				},
			},
			"mounting_depth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum depth of a mounted device, in millimetres.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"mounting_depth_mm": schema.Int64Attribute{
				Optional:           true,
				Computed:           true,
				DeprecationMessage: "Use mounting_depth instead.",
				Description:        "Deprecated alias of mounting_depth.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"weight": schema.Float64Attribute{
				Optional:    true,
				Description: "Weight of the rack (requires weight_unit).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"max_weight": schema.Int64Attribute{
				Optional:    true,
				Description: "Maximum load capacity (requires weight_unit).",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"weight_unit": schema.StringAttribute{
				Optional:    true,
				Description: "Unit of weight and max_weight. One of: kg, g, lb, oz.",
				Validators: []validator.String{
					stringvalidator.OneOf("kg", "g", "lb", "oz"),
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
			"rack_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of racks of the rack type.",
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

func (r *rackTypeResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *rackTypeResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var config, plan rackTypeResourceModel
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
	// Alias sync, after the backend fragment and the companion hook (both rewrite resp.Plan
	// wholesale from req.Plan, so an earlier sync would be clobbered): fold each configured
	// deprecated alias into its canonical attribute and keep the two equal in the plan.
	if !resp.Diagnostics.HasError() && !resp.Plan.Raw.IsNull() {
		var config, plan rackTypeResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		switch {
		case !config.MountingDepthMm.IsNull() && !config.MountingDepth.IsNull():
			if !config.MountingDepthMm.IsUnknown() && !config.MountingDepth.IsUnknown() && !config.MountingDepthMm.Equal(config.MountingDepth) {
				resp.Diagnostics.AddError("Conflicting attribute values", "mounting_depth_mm is a deprecated alias of mounting_depth; both are set and the values differ.")
			}
		case !config.MountingDepthMm.IsNull():
			plan.MountingDepth = plan.MountingDepthMm
		case !config.MountingDepth.IsNull():
			plan.MountingDepthMm = plan.MountingDepth
		default:
			// mounting_depth was plain optional before it gained the alias: removed from the
			// config (under either name), it still clears, despite the optional+computed
			// demotion keeping the prior value in the plan.
			plan.MountingDepth = types.Int64Null()
			plan.MountingDepthMm = plan.MountingDepth
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
		var config, plan, state rackTypeResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		if config.RackCount.IsNull() {
			plan.RackCount = state.RackCount
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		if config.RackCount.IsNull() {
			plan.RackCount = types.Int64Unknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *rackTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data rackTypeResourceModel

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
	requestDTO, diags := expandRackType(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_rack_type custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Dcim.DcimRackTypesCreateContext(ctx, dcim.NewDcimRackTypesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_rack_type", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenRackType(ctx, netboxapi.RackTypeResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *rackTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data rackTypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Dcim.DcimRackTypesRetrieveContext(ctx, dcim.NewDcimRackTypesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_rack_type", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenRackType(ctx, netboxapi.RackTypeResponseDTOFromGoNetbox(res.Payload), state)...)
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

func (r *rackTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data rackTypeResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandRackType(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior rackTypeResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_rack_type custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Dcim.DcimRackTypesUpdateContext(ctx, dcim.NewDcimRackTypesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_rack_type", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenRackType(ctx, netboxapi.RackTypeResponseDTOFromGoNetbox(res.Payload), plan)...)
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

func (r *rackTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data rackTypeResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimRackTypesDestroyContext(ctx, dcim.NewDcimRackTypesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_rack_type", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *rackTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
