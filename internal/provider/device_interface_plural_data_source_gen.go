// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*deviceInterfacesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceInterfacesDataSource)(nil)
)

// NewDeviceInterfacesDataSource returns a new device_interfaces data source, which
// lists device_interface objects matching its filters.
func NewDeviceInterfacesDataSource() datasource.DataSource {
	return &deviceInterfacesDataSource{}
}

type deviceInterfacesDataSource struct {
	client *netboxapi.Client
}

// deviceInterfacesDataSourceModel is the data source model: the filters, the limit and the matching
// device_interfaces.
type deviceInterfacesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"device_interfaces"`
}

// deviceInterfaceResourceAttrTypes is the attribute type map of deviceInterfaceResourceModel.
func deviceInterfaceResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                         types.Int64Type,
		"device_id":                  types.Int64Type,
		"name":                       types.StringType,
		"type":                       types.StringType,
		"label":                      types.StringType,
		"enabled":                    types.BoolType,
		"mgmt_only":                  types.BoolType,
		"mgmtonly":                   types.BoolType,
		"mark_connected":             types.BoolType,
		"mtu":                        types.Int64Type,
		"speed":                      types.Int64Type,
		"duplex":                     types.StringType,
		"poe_mode":                   types.StringType,
		"poe_type":                   types.StringType,
		"rf_role":                    types.StringType,
		"rf_channel":                 types.StringType,
		"rf_channel_frequency":       types.Float64Type,
		"rf_channel_width":           types.Float64Type,
		"tx_power":                   types.Int64Type,
		"wwn":                        types.StringType,
		"mode":                       types.StringType,
		"untagged_vlan_id":           types.Int64Type,
		"untagged_vlan":              types.Int64Type,
		"tagged_vlan_ids":            types.SetType{ElemType: types.Int64Type},
		"tagged_vlans":               types.SetType{ElemType: types.Int64Type},
		"qinq_svlan_id":              types.Int64Type,
		"vlan_translation_policy_id": types.Int64Type,
		"module_id":                  types.Int64Type,
		"lag_device_interface_id":    types.Int64Type,
		"parent_device_interface_id": types.Int64Type,
		"bridge_id":                  types.Int64Type,
		"vrf_id":                     types.Int64Type,
		"vdc_ids":                    types.SetType{ElemType: types.Int64Type},
		"wireless_lan_ids":           types.SetType{ElemType: types.Int64Type},
		"description":                types.StringType,
		"owner_id":                   types.Int64Type,
		"created":                    types.StringType,
		"last_updated":               types.StringType,
		"url":                        types.StringType,
		"tags":                       types.SetType{ElemType: types.StringType},
		"tags_all":                   types.SetType{ElemType: types.StringType},
		"custom_fields":              types.MapType{ElemType: types.StringType},
	}
}

func (d *deviceInterfacesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_interfaces"
}

func (d *deviceInterfacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists device_interface objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: bridge_id, bridge_id__n, cable_connector, cable_connector__empty, cable_connector__gt, cable_connector__gte, cable_connector__lt, cable_connector__lte, cable_connector__n, cable_end, cable_end__empty, cable_end__ic, cable_end__ie, cable_end__iew, cable_end__iregex, cable_end__isw, cable_end__n, cable_end__nic, cable_end__nie, cable_end__niew, cable_end__nisw, cable_end__regex, cable_id, cable_id__n, cabled, connected, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device__n, device_id, device_id__n, device_role, device_role__n, device_role_id, device_role_id__n, device_status, device_status__empty, device_status__ic, device_status__ie, device_status__iew, device_status__iregex, device_status__isw, device_status__n, device_status__nic, device_status__nie, device_status__niew, device_status__nisw, device_status__regex, device_type, device_type__n, device_type_id, device_type_id__n, duplex, duplex__empty, duplex__ic, duplex__ie, duplex__iew, duplex__iregex, duplex__isw, duplex__n, duplex__nic, duplex__nie, duplex__niew, duplex__nisw, duplex__regex, enabled, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, kind, l2vpn, l2vpn__n, l2vpn_id, l2vpn_id__n, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, lag_id, lag_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, mac_address, mac_address__ic, mac_address__ie, mac_address__iew, mac_address__iregex, mac_address__isw, mac_address__n, mac_address__nic, mac_address__nie, mac_address__niew, mac_address__nisw, mac_address__regex, mark_connected, mgmt_only, mode, mode__empty, mode__ic, mode__ie, mode__iew, mode__iregex, mode__isw, mode__n, mode__nic, mode__nie, mode__niew, mode__nisw, mode__regex, module_id, module_id__n, mtu, mtu__empty, mtu__gt, mtu__gte, mtu__lt, mtu__lte, mtu__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, occupied, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, parent_id, parent_id__n, poe_mode, poe_mode__empty, poe_mode__ic, poe_mode__ie, poe_mode__iew, poe_mode__iregex, poe_mode__isw, poe_mode__n, poe_mode__nic, poe_mode__nie, poe_mode__niew, poe_mode__nisw, poe_mode__regex, poe_type, poe_type__empty, poe_type__ic, poe_type__ie, poe_type__iew, poe_type__iregex, poe_type__isw, poe_type__n, poe_type__nic, poe_type__nie, poe_type__niew, poe_type__nisw, poe_type__regex, primary_mac_address, primary_mac_address__n, primary_mac_address_id, primary_mac_address_id__n, q, rack, rack__n, rack_id, rack_id__n, region, region__n, region_id, region_id__n, rf_channel, rf_channel__empty, rf_channel__ic, rf_channel__ie, rf_channel__iew, rf_channel__iregex, rf_channel__isw, rf_channel__n, rf_channel__nic, rf_channel__nie, rf_channel__niew, rf_channel__nisw, rf_channel__regex, rf_channel_frequency, rf_channel_frequency__empty, rf_channel_frequency__gt, rf_channel_frequency__gte, rf_channel_frequency__lt, rf_channel_frequency__lte, rf_channel_frequency__n, rf_channel_width, rf_channel_width__empty, rf_channel_width__gt, rf_channel_width__gte, rf_channel_width__lt, rf_channel_width__lte, rf_channel_width__n, rf_role, rf_role__empty, rf_role__ic, rf_role__ie, rf_role__iew, rf_role__iregex, rf_role__isw, rf_role__n, rf_role__nic, rf_role__nie, rf_role__niew, rf_role__nisw, rf_role__regex, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, speed, speed__empty, speed__gt, speed__gte, speed__lt, speed__lte, speed__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_id, tenant_id__n, tx_power, tx_power__empty, tx_power__gt, tx_power__gte, tx_power__lt, tx_power__lte, tx_power__n, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex, vdc, vdc__n, vdc_id, vdc_id__n, vdc_identifier, vdc_identifier__n, virtual_chassis, virtual_chassis__n, virtual_chassis_id, virtual_chassis_id__n, virtual_chassis_member, virtual_chassis_member_id, virtual_chassis_member_or_master, virtual_chassis_member_or_master_id, virtual_circuit_id, virtual_circuit_id__n, virtual_circuit_termination_id, virtual_circuit_termination_id__n, vlan, vlan_id, vlan_translation_policy, vlan_translation_policy__n, vlan_translation_policy_id, vlan_translation_policy_id__n, vrf, vrf__n, vrf_id, vrf_id__n, wireless_lan_id, wireless_lan_id__n, wireless_link_id, wireless_link_id__n, wwn, wwn__ic, wwn__ie, wwn__iew, wwn__iregex, wwn__isw, wwn__n, wwn__nic, wwn__nie, wwn__niew, wwn__nisw, wwn__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"device_interfaces": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching device_interfaces.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"device_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Interface type as its NetBox slug, e.g. virtual, lag, bridge, 1000base-t, 10gbase-x-sfpp (any of NetBox's interface type choices).",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Physical label.",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the interface is enabled.",
						},
						"mgmt_only": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the interface is used for out-of-band management only.",
						},
						"mgmtonly": schema.BoolAttribute{
							Computed:           true,
							DeprecationMessage: "Use mgmt_only instead.",
							Description:        "Deprecated alias of mgmt_only.",
						},
						"mark_connected": schema.BoolAttribute{
							Computed:    true,
							Description: "Treat as if a cable is connected.",
						},
						"mtu": schema.Int64Attribute{
							Computed:    true,
							Description: "Must be a whole number.",
						},
						"speed": schema.Int64Attribute{
							Computed:    true,
							Description: "Speed in kbps. Must be a whole number.",
						},
						"duplex": schema.StringAttribute{
							Computed:    true,
							Description: "One of: half, full, auto.",
						},
						"poe_mode": schema.StringAttribute{
							Computed:    true,
							Description: "Power over Ethernet mode. One of: pd, pse.",
						},
						"poe_type": schema.StringAttribute{
							Computed:    true,
							Description: "Power over Ethernet type as its NetBox slug, e.g. type1-ieee802.3af, type3-ieee802.3bt, passive-24v-2pair (any of NetBox's PoE type choices).",
						},
						"rf_role": schema.StringAttribute{
							Computed:    true,
							Description: "Wireless role. One of: ap, station.",
						},
						"rf_channel": schema.StringAttribute{
							Computed:    true,
							Description: "Wireless channel as its NetBox slug, e.g. 2.4g-1-2412-22, 5g-36-5180-20 (any of NetBox's wireless channel choices).",
						},
						"rf_channel_frequency": schema.Float64Attribute{
							Computed:    true,
							Description: "Channel frequency in MHz. NetBox derives it from rf_channel when that is set.",
						},
						"rf_channel_width": schema.Float64Attribute{
							Computed:    true,
							Description: "Channel width in MHz. NetBox derives it from rf_channel when that is set.",
						},
						"tx_power": schema.Int64Attribute{
							Computed:    true,
							Description: "Transmit power in dBm. Must be a whole number.",
						},
						"wwn": schema.StringAttribute{
							Computed:    true,
							Description: "64-bit World Wide Name.",
						},
						"mode": schema.StringAttribute{
							Computed:    true,
							Description: "802.1Q tagging mode. One of: access, tagged, tagged-all, q-in-q.",
						},
						"untagged_vlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the untagged VLAN.",
						},
						"untagged_vlan": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use untagged_vlan_id instead.",
							Description:        "Deprecated alias of untagged_vlan_id.",
						},
						"tagged_vlan_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the tagged VLANs.",
						},
						"tagged_vlans": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use tagged_vlan_ids instead.",
							Description:        "Deprecated alias of tagged_vlan_ids.",
						},
						"qinq_svlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the Q-in-Q service VLAN.",
						},
						"vlan_translation_policy_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VLAN translation policy.",
						},
						"module_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the installed module this interface belongs to.",
						},
						"lag_device_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the parent LAG interface.",
						},
						"parent_device_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the parent interface.",
						},
						"bridge_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the bridge interface.",
						},
						"vrf_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VRF.",
						},
						"vdc_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the virtual device contexts.",
						},
						"wireless_lan_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the wireless LANs.",
						},
						"description": schema.StringAttribute{
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

func (d *deviceInterfacesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceInterfacesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceInterfacesDataSourceModel

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
	var responseDTOs []*netboxapi.DeviceInterfaceResponseDTO

	params := dcim.NewDcimInterfacesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                                  netboxapi.FilterInt64s(params.SetID),
		"id__n":                               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                           netboxapi.FilterBool(params.SetIDEmpty),
		"name":                                netboxapi.FilterStrings(params.SetName),
		"name__n":                             netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                            netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                           netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                            netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                           netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                           netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                          netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                           netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                          netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":                         netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":                         netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":                        netboxapi.FilterStrings(params.SetNameIregex),
		"device_id":                           netboxapi.FilterInt64s(params.SetDeviceID),
		"device_id__n":                        netboxapi.FilterInt64s(params.SetDeviceIDn),
		"type":                                netboxapi.FilterStrings(params.SetType),
		"type__n":                             netboxapi.FilterStrings(params.SetTypen),
		"type__ic":                            netboxapi.FilterStrings(params.SetTypeIc),
		"type__nic":                           netboxapi.FilterStrings(params.SetTypeNic),
		"type__ie":                            netboxapi.FilterStrings(params.SetTypeIe),
		"type__nie":                           netboxapi.FilterStrings(params.SetTypeNie),
		"type__isw":                           netboxapi.FilterStrings(params.SetTypeIsw),
		"type__nisw":                          netboxapi.FilterStrings(params.SetTypeNisw),
		"type__iew":                           netboxapi.FilterStrings(params.SetTypeIew),
		"type__niew":                          netboxapi.FilterStrings(params.SetTypeNiew),
		"type__empty":                         netboxapi.FilterBool(params.SetTypeEmpty),
		"type__regex":                         netboxapi.FilterStrings(params.SetTypeRegex),
		"type__iregex":                        netboxapi.FilterStrings(params.SetTypeIregex),
		"mode":                                netboxapi.FilterStrings(params.SetMode),
		"mode__n":                             netboxapi.FilterStrings(params.SetModen),
		"mode__ic":                            netboxapi.FilterStrings(params.SetModeIc),
		"mode__nic":                           netboxapi.FilterStrings(params.SetModeNic),
		"mode__ie":                            netboxapi.FilterStrings(params.SetModeIe),
		"mode__nie":                           netboxapi.FilterStrings(params.SetModeNie),
		"mode__isw":                           netboxapi.FilterStrings(params.SetModeIsw),
		"mode__nisw":                          netboxapi.FilterStrings(params.SetModeNisw),
		"mode__iew":                           netboxapi.FilterStrings(params.SetModeIew),
		"mode__niew":                          netboxapi.FilterStrings(params.SetModeNiew),
		"mode__empty":                         netboxapi.FilterBool(params.SetModeEmpty),
		"mode__regex":                         netboxapi.FilterStrings(params.SetModeRegex),
		"mode__iregex":                        netboxapi.FilterStrings(params.SetModeIregex),
		"enabled":                             netboxapi.FilterBool(params.SetEnabled),
		"lag_id":                              netboxapi.FilterInt64s(params.SetLagID),
		"lag_id__n":                           netboxapi.FilterInt64s(params.SetLagIDn),
		"mac_address":                         netboxapi.FilterStrings(params.SetMacAddress),
		"mac_address__n":                      netboxapi.FilterStrings(params.SetMacAddressn),
		"mac_address__ic":                     netboxapi.FilterStrings(params.SetMacAddressIc),
		"mac_address__nic":                    netboxapi.FilterStrings(params.SetMacAddressNic),
		"mac_address__ie":                     netboxapi.FilterStrings(params.SetMacAddressIe),
		"mac_address__nie":                    netboxapi.FilterStrings(params.SetMacAddressNie),
		"mac_address__isw":                    netboxapi.FilterStrings(params.SetMacAddressIsw),
		"mac_address__nisw":                   netboxapi.FilterStrings(params.SetMacAddressNisw),
		"mac_address__iew":                    netboxapi.FilterStrings(params.SetMacAddressIew),
		"mac_address__niew":                   netboxapi.FilterStrings(params.SetMacAddressNiew),
		"mac_address__regex":                  netboxapi.FilterStrings(params.SetMacAddressRegex),
		"mac_address__iregex":                 netboxapi.FilterStrings(params.SetMacAddressIregex),
		"owner_id":                            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                         netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":                              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                            netboxapi.FilterStrings(params.SetTagAny),
		"bridge_id":                           netboxapi.FilterInt64s(params.SetBridgeID),
		"bridge_id__n":                        netboxapi.FilterInt64s(params.SetBridgeIDn),
		"cable_connector":                     netboxapi.FilterInt64s(params.SetCableConnector),
		"cable_connector__n":                  netboxapi.FilterInt64s(params.SetCableConnectorn),
		"cable_connector__lt":                 netboxapi.FilterInt64s(params.SetCableConnectorLt),
		"cable_connector__lte":                netboxapi.FilterInt64s(params.SetCableConnectorLte),
		"cable_connector__gt":                 netboxapi.FilterInt64s(params.SetCableConnectorGt),
		"cable_connector__gte":                netboxapi.FilterInt64s(params.SetCableConnectorGte),
		"cable_connector__empty":              netboxapi.FilterBool(params.SetCableConnectorEmpty),
		"cable_end":                           netboxapi.FilterString(params.SetCableEnd),
		"cable_end__n":                        netboxapi.FilterString(params.SetCableEndn),
		"cable_end__ic":                       netboxapi.FilterStrings(params.SetCableEndIc),
		"cable_end__nic":                      netboxapi.FilterStrings(params.SetCableEndNic),
		"cable_end__ie":                       netboxapi.FilterStrings(params.SetCableEndIe),
		"cable_end__nie":                      netboxapi.FilterStrings(params.SetCableEndNie),
		"cable_end__isw":                      netboxapi.FilterStrings(params.SetCableEndIsw),
		"cable_end__nisw":                     netboxapi.FilterStrings(params.SetCableEndNisw),
		"cable_end__iew":                      netboxapi.FilterStrings(params.SetCableEndIew),
		"cable_end__niew":                     netboxapi.FilterStrings(params.SetCableEndNiew),
		"cable_end__empty":                    netboxapi.FilterBool(params.SetCableEndEmpty),
		"cable_end__regex":                    netboxapi.FilterStrings(params.SetCableEndRegex),
		"cable_end__iregex":                   netboxapi.FilterStrings(params.SetCableEndIregex),
		"cable_id":                            netboxapi.FilterInt64s(params.SetCableID),
		"cable_id__n":                         netboxapi.FilterInt64s(params.SetCableIDn),
		"cabled":                              netboxapi.FilterBool(params.SetCabled),
		"connected":                           netboxapi.FilterBool(params.SetConnected),
		"created":                             netboxapi.FilterStrings(params.SetCreated),
		"created__n":                          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":                        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":                        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":                      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":                         netboxapi.FilterStrings(params.SetDescription),
		"description__n":                      netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":                     netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":                    netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":                     netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":                    netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":                    netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":                   netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":                    netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":                   netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":                  netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":                  netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":                 netboxapi.FilterStrings(params.SetDescriptionIregex),
		"device":                              netboxapi.FilterStrings(params.SetDevice),
		"device__n":                           netboxapi.FilterStrings(params.SetDevicen),
		"device_role":                         netboxapi.FilterStrings(params.SetDeviceRole),
		"device_role__n":                      netboxapi.FilterStrings(params.SetDeviceRolen),
		"device_role_id":                      netboxapi.FilterInt64s(params.SetDeviceRoleID),
		"device_role_id__n":                   netboxapi.FilterInt64s(params.SetDeviceRoleIDn),
		"device_status":                       netboxapi.FilterStrings(params.SetDeviceStatus),
		"device_status__n":                    netboxapi.FilterStrings(params.SetDeviceStatusn),
		"device_status__ic":                   netboxapi.FilterStrings(params.SetDeviceStatusIc),
		"device_status__nic":                  netboxapi.FilterStrings(params.SetDeviceStatusNic),
		"device_status__ie":                   netboxapi.FilterStrings(params.SetDeviceStatusIe),
		"device_status__nie":                  netboxapi.FilterStrings(params.SetDeviceStatusNie),
		"device_status__isw":                  netboxapi.FilterStrings(params.SetDeviceStatusIsw),
		"device_status__nisw":                 netboxapi.FilterStrings(params.SetDeviceStatusNisw),
		"device_status__iew":                  netboxapi.FilterStrings(params.SetDeviceStatusIew),
		"device_status__niew":                 netboxapi.FilterStrings(params.SetDeviceStatusNiew),
		"device_status__empty":                netboxapi.FilterBool(params.SetDeviceStatusEmpty),
		"device_status__regex":                netboxapi.FilterStrings(params.SetDeviceStatusRegex),
		"device_status__iregex":               netboxapi.FilterStrings(params.SetDeviceStatusIregex),
		"device_type":                         netboxapi.FilterStrings(params.SetDeviceType),
		"device_type__n":                      netboxapi.FilterStrings(params.SetDeviceTypen),
		"device_type_id":                      netboxapi.FilterInt64s(params.SetDeviceTypeID),
		"device_type_id__n":                   netboxapi.FilterInt64s(params.SetDeviceTypeIDn),
		"duplex":                              netboxapi.FilterStrings(params.SetDuplex),
		"duplex__n":                           netboxapi.FilterStrings(params.SetDuplexn),
		"duplex__ic":                          netboxapi.FilterStrings(params.SetDuplexIc),
		"duplex__nic":                         netboxapi.FilterStrings(params.SetDuplexNic),
		"duplex__ie":                          netboxapi.FilterStrings(params.SetDuplexIe),
		"duplex__nie":                         netboxapi.FilterStrings(params.SetDuplexNie),
		"duplex__isw":                         netboxapi.FilterStrings(params.SetDuplexIsw),
		"duplex__nisw":                        netboxapi.FilterStrings(params.SetDuplexNisw),
		"duplex__iew":                         netboxapi.FilterStrings(params.SetDuplexIew),
		"duplex__niew":                        netboxapi.FilterStrings(params.SetDuplexNiew),
		"duplex__empty":                       netboxapi.FilterBool(params.SetDuplexEmpty),
		"duplex__regex":                       netboxapi.FilterStrings(params.SetDuplexRegex),
		"duplex__iregex":                      netboxapi.FilterStrings(params.SetDuplexIregex),
		"kind":                                netboxapi.FilterString(params.SetKind),
		"l2vpn":                               netboxapi.FilterInt64s(params.SetL2vpn),
		"l2vpn__n":                            netboxapi.FilterInt64s(params.SetL2vpnn),
		"l2vpn_id":                            netboxapi.FilterInt64s(params.SetL2vpnID),
		"l2vpn_id__n":                         netboxapi.FilterInt64s(params.SetL2vpnIDn),
		"label":                               netboxapi.FilterStrings(params.SetLabel),
		"label__n":                            netboxapi.FilterStrings(params.SetLabeln),
		"label__ic":                           netboxapi.FilterStrings(params.SetLabelIc),
		"label__nic":                          netboxapi.FilterStrings(params.SetLabelNic),
		"label__ie":                           netboxapi.FilterStrings(params.SetLabelIe),
		"label__nie":                          netboxapi.FilterStrings(params.SetLabelNie),
		"label__isw":                          netboxapi.FilterStrings(params.SetLabelIsw),
		"label__nisw":                         netboxapi.FilterStrings(params.SetLabelNisw),
		"label__iew":                          netboxapi.FilterStrings(params.SetLabelIew),
		"label__niew":                         netboxapi.FilterStrings(params.SetLabelNiew),
		"label__empty":                        netboxapi.FilterBool(params.SetLabelEmpty),
		"label__regex":                        netboxapi.FilterStrings(params.SetLabelRegex),
		"label__iregex":                       netboxapi.FilterStrings(params.SetLabelIregex),
		"last_updated":                        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":                     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":                    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":                   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":                    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":                   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":                 netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":                            netboxapi.FilterStrings(params.SetLocation),
		"location__n":                         netboxapi.FilterStrings(params.SetLocationn),
		"location_id":                         netboxapi.FilterInt64s(params.SetLocationID),
		"location_id__n":                      netboxapi.FilterInt64s(params.SetLocationIDn),
		"mark_connected":                      netboxapi.FilterBool(params.SetMarkConnected),
		"mgmt_only":                           netboxapi.FilterBool(params.SetMgmtOnly),
		"module_id":                           netboxapi.FilterInt64s(params.SetModuleID),
		"module_id__n":                        netboxapi.FilterInt64s(params.SetModuleIDn),
		"mtu":                                 netboxapi.FilterInt64s(params.SetMtu),
		"mtu__n":                              netboxapi.FilterInt64s(params.SetMtun),
		"mtu__lt":                             netboxapi.FilterInt64s(params.SetMtuLt),
		"mtu__lte":                            netboxapi.FilterInt64s(params.SetMtuLte),
		"mtu__gt":                             netboxapi.FilterInt64s(params.SetMtuGt),
		"mtu__gte":                            netboxapi.FilterInt64s(params.SetMtuGte),
		"mtu__empty":                          netboxapi.FilterBool(params.SetMtuEmpty),
		"occupied":                            netboxapi.FilterBool(params.SetOccupied),
		"owner":                               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":                      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":                      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":                   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"parent_id":                           netboxapi.FilterInt64s(params.SetParentID),
		"parent_id__n":                        netboxapi.FilterInt64s(params.SetParentIDn),
		"poe_mode":                            netboxapi.FilterStrings(params.SetPoeMode),
		"poe_mode__n":                         netboxapi.FilterStrings(params.SetPoeModen),
		"poe_mode__ic":                        netboxapi.FilterStrings(params.SetPoeModeIc),
		"poe_mode__nic":                       netboxapi.FilterStrings(params.SetPoeModeNic),
		"poe_mode__ie":                        netboxapi.FilterStrings(params.SetPoeModeIe),
		"poe_mode__nie":                       netboxapi.FilterStrings(params.SetPoeModeNie),
		"poe_mode__isw":                       netboxapi.FilterStrings(params.SetPoeModeIsw),
		"poe_mode__nisw":                      netboxapi.FilterStrings(params.SetPoeModeNisw),
		"poe_mode__iew":                       netboxapi.FilterStrings(params.SetPoeModeIew),
		"poe_mode__niew":                      netboxapi.FilterStrings(params.SetPoeModeNiew),
		"poe_mode__empty":                     netboxapi.FilterBool(params.SetPoeModeEmpty),
		"poe_mode__regex":                     netboxapi.FilterStrings(params.SetPoeModeRegex),
		"poe_mode__iregex":                    netboxapi.FilterStrings(params.SetPoeModeIregex),
		"poe_type":                            netboxapi.FilterStrings(params.SetPoeType),
		"poe_type__n":                         netboxapi.FilterStrings(params.SetPoeTypen),
		"poe_type__ic":                        netboxapi.FilterStrings(params.SetPoeTypeIc),
		"poe_type__nic":                       netboxapi.FilterStrings(params.SetPoeTypeNic),
		"poe_type__ie":                        netboxapi.FilterStrings(params.SetPoeTypeIe),
		"poe_type__nie":                       netboxapi.FilterStrings(params.SetPoeTypeNie),
		"poe_type__isw":                       netboxapi.FilterStrings(params.SetPoeTypeIsw),
		"poe_type__nisw":                      netboxapi.FilterStrings(params.SetPoeTypeNisw),
		"poe_type__iew":                       netboxapi.FilterStrings(params.SetPoeTypeIew),
		"poe_type__niew":                      netboxapi.FilterStrings(params.SetPoeTypeNiew),
		"poe_type__empty":                     netboxapi.FilterBool(params.SetPoeTypeEmpty),
		"poe_type__regex":                     netboxapi.FilterStrings(params.SetPoeTypeRegex),
		"poe_type__iregex":                    netboxapi.FilterStrings(params.SetPoeTypeIregex),
		"primary_mac_address":                 netboxapi.FilterStrings(params.SetPrimaryMacAddress),
		"primary_mac_address__n":              netboxapi.FilterStrings(params.SetPrimaryMacAddressn),
		"primary_mac_address_id":              netboxapi.FilterInt64s(params.SetPrimaryMacAddressID),
		"primary_mac_address_id__n":           netboxapi.FilterInt64s(params.SetPrimaryMacAddressIDn),
		"q":                                   netboxapi.FilterString(params.SetQ),
		"rack":                                netboxapi.FilterStrings(params.SetRack),
		"rack__n":                             netboxapi.FilterStrings(params.SetRackn),
		"rack_id":                             netboxapi.FilterInt64s(params.SetRackID),
		"rack_id__n":                          netboxapi.FilterInt64s(params.SetRackIDn),
		"region":                              netboxapi.FilterStrings(params.SetRegion),
		"region__n":                           netboxapi.FilterStrings(params.SetRegionn),
		"region_id":                           netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":                        netboxapi.FilterStrings(params.SetRegionIDn),
		"rf_channel":                          netboxapi.FilterStrings(params.SetRfChannel),
		"rf_channel__n":                       netboxapi.FilterStrings(params.SetRfChanneln),
		"rf_channel__ic":                      netboxapi.FilterStrings(params.SetRfChannelIc),
		"rf_channel__nic":                     netboxapi.FilterStrings(params.SetRfChannelNic),
		"rf_channel__ie":                      netboxapi.FilterStrings(params.SetRfChannelIe),
		"rf_channel__nie":                     netboxapi.FilterStrings(params.SetRfChannelNie),
		"rf_channel__isw":                     netboxapi.FilterStrings(params.SetRfChannelIsw),
		"rf_channel__nisw":                    netboxapi.FilterStrings(params.SetRfChannelNisw),
		"rf_channel__iew":                     netboxapi.FilterStrings(params.SetRfChannelIew),
		"rf_channel__niew":                    netboxapi.FilterStrings(params.SetRfChannelNiew),
		"rf_channel__empty":                   netboxapi.FilterBool(params.SetRfChannelEmpty),
		"rf_channel__regex":                   netboxapi.FilterStrings(params.SetRfChannelRegex),
		"rf_channel__iregex":                  netboxapi.FilterStrings(params.SetRfChannelIregex),
		"rf_channel_frequency":                netboxapi.FilterFloat64s(params.SetRfChannelFrequency),
		"rf_channel_frequency__n":             netboxapi.FilterFloat64s(params.SetRfChannelFrequencyn),
		"rf_channel_frequency__lt":            netboxapi.FilterFloat64s(params.SetRfChannelFrequencyLt),
		"rf_channel_frequency__lte":           netboxapi.FilterFloat64s(params.SetRfChannelFrequencyLte),
		"rf_channel_frequency__gt":            netboxapi.FilterFloat64s(params.SetRfChannelFrequencyGt),
		"rf_channel_frequency__gte":           netboxapi.FilterFloat64s(params.SetRfChannelFrequencyGte),
		"rf_channel_frequency__empty":         netboxapi.FilterBool(params.SetRfChannelFrequencyEmpty),
		"rf_channel_width":                    netboxapi.FilterFloat64s(params.SetRfChannelWidth),
		"rf_channel_width__n":                 netboxapi.FilterFloat64s(params.SetRfChannelWidthn),
		"rf_channel_width__lt":                netboxapi.FilterFloat64s(params.SetRfChannelWidthLt),
		"rf_channel_width__lte":               netboxapi.FilterFloat64s(params.SetRfChannelWidthLte),
		"rf_channel_width__gt":                netboxapi.FilterFloat64s(params.SetRfChannelWidthGt),
		"rf_channel_width__gte":               netboxapi.FilterFloat64s(params.SetRfChannelWidthGte),
		"rf_channel_width__empty":             netboxapi.FilterBool(params.SetRfChannelWidthEmpty),
		"rf_role":                             netboxapi.FilterStrings(params.SetRfRole),
		"rf_role__n":                          netboxapi.FilterStrings(params.SetRfRolen),
		"rf_role__ic":                         netboxapi.FilterStrings(params.SetRfRoleIc),
		"rf_role__nic":                        netboxapi.FilterStrings(params.SetRfRoleNic),
		"rf_role__ie":                         netboxapi.FilterStrings(params.SetRfRoleIe),
		"rf_role__nie":                        netboxapi.FilterStrings(params.SetRfRoleNie),
		"rf_role__isw":                        netboxapi.FilterStrings(params.SetRfRoleIsw),
		"rf_role__nisw":                       netboxapi.FilterStrings(params.SetRfRoleNisw),
		"rf_role__iew":                        netboxapi.FilterStrings(params.SetRfRoleIew),
		"rf_role__niew":                       netboxapi.FilterStrings(params.SetRfRoleNiew),
		"rf_role__empty":                      netboxapi.FilterBool(params.SetRfRoleEmpty),
		"rf_role__regex":                      netboxapi.FilterStrings(params.SetRfRoleRegex),
		"rf_role__iregex":                     netboxapi.FilterStrings(params.SetRfRoleIregex),
		"site":                                netboxapi.FilterStrings(params.SetSite),
		"site__n":                             netboxapi.FilterStrings(params.SetSiten),
		"site_group":                          netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":                       netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":                       netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":                    netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"site_id":                             netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":                          netboxapi.FilterInt64s(params.SetSiteIDn),
		"speed":                               netboxapi.FilterInt64s(params.SetSpeed),
		"speed__n":                            netboxapi.FilterInt64s(params.SetSpeedn),
		"speed__lt":                           netboxapi.FilterInt64s(params.SetSpeedLt),
		"speed__lte":                          netboxapi.FilterInt64s(params.SetSpeedLte),
		"speed__gt":                           netboxapi.FilterInt64s(params.SetSpeedGt),
		"speed__gte":                          netboxapi.FilterInt64s(params.SetSpeedGte),
		"speed__empty":                        netboxapi.FilterInt64s(params.SetSpeedEmpty),
		"tag_id":                              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                         netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                              netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":                           netboxapi.FilterStrings(params.SetTenantn),
		"tenant_id":                           netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":                        netboxapi.FilterInt64s(params.SetTenantIDn),
		"tx_power":                            netboxapi.FilterInt64s(params.SetTxPower),
		"tx_power__n":                         netboxapi.FilterInt64s(params.SetTxPowern),
		"tx_power__lt":                        netboxapi.FilterInt64s(params.SetTxPowerLt),
		"tx_power__lte":                       netboxapi.FilterInt64s(params.SetTxPowerLte),
		"tx_power__gt":                        netboxapi.FilterInt64s(params.SetTxPowerGt),
		"tx_power__gte":                       netboxapi.FilterInt64s(params.SetTxPowerGte),
		"tx_power__empty":                     netboxapi.FilterBool(params.SetTxPowerEmpty),
		"vdc":                                 netboxapi.FilterStrings(params.SetVdc),
		"vdc__n":                              netboxapi.FilterStrings(params.SetVdcn),
		"vdc_id":                              netboxapi.FilterInt64s(params.SetVdcID),
		"vdc_id__n":                           netboxapi.FilterInt64s(params.SetVdcIDn),
		"vdc_identifier":                      netboxapi.FilterInt64s(params.SetVdcIdentifier),
		"vdc_identifier__n":                   netboxapi.FilterInt64s(params.SetVdcIdentifiern),
		"virtual_chassis":                     netboxapi.FilterStrings(params.SetVirtualChassis),
		"virtual_chassis__n":                  netboxapi.FilterStrings(params.SetVirtualChassisn),
		"virtual_chassis_id":                  netboxapi.FilterInt64s(params.SetVirtualChassisID),
		"virtual_chassis_id__n":               netboxapi.FilterInt64s(params.SetVirtualChassisIDn),
		"virtual_chassis_member":              netboxapi.FilterStrings(params.SetVirtualChassisMember),
		"virtual_chassis_member_id":           netboxapi.FilterInt64s(params.SetVirtualChassisMemberID),
		"virtual_chassis_member_or_master":    netboxapi.FilterStrings(params.SetVirtualChassisMemberOrMaster),
		"virtual_chassis_member_or_master_id": netboxapi.FilterInt64s(params.SetVirtualChassisMemberOrMasterID),
		"virtual_circuit_id":                  netboxapi.FilterInt64s(params.SetVirtualCircuitID),
		"virtual_circuit_id__n":               netboxapi.FilterInt64s(params.SetVirtualCircuitIDn),
		"virtual_circuit_termination_id":      netboxapi.FilterInt64s(params.SetVirtualCircuitTerminationID),
		"virtual_circuit_termination_id__n":   netboxapi.FilterInt64s(params.SetVirtualCircuitTerminationIDn),
		"vlan":                                netboxapi.FilterString(params.SetVlan),
		"vlan_id":                             netboxapi.FilterString(params.SetVlanID),
		"vlan_translation_policy":             netboxapi.FilterStrings(params.SetVlanTranslationPolicy),
		"vlan_translation_policy__n":          netboxapi.FilterStrings(params.SetVlanTranslationPolicyn),
		"vlan_translation_policy_id":          netboxapi.FilterInt64s(params.SetVlanTranslationPolicyID),
		"vlan_translation_policy_id__n":       netboxapi.FilterInt64s(params.SetVlanTranslationPolicyIDn),
		"vrf":                                 netboxapi.FilterStrings(params.SetVrf),
		"vrf__n":                              netboxapi.FilterStrings(params.SetVrfn),
		"vrf_id":                              netboxapi.FilterInt64s(params.SetVrfID),
		"vrf_id__n":                           netboxapi.FilterInt64s(params.SetVrfIDn),
		"wireless_lan_id":                     netboxapi.FilterInt64s(params.SetWirelessLanID),
		"wireless_lan_id__n":                  netboxapi.FilterInt64s(params.SetWirelessLanIDn),
		"wireless_link_id":                    netboxapi.FilterInt64s(params.SetWirelessLinkID),
		"wireless_link_id__n":                 netboxapi.FilterInt64s(params.SetWirelessLinkIDn),
		"wwn":                                 netboxapi.FilterStrings(params.SetWwn),
		"wwn__n":                              netboxapi.FilterStrings(params.SetWwnn),
		"wwn__ic":                             netboxapi.FilterStrings(params.SetWwnIc),
		"wwn__nic":                            netboxapi.FilterStrings(params.SetWwnNic),
		"wwn__ie":                             netboxapi.FilterStrings(params.SetWwnIe),
		"wwn__nie":                            netboxapi.FilterStrings(params.SetWwnNie),
		"wwn__isw":                            netboxapi.FilterStrings(params.SetWwnIsw),
		"wwn__nisw":                           netboxapi.FilterStrings(params.SetWwnNisw),
		"wwn__iew":                            netboxapi.FilterStrings(params.SetWwnIew),
		"wwn__niew":                           netboxapi.FilterStrings(params.SetWwnNiew),
		"wwn__regex":                          netboxapi.FilterStrings(params.SetWwnRegex),
		"wwn__iregex":                         netboxapi.FilterStrings(params.SetWwnIregex),
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
		res, err := d.client.Dcim.DcimInterfacesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_device_interfaces", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.DeviceInterfaceResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]deviceInterfaceResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model deviceInterfaceResourceModel
		resp.Diagnostics.Append(flattenDeviceInterface(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: deviceInterfaceResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
