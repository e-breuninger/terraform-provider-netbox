// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// asnResourceModel is the Terraform state/plan model of the "asn" resource.
type asnResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	Asn           types.Int64  `tfsdk:"asn"`
	RirID         types.Int64  `tfsdk:"rir_id"`
	TenantID      types.Int64  `tfsdk:"tenant_id"`
	RoleID        types.Int64  `tfsdk:"role_id"`
	SiteIds       types.Set    `tfsdk:"site_ids"`
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

// expandAsn converts the Terraform model of type "asn" into an API request.
func expandAsn(ctx context.Context, model *asnResourceModel) (*netboxapi.AsnRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.AsnRequestDTO{}
	requestDTO.Asn = conv.Int64Ptr(model.Asn)
	requestDTO.Rir = conv.Int64Ptr(model.RirID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Sites = conv.SetTo[int64](ctx, model.SiteIds, &diags)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenAsn refreshes the Terraform model of type "asn" from an API response.
func flattenAsn(ctx context.Context, responseDTO *netboxapi.AsnResponseDTO, model *asnResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Asn = conv.FromInt64Ptr(responseDTO.Asn)
	model.RirID = conv.FromInt64Ptr(responseDTO.Rir)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.SiteIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Sites, model.SiteIds.IsNull() || model.SiteIds.IsUnknown(), &diags)
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
