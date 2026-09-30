// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// deviceRoleResourceModel is the Terraform state/plan model of the "device_role" resource.
type deviceRoleResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Slug                types.String `tfsdk:"slug"`
	Description         types.String `tfsdk:"description"`
	Comments            types.String `tfsdk:"comments"`
	ColorHex            types.String `tfsdk:"color_hex"`
	VmRole              types.Bool   `tfsdk:"vm_role"`
	ParentID            types.Int64  `tfsdk:"parent_id"`
	ConfigTemplateID    types.Int64  `tfsdk:"config_template_id"`
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

// expandDeviceRole converts the Terraform model of type "device_role" into an API request.
func expandDeviceRole(ctx context.Context, model *deviceRoleResourceModel) (*netboxapi.DeviceRoleRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceRoleRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Color = conv.StringPtr(model.ColorHex)
	requestDTO.VMRole = conv.BoolPtr(model.VmRole)
	requestDTO.Parent = conv.Int64Ptr(model.ParentID)
	requestDTO.ConfigTemplate = conv.Int64Ptr(model.ConfigTemplateID)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDeviceRole refreshes the Terraform model of type "device_role" from an API response.
func flattenDeviceRole(ctx context.Context, responseDTO *netboxapi.DeviceRoleResponseDTO, model *deviceRoleResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.ColorHex = conv.FromStringPtr(responseDTO.Color)
	model.VmRole = conv.FromBoolPtr(responseDTO.VMRole)
	model.ParentID = conv.FromInt64Ptr(responseDTO.Parent)
	model.ConfigTemplateID = conv.FromInt64Ptr(responseDTO.ConfigTemplate)
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
