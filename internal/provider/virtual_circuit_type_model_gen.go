// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// virtualCircuitTypeResourceModel is the Terraform state/plan model of the "virtual_circuit_type" resource.
type virtualCircuitTypeResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Slug                types.String `tfsdk:"slug"`
	Description         types.String `tfsdk:"description"`
	Comments            types.String `tfsdk:"comments"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	Created             types.String `tfsdk:"created"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	URL                 types.String `tfsdk:"url"`
	VirtualCircuitCount types.Int64  `tfsdk:"virtual_circuit_count"`
	Tags                types.Set    `tfsdk:"tags"`
	TagsAll             types.Set    `tfsdk:"tags_all"`
	CustomFields        types.Map    `tfsdk:"custom_fields"`
}

// expandVirtualCircuitType converts the Terraform model of type "virtual_circuit_type" into an API request.
func expandVirtualCircuitType(ctx context.Context, model *virtualCircuitTypeResourceModel) (*netboxapi.VirtualCircuitTypeRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualCircuitTypeRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualCircuitType refreshes the Terraform model of type "virtual_circuit_type" from an API response.
func flattenVirtualCircuitType(ctx context.Context, responseDTO *netboxapi.VirtualCircuitTypeResponseDTO, model *virtualCircuitTypeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.VirtualCircuitCount = conv.FromInt64Ptr(responseDTO.VirtualCircuitCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
