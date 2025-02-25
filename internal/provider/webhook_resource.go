package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/booldefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int32planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
	"strconv"
)

type WebhookModel struct {
	AdditionalHeaders types.String  `tfsdk:"additional_headers"`
	BodyTemplate      types.String  `tfsdk:"body_template"`
	CaFilePath        types.String  `tfsdk:"ca_file_path"`
	CustomFields      types.Dynamic `tfsdk:"custom_fields"`
	Description       types.String  `tfsdk:"description"`
	HttpContentType   types.String  `tfsdk:"http_content_type"`
	HttpMethod        types.String  `tfsdk:"http_method"`
	Id                types.Int32   `tfsdk:"id"`
	Name              types.String  `tfsdk:"name"`
	PayloadUrl        types.String  `tfsdk:"payload_url"`
	Secret            types.String  `tfsdk:"secret"`
	SslVerification   types.Bool    `tfsdk:"ssl_verification"`
	Tags              types.List    `tfsdk:"tags"`
}

var _ resource.Resource = (*webhookResource)(nil)

func NewWebhookResource() resource.Resource {
	return &webhookResource{}
}

type webhookResource struct {
	provider *netboxProvider
}

func (r *webhookResource) readAPI(ctx context.Context, data *WebhookModel, webhook *netbox.Webhook) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	data.Id = types.Int32Value(webhook.Id)
	data.Name = types.StringValue(webhook.Name)
	data.PayloadUrl = types.StringValue(webhook.PayloadUrl)
	data.BodyTemplate = types.StringPointerValue(webhook.BodyTemplate)
	data.HttpMethod = types.StringValue(string(*webhook.HttpMethod))
	data.HttpContentType = types.StringPointerValue(webhook.HttpContentType)
	data.AdditionalHeaders = types.StringPointerValue(webhook.AdditionalHeaders)
	data.SslVerification = types.BoolPointerValue(webhook.SslVerification)
	if *webhook.Secret != "" {
		data.Secret = types.StringPointerValue(webhook.Secret)
	}

	data.Description = types.StringPointerValue(webhook.Description)
	if webhook.CaFilePath.IsSet() {
		data.CaFilePath = types.StringPointerValue(webhook.CaFilePath.Get())
	}

	tags := readTags(webhook.Tags)
	tagsdata, diagdata := types.ListValueFrom(ctx, types.StringType, tags)
	if diagdata.HasError() {
		diags.AddError(
			"Error while reading Tags",
			"") //TODO Better handling
		return diags
	}
	data.Tags = tagsdata

	customFieldResults, diags := readCustomFieldsFromAPI(webhook.CustomFields)
	if diags.HasError() {
		return diags
	}
	data.CustomFields = types.DynamicValue(customFieldResults)
	return nil
}

func (r *webhookResource) writeAPI(ctx context.Context, data *WebhookModel) *netbox.WebhookRequest {
	webhookRequest := netbox.NewWebhookRequestWithDefaults()
	webhookRequest.Name = data.Name.ValueString()
	webhookRequest.PayloadUrl = data.PayloadUrl.ValueString()
	webhookRequest.BodyTemplate = data.BodyTemplate.ValueStringPointer()

	httpMethod := netbox.PatchedWebhookRequestHttpMethod(data.HttpMethod.ValueString())
	webhookRequest.HttpMethod = &httpMethod

	webhookRequest.HttpContentType = data.HttpContentType.ValueStringPointer()
	webhookRequest.AdditionalHeaders = data.AdditionalHeaders.ValueStringPointer()
	webhookRequest.SslVerification = data.SslVerification.ValueBoolPointer()
	webhookRequest.Secret = data.Secret.ValueStringPointer()
	webhookRequest.Description = data.Description.ValueStringPointer()
	caFilePath := netbox.NullableString{}
	caFilePath.Set(data.CaFilePath.ValueStringPointer())
	webhookRequest.CaFilePath = caFilePath
	tag_list := []netbox.NestedTagRequest{}

	for _, element := range data.Tags.Elements() {
		tag := netbox.NewNestedTagRequestWithDefaults()
		tag.Name = element.String()
		tag_list = append(tag_list, *tag)
	}
	webhookRequest.Tags = tag_list
	return webhookRequest
}

