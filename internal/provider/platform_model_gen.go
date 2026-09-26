// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// platformResourceModel is the Terraform state/plan model of the "platform" resource.
type platformResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Slug                types.String `tfsdk:"slug"`
	Description         types.String `tfsdk:"description"`
	ManufacturerID      types.Int64  `tfsdk:"manufacturer_id"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	Created             types.String `tfsdk:"created"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	URL                 types.String `tfsdk:"url"`
	DeviceCount         types.Int64  `tfsdk:"device_count"`
	VirtualMachineCount types.Int64  `tfsdk:"virtual_machine_count"`
	Tags                types.Set    `tfsdk:"tags"`
	TagsAll             types.Set    `tfsdk:"tags_all"`
	CustomFields        types.Map    `tfsdk:"custom_fields"`
}

// expandPlatform converts the Terraform model of type "platform" into an API request.
func expandPlatform(ctx context.Context, model *platformResourceModel) (*netboxapi.PlatformRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.PlatformRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Manufacturer = conv.Int64Ptr(model.ManufacturerID)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenPlatform refreshes the Terraform model of type "platform" from an API response.
func flattenPlatform(ctx context.Context, responseDTO *netboxapi.PlatformResponseDTO, model *platformResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.ManufacturerID = conv.FromInt64Ptr(responseDTO.Manufacturer)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.DeviceCount = conv.FromInt64Ptr(responseDTO.DeviceCount)
	model.VirtualMachineCount = conv.FromInt64Ptr(responseDTO.VirtualmachineCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
