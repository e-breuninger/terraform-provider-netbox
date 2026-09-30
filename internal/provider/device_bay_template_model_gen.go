// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// deviceBayTemplateResourceModel is the Terraform state/plan model of the "device_bay_template" resource.
type deviceBayTemplateResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	DeviceTypeID types.Int64  `tfsdk:"device_type_id"`
	Name         types.String `tfsdk:"name"`
	Label        types.String `tfsdk:"label"`
	Enabled      types.Bool   `tfsdk:"enabled"`
	Description  types.String `tfsdk:"description"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
}

// expandDeviceBayTemplate converts the Terraform model of type "device_bay_template" into an API request.
func expandDeviceBayTemplate(ctx context.Context, model *deviceBayTemplateResourceModel) (*netboxapi.DeviceBayTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceBayTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenDeviceBayTemplate refreshes the Terraform model of type "device_bay_template" from an API response.
func flattenDeviceBayTemplate(ctx context.Context, responseDTO *netboxapi.DeviceBayTemplateResponseDTO, model *deviceBayTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
