// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*customFieldChoiceSetResource)(nil)
	_ resource.ResourceWithConfigure   = (*customFieldChoiceSetResource)(nil)
	_ resource.ResourceWithImportState = (*customFieldChoiceSetResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*customFieldChoiceSetResource)(nil)
)

// NewCustomFieldChoiceSetResource returns a new custom_field_choice_set resource.
func NewCustomFieldChoiceSetResource() resource.Resource {
	return &customFieldChoiceSetResource{}
}

type customFieldChoiceSetResource struct {
	client *netboxapi.Client
}

func (r *customFieldChoiceSetResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_field_choice_set"
}

func (r *customFieldChoiceSetResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox custom field choice set (extras.customfieldchoiceset).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/extras/customfieldchoiceset/):\n\nSingle- and multi-selection custom fields must define a set of valid choices from which the user may choose when defining the field value. These choices are defined as sets that may be reused among multiple custom fields.\n\nA choice set must define a base choice set and/or a set of arbitrary extra choices.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id of the choice set.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"base_choices": schema.StringAttribute{
				Optional:    true,
				Description: "Predefined set of choices to start from. One of: IATA, ISO_3166, UN_LOCODE.",
				Validators: []validator.String{
					stringvalidator.OneOf("IATA", "ISO_3166", "UN_LOCODE"),
				},
			},
			"extra_choices": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"value": schema.StringAttribute{
							Required: true,
						},
						"label": schema.StringAttribute{
							Required: true,
						},
					},
				},
				Optional:    true,
				Description: "Additional choices as value/label pairs.",
			},
			"order_alphabetically": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sort choices alphabetically instead of by definition order.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"choices_count": schema.Int64Attribute{
				Computed: true,
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
		},
	}
}

func (r *customFieldChoiceSetResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *customFieldChoiceSetResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state customFieldChoiceSetResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.ChoicesCount.IsNull() {
			plan.ChoicesCount = state.ChoicesCount
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.ChoicesCount.IsNull() {
			plan.ChoicesCount = types.Int64Unknown()
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *customFieldChoiceSetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data customFieldChoiceSetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	priorExtraChoices := plan.ExtraChoices
	requestDTO, diags := expandCustomFieldChoiceSet(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Extras.ExtrasCustomFieldChoiceSetsCreateContext(ctx, extras.NewExtrasCustomFieldChoiceSetsCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_custom_field_choice_set", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCustomFieldChoiceSet(ctx, netboxapi.CustomFieldChoiceSetResponseDTOFromGoNetbox(res.Payload), plan)...)
	plan.ExtraChoices = reorderToPrior(priorExtraChoices, plan.ExtraChoices)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *customFieldChoiceSetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data customFieldChoiceSetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	priorExtraChoices := state.ExtraChoices
	res, err := r.client.Extras.ExtrasCustomFieldChoiceSetsRetrieveContext(ctx, extras.NewExtrasCustomFieldChoiceSetsRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_custom_field_choice_set", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCustomFieldChoiceSet(ctx, netboxapi.CustomFieldChoiceSetResponseDTOFromGoNetbox(res.Payload), state)...)
	state.ExtraChoices = reorderToPrior(priorExtraChoices, state.ExtraChoices)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customFieldChoiceSetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data customFieldChoiceSetResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	priorExtraChoices := plan.ExtraChoices
	requestDTO, diags := expandCustomFieldChoiceSet(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	res, err := r.client.Extras.ExtrasCustomFieldChoiceSetsUpdateContext(ctx, extras.NewExtrasCustomFieldChoiceSetsUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_custom_field_choice_set", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCustomFieldChoiceSet(ctx, netboxapi.CustomFieldChoiceSetResponseDTOFromGoNetbox(res.Payload), plan)...)
	plan.ExtraChoices = reorderToPrior(priorExtraChoices, plan.ExtraChoices)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *customFieldChoiceSetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data customFieldChoiceSetResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Extras.ExtrasCustomFieldChoiceSetsDestroyContext(ctx, extras.NewExtrasCustomFieldChoiceSetsDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_custom_field_choice_set", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *customFieldChoiceSetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
