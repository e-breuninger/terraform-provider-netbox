// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
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
	_ datasource.DataSource              = (*prefixesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*prefixesDataSource)(nil)
)

// NewPrefixesDataSource returns a new prefixes data source, which
// lists prefix objects matching its filters.
func NewPrefixesDataSource() datasource.DataSource {
	return &prefixesDataSource{}
}

type prefixesDataSource struct {
	client *netboxapi.Client
}

// prefixesDataSourceModel is the data source model: the filters, the limit and the matching
// prefixes.
type prefixesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"prefixes"`
}

// prefixResourceAttrTypes is the attribute type map of prefixResourceModel.
func prefixResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"prefix":        conv.CIDRType{},
		"status":        types.StringType,
		"role_id":       types.Int64Type,
		"tenant_id":     types.Int64Type,
		"vlan_id":       types.Int64Type,
		"vrf_id":        types.Int64Type,
		"is_pool":       types.BoolType,
		"mark_utilized": types.BoolType,
		"family":        types.Int64Type,
		"description":   types.StringType,
		"comments":      types.StringType,
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
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *prefixesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_prefixes"
}

func (d *prefixesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists prefix objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: children, children__empty, children__gt, children__gte, children__lt, children__lte, children__n, contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, contains, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, depth, depth__empty, depth__gt, depth__gte, depth__lt, depth__lte, depth__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, family, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, is_pool, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, mark_utilized, mask_length, mask_length__gte, mask_length__lte, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, prefix, present_in_vrf, present_in_vrf_id, q, region, region__n, region_id, region_id__n, role, role__n, role_id, role_id__n, scope_id, scope_id__empty, scope_id__gt, scope_id__gte, scope_id__lt, scope_id__lte, scope_id__n, scope_type, scope_type__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, vlan_group, vlan_group__n, vlan_group_id, vlan_group_id__n, vlan_id, vlan_id__n, vlan_vid, vlan_vid__empty, vlan_vid__gt, vlan_vid__gte, vlan_vid__lt, vlan_vid__lte, vlan_vid__n, vrf, vrf__n, vrf_id, vrf_id__n, within, within_include. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"prefixes": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching prefixes.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"prefix": schema.StringAttribute{
							CustomType:  conv.CIDRType{},
							Computed:    true,
							Description: "IPv4 or IPv6 network with mask, e.g. 10.0.0.0/24.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: container, active, reserved, deprecated.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the role.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"vlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VLAN.",
						},
						"vrf_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VRF.",
						},
						"is_pool": schema.BoolAttribute{
							Computed:    true,
							Description: "All addresses within this prefix are usable.",
						},
						"mark_utilized": schema.BoolAttribute{
							Computed:    true,
							Description: "Report space as fully utilized.",
						},
						"family": schema.Int64Attribute{
							Computed:    true,
							Description: "Address family: 4 or 6.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
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

func (d *prefixesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *prefixesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data prefixesDataSourceModel

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
	var responseDTOs []*netboxapi.PrefixResponseDTO

	params := ipam.NewIpamPrefixesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"prefix":              netboxapi.FilterStrings(params.SetPrefix),
		"vrf_id":              netboxapi.FilterInt64s(params.SetVrfID),
		"vrf_id__n":           netboxapi.FilterInt64s(params.SetVrfIDn),
		"contains":            netboxapi.FilterString(params.SetContains),
		"status":              netboxapi.FilterStrings(params.SetStatus),
		"status__n":           netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":          netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":         netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":          netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":         netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":         netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":        netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":         netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":        netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":       netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":       netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":      netboxapi.FilterStrings(params.SetStatusIregex),
		"role_id":             netboxapi.FilterInt64s(params.SetRoleID),
		"role_id__n":          netboxapi.FilterInt64s(params.SetRoleIDn),
		"tenant_id":           netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":        netboxapi.FilterInt64s(params.SetTenantIDn),
		"vlan_id":             netboxapi.FilterInt64s(params.SetVlanID),
		"vlan_id__n":          netboxapi.FilterInt64s(params.SetVlanIDn),
		"vlan_vid":            netboxapi.FilterInt64(params.SetVlanVid),
		"vlan_vid__n":         netboxapi.FilterInt64(params.SetVlanVidn),
		"vlan_vid__lt":        netboxapi.FilterInt64(params.SetVlanVidLt),
		"vlan_vid__lte":       netboxapi.FilterInt64(params.SetVlanVidLte),
		"vlan_vid__gt":        netboxapi.FilterInt64(params.SetVlanVidGt),
		"vlan_vid__gte":       netboxapi.FilterInt64(params.SetVlanVidGte),
		"vlan_vid__empty":     netboxapi.FilterInt64(params.SetVlanVidEmpty),
		"site_id":             netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":          netboxapi.FilterInt64s(params.SetSiteIDn),
		"region_id":           netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":        netboxapi.FilterStrings(params.SetRegionIDn),
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
		"owner_id":            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":         netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":            netboxapi.FilterStrings(params.SetTagAny),
		"children":            netboxapi.FilterInt64s(params.SetChildren),
		"children__n":         netboxapi.FilterInt64s(params.SetChildrenn),
		"children__lt":        netboxapi.FilterInt64s(params.SetChildrenLt),
		"children__lte":       netboxapi.FilterInt64s(params.SetChildrenLte),
		"children__gt":        netboxapi.FilterInt64s(params.SetChildrenGt),
		"children__gte":       netboxapi.FilterInt64s(params.SetChildrenGte),
		"children__empty":     netboxapi.FilterInt64s(params.SetChildrenEmpty),
		"contact":             netboxapi.FilterInt64s(params.SetContact),
		"contact__n":          netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":       netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":    netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":        netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":     netboxapi.FilterInt64s(params.SetContactRolen),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"depth":               netboxapi.FilterInt64s(params.SetDepth),
		"depth__n":            netboxapi.FilterInt64s(params.SetDepthn),
		"depth__lt":           netboxapi.FilterInt64s(params.SetDepthLt),
		"depth__lte":          netboxapi.FilterInt64s(params.SetDepthLte),
		"depth__gt":           netboxapi.FilterInt64s(params.SetDepthGt),
		"depth__gte":          netboxapi.FilterInt64s(params.SetDepthGte),
		"depth__empty":        netboxapi.FilterInt64s(params.SetDepthEmpty),
		"family":              netboxapi.FilterFloat64(params.SetFamily),
		"is_pool":             netboxapi.FilterBool(params.SetIsPool),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":            netboxapi.FilterStrings(params.SetLocation),
		"location__n":         netboxapi.FilterStrings(params.SetLocationn),
		"location_id":         netboxapi.FilterStrings(params.SetLocationID),
		"location_id__n":      netboxapi.FilterStrings(params.SetLocationIDn),
		"mark_utilized":       netboxapi.FilterBool(params.SetMarkUtilized),
		"mask_length__gte":    netboxapi.FilterFloat64(params.SetMaskLengthGte),
		"mask_length__lte":    netboxapi.FilterFloat64(params.SetMaskLengthLte),
		"mask_length":         netboxapi.FilterInt64s(params.SetMaskLength),
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"present_in_vrf":      netboxapi.FilterString(params.SetPresentInVrf),
		"present_in_vrf_id":   netboxapi.FilterString(params.SetPresentInVrfID),
		"q":                   netboxapi.FilterString(params.SetQ),
		"region":              netboxapi.FilterStrings(params.SetRegion),
		"region__n":           netboxapi.FilterStrings(params.SetRegionn),
		"role":                netboxapi.FilterStrings(params.SetRole),
		"role__n":             netboxapi.FilterStrings(params.SetRolen),
		"scope_id":            netboxapi.FilterInt64s(params.SetScopeID),
		"scope_id__n":         netboxapi.FilterInt64s(params.SetScopeIDn),
		"scope_id__lt":        netboxapi.FilterInt64s(params.SetScopeIDLt),
		"scope_id__lte":       netboxapi.FilterInt64s(params.SetScopeIDLte),
		"scope_id__gt":        netboxapi.FilterInt64s(params.SetScopeIDGt),
		"scope_id__gte":       netboxapi.FilterInt64s(params.SetScopeIDGte),
		"scope_id__empty":     netboxapi.FilterBool(params.SetScopeIDEmpty),
		"scope_type":          netboxapi.FilterStrings(params.SetScopeType),
		"scope_type__n":       netboxapi.FilterStrings(params.SetScopeTypen),
		"site":                netboxapi.FilterStrings(params.SetSite),
		"site__n":             netboxapi.FilterStrings(params.SetSiten),
		"site_group":          netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":       netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":       netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":    netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":              netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":           netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":        netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":     netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":     netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":  netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"vlan_group":          netboxapi.FilterStrings(params.SetVlanGroup),
		"vlan_group__n":       netboxapi.FilterStrings(params.SetVlanGroupn),
		"vlan_group_id":       netboxapi.FilterInt64s(params.SetVlanGroupID),
		"vlan_group_id__n":    netboxapi.FilterInt64s(params.SetVlanGroupIDn),
		"vrf":                 netboxapi.FilterStrings(params.SetVrf),
		"vrf__n":              netboxapi.FilterStrings(params.SetVrfn),
		"within":              netboxapi.FilterString(params.SetWithin),
		"within_include":      netboxapi.FilterString(params.SetWithinInclude),
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
		res, err := d.client.Ipam.IpamPrefixesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_prefixes", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.PrefixResponseDTOFromGoNetbox(goNetboxModel))
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
	resourceWithHooks := &prefixResource{client: d.client}
	items := make([]prefixResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model prefixResourceModel
		resp.Diagnostics.Append(flattenPrefix(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: prefixResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
