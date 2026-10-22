// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// virtualDiskResourceModel is the Terraform state/plan model of the "virtual_disk" resource.
type virtualDiskResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	VirtualMachineID types.Int64  `tfsdk:"virtual_machine_id"`
	Name             types.String `tfsdk:"name"`
	SizeMb           types.Int64  `tfsdk:"size_mb"`
	Description      types.String `tfsdk:"description"`
	OwnerID          types.Int64  `tfsdk:"owner_id"`
	Created          types.String `tfsdk:"created"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	URL              types.String `tfsdk:"url"`
	Tags             types.Set    `tfsdk:"tags"`
	TagsAll          types.Set    `tfsdk:"tags_all"`
	CustomFields     types.Map    `tfsdk:"custom_fields"`
}

// expandVirtualDisk converts the Terraform model of type "virtual_disk" into an API request.
func expandVirtualDisk(ctx context.Context, model *virtualDiskResourceModel) (*netboxapi.VirtualDiskRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualDiskRequestDTO{}
	requestDTO.VirtualMachine = conv.Int64Ptr(model.VirtualMachineID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Size = conv.Int64Ptr(model.SizeMb)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualDisk refreshes the Terraform model of type "virtual_disk" from an API response.
func flattenVirtualDisk(ctx context.Context, responseDTO *netboxapi.VirtualDiskResponseDTO, model *virtualDiskResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.VirtualMachineID = conv.FromInt64Ptr(responseDTO.VirtualMachine)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.SizeMb = conv.FromInt64Ptr(responseDTO.Size)
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
