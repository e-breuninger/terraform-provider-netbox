// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                               netboxapi.FilterInt64s(params.SetID),
		"id__n":                            netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                           netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                          netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                           netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                          netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                        netboxapi.FilterBool(params.SetIDEmpty),
		"name":                             netboxapi.FilterStrings(params.SetName),
		"name__n":                          netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                         netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                        netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                         netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                        netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                        netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                       netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                        netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                       netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":                      netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":                      netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":                     netboxapi.FilterStrings(params.SetNameIregex),
		"serial":                           netboxapi.FilterStrings(params.SetSerial),
		"serial__n":                        netboxapi.FilterStrings(params.SetSerialn),
		"serial__ic":                       netboxapi.FilterStrings(params.SetSerialIc),
		"serial__nic":                      netboxapi.FilterStrings(params.SetSerialNic),
		"serial__ie":                       netboxapi.FilterStrings(params.SetSerialIe),
		"serial__nie":                      netboxapi.FilterStrings(params.SetSerialNie),
		"serial__isw":                      netboxapi.FilterStrings(params.SetSerialIsw),
		"serial__nisw":                     netboxapi.FilterStrings(params.SetSerialNisw),
		"serial__iew":                      netboxapi.FilterStrings(params.SetSerialIew),
		"serial__niew":                     netboxapi.FilterStrings(params.SetSerialNiew),
		"serial__empty":                    netboxapi.FilterBool(params.SetSerialEmpty),
		"serial__regex":                    netboxapi.FilterStrings(params.SetSerialRegex),
		"serial__iregex":                   netboxapi.FilterStrings(params.SetSerialIregex),
		"asset_tag":                        netboxapi.FilterStrings(params.SetAssetTag),
		"asset_tag__n":                     netboxapi.FilterStrings(params.SetAssetTagn),
		"asset_tag__ic":                    netboxapi.FilterStrings(params.SetAssetTagIc),
		"asset_tag__nic":                   netboxapi.FilterStrings(params.SetAssetTagNic),
		"asset_tag__ie":                    netboxapi.FilterStrings(params.SetAssetTagIe),
		"asset_tag__nie":                   netboxapi.FilterStrings(params.SetAssetTagNie),
		"asset_tag__isw":                   netboxapi.FilterStrings(params.SetAssetTagIsw),
		"asset_tag__nisw":                  netboxapi.FilterStrings(params.SetAssetTagNisw),
		"asset_tag__iew":                   netboxapi.FilterStrings(params.SetAssetTagIew),
		"asset_tag__niew":                  netboxapi.FilterStrings(params.SetAssetTagNiew),
		"asset_tag__empty":                 netboxapi.FilterBool(params.SetAssetTagEmpty),
		"asset_tag__regex":                 netboxapi.FilterStrings(params.SetAssetTagRegex),
		"asset_tag__iregex":                netboxapi.FilterStrings(params.SetAssetTagIregex),
		"status":                           netboxapi.FilterStrings(params.SetStatus),
		"status__n":                        netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":                       netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":                      netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":                       netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":                      netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":                      netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":                     netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":                      netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":                     netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":                    netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":                    netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":                   netboxapi.FilterStrings(params.SetStatusIregex),
		"site_id":                          netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":                       netboxapi.FilterInt64s(params.SetSiteIDn),
		"role_id":                          netboxapi.FilterStrings(params.SetRoleID),
		"role_id__n":                       netboxapi.FilterStrings(params.SetRoleIDn),
		"device_type_id":                   netboxapi.FilterInt64s(params.SetDeviceTypeID),
		"device_type_id__n":                netboxapi.FilterInt64s(params.SetDeviceTypeIDn),
		"tenant_id":                        netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":                     netboxapi.FilterInt64s(params.SetTenantIDn),
		"rack_id":                          netboxapi.FilterInt64s(params.SetRackID),
		"rack_id__n":                       netboxapi.FilterInt64s(params.SetRackIDn),
		"location_id":                      netboxapi.FilterStrings(params.SetLocationID),
		"location_id__n":                   netboxapi.FilterStrings(params.SetLocationIDn),
		"cluster_id":                       netboxapi.FilterInt64s(params.SetClusterID),
		"cluster_id__n":                    netboxapi.FilterInt64s(params.SetClusterIDn),
		"platform_id":                      netboxapi.FilterStrings(params.SetPlatformID),
		"platform_id__n":                   netboxapi.FilterStrings(params.SetPlatformIDn),
		"region":                           netboxapi.FilterStrings(params.SetRegion),
		"region__n":                        netboxapi.FilterStrings(params.SetRegionn),
		"owner_id":                         netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                      netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                              netboxapi.FilterStrings(params.SetTag),
		"tag__n":                           netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                         netboxapi.FilterStrings(params.SetTagAny),
		"airflow":                          netboxapi.FilterString(params.SetAirflow),
		"airflow__n":                       netboxapi.FilterString(params.SetAirflown),
		"airflow__ic":                      netboxapi.FilterStrings(params.SetAirflowIc),
		"airflow__nic":                     netboxapi.FilterStrings(params.SetAirflowNic),
		"airflow__ie":                      netboxapi.FilterStrings(params.SetAirflowIe),
		"airflow__nie":                     netboxapi.FilterStrings(params.SetAirflowNie),
		"airflow__isw":                     netboxapi.FilterStrings(params.SetAirflowIsw),
		"airflow__nisw":                    netboxapi.FilterStrings(params.SetAirflowNisw),
		"airflow__iew":                     netboxapi.FilterStrings(params.SetAirflowIew),
		"airflow__niew":                    netboxapi.FilterStrings(params.SetAirflowNiew),
		"airflow__empty":                   netboxapi.FilterBool(params.SetAirflowEmpty),
		"airflow__regex":                   netboxapi.FilterStrings(params.SetAirflowRegex),
		"airflow__iregex":                  netboxapi.FilterStrings(params.SetAirflowIregex),
		"cluster_group":                    netboxapi.FilterStrings(params.SetClusterGroup),
		"cluster_group__n":                 netboxapi.FilterStrings(params.SetClusterGroupn),
		"cluster_group_id":                 netboxapi.FilterInt64s(params.SetClusterGroupID),
		"cluster_group_id__n":              netboxapi.FilterInt64s(params.SetClusterGroupIDn),
		"config_template_id":               netboxapi.FilterInt64s(params.SetConfigTemplateID),
		"config_template_id__n":            netboxapi.FilterInt64s(params.SetConfigTemplateIDn),
		"console_port_count":               netboxapi.FilterInt64s(params.SetConsolePortCount),
		"console_port_count__n":            netboxapi.FilterInt64s(params.SetConsolePortCountn),
		"console_port_count__lt":           netboxapi.FilterInt64s(params.SetConsolePortCountLt),
		"console_port_count__lte":          netboxapi.FilterInt64s(params.SetConsolePortCountLte),
		"console_port_count__gt":           netboxapi.FilterInt64s(params.SetConsolePortCountGt),
		"console_port_count__gte":          netboxapi.FilterInt64s(params.SetConsolePortCountGte),
		"console_port_count__empty":        netboxapi.FilterBool(params.SetConsolePortCountEmpty),
		"console_ports":                    netboxapi.FilterBool(params.SetConsolePorts),
		"console_server_port_count":        netboxapi.FilterInt64s(params.SetConsoleServerPortCount),
		"console_server_port_count__n":     netboxapi.FilterInt64s(params.SetConsoleServerPortCountn),
		"console_server_port_count__lt":    netboxapi.FilterInt64s(params.SetConsoleServerPortCountLt),
		"console_server_port_count__lte":   netboxapi.FilterInt64s(params.SetConsoleServerPortCountLte),
		"console_server_port_count__gt":    netboxapi.FilterInt64s(params.SetConsoleServerPortCountGt),
		"console_server_port_count__gte":   netboxapi.FilterInt64s(params.SetConsoleServerPortCountGte),
		"console_server_port_count__empty": netboxapi.FilterBool(params.SetConsoleServerPortCountEmpty),
		"console_server_ports":             netboxapi.FilterBool(params.SetConsoleServerPorts),
		"contact":                          netboxapi.FilterInt64s(params.SetContact),
		"contact__n":                       netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":                    netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":                 netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":                     netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":                  netboxapi.FilterInt64s(params.SetContactRolen),
		"created":                          netboxapi.FilterStrings(params.SetCreated),
		"created__n":                       netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                      netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":                     netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                      netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":                     netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":                   netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":                      netboxapi.FilterStrings(params.SetDescription),
		"description__n":                   netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":                  netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":                 netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":                  netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":                 netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":                 netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":                netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":                 netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":                netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":               netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":               netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":              netboxapi.FilterStrings(params.SetDescriptionIregex),
		"device_bay_count":                 netboxapi.FilterInt64s(params.SetDeviceBayCount),
		"device_bay_count__n":              netboxapi.FilterInt64s(params.SetDeviceBayCountn),
		"device_bay_count__lt":             netboxapi.FilterInt64s(params.SetDeviceBayCountLt),
		"device_bay_count__lte":            netboxapi.FilterInt64s(params.SetDeviceBayCountLte),
		"device_bay_count__gt":             netboxapi.FilterInt64s(params.SetDeviceBayCountGt),
		"device_bay_count__gte":            netboxapi.FilterInt64s(params.SetDeviceBayCountGte),
		"device_bay_count__empty":          netboxapi.FilterBool(params.SetDeviceBayCountEmpty),
		"device_bays":                      netboxapi.FilterBool(params.SetDeviceBays),
		"device_type":                      netboxapi.FilterStrings(params.SetDeviceType),
		"device_type__n":                   netboxapi.FilterStrings(params.SetDeviceTypen),
		"face":                             netboxapi.FilterString(params.SetFace),
		"face__n":                          netboxapi.FilterString(params.SetFacen),
		"face__ic":                         netboxapi.FilterStrings(params.SetFaceIc),
		"face__nic":                        netboxapi.FilterStrings(params.SetFaceNic),
		"face__ie":                         netboxapi.FilterStrings(params.SetFaceIe),
		"face__nie":                        netboxapi.FilterStrings(params.SetFaceNie),
		"face__isw":                        netboxapi.FilterStrings(params.SetFaceIsw),
		"face__nisw":                       netboxapi.FilterStrings(params.SetFaceNisw),
		"face__iew":                        netboxapi.FilterStrings(params.SetFaceIew),
		"face__niew":                       netboxapi.FilterStrings(params.SetFaceNiew),
		"face__empty":                      netboxapi.FilterBool(params.SetFaceEmpty),
		"face__regex":                      netboxapi.FilterStrings(params.SetFaceRegex),
		"face__iregex":                     netboxapi.FilterStrings(params.SetFaceIregex),
		"front_port_count":                 netboxapi.FilterInt64s(params.SetFrontPortCount),
		"front_port_count__n":              netboxapi.FilterInt64s(params.SetFrontPortCountn),
		"front_port_count__lt":             netboxapi.FilterInt64s(params.SetFrontPortCountLt),
		"front_port_count__lte":            netboxapi.FilterInt64s(params.SetFrontPortCountLte),
		"front_port_count__gt":             netboxapi.FilterInt64s(params.SetFrontPortCountGt),
		"front_port_count__gte":            netboxapi.FilterInt64s(params.SetFrontPortCountGte),
		"front_port_count__empty":          netboxapi.FilterBool(params.SetFrontPortCountEmpty),
		"has_oob_ip":                       netboxapi.FilterBool(params.SetHasOobIP),
		"has_primary_ip":                   netboxapi.FilterBool(params.SetHasPrimaryIP),
		"has_virtual_device_context":       netboxapi.FilterBool(params.SetHasVirtualDeviceContext),
		"interface_count":                  netboxapi.FilterInt64s(params.SetInterfaceCount),
		"interface_count__n":               netboxapi.FilterInt64s(params.SetInterfaceCountn),
		"interface_count__lt":              netboxapi.FilterInt64s(params.SetInterfaceCountLt),
		"interface_count__lte":             netboxapi.FilterInt64s(params.SetInterfaceCountLte),
		"interface_count__gt":              netboxapi.FilterInt64s(params.SetInterfaceCountGt),
		"interface_count__gte":             netboxapi.FilterInt64s(params.SetInterfaceCountGte),
		"interface_count__empty":           netboxapi.FilterBool(params.SetInterfaceCountEmpty),
		"interfaces":                       netboxapi.FilterBool(params.SetInterfaces),
		"inventory_item_count":             netboxapi.FilterInt64s(params.SetInventoryItemCount),
		"inventory_item_count__n":          netboxapi.FilterInt64s(params.SetInventoryItemCountn),
		"inventory_item_count__lt":         netboxapi.FilterInt64s(params.SetInventoryItemCountLt),
		"inventory_item_count__lte":        netboxapi.FilterInt64s(params.SetInventoryItemCountLte),
		"inventory_item_count__gt":         netboxapi.FilterInt64s(params.SetInventoryItemCountGt),
		"inventory_item_count__gte":        netboxapi.FilterInt64s(params.SetInventoryItemCountGte),
		"inventory_item_count__empty":      netboxapi.FilterBool(params.SetInventoryItemCountEmpty),
		"is_full_depth":                    netboxapi.FilterBool(params.SetIsFullDepth),
		"last_updated":                     netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":                  netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":                 netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":                netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":                 netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":                netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":              netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"latitude":                         netboxapi.FilterFloat64s(params.SetLatitude),
		"latitude__n":                      netboxapi.FilterFloat64s(params.SetLatituden),
		"latitude__lt":                     netboxapi.FilterFloat64s(params.SetLatitudeLt),
		"latitude__lte":                    netboxapi.FilterFloat64s(params.SetLatitudeLte),
		"latitude__gt":                     netboxapi.FilterFloat64s(params.SetLatitudeGt),
		"latitude__gte":                    netboxapi.FilterFloat64s(params.SetLatitudeGte),
		"latitude__empty":                  netboxapi.FilterBool(params.SetLatitudeEmpty),
		"local_context_data":               netboxapi.FilterBool(params.SetLocalContextData),
		"location":                         netboxapi.FilterStrings(params.SetLocation),
		"location__n":                      netboxapi.FilterStrings(params.SetLocationn),
		"longitude":                        netboxapi.FilterFloat64s(params.SetLongitude),
		"longitude__n":                     netboxapi.FilterFloat64s(params.SetLongituden),
		"longitude__lt":                    netboxapi.FilterFloat64s(params.SetLongitudeLt),
		"longitude__lte":                   netboxapi.FilterFloat64s(params.SetLongitudeLte),
		"longitude__gt":                    netboxapi.FilterFloat64s(params.SetLongitudeGt),
		"longitude__gte":                   netboxapi.FilterFloat64s(params.SetLongitudeGte),
		"longitude__empty":                 netboxapi.FilterBool(params.SetLongitudeEmpty),
		"mac_address":                      netboxapi.FilterStrings(params.SetMacAddress),
		"mac_address__n":                   netboxapi.FilterStrings(params.SetMacAddressn),
		"mac_address__ic":                  netboxapi.FilterStrings(params.SetMacAddressIc),
		"mac_address__nic":                 netboxapi.FilterStrings(params.SetMacAddressNic),
		"mac_address__ie":                  netboxapi.FilterStrings(params.SetMacAddressIe),
		"mac_address__nie":                 netboxapi.FilterStrings(params.SetMacAddressNie),
		"mac_address__isw":                 netboxapi.FilterStrings(params.SetMacAddressIsw),
		"mac_address__nisw":                netboxapi.FilterStrings(params.SetMacAddressNisw),
		"mac_address__iew":                 netboxapi.FilterStrings(params.SetMacAddressIew),
		"mac_address__niew":                netboxapi.FilterStrings(params.SetMacAddressNiew),
		"mac_address__regex":               netboxapi.FilterStrings(params.SetMacAddressRegex),
		"mac_address__iregex":              netboxapi.FilterStrings(params.SetMacAddressIregex),
		"manufacturer":                     netboxapi.FilterStrings(params.SetManufacturer),
		"manufacturer__n":                  netboxapi.FilterStrings(params.SetManufacturern),
		"manufacturer_id":                  netboxapi.FilterInt64s(params.SetManufacturerID),
		"manufacturer_id__n":               netboxapi.FilterInt64s(params.SetManufacturerIDn),
		"model":                            netboxapi.FilterStrings(params.SetModel),
		"model__n":                         netboxapi.FilterStrings(params.SetModeln),
		"module_bay_count":                 netboxapi.FilterInt64s(params.SetModuleBayCount),
		"module_bay_count__n":              netboxapi.FilterInt64s(params.SetModuleBayCountn),
		"module_bay_count__lt":             netboxapi.FilterInt64s(params.SetModuleBayCountLt),
		"module_bay_count__lte":            netboxapi.FilterInt64s(params.SetModuleBayCountLte),
		"module_bay_count__gt":             netboxapi.FilterInt64s(params.SetModuleBayCountGt),
		"module_bay_count__gte":            netboxapi.FilterInt64s(params.SetModuleBayCountGte),
		"module_bay_count__empty":          netboxapi.FilterBool(params.SetModuleBayCountEmpty),
		"module_bays":                      netboxapi.FilterBool(params.SetModuleBays),
		"oob_ip_id":                        netboxapi.FilterInt64s(params.SetOobIPID),
		"oob_ip_id__n":                     netboxapi.FilterInt64s(params.SetOobIPIDn),
		"owner":                            netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                         netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                      netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":                   netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":                   netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":                netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"parent_bay_id":                    netboxapi.FilterInt64s(params.SetParentBayID),
		"parent_bay_id__n":                 netboxapi.FilterInt64s(params.SetParentBayIDn),
		"parent_device_id":                 netboxapi.FilterInt64s(params.SetParentDeviceID),
		"parent_device_id__n":              netboxapi.FilterInt64s(params.SetParentDeviceIDn),
		"pass_through_ports":               netboxapi.FilterBool(params.SetPassThroughPorts),
		"platform":                         netboxapi.FilterStrings(params.SetPlatform),
		"platform__n":                      netboxapi.FilterStrings(params.SetPlatformn),
		"position":                         netboxapi.FilterFloat64s(params.SetPosition),
		"position__n":                      netboxapi.FilterFloat64s(params.SetPositionn),
		"position__lt":                     netboxapi.FilterFloat64s(params.SetPositionLt),
		"position__lte":                    netboxapi.FilterFloat64s(params.SetPositionLte),
		"position__gt":                     netboxapi.FilterFloat64s(params.SetPositionGt),
		"position__gte":                    netboxapi.FilterFloat64s(params.SetPositionGte),
		"position__empty":                  netboxapi.FilterBool(params.SetPositionEmpty),
		"power_outlet_count":               netboxapi.FilterInt64s(params.SetPowerOutletCount),
		"power_outlet_count__n":            netboxapi.FilterInt64s(params.SetPowerOutletCountn),
		"power_outlet_count__lt":           netboxapi.FilterInt64s(params.SetPowerOutletCountLt),
		"power_outlet_count__lte":          netboxapi.FilterInt64s(params.SetPowerOutletCountLte),
		"power_outlet_count__gt":           netboxapi.FilterInt64s(params.SetPowerOutletCountGt),
		"power_outlet_count__gte":          netboxapi.FilterInt64s(params.SetPowerOutletCountGte),
		"power_outlet_count__empty":        netboxapi.FilterBool(params.SetPowerOutletCountEmpty),
		"power_outlets":                    netboxapi.FilterBool(params.SetPowerOutlets),
		"power_port_count":                 netboxapi.FilterInt64s(params.SetPowerPortCount),
		"power_port_count__n":              netboxapi.FilterInt64s(params.SetPowerPortCountn),
		"power_port_count__lt":             netboxapi.FilterInt64s(params.SetPowerPortCountLt),
		"power_port_count__lte":            netboxapi.FilterInt64s(params.SetPowerPortCountLte),
		"power_port_count__gt":             netboxapi.FilterInt64s(params.SetPowerPortCountGt),
		"power_port_count__gte":            netboxapi.FilterInt64s(params.SetPowerPortCountGte),
		"power_port_count__empty":          netboxapi.FilterBool(params.SetPowerPortCountEmpty),
		"power_ports":                      netboxapi.FilterBool(params.SetPowerPorts),
		"primary_ip4":                      netboxapi.FilterStrings(params.SetPrimaryIp4),
		"primary_ip4__n":                   netboxapi.FilterStrings(params.SetPrimaryIp4n),
		"primary_ip4_id":                   netboxapi.FilterInt64s(params.SetPrimaryIp4ID),
		"primary_ip4_id__n":                netboxapi.FilterInt64s(params.SetPrimaryIp4IDn),
		"primary_ip6":                      netboxapi.FilterStrings(params.SetPrimaryIp6),
		"primary_ip6__n":                   netboxapi.FilterStrings(params.SetPrimaryIp6n),
		"primary_ip6_id":                   netboxapi.FilterInt64s(params.SetPrimaryIp6ID),
		"primary_ip6_id__n":                netboxapi.FilterInt64s(params.SetPrimaryIp6IDn),
		"q":                                netboxapi.FilterString(params.SetQ),
		"rear_port_count":                  netboxapi.FilterInt64s(params.SetRearPortCount),
		"rear_port_count__n":               netboxapi.FilterInt64s(params.SetRearPortCountn),
		"rear_port_count__lt":              netboxapi.FilterInt64s(params.SetRearPortCountLt),
		"rear_port_count__lte":             netboxapi.FilterInt64s(params.SetRearPortCountLte),
		"rear_port_count__gt":              netboxapi.FilterInt64s(params.SetRearPortCountGt),
		"rear_port_count__gte":             netboxapi.FilterInt64s(params.SetRearPortCountGte),
		"rear_port_count__empty":           netboxapi.FilterBool(params.SetRearPortCountEmpty),
		"region_id":                        netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":                     netboxapi.FilterStrings(params.SetRegionIDn),
		"role":                             netboxapi.FilterStrings(params.SetRole),
		"role__n":                          netboxapi.FilterStrings(params.SetRolen),
		"site":                             netboxapi.FilterStrings(params.SetSite),
		"site__n":                          netboxapi.FilterStrings(params.SetSiten),
		"site_group":                       netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":                    netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":                    netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":                 netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"tag_id":                           netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                        netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                      netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                           netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":                        netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":                     netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":                  netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":                  netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":               netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"vc_position":                      netboxapi.FilterInt64s(params.SetVcPosition),
		"vc_position__n":                   netboxapi.FilterInt64s(params.SetVcPositionn),
		"vc_position__lt":                  netboxapi.FilterInt64s(params.SetVcPositionLt),
		"vc_position__lte":                 netboxapi.FilterInt64s(params.SetVcPositionLte),
		"vc_position__gt":                  netboxapi.FilterInt64s(params.SetVcPositionGt),
		"vc_position__gte":                 netboxapi.FilterInt64s(params.SetVcPositionGte),
		"vc_position__empty":               netboxapi.FilterBool(params.SetVcPositionEmpty),
		"vc_priority":                      netboxapi.FilterInt64s(params.SetVcPriority),
		"vc_priority__n":                   netboxapi.FilterInt64s(params.SetVcPriorityn),
		"vc_priority__lt":                  netboxapi.FilterInt64s(params.SetVcPriorityLt),
		"vc_priority__lte":                 netboxapi.FilterInt64s(params.SetVcPriorityLte),
		"vc_priority__gt":                  netboxapi.FilterInt64s(params.SetVcPriorityGt),
		"vc_priority__gte":                 netboxapi.FilterInt64s(params.SetVcPriorityGte),
		"vc_priority__empty":               netboxapi.FilterBool(params.SetVcPriorityEmpty),
		"virtual_chassis_id":               netboxapi.FilterInt64s(params.SetVirtualChassisID),
		"virtual_chassis_id__n":            netboxapi.FilterInt64s(params.SetVirtualChassisIDn),
		"virtual_chassis_member":           netboxapi.FilterBool(params.SetVirtualChassisMember),
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
