// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// fhrpGroupAssignmentResourceModel is the Terraform state/plan model of the "fhrp_group_assignment" resource.
type fhrpGroupAssignmentResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	GroupID       types.Int64  `tfsdk:"group_id"`
	InterfaceType types.String `tfsdk:"interface_type"`
	InterfaceID   types.Int64  `tfsdk:"interface_id"`
	Priority      types.Int64  `tfsdk:"priority"`
	Created       types.String `tfsdk:"created"`
	LastUpdated   types.String `tfsdk:"last_updated"`
	URL           types.String `tfsdk:"url"`
}

// expandFhrpGroupAssignment converts the Terraform model of type "fhrp_group_assignment" into an API request.
func expandFhrpGroupAssignment(ctx context.Context, model *fhrpGroupAssignmentResourceModel) (*netboxapi.FhrpGroupAssignmentRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.FhrpGroupAssignmentRequestDTO{}
	requestDTO.Group = conv.Int64Ptr(model.GroupID)
	requestDTO.InterfaceType = conv.StringPtr(model.InterfaceType)
	requestDTO.InterfaceID = conv.Int64Ptr(model.InterfaceID)
	requestDTO.Priority = conv.Int64Ptr(model.Priority)
	return requestDTO, diags
}

// flattenFhrpGroupAssignment refreshes the Terraform model of type "fhrp_group_assignment" from an API response.
func flattenFhrpGroupAssignment(ctx context.Context, responseDTO *netboxapi.FhrpGroupAssignmentResponseDTO, model *fhrpGroupAssignmentResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.GroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.InterfaceType = conv.FromStringPtr(responseDTO.InterfaceType)
	model.InterfaceID = conv.FromInt64Ptr(responseDTO.InterfaceID)
	model.Priority = conv.FromInt64Ptr(responseDTO.Priority)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
