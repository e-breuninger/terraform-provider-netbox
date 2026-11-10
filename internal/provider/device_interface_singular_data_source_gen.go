// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*deviceInterfaceDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceInterfaceDataSource)(nil)
)

// NewDeviceInterfaceDataSource returns a new device_interface data source.
func NewDeviceInterfaceDataSource() datasource.DataSource {
	return &deviceInterfaceDataSource{}
}

type deviceInterfaceDataSource struct {
	client *netboxapi.Client
}

// deviceInterfaceDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type deviceInterfaceDataSourceModel struct {
	deviceInterfaceResourceModel
	MacAddress                     types.String `tfsdk:"mac_address"`
	TagSlug                        types.String `tfsdk:"tag_slug"`
	CableConnector                 types.Int64  `tfsdk:"cable_connector"`
	CableEnd                       types.String `tfsdk:"cable_end"`
	CableID                        types.Int64  `tfsdk:"cable_id"`
	Cabled                         types.Bool   `tfsdk:"cabled"`
	Connected                      types.Bool   `tfsdk:"connected"`
	DeviceName                     types.String `tfsdk:"device_name"`
	DeviceRoleSlug                 types.String `tfsdk:"device_role_slug"`
	DeviceRoleID                   types.Int64  `tfsdk:"device_role_id"`
	DeviceStatus                   types.String `tfsdk:"device_status"`
	DeviceTypeModel                types.String `tfsdk:"device_type_model"`
	DeviceTypeID                   types.Int64  `tfsdk:"device_type_id"`
	Kind                           types.String `tfsdk:"kind"`
	L2vpn                          types.Int64  `tfsdk:"l2vpn"`
	L2vpnID                        types.Int64  `tfsdk:"l2vpn_id"`
	LocationSlug                   types.String `tfsdk:"location_slug"`
	LocationID                     types.Int64  `tfsdk:"location_id"`
	Occupied                       types.Bool   `tfsdk:"occupied"`
	OwnerName                      types.String `tfsdk:"owner_name"`
	OwnerGroupName                 types.String `tfsdk:"owner_group_name"`
	OwnerGroupID                   types.Int64  `tfsdk:"owner_group_id"`
	ParentID                       types.Int64  `tfsdk:"parent_id"`
	PrimaryMacAddress              types.String `tfsdk:"primary_mac_address"`
	PrimaryMacAddressID            types.Int64  `tfsdk:"primary_mac_address_id"`
	Q                              types.String `tfsdk:"q"`
	RackName                       types.String `tfsdk:"rack_name"`
	RackID                         types.Int64  `tfsdk:"rack_id"`
	RegionSlug                     types.String `tfsdk:"region_slug"`
	RegionID                       types.String `tfsdk:"region_id"`
	SiteSlug                       types.String `tfsdk:"site_slug"`
	SiteGroupSlug                  types.String `tfsdk:"site_group_slug"`
	SiteGroupID                    types.String `tfsdk:"site_group_id"`
	SiteID                         types.Int64  `tfsdk:"site_id"`
	TagID                          types.Int64  `tfsdk:"tag_id"`
	TenantSlug                     types.String `tfsdk:"tenant_slug"`
	TenantID                       types.Int64  `tfsdk:"tenant_id"`
	Vdc                            types.String `tfsdk:"vdc"`
	VdcID                          types.Int64  `tfsdk:"vdc_id"`
	VdcIdentifier                  types.Int64  `tfsdk:"vdc_identifier"`
	VirtualChassis                 types.String `tfsdk:"virtual_chassis"`
	VirtualChassisID               types.Int64  `tfsdk:"virtual_chassis_id"`
	VirtualChassisMember           types.String `tfsdk:"virtual_chassis_member"`
	VirtualChassisMemberID         types.Int64  `tfsdk:"virtual_chassis_member_id"`
	VirtualChassisMemberOrMaster   types.String `tfsdk:"virtual_chassis_member_or_master"`
	VirtualChassisMemberOrMasterID types.Int64  `tfsdk:"virtual_chassis_member_or_master_id"`
	VirtualCircuitID               types.Int64  `tfsdk:"virtual_circuit_id"`
	VirtualCircuitTerminationID    types.Int64  `tfsdk:"virtual_circuit_termination_id"`
	Vlan                           types.String `tfsdk:"vlan"`
	VlanID                         types.String `tfsdk:"vlan_id"`
	VlanTranslationPolicy          types.String `tfsdk:"vlan_translation_policy"`
	Vrf                            types.String `tfsdk:"vrf"`
	WirelessLanID                  types.Int64  `tfsdk:"wireless_lan_id"`
	WirelessLinkID                 types.Int64  `tfsdk:"wireless_link_id"`
	CustomFieldFilters             types.Map    `tfsdk:"custom_field_filters"`
}

