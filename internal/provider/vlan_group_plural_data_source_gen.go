// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*vlanGroupsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlanGroupsDataSource)(nil)
)

// NewVlanGroupsDataSource returns a new vlan_groups data source, which
// lists vlan_group objects matching its filters.
func NewVlanGroupsDataSource() datasource.DataSource {
	return &vlanGroupsDataSource{}
}

type vlanGroupsDataSource struct {
	client *netboxapi.Client
}

// vlanGroupsDataSourceModel is the data source model: the filters, the limit and the matching
// vlan_groups.
type vlanGroupsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"vlan_groups"`
}

// vlanGroupResourceAttrTypes is the attribute type map of vlanGroupResourceModel.
func vlanGroupResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"name":          types.StringType,
		"slug":          types.StringType,
		"description":   types.StringType,
		"comments":      types.StringType,
		"vid_ranges":    types.ListType{ElemType: types.ObjectType{AttrTypes: vlanGroupVidRangesAttrTypes()}},
		"tenant_id":     types.Int64Type,
		"scope_type":    types.StringType,
		"scope_id":      types.Int64Type,
		"site_id":       types.Int64Type,
		"location_id":   types.Int64Type,
		"region_id":     types.Int64Type,
		"site_group_id": types.Int64Type,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"vlan_count":    types.Int64Type,
		"utilization":   types.StringType,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *vlanGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_groups"
}

func (d *vlanGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists vlan_group objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: cluster, cluster_group, contains_vid, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, rack, rack_group, region, scope_id, scope_id__empty, scope_id__gt, scope_id__gte, scope_id__lt, scope_id__lte, scope_id__n, scope_type, scope_type__n, site, site_group, slug, slug__empty, slug__ic, slug__ie, slug__iew, slug__iregex, slug__isw, slug__n, slug__nic, slug__nie, slug__niew, slug__nisw, slug__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, total_vlan_ids, total_vlan_ids__empty, total_vlan_ids__gt, total_vlan_ids__gte, total_vlan_ids__lt, total_vlan_ids__lte, total_vlan_ids__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"vlan_groups": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching vlan_groups.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"slug": schema.StringAttribute{
							Computed:    true,
							Description: "URL-friendly unique shorthand; derived from name when not set.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"vid_ranges": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"start": schema.Int64Attribute{
										Computed:    true,
										Description: "First VLAN id of the range. Must be a whole number.",
									},
									"end": schema.Int64Attribute{
										Computed:    true,
										Description: "Last VLAN id of the range. Must be a whole number.",
									},
								},
							},
							Computed:    true,
							Description: "VLAN id ranges of this group as {start, end} pairs.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"scope_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the scope. Derived from site_id, location_id, region_id or site_group_id when one of those is set; set it together with scope_id otherwise. One of: dcim.site, dcim.location, dcim.region, dcim.sitegroup.",
						},
						"scope_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the scope object (see scope_type).",
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site the object is scoped to (scope_type dcim.site). Conflicts with the other scope aliases and with setting the scope_* pair directly.",
						},
						"location_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the location the object is scoped to (scope_type dcim.location).",
						},
						"region_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the region the object is scoped to (scope_type dcim.region).",
						},
						"site_group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site group the object is scoped to (scope_type dcim.sitegroup).",
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
						"vlan_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of VLANs in the group.",
						},
						"utilization": schema.StringAttribute{
							Computed:    true,
							Description: "Share of the VID ranges in use, as NetBox reports it, e.g. \"12.50%\".",
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

func (d *vlanGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vlanGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlanGroupsDataSourceModel

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
	var responseDTOs []*netboxapi.VlanGroupResponseDTO

	params := ipam.NewIpamVlanGroupsListParams()
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
		"slug":                  netboxapi.FilterStrings(params.SetSlug),
		"slug__n":               netboxapi.FilterStrings(params.SetSlugn),
		"slug__ic":              netboxapi.FilterStrings(params.SetSlugIc),
		"slug__nic":             netboxapi.FilterStrings(params.SetSlugNic),
		"slug__ie":              netboxapi.FilterStrings(params.SetSlugIe),
		"slug__nie":             netboxapi.FilterStrings(params.SetSlugNie),
		"slug__isw":             netboxapi.FilterStrings(params.SetSlugIsw),
		"slug__nisw":            netboxapi.FilterStrings(params.SetSlugNisw),
		"slug__iew":             netboxapi.FilterStrings(params.SetSlugIew),
		"slug__niew":            netboxapi.FilterStrings(params.SetSlugNiew),
		"slug__empty":           netboxapi.FilterBool(params.SetSlugEmpty),
		"slug__regex":           netboxapi.FilterStrings(params.SetSlugRegex),
		"slug__iregex":          netboxapi.FilterStrings(params.SetSlugIregex),
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
		"scope_type":            netboxapi.FilterStrings(params.SetScopeType),
		"scope_type__n":         netboxapi.FilterStrings(params.SetScopeTypen),
		"scope_id":              netboxapi.FilterInt64s(params.SetScopeID),
		"scope_id__n":           netboxapi.FilterInt64s(params.SetScopeIDn),
		"scope_id__lt":          netboxapi.FilterInt64s(params.SetScopeIDLt),
		"scope_id__lte":         netboxapi.FilterInt64s(params.SetScopeIDLte),
		"scope_id__gt":          netboxapi.FilterInt64s(params.SetScopeIDGt),
		"scope_id__gte":         netboxapi.FilterInt64s(params.SetScopeIDGte),
		"scope_id__empty":       netboxapi.FilterBool(params.SetScopeIDEmpty),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"cluster":               netboxapi.FilterInt64(params.SetCluster),
		"cluster_group":         netboxapi.FilterInt64(params.SetClusterGroup),
		"contains_vid":          netboxapi.FilterFloat64(params.SetContainsVid),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":              netboxapi.FilterInt64(params.SetLocation),
		"owner":                 netboxapi.FilterStrings(params.SetOwner),
		"owner__n":              netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":           netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":        netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":        netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":     netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                     netboxapi.FilterString(params.SetQ),
		"rack":                  netboxapi.FilterInt64(params.SetRack),
		"rack_group":            netboxapi.FilterFloat64(params.SetRackGroup),
		"region":                netboxapi.FilterInt64(params.SetRegion),
		"site":                  netboxapi.FilterInt64(params.SetSite),
		"site_group":            netboxapi.FilterInt64(params.SetSiteGroup),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":             netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":          netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":       netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":       netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":    netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"tenant_id":             netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":          netboxapi.FilterInt64s(params.SetTenantIDn),
		"total_vlan_ids":        netboxapi.FilterInt64s(params.SetTotalVlanIds),
		"total_vlan_ids__n":     netboxapi.FilterInt64s(params.SetTotalVlanIdsn),
		"total_vlan_ids__lt":    netboxapi.FilterInt64s(params.SetTotalVlanIdsLt),
		"total_vlan_ids__lte":   netboxapi.FilterInt64s(params.SetTotalVlanIdsLte),
		"total_vlan_ids__gt":    netboxapi.FilterInt64s(params.SetTotalVlanIdsGt),
		"total_vlan_ids__gte":   netboxapi.FilterInt64s(params.SetTotalVlanIdsGte),
		"total_vlan_ids__empty": netboxapi.FilterBool(params.SetTotalVlanIdsEmpty),
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
		res, err := d.client.Ipam.IpamVlanGroupsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_vlan_groups", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VlanGroupResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) runs on every item, as on the resource
	// itself (see the singular data source).
	resourceWithHooks := &vlanGroupResource{client: d.client}
	items := make([]vlanGroupResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model vlanGroupResourceModel
		resp.Diagnostics.Append(flattenVlanGroup(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: vlanGroupResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
