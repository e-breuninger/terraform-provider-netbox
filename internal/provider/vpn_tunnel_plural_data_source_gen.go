// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
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
		"group_id":              netboxapi.FilterInt64s(params.SetGroupID),
		"group_id__n":           netboxapi.FilterInt64s(params.SetGroupIDn),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"contact":               netboxapi.FilterInt64s(params.SetContact),
		"contact__n":            netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":         netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":      netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":          netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":       netboxapi.FilterInt64s(params.SetContactRolen),
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
		"encapsulation":         netboxapi.FilterStrings(params.SetEncapsulation),
		"encapsulation__n":      netboxapi.FilterStrings(params.SetEncapsulationn),
		"encapsulation__ic":     netboxapi.FilterStrings(params.SetEncapsulationIc),
		"encapsulation__nic":    netboxapi.FilterStrings(params.SetEncapsulationNic),
		"encapsulation__ie":     netboxapi.FilterStrings(params.SetEncapsulationIe),
		"encapsulation__nie":    netboxapi.FilterStrings(params.SetEncapsulationNie),
		"encapsulation__isw":    netboxapi.FilterStrings(params.SetEncapsulationIsw),
		"encapsulation__nisw":   netboxapi.FilterStrings(params.SetEncapsulationNisw),
		"encapsulation__iew":    netboxapi.FilterStrings(params.SetEncapsulationIew),
		"encapsulation__niew":   netboxapi.FilterStrings(params.SetEncapsulationNiew),
		"encapsulation__empty":  netboxapi.FilterBool(params.SetEncapsulationEmpty),
		"encapsulation__regex":  netboxapi.FilterStrings(params.SetEncapsulationRegex),
		"encapsulation__iregex": netboxapi.FilterStrings(params.SetEncapsulationIregex),
		"group":                 netboxapi.FilterStrings(params.SetGroup),
		"group__n":              netboxapi.FilterStrings(params.SetGroupn),
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"ipsec_profile":         netboxapi.FilterStrings(params.SetIpsecProfile),
		"ipsec_profile__n":      netboxapi.FilterStrings(params.SetIpsecProfilen),
		"ipsec_profile_id":      netboxapi.FilterInt64s(params.SetIpsecProfileID),
		"ipsec_profile_id__n":   netboxapi.FilterInt64s(params.SetIpsecProfileIDn),
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
		"tunnel_id":             netboxapi.FilterInt64s(params.SetTunnelID),
		"tunnel_id__n":          netboxapi.FilterInt64s(params.SetTunnelIDn),
		"tunnel_id__lt":         netboxapi.FilterInt64s(params.SetTunnelIDLt),
		"tunnel_id__lte":        netboxapi.FilterInt64s(params.SetTunnelIDLte),
		"tunnel_id__gt":         netboxapi.FilterInt64s(params.SetTunnelIDGt),
		"tunnel_id__gte":        netboxapi.FilterInt64s(params.SetTunnelIDGte),
		"tunnel_id__empty":      netboxapi.FilterBool(params.SetTunnelIDEmpty),
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
