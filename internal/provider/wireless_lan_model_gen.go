// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// wirelessLanResourceModel is the Terraform state/plan model of the "wireless_lan" resource.
type wirelessLanResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Ssid         types.String `tfsdk:"ssid"`
	Description  types.String `tfsdk:"description"`
	GroupID      types.Int64  `tfsdk:"group_id"`
	Status       types.String `tfsdk:"status"`
	VlanID       types.Int64  `tfsdk:"vlan_id"`
	TenantID     types.Int64  `tfsdk:"tenant_id"`
	AuthType     types.String `tfsdk:"auth_type"`
	AuthCipher   types.String `tfsdk:"auth_cipher"`
	AuthPsk      types.String `tfsdk:"auth_psk"`
	ScopeType    types.String `tfsdk:"scope_type"`
	ScopeID      types.Int64  `tfsdk:"scope_id"`
	SiteID       types.Int64  `tfsdk:"site_id"`
	LocationID   types.Int64  `tfsdk:"location_id"`
	RegionID     types.Int64  `tfsdk:"region_id"`
	SiteGroupID  types.Int64  `tfsdk:"site_group_id"`
	Comments     types.String `tfsdk:"comments"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandWirelessLan converts the Terraform model of type "wireless_lan" into an API request.
func expandWirelessLan(ctx context.Context, model *wirelessLanResourceModel) (*netboxapi.WirelessLanRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.WirelessLanRequestDTO{}
	requestDTO.Ssid = conv.StringPtr(model.Ssid)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Group = conv.Int64Ptr(model.GroupID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Vlan = conv.Int64Ptr(model.VlanID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.AuthType = conv.StringPtr(model.AuthType)
	requestDTO.AuthCipher = conv.StringPtr(model.AuthCipher)
	requestDTO.AuthPsk = conv.StringPtr(model.AuthPsk)
	requestDTO.ScopeType = conv.StringPtr(model.ScopeType)
	requestDTO.ScopeID = conv.Int64Ptr(model.ScopeID)
	requestDTO.SiteID = conv.Int64Ptr(model.SiteID)
	requestDTO.LocationID = conv.Int64Ptr(model.LocationID)
	requestDTO.RegionID = conv.Int64Ptr(model.RegionID)
	requestDTO.SiteGroupID = conv.Int64Ptr(model.SiteGroupID)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenWirelessLan refreshes the Terraform model of type "wireless_lan" from an API response.
func flattenWirelessLan(ctx context.Context, responseDTO *netboxapi.WirelessLanResponseDTO, model *wirelessLanResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Ssid = conv.FromStringPtr(responseDTO.Ssid)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.GroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.VlanID = conv.FromInt64Ptr(responseDTO.Vlan)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.AuthType = conv.FromStringPtr(responseDTO.AuthType)
	model.AuthCipher = conv.FromStringPtr(responseDTO.AuthCipher)
	model.AuthPsk = conv.FromStringPtr(responseDTO.AuthPsk)
	model.ScopeType = conv.FromStringPtr(responseDTO.ScopeType)
	model.ScopeID = conv.FromInt64Ptr(responseDTO.ScopeID)
	model.SiteID = conv.FromInt64Ptr(responseDTO.SiteID)
	model.LocationID = conv.FromInt64Ptr(responseDTO.LocationID)
	model.RegionID = conv.FromInt64Ptr(responseDTO.RegionID)
	model.SiteGroupID = conv.FromInt64Ptr(responseDTO.SiteGroupID)
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
