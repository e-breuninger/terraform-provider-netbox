// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vpnTunnelResourceModel is the Terraform state/plan model of the "vpn_tunnel" resource.
type vpnTunnelResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Status           types.String `tfsdk:"status"`
	Encapsulation    types.String `tfsdk:"encapsulation"`
	VpnTunnelGroupID types.Int64  `tfsdk:"vpn_tunnel_group_id"`
	TunnelGroupID    types.Int64  `tfsdk:"tunnel_group_id"`
	TenantID         types.Int64  `tfsdk:"tenant_id"`
	TunnelID         types.Int64  `tfsdk:"tunnel_id"`
	Description      types.String `tfsdk:"description"`
	Comments         types.String `tfsdk:"comments"`
	OwnerID          types.Int64  `tfsdk:"owner_id"`
	Created          types.String `tfsdk:"created"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	URL              types.String `tfsdk:"url"`
	TerminationCount types.Int64  `tfsdk:"termination_count"`
	Tags             types.Set    `tfsdk:"tags"`
	TagsAll          types.Set    `tfsdk:"tags_all"`
	CustomFields     types.Map    `tfsdk:"custom_fields"`
}

// expandVpnTunnel converts the Terraform model of type "vpn_tunnel" into an API request.
func expandVpnTunnel(ctx context.Context, model *vpnTunnelResourceModel) (*netboxapi.VpnTunnelRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VpnTunnelRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Encapsulation = conv.StringPtr(model.Encapsulation)
	requestDTO.Group = conv.Int64Ptr(model.VpnTunnelGroupID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.TunnelID = conv.Int64Ptr(model.TunnelID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVpnTunnel refreshes the Terraform model of type "vpn_tunnel" from an API response.
func flattenVpnTunnel(ctx context.Context, responseDTO *netboxapi.VpnTunnelResponseDTO, model *vpnTunnelResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Encapsulation = conv.FromStringPtr(responseDTO.Encapsulation)
	model.VpnTunnelGroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.TunnelID = conv.FromInt64Ptr(responseDTO.TunnelID)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.TerminationCount = conv.FromInt64Ptr(responseDTO.TerminationsCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	model.TunnelGroupID = model.VpnTunnelGroupID
	return diags
}
