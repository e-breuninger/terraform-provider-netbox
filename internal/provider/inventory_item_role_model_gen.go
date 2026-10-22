// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// inventoryItemRoleResourceModel is the Terraform state/plan model of the "inventory_item_role" resource.
type inventoryItemRoleResourceModel struct {
	ID                 types.Int64  `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Slug               types.String `tfsdk:"slug"`
	ColorHex           types.String `tfsdk:"color_hex"`
	Description        types.String `tfsdk:"description"`
	Comments           types.String `tfsdk:"comments"`
	OwnerID            types.Int64  `tfsdk:"owner_id"`
	Created            types.String `tfsdk:"created"`
	LastUpdated        types.String `tfsdk:"last_updated"`
	URL                types.String `tfsdk:"url"`
	InventoryItemCount types.Int64  `tfsdk:"inventory_item_count"`
	Tags               types.Set    `tfsdk:"tags"`
	TagsAll            types.Set    `tfsdk:"tags_all"`
	CustomFields       types.Map    `tfsdk:"custom_fields"`
}

// expandInventoryItemRole converts the Terraform model of type "inventory_item_role" into an API request.
func expandInventoryItemRole(ctx context.Context, model *inventoryItemRoleResourceModel) (*netboxapi.InventoryItemRoleRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.InventoryItemRoleRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Color = conv.StringPtr(model.ColorHex)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenInventoryItemRole refreshes the Terraform model of type "inventory_item_role" from an API response.
func flattenInventoryItemRole(ctx context.Context, responseDTO *netboxapi.InventoryItemRoleResponseDTO, model *inventoryItemRoleResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.ColorHex = conv.FromStringPtr(responseDTO.Color)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.InventoryItemCount = conv.FromInt64Ptr(responseDTO.InventoryitemCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
