package provider

import (
	"context"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*webhookDataSource)(nil)

func NewWebhookDataSource() datasource.DataSource {
	return &webhookDataSource{}
}

type webhookDataSource struct {
	NetboxDataSource
}

func (d *webhookDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

func (d *webhookDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A webhook is a mechanism for conveying to some external system a change that took place in NetBox. For example, you may want to notify a monitoring system whenever the status of a device is updated in NetBox. This can be done by creating a webhook for the device model in NetBox and identifying the webhook receiver. When NetBox detects a change to a device, an HTTP request containing the details of the change and who made it be sent to the specified receiver.",
		Attributes: map[string]schema.Attribute{
			"additional_headers": schema.StringAttribute{
				Computed:            true,
				Description:         "User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).",
				MarkdownDescription: "User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).",
			},
			"body_template": schema.StringAttribute{
				Computed:            true,
				Description:         "Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.",
				MarkdownDescription: "Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.",
			},
			"ca_file_path": schema.StringAttribute{
				Computed:            true,
				Description:         "The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.",
				MarkdownDescription: "The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.",
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"http_content_type": schema.StringAttribute{
				Computed:            true,
				Description:         "The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.",
				MarkdownDescription: "The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.",
			},
			"http_method": schema.StringAttribute{
				Computed:            true,
				Description:         "* `GET` - GET\n* `POST` - POST\n* `PUT` - PUT\n* `PATCH` - PATCH\n* `DELETE` - DELETE",
				MarkdownDescription: "* `GET` - GET\n* `POST` - POST\n* `PUT` - PUT\n* `PATCH` - PATCH\n* `DELETE` - DELETE",
			},
			"id": schema.Int32Attribute{
				Required:            true,
				Description:         "A unique integer value identifying this webhook.",
				MarkdownDescription: "A unique integer value identifying this webhook.",
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"secret": schema.StringAttribute{
				Computed: true,
			},
			"payload_url": schema.StringAttribute{
				Computed:            true,
				Description:         "This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body.",
				MarkdownDescription: "This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body.",
			},
			"ssl_verification": schema.BoolAttribute{
				Computed:            true,
				Description:         "Enable SSL certificate verification. Disable with caution!",
				MarkdownDescription: "Enable SSL certificate verification. Disable with caution!",
			},
			"tags": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
		},
	}
}

func (d *webhookDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.WebhookTerraformModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(d.provider.client)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	webhook, httpCode, err := d.provider.client.ExtrasAPI.ExtrasWebhooksRetrieve(ctx, data.Id.ValueInt32()).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to retrieve Webhook value.",
				err.Error(),
			)
			return
		}
	}
	errors := data.ReadAPI(ctx, webhook)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
