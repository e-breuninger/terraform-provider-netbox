// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// virtualDeviceContextResourceModel is the Terraform state/plan model of the "virtual_device_context" resource.
type virtualDeviceContextResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	DeviceID       types.Int64  `tfsdk:"device_id"`
	Name           types.String `tfsdk:"name"`
	Identifier     types.Int64  `tfsdk:"identifier"`
	Status         types.String `tfsdk:"status"`
	TenantID       types.Int64  `tfsdk:"tenant_id"`
	PrimaryIp4ID   types.Int64  `tfsdk:"primary_ip4_id"`
	PrimaryIp6ID   types.Int64  `tfsdk:"primary_ip6_id"`
	InterfaceCount types.Int64  `tfsdk:"interface_count"`
	Description    types.String `tfsdk:"description"`
	Comments       types.String `tfsdk:"comments"`
	OwnerID        types.Int64  `tfsdk:"owner_id"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
	Tags           types.Set    `tfsdk:"tags"`
	TagsAll        types.Set    `tfsdk:"tags_all"`
	CustomFields   types.Map    `tfsdk:"custom_fields"`
}

// expandVirtualDeviceContext converts the Terraform model of type "virtual_device_context" into an API request.
func expandVirtualDeviceContext(ctx context.Context, model *virtualDeviceContextResourceModel) (*netboxapi.VirtualDeviceContextRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualDeviceContextRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Identifier = conv.Int64Ptr(model.Identifier)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.PrimaryIp4 = conv.Int64Ptr(model.PrimaryIp4ID)
	requestDTO.PrimaryIp6 = conv.Int64Ptr(model.PrimaryIp6ID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualDeviceContext refreshes the Terraform model of type "virtual_device_context" from an API response.
func flattenVirtualDeviceContext(ctx context.Context, responseDTO *netboxapi.VirtualDeviceContextResponseDTO, model *virtualDeviceContextResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Identifier = conv.FromInt64Ptr(responseDTO.Identifier)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.PrimaryIp4ID = conv.FromInt64Ptr(responseDTO.PrimaryIp4)
	model.PrimaryIp6ID = conv.FromInt64Ptr(responseDTO.PrimaryIp6)
	model.InterfaceCount = conv.FromInt64Ptr(responseDTO.InterfaceCount)
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
