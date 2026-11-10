// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vPrefix []string
	var vVrfID []int64
	var vVrfIDn []int64
	var vContains *string
	var vStatus []string
	var vStatusn []string
	var vStatusIc []string
	var vStatusNic []string
	var vStatusIe []string
	var vStatusNie []string
	var vStatusIsw []string
	var vStatusNisw []string
	var vStatusIew []string
	var vStatusNiew []string
	var vStatusEmpty *bool
	var vStatusRegex []string
	var vStatusIregex []string
	var vRoleID []int64
	var vRoleIDn []int64
	var vTenantID []int64
	var vTenantIDn []int64
	var vVlanID []int64
	var vVlanIDn []int64
	var vVlanVid *int64
	var vVlanVidn *int64
	var vVlanVidLt *int64
	var vVlanVidLte *int64
	var vVlanVidGt *int64
	var vVlanVidGte *int64
	var vVlanVidEmpty *int64
	var vSiteID []int64
	var vSiteIDn []int64
	var vRegionID []string
	var vRegionIDn []string
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
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vChildren []int64
	var vChildrenn []int64
	var vChildrenLt []int64
	var vChildrenLte []int64
	var vChildrenGt []int64
	var vChildrenGte []int64
	var vChildrenEmpty []int64
	var vContact []int64
	var vContactn []int64
	var vContactGroup []string
	var vContactGroupn []string
	var vContactRole []int64
	var vContactRolen []int64
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vDepth []int64
	var vDepthn []int64
	var vDepthLt []int64
	var vDepthLte []int64
	var vDepthGt []int64
	var vDepthGte []int64
	var vDepthEmpty []int64
	var vFamily *float64
	var vIsPool *bool
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLocation []string
	var vLocationn []string
	var vLocationID []string
	var vLocationIDn []string
	var vMarkUtilized *bool
	var vMaskLengthGte *float64
	var vMaskLengthLte *float64
	var vMaskLength []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPresentInVrf *string
	var vPresentInVrfID *string
	var vQ *string
	var vRegion []string
	var vRegionn []string
	var vRole []string
	var vRolen []string
	var vScopeID []int64
	var vScopeIDn []int64
	var vScopeIDLt []int64
	var vScopeIDLte []int64
	var vScopeIDGt []int64
	var vScopeIDGte []int64
	var vScopeIDEmpty *bool
	var vScopeType []string
	var vScopeTypen []string
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []string
	var vSiteGroupIDn []string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
	var vVlanGroup []string
	var vVlanGroupn []string
	var vVlanGroupID []int64
	var vVlanGroupIDn []int64
	var vVrf []string
	var vVrfn []string
	var vWithin *string
	var vWithinInclude *string
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
		case "prefix":
			v := value
			vPrefix = append(vPrefix, v)
		case "vrf_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vrf_id' takes an integer, got %q.", value))
				return
			}
			vVrfID = append(vVrfID, v)
		case "vrf_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vrf_id__n' takes an integer, got %q.", value))
				return
			}
			vVrfIDn = append(vVrfIDn, v)
		case "contains":
			if vContains != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'contains' takes a single value.")
				return
			}
			v := value
			vContains = &v
		case "status":
			v := value
			vStatus = append(vStatus, v)
		case "status__n":
			v := value
			vStatusn = append(vStatusn, v)
		case "status__ic":
			v := value
			vStatusIc = append(vStatusIc, v)
		case "status__nic":
			v := value
			vStatusNic = append(vStatusNic, v)
		case "status__ie":
			v := value
			vStatusIe = append(vStatusIe, v)
		case "status__nie":
			v := value
			vStatusNie = append(vStatusNie, v)
		case "status__isw":
			v := value
			vStatusIsw = append(vStatusIsw, v)
		case "status__nisw":
			v := value
			vStatusNisw = append(vStatusNisw, v)
		case "status__iew":
			v := value
			vStatusIew = append(vStatusIew, v)
		case "status__niew":
			v := value
			vStatusNiew = append(vStatusNiew, v)
		case "status__empty":
			if vStatusEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'status__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'status__empty' takes a boolean, got %q.", value))
				return
			}
			vStatusEmpty = &v
		case "status__regex":
			v := value
			vStatusRegex = append(vStatusRegex, v)
		case "status__iregex":
			v := value
			vStatusIregex = append(vStatusIregex, v)
		case "role_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id' takes an integer, got %q.", value))
				return
			}
			vRoleID = append(vRoleID, v)
		case "role_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id__n' takes an integer, got %q.", value))
				return
			}
			vRoleIDn = append(vRoleIDn, v)
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
		case "vlan_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_id' takes an integer, got %q.", value))
				return
			}
			vVlanID = append(vVlanID, v)
		case "vlan_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_id__n' takes an integer, got %q.", value))
				return
			}
			vVlanIDn = append(vVlanIDn, v)
		case "vlan_vid":
			if vVlanVid != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid' takes an integer, got %q.", value))
				return
			}
			vVlanVid = &v
		case "vlan_vid__n":
			if vVlanVidn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__n' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__n' takes an integer, got %q.", value))
				return
			}
			vVlanVidn = &v
		case "vlan_vid__lt":
			if vVlanVidLt != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__lt' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__lt' takes an integer, got %q.", value))
				return
			}
			vVlanVidLt = &v
		case "vlan_vid__lte":
			if vVlanVidLte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__lte' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__lte' takes an integer, got %q.", value))
				return
			}
			vVlanVidLte = &v
		case "vlan_vid__gt":
			if vVlanVidGt != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__gt' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__gt' takes an integer, got %q.", value))
				return
			}
			vVlanVidGt = &v
		case "vlan_vid__gte":
			if vVlanVidGte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__gte' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__gte' takes an integer, got %q.", value))
				return
			}
			vVlanVidGte = &v
		case "vlan_vid__empty":
			if vVlanVidEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__empty' takes an integer, got %q.", value))
				return
			}
			vVlanVidEmpty = &v
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
		case "region_id":
			v := value
			vRegionID = append(vRegionID, v)
		case "region_id__n":
			v := value
			vRegionIDn = append(vRegionIDn, v)
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
		case "children":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'children' takes an integer, got %q.", value))
				return
			}
			vChildren = append(vChildren, v)
		case "children__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'children__n' takes an integer, got %q.", value))
				return
			}
			vChildrenn = append(vChildrenn, v)
		case "children__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'children__lt' takes an integer, got %q.", value))
				return
			}
			vChildrenLt = append(vChildrenLt, v)
		case "children__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'children__lte' takes an integer, got %q.", value))
				return
			}
			vChildrenLte = append(vChildrenLte, v)
		case "children__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'children__gt' takes an integer, got %q.", value))
				return
			}
			vChildrenGt = append(vChildrenGt, v)
		case "children__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'children__gte' takes an integer, got %q.", value))
				return
			}
			vChildrenGte = append(vChildrenGte, v)
		case "children__empty":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'children__empty' takes an integer, got %q.", value))
				return
			}
			vChildrenEmpty = append(vChildrenEmpty, v)
		case "contact":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact' takes an integer, got %q.", value))
				return
			}
			vContact = append(vContact, v)
		case "contact__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact__n' takes an integer, got %q.", value))
				return
			}
			vContactn = append(vContactn, v)
		case "contact_group":
			v := value
			vContactGroup = append(vContactGroup, v)
		case "contact_group__n":
			v := value
			vContactGroupn = append(vContactGroupn, v)
		case "contact_role":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_role' takes an integer, got %q.", value))
				return
			}
			vContactRole = append(vContactRole, v)
		case "contact_role__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_role__n' takes an integer, got %q.", value))
				return
			}
			vContactRolen = append(vContactRolen, v)
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
		case "depth":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'depth' takes an integer, got %q.", value))
				return
			}
			vDepth = append(vDepth, v)
		case "depth__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'depth__n' takes an integer, got %q.", value))
				return
			}
			vDepthn = append(vDepthn, v)
		case "depth__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'depth__lt' takes an integer, got %q.", value))
				return
			}
			vDepthLt = append(vDepthLt, v)
		case "depth__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'depth__lte' takes an integer, got %q.", value))
				return
			}
			vDepthLte = append(vDepthLte, v)
		case "depth__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'depth__gt' takes an integer, got %q.", value))
				return
			}
			vDepthGt = append(vDepthGt, v)
		case "depth__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'depth__gte' takes an integer, got %q.", value))
				return
			}
			vDepthGte = append(vDepthGte, v)
		case "depth__empty":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'depth__empty' takes an integer, got %q.", value))
				return
			}
			vDepthEmpty = append(vDepthEmpty, v)
		case "family":
			if vFamily != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'family' takes a single value.")
				return
			}
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'family' takes a number, got %q.", value))
				return
			}
			vFamily = &v
		case "is_pool":
			if vIsPool != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'is_pool' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'is_pool' takes a boolean, got %q.", value))
				return
			}
			vIsPool = &v
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
			v := value
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v := value
			vLocationIDn = append(vLocationIDn, v)
		case "mark_utilized":
			if vMarkUtilized != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mark_utilized' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mark_utilized' takes a boolean, got %q.", value))
				return
			}
			vMarkUtilized = &v
		case "mask_length__gte":
			if vMaskLengthGte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mask_length__gte' takes a single value.")
				return
			}
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mask_length__gte' takes a number, got %q.", value))
				return
			}
			vMaskLengthGte = &v
		case "mask_length__lte":
			if vMaskLengthLte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mask_length__lte' takes a single value.")
				return
			}
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mask_length__lte' takes a number, got %q.", value))
				return
			}
			vMaskLengthLte = &v
		case "mask_length":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mask_length' takes an integer, got %q.", value))
				return
			}
			vMaskLength = append(vMaskLength, v)
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
		case "present_in_vrf":
			if vPresentInVrf != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'present_in_vrf' takes a single value.")
				return
			}
			v := value
			vPresentInVrf = &v
		case "present_in_vrf_id":
			if vPresentInVrfID != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'present_in_vrf_id' takes a single value.")
				return
			}
			v := value
			vPresentInVrfID = &v
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
		case "role":
			v := value
			vRole = append(vRole, v)
		case "role__n":
			v := value
			vRolen = append(vRolen, v)
		case "scope_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'scope_id' takes an integer, got %q.", value))
				return
			}
			vScopeID = append(vScopeID, v)
		case "scope_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'scope_id__n' takes an integer, got %q.", value))
				return
			}
			vScopeIDn = append(vScopeIDn, v)
		case "scope_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'scope_id__lt' takes an integer, got %q.", value))
				return
			}
			vScopeIDLt = append(vScopeIDLt, v)
		case "scope_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'scope_id__lte' takes an integer, got %q.", value))
				return
			}
			vScopeIDLte = append(vScopeIDLte, v)
		case "scope_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'scope_id__gt' takes an integer, got %q.", value))
				return
			}
			vScopeIDGt = append(vScopeIDGt, v)
		case "scope_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'scope_id__gte' takes an integer, got %q.", value))
				return
			}
			vScopeIDGte = append(vScopeIDGte, v)
		case "scope_id__empty":
			if vScopeIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'scope_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'scope_id__empty' takes a boolean, got %q.", value))
				return
			}
			vScopeIDEmpty = &v
		case "scope_type":
			v := value
			vScopeType = append(vScopeType, v)
		case "scope_type__n":
			v := value
			vScopeTypen = append(vScopeTypen, v)
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
			v := value
			vSiteGroupID = append(vSiteGroupID, v)
		case "site_group_id__n":
			v := value
			vSiteGroupIDn = append(vSiteGroupIDn, v)
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
			v := value
			vTenantGroupID = append(vTenantGroupID, v)
		case "tenant_group_id__n":
			v := value
			vTenantGroupIDn = append(vTenantGroupIDn, v)
		case "vlan_group":
			v := value
			vVlanGroup = append(vVlanGroup, v)
		case "vlan_group__n":
			v := value
			vVlanGroupn = append(vVlanGroupn, v)
		case "vlan_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_group_id' takes an integer, got %q.", value))
				return
			}
			vVlanGroupID = append(vVlanGroupID, v)
		case "vlan_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_group_id__n' takes an integer, got %q.", value))
				return
			}
			vVlanGroupIDn = append(vVlanGroupIDn, v)
		case "vrf":
			v := value
			vVrf = append(vVrf, v)
		case "vrf__n":
			v := value
			vVrfn = append(vVrfn, v)
		case "within":
			if vWithin != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'within' takes a single value.")
				return
			}
			v := value
			vWithin = &v
		case "within_include":
			if vWithinInclude != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'within_include' takes a single value.")
				return
			}
			v := value
			vWithinInclude = &v
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
	if len(vPrefix) > 0 {
		params.SetPrefix(vPrefix)
	}
	if len(vVrfID) > 0 {
		params.SetVrfID(vVrfID)
	}
	if len(vVrfIDn) > 0 {
		params.SetVrfIDn(vVrfIDn)
	}
	if vContains != nil {
		params.SetContains(vContains)
	}
	if len(vStatus) > 0 {
		params.SetStatus(vStatus)
	}
	if len(vStatusn) > 0 {
		params.SetStatusn(vStatusn)
	}
	if len(vStatusIc) > 0 {
		params.SetStatusIc(vStatusIc)
	}
	if len(vStatusNic) > 0 {
		params.SetStatusNic(vStatusNic)
	}
	if len(vStatusIe) > 0 {
		params.SetStatusIe(vStatusIe)
	}
	if len(vStatusNie) > 0 {
		params.SetStatusNie(vStatusNie)
	}
	if len(vStatusIsw) > 0 {
		params.SetStatusIsw(vStatusIsw)
	}
	if len(vStatusNisw) > 0 {
		params.SetStatusNisw(vStatusNisw)
	}
	if len(vStatusIew) > 0 {
		params.SetStatusIew(vStatusIew)
	}
	if len(vStatusNiew) > 0 {
		params.SetStatusNiew(vStatusNiew)
	}
	if vStatusEmpty != nil {
		params.SetStatusEmpty(vStatusEmpty)
	}
	if len(vStatusRegex) > 0 {
		params.SetStatusRegex(vStatusRegex)
	}
	if len(vStatusIregex) > 0 {
		params.SetStatusIregex(vStatusIregex)
	}
	if len(vRoleID) > 0 {
		params.SetRoleID(vRoleID)
	}
	if len(vRoleIDn) > 0 {
		params.SetRoleIDn(vRoleIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vVlanID) > 0 {
		params.SetVlanID(vVlanID)
	}
	if len(vVlanIDn) > 0 {
		params.SetVlanIDn(vVlanIDn)
	}
	if vVlanVid != nil {
		params.SetVlanVid(vVlanVid)
	}
	if vVlanVidn != nil {
		params.SetVlanVidn(vVlanVidn)
	}
	if vVlanVidLt != nil {
		params.SetVlanVidLt(vVlanVidLt)
	}
	if vVlanVidLte != nil {
		params.SetVlanVidLte(vVlanVidLte)
	}
	if vVlanVidGt != nil {
		params.SetVlanVidGt(vVlanVidGt)
	}
	if vVlanVidGte != nil {
		params.SetVlanVidGte(vVlanVidGte)
	}
	if vVlanVidEmpty != nil {
		params.SetVlanVidEmpty(vVlanVidEmpty)
	}
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
	}
	if len(vRegionID) > 0 {
		params.SetRegionID(vRegionID)
	}
	if len(vRegionIDn) > 0 {
		params.SetRegionIDn(vRegionIDn)
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
	if len(vChildren) > 0 {
		params.SetChildren(vChildren)
	}
	if len(vChildrenn) > 0 {
		params.SetChildrenn(vChildrenn)
	}
	if len(vChildrenLt) > 0 {
		params.SetChildrenLt(vChildrenLt)
	}
	if len(vChildrenLte) > 0 {
		params.SetChildrenLte(vChildrenLte)
	}
	if len(vChildrenGt) > 0 {
		params.SetChildrenGt(vChildrenGt)
	}
	if len(vChildrenGte) > 0 {
		params.SetChildrenGte(vChildrenGte)
	}
	if len(vChildrenEmpty) > 0 {
		params.SetChildrenEmpty(vChildrenEmpty)
	}
	if len(vContact) > 0 {
		params.SetContact(vContact)
	}
	if len(vContactn) > 0 {
		params.SetContactn(vContactn)
	}
	if len(vContactGroup) > 0 {
		params.SetContactGroup(vContactGroup)
	}
	if len(vContactGroupn) > 0 {
		params.SetContactGroupn(vContactGroupn)
	}
	if len(vContactRole) > 0 {
		params.SetContactRole(vContactRole)
	}
	if len(vContactRolen) > 0 {
		params.SetContactRolen(vContactRolen)
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
	if len(vDepth) > 0 {
		params.SetDepth(vDepth)
	}
	if len(vDepthn) > 0 {
		params.SetDepthn(vDepthn)
	}
	if len(vDepthLt) > 0 {
		params.SetDepthLt(vDepthLt)
	}
	if len(vDepthLte) > 0 {
		params.SetDepthLte(vDepthLte)
	}
	if len(vDepthGt) > 0 {
		params.SetDepthGt(vDepthGt)
	}
	if len(vDepthGte) > 0 {
		params.SetDepthGte(vDepthGte)
	}
	if len(vDepthEmpty) > 0 {
		params.SetDepthEmpty(vDepthEmpty)
	}
	if vFamily != nil {
		params.SetFamily(vFamily)
	}
	if vIsPool != nil {
		params.SetIsPool(vIsPool)
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
	if vMarkUtilized != nil {
		params.SetMarkUtilized(vMarkUtilized)
	}
	if vMaskLengthGte != nil {
		params.SetMaskLengthGte(vMaskLengthGte)
	}
	if vMaskLengthLte != nil {
		params.SetMaskLengthLte(vMaskLengthLte)
	}
	if len(vMaskLength) > 0 {
		params.SetMaskLength(vMaskLength)
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
	if vPresentInVrf != nil {
		params.SetPresentInVrf(vPresentInVrf)
	}
	if vPresentInVrfID != nil {
		params.SetPresentInVrfID(vPresentInVrfID)
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
	if len(vRole) > 0 {
		params.SetRole(vRole)
	}
	if len(vRolen) > 0 {
		params.SetRolen(vRolen)
	}
	if len(vScopeID) > 0 {
		params.SetScopeID(vScopeID)
	}
	if len(vScopeIDn) > 0 {
		params.SetScopeIDn(vScopeIDn)
	}
	if len(vScopeIDLt) > 0 {
		params.SetScopeIDLt(vScopeIDLt)
	}
	if len(vScopeIDLte) > 0 {
		params.SetScopeIDLte(vScopeIDLte)
	}
	if len(vScopeIDGt) > 0 {
		params.SetScopeIDGt(vScopeIDGt)
	}
	if len(vScopeIDGte) > 0 {
		params.SetScopeIDGte(vScopeIDGte)
	}
	if vScopeIDEmpty != nil {
		params.SetScopeIDEmpty(vScopeIDEmpty)
	}
	if len(vScopeType) > 0 {
		params.SetScopeType(vScopeType)
	}
	if len(vScopeTypen) > 0 {
		params.SetScopeTypen(vScopeTypen)
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
	if len(vTagID) > 0 {
		params.SetTagID(vTagID)
	}
	if len(vTagIDn) > 0 {
		params.SetTagIDn(vTagIDn)
	}
	if len(vTagIDAny) > 0 {
		params.SetTagIDAny(vTagIDAny)
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
	if len(vVlanGroup) > 0 {
		params.SetVlanGroup(vVlanGroup)
	}
	if len(vVlanGroupn) > 0 {
		params.SetVlanGroupn(vVlanGroupn)
	}
	if len(vVlanGroupID) > 0 {
		params.SetVlanGroupID(vVlanGroupID)
	}
	if len(vVlanGroupIDn) > 0 {
		params.SetVlanGroupIDn(vVlanGroupIDn)
	}
	if len(vVrf) > 0 {
		params.SetVrf(vVrf)
	}
	if len(vVrfn) > 0 {
		params.SetVrfn(vVrfn)
	}
	if vWithin != nil {
		params.SetWithin(vWithin)
	}
	if vWithinInclude != nil {
		params.SetWithinInclude(vWithinInclude)
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
