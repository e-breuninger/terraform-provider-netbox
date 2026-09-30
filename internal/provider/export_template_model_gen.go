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

// exportTemplateResourceModel is the Terraform state/plan model of the "export_template" resource.
type exportTemplateResourceModel struct {
	ID                types.Int64          `tfsdk:"id"`
	Name              types.String         `tfsdk:"name"`
	ObjectTypes       types.Set            `tfsdk:"object_types"`
	TemplateCode      types.String         `tfsdk:"template_code"`
	MimeType          types.String         `tfsdk:"mime_type"`
	FileName          types.String         `tfsdk:"file_name"`
	FileExtension     types.String         `tfsdk:"file_extension"`
	AsAttachment      types.Bool           `tfsdk:"as_attachment"`
	EnvironmentParams jsontypes.Normalized `tfsdk:"environment_params"`
	Description       types.String         `tfsdk:"description"`
	OwnerID           types.Int64          `tfsdk:"owner_id"`
	Created           types.String         `tfsdk:"created"`
	LastUpdated       types.String         `tfsdk:"last_updated"`
	URL               types.String         `tfsdk:"url"`
}

// expandExportTemplate converts the Terraform model of type "export_template" into an API request.
func expandExportTemplate(ctx context.Context, model *exportTemplateResourceModel) (*netboxapi.ExportTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ExportTemplateRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.ObjectTypes = conv.SetTo[string](ctx, model.ObjectTypes, &diags)
	requestDTO.TemplateCode = conv.StringPtr(model.TemplateCode)
	requestDTO.MimeType = conv.StringPtr(model.MimeType)
	requestDTO.FileName = conv.StringPtr(model.FileName)
	requestDTO.FileExtension = conv.StringPtr(model.FileExtension)
	requestDTO.AsAttachment = conv.BoolPtr(model.AsAttachment)
	requestDTO.EnvironmentParams = conv.StringPtr(model.EnvironmentParams.StringValue)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	return requestDTO, diags
}

// flattenExportTemplate refreshes the Terraform model of type "export_template" from an API response.
func flattenExportTemplate(ctx context.Context, responseDTO *netboxapi.ExportTemplateResponseDTO, model *exportTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.ObjectTypes = conv.SetFrom(ctx, types.StringType, responseDTO.ObjectTypes, false, &diags)
	model.TemplateCode = conv.FromStringPtr(responseDTO.TemplateCode)
	model.MimeType = conv.FromStringPtr(responseDTO.MimeType)
	model.FileName = conv.FromStringPtr(responseDTO.FileName)
	model.FileExtension = conv.FromStringPtr(responseDTO.FileExtension)
	model.AsAttachment = conv.FromBoolPtr(responseDTO.AsAttachment)
	model.EnvironmentParams = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.EnvironmentParams))
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
