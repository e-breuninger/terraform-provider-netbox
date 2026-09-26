// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ipsecPolicyResourceModel is the Terraform state/plan model of the "ipsec_policy" resource.
type ipsecPolicyResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	IpsecProposalIds types.Set    `tfsdk:"ipsec_proposal_ids"`
	PfsGroup         types.Int64  `tfsdk:"pfs_group"`
	Description      types.String `tfsdk:"description"`
	Comments         types.String `tfsdk:"comments"`
	OwnerID          types.Int64  `tfsdk:"owner_id"`
	Created          types.String `tfsdk:"created"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	URL              types.String `tfsdk:"url"`
	Tags             types.Set    `tfsdk:"tags"`
	TagsAll          types.Set    `tfsdk:"tags_all"`
	CustomFields     types.Map    `tfsdk:"custom_fields"`
}

// expandIpsecPolicy converts the Terraform model of type "ipsec_policy" into an API request.
func expandIpsecPolicy(ctx context.Context, model *ipsecPolicyResourceModel) (*netboxapi.IpsecPolicyRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.IpsecPolicyRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Proposals = conv.SetTo[int64](ctx, model.IpsecProposalIds, &diags)
	requestDTO.PfsGroup = conv.Int64Ptr(model.PfsGroup)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenIpsecPolicy refreshes the Terraform model of type "ipsec_policy" from an API response.
func flattenIpsecPolicy(ctx context.Context, responseDTO *netboxapi.IpsecPolicyResponseDTO, model *ipsecPolicyResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.IpsecProposalIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Proposals, false, &diags)
	model.PfsGroup = conv.FromInt64Ptr(responseDTO.PfsGroup)
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
