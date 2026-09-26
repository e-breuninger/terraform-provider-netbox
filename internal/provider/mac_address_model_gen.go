// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// macAddressResourceModel is the Terraform state/plan model of the "mac_address" resource.
type macAddressResourceModel struct {
	ID                        types.Int64     `tfsdk:"id"`
	MacAddress                conv.MACAddress `tfsdk:"mac_address"`
	AssignedObjectType        types.String    `tfsdk:"assigned_object_type"`
	AssignedObjectID          types.Int64     `tfsdk:"assigned_object_id"`
	DeviceInterfaceID         types.Int64     `tfsdk:"device_interface_id"`
	VirtualMachineInterfaceID types.Int64     `tfsdk:"virtual_machine_interface_id"`
	Description               types.String    `tfsdk:"description"`
	Comments                  types.String    `tfsdk:"comments"`
	OwnerID                   types.Int64     `tfsdk:"owner_id"`
	LastUpdated               types.String    `tfsdk:"last_updated"`
	URL                       types.String    `tfsdk:"url"`
	Tags                      types.Set       `tfsdk:"tags"`
	TagsAll                   types.Set       `tfsdk:"tags_all"`
	CustomFields              types.Map       `tfsdk:"custom_fields"`
}

// expandMacAddress converts the Terraform model of type "mac_address" into an API request.
func expandMacAddress(ctx context.Context, model *macAddressResourceModel) (*netboxapi.MacAddressRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.MacAddressRequestDTO{}
	requestDTO.MacAddress = conv.StringPtr(model.MacAddress.StringValue)
	requestDTO.AssignedObjectType = conv.StringPtr(model.AssignedObjectType)
	requestDTO.AssignedObjectID = conv.Int64Ptr(model.AssignedObjectID)
	requestDTO.DeviceInterfaceID = conv.Int64Ptr(model.DeviceInterfaceID)
	requestDTO.VirtualMachineInterfaceID = conv.Int64Ptr(model.VirtualMachineInterfaceID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenMacAddress refreshes the Terraform model of type "mac_address" from an API response.
func flattenMacAddress(ctx context.Context, responseDTO *netboxapi.MacAddressResponseDTO, model *macAddressResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.MacAddress = conv.NewMACAddressPointerValue(conv.EmptyStringAsNil(responseDTO.MacAddress))
	model.AssignedObjectType = conv.FromStringPtr(responseDTO.AssignedObjectType)
	model.AssignedObjectID = conv.FromInt64Ptr(responseDTO.AssignedObjectID)
	model.DeviceInterfaceID = conv.FromInt64Ptr(responseDTO.DeviceInterfaceID)
	model.VirtualMachineInterfaceID = conv.FromInt64Ptr(responseDTO.VirtualMachineInterfaceID)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
