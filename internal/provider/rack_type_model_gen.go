// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rackTypeResourceModel is the Terraform state/plan model of the "rack_type" resource.
type rackTypeResourceModel struct {
	ID              types.Int64   `tfsdk:"id"`
	ManufacturerID  types.Int64   `tfsdk:"manufacturer_id"`
	Model           types.String  `tfsdk:"model"`
	Slug            types.String  `tfsdk:"slug"`
	Description     types.String  `tfsdk:"description"`
	FormFactor      types.String  `tfsdk:"form_factor"`
	Width           types.Int64   `tfsdk:"width"`
	UHeight         types.Int64   `tfsdk:"u_height"`
	StartingUnit    types.Int64   `tfsdk:"starting_unit"`
	DescUnits       types.Bool    `tfsdk:"desc_units"`
	OuterWidth      types.Int64   `tfsdk:"outer_width"`
	OuterHeight     types.Int64   `tfsdk:"outer_height"`
	OuterDepth      types.Int64   `tfsdk:"outer_depth"`
	OuterUnit       types.String  `tfsdk:"outer_unit"`
	MountingDepth   types.Int64   `tfsdk:"mounting_depth"`
	MountingDepthMm types.Int64   `tfsdk:"mounting_depth_mm"`
	Weight          types.Float64 `tfsdk:"weight"`
	MaxWeight       types.Int64   `tfsdk:"max_weight"`
	WeightUnit      types.String  `tfsdk:"weight_unit"`
	Comments        types.String  `tfsdk:"comments"`
	OwnerID         types.Int64   `tfsdk:"owner_id"`
	Created         types.String  `tfsdk:"created"`
	LastUpdated     types.String  `tfsdk:"last_updated"`
	URL             types.String  `tfsdk:"url"`
	RackCount       types.Int64   `tfsdk:"rack_count"`
	Tags            types.Set     `tfsdk:"tags"`
	TagsAll         types.Set     `tfsdk:"tags_all"`
	CustomFields    types.Map     `tfsdk:"custom_fields"`
}

// expandRackType converts the Terraform model of type "rack_type" into an API request.
func expandRackType(ctx context.Context, model *rackTypeResourceModel) (*netboxapi.RackTypeRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.RackTypeRequestDTO{}
	requestDTO.Manufacturer = conv.Int64Ptr(model.ManufacturerID)
	requestDTO.Model = conv.StringPtr(model.Model)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.FormFactor = conv.StringPtr(model.FormFactor)
	requestDTO.Width = conv.Int64Ptr(model.Width)
	requestDTO.UHeight = conv.Int64Ptr(model.UHeight)
	requestDTO.StartingUnit = conv.Int64Ptr(model.StartingUnit)
	requestDTO.DescUnits = conv.BoolPtr(model.DescUnits)
	requestDTO.OuterWidth = conv.Int64Ptr(model.OuterWidth)
	requestDTO.OuterHeight = conv.Int64Ptr(model.OuterHeight)
	requestDTO.OuterDepth = conv.Int64Ptr(model.OuterDepth)
	requestDTO.OuterUnit = conv.StringPtr(model.OuterUnit)
	requestDTO.MountingDepth = conv.Int64Ptr(model.MountingDepth)
	requestDTO.Weight = conv.Float64Ptr(model.Weight)
	requestDTO.MaxWeight = conv.Int64Ptr(model.MaxWeight)
	requestDTO.WeightUnit = conv.StringPtr(model.WeightUnit)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenRackType refreshes the Terraform model of type "rack_type" from an API response.
func flattenRackType(ctx context.Context, responseDTO *netboxapi.RackTypeResponseDTO, model *rackTypeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.ManufacturerID = conv.FromInt64Ptr(responseDTO.Manufacturer)
	model.Model = conv.FromStringPtr(responseDTO.Model)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.FormFactor = conv.FromStringPtr(responseDTO.FormFactor)
	model.Width = conv.FromInt64Ptr(responseDTO.Width)
	model.UHeight = conv.FromInt64Ptr(responseDTO.UHeight)
	model.StartingUnit = conv.FromInt64Ptr(responseDTO.StartingUnit)
	model.DescUnits = conv.FromBoolPtr(responseDTO.DescUnits)
	model.OuterWidth = conv.FromInt64Ptr(responseDTO.OuterWidth)
	model.OuterHeight = conv.FromInt64Ptr(responseDTO.OuterHeight)
	model.OuterDepth = conv.FromInt64Ptr(responseDTO.OuterDepth)
	model.OuterUnit = conv.FromStringPtr(responseDTO.OuterUnit)
	model.MountingDepth = conv.FromInt64Ptr(responseDTO.MountingDepth)
	model.Weight = conv.FromFloat64Ptr(responseDTO.Weight)
	model.MaxWeight = conv.FromInt64Ptr(responseDTO.MaxWeight)
	model.WeightUnit = conv.FromStringPtr(responseDTO.WeightUnit)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.RackCount = conv.FromInt64Ptr(responseDTO.RackCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	model.MountingDepthMm = model.MountingDepth
	return diags
}
