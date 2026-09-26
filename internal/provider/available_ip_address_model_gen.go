// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// availableIPAddressResourceModel is the Terraform state/plan model of the "available_ip_address" resource.
type availableIPAddressResourceModel struct {
	ID                        types.Int64  `tfsdk:"id"`
	PrefixID                  types.Int64  `tfsdk:"prefix_id"`
	IPRangeID                 types.Int64  `tfsdk:"ip_range_id"`
	IPAddress                 types.String `tfsdk:"ip_address"`
	Status                    types.String `tfsdk:"status"`
	Role                      types.String `tfsdk:"role"`
	DNSName                   types.String `tfsdk:"dns_name"`
	Description               types.String `tfsdk:"description"`
	Comments                  types.String `tfsdk:"comments"`
	TenantID                  types.Int64  `tfsdk:"tenant_id"`
	VrfID                     types.Int64  `tfsdk:"vrf_id"`
	NatInsideID               types.Int64  `tfsdk:"nat_inside_id"`
	AssignedObjectType        types.String `tfsdk:"assigned_object_type"`
	AssignedObjectID          types.Int64  `tfsdk:"assigned_object_id"`
	DeviceInterfaceID         types.Int64  `tfsdk:"device_interface_id"`
	VirtualMachineInterfaceID types.Int64  `tfsdk:"virtual_machine_interface_id"`
	AssignedObject            types.Object `tfsdk:"assigned_object"`
	Family                    types.Int64  `tfsdk:"family"`
	OwnerID                   types.Int64  `tfsdk:"owner_id"`
	Created                   types.String `tfsdk:"created"`
	LastUpdated               types.String `tfsdk:"last_updated"`
	URL                       types.String `tfsdk:"url"`
	Tags                      types.Set    `tfsdk:"tags"`
	TagsAll                   types.Set    `tfsdk:"tags_all"`
	CustomFields              types.Map    `tfsdk:"custom_fields"`
}

// expandAvailableIPAddress converts the Terraform model of type "available_ip_address" into an API request.
func expandAvailableIPAddress(ctx context.Context, model *availableIPAddressResourceModel) (*netboxapi.AvailableIPAddressRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.AvailableIPAddressRequestDTO{}
	requestDTO.PrefixID = conv.Int64Ptr(model.PrefixID)
	requestDTO.IPRangeID = conv.Int64Ptr(model.IPRangeID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Role = conv.StringPtr(model.Role)
	requestDTO.DNSName = conv.StringPtr(model.DNSName)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Vrf = conv.Int64Ptr(model.VrfID)
	requestDTO.NatInside = conv.Int64Ptr(model.NatInsideID)
	requestDTO.AssignedObjectType = conv.StringPtr(model.AssignedObjectType)
	requestDTO.AssignedObjectID = conv.Int64Ptr(model.AssignedObjectID)
	requestDTO.DeviceInterfaceID = conv.Int64Ptr(model.DeviceInterfaceID)
	requestDTO.VirtualMachineInterfaceID = conv.Int64Ptr(model.VirtualMachineInterfaceID)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenAvailableIPAddress refreshes the Terraform model of type "available_ip_address" from an API response.
func flattenAvailableIPAddress(ctx context.Context, responseDTO *netboxapi.AvailableIPAddressResponseDTO, model *availableIPAddressResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.IPAddress = conv.FromStringPtr(responseDTO.Address)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Role = conv.FromStringPtr(responseDTO.Role)
	model.DNSName = conv.FromStringPtr(responseDTO.DNSName)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.VrfID = conv.FromInt64Ptr(responseDTO.Vrf)
	model.NatInsideID = conv.FromInt64Ptr(responseDTO.NatInside)
	model.AssignedObjectType = conv.FromStringPtr(responseDTO.AssignedObjectType)
	model.AssignedObjectID = conv.FromInt64Ptr(responseDTO.AssignedObjectID)
	model.DeviceInterfaceID = conv.FromInt64Ptr(responseDTO.DeviceInterfaceID)
	model.VirtualMachineInterfaceID = conv.FromInt64Ptr(responseDTO.VirtualMachineInterfaceID)
	if responseDTO.AssignedObject == nil {
		model.AssignedObject = types.ObjectNull(availableIPAddressAssignedObjectAttrTypes())
	} else {
		var o availableIPAddressAssignedObjectModel
		diags.Append(flattenAvailableIPAddressAssignedObject(ctx, responseDTO.AssignedObject, &o)...)
		v, d := types.ObjectValueFrom(ctx, availableIPAddressAssignedObjectAttrTypes(), o)
		diags.Append(d...)
		model.AssignedObject = v
	}
	model.Family = conv.FromInt64Ptr(responseDTO.Family)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}

// availableIPAddressAssignedObjectModel is the model of the assigned_object object.
type availableIPAddressAssignedObjectModel struct {
	ID     types.Int64  `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Device types.Object `tfsdk:"device"`
}

// availableIPAddressAssignedObjectAttrTypes is the attribute type map of availableIPAddressAssignedObjectModel.
func availableIPAddressAssignedObjectAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":     types.Int64Type,
		"name":   types.StringType,
		"device": types.ObjectType{AttrTypes: availableIPAddressAssignedObjectDeviceAttrTypes()},
	}
}

func flattenAvailableIPAddressAssignedObject(ctx context.Context, responseDTO *netboxapi.AvailableIPAddressResponseDTOAssignedObject, model *availableIPAddressAssignedObjectModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.AvailableIPAddressResponseDTOAssignedObject{}
	}
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	if responseDTO.Device == nil {
		model.Device = types.ObjectNull(availableIPAddressAssignedObjectDeviceAttrTypes())
	} else {
		var o availableIPAddressAssignedObjectDeviceModel
		diags.Append(flattenAvailableIPAddressAssignedObjectDevice(ctx, responseDTO.Device, &o)...)
		v, d := types.ObjectValueFrom(ctx, availableIPAddressAssignedObjectDeviceAttrTypes(), o)
		diags.Append(d...)
		model.Device = v
	}
	return diags
}

// availableIPAddressAssignedObjectDeviceModel is the model of the device object.
type availableIPAddressAssignedObjectDeviceModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// availableIPAddressAssignedObjectDeviceAttrTypes is the attribute type map of availableIPAddressAssignedObjectDeviceModel.
func availableIPAddressAssignedObjectDeviceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   types.Int64Type,
		"name": types.StringType,
	}
}

func flattenAvailableIPAddressAssignedObjectDevice(ctx context.Context, responseDTO *netboxapi.AvailableIPAddressResponseDTOAssignedObjectDevice, model *availableIPAddressAssignedObjectDeviceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.AvailableIPAddressResponseDTOAssignedObjectDevice{}
	}
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	return diags
}
