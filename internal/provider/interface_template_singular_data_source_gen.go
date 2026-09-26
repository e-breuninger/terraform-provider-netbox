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
	_ datasource.DataSource              = (*interfaceTemplateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*interfaceTemplateDataSource)(nil)
)

// NewInterfaceTemplateDataSource returns a new interface_template data source.
func NewInterfaceTemplateDataSource() datasource.DataSource {
	return &interfaceTemplateDataSource{}
}

type interfaceTemplateDataSource struct {
	client *netboxapi.Client
}

// interfaceTemplateDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type interfaceTemplateDataSourceModel struct {
	interfaceTemplateResourceModel
	NameContains types.String `tfsdk:"name_contains"`
}

func (d *interfaceTemplateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_template"
}

func (d *interfaceTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):An interface template of a device type or module type (dcim.interfacetemplate).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/interfacetemplate/):\n\n> A template for a network interface that will be created on all instantiations of the parent device type. See the interface documentation for more detail.",
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
				Description: "Interface type as its NetBox slug, e.g. virtual, lag, bridge, 1000base-t, 10gbase-x-sfpp (any of NetBox's interface type choices).",
			},
			"mgmt_only": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the interface is used for out-of-band management only.",
			},
			"poe_mode": schema.StringAttribute{
				Computed:    true,
				Description: "Power over Ethernet mode. One of: pd, pse.",
			},
			"poe_type": schema.StringAttribute{
				Computed:    true,
				Description: "Power over Ethernet type as its NetBox slug, e.g. type1-ieee802.3af, type3-ieee802.3bt, passive-24v-2pair (any of NetBox's PoE type choices).",
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

func (d *interfaceTemplateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *interfaceTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data interfaceTemplateDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimInterfaceTemplatesListParams()
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
	if !state.MgmtOnly.IsNull() {
		v := state.MgmtOnly.ValueBool()
		params.SetMgmtOnly(&v)
		hasInput = true
	}
	if !state.ModuleTypeID.IsNull() {
		v := state.ModuleTypeID.ValueInt64()
		params.SetModuleTypeID([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or name_contains and/or type and/or mgmt_only and/or module_type_id to look up a netbox_interface_template.")
		return
	}
	res, err := d.client.Dcim.DcimInterfaceTemplatesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_interface_template", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_interface_template",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenInterfaceTemplate(ctx, netboxapi.InterfaceTemplateResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.interfaceTemplateResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
