// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

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
	_ datasource.DataSource              = (*eventRulesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*eventRulesDataSource)(nil)
)

// NewEventRulesDataSource returns a new event_rules data source, which
// lists event_rule objects matching its filters.
func NewEventRulesDataSource() datasource.DataSource {
	return &eventRulesDataSource{}
}

type eventRulesDataSource struct {
	client *netboxapi.Client
}

// eventRulesDataSourceModel is the data source model: the filters, the limit and the matching
// event_rules.
type eventRulesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"event_rules"`
}

// eventRuleResourceAttrTypes is the attribute type map of eventRuleResourceModel.
func eventRuleResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                 types.Int64Type,
		"name":               types.StringType,
		"object_types":       types.SetType{ElemType: types.StringType},
		"event_types":        types.SetType{ElemType: types.StringType},
		"enabled":            types.BoolType,
		"conditions":         jsontypes.NormalizedType{},
		"action_type":        types.StringType,
		"action_object_type": types.StringType,
		"action_object_id":   types.Int64Type,
		"description":        types.StringType,
		"owner_id":           types.Int64Type,
		"created":            types.StringType,
		"last_updated":       types.StringType,
		"url":                types.StringType,
		"tags":               types.SetType{ElemType: types.StringType},
		"tags_all":           types.SetType{ElemType: types.StringType},
		"custom_fields":      types.MapType{ElemType: types.StringType},
	}
}

func (d *eventRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_event_rules"
}

func (d *eventRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists event_rule objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: action_object_id, action_object_id__empty, action_object_id__gt, action_object_id__gte, action_object_id__lt, action_object_id__lte, action_object_id__n, action_object_type, action_object_type__n, action_type, action_type__empty, action_type__ic, action_type__ie, action_type__iew, action_type__iregex, action_type__isw, action_type__n, action_type__nic, action_type__nie, action_type__niew, action_type__nisw, action_type__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, enabled, event_type, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, object_type, object_type__ic, object_type__ie, object_type__iew, object_type__iregex, object_type__isw, object_type__n, object_type__nic, object_type__nie, object_type__niew, object_type__nisw, object_type__regex, object_type_id, object_type_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"event_rules": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching event_rules.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"object_types": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Object types the rule watches, as app-labelled content types (dcim.device).",
						},
						"event_types": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Events that trigger the rule: object_created, object_updated, object_deleted, job_started, job_completed, job_failed or job_errored.",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the rule fires.",
						},
						"conditions": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Conditions the changed object must satisfy, as JSON text (use jsonencode()).",
						},
						"action_type": schema.StringAttribute{
							Computed:    true,
							Description: "What the rule runs. One of: webhook, script, notification.",
						},
						"action_object_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the action object: extras.webhook or extras.script.",
						},
						"action_object_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the webhook or script to run.",
						},
						"description": schema.StringAttribute{
							Computed: true,
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
						"tags": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Slugs of the tags assigned to the object (the provider's default_tags are added on top, see tags_all).",
						},
						"tags_all": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Slugs of all tags on the object, including the provider's default_tags.",
						},
						"custom_fields": schema.MapAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *eventRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *eventRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data eventRulesDataSourceModel

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
	var responseDTOs []*netboxapi.EventRuleResponseDTO

	params := extras.NewExtrasEventRulesListParams()
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
	var vEnabled *bool
	var vActionType []string
	var vActionTypen []string
	var vActionTypeIc []string
	var vActionTypeNic []string
	var vActionTypeIe []string
	var vActionTypeNie []string
	var vActionTypeIsw []string
	var vActionTypeNisw []string
	var vActionTypeIew []string
	var vActionTypeNiew []string
	var vActionTypeEmpty *bool
	var vActionTypeRegex []string
	var vActionTypeIregex []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vActionObjectID []int64
	var vActionObjectIDn []int64
	var vActionObjectIDLt []int64
	var vActionObjectIDLte []int64
	var vActionObjectIDGt []int64
	var vActionObjectIDGte []int64
	var vActionObjectIDEmpty []int64
	var vActionObjectType []string
	var vActionObjectTypen []string
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
	var vEventType []string
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
	var vQ *string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	customFieldQuery := url.Values{}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		// Custom fields are queried as cf_<field name> with the field's own filter logic; NetBox ignores
		// names it has no custom field for.
		if strings.HasPrefix(name, "cf_") {
			customFieldQuery.Add(name, value)
			continue
		}
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
		case "enabled":
			if vEnabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'enabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'enabled' takes a boolean, got %q.", value))
				return
			}
			vEnabled = &v
		case "action_type":
			v := value
			vActionType = append(vActionType, v)
		case "action_type__n":
			v := value
			vActionTypen = append(vActionTypen, v)
		case "action_type__ic":
			v := value
			vActionTypeIc = append(vActionTypeIc, v)
		case "action_type__nic":
			v := value
			vActionTypeNic = append(vActionTypeNic, v)
		case "action_type__ie":
			v := value
			vActionTypeIe = append(vActionTypeIe, v)
		case "action_type__nie":
			v := value
			vActionTypeNie = append(vActionTypeNie, v)
		case "action_type__isw":
			v := value
			vActionTypeIsw = append(vActionTypeIsw, v)
		case "action_type__nisw":
			v := value
			vActionTypeNisw = append(vActionTypeNisw, v)
		case "action_type__iew":
			v := value
			vActionTypeIew = append(vActionTypeIew, v)
		case "action_type__niew":
			v := value
			vActionTypeNiew = append(vActionTypeNiew, v)
		case "action_type__empty":
			if vActionTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'action_type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_type__empty' takes a boolean, got %q.", value))
				return
			}
			vActionTypeEmpty = &v
		case "action_type__regex":
			v := value
			vActionTypeRegex = append(vActionTypeRegex, v)
		case "action_type__iregex":
			v := value
			vActionTypeIregex = append(vActionTypeIregex, v)
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
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
		case "action_object_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_object_id' takes an integer, got %q.", value))
				return
			}
			vActionObjectID = append(vActionObjectID, v)
		case "action_object_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_object_id__n' takes an integer, got %q.", value))
				return
			}
			vActionObjectIDn = append(vActionObjectIDn, v)
		case "action_object_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_object_id__lt' takes an integer, got %q.", value))
				return
			}
			vActionObjectIDLt = append(vActionObjectIDLt, v)
		case "action_object_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_object_id__lte' takes an integer, got %q.", value))
				return
			}
			vActionObjectIDLte = append(vActionObjectIDLte, v)
		case "action_object_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_object_id__gt' takes an integer, got %q.", value))
				return
			}
			vActionObjectIDGt = append(vActionObjectIDGt, v)
		case "action_object_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_object_id__gte' takes an integer, got %q.", value))
				return
			}
			vActionObjectIDGte = append(vActionObjectIDGte, v)
		case "action_object_id__empty":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'action_object_id__empty' takes an integer, got %q.", value))
				return
			}
			vActionObjectIDEmpty = append(vActionObjectIDEmpty, v)
		case "action_object_type":
			v := value
			vActionObjectType = append(vActionObjectType, v)
		case "action_object_type__n":
			v := value
			vActionObjectTypen = append(vActionObjectTypen, v)
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
		case "event_type":
			v := value
			vEventType = append(vEventType, v)
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
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "tag_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id' takes an integer, got %q.", value))
				return
			}
			vTagID = append(vTagID, v)
		case "tag_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__n' takes an integer, got %q.", value))
				return
			}
			vTagIDn = append(vTagIDn, v)
		case "tag_id__any":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__any' takes an integer, got %q.", value))
				return
			}
			vTagIDAny = append(vTagIDAny, v)
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
	if vEnabled != nil {
		params.SetEnabled(vEnabled)
	}
	if len(vActionType) > 0 {
		params.SetActionType(vActionType)
	}
	if len(vActionTypen) > 0 {
		params.SetActionTypen(vActionTypen)
	}
	if len(vActionTypeIc) > 0 {
		params.SetActionTypeIc(vActionTypeIc)
	}
	if len(vActionTypeNic) > 0 {
		params.SetActionTypeNic(vActionTypeNic)
	}
	if len(vActionTypeIe) > 0 {
		params.SetActionTypeIe(vActionTypeIe)
	}
	if len(vActionTypeNie) > 0 {
		params.SetActionTypeNie(vActionTypeNie)
	}
	if len(vActionTypeIsw) > 0 {
		params.SetActionTypeIsw(vActionTypeIsw)
	}
	if len(vActionTypeNisw) > 0 {
		params.SetActionTypeNisw(vActionTypeNisw)
	}
	if len(vActionTypeIew) > 0 {
		params.SetActionTypeIew(vActionTypeIew)
	}
	if len(vActionTypeNiew) > 0 {
		params.SetActionTypeNiew(vActionTypeNiew)
	}
	if vActionTypeEmpty != nil {
		params.SetActionTypeEmpty(vActionTypeEmpty)
	}
	if len(vActionTypeRegex) > 0 {
		params.SetActionTypeRegex(vActionTypeRegex)
	}
	if len(vActionTypeIregex) > 0 {
		params.SetActionTypeIregex(vActionTypeIregex)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
	}
	if len(vTag) > 0 {
		params.SetTag(vTag)
	}
	if len(vTagn) > 0 {
		params.SetTagn(vTagn)
	}
	if len(vTagAny) > 0 {
		params.SetTagAny(vTagAny)
	}
	if len(vActionObjectID) > 0 {
		params.SetActionObjectID(vActionObjectID)
	}
	if len(vActionObjectIDn) > 0 {
		params.SetActionObjectIDn(vActionObjectIDn)
	}
	if len(vActionObjectIDLt) > 0 {
		params.SetActionObjectIDLt(vActionObjectIDLt)
	}
	if len(vActionObjectIDLte) > 0 {
		params.SetActionObjectIDLte(vActionObjectIDLte)
	}
	if len(vActionObjectIDGt) > 0 {
		params.SetActionObjectIDGt(vActionObjectIDGt)
	}
	if len(vActionObjectIDGte) > 0 {
		params.SetActionObjectIDGte(vActionObjectIDGte)
	}
	if len(vActionObjectIDEmpty) > 0 {
		params.SetActionObjectIDEmpty(vActionObjectIDEmpty)
	}
	if len(vActionObjectType) > 0 {
		params.SetActionObjectType(vActionObjectType)
	}
	if len(vActionObjectTypen) > 0 {
		params.SetActionObjectTypen(vActionObjectTypen)
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
	if len(vEventType) > 0 {
		params.SetEventType(vEventType)
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
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vTagID) > 0 {
		params.SetTagID(vTagID)
	}
	if len(vTagIDn) > 0 {
		params.SetTagIDn(vTagIDn)
	}
	if len(vTagIDAny) > 0 {
		params.SetTagIDAny(vTagIDAny)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Extras.ExtrasEventRulesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_event_rules", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.EventRuleResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]eventRuleResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model eventRuleResourceModel
		resp.Diagnostics.Append(flattenEventRule(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: eventRuleResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
