// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*deviceDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceDataSource)(nil)
)

// NewDeviceDataSource returns a new device data source.
func NewDeviceDataSource() datasource.DataSource {
	return &deviceDataSource{}
}

type deviceDataSource struct {
	client *netboxapi.Client
}

// deviceDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type deviceDataSourceModel struct {
	deviceResourceModel
	RegionSlug              types.String  `tfsdk:"region_slug"`
	TagSlug                 types.String  `tfsdk:"tag_slug"`
	ClusterGroupSlug        types.String  `tfsdk:"cluster_group_slug"`
	ClusterGroupID          types.Int64   `tfsdk:"cluster_group_id"`
	ConsolePorts            types.Bool    `tfsdk:"console_ports"`
	ConsoleServerPorts      types.Bool    `tfsdk:"console_server_ports"`
	Contact                 types.Int64   `tfsdk:"contact"`
	ContactGroupSlug        types.String  `tfsdk:"contact_group_slug"`
	ContactRole             types.Int64   `tfsdk:"contact_role"`
	DeviceBays              types.Bool    `tfsdk:"device_bays"`
	DeviceTypeSlug          types.String  `tfsdk:"device_type_slug"`
	Face                    types.String  `tfsdk:"face"`
	HasOobIP                types.Bool    `tfsdk:"has_oob_ip"`
	HasPrimaryIP            types.Bool    `tfsdk:"has_primary_ip"`
	HasVirtualDeviceContext types.Bool    `tfsdk:"has_virtual_device_context"`
	Interfaces              types.Bool    `tfsdk:"interfaces"`
	IsFullDepth             types.Bool    `tfsdk:"is_full_depth"`
	HasLocalContextData     types.Bool    `tfsdk:"has_local_context_data"`
	LocationSlug            types.String  `tfsdk:"location_slug"`
	MacAddress              types.String  `tfsdk:"mac_address"`
	ManufacturerSlug        types.String  `tfsdk:"manufacturer_slug"`
	ManufacturerID          types.Int64   `tfsdk:"manufacturer_id"`
	ModelSlug               types.String  `tfsdk:"model_slug"`
	ModuleBays              types.Bool    `tfsdk:"module_bays"`
	OwnerName               types.String  `tfsdk:"owner_name"`
	OwnerGroupName          types.String  `tfsdk:"owner_group_name"`
	OwnerGroupID            types.Int64   `tfsdk:"owner_group_id"`
	ParentBayID             types.Int64   `tfsdk:"parent_bay_id"`
	ParentDeviceID          types.Int64   `tfsdk:"parent_device_id"`
	PassThroughPorts        types.Bool    `tfsdk:"pass_through_ports"`
	Platform                types.String  `tfsdk:"platform"`
	Position                types.Float64 `tfsdk:"position"`
	PowerOutlets            types.Bool    `tfsdk:"power_outlets"`
	PowerPorts              types.Bool    `tfsdk:"power_ports"`
	PrimaryIp4              types.String  `tfsdk:"primary_ip4"`
	PrimaryIp6              types.String  `tfsdk:"primary_ip6"`
	Q                       types.String  `tfsdk:"q"`
	RegionID                types.String  `tfsdk:"region_id"`
	RoleSlug                types.String  `tfsdk:"role_slug"`
	SiteSlug                types.String  `tfsdk:"site_slug"`
	SiteGroupSlug           types.String  `tfsdk:"site_group_slug"`
	SiteGroupID             types.String  `tfsdk:"site_group_id"`
	TagID                   types.Int64   `tfsdk:"tag_id"`
	TenantSlug              types.String  `tfsdk:"tenant_slug"`
	TenantGroupSlug         types.String  `tfsdk:"tenant_group_slug"`
	TenantGroupID           types.String  `tfsdk:"tenant_group_id"`
	VcPosition              types.Int64   `tfsdk:"vc_position"`
	VcPriority              types.Int64   `tfsdk:"vc_priority"`
	VirtualChassisMember    types.Bool    `tfsdk:"virtual_chassis_member"`
	CustomFieldFilters      types.Map     `tfsdk:"custom_field_filters"`
}

