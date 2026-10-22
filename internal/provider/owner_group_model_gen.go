// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ownerGroupResourceModel is the Terraform state/plan model of the "owner_group" resource.
type ownerGroupResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	URL         types.String `tfsdk:"url"`
	MemberCount types.Int64  `tfsdk:"member_count"`
}

// expandOwnerGroup converts the Terraform model of type "owner_group" into an API request.
func expandOwnerGroup(ctx context.Context, model *ownerGroupResourceModel) (*netboxapi.OwnerGroupRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.OwnerGroupRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenOwnerGroup refreshes the Terraform model of type "owner_group" from an API response.
func flattenOwnerGroup(ctx context.Context, responseDTO *netboxapi.OwnerGroupResponseDTO, model *ownerGroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.MemberCount = conv.FromInt64Ptr(responseDTO.MemberCount)
	return diags
}
