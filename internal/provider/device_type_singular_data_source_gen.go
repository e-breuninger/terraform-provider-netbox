// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*deviceTypeDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceTypeDataSource)(nil)
)

// NewDeviceTypeDataSource returns a new device_type data source.
func NewDeviceTypeDataSource() datasource.DataSource {
	return &deviceTypeDataSource{}
}

type deviceTypeDataSource struct {
	client *netboxapi.Client
}

// deviceTypeDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type deviceTypeDataSourceModel struct {
	deviceTypeResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	ConsolePorts       types.Bool   `tfsdk:"console_ports"`
	ConsoleServerPorts types.Bool   `tfsdk:"console_server_ports"`
	DefaultPlatform    types.String `tfsdk:"default_platform"`
	DeviceBays         types.Bool   `tfsdk:"device_bays"`
	HasFrontImage      types.Bool   `tfsdk:"has_front_image"`
	HasRearImage       types.Bool   `tfsdk:"has_rear_image"`
	Interfaces         types.Bool   `tfsdk:"interfaces"`
	InventoryItems     types.Bool   `tfsdk:"inventory_items"`
	ManufacturerSlug   types.String `tfsdk:"manufacturer_slug"`
	ModuleBays         types.Bool   `tfsdk:"module_bays"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	PassThroughPorts   types.Bool   `tfsdk:"pass_through_ports"`
	PowerOutlets       types.Bool   `tfsdk:"power_outlets"`
	PowerPorts         types.Bool   `tfsdk:"power_ports"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *deviceTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_type"
}

