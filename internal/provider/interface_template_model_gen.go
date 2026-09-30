// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// interfaceTemplateResourceModel is the Terraform state/plan model of the "interface_template" resource.
type interfaceTemplateResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	DeviceTypeID types.Int64  `tfsdk:"device_type_id"`
	ModuleTypeID types.Int64  `tfsdk:"module_type_id"`
	Name         types.String `tfsdk:"name"`
	Type         types.String `tfsdk:"type"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	MgmtOnly     types.Bool   `tfsdk:"mgmt_only"`
	BridgeID     types.Int64  `tfsdk:"bridge_id"`
	PoeMode      types.String `tfsdk:"poe_mode"`
	PoeType      types.String `tfsdk:"poe_type"`
	RfRole       types.String `tfsdk:"rf_role"`
	Label        types.String `tfsdk:"label"`
	Description  types.String `tfsdk:"description"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
}

// expandInterfaceTemplate converts the Terraform model of type "interface_template" into an API request.
func expandInterfaceTemplate(ctx context.Context, model *interfaceTemplateResourceModel) (*netboxapi.InterfaceTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.InterfaceTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.ModuleType = conv.Int64Ptr(model.ModuleTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.MgmtOnly = conv.BoolPtr(model.MgmtOnly)
	requestDTO.Bridge = conv.Int64Ptr(model.BridgeID)
	requestDTO.PoeMode = conv.StringPtr(model.PoeMode)
	requestDTO.PoeType = conv.StringPtr(model.PoeType)
	requestDTO.RfRole = conv.StringPtr(model.RfRole)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenInterfaceTemplate refreshes the Terraform model of type "interface_template" from an API response.
func flattenInterfaceTemplate(ctx context.Context, responseDTO *netboxapi.InterfaceTemplateResponseDTO, model *interfaceTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.ModuleTypeID = conv.FromInt64Ptr(responseDTO.ModuleType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.MgmtOnly = conv.FromBoolPtr(responseDTO.MgmtOnly)
	model.BridgeID = conv.FromInt64Ptr(responseDTO.Bridge)
	model.PoeMode = conv.FromStringPtr(responseDTO.PoeMode)
	model.PoeType = conv.FromStringPtr(responseDTO.PoeType)
	model.RfRole = conv.FromStringPtr(responseDTO.RfRole)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
