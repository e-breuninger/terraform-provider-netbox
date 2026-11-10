// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vDeviceID []int64
	var vDeviceIDn []int64
	var vType []string
	var vTypen []string
	var vTypeIc []string
	var vTypeNic []string
	var vTypeIe []string
	var vTypeNie []string
	var vTypeIsw []string
	var vTypeNisw []string
	var vTypeIew []string
	var vTypeNiew []string
	var vTypeEmpty *bool
	var vTypeRegex []string
	var vTypeIregex []string
	var vMode []string
	var vModen []string
	var vModeIc []string
	var vModeNic []string
	var vModeIe []string
	var vModeNie []string
	var vModeIsw []string
	var vModeNisw []string
	var vModeIew []string
	var vModeNiew []string
	var vModeEmpty *bool
	var vModeRegex []string
	var vModeIregex []string
	var vEnabled *bool
	var vLagID []int64
	var vLagIDn []int64
	var vMacAddress []string
	var vMacAddressn []string
	var vMacAddressIc []string
	var vMacAddressNic []string
	var vMacAddressIe []string
	var vMacAddressNie []string
	var vMacAddressIsw []string
	var vMacAddressNisw []string
	var vMacAddressIew []string
	var vMacAddressNiew []string
	var vMacAddressRegex []string
	var vMacAddressIregex []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vBridgeID []int64
	var vBridgeIDn []int64
	var vCableConnector []int64
	var vCableConnectorn []int64
	var vCableConnectorLt []int64
	var vCableConnectorLte []int64
	var vCableConnectorGt []int64
	var vCableConnectorGte []int64
	var vCableConnectorEmpty *bool
	var vCableEnd *string
	var vCableEndn *string
	var vCableEndIc []string
	var vCableEndNic []string
	var vCableEndIe []string
	var vCableEndNie []string
	var vCableEndIsw []string
	var vCableEndNisw []string
	var vCableEndIew []string
	var vCableEndNiew []string
	var vCableEndEmpty *bool
	var vCableEndRegex []string
	var vCableEndIregex []string
	var vCableID []int64
	var vCableIDn []int64
	var vCabled *bool
	var vConnected *bool
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
	var vDevice []string
	var vDevicen []string
	var vDeviceRole []string
	var vDeviceRolen []string
	var vDeviceRoleID []int64
	var vDeviceRoleIDn []int64
	var vDeviceStatus []string
	var vDeviceStatusn []string
	var vDeviceStatusIc []string
	var vDeviceStatusNic []string
	var vDeviceStatusIe []string
	var vDeviceStatusNie []string
	var vDeviceStatusIsw []string
	var vDeviceStatusNisw []string
	var vDeviceStatusIew []string
	var vDeviceStatusNiew []string
	var vDeviceStatusEmpty *bool
	var vDeviceStatusRegex []string
	var vDeviceStatusIregex []string
	var vDeviceType []string
	var vDeviceTypen []string
	var vDeviceTypeID []int64
	var vDeviceTypeIDn []int64
	var vDuplex []string
	var vDuplexn []string
	var vDuplexIc []string
	var vDuplexNic []string
	var vDuplexIe []string
	var vDuplexNie []string
	var vDuplexIsw []string
	var vDuplexNisw []string
	var vDuplexIew []string
	var vDuplexNiew []string
	var vDuplexEmpty *bool
	var vDuplexRegex []string
	var vDuplexIregex []string
	var vKind *string
	var vL2vpn []int64
	var vL2vpnn []int64
	var vL2vpnID []int64
	var vL2vpnIDn []int64
	var vLabel []string
	var vLabeln []string
	var vLabelIc []string
	var vLabelNic []string
	var vLabelIe []string
	var vLabelNie []string
	var vLabelIsw []string
	var vLabelNisw []string
	var vLabelIew []string
	var vLabelNiew []string
	var vLabelEmpty *bool
	var vLabelRegex []string
	var vLabelIregex []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLocation []string
	var vLocationn []string
	var vLocationID []int64
	var vLocationIDn []int64
	var vMarkConnected *bool
	var vMgmtOnly *bool
	var vModuleID []int64
	var vModuleIDn []int64
	var vMtu []int64
	var vMtun []int64
	var vMtuLt []int64
	var vMtuLte []int64
	var vMtuGt []int64
	var vMtuGte []int64
	var vMtuEmpty *bool
	var vOccupied *bool
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vParentID []int64
	var vParentIDn []int64
	var vPoeMode []string
	var vPoeModen []string
	var vPoeModeIc []string
	var vPoeModeNic []string
	var vPoeModeIe []string
	var vPoeModeNie []string
	var vPoeModeIsw []string
	var vPoeModeNisw []string
	var vPoeModeIew []string
	var vPoeModeNiew []string
	var vPoeModeEmpty *bool
	var vPoeModeRegex []string
	var vPoeModeIregex []string
	var vPoeType []string
	var vPoeTypen []string
	var vPoeTypeIc []string
	var vPoeTypeNic []string
	var vPoeTypeIe []string
	var vPoeTypeNie []string
	var vPoeTypeIsw []string
	var vPoeTypeNisw []string
	var vPoeTypeIew []string
	var vPoeTypeNiew []string
	var vPoeTypeEmpty *bool
	var vPoeTypeRegex []string
	var vPoeTypeIregex []string
	var vPrimaryMacAddress []string
	var vPrimaryMacAddressn []string
	var vPrimaryMacAddressID []int64
	var vPrimaryMacAddressIDn []int64
	var vQ *string
	var vRack []string
	var vRackn []string
	var vRackID []int64
	var vRackIDn []int64
	var vRegion []string
	var vRegionn []string
	var vRegionID []string
	var vRegionIDn []string
	var vRfChannel []string
	var vRfChanneln []string
	var vRfChannelIc []string
	var vRfChannelNic []string
	var vRfChannelIe []string
	var vRfChannelNie []string
	var vRfChannelIsw []string
	var vRfChannelNisw []string
	var vRfChannelIew []string
	var vRfChannelNiew []string
	var vRfChannelEmpty *bool
	var vRfChannelRegex []string
	var vRfChannelIregex []string
	var vRfChannelFrequency []float64
	var vRfChannelFrequencyn []float64
	var vRfChannelFrequencyLt []float64
	var vRfChannelFrequencyLte []float64
	var vRfChannelFrequencyGt []float64
	var vRfChannelFrequencyGte []float64
	var vRfChannelFrequencyEmpty *bool
	var vRfChannelWidth []float64
	var vRfChannelWidthn []float64
	var vRfChannelWidthLt []float64
	var vRfChannelWidthLte []float64
	var vRfChannelWidthGt []float64
	var vRfChannelWidthGte []float64
	var vRfChannelWidthEmpty *bool
	var vRfRole []string
	var vRfRolen []string
	var vRfRoleIc []string
	var vRfRoleNic []string
	var vRfRoleIe []string
	var vRfRoleNie []string
	var vRfRoleIsw []string
	var vRfRoleNisw []string
	var vRfRoleIew []string
	var vRfRoleNiew []string
	var vRfRoleEmpty *bool
	var vRfRoleRegex []string
	var vRfRoleIregex []string
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []string
	var vSiteGroupIDn []string
	var vSiteID []int64
	var vSiteIDn []int64
	var vSpeed []int64
	var vSpeedn []int64
	var vSpeedLt []int64
	var vSpeedLte []int64
	var vSpeedGt []int64
	var vSpeedGte []int64
	var vSpeedEmpty []int64
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantID []int64
	var vTenantIDn []int64
	var vTxPower []int64
	var vTxPowern []int64
	var vTxPowerLt []int64
	var vTxPowerLte []int64
	var vTxPowerGt []int64
	var vTxPowerGte []int64
	var vTxPowerEmpty *bool
	var vVdc []string
	var vVdcn []string
	var vVdcID []int64
	var vVdcIDn []int64
	var vVdcIdentifier []int64
	var vVdcIdentifiern []int64
	var vVirtualChassis []string
	var vVirtualChassisn []string
	var vVirtualChassisID []int64
	var vVirtualChassisIDn []int64
	var vVirtualChassisMember []string
	var vVirtualChassisMemberID []int64
	var vVirtualChassisMemberOrMaster []string
	var vVirtualChassisMemberOrMasterID []int64
	var vVirtualCircuitID []int64
	var vVirtualCircuitIDn []int64
	var vVirtualCircuitTerminationID []int64
	var vVirtualCircuitTerminationIDn []int64
	var vVlan *string
	var vVlanID *string
	var vVlanTranslationPolicy []string
	var vVlanTranslationPolicyn []string
	var vVlanTranslationPolicyID []int64
	var vVlanTranslationPolicyIDn []int64
	var vVrf []string
	var vVrfn []string
	var vVrfID []int64
	var vVrfIDn []int64
	var vWirelessLanID []int64
	var vWirelessLanIDn []int64
	var vWirelessLinkID []int64
	var vWirelessLinkIDn []int64
	var vWwn []string
	var vWwnn []string
	var vWwnIc []string
	var vWwnNic []string
	var vWwnIe []string
	var vWwnNie []string
	var vWwnIsw []string
	var vWwnNisw []string
	var vWwnIew []string
	var vWwnNiew []string
	var vWwnRegex []string
	var vWwnIregex []string
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
		case "device_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id' takes an integer, got %q.", value))
				return
			}
			vDeviceID = append(vDeviceID, v)
		case "device_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceIDn = append(vDeviceIDn, v)
		case "type":
			v := value
			vType = append(vType, v)
		case "type__n":
			v := value
			vTypen = append(vTypen, v)
		case "type__ic":
			v := value
			vTypeIc = append(vTypeIc, v)
		case "type__nic":
			v := value
			vTypeNic = append(vTypeNic, v)
		case "type__ie":
			v := value
			vTypeIe = append(vTypeIe, v)
		case "type__nie":
			v := value
			vTypeNie = append(vTypeNie, v)
		case "type__isw":
			v := value
			vTypeIsw = append(vTypeIsw, v)
		case "type__nisw":
			v := value
			vTypeNisw = append(vTypeNisw, v)
		case "type__iew":
			v := value
			vTypeIew = append(vTypeIew, v)
		case "type__niew":
			v := value
			vTypeNiew = append(vTypeNiew, v)
		case "type__empty":
			if vTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'type__empty' takes a boolean, got %q.", value))
				return
			}
			vTypeEmpty = &v
		case "type__regex":
			v := value
			vTypeRegex = append(vTypeRegex, v)
		case "type__iregex":
			v := value
			vTypeIregex = append(vTypeIregex, v)
		case "mode":
			v := value
			vMode = append(vMode, v)
		case "mode__n":
			v := value
			vModen = append(vModen, v)
		case "mode__ic":
			v := value
			vModeIc = append(vModeIc, v)
		case "mode__nic":
			v := value
			vModeNic = append(vModeNic, v)
		case "mode__ie":
			v := value
			vModeIe = append(vModeIe, v)
		case "mode__nie":
			v := value
			vModeNie = append(vModeNie, v)
		case "mode__isw":
			v := value
			vModeIsw = append(vModeIsw, v)
		case "mode__nisw":
			v := value
			vModeNisw = append(vModeNisw, v)
		case "mode__iew":
			v := value
			vModeIew = append(vModeIew, v)
		case "mode__niew":
			v := value
			vModeNiew = append(vModeNiew, v)
		case "mode__empty":
			if vModeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mode__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mode__empty' takes a boolean, got %q.", value))
				return
			}
			vModeEmpty = &v
		case "mode__regex":
			v := value
			vModeRegex = append(vModeRegex, v)
		case "mode__iregex":
			v := value
			vModeIregex = append(vModeIregex, v)
		case "enabled":
			if vEnabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'enabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'enabled' takes a boolean, got %q.", value))
				return
			}
			vEnabled = &v
		case "lag_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'lag_id' takes an integer, got %q.", value))
				return
			}
			vLagID = append(vLagID, v)
		case "lag_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'lag_id__n' takes an integer, got %q.", value))
				return
			}
			vLagIDn = append(vLagIDn, v)
		case "mac_address":
			v := value
			vMacAddress = append(vMacAddress, v)
		case "mac_address__n":
			v := value
			vMacAddressn = append(vMacAddressn, v)
		case "mac_address__ic":
			v := value
			vMacAddressIc = append(vMacAddressIc, v)
		case "mac_address__nic":
			v := value
			vMacAddressNic = append(vMacAddressNic, v)
		case "mac_address__ie":
			v := value
			vMacAddressIe = append(vMacAddressIe, v)
		case "mac_address__nie":
			v := value
			vMacAddressNie = append(vMacAddressNie, v)
		case "mac_address__isw":
			v := value
			vMacAddressIsw = append(vMacAddressIsw, v)
		case "mac_address__nisw":
			v := value
			vMacAddressNisw = append(vMacAddressNisw, v)
		case "mac_address__iew":
			v := value
			vMacAddressIew = append(vMacAddressIew, v)
		case "mac_address__niew":
			v := value
			vMacAddressNiew = append(vMacAddressNiew, v)
		case "mac_address__regex":
			v := value
			vMacAddressRegex = append(vMacAddressRegex, v)
		case "mac_address__iregex":
			v := value
			vMacAddressIregex = append(vMacAddressIregex, v)
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
		case "bridge_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'bridge_id' takes an integer, got %q.", value))
				return
			}
			vBridgeID = append(vBridgeID, v)
		case "bridge_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'bridge_id__n' takes an integer, got %q.", value))
				return
			}
			vBridgeIDn = append(vBridgeIDn, v)
		case "cable_connector":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector' takes an integer, got %q.", value))
				return
			}
			vCableConnector = append(vCableConnector, v)
		case "cable_connector__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__n' takes an integer, got %q.", value))
				return
			}
			vCableConnectorn = append(vCableConnectorn, v)
		case "cable_connector__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__lt' takes an integer, got %q.", value))
				return
			}
			vCableConnectorLt = append(vCableConnectorLt, v)
		case "cable_connector__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__lte' takes an integer, got %q.", value))
				return
			}
			vCableConnectorLte = append(vCableConnectorLte, v)
		case "cable_connector__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__gt' takes an integer, got %q.", value))
				return
			}
			vCableConnectorGt = append(vCableConnectorGt, v)
		case "cable_connector__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__gte' takes an integer, got %q.", value))
				return
			}
			vCableConnectorGte = append(vCableConnectorGte, v)
		case "cable_connector__empty":
			if vCableConnectorEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_connector__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__empty' takes a boolean, got %q.", value))
				return
			}
			vCableConnectorEmpty = &v
		case "cable_end":
			if vCableEnd != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end' takes a single value.")
				return
			}
			v := value
			vCableEnd = &v
		case "cable_end__n":
			if vCableEndn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end__n' takes a single value.")
				return
			}
			v := value
			vCableEndn = &v
		case "cable_end__ic":
			v := value
			vCableEndIc = append(vCableEndIc, v)
		case "cable_end__nic":
			v := value
			vCableEndNic = append(vCableEndNic, v)
		case "cable_end__ie":
			v := value
			vCableEndIe = append(vCableEndIe, v)
		case "cable_end__nie":
			v := value
			vCableEndNie = append(vCableEndNie, v)
		case "cable_end__isw":
			v := value
			vCableEndIsw = append(vCableEndIsw, v)
		case "cable_end__nisw":
			v := value
			vCableEndNisw = append(vCableEndNisw, v)
		case "cable_end__iew":
			v := value
			vCableEndIew = append(vCableEndIew, v)
		case "cable_end__niew":
			v := value
			vCableEndNiew = append(vCableEndNiew, v)
		case "cable_end__empty":
			if vCableEndEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_end__empty' takes a boolean, got %q.", value))
				return
			}
			vCableEndEmpty = &v
		case "cable_end__regex":
			v := value
			vCableEndRegex = append(vCableEndRegex, v)
		case "cable_end__iregex":
			v := value
			vCableEndIregex = append(vCableEndIregex, v)
		case "cable_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_id' takes an integer, got %q.", value))
				return
			}
			vCableID = append(vCableID, v)
		case "cable_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_id__n' takes an integer, got %q.", value))
				return
			}
			vCableIDn = append(vCableIDn, v)
		case "cabled":
			if vCabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cabled' takes a boolean, got %q.", value))
				return
			}
			vCabled = &v
		case "connected":
			if vConnected != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'connected' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'connected' takes a boolean, got %q.", value))
				return
			}
			vConnected = &v
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
		case "device":
			v := value
			vDevice = append(vDevice, v)
		case "device__n":
			v := value
			vDevicen = append(vDevicen, v)
		case "device_role":
			v := value
			vDeviceRole = append(vDeviceRole, v)
		case "device_role__n":
			v := value
			vDeviceRolen = append(vDeviceRolen, v)
		case "device_role_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_role_id' takes an integer, got %q.", value))
				return
			}
			vDeviceRoleID = append(vDeviceRoleID, v)
		case "device_role_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_role_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceRoleIDn = append(vDeviceRoleIDn, v)
		case "device_status":
			v := value
			vDeviceStatus = append(vDeviceStatus, v)
		case "device_status__n":
			v := value
			vDeviceStatusn = append(vDeviceStatusn, v)
		case "device_status__ic":
			v := value
			vDeviceStatusIc = append(vDeviceStatusIc, v)
		case "device_status__nic":
			v := value
			vDeviceStatusNic = append(vDeviceStatusNic, v)
		case "device_status__ie":
			v := value
			vDeviceStatusIe = append(vDeviceStatusIe, v)
		case "device_status__nie":
			v := value
			vDeviceStatusNie = append(vDeviceStatusNie, v)
		case "device_status__isw":
			v := value
			vDeviceStatusIsw = append(vDeviceStatusIsw, v)
		case "device_status__nisw":
			v := value
			vDeviceStatusNisw = append(vDeviceStatusNisw, v)
		case "device_status__iew":
			v := value
			vDeviceStatusIew = append(vDeviceStatusIew, v)
		case "device_status__niew":
			v := value
			vDeviceStatusNiew = append(vDeviceStatusNiew, v)
		case "device_status__empty":
			if vDeviceStatusEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'device_status__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_status__empty' takes a boolean, got %q.", value))
				return
			}
			vDeviceStatusEmpty = &v
		case "device_status__regex":
			v := value
			vDeviceStatusRegex = append(vDeviceStatusRegex, v)
		case "device_status__iregex":
			v := value
			vDeviceStatusIregex = append(vDeviceStatusIregex, v)
		case "device_type":
			v := value
			vDeviceType = append(vDeviceType, v)
		case "device_type__n":
			v := value
			vDeviceTypen = append(vDeviceTypen, v)
		case "device_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_type_id' takes an integer, got %q.", value))
				return
			}
			vDeviceTypeID = append(vDeviceTypeID, v)
		case "device_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_type_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceTypeIDn = append(vDeviceTypeIDn, v)
		case "duplex":
			v := value
			vDuplex = append(vDuplex, v)
		case "duplex__n":
			v := value
			vDuplexn = append(vDuplexn, v)
		case "duplex__ic":
			v := value
			vDuplexIc = append(vDuplexIc, v)
		case "duplex__nic":
			v := value
			vDuplexNic = append(vDuplexNic, v)
		case "duplex__ie":
			v := value
			vDuplexIe = append(vDuplexIe, v)
		case "duplex__nie":
			v := value
			vDuplexNie = append(vDuplexNie, v)
		case "duplex__isw":
			v := value
			vDuplexIsw = append(vDuplexIsw, v)
		case "duplex__nisw":
			v := value
			vDuplexNisw = append(vDuplexNisw, v)
		case "duplex__iew":
			v := value
			vDuplexIew = append(vDuplexIew, v)
		case "duplex__niew":
			v := value
			vDuplexNiew = append(vDuplexNiew, v)
		case "duplex__empty":
			if vDuplexEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'duplex__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'duplex__empty' takes a boolean, got %q.", value))
				return
			}
			vDuplexEmpty = &v
		case "duplex__regex":
			v := value
			vDuplexRegex = append(vDuplexRegex, v)
		case "duplex__iregex":
			v := value
			vDuplexIregex = append(vDuplexIregex, v)
		case "kind":
			if vKind != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'kind' takes a single value.")
				return
			}
			v := value
			vKind = &v
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
		case "label":
			v := value
			vLabel = append(vLabel, v)
		case "label__n":
			v := value
			vLabeln = append(vLabeln, v)
		case "label__ic":
			v := value
			vLabelIc = append(vLabelIc, v)
		case "label__nic":
			v := value
			vLabelNic = append(vLabelNic, v)
		case "label__ie":
			v := value
			vLabelIe = append(vLabelIe, v)
		case "label__nie":
			v := value
			vLabelNie = append(vLabelNie, v)
		case "label__isw":
			v := value
			vLabelIsw = append(vLabelIsw, v)
		case "label__nisw":
			v := value
			vLabelNisw = append(vLabelNisw, v)
		case "label__iew":
			v := value
			vLabelIew = append(vLabelIew, v)
		case "label__niew":
			v := value
			vLabelNiew = append(vLabelNiew, v)
		case "label__empty":
			if vLabelEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'label__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'label__empty' takes a boolean, got %q.", value))
				return
			}
			vLabelEmpty = &v
		case "label__regex":
			v := value
			vLabelRegex = append(vLabelRegex, v)
		case "label__iregex":
			v := value
			vLabelIregex = append(vLabelIregex, v)
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
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id' takes an integer, got %q.", value))
				return
			}
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id__n' takes an integer, got %q.", value))
				return
			}
			vLocationIDn = append(vLocationIDn, v)
		case "mark_connected":
			if vMarkConnected != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mark_connected' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mark_connected' takes a boolean, got %q.", value))
				return
			}
			vMarkConnected = &v
		case "mgmt_only":
			if vMgmtOnly != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mgmt_only' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mgmt_only' takes a boolean, got %q.", value))
				return
			}
			vMgmtOnly = &v
		case "module_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_id' takes an integer, got %q.", value))
				return
			}
			vModuleID = append(vModuleID, v)
		case "module_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_id__n' takes an integer, got %q.", value))
				return
			}
			vModuleIDn = append(vModuleIDn, v)
		case "mtu":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu' takes an integer, got %q.", value))
				return
			}
			vMtu = append(vMtu, v)
		case "mtu__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__n' takes an integer, got %q.", value))
				return
			}
			vMtun = append(vMtun, v)
		case "mtu__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__lt' takes an integer, got %q.", value))
				return
			}
			vMtuLt = append(vMtuLt, v)
		case "mtu__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__lte' takes an integer, got %q.", value))
				return
			}
			vMtuLte = append(vMtuLte, v)
		case "mtu__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__gt' takes an integer, got %q.", value))
				return
			}
			vMtuGt = append(vMtuGt, v)
		case "mtu__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__gte' takes an integer, got %q.", value))
				return
			}
			vMtuGte = append(vMtuGte, v)
		case "mtu__empty":
			if vMtuEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mtu__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__empty' takes a boolean, got %q.", value))
				return
			}
			vMtuEmpty = &v
		case "occupied":
			if vOccupied != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'occupied' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'occupied' takes a boolean, got %q.", value))
				return
			}
			vOccupied = &v
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
		case "parent_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_id' takes an integer, got %q.", value))
				return
			}
			vParentID = append(vParentID, v)
		case "parent_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_id__n' takes an integer, got %q.", value))
				return
			}
			vParentIDn = append(vParentIDn, v)
		case "poe_mode":
			v := value
			vPoeMode = append(vPoeMode, v)
		case "poe_mode__n":
			v := value
			vPoeModen = append(vPoeModen, v)
		case "poe_mode__ic":
			v := value
			vPoeModeIc = append(vPoeModeIc, v)
		case "poe_mode__nic":
			v := value
			vPoeModeNic = append(vPoeModeNic, v)
		case "poe_mode__ie":
			v := value
			vPoeModeIe = append(vPoeModeIe, v)
		case "poe_mode__nie":
			v := value
			vPoeModeNie = append(vPoeModeNie, v)
		case "poe_mode__isw":
			v := value
			vPoeModeIsw = append(vPoeModeIsw, v)
		case "poe_mode__nisw":
			v := value
			vPoeModeNisw = append(vPoeModeNisw, v)
		case "poe_mode__iew":
			v := value
			vPoeModeIew = append(vPoeModeIew, v)
		case "poe_mode__niew":
			v := value
			vPoeModeNiew = append(vPoeModeNiew, v)
		case "poe_mode__empty":
			if vPoeModeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'poe_mode__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'poe_mode__empty' takes a boolean, got %q.", value))
				return
			}
			vPoeModeEmpty = &v
		case "poe_mode__regex":
			v := value
			vPoeModeRegex = append(vPoeModeRegex, v)
		case "poe_mode__iregex":
			v := value
			vPoeModeIregex = append(vPoeModeIregex, v)
		case "poe_type":
			v := value
			vPoeType = append(vPoeType, v)
		case "poe_type__n":
			v := value
			vPoeTypen = append(vPoeTypen, v)
		case "poe_type__ic":
			v := value
			vPoeTypeIc = append(vPoeTypeIc, v)
		case "poe_type__nic":
			v := value
			vPoeTypeNic = append(vPoeTypeNic, v)
		case "poe_type__ie":
			v := value
			vPoeTypeIe = append(vPoeTypeIe, v)
		case "poe_type__nie":
			v := value
			vPoeTypeNie = append(vPoeTypeNie, v)
		case "poe_type__isw":
			v := value
			vPoeTypeIsw = append(vPoeTypeIsw, v)
		case "poe_type__nisw":
			v := value
			vPoeTypeNisw = append(vPoeTypeNisw, v)
		case "poe_type__iew":
			v := value
			vPoeTypeIew = append(vPoeTypeIew, v)
		case "poe_type__niew":
			v := value
			vPoeTypeNiew = append(vPoeTypeNiew, v)
		case "poe_type__empty":
			if vPoeTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'poe_type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'poe_type__empty' takes a boolean, got %q.", value))
				return
			}
			vPoeTypeEmpty = &v
		case "poe_type__regex":
			v := value
			vPoeTypeRegex = append(vPoeTypeRegex, v)
		case "poe_type__iregex":
			v := value
			vPoeTypeIregex = append(vPoeTypeIregex, v)
		case "primary_mac_address":
			v := value
			vPrimaryMacAddress = append(vPrimaryMacAddress, v)
		case "primary_mac_address__n":
			v := value
			vPrimaryMacAddressn = append(vPrimaryMacAddressn, v)
		case "primary_mac_address_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_mac_address_id' takes an integer, got %q.", value))
				return
			}
			vPrimaryMacAddressID = append(vPrimaryMacAddressID, v)
		case "primary_mac_address_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_mac_address_id__n' takes an integer, got %q.", value))
				return
			}
			vPrimaryMacAddressIDn = append(vPrimaryMacAddressIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "rack":
			v := value
			vRack = append(vRack, v)
		case "rack__n":
			v := value
			vRackn = append(vRackn, v)
		case "rack_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_id' takes an integer, got %q.", value))
				return
			}
			vRackID = append(vRackID, v)
		case "rack_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_id__n' takes an integer, got %q.", value))
				return
			}
			vRackIDn = append(vRackIDn, v)
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
		case "rf_channel":
			v := value
			vRfChannel = append(vRfChannel, v)
		case "rf_channel__n":
			v := value
			vRfChanneln = append(vRfChanneln, v)
		case "rf_channel__ic":
			v := value
			vRfChannelIc = append(vRfChannelIc, v)
		case "rf_channel__nic":
			v := value
			vRfChannelNic = append(vRfChannelNic, v)
		case "rf_channel__ie":
			v := value
			vRfChannelIe = append(vRfChannelIe, v)
		case "rf_channel__nie":
			v := value
			vRfChannelNie = append(vRfChannelNie, v)
		case "rf_channel__isw":
			v := value
			vRfChannelIsw = append(vRfChannelIsw, v)
		case "rf_channel__nisw":
			v := value
			vRfChannelNisw = append(vRfChannelNisw, v)
		case "rf_channel__iew":
			v := value
			vRfChannelIew = append(vRfChannelIew, v)
		case "rf_channel__niew":
			v := value
			vRfChannelNiew = append(vRfChannelNiew, v)
		case "rf_channel__empty":
			if vRfChannelEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'rf_channel__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel__empty' takes a boolean, got %q.", value))
				return
			}
			vRfChannelEmpty = &v
		case "rf_channel__regex":
			v := value
			vRfChannelRegex = append(vRfChannelRegex, v)
		case "rf_channel__iregex":
			v := value
			vRfChannelIregex = append(vRfChannelIregex, v)
		case "rf_channel_frequency":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_frequency' takes a number, got %q.", value))
				return
			}
			vRfChannelFrequency = append(vRfChannelFrequency, v)
		case "rf_channel_frequency__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_frequency__n' takes a number, got %q.", value))
				return
			}
			vRfChannelFrequencyn = append(vRfChannelFrequencyn, v)
		case "rf_channel_frequency__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_frequency__lt' takes a number, got %q.", value))
				return
			}
			vRfChannelFrequencyLt = append(vRfChannelFrequencyLt, v)
		case "rf_channel_frequency__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_frequency__lte' takes a number, got %q.", value))
				return
			}
			vRfChannelFrequencyLte = append(vRfChannelFrequencyLte, v)
		case "rf_channel_frequency__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_frequency__gt' takes a number, got %q.", value))
				return
			}
			vRfChannelFrequencyGt = append(vRfChannelFrequencyGt, v)
		case "rf_channel_frequency__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_frequency__gte' takes a number, got %q.", value))
				return
			}
			vRfChannelFrequencyGte = append(vRfChannelFrequencyGte, v)
		case "rf_channel_frequency__empty":
			if vRfChannelFrequencyEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'rf_channel_frequency__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_frequency__empty' takes a boolean, got %q.", value))
				return
			}
			vRfChannelFrequencyEmpty = &v
		case "rf_channel_width":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_width' takes a number, got %q.", value))
				return
			}
			vRfChannelWidth = append(vRfChannelWidth, v)
		case "rf_channel_width__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_width__n' takes a number, got %q.", value))
				return
			}
			vRfChannelWidthn = append(vRfChannelWidthn, v)
		case "rf_channel_width__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_width__lt' takes a number, got %q.", value))
				return
			}
			vRfChannelWidthLt = append(vRfChannelWidthLt, v)
		case "rf_channel_width__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_width__lte' takes a number, got %q.", value))
				return
			}
			vRfChannelWidthLte = append(vRfChannelWidthLte, v)
		case "rf_channel_width__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_width__gt' takes a number, got %q.", value))
				return
			}
			vRfChannelWidthGt = append(vRfChannelWidthGt, v)
		case "rf_channel_width__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_width__gte' takes a number, got %q.", value))
				return
			}
			vRfChannelWidthGte = append(vRfChannelWidthGte, v)
		case "rf_channel_width__empty":
			if vRfChannelWidthEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'rf_channel_width__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_channel_width__empty' takes a boolean, got %q.", value))
				return
			}
			vRfChannelWidthEmpty = &v
		case "rf_role":
			v := value
			vRfRole = append(vRfRole, v)
		case "rf_role__n":
			v := value
			vRfRolen = append(vRfRolen, v)
		case "rf_role__ic":
			v := value
			vRfRoleIc = append(vRfRoleIc, v)
		case "rf_role__nic":
			v := value
			vRfRoleNic = append(vRfRoleNic, v)
		case "rf_role__ie":
			v := value
			vRfRoleIe = append(vRfRoleIe, v)
		case "rf_role__nie":
			v := value
			vRfRoleNie = append(vRfRoleNie, v)
		case "rf_role__isw":
			v := value
			vRfRoleIsw = append(vRfRoleIsw, v)
		case "rf_role__nisw":
			v := value
			vRfRoleNisw = append(vRfRoleNisw, v)
		case "rf_role__iew":
			v := value
			vRfRoleIew = append(vRfRoleIew, v)
		case "rf_role__niew":
			v := value
			vRfRoleNiew = append(vRfRoleNiew, v)
		case "rf_role__empty":
			if vRfRoleEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'rf_role__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rf_role__empty' takes a boolean, got %q.", value))
				return
			}
			vRfRoleEmpty = &v
		case "rf_role__regex":
			v := value
			vRfRoleRegex = append(vRfRoleRegex, v)
		case "rf_role__iregex":
			v := value
			vRfRoleIregex = append(vRfRoleIregex, v)
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
		case "speed":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed' takes an integer, got %q.", value))
				return
			}
			vSpeed = append(vSpeed, v)
		case "speed__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__n' takes an integer, got %q.", value))
				return
			}
			vSpeedn = append(vSpeedn, v)
		case "speed__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__lt' takes an integer, got %q.", value))
				return
			}
			vSpeedLt = append(vSpeedLt, v)
		case "speed__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__lte' takes an integer, got %q.", value))
				return
			}
			vSpeedLte = append(vSpeedLte, v)
		case "speed__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__gt' takes an integer, got %q.", value))
				return
			}
			vSpeedGt = append(vSpeedGt, v)
		case "speed__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__gte' takes an integer, got %q.", value))
				return
			}
			vSpeedGte = append(vSpeedGte, v)
		case "speed__empty":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__empty' takes an integer, got %q.", value))
				return
			}
			vSpeedEmpty = append(vSpeedEmpty, v)
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
		case "tx_power":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tx_power' takes an integer, got %q.", value))
				return
			}
			vTxPower = append(vTxPower, v)
		case "tx_power__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tx_power__n' takes an integer, got %q.", value))
				return
			}
			vTxPowern = append(vTxPowern, v)
		case "tx_power__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tx_power__lt' takes an integer, got %q.", value))
				return
			}
			vTxPowerLt = append(vTxPowerLt, v)
		case "tx_power__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tx_power__lte' takes an integer, got %q.", value))
				return
			}
			vTxPowerLte = append(vTxPowerLte, v)
		case "tx_power__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tx_power__gt' takes an integer, got %q.", value))
				return
			}
			vTxPowerGt = append(vTxPowerGt, v)
		case "tx_power__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tx_power__gte' takes an integer, got %q.", value))
				return
			}
			vTxPowerGte = append(vTxPowerGte, v)
		case "tx_power__empty":
			if vTxPowerEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'tx_power__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tx_power__empty' takes a boolean, got %q.", value))
				return
			}
			vTxPowerEmpty = &v
		case "vdc":
			v := value
			vVdc = append(vVdc, v)
		case "vdc__n":
			v := value
			vVdcn = append(vVdcn, v)
		case "vdc_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vdc_id' takes an integer, got %q.", value))
				return
			}
			vVdcID = append(vVdcID, v)
		case "vdc_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vdc_id__n' takes an integer, got %q.", value))
				return
			}
			vVdcIDn = append(vVdcIDn, v)
		case "vdc_identifier":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vdc_identifier' takes an integer, got %q.", value))
				return
			}
			vVdcIdentifier = append(vVdcIdentifier, v)
		case "vdc_identifier__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vdc_identifier__n' takes an integer, got %q.", value))
				return
			}
			vVdcIdentifiern = append(vVdcIdentifiern, v)
		case "virtual_chassis":
			v := value
			vVirtualChassis = append(vVirtualChassis, v)
		case "virtual_chassis__n":
			v := value
			vVirtualChassisn = append(vVirtualChassisn, v)
		case "virtual_chassis_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_chassis_id' takes an integer, got %q.", value))
				return
			}
			vVirtualChassisID = append(vVirtualChassisID, v)
		case "virtual_chassis_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_chassis_id__n' takes an integer, got %q.", value))
				return
			}
			vVirtualChassisIDn = append(vVirtualChassisIDn, v)
		case "virtual_chassis_member":
			v := value
			vVirtualChassisMember = append(vVirtualChassisMember, v)
		case "virtual_chassis_member_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_chassis_member_id' takes an integer, got %q.", value))
				return
			}
			vVirtualChassisMemberID = append(vVirtualChassisMemberID, v)
		case "virtual_chassis_member_or_master":
			v := value
			vVirtualChassisMemberOrMaster = append(vVirtualChassisMemberOrMaster, v)
		case "virtual_chassis_member_or_master_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_chassis_member_or_master_id' takes an integer, got %q.", value))
				return
			}
			vVirtualChassisMemberOrMasterID = append(vVirtualChassisMemberOrMasterID, v)
		case "virtual_circuit_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_circuit_id' takes an integer, got %q.", value))
				return
			}
			vVirtualCircuitID = append(vVirtualCircuitID, v)
		case "virtual_circuit_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_circuit_id__n' takes an integer, got %q.", value))
				return
			}
			vVirtualCircuitIDn = append(vVirtualCircuitIDn, v)
		case "virtual_circuit_termination_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_circuit_termination_id' takes an integer, got %q.", value))
				return
			}
			vVirtualCircuitTerminationID = append(vVirtualCircuitTerminationID, v)
		case "virtual_circuit_termination_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_circuit_termination_id__n' takes an integer, got %q.", value))
				return
			}
			vVirtualCircuitTerminationIDn = append(vVirtualCircuitTerminationIDn, v)
		case "vlan":
			if vVlan != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan' takes a single value.")
				return
			}
			v := value
			vVlan = &v
		case "vlan_id":
			if vVlanID != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_id' takes a single value.")
				return
			}
			v := value
			vVlanID = &v
		case "vlan_translation_policy":
			v := value
			vVlanTranslationPolicy = append(vVlanTranslationPolicy, v)
		case "vlan_translation_policy__n":
			v := value
			vVlanTranslationPolicyn = append(vVlanTranslationPolicyn, v)
		case "vlan_translation_policy_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_translation_policy_id' takes an integer, got %q.", value))
				return
			}
			vVlanTranslationPolicyID = append(vVlanTranslationPolicyID, v)
		case "vlan_translation_policy_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_translation_policy_id__n' takes an integer, got %q.", value))
				return
			}
			vVlanTranslationPolicyIDn = append(vVlanTranslationPolicyIDn, v)
		case "vrf":
			v := value
			vVrf = append(vVrf, v)
		case "vrf__n":
			v := value
			vVrfn = append(vVrfn, v)
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
		case "wireless_lan_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'wireless_lan_id' takes an integer, got %q.", value))
				return
			}
			vWirelessLanID = append(vWirelessLanID, v)
		case "wireless_lan_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'wireless_lan_id__n' takes an integer, got %q.", value))
				return
			}
			vWirelessLanIDn = append(vWirelessLanIDn, v)
		case "wireless_link_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'wireless_link_id' takes an integer, got %q.", value))
				return
			}
			vWirelessLinkID = append(vWirelessLinkID, v)
		case "wireless_link_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'wireless_link_id__n' takes an integer, got %q.", value))
				return
			}
			vWirelessLinkIDn = append(vWirelessLinkIDn, v)
		case "wwn":
			v := value
			vWwn = append(vWwn, v)
		case "wwn__n":
			v := value
			vWwnn = append(vWwnn, v)
		case "wwn__ic":
			v := value
			vWwnIc = append(vWwnIc, v)
		case "wwn__nic":
			v := value
			vWwnNic = append(vWwnNic, v)
		case "wwn__ie":
			v := value
			vWwnIe = append(vWwnIe, v)
		case "wwn__nie":
			v := value
			vWwnNie = append(vWwnNie, v)
		case "wwn__isw":
			v := value
			vWwnIsw = append(vWwnIsw, v)
		case "wwn__nisw":
			v := value
			vWwnNisw = append(vWwnNisw, v)
		case "wwn__iew":
			v := value
			vWwnIew = append(vWwnIew, v)
		case "wwn__niew":
			v := value
			vWwnNiew = append(vWwnNiew, v)
		case "wwn__regex":
			v := value
			vWwnRegex = append(vWwnRegex, v)
		case "wwn__iregex":
			v := value
			vWwnIregex = append(vWwnIregex, v)
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
	if len(vDeviceID) > 0 {
		params.SetDeviceID(vDeviceID)
	}
	if len(vDeviceIDn) > 0 {
		params.SetDeviceIDn(vDeviceIDn)
	}
	if len(vType) > 0 {
		params.SetType(vType)
	}
	if len(vTypen) > 0 {
		params.SetTypen(vTypen)
	}
	if len(vTypeIc) > 0 {
		params.SetTypeIc(vTypeIc)
	}
	if len(vTypeNic) > 0 {
		params.SetTypeNic(vTypeNic)
	}
	if len(vTypeIe) > 0 {
		params.SetTypeIe(vTypeIe)
	}
	if len(vTypeNie) > 0 {
		params.SetTypeNie(vTypeNie)
	}
	if len(vTypeIsw) > 0 {
		params.SetTypeIsw(vTypeIsw)
	}
	if len(vTypeNisw) > 0 {
		params.SetTypeNisw(vTypeNisw)
	}
	if len(vTypeIew) > 0 {
		params.SetTypeIew(vTypeIew)
	}
	if len(vTypeNiew) > 0 {
		params.SetTypeNiew(vTypeNiew)
	}
	if vTypeEmpty != nil {
		params.SetTypeEmpty(vTypeEmpty)
	}
	if len(vTypeRegex) > 0 {
		params.SetTypeRegex(vTypeRegex)
	}
	if len(vTypeIregex) > 0 {
		params.SetTypeIregex(vTypeIregex)
	}
	if len(vMode) > 0 {
		params.SetMode(vMode)
	}
	if len(vModen) > 0 {
		params.SetModen(vModen)
	}
	if len(vModeIc) > 0 {
		params.SetModeIc(vModeIc)
	}
	if len(vModeNic) > 0 {
		params.SetModeNic(vModeNic)
	}
	if len(vModeIe) > 0 {
		params.SetModeIe(vModeIe)
	}
	if len(vModeNie) > 0 {
		params.SetModeNie(vModeNie)
	}
	if len(vModeIsw) > 0 {
		params.SetModeIsw(vModeIsw)
	}
	if len(vModeNisw) > 0 {
		params.SetModeNisw(vModeNisw)
	}
	if len(vModeIew) > 0 {
		params.SetModeIew(vModeIew)
	}
	if len(vModeNiew) > 0 {
		params.SetModeNiew(vModeNiew)
	}
	if vModeEmpty != nil {
		params.SetModeEmpty(vModeEmpty)
	}
	if len(vModeRegex) > 0 {
		params.SetModeRegex(vModeRegex)
	}
	if len(vModeIregex) > 0 {
		params.SetModeIregex(vModeIregex)
	}
	if vEnabled != nil {
		params.SetEnabled(vEnabled)
	}
	if len(vLagID) > 0 {
		params.SetLagID(vLagID)
	}
	if len(vLagIDn) > 0 {
		params.SetLagIDn(vLagIDn)
	}
	if len(vMacAddress) > 0 {
		params.SetMacAddress(vMacAddress)
	}
	if len(vMacAddressn) > 0 {
		params.SetMacAddressn(vMacAddressn)
	}
	if len(vMacAddressIc) > 0 {
		params.SetMacAddressIc(vMacAddressIc)
	}
	if len(vMacAddressNic) > 0 {
		params.SetMacAddressNic(vMacAddressNic)
	}
	if len(vMacAddressIe) > 0 {
		params.SetMacAddressIe(vMacAddressIe)
	}
	if len(vMacAddressNie) > 0 {
		params.SetMacAddressNie(vMacAddressNie)
	}
	if len(vMacAddressIsw) > 0 {
		params.SetMacAddressIsw(vMacAddressIsw)
	}
	if len(vMacAddressNisw) > 0 {
		params.SetMacAddressNisw(vMacAddressNisw)
	}
	if len(vMacAddressIew) > 0 {
		params.SetMacAddressIew(vMacAddressIew)
	}
	if len(vMacAddressNiew) > 0 {
		params.SetMacAddressNiew(vMacAddressNiew)
	}
	if len(vMacAddressRegex) > 0 {
		params.SetMacAddressRegex(vMacAddressRegex)
	}
	if len(vMacAddressIregex) > 0 {
		params.SetMacAddressIregex(vMacAddressIregex)
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
	if len(vBridgeID) > 0 {
		params.SetBridgeID(vBridgeID)
	}
	if len(vBridgeIDn) > 0 {
		params.SetBridgeIDn(vBridgeIDn)
	}
	if len(vCableConnector) > 0 {
		params.SetCableConnector(vCableConnector)
	}
	if len(vCableConnectorn) > 0 {
		params.SetCableConnectorn(vCableConnectorn)
	}
	if len(vCableConnectorLt) > 0 {
		params.SetCableConnectorLt(vCableConnectorLt)
	}
	if len(vCableConnectorLte) > 0 {
		params.SetCableConnectorLte(vCableConnectorLte)
	}
	if len(vCableConnectorGt) > 0 {
		params.SetCableConnectorGt(vCableConnectorGt)
	}
	if len(vCableConnectorGte) > 0 {
		params.SetCableConnectorGte(vCableConnectorGte)
	}
	if vCableConnectorEmpty != nil {
		params.SetCableConnectorEmpty(vCableConnectorEmpty)
	}
	if vCableEnd != nil {
		params.SetCableEnd(vCableEnd)
	}
	if vCableEndn != nil {
		params.SetCableEndn(vCableEndn)
	}
	if len(vCableEndIc) > 0 {
		params.SetCableEndIc(vCableEndIc)
	}
	if len(vCableEndNic) > 0 {
		params.SetCableEndNic(vCableEndNic)
	}
	if len(vCableEndIe) > 0 {
		params.SetCableEndIe(vCableEndIe)
	}
	if len(vCableEndNie) > 0 {
		params.SetCableEndNie(vCableEndNie)
	}
	if len(vCableEndIsw) > 0 {
		params.SetCableEndIsw(vCableEndIsw)
	}
	if len(vCableEndNisw) > 0 {
		params.SetCableEndNisw(vCableEndNisw)
	}
	if len(vCableEndIew) > 0 {
		params.SetCableEndIew(vCableEndIew)
	}
	if len(vCableEndNiew) > 0 {
		params.SetCableEndNiew(vCableEndNiew)
	}
	if vCableEndEmpty != nil {
		params.SetCableEndEmpty(vCableEndEmpty)
	}
	if len(vCableEndRegex) > 0 {
		params.SetCableEndRegex(vCableEndRegex)
	}
	if len(vCableEndIregex) > 0 {
		params.SetCableEndIregex(vCableEndIregex)
	}
	if len(vCableID) > 0 {
		params.SetCableID(vCableID)
	}
	if len(vCableIDn) > 0 {
		params.SetCableIDn(vCableIDn)
	}
	if vCabled != nil {
		params.SetCabled(vCabled)
	}
	if vConnected != nil {
		params.SetConnected(vConnected)
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
	if len(vDevice) > 0 {
		params.SetDevice(vDevice)
	}
	if len(vDevicen) > 0 {
		params.SetDevicen(vDevicen)
	}
	if len(vDeviceRole) > 0 {
		params.SetDeviceRole(vDeviceRole)
	}
	if len(vDeviceRolen) > 0 {
		params.SetDeviceRolen(vDeviceRolen)
	}
	if len(vDeviceRoleID) > 0 {
		params.SetDeviceRoleID(vDeviceRoleID)
	}
	if len(vDeviceRoleIDn) > 0 {
		params.SetDeviceRoleIDn(vDeviceRoleIDn)
	}
	if len(vDeviceStatus) > 0 {
		params.SetDeviceStatus(vDeviceStatus)
	}
	if len(vDeviceStatusn) > 0 {
		params.SetDeviceStatusn(vDeviceStatusn)
	}
	if len(vDeviceStatusIc) > 0 {
		params.SetDeviceStatusIc(vDeviceStatusIc)
	}
	if len(vDeviceStatusNic) > 0 {
		params.SetDeviceStatusNic(vDeviceStatusNic)
	}
	if len(vDeviceStatusIe) > 0 {
		params.SetDeviceStatusIe(vDeviceStatusIe)
	}
	if len(vDeviceStatusNie) > 0 {
		params.SetDeviceStatusNie(vDeviceStatusNie)
	}
	if len(vDeviceStatusIsw) > 0 {
		params.SetDeviceStatusIsw(vDeviceStatusIsw)
	}
	if len(vDeviceStatusNisw) > 0 {
		params.SetDeviceStatusNisw(vDeviceStatusNisw)
	}
	if len(vDeviceStatusIew) > 0 {
		params.SetDeviceStatusIew(vDeviceStatusIew)
	}
	if len(vDeviceStatusNiew) > 0 {
		params.SetDeviceStatusNiew(vDeviceStatusNiew)
	}
	if vDeviceStatusEmpty != nil {
		params.SetDeviceStatusEmpty(vDeviceStatusEmpty)
	}
	if len(vDeviceStatusRegex) > 0 {
		params.SetDeviceStatusRegex(vDeviceStatusRegex)
	}
	if len(vDeviceStatusIregex) > 0 {
		params.SetDeviceStatusIregex(vDeviceStatusIregex)
	}
	if len(vDeviceType) > 0 {
		params.SetDeviceType(vDeviceType)
	}
	if len(vDeviceTypen) > 0 {
		params.SetDeviceTypen(vDeviceTypen)
	}
	if len(vDeviceTypeID) > 0 {
		params.SetDeviceTypeID(vDeviceTypeID)
	}
	if len(vDeviceTypeIDn) > 0 {
		params.SetDeviceTypeIDn(vDeviceTypeIDn)
	}
	if len(vDuplex) > 0 {
		params.SetDuplex(vDuplex)
	}
	if len(vDuplexn) > 0 {
		params.SetDuplexn(vDuplexn)
	}
	if len(vDuplexIc) > 0 {
		params.SetDuplexIc(vDuplexIc)
	}
	if len(vDuplexNic) > 0 {
		params.SetDuplexNic(vDuplexNic)
	}
	if len(vDuplexIe) > 0 {
		params.SetDuplexIe(vDuplexIe)
	}
	if len(vDuplexNie) > 0 {
		params.SetDuplexNie(vDuplexNie)
	}
	if len(vDuplexIsw) > 0 {
		params.SetDuplexIsw(vDuplexIsw)
	}
	if len(vDuplexNisw) > 0 {
		params.SetDuplexNisw(vDuplexNisw)
	}
	if len(vDuplexIew) > 0 {
		params.SetDuplexIew(vDuplexIew)
	}
	if len(vDuplexNiew) > 0 {
		params.SetDuplexNiew(vDuplexNiew)
	}
	if vDuplexEmpty != nil {
		params.SetDuplexEmpty(vDuplexEmpty)
	}
	if len(vDuplexRegex) > 0 {
		params.SetDuplexRegex(vDuplexRegex)
	}
	if len(vDuplexIregex) > 0 {
		params.SetDuplexIregex(vDuplexIregex)
	}
	if vKind != nil {
		params.SetKind(vKind)
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
	if len(vLabel) > 0 {
		params.SetLabel(vLabel)
	}
	if len(vLabeln) > 0 {
		params.SetLabeln(vLabeln)
	}
	if len(vLabelIc) > 0 {
		params.SetLabelIc(vLabelIc)
	}
	if len(vLabelNic) > 0 {
		params.SetLabelNic(vLabelNic)
	}
	if len(vLabelIe) > 0 {
		params.SetLabelIe(vLabelIe)
	}
	if len(vLabelNie) > 0 {
		params.SetLabelNie(vLabelNie)
	}
	if len(vLabelIsw) > 0 {
		params.SetLabelIsw(vLabelIsw)
	}
	if len(vLabelNisw) > 0 {
		params.SetLabelNisw(vLabelNisw)
	}
	if len(vLabelIew) > 0 {
		params.SetLabelIew(vLabelIew)
	}
	if len(vLabelNiew) > 0 {
		params.SetLabelNiew(vLabelNiew)
	}
	if vLabelEmpty != nil {
		params.SetLabelEmpty(vLabelEmpty)
	}
	if len(vLabelRegex) > 0 {
		params.SetLabelRegex(vLabelRegex)
	}
	if len(vLabelIregex) > 0 {
		params.SetLabelIregex(vLabelIregex)
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
	if vMarkConnected != nil {
		params.SetMarkConnected(vMarkConnected)
	}
	if vMgmtOnly != nil {
		params.SetMgmtOnly(vMgmtOnly)
	}
	if len(vModuleID) > 0 {
		params.SetModuleID(vModuleID)
	}
	if len(vModuleIDn) > 0 {
		params.SetModuleIDn(vModuleIDn)
	}
	if len(vMtu) > 0 {
		params.SetMtu(vMtu)
	}
	if len(vMtun) > 0 {
		params.SetMtun(vMtun)
	}
	if len(vMtuLt) > 0 {
		params.SetMtuLt(vMtuLt)
	}
	if len(vMtuLte) > 0 {
		params.SetMtuLte(vMtuLte)
	}
	if len(vMtuGt) > 0 {
		params.SetMtuGt(vMtuGt)
	}
	if len(vMtuGte) > 0 {
		params.SetMtuGte(vMtuGte)
	}
	if vMtuEmpty != nil {
		params.SetMtuEmpty(vMtuEmpty)
	}
	if vOccupied != nil {
		params.SetOccupied(vOccupied)
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
	if len(vParentID) > 0 {
		params.SetParentID(vParentID)
	}
	if len(vParentIDn) > 0 {
		params.SetParentIDn(vParentIDn)
	}
	if len(vPoeMode) > 0 {
		params.SetPoeMode(vPoeMode)
	}
	if len(vPoeModen) > 0 {
		params.SetPoeModen(vPoeModen)
	}
	if len(vPoeModeIc) > 0 {
		params.SetPoeModeIc(vPoeModeIc)
	}
	if len(vPoeModeNic) > 0 {
		params.SetPoeModeNic(vPoeModeNic)
	}
	if len(vPoeModeIe) > 0 {
		params.SetPoeModeIe(vPoeModeIe)
	}
	if len(vPoeModeNie) > 0 {
		params.SetPoeModeNie(vPoeModeNie)
	}
	if len(vPoeModeIsw) > 0 {
		params.SetPoeModeIsw(vPoeModeIsw)
	}
	if len(vPoeModeNisw) > 0 {
		params.SetPoeModeNisw(vPoeModeNisw)
	}
	if len(vPoeModeIew) > 0 {
		params.SetPoeModeIew(vPoeModeIew)
	}
	if len(vPoeModeNiew) > 0 {
		params.SetPoeModeNiew(vPoeModeNiew)
	}
	if vPoeModeEmpty != nil {
		params.SetPoeModeEmpty(vPoeModeEmpty)
	}
	if len(vPoeModeRegex) > 0 {
		params.SetPoeModeRegex(vPoeModeRegex)
	}
	if len(vPoeModeIregex) > 0 {
		params.SetPoeModeIregex(vPoeModeIregex)
	}
	if len(vPoeType) > 0 {
		params.SetPoeType(vPoeType)
	}
	if len(vPoeTypen) > 0 {
		params.SetPoeTypen(vPoeTypen)
	}
	if len(vPoeTypeIc) > 0 {
		params.SetPoeTypeIc(vPoeTypeIc)
	}
	if len(vPoeTypeNic) > 0 {
		params.SetPoeTypeNic(vPoeTypeNic)
	}
	if len(vPoeTypeIe) > 0 {
		params.SetPoeTypeIe(vPoeTypeIe)
	}
	if len(vPoeTypeNie) > 0 {
		params.SetPoeTypeNie(vPoeTypeNie)
	}
	if len(vPoeTypeIsw) > 0 {
		params.SetPoeTypeIsw(vPoeTypeIsw)
	}
	if len(vPoeTypeNisw) > 0 {
		params.SetPoeTypeNisw(vPoeTypeNisw)
	}
	if len(vPoeTypeIew) > 0 {
		params.SetPoeTypeIew(vPoeTypeIew)
	}
	if len(vPoeTypeNiew) > 0 {
		params.SetPoeTypeNiew(vPoeTypeNiew)
	}
	if vPoeTypeEmpty != nil {
		params.SetPoeTypeEmpty(vPoeTypeEmpty)
	}
	if len(vPoeTypeRegex) > 0 {
		params.SetPoeTypeRegex(vPoeTypeRegex)
	}
	if len(vPoeTypeIregex) > 0 {
		params.SetPoeTypeIregex(vPoeTypeIregex)
	}
	if len(vPrimaryMacAddress) > 0 {
		params.SetPrimaryMacAddress(vPrimaryMacAddress)
	}
	if len(vPrimaryMacAddressn) > 0 {
		params.SetPrimaryMacAddressn(vPrimaryMacAddressn)
	}
	if len(vPrimaryMacAddressID) > 0 {
		params.SetPrimaryMacAddressID(vPrimaryMacAddressID)
	}
	if len(vPrimaryMacAddressIDn) > 0 {
		params.SetPrimaryMacAddressIDn(vPrimaryMacAddressIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRack) > 0 {
		params.SetRack(vRack)
	}
	if len(vRackn) > 0 {
		params.SetRackn(vRackn)
	}
	if len(vRackID) > 0 {
		params.SetRackID(vRackID)
	}
	if len(vRackIDn) > 0 {
		params.SetRackIDn(vRackIDn)
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
	if len(vRfChannel) > 0 {
		params.SetRfChannel(vRfChannel)
	}
	if len(vRfChanneln) > 0 {
		params.SetRfChanneln(vRfChanneln)
	}
	if len(vRfChannelIc) > 0 {
		params.SetRfChannelIc(vRfChannelIc)
	}
	if len(vRfChannelNic) > 0 {
		params.SetRfChannelNic(vRfChannelNic)
	}
	if len(vRfChannelIe) > 0 {
		params.SetRfChannelIe(vRfChannelIe)
	}
	if len(vRfChannelNie) > 0 {
		params.SetRfChannelNie(vRfChannelNie)
	}
	if len(vRfChannelIsw) > 0 {
		params.SetRfChannelIsw(vRfChannelIsw)
	}
	if len(vRfChannelNisw) > 0 {
		params.SetRfChannelNisw(vRfChannelNisw)
	}
	if len(vRfChannelIew) > 0 {
		params.SetRfChannelIew(vRfChannelIew)
	}
	if len(vRfChannelNiew) > 0 {
		params.SetRfChannelNiew(vRfChannelNiew)
	}
	if vRfChannelEmpty != nil {
		params.SetRfChannelEmpty(vRfChannelEmpty)
	}
	if len(vRfChannelRegex) > 0 {
		params.SetRfChannelRegex(vRfChannelRegex)
	}
	if len(vRfChannelIregex) > 0 {
		params.SetRfChannelIregex(vRfChannelIregex)
	}
	if len(vRfChannelFrequency) > 0 {
		params.SetRfChannelFrequency(vRfChannelFrequency)
	}
	if len(vRfChannelFrequencyn) > 0 {
		params.SetRfChannelFrequencyn(vRfChannelFrequencyn)
	}
	if len(vRfChannelFrequencyLt) > 0 {
		params.SetRfChannelFrequencyLt(vRfChannelFrequencyLt)
	}
	if len(vRfChannelFrequencyLte) > 0 {
		params.SetRfChannelFrequencyLte(vRfChannelFrequencyLte)
	}
	if len(vRfChannelFrequencyGt) > 0 {
		params.SetRfChannelFrequencyGt(vRfChannelFrequencyGt)
	}
	if len(vRfChannelFrequencyGte) > 0 {
		params.SetRfChannelFrequencyGte(vRfChannelFrequencyGte)
	}
	if vRfChannelFrequencyEmpty != nil {
		params.SetRfChannelFrequencyEmpty(vRfChannelFrequencyEmpty)
	}
	if len(vRfChannelWidth) > 0 {
		params.SetRfChannelWidth(vRfChannelWidth)
	}
	if len(vRfChannelWidthn) > 0 {
		params.SetRfChannelWidthn(vRfChannelWidthn)
	}
	if len(vRfChannelWidthLt) > 0 {
		params.SetRfChannelWidthLt(vRfChannelWidthLt)
	}
	if len(vRfChannelWidthLte) > 0 {
		params.SetRfChannelWidthLte(vRfChannelWidthLte)
	}
	if len(vRfChannelWidthGt) > 0 {
		params.SetRfChannelWidthGt(vRfChannelWidthGt)
	}
	if len(vRfChannelWidthGte) > 0 {
		params.SetRfChannelWidthGte(vRfChannelWidthGte)
	}
	if vRfChannelWidthEmpty != nil {
		params.SetRfChannelWidthEmpty(vRfChannelWidthEmpty)
	}
	if len(vRfRole) > 0 {
		params.SetRfRole(vRfRole)
	}
	if len(vRfRolen) > 0 {
		params.SetRfRolen(vRfRolen)
	}
	if len(vRfRoleIc) > 0 {
		params.SetRfRoleIc(vRfRoleIc)
	}
	if len(vRfRoleNic) > 0 {
		params.SetRfRoleNic(vRfRoleNic)
	}
	if len(vRfRoleIe) > 0 {
		params.SetRfRoleIe(vRfRoleIe)
	}
	if len(vRfRoleNie) > 0 {
		params.SetRfRoleNie(vRfRoleNie)
	}
	if len(vRfRoleIsw) > 0 {
		params.SetRfRoleIsw(vRfRoleIsw)
	}
	if len(vRfRoleNisw) > 0 {
		params.SetRfRoleNisw(vRfRoleNisw)
	}
	if len(vRfRoleIew) > 0 {
		params.SetRfRoleIew(vRfRoleIew)
	}
	if len(vRfRoleNiew) > 0 {
		params.SetRfRoleNiew(vRfRoleNiew)
	}
	if vRfRoleEmpty != nil {
		params.SetRfRoleEmpty(vRfRoleEmpty)
	}
	if len(vRfRoleRegex) > 0 {
		params.SetRfRoleRegex(vRfRoleRegex)
	}
	if len(vRfRoleIregex) > 0 {
		params.SetRfRoleIregex(vRfRoleIregex)
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
	if len(vSpeed) > 0 {
		params.SetSpeed(vSpeed)
	}
	if len(vSpeedn) > 0 {
		params.SetSpeedn(vSpeedn)
	}
	if len(vSpeedLt) > 0 {
		params.SetSpeedLt(vSpeedLt)
	}
	if len(vSpeedLte) > 0 {
		params.SetSpeedLte(vSpeedLte)
	}
	if len(vSpeedGt) > 0 {
		params.SetSpeedGt(vSpeedGt)
	}
	if len(vSpeedGte) > 0 {
		params.SetSpeedGte(vSpeedGte)
	}
	if len(vSpeedEmpty) > 0 {
		params.SetSpeedEmpty(vSpeedEmpty)
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
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vTxPower) > 0 {
		params.SetTxPower(vTxPower)
	}
	if len(vTxPowern) > 0 {
		params.SetTxPowern(vTxPowern)
	}
	if len(vTxPowerLt) > 0 {
		params.SetTxPowerLt(vTxPowerLt)
	}
	if len(vTxPowerLte) > 0 {
		params.SetTxPowerLte(vTxPowerLte)
	}
	if len(vTxPowerGt) > 0 {
		params.SetTxPowerGt(vTxPowerGt)
	}
	if len(vTxPowerGte) > 0 {
		params.SetTxPowerGte(vTxPowerGte)
	}
	if vTxPowerEmpty != nil {
		params.SetTxPowerEmpty(vTxPowerEmpty)
	}
	if len(vVdc) > 0 {
		params.SetVdc(vVdc)
	}
	if len(vVdcn) > 0 {
		params.SetVdcn(vVdcn)
	}
	if len(vVdcID) > 0 {
		params.SetVdcID(vVdcID)
	}
	if len(vVdcIDn) > 0 {
		params.SetVdcIDn(vVdcIDn)
	}
	if len(vVdcIdentifier) > 0 {
		params.SetVdcIdentifier(vVdcIdentifier)
	}
	if len(vVdcIdentifiern) > 0 {
		params.SetVdcIdentifiern(vVdcIdentifiern)
	}
	if len(vVirtualChassis) > 0 {
		params.SetVirtualChassis(vVirtualChassis)
	}
	if len(vVirtualChassisn) > 0 {
		params.SetVirtualChassisn(vVirtualChassisn)
	}
	if len(vVirtualChassisID) > 0 {
		params.SetVirtualChassisID(vVirtualChassisID)
	}
	if len(vVirtualChassisIDn) > 0 {
		params.SetVirtualChassisIDn(vVirtualChassisIDn)
	}
	if len(vVirtualChassisMember) > 0 {
		params.SetVirtualChassisMember(vVirtualChassisMember)
	}
	if len(vVirtualChassisMemberID) > 0 {
		params.SetVirtualChassisMemberID(vVirtualChassisMemberID)
	}
	if len(vVirtualChassisMemberOrMaster) > 0 {
		params.SetVirtualChassisMemberOrMaster(vVirtualChassisMemberOrMaster)
	}
	if len(vVirtualChassisMemberOrMasterID) > 0 {
		params.SetVirtualChassisMemberOrMasterID(vVirtualChassisMemberOrMasterID)
	}
	if len(vVirtualCircuitID) > 0 {
		params.SetVirtualCircuitID(vVirtualCircuitID)
	}
	if len(vVirtualCircuitIDn) > 0 {
		params.SetVirtualCircuitIDn(vVirtualCircuitIDn)
	}
	if len(vVirtualCircuitTerminationID) > 0 {
		params.SetVirtualCircuitTerminationID(vVirtualCircuitTerminationID)
	}
	if len(vVirtualCircuitTerminationIDn) > 0 {
		params.SetVirtualCircuitTerminationIDn(vVirtualCircuitTerminationIDn)
	}
	if vVlan != nil {
		params.SetVlan(vVlan)
	}
	if vVlanID != nil {
		params.SetVlanID(vVlanID)
	}
	if len(vVlanTranslationPolicy) > 0 {
		params.SetVlanTranslationPolicy(vVlanTranslationPolicy)
	}
	if len(vVlanTranslationPolicyn) > 0 {
		params.SetVlanTranslationPolicyn(vVlanTranslationPolicyn)
	}
	if len(vVlanTranslationPolicyID) > 0 {
		params.SetVlanTranslationPolicyID(vVlanTranslationPolicyID)
	}
	if len(vVlanTranslationPolicyIDn) > 0 {
		params.SetVlanTranslationPolicyIDn(vVlanTranslationPolicyIDn)
	}
	if len(vVrf) > 0 {
		params.SetVrf(vVrf)
	}
	if len(vVrfn) > 0 {
		params.SetVrfn(vVrfn)
	}
	if len(vVrfID) > 0 {
		params.SetVrfID(vVrfID)
	}
	if len(vVrfIDn) > 0 {
		params.SetVrfIDn(vVrfIDn)
	}
	if len(vWirelessLanID) > 0 {
		params.SetWirelessLanID(vWirelessLanID)
	}
	if len(vWirelessLanIDn) > 0 {
		params.SetWirelessLanIDn(vWirelessLanIDn)
	}
	if len(vWirelessLinkID) > 0 {
		params.SetWirelessLinkID(vWirelessLinkID)
	}
	if len(vWirelessLinkIDn) > 0 {
		params.SetWirelessLinkIDn(vWirelessLinkIDn)
	}
	if len(vWwn) > 0 {
		params.SetWwn(vWwn)
	}
	if len(vWwnn) > 0 {
		params.SetWwnn(vWwnn)
	}
	if len(vWwnIc) > 0 {
		params.SetWwnIc(vWwnIc)
	}
	if len(vWwnNic) > 0 {
		params.SetWwnNic(vWwnNic)
	}
	if len(vWwnIe) > 0 {
		params.SetWwnIe(vWwnIe)
	}
	if len(vWwnNie) > 0 {
		params.SetWwnNie(vWwnNie)
	}
	if len(vWwnIsw) > 0 {
		params.SetWwnIsw(vWwnIsw)
	}
	if len(vWwnNisw) > 0 {
		params.SetWwnNisw(vWwnNisw)
	}
	if len(vWwnIew) > 0 {
		params.SetWwnIew(vWwnIew)
	}
	if len(vWwnNiew) > 0 {
		params.SetWwnNiew(vWwnNiew)
	}
	if len(vWwnRegex) > 0 {
		params.SetWwnRegex(vWwnRegex)
	}
	if len(vWwnIregex) > 0 {
		params.SetWwnIregex(vWwnIregex)
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
