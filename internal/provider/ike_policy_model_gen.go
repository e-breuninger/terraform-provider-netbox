// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ikePolicyResourceModel is the Terraform state/plan model of the "ike_policy" resource.
type ikePolicyResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Version        types.Int64  `tfsdk:"version"`
	Mode           types.String `tfsdk:"mode"`
	IkeProposalIds types.Set    `tfsdk:"ike_proposal_ids"`
	PresharedKey   types.String `tfsdk:"preshared_key"`
	Description    types.String `tfsdk:"description"`
	Comments       types.String `tfsdk:"comments"`
	OwnerID        types.Int64  `tfsdk:"owner_id"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
	Tags           types.Set    `tfsdk:"tags"`
	TagsAll        types.Set    `tfsdk:"tags_all"`
	CustomFields   types.Map    `tfsdk:"custom_fields"`
}

// expandIkePolicy converts the Terraform model of type "ike_policy" into an API request.
func expandIkePolicy(ctx context.Context, model *ikePolicyResourceModel) (*netboxapi.IkePolicyRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.IkePolicyRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Version = conv.Int64Ptr(model.Version)
	requestDTO.Mode = conv.StringPtr(model.Mode)
	requestDTO.Proposals = conv.SetTo[int64](ctx, model.IkeProposalIds, &diags)
	requestDTO.PresharedKey = conv.StringPtr(model.PresharedKey)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenIkePolicy refreshes the Terraform model of type "ike_policy" from an API response.
func flattenIkePolicy(ctx context.Context, responseDTO *netboxapi.IkePolicyResponseDTO, model *ikePolicyResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Version = conv.FromInt64Ptr(responseDTO.Version)
	model.Mode = conv.FromStringPtr(responseDTO.Mode)
	model.IkeProposalIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Proposals, false, &diags)
	model.PresharedKey = conv.FromStringPtr(responseDTO.PresharedKey)
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
