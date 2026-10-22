// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// consolePortTemplateResourceModel is the Terraform state/plan model of the "console_port_template" resource.
type consolePortTemplateResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	DeviceTypeID types.Int64  `tfsdk:"device_type_id"`
	ModuleTypeID types.Int64  `tfsdk:"module_type_id"`
	Name         types.String `tfsdk:"name"`
	Type         types.String `tfsdk:"type"`
	Label        types.String `tfsdk:"label"`
	Description  types.String `tfsdk:"description"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
}

// expandConsolePortTemplate converts the Terraform model of type "console_port_template" into an API request.
func expandConsolePortTemplate(ctx context.Context, model *consolePortTemplateResourceModel) (*netboxapi.ConsolePortTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ConsolePortTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.ModuleType = conv.Int64Ptr(model.ModuleTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenConsolePortTemplate refreshes the Terraform model of type "console_port_template" from an API response.
func flattenConsolePortTemplate(ctx context.Context, responseDTO *netboxapi.ConsolePortTemplateResponseDTO, model *consolePortTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.ModuleTypeID = conv.FromInt64Ptr(responseDTO.ModuleType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
