// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                      netboxapi.FilterInt64s(params.SetID),
		"id__n":                   netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                  netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                 netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                  netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                 netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":               netboxapi.FilterBool(params.SetIDEmpty),
		"name":                    netboxapi.FilterStrings(params.SetName),
		"name__n":                 netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":               netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":               netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":               netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":              netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":               netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":              netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":             netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":             netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":            netboxapi.FilterStrings(params.SetNameIregex),
		"enabled":                 netboxapi.FilterBool(params.SetEnabled),
		"action_type":             netboxapi.FilterStrings(params.SetActionType),
		"action_type__n":          netboxapi.FilterStrings(params.SetActionTypen),
		"action_type__ic":         netboxapi.FilterStrings(params.SetActionTypeIc),
		"action_type__nic":        netboxapi.FilterStrings(params.SetActionTypeNic),
		"action_type__ie":         netboxapi.FilterStrings(params.SetActionTypeIe),
		"action_type__nie":        netboxapi.FilterStrings(params.SetActionTypeNie),
		"action_type__isw":        netboxapi.FilterStrings(params.SetActionTypeIsw),
		"action_type__nisw":       netboxapi.FilterStrings(params.SetActionTypeNisw),
		"action_type__iew":        netboxapi.FilterStrings(params.SetActionTypeIew),
		"action_type__niew":       netboxapi.FilterStrings(params.SetActionTypeNiew),
		"action_type__empty":      netboxapi.FilterBool(params.SetActionTypeEmpty),
		"action_type__regex":      netboxapi.FilterStrings(params.SetActionTypeRegex),
		"action_type__iregex":     netboxapi.FilterStrings(params.SetActionTypeIregex),
		"owner_id":                netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":             netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                     netboxapi.FilterStrings(params.SetTag),
		"tag__n":                  netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                netboxapi.FilterStrings(params.SetTagAny),
		"action_object_id":        netboxapi.FilterInt64s(params.SetActionObjectID),
		"action_object_id__n":     netboxapi.FilterInt64s(params.SetActionObjectIDn),
		"action_object_id__lt":    netboxapi.FilterInt64s(params.SetActionObjectIDLt),
		"action_object_id__lte":   netboxapi.FilterInt64s(params.SetActionObjectIDLte),
		"action_object_id__gt":    netboxapi.FilterInt64s(params.SetActionObjectIDGt),
		"action_object_id__gte":   netboxapi.FilterInt64s(params.SetActionObjectIDGte),
		"action_object_id__empty": netboxapi.FilterInt64s(params.SetActionObjectIDEmpty),
		"action_object_type":      netboxapi.FilterStrings(params.SetActionObjectType),
		"action_object_type__n":   netboxapi.FilterStrings(params.SetActionObjectTypen),
		"created":                 netboxapi.FilterStrings(params.SetCreated),
		"created__n":              netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":             netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":            netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":             netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":            netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":          netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":             netboxapi.FilterStrings(params.SetDescription),
		"description__n":          netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":         netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":        netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":         netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":        netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":        netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":       netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":        netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":       netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":      netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":      netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":     netboxapi.FilterStrings(params.SetDescriptionIregex),
		"event_type":              netboxapi.FilterStrings(params.SetEventType),
		"last_updated":            netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":         netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":        netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":       netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":        netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":       netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":     netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"object_type":             netboxapi.FilterStrings(params.SetObjectType),
		"object_type__n":          netboxapi.FilterStrings(params.SetObjectTypen),
		"object_type__ic":         netboxapi.FilterStrings(params.SetObjectTypeIc),
		"object_type__nic":        netboxapi.FilterStrings(params.SetObjectTypeNic),
		"object_type__ie":         netboxapi.FilterStrings(params.SetObjectTypeIe),
		"object_type__nie":        netboxapi.FilterStrings(params.SetObjectTypeNie),
		"object_type__isw":        netboxapi.FilterStrings(params.SetObjectTypeIsw),
		"object_type__nisw":       netboxapi.FilterStrings(params.SetObjectTypeNisw),
		"object_type__iew":        netboxapi.FilterStrings(params.SetObjectTypeIew),
		"object_type__niew":       netboxapi.FilterStrings(params.SetObjectTypeNiew),
		"object_type__regex":      netboxapi.FilterStrings(params.SetObjectTypeRegex),
		"object_type__iregex":     netboxapi.FilterStrings(params.SetObjectTypeIregex),
		"object_type_id":          netboxapi.FilterInt64s(params.SetObjectTypeID),
		"object_type_id__n":       netboxapi.FilterInt64s(params.SetObjectTypeIDn),
		"owner":                   netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":             netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":          netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":          netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":       netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                       netboxapi.FilterString(params.SetQ),
		"tag_id":                  netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":               netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":             netboxapi.FilterInt64s(params.SetTagIDAny),
	}
	customFieldQuery := url.Values{}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		// Custom fields are queried as cf_<field name> with the field's own filter logic; NetBox ignores
		// names it has no custom field for.
		if strings.HasPrefix(name, "cf_") {
			customFieldQuery.Add(name, value)
			continue
		}
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
