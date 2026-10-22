// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// virtualMachineTypeResourceModel is the Terraform state/plan model of the "virtual_machine_type" resource.
type virtualMachineTypeResourceModel struct {
	ID                  types.Int64   `tfsdk:"id"`
	Name                types.String  `tfsdk:"name"`
	Slug                types.String  `tfsdk:"slug"`
	DefaultPlatformID   types.Int64   `tfsdk:"default_platform_id"`
	DefaultVcpus        types.Float64 `tfsdk:"default_vcpus"`
	DefaultMemory       types.Int64   `tfsdk:"default_memory"`
	Description         types.String  `tfsdk:"description"`
	Comments            types.String  `tfsdk:"comments"`
	OwnerID             types.Int64   `tfsdk:"owner_id"`
	Created             types.String  `tfsdk:"created"`
	LastUpdated         types.String  `tfsdk:"last_updated"`
	URL                 types.String  `tfsdk:"url"`
	VirtualMachineCount types.Int64   `tfsdk:"virtual_machine_count"`
	Tags                types.Set     `tfsdk:"tags"`
	TagsAll             types.Set     `tfsdk:"tags_all"`
	CustomFields        types.Map     `tfsdk:"custom_fields"`
}

// expandVirtualMachineType converts the Terraform model of type "virtual_machine_type" into an API request.
func expandVirtualMachineType(ctx context.Context, model *virtualMachineTypeResourceModel) (*netboxapi.VirtualMachineTypeRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualMachineTypeRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.DefaultPlatform = conv.Int64Ptr(model.DefaultPlatformID)
	requestDTO.DefaultVcpus = conv.Float64Ptr(model.DefaultVcpus)
	requestDTO.DefaultMemory = conv.Int64Ptr(model.DefaultMemory)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualMachineType refreshes the Terraform model of type "virtual_machine_type" from an API response.
func flattenVirtualMachineType(ctx context.Context, responseDTO *netboxapi.VirtualMachineTypeResponseDTO, model *virtualMachineTypeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.DefaultPlatformID = conv.FromInt64Ptr(responseDTO.DefaultPlatform)
	model.DefaultVcpus = conv.FromFloat64Ptr(responseDTO.DefaultVcpus)
	model.DefaultMemory = conv.FromInt64Ptr(responseDTO.DefaultMemory)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.VirtualMachineCount = conv.FromInt64Ptr(responseDTO.VirtualMachineCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
