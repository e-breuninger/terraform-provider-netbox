// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// userResourceModel is the Terraform state/plan model of the "user" resource.
type userResourceModel struct {
	ID         types.Int64  `tfsdk:"id"`
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	FirstName  types.String `tfsdk:"first_name"`
	LastName   types.String `tfsdk:"last_name"`
	Email      types.String `tfsdk:"email"`
	IsActive   types.Bool   `tfsdk:"is_active"`
	Active     types.Bool   `tfsdk:"active"`
	GroupIds   types.Set    `tfsdk:"group_ids"`
	DateJoined types.String `tfsdk:"date_joined"`
	LastLogin  types.String `tfsdk:"last_login"`
	URL        types.String `tfsdk:"url"`
}

// expandUser converts the Terraform model of type "user" into an API request.
func expandUser(ctx context.Context, model *userResourceModel) (*netboxapi.UserRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.UserRequestDTO{}
	requestDTO.Username = conv.StringPtr(model.Username)
	requestDTO.Password = conv.StringPtr(model.Password)
	requestDTO.FirstName = conv.StringPtr(model.FirstName)
	requestDTO.LastName = conv.StringPtr(model.LastName)
	requestDTO.Email = conv.StringPtr(model.Email)
	requestDTO.IsActive = conv.BoolPtr(model.IsActive)
	requestDTO.Groups = conv.SetTo[int64](ctx, model.GroupIds, &diags)
	return requestDTO, diags
}

// flattenUser refreshes the Terraform model of type "user" from an API response.
func flattenUser(ctx context.Context, responseDTO *netboxapi.UserResponseDTO, model *userResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Username = conv.FromStringPtr(responseDTO.Username)
	model.FirstName = conv.FromStringPtr(responseDTO.FirstName)
	model.LastName = conv.FromStringPtr(responseDTO.LastName)
	model.Email = conv.FromStringPtr(responseDTO.Email)
	model.IsActive = conv.FromBoolPtr(responseDTO.IsActive)
	model.GroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Groups, model.GroupIds.IsNull() || model.GroupIds.IsUnknown(), &diags)
	model.DateJoined = conv.FromStringPtr(responseDTO.DateJoined)
	model.LastLogin = conv.FromStringPtr(responseDTO.LastLogin)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Active = model.IsActive
	return diags
}
