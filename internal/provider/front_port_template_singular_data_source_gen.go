// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

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
	_ datasource.DataSource              = (*frontPortTemplateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*frontPortTemplateDataSource)(nil)
)

// NewFrontPortTemplateDataSource returns a new front_port_template data source.
func NewFrontPortTemplateDataSource() datasource.DataSource {
	return &frontPortTemplateDataSource{}
}

type frontPortTemplateDataSource struct {
	client *netboxapi.Client
}

// frontPortTemplateDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type frontPortTemplateDataSourceModel struct {
	frontPortTemplateResourceModel
	NameContains types.String `tfsdk:"name_contains"`
}

func (d *frontPortTemplateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_front_port_template"
}

func (d *frontPortTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A front port template of a device type or module type (dcim.frontporttemplate), mapped position by position onto rear port templates of the same type.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/frontporttemplate/):\n\n> A template for a front-facing pass-through port that will be created on all instantiations of the parent device type. See the [front port](https://netboxlabs.com/docs/netbox/models/dcim/frontport/) documentation for more detail.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"device_type_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the device type. Set exactly one of device_type_id and module_type_id; NetBox rejects both and neither.",
			},
			"module_type_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the module type. Set exactly one of device_type_id and module_type_id; NetBox rejects both and neither.",
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
				Description: "Port type as its NetBox slug, e.g. 8p8c, lc, mpo (any of NetBox's port type choices).",
			},
			"positions": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of positions on the front port, at least the number of rear_ports mappings. NetBox checks a lower value against the mappings it already holds, so remove mappings in one apply and lower positions in the next.",
			},
			"rear_ports": schema.SetNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"position": schema.Int64Attribute{
							Computed:    true,
							Description: "Position on the front port, 1-based and unique per port.",
						},
						"rear_port_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the rear port template the position maps to.",
						},
						"rear_port_position": schema.Int64Attribute{
							Computed:    true,
							Description: "Position on that rear port, 1-based; each rear port position takes one mapping.",
						},
					},
				},
				Computed:    true,
				Description: "Mappings of the front port's positions onto rear port templates of the same device type or module type, one object per position. Derived from rear_port_id and rear_port_position when those are set; an empty set, or neither form, leaves the port unmapped. Changing the set recreates every mapping, since NetBox 4.6 rejects an update that resends an existing one.",
			},
			"rear_port_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Rear port of a single-position front port: the same as rear_ports = [{position = 1, rear_port_id = ..., rear_port_position = ...}]. Conflicts with rear_ports and with positions > 1.",
			},
			"rear_port_position": schema.Int64Attribute{
				Computed:    true,
				Description: "Position on rear_port_id. Requires rear_port_id.",
			},
			"color_hex": schema.StringAttribute{
				Computed:    true,
				Description: "RGB color in hex (e.g. ff0000).",
			},
			"label": schema.StringAttribute{
				Computed:    true,
				Description: "Physical label.",
			},
			"description": schema.StringAttribute{
				Computed: true,
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
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
			},
		},
	}
}

func (d *frontPortTemplateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *frontPortTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data frontPortTemplateDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimFrontPortTemplatesListParams()
	hasInput := false
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
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.ModuleTypeID.IsNull() {
		v := state.ModuleTypeID.ValueInt64()
		params.SetModuleTypeID([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or name_contains and/or type and/or module_type_id to look up a netbox_front_port_template.")
		return
	}
	res, err := d.client.Dcim.DcimFrontPortTemplatesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_front_port_template", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_front_port_template",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenFrontPortTemplate(ctx, netboxapi.FrontPortTemplateResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.frontPortTemplateResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&frontPortTemplateResource{client: d.client}).postRead(ctx, &state.frontPortTemplateResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
