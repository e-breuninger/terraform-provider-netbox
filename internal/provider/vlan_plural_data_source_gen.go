// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	_ datasource.DataSource              = (*vlansDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlansDataSource)(nil)
)

// NewVlansDataSource returns a new vlans data source, which
// lists vlan objects matching its filters.
func NewVlansDataSource() datasource.DataSource {
	return &vlansDataSource{}
}

type vlansDataSource struct {
	client *netboxapi.Client
}

// vlansDataSourceModel is the data source model: the filters, the limit and the matching
// vlans.
type vlansDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"vlans"`
}

// vlanResourceAttrTypes is the attribute type map of vlanResourceModel.
func vlanResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"name":          types.StringType,
		"vid":           types.Int64Type,
		"status":        types.StringType,
		"site_id":       types.Int64Type,
		"group_id":      types.Int64Type,
		"tenant_id":     types.Int64Type,
		"role_id":       types.Int64Type,
		"qinq_role":     types.StringType,
		"qinq_svlan_id": types.Int64Type,
		"description":   types.StringType,
		"comments":      types.StringType,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"prefix_count":  types.Int64Type,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *vlansDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlans"
}

func (d *vlansDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists vlan objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: available_at_site, available_at_site_group, available_on_device, available_on_virtualmachine, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_id, l2vpn, l2vpn__n, l2vpn_id, l2vpn_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, qinq_role, qinq_role__empty, qinq_role__ic, qinq_role__ie, qinq_role__iew, qinq_role__iregex, qinq_role__isw, qinq_role__n, qinq_role__nic, qinq_role__nie, qinq_role__niew, qinq_role__nisw, qinq_role__regex, qinq_svlan_id, qinq_svlan_id__n, qinq_svlan_vid, qinq_svlan_vid__empty, qinq_svlan_vid__gt, qinq_svlan_vid__gte, qinq_svlan_vid__lt, qinq_svlan_vid__lte, qinq_svlan_vid__n, region, region__n, region_id, region_id__n, role, role__n, role_id, role_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, vid, vid__empty, vid__gt, vid__gte, vid__lt, vid__lte, vid__n, vminterface_id. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"vlans": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching vlans.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"vid": schema.Int64Attribute{
							Computed:    true,
							Description: "VLAN id (1-4094). Must be a whole number.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: active, reserved, deprecated.",
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site.",
						},
						"group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VLAN group.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the role.",
						},
						"qinq_role": schema.StringAttribute{
							Computed:    true,
							Description: "Q-in-Q role: a service VLAN (svlan) carrying customer VLANs, or a customer VLAN (cvlan) inside one. One of: svlan, cvlan.",
						},
						"qinq_svlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the service VLAN this customer VLAN is carried in (qinq_role cvlan).",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"comments": schema.StringAttribute{
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
						"prefix_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of prefixes on the VLAN.",
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

func (d *vlansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlansDataSourceModel

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
	var responseDTOs []*netboxapi.VlanResponseDTO

	params := ipam.NewIpamVlansListParams()
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
	var vVid []int64
	var vVidn []int64
	var vVidLt []int64
	var vVidLte []int64
	var vVidGt []int64
	var vVidGte []int64
	var vVidEmpty *bool
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
	var vSiteID []int64
	var vSiteIDn []int64
	var vGroupID []int64
	var vGroupIDn []int64
	var vTenantID []int64
	var vTenantIDn []int64
	var vGroup []string
	var vGroupn []string
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAvailableAtSite *string
	var vAvailableAtSiteGroup *string
	var vAvailableOnDevice *string
	var vAvailableOnVirtualmachine *string
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
	var vInterfaceID *int64
	var vL2vpn []int64
	var vL2vpnn []int64
	var vL2vpnID []int64
	var vL2vpnIDn []int64
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vQ *string
	var vQinqRole []string
	var vQinqRolen []string
	var vQinqRoleIc []string
	var vQinqRoleNic []string
	var vQinqRoleIe []string
	var vQinqRoleNie []string
	var vQinqRoleIsw []string
	var vQinqRoleNisw []string
	var vQinqRoleIew []string
	var vQinqRoleNiew []string
	var vQinqRoleEmpty *bool
	var vQinqRoleRegex []string
	var vQinqRoleIregex []string
	var vQinqSvlanID []int64
	var vQinqSvlanIDn []int64
	var vQinqSvlanVid []int64
	var vQinqSvlanVidn []int64
	var vQinqSvlanVidLt []int64
	var vQinqSvlanVidLte []int64
	var vQinqSvlanVidGt []int64
	var vQinqSvlanVidGte []int64
	var vQinqSvlanVidEmpty []int64
	var vRegion []string
	var vRegionn []string
	var vRegionID []string
	var vRegionIDn []string
	var vRole []string
	var vRolen []string
	var vRoleID []int64
	var vRoleIDn []int64
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []string
	var vSiteGroupIDn []string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vVminterfaceID *int64
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
		case "vid":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vid' takes an integer, got %q.", value))
				return
			}
			vVid = append(vVid, v)
		case "vid__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vid__n' takes an integer, got %q.", value))
				return
			}
			vVidn = append(vVidn, v)
		case "vid__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vid__lt' takes an integer, got %q.", value))
				return
			}
			vVidLt = append(vVidLt, v)
		case "vid__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vid__lte' takes an integer, got %q.", value))
				return
			}
			vVidLte = append(vVidLte, v)
		case "vid__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vid__gt' takes an integer, got %q.", value))
				return
			}
			vVidGt = append(vVidGt, v)
		case "vid__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vid__gte' takes an integer, got %q.", value))
				return
			}
			vVidGte = append(vVidGte, v)
		case "vid__empty":
			if vVidEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vid__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vid__empty' takes a boolean, got %q.", value))
				return
			}
			vVidEmpty = &v
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
		case "group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id' takes an integer, got %q.", value))
				return
			}
			vGroupID = append(vGroupID, v)
		case "group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__n' takes an integer, got %q.", value))
				return
			}
			vGroupIDn = append(vGroupIDn, v)
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
		case "group":
			v := value
			vGroup = append(vGroup, v)
		case "group__n":
			v := value
			vGroupn = append(vGroupn, v)
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
		case "available_at_site":
			if vAvailableAtSite != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'available_at_site' takes a single value.")
				return
			}
			v := value
			vAvailableAtSite = &v
		case "available_at_site_group":
			if vAvailableAtSiteGroup != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'available_at_site_group' takes a single value.")
				return
			}
			v := value
			vAvailableAtSiteGroup = &v
		case "available_on_device":
			if vAvailableOnDevice != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'available_on_device' takes a single value.")
				return
			}
			v := value
			vAvailableOnDevice = &v
		case "available_on_virtualmachine":
			if vAvailableOnVirtualmachine != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'available_on_virtualmachine' takes a single value.")
				return
			}
			v := value
			vAvailableOnVirtualmachine = &v
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
		case "interface_id":
			if vInterfaceID != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'interface_id' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_id' takes an integer, got %q.", value))
				return
			}
			vInterfaceID = &v
		case "l2vpn":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn' takes an integer, got %q.", value))
				return
			}
			vL2vpn = append(vL2vpn, v)
		case "l2vpn__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn__n' takes an integer, got %q.", value))
				return
			}
			vL2vpnn = append(vL2vpnn, v)
		case "l2vpn_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn_id' takes an integer, got %q.", value))
				return
			}
			vL2vpnID = append(vL2vpnID, v)
		case "l2vpn_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn_id__n' takes an integer, got %q.", value))
				return
			}
			vL2vpnIDn = append(vL2vpnIDn, v)
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
		case "qinq_role":
			v := value
			vQinqRole = append(vQinqRole, v)
		case "qinq_role__n":
			v := value
			vQinqRolen = append(vQinqRolen, v)
		case "qinq_role__ic":
			v := value
			vQinqRoleIc = append(vQinqRoleIc, v)
		case "qinq_role__nic":
			v := value
			vQinqRoleNic = append(vQinqRoleNic, v)
		case "qinq_role__ie":
			v := value
			vQinqRoleIe = append(vQinqRoleIe, v)
		case "qinq_role__nie":
			v := value
			vQinqRoleNie = append(vQinqRoleNie, v)
		case "qinq_role__isw":
			v := value
			vQinqRoleIsw = append(vQinqRoleIsw, v)
		case "qinq_role__nisw":
			v := value
			vQinqRoleNisw = append(vQinqRoleNisw, v)
		case "qinq_role__iew":
			v := value
			vQinqRoleIew = append(vQinqRoleIew, v)
		case "qinq_role__niew":
			v := value
			vQinqRoleNiew = append(vQinqRoleNiew, v)
		case "qinq_role__empty":
			if vQinqRoleEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'qinq_role__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_role__empty' takes a boolean, got %q.", value))
				return
			}
			vQinqRoleEmpty = &v
		case "qinq_role__regex":
			v := value
			vQinqRoleRegex = append(vQinqRoleRegex, v)
		case "qinq_role__iregex":
			v := value
			vQinqRoleIregex = append(vQinqRoleIregex, v)
		case "qinq_svlan_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_id' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanID = append(vQinqSvlanID, v)
		case "qinq_svlan_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_id__n' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanIDn = append(vQinqSvlanIDn, v)
		case "qinq_svlan_vid":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_vid' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanVid = append(vQinqSvlanVid, v)
		case "qinq_svlan_vid__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_vid__n' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanVidn = append(vQinqSvlanVidn, v)
		case "qinq_svlan_vid__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_vid__lt' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanVidLt = append(vQinqSvlanVidLt, v)
		case "qinq_svlan_vid__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_vid__lte' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanVidLte = append(vQinqSvlanVidLte, v)
		case "qinq_svlan_vid__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_vid__gt' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanVidGt = append(vQinqSvlanVidGt, v)
		case "qinq_svlan_vid__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_vid__gte' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanVidGte = append(vQinqSvlanVidGte, v)
		case "qinq_svlan_vid__empty":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'qinq_svlan_vid__empty' takes an integer, got %q.", value))
				return
			}
			vQinqSvlanVidEmpty = append(vQinqSvlanVidEmpty, v)
		case "region":
			v := value
			vRegion = append(vRegion, v)
		case "region__n":
			v := value
			vRegionn = append(vRegionn, v)
		case "region_id":
			v := value
			vRegionID = append(vRegionID, v)
		case "region_id__n":
			v := value
			vRegionIDn = append(vRegionIDn, v)
		case "role":
			v := value
			vRole = append(vRole, v)
		case "role__n":
			v := value
			vRolen = append(vRolen, v)
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
		case "vminterface_id":
			if vVminterfaceID != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vminterface_id' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vminterface_id' takes an integer, got %q.", value))
				return
			}
			vVminterfaceID = &v
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
	if len(vVid) > 0 {
		params.SetVid(vVid)
	}
	if len(vVidn) > 0 {
		params.SetVidn(vVidn)
	}
	if len(vVidLt) > 0 {
		params.SetVidLt(vVidLt)
	}
	if len(vVidLte) > 0 {
		params.SetVidLte(vVidLte)
	}
	if len(vVidGt) > 0 {
		params.SetVidGt(vVidGt)
	}
	if len(vVidGte) > 0 {
		params.SetVidGte(vVidGte)
	}
	if vVidEmpty != nil {
		params.SetVidEmpty(vVidEmpty)
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
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
	}
	if len(vGroupID) > 0 {
		params.SetGroupID(vGroupID)
	}
	if len(vGroupIDn) > 0 {
		params.SetGroupIDn(vGroupIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vGroup) > 0 {
		params.SetGroup(vGroup)
	}
	if len(vGroupn) > 0 {
		params.SetGroupn(vGroupn)
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
	if vAvailableAtSite != nil {
		params.SetAvailableAtSite(vAvailableAtSite)
	}
	if vAvailableAtSiteGroup != nil {
		params.SetAvailableAtSiteGroup(vAvailableAtSiteGroup)
	}
	if vAvailableOnDevice != nil {
		params.SetAvailableOnDevice(vAvailableOnDevice)
	}
	if vAvailableOnVirtualmachine != nil {
		params.SetAvailableOnVirtualmachine(vAvailableOnVirtualmachine)
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
	if vInterfaceID != nil {
		params.SetInterfaceID(vInterfaceID)
	}
	if len(vL2vpn) > 0 {
		params.SetL2vpn(vL2vpn)
	}
	if len(vL2vpnn) > 0 {
		params.SetL2vpnn(vL2vpnn)
	}
	if len(vL2vpnID) > 0 {
		params.SetL2vpnID(vL2vpnID)
	}
	if len(vL2vpnIDn) > 0 {
		params.SetL2vpnIDn(vL2vpnIDn)
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
	if len(vQinqRole) > 0 {
		params.SetQinqRole(vQinqRole)
	}
	if len(vQinqRolen) > 0 {
		params.SetQinqRolen(vQinqRolen)
	}
	if len(vQinqRoleIc) > 0 {
		params.SetQinqRoleIc(vQinqRoleIc)
	}
	if len(vQinqRoleNic) > 0 {
		params.SetQinqRoleNic(vQinqRoleNic)
	}
	if len(vQinqRoleIe) > 0 {
		params.SetQinqRoleIe(vQinqRoleIe)
	}
	if len(vQinqRoleNie) > 0 {
		params.SetQinqRoleNie(vQinqRoleNie)
	}
	if len(vQinqRoleIsw) > 0 {
		params.SetQinqRoleIsw(vQinqRoleIsw)
	}
	if len(vQinqRoleNisw) > 0 {
		params.SetQinqRoleNisw(vQinqRoleNisw)
	}
	if len(vQinqRoleIew) > 0 {
		params.SetQinqRoleIew(vQinqRoleIew)
	}
	if len(vQinqRoleNiew) > 0 {
		params.SetQinqRoleNiew(vQinqRoleNiew)
	}
	if vQinqRoleEmpty != nil {
		params.SetQinqRoleEmpty(vQinqRoleEmpty)
	}
	if len(vQinqRoleRegex) > 0 {
		params.SetQinqRoleRegex(vQinqRoleRegex)
	}
	if len(vQinqRoleIregex) > 0 {
		params.SetQinqRoleIregex(vQinqRoleIregex)
	}
	if len(vQinqSvlanID) > 0 {
		params.SetQinqSvlanID(vQinqSvlanID)
	}
	if len(vQinqSvlanIDn) > 0 {
		params.SetQinqSvlanIDn(vQinqSvlanIDn)
	}
	if len(vQinqSvlanVid) > 0 {
		params.SetQinqSvlanVid(vQinqSvlanVid)
	}
	if len(vQinqSvlanVidn) > 0 {
		params.SetQinqSvlanVidn(vQinqSvlanVidn)
	}
	if len(vQinqSvlanVidLt) > 0 {
		params.SetQinqSvlanVidLt(vQinqSvlanVidLt)
	}
	if len(vQinqSvlanVidLte) > 0 {
		params.SetQinqSvlanVidLte(vQinqSvlanVidLte)
	}
	if len(vQinqSvlanVidGt) > 0 {
		params.SetQinqSvlanVidGt(vQinqSvlanVidGt)
	}
	if len(vQinqSvlanVidGte) > 0 {
		params.SetQinqSvlanVidGte(vQinqSvlanVidGte)
	}
	if len(vQinqSvlanVidEmpty) > 0 {
		params.SetQinqSvlanVidEmpty(vQinqSvlanVidEmpty)
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
	if len(vRole) > 0 {
		params.SetRole(vRole)
	}
	if len(vRolen) > 0 {
		params.SetRolen(vRolen)
	}
	if len(vRoleID) > 0 {
		params.SetRoleID(vRoleID)
	}
	if len(vRoleIDn) > 0 {
		params.SetRoleIDn(vRoleIDn)
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
	if vVminterfaceID != nil {
		params.SetVminterfaceID(vVminterfaceID)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Ipam.IpamVlansListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_vlans", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VlanResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]vlanResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model vlanResourceModel
		resp.Diagnostics.Append(flattenVlan(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: vlanResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
