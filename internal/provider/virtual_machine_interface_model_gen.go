// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// virtualMachineInterfaceResourceModel is the Terraform state/plan model of the "virtual_machine_interface" resource.
type virtualMachineInterfaceResourceModel struct {
	ID                      types.Int64  `tfsdk:"id"`
	VirtualMachineID        types.Int64  `tfsdk:"virtual_machine_id"`
	Name                    types.String `tfsdk:"name"`
	Enabled                 types.Bool   `tfsdk:"enabled"`
	Mtu                     types.Int64  `tfsdk:"mtu"`
	Mode                    types.String `tfsdk:"mode"`
	UntaggedVlanID          types.Int64  `tfsdk:"untagged_vlan_id"`
	TaggedVlanIds           types.Set    `tfsdk:"tagged_vlan_ids"`
	QinqSvlanID             types.Int64  `tfsdk:"qinq_svlan_id"`
	VlanTranslationPolicyID types.Int64  `tfsdk:"vlan_translation_policy_id"`
	VrfID                   types.Int64  `tfsdk:"vrf_id"`
	ParentID                types.Int64  `tfsdk:"parent_id"`
	BridgeID                types.Int64  `tfsdk:"bridge_id"`
	Description             types.String `tfsdk:"description"`
	OwnerID                 types.Int64  `tfsdk:"owner_id"`
	Created                 types.String `tfsdk:"created"`
	LastUpdated             types.String `tfsdk:"last_updated"`
	URL                     types.String `tfsdk:"url"`
	Tags                    types.Set    `tfsdk:"tags"`
	TagsAll                 types.Set    `tfsdk:"tags_all"`
	CustomFields            types.Map    `tfsdk:"custom_fields"`
}

// expandVirtualMachineInterface converts the Terraform model of type "virtual_machine_interface" into an API request.
func expandVirtualMachineInterface(ctx context.Context, model *virtualMachineInterfaceResourceModel) (*netboxapi.VirtualMachineInterfaceRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualMachineInterfaceRequestDTO{}
	requestDTO.VirtualMachine = conv.Int64Ptr(model.VirtualMachineID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.Mtu = conv.Int64Ptr(model.Mtu)
	requestDTO.Mode = conv.StringPtr(model.Mode)
	requestDTO.UntaggedVlan = conv.Int64Ptr(model.UntaggedVlanID)
	requestDTO.TaggedVlans = conv.SetTo[int64](ctx, model.TaggedVlanIds, &diags)
	requestDTO.QinqSvlan = conv.Int64Ptr(model.QinqSvlanID)
	requestDTO.VlanTranslationPolicy = conv.Int64Ptr(model.VlanTranslationPolicyID)
	requestDTO.Vrf = conv.Int64Ptr(model.VrfID)
	requestDTO.Parent = conv.Int64Ptr(model.ParentID)
	requestDTO.Bridge = conv.Int64Ptr(model.BridgeID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualMachineInterface refreshes the Terraform model of type "virtual_machine_interface" from an API response.
func flattenVirtualMachineInterface(ctx context.Context, responseDTO *netboxapi.VirtualMachineInterfaceResponseDTO, model *virtualMachineInterfaceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.VirtualMachineID = conv.FromInt64Ptr(responseDTO.VirtualMachine)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.Mtu = conv.FromInt64Ptr(responseDTO.Mtu)
	model.Mode = conv.FromStringPtr(responseDTO.Mode)
	model.UntaggedVlanID = conv.FromInt64Ptr(responseDTO.UntaggedVlan)
	model.TaggedVlanIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.TaggedVlans, model.TaggedVlanIds.IsNull() || model.TaggedVlanIds.IsUnknown(), &diags)
	model.QinqSvlanID = conv.FromInt64Ptr(responseDTO.QinqSvlan)
	model.VlanTranslationPolicyID = conv.FromInt64Ptr(responseDTO.VlanTranslationPolicy)
	model.VrfID = conv.FromInt64Ptr(responseDTO.Vrf)
	model.ParentID = conv.FromInt64Ptr(responseDTO.Parent)
	model.BridgeID = conv.FromInt64Ptr(responseDTO.Bridge)
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
