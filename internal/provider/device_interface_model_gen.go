// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// deviceInterfaceResourceModel is the Terraform state/plan model of the "device_interface" resource.
type deviceInterfaceResourceModel struct {
	ID                      types.Int64  `tfsdk:"id"`
	DeviceID                types.Int64  `tfsdk:"device_id"`
	Name                    types.String `tfsdk:"name"`
	Type                    types.String `tfsdk:"type"`
	Label                   types.String `tfsdk:"label"`
	Enabled                 types.Bool   `tfsdk:"enabled"`
	MgmtOnly                types.Bool   `tfsdk:"mgmt_only"`
	Mgmtonly                types.Bool   `tfsdk:"mgmtonly"`
	MarkConnected           types.Bool   `tfsdk:"mark_connected"`
	Mtu                     types.Int64  `tfsdk:"mtu"`
	Speed                   types.Int64  `tfsdk:"speed"`
	Duplex                  types.String `tfsdk:"duplex"`
	Wwn                     types.String `tfsdk:"wwn"`
	Mode                    types.String `tfsdk:"mode"`
	UntaggedVlanID          types.Int64  `tfsdk:"untagged_vlan_id"`
	UntaggedVlan            types.Int64  `tfsdk:"untagged_vlan"`
	TaggedVlanIds           types.Set    `tfsdk:"tagged_vlan_ids"`
	TaggedVlans             types.Set    `tfsdk:"tagged_vlans"`
	ModuleID                types.Int64  `tfsdk:"module_id"`
	LagDeviceInterfaceID    types.Int64  `tfsdk:"lag_device_interface_id"`
	ParentDeviceInterfaceID types.Int64  `tfsdk:"parent_device_interface_id"`
	BridgeID                types.Int64  `tfsdk:"bridge_id"`
	VrfID                   types.Int64  `tfsdk:"vrf_id"`
	VdcIds                  types.Set    `tfsdk:"vdc_ids"`
	Description             types.String `tfsdk:"description"`
	OwnerID                 types.Int64  `tfsdk:"owner_id"`
	Created                 types.String `tfsdk:"created"`
	LastUpdated             types.String `tfsdk:"last_updated"`
	URL                     types.String `tfsdk:"url"`
	Tags                    types.Set    `tfsdk:"tags"`
	TagsAll                 types.Set    `tfsdk:"tags_all"`
	CustomFields            types.Map    `tfsdk:"custom_fields"`
}

// expandDeviceInterface converts the Terraform model of type "device_interface" into an API request.
func expandDeviceInterface(ctx context.Context, model *deviceInterfaceResourceModel) (*netboxapi.DeviceInterfaceRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceInterfaceRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.MgmtOnly = conv.BoolPtr(model.MgmtOnly)
	requestDTO.MarkConnected = conv.BoolPtr(model.MarkConnected)
	requestDTO.Mtu = conv.Int64Ptr(model.Mtu)
	requestDTO.Speed = conv.Int64Ptr(model.Speed)
	requestDTO.Duplex = conv.StringPtr(model.Duplex)
	requestDTO.Wwn = conv.StringPtr(model.Wwn)
	requestDTO.Mode = conv.StringPtr(model.Mode)
	requestDTO.UntaggedVlan = conv.Int64Ptr(model.UntaggedVlanID)
	requestDTO.TaggedVlans = conv.SetTo[int64](ctx, model.TaggedVlanIds, &diags)
	requestDTO.Module = conv.Int64Ptr(model.ModuleID)
	requestDTO.Lag = conv.Int64Ptr(model.LagDeviceInterfaceID)
	requestDTO.Parent = conv.Int64Ptr(model.ParentDeviceInterfaceID)
	requestDTO.Bridge = conv.Int64Ptr(model.BridgeID)
	requestDTO.Vrf = conv.Int64Ptr(model.VrfID)
	requestDTO.Vdcs = conv.SetTo[int64](ctx, model.VdcIds, &diags)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDeviceInterface refreshes the Terraform model of type "device_interface" from an API response.
func flattenDeviceInterface(ctx context.Context, responseDTO *netboxapi.DeviceInterfaceResponseDTO, model *deviceInterfaceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.MgmtOnly = conv.FromBoolPtr(responseDTO.MgmtOnly)
	model.MarkConnected = conv.FromBoolPtr(responseDTO.MarkConnected)
	model.Mtu = conv.FromInt64Ptr(responseDTO.Mtu)
	model.Speed = conv.FromInt64Ptr(responseDTO.Speed)
	model.Duplex = conv.FromStringPtr(responseDTO.Duplex)
	model.Wwn = conv.FromStringPtr(responseDTO.Wwn)
	model.Mode = conv.FromStringPtr(responseDTO.Mode)
	model.UntaggedVlanID = conv.FromInt64Ptr(responseDTO.UntaggedVlan)
	model.TaggedVlanIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.TaggedVlans, model.TaggedVlanIds.IsNull() || model.TaggedVlanIds.IsUnknown(), &diags)
	model.ModuleID = conv.FromInt64Ptr(responseDTO.Module)
	model.LagDeviceInterfaceID = conv.FromInt64Ptr(responseDTO.Lag)
	model.ParentDeviceInterfaceID = conv.FromInt64Ptr(responseDTO.Parent)
	model.BridgeID = conv.FromInt64Ptr(responseDTO.Bridge)
	model.VrfID = conv.FromInt64Ptr(responseDTO.Vrf)
	model.VdcIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Vdcs, model.VdcIds.IsNull() || model.VdcIds.IsUnknown(), &diags)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	model.Mgmtonly = model.MgmtOnly
	model.UntaggedVlan = model.UntaggedVlanID
	model.TaggedVlans = model.TaggedVlanIds
	return diags
}