func (d *deviceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device"
}

func (d *deviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A physical device (dcim.device). primary_ip4_id/primary_ip6_id are set through netbox_primary_ip.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/device/):\n\n> Every piece of hardware which is installed within a site or rack exists in NetBox as a device. Devices are measured in rack units (U) and can be half depth or full depth. A device may have a height of 0U: These devices do not consume vertical rack space and cannot be assigned to a particular rack unit. A common example of a 0U device is a vertically-mounted PDU.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Device name; NetBox allows unnamed devices.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"device_type_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device type.",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device role.",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the site.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the location within the site.",
			},
			"rack_id": schema.Int64Attribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "GPS coordinate in decimal format (xx.yyyyyy).",
				Validators: []validator.Float64{
					float64validator.Between(-90.0, 90.0),
				},
			},
			"longitude": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "GPS coordinate in decimal format (xx.yyyyyy).",
				Validators: []validator.Float64{
					float64validator.Between(-180.0, 180.0),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: offline, active, planned, staged, failed, inventory, decommissioning.",
				Validators: []validator.String{
					stringvalidator.OneOf("offline", "active", "planned", "staged", "failed", "inventory", "decommissioning"),
				},
			},
			"airflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
				Validators: []validator.String{
					stringvalidator.OneOf("front-to-rear", "rear-to-front", "left-to-right", "right-to-left", "side-to-rear", "passive", "mixed", "rear-to-side", "bottom-to-top", "top-to-bottom"),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"platform_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the platform.",
			},
			"cluster_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the virtualization cluster the device hosts.",
			},
			"virtual_chassis_id": schema.Int64Attribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "Id of the config template.",
			},
			"serial": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Chassis serial number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"asset_tag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique asset tag.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"primary_ip4_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the primary IPv4 address (managed by netbox_primary_ip).",
			},
			"primary_ip6_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the primary IPv6 address (managed by netbox_primary_ip).",
			},
			"oob_ip_id": schema.Int64Attribute{
				Optional:    true,
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
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Computed: true,
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
			"console_port_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of console ports on the device.",
			},
			"console_server_port_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of console server ports on the device.",
			},
			"power_port_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of power ports on the device.",
			},
			"power_outlet_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of power outlets on the device.",
			},
			"interface_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of interfaces on the device.",
			},
			"front_port_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of front ports on the device.",
			},
			"rear_port_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of rear ports on the device.",
			},
			"device_bay_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of device bays on the device.",
			},
			"module_bay_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of module bays on the device.",
			},
			"inventory_item_count": schema.Int64Attribute{
				Optional:    true,
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
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region of the site.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"cluster_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the cluster group.",
			},
			"cluster_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster group.",
			},
			"console_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has console ports.",
			},
			"console_server_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has console server ports.",
			},
			"contact": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact.",
			},
			"contact_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the contact group.",
			},
			"contact_role": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact role.",
			},
			"device_bays": schema.BoolAttribute{
				Optional:    true,
				Description: "Has device bays.",
			},
			"device_type_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the device type.",
			},
			"face": schema.StringAttribute{
				Optional:    true,
				Description: "Face. One of: front, rear.",
			},
			"has_oob_ip": schema.BoolAttribute{
				Optional:    true,
				Description: "Has an out-of-band IP.",
			},
			"has_primary_ip": schema.BoolAttribute{
				Optional:    true,
				Description: "Has a primary IP.",
			},
			"has_virtual_device_context": schema.BoolAttribute{
				Optional:    true,
				Description: "Has virtual device context.",
			},
			"interfaces": schema.BoolAttribute{
				Optional:    true,
				Description: "Has interfaces.",
			},
			"is_full_depth": schema.BoolAttribute{
				Optional:    true,
				Description: "Is full depth.",
			},
			"has_local_context_data": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the object has local config context data.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Description: "Mac address to match.",
			},
			"manufacturer_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the manufacturer.",
			},
			"manufacturer_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the manufacturer.",
			},
			"model_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the device model.",
			},
			"module_bays": schema.BoolAttribute{
				Optional:    true,
				Description: "Has module bays.",
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
			"parent_bay_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the parent bay.",
			},
			"parent_device_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the parent device.",
			},
			"pass_through_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has pass-through ports.",
			},
			"platform": schema.StringAttribute{
				Optional:    true,
				Description: "Platform to match.",
			},
			"position": schema.Float64Attribute{
				Optional:    true,
				Description: "Position to match.",
			},
			"power_outlets": schema.BoolAttribute{
				Optional:    true,
				Description: "Has power outlets.",
			},
			"power_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has power ports.",
			},
			"primary_ip4": schema.StringAttribute{
				Optional:    true,
				Description: "Primary IPv4 (address).",
			},
			"primary_ip6": schema.StringAttribute{
				Optional:    true,
				Description: "Primary IPv6 (address).",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"region_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"role_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the role.",
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
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"tenant_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant.",
			},
			"tenant_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant group.",
			},
			"tenant_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the tenant group.",
			},
			"vc_position": schema.Int64Attribute{
				Optional:    true,
				Description: "Vc position to match.",
			},
			"vc_priority": schema.Int64Attribute{
				Optional:    true,
				Description: "Vc priority to match.",
			},
			"virtual_chassis_member": schema.BoolAttribute{
				Optional:    true,
				Description: "Is a virtual chassis member.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *deviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimDevicesListParams()
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
	if !state.Serial.IsNull() {
		v := state.Serial.ValueString()
		params.SetSerial([]string{v})
		hasInput = true
	}
	if !state.AssetTag.IsNull() {
		v := state.AssetTag.ValueString()
		params.SetAssetTag([]string{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.RoleID.IsNull() {
		v := strconv.FormatInt(state.RoleID.ValueInt64(), 10)
		params.SetRoleID([]string{v})
		hasInput = true
	}
	if !state.DeviceTypeID.IsNull() {
		v := state.DeviceTypeID.ValueInt64()
		params.SetDeviceTypeID([]int64{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.RackID.IsNull() {
		v := state.RackID.ValueInt64()
		params.SetRackID([]int64{v})
		hasInput = true
	}
	if !state.LocationID.IsNull() {
		v := strconv.FormatInt(state.LocationID.ValueInt64(), 10)
		params.SetLocationID([]string{v})
		hasInput = true
	}
	if !state.ClusterID.IsNull() {
		v := state.ClusterID.ValueInt64()
		params.SetClusterID([]int64{v})
		hasInput = true
	}
	if !state.PlatformID.IsNull() {
		v := strconv.FormatInt(state.PlatformID.ValueInt64(), 10)
		params.SetPlatformID([]string{v})
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
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
	if !state.Airflow.IsNull() {
		v := state.Airflow.ValueString()
		params.SetAirflow(&v)
		hasInput = true
	}
	if !state.ClusterGroupSlug.IsNull() {
		v := state.ClusterGroupSlug.ValueString()
		params.SetClusterGroup([]string{v})
		hasInput = true
	}
	if !state.ClusterGroupID.IsNull() {
		v := state.ClusterGroupID.ValueInt64()
		params.SetClusterGroupID([]int64{v})
		hasInput = true
	}
	if !state.ConfigTemplateID.IsNull() {
		v := state.ConfigTemplateID.ValueInt64()
		params.SetConfigTemplateID([]int64{v})
		hasInput = true
	}
	if !state.ConsolePortCount.IsNull() {
		v := state.ConsolePortCount.ValueInt64()
		params.SetConsolePortCount([]int64{v})
		hasInput = true
	}
	if !state.ConsolePorts.IsNull() {
		v := state.ConsolePorts.ValueBool()
		params.SetConsolePorts(&v)
		hasInput = true
	}
	if !state.ConsoleServerPortCount.IsNull() {
		v := state.ConsoleServerPortCount.ValueInt64()
		params.SetConsoleServerPortCount([]int64{v})
		hasInput = true
	}
	if !state.ConsoleServerPorts.IsNull() {
		v := state.ConsoleServerPorts.ValueBool()
		params.SetConsoleServerPorts(&v)
		hasInput = true
	}
	if !state.Contact.IsNull() {
		v := state.Contact.ValueInt64()
		params.SetContact([]int64{v})
		hasInput = true
	}
	if !state.ContactGroupSlug.IsNull() {
		v := state.ContactGroupSlug.ValueString()
		params.SetContactGroup([]string{v})
		hasInput = true
	}
	if !state.ContactRole.IsNull() {
		v := state.ContactRole.ValueInt64()
		params.SetContactRole([]int64{v})
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
	if !state.DeviceBayCount.IsNull() {
		v := state.DeviceBayCount.ValueInt64()
		params.SetDeviceBayCount([]int64{v})
		hasInput = true
	}
	if !state.DeviceBays.IsNull() {
		v := state.DeviceBays.ValueBool()
		params.SetDeviceBays(&v)
		hasInput = true
	}
	if !state.DeviceTypeSlug.IsNull() {
		v := state.DeviceTypeSlug.ValueString()
		params.SetDeviceType([]string{v})
		hasInput = true
	}
	if !state.Face.IsNull() {
		v := state.Face.ValueString()
		params.SetFace(&v)
		hasInput = true
	}
	if !state.FrontPortCount.IsNull() {
		v := state.FrontPortCount.ValueInt64()
		params.SetFrontPortCount([]int64{v})
		hasInput = true
	}
	if !state.HasOobIP.IsNull() {
		v := state.HasOobIP.ValueBool()
		params.SetHasOobIP(&v)
		hasInput = true
	}
	if !state.HasPrimaryIP.IsNull() {
		v := state.HasPrimaryIP.ValueBool()
		params.SetHasPrimaryIP(&v)
		hasInput = true
	}
	if !state.HasVirtualDeviceContext.IsNull() {
		v := state.HasVirtualDeviceContext.ValueBool()
		params.SetHasVirtualDeviceContext(&v)
		hasInput = true
	}
	if !state.InterfaceCount.IsNull() {
		v := state.InterfaceCount.ValueInt64()
		params.SetInterfaceCount([]int64{v})
		hasInput = true
	}
	if !state.Interfaces.IsNull() {
		v := state.Interfaces.ValueBool()
		params.SetInterfaces(&v)
		hasInput = true
	}
	if !state.InventoryItemCount.IsNull() {
		v := state.InventoryItemCount.ValueInt64()
		params.SetInventoryItemCount([]int64{v})
		hasInput = true
	}
	if !state.IsFullDepth.IsNull() {
		v := state.IsFullDepth.ValueBool()
		params.SetIsFullDepth(&v)
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.Latitude.IsNull() {
		v := state.Latitude.ValueFloat64()
		params.SetLatitude([]float64{v})
		hasInput = true
	}
	if !state.HasLocalContextData.IsNull() {
		v := state.HasLocalContextData.ValueBool()
		params.SetLocalContextData(&v)
		hasInput = true
	}
	if !state.LocationSlug.IsNull() {
		v := state.LocationSlug.ValueString()
		params.SetLocation([]string{v})
		hasInput = true
	}
	if !state.Longitude.IsNull() {
		v := state.Longitude.ValueFloat64()
		params.SetLongitude([]float64{v})
		hasInput = true
	}
	if !state.MacAddress.IsNull() {
		v := state.MacAddress.ValueString()
		params.SetMacAddress([]string{v})
		hasInput = true
	}
	if !state.ManufacturerSlug.IsNull() {
		v := state.ManufacturerSlug.ValueString()
		params.SetManufacturer([]string{v})
		hasInput = true
	}
	if !state.ManufacturerID.IsNull() {
		v := state.ManufacturerID.ValueInt64()
		params.SetManufacturerID([]int64{v})
		hasInput = true
	}
	if !state.ModelSlug.IsNull() {
		v := state.ModelSlug.ValueString()
		params.SetModel([]string{v})
		hasInput = true
	}
	if !state.ModuleBayCount.IsNull() {
		v := state.ModuleBayCount.ValueInt64()
		params.SetModuleBayCount([]int64{v})
		hasInput = true
	}
	if !state.ModuleBays.IsNull() {
		v := state.ModuleBays.ValueBool()
		params.SetModuleBays(&v)
		hasInput = true
	}
	if !state.OobIPID.IsNull() {
		v := state.OobIPID.ValueInt64()
		params.SetOobIPID([]int64{v})
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
	if !state.ParentBayID.IsNull() {
		v := state.ParentBayID.ValueInt64()
		params.SetParentBayID([]int64{v})
		hasInput = true
	}
	if !state.ParentDeviceID.IsNull() {
		v := state.ParentDeviceID.ValueInt64()
		params.SetParentDeviceID([]int64{v})
		hasInput = true
	}
	if !state.PassThroughPorts.IsNull() {
		v := state.PassThroughPorts.ValueBool()
		params.SetPassThroughPorts(&v)
		hasInput = true
	}
	if !state.Platform.IsNull() {
		v := state.Platform.ValueString()
		params.SetPlatform([]string{v})
		hasInput = true
	}
	if !state.Position.IsNull() {
		v := state.Position.ValueFloat64()
		params.SetPosition([]float64{v})
		hasInput = true
	}
	if !state.PowerOutletCount.IsNull() {
		v := state.PowerOutletCount.ValueInt64()
		params.SetPowerOutletCount([]int64{v})
		hasInput = true
	}
	if !state.PowerOutlets.IsNull() {
		v := state.PowerOutlets.ValueBool()
		params.SetPowerOutlets(&v)
		hasInput = true
	}
	if !state.PowerPortCount.IsNull() {
		v := state.PowerPortCount.ValueInt64()
		params.SetPowerPortCount([]int64{v})
		hasInput = true
	}
	if !state.PowerPorts.IsNull() {
		v := state.PowerPorts.ValueBool()
		params.SetPowerPorts(&v)
		hasInput = true
	}
	if !state.PrimaryIp4.IsNull() {
		v := state.PrimaryIp4.ValueString()
		params.SetPrimaryIp4([]string{v})
		hasInput = true
	}
	if !state.PrimaryIp4ID.IsNull() {
		v := state.PrimaryIp4ID.ValueInt64()
		params.SetPrimaryIp4ID([]int64{v})
		hasInput = true
	}
	if !state.PrimaryIp6.IsNull() {
		v := state.PrimaryIp6.ValueString()
		params.SetPrimaryIp6([]string{v})
		hasInput = true
	}
	if !state.PrimaryIp6ID.IsNull() {
		v := state.PrimaryIp6ID.ValueInt64()
		params.SetPrimaryIp6ID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RearPortCount.IsNull() {
		v := state.RearPortCount.ValueInt64()
		params.SetRearPortCount([]int64{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueString()
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.RoleSlug.IsNull() {
		v := state.RoleSlug.ValueString()
		params.SetRole([]string{v})
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
	if !state.TenantGroupSlug.IsNull() {
		v := state.TenantGroupSlug.ValueString()
		params.SetTenantGroup([]string{v})
		hasInput = true
	}
	if !state.TenantGroupID.IsNull() {
		v := state.TenantGroupID.ValueString()
		params.SetTenantGroupID([]string{v})
		hasInput = true
	}
	if !state.VcPosition.IsNull() {
		v := state.VcPosition.ValueInt64()
		params.SetVcPosition([]int64{v})
		hasInput = true
	}
	if !state.VcPriority.IsNull() {
		v := state.VcPriority.ValueInt64()
		params.SetVcPriority([]int64{v})
		hasInput = true
	}
	if !state.VirtualChassisID.IsNull() {
		v := state.VirtualChassisID.ValueInt64()
		params.SetVirtualChassisID([]int64{v})
		hasInput = true
	}
	if !state.VirtualChassisMember.IsNull() {
		v := state.VirtualChassisMember.ValueBool()
		params.SetVirtualChassisMember(&v)
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_device.")
		return
	}
	res, err := d.client.Dcim.DcimDevicesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_device",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenDevice(ctx, netboxapi.DeviceResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.deviceResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
