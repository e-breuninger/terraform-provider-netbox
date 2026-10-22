// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// tagResourceModel is the Terraform state/plan model of the "tag" resource.
type tagResourceModel struct {
	ID          types.Int64  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Slug        types.String `tfsdk:"slug"`
	ColorHex    types.String `tfsdk:"color_hex"`
	Description types.String `tfsdk:"description"`
	Weight      types.Int64  `tfsdk:"weight"`
	ObjectTypes types.Set    `tfsdk:"object_types"`
	Created     types.String `tfsdk:"created"`
	LastUpdated types.String `tfsdk:"last_updated"`
	URL         types.String `tfsdk:"url"`
}

// expandTag converts the Terraform model of type "tag" into an API request.
func expandTag(ctx context.Context, model *tagResourceModel) (*netboxapi.TagRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.TagRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Color = conv.StringPtr(model.ColorHex)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Weight = conv.Int64Ptr(model.Weight)
	requestDTO.ObjectTypes = conv.SetTo[string](ctx, model.ObjectTypes, &diags)
	return requestDTO, diags
}

// flattenTag refreshes the Terraform model of type "tag" from an API response.
func flattenTag(ctx context.Context, responseDTO *netboxapi.TagResponseDTO, model *tagResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.ColorHex = conv.FromStringPtr(responseDTO.Color)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Weight = conv.FromInt64Ptr(responseDTO.Weight)
	model.ObjectTypes = conv.SetFrom(ctx, types.StringType, responseDTO.ObjectTypes, model.ObjectTypes.IsNull() || model.ObjectTypes.IsUnknown(), &diags)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}
