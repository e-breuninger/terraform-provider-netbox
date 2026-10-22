// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// inventoryItemResourceModel is the Terraform state/plan model of the "inventory_item" resource.
type inventoryItemResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	DeviceID       types.Int64  `tfsdk:"device_id"`
	Name           types.String `tfsdk:"name"`
	ParentID       types.Int64  `tfsdk:"parent_id"`
	RoleID         types.Int64  `tfsdk:"role_id"`
	Status         types.String `tfsdk:"status"`
	ManufacturerID types.Int64  `tfsdk:"manufacturer_id"`
	PartID         types.String `tfsdk:"part_id"`
	Serial         types.String `tfsdk:"serial"`
	AssetTag       types.String `tfsdk:"asset_tag"`
	Discovered     types.Bool   `tfsdk:"discovered"`
	ComponentType  types.String `tfsdk:"component_type"`
	ComponentID    types.Int64  `tfsdk:"component_id"`
	Label          types.String `tfsdk:"label"`
	Description    types.String `tfsdk:"description"`
	OwnerID        types.Int64  `tfsdk:"owner_id"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
	Tags           types.Set    `tfsdk:"tags"`
	TagsAll        types.Set    `tfsdk:"tags_all"`
	CustomFields   types.Map    `tfsdk:"custom_fields"`
}

// expandInventoryItem converts the Terraform model of type "inventory_item" into an API request.
func expandInventoryItem(ctx context.Context, model *inventoryItemResourceModel) (*netboxapi.InventoryItemRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.InventoryItemRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Parent = conv.Int64Ptr(model.ParentID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Manufacturer = conv.Int64Ptr(model.ManufacturerID)
	requestDTO.PartID = conv.StringPtr(model.PartID)
	requestDTO.Serial = conv.StringPtr(model.Serial)
	requestDTO.AssetTag = conv.StringPtr(model.AssetTag)
	requestDTO.Discovered = conv.BoolPtr(model.Discovered)
	requestDTO.ComponentType = conv.StringPtr(model.ComponentType)
	requestDTO.ComponentID = conv.Int64Ptr(model.ComponentID)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenInventoryItem refreshes the Terraform model of type "inventory_item" from an API response.
func flattenInventoryItem(ctx context.Context, responseDTO *netboxapi.InventoryItemResponseDTO, model *inventoryItemResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.ParentID = conv.FromInt64Ptr(responseDTO.Parent)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.ManufacturerID = conv.FromInt64Ptr(responseDTO.Manufacturer)
	model.PartID = conv.FromStringPtr(responseDTO.PartID)
	model.Serial = conv.FromStringPtr(responseDTO.Serial)
	model.AssetTag = conv.FromStringPtr(responseDTO.AssetTag)
	model.Discovered = conv.FromBoolPtr(responseDTO.Discovered)
	model.ComponentType = conv.FromStringPtr(responseDTO.ComponentType)
	model.ComponentID = conv.FromInt64Ptr(responseDTO.ComponentID)
	model.Label = conv.FromStringPtr(responseDTO.Label)
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
