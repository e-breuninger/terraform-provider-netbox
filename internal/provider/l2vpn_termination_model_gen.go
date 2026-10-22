// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// l2vpnTerminationResourceModel is the Terraform state/plan model of the "l2vpn_termination" resource.
type l2vpnTerminationResourceModel struct {
	ID                        types.Int64  `tfsdk:"id"`
	L2vpnID                   types.Int64  `tfsdk:"l2vpn_id"`
	AssignedObjectType        types.String `tfsdk:"assigned_object_type"`
	AssignedObjectID          types.Int64  `tfsdk:"assigned_object_id"`
	DeviceInterfaceID         types.Int64  `tfsdk:"device_interface_id"`
	VirtualMachineInterfaceID types.Int64  `tfsdk:"virtual_machine_interface_id"`
	VlanID                    types.Int64  `tfsdk:"vlan_id"`
	Created                   types.String `tfsdk:"created"`
	LastUpdated               types.String `tfsdk:"last_updated"`
	URL                       types.String `tfsdk:"url"`
	Tags                      types.Set    `tfsdk:"tags"`
	TagsAll                   types.Set    `tfsdk:"tags_all"`
	CustomFields              types.Map    `tfsdk:"custom_fields"`
}

// expandL2vpnTermination converts the Terraform model of type "l2vpn_termination" into an API request.
func expandL2vpnTermination(ctx context.Context, model *l2vpnTerminationResourceModel) (*netboxapi.L2vpnTerminationRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.L2vpnTerminationRequestDTO{}
	requestDTO.L2vpn = conv.Int64Ptr(model.L2vpnID)
	requestDTO.AssignedObjectType = conv.StringPtr(model.AssignedObjectType)
	requestDTO.AssignedObjectID = conv.Int64Ptr(model.AssignedObjectID)
	requestDTO.DeviceInterfaceID = conv.Int64Ptr(model.DeviceInterfaceID)
	requestDTO.VirtualMachineInterfaceID = conv.Int64Ptr(model.VirtualMachineInterfaceID)
	requestDTO.VlanID = conv.Int64Ptr(model.VlanID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenL2vpnTermination refreshes the Terraform model of type "l2vpn_termination" from an API response.
func flattenL2vpnTermination(ctx context.Context, responseDTO *netboxapi.L2vpnTerminationResponseDTO, model *l2vpnTerminationResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.L2vpnID = conv.FromInt64Ptr(responseDTO.L2vpn)
	model.AssignedObjectType = conv.FromStringPtr(responseDTO.AssignedObjectType)
	model.AssignedObjectID = conv.FromInt64Ptr(responseDTO.AssignedObjectID)
	model.DeviceInterfaceID = conv.FromInt64Ptr(responseDTO.DeviceInterfaceID)
	model.VirtualMachineInterfaceID = conv.FromInt64Ptr(responseDTO.VirtualMachineInterfaceID)
	model.VlanID = conv.FromInt64Ptr(responseDTO.VlanID)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
