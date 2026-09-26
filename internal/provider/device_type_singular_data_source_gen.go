// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

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
	ModelContains      types.String `tfsdk:"model_contains"`
	TagSlug            types.String `tfsdk:"tag_slug"`
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
				Computed:    true,
				Description: "Discrete part number.",
			},
			"u_height": schema.Float64Attribute{
				Computed:    true,
				Description: "Height in rack units (0.5 steps).",
			},
			"is_full_depth": schema.BoolAttribute{
				Computed:    true,
				Description: "Device consumes both the front and rear rack faces.",
			},
			"subdevice_role": schema.StringAttribute{
				Computed:    true,
				Description: "Parent devices house child devices in device bays. One of: parent, child.",
			},
			"airflow": schema.StringAttribute{
				Computed:    true,
				Description: "Airflow direction. One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
			},
			"weight": schema.Float64Attribute{
				Computed:    true,
				Description: "Weight of the device type (requires weight_unit).",
			},
			"weight_unit": schema.StringAttribute{
				Computed:    true,
				Description: "Unit of weight. One of: kg, g, lb, oz.",
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
			"device_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of devices of the device type.",
			},
			"console_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of console port templates of the device type.",
			},
			"console_server_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of console server port templates of the device type.",
			},
			"power_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of power port templates of the device type.",
			},
			"power_outlet_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of power outlet templates of the device type.",
			},
			"interface_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of interface templates of the device type.",
			},
			"front_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of front port templates of the device type.",
			},
			"rear_port_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of rear port templates of the device type.",
			},
			"device_bay_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of device bay templates of the device type.",
			},
			"module_bay_template_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of module bay templates of the device type.",
			},
			"inventory_item_template_count": schema.Int64Attribute{
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
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"model_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the model.",
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
	if !state.ModelContains.IsNull() {
		v := state.ModelContains.ValueString()
		params.SetModelIc([]string{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or model and/or slug and/or model_contains and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_device_type.")
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
