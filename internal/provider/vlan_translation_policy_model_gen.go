// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vlanTranslationPolicyResourceModel is the Terraform state/plan model of the "vlan_translation_policy" resource.
type vlanTranslationPolicyResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	Comments    types.String `tfsdk:"comments"`
	OwnerID     types.Int64  `tfsdk:"owner_id"`
	URL         types.String `tfsdk:"url"`
}

// expandVlanTranslationPolicy converts the Terraform model of type "vlan_translation_policy" into an API request.
func expandVlanTranslationPolicy(ctx context.Context, model *vlanTranslationPolicyResourceModel) (*netboxapi.VlanTranslationPolicyRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VlanTranslationPolicyRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	return requestDTO, diags
}

// flattenVlanTranslationPolicy refreshes the Terraform model of type "vlan_translation_policy" from an API response.
func flattenVlanTranslationPolicy(ctx context.Context, responseDTO *netboxapi.VlanTranslationPolicyResponseDTO, model *vlanTranslationPolicyResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
