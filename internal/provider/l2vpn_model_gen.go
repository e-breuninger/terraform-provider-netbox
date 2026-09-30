// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// l2vpnResourceModel is the Terraform state/plan model of the "l2vpn" resource.
type l2vpnResourceModel struct {
	ID              types.Int64  `tfsdk:"id"`
	Name            types.String `tfsdk:"name"`
	Slug            types.String `tfsdk:"slug"`
	Type            types.String `tfsdk:"type"`
	Status          types.String `tfsdk:"status"`
	Identifier      types.Int64  `tfsdk:"identifier"`
	TenantID        types.Int64  `tfsdk:"tenant_id"`
	ImportTargetIds types.Set    `tfsdk:"import_target_ids"`
	ExportTargetIds types.Set    `tfsdk:"export_target_ids"`
	Description     types.String `tfsdk:"description"`
	Comments        types.String `tfsdk:"comments"`
	OwnerID         types.Int64  `tfsdk:"owner_id"`
	Created         types.String `tfsdk:"created"`
	LastUpdated     types.String `tfsdk:"last_updated"`
	URL             types.String `tfsdk:"url"`
	Tags            types.Set    `tfsdk:"tags"`
	TagsAll         types.Set    `tfsdk:"tags_all"`
	CustomFields    types.Map    `tfsdk:"custom_fields"`
}

// expandL2vpn converts the Terraform model of type "l2vpn" into an API request.
func expandL2vpn(ctx context.Context, model *l2vpnResourceModel) (*netboxapi.L2vpnRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.L2vpnRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Identifier = conv.Int64Ptr(model.Identifier)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.ImportTargets = conv.SetTo[int64](ctx, model.ImportTargetIds, &diags)
	requestDTO.ExportTargets = conv.SetTo[int64](ctx, model.ExportTargetIds, &diags)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenL2vpn refreshes the Terraform model of type "l2vpn" from an API response.
func flattenL2vpn(ctx context.Context, responseDTO *netboxapi.L2vpnResponseDTO, model *l2vpnResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Identifier = conv.FromInt64Ptr(responseDTO.Identifier)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.ImportTargetIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.ImportTargets, model.ImportTargetIds.IsNull() || model.ImportTargetIds.IsUnknown(), &diags)
	model.ExportTargetIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.ExportTargets, model.ExportTargetIds.IsNull() || model.ExportTargetIds.IsUnknown(), &diags)
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
