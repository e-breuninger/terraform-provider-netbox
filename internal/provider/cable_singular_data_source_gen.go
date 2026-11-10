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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*cableDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*cableDataSource)(nil)
)

// NewCableDataSource returns a new cable data source.
func NewCableDataSource() datasource.DataSource {
	return &cableDataSource{}
}

type cableDataSource struct {
	client *netboxapi.Client
}

// cableDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type cableDataSourceModel struct {
	cableResourceModel
	TagSlug              types.String `tfsdk:"tag_slug"`
	SiteID               types.Int64  `tfsdk:"site_id"`
	SiteSlug             types.String `tfsdk:"site_slug"`
	LocationID           types.Int64  `tfsdk:"location_id"`
	RackID               types.Int64  `tfsdk:"rack_id"`
	DeviceID             types.Int64  `tfsdk:"device_id"`
	DeviceName           types.String `tfsdk:"device_name"`
	BundleName           types.String `tfsdk:"bundle_name"`
	CircuitterminationID types.Int64  `tfsdk:"circuittermination_id"`
	Color                types.String `tfsdk:"color"`
	ConsoleportID        types.Int64  `tfsdk:"consoleport_id"`
	ConsoleserverportID  types.Int64  `tfsdk:"consoleserverport_id"`
	FrontportID          types.Int64  `tfsdk:"frontport_id"`
	InterfaceID          types.Int64  `tfsdk:"interface_id"`
	LocationSlug         types.String `tfsdk:"location_slug"`
	OwnerName            types.String `tfsdk:"owner_name"`
	OwnerGroupName       types.String `tfsdk:"owner_group_name"`
	OwnerGroupID         types.Int64  `tfsdk:"owner_group_id"`
	PowerfeedID          types.Int64  `tfsdk:"powerfeed_id"`
	PoweroutletID        types.Int64  `tfsdk:"poweroutlet_id"`
	PowerportID          types.Int64  `tfsdk:"powerport_id"`
	Q                    types.String `tfsdk:"q"`
	RackName             types.String `tfsdk:"rack_name"`
	RearportID           types.Int64  `tfsdk:"rearport_id"`
	TagID                types.Int64  `tfsdk:"tag_id"`
	TenantSlug           types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug      types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID        types.String `tfsdk:"tenant_group_id"`
	TerminationAID       types.Int64  `tfsdk:"termination_a_id"`
	TerminationAType     types.String `tfsdk:"termination_a_type"`
	TerminationBID       types.Int64  `tfsdk:"termination_b_id"`
	TerminationBType     types.String `tfsdk:"termination_b_type"`
	Unterminated         types.Bool   `tfsdk:"unterminated"`
	CustomFieldFilters   types.Map    `tfsdk:"custom_field_filters"`
}

func (d *cableDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cable"
}

