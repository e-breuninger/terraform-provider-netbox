// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ipsecProfileResourceModel is the Terraform state/plan model of the "ipsec_profile" resource.
type ipsecProfileResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	Name          types.String `tfsdk:"name"`
	Mode          types.String `tfsdk:"mode"`
	IkePolicyID   types.Int64  `tfsdk:"ike_policy_id"`
	IpsecPolicyID types.Int64  `tfsdk:"ipsec_policy_id"`
	Description   types.String `tfsdk:"description"`
	Comments      types.String `tfsdk:"comments"`
	OwnerID       types.Int64  `tfsdk:"owner_id"`
	Created       types.String `tfsdk:"created"`
	LastUpdated   types.String `tfsdk:"last_updated"`
	URL           types.String `tfsdk:"url"`
	Tags          types.Set    `tfsdk:"tags"`
	TagsAll       types.Set    `tfsdk:"tags_all"`
	CustomFields  types.Map    `tfsdk:"custom_fields"`
}

// expandIpsecProfile converts the Terraform model of type "ipsec_profile" into an API request.
func expandIpsecProfile(ctx context.Context, model *ipsecProfileResourceModel) (*netboxapi.IpsecProfileRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.IpsecProfileRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Mode = conv.StringPtr(model.Mode)
	requestDTO.IkePolicy = conv.Int64Ptr(model.IkePolicyID)
	requestDTO.IpsecPolicy = conv.Int64Ptr(model.IpsecPolicyID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenIpsecProfile refreshes the Terraform model of type "ipsec_profile" from an API response.
func flattenIpsecProfile(ctx context.Context, responseDTO *netboxapi.IpsecProfileResponseDTO, model *ipsecProfileResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Mode = conv.FromStringPtr(responseDTO.Mode)
	model.IkePolicyID = conv.FromInt64Ptr(responseDTO.IkePolicy)
	model.IpsecPolicyID = conv.FromInt64Ptr(responseDTO.IpsecPolicy)
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
