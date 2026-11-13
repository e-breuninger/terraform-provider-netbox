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
	_ datasource.DataSource              = (*configContextsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*configContextsDataSource)(nil)
)

// NewConfigContextsDataSource returns a new config_contexts data source, which
// lists config_context objects matching its filters.
func NewConfigContextsDataSource() datasource.DataSource {
	return &configContextsDataSource{}
}

type configContextsDataSource struct {
	client *netboxapi.Client
}

// configContextsDataSourceModel is the data source model: the filters, the limit and the matching
// config_contexts.
type configContextsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"config_contexts"`
}

// configContextResourceAttrTypes is the attribute type map of configContextResourceModel.
func configContextResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                types.Int64Type,
		"name":              types.StringType,
		"weight":            types.Int64Type,
		"description":       types.StringType,
		"is_active":         types.BoolType,
		"data":              jsontypes.NormalizedType{},
		"profile_id":        types.Int64Type,
		"region_ids":        types.SetType{ElemType: types.Int64Type},
		"regions":           types.SetType{ElemType: types.Int64Type},
		"site_group_ids":    types.SetType{ElemType: types.Int64Type},
		"site_groups":       types.SetType{ElemType: types.Int64Type},
		"site_ids":          types.SetType{ElemType: types.Int64Type},
		"sites":             types.SetType{ElemType: types.Int64Type},
		"location_ids":      types.SetType{ElemType: types.Int64Type},
		"locations":         types.SetType{ElemType: types.Int64Type},
		"device_type_ids":   types.SetType{ElemType: types.Int64Type},
		"device_types":      types.SetType{ElemType: types.Int64Type},
		"device_role_ids":   types.SetType{ElemType: types.Int64Type},
		"roles":             types.SetType{ElemType: types.Int64Type},
		"platform_ids":      types.SetType{ElemType: types.Int64Type},
		"platforms":         types.SetType{ElemType: types.Int64Type},
		"cluster_type_ids":  types.SetType{ElemType: types.Int64Type},
		"cluster_types":     types.SetType{ElemType: types.Int64Type},
		"cluster_group_ids": types.SetType{ElemType: types.Int64Type},
		"cluster_groups":    types.SetType{ElemType: types.Int64Type},
		"cluster_ids":       types.SetType{ElemType: types.Int64Type},
		"clusters":          types.SetType{ElemType: types.Int64Type},
		"tenant_group_ids":  types.SetType{ElemType: types.Int64Type},
		"tenant_groups":     types.SetType{ElemType: types.Int64Type},
		"tenant_ids":        types.SetType{ElemType: types.Int64Type},
		"tenants":           types.SetType{ElemType: types.Int64Type},
		"tag_slugs":         types.SetType{ElemType: types.StringType},
		"owner_id":          types.Int64Type,
		"created":           types.StringType,
		"last_updated":      types.StringType,
		"url":               types.StringType,
	}
}

func (d *configContextsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_contexts"
}

func (d *configContextsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists config_context objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: auto_sync_enabled, cluster_group, cluster_group__n, cluster_group_id, cluster_group_id__n, cluster_id, cluster_id__n, cluster_type, cluster_type__n, cluster_type_id, cluster_type_id__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, data_file_id, data_file_id__n, data_source_id, data_source_id__n, data_synced, data_synced__empty, data_synced__gt, data_synced__gte, data_synced__lt, data_synced__lte, data_synced__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device_role, device_role__n, device_role_id, device_role_id__n, device_type_id, device_type_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, is_active, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, platform, platform__n, platform_id, platform_id__n, profile, profile__n, profile_id, profile_id__n, q, region, region__n, region_id, region_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, tag, tag__n, tag_id, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n. Repeating a name sends that parameter once per value.",
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
			"config_contexts": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching config_contexts.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"weight": schema.Int64Attribute{
							Computed:    true,
							Description: "Merge order; higher weights are applied later and win. Must be a whole number.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"is_active": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the context is applied.",
						},
						"data": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "The context data as JSON text (use jsonencode()).",
						},
						"profile_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the config context profile whose schema the data must satisfy.",
						},
						"region_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the regions the context applies to; empty means all.",
						},
						"regions": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use region_ids instead.",
							Description:        "Deprecated alias of region_ids.",
						},
						"site_group_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the site groups the context applies to; empty means all.",
						},
						"site_groups": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use site_group_ids instead.",
							Description:        "Deprecated alias of site_group_ids.",
						},
						"site_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the sites the context applies to; empty means all.",
						},
						"sites": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use site_ids instead.",
							Description:        "Deprecated alias of site_ids.",
						},
						"location_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the locations the context applies to; empty means all.",
						},
						"locations": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use location_ids instead.",
							Description:        "Deprecated alias of location_ids.",
						},
						"device_type_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the device types the context applies to; empty means all.",
						},
						"device_types": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use device_type_ids instead.",
							Description:        "Deprecated alias of device_type_ids.",
						},
						"device_role_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the device roles the context applies to; empty means all.",
						},
						"roles": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use device_role_ids instead.",
							Description:        "Deprecated alias of device_role_ids.",
						},
						"platform_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the platforms the context applies to; empty means all.",
						},
						"platforms": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use platform_ids instead.",
							Description:        "Deprecated alias of platform_ids.",
						},
						"cluster_type_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the cluster types the context applies to; empty means all.",
						},
						"cluster_types": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use cluster_type_ids instead.",
							Description:        "Deprecated alias of cluster_type_ids.",
						},
						"cluster_group_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the cluster groups the context applies to; empty means all.",
						},
						"cluster_groups": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use cluster_group_ids instead.",
							Description:        "Deprecated alias of cluster_group_ids.",
						},
						"cluster_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the clusters the context applies to; empty means all.",
						},
						"clusters": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use cluster_ids instead.",
							Description:        "Deprecated alias of cluster_ids.",
						},
						"tenant_group_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the tenant groups the context applies to; empty means all.",
						},
						"tenant_groups": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use tenant_group_ids instead.",
							Description:        "Deprecated alias of tenant_group_ids.",
						},
						"tenant_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the tenants the context applies to; empty means all.",
						},
						"tenants": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use tenant_ids instead.",
							Description:        "Deprecated alias of tenant_ids.",
						},
						"tag_slugs": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Slugs of the tags an object must carry for the context to apply; empty means any. This is a scope, not the context's own tags.",
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

func (d *configContextsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *configContextsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data configContextsDataSourceModel

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
	var responseDTOs []*netboxapi.ConfigContextResponseDTO

	params := extras.NewExtrasConfigContextsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"name":                netboxapi.FilterStrings(params.SetName),
		"name__n":             netboxapi.FilterStrings(params.SetNamen),
		"name__ic":            netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":           netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":            netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":           netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":           netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":          netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":           netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":          netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":         netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":         netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":        netboxapi.FilterStrings(params.SetNameIregex),
		"is_active":           netboxapi.FilterBool(params.SetIsActive),
		"owner_id":            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":         netboxapi.FilterInt64s(params.SetOwnerIDn),
		"auto_sync_enabled":   netboxapi.FilterBool(params.SetAutoSyncEnabled),
		"cluster_group":       netboxapi.FilterStrings(params.SetClusterGroup),
		"cluster_group__n":    netboxapi.FilterStrings(params.SetClusterGroupn),
		"cluster_group_id":    netboxapi.FilterInt64s(params.SetClusterGroupID),
		"cluster_group_id__n": netboxapi.FilterInt64s(params.SetClusterGroupIDn),
		"cluster_id":          netboxapi.FilterInt64s(params.SetClusterID),
		"cluster_id__n":       netboxapi.FilterInt64s(params.SetClusterIDn),
		"cluster_type":        netboxapi.FilterStrings(params.SetClusterType),
		"cluster_type__n":     netboxapi.FilterStrings(params.SetClusterTypen),
		"cluster_type_id":     netboxapi.FilterInt64s(params.SetClusterTypeID),
		"cluster_type_id__n":  netboxapi.FilterInt64s(params.SetClusterTypeIDn),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"data_file_id":        netboxapi.FilterInt64s(params.SetDataFileID),
		"data_file_id__n":     netboxapi.FilterInt64s(params.SetDataFileIDn),
		"data_source_id":      netboxapi.FilterInt64s(params.SetDataSourceID),
		"data_source_id__n":   netboxapi.FilterInt64s(params.SetDataSourceIDn),
		"data_synced":         netboxapi.FilterStrings(params.SetDataSynced),
		"data_synced__n":      netboxapi.FilterStrings(params.SetDataSyncedn),
		"data_synced__lt":     netboxapi.FilterStrings(params.SetDataSyncedLt),
		"data_synced__lte":    netboxapi.FilterStrings(params.SetDataSyncedLte),
		"data_synced__gt":     netboxapi.FilterStrings(params.SetDataSyncedGt),
		"data_synced__gte":    netboxapi.FilterStrings(params.SetDataSyncedGte),
		"data_synced__empty":  netboxapi.FilterBool(params.SetDataSyncedEmpty),
		"description":         netboxapi.FilterStrings(params.SetDescription),
		"description__n":      netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":     netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":    netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":     netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":    netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":    netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":   netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":    netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":   netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":  netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":  netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex": netboxapi.FilterStrings(params.SetDescriptionIregex),
		"device_role":         netboxapi.FilterStrings(params.SetDeviceRole),
		"device_role__n":      netboxapi.FilterStrings(params.SetDeviceRolen),
		"device_role_id":      netboxapi.FilterInt64s(params.SetDeviceRoleID),
		"device_role_id__n":   netboxapi.FilterInt64s(params.SetDeviceRoleIDn),
		"device_type_id":      netboxapi.FilterInt64s(params.SetDeviceTypeID),
		"device_type_id__n":   netboxapi.FilterInt64s(params.SetDeviceTypeIDn),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":            netboxapi.FilterStrings(params.SetLocation),
		"location__n":         netboxapi.FilterStrings(params.SetLocationn),
		"location_id":         netboxapi.FilterInt64s(params.SetLocationID),
		"location_id__n":      netboxapi.FilterInt64s(params.SetLocationIDn),
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"platform":            netboxapi.FilterStrings(params.SetPlatform),
		"platform__n":         netboxapi.FilterStrings(params.SetPlatformn),
		"platform_id":         netboxapi.FilterInt64s(params.SetPlatformID),
		"platform_id__n":      netboxapi.FilterInt64s(params.SetPlatformIDn),
		"profile":             netboxapi.FilterStrings(params.SetProfile),
		"profile__n":          netboxapi.FilterStrings(params.SetProfilen),
		"profile_id":          netboxapi.FilterInt64s(params.SetProfileID),
		"profile_id__n":       netboxapi.FilterInt64s(params.SetProfileIDn),
		"q":                   netboxapi.FilterString(params.SetQ),
		"region":              netboxapi.FilterStrings(params.SetRegion),
		"region__n":           netboxapi.FilterStrings(params.SetRegionn),
		"region_id":           netboxapi.FilterInt64s(params.SetRegionID),
		"region_id__n":        netboxapi.FilterInt64s(params.SetRegionIDn),
		"site":                netboxapi.FilterStrings(params.SetSite),
		"site__n":             netboxapi.FilterStrings(params.SetSiten),
		"site_group":          netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":       netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":       netboxapi.FilterInt64s(params.SetSiteGroupID),
		"site_group_id__n":    netboxapi.FilterInt64s(params.SetSiteGroupIDn),
		"site_id":             netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":          netboxapi.FilterInt64s(params.SetSiteIDn),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tenant":              netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":           netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":        netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":     netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":     netboxapi.FilterInt64s(params.SetTenantGroupID),
		"tenant_group_id__n":  netboxapi.FilterInt64s(params.SetTenantGroupIDn),
		"tenant_id":           netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":        netboxapi.FilterInt64s(params.SetTenantIDn),
		"weight":              netboxapi.FilterInt64s(params.SetWeight),
		"weight__n":           netboxapi.FilterInt64s(params.SetWeightn),
		"weight__lt":          netboxapi.FilterInt64s(params.SetWeightLt),
		"weight__lte":         netboxapi.FilterInt64s(params.SetWeightLte),
		"weight__gt":          netboxapi.FilterInt64s(params.SetWeightGt),
		"weight__gte":         netboxapi.FilterInt64s(params.SetWeightGte),
		"weight__empty":       netboxapi.FilterBool(params.SetWeightEmpty),
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
		res, err := d.client.Extras.ExtrasConfigContextsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_config_contexts", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ConfigContextResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]configContextResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model configContextResourceModel
		resp.Diagnostics.Append(flattenConfigContext(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: configContextResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
