// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ikeProposalResourceModel is the Terraform state/plan model of the "ike_proposal" resource.
type ikeProposalResourceModel struct {
	ID                      types.Int64  `tfsdk:"id"`
	Name                    types.String `tfsdk:"name"`
	AuthenticationMethod    types.String `tfsdk:"authentication_method"`
	EncryptionAlgorithm     types.String `tfsdk:"encryption_algorithm"`
	AuthenticationAlgorithm types.String `tfsdk:"authentication_algorithm"`
	Group                   types.Int64  `tfsdk:"group"`
	SaLifetime              types.Int64  `tfsdk:"sa_lifetime"`
	Description             types.String `tfsdk:"description"`
	Comments                types.String `tfsdk:"comments"`
	OwnerID                 types.Int64  `tfsdk:"owner_id"`
	Created                 types.String `tfsdk:"created"`
	LastUpdated             types.String `tfsdk:"last_updated"`
	URL                     types.String `tfsdk:"url"`
	Tags                    types.Set    `tfsdk:"tags"`
	TagsAll                 types.Set    `tfsdk:"tags_all"`
	CustomFields            types.Map    `tfsdk:"custom_fields"`
}

// expandIkeProposal converts the Terraform model of type "ike_proposal" into an API request.
func expandIkeProposal(ctx context.Context, model *ikeProposalResourceModel) (*netboxapi.IkeProposalRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.IkeProposalRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.AuthenticationMethod = conv.StringPtr(model.AuthenticationMethod)
	requestDTO.EncryptionAlgorithm = conv.StringPtr(model.EncryptionAlgorithm)
	requestDTO.AuthenticationAlgorithm = conv.StringPtr(model.AuthenticationAlgorithm)
	requestDTO.Group = conv.Int64Ptr(model.Group)
	requestDTO.SaLifetime = conv.Int64Ptr(model.SaLifetime)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenIkeProposal refreshes the Terraform model of type "ike_proposal" from an API response.
func flattenIkeProposal(ctx context.Context, responseDTO *netboxapi.IkeProposalResponseDTO, model *ikeProposalResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.AuthenticationMethod = conv.FromStringPtr(responseDTO.AuthenticationMethod)
	model.EncryptionAlgorithm = conv.FromStringPtr(responseDTO.EncryptionAlgorithm)
	model.AuthenticationAlgorithm = conv.FromStringPtr(responseDTO.AuthenticationAlgorithm)
	model.Group = conv.FromInt64Ptr(responseDTO.Group)
	model.SaLifetime = conv.FromInt64Ptr(responseDTO.SaLifetime)
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
