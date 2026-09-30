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
	NameContains       types.String `tfsdk:"name_contains"`
	RegionSlug         types.String `tfsdk:"region_slug"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
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
				Computed:    true,
				Description: "GPS coordinate in decimal format (xx.yyyyyy).",
			},
			"longitude": schema.Float64Attribute{
				Computed:    true,
				Description: "GPS coordinate in decimal format (xx.yyyyyy).",
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
				Computed:    true,
				Description: "One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
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
				Computed:    true,
				Description: "Id of the virtual chassis.",
			},
			"virtual_chassis_position": schema.Int64Attribute{
				Computed:    true,
				Description: "Position in the virtual chassis (0-255).",
			},
			"virtual_chassis_priority": schema.Int64Attribute{
				Computed:    true,
				Description: "Master election priority in the virtual chassis (0-255).",
			},
			"config_template_id": schema.Int64Attribute{
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
			"description": schema.StringAttribute{
				Computed: true,
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
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region of the site.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
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
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or name_contains and/or serial and/or asset_tag and/or status and/or site_id and/or role_id and/or device_type_id and/or tenant_id and/or rack_id and/or location_id and/or cluster_id and/or platform_id and/or region_slug and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_device.")
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
