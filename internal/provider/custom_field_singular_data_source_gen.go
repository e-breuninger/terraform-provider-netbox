// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*customFieldDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customFieldDataSource)(nil)
)

// NewCustomFieldDataSource returns a new custom_field data source.
func NewCustomFieldDataSource() datasource.DataSource {
	return &customFieldDataSource{}
}

type customFieldDataSource struct {
	client *netboxapi.Client
}

// customFieldDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type customFieldDataSourceModel struct {
	customFieldResourceModel
	ChoiceSet           types.String `tfsdk:"choice_set"`
	ObjectType          types.String `tfsdk:"object_type"`
	ObjectTypeID        types.Int64  `tfsdk:"object_type_id"`
	OwnerName           types.String `tfsdk:"owner_name"`
	OwnerGroupName      types.String `tfsdk:"owner_group_name"`
	OwnerGroupID        types.Int64  `tfsdk:"owner_group_id"`
	Q                   types.String `tfsdk:"q"`
	RelatedObjectTypeID types.Int64  `tfsdk:"related_object_type_id"`
}

func (d *customFieldDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_field"
}

func (d *customFieldDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox custom field definition (extras.customfield). Values are set per object via the custom_fields attribute.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/customization/custom-fields/#custom-fields):\n\n> Each model in NetBox is represented in the database as a discrete table, and each attribute of a model exists as a column within its table. For example, sites are stored in the dcim_site table, which has columns named name, facility, physical_address, and so on. As new attributes are added to objects throughout the development of NetBox, tables are expanded to include new rows.\n>\n> However, some users might want to store additional object attributes that are somewhat esoteric in nature, and that would not make sense to include in the core NetBox database schema. For instance, suppose your organization needs to associate each device with a ticket number correlating it with an internal support system record. This is certainly a legitimate use for NetBox, but it's not a common enough need to warrant including a field for every NetBox installation. Instead, you can create a custom field to hold this data.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id of the custom field.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Internal field name (also the key in custom_fields).",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 50),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[a-z0-9_]+$`), ""),
				},
			},
			"object_types": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Object types the field applies to, e.g. dcim.site.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Field data type. One of: text, longtext, integer, decimal, boolean, date, datetime, url, json, select, multiselect, object, multiobject.",
				Validators: []validator.String{
					stringvalidator.OneOf("text", "longtext", "integer", "decimal", "boolean", "date", "datetime", "url", "json", "select", "multiselect", "object", "multiobject"),
				},
			},
			"related_object_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the related objects for object and multiobject fields, e.g. dcim.device.",
			},
			"related_object_filter": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Query parameters that narrow the objects offered for object and multiobject fields, as JSON text, e.g. jsonencode({ status = \"active\" }).",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Human-readable label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"group_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Custom fields in the same group are shown together.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"required": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether a value is required when creating objects.",
			},
			"unique": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the value must be unique among the objects the field is assigned to.",
			},
			"filter_logic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How the field is matched when filtering. One of: disabled, loose, exact.",
				Validators: []validator.String{
					stringvalidator.OneOf("disabled", "loose", "exact"),
				},
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Order within a group. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"search_weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of the field in the global search; lower values rank higher and 0 leaves the field out. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"ui_visible": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the field is displayed in the UI. One of: always, if-set, hidden.",
				Validators: []validator.String{
					stringvalidator.OneOf("always", "if-set", "hidden"),
				},
			},
			"ui_editable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the field value can be edited in the UI. One of: yes, no, hidden.",
				Validators: []validator.String{
					stringvalidator.OneOf("yes", "no", "hidden"),
				},
			},
			"is_cloneable": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the value is copied when an object is cloned.",
			},
			"validation_minimum": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimum allowed value (numeric fields).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000000000, 1000000000000),
				},
			},
			"validation_maximum": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum allowed value (numeric fields).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000000000, 1000000000000),
				},
			},
			"validation_regex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Regular expression text values must match.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(500),
				},
			},
			"validation_schema": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "JSON schema that values of a json field must match, as JSON text (use jsonencode()).",
			},
			"default": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Default value as JSON text, e.g. jsonencode(\"foo\") or jsonencode(5).",
			},
			"choice_set_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the choice set (select and multiselect fields).",
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"choice_set": schema.StringAttribute{
				Optional:    true,
				Description: "Choice set to match.",
			},
			"object_type": schema.StringAttribute{
				Optional:    true,
				Description: "Object type to match.",
			},
			"object_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the object type.",
			},
			"owner_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner.",
			},
			"owner_group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner group.",
			},
			"owner_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner group.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"related_object_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the related object type.",
			},
		},
	}
}