func (d *deviceInterfaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_interface"
}

func (d *deviceInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A device interface (dcim.interface).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/interface/):\n\n> Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface type as its NetBox slug, e.g. virtual, lag, bridge, 1000base-t, 10gbase-x-sfpp (any of NetBox's interface type choices).",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the interface is enabled.",
			},
			"mgmt_only": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the interface is used for out-of-band management only.",
			},
			"mgmtonly": schema.BoolAttribute{
				Computed:           true,
				DeprecationMessage: "Use mgmt_only instead.",
				Description:        "Deprecated alias of mgmt_only.",
			},
			"mark_connected": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat as if a cable is connected.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(1, 65536),
				},
			},
			"speed": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Speed in kbps. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"duplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "One of: half, full, auto.",
				Validators: []validator.String{
					stringvalidator.OneOf("half", "full", "auto"),
				},
			},
			"poe_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Power over Ethernet mode. One of: pd, pse.",
				Validators: []validator.String{
					stringvalidator.OneOf("pd", "pse"),
				},
			},
			"poe_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Power over Ethernet type as its NetBox slug, e.g. type1-ieee802.3af, type3-ieee802.3bt, passive-24v-2pair (any of NetBox's PoE type choices).",
			},
			"rf_role": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Wireless role. One of: ap, station.",
				Validators: []validator.String{
					stringvalidator.OneOf("ap", "station"),
				},
			},
			"rf_channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Wireless channel as its NetBox slug, e.g. 2.4g-1-2412-22, 5g-36-5180-20 (any of NetBox's wireless channel choices).",
			},
			"rf_channel_frequency": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Channel frequency in MHz. NetBox derives it from rf_channel when that is set.",
				Validators: []validator.Float64{
					float64validator.Between(-100000, 100000),
				},
			},
			"rf_channel_width": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Channel width in MHz. NetBox derives it from rf_channel when that is set.",
				Validators: []validator.Float64{
					float64validator.Between(-10000, 10000),
				},
			},
			"tx_power": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Transmit power in dBm. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(-40, 127),
				},
			},
			"wwn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "64-bit World Wide Name.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "802.1Q tagging mode. One of: access, tagged, tagged-all, q-in-q.",
				Validators: []validator.String{
					stringvalidator.OneOf("access", "tagged", "tagged-all", "q-in-q"),
				},
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
				Optional:    true,
				Computed:    true,
				Description: "Id of the VLAN translation policy.",
			},
			"module_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the installed module this interface belongs to.",
			},
			"lag_device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the parent LAG interface.",
			},
			"parent_device_interface_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the parent interface.",
			},
			"bridge_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the bridge interface.",
			},
			"vrf_id": schema.Int64Attribute{
				Optional:    true,
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
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Optional: true,
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
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Description: "MAC address assigned to the interface.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"cable_connector": schema.Int64Attribute{
				Optional:    true,
				Description: "Cable connector to match.",
			},
			"cable_end": schema.StringAttribute{
				Optional:    true,
				Description: "Cable end. One of: A, B.",
			},
			"cable_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cable.",
			},
			"cabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether cabled is set.",
			},
			"connected": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether connected is set.",
			},
			"device_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the device.",
			},
			"device_role_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the device role.",
			},
			"device_role_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device role.",
			},
			"device_status": schema.StringAttribute{
				Optional:    true,
				Description: "Device status to match.",
			},
			"device_type_model": schema.StringAttribute{
				Optional:    true,
				Description: "Model of the device type.",
			},
			"device_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device type.",
			},
			"kind": schema.StringAttribute{
				Optional:    true,
				Description: "Kind of interface.",
			},
			"l2vpn": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the L2VPN.",
			},
			"l2vpn_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the L2VPN.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the location.",
			},
			"occupied": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether occupied is set.",
			},
			"owner_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner.",
			},
			"owner_group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner group.",
			},
			"owner_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner group.",
			},
			"parent_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the parent interface.",
			},
			"primary_mac_address": schema.StringAttribute{
				Optional:    true,
				Description: "Primary MAC address.",
			},
			"primary_mac_address_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the primary MAC address.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"rack_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the rack.",
			},
			"rack_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the rack.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
			},
			"region_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"site_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site.",
			},
			"site_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site group.",
			},
			"site_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the site group.",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"tenant_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant.",
			},
			"vdc": schema.StringAttribute{
				Optional:    true,
				Description: "Virtual Device Context.",
			},
			"vdc_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual device context.",
			},
			"vdc_identifier": schema.Int64Attribute{
				Optional:    true,
				Description: "Virtual Device Context (Identifier).",
			},
			"virtual_chassis": schema.StringAttribute{
				Optional:    true,
				Description: "Virtual Chassis.",
			},
			"virtual_chassis_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual chassis.",
			},
			"virtual_chassis_member": schema.StringAttribute{
				Optional:    true,
				Description: "Virtual chassis member to match.",
			},
			"virtual_chassis_member_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual chassis member.",
			},
			"virtual_chassis_member_or_master": schema.StringAttribute{
				Optional:    true,
				Description: "Virtual chassis member or master to match.",
			},
			"virtual_chassis_member_or_master_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual chassis member or master.",
			},
			"virtual_circuit_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual circuit.",
			},
			"virtual_circuit_termination_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual circuit termination.",
			},
			"vlan": schema.StringAttribute{
				Optional:    true,
				Description: "Assigned VID.",
			},
			"vlan_id": schema.StringAttribute{
				Optional:    true,
				Description: "Assigned VLAN.",
			},
			"vlan_translation_policy": schema.StringAttribute{
				Optional:    true,
				Description: "VLAN Translation Policy.",
			},
			"vrf": schema.StringAttribute{
				Optional:    true,
				Description: "VRF (RD).",
			},
			"wireless_lan_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the wireless LAN.",
			},
			"wireless_link_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the wireless link.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *deviceInterfaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceInterfaceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimInterfacesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.Mode.IsNull() {
		v := state.Mode.ValueString()
		params.SetMode([]string{v})
		hasInput = true
	}
	if !state.Enabled.IsNull() {
		v := state.Enabled.ValueBool()
		params.SetEnabled(&v)
		hasInput = true
	}
	if !state.LagDeviceInterfaceID.IsNull() {
		v := state.LagDeviceInterfaceID.ValueInt64()
		params.SetLagID([]int64{v})
		hasInput = true
	}
	if !state.MacAddress.IsNull() {
		v := state.MacAddress.ValueString()
		params.SetMacAddress([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
		hasInput = true
	}
	if !state.BridgeID.IsNull() {
		v := state.BridgeID.ValueInt64()
		params.SetBridgeID([]int64{v})
		hasInput = true
	}
	if !state.CableConnector.IsNull() {
		v := state.CableConnector.ValueInt64()
		params.SetCableConnector([]int64{v})
		hasInput = true
	}
	if !state.CableEnd.IsNull() {
		v := state.CableEnd.ValueString()
		params.SetCableEnd(&v)
		hasInput = true
	}
	if !state.CableID.IsNull() {
		v := state.CableID.ValueInt64()
		params.SetCableID([]int64{v})
		hasInput = true
	}
	if !state.Cabled.IsNull() {
		v := state.Cabled.ValueBool()
		params.SetCabled(&v)
		hasInput = true
	}
	if !state.Connected.IsNull() {
		v := state.Connected.ValueBool()
		params.SetConnected(&v)
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.DeviceName.IsNull() {
		v := state.DeviceName.ValueString()
		params.SetDevice([]string{v})
		hasInput = true
	}
	if !state.DeviceRoleSlug.IsNull() {
		v := state.DeviceRoleSlug.ValueString()
		params.SetDeviceRole([]string{v})
		hasInput = true
	}
	if !state.DeviceRoleID.IsNull() {
		v := state.DeviceRoleID.ValueInt64()
		params.SetDeviceRoleID([]int64{v})
		hasInput = true
	}
	if !state.DeviceStatus.IsNull() {
		v := state.DeviceStatus.ValueString()
		params.SetDeviceStatus([]string{v})
		hasInput = true
	}
	if !state.DeviceTypeModel.IsNull() {
		v := state.DeviceTypeModel.ValueString()
		params.SetDeviceType([]string{v})
		hasInput = true
	}
	if !state.DeviceTypeID.IsNull() {
		v := state.DeviceTypeID.ValueInt64()
		params.SetDeviceTypeID([]int64{v})
		hasInput = true
	}
	if !state.Duplex.IsNull() {
		v := state.Duplex.ValueString()
		params.SetDuplex([]string{v})
		hasInput = true
	}
	if !state.Kind.IsNull() {
		v := state.Kind.ValueString()
		params.SetKind(&v)
		hasInput = true
	}
	if !state.L2vpn.IsNull() {
		v := state.L2vpn.ValueInt64()
		params.SetL2vpn([]int64{v})
		hasInput = true
	}
	if !state.L2vpnID.IsNull() {
		v := state.L2vpnID.ValueInt64()
		params.SetL2vpnID([]int64{v})
		hasInput = true
	}
	if !state.Label.IsNull() {
		v := state.Label.ValueString()
		params.SetLabel([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.LocationSlug.IsNull() {
		v := state.LocationSlug.ValueString()
		params.SetLocation([]string{v})
		hasInput = true
	}
	if !state.LocationID.IsNull() {
		v := state.LocationID.ValueInt64()
		params.SetLocationID([]int64{v})
		hasInput = true
	}
	if !state.MarkConnected.IsNull() {
		v := state.MarkConnected.ValueBool()
		params.SetMarkConnected(&v)
		hasInput = true
	}
	if !state.MgmtOnly.IsNull() {
		v := state.MgmtOnly.ValueBool()
		params.SetMgmtOnly(&v)
		hasInput = true
	}
	if !state.ModuleID.IsNull() {
		v := state.ModuleID.ValueInt64()
		params.SetModuleID([]int64{v})
		hasInput = true
	}
	if !state.Mtu.IsNull() {
		v := state.Mtu.ValueInt64()
		params.SetMtu([]int64{v})
		hasInput = true
	}
	if !state.Occupied.IsNull() {
		v := state.Occupied.ValueBool()
		params.SetOccupied(&v)
		hasInput = true
	}
	if !state.OwnerName.IsNull() {
		v := state.OwnerName.ValueString()
		params.SetOwner([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupName.IsNull() {
		v := state.OwnerGroupName.ValueString()
		params.SetOwnerGroup([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupID.IsNull() {
		v := state.OwnerGroupID.ValueInt64()
		params.SetOwnerGroupID([]int64{v})
		hasInput = true
	}
	if !state.ParentID.IsNull() {
		v := state.ParentID.ValueInt64()
		params.SetParentID([]int64{v})
		hasInput = true
	}
	if !state.PoeMode.IsNull() {
		v := state.PoeMode.ValueString()
		params.SetPoeMode([]string{v})
		hasInput = true
	}
	if !state.PoeType.IsNull() {
		v := state.PoeType.ValueString()
		params.SetPoeType([]string{v})
		hasInput = true
	}
	if !state.PrimaryMacAddress.IsNull() {
		v := state.PrimaryMacAddress.ValueString()
		params.SetPrimaryMacAddress([]string{v})
		hasInput = true
	}
	if !state.PrimaryMacAddressID.IsNull() {
		v := state.PrimaryMacAddressID.ValueInt64()
		params.SetPrimaryMacAddressID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RackName.IsNull() {
		v := state.RackName.ValueString()
		params.SetRack([]string{v})
		hasInput = true
	}
	if !state.RackID.IsNull() {
		v := state.RackID.ValueInt64()
		params.SetRackID([]int64{v})
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueString()
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.RfChannel.IsNull() {
		v := state.RfChannel.ValueString()
		params.SetRfChannel([]string{v})
		hasInput = true
	}
	if !state.RfChannelFrequency.IsNull() {
		v := state.RfChannelFrequency.ValueFloat64()
		params.SetRfChannelFrequency([]float64{v})
		hasInput = true
	}
	if !state.RfChannelWidth.IsNull() {
		v := state.RfChannelWidth.ValueFloat64()
		params.SetRfChannelWidth([]float64{v})
		hasInput = true
	}
	if !state.RfRole.IsNull() {
		v := state.RfRole.ValueString()
		params.SetRfRole([]string{v})
		hasInput = true
	}
	if !state.SiteSlug.IsNull() {
		v := state.SiteSlug.ValueString()
		params.SetSite([]string{v})
		hasInput = true
	}
	if !state.SiteGroupSlug.IsNull() {
		v := state.SiteGroupSlug.ValueString()
		params.SetSiteGroup([]string{v})
		hasInput = true
	}
	if !state.SiteGroupID.IsNull() {
		v := state.SiteGroupID.ValueString()
		params.SetSiteGroupID([]string{v})
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.Speed.IsNull() {
		v := state.Speed.ValueInt64()
		params.SetSpeed([]int64{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.TenantSlug.IsNull() {
		v := state.TenantSlug.ValueString()
		params.SetTenant([]string{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.TxPower.IsNull() {
		v := state.TxPower.ValueInt64()
		params.SetTxPower([]int64{v})
		hasInput = true
	}
	if !state.Vdc.IsNull() {
		v := state.Vdc.ValueString()
		params.SetVdc([]string{v})
		hasInput = true
	}
	if !state.VdcID.IsNull() {
		v := state.VdcID.ValueInt64()
		params.SetVdcID([]int64{v})
		hasInput = true
	}
	if !state.VdcIdentifier.IsNull() {
		v := state.VdcIdentifier.ValueInt64()
		params.SetVdcIdentifier([]int64{v})
		hasInput = true
	}
	if !state.VirtualChassis.IsNull() {
		v := state.VirtualChassis.ValueString()
		params.SetVirtualChassis([]string{v})
		hasInput = true
	}
	if !state.VirtualChassisID.IsNull() {
		v := state.VirtualChassisID.ValueInt64()
		params.SetVirtualChassisID([]int64{v})
		hasInput = true
	}
	if !state.VirtualChassisMember.IsNull() {
		v := state.VirtualChassisMember.ValueString()
		params.SetVirtualChassisMember([]string{v})
		hasInput = true
	}
	if !state.VirtualChassisMemberID.IsNull() {
		v := state.VirtualChassisMemberID.ValueInt64()
		params.SetVirtualChassisMemberID([]int64{v})
		hasInput = true
	}
	if !state.VirtualChassisMemberOrMaster.IsNull() {
		v := state.VirtualChassisMemberOrMaster.ValueString()
		params.SetVirtualChassisMemberOrMaster([]string{v})
		hasInput = true
	}
	if !state.VirtualChassisMemberOrMasterID.IsNull() {
		v := state.VirtualChassisMemberOrMasterID.ValueInt64()
		params.SetVirtualChassisMemberOrMasterID([]int64{v})
		hasInput = true
	}
	if !state.VirtualCircuitID.IsNull() {
		v := state.VirtualCircuitID.ValueInt64()
		params.SetVirtualCircuitID([]int64{v})
		hasInput = true
	}
	if !state.VirtualCircuitTerminationID.IsNull() {
		v := state.VirtualCircuitTerminationID.ValueInt64()
		params.SetVirtualCircuitTerminationID([]int64{v})
		hasInput = true
	}
	if !state.Vlan.IsNull() {
		v := state.Vlan.ValueString()
		params.SetVlan(&v)
		hasInput = true
	}
	if !state.VlanID.IsNull() {
		v := state.VlanID.ValueString()
		params.SetVlanID(&v)
		hasInput = true
	}
	if !state.VlanTranslationPolicy.IsNull() {
		v := state.VlanTranslationPolicy.ValueString()
		params.SetVlanTranslationPolicy([]string{v})
		hasInput = true
	}
	if !state.VlanTranslationPolicyID.IsNull() {
		v := state.VlanTranslationPolicyID.ValueInt64()
		params.SetVlanTranslationPolicyID([]int64{v})
		hasInput = true
	}
	if !state.Vrf.IsNull() {
		v := state.Vrf.ValueString()
		params.SetVrf([]string{v})
		hasInput = true
	}
	if !state.VrfID.IsNull() {
		v := state.VrfID.ValueInt64()
		params.SetVrfID([]int64{v})
		hasInput = true
	}
	if !state.WirelessLanID.IsNull() {
		v := state.WirelessLanID.ValueInt64()
		params.SetWirelessLanID([]int64{v})
		hasInput = true
	}
	if !state.WirelessLinkID.IsNull() {
		v := state.WirelessLinkID.ValueInt64()
		params.SetWirelessLinkID([]int64{v})
		hasInput = true
	}
	if !state.Wwn.IsNull() {
		v := state.Wwn.ValueString()
		params.SetWwn([]string{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_device_interface.")
		return
	}
	res, err := d.client.Dcim.DcimInterfacesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device_interface", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_device_interface",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenDeviceInterface(ctx, netboxapi.DeviceInterfaceResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.deviceInterfaceResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
