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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vTunnelID []int64
	var vTunnelIDn []int64
	var vInterfaceID []int64
	var vInterfaceIDn []int64
	var vVminterfaceID []int64
	var vVminterfaceIDn []int64
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
	var vInterface []string
	var vInterfacen []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vOutsideIPID []int64
	var vOutsideIPIDn []int64
	var vQ *string
	var vRole []string
	var vRolen []string
	var vRoleIc []string
	var vRoleNic []string
	var vRoleIe []string
	var vRoleNie []string
	var vRoleIsw []string
	var vRoleNisw []string
	var vRoleIew []string
	var vRoleNiew []string
	var vRoleEmpty *bool
	var vRoleRegex []string
	var vRoleIregex []string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTerminationID []int64
	var vTerminationIDn []int64
	var vTerminationIDLt []int64
	var vTerminationIDLte []int64
	var vTerminationIDGt []int64
	var vTerminationIDGte []int64
	var vTerminationIDEmpty *bool
	var vTerminationType []string
	var vTerminationTypen []string
	var vTunnel []string
	var vTunneln []string
	var vVminterface []string
	var vVminterfacen []string
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
		case "interface_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_id' takes an integer, got %q.", value))
				return
			}
			vInterfaceID = append(vInterfaceID, v)
		case "interface_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_id__n' takes an integer, got %q.", value))
				return
			}
			vInterfaceIDn = append(vInterfaceIDn, v)
		case "vminterface_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vminterface_id' takes an integer, got %q.", value))
				return
			}
			vVminterfaceID = append(vVminterfaceID, v)
		case "vminterface_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vminterface_id__n' takes an integer, got %q.", value))
				return
			}
			vVminterfaceIDn = append(vVminterfaceIDn, v)
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
		case "interface":
			v := value
			vInterface = append(vInterface, v)
		case "interface__n":
			v := value
			vInterfacen = append(vInterfacen, v)
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
		case "outside_ip_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outside_ip_id' takes an integer, got %q.", value))
				return
			}
			vOutsideIPID = append(vOutsideIPID, v)
		case "outside_ip_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outside_ip_id__n' takes an integer, got %q.", value))
				return
			}
			vOutsideIPIDn = append(vOutsideIPIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "role":
			v := value
			vRole = append(vRole, v)
		case "role__n":
			v := value
			vRolen = append(vRolen, v)
		case "role__ic":
			v := value
			vRoleIc = append(vRoleIc, v)
		case "role__nic":
			v := value
			vRoleNic = append(vRoleNic, v)
		case "role__ie":
			v := value
			vRoleIe = append(vRoleIe, v)
		case "role__nie":
			v := value
			vRoleNie = append(vRoleNie, v)
		case "role__isw":
			v := value
			vRoleIsw = append(vRoleIsw, v)
		case "role__nisw":
			v := value
			vRoleNisw = append(vRoleNisw, v)
		case "role__iew":
			v := value
			vRoleIew = append(vRoleIew, v)
		case "role__niew":
			v := value
			vRoleNiew = append(vRoleNiew, v)
		case "role__empty":
			if vRoleEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'role__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role__empty' takes a boolean, got %q.", value))
				return
			}
			vRoleEmpty = &v
		case "role__regex":
			v := value
			vRoleRegex = append(vRoleRegex, v)
		case "role__iregex":
			v := value
			vRoleIregex = append(vRoleIregex, v)
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
		case "termination_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id' takes an integer, got %q.", value))
				return
			}
			vTerminationID = append(vTerminationID, v)
		case "termination_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__n' takes an integer, got %q.", value))
				return
			}
			vTerminationIDn = append(vTerminationIDn, v)
		case "termination_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__lt' takes an integer, got %q.", value))
				return
			}
			vTerminationIDLt = append(vTerminationIDLt, v)
		case "termination_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__lte' takes an integer, got %q.", value))
				return
			}
			vTerminationIDLte = append(vTerminationIDLte, v)
		case "termination_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__gt' takes an integer, got %q.", value))
				return
			}
			vTerminationIDGt = append(vTerminationIDGt, v)
		case "termination_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__gte' takes an integer, got %q.", value))
				return
			}
			vTerminationIDGte = append(vTerminationIDGte, v)
		case "termination_id__empty":
			if vTerminationIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'termination_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__empty' takes a boolean, got %q.", value))
				return
			}
			vTerminationIDEmpty = &v
		case "termination_type":
			v := value
			vTerminationType = append(vTerminationType, v)
		case "termination_type__n":
			v := value
			vTerminationTypen = append(vTerminationTypen, v)
		case "tunnel":
			v := value
			vTunnel = append(vTunnel, v)
		case "tunnel__n":
			v := value
			vTunneln = append(vTunneln, v)
		case "vminterface":
			v := value
			vVminterface = append(vVminterface, v)
		case "vminterface__n":
			v := value
			vVminterfacen = append(vVminterfacen, v)
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
	if len(vTunnelID) > 0 {
		params.SetTunnelID(vTunnelID)
	}
	if len(vTunnelIDn) > 0 {
		params.SetTunnelIDn(vTunnelIDn)
	}
	if len(vInterfaceID) > 0 {
		params.SetInterfaceID(vInterfaceID)
	}
	if len(vInterfaceIDn) > 0 {
		params.SetInterfaceIDn(vInterfaceIDn)
	}
	if len(vVminterfaceID) > 0 {
		params.SetVminterfaceID(vVminterfaceID)
	}
	if len(vVminterfaceIDn) > 0 {
		params.SetVminterfaceIDn(vVminterfaceIDn)
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
	if len(vInterface) > 0 {
		params.SetInterface(vInterface)
	}
	if len(vInterfacen) > 0 {
		params.SetInterfacen(vInterfacen)
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
	if len(vOutsideIPID) > 0 {
		params.SetOutsideIPID(vOutsideIPID)
	}
	if len(vOutsideIPIDn) > 0 {
		params.SetOutsideIPIDn(vOutsideIPIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRole) > 0 {
		params.SetRole(vRole)
	}
	if len(vRolen) > 0 {
		params.SetRolen(vRolen)
	}
	if len(vRoleIc) > 0 {
		params.SetRoleIc(vRoleIc)
	}
	if len(vRoleNic) > 0 {
		params.SetRoleNic(vRoleNic)
	}
	if len(vRoleIe) > 0 {
		params.SetRoleIe(vRoleIe)
	}
	if len(vRoleNie) > 0 {
		params.SetRoleNie(vRoleNie)
	}
	if len(vRoleIsw) > 0 {
		params.SetRoleIsw(vRoleIsw)
	}
	if len(vRoleNisw) > 0 {
		params.SetRoleNisw(vRoleNisw)
	}
	if len(vRoleIew) > 0 {
		params.SetRoleIew(vRoleIew)
	}
	if len(vRoleNiew) > 0 {
		params.SetRoleNiew(vRoleNiew)
	}
	if vRoleEmpty != nil {
		params.SetRoleEmpty(vRoleEmpty)
	}
	if len(vRoleRegex) > 0 {
		params.SetRoleRegex(vRoleRegex)
	}
	if len(vRoleIregex) > 0 {
		params.SetRoleIregex(vRoleIregex)
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
	if len(vTerminationID) > 0 {
		params.SetTerminationID(vTerminationID)
	}
	if len(vTerminationIDn) > 0 {
		params.SetTerminationIDn(vTerminationIDn)
	}
	if len(vTerminationIDLt) > 0 {
		params.SetTerminationIDLt(vTerminationIDLt)
	}
	if len(vTerminationIDLte) > 0 {
		params.SetTerminationIDLte(vTerminationIDLte)
	}
	if len(vTerminationIDGt) > 0 {
		params.SetTerminationIDGt(vTerminationIDGt)
	}
	if len(vTerminationIDGte) > 0 {
		params.SetTerminationIDGte(vTerminationIDGte)
	}
	if vTerminationIDEmpty != nil {
		params.SetTerminationIDEmpty(vTerminationIDEmpty)
	}
	if len(vTerminationType) > 0 {
		params.SetTerminationType(vTerminationType)
	}
	if len(vTerminationTypen) > 0 {
		params.SetTerminationTypen(vTerminationTypen)
	}
	if len(vTunnel) > 0 {
		params.SetTunnel(vTunnel)
	}
	if len(vTunneln) > 0 {
		params.SetTunneln(vTunneln)
	}
	if len(vVminterface) > 0 {
		params.SetVminterface(vVminterface)
	}
	if len(vVminterfacen) > 0 {
		params.SetVminterfacen(vVminterfacen)
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
