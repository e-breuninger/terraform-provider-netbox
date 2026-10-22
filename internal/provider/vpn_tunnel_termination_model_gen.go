// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vpnTunnelTerminationResourceModel is the Terraform state/plan model of the "vpn_tunnel_termination" resource.
type vpnTunnelTerminationResourceModel struct {
	ID                        types.Int64  `tfsdk:"id"`
	VpnTunnelID               types.Int64  `tfsdk:"vpn_tunnel_id"`
	Role                      types.String `tfsdk:"role"`
	TerminationType           types.String `tfsdk:"termination_type"`
	TerminationID             types.Int64  `tfsdk:"termination_id"`
	DeviceInterfaceID         types.Int64  `tfsdk:"device_interface_id"`
	VirtualMachineInterfaceID types.Int64  `tfsdk:"virtual_machine_interface_id"`
	OutsideIPAddressID        types.Int64  `tfsdk:"outside_ip_address_id"`
	Created                   types.String `tfsdk:"created"`
	LastUpdated               types.String `tfsdk:"last_updated"`
	URL                       types.String `tfsdk:"url"`
	Tags                      types.Set    `tfsdk:"tags"`
	TagsAll                   types.Set    `tfsdk:"tags_all"`
	CustomFields              types.Map    `tfsdk:"custom_fields"`
}

// expandVpnTunnelTermination converts the Terraform model of type "vpn_tunnel_termination" into an API request.
func expandVpnTunnelTermination(ctx context.Context, model *vpnTunnelTerminationResourceModel) (*netboxapi.VpnTunnelTerminationRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VpnTunnelTerminationRequestDTO{}
	requestDTO.Tunnel = conv.Int64Ptr(model.VpnTunnelID)
	requestDTO.Role = conv.StringPtr(model.Role)
	requestDTO.TerminationType = conv.StringPtr(model.TerminationType)
	requestDTO.TerminationID = conv.Int64Ptr(model.TerminationID)
	requestDTO.DeviceInterfaceID = conv.Int64Ptr(model.DeviceInterfaceID)
	requestDTO.VirtualMachineInterfaceID = conv.Int64Ptr(model.VirtualMachineInterfaceID)
	requestDTO.OutsideIP = conv.Int64Ptr(model.OutsideIPAddressID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVpnTunnelTermination refreshes the Terraform model of type "vpn_tunnel_termination" from an API response.
func flattenVpnTunnelTermination(ctx context.Context, responseDTO *netboxapi.VpnTunnelTerminationResponseDTO, model *vpnTunnelTerminationResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.VpnTunnelID = conv.FromInt64Ptr(responseDTO.Tunnel)
	model.Role = conv.FromStringPtr(responseDTO.Role)
	model.TerminationType = conv.FromStringPtr(responseDTO.TerminationType)
	model.TerminationID = conv.FromInt64Ptr(responseDTO.TerminationID)
	model.DeviceInterfaceID = conv.FromInt64Ptr(responseDTO.DeviceInterfaceID)
	model.VirtualMachineInterfaceID = conv.FromInt64Ptr(responseDTO.VirtualMachineInterfaceID)
	model.OutsideIPAddressID = conv.FromInt64Ptr(responseDTO.OutsideIP)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
