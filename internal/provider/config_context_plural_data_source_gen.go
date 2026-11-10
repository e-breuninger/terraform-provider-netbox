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
	var vIsActive *bool
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vAutoSyncEnabled *bool
	var vClusterGroup []string
	var vClusterGroupn []string
	var vClusterGroupID []int64
	var vClusterGroupIDn []int64
	var vClusterID []int64
	var vClusterIDn []int64
	var vClusterType []string
	var vClusterTypen []string
	var vClusterTypeID []int64
	var vClusterTypeIDn []int64
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vDataFileID []int64
	var vDataFileIDn []int64
	var vDataSourceID []int64
	var vDataSourceIDn []int64
	var vDataSynced []string
	var vDataSyncedn []string
	var vDataSyncedLt []string
	var vDataSyncedLte []string
	var vDataSyncedGt []string
	var vDataSyncedGte []string
	var vDataSyncedEmpty *bool
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
	var vDeviceRole []string
	var vDeviceRolen []string
	var vDeviceRoleID []int64
	var vDeviceRoleIDn []int64
	var vDeviceTypeID []int64
	var vDeviceTypeIDn []int64
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLocation []string
	var vLocationn []string
	var vLocationID []int64
	var vLocationIDn []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPlatform []string
	var vPlatformn []string
	var vPlatformID []int64
	var vPlatformIDn []int64
	var vProfile []string
	var vProfilen []string
	var vProfileID []int64
	var vProfileIDn []int64
	var vQ *string
	var vRegion []string
	var vRegionn []string
	var vRegionID []int64
	var vRegionIDn []int64
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []int64
	var vSiteGroupIDn []int64
	var vSiteID []int64
	var vSiteIDn []int64
	var vTag []string
	var vTagn []string
	var vTagID []int64
	var vTagIDn []int64
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []int64
	var vTenantGroupIDn []int64
	var vTenantID []int64
	var vTenantIDn []int64
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
		case "is_active":
			if vIsActive != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'is_active' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'is_active' takes a boolean, got %q.", value))
				return
			}
			vIsActive = &v
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
		case "auto_sync_enabled":
			if vAutoSyncEnabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'auto_sync_enabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'auto_sync_enabled' takes a boolean, got %q.", value))
				return
			}
			vAutoSyncEnabled = &v
		case "cluster_group":
			v := value
			vClusterGroup = append(vClusterGroup, v)
		case "cluster_group__n":
			v := value
			vClusterGroupn = append(vClusterGroupn, v)
		case "cluster_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_group_id' takes an integer, got %q.", value))
				return
			}
			vClusterGroupID = append(vClusterGroupID, v)
		case "cluster_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_group_id__n' takes an integer, got %q.", value))
				return
			}
			vClusterGroupIDn = append(vClusterGroupIDn, v)
		case "cluster_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_id' takes an integer, got %q.", value))
				return
			}
			vClusterID = append(vClusterID, v)
		case "cluster_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_id__n' takes an integer, got %q.", value))
				return
			}
			vClusterIDn = append(vClusterIDn, v)
		case "cluster_type":
			v := value
			vClusterType = append(vClusterType, v)
		case "cluster_type__n":
			v := value
			vClusterTypen = append(vClusterTypen, v)
		case "cluster_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_type_id' takes an integer, got %q.", value))
				return
			}
			vClusterTypeID = append(vClusterTypeID, v)
		case "cluster_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_type_id__n' takes an integer, got %q.", value))
				return
			}
			vClusterTypeIDn = append(vClusterTypeIDn, v)
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
		case "data_file_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'data_file_id' takes an integer, got %q.", value))
				return
			}
			vDataFileID = append(vDataFileID, v)
		case "data_file_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'data_file_id__n' takes an integer, got %q.", value))
				return
			}
			vDataFileIDn = append(vDataFileIDn, v)
		case "data_source_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'data_source_id' takes an integer, got %q.", value))
				return
			}
			vDataSourceID = append(vDataSourceID, v)
		case "data_source_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'data_source_id__n' takes an integer, got %q.", value))
				return
			}
			vDataSourceIDn = append(vDataSourceIDn, v)
		case "data_synced":
			v := value
			vDataSynced = append(vDataSynced, v)
		case "data_synced__n":
			v := value
			vDataSyncedn = append(vDataSyncedn, v)
		case "data_synced__lt":
			v := value
			vDataSyncedLt = append(vDataSyncedLt, v)
		case "data_synced__lte":
			v := value
			vDataSyncedLte = append(vDataSyncedLte, v)
		case "data_synced__gt":
			v := value
			vDataSyncedGt = append(vDataSyncedGt, v)
		case "data_synced__gte":
			v := value
			vDataSyncedGte = append(vDataSyncedGte, v)
		case "data_synced__empty":
			if vDataSyncedEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'data_synced__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'data_synced__empty' takes a boolean, got %q.", value))
				return
			}
			vDataSyncedEmpty = &v
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
		case "device_role":
			v := value
			vDeviceRole = append(vDeviceRole, v)
		case "device_role__n":
			v := value
			vDeviceRolen = append(vDeviceRolen, v)
		case "device_role_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_role_id' takes an integer, got %q.", value))
				return
			}
			vDeviceRoleID = append(vDeviceRoleID, v)
		case "device_role_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_role_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceRoleIDn = append(vDeviceRoleIDn, v)
		case "device_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_type_id' takes an integer, got %q.", value))
				return
			}
			vDeviceTypeID = append(vDeviceTypeID, v)
		case "device_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_type_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceTypeIDn = append(vDeviceTypeIDn, v)
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
		case "location":
			v := value
			vLocation = append(vLocation, v)
		case "location__n":
			v := value
			vLocationn = append(vLocationn, v)
		case "location_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id' takes an integer, got %q.", value))
				return
			}
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id__n' takes an integer, got %q.", value))
				return
			}
			vLocationIDn = append(vLocationIDn, v)
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
		case "platform":
			v := value
			vPlatform = append(vPlatform, v)
		case "platform__n":
			v := value
			vPlatformn = append(vPlatformn, v)
		case "platform_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'platform_id' takes an integer, got %q.", value))
				return
			}
			vPlatformID = append(vPlatformID, v)
		case "platform_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'platform_id__n' takes an integer, got %q.", value))
				return
			}
			vPlatformIDn = append(vPlatformIDn, v)
		case "profile":
			v := value
			vProfile = append(vProfile, v)
		case "profile__n":
			v := value
			vProfilen = append(vProfilen, v)
		case "profile_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'profile_id' takes an integer, got %q.", value))
				return
			}
			vProfileID = append(vProfileID, v)
		case "profile_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'profile_id__n' takes an integer, got %q.", value))
				return
			}
			vProfileIDn = append(vProfileIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "region":
			v := value
			vRegion = append(vRegion, v)
		case "region__n":
			v := value
			vRegionn = append(vRegionn, v)
		case "region_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'region_id' takes an integer, got %q.", value))
				return
			}
			vRegionID = append(vRegionID, v)
		case "region_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'region_id__n' takes an integer, got %q.", value))
				return
			}
			vRegionIDn = append(vRegionIDn, v)
		case "site":
			v := value
			vSite = append(vSite, v)
		case "site__n":
			v := value
			vSiten = append(vSiten, v)
		case "site_group":
			v := value
			vSiteGroup = append(vSiteGroup, v)
		case "site_group__n":
			v := value
			vSiteGroupn = append(vSiteGroupn, v)
		case "site_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_group_id' takes an integer, got %q.", value))
				return
			}
			vSiteGroupID = append(vSiteGroupID, v)
		case "site_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_group_id__n' takes an integer, got %q.", value))
				return
			}
			vSiteGroupIDn = append(vSiteGroupIDn, v)
		case "site_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id' takes an integer, got %q.", value))
				return
			}
			vSiteID = append(vSiteID, v)
		case "site_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id__n' takes an integer, got %q.", value))
				return
			}
			vSiteIDn = append(vSiteIDn, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
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
		case "tenant":
			v := value
			vTenant = append(vTenant, v)
		case "tenant__n":
			v := value
			vTenantn = append(vTenantn, v)
		case "tenant_group":
			v := value
			vTenantGroup = append(vTenantGroup, v)
		case "tenant_group__n":
			v := value
			vTenantGroupn = append(vTenantGroupn, v)
		case "tenant_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_group_id' takes an integer, got %q.", value))
				return
			}
			vTenantGroupID = append(vTenantGroupID, v)
		case "tenant_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_group_id__n' takes an integer, got %q.", value))
				return
			}
			vTenantGroupIDn = append(vTenantGroupIDn, v)
		case "tenant_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id' takes an integer, got %q.", value))
				return
			}
			vTenantID = append(vTenantID, v)
		case "tenant_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id__n' takes an integer, got %q.", value))
				return
			}
			vTenantIDn = append(vTenantIDn, v)
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
	if vIsActive != nil {
		params.SetIsActive(vIsActive)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
	}
	if vAutoSyncEnabled != nil {
		params.SetAutoSyncEnabled(vAutoSyncEnabled)
	}
	if len(vClusterGroup) > 0 {
		params.SetClusterGroup(vClusterGroup)
	}
	if len(vClusterGroupn) > 0 {
		params.SetClusterGroupn(vClusterGroupn)
	}
	if len(vClusterGroupID) > 0 {
		params.SetClusterGroupID(vClusterGroupID)
	}
	if len(vClusterGroupIDn) > 0 {
		params.SetClusterGroupIDn(vClusterGroupIDn)
	}
	if len(vClusterID) > 0 {
		params.SetClusterID(vClusterID)
	}
	if len(vClusterIDn) > 0 {
		params.SetClusterIDn(vClusterIDn)
	}
	if len(vClusterType) > 0 {
		params.SetClusterType(vClusterType)
	}
	if len(vClusterTypen) > 0 {
		params.SetClusterTypen(vClusterTypen)
	}
	if len(vClusterTypeID) > 0 {
		params.SetClusterTypeID(vClusterTypeID)
	}
	if len(vClusterTypeIDn) > 0 {
		params.SetClusterTypeIDn(vClusterTypeIDn)
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
	if len(vDataFileID) > 0 {
		params.SetDataFileID(vDataFileID)
	}
	if len(vDataFileIDn) > 0 {
		params.SetDataFileIDn(vDataFileIDn)
	}
	if len(vDataSourceID) > 0 {
		params.SetDataSourceID(vDataSourceID)
	}
	if len(vDataSourceIDn) > 0 {
		params.SetDataSourceIDn(vDataSourceIDn)
	}
	if len(vDataSynced) > 0 {
		params.SetDataSynced(vDataSynced)
	}
	if len(vDataSyncedn) > 0 {
		params.SetDataSyncedn(vDataSyncedn)
	}
	if len(vDataSyncedLt) > 0 {
		params.SetDataSyncedLt(vDataSyncedLt)
	}
	if len(vDataSyncedLte) > 0 {
		params.SetDataSyncedLte(vDataSyncedLte)
	}
	if len(vDataSyncedGt) > 0 {
		params.SetDataSyncedGt(vDataSyncedGt)
	}
	if len(vDataSyncedGte) > 0 {
		params.SetDataSyncedGte(vDataSyncedGte)
	}
	if vDataSyncedEmpty != nil {
		params.SetDataSyncedEmpty(vDataSyncedEmpty)
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
	if len(vDeviceRole) > 0 {
		params.SetDeviceRole(vDeviceRole)
	}
	if len(vDeviceRolen) > 0 {
		params.SetDeviceRolen(vDeviceRolen)
	}
	if len(vDeviceRoleID) > 0 {
		params.SetDeviceRoleID(vDeviceRoleID)
	}
	if len(vDeviceRoleIDn) > 0 {
		params.SetDeviceRoleIDn(vDeviceRoleIDn)
	}
	if len(vDeviceTypeID) > 0 {
		params.SetDeviceTypeID(vDeviceTypeID)
	}
	if len(vDeviceTypeIDn) > 0 {
		params.SetDeviceTypeIDn(vDeviceTypeIDn)
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
	if len(vLocation) > 0 {
		params.SetLocation(vLocation)
	}
	if len(vLocationn) > 0 {
		params.SetLocationn(vLocationn)
	}
	if len(vLocationID) > 0 {
		params.SetLocationID(vLocationID)
	}
	if len(vLocationIDn) > 0 {
		params.SetLocationIDn(vLocationIDn)
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
	if len(vPlatform) > 0 {
		params.SetPlatform(vPlatform)
	}
	if len(vPlatformn) > 0 {
		params.SetPlatformn(vPlatformn)
	}
	if len(vPlatformID) > 0 {
		params.SetPlatformID(vPlatformID)
	}
	if len(vPlatformIDn) > 0 {
		params.SetPlatformIDn(vPlatformIDn)
	}
	if len(vProfile) > 0 {
		params.SetProfile(vProfile)
	}
	if len(vProfilen) > 0 {
		params.SetProfilen(vProfilen)
	}
	if len(vProfileID) > 0 {
		params.SetProfileID(vProfileID)
	}
	if len(vProfileIDn) > 0 {
		params.SetProfileIDn(vProfileIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRegion) > 0 {
		params.SetRegion(vRegion)
	}
	if len(vRegionn) > 0 {
		params.SetRegionn(vRegionn)
	}
	if len(vRegionID) > 0 {
		params.SetRegionID(vRegionID)
	}
	if len(vRegionIDn) > 0 {
		params.SetRegionIDn(vRegionIDn)
	}
	if len(vSite) > 0 {
		params.SetSite(vSite)
	}
	if len(vSiten) > 0 {
		params.SetSiten(vSiten)
	}
	if len(vSiteGroup) > 0 {
		params.SetSiteGroup(vSiteGroup)
	}
	if len(vSiteGroupn) > 0 {
		params.SetSiteGroupn(vSiteGroupn)
	}
	if len(vSiteGroupID) > 0 {
		params.SetSiteGroupID(vSiteGroupID)
	}
	if len(vSiteGroupIDn) > 0 {
		params.SetSiteGroupIDn(vSiteGroupIDn)
	}
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
	}
	if len(vTag) > 0 {
		params.SetTag(vTag)
	}
	if len(vTagn) > 0 {
		params.SetTagn(vTagn)
	}
	if len(vTagID) > 0 {
		params.SetTagID(vTagID)
	}
	if len(vTagIDn) > 0 {
		params.SetTagIDn(vTagIDn)
	}
	if len(vTenant) > 0 {
		params.SetTenant(vTenant)
	}
	if len(vTenantn) > 0 {
		params.SetTenantn(vTenantn)
	}
	if len(vTenantGroup) > 0 {
		params.SetTenantGroup(vTenantGroup)
	}
	if len(vTenantGroupn) > 0 {
		params.SetTenantGroupn(vTenantGroupn)
	}
	if len(vTenantGroupID) > 0 {
		params.SetTenantGroupID(vTenantGroupID)
	}
	if len(vTenantGroupIDn) > 0 {
		params.SetTenantGroupIDn(vTenantGroupIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
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
