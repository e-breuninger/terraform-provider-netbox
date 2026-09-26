// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// availableAsnResourceModel is the Terraform state/plan model of the "available_asn" resource.
type availableAsnResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	AsnRangeID    types.Int64  `tfsdk:"asn_range_id"`
	Asn           types.Int64  `tfsdk:"asn"`
	RirID         types.Int64  `tfsdk:"rir_id"`
	TenantID      types.Int64  `tfsdk:"tenant_id"`
	RoleID        types.Int64  `tfsdk:"role_id"`
	Description   types.String `tfsdk:"description"`
	Comments      types.String `tfsdk:"comments"`
	OwnerID       types.Int64  `tfsdk:"owner_id"`
	Created       types.String `tfsdk:"created"`
	LastUpdated   types.String `tfsdk:"last_updated"`
	URL           types.String `tfsdk:"url"`
	SiteCount     types.Int64  `tfsdk:"site_count"`
	ProviderCount types.Int64  `tfsdk:"provider_count"`
	Tags          types.Set    `tfsdk:"tags"`
	TagsAll       types.Set    `tfsdk:"tags_all"`
	CustomFields  types.Map    `tfsdk:"custom_fields"`
}

// expandAvailableAsn converts the Terraform model of type "available_asn" into an API request.
func expandAvailableAsn(ctx context.Context, model *availableAsnResourceModel) (*netboxapi.AvailableAsnRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.AvailableAsnRequestDTO{}
	requestDTO.AsnRangeID = conv.Int64Ptr(model.AsnRangeID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenAvailableAsn refreshes the Terraform model of type "available_asn" from an API response.
func flattenAvailableAsn(ctx context.Context, responseDTO *netboxapi.AvailableAsnResponseDTO, model *availableAsnResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Asn = conv.FromInt64Ptr(responseDTO.Asn)
	model.RirID = conv.FromInt64Ptr(responseDTO.Rir)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.SiteCount = conv.FromInt64Ptr(responseDTO.SiteCount)
	model.ProviderCount = conv.FromInt64Ptr(responseDTO.ProviderCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
