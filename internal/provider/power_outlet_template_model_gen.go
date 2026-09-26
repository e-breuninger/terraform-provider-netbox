// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// powerOutletTemplateResourceModel is the Terraform state/plan model of the "power_outlet_template" resource.
type powerOutletTemplateResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	DeviceTypeID        types.Int64  `tfsdk:"device_type_id"`
	ModuleTypeID        types.Int64  `tfsdk:"module_type_id"`
	Name                types.String `tfsdk:"name"`
	Type                types.String `tfsdk:"type"`
	PowerPortTemplateID types.Int64  `tfsdk:"power_port_template_id"`
	PowerPortID         types.Int64  `tfsdk:"power_port_id"`
	FeedLeg             types.String `tfsdk:"feed_leg"`
	Label               types.String `tfsdk:"label"`
	Description         types.String `tfsdk:"description"`
	Created             types.String `tfsdk:"created"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	URL                 types.String `tfsdk:"url"`
}

// expandPowerOutletTemplate converts the Terraform model of type "power_outlet_template" into an API request.
func expandPowerOutletTemplate(ctx context.Context, model *powerOutletTemplateResourceModel) (*netboxapi.PowerOutletTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.PowerOutletTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.ModuleType = conv.Int64Ptr(model.ModuleTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.PowerPort = conv.Int64Ptr(model.PowerPortTemplateID)
	requestDTO.FeedLeg = conv.StringPtr(model.FeedLeg)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenPowerOutletTemplate refreshes the Terraform model of type "power_outlet_template" from an API response.
func flattenPowerOutletTemplate(ctx context.Context, responseDTO *netboxapi.PowerOutletTemplateResponseDTO, model *powerOutletTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.ModuleTypeID = conv.FromInt64Ptr(responseDTO.ModuleType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.PowerPortTemplateID = conv.FromInt64Ptr(responseDTO.PowerPort)
	model.FeedLeg = conv.FromStringPtr(responseDTO.FeedLeg)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.PowerPortID = model.PowerPortTemplateID
	return diags
}