func (d *cableDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A cable between two sets of terminations (dcim.cable): interfaces, front and rear ports, console and power ports, power feeds or circuit terminations.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/cable/):\n\n> All connections between device components in NetBox are represented using cables. A cable represents a direct physical connection between two sets of endpoints (A and B), such as a console port and a patch panel port, or between two network interfaces.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"a_side": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"object_type": schema.StringAttribute{
						Computed:    true,
						Description: "Content type of the terminating objects. Derived from the *_ids list that is set; set it together with ids otherwise. One of: dcim.interface, dcim.frontport, dcim.rearport, dcim.consoleport, dcim.consoleserverport, dcim.powerport, dcim.poweroutlet, dcim.powerfeed, circuits.circuittermination.",
					},
					"ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the terminating objects, in order (see object_type).",
					},
					"device_interface_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the device interfaces this side terminates on (object_type dcim.interface). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"front_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the front ports this side terminates on (object_type dcim.frontport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"rear_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the rear ports this side terminates on (object_type dcim.rearport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"console_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the console ports this side terminates on (object_type dcim.consoleport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"console_server_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the console server ports this side terminates on (object_type dcim.consoleserverport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"power_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the power ports this side terminates on (object_type dcim.powerport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"power_outlet_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the power outlets this side terminates on (object_type dcim.poweroutlet). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"power_feed_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the power feeds this side terminates on (object_type dcim.powerfeed). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"circuit_termination_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the circuit terminations this side terminates on (object_type circuits.circuittermination). Conflicts with the other *_ids lists and with object_type/ids.",
					},
				},
				Computed:    true,
				Description: "What the A side terminates on: objects of one type (NetBox's rule), one or several for a breakout, in order (with a profile the position is the connector). Give object_type and ids, or exactly one of the typed *_ids lists, which is an alias of the same thing.",
			},
			"b_side": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"object_type": schema.StringAttribute{
						Computed:    true,
						Description: "Content type of the terminating objects. Derived from the *_ids list that is set; set it together with ids otherwise. One of: dcim.interface, dcim.frontport, dcim.rearport, dcim.consoleport, dcim.consoleserverport, dcim.powerport, dcim.poweroutlet, dcim.powerfeed, circuits.circuittermination.",
					},
					"ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the terminating objects, in order (see object_type).",
					},
					"device_interface_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the device interfaces this side terminates on (object_type dcim.interface). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"front_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the front ports this side terminates on (object_type dcim.frontport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"rear_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the rear ports this side terminates on (object_type dcim.rearport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"console_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the console ports this side terminates on (object_type dcim.consoleport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"console_server_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the console server ports this side terminates on (object_type dcim.consoleserverport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"power_port_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the power ports this side terminates on (object_type dcim.powerport). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"power_outlet_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the power outlets this side terminates on (object_type dcim.poweroutlet). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"power_feed_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the power feeds this side terminates on (object_type dcim.powerfeed). Conflicts with the other *_ids lists and with object_type/ids.",
					},
					"circuit_termination_ids": schema.ListAttribute{
						ElementType: types.Int64Type,
						Computed:    true,
						Description: "Ids of the circuit terminations this side terminates on (object_type circuits.circuittermination). Conflicts with the other *_ids lists and with object_type/ids.",
					},
				},
				Computed:    true,
				Description: "What the B side terminates on: objects of one type (NetBox's rule), one or several for a breakout, in order (with a profile the position is the connector). Give object_type and ids, or exactly one of the typed *_ids lists, which is an alias of the same thing.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cable type as its NetBox slug, e.g. cat6, smf-os2, power (any of NetBox's cable type choices).",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cable status. One of: connected, planned, decommissioning.",
				Validators: []validator.String{
					stringvalidator.OneOf("connected", "planned", "decommissioning"),
				},
			},
			"profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cable profile (connectors and positions per side); with one set, each termination's position in the list is its connector. NetBox rejects clearing it, so unset keeps the current value. One of: single-1c1p, single-1c2p, single-1c4p, single-1c6p, single-1c8p, single-1c12p, single-1c16p, trunk-2c1p, trunk-2c2p, trunk-2c4p, trunk-2c4p-shuffle, trunk-2c6p, trunk-2c8p, trunk-2c12p, trunk-4c1p, trunk-4c2p, trunk-4c4p, trunk-4c4p-shuffle, trunk-4c6p, trunk-4c8p, trunk-8c4p, breakout-1c2p-2c1p, breakout-1c4p-4c1p, breakout-1c6p-6c1p, breakout-1c8p-8c1p, breakout-2c4p-8c1p-shuffle.",
				Validators: []validator.String{
					stringvalidator.OneOf("single-1c1p", "single-1c2p", "single-1c4p", "single-1c6p", "single-1c8p", "single-1c12p", "single-1c16p", "trunk-2c1p", "trunk-2c2p", "trunk-2c4p", "trunk-2c4p-shuffle", "trunk-2c6p", "trunk-2c8p", "trunk-2c12p", "trunk-4c1p", "trunk-4c2p", "trunk-4c4p", "trunk-4c4p-shuffle", "trunk-4c6p", "trunk-4c8p", "trunk-8c4p", "breakout-1c2p-2c1p", "breakout-1c4p-4c1p", "breakout-1c6p-6c1p", "breakout-1c8p-8c1p", "breakout-2c4p-8c1p-shuffle"),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"bundle_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the cable bundle.",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(100),
				},
			},
			"color_hex": schema.StringAttribute{
				Computed:    true,
				Description: "RGB color in hex (e.g. ff0000).",
			},
			"length": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Cable length; requires length_unit.",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"length_unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of length. One of: km, m, cm, mi, ft, in.",
				Validators: []validator.String{
					stringvalidator.OneOf("km", "m", "cm", "mi", "ft", "in"),
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
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site a termination of the cable is in.",
			},
			"site_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site a termination of the cable is in.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the location a termination of the cable is in.",
			},
			"rack_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the rack a termination of the cable is in.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a device the cable is connected to.",
			},
			"device_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of a device the cable is connected to.",
			},
			"bundle_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the cable bundle.",
			},
			"circuittermination_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the circuittermination.",
			},
			"color": schema.StringAttribute{
				Optional:    true,
				Description: "Color to match.",
			},
			"consoleport_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the consoleport.",
			},
			"consoleserverport_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the consoleserverport.",
			},
			"frontport_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the frontport.",
			},
			"interface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the interface.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
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
			"powerfeed_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the powerfeed.",
			},
			"poweroutlet_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the poweroutlet.",
			},
			"powerport_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the powerport.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"rack_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the rack.",
			},
			"rearport_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the rearport.",
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
			"termination_a_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the termination a.",
			},
			"termination_a_type": schema.StringAttribute{
				Optional:    true,
				Description: "Termination a type to match.",
			},
			"termination_b_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the termination b.",
			},
			"termination_b_type": schema.StringAttribute{
				Optional:    true,
				Description: "Termination b type to match.",
			},
			"unterminated": schema.BoolAttribute{
				Optional:    true,
				Description: "Unterminated.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *cableDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *cableDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data cableDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimCablesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Label.IsNull() {
		v := state.Label.ValueString()
		params.SetLabel([]string{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
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
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.SiteSlug.IsNull() {
		v := state.SiteSlug.ValueString()
		params.SetSite([]string{v})
		hasInput = true
	}
	if !state.LocationID.IsNull() {
		v := state.LocationID.ValueInt64()
		params.SetLocationID([]int64{v})
		hasInput = true
	}
	if !state.RackID.IsNull() {
		v := state.RackID.ValueInt64()
		params.SetRackID([]int64{v})
		hasInput = true
	}
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.DeviceName.IsNull() {
		v := state.DeviceName.ValueString()
		params.SetDevice([]string{v})
		hasInput = true
	}
	if !state.BundleName.IsNull() {
		v := state.BundleName.ValueString()
		params.SetBundle([]string{v})
		hasInput = true
	}
	if !state.BundleID.IsNull() {
		v := state.BundleID.ValueInt64()
		params.SetBundleID([]int64{v})
		hasInput = true
	}
	if !state.CircuitterminationID.IsNull() {
		v := state.CircuitterminationID.ValueInt64()
		params.SetCircuitterminationID([]int64{v})
		hasInput = true
	}
	if !state.Color.IsNull() {
		v := state.Color.ValueString()
		params.SetColor([]string{v})
		hasInput = true
	}
	if !state.ConsoleportID.IsNull() {
		v := state.ConsoleportID.ValueInt64()
		params.SetConsoleportID([]int64{v})
		hasInput = true
	}
	if !state.ConsoleserverportID.IsNull() {
		v := state.ConsoleserverportID.ValueInt64()
		params.SetConsoleserverportID([]int64{v})
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
	if !state.FrontportID.IsNull() {
		v := state.FrontportID.ValueInt64()
		params.SetFrontportID([]int64{v})
		hasInput = true
	}
	if !state.InterfaceID.IsNull() {
		v := state.InterfaceID.ValueInt64()
		params.SetInterfaceID([]int64{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.Length.IsNull() {
		v := state.Length.ValueFloat64()
		params.SetLength([]float64{v})
		hasInput = true
	}
	if !state.LengthUnit.IsNull() {
		v := state.LengthUnit.ValueString()
		params.SetLengthUnit(&v)
		hasInput = true
	}
	if !state.LocationSlug.IsNull() {
		v := state.LocationSlug.ValueString()
		params.SetLocation([]string{v})
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
	if !state.PowerfeedID.IsNull() {
		v := state.PowerfeedID.ValueInt64()
		params.SetPowerfeedID([]int64{v})
		hasInput = true
	}
	if !state.PoweroutletID.IsNull() {
		v := state.PoweroutletID.ValueInt64()
		params.SetPoweroutletID([]int64{v})
		hasInput = true
	}
	if !state.PowerportID.IsNull() {
		v := state.PowerportID.ValueInt64()
		params.SetPowerportID([]int64{v})
		hasInput = true
	}
	if !state.Profile.IsNull() {
		v := state.Profile.ValueString()
		params.SetProfile([]string{v})
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
	if !state.RearportID.IsNull() {
		v := state.RearportID.ValueInt64()
		params.SetRearportID([]int64{v})
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
	if !state.TerminationAID.IsNull() {
		v := state.TerminationAID.ValueInt64()
		params.SetTerminationaID([]int64{v})
		hasInput = true
	}
	if !state.TerminationAType.IsNull() {
		v := state.TerminationAType.ValueString()
		params.SetTerminationaType([]string{v})
		hasInput = true
	}
	if !state.TerminationBID.IsNull() {
		v := state.TerminationBID.ValueInt64()
		params.SetTerminationbID([]int64{v})
		hasInput = true
	}
	if !state.TerminationBType.IsNull() {
		v := state.TerminationBType.ValueString()
		params.SetTerminationbType([]string{v})
		hasInput = true
	}
	if !state.Unterminated.IsNull() {
		v := state.Unterminated.ValueBool()
		params.SetUnterminated(&v)
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_cable.")
		return
	}
	res, err := d.client.Dcim.DcimCablesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_cable", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_cable",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCable(ctx, netboxapi.CableResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.cableResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&cableResource{client: d.client}).postReadHook(ctx, &state.cableResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
