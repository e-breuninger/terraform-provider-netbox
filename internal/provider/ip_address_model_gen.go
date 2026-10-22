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

// ipAddressResourceModel is the Terraform state/plan model of the "ip_address" resource.
type ipAddressResourceModel struct {
	ID                        types.Int64    `tfsdk:"id"`
	IPAddress                 conv.IPAddress `tfsdk:"ip_address"`
	Status                    types.String   `tfsdk:"status"`
	Role                      types.String   `tfsdk:"role"`
	DNSName                   types.String   `tfsdk:"dns_name"`
	Description               types.String   `tfsdk:"description"`
	Comments                  types.String   `tfsdk:"comments"`
	TenantID                  types.Int64    `tfsdk:"tenant_id"`
	VrfID                     types.Int64    `tfsdk:"vrf_id"`
	NatInsideID               types.Int64    `tfsdk:"nat_inside_id"`
	NatInsideAddressID        types.Int64    `tfsdk:"nat_inside_address_id"`
	NatOutsideIds             types.Set      `tfsdk:"nat_outside_ids"`
	AssignedObjectType        types.String   `tfsdk:"assigned_object_type"`
	AssignedObjectID          types.Int64    `tfsdk:"assigned_object_id"`
	DeviceInterfaceID         types.Int64    `tfsdk:"device_interface_id"`
	VirtualMachineInterfaceID types.Int64    `tfsdk:"virtual_machine_interface_id"`
	AssignedObject            types.Object   `tfsdk:"assigned_object"`
	Family                    types.Int64    `tfsdk:"family"`
	OwnerID                   types.Int64    `tfsdk:"owner_id"`
	Created                   types.String   `tfsdk:"created"`
	LastUpdated               types.String   `tfsdk:"last_updated"`
	URL                       types.String   `tfsdk:"url"`
	Tags                      types.Set      `tfsdk:"tags"`
	TagsAll                   types.Set      `tfsdk:"tags_all"`
	CustomFields              types.Map      `tfsdk:"custom_fields"`
}

// expandIPAddress converts the Terraform model of type "ip_address" into an API request.
func expandIPAddress(ctx context.Context, model *ipAddressResourceModel) (*netboxapi.IPAddressRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.IPAddressRequestDTO{}
	requestDTO.Address = conv.StringPtr(model.IPAddress.StringValue)
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

// flattenIPAddress refreshes the Terraform model of type "ip_address" from an API response.
func flattenIPAddress(ctx context.Context, responseDTO *netboxapi.IPAddressResponseDTO, model *ipAddressResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.IPAddress = conv.NewIPAddressPointerValue(conv.EmptyStringAsNil(responseDTO.Address))
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Role = conv.FromStringPtr(responseDTO.Role)
	model.DNSName = conv.FromStringPtr(responseDTO.DNSName)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.VrfID = conv.FromInt64Ptr(responseDTO.Vrf)
	model.NatInsideID = conv.FromInt64Ptr(responseDTO.NatInside)
	model.NatOutsideIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.NatOutside, false, &diags)
	model.AssignedObjectType = conv.FromStringPtr(responseDTO.AssignedObjectType)
	model.AssignedObjectID = conv.FromInt64Ptr(responseDTO.AssignedObjectID)
	model.DeviceInterfaceID = conv.FromInt64Ptr(responseDTO.DeviceInterfaceID)
	model.VirtualMachineInterfaceID = conv.FromInt64Ptr(responseDTO.VirtualMachineInterfaceID)
	if responseDTO.AssignedObject == nil {
		model.AssignedObject = types.ObjectNull(ipAddressAssignedObjectAttrTypes())
	} else {
		var o ipAddressAssignedObjectModel
		diags.Append(flattenIPAddressAssignedObject(ctx, responseDTO.AssignedObject, &o)...)
		v, d := types.ObjectValueFrom(ctx, ipAddressAssignedObjectAttrTypes(), o)
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
	model.NatInsideAddressID = model.NatInsideID
	return diags
}

// ipAddressAssignedObjectModel is the model of the assigned_object object.
type ipAddressAssignedObjectModel struct {
	ID     types.Int64  `tfsdk:"id"`
	Name   types.String `tfsdk:"name"`
	Device types.Object `tfsdk:"device"`
}

// ipAddressAssignedObjectAttrTypes is the attribute type map of ipAddressAssignedObjectModel.
func ipAddressAssignedObjectAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":     types.Int64Type,
		"name":   types.StringType,
		"device": types.ObjectType{AttrTypes: ipAddressAssignedObjectDeviceAttrTypes()},
	}
}

func flattenIPAddressAssignedObject(ctx context.Context, responseDTO *netboxapi.IPAddressResponseDTOAssignedObject, model *ipAddressAssignedObjectModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.IPAddressResponseDTOAssignedObject{}
	}
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	if responseDTO.Device == nil {
		model.Device = types.ObjectNull(ipAddressAssignedObjectDeviceAttrTypes())
	} else {
		var o ipAddressAssignedObjectDeviceModel
		diags.Append(flattenIPAddressAssignedObjectDevice(ctx, responseDTO.Device, &o)...)
		v, d := types.ObjectValueFrom(ctx, ipAddressAssignedObjectDeviceAttrTypes(), o)
		diags.Append(d...)
		model.Device = v
	}
	return diags
}

// ipAddressAssignedObjectDeviceModel is the model of the device object.
type ipAddressAssignedObjectDeviceModel struct {
	ID   types.Int64  `tfsdk:"id"`
	Name types.String `tfsdk:"name"`
}

// ipAddressAssignedObjectDeviceAttrTypes is the attribute type map of ipAddressAssignedObjectDeviceModel.
func ipAddressAssignedObjectDeviceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":   types.Int64Type,
		"name": types.StringType,
	}
}

func flattenIPAddressAssignedObjectDevice(ctx context.Context, responseDTO *netboxapi.IPAddressResponseDTOAssignedObjectDevice, model *ipAddressAssignedObjectDeviceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.IPAddressResponseDTOAssignedObjectDevice{}
	}
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	return diags
}
