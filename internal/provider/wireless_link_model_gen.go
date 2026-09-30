// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// wirelessLinkResourceModel is the Terraform state/plan model of the "wireless_link" resource.
type wirelessLinkResourceModel struct {
	ID           types.Int64   `tfsdk:"id"`
	InterfaceAID types.Int64   `tfsdk:"interface_a_id"`
	InterfaceBID types.Int64   `tfsdk:"interface_b_id"`
	Ssid         types.String  `tfsdk:"ssid"`
	Status       types.String  `tfsdk:"status"`
	TenantID     types.Int64   `tfsdk:"tenant_id"`
	AuthType     types.String  `tfsdk:"auth_type"`
	AuthCipher   types.String  `tfsdk:"auth_cipher"`
	AuthPsk      types.String  `tfsdk:"auth_psk"`
	Distance     types.Float64 `tfsdk:"distance"`
	DistanceUnit types.String  `tfsdk:"distance_unit"`
	Description  types.String  `tfsdk:"description"`
	Comments     types.String  `tfsdk:"comments"`
	OwnerID      types.Int64   `tfsdk:"owner_id"`
	Created      types.String  `tfsdk:"created"`
	LastUpdated  types.String  `tfsdk:"last_updated"`
	URL          types.String  `tfsdk:"url"`
	Tags         types.Set     `tfsdk:"tags"`
	TagsAll      types.Set     `tfsdk:"tags_all"`
	CustomFields types.Map     `tfsdk:"custom_fields"`
}

// expandWirelessLink converts the Terraform model of type "wireless_link" into an API request.
func expandWirelessLink(ctx context.Context, model *wirelessLinkResourceModel) (*netboxapi.WirelessLinkRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.WirelessLinkRequestDTO{}
	requestDTO.Interfacea = conv.Int64Ptr(model.InterfaceAID)
	requestDTO.Interfaceb = conv.Int64Ptr(model.InterfaceBID)
	requestDTO.Ssid = conv.StringPtr(model.Ssid)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.AuthType = conv.StringPtr(model.AuthType)
	requestDTO.AuthCipher = conv.StringPtr(model.AuthCipher)
	requestDTO.AuthPsk = conv.StringPtr(model.AuthPsk)
	requestDTO.Distance = conv.Float64Ptr(model.Distance)
	requestDTO.DistanceUnit = conv.StringPtr(model.DistanceUnit)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenWirelessLink refreshes the Terraform model of type "wireless_link" from an API response.
func flattenWirelessLink(ctx context.Context, responseDTO *netboxapi.WirelessLinkResponseDTO, model *wirelessLinkResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.InterfaceAID = conv.FromInt64Ptr(responseDTO.Interfacea)
	model.InterfaceBID = conv.FromInt64Ptr(responseDTO.Interfaceb)
	model.Ssid = conv.FromStringPtr(responseDTO.Ssid)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.AuthType = conv.FromStringPtr(responseDTO.AuthType)
	model.AuthCipher = conv.FromStringPtr(responseDTO.AuthCipher)
	model.AuthPsk = conv.FromStringPtr(responseDTO.AuthPsk)
	model.Distance = conv.FromFloat64Ptr(responseDTO.Distance)
	model.DistanceUnit = conv.FromStringPtr(responseDTO.DistanceUnit)
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