func (d *deviceTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A device type (dcim.devicetype): a hardware model made by a manufacturer.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/devicetype/):\n\n> A device type represents a particular make and model of hardware that exists in the real world. Device types define the physical attributes of a device (rack height and depth) and its individual components (console, power, network interfaces, and so on).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"manufacturer_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the manufacturer.",
			},
			"model": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Model name (unique per manufacturer).",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from model when not set.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"part_number": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Discrete part number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"default_platform_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the platform assigned to new devices of this type.",
			},
			"u_height": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Height in rack units (0.5 steps).",
				Validators: []validator.Float64{
					float64validator.Between(0, 1000),
				},
			},
			"is_full_depth": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Device consumes both the front and rear rack faces.",
			},
			"exclude_from_utilization": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether devices of this type are left out of the rack utilization.",
			},
			"subdevice_role": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parent devices house child devices in device bays. One of: parent, child.",
				Validators: []validator.String{
					stringvalidator.OneOf("parent", "child"),
				},
			},
			"airflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Airflow direction. One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
				Validators: []validator.String{
					stringvalidator.OneOf("front-to-rear", "rear-to-front", "left-to-right", "right-to-left", "side-to-rear", "passive", "mixed", "rear-to-side", "bottom-to-top", "top-to-bottom"),
				},
			},
			"weight": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of the device type (requires weight_unit).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"weight_unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of weight. One of: kg, g, lb, oz.",
				Validators: []validator.String{
					stringvalidator.OneOf("kg", "g", "lb", "oz"),
				},
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
			"device_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of devices of the device type.",
			},
			"console_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of console port templates of the device type.",
			},
			"console_server_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of console server port templates of the device type.",
			},
			"power_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of power port templates of the device type.",
			},
			"power_outlet_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of power outlet templates of the device type.",
			},
			"interface_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of interface templates of the device type.",
			},
			"front_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of front port templates of the device type.",
			},
			"rear_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of rear port templates of the device type.",
			},
			"device_bay_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of device bay templates of the device type.",
			},
			"module_bay_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of module bay templates of the device type.",
			},
			"inventory_item_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of inventory item templates of the device type.",
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
			"console_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has console ports.",
			},
			"console_server_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has console server ports.",
			},
			"default_platform": schema.StringAttribute{
				Optional:    true,
				Description: "Default platform to match.",
			},
			"device_bays": schema.BoolAttribute{
				Optional:    true,
				Description: "Has device bays.",
			},
			"has_front_image": schema.BoolAttribute{
				Optional:    true,
				Description: "Has a front image.",
			},
			"has_rear_image": schema.BoolAttribute{
				Optional:    true,
				Description: "Has a rear image.",
			},
			"interfaces": schema.BoolAttribute{
				Optional:    true,
				Description: "Has interfaces.",
			},
			"inventory_items": schema.BoolAttribute{
				Optional:    true,
				Description: "Has inventory items.",
			},
			"manufacturer_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the manufacturer.",
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
			"pass_through_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has pass-through ports.",
			},
			"power_outlets": schema.BoolAttribute{
				Optional:    true,
				Description: "Has power outlets.",
			},
			"power_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has power ports.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *deviceTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceTypeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimDeviceTypesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Model.IsNull() {
		v := state.Model.ValueString()
		params.SetModel([]string{v})
		hasInput = true
	}
	if !state.Slug.IsNull() {
		v := state.Slug.ValueString()
		params.SetSlug([]string{v})
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
	if !state.ConsolePortTemplateCount.IsNull() {
		v := state.ConsolePortTemplateCount.ValueInt64()
		params.SetConsolePortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.ConsolePorts.IsNull() {
		v := state.ConsolePorts.ValueBool()
		params.SetConsolePorts(&v)
		hasInput = true
	}
	if !state.ConsoleServerPortTemplateCount.IsNull() {
		v := state.ConsoleServerPortTemplateCount.ValueInt64()
		params.SetConsoleServerPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.ConsoleServerPorts.IsNull() {
		v := state.ConsoleServerPorts.ValueBool()
		params.SetConsoleServerPorts(&v)
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.DefaultPlatform.IsNull() {
		v := state.DefaultPlatform.ValueString()
		params.SetDefaultPlatform([]string{v})
		hasInput = true
	}
	if !state.DefaultPlatformID.IsNull() {
		v := strconv.FormatInt(state.DefaultPlatformID.ValueInt64(), 10)
		params.SetDefaultPlatformID([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.DeviceBayTemplateCount.IsNull() {
		v := state.DeviceBayTemplateCount.ValueInt64()
		params.SetDeviceBayTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.DeviceBays.IsNull() {
		v := state.DeviceBays.ValueBool()
		params.SetDeviceBays(&v)
		hasInput = true
	}
	if !state.DeviceCount.IsNull() {
		v := state.DeviceCount.ValueInt64()
		params.SetDeviceCount([]int64{v})
		hasInput = true
	}
	if !state.ExcludeFromUtilization.IsNull() {
		v := state.ExcludeFromUtilization.ValueBool()
		params.SetExcludeFromUtilization(&v)
		hasInput = true
	}
	if !state.FrontPortTemplateCount.IsNull() {
		v := state.FrontPortTemplateCount.ValueInt64()
		params.SetFrontPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.HasFrontImage.IsNull() {
		v := state.HasFrontImage.ValueBool()
		params.SetHasFrontImage(&v)
		hasInput = true
	}
	if !state.HasRearImage.IsNull() {
		v := state.HasRearImage.ValueBool()
		params.SetHasRearImage(&v)
		hasInput = true
	}
	if !state.InterfaceTemplateCount.IsNull() {
		v := state.InterfaceTemplateCount.ValueInt64()
		params.SetInterfaceTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.Interfaces.IsNull() {
		v := state.Interfaces.ValueBool()
		params.SetInterfaces(&v)
		hasInput = true
	}
	if !state.InventoryItemTemplateCount.IsNull() {
		v := state.InventoryItemTemplateCount.ValueInt64()
		params.SetInventoryItemTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.InventoryItems.IsNull() {
		v := state.InventoryItems.ValueBool()
		params.SetInventoryItems(&v)
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
	if !state.ModuleBayTemplateCount.IsNull() {
		v := state.ModuleBayTemplateCount.ValueInt64()
		params.SetModuleBayTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.ModuleBays.IsNull() {
		v := state.ModuleBays.ValueBool()
		params.SetModuleBays(&v)
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
	if !state.PartNumber.IsNull() {
		v := state.PartNumber.ValueString()
		params.SetPartNumber([]string{v})
		hasInput = true
	}
	if !state.PassThroughPorts.IsNull() {
		v := state.PassThroughPorts.ValueBool()
		params.SetPassThroughPorts(&v)
		hasInput = true
	}
	if !state.PowerOutletTemplateCount.IsNull() {
		v := state.PowerOutletTemplateCount.ValueInt64()
		params.SetPowerOutletTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.PowerOutlets.IsNull() {
		v := state.PowerOutlets.ValueBool()
		params.SetPowerOutlets(&v)
		hasInput = true
	}
	if !state.PowerPortTemplateCount.IsNull() {
		v := state.PowerPortTemplateCount.ValueInt64()
		params.SetPowerPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.PowerPorts.IsNull() {
		v := state.PowerPorts.ValueBool()
		params.SetPowerPorts(&v)
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RearPortTemplateCount.IsNull() {
		v := state.RearPortTemplateCount.ValueInt64()
		params.SetRearPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.SubdeviceRole.IsNull() {
		v := state.SubdeviceRole.ValueString()
		params.SetSubdeviceRole(&v)
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.UHeight.IsNull() {
		v := state.UHeight.ValueFloat64()
		params.SetUHeight([]float64{v})
		hasInput = true
	}
	if !state.Weight.IsNull() {
		v := state.Weight.ValueFloat64()
		params.SetWeight([]float64{v})
		hasInput = true
	}
	if !state.WeightUnit.IsNull() {
		v := state.WeightUnit.ValueString()
		params.SetWeightUnit(&v)
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_device_type.")
		return
	}
	res, err := d.client.Dcim.DcimDeviceTypesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device_type", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_device_type",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenDeviceType(ctx, netboxapi.DeviceTypeResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.deviceTypeResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
