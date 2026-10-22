// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// permissionResourceModel is the Terraform state/plan model of the "permission" resource.
type permissionResourceModel struct {
	ID          types.Int64          `tfsdk:"id"`
	Name        types.String         `tfsdk:"name"`
	Description types.String         `tfsdk:"description"`
	Enabled     types.Bool           `tfsdk:"enabled"`
	ObjectTypes types.Set            `tfsdk:"object_types"`
	Actions     types.Set            `tfsdk:"actions"`
	GroupIds    types.Set            `tfsdk:"group_ids"`
	Groups      types.Set            `tfsdk:"groups"`
	UserIds     types.Set            `tfsdk:"user_ids"`
	Users       types.Set            `tfsdk:"users"`
	Constraints jsontypes.Normalized `tfsdk:"constraints"`
	URL         types.String         `tfsdk:"url"`
}

// expandPermission converts the Terraform model of type "permission" into an API request.
func expandPermission(ctx context.Context, model *permissionResourceModel) (*netboxapi.PermissionRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.PermissionRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.ObjectTypes = conv.SetTo[string](ctx, model.ObjectTypes, &diags)
	requestDTO.Actions = conv.SetTo[string](ctx, model.Actions, &diags)
	requestDTO.Groups = conv.SetTo[int64](ctx, model.GroupIds, &diags)
	requestDTO.Users = conv.SetTo[int64](ctx, model.UserIds, &diags)
	requestDTO.Constraints = conv.StringPtr(model.Constraints.StringValue)
	return requestDTO, diags
}

// flattenPermission refreshes the Terraform model of type "permission" from an API response.
func flattenPermission(ctx context.Context, responseDTO *netboxapi.PermissionResponseDTO, model *permissionResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.ObjectTypes = conv.SetFrom(ctx, types.StringType, responseDTO.ObjectTypes, false, &diags)
	model.Actions = conv.SetFrom(ctx, types.StringType, responseDTO.Actions, false, &diags)
	model.GroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Groups, model.GroupIds.IsNull() || model.GroupIds.IsUnknown(), &diags)
	model.UserIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Users, model.UserIds.IsNull() || model.UserIds.IsUnknown(), &diags)
	model.Constraints = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.Constraints))
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Groups = model.GroupIds
	model.Users = model.UserIds
	return diags
}
