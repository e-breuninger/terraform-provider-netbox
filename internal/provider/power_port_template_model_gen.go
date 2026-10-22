// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// powerPortTemplateResourceModel is the Terraform state/plan model of the "power_port_template" resource.
type powerPortTemplateResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	DeviceTypeID  types.Int64  `tfsdk:"device_type_id"`
	ModuleTypeID  types.Int64  `tfsdk:"module_type_id"`
	Name          types.String `tfsdk:"name"`
	Type          types.String `tfsdk:"type"`
	MaximumDraw   types.Int64  `tfsdk:"maximum_draw"`
	AllocatedDraw types.Int64  `tfsdk:"allocated_draw"`
	Label         types.String `tfsdk:"label"`
	Description   types.String `tfsdk:"description"`
	Created       types.String `tfsdk:"created"`
	LastUpdated   types.String `tfsdk:"last_updated"`
	URL           types.String `tfsdk:"url"`
}

// expandPowerPortTemplate converts the Terraform model of type "power_port_template" into an API request.
func expandPowerPortTemplate(ctx context.Context, model *powerPortTemplateResourceModel) (*netboxapi.PowerPortTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.PowerPortTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.ModuleType = conv.Int64Ptr(model.ModuleTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.MaximumDraw = conv.Int64Ptr(model.MaximumDraw)
	requestDTO.AllocatedDraw = conv.Int64Ptr(model.AllocatedDraw)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenPowerPortTemplate refreshes the Terraform model of type "power_port_template" from an API response.
func flattenPowerPortTemplate(ctx context.Context, responseDTO *netboxapi.PowerPortTemplateResponseDTO, model *powerPortTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.ModuleTypeID = conv.FromInt64Ptr(responseDTO.ModuleType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.MaximumDraw = conv.FromInt64Ptr(responseDTO.MaximumDraw)
	model.AllocatedDraw = conv.FromInt64Ptr(responseDTO.AllocatedDraw)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
