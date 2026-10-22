// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// groupResourceModel is the Terraform state/plan model of the "group" resource.
type groupResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Description types.String `tfsdk:"description"`
	UserCount   types.Int64  `tfsdk:"user_count"`
	URL         types.String `tfsdk:"url"`
}

// expandGroup converts the Terraform model of type "group" into an API request.
func expandGroup(ctx context.Context, model *groupResourceModel) (*netboxapi.GroupRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.GroupRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenGroup refreshes the Terraform model of type "group" from an API response.
func flattenGroup(ctx context.Context, responseDTO *netboxapi.GroupResponseDTO, model *groupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.UserCount = conv.FromInt64Ptr(responseDTO.UserCount)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
