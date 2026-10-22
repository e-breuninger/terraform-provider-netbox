// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// webhookResourceModel is the Terraform state/plan model of the "webhook" resource.
type webhookResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	Name              types.String `tfsdk:"name"`
	PayloadURL        types.String `tfsdk:"payload_url"`
	HTTPMethod        types.String `tfsdk:"http_method"`
	HTTPContentType   types.String `tfsdk:"http_content_type"`
	AdditionalHeaders types.String `tfsdk:"additional_headers"`
	BodyTemplate      types.String `tfsdk:"body_template"`
	Secret            types.String `tfsdk:"secret"`
	SslVerification   types.Bool   `tfsdk:"ssl_verification"`
	CaFilePath        types.String `tfsdk:"ca_file_path"`
	Description       types.String `tfsdk:"description"`
	OwnerID           types.Int64  `tfsdk:"owner_id"`
	Created           types.String `tfsdk:"created"`
	LastUpdated       types.String `tfsdk:"last_updated"`
	URL               types.String `tfsdk:"url"`
	Tags              types.Set    `tfsdk:"tags"`
	TagsAll           types.Set    `tfsdk:"tags_all"`
	CustomFields      types.Map    `tfsdk:"custom_fields"`
}

// expandWebhook converts the Terraform model of type "webhook" into an API request.
func expandWebhook(ctx context.Context, model *webhookResourceModel) (*netboxapi.WebhookRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.WebhookRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.PayloadURL = conv.StringPtr(model.PayloadURL)
	requestDTO.HTTPMethod = conv.StringPtr(model.HTTPMethod)
	requestDTO.HTTPContentType = conv.StringPtr(model.HTTPContentType)
	requestDTO.AdditionalHeaders = conv.StringPtr(model.AdditionalHeaders)
	requestDTO.BodyTemplate = conv.StringPtr(model.BodyTemplate)
	requestDTO.Secret = conv.StringPtr(model.Secret)
	requestDTO.SslVerification = conv.BoolPtr(model.SslVerification)
	requestDTO.CaFilePath = conv.StringPtr(model.CaFilePath)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenWebhook refreshes the Terraform model of type "webhook" from an API response.
func flattenWebhook(ctx context.Context, responseDTO *netboxapi.WebhookResponseDTO, model *webhookResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.PayloadURL = conv.FromStringPtr(responseDTO.PayloadURL)
	model.HTTPMethod = conv.FromStringPtr(responseDTO.HTTPMethod)
	model.HTTPContentType = conv.FromStringPtr(responseDTO.HTTPContentType)
	model.AdditionalHeaders = conv.FromStringPtr(responseDTO.AdditionalHeaders)
	model.BodyTemplate = conv.FromStringPtr(responseDTO.BodyTemplate)
	model.Secret = conv.FromStringPtr(responseDTO.Secret)
	model.SslVerification = conv.FromBoolPtr(responseDTO.SslVerification)
	model.CaFilePath = conv.FromStringPtr(responseDTO.CaFilePath)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
