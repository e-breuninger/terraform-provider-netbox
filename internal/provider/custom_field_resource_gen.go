// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	_ resource.Resource                = (*customFieldResource)(nil)
	_ resource.ResourceWithConfigure   = (*customFieldResource)(nil)
	_ resource.ResourceWithImportState = (*customFieldResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*customFieldResource)(nil)
)

// NewCustomFieldResource returns a new custom_field resource.
func NewCustomFieldResource() resource.Resource {
	return &customFieldResource{}
}

type customFieldResource struct {
	client *netboxapi.Client
}

func (r *customFieldResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_field"
}

func (r *customFieldResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox custom field definition (extras.customfield). Values are set per object via the custom_fields attribute.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/customization/custom-fields/#custom-fields):\n\n> Each model in NetBox is represented in the database as a discrete table, and each attribute of a model exists as a column within its table. For example, sites are stored in the dcim_site table, which has columns named name, facility, physical_address, and so on. As new attributes are added to objects throughout the development of NetBox, tables are expanded to include new rows.\n>\n> However, some users might want to store additional object attributes that are somewhat esoteric in nature, and that would not make sense to include in the core NetBox database schema. For instance, suppose your organization needs to associate each device with a ticket number correlating it with an internal support system record. This is certainly a legitimate use for NetBox, but it's not a common enough need to warrant including a field for every NetBox installation. Instead, you can create a custom field to hold this data.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id of the custom field.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Internal field name (also the key in custom_fields).",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 50),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9_]+$`), ""),
				},
			},
			"object_types": schema.SetAttribute{
				ElementType: types.StringType,
				Required:    true,
				Description: "Object types the field applies to, e.g. dcim.site.",
			},
			"type": schema.StringAttribute{
				Required:    true,
				Description: "Field data type. One of: text, longtext, integer, decimal, boolean, date, datetime, url, json, select, multiselect, object, multiobject.",
				Validators: []validator.String{
					stringvalidator.OneOf("text", "longtext", "integer", "decimal", "boolean", "date", "datetime", "url", "json", "select", "multiselect", "object", "multiobject"),
				},
			},
			"related_object_type": schema.StringAttribute{
				Optional:    true,
				Description: "Content type of the related objects for object and multiobject fields, e.g. dcim.device.",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Description: "Human-readable label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Custom fields in the same group are shown together.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"required": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether a value is required when creating objects.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"filter_logic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How the field is matched when filtering. One of: disabled, loose, exact.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "loose", "exact"),
				},
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order within a group.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"validation_minimum": schema.Float64Attribute{
				Optional:    true,
				Description: "Minimum allowed value (numeric fields).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000000000, 1000000000000),
				},
			},
			"validation_maximum": schema.Float64Attribute{
				Optional:    true,
				Description: "Maximum allowed value (numeric fields).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000000000, 1000000000000),
				},
			},
			"validation_regex": schema.StringAttribute{
				Optional:    true,
				Description: "Regular expression text values must match.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(500),
				},
			},
			"default": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Optional:    true,
				Description: "Default value as JSON text, e.g. jsonencode(\"foo\") or jsonencode(5).",
			},
			"choice_set_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the choice set (select and multiselect fields).",
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

func (r *customFieldResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *customFieldResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state customFieldResourceModel
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

func (r *customFieldResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data customFieldResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	requestDTO, diags := expandCustomField(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Extras.ExtrasCustomFieldsCreateContext(ctx, extras.NewExtrasCustomFieldsCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_custom_field", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCustomField(ctx, netboxapi.CustomFieldResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *customFieldResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data customFieldResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	res, err := r.client.Extras.ExtrasCustomFieldsRetrieveContext(ctx, extras.NewExtrasCustomFieldsRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_custom_field", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCustomField(ctx, netboxapi.CustomFieldResponseDTOFromGoNetbox(res.Payload), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *customFieldResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data customFieldResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	requestDTO, diags := expandCustomField(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	res, err := r.client.Extras.ExtrasCustomFieldsUpdateContext(ctx, extras.NewExtrasCustomFieldsUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_custom_field", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenCustomField(ctx, netboxapi.CustomFieldResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *customFieldResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data customFieldResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Extras.ExtrasCustomFieldsDestroyContext(ctx, extras.NewExtrasCustomFieldsDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_custom_field", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *customFieldResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
