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
	_ datasource.DataSource              = (*powerOutletTemplateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*powerOutletTemplateDataSource)(nil)
)

// NewPowerOutletTemplateDataSource returns a new power_outlet_template data source.
func NewPowerOutletTemplateDataSource() datasource.DataSource {
	return &powerOutletTemplateDataSource{}
}

type powerOutletTemplateDataSource struct {
	client *netboxapi.Client
}

// powerOutletTemplateDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type powerOutletTemplateDataSourceModel struct {
	powerOutletTemplateResourceModel
	Color types.String `tfsdk:"color"`
	Q     types.String `tfsdk:"q"`
}

func (d *powerOutletTemplateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_power_outlet_template"
}

func (d *powerOutletTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A power outlet template of a device type or module type (dcim.poweroutlettemplate).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/poweroutlettemplate/):\n\n> A template for a power outlet that will be created on all instantiations of the parent device type. See the power outlet documentation for more detail.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"device_type_id": schema.Int64Attribute{
				Optional:    true,
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
				Description: "Outlet type as its NetBox slug, e.g. iec-60320-c13, nema-5-15r, CS8364C (any of NetBox's power outlet type choices).",
			},
			"power_port_template_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the power port template this outlet draws from.",
			},
			"power_port_id": schema.Int64Attribute{
				Computed:           true,
				DeprecationMessage: "Use power_port_template_id instead.",
				Description:        "Deprecated alias of power_port_template_id.",
			},
			"feed_leg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Phase of the upstream feed this outlet is mapped to. One of: A, B, C.",
				Validators: []validator.String{
					stringvalidator.OneOf("A", "B", "C"),
				},
			},
			"color_hex": schema.StringAttribute{
				Computed:    true,
				Description: "RGB color in hex (e.g. ff0000).",
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
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
			"color": schema.StringAttribute{
				Optional:    true,
				Description: "Color to match.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
		},
	}
}

func (d *powerOutletTemplateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *powerOutletTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data powerOutletTemplateDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimPowerOutletTemplatesListParams()
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
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType(&v)
		hasInput = true
	}
	if !state.FeedLeg.IsNull() {
		v := state.FeedLeg.ValueString()
		params.SetFeedLeg([]string{v})
		hasInput = true
	}
	if !state.ModuleTypeID.IsNull() {
		v := state.ModuleTypeID.ValueInt64()
		params.SetModuleTypeID([]int64{v})
		hasInput = true
	}
	if !state.Color.IsNull() {
		v := state.Color.ValueString()
		params.SetColor([]string{v})
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
	if !state.DeviceTypeID.IsNull() {
		v := state.DeviceTypeID.ValueInt64()
		params.SetDeviceTypeID([]int64{v})
		hasInput = true
	}
	if !state.Label.IsNull() {
		v := state.Label.ValueString()
		params.SetLabel([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.PowerPortTemplateID.IsNull() {
		v := state.PowerPortTemplateID.ValueInt64()
		params.SetPowerPortID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_power_outlet_template.")
		return
	}
	res, err := d.client.Dcim.DcimPowerOutletTemplatesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_power_outlet_template", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_power_outlet_template",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenPowerOutletTemplate(ctx, netboxapi.PowerOutletTemplateResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.powerOutletTemplateResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
