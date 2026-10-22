// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// inventoryItemTemplateResourceModel is the Terraform state/plan model of the "inventory_item_template" resource.
type inventoryItemTemplateResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	DeviceTypeID   types.Int64  `tfsdk:"device_type_id"`
	Name           types.String `tfsdk:"name"`
	ParentID       types.Int64  `tfsdk:"parent_id"`
	RoleID         types.Int64  `tfsdk:"role_id"`
	ManufacturerID types.Int64  `tfsdk:"manufacturer_id"`
	PartID         types.String `tfsdk:"part_id"`
	Label          types.String `tfsdk:"label"`
	ComponentType  types.String `tfsdk:"component_type"`
	ComponentID    types.Int64  `tfsdk:"component_id"`
	Description    types.String `tfsdk:"description"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
}

// expandInventoryItemTemplate converts the Terraform model of type "inventory_item_template" into an API request.
func expandInventoryItemTemplate(ctx context.Context, model *inventoryItemTemplateResourceModel) (*netboxapi.InventoryItemTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.InventoryItemTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Parent = conv.Int64Ptr(model.ParentID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Manufacturer = conv.Int64Ptr(model.ManufacturerID)
	requestDTO.PartID = conv.StringPtr(model.PartID)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.ComponentType = conv.StringPtr(model.ComponentType)
	requestDTO.ComponentID = conv.Int64Ptr(model.ComponentID)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenInventoryItemTemplate refreshes the Terraform model of type "inventory_item_template" from an API response.
func flattenInventoryItemTemplate(ctx context.Context, responseDTO *netboxapi.InventoryItemTemplateResponseDTO, model *inventoryItemTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.ParentID = conv.FromInt64Ptr(responseDTO.Parent)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.ManufacturerID = conv.FromInt64Ptr(responseDTO.Manufacturer)
	model.PartID = conv.FromStringPtr(responseDTO.PartID)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.ComponentType = conv.FromStringPtr(responseDTO.ComponentType)
	model.ComponentID = conv.FromInt64Ptr(responseDTO.ComponentID)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
