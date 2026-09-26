// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// circuitGroupAssignmentResourceModel is the Terraform state/plan model of the "circuit_group_assignment" resource.
type circuitGroupAssignmentResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	CircuitGroupID types.Int64  `tfsdk:"circuit_group_id"`
	MemberType     types.String `tfsdk:"member_type"`
	MemberID       types.Int64  `tfsdk:"member_id"`
	Priority       types.String `tfsdk:"priority"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
	Tags           types.Set    `tfsdk:"tags"`
	TagsAll        types.Set    `tfsdk:"tags_all"`
}

// expandCircuitGroupAssignment converts the Terraform model of type "circuit_group_assignment" into an API request.
func expandCircuitGroupAssignment(ctx context.Context, model *circuitGroupAssignmentResourceModel) (*netboxapi.CircuitGroupAssignmentRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CircuitGroupAssignmentRequestDTO{}
	requestDTO.Group = conv.Int64Ptr(model.CircuitGroupID)
	requestDTO.MemberType = conv.StringPtr(model.MemberType)
	requestDTO.MemberID = conv.Int64Ptr(model.MemberID)
	requestDTO.Priority = conv.StringPtr(model.Priority)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	return requestDTO, diags
}

// flattenCircuitGroupAssignment refreshes the Terraform model of type "circuit_group_assignment" from an API response.
func flattenCircuitGroupAssignment(ctx context.Context, responseDTO *netboxapi.CircuitGroupAssignmentResponseDTO, model *circuitGroupAssignmentResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.CircuitGroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.MemberType = conv.FromStringPtr(responseDTO.MemberType)
	model.MemberID = conv.FromInt64Ptr(responseDTO.MemberID)
	model.Priority = conv.FromStringPtr(responseDTO.Priority)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	return diags
}