func (r *webhookResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(*netboxProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *netbox.apiCLient, got: %T, Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	r.provider = provider
}

func (r *webhookResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhook"
}

func (r *webhookResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"additional_headers": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).",
				MarkdownDescription: "User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).",
			},
			"body_template": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.",
				MarkdownDescription: "Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.",
			},
			"ca_file_path": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.",
				MarkdownDescription: "The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(4096),
				},
			},
			"custom_fields": schema.DynamicAttribute{
				Optional: true,
				Computed: true,
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"http_content_type": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.",
				MarkdownDescription: "The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
				Default: stringdefault.StaticString("application/json"),
			},
			"http_method": schema.StringAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "* `GET` - GET\n* `POST` - POST\n* `PUT` - PUT\n* `PATCH` - PATCH\n* `DELETE` - DELETE",
				MarkdownDescription: "* `GET` - GET\n* `POST` - POST\n* `PUT` - PUT\n* `PATCH` - PATCH\n* `DELETE` - DELETE",
				Validators: []validator.String{
					stringvalidator.OneOf(
						"GET",
						"POST",
						"PUT",
						"PATCH",
						"DELETE",
					),
				},
				Default: stringdefault.StaticString("POST"),
			},
			"id": schema.Int32Attribute{
				Computed:            true,
				Description:         "A unique integer value identifying this webhook.",
				MarkdownDescription: "A unique integer value identifying this webhook.",
				PlanModifiers: []planmodifier.Int32{
					int32planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 150),
				},
			},
			"payload_url": schema.StringAttribute{
				Required:            true,
				Description:         "This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body.",
				MarkdownDescription: "This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 500),
				},
			},
			"secret": schema.StringAttribute{
				Optional:            true,
				Description:         "When provided, the request will include a <code>X-Hook-Signature</code> header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.",
				MarkdownDescription: "When provided, the request will include a <code>X-Hook-Signature</code> header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
				},
				Sensitive: true,
			},
			"ssl_verification": schema.BoolAttribute{
				Optional:            true,
				Computed:            true,
				Description:         "Enable SSL certificate verification. Disable with caution!",
				MarkdownDescription: "Enable SSL certificate verification. Disable with caution!",
				Default:             booldefault.StaticBool(true),
			},
			"tags": schema.ListAttribute{
				ElementType: types.StringType,
				Optional:    true,
			},
		},
	}
}
func (r *webhookResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.Atoi(req.ID)
	if err != nil {
		resp.Diagnostics.AddError(
			"Unable to import Webhook. This method requires a integer.",
			err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}

func (r *webhookResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data WebhookModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Create API call logic
	webhookRequest := r.writeAPI(ctx, &data)

	api_res, _, err := r.provider.client.ExtrasAPI.
		ExtrasWebhooksCreate(ctx).
		WebhookRequest(*webhookRequest).
		Execute()

	if err != nil {
		resp.Diagnostics.AddError(
			"Error while creating the webhook.",
			err.Error(),
		)
		return
	}

	// Read API call logic
	errors := r.readAPI(ctx, &data, api_res)
	if errors != nil {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *webhookResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data WebhookModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}
	webhook, httpCode, err := r.provider.client.ExtrasAPI.ExtrasWebhooksRetrieve(ctx, data.Id.ValueInt32()).Execute()

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

	// Read API call logic
	errors := r.readAPI(ctx, &data, webhook)
	if errors != nil {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *webhookResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data WebhookModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(r.provider.client)...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Update API call logic
	webhookRequest := r.writeAPI(ctx, &data)

	webhook, httpCode, err := r.provider.client.ExtrasAPI.ExtrasWebhooksUpdate(ctx, data.Id.ValueInt32()).WebhookRequest(*webhookRequest).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.Diagnostics.AddError(
				"Webhook no longer exists",
				err.Error())
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to update Webhook.",
				err.Error(),
			)
			return
		}
	}

	// Read API call logic
	errors := r.readAPI(ctx, &data, webhook)
	if errors != nil {
		resp.Diagnostics.Append(errors...)
		return
	}
	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

func (r *webhookResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data WebhookModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	//Validate that the APIClient exist.
	if r.provider.client == nil {
		resp.Diagnostics.AddError(
			"Create: Unconfigured API Client",
			"Expected configured API Client. Please report this issue to the provider developers.",
		)
		return
	}
	// Delete API call logic
	httpCode, err := r.provider.client.ExtrasAPI.ExtrasWebhooksDestroy(ctx, data.Id.ValueInt32()).Execute()
	if err != nil {
		if httpCode != nil && httpCode.StatusCode != 404 {
			resp.Diagnostics.AddError(
				"Unable to update Webhook.",
				err.Error(),
			)
			return
		}
	}
}
