// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// configTemplateResourceModel is the Terraform state/plan model of the "config_template" resource.
type configTemplateResourceModel struct {
	ID                types.Int64          `tfsdk:"id"`
	Name              types.String         `tfsdk:"name"`
	Description       types.String         `tfsdk:"description"`
	TemplateCode      types.String         `tfsdk:"template_code"`
	EnvironmentParams jsontypes.Normalized `tfsdk:"environment_params"`
	MimeType          types.String         `tfsdk:"mime_type"`
	FileName          types.String         `tfsdk:"file_name"`
	FileExtension     types.String         `tfsdk:"file_extension"`
	AsAttachment      types.Bool           `tfsdk:"as_attachment"`
	Debug             types.Bool           `tfsdk:"debug"`
	OwnerID           types.Int64          `tfsdk:"owner_id"`
	Created           types.String         `tfsdk:"created"`
	LastUpdated       types.String         `tfsdk:"last_updated"`
	URL               types.String         `tfsdk:"url"`
	Tags              types.Set            `tfsdk:"tags"`
	TagsAll           types.Set            `tfsdk:"tags_all"`
}

// expandConfigTemplate converts the Terraform model of type "config_template" into an API request.
func expandConfigTemplate(ctx context.Context, model *configTemplateResourceModel) (*netboxapi.ConfigTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ConfigTemplateRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.TemplateCode = conv.StringPtr(model.TemplateCode)
	requestDTO.EnvironmentParams = conv.StringPtr(model.EnvironmentParams.StringValue)
	requestDTO.MimeType = conv.StringPtr(model.MimeType)
	requestDTO.FileName = conv.StringPtr(model.FileName)
	requestDTO.FileExtension = conv.StringPtr(model.FileExtension)
	requestDTO.AsAttachment = conv.BoolPtr(model.AsAttachment)
	requestDTO.Debug = conv.BoolPtr(model.Debug)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	return requestDTO, diags
}

// flattenConfigTemplate refreshes the Terraform model of type "config_template" from an API response.
func flattenConfigTemplate(ctx context.Context, responseDTO *netboxapi.ConfigTemplateResponseDTO, model *configTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.TemplateCode = conv.FromStringPtr(responseDTO.TemplateCode)
	model.EnvironmentParams = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.EnvironmentParams))
	model.MimeType = conv.FromStringPtr(responseDTO.MimeType)
	model.FileName = conv.FromStringPtr(responseDTO.FileName)
	model.FileExtension = conv.FromStringPtr(responseDTO.FileExtension)
	model.AsAttachment = conv.FromBoolPtr(responseDTO.AsAttachment)
	model.Debug = conv.FromBoolPtr(responseDTO.Debug)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	return diags
}
