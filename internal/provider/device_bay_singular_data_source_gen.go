// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*deviceBayDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceBayDataSource)(nil)
)

// NewDeviceBayDataSource returns a new device_bay data source.
func NewDeviceBayDataSource() datasource.DataSource {
	return &deviceBayDataSource{}
}

type deviceBayDataSource struct {
	client *netboxapi.Client
}

// deviceBayDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type deviceBayDataSourceModel struct {
	deviceBayResourceModel
	TagSlug             types.String `tfsdk:"tag_slug"`
	DeviceName          types.String `tfsdk:"device_name"`
	DeviceRoleSlug      types.String `tfsdk:"device_role_slug"`
	DeviceRoleID        types.Int64  `tfsdk:"device_role_id"`
	DeviceStatus        types.String `tfsdk:"device_status"`
	DeviceTypeModel     types.String `tfsdk:"device_type_model"`
	DeviceTypeID        types.Int64  `tfsdk:"device_type_id"`
	InstalledDeviceName types.String `tfsdk:"installed_device_name"`
	LocationSlug        types.String `tfsdk:"location_slug"`
	LocationID          types.Int64  `tfsdk:"location_id"`
	OwnerName           types.String `tfsdk:"owner_name"`
	OwnerGroupName      types.String `tfsdk:"owner_group_name"`
	OwnerGroupID        types.Int64  `tfsdk:"owner_group_id"`
	Q                   types.String `tfsdk:"q"`
	RackName            types.String `tfsdk:"rack_name"`
	RackID              types.Int64  `tfsdk:"rack_id"`
	RegionSlug          types.String `tfsdk:"region_slug"`
	RegionID            types.String `tfsdk:"region_id"`
	SiteSlug            types.String `tfsdk:"site_slug"`
	SiteGroupSlug       types.String `tfsdk:"site_group_slug"`
	SiteGroupID         types.String `tfsdk:"site_group_id"`
	SiteID              types.Int64  `tfsdk:"site_id"`
	TagID               types.Int64  `tfsdk:"tag_id"`
	TenantSlug          types.String `tfsdk:"tenant_slug"`
	TenantID            types.Int64  `tfsdk:"tenant_id"`
	VirtualChassis      types.String `tfsdk:"virtual_chassis"`
	VirtualChassisID    types.Int64  `tfsdk:"virtual_chassis_id"`
	CustomFieldFilters  types.Map    `tfsdk:"custom_field_filters"`
}

func (d *deviceBayDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_bay"
}

func (d *deviceBayDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A device bay (dcim.devicebay).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/devicebay/):\n\n> Device bays represent a space or slot within a device in which a field-replaceable device may be installed. A common example is that of a chassis-based server that holds a number of blades which may contain device components that don't fit the module pattern. Devices in turn hold additional components that become available to the parent device.",
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
				Description: "Whether the device bay is enabled.",
			},
			"installed_device_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device installed in this bay.",
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
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
			"installed_device_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the installed device.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the location.",
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
			"virtual_chassis": schema.StringAttribute{
				Optional:    true,
				Description: "Virtual Chassis.",
			},
			"virtual_chassis_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual chassis.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *deviceBayDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceBayDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceBayDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimDeviceBaysListParams()
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
	if !state.Label.IsNull() {
		v := state.Label.ValueString()
		params.SetLabel([]string{v})
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
	if !state.Enabled.IsNull() {
		v := state.Enabled.ValueBool()
		params.SetEnabled(&v)
		hasInput = true
	}
	if !state.InstalledDeviceName.IsNull() {
		v := state.InstalledDeviceName.ValueString()
		params.SetInstalledDevice([]string{v})
		hasInput = true
	}
	if !state.InstalledDeviceID.IsNull() {
		v := state.InstalledDeviceID.ValueInt64()
		params.SetInstalledDeviceID([]int64{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_device_bay.")
		return
	}
	res, err := d.client.Dcim.DcimDeviceBaysListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device_bay", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_device_bay",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenDeviceBay(ctx, netboxapi.DeviceBayResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.deviceBayResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