func (d *customFieldDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected data source configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = c
}

func (d *customFieldDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customFieldDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasCustomFieldsListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.ChoiceSet.IsNull() {
		v := state.ChoiceSet.ValueString()
		params.SetChoiceSet([]string{v})
		hasInput = true
	}
	if !state.ChoiceSetID.IsNull() {
		v := state.ChoiceSetID.ValueInt64()
		params.SetChoiceSetID([]int64{v})
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.FilterLogic.IsNull() {
		v := state.FilterLogic.ValueString()
		params.SetFilterLogic(&v)
		hasInput = true
	}
	if !state.GroupName.IsNull() {
		v := state.GroupName.ValueString()
		params.SetGroupName([]string{v})
		hasInput = true
	}
	if !state.IsCloneable.IsNull() {
		v := state.IsCloneable.ValueBool()
		params.SetIsCloneable(&v)
		hasInput = true
	}
	if !state.Label.IsNull() {
		v := state.Label.ValueString()
		params.SetLabel([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.ObjectType.IsNull() {
		v := state.ObjectType.ValueString()
		params.SetObjectType([]string{v})
		hasInput = true
	}
	if !state.ObjectTypeID.IsNull() {
		v := state.ObjectTypeID.ValueInt64()
		params.SetObjectTypeID([]int64{v})
		hasInput = true
	}
	if !state.OwnerName.IsNull() {
		v := state.OwnerName.ValueString()
		params.SetOwner([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupName.IsNull() {
		v := state.OwnerGroupName.ValueString()
		params.SetOwnerGroup([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupID.IsNull() {
		v := state.OwnerGroupID.ValueInt64()
		params.SetOwnerGroupID([]int64{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RelatedObjectType.IsNull() {
		v := state.RelatedObjectType.ValueString()
		params.SetRelatedObjectType([]string{v})
		hasInput = true
	}
	if !state.RelatedObjectTypeID.IsNull() {
		v := state.RelatedObjectTypeID.ValueInt64()
		params.SetRelatedObjectTypeID([]int64{v})
		hasInput = true
	}
	if !state.Required.IsNull() {
		v := state.Required.ValueBool()
		params.SetRequired(&v)
		hasInput = true
	}
	if !state.SearchWeight.IsNull() {
		v := state.SearchWeight.ValueInt64()
		params.SetSearchWeight([]int64{v})
		hasInput = true
	}
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.UiEditable.IsNull() {
		v := state.UiEditable.ValueString()
		params.SetUIEditable(&v)
		hasInput = true
	}
	if !state.UiVisible.IsNull() {
		v := state.UiVisible.ValueString()
		params.SetUIVisible(&v)
		hasInput = true
	}
	if !state.Unique.IsNull() {
		v := state.Unique.ValueBool()
		params.SetUnique(&v)
		hasInput = true
	}
	if !state.ValidationMaximum.IsNull() {
		v := state.ValidationMaximum.ValueFloat64()
		params.SetValidationMaximum([]float64{v})
		hasInput = true
	}
	if !state.ValidationMinimum.IsNull() {
		v := state.ValidationMinimum.ValueFloat64()
		params.SetValidationMinimum([]float64{v})
		hasInput = true
	}
	if !state.ValidationRegex.IsNull() {
		v := state.ValidationRegex.ValueString()
		params.SetValidationRegex([]string{v})
		hasInput = true
	}
	if !state.Weight.IsNull() {
		v := state.Weight.ValueInt64()
		params.SetWeight([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_custom_field.")
		return
	}
	res, err := d.client.Extras.ExtrasCustomFieldsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_custom_field", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_custom_field",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCustomField(ctx, netboxapi.CustomFieldResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.customFieldResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
