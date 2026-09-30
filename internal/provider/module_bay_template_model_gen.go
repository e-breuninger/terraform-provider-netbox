// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// moduleBayTemplateResourceModel is the Terraform state/plan model of the "module_bay_template" resource.
type moduleBayTemplateResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	DeviceTypeID types.Int64  `tfsdk:"device_type_id"`
	ModuleTypeID types.Int64  `tfsdk:"module_type_id"`
	Name         types.String `tfsdk:"name"`
	Position     types.String `tfsdk:"position"`
	Label        types.String `tfsdk:"label"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	Description  types.String `tfsdk:"description"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
}

// expandModuleBayTemplate converts the Terraform model of type "module_bay_template" into an API request.
func expandModuleBayTemplate(ctx context.Context, model *moduleBayTemplateResourceModel) (*netboxapi.ModuleBayTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ModuleBayTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.ModuleType = conv.Int64Ptr(model.ModuleTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Position = conv.StringPtr(model.Position)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenModuleBayTemplate refreshes the Terraform model of type "module_bay_template" from an API response.
func flattenModuleBayTemplate(ctx context.Context, responseDTO *netboxapi.ModuleBayTemplateResponseDTO, model *moduleBayTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.ModuleTypeID = conv.FromInt64Ptr(responseDTO.ModuleType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Position = conv.FromStringPtr(responseDTO.Position)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
