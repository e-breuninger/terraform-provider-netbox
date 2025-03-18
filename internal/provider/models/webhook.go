package models

import (
	"context"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider/helpers"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
)

type WebhookTerraformModel struct {
	AdditionalHeaders types.String `tfsdk:"additional_headers"`
	BodyTemplate      types.String `tfsdk:"body_template"`
	CaFilePath        types.String `tfsdk:"ca_file_path"`
	CustomFields      types.Map    `tfsdk:"custom_fields"`
	Description       types.String `tfsdk:"description"`
	HttpContentType   types.String `tfsdk:"http_content_type"`
	HttpMethod        types.String `tfsdk:"http_method"`
	Id                types.Int32  `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	PayloadUrl        types.String `tfsdk:"payload_url"`
	Secret            types.String `tfsdk:"secret"`
	SslVerification   types.Bool   `tfsdk:"ssl_verification"`
	Tags              types.List   `tfsdk:"tags"`
}

func (data *WebhookTerraformModel) ReadAPI(ctx context.Context, webhook *netbox.Webhook) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	data.Id = types.Int32Value(webhook.Id)
	data.Name = types.StringValue(webhook.Name)
	data.PayloadUrl = types.StringValue(webhook.PayloadUrl)
	data.BodyTemplate = types.StringPointerValue(webhook.BodyTemplate)
	data.HttpMethod = types.StringValue(string(*webhook.HttpMethod))
	data.HttpContentType = types.StringPointerValue(webhook.HttpContentType)
	data.AdditionalHeaders = types.StringPointerValue(webhook.AdditionalHeaders)
	data.SslVerification = types.BoolPointerValue(webhook.SslVerification)

	data.Description = types.StringPointerValue(webhook.Description)
	if webhook.CaFilePath.IsSet() {
		data.CaFilePath = types.StringPointerValue(webhook.CaFilePath.Get())
	}

	tags := helpers.ReadTagsFromAPI(webhook.Tags)
	tagsdata, diagdata := types.ListValueFrom(ctx, types.StringType, tags)
	if diagdata.HasError() {
		diags.AddError(
			"Error while reading Webhook",
			"") //TODO Better handling
		return diags
	}
	data.Tags = tagsdata

	customFields, diagData := types.MapValueFrom(ctx, types.StringType, helpers.ReadCustomFieldsFromAPI(webhook.CustomFields))
	if diagData.HasError() {
		diags.Append()
	}

	data.CustomFields = customFields
	return nil
}
