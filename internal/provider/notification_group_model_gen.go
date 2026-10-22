// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// notificationGroupResourceModel is the Terraform state/plan model of the "notification_group" resource.
type notificationGroupResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	GroupIds    types.Set    `tfsdk:"group_ids"`
	UserIds     types.Set    `tfsdk:"user_ids"`
	URL         types.String `tfsdk:"url"`
}

// expandNotificationGroup converts the Terraform model of type "notification_group" into an API request.
func expandNotificationGroup(ctx context.Context, model *notificationGroupResourceModel) (*netboxapi.NotificationGroupRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.NotificationGroupRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Groups = conv.SetTo[int64](ctx, model.GroupIds, &diags)
	requestDTO.Users = conv.SetTo[int64](ctx, model.UserIds, &diags)
	return requestDTO, diags
}

// flattenNotificationGroup refreshes the Terraform model of type "notification_group" from an API response.
func flattenNotificationGroup(ctx context.Context, responseDTO *netboxapi.NotificationGroupResponseDTO, model *notificationGroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.GroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Groups, model.GroupIds.IsNull() || model.GroupIds.IsUnknown(), &diags)
	model.UserIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Users, model.UserIds.IsNull() || model.UserIds.IsUnknown(), &diags)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
