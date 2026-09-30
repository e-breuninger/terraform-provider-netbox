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
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
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
				Computed:    true,
				Description: "Cable profile (connectors and positions per side); with one set, each termination's position in the list is its connector. NetBox rejects clearing it, so unset keeps the current value. One of: single-1c1p, single-1c2p, single-1c4p, single-1c6p, single-1c8p, single-1c12p, single-1c16p, trunk-2c1p, trunk-2c2p, trunk-2c4p, trunk-2c4p-shuffle, trunk-2c6p, trunk-2c8p, trunk-2c12p, trunk-4c1p, trunk-4c2p, trunk-4c4p, trunk-4c4p-shuffle, trunk-4c6p, trunk-4c8p, trunk-8c4p, breakout-1c2p-2c1p, breakout-1c4p-4c1p, breakout-1c6p-6c1p, breakout-1c8p-8c1p, breakout-2c4p-8c1p-shuffle.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"bundle_id": schema.Int64Attribute{
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
				Computed:    true,
				Description: "Cable length; requires length_unit.",
			},
			"length_unit": schema.StringAttribute{
				Computed:    true,
				Description: "Unit of length. One of: km, m, cm, mi, ft, in.",
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or label and/or status and/or type and/or tenant_id and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_cable.")
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
