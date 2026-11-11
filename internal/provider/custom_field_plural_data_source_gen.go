// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                        netboxapi.FilterInt64s(params.SetID),
		"id__n":                     netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                    netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                   netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                    netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                   netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                 netboxapi.FilterBool(params.SetIDEmpty),
		"name":                      netboxapi.FilterStrings(params.SetName),
		"name__n":                   netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                  netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                 netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                  netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                 netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                 netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                 netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":               netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":               netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":              netboxapi.FilterStrings(params.SetNameIregex),
		"choice_set":                netboxapi.FilterStrings(params.SetChoiceSet),
		"choice_set__n":             netboxapi.FilterStrings(params.SetChoiceSetn),
		"choice_set_id":             netboxapi.FilterInt64s(params.SetChoiceSetID),
		"choice_set_id__n":          netboxapi.FilterInt64s(params.SetChoiceSetIDn),
		"created":                   netboxapi.FilterStrings(params.SetCreated),
		"created__n":                netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":               netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":              netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":               netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":              netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":            netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":               netboxapi.FilterStrings(params.SetDescription),
		"description__n":            netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":           netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":          netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":           netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":          netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":          netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":         netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":          netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":         netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":        netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":        netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":       netboxapi.FilterStrings(params.SetDescriptionIregex),
		"filter_logic":              netboxapi.FilterString(params.SetFilterLogic),
		"filter_logic__n":           netboxapi.FilterString(params.SetFilterLogicn),
		"filter_logic__ic":          netboxapi.FilterStrings(params.SetFilterLogicIc),
		"filter_logic__nic":         netboxapi.FilterStrings(params.SetFilterLogicNic),
		"filter_logic__ie":          netboxapi.FilterStrings(params.SetFilterLogicIe),
		"filter_logic__nie":         netboxapi.FilterStrings(params.SetFilterLogicNie),
		"filter_logic__isw":         netboxapi.FilterStrings(params.SetFilterLogicIsw),
		"filter_logic__nisw":        netboxapi.FilterStrings(params.SetFilterLogicNisw),
		"filter_logic__iew":         netboxapi.FilterStrings(params.SetFilterLogicIew),
		"filter_logic__niew":        netboxapi.FilterStrings(params.SetFilterLogicNiew),
		"filter_logic__empty":       netboxapi.FilterBool(params.SetFilterLogicEmpty),
		"filter_logic__regex":       netboxapi.FilterStrings(params.SetFilterLogicRegex),
		"filter_logic__iregex":      netboxapi.FilterStrings(params.SetFilterLogicIregex),
		"group_name":                netboxapi.FilterStrings(params.SetGroupName),
		"group_name__n":             netboxapi.FilterStrings(params.SetGroupNamen),
		"group_name__ic":            netboxapi.FilterStrings(params.SetGroupNameIc),
		"group_name__nic":           netboxapi.FilterStrings(params.SetGroupNameNic),
		"group_name__ie":            netboxapi.FilterStrings(params.SetGroupNameIe),
		"group_name__nie":           netboxapi.FilterStrings(params.SetGroupNameNie),
		"group_name__isw":           netboxapi.FilterStrings(params.SetGroupNameIsw),
		"group_name__nisw":          netboxapi.FilterStrings(params.SetGroupNameNisw),
		"group_name__iew":           netboxapi.FilterStrings(params.SetGroupNameIew),
		"group_name__niew":          netboxapi.FilterStrings(params.SetGroupNameNiew),
		"group_name__empty":         netboxapi.FilterBool(params.SetGroupNameEmpty),
		"group_name__regex":         netboxapi.FilterStrings(params.SetGroupNameRegex),
		"group_name__iregex":        netboxapi.FilterStrings(params.SetGroupNameIregex),
		"is_cloneable":              netboxapi.FilterBool(params.SetIsCloneable),
		"label":                     netboxapi.FilterStrings(params.SetLabel),
		"label__n":                  netboxapi.FilterStrings(params.SetLabeln),
		"label__ic":                 netboxapi.FilterStrings(params.SetLabelIc),
		"label__nic":                netboxapi.FilterStrings(params.SetLabelNic),
		"label__ie":                 netboxapi.FilterStrings(params.SetLabelIe),
		"label__nie":                netboxapi.FilterStrings(params.SetLabelNie),
		"label__isw":                netboxapi.FilterStrings(params.SetLabelIsw),
		"label__nisw":               netboxapi.FilterStrings(params.SetLabelNisw),
		"label__iew":                netboxapi.FilterStrings(params.SetLabelIew),
		"label__niew":               netboxapi.FilterStrings(params.SetLabelNiew),
		"label__empty":              netboxapi.FilterBool(params.SetLabelEmpty),
		"label__regex":              netboxapi.FilterStrings(params.SetLabelRegex),
		"label__iregex":             netboxapi.FilterStrings(params.SetLabelIregex),
		"last_updated":              netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":           netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":          netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":         netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":          netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":         netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":       netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"object_type":               netboxapi.FilterStrings(params.SetObjectType),
		"object_type__n":            netboxapi.FilterStrings(params.SetObjectTypen),
		"object_type__ic":           netboxapi.FilterStrings(params.SetObjectTypeIc),
		"object_type__nic":          netboxapi.FilterStrings(params.SetObjectTypeNic),
		"object_type__ie":           netboxapi.FilterStrings(params.SetObjectTypeIe),
		"object_type__nie":          netboxapi.FilterStrings(params.SetObjectTypeNie),
		"object_type__isw":          netboxapi.FilterStrings(params.SetObjectTypeIsw),
		"object_type__nisw":         netboxapi.FilterStrings(params.SetObjectTypeNisw),
		"object_type__iew":          netboxapi.FilterStrings(params.SetObjectTypeIew),
		"object_type__niew":         netboxapi.FilterStrings(params.SetObjectTypeNiew),
		"object_type__regex":        netboxapi.FilterStrings(params.SetObjectTypeRegex),
		"object_type__iregex":       netboxapi.FilterStrings(params.SetObjectTypeIregex),
		"object_type_id":            netboxapi.FilterInt64s(params.SetObjectTypeID),
		"object_type_id__n":         netboxapi.FilterInt64s(params.SetObjectTypeIDn),
		"owner":                     netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                  netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":               netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":            netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":            netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":         netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"owner_id":                  netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":               netboxapi.FilterInt64s(params.SetOwnerIDn),
		"q":                         netboxapi.FilterString(params.SetQ),
		"related_object_type":       netboxapi.FilterStrings(params.SetRelatedObjectType),
		"related_object_type__n":    netboxapi.FilterStrings(params.SetRelatedObjectTypen),
		"related_object_type_id":    netboxapi.FilterInt64s(params.SetRelatedObjectTypeID),
		"related_object_type_id__n": netboxapi.FilterInt64s(params.SetRelatedObjectTypeIDn),
		"required":                  netboxapi.FilterBool(params.SetRequired),
		"search_weight":             netboxapi.FilterInt64s(params.SetSearchWeight),
		"search_weight__n":          netboxapi.FilterInt64s(params.SetSearchWeightn),
		"search_weight__lt":         netboxapi.FilterInt64s(params.SetSearchWeightLt),
		"search_weight__lte":        netboxapi.FilterInt64s(params.SetSearchWeightLte),
		"search_weight__gt":         netboxapi.FilterInt64s(params.SetSearchWeightGt),
		"search_weight__gte":        netboxapi.FilterInt64s(params.SetSearchWeightGte),
		"search_weight__empty":      netboxapi.FilterBool(params.SetSearchWeightEmpty),
		"type":                      netboxapi.FilterStrings(params.SetType),
		"type__n":                   netboxapi.FilterStrings(params.SetTypen),
		"type__ic":                  netboxapi.FilterStrings(params.SetTypeIc),
		"type__nic":                 netboxapi.FilterStrings(params.SetTypeNic),
		"type__ie":                  netboxapi.FilterStrings(params.SetTypeIe),
		"type__nie":                 netboxapi.FilterStrings(params.SetTypeNie),
		"type__isw":                 netboxapi.FilterStrings(params.SetTypeIsw),
		"type__nisw":                netboxapi.FilterStrings(params.SetTypeNisw),
		"type__iew":                 netboxapi.FilterStrings(params.SetTypeIew),
		"type__niew":                netboxapi.FilterStrings(params.SetTypeNiew),
		"type__empty":               netboxapi.FilterBool(params.SetTypeEmpty),
		"type__regex":               netboxapi.FilterStrings(params.SetTypeRegex),
		"type__iregex":              netboxapi.FilterStrings(params.SetTypeIregex),
		"ui_editable":               netboxapi.FilterString(params.SetUIEditable),
		"ui_editable__n":            netboxapi.FilterString(params.SetUIEditablen),
		"ui_editable__ic":           netboxapi.FilterStrings(params.SetUIEditableIc),
		"ui_editable__nic":          netboxapi.FilterStrings(params.SetUIEditableNic),
		"ui_editable__ie":           netboxapi.FilterStrings(params.SetUIEditableIe),
		"ui_editable__nie":          netboxapi.FilterStrings(params.SetUIEditableNie),
		"ui_editable__isw":          netboxapi.FilterStrings(params.SetUIEditableIsw),
		"ui_editable__nisw":         netboxapi.FilterStrings(params.SetUIEditableNisw),
		"ui_editable__iew":          netboxapi.FilterStrings(params.SetUIEditableIew),
		"ui_editable__niew":         netboxapi.FilterStrings(params.SetUIEditableNiew),
		"ui_editable__empty":        netboxapi.FilterBool(params.SetUIEditableEmpty),
		"ui_editable__regex":        netboxapi.FilterStrings(params.SetUIEditableRegex),
		"ui_editable__iregex":       netboxapi.FilterStrings(params.SetUIEditableIregex),
		"ui_visible":                netboxapi.FilterString(params.SetUIVisible),
		"ui_visible__n":             netboxapi.FilterString(params.SetUIVisiblen),
		"ui_visible__ic":            netboxapi.FilterStrings(params.SetUIVisibleIc),
		"ui_visible__nic":           netboxapi.FilterStrings(params.SetUIVisibleNic),
		"ui_visible__ie":            netboxapi.FilterStrings(params.SetUIVisibleIe),
		"ui_visible__nie":           netboxapi.FilterStrings(params.SetUIVisibleNie),
		"ui_visible__isw":           netboxapi.FilterStrings(params.SetUIVisibleIsw),
		"ui_visible__nisw":          netboxapi.FilterStrings(params.SetUIVisibleNisw),
		"ui_visible__iew":           netboxapi.FilterStrings(params.SetUIVisibleIew),
		"ui_visible__niew":          netboxapi.FilterStrings(params.SetUIVisibleNiew),
		"ui_visible__empty":         netboxapi.FilterBool(params.SetUIVisibleEmpty),
		"ui_visible__regex":         netboxapi.FilterStrings(params.SetUIVisibleRegex),
		"ui_visible__iregex":        netboxapi.FilterStrings(params.SetUIVisibleIregex),
		"unique":                    netboxapi.FilterBool(params.SetUnique),
		"validation_maximum":        netboxapi.FilterFloat64s(params.SetValidationMaximum),
		"validation_maximum__n":     netboxapi.FilterFloat64s(params.SetValidationMaximumn),
		"validation_maximum__lt":    netboxapi.FilterFloat64s(params.SetValidationMaximumLt),
		"validation_maximum__lte":   netboxapi.FilterFloat64s(params.SetValidationMaximumLte),
		"validation_maximum__gt":    netboxapi.FilterFloat64s(params.SetValidationMaximumGt),
		"validation_maximum__gte":   netboxapi.FilterFloat64s(params.SetValidationMaximumGte),
		"validation_maximum__empty": netboxapi.FilterBool(params.SetValidationMaximumEmpty),
		"validation_minimum":        netboxapi.FilterFloat64s(params.SetValidationMinimum),
		"validation_minimum__n":     netboxapi.FilterFloat64s(params.SetValidationMinimumn),
		"validation_minimum__lt":    netboxapi.FilterFloat64s(params.SetValidationMinimumLt),
		"validation_minimum__lte":   netboxapi.FilterFloat64s(params.SetValidationMinimumLte),
		"validation_minimum__gt":    netboxapi.FilterFloat64s(params.SetValidationMinimumGt),
		"validation_minimum__gte":   netboxapi.FilterFloat64s(params.SetValidationMinimumGte),
		"validation_minimum__empty": netboxapi.FilterBool(params.SetValidationMinimumEmpty),
		"validation_regex":          netboxapi.FilterStrings(params.SetValidationRegex),
		"validation_regex__n":       netboxapi.FilterStrings(params.SetValidationRegexn),
		"validation_regex__ic":      netboxapi.FilterStrings(params.SetValidationRegexIc),
		"validation_regex__nic":     netboxapi.FilterStrings(params.SetValidationRegexNic),
		"validation_regex__ie":      netboxapi.FilterStrings(params.SetValidationRegexIe),
		"validation_regex__nie":     netboxapi.FilterStrings(params.SetValidationRegexNie),
		"validation_regex__isw":     netboxapi.FilterStrings(params.SetValidationRegexIsw),
		"validation_regex__nisw":    netboxapi.FilterStrings(params.SetValidationRegexNisw),
		"validation_regex__iew":     netboxapi.FilterStrings(params.SetValidationRegexIew),
		"validation_regex__niew":    netboxapi.FilterStrings(params.SetValidationRegexNiew),
		"validation_regex__empty":   netboxapi.FilterBool(params.SetValidationRegexEmpty),
		"validation_regex__regex":   netboxapi.FilterStrings(params.SetValidationRegexRegex),
		"validation_regex__iregex":  netboxapi.FilterStrings(params.SetValidationRegexIregex),
		"weight":                    netboxapi.FilterInt64s(params.SetWeight),
		"weight__n":                 netboxapi.FilterInt64s(params.SetWeightn),
		"weight__lt":                netboxapi.FilterInt64s(params.SetWeightLt),
		"weight__lte":               netboxapi.FilterInt64s(params.SetWeightLte),
		"weight__gt":                netboxapi.FilterInt64s(params.SetWeightGt),
		"weight__gte":               netboxapi.FilterInt64s(params.SetWeightGte),
		"weight__empty":             netboxapi.FilterBool(params.SetWeightEmpty),
	}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		listFilter, ok := listFilters[name]
		if !ok {
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
		if problem := listFilter.Add(name, value); problem != nil {
			resp.Diagnostics.AddError(problem.Summary, problem.Detail)
			return
		}
	}
	for _, listFilter := range listFilters {
		listFilter.Apply()
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
