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
	_ datasource.DataSource              = (*asnsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*asnsDataSource)(nil)
)

// NewAsnsDataSource returns a new asns data source, which
// lists asn objects matching its filters.
func NewAsnsDataSource() datasource.DataSource {
	return &asnsDataSource{}
}

type asnsDataSource struct {
	client *netboxapi.Client
}

// asnsDataSourceModel is the data source model: the filters, the limit and the matching
// asns.
type asnsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"asns"`
}

// asnResourceAttrTypes is the attribute type map of asnResourceModel.
func asnResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             types.Int64Type,
		"asn":            types.Int64Type,
		"rir_id":         types.Int64Type,
		"tenant_id":      types.Int64Type,
		"role_id":        types.Int64Type,
		"site_ids":       types.SetType{ElemType: types.Int64Type},
		"description":    types.StringType,
		"comments":       types.StringType,
		"owner_id":       types.Int64Type,
		"created":        types.StringType,
		"last_updated":   types.StringType,
		"url":            types.StringType,
		"site_count":     types.Int64Type,
		"provider_count": types.Int64Type,
		"tags":           types.SetType{ElemType: types.StringType},
		"tags_all":       types.SetType{ElemType: types.StringType},
		"custom_fields":  types.MapType{ElemType: types.StringType},
	}
}

func (d *asnsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_asns"
}

func (d *asnsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists asn objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: asn, asn__empty, asn__gt, asn__gte, asn__lt, asn__lte, asn__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, provider, provider__n, provider_id, provider_id__n, q, rir, rir__n, rir_id, rir_id__n, role, role__n, role_id, role_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"asns": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching asns.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"asn": schema.Int64Attribute{
							Computed:    true,
							Description: "AS number. Must be a whole number.",
						},
						"rir_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the RIR the ASN belongs to.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the role (ipam.role).",
						},
						"site_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of sites the ASN is assigned to.",
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
						"site_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of sites using the ASN.",
						},
						"provider_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of providers using the ASN.",
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

func (d *asnsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *asnsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data asnsDataSourceModel

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
	var responseDTOs []*netboxapi.AsnResponseDTO

	params := ipam.NewIpamAsnsListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vAsn []int64
	var vAsnn []int64
	var vAsnLt []int64
	var vAsnLte []int64
	var vAsnGt []int64
	var vAsnGte []int64
	var vAsnEmpty *bool
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
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
	var vProvider []string
	var vProvidern []string
	var vProviderID []int64
	var vProviderIDn []int64
	var vQ *string
	var vRir []string
	var vRirn []string
	var vRirID []int64
	var vRirIDn []int64
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
	var vSiteID []int64
	var vSiteIDn []int64
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
	var vTenantID []int64
	var vTenantIDn []int64
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
		case "asn":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asn' takes an integer, got %q.", value))
				return
			}
			vAsn = append(vAsn, v)
		case "asn__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asn__n' takes an integer, got %q.", value))
				return
			}
			vAsnn = append(vAsnn, v)
		case "asn__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asn__lt' takes an integer, got %q.", value))
				return
			}
			vAsnLt = append(vAsnLt, v)
		case "asn__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asn__lte' takes an integer, got %q.", value))
				return
			}
			vAsnLte = append(vAsnLte, v)
		case "asn__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asn__gt' takes an integer, got %q.", value))
				return
			}
			vAsnGt = append(vAsnGt, v)
		case "asn__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asn__gte' takes an integer, got %q.", value))
				return
			}
			vAsnGte = append(vAsnGte, v)
		case "asn__empty":
			if vAsnEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'asn__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asn__empty' takes a boolean, got %q.", value))
				return
			}
			vAsnEmpty = &v
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
		case "provider":
			v := value
			vProvider = append(vProvider, v)
		case "provider__n":
			v := value
			vProvidern = append(vProvidern, v)
		case "provider_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_id' takes an integer, got %q.", value))
				return
			}
			vProviderID = append(vProviderID, v)
		case "provider_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_id__n' takes an integer, got %q.", value))
				return
			}
			vProviderIDn = append(vProviderIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "rir":
			v := value
			vRir = append(vRir, v)
		case "rir__n":
			v := value
			vRirn = append(vRirn, v)
		case "rir_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rir_id' takes an integer, got %q.", value))
				return
			}
			vRirID = append(vRirID, v)
		case "rir_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rir_id__n' takes an integer, got %q.", value))
				return
			}
			vRirIDn = append(vRirIDn, v)
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
	if len(vAsn) > 0 {
		params.SetAsn(vAsn)
	}
	if len(vAsnn) > 0 {
		params.SetAsnn(vAsnn)
	}
	if len(vAsnLt) > 0 {
		params.SetAsnLt(vAsnLt)
	}
	if len(vAsnLte) > 0 {
		params.SetAsnLte(vAsnLte)
	}
	if len(vAsnGt) > 0 {
		params.SetAsnGt(vAsnGt)
	}
	if len(vAsnGte) > 0 {
		params.SetAsnGte(vAsnGte)
	}
	if vAsnEmpty != nil {
		params.SetAsnEmpty(vAsnEmpty)
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
	if len(vProvider) > 0 {
		params.SetProvider(vProvider)
	}
	if len(vProvidern) > 0 {
		params.SetProvidern(vProvidern)
	}
	if len(vProviderID) > 0 {
		params.SetProviderID(vProviderID)
	}
	if len(vProviderIDn) > 0 {
		params.SetProviderIDn(vProviderIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRir) > 0 {
		params.SetRir(vRir)
	}
	if len(vRirn) > 0 {
		params.SetRirn(vRirn)
	}
	if len(vRirID) > 0 {
		params.SetRirID(vRirID)
	}
	if len(vRirIDn) > 0 {
		params.SetRirIDn(vRirIDn)
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
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
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
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Ipam.IpamAsnsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_asns", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.AsnResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]asnResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model asnResourceModel
		resp.Diagnostics.Append(flattenAsn(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: asnResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
