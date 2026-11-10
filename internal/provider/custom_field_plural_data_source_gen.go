// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*customFieldsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customFieldsDataSource)(nil)
)

// NewCustomFieldsDataSource returns a new custom_fields data source, which
// lists custom_field objects matching its filters.
func NewCustomFieldsDataSource() datasource.DataSource {
	return &customFieldsDataSource{}
}

type customFieldsDataSource struct {
	client *netboxapi.Client
}

// customFieldsDataSourceModel is the data source model: the filters, the limit and the matching
// custom_fields.
type customFieldsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"custom_fields"`
}

// customFieldResourceAttrTypes is the attribute type map of customFieldResourceModel.
func customFieldResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                    types.Int64Type,
		"name":                  types.StringType,
		"object_types":          types.SetType{ElemType: types.StringType},
		"type":                  types.StringType,
		"related_object_type":   types.StringType,
		"related_object_filter": jsontypes.NormalizedType{},
		"label":                 types.StringType,
		"description":           types.StringType,
		"comments":              types.StringType,
		"group_name":            types.StringType,
		"required":              types.BoolType,
		"unique":                types.BoolType,
		"filter_logic":          types.StringType,
		"weight":                types.Int64Type,
		"search_weight":         types.Int64Type,
		"ui_visible":            types.StringType,
		"ui_editable":           types.StringType,
		"is_cloneable":          types.BoolType,
		"validation_minimum":    types.Float64Type,
		"validation_maximum":    types.Float64Type,
		"validation_regex":      types.StringType,
		"validation_schema":     jsontypes.NormalizedType{},
		"default":               jsontypes.NormalizedType{},
		"choice_set_id":         types.Int64Type,
		"owner_id":              types.Int64Type,
		"created":               types.StringType,
		"last_updated":          types.StringType,
		"url":                   types.StringType,
	}
}

func (d *customFieldsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_fields"
}

func (d *customFieldsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists custom_field objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: choice_set, choice_set__n, choice_set_id, choice_set_id__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, filter_logic, filter_logic__empty, filter_logic__ic, filter_logic__ie, filter_logic__iew, filter_logic__iregex, filter_logic__isw, filter_logic__n, filter_logic__nic, filter_logic__nie, filter_logic__niew, filter_logic__nisw, filter_logic__regex, group_name, group_name__empty, group_name__ic, group_name__ie, group_name__iew, group_name__iregex, group_name__isw, group_name__n, group_name__nic, group_name__nie, group_name__niew, group_name__nisw, group_name__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, is_cloneable, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, object_type, object_type__ic, object_type__ie, object_type__iew, object_type__iregex, object_type__isw, object_type__n, object_type__nic, object_type__nie, object_type__niew, object_type__nisw, object_type__regex, object_type_id, object_type_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, related_object_type, related_object_type__n, related_object_type_id, related_object_type_id__n, required, search_weight, search_weight__empty, search_weight__gt, search_weight__gte, search_weight__lt, search_weight__lte, search_weight__n, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex, ui_editable, ui_editable__empty, ui_editable__ic, ui_editable__ie, ui_editable__iew, ui_editable__iregex, ui_editable__isw, ui_editable__n, ui_editable__nic, ui_editable__nie, ui_editable__niew, ui_editable__nisw, ui_editable__regex, ui_visible, ui_visible__empty, ui_visible__ic, ui_visible__ie, ui_visible__iew, ui_visible__iregex, ui_visible__isw, ui_visible__n, ui_visible__nic, ui_visible__nie, ui_visible__niew, ui_visible__nisw, ui_visible__regex, unique, validation_maximum, validation_maximum__empty, validation_maximum__gt, validation_maximum__gte, validation_maximum__lt, validation_maximum__lte, validation_maximum__n, validation_minimum, validation_minimum__empty, validation_minimum__gt, validation_minimum__gte, validation_minimum__lt, validation_minimum__lte, validation_minimum__n, validation_regex, validation_regex__empty, validation_regex__ic, validation_regex__ie, validation_regex__iew, validation_regex__iregex, validation_regex__isw, validation_regex__n, validation_regex__nic, validation_regex__nie, validation_regex__niew, validation_regex__nisw, validation_regex__regex, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n. Repeating a name sends that parameter once per value.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "Name of the query filter (the API list parameter).",
						},
						"value": schema.StringAttribute{
							Required:    true,
							Description: "Value to filter by.",
						},
					},
				},
			},
			"limit": schema.Int64Attribute{
				Optional:    true,
				Description: "The maximum number of objects to return from the API lookup. Defaults to 1000.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"custom_fields": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching custom_fields.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id of the custom field.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Internal field name (also the key in custom_fields).",
						},
						"object_types": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Object types the field applies to, e.g. dcim.site.",
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Field data type. One of: text, longtext, integer, decimal, boolean, date, datetime, url, json, select, multiselect, object, multiobject.",
						},
						"related_object_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the related objects for object and multiobject fields, e.g. dcim.device.",
						},
						"related_object_filter": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Query parameters that narrow the objects offered for object and multiobject fields, as JSON text, e.g. jsonencode({ status = \"active\" }).",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Human-readable label.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"group_name": schema.StringAttribute{
							Computed:    true,
							Description: "Custom fields in the same group are shown together.",
						},
						"required": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether a value is required when creating objects.",
						},
						"unique": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the value must be unique among the objects the field is assigned to.",
						},
						"filter_logic": schema.StringAttribute{
							Computed:    true,
							Description: "How the field is matched when filtering. One of: disabled, loose, exact.",
						},
						"weight": schema.Int64Attribute{
							Computed:    true,
							Description: "Order within a group. Must be a whole number.",
						},
						"search_weight": schema.Int64Attribute{
							Computed:    true,
							Description: "Weight of the field in the global search; lower values rank higher and 0 leaves the field out. Must be a whole number.",
						},
						"ui_visible": schema.StringAttribute{
							Computed:    true,
							Description: "Whether the field is displayed in the UI. One of: always, if-set, hidden.",
						},
						"ui_editable": schema.StringAttribute{
							Computed:    true,
							Description: "Whether the field value can be edited in the UI. One of: yes, no, hidden.",
						},
						"is_cloneable": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the value is copied when an object is cloned.",
						},
						"validation_minimum": schema.Float64Attribute{
							Computed:    true,
							Description: "Minimum allowed value (numeric fields).",
						},
						"validation_maximum": schema.Float64Attribute{
							Computed:    true,
							Description: "Maximum allowed value (numeric fields).",
						},
						"validation_regex": schema.StringAttribute{
							Computed:    true,
							Description: "Regular expression text values must match.",
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
							Computed:    true,
							Description: "Id of the choice set (select and multiselect fields).",
						},
						"owner_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the owner the object is assigned to.",
						},
						"created": schema.StringAttribute{
							Computed: true,
						},
						"last_updated": schema.StringAttribute{
							Computed: true,
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *customFieldsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *customFieldsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customFieldsDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	limit := int64(1000)
	if !state.Limit.IsNull() && !state.Limit.IsUnknown() {
		limit = state.Limit.ValueInt64()
	}
	_ = limit

	// The filters entries, in no particular order. Entries sharing a name are distinct values of
	// the same API parameter.
	var filters []struct {
		Name  types.String `tfsdk:"name"`
		Value types.String `tfsdk:"value"`
	}
	if !state.Filters.IsNull() {
		resp.Diagnostics.Append(state.Filters.ElementsAs(ctx, &filters, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	_ = filters
	var responseDTOs []*netboxapi.CustomFieldResponseDTO

	params := extras.NewExtrasCustomFieldsListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vName []string
	var vNamen []string
	var vNameIc []string
	var vNameNic []string
	var vNameIe []string
	var vNameNie []string
	var vNameIsw []string
	var vNameNisw []string
	var vNameIew []string
	var vNameNiew []string
	var vNameEmpty *bool
	var vNameRegex []string
	var vNameIregex []string
	var vChoiceSet []string
	var vChoiceSetn []string
	var vChoiceSetID []int64
	var vChoiceSetIDn []int64
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vDescription []string
	var vDescriptionn []string
	var vDescriptionIc []string
	var vDescriptionNic []string
	var vDescriptionIe []string
	var vDescriptionNie []string
	var vDescriptionIsw []string
	var vDescriptionNisw []string
	var vDescriptionIew []string
	var vDescriptionNiew []string
	var vDescriptionEmpty *bool
	var vDescriptionRegex []string
	var vDescriptionIregex []string
	var vFilterLogic *string
	var vFilterLogicn *string
	var vFilterLogicIc []string
	var vFilterLogicNic []string
	var vFilterLogicIe []string
	var vFilterLogicNie []string
	var vFilterLogicIsw []string
	var vFilterLogicNisw []string
	var vFilterLogicIew []string
	var vFilterLogicNiew []string
	var vFilterLogicEmpty *bool
	var vFilterLogicRegex []string
	var vFilterLogicIregex []string
	var vGroupName []string
	var vGroupNamen []string
	var vGroupNameIc []string
	var vGroupNameNic []string
	var vGroupNameIe []string
	var vGroupNameNie []string
	var vGroupNameIsw []string
	var vGroupNameNisw []string
	var vGroupNameIew []string
	var vGroupNameNiew []string
	var vGroupNameEmpty *bool
	var vGroupNameRegex []string
	var vGroupNameIregex []string
	var vIsCloneable *bool
	var vLabel []string
	var vLabeln []string
	var vLabelIc []string
	var vLabelNic []string
	var vLabelIe []string
	var vLabelNie []string
	var vLabelIsw []string
	var vLabelNisw []string
	var vLabelIew []string
	var vLabelNiew []string
	var vLabelEmpty *bool
	var vLabelRegex []string
	var vLabelIregex []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vObjectType []string
	var vObjectTypen []string
	var vObjectTypeIc []string
	var vObjectTypeNic []string
	var vObjectTypeIe []string
	var vObjectTypeNie []string
	var vObjectTypeIsw []string
	var vObjectTypeNisw []string
	var vObjectTypeIew []string
	var vObjectTypeNiew []string
	var vObjectTypeRegex []string
	var vObjectTypeIregex []string
	var vObjectTypeID []int64
	var vObjectTypeIDn []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vQ *string
	var vRelatedObjectType []string
	var vRelatedObjectTypen []string
	var vRelatedObjectTypeID []int64
	var vRelatedObjectTypeIDn []int64
	var vRequired *bool
	var vSearchWeight []int64
	var vSearchWeightn []int64
	var vSearchWeightLt []int64
	var vSearchWeightLte []int64
	var vSearchWeightGt []int64
	var vSearchWeightGte []int64
	var vSearchWeightEmpty *bool
	var vType []string
	var vTypen []string
	var vTypeIc []string
	var vTypeNic []string
	var vTypeIe []string
	var vTypeNie []string
	var vTypeIsw []string
	var vTypeNisw []string
	var vTypeIew []string
	var vTypeNiew []string
	var vTypeEmpty *bool
	var vTypeRegex []string
	var vTypeIregex []string
	var vUIEditable *string
	var vUIEditablen *string
	var vUIEditableIc []string
	var vUIEditableNic []string
	var vUIEditableIe []string
	var vUIEditableNie []string
	var vUIEditableIsw []string
	var vUIEditableNisw []string
	var vUIEditableIew []string
	var vUIEditableNiew []string
	var vUIEditableEmpty *bool
	var vUIEditableRegex []string
	var vUIEditableIregex []string
	var vUIVisible *string
	var vUIVisiblen *string
	var vUIVisibleIc []string
	var vUIVisibleNic []string
	var vUIVisibleIe []string
	var vUIVisibleNie []string
	var vUIVisibleIsw []string
	var vUIVisibleNisw []string
	var vUIVisibleIew []string
	var vUIVisibleNiew []string
	var vUIVisibleEmpty *bool
	var vUIVisibleRegex []string
	var vUIVisibleIregex []string
	var vUnique *bool
	var vValidationMaximum []float64
	var vValidationMaximumn []float64
	var vValidationMaximumLt []float64
	var vValidationMaximumLte []float64
	var vValidationMaximumGt []float64
	var vValidationMaximumGte []float64
	var vValidationMaximumEmpty *bool
	var vValidationMinimum []float64
	var vValidationMinimumn []float64
	var vValidationMinimumLt []float64
	var vValidationMinimumLte []float64
	var vValidationMinimumGt []float64
	var vValidationMinimumGte []float64
	var vValidationMinimumEmpty *bool
	var vValidationRegex []string
	var vValidationRegexn []string
	var vValidationRegexIc []string
	var vValidationRegexNic []string
	var vValidationRegexIe []string
	var vValidationRegexNie []string
	var vValidationRegexIsw []string
	var vValidationRegexNisw []string
	var vValidationRegexIew []string
	var vValidationRegexNiew []string
	var vValidationRegexEmpty *bool
	var vValidationRegexRegex []string
	var vValidationRegexIregex []string
	var vWeight []int64
	var vWeightn []int64
	var vWeightLt []int64
	var vWeightLte []int64
	var vWeightGt []int64
	var vWeightGte []int64
	var vWeightEmpty *bool
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		switch name {
		case "id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id' takes an integer, got %q.", value))
				return
			}
			vID = append(vID, v)
		case "id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__n' takes an integer, got %q.", value))
				return
			}
			vIDn = append(vIDn, v)
		case "id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lt' takes an integer, got %q.", value))
				return
			}
			vIDLt = append(vIDLt, v)
		case "id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lte' takes an integer, got %q.", value))
				return
			}
			vIDLte = append(vIDLte, v)
		case "id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gt' takes an integer, got %q.", value))
				return
			}
			vIDGt = append(vIDGt, v)
		case "id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gte' takes an integer, got %q.", value))
				return
			}
			vIDGte = append(vIDGte, v)
		case "id__empty":
			if vIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__empty' takes a boolean, got %q.", value))
				return
			}
			vIDEmpty = &v
		case "name":
			v := value
			vName = append(vName, v)
		case "name__n":
			v := value
			vNamen = append(vNamen, v)
		case "name__ic":
			v := value
			vNameIc = append(vNameIc, v)
		case "name__nic":
			v := value
			vNameNic = append(vNameNic, v)
		case "name__ie":
			v := value
			vNameIe = append(vNameIe, v)
		case "name__nie":
			v := value
			vNameNie = append(vNameNie, v)
		case "name__isw":
			v := value
			vNameIsw = append(vNameIsw, v)
		case "name__nisw":
			v := value
			vNameNisw = append(vNameNisw, v)
		case "name__iew":
			v := value
			vNameIew = append(vNameIew, v)
		case "name__niew":
			v := value
			vNameNiew = append(vNameNiew, v)
		case "name__empty":
			if vNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'name__empty' takes a boolean, got %q.", value))
				return
			}
			vNameEmpty = &v
		case "name__regex":
			v := value
			vNameRegex = append(vNameRegex, v)
		case "name__iregex":
			v := value
			vNameIregex = append(vNameIregex, v)
		case "choice_set":
			v := value
			vChoiceSet = append(vChoiceSet, v)
		case "choice_set__n":
			v := value
			vChoiceSetn = append(vChoiceSetn, v)
		case "choice_set_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'choice_set_id' takes an integer, got %q.", value))
				return
			}
			vChoiceSetID = append(vChoiceSetID, v)
		case "choice_set_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'choice_set_id__n' takes an integer, got %q.", value))
				return
			}
			vChoiceSetIDn = append(vChoiceSetIDn, v)
		case "created":
			v := value
			vCreated = append(vCreated, v)
		case "created__n":
			v := value
			vCreatedn = append(vCreatedn, v)
		case "created__lt":
			v := value
			vCreatedLt = append(vCreatedLt, v)
		case "created__lte":
			v := value
			vCreatedLte = append(vCreatedLte, v)
		case "created__gt":
			v := value
			vCreatedGt = append(vCreatedGt, v)
		case "created__gte":
			v := value
			vCreatedGte = append(vCreatedGte, v)
		case "created__empty":
			v := value
			vCreatedEmpty = append(vCreatedEmpty, v)
		case "description":
			v := value
			vDescription = append(vDescription, v)
		case "description__n":
			v := value
			vDescriptionn = append(vDescriptionn, v)
		case "description__ic":
			v := value
			vDescriptionIc = append(vDescriptionIc, v)
		case "description__nic":
			v := value
			vDescriptionNic = append(vDescriptionNic, v)
		case "description__ie":
			v := value
			vDescriptionIe = append(vDescriptionIe, v)
		case "description__nie":
			v := value
			vDescriptionNie = append(vDescriptionNie, v)
		case "description__isw":
			v := value
			vDescriptionIsw = append(vDescriptionIsw, v)
		case "description__nisw":
			v := value
			vDescriptionNisw = append(vDescriptionNisw, v)
		case "description__iew":
			v := value
			vDescriptionIew = append(vDescriptionIew, v)
		case "description__niew":
			v := value
			vDescriptionNiew = append(vDescriptionNiew, v)
		case "description__empty":
			if vDescriptionEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'description__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'description__empty' takes a boolean, got %q.", value))
				return
			}
			vDescriptionEmpty = &v
		case "description__regex":
			v := value
			vDescriptionRegex = append(vDescriptionRegex, v)
		case "description__iregex":
			v := value
			vDescriptionIregex = append(vDescriptionIregex, v)
		case "filter_logic":
			if vFilterLogic != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'filter_logic' takes a single value.")
				return
			}
			v := value
			vFilterLogic = &v
		case "filter_logic__n":
			if vFilterLogicn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'filter_logic__n' takes a single value.")
				return
			}
			v := value
			vFilterLogicn = &v
		case "filter_logic__ic":
			v := value
			vFilterLogicIc = append(vFilterLogicIc, v)
		case "filter_logic__nic":
			v := value
			vFilterLogicNic = append(vFilterLogicNic, v)
		case "filter_logic__ie":
			v := value
			vFilterLogicIe = append(vFilterLogicIe, v)
		case "filter_logic__nie":
			v := value
			vFilterLogicNie = append(vFilterLogicNie, v)
		case "filter_logic__isw":
			v := value
			vFilterLogicIsw = append(vFilterLogicIsw, v)
		case "filter_logic__nisw":
			v := value
			vFilterLogicNisw = append(vFilterLogicNisw, v)
		case "filter_logic__iew":
			v := value
			vFilterLogicIew = append(vFilterLogicIew, v)
		case "filter_logic__niew":
			v := value
			vFilterLogicNiew = append(vFilterLogicNiew, v)
		case "filter_logic__empty":
			if vFilterLogicEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'filter_logic__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'filter_logic__empty' takes a boolean, got %q.", value))
				return
			}
			vFilterLogicEmpty = &v
		case "filter_logic__regex":
			v := value
			vFilterLogicRegex = append(vFilterLogicRegex, v)
		case "filter_logic__iregex":
			v := value
			vFilterLogicIregex = append(vFilterLogicIregex, v)
		case "group_name":
			v := value
			vGroupName = append(vGroupName, v)
		case "group_name__n":
			v := value
			vGroupNamen = append(vGroupNamen, v)
		case "group_name__ic":
			v := value
			vGroupNameIc = append(vGroupNameIc, v)
		case "group_name__nic":
			v := value
			vGroupNameNic = append(vGroupNameNic, v)
		case "group_name__ie":
			v := value
			vGroupNameIe = append(vGroupNameIe, v)
		case "group_name__nie":
			v := value
			vGroupNameNie = append(vGroupNameNie, v)
		case "group_name__isw":
			v := value
			vGroupNameIsw = append(vGroupNameIsw, v)
		case "group_name__nisw":
			v := value
			vGroupNameNisw = append(vGroupNameNisw, v)
		case "group_name__iew":
			v := value
			vGroupNameIew = append(vGroupNameIew, v)
		case "group_name__niew":
			v := value
			vGroupNameNiew = append(vGroupNameNiew, v)
		case "group_name__empty":
			if vGroupNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'group_name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_name__empty' takes a boolean, got %q.", value))
				return
			}
			vGroupNameEmpty = &v
		case "group_name__regex":
			v := value
			vGroupNameRegex = append(vGroupNameRegex, v)
		case "group_name__iregex":
			v := value
			vGroupNameIregex = append(vGroupNameIregex, v)
		case "is_cloneable":
			if vIsCloneable != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'is_cloneable' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'is_cloneable' takes a boolean, got %q.", value))
				return
			}
			vIsCloneable = &v
		case "label":
			v := value
			vLabel = append(vLabel, v)
		case "label__n":
			v := value
			vLabeln = append(vLabeln, v)
		case "label__ic":
			v := value
			vLabelIc = append(vLabelIc, v)
		case "label__nic":
			v := value
			vLabelNic = append(vLabelNic, v)
		case "label__ie":
			v := value
			vLabelIe = append(vLabelIe, v)
		case "label__nie":
			v := value
			vLabelNie = append(vLabelNie, v)
		case "label__isw":
			v := value
			vLabelIsw = append(vLabelIsw, v)
		case "label__nisw":
			v := value
			vLabelNisw = append(vLabelNisw, v)
		case "label__iew":
			v := value
			vLabelIew = append(vLabelIew, v)
		case "label__niew":
			v := value
			vLabelNiew = append(vLabelNiew, v)
		case "label__empty":
			if vLabelEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'label__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'label__empty' takes a boolean, got %q.", value))
				return
			}
			vLabelEmpty = &v
		case "label__regex":
			v := value
			vLabelRegex = append(vLabelRegex, v)
		case "label__iregex":
			v := value
			vLabelIregex = append(vLabelIregex, v)
		case "last_updated":
			v := value
			vLastUpdated = append(vLastUpdated, v)
		case "last_updated__n":
			v := value
			vLastUpdatedn = append(vLastUpdatedn, v)
		case "last_updated__lt":
			v := value
			vLastUpdatedLt = append(vLastUpdatedLt, v)
		case "last_updated__lte":
			v := value
			vLastUpdatedLte = append(vLastUpdatedLte, v)
		case "last_updated__gt":
			v := value
			vLastUpdatedGt = append(vLastUpdatedGt, v)
		case "last_updated__gte":
			v := value
			vLastUpdatedGte = append(vLastUpdatedGte, v)
		case "last_updated__empty":
			v := value
			vLastUpdatedEmpty = append(vLastUpdatedEmpty, v)
		case "object_type":
			v := value
			vObjectType = append(vObjectType, v)
		case "object_type__n":
			v := value
			vObjectTypen = append(vObjectTypen, v)
		case "object_type__ic":
			v := value
			vObjectTypeIc = append(vObjectTypeIc, v)
		case "object_type__nic":
			v := value
			vObjectTypeNic = append(vObjectTypeNic, v)
		case "object_type__ie":
			v := value
			vObjectTypeIe = append(vObjectTypeIe, v)
		case "object_type__nie":
			v := value
			vObjectTypeNie = append(vObjectTypeNie, v)
		case "object_type__isw":
			v := value
			vObjectTypeIsw = append(vObjectTypeIsw, v)
		case "object_type__nisw":
			v := value
			vObjectTypeNisw = append(vObjectTypeNisw, v)
		case "object_type__iew":
			v := value
			vObjectTypeIew = append(vObjectTypeIew, v)
		case "object_type__niew":
			v := value
			vObjectTypeNiew = append(vObjectTypeNiew, v)
		case "object_type__regex":
			v := value
			vObjectTypeRegex = append(vObjectTypeRegex, v)
		case "object_type__iregex":
			v := value
			vObjectTypeIregex = append(vObjectTypeIregex, v)
		case "object_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id' takes an integer, got %q.", value))
				return
			}
			vObjectTypeID = append(vObjectTypeID, v)
		case "object_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id__n' takes an integer, got %q.", value))
				return
			}
			vObjectTypeIDn = append(vObjectTypeIDn, v)
		case "owner":
			v := value
			vOwner = append(vOwner, v)
		case "owner__n":
			v := value
			vOwnern = append(vOwnern, v)
		case "owner_group":
			v := value
			vOwnerGroup = append(vOwnerGroup, v)
		case "owner_group__n":
			v := value
			vOwnerGroupn = append(vOwnerGroupn, v)
		case "owner_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupID = append(vOwnerGroupID, v)
		case "owner_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupIDn = append(vOwnerGroupIDn, v)
		case "owner_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id' takes an integer, got %q.", value))
				return
			}
			vOwnerID = append(vOwnerID, v)
		case "owner_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerIDn = append(vOwnerIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "related_object_type":
			v := value
			vRelatedObjectType = append(vRelatedObjectType, v)
		case "related_object_type__n":
			v := value
			vRelatedObjectTypen = append(vRelatedObjectTypen, v)
		case "related_object_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'related_object_type_id' takes an integer, got %q.", value))
				return
			}
			vRelatedObjectTypeID = append(vRelatedObjectTypeID, v)
		case "related_object_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'related_object_type_id__n' takes an integer, got %q.", value))
				return
			}
			vRelatedObjectTypeIDn = append(vRelatedObjectTypeIDn, v)
		case "required":
			if vRequired != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'required' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'required' takes a boolean, got %q.", value))
				return
			}
			vRequired = &v
		case "search_weight":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'search_weight' takes an integer, got %q.", value))
				return
			}
			vSearchWeight = append(vSearchWeight, v)
		case "search_weight__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'search_weight__n' takes an integer, got %q.", value))
				return
			}
			vSearchWeightn = append(vSearchWeightn, v)
		case "search_weight__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'search_weight__lt' takes an integer, got %q.", value))
				return
			}
			vSearchWeightLt = append(vSearchWeightLt, v)
		case "search_weight__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'search_weight__lte' takes an integer, got %q.", value))
				return
			}
			vSearchWeightLte = append(vSearchWeightLte, v)
		case "search_weight__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'search_weight__gt' takes an integer, got %q.", value))
				return
			}
			vSearchWeightGt = append(vSearchWeightGt, v)
		case "search_weight__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'search_weight__gte' takes an integer, got %q.", value))
				return
			}
			vSearchWeightGte = append(vSearchWeightGte, v)
		case "search_weight__empty":
			if vSearchWeightEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'search_weight__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'search_weight__empty' takes a boolean, got %q.", value))
				return
			}
			vSearchWeightEmpty = &v
		case "type":
			v := value
			vType = append(vType, v)
		case "type__n":
			v := value
			vTypen = append(vTypen, v)
		case "type__ic":
			v := value
			vTypeIc = append(vTypeIc, v)
		case "type__nic":
			v := value
			vTypeNic = append(vTypeNic, v)
		case "type__ie":
			v := value
			vTypeIe = append(vTypeIe, v)
		case "type__nie":
			v := value
			vTypeNie = append(vTypeNie, v)
		case "type__isw":
			v := value
			vTypeIsw = append(vTypeIsw, v)
		case "type__nisw":
			v := value
			vTypeNisw = append(vTypeNisw, v)
		case "type__iew":
			v := value
			vTypeIew = append(vTypeIew, v)
		case "type__niew":
			v := value
			vTypeNiew = append(vTypeNiew, v)
		case "type__empty":
			if vTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'type__empty' takes a boolean, got %q.", value))
				return
			}
			vTypeEmpty = &v
		case "type__regex":
			v := value
			vTypeRegex = append(vTypeRegex, v)
		case "type__iregex":
			v := value
			vTypeIregex = append(vTypeIregex, v)
		case "ui_editable":
			if vUIEditable != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ui_editable' takes a single value.")
				return
			}
			v := value
			vUIEditable = &v
		case "ui_editable__n":
			if vUIEditablen != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ui_editable__n' takes a single value.")
				return
			}
			v := value
			vUIEditablen = &v
		case "ui_editable__ic":
			v := value
			vUIEditableIc = append(vUIEditableIc, v)
		case "ui_editable__nic":
			v := value
			vUIEditableNic = append(vUIEditableNic, v)
		case "ui_editable__ie":
			v := value
			vUIEditableIe = append(vUIEditableIe, v)
		case "ui_editable__nie":
			v := value
			vUIEditableNie = append(vUIEditableNie, v)
		case "ui_editable__isw":
			v := value
			vUIEditableIsw = append(vUIEditableIsw, v)
		case "ui_editable__nisw":
			v := value
			vUIEditableNisw = append(vUIEditableNisw, v)
		case "ui_editable__iew":
			v := value
			vUIEditableIew = append(vUIEditableIew, v)
		case "ui_editable__niew":
			v := value
			vUIEditableNiew = append(vUIEditableNiew, v)
		case "ui_editable__empty":
			if vUIEditableEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ui_editable__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ui_editable__empty' takes a boolean, got %q.", value))
				return
			}
			vUIEditableEmpty = &v
		case "ui_editable__regex":
			v := value
			vUIEditableRegex = append(vUIEditableRegex, v)
		case "ui_editable__iregex":
			v := value
			vUIEditableIregex = append(vUIEditableIregex, v)
		case "ui_visible":
			if vUIVisible != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ui_visible' takes a single value.")
				return
			}
			v := value
			vUIVisible = &v
		case "ui_visible__n":
			if vUIVisiblen != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ui_visible__n' takes a single value.")
				return
			}
			v := value
			vUIVisiblen = &v
		case "ui_visible__ic":
			v := value
			vUIVisibleIc = append(vUIVisibleIc, v)
		case "ui_visible__nic":
			v := value
			vUIVisibleNic = append(vUIVisibleNic, v)
		case "ui_visible__ie":
			v := value
			vUIVisibleIe = append(vUIVisibleIe, v)
		case "ui_visible__nie":
			v := value
			vUIVisibleNie = append(vUIVisibleNie, v)
		case "ui_visible__isw":
			v := value
			vUIVisibleIsw = append(vUIVisibleIsw, v)
		case "ui_visible__nisw":
			v := value
			vUIVisibleNisw = append(vUIVisibleNisw, v)
		case "ui_visible__iew":
			v := value
			vUIVisibleIew = append(vUIVisibleIew, v)
		case "ui_visible__niew":
			v := value
			vUIVisibleNiew = append(vUIVisibleNiew, v)
		case "ui_visible__empty":
			if vUIVisibleEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ui_visible__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ui_visible__empty' takes a boolean, got %q.", value))
				return
			}
			vUIVisibleEmpty = &v
		case "ui_visible__regex":
			v := value
			vUIVisibleRegex = append(vUIVisibleRegex, v)
		case "ui_visible__iregex":
			v := value
			vUIVisibleIregex = append(vUIVisibleIregex, v)
		case "unique":
			if vUnique != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'unique' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'unique' takes a boolean, got %q.", value))
				return
			}
			vUnique = &v
		case "validation_maximum":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_maximum' takes a number, got %q.", value))
				return
			}
			vValidationMaximum = append(vValidationMaximum, v)
		case "validation_maximum__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_maximum__n' takes a number, got %q.", value))
				return
			}
			vValidationMaximumn = append(vValidationMaximumn, v)
		case "validation_maximum__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_maximum__lt' takes a number, got %q.", value))
				return
			}
			vValidationMaximumLt = append(vValidationMaximumLt, v)
		case "validation_maximum__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_maximum__lte' takes a number, got %q.", value))
				return
			}
			vValidationMaximumLte = append(vValidationMaximumLte, v)
		case "validation_maximum__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_maximum__gt' takes a number, got %q.", value))
				return
			}
			vValidationMaximumGt = append(vValidationMaximumGt, v)
		case "validation_maximum__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_maximum__gte' takes a number, got %q.", value))
				return
			}
			vValidationMaximumGte = append(vValidationMaximumGte, v)
		case "validation_maximum__empty":
			if vValidationMaximumEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'validation_maximum__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_maximum__empty' takes a boolean, got %q.", value))
				return
			}
			vValidationMaximumEmpty = &v
		case "validation_minimum":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_minimum' takes a number, got %q.", value))
				return
			}
			vValidationMinimum = append(vValidationMinimum, v)
		case "validation_minimum__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_minimum__n' takes a number, got %q.", value))
				return
			}
			vValidationMinimumn = append(vValidationMinimumn, v)
		case "validation_minimum__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_minimum__lt' takes a number, got %q.", value))
				return
			}
			vValidationMinimumLt = append(vValidationMinimumLt, v)
		case "validation_minimum__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_minimum__lte' takes a number, got %q.", value))
				return
			}
			vValidationMinimumLte = append(vValidationMinimumLte, v)
		case "validation_minimum__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_minimum__gt' takes a number, got %q.", value))
				return
			}
			vValidationMinimumGt = append(vValidationMinimumGt, v)
		case "validation_minimum__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_minimum__gte' takes a number, got %q.", value))
				return
			}
			vValidationMinimumGte = append(vValidationMinimumGte, v)
		case "validation_minimum__empty":
			if vValidationMinimumEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'validation_minimum__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_minimum__empty' takes a boolean, got %q.", value))
				return
			}
			vValidationMinimumEmpty = &v
		case "validation_regex":
			v := value
			vValidationRegex = append(vValidationRegex, v)
		case "validation_regex__n":
			v := value
			vValidationRegexn = append(vValidationRegexn, v)
		case "validation_regex__ic":
			v := value
			vValidationRegexIc = append(vValidationRegexIc, v)
		case "validation_regex__nic":
			v := value
			vValidationRegexNic = append(vValidationRegexNic, v)
		case "validation_regex__ie":
			v := value
			vValidationRegexIe = append(vValidationRegexIe, v)
		case "validation_regex__nie":
			v := value
			vValidationRegexNie = append(vValidationRegexNie, v)
		case "validation_regex__isw":
			v := value
			vValidationRegexIsw = append(vValidationRegexIsw, v)
		case "validation_regex__nisw":
			v := value
			vValidationRegexNisw = append(vValidationRegexNisw, v)
		case "validation_regex__iew":
			v := value
			vValidationRegexIew = append(vValidationRegexIew, v)
		case "validation_regex__niew":
			v := value
			vValidationRegexNiew = append(vValidationRegexNiew, v)
		case "validation_regex__empty":
			if vValidationRegexEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'validation_regex__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'validation_regex__empty' takes a boolean, got %q.", value))
				return
			}
			vValidationRegexEmpty = &v
		case "validation_regex__regex":
			v := value
			vValidationRegexRegex = append(vValidationRegexRegex, v)
		case "validation_regex__iregex":
			v := value
			vValidationRegexIregex = append(vValidationRegexIregex, v)
		case "weight":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight' takes an integer, got %q.", value))
				return
			}
			vWeight = append(vWeight, v)
		case "weight__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__n' takes an integer, got %q.", value))
				return
			}
			vWeightn = append(vWeightn, v)
		case "weight__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__lt' takes an integer, got %q.", value))
				return
			}
			vWeightLt = append(vWeightLt, v)
		case "weight__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__lte' takes an integer, got %q.", value))
				return
			}
			vWeightLte = append(vWeightLte, v)
		case "weight__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__gt' takes an integer, got %q.", value))
				return
			}
			vWeightGt = append(vWeightGt, v)
		case "weight__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__gte' takes an integer, got %q.", value))
				return
			}
			vWeightGte = append(vWeightGte, v)
		case "weight__empty":
			if vWeightEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'weight__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__empty' takes a boolean, got %q.", value))
				return
			}
			vWeightEmpty = &v
		default:
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
	}
	if len(vID) > 0 {
		params.SetID(vID)
	}
	if len(vIDn) > 0 {
		params.SetIDn(vIDn)
	}
	if len(vIDLt) > 0 {
		params.SetIDLt(vIDLt)
	}
	if len(vIDLte) > 0 {
		params.SetIDLte(vIDLte)
	}
	if len(vIDGt) > 0 {
		params.SetIDGt(vIDGt)
	}
	if len(vIDGte) > 0 {
		params.SetIDGte(vIDGte)
	}
	if vIDEmpty != nil {
		params.SetIDEmpty(vIDEmpty)
	}
	if len(vName) > 0 {
		params.SetName(vName)
	}
	if len(vNamen) > 0 {
		params.SetNamen(vNamen)
	}
	if len(vNameIc) > 0 {
		params.SetNameIc(vNameIc)
	}
	if len(vNameNic) > 0 {
		params.SetNameNic(vNameNic)
	}
	if len(vNameIe) > 0 {
		params.SetNameIe(vNameIe)
	}
	if len(vNameNie) > 0 {
		params.SetNameNie(vNameNie)
	}
	if len(vNameIsw) > 0 {
		params.SetNameIsw(vNameIsw)
	}
	if len(vNameNisw) > 0 {
		params.SetNameNisw(vNameNisw)
	}
	if len(vNameIew) > 0 {
		params.SetNameIew(vNameIew)
	}
	if len(vNameNiew) > 0 {
		params.SetNameNiew(vNameNiew)
	}
	if vNameEmpty != nil {
		params.SetNameEmpty(vNameEmpty)
	}
	if len(vNameRegex) > 0 {
		params.SetNameRegex(vNameRegex)
	}
	if len(vNameIregex) > 0 {
		params.SetNameIregex(vNameIregex)
	}
	if len(vChoiceSet) > 0 {
		params.SetChoiceSet(vChoiceSet)
	}
	if len(vChoiceSetn) > 0 {
		params.SetChoiceSetn(vChoiceSetn)
	}
	if len(vChoiceSetID) > 0 {
		params.SetChoiceSetID(vChoiceSetID)
	}
	if len(vChoiceSetIDn) > 0 {
		params.SetChoiceSetIDn(vChoiceSetIDn)
	}
	if len(vCreated) > 0 {
		params.SetCreated(vCreated)
	}
	if len(vCreatedn) > 0 {
		params.SetCreatedn(vCreatedn)
	}
	if len(vCreatedLt) > 0 {
		params.SetCreatedLt(vCreatedLt)
	}
	if len(vCreatedLte) > 0 {
		params.SetCreatedLte(vCreatedLte)
	}
	if len(vCreatedGt) > 0 {
		params.SetCreatedGt(vCreatedGt)
	}
	if len(vCreatedGte) > 0 {
		params.SetCreatedGte(vCreatedGte)
	}
	if len(vCreatedEmpty) > 0 {
		params.SetCreatedEmpty(vCreatedEmpty)
	}
	if len(vDescription) > 0 {
		params.SetDescription(vDescription)
	}
	if len(vDescriptionn) > 0 {
		params.SetDescriptionn(vDescriptionn)
	}
	if len(vDescriptionIc) > 0 {
		params.SetDescriptionIc(vDescriptionIc)
	}
	if len(vDescriptionNic) > 0 {
		params.SetDescriptionNic(vDescriptionNic)
	}
	if len(vDescriptionIe) > 0 {
		params.SetDescriptionIe(vDescriptionIe)
	}
	if len(vDescriptionNie) > 0 {
		params.SetDescriptionNie(vDescriptionNie)
	}
	if len(vDescriptionIsw) > 0 {
		params.SetDescriptionIsw(vDescriptionIsw)
	}
	if len(vDescriptionNisw) > 0 {
		params.SetDescriptionNisw(vDescriptionNisw)
	}
	if len(vDescriptionIew) > 0 {
		params.SetDescriptionIew(vDescriptionIew)
	}
	if len(vDescriptionNiew) > 0 {
		params.SetDescriptionNiew(vDescriptionNiew)
	}
	if vDescriptionEmpty != nil {
		params.SetDescriptionEmpty(vDescriptionEmpty)
	}
	if len(vDescriptionRegex) > 0 {
		params.SetDescriptionRegex(vDescriptionRegex)
	}
	if len(vDescriptionIregex) > 0 {
		params.SetDescriptionIregex(vDescriptionIregex)
	}
	if vFilterLogic != nil {
		params.SetFilterLogic(vFilterLogic)
	}
	if vFilterLogicn != nil {
		params.SetFilterLogicn(vFilterLogicn)
	}
	if len(vFilterLogicIc) > 0 {
		params.SetFilterLogicIc(vFilterLogicIc)
	}
	if len(vFilterLogicNic) > 0 {
		params.SetFilterLogicNic(vFilterLogicNic)
	}
	if len(vFilterLogicIe) > 0 {
		params.SetFilterLogicIe(vFilterLogicIe)
	}
	if len(vFilterLogicNie) > 0 {
		params.SetFilterLogicNie(vFilterLogicNie)
	}
	if len(vFilterLogicIsw) > 0 {
		params.SetFilterLogicIsw(vFilterLogicIsw)
	}
	if len(vFilterLogicNisw) > 0 {
		params.SetFilterLogicNisw(vFilterLogicNisw)
	}
	if len(vFilterLogicIew) > 0 {
		params.SetFilterLogicIew(vFilterLogicIew)
	}
	if len(vFilterLogicNiew) > 0 {
		params.SetFilterLogicNiew(vFilterLogicNiew)
	}
	if vFilterLogicEmpty != nil {
		params.SetFilterLogicEmpty(vFilterLogicEmpty)
	}
	if len(vFilterLogicRegex) > 0 {
		params.SetFilterLogicRegex(vFilterLogicRegex)
	}
	if len(vFilterLogicIregex) > 0 {
		params.SetFilterLogicIregex(vFilterLogicIregex)
	}
	if len(vGroupName) > 0 {
		params.SetGroupName(vGroupName)
	}
	if len(vGroupNamen) > 0 {
		params.SetGroupNamen(vGroupNamen)
	}
	if len(vGroupNameIc) > 0 {
		params.SetGroupNameIc(vGroupNameIc)
	}
	if len(vGroupNameNic) > 0 {
		params.SetGroupNameNic(vGroupNameNic)
	}
	if len(vGroupNameIe) > 0 {
		params.SetGroupNameIe(vGroupNameIe)
	}
	if len(vGroupNameNie) > 0 {
		params.SetGroupNameNie(vGroupNameNie)
	}
	if len(vGroupNameIsw) > 0 {
		params.SetGroupNameIsw(vGroupNameIsw)
	}
	if len(vGroupNameNisw) > 0 {
		params.SetGroupNameNisw(vGroupNameNisw)
	}
	if len(vGroupNameIew) > 0 {
		params.SetGroupNameIew(vGroupNameIew)
	}
	if len(vGroupNameNiew) > 0 {
		params.SetGroupNameNiew(vGroupNameNiew)
	}
	if vGroupNameEmpty != nil {
		params.SetGroupNameEmpty(vGroupNameEmpty)
	}
	if len(vGroupNameRegex) > 0 {
		params.SetGroupNameRegex(vGroupNameRegex)
	}
	if len(vGroupNameIregex) > 0 {
		params.SetGroupNameIregex(vGroupNameIregex)
	}
	if vIsCloneable != nil {
		params.SetIsCloneable(vIsCloneable)
	}
	if len(vLabel) > 0 {
		params.SetLabel(vLabel)
	}
	if len(vLabeln) > 0 {
		params.SetLabeln(vLabeln)
	}
	if len(vLabelIc) > 0 {
		params.SetLabelIc(vLabelIc)
	}
	if len(vLabelNic) > 0 {
		params.SetLabelNic(vLabelNic)
	}
	if len(vLabelIe) > 0 {
		params.SetLabelIe(vLabelIe)
	}
	if len(vLabelNie) > 0 {
		params.SetLabelNie(vLabelNie)
	}
	if len(vLabelIsw) > 0 {
		params.SetLabelIsw(vLabelIsw)
	}
	if len(vLabelNisw) > 0 {
		params.SetLabelNisw(vLabelNisw)
	}
	if len(vLabelIew) > 0 {
		params.SetLabelIew(vLabelIew)
	}
	if len(vLabelNiew) > 0 {
		params.SetLabelNiew(vLabelNiew)
	}
	if vLabelEmpty != nil {
		params.SetLabelEmpty(vLabelEmpty)
	}
	if len(vLabelRegex) > 0 {
		params.SetLabelRegex(vLabelRegex)
	}
	if len(vLabelIregex) > 0 {
		params.SetLabelIregex(vLabelIregex)
	}
	if len(vLastUpdated) > 0 {
		params.SetLastUpdated(vLastUpdated)
	}
	if len(vLastUpdatedn) > 0 {
		params.SetLastUpdatedn(vLastUpdatedn)
	}
	if len(vLastUpdatedLt) > 0 {
		params.SetLastUpdatedLt(vLastUpdatedLt)
	}
	if len(vLastUpdatedLte) > 0 {
		params.SetLastUpdatedLte(vLastUpdatedLte)
	}
	if len(vLastUpdatedGt) > 0 {
		params.SetLastUpdatedGt(vLastUpdatedGt)
	}
	if len(vLastUpdatedGte) > 0 {
		params.SetLastUpdatedGte(vLastUpdatedGte)
	}
	if len(vLastUpdatedEmpty) > 0 {
		params.SetLastUpdatedEmpty(vLastUpdatedEmpty)
	}
	if len(vObjectType) > 0 {
		params.SetObjectType(vObjectType)
	}
	if len(vObjectTypen) > 0 {
		params.SetObjectTypen(vObjectTypen)
	}
	if len(vObjectTypeIc) > 0 {
		params.SetObjectTypeIc(vObjectTypeIc)
	}
	if len(vObjectTypeNic) > 0 {
		params.SetObjectTypeNic(vObjectTypeNic)
	}
	if len(vObjectTypeIe) > 0 {
		params.SetObjectTypeIe(vObjectTypeIe)
	}
	if len(vObjectTypeNie) > 0 {
		params.SetObjectTypeNie(vObjectTypeNie)
	}
	if len(vObjectTypeIsw) > 0 {
		params.SetObjectTypeIsw(vObjectTypeIsw)
	}
	if len(vObjectTypeNisw) > 0 {
		params.SetObjectTypeNisw(vObjectTypeNisw)
	}
	if len(vObjectTypeIew) > 0 {
		params.SetObjectTypeIew(vObjectTypeIew)
	}
	if len(vObjectTypeNiew) > 0 {
		params.SetObjectTypeNiew(vObjectTypeNiew)
	}
	if len(vObjectTypeRegex) > 0 {
		params.SetObjectTypeRegex(vObjectTypeRegex)
	}
	if len(vObjectTypeIregex) > 0 {
		params.SetObjectTypeIregex(vObjectTypeIregex)
	}
	if len(vObjectTypeID) > 0 {
		params.SetObjectTypeID(vObjectTypeID)
	}
	if len(vObjectTypeIDn) > 0 {
		params.SetObjectTypeIDn(vObjectTypeIDn)
	}
	if len(vOwner) > 0 {
		params.SetOwner(vOwner)
	}
	if len(vOwnern) > 0 {
		params.SetOwnern(vOwnern)
	}
	if len(vOwnerGroup) > 0 {
		params.SetOwnerGroup(vOwnerGroup)
	}
	if len(vOwnerGroupn) > 0 {
		params.SetOwnerGroupn(vOwnerGroupn)
	}
	if len(vOwnerGroupID) > 0 {
		params.SetOwnerGroupID(vOwnerGroupID)
	}
	if len(vOwnerGroupIDn) > 0 {
		params.SetOwnerGroupIDn(vOwnerGroupIDn)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRelatedObjectType) > 0 {
		params.SetRelatedObjectType(vRelatedObjectType)
	}
	if len(vRelatedObjectTypen) > 0 {
		params.SetRelatedObjectTypen(vRelatedObjectTypen)
	}
	if len(vRelatedObjectTypeID) > 0 {
		params.SetRelatedObjectTypeID(vRelatedObjectTypeID)
	}
	if len(vRelatedObjectTypeIDn) > 0 {
		params.SetRelatedObjectTypeIDn(vRelatedObjectTypeIDn)
	}
	if vRequired != nil {
		params.SetRequired(vRequired)
	}
	if len(vSearchWeight) > 0 {
		params.SetSearchWeight(vSearchWeight)
	}
	if len(vSearchWeightn) > 0 {
		params.SetSearchWeightn(vSearchWeightn)
	}
	if len(vSearchWeightLt) > 0 {
		params.SetSearchWeightLt(vSearchWeightLt)
	}
	if len(vSearchWeightLte) > 0 {
		params.SetSearchWeightLte(vSearchWeightLte)
	}
	if len(vSearchWeightGt) > 0 {
		params.SetSearchWeightGt(vSearchWeightGt)
	}
	if len(vSearchWeightGte) > 0 {
		params.SetSearchWeightGte(vSearchWeightGte)
	}
	if vSearchWeightEmpty != nil {
		params.SetSearchWeightEmpty(vSearchWeightEmpty)
	}
	if len(vType) > 0 {
		params.SetType(vType)
	}
	if len(vTypen) > 0 {
		params.SetTypen(vTypen)
	}
	if len(vTypeIc) > 0 {
		params.SetTypeIc(vTypeIc)
	}
	if len(vTypeNic) > 0 {
		params.SetTypeNic(vTypeNic)
	}
	if len(vTypeIe) > 0 {
		params.SetTypeIe(vTypeIe)
	}
	if len(vTypeNie) > 0 {
		params.SetTypeNie(vTypeNie)
	}
	if len(vTypeIsw) > 0 {
		params.SetTypeIsw(vTypeIsw)
	}
	if len(vTypeNisw) > 0 {
		params.SetTypeNisw(vTypeNisw)
	}
	if len(vTypeIew) > 0 {
		params.SetTypeIew(vTypeIew)
	}
	if len(vTypeNiew) > 0 {
		params.SetTypeNiew(vTypeNiew)
	}
	if vTypeEmpty != nil {
		params.SetTypeEmpty(vTypeEmpty)
	}
	if len(vTypeRegex) > 0 {
		params.SetTypeRegex(vTypeRegex)
	}
	if len(vTypeIregex) > 0 {
		params.SetTypeIregex(vTypeIregex)
	}
	if vUIEditable != nil {
		params.SetUIEditable(vUIEditable)
	}
	if vUIEditablen != nil {
		params.SetUIEditablen(vUIEditablen)
	}
	if len(vUIEditableIc) > 0 {
		params.SetUIEditableIc(vUIEditableIc)
	}
	if len(vUIEditableNic) > 0 {
		params.SetUIEditableNic(vUIEditableNic)
	}
	if len(vUIEditableIe) > 0 {
		params.SetUIEditableIe(vUIEditableIe)
	}
	if len(vUIEditableNie) > 0 {
		params.SetUIEditableNie(vUIEditableNie)
	}
	if len(vUIEditableIsw) > 0 {
		params.SetUIEditableIsw(vUIEditableIsw)
	}
	if len(vUIEditableNisw) > 0 {
		params.SetUIEditableNisw(vUIEditableNisw)
	}
	if len(vUIEditableIew) > 0 {
		params.SetUIEditableIew(vUIEditableIew)
	}
	if len(vUIEditableNiew) > 0 {
		params.SetUIEditableNiew(vUIEditableNiew)
	}
	if vUIEditableEmpty != nil {
		params.SetUIEditableEmpty(vUIEditableEmpty)
	}
	if len(vUIEditableRegex) > 0 {
		params.SetUIEditableRegex(vUIEditableRegex)
	}
	if len(vUIEditableIregex) > 0 {
		params.SetUIEditableIregex(vUIEditableIregex)
	}
	if vUIVisible != nil {
		params.SetUIVisible(vUIVisible)
	}
	if vUIVisiblen != nil {
		params.SetUIVisiblen(vUIVisiblen)
	}
	if len(vUIVisibleIc) > 0 {
		params.SetUIVisibleIc(vUIVisibleIc)
	}
	if len(vUIVisibleNic) > 0 {
		params.SetUIVisibleNic(vUIVisibleNic)
	}
	if len(vUIVisibleIe) > 0 {
		params.SetUIVisibleIe(vUIVisibleIe)
	}
	if len(vUIVisibleNie) > 0 {
		params.SetUIVisibleNie(vUIVisibleNie)
	}
	if len(vUIVisibleIsw) > 0 {
		params.SetUIVisibleIsw(vUIVisibleIsw)
	}
	if len(vUIVisibleNisw) > 0 {
		params.SetUIVisibleNisw(vUIVisibleNisw)
	}
	if len(vUIVisibleIew) > 0 {
		params.SetUIVisibleIew(vUIVisibleIew)
	}
	if len(vUIVisibleNiew) > 0 {
		params.SetUIVisibleNiew(vUIVisibleNiew)
	}
	if vUIVisibleEmpty != nil {
		params.SetUIVisibleEmpty(vUIVisibleEmpty)
	}
	if len(vUIVisibleRegex) > 0 {
		params.SetUIVisibleRegex(vUIVisibleRegex)
	}
	if len(vUIVisibleIregex) > 0 {
		params.SetUIVisibleIregex(vUIVisibleIregex)
	}
	if vUnique != nil {
		params.SetUnique(vUnique)
	}
	if len(vValidationMaximum) > 0 {
		params.SetValidationMaximum(vValidationMaximum)
	}
	if len(vValidationMaximumn) > 0 {
		params.SetValidationMaximumn(vValidationMaximumn)
	}
	if len(vValidationMaximumLt) > 0 {
		params.SetValidationMaximumLt(vValidationMaximumLt)
	}
	if len(vValidationMaximumLte) > 0 {
		params.SetValidationMaximumLte(vValidationMaximumLte)
	}
	if len(vValidationMaximumGt) > 0 {
		params.SetValidationMaximumGt(vValidationMaximumGt)
	}
	if len(vValidationMaximumGte) > 0 {
		params.SetValidationMaximumGte(vValidationMaximumGte)
	}
	if vValidationMaximumEmpty != nil {
		params.SetValidationMaximumEmpty(vValidationMaximumEmpty)
	}
	if len(vValidationMinimum) > 0 {
		params.SetValidationMinimum(vValidationMinimum)
	}
	if len(vValidationMinimumn) > 0 {
		params.SetValidationMinimumn(vValidationMinimumn)
	}
	if len(vValidationMinimumLt) > 0 {
		params.SetValidationMinimumLt(vValidationMinimumLt)
	}
	if len(vValidationMinimumLte) > 0 {
		params.SetValidationMinimumLte(vValidationMinimumLte)
	}
	if len(vValidationMinimumGt) > 0 {
		params.SetValidationMinimumGt(vValidationMinimumGt)
	}
	if len(vValidationMinimumGte) > 0 {
		params.SetValidationMinimumGte(vValidationMinimumGte)
	}
	if vValidationMinimumEmpty != nil {
		params.SetValidationMinimumEmpty(vValidationMinimumEmpty)
	}
	if len(vValidationRegex) > 0 {
		params.SetValidationRegex(vValidationRegex)
	}
	if len(vValidationRegexn) > 0 {
		params.SetValidationRegexn(vValidationRegexn)
	}
	if len(vValidationRegexIc) > 0 {
		params.SetValidationRegexIc(vValidationRegexIc)
	}
	if len(vValidationRegexNic) > 0 {
		params.SetValidationRegexNic(vValidationRegexNic)
	}
	if len(vValidationRegexIe) > 0 {
		params.SetValidationRegexIe(vValidationRegexIe)
	}
	if len(vValidationRegexNie) > 0 {
		params.SetValidationRegexNie(vValidationRegexNie)
	}
	if len(vValidationRegexIsw) > 0 {
		params.SetValidationRegexIsw(vValidationRegexIsw)
	}
	if len(vValidationRegexNisw) > 0 {
		params.SetValidationRegexNisw(vValidationRegexNisw)
	}
	if len(vValidationRegexIew) > 0 {
		params.SetValidationRegexIew(vValidationRegexIew)
	}
	if len(vValidationRegexNiew) > 0 {
		params.SetValidationRegexNiew(vValidationRegexNiew)
	}
	if vValidationRegexEmpty != nil {
		params.SetValidationRegexEmpty(vValidationRegexEmpty)
	}
	if len(vValidationRegexRegex) > 0 {
		params.SetValidationRegexRegex(vValidationRegexRegex)
	}
	if len(vValidationRegexIregex) > 0 {
		params.SetValidationRegexIregex(vValidationRegexIregex)
	}
	if len(vWeight) > 0 {
		params.SetWeight(vWeight)
	}
	if len(vWeightn) > 0 {
		params.SetWeightn(vWeightn)
	}
	if len(vWeightLt) > 0 {
		params.SetWeightLt(vWeightLt)
	}
	if len(vWeightLte) > 0 {
		params.SetWeightLte(vWeightLte)
	}
	if len(vWeightGt) > 0 {
		params.SetWeightGt(vWeightGt)
	}
	if len(vWeightGte) > 0 {
		params.SetWeightGte(vWeightGte)
	}
	if vWeightEmpty != nil {
		params.SetWeightEmpty(vWeightEmpty)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Extras.ExtrasCustomFieldsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_custom_fields", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CustomFieldResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]customFieldResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model customFieldResourceModel
		resp.Diagnostics.Append(flattenCustomField(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: customFieldResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
