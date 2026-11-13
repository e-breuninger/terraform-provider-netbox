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
	_ datasource.DataSource              = (*vpnTunnelTerminationsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vpnTunnelTerminationsDataSource)(nil)
)

// NewVpnTunnelTerminationsDataSource returns a new vpn_tunnel_terminations data source, which
// lists vpn_tunnel_termination objects matching its filters.
func NewVpnTunnelTerminationsDataSource() datasource.DataSource {
	return &vpnTunnelTerminationsDataSource{}
}

type vpnTunnelTerminationsDataSource struct {
	client *netboxapi.Client
}

// vpnTunnelTerminationsDataSourceModel is the data source model: the filters, the limit and the matching
// vpn_tunnel_terminations.
type vpnTunnelTerminationsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"vpn_tunnel_terminations"`
}

// vpnTunnelTerminationResourceAttrTypes is the attribute type map of vpnTunnelTerminationResourceModel.
func vpnTunnelTerminationResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                           types.Int64Type,
		"vpn_tunnel_id":                types.Int64Type,
		"role":                         types.StringType,
		"termination_type":             types.StringType,
		"termination_id":               types.Int64Type,
		"device_interface_id":          types.Int64Type,
		"virtual_machine_interface_id": types.Int64Type,
		"outside_ip_address_id":        types.Int64Type,
		"created":                      types.StringType,
		"last_updated":                 types.StringType,
		"url":                          types.StringType,
		"tags":                         types.SetType{ElemType: types.StringType},
		"tags_all":                     types.SetType{ElemType: types.StringType},
		"custom_fields":                types.MapType{ElemType: types.StringType},
	}
}

func (d *vpnTunnelTerminationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_tunnel_terminations"
}

func (d *vpnTunnelTerminationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:Lists vpn_tunnel_termination objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface, interface__n, interface_id, interface_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, outside_ip_id, outside_ip_id__n, q, role, role__empty, role__ic, role__ie, role__iew, role__iregex, role__isw, role__n, role__nic, role__nie, role__niew, role__nisw, role__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, termination_id, termination_id__empty, termination_id__gt, termination_id__gte, termination_id__lt, termination_id__lte, termination_id__n, termination_type, termination_type__n, tunnel, tunnel__n, tunnel_id, tunnel_id__n, vminterface, vminterface__n, vminterface_id, vminterface_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"vpn_tunnel_terminations": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching vpn_tunnel_terminations.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"vpn_tunnel_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tunnel.",
						},
						"role": schema.StringAttribute{
							Computed:    true,
							Description: "Role of the termination in the tunnel topology. One of: peer, hub, spoke.",
						},
						"termination_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the terminating interface. Derived from device_interface_id or virtual_machine_interface_id when one of those is set; set it together with termination_id otherwise. One assignment is required. One of: dcim.interface, virtualization.vminterface.",
						},
						"termination_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating interface (see termination_type).",
						},
						"device_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating device interface (termination_type dcim.interface). Conflicts with virtual_machine_interface_id and with setting the termination_* pair directly.",
						},
						"virtual_machine_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating virtual machine interface (termination_type virtualization.vminterface). Conflicts with device_interface_id and with setting the termination_* pair directly.",
						},
						"outside_ip_address_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the outside IP address of the termination.",
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

func (d *vpnTunnelTerminationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vpnTunnelTerminationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vpnTunnelTerminationsDataSourceModel

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
	var responseDTOs []*netboxapi.VpnTunnelTerminationResponseDTO

	params := vpn.NewVpnTunnelTerminationsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"tunnel_id":             netboxapi.FilterInt64s(params.SetTunnelID),
		"tunnel_id__n":          netboxapi.FilterInt64s(params.SetTunnelIDn),
		"interface_id":          netboxapi.FilterInt64s(params.SetInterfaceID),
		"interface_id__n":       netboxapi.FilterInt64s(params.SetInterfaceIDn),
		"vminterface_id":        netboxapi.FilterInt64s(params.SetVminterfaceID),
		"vminterface_id__n":     netboxapi.FilterInt64s(params.SetVminterfaceIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
		"interface":             netboxapi.FilterStrings(params.SetInterface),
		"interface__n":          netboxapi.FilterStrings(params.SetInterfacen),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"outside_ip_id":         netboxapi.FilterInt64s(params.SetOutsideIPID),
		"outside_ip_id__n":      netboxapi.FilterInt64s(params.SetOutsideIPIDn),
		"q":                     netboxapi.FilterString(params.SetQ),
		"role":                  netboxapi.FilterStrings(params.SetRole),
		"role__n":               netboxapi.FilterStrings(params.SetRolen),
		"role__ic":              netboxapi.FilterStrings(params.SetRoleIc),
		"role__nic":             netboxapi.FilterStrings(params.SetRoleNic),
		"role__ie":              netboxapi.FilterStrings(params.SetRoleIe),
		"role__nie":             netboxapi.FilterStrings(params.SetRoleNie),
		"role__isw":             netboxapi.FilterStrings(params.SetRoleIsw),
		"role__nisw":            netboxapi.FilterStrings(params.SetRoleNisw),
		"role__iew":             netboxapi.FilterStrings(params.SetRoleIew),
		"role__niew":            netboxapi.FilterStrings(params.SetRoleNiew),
		"role__empty":           netboxapi.FilterBool(params.SetRoleEmpty),
		"role__regex":           netboxapi.FilterStrings(params.SetRoleRegex),
		"role__iregex":          netboxapi.FilterStrings(params.SetRoleIregex),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"termination_id":        netboxapi.FilterInt64s(params.SetTerminationID),
		"termination_id__n":     netboxapi.FilterInt64s(params.SetTerminationIDn),
		"termination_id__lt":    netboxapi.FilterInt64s(params.SetTerminationIDLt),
		"termination_id__lte":   netboxapi.FilterInt64s(params.SetTerminationIDLte),
		"termination_id__gt":    netboxapi.FilterInt64s(params.SetTerminationIDGt),
		"termination_id__gte":   netboxapi.FilterInt64s(params.SetTerminationIDGte),
		"termination_id__empty": netboxapi.FilterBool(params.SetTerminationIDEmpty),
		"termination_type":      netboxapi.FilterStrings(params.SetTerminationType),
		"termination_type__n":   netboxapi.FilterStrings(params.SetTerminationTypen),
		"tunnel":                netboxapi.FilterStrings(params.SetTunnel),
		"tunnel__n":             netboxapi.FilterStrings(params.SetTunneln),
		"vminterface":           netboxapi.FilterStrings(params.SetVminterface),
		"vminterface__n":        netboxapi.FilterStrings(params.SetVminterfacen),
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
		res, err := d.client.Vpn.VpnTunnelTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_vpn_tunnel_terminations", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VpnTunnelTerminationResponseDTOFromGoNetbox(goNetboxModel))
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
	resourceWithHooks := &vpnTunnelTerminationResource{client: d.client}
	items := make([]vpnTunnelTerminationResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model vpnTunnelTerminationResourceModel
		resp.Diagnostics.Append(flattenVpnTunnelTermination(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: vpnTunnelTerminationResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
