// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vlanTranslationRuleResourceModel is the Terraform state/plan model of the "vlan_translation_rule" resource.
type vlanTranslationRuleResourceModel struct {
	ID                      types.Int64  `tfsdk:"id"`
	VlanTranslationPolicyID types.Int64  `tfsdk:"vlan_translation_policy_id"`
	LocalVid                types.Int64  `tfsdk:"local_vid"`
	RemoteVid               types.Int64  `tfsdk:"remote_vid"`
	Description             types.String `tfsdk:"description"`
	URL                     types.String `tfsdk:"url"`
}

// expandVlanTranslationRule converts the Terraform model of type "vlan_translation_rule" into an API request.
func expandVlanTranslationRule(ctx context.Context, model *vlanTranslationRuleResourceModel) (*netboxapi.VlanTranslationRuleRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VlanTranslationRuleRequestDTO{}
	requestDTO.Policy = conv.Int64Ptr(model.VlanTranslationPolicyID)
	requestDTO.LocalVid = conv.Int64Ptr(model.LocalVid)
	requestDTO.RemoteVid = conv.Int64Ptr(model.RemoteVid)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenVlanTranslationRule refreshes the Terraform model of type "vlan_translation_rule" from an API response.
func flattenVlanTranslationRule(ctx context.Context, responseDTO *netboxapi.VlanTranslationRuleResponseDTO, model *vlanTranslationRuleResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.VlanTranslationPolicyID = conv.FromInt64Ptr(responseDTO.Policy)
	model.LocalVid = conv.FromInt64Ptr(responseDTO.LocalVid)
	model.RemoteVid = conv.FromInt64Ptr(responseDTO.RemoteVid)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
