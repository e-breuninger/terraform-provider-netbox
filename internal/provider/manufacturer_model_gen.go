// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// manufacturerResourceModel is the Terraform state/plan model of the "manufacturer" resource.
type manufacturerResourceModel struct {
	ID                 types.Int64  `tfsdk:"id"`
	Name               types.String `tfsdk:"name"`
	Slug               types.String `tfsdk:"slug"`
	Description        types.String `tfsdk:"description"`
	OwnerID            types.Int64  `tfsdk:"owner_id"`
	Created            types.String `tfsdk:"created"`
	LastUpdated        types.String `tfsdk:"last_updated"`
	URL                types.String `tfsdk:"url"`
	DeviceTypeCount    types.Int64  `tfsdk:"device_type_count"`
	ModuleTypeCount    types.Int64  `tfsdk:"module_type_count"`
	InventoryItemCount types.Int64  `tfsdk:"inventory_item_count"`
	PlatformCount      types.Int64  `tfsdk:"platform_count"`
	Tags               types.Set    `tfsdk:"tags"`
	TagsAll            types.Set    `tfsdk:"tags_all"`
	CustomFields       types.Map    `tfsdk:"custom_fields"`
}

// expandManufacturer converts the Terraform model of type "manufacturer" into an API request.
func expandManufacturer(ctx context.Context, model *manufacturerResourceModel) (*netboxapi.ManufacturerRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ManufacturerRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenManufacturer refreshes the Terraform model of type "manufacturer" from an API response.
func flattenManufacturer(ctx context.Context, responseDTO *netboxapi.ManufacturerResponseDTO, model *manufacturerResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.DeviceTypeCount = conv.FromInt64Ptr(responseDTO.DevicetypeCount)
	model.ModuleTypeCount = conv.FromInt64Ptr(responseDTO.ModuletypeCount)
	model.InventoryItemCount = conv.FromInt64Ptr(responseDTO.InventoryitemCount)
	model.PlatformCount = conv.FromInt64Ptr(responseDTO.PlatformCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
