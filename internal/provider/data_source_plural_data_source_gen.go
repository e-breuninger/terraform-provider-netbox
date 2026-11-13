// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/core"
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
	_ datasource.DataSource              = (*dataSourcesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*dataSourcesDataSource)(nil)
)

// NewDataSourcesDataSource returns a new data_sources data source, which
// lists data_source objects matching its filters.
func NewDataSourcesDataSource() datasource.DataSource {
	return &dataSourcesDataSource{}
}

type dataSourcesDataSource struct {
	client *netboxapi.Client
}

// dataSourcesDataSourceModel is the data source model: the filters, the limit and the matching
// data_sources.
type dataSourcesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"data_sources"`
}

// dataSourceResourceAttrTypes is the attribute type map of dataSourceResourceModel.
func dataSourceResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"name":          types.StringType,
		"type":          types.StringType,
		"source_url":    types.StringType,
		"enabled":       types.BoolType,
		"ignore_rules":  types.StringType,
		"parameters":    jsontypes.NormalizedType{},
		"sync_interval": types.Int64Type,
		"description":   types.StringType,
		"comments":      types.StringType,
		"status":        types.StringType,
		"last_synced":   types.StringType,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"file_count":    types.Int64Type,
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *dataSourcesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_data_sources"
}

func (d *dataSourcesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists data_source objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, enabled, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_synced, last_synced__empty, last_synced__gt, last_synced__gte, last_synced__lt, last_synced__lte, last_synced__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, source_url, source_url__empty, source_url__ic, source_url__ie, source_url__iew, source_url__iregex, source_url__isw, source_url__n, source_url__nic, source_url__nie, source_url__niew, source_url__nisw, source_url__regex, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, sync_interval, sync_interval__ic, sync_interval__ie, sync_interval__iew, sync_interval__iregex, sync_interval__isw, sync_interval__n, sync_interval__nic, sync_interval__nie, sync_interval__niew, sync_interval__nisw, sync_interval__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"data_sources": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching data_sources.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Kind of backing store. One of: local, git, amazon-s3.",
						},
						"source_url": schema.StringAttribute{
							Computed:    true,
							Description: "URL or path of the source.",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether NetBox synchronises from this source.",
						},
						"ignore_rules": schema.StringAttribute{
							Computed:    true,
							Description: "Newline-separated glob patterns of files to ignore when synchronising.",
						},
						"parameters": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Backend-specific parameters as JSON text (use jsonencode()).",
						},
						"sync_interval": schema.Int64Attribute{
							Computed:    true,
							Description: "Automatic synchronisation interval in minutes. One of: 1, 60, 720, 1440, 10080, 43200.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Synchronisation status reported by NetBox.",
						},
						"last_synced": schema.StringAttribute{
							Computed:    true,
							Description: "When NetBox last synchronised the source.",
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
						"file_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of files synced from the data source.",
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

func (d *dataSourcesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *dataSourcesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data dataSourcesDataSourceModel

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
	var responseDTOs []*netboxapi.DataSourceResponseDTO

	params := core.NewCoreDataSourcesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"name":                  netboxapi.FilterStrings(params.SetName),
		"name__n":               netboxapi.FilterStrings(params.SetNamen),
		"name__ic":              netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":             netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":              netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":             netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":             netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":            netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":             netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":            netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":           netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":           netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":          netboxapi.FilterStrings(params.SetNameIregex),
		"status":                netboxapi.FilterStrings(params.SetStatus),
		"status__n":             netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":            netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":           netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":            netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":           netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":           netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":          netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":           netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":          netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":         netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":         netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":        netboxapi.FilterStrings(params.SetStatusIregex),
		"enabled":               netboxapi.FilterBool(params.SetEnabled),
		"source_url":            netboxapi.FilterStrings(params.SetSourceURL),
		"source_url__n":         netboxapi.FilterStrings(params.SetSourceURLn),
		"source_url__ic":        netboxapi.FilterStrings(params.SetSourceURLIc),
		"source_url__nic":       netboxapi.FilterStrings(params.SetSourceURLNic),
		"source_url__ie":        netboxapi.FilterStrings(params.SetSourceURLIe),
		"source_url__nie":       netboxapi.FilterStrings(params.SetSourceURLNie),
		"source_url__isw":       netboxapi.FilterStrings(params.SetSourceURLIsw),
		"source_url__nisw":      netboxapi.FilterStrings(params.SetSourceURLNisw),
		"source_url__iew":       netboxapi.FilterStrings(params.SetSourceURLIew),
		"source_url__niew":      netboxapi.FilterStrings(params.SetSourceURLNiew),
		"source_url__empty":     netboxapi.FilterBool(params.SetSourceURLEmpty),
		"source_url__regex":     netboxapi.FilterStrings(params.SetSourceURLRegex),
		"source_url__iregex":    netboxapi.FilterStrings(params.SetSourceURLIregex),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":           netboxapi.FilterStrings(params.SetDescription),
		"description__n":        netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":       netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":      netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":       netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":      netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":      netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":     netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":      netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":     netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":    netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":    netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":   netboxapi.FilterStrings(params.SetDescriptionIregex),
		"last_synced":           netboxapi.FilterStrings(params.SetLastSynced),
		"last_synced__n":        netboxapi.FilterStrings(params.SetLastSyncedn),
		"last_synced__lt":       netboxapi.FilterStrings(params.SetLastSyncedLt),
		"last_synced__lte":      netboxapi.FilterStrings(params.SetLastSyncedLte),
		"last_synced__gt":       netboxapi.FilterStrings(params.SetLastSyncedGt),
		"last_synced__gte":      netboxapi.FilterStrings(params.SetLastSyncedGte),
		"last_synced__empty":    netboxapi.FilterBool(params.SetLastSyncedEmpty),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                 netboxapi.FilterStrings(params.SetOwner),
		"owner__n":              netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":           netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":        netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":        netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":     netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                     netboxapi.FilterString(params.SetQ),
		"sync_interval":         netboxapi.FilterInt64s(params.SetSyncInterval),
		"sync_interval__n":      netboxapi.FilterInt64s(params.SetSyncIntervaln),
		"sync_interval__ic":     netboxapi.FilterInt64s(params.SetSyncIntervalIc),
		"sync_interval__nic":    netboxapi.FilterInt64s(params.SetSyncIntervalNic),
		"sync_interval__ie":     netboxapi.FilterInt64s(params.SetSyncIntervalIe),
		"sync_interval__nie":    netboxapi.FilterInt64s(params.SetSyncIntervalNie),
		"sync_interval__isw":    netboxapi.FilterInt64s(params.SetSyncIntervalIsw),
		"sync_interval__nisw":   netboxapi.FilterInt64s(params.SetSyncIntervalNisw),
		"sync_interval__iew":    netboxapi.FilterInt64s(params.SetSyncIntervalIew),
		"sync_interval__niew":   netboxapi.FilterInt64s(params.SetSyncIntervalNiew),
		"sync_interval__regex":  netboxapi.FilterInt64s(params.SetSyncIntervalRegex),
		"sync_interval__iregex": netboxapi.FilterInt64s(params.SetSyncIntervalIregex),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"type":                  netboxapi.FilterStrings(params.SetType),
		"type__n":               netboxapi.FilterStrings(params.SetTypen),
		"type__ic":              netboxapi.FilterStrings(params.SetTypeIc),
		"type__nic":             netboxapi.FilterStrings(params.SetTypeNic),
		"type__ie":              netboxapi.FilterStrings(params.SetTypeIe),
		"type__nie":             netboxapi.FilterStrings(params.SetTypeNie),
		"type__isw":             netboxapi.FilterStrings(params.SetTypeIsw),
		"type__nisw":            netboxapi.FilterStrings(params.SetTypeNisw),
		"type__iew":             netboxapi.FilterStrings(params.SetTypeIew),
		"type__niew":            netboxapi.FilterStrings(params.SetTypeNiew),
		"type__empty":           netboxapi.FilterBool(params.SetTypeEmpty),
		"type__regex":           netboxapi.FilterStrings(params.SetTypeRegex),
		"type__iregex":          netboxapi.FilterStrings(params.SetTypeIregex),
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
		res, err := d.client.Core.CoreDataSourcesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_data_sources", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.DataSourceResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]dataSourceResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model dataSourceResourceModel
		resp.Diagnostics.Append(flattenDataSource(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: dataSourceResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
