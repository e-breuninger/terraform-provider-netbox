// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*vpnTunnelsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vpnTunnelsDataSource)(nil)
)

// NewVpnTunnelsDataSource returns a new vpn_tunnels data source, which
// lists vpn_tunnel objects matching its filters.
func NewVpnTunnelsDataSource() datasource.DataSource {
	return &vpnTunnelsDataSource{}
}

type vpnTunnelsDataSource struct {
	client *netboxapi.Client
}

// vpnTunnelsDataSourceModel is the data source model: the filters, the limit and the matching
// vpn_tunnels.
type vpnTunnelsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"vpn_tunnels"`
}

// vpnTunnelResourceAttrTypes is the attribute type map of vpnTunnelResourceModel.
func vpnTunnelResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                  types.Int64Type,
		"name":                types.StringType,
		"status":              types.StringType,
		"encapsulation":       types.StringType,
		"vpn_tunnel_group_id": types.Int64Type,
		"tunnel_group_id":     types.Int64Type,
		"ipsec_profile_id":    types.Int64Type,
		"tenant_id":           types.Int64Type,
		"tunnel_id":           types.Int64Type,
		"description":         types.StringType,
		"comments":            types.StringType,
		"owner_id":            types.Int64Type,
		"created":             types.StringType,
		"last_updated":        types.StringType,
		"url":                 types.StringType,
		"termination_count":   types.Int64Type,
		"tags":                types.SetType{ElemType: types.StringType},
		"tags_all":            types.SetType{ElemType: types.StringType},
		"custom_fields":       types.MapType{ElemType: types.StringType},
	}
}

func (d *vpnTunnelsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_tunnels"
}

func (d *vpnTunnelsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:Lists vpn_tunnel objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, encapsulation, encapsulation__empty, encapsulation__ic, encapsulation__ie, encapsulation__iew, encapsulation__iregex, encapsulation__isw, encapsulation__n, encapsulation__nic, encapsulation__nie, encapsulation__niew, encapsulation__nisw, encapsulation__regex, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, ipsec_profile, ipsec_profile__n, ipsec_profile_id, ipsec_profile_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, tunnel_id, tunnel_id__empty, tunnel_id__gt, tunnel_id__gte, tunnel_id__lt, tunnel_id__lte, tunnel_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"vpn_tunnels": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching vpn_tunnels.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: planned, active, disabled.",
						},
						"encapsulation": schema.StringAttribute{
							Computed:    true,
							Description: "Tunnel encapsulation protocol. One of: ipsec-transport, ipsec-tunnel, ip-ip, gre, wireguard, openvpn, l2tp, pptp.",
						},
						"vpn_tunnel_group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tunnel group.",
						},
						"tunnel_group_id": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use vpn_tunnel_group_id instead.",
							Description:        "Deprecated alias of vpn_tunnel_group_id.",
						},
						"ipsec_profile_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the IPsec profile.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"tunnel_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Numeric tunnel identifier on the wire (e.g. a GRE key), not the NetBox id.",
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
						"termination_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of terminations of the tunnel.",
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

func (d *vpnTunnelsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vpnTunnelsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vpnTunnelsDataSourceModel

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
	var responseDTOs []*netboxapi.VpnTunnelResponseDTO

	params := vpn.NewVpnTunnelsListParams()
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
	var vGroupID []int64
	var vGroupIDn []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
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
	var vEncapsulation []string
	var vEncapsulationn []string
	var vEncapsulationIc []string
	var vEncapsulationNic []string
	var vEncapsulationIe []string
	var vEncapsulationNie []string
	var vEncapsulationIsw []string
	var vEncapsulationNisw []string
	var vEncapsulationIew []string
	var vEncapsulationNiew []string
	var vEncapsulationEmpty *bool
	var vEncapsulationRegex []string
	var vEncapsulationIregex []string
	var vGroup []string
	var vGroupn []string
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vIpsecProfile []string
	var vIpsecProfilen []string
	var vIpsecProfileID []int64
	var vIpsecProfileIDn []int64
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
	var vTunnelID []int64
	var vTunnelIDn []int64
	var vTunnelIDLt []int64
	var vTunnelIDLte []int64
	var vTunnelIDGt []int64
	var vTunnelIDGte []int64
	var vTunnelIDEmpty *bool
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
		case "encapsulation":
			v := value
			vEncapsulation = append(vEncapsulation, v)
		case "encapsulation__n":
			v := value
			vEncapsulationn = append(vEncapsulationn, v)
		case "encapsulation__ic":
			v := value
			vEncapsulationIc = append(vEncapsulationIc, v)
		case "encapsulation__nic":
			v := value
			vEncapsulationNic = append(vEncapsulationNic, v)
		case "encapsulation__ie":
			v := value
			vEncapsulationIe = append(vEncapsulationIe, v)
		case "encapsulation__nie":
			v := value
			vEncapsulationNie = append(vEncapsulationNie, v)
		case "encapsulation__isw":
			v := value
			vEncapsulationIsw = append(vEncapsulationIsw, v)
		case "encapsulation__nisw":
			v := value
			vEncapsulationNisw = append(vEncapsulationNisw, v)
		case "encapsulation__iew":
			v := value
			vEncapsulationIew = append(vEncapsulationIew, v)
		case "encapsulation__niew":
			v := value
			vEncapsulationNiew = append(vEncapsulationNiew, v)
		case "encapsulation__empty":
			if vEncapsulationEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'encapsulation__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'encapsulation__empty' takes a boolean, got %q.", value))
				return
			}
			vEncapsulationEmpty = &v
		case "encapsulation__regex":
			v := value
			vEncapsulationRegex = append(vEncapsulationRegex, v)
		case "encapsulation__iregex":
			v := value
			vEncapsulationIregex = append(vEncapsulationIregex, v)
		case "group":
			v := value
			vGroup = append(vGroup, v)
		case "group__n":
			v := value
			vGroupn = append(vGroupn, v)
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
		case "ipsec_profile":
			v := value
			vIpsecProfile = append(vIpsecProfile, v)
		case "ipsec_profile__n":
			v := value
			vIpsecProfilen = append(vIpsecProfilen, v)
		case "ipsec_profile_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ipsec_profile_id' takes an integer, got %q.", value))
				return
			}
			vIpsecProfileID = append(vIpsecProfileID, v)
		case "ipsec_profile_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ipsec_profile_id__n' takes an integer, got %q.", value))
				return
			}
			vIpsecProfileIDn = append(vIpsecProfileIDn, v)
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
		case "tunnel_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tunnel_id' takes an integer, got %q.", value))
				return
			}
			vTunnelID = append(vTunnelID, v)
		case "tunnel_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tunnel_id__n' takes an integer, got %q.", value))
				return
			}
			vTunnelIDn = append(vTunnelIDn, v)
		case "tunnel_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tunnel_id__lt' takes an integer, got %q.", value))
				return
			}
			vTunnelIDLt = append(vTunnelIDLt, v)
		case "tunnel_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tunnel_id__lte' takes an integer, got %q.", value))
				return
			}
			vTunnelIDLte = append(vTunnelIDLte, v)
		case "tunnel_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tunnel_id__gt' takes an integer, got %q.", value))
				return
			}
			vTunnelIDGt = append(vTunnelIDGt, v)
		case "tunnel_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tunnel_id__gte' takes an integer, got %q.", value))
				return
			}
			vTunnelIDGte = append(vTunnelIDGte, v)
		case "tunnel_id__empty":
			if vTunnelIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'tunnel_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tunnel_id__empty' takes a boolean, got %q.", value))
				return
			}
			vTunnelIDEmpty = &v
		default:
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
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
	if len(vGroupID) > 0 {
		params.SetGroupID(vGroupID)
	}
	if len(vGroupIDn) > 0 {
		params.SetGroupIDn(vGroupIDn)
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
	if len(vEncapsulation) > 0 {
		params.SetEncapsulation(vEncapsulation)
	}
	if len(vEncapsulationn) > 0 {
		params.SetEncapsulationn(vEncapsulationn)
	}
	if len(vEncapsulationIc) > 0 {
		params.SetEncapsulationIc(vEncapsulationIc)
	}
	if len(vEncapsulationNic) > 0 {
		params.SetEncapsulationNic(vEncapsulationNic)
	}
	if len(vEncapsulationIe) > 0 {
		params.SetEncapsulationIe(vEncapsulationIe)
	}
	if len(vEncapsulationNie) > 0 {
		params.SetEncapsulationNie(vEncapsulationNie)
	}
	if len(vEncapsulationIsw) > 0 {
		params.SetEncapsulationIsw(vEncapsulationIsw)
	}
	if len(vEncapsulationNisw) > 0 {
		params.SetEncapsulationNisw(vEncapsulationNisw)
	}
	if len(vEncapsulationIew) > 0 {
		params.SetEncapsulationIew(vEncapsulationIew)
	}
	if len(vEncapsulationNiew) > 0 {
		params.SetEncapsulationNiew(vEncapsulationNiew)
	}
	if vEncapsulationEmpty != nil {
		params.SetEncapsulationEmpty(vEncapsulationEmpty)
	}
	if len(vEncapsulationRegex) > 0 {
		params.SetEncapsulationRegex(vEncapsulationRegex)
	}
	if len(vEncapsulationIregex) > 0 {
		params.SetEncapsulationIregex(vEncapsulationIregex)
	}
	if len(vGroup) > 0 {
		params.SetGroup(vGroup)
	}
	if len(vGroupn) > 0 {
		params.SetGroupn(vGroupn)
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
	if len(vIpsecProfile) > 0 {
		params.SetIpsecProfile(vIpsecProfile)
	}
	if len(vIpsecProfilen) > 0 {
		params.SetIpsecProfilen(vIpsecProfilen)
	}
	if len(vIpsecProfileID) > 0 {
		params.SetIpsecProfileID(vIpsecProfileID)
	}
	if len(vIpsecProfileIDn) > 0 {
		params.SetIpsecProfileIDn(vIpsecProfileIDn)
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
	if len(vTunnelID) > 0 {
		params.SetTunnelID(vTunnelID)
	}
	if len(vTunnelIDn) > 0 {
		params.SetTunnelIDn(vTunnelIDn)
	}
	if len(vTunnelIDLt) > 0 {
		params.SetTunnelIDLt(vTunnelIDLt)
	}
	if len(vTunnelIDLte) > 0 {
		params.SetTunnelIDLte(vTunnelIDLte)
	}
	if len(vTunnelIDGt) > 0 {
		params.SetTunnelIDGt(vTunnelIDGt)
	}
	if len(vTunnelIDGte) > 0 {
		params.SetTunnelIDGte(vTunnelIDGte)
	}
	if vTunnelIDEmpty != nil {
		params.SetTunnelIDEmpty(vTunnelIDEmpty)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Vpn.VpnTunnelsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_vpn_tunnels", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VpnTunnelResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]vpnTunnelResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model vpnTunnelResourceModel
		resp.Diagnostics.Append(flattenVpnTunnel(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: vpnTunnelResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
