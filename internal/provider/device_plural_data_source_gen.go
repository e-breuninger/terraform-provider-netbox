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
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*devicesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*devicesDataSource)(nil)
)

// NewDevicesDataSource returns a new devices data source, which
// lists device objects matching its filters.
func NewDevicesDataSource() datasource.DataSource {
	return &devicesDataSource{}
}

type devicesDataSource struct {
	client *netboxapi.Client
}

// devicesDataSourceModel is the data source model: the filters, the limit and the matching
// devices.
type devicesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"devices"`
}

// deviceResourceAttrTypes is the attribute type map of deviceResourceModel.
func deviceResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                        types.Int64Type,
		"name":                      types.StringType,
		"device_type_id":            types.Int64Type,
		"role_id":                   types.Int64Type,
		"site_id":                   types.Int64Type,
		"location_id":               types.Int64Type,
		"rack_id":                   types.Int64Type,
		"rack_position":             types.Float64Type,
		"rack_face":                 types.StringType,
		"latitude":                  types.Float64Type,
		"longitude":                 types.Float64Type,
		"status":                    types.StringType,
		"airflow":                   types.StringType,
		"tenant_id":                 types.Int64Type,
		"platform_id":               types.Int64Type,
		"cluster_id":                types.Int64Type,
		"virtual_chassis_id":        types.Int64Type,
		"virtual_chassis_position":  types.Int64Type,
		"virtual_chassis_priority":  types.Int64Type,
		"config_template_id":        types.Int64Type,
		"serial":                    types.StringType,
		"asset_tag":                 types.StringType,
		"primary_ip4_id":            types.Int64Type,
		"primary_ip6_id":            types.Int64Type,
		"oob_ip_id":                 types.Int64Type,
		"local_context_data":        jsontypes.NormalizedType{},
		"config_context":            jsontypes.NormalizedType{},
		"description":               types.StringType,
		"comments":                  types.StringType,
		"owner_id":                  types.Int64Type,
		"created":                   types.StringType,
		"last_updated":              types.StringType,
		"url":                       types.StringType,
		"console_port_count":        types.Int64Type,
		"console_server_port_count": types.Int64Type,
		"power_port_count":          types.Int64Type,
		"power_outlet_count":        types.Int64Type,
		"interface_count":           types.Int64Type,
		"front_port_count":          types.Int64Type,
		"rear_port_count":           types.Int64Type,
		"device_bay_count":          types.Int64Type,
		"module_bay_count":          types.Int64Type,
		"inventory_item_count":      types.Int64Type,
		"tags":                      types.SetType{ElemType: types.StringType},
		"tags_all":                  types.SetType{ElemType: types.StringType},
		"custom_fields":             types.MapType{ElemType: types.StringType},
	}
}

func (d *devicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_devices"
}

func (d *devicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists device objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: airflow, airflow__empty, airflow__ic, airflow__ie, airflow__iew, airflow__iregex, airflow__isw, airflow__n, airflow__nic, airflow__nie, airflow__niew, airflow__nisw, airflow__regex, asset_tag, asset_tag__empty, asset_tag__ic, asset_tag__ie, asset_tag__iew, asset_tag__iregex, asset_tag__isw, asset_tag__n, asset_tag__nic, asset_tag__nie, asset_tag__niew, asset_tag__nisw, asset_tag__regex, cluster_group, cluster_group__n, cluster_group_id, cluster_group_id__n, cluster_id, cluster_id__n, config_template_id, config_template_id__n, console_port_count, console_port_count__empty, console_port_count__gt, console_port_count__gte, console_port_count__lt, console_port_count__lte, console_port_count__n, console_ports, console_server_port_count, console_server_port_count__empty, console_server_port_count__gt, console_server_port_count__gte, console_server_port_count__lt, console_server_port_count__lte, console_server_port_count__n, console_server_ports, contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device_bay_count, device_bay_count__empty, device_bay_count__gt, device_bay_count__gte, device_bay_count__lt, device_bay_count__lte, device_bay_count__n, device_bays, device_type, device_type__n, device_type_id, device_type_id__n, face, face__empty, face__ic, face__ie, face__iew, face__iregex, face__isw, face__n, face__nic, face__nie, face__niew, face__nisw, face__regex, front_port_count, front_port_count__empty, front_port_count__gt, front_port_count__gte, front_port_count__lt, front_port_count__lte, front_port_count__n, has_oob_ip, has_primary_ip, has_virtual_device_context, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_count, interface_count__empty, interface_count__gt, interface_count__gte, interface_count__lt, interface_count__lte, interface_count__n, interfaces, inventory_item_count, inventory_item_count__empty, inventory_item_count__gt, inventory_item_count__gte, inventory_item_count__lt, inventory_item_count__lte, inventory_item_count__n, is_full_depth, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, latitude, latitude__empty, latitude__gt, latitude__gte, latitude__lt, latitude__lte, latitude__n, local_context_data, location, location__n, location_id, location_id__n, longitude, longitude__empty, longitude__gt, longitude__gte, longitude__lt, longitude__lte, longitude__n, mac_address, mac_address__ic, mac_address__ie, mac_address__iew, mac_address__iregex, mac_address__isw, mac_address__n, mac_address__nic, mac_address__nie, mac_address__niew, mac_address__nisw, mac_address__regex, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, model, model__n, module_bay_count, module_bay_count__empty, module_bay_count__gt, module_bay_count__gte, module_bay_count__lt, module_bay_count__lte, module_bay_count__n, module_bays, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, oob_ip_id, oob_ip_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, parent_bay_id, parent_bay_id__n, parent_device_id, parent_device_id__n, pass_through_ports, platform, platform__n, platform_id, platform_id__n, position, position__empty, position__gt, position__gte, position__lt, position__lte, position__n, power_outlet_count, power_outlet_count__empty, power_outlet_count__gt, power_outlet_count__gte, power_outlet_count__lt, power_outlet_count__lte, power_outlet_count__n, power_outlets, power_port_count, power_port_count__empty, power_port_count__gt, power_port_count__gte, power_port_count__lt, power_port_count__lte, power_port_count__n, power_ports, primary_ip4, primary_ip4__n, primary_ip4_id, primary_ip4_id__n, primary_ip6, primary_ip6__n, primary_ip6_id, primary_ip6_id__n, q, rack_id, rack_id__n, rear_port_count, rear_port_count__empty, rear_port_count__gt, rear_port_count__gte, rear_port_count__lt, rear_port_count__lte, rear_port_count__n, region, region__n, region_id, region_id__n, role, role__n, role_id, role_id__n, serial, serial__empty, serial__ic, serial__ie, serial__iew, serial__iregex, serial__isw, serial__n, serial__nic, serial__nie, serial__niew, serial__nisw, serial__regex, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, vc_position, vc_position__empty, vc_position__gt, vc_position__gte, vc_position__lt, vc_position__lte, vc_position__n, vc_priority, vc_priority__empty, vc_priority__gt, vc_priority__gte, vc_priority__lt, vc_priority__lte, vc_priority__n, virtual_chassis_id, virtual_chassis_id__n, virtual_chassis_member. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"devices": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching devices.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed:    true,
							Description: "Device name; NetBox allows unnamed devices.",
						},
						"device_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device type.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device role.",
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site.",
						},
						"location_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the location within the site.",
						},
						"rack_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the rack.",
						},
						"rack_position": schema.Float64Attribute{
							Computed:    true,
							Description: "Lowest rack unit occupied by the device; half units (e.g. 1.5) are allowed.",
						},
						"rack_face": schema.StringAttribute{
							Computed:    true,
							Description: "Rack face the device is mounted on. One of: front, rear.",
						},
						"latitude": schema.Float64Attribute{
							Computed:    true,
							Description: "GPS coordinate in decimal format (xx.yyyyyy).",
						},
						"longitude": schema.Float64Attribute{
							Computed:    true,
							Description: "GPS coordinate in decimal format (xx.yyyyyy).",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: offline, active, planned, staged, failed, inventory, decommissioning.",
						},
						"airflow": schema.StringAttribute{
							Computed:    true,
							Description: "One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"platform_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the platform.",
						},
						"cluster_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtualization cluster the device hosts.",
						},
						"virtual_chassis_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual chassis.",
						},
						"virtual_chassis_position": schema.Int64Attribute{
							Computed:    true,
							Description: "Position in the virtual chassis (0-255). Must be a whole number.",
						},
						"virtual_chassis_priority": schema.Int64Attribute{
							Computed:    true,
							Description: "Master election priority in the virtual chassis (0-255). Must be a whole number.",
						},
						"config_template_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the config template.",
						},
						"serial": schema.StringAttribute{
							Computed:    true,
							Description: "Chassis serial number.",
						},
						"asset_tag": schema.StringAttribute{
							Computed:    true,
							Description: "Unique asset tag.",
						},
						"primary_ip4_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the primary IPv4 address (managed by netbox_primary_ip).",
						},
						"primary_ip6_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the primary IPv6 address (managed by netbox_primary_ip).",
						},
						"oob_ip_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the out-of-band management IP address (managed by netbox_device_oob_ip).",
						},
						"local_context_data": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Local config context data as JSON text (use jsonencode()).",
						},
						"config_context": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Rendered config context (local_context_data merged with the context assigned through cluster, platform, role, site, tenant and tags) as JSON text, decode with jsondecode().",
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
						"console_port_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of console ports on the device.",
						},
						"console_server_port_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of console server ports on the device.",
						},
						"power_port_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of power ports on the device.",
						},
						"power_outlet_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of power outlets on the device.",
						},
						"interface_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of interfaces on the device.",
						},
						"front_port_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of front ports on the device.",
						},
						"rear_port_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of rear ports on the device.",
						},
						"device_bay_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of device bays on the device.",
						},
						"module_bay_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of module bays on the device.",
						},
						"inventory_item_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of inventory items on the device.",
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

func (d *devicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *devicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data devicesDataSourceModel

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
	var responseDTOs []*netboxapi.DeviceResponseDTO

	params := dcim.NewDcimDevicesListParams()
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
	var vSerial []string
	var vSerialn []string
	var vSerialIc []string
	var vSerialNic []string
	var vSerialIe []string
	var vSerialNie []string
	var vSerialIsw []string
	var vSerialNisw []string
	var vSerialIew []string
	var vSerialNiew []string
	var vSerialEmpty *bool
	var vSerialRegex []string
	var vSerialIregex []string
	var vAssetTag []string
	var vAssetTagn []string
	var vAssetTagIc []string
	var vAssetTagNic []string
	var vAssetTagIe []string
	var vAssetTagNie []string
	var vAssetTagIsw []string
	var vAssetTagNisw []string
	var vAssetTagIew []string
	var vAssetTagNiew []string
	var vAssetTagEmpty *bool
	var vAssetTagRegex []string
	var vAssetTagIregex []string
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
	var vRoleID []string
	var vRoleIDn []string
	var vDeviceTypeID []int64
	var vDeviceTypeIDn []int64
	var vTenantID []int64
	var vTenantIDn []int64
	var vRackID []int64
	var vRackIDn []int64
	var vLocationID []string
	var vLocationIDn []string
	var vClusterID []int64
	var vClusterIDn []int64
	var vPlatformID []string
	var vPlatformIDn []string
	var vRegion []string
	var vRegionn []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAirflow *string
	var vAirflown *string
	var vAirflowIc []string
	var vAirflowNic []string
	var vAirflowIe []string
	var vAirflowNie []string
	var vAirflowIsw []string
	var vAirflowNisw []string
	var vAirflowIew []string
	var vAirflowNiew []string
	var vAirflowEmpty *bool
	var vAirflowRegex []string
	var vAirflowIregex []string
	var vClusterGroup []string
	var vClusterGroupn []string
	var vClusterGroupID []int64
	var vClusterGroupIDn []int64
	var vConfigTemplateID []int64
	var vConfigTemplateIDn []int64
	var vConsolePortCount []int64
	var vConsolePortCountn []int64
	var vConsolePortCountLt []int64
	var vConsolePortCountLte []int64
	var vConsolePortCountGt []int64
	var vConsolePortCountGte []int64
	var vConsolePortCountEmpty *bool
	var vConsolePorts *bool
	var vConsoleServerPortCount []int64
	var vConsoleServerPortCountn []int64
	var vConsoleServerPortCountLt []int64
	var vConsoleServerPortCountLte []int64
	var vConsoleServerPortCountGt []int64
	var vConsoleServerPortCountGte []int64
	var vConsoleServerPortCountEmpty *bool
	var vConsoleServerPorts *bool
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
	var vDeviceBayCount []int64
	var vDeviceBayCountn []int64
	var vDeviceBayCountLt []int64
	var vDeviceBayCountLte []int64
	var vDeviceBayCountGt []int64
	var vDeviceBayCountGte []int64
	var vDeviceBayCountEmpty *bool
	var vDeviceBays *bool
	var vDeviceType []string
	var vDeviceTypen []string
	var vFace *string
	var vFacen *string
	var vFaceIc []string
	var vFaceNic []string
	var vFaceIe []string
	var vFaceNie []string
	var vFaceIsw []string
	var vFaceNisw []string
	var vFaceIew []string
	var vFaceNiew []string
	var vFaceEmpty *bool
	var vFaceRegex []string
	var vFaceIregex []string
	var vFrontPortCount []int64
	var vFrontPortCountn []int64
	var vFrontPortCountLt []int64
	var vFrontPortCountLte []int64
	var vFrontPortCountGt []int64
	var vFrontPortCountGte []int64
	var vFrontPortCountEmpty *bool
	var vHasOobIP *bool
	var vHasPrimaryIP *bool
	var vHasVirtualDeviceContext *bool
	var vInterfaceCount []int64
	var vInterfaceCountn []int64
	var vInterfaceCountLt []int64
	var vInterfaceCountLte []int64
	var vInterfaceCountGt []int64
	var vInterfaceCountGte []int64
	var vInterfaceCountEmpty *bool
	var vInterfaces *bool
	var vInventoryItemCount []int64
	var vInventoryItemCountn []int64
	var vInventoryItemCountLt []int64
	var vInventoryItemCountLte []int64
	var vInventoryItemCountGt []int64
	var vInventoryItemCountGte []int64
	var vInventoryItemCountEmpty *bool
	var vIsFullDepth *bool
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLatitude []float64
	var vLatituden []float64
	var vLatitudeLt []float64
	var vLatitudeLte []float64
	var vLatitudeGt []float64
	var vLatitudeGte []float64
	var vLatitudeEmpty *bool
	var vLocalContextData *bool
	var vLocation []string
	var vLocationn []string
	var vLongitude []float64
	var vLongituden []float64
	var vLongitudeLt []float64
	var vLongitudeLte []float64
	var vLongitudeGt []float64
	var vLongitudeGte []float64
	var vLongitudeEmpty *bool
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
	var vManufacturer []string
	var vManufacturern []string
	var vManufacturerID []int64
	var vManufacturerIDn []int64
	var vModel []string
	var vModeln []string
	var vModuleBayCount []int64
	var vModuleBayCountn []int64
	var vModuleBayCountLt []int64
	var vModuleBayCountLte []int64
	var vModuleBayCountGt []int64
	var vModuleBayCountGte []int64
	var vModuleBayCountEmpty *bool
	var vModuleBays *bool
	var vOobIPID []int64
	var vOobIPIDn []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vParentBayID []int64
	var vParentBayIDn []int64
	var vParentDeviceID []int64
	var vParentDeviceIDn []int64
	var vPassThroughPorts *bool
	var vPlatform []string
	var vPlatformn []string
	var vPosition []float64
	var vPositionn []float64
	var vPositionLt []float64
	var vPositionLte []float64
	var vPositionGt []float64
	var vPositionGte []float64
	var vPositionEmpty *bool
	var vPowerOutletCount []int64
	var vPowerOutletCountn []int64
	var vPowerOutletCountLt []int64
	var vPowerOutletCountLte []int64
	var vPowerOutletCountGt []int64
	var vPowerOutletCountGte []int64
	var vPowerOutletCountEmpty *bool
	var vPowerOutlets *bool
	var vPowerPortCount []int64
	var vPowerPortCountn []int64
	var vPowerPortCountLt []int64
	var vPowerPortCountLte []int64
	var vPowerPortCountGt []int64
	var vPowerPortCountGte []int64
	var vPowerPortCountEmpty *bool
	var vPowerPorts *bool
	var vPrimaryIp4 []string
	var vPrimaryIp4n []string
	var vPrimaryIp4ID []int64
	var vPrimaryIp4IDn []int64
	var vPrimaryIp6 []string
	var vPrimaryIp6n []string
	var vPrimaryIp6ID []int64
	var vPrimaryIp6IDn []int64
	var vQ *string
	var vRearPortCount []int64
	var vRearPortCountn []int64
	var vRearPortCountLt []int64
	var vRearPortCountLte []int64
	var vRearPortCountGt []int64
	var vRearPortCountGte []int64
	var vRearPortCountEmpty *bool
	var vRegionID []string
	var vRegionIDn []string
	var vRole []string
	var vRolen []string
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
	var vVcPosition []int64
	var vVcPositionn []int64
	var vVcPositionLt []int64
	var vVcPositionLte []int64
	var vVcPositionGt []int64
	var vVcPositionGte []int64
	var vVcPositionEmpty *bool
	var vVcPriority []int64
	var vVcPriorityn []int64
	var vVcPriorityLt []int64
	var vVcPriorityLte []int64
	var vVcPriorityGt []int64
	var vVcPriorityGte []int64
	var vVcPriorityEmpty *bool
	var vVirtualChassisID []int64
	var vVirtualChassisIDn []int64
	var vVirtualChassisMember *bool
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
		case "serial":
			v := value
			vSerial = append(vSerial, v)
		case "serial__n":
			v := value
			vSerialn = append(vSerialn, v)
		case "serial__ic":
			v := value
			vSerialIc = append(vSerialIc, v)
		case "serial__nic":
			v := value
			vSerialNic = append(vSerialNic, v)
		case "serial__ie":
			v := value
			vSerialIe = append(vSerialIe, v)
		case "serial__nie":
			v := value
			vSerialNie = append(vSerialNie, v)
		case "serial__isw":
			v := value
			vSerialIsw = append(vSerialIsw, v)
		case "serial__nisw":
			v := value
			vSerialNisw = append(vSerialNisw, v)
		case "serial__iew":
			v := value
			vSerialIew = append(vSerialIew, v)
		case "serial__niew":
			v := value
			vSerialNiew = append(vSerialNiew, v)
		case "serial__empty":
			if vSerialEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'serial__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'serial__empty' takes a boolean, got %q.", value))
				return
			}
			vSerialEmpty = &v
		case "serial__regex":
			v := value
			vSerialRegex = append(vSerialRegex, v)
		case "serial__iregex":
			v := value
			vSerialIregex = append(vSerialIregex, v)
		case "asset_tag":
			v := value
			vAssetTag = append(vAssetTag, v)
		case "asset_tag__n":
			v := value
			vAssetTagn = append(vAssetTagn, v)
		case "asset_tag__ic":
			v := value
			vAssetTagIc = append(vAssetTagIc, v)
		case "asset_tag__nic":
			v := value
			vAssetTagNic = append(vAssetTagNic, v)
		case "asset_tag__ie":
			v := value
			vAssetTagIe = append(vAssetTagIe, v)
		case "asset_tag__nie":
			v := value
			vAssetTagNie = append(vAssetTagNie, v)
		case "asset_tag__isw":
			v := value
			vAssetTagIsw = append(vAssetTagIsw, v)
		case "asset_tag__nisw":
			v := value
			vAssetTagNisw = append(vAssetTagNisw, v)
		case "asset_tag__iew":
			v := value
			vAssetTagIew = append(vAssetTagIew, v)
		case "asset_tag__niew":
			v := value
			vAssetTagNiew = append(vAssetTagNiew, v)
		case "asset_tag__empty":
			if vAssetTagEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'asset_tag__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asset_tag__empty' takes a boolean, got %q.", value))
				return
			}
			vAssetTagEmpty = &v
		case "asset_tag__regex":
			v := value
			vAssetTagRegex = append(vAssetTagRegex, v)
		case "asset_tag__iregex":
			v := value
			vAssetTagIregex = append(vAssetTagIregex, v)
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
		case "role_id":
			v := value
			vRoleID = append(vRoleID, v)
		case "role_id__n":
			v := value
			vRoleIDn = append(vRoleIDn, v)
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
		case "location_id":
			v := value
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v := value
			vLocationIDn = append(vLocationIDn, v)
		case "cluster_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_id' takes an integer, got %q.", value))
				return
			}
			vClusterID = append(vClusterID, v)
		case "cluster_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_id__n' takes an integer, got %q.", value))
				return
			}
			vClusterIDn = append(vClusterIDn, v)
		case "platform_id":
			v := value
			vPlatformID = append(vPlatformID, v)
		case "platform_id__n":
			v := value
			vPlatformIDn = append(vPlatformIDn, v)
		case "region":
			v := value
			vRegion = append(vRegion, v)
		case "region__n":
			v := value
			vRegionn = append(vRegionn, v)
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
		case "airflow":
			if vAirflow != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'airflow' takes a single value.")
				return
			}
			v := value
			vAirflow = &v
		case "airflow__n":
			if vAirflown != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'airflow__n' takes a single value.")
				return
			}
			v := value
			vAirflown = &v
		case "airflow__ic":
			v := value
			vAirflowIc = append(vAirflowIc, v)
		case "airflow__nic":
			v := value
			vAirflowNic = append(vAirflowNic, v)
		case "airflow__ie":
			v := value
			vAirflowIe = append(vAirflowIe, v)
		case "airflow__nie":
			v := value
			vAirflowNie = append(vAirflowNie, v)
		case "airflow__isw":
			v := value
			vAirflowIsw = append(vAirflowIsw, v)
		case "airflow__nisw":
			v := value
			vAirflowNisw = append(vAirflowNisw, v)
		case "airflow__iew":
			v := value
			vAirflowIew = append(vAirflowIew, v)
		case "airflow__niew":
			v := value
			vAirflowNiew = append(vAirflowNiew, v)
		case "airflow__empty":
			if vAirflowEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'airflow__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'airflow__empty' takes a boolean, got %q.", value))
				return
			}
			vAirflowEmpty = &v
		case "airflow__regex":
			v := value
			vAirflowRegex = append(vAirflowRegex, v)
		case "airflow__iregex":
			v := value
			vAirflowIregex = append(vAirflowIregex, v)
		case "cluster_group":
			v := value
			vClusterGroup = append(vClusterGroup, v)
		case "cluster_group__n":
			v := value
			vClusterGroupn = append(vClusterGroupn, v)
		case "cluster_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_group_id' takes an integer, got %q.", value))
				return
			}
			vClusterGroupID = append(vClusterGroupID, v)
		case "cluster_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_group_id__n' takes an integer, got %q.", value))
				return
			}
			vClusterGroupIDn = append(vClusterGroupIDn, v)
		case "config_template_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'config_template_id' takes an integer, got %q.", value))
				return
			}
			vConfigTemplateID = append(vConfigTemplateID, v)
		case "config_template_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'config_template_id__n' takes an integer, got %q.", value))
				return
			}
			vConfigTemplateIDn = append(vConfigTemplateIDn, v)
		case "console_port_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_count' takes an integer, got %q.", value))
				return
			}
			vConsolePortCount = append(vConsolePortCount, v)
		case "console_port_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_count__n' takes an integer, got %q.", value))
				return
			}
			vConsolePortCountn = append(vConsolePortCountn, v)
		case "console_port_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_count__lt' takes an integer, got %q.", value))
				return
			}
			vConsolePortCountLt = append(vConsolePortCountLt, v)
		case "console_port_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_count__lte' takes an integer, got %q.", value))
				return
			}
			vConsolePortCountLte = append(vConsolePortCountLte, v)
		case "console_port_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_count__gt' takes an integer, got %q.", value))
				return
			}
			vConsolePortCountGt = append(vConsolePortCountGt, v)
		case "console_port_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_count__gte' takes an integer, got %q.", value))
				return
			}
			vConsolePortCountGte = append(vConsolePortCountGte, v)
		case "console_port_count__empty":
			if vConsolePortCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_port_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_count__empty' takes a boolean, got %q.", value))
				return
			}
			vConsolePortCountEmpty = &v
		case "console_ports":
			if vConsolePorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_ports' takes a boolean, got %q.", value))
				return
			}
			vConsolePorts = &v
		case "console_server_port_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_count' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortCount = append(vConsoleServerPortCount, v)
		case "console_server_port_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_count__n' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortCountn = append(vConsoleServerPortCountn, v)
		case "console_server_port_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_count__lt' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortCountLt = append(vConsoleServerPortCountLt, v)
		case "console_server_port_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_count__lte' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortCountLte = append(vConsoleServerPortCountLte, v)
		case "console_server_port_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_count__gt' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortCountGt = append(vConsoleServerPortCountGt, v)
		case "console_server_port_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_count__gte' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortCountGte = append(vConsoleServerPortCountGte, v)
		case "console_server_port_count__empty":
			if vConsoleServerPortCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_server_port_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_count__empty' takes a boolean, got %q.", value))
				return
			}
			vConsoleServerPortCountEmpty = &v
		case "console_server_ports":
			if vConsoleServerPorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_server_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_ports' takes a boolean, got %q.", value))
				return
			}
			vConsoleServerPorts = &v
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
		case "device_bay_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bay_count' takes an integer, got %q.", value))
				return
			}
			vDeviceBayCount = append(vDeviceBayCount, v)
		case "device_bay_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bay_count__n' takes an integer, got %q.", value))
				return
			}
			vDeviceBayCountn = append(vDeviceBayCountn, v)
		case "device_bay_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bay_count__lt' takes an integer, got %q.", value))
				return
			}
			vDeviceBayCountLt = append(vDeviceBayCountLt, v)
		case "device_bay_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bay_count__lte' takes an integer, got %q.", value))
				return
			}
			vDeviceBayCountLte = append(vDeviceBayCountLte, v)
		case "device_bay_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bay_count__gt' takes an integer, got %q.", value))
				return
			}
			vDeviceBayCountGt = append(vDeviceBayCountGt, v)
		case "device_bay_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bay_count__gte' takes an integer, got %q.", value))
				return
			}
			vDeviceBayCountGte = append(vDeviceBayCountGte, v)
		case "device_bay_count__empty":
			if vDeviceBayCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'device_bay_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bay_count__empty' takes a boolean, got %q.", value))
				return
			}
			vDeviceBayCountEmpty = &v
		case "device_bays":
			if vDeviceBays != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'device_bays' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_bays' takes a boolean, got %q.", value))
				return
			}
			vDeviceBays = &v
		case "device_type":
			v := value
			vDeviceType = append(vDeviceType, v)
		case "device_type__n":
			v := value
			vDeviceTypen = append(vDeviceTypen, v)
		case "face":
			if vFace != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'face' takes a single value.")
				return
			}
			v := value
			vFace = &v
		case "face__n":
			if vFacen != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'face__n' takes a single value.")
				return
			}
			v := value
			vFacen = &v
		case "face__ic":
			v := value
			vFaceIc = append(vFaceIc, v)
		case "face__nic":
			v := value
			vFaceNic = append(vFaceNic, v)
		case "face__ie":
			v := value
			vFaceIe = append(vFaceIe, v)
		case "face__nie":
			v := value
			vFaceNie = append(vFaceNie, v)
		case "face__isw":
			v := value
			vFaceIsw = append(vFaceIsw, v)
		case "face__nisw":
			v := value
			vFaceNisw = append(vFaceNisw, v)
		case "face__iew":
			v := value
			vFaceIew = append(vFaceIew, v)
		case "face__niew":
			v := value
			vFaceNiew = append(vFaceNiew, v)
		case "face__empty":
			if vFaceEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'face__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'face__empty' takes a boolean, got %q.", value))
				return
			}
			vFaceEmpty = &v
		case "face__regex":
			v := value
			vFaceRegex = append(vFaceRegex, v)
		case "face__iregex":
			v := value
			vFaceIregex = append(vFaceIregex, v)
		case "front_port_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_count' takes an integer, got %q.", value))
				return
			}
			vFrontPortCount = append(vFrontPortCount, v)
		case "front_port_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_count__n' takes an integer, got %q.", value))
				return
			}
			vFrontPortCountn = append(vFrontPortCountn, v)
		case "front_port_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_count__lt' takes an integer, got %q.", value))
				return
			}
			vFrontPortCountLt = append(vFrontPortCountLt, v)
		case "front_port_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_count__lte' takes an integer, got %q.", value))
				return
			}
			vFrontPortCountLte = append(vFrontPortCountLte, v)
		case "front_port_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_count__gt' takes an integer, got %q.", value))
				return
			}
			vFrontPortCountGt = append(vFrontPortCountGt, v)
		case "front_port_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_count__gte' takes an integer, got %q.", value))
				return
			}
			vFrontPortCountGte = append(vFrontPortCountGte, v)
		case "front_port_count__empty":
			if vFrontPortCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'front_port_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_count__empty' takes a boolean, got %q.", value))
				return
			}
			vFrontPortCountEmpty = &v
		case "has_oob_ip":
			if vHasOobIP != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'has_oob_ip' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'has_oob_ip' takes a boolean, got %q.", value))
				return
			}
			vHasOobIP = &v
		case "has_primary_ip":
			if vHasPrimaryIP != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'has_primary_ip' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'has_primary_ip' takes a boolean, got %q.", value))
				return
			}
			vHasPrimaryIP = &v
		case "has_virtual_device_context":
			if vHasVirtualDeviceContext != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'has_virtual_device_context' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'has_virtual_device_context' takes a boolean, got %q.", value))
				return
			}
			vHasVirtualDeviceContext = &v
		case "interface_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_count' takes an integer, got %q.", value))
				return
			}
			vInterfaceCount = append(vInterfaceCount, v)
		case "interface_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_count__n' takes an integer, got %q.", value))
				return
			}
			vInterfaceCountn = append(vInterfaceCountn, v)
		case "interface_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_count__lt' takes an integer, got %q.", value))
				return
			}
			vInterfaceCountLt = append(vInterfaceCountLt, v)
		case "interface_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_count__lte' takes an integer, got %q.", value))
				return
			}
			vInterfaceCountLte = append(vInterfaceCountLte, v)
		case "interface_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_count__gt' takes an integer, got %q.", value))
				return
			}
			vInterfaceCountGt = append(vInterfaceCountGt, v)
		case "interface_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_count__gte' takes an integer, got %q.", value))
				return
			}
			vInterfaceCountGte = append(vInterfaceCountGte, v)
		case "interface_count__empty":
			if vInterfaceCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'interface_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_count__empty' takes a boolean, got %q.", value))
				return
			}
			vInterfaceCountEmpty = &v
		case "interfaces":
			if vInterfaces != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'interfaces' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interfaces' takes a boolean, got %q.", value))
				return
			}
			vInterfaces = &v
		case "inventory_item_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'inventory_item_count' takes an integer, got %q.", value))
				return
			}
			vInventoryItemCount = append(vInventoryItemCount, v)
		case "inventory_item_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'inventory_item_count__n' takes an integer, got %q.", value))
				return
			}
			vInventoryItemCountn = append(vInventoryItemCountn, v)
		case "inventory_item_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'inventory_item_count__lt' takes an integer, got %q.", value))
				return
			}
			vInventoryItemCountLt = append(vInventoryItemCountLt, v)
		case "inventory_item_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'inventory_item_count__lte' takes an integer, got %q.", value))
				return
			}
			vInventoryItemCountLte = append(vInventoryItemCountLte, v)
		case "inventory_item_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'inventory_item_count__gt' takes an integer, got %q.", value))
				return
			}
			vInventoryItemCountGt = append(vInventoryItemCountGt, v)
		case "inventory_item_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'inventory_item_count__gte' takes an integer, got %q.", value))
				return
			}
			vInventoryItemCountGte = append(vInventoryItemCountGte, v)
		case "inventory_item_count__empty":
			if vInventoryItemCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'inventory_item_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'inventory_item_count__empty' takes a boolean, got %q.", value))
				return
			}
			vInventoryItemCountEmpty = &v
		case "is_full_depth":
			if vIsFullDepth != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'is_full_depth' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'is_full_depth' takes a boolean, got %q.", value))
				return
			}
			vIsFullDepth = &v
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
		case "latitude":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'latitude' takes a number, got %q.", value))
				return
			}
			vLatitude = append(vLatitude, v)
		case "latitude__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'latitude__n' takes a number, got %q.", value))
				return
			}
			vLatituden = append(vLatituden, v)
		case "latitude__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'latitude__lt' takes a number, got %q.", value))
				return
			}
			vLatitudeLt = append(vLatitudeLt, v)
		case "latitude__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'latitude__lte' takes a number, got %q.", value))
				return
			}
			vLatitudeLte = append(vLatitudeLte, v)
		case "latitude__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'latitude__gt' takes a number, got %q.", value))
				return
			}
			vLatitudeGt = append(vLatitudeGt, v)
		case "latitude__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'latitude__gte' takes a number, got %q.", value))
				return
			}
			vLatitudeGte = append(vLatitudeGte, v)
		case "latitude__empty":
			if vLatitudeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'latitude__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'latitude__empty' takes a boolean, got %q.", value))
				return
			}
			vLatitudeEmpty = &v
		case "local_context_data":
			if vLocalContextData != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'local_context_data' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_context_data' takes a boolean, got %q.", value))
				return
			}
			vLocalContextData = &v
		case "location":
			v := value
			vLocation = append(vLocation, v)
		case "location__n":
			v := value
			vLocationn = append(vLocationn, v)
		case "longitude":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'longitude' takes a number, got %q.", value))
				return
			}
			vLongitude = append(vLongitude, v)
		case "longitude__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'longitude__n' takes a number, got %q.", value))
				return
			}
			vLongituden = append(vLongituden, v)
		case "longitude__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'longitude__lt' takes a number, got %q.", value))
				return
			}
			vLongitudeLt = append(vLongitudeLt, v)
		case "longitude__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'longitude__lte' takes a number, got %q.", value))
				return
			}
			vLongitudeLte = append(vLongitudeLte, v)
		case "longitude__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'longitude__gt' takes a number, got %q.", value))
				return
			}
			vLongitudeGt = append(vLongitudeGt, v)
		case "longitude__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'longitude__gte' takes a number, got %q.", value))
				return
			}
			vLongitudeGte = append(vLongitudeGte, v)
		case "longitude__empty":
			if vLongitudeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'longitude__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'longitude__empty' takes a boolean, got %q.", value))
				return
			}
			vLongitudeEmpty = &v
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
		case "manufacturer":
			v := value
			vManufacturer = append(vManufacturer, v)
		case "manufacturer__n":
			v := value
			vManufacturern = append(vManufacturern, v)
		case "manufacturer_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'manufacturer_id' takes an integer, got %q.", value))
				return
			}
			vManufacturerID = append(vManufacturerID, v)
		case "manufacturer_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'manufacturer_id__n' takes an integer, got %q.", value))
				return
			}
			vManufacturerIDn = append(vManufacturerIDn, v)
		case "model":
			v := value
			vModel = append(vModel, v)
		case "model__n":
			v := value
			vModeln = append(vModeln, v)
		case "module_bay_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_count' takes an integer, got %q.", value))
				return
			}
			vModuleBayCount = append(vModuleBayCount, v)
		case "module_bay_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_count__n' takes an integer, got %q.", value))
				return
			}
			vModuleBayCountn = append(vModuleBayCountn, v)
		case "module_bay_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_count__lt' takes an integer, got %q.", value))
				return
			}
			vModuleBayCountLt = append(vModuleBayCountLt, v)
		case "module_bay_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_count__lte' takes an integer, got %q.", value))
				return
			}
			vModuleBayCountLte = append(vModuleBayCountLte, v)
		case "module_bay_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_count__gt' takes an integer, got %q.", value))
				return
			}
			vModuleBayCountGt = append(vModuleBayCountGt, v)
		case "module_bay_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_count__gte' takes an integer, got %q.", value))
				return
			}
			vModuleBayCountGte = append(vModuleBayCountGte, v)
		case "module_bay_count__empty":
			if vModuleBayCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'module_bay_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_count__empty' takes a boolean, got %q.", value))
				return
			}
			vModuleBayCountEmpty = &v
		case "module_bays":
			if vModuleBays != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'module_bays' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bays' takes a boolean, got %q.", value))
				return
			}
			vModuleBays = &v
		case "oob_ip_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'oob_ip_id' takes an integer, got %q.", value))
				return
			}
			vOobIPID = append(vOobIPID, v)
		case "oob_ip_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'oob_ip_id__n' takes an integer, got %q.", value))
				return
			}
			vOobIPIDn = append(vOobIPIDn, v)
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
		case "parent_bay_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_bay_id' takes an integer, got %q.", value))
				return
			}
			vParentBayID = append(vParentBayID, v)
		case "parent_bay_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_bay_id__n' takes an integer, got %q.", value))
				return
			}
			vParentBayIDn = append(vParentBayIDn, v)
		case "parent_device_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_device_id' takes an integer, got %q.", value))
				return
			}
			vParentDeviceID = append(vParentDeviceID, v)
		case "parent_device_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_device_id__n' takes an integer, got %q.", value))
				return
			}
			vParentDeviceIDn = append(vParentDeviceIDn, v)
		case "pass_through_ports":
			if vPassThroughPorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'pass_through_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pass_through_ports' takes a boolean, got %q.", value))
				return
			}
			vPassThroughPorts = &v
		case "platform":
			v := value
			vPlatform = append(vPlatform, v)
		case "platform__n":
			v := value
			vPlatformn = append(vPlatformn, v)
		case "position":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'position' takes a number, got %q.", value))
				return
			}
			vPosition = append(vPosition, v)
		case "position__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'position__n' takes a number, got %q.", value))
				return
			}
			vPositionn = append(vPositionn, v)
		case "position__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'position__lt' takes a number, got %q.", value))
				return
			}
			vPositionLt = append(vPositionLt, v)
		case "position__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'position__lte' takes a number, got %q.", value))
				return
			}
			vPositionLte = append(vPositionLte, v)
		case "position__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'position__gt' takes a number, got %q.", value))
				return
			}
			vPositionGt = append(vPositionGt, v)
		case "position__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'position__gte' takes a number, got %q.", value))
				return
			}
			vPositionGte = append(vPositionGte, v)
		case "position__empty":
			if vPositionEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'position__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'position__empty' takes a boolean, got %q.", value))
				return
			}
			vPositionEmpty = &v
		case "power_outlet_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_count' takes an integer, got %q.", value))
				return
			}
			vPowerOutletCount = append(vPowerOutletCount, v)
		case "power_outlet_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_count__n' takes an integer, got %q.", value))
				return
			}
			vPowerOutletCountn = append(vPowerOutletCountn, v)
		case "power_outlet_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_count__lt' takes an integer, got %q.", value))
				return
			}
			vPowerOutletCountLt = append(vPowerOutletCountLt, v)
		case "power_outlet_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_count__lte' takes an integer, got %q.", value))
				return
			}
			vPowerOutletCountLte = append(vPowerOutletCountLte, v)
		case "power_outlet_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_count__gt' takes an integer, got %q.", value))
				return
			}
			vPowerOutletCountGt = append(vPowerOutletCountGt, v)
		case "power_outlet_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_count__gte' takes an integer, got %q.", value))
				return
			}
			vPowerOutletCountGte = append(vPowerOutletCountGte, v)
		case "power_outlet_count__empty":
			if vPowerOutletCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_outlet_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_count__empty' takes a boolean, got %q.", value))
				return
			}
			vPowerOutletCountEmpty = &v
		case "power_outlets":
			if vPowerOutlets != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_outlets' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlets' takes a boolean, got %q.", value))
				return
			}
			vPowerOutlets = &v
		case "power_port_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_count' takes an integer, got %q.", value))
				return
			}
			vPowerPortCount = append(vPowerPortCount, v)
		case "power_port_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_count__n' takes an integer, got %q.", value))
				return
			}
			vPowerPortCountn = append(vPowerPortCountn, v)
		case "power_port_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_count__lt' takes an integer, got %q.", value))
				return
			}
			vPowerPortCountLt = append(vPowerPortCountLt, v)
		case "power_port_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_count__lte' takes an integer, got %q.", value))
				return
			}
			vPowerPortCountLte = append(vPowerPortCountLte, v)
		case "power_port_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_count__gt' takes an integer, got %q.", value))
				return
			}
			vPowerPortCountGt = append(vPowerPortCountGt, v)
		case "power_port_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_count__gte' takes an integer, got %q.", value))
				return
			}
			vPowerPortCountGte = append(vPowerPortCountGte, v)
		case "power_port_count__empty":
			if vPowerPortCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_port_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_count__empty' takes a boolean, got %q.", value))
				return
			}
			vPowerPortCountEmpty = &v
		case "power_ports":
			if vPowerPorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_ports' takes a boolean, got %q.", value))
				return
			}
			vPowerPorts = &v
		case "primary_ip4":
			v := value
			vPrimaryIp4 = append(vPrimaryIp4, v)
		case "primary_ip4__n":
			v := value
			vPrimaryIp4n = append(vPrimaryIp4n, v)
		case "primary_ip4_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_ip4_id' takes an integer, got %q.", value))
				return
			}
			vPrimaryIp4ID = append(vPrimaryIp4ID, v)
		case "primary_ip4_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_ip4_id__n' takes an integer, got %q.", value))
				return
			}
			vPrimaryIp4IDn = append(vPrimaryIp4IDn, v)
		case "primary_ip6":
			v := value
			vPrimaryIp6 = append(vPrimaryIp6, v)
		case "primary_ip6__n":
			v := value
			vPrimaryIp6n = append(vPrimaryIp6n, v)
		case "primary_ip6_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_ip6_id' takes an integer, got %q.", value))
				return
			}
			vPrimaryIp6ID = append(vPrimaryIp6ID, v)
		case "primary_ip6_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_ip6_id__n' takes an integer, got %q.", value))
				return
			}
			vPrimaryIp6IDn = append(vPrimaryIp6IDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "rear_port_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_count' takes an integer, got %q.", value))
				return
			}
			vRearPortCount = append(vRearPortCount, v)
		case "rear_port_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_count__n' takes an integer, got %q.", value))
				return
			}
			vRearPortCountn = append(vRearPortCountn, v)
		case "rear_port_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_count__lt' takes an integer, got %q.", value))
				return
			}
			vRearPortCountLt = append(vRearPortCountLt, v)
		case "rear_port_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_count__lte' takes an integer, got %q.", value))
				return
			}
			vRearPortCountLte = append(vRearPortCountLte, v)
		case "rear_port_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_count__gt' takes an integer, got %q.", value))
				return
			}
			vRearPortCountGt = append(vRearPortCountGt, v)
		case "rear_port_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_count__gte' takes an integer, got %q.", value))
				return
			}
			vRearPortCountGte = append(vRearPortCountGte, v)
		case "rear_port_count__empty":
			if vRearPortCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'rear_port_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_count__empty' takes a boolean, got %q.", value))
				return
			}
			vRearPortCountEmpty = &v
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
		case "vc_position":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_position' takes an integer, got %q.", value))
				return
			}
			vVcPosition = append(vVcPosition, v)
		case "vc_position__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_position__n' takes an integer, got %q.", value))
				return
			}
			vVcPositionn = append(vVcPositionn, v)
		case "vc_position__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_position__lt' takes an integer, got %q.", value))
				return
			}
			vVcPositionLt = append(vVcPositionLt, v)
		case "vc_position__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_position__lte' takes an integer, got %q.", value))
				return
			}
			vVcPositionLte = append(vVcPositionLte, v)
		case "vc_position__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_position__gt' takes an integer, got %q.", value))
				return
			}
			vVcPositionGt = append(vVcPositionGt, v)
		case "vc_position__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_position__gte' takes an integer, got %q.", value))
				return
			}
			vVcPositionGte = append(vVcPositionGte, v)
		case "vc_position__empty":
			if vVcPositionEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vc_position__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_position__empty' takes a boolean, got %q.", value))
				return
			}
			vVcPositionEmpty = &v
		case "vc_priority":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_priority' takes an integer, got %q.", value))
				return
			}
			vVcPriority = append(vVcPriority, v)
		case "vc_priority__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_priority__n' takes an integer, got %q.", value))
				return
			}
			vVcPriorityn = append(vVcPriorityn, v)
		case "vc_priority__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_priority__lt' takes an integer, got %q.", value))
				return
			}
			vVcPriorityLt = append(vVcPriorityLt, v)
		case "vc_priority__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_priority__lte' takes an integer, got %q.", value))
				return
			}
			vVcPriorityLte = append(vVcPriorityLte, v)
		case "vc_priority__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_priority__gt' takes an integer, got %q.", value))
				return
			}
			vVcPriorityGt = append(vVcPriorityGt, v)
		case "vc_priority__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_priority__gte' takes an integer, got %q.", value))
				return
			}
			vVcPriorityGte = append(vVcPriorityGte, v)
		case "vc_priority__empty":
			if vVcPriorityEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vc_priority__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vc_priority__empty' takes a boolean, got %q.", value))
				return
			}
			vVcPriorityEmpty = &v
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
			if vVirtualChassisMember != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'virtual_chassis_member' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_chassis_member' takes a boolean, got %q.", value))
				return
			}
			vVirtualChassisMember = &v
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
	if len(vSerial) > 0 {
		params.SetSerial(vSerial)
	}
	if len(vSerialn) > 0 {
		params.SetSerialn(vSerialn)
	}
	if len(vSerialIc) > 0 {
		params.SetSerialIc(vSerialIc)
	}
	if len(vSerialNic) > 0 {
		params.SetSerialNic(vSerialNic)
	}
	if len(vSerialIe) > 0 {
		params.SetSerialIe(vSerialIe)
	}
	if len(vSerialNie) > 0 {
		params.SetSerialNie(vSerialNie)
	}
	if len(vSerialIsw) > 0 {
		params.SetSerialIsw(vSerialIsw)
	}
	if len(vSerialNisw) > 0 {
		params.SetSerialNisw(vSerialNisw)
	}
	if len(vSerialIew) > 0 {
		params.SetSerialIew(vSerialIew)
	}
	if len(vSerialNiew) > 0 {
		params.SetSerialNiew(vSerialNiew)
	}
	if vSerialEmpty != nil {
		params.SetSerialEmpty(vSerialEmpty)
	}
	if len(vSerialRegex) > 0 {
		params.SetSerialRegex(vSerialRegex)
	}
	if len(vSerialIregex) > 0 {
		params.SetSerialIregex(vSerialIregex)
	}
	if len(vAssetTag) > 0 {
		params.SetAssetTag(vAssetTag)
	}
	if len(vAssetTagn) > 0 {
		params.SetAssetTagn(vAssetTagn)
	}
	if len(vAssetTagIc) > 0 {
		params.SetAssetTagIc(vAssetTagIc)
	}
	if len(vAssetTagNic) > 0 {
		params.SetAssetTagNic(vAssetTagNic)
	}
	if len(vAssetTagIe) > 0 {
		params.SetAssetTagIe(vAssetTagIe)
	}
	if len(vAssetTagNie) > 0 {
		params.SetAssetTagNie(vAssetTagNie)
	}
	if len(vAssetTagIsw) > 0 {
		params.SetAssetTagIsw(vAssetTagIsw)
	}
	if len(vAssetTagNisw) > 0 {
		params.SetAssetTagNisw(vAssetTagNisw)
	}
	if len(vAssetTagIew) > 0 {
		params.SetAssetTagIew(vAssetTagIew)
	}
	if len(vAssetTagNiew) > 0 {
		params.SetAssetTagNiew(vAssetTagNiew)
	}
	if vAssetTagEmpty != nil {
		params.SetAssetTagEmpty(vAssetTagEmpty)
	}
	if len(vAssetTagRegex) > 0 {
		params.SetAssetTagRegex(vAssetTagRegex)
	}
	if len(vAssetTagIregex) > 0 {
		params.SetAssetTagIregex(vAssetTagIregex)
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
	if len(vRoleID) > 0 {
		params.SetRoleID(vRoleID)
	}
	if len(vRoleIDn) > 0 {
		params.SetRoleIDn(vRoleIDn)
	}
	if len(vDeviceTypeID) > 0 {
		params.SetDeviceTypeID(vDeviceTypeID)
	}
	if len(vDeviceTypeIDn) > 0 {
		params.SetDeviceTypeIDn(vDeviceTypeIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vRackID) > 0 {
		params.SetRackID(vRackID)
	}
	if len(vRackIDn) > 0 {
		params.SetRackIDn(vRackIDn)
	}
	if len(vLocationID) > 0 {
		params.SetLocationID(vLocationID)
	}
	if len(vLocationIDn) > 0 {
		params.SetLocationIDn(vLocationIDn)
	}
	if len(vClusterID) > 0 {
		params.SetClusterID(vClusterID)
	}
	if len(vClusterIDn) > 0 {
		params.SetClusterIDn(vClusterIDn)
	}
	if len(vPlatformID) > 0 {
		params.SetPlatformID(vPlatformID)
	}
	if len(vPlatformIDn) > 0 {
		params.SetPlatformIDn(vPlatformIDn)
	}
	if len(vRegion) > 0 {
		params.SetRegion(vRegion)
	}
	if len(vRegionn) > 0 {
		params.SetRegionn(vRegionn)
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
	if vAirflow != nil {
		params.SetAirflow(vAirflow)
	}
	if vAirflown != nil {
		params.SetAirflown(vAirflown)
	}
	if len(vAirflowIc) > 0 {
		params.SetAirflowIc(vAirflowIc)
	}
	if len(vAirflowNic) > 0 {
		params.SetAirflowNic(vAirflowNic)
	}
	if len(vAirflowIe) > 0 {
		params.SetAirflowIe(vAirflowIe)
	}
	if len(vAirflowNie) > 0 {
		params.SetAirflowNie(vAirflowNie)
	}
	if len(vAirflowIsw) > 0 {
		params.SetAirflowIsw(vAirflowIsw)
	}
	if len(vAirflowNisw) > 0 {
		params.SetAirflowNisw(vAirflowNisw)
	}
	if len(vAirflowIew) > 0 {
		params.SetAirflowIew(vAirflowIew)
	}
	if len(vAirflowNiew) > 0 {
		params.SetAirflowNiew(vAirflowNiew)
	}
	if vAirflowEmpty != nil {
		params.SetAirflowEmpty(vAirflowEmpty)
	}
	if len(vAirflowRegex) > 0 {
		params.SetAirflowRegex(vAirflowRegex)
	}
	if len(vAirflowIregex) > 0 {
		params.SetAirflowIregex(vAirflowIregex)
	}
	if len(vClusterGroup) > 0 {
		params.SetClusterGroup(vClusterGroup)
	}
	if len(vClusterGroupn) > 0 {
		params.SetClusterGroupn(vClusterGroupn)
	}
	if len(vClusterGroupID) > 0 {
		params.SetClusterGroupID(vClusterGroupID)
	}
	if len(vClusterGroupIDn) > 0 {
		params.SetClusterGroupIDn(vClusterGroupIDn)
	}
	if len(vConfigTemplateID) > 0 {
		params.SetConfigTemplateID(vConfigTemplateID)
	}
	if len(vConfigTemplateIDn) > 0 {
		params.SetConfigTemplateIDn(vConfigTemplateIDn)
	}
	if len(vConsolePortCount) > 0 {
		params.SetConsolePortCount(vConsolePortCount)
	}
	if len(vConsolePortCountn) > 0 {
		params.SetConsolePortCountn(vConsolePortCountn)
	}
	if len(vConsolePortCountLt) > 0 {
		params.SetConsolePortCountLt(vConsolePortCountLt)
	}
	if len(vConsolePortCountLte) > 0 {
		params.SetConsolePortCountLte(vConsolePortCountLte)
	}
	if len(vConsolePortCountGt) > 0 {
		params.SetConsolePortCountGt(vConsolePortCountGt)
	}
	if len(vConsolePortCountGte) > 0 {
		params.SetConsolePortCountGte(vConsolePortCountGte)
	}
	if vConsolePortCountEmpty != nil {
		params.SetConsolePortCountEmpty(vConsolePortCountEmpty)
	}
	if vConsolePorts != nil {
		params.SetConsolePorts(vConsolePorts)
	}
	if len(vConsoleServerPortCount) > 0 {
		params.SetConsoleServerPortCount(vConsoleServerPortCount)
	}
	if len(vConsoleServerPortCountn) > 0 {
		params.SetConsoleServerPortCountn(vConsoleServerPortCountn)
	}
	if len(vConsoleServerPortCountLt) > 0 {
		params.SetConsoleServerPortCountLt(vConsoleServerPortCountLt)
	}
	if len(vConsoleServerPortCountLte) > 0 {
		params.SetConsoleServerPortCountLte(vConsoleServerPortCountLte)
	}
	if len(vConsoleServerPortCountGt) > 0 {
		params.SetConsoleServerPortCountGt(vConsoleServerPortCountGt)
	}
	if len(vConsoleServerPortCountGte) > 0 {
		params.SetConsoleServerPortCountGte(vConsoleServerPortCountGte)
	}
	if vConsoleServerPortCountEmpty != nil {
		params.SetConsoleServerPortCountEmpty(vConsoleServerPortCountEmpty)
	}
	if vConsoleServerPorts != nil {
		params.SetConsoleServerPorts(vConsoleServerPorts)
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
	if len(vDeviceBayCount) > 0 {
		params.SetDeviceBayCount(vDeviceBayCount)
	}
	if len(vDeviceBayCountn) > 0 {
		params.SetDeviceBayCountn(vDeviceBayCountn)
	}
	if len(vDeviceBayCountLt) > 0 {
		params.SetDeviceBayCountLt(vDeviceBayCountLt)
	}
	if len(vDeviceBayCountLte) > 0 {
		params.SetDeviceBayCountLte(vDeviceBayCountLte)
	}
	if len(vDeviceBayCountGt) > 0 {
		params.SetDeviceBayCountGt(vDeviceBayCountGt)
	}
	if len(vDeviceBayCountGte) > 0 {
		params.SetDeviceBayCountGte(vDeviceBayCountGte)
	}
	if vDeviceBayCountEmpty != nil {
		params.SetDeviceBayCountEmpty(vDeviceBayCountEmpty)
	}
	if vDeviceBays != nil {
		params.SetDeviceBays(vDeviceBays)
	}
	if len(vDeviceType) > 0 {
		params.SetDeviceType(vDeviceType)
	}
	if len(vDeviceTypen) > 0 {
		params.SetDeviceTypen(vDeviceTypen)
	}
	if vFace != nil {
		params.SetFace(vFace)
	}
	if vFacen != nil {
		params.SetFacen(vFacen)
	}
	if len(vFaceIc) > 0 {
		params.SetFaceIc(vFaceIc)
	}
	if len(vFaceNic) > 0 {
		params.SetFaceNic(vFaceNic)
	}
	if len(vFaceIe) > 0 {
		params.SetFaceIe(vFaceIe)
	}
	if len(vFaceNie) > 0 {
		params.SetFaceNie(vFaceNie)
	}
	if len(vFaceIsw) > 0 {
		params.SetFaceIsw(vFaceIsw)
	}
	if len(vFaceNisw) > 0 {
		params.SetFaceNisw(vFaceNisw)
	}
	if len(vFaceIew) > 0 {
		params.SetFaceIew(vFaceIew)
	}
	if len(vFaceNiew) > 0 {
		params.SetFaceNiew(vFaceNiew)
	}
	if vFaceEmpty != nil {
		params.SetFaceEmpty(vFaceEmpty)
	}
	if len(vFaceRegex) > 0 {
		params.SetFaceRegex(vFaceRegex)
	}
	if len(vFaceIregex) > 0 {
		params.SetFaceIregex(vFaceIregex)
	}
	if len(vFrontPortCount) > 0 {
		params.SetFrontPortCount(vFrontPortCount)
	}
	if len(vFrontPortCountn) > 0 {
		params.SetFrontPortCountn(vFrontPortCountn)
	}
	if len(vFrontPortCountLt) > 0 {
		params.SetFrontPortCountLt(vFrontPortCountLt)
	}
	if len(vFrontPortCountLte) > 0 {
		params.SetFrontPortCountLte(vFrontPortCountLte)
	}
	if len(vFrontPortCountGt) > 0 {
		params.SetFrontPortCountGt(vFrontPortCountGt)
	}
	if len(vFrontPortCountGte) > 0 {
		params.SetFrontPortCountGte(vFrontPortCountGte)
	}
	if vFrontPortCountEmpty != nil {
		params.SetFrontPortCountEmpty(vFrontPortCountEmpty)
	}
	if vHasOobIP != nil {
		params.SetHasOobIP(vHasOobIP)
	}
	if vHasPrimaryIP != nil {
		params.SetHasPrimaryIP(vHasPrimaryIP)
	}
	if vHasVirtualDeviceContext != nil {
		params.SetHasVirtualDeviceContext(vHasVirtualDeviceContext)
	}
	if len(vInterfaceCount) > 0 {
		params.SetInterfaceCount(vInterfaceCount)
	}
	if len(vInterfaceCountn) > 0 {
		params.SetInterfaceCountn(vInterfaceCountn)
	}
	if len(vInterfaceCountLt) > 0 {
		params.SetInterfaceCountLt(vInterfaceCountLt)
	}
	if len(vInterfaceCountLte) > 0 {
		params.SetInterfaceCountLte(vInterfaceCountLte)
	}
	if len(vInterfaceCountGt) > 0 {
		params.SetInterfaceCountGt(vInterfaceCountGt)
	}
	if len(vInterfaceCountGte) > 0 {
		params.SetInterfaceCountGte(vInterfaceCountGte)
	}
	if vInterfaceCountEmpty != nil {
		params.SetInterfaceCountEmpty(vInterfaceCountEmpty)
	}
	if vInterfaces != nil {
		params.SetInterfaces(vInterfaces)
	}
	if len(vInventoryItemCount) > 0 {
		params.SetInventoryItemCount(vInventoryItemCount)
	}
	if len(vInventoryItemCountn) > 0 {
		params.SetInventoryItemCountn(vInventoryItemCountn)
	}
	if len(vInventoryItemCountLt) > 0 {
		params.SetInventoryItemCountLt(vInventoryItemCountLt)
	}
	if len(vInventoryItemCountLte) > 0 {
		params.SetInventoryItemCountLte(vInventoryItemCountLte)
	}
	if len(vInventoryItemCountGt) > 0 {
		params.SetInventoryItemCountGt(vInventoryItemCountGt)
	}
	if len(vInventoryItemCountGte) > 0 {
		params.SetInventoryItemCountGte(vInventoryItemCountGte)
	}
	if vInventoryItemCountEmpty != nil {
		params.SetInventoryItemCountEmpty(vInventoryItemCountEmpty)
	}
	if vIsFullDepth != nil {
		params.SetIsFullDepth(vIsFullDepth)
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
	if len(vLatitude) > 0 {
		params.SetLatitude(vLatitude)
	}
	if len(vLatituden) > 0 {
		params.SetLatituden(vLatituden)
	}
	if len(vLatitudeLt) > 0 {
		params.SetLatitudeLt(vLatitudeLt)
	}
	if len(vLatitudeLte) > 0 {
		params.SetLatitudeLte(vLatitudeLte)
	}
	if len(vLatitudeGt) > 0 {
		params.SetLatitudeGt(vLatitudeGt)
	}
	if len(vLatitudeGte) > 0 {
		params.SetLatitudeGte(vLatitudeGte)
	}
	if vLatitudeEmpty != nil {
		params.SetLatitudeEmpty(vLatitudeEmpty)
	}
	if vLocalContextData != nil {
		params.SetLocalContextData(vLocalContextData)
	}
	if len(vLocation) > 0 {
		params.SetLocation(vLocation)
	}
	if len(vLocationn) > 0 {
		params.SetLocationn(vLocationn)
	}
	if len(vLongitude) > 0 {
		params.SetLongitude(vLongitude)
	}
	if len(vLongituden) > 0 {
		params.SetLongituden(vLongituden)
	}
	if len(vLongitudeLt) > 0 {
		params.SetLongitudeLt(vLongitudeLt)
	}
	if len(vLongitudeLte) > 0 {
		params.SetLongitudeLte(vLongitudeLte)
	}
	if len(vLongitudeGt) > 0 {
		params.SetLongitudeGt(vLongitudeGt)
	}
	if len(vLongitudeGte) > 0 {
		params.SetLongitudeGte(vLongitudeGte)
	}
	if vLongitudeEmpty != nil {
		params.SetLongitudeEmpty(vLongitudeEmpty)
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
	if len(vManufacturer) > 0 {
		params.SetManufacturer(vManufacturer)
	}
	if len(vManufacturern) > 0 {
		params.SetManufacturern(vManufacturern)
	}
	if len(vManufacturerID) > 0 {
		params.SetManufacturerID(vManufacturerID)
	}
	if len(vManufacturerIDn) > 0 {
		params.SetManufacturerIDn(vManufacturerIDn)
	}
	if len(vModel) > 0 {
		params.SetModel(vModel)
	}
	if len(vModeln) > 0 {
		params.SetModeln(vModeln)
	}
	if len(vModuleBayCount) > 0 {
		params.SetModuleBayCount(vModuleBayCount)
	}
	if len(vModuleBayCountn) > 0 {
		params.SetModuleBayCountn(vModuleBayCountn)
	}
	if len(vModuleBayCountLt) > 0 {
		params.SetModuleBayCountLt(vModuleBayCountLt)
	}
	if len(vModuleBayCountLte) > 0 {
		params.SetModuleBayCountLte(vModuleBayCountLte)
	}
	if len(vModuleBayCountGt) > 0 {
		params.SetModuleBayCountGt(vModuleBayCountGt)
	}
	if len(vModuleBayCountGte) > 0 {
		params.SetModuleBayCountGte(vModuleBayCountGte)
	}
	if vModuleBayCountEmpty != nil {
		params.SetModuleBayCountEmpty(vModuleBayCountEmpty)
	}
	if vModuleBays != nil {
		params.SetModuleBays(vModuleBays)
	}
	if len(vOobIPID) > 0 {
		params.SetOobIPID(vOobIPID)
	}
	if len(vOobIPIDn) > 0 {
		params.SetOobIPIDn(vOobIPIDn)
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
	if len(vParentBayID) > 0 {
		params.SetParentBayID(vParentBayID)
	}
	if len(vParentBayIDn) > 0 {
		params.SetParentBayIDn(vParentBayIDn)
	}
	if len(vParentDeviceID) > 0 {
		params.SetParentDeviceID(vParentDeviceID)
	}
	if len(vParentDeviceIDn) > 0 {
		params.SetParentDeviceIDn(vParentDeviceIDn)
	}
	if vPassThroughPorts != nil {
		params.SetPassThroughPorts(vPassThroughPorts)
	}
	if len(vPlatform) > 0 {
		params.SetPlatform(vPlatform)
	}
	if len(vPlatformn) > 0 {
		params.SetPlatformn(vPlatformn)
	}
	if len(vPosition) > 0 {
		params.SetPosition(vPosition)
	}
	if len(vPositionn) > 0 {
		params.SetPositionn(vPositionn)
	}
	if len(vPositionLt) > 0 {
		params.SetPositionLt(vPositionLt)
	}
	if len(vPositionLte) > 0 {
		params.SetPositionLte(vPositionLte)
	}
	if len(vPositionGt) > 0 {
		params.SetPositionGt(vPositionGt)
	}
	if len(vPositionGte) > 0 {
		params.SetPositionGte(vPositionGte)
	}
	if vPositionEmpty != nil {
		params.SetPositionEmpty(vPositionEmpty)
	}
	if len(vPowerOutletCount) > 0 {
		params.SetPowerOutletCount(vPowerOutletCount)
	}
	if len(vPowerOutletCountn) > 0 {
		params.SetPowerOutletCountn(vPowerOutletCountn)
	}
	if len(vPowerOutletCountLt) > 0 {
		params.SetPowerOutletCountLt(vPowerOutletCountLt)
	}
	if len(vPowerOutletCountLte) > 0 {
		params.SetPowerOutletCountLte(vPowerOutletCountLte)
	}
	if len(vPowerOutletCountGt) > 0 {
		params.SetPowerOutletCountGt(vPowerOutletCountGt)
	}
	if len(vPowerOutletCountGte) > 0 {
		params.SetPowerOutletCountGte(vPowerOutletCountGte)
	}
	if vPowerOutletCountEmpty != nil {
		params.SetPowerOutletCountEmpty(vPowerOutletCountEmpty)
	}
	if vPowerOutlets != nil {
		params.SetPowerOutlets(vPowerOutlets)
	}
	if len(vPowerPortCount) > 0 {
		params.SetPowerPortCount(vPowerPortCount)
	}
	if len(vPowerPortCountn) > 0 {
		params.SetPowerPortCountn(vPowerPortCountn)
	}
	if len(vPowerPortCountLt) > 0 {
		params.SetPowerPortCountLt(vPowerPortCountLt)
	}
	if len(vPowerPortCountLte) > 0 {
		params.SetPowerPortCountLte(vPowerPortCountLte)
	}
	if len(vPowerPortCountGt) > 0 {
		params.SetPowerPortCountGt(vPowerPortCountGt)
	}
	if len(vPowerPortCountGte) > 0 {
		params.SetPowerPortCountGte(vPowerPortCountGte)
	}
	if vPowerPortCountEmpty != nil {
		params.SetPowerPortCountEmpty(vPowerPortCountEmpty)
	}
	if vPowerPorts != nil {
		params.SetPowerPorts(vPowerPorts)
	}
	if len(vPrimaryIp4) > 0 {
		params.SetPrimaryIp4(vPrimaryIp4)
	}
	if len(vPrimaryIp4n) > 0 {
		params.SetPrimaryIp4n(vPrimaryIp4n)
	}
	if len(vPrimaryIp4ID) > 0 {
		params.SetPrimaryIp4ID(vPrimaryIp4ID)
	}
	if len(vPrimaryIp4IDn) > 0 {
		params.SetPrimaryIp4IDn(vPrimaryIp4IDn)
	}
	if len(vPrimaryIp6) > 0 {
		params.SetPrimaryIp6(vPrimaryIp6)
	}
	if len(vPrimaryIp6n) > 0 {
		params.SetPrimaryIp6n(vPrimaryIp6n)
	}
	if len(vPrimaryIp6ID) > 0 {
		params.SetPrimaryIp6ID(vPrimaryIp6ID)
	}
	if len(vPrimaryIp6IDn) > 0 {
		params.SetPrimaryIp6IDn(vPrimaryIp6IDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRearPortCount) > 0 {
		params.SetRearPortCount(vRearPortCount)
	}
	if len(vRearPortCountn) > 0 {
		params.SetRearPortCountn(vRearPortCountn)
	}
	if len(vRearPortCountLt) > 0 {
		params.SetRearPortCountLt(vRearPortCountLt)
	}
	if len(vRearPortCountLte) > 0 {
		params.SetRearPortCountLte(vRearPortCountLte)
	}
	if len(vRearPortCountGt) > 0 {
		params.SetRearPortCountGt(vRearPortCountGt)
	}
	if len(vRearPortCountGte) > 0 {
		params.SetRearPortCountGte(vRearPortCountGte)
	}
	if vRearPortCountEmpty != nil {
		params.SetRearPortCountEmpty(vRearPortCountEmpty)
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
	if len(vVcPosition) > 0 {
		params.SetVcPosition(vVcPosition)
	}
	if len(vVcPositionn) > 0 {
		params.SetVcPositionn(vVcPositionn)
	}
	if len(vVcPositionLt) > 0 {
		params.SetVcPositionLt(vVcPositionLt)
	}
	if len(vVcPositionLte) > 0 {
		params.SetVcPositionLte(vVcPositionLte)
	}
	if len(vVcPositionGt) > 0 {
		params.SetVcPositionGt(vVcPositionGt)
	}
	if len(vVcPositionGte) > 0 {
		params.SetVcPositionGte(vVcPositionGte)
	}
	if vVcPositionEmpty != nil {
		params.SetVcPositionEmpty(vVcPositionEmpty)
	}
	if len(vVcPriority) > 0 {
		params.SetVcPriority(vVcPriority)
	}
	if len(vVcPriorityn) > 0 {
		params.SetVcPriorityn(vVcPriorityn)
	}
	if len(vVcPriorityLt) > 0 {
		params.SetVcPriorityLt(vVcPriorityLt)
	}
	if len(vVcPriorityLte) > 0 {
		params.SetVcPriorityLte(vVcPriorityLte)
	}
	if len(vVcPriorityGt) > 0 {
		params.SetVcPriorityGt(vVcPriorityGt)
	}
	if len(vVcPriorityGte) > 0 {
		params.SetVcPriorityGte(vVcPriorityGte)
	}
	if vVcPriorityEmpty != nil {
		params.SetVcPriorityEmpty(vVcPriorityEmpty)
	}
	if len(vVirtualChassisID) > 0 {
		params.SetVirtualChassisID(vVirtualChassisID)
	}
	if len(vVirtualChassisIDn) > 0 {
		params.SetVirtualChassisIDn(vVirtualChassisIDn)
	}
	if vVirtualChassisMember != nil {
		params.SetVirtualChassisMember(vVirtualChassisMember)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Dcim.DcimDevicesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_devices", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.DeviceResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]deviceResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model deviceResourceModel
		resp.Diagnostics.Append(flattenDevice(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: deviceResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
