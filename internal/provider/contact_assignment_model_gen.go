// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// contactAssignmentResourceModel is the Terraform state/plan model of the "contact_assignment" resource.
type contactAssignmentResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	ObjectType   types.String `tfsdk:"object_type"`
	ObjectID     types.Int64  `tfsdk:"object_id"`
	ContactID    types.Int64  `tfsdk:"contact_id"`
	RoleID       types.Int64  `tfsdk:"role_id"`
	Priority     types.String `tfsdk:"priority"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandContactAssignment converts the Terraform model of type "contact_assignment" into an API request.
func expandContactAssignment(ctx context.Context, model *contactAssignmentResourceModel) (*netboxapi.ContactAssignmentRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ContactAssignmentRequestDTO{}
	requestDTO.ObjectType = conv.StringPtr(model.ObjectType)
	requestDTO.ObjectID = conv.Int64Ptr(model.ObjectID)
	requestDTO.Contact = conv.Int64Ptr(model.ContactID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Priority = conv.StringPtr(model.Priority)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenContactAssignment refreshes the Terraform model of type "contact_assignment" from an API response.
func flattenContactAssignment(ctx context.Context, responseDTO *netboxapi.ContactAssignmentResponseDTO, model *contactAssignmentResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.ObjectType = conv.FromStringPtr(responseDTO.ObjectType)
	model.ObjectID = conv.FromInt64Ptr(responseDTO.ObjectID)
	model.ContactID = conv.FromInt64Ptr(responseDTO.Contact)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.Priority = conv.FromStringPtr(responseDTO.Priority)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
