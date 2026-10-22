// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ownerResourceModel is the Terraform state/plan model of the "owner" resource.
type ownerResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	OwnerGroupID types.Int64  `tfsdk:"owner_group_id"`
	UserIds      types.Set    `tfsdk:"user_ids"`
	UserGroupIds types.Set    `tfsdk:"user_group_ids"`
	Description  types.String `tfsdk:"description"`
	URL          types.String `tfsdk:"url"`
}

// expandOwner converts the Terraform model of type "owner" into an API request.
func expandOwner(ctx context.Context, model *ownerResourceModel) (*netboxapi.OwnerRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.OwnerRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Group = conv.Int64Ptr(model.OwnerGroupID)
	requestDTO.Users = conv.SetTo[int64](ctx, model.UserIds, &diags)
	requestDTO.UserGroups = conv.SetTo[int64](ctx, model.UserGroupIds, &diags)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenOwner refreshes the Terraform model of type "owner" from an API response.
func flattenOwner(ctx context.Context, responseDTO *netboxapi.OwnerResponseDTO, model *ownerResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.OwnerGroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.UserIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Users, model.UserIds.IsNull() || model.UserIds.IsUnknown(), &diags)
	model.UserGroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.UserGroups, model.UserGroupIds.IsNull() || model.UserGroupIds.IsUnknown(), &diags)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
