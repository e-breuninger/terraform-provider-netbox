// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// fhrpGroupResourceModel is the Terraform state/plan model of the "fhrp_group" resource.
type fhrpGroupResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Protocol     types.String `tfsdk:"protocol"`
	GroupID      types.Int64  `tfsdk:"group_id"`
	AuthType     types.String `tfsdk:"auth_type"`
	AuthKey      types.String `tfsdk:"auth_key"`
	Description  types.String `tfsdk:"description"`
	Comments     types.String `tfsdk:"comments"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandFhrpGroup converts the Terraform model of type "fhrp_group" into an API request.
func expandFhrpGroup(ctx context.Context, model *fhrpGroupResourceModel) (*netboxapi.FhrpGroupRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.FhrpGroupRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Protocol = conv.StringPtr(model.Protocol)
	requestDTO.GroupID = conv.Int64Ptr(model.GroupID)
	requestDTO.AuthType = conv.StringPtr(model.AuthType)
	requestDTO.AuthKey = conv.StringPtr(model.AuthKey)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenFhrpGroup refreshes the Terraform model of type "fhrp_group" from an API response.
func flattenFhrpGroup(ctx context.Context, responseDTO *netboxapi.FhrpGroupResponseDTO, model *fhrpGroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Protocol = conv.FromStringPtr(responseDTO.Protocol)
	model.GroupID = conv.FromInt64Ptr(responseDTO.GroupID)
	model.AuthType = conv.FromStringPtr(responseDTO.AuthType)
	model.AuthKey = conv.FromStringPtr(responseDTO.AuthKey)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
