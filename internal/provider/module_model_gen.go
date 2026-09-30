// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// moduleResourceModel is the Terraform state/plan model of the "module" resource.
type moduleResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	DeviceID            types.Int64  `tfsdk:"device_id"`
	ModuleBayID         types.Int64  `tfsdk:"module_bay_id"`
	ModuleTypeID        types.Int64  `tfsdk:"module_type_id"`
	Status              types.String `tfsdk:"status"`
	Serial              types.String `tfsdk:"serial"`
	AssetTag            types.String `tfsdk:"asset_tag"`
	ReplicateComponents types.Bool   `tfsdk:"replicate_components"`
	AdoptComponents     types.Bool   `tfsdk:"adopt_components"`
	Description         types.String `tfsdk:"description"`
	Comments            types.String `tfsdk:"comments"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	Created             types.String `tfsdk:"created"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	URL                 types.String `tfsdk:"url"`
	Tags                types.Set    `tfsdk:"tags"`
	TagsAll             types.Set    `tfsdk:"tags_all"`
	CustomFields        types.Map    `tfsdk:"custom_fields"`
}

// expandModule converts the Terraform model of type "module" into an API request.
func expandModule(ctx context.Context, model *moduleResourceModel) (*netboxapi.ModuleRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ModuleRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.ModuleBay = conv.Int64Ptr(model.ModuleBayID)
	requestDTO.ModuleType = conv.Int64Ptr(model.ModuleTypeID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Serial = conv.StringPtr(model.Serial)
	requestDTO.AssetTag = conv.StringPtr(model.AssetTag)
	requestDTO.ReplicateComponents = conv.BoolPtr(model.ReplicateComponents)
	requestDTO.AdoptComponents = conv.BoolPtr(model.AdoptComponents)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenModule refreshes the Terraform model of type "module" from an API response.
func flattenModule(ctx context.Context, responseDTO *netboxapi.ModuleResponseDTO, model *moduleResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.ModuleBayID = conv.FromInt64Ptr(responseDTO.ModuleBay)
	model.ModuleTypeID = conv.FromInt64Ptr(responseDTO.ModuleType)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Serial = conv.FromStringPtr(responseDTO.Serial)
	model.AssetTag = conv.FromStringPtr(responseDTO.AssetTag)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
