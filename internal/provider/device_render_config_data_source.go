package provider

import (
	"context"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// netbox_device_render_config is a hand-written companion data source: it renders a device's
// configuration through NetBox's render-config endpoint, which has no object of its own for the
// spec to describe.

func init() {
	companionDataSources = append(companionDataSources, NewDeviceRenderConfigDataSource)
}

var (
	_ datasource.DataSource              = (*deviceRenderConfigDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceRenderConfigDataSource)(nil)
)

func NewDeviceRenderConfigDataSource() datasource.DataSource {
	return &deviceRenderConfigDataSource{}
}

type deviceRenderConfigDataSource struct {
	client *netboxapi.Client
}

type deviceRenderConfigDataSourceModel struct {
	DeviceID           types.Int64  `tfsdk:"device_id"`
	ConfigTemplateID   types.Int64  `tfsdk:"config_template_id"`
	ConfigTemplateName types.String `tfsdk:"config_template_name"`
	Content            types.String `tfsdk:"content"`
}

func (dataSource *deviceRenderConfigDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_render_config"
}

func (dataSource *deviceRenderConfigDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):The configuration NetBox renders for a device (dcim.device) from its config template, through the render-config endpoint. The template comes from the device, its role or its platform unless config_template_id names one.",
		Attributes: map[string]schema.Attribute{
			"device_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the device to render the configuration for.",
			},
			"config_template_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the config template to render with. Unset renders with the template NetBox resolves for the device, and reads back as that template's id.",
			},
			"config_template_name": schema.StringAttribute{
				Computed:    true,
				Description: "Name of the config template the configuration was rendered with.",
			},
			"content": schema.StringAttribute{
				Computed:    true,
				Description: "The rendered configuration.",
			},
		},
	}
}

func (dataSource *deviceRenderConfigDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected data source configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	dataSource.client = client
}

func (dataSource *deviceRenderConfigDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model deviceRenderConfigDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	deviceID := model.DeviceID.ValueInt64()
	params := dcim.NewDcimDevicesRenderConfigCreateParams().WithID(deviceID)
	if !model.ConfigTemplateID.IsNull() && !model.ConfigTemplateID.IsUnknown() {
		params = params.WithData(&models.RenderConfigInput{ConfigTemplateID: model.ConfigTemplateID.ValueInt64()})
	}
	res, err := dataSource.client.Dcim.DcimDevicesRenderConfigCreateContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error rendering the configuration of device %d", deviceID), err.Error())
		return
	}
	rendered := res.Payload
	if rendered == nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error rendering the configuration of device %d", deviceID), "NetBox returned no rendered configuration.")
		return
	}
	model.Content = types.StringValue(rendered.Content)
	model.ConfigTemplateID = types.Int64Null()
	model.ConfigTemplateName = types.StringNull()
	if rendered.Configtemplate != nil {
		model.ConfigTemplateID = types.Int64Value(rendered.Configtemplate.ID)
		if rendered.Configtemplate.Name != nil {
			model.ConfigTemplateName = types.StringValue(*rendered.Configtemplate.Name)
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
