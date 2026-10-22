// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// asnRangeResourceModel is the Terraform state/plan model of the "asn_range" resource.
type asnRangeResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Slug         types.String `tfsdk:"slug"`
	RirID        types.Int64  `tfsdk:"rir_id"`
	Start        types.Int64  `tfsdk:"start"`
	End          types.Int64  `tfsdk:"end"`
	TenantID     types.Int64  `tfsdk:"tenant_id"`
	Description  types.String `tfsdk:"description"`
	Comments     types.String `tfsdk:"comments"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	AsnCount     types.Int64  `tfsdk:"asn_count"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandAsnRange converts the Terraform model of type "asn_range" into an API request.
func expandAsnRange(ctx context.Context, model *asnRangeResourceModel) (*netboxapi.AsnRangeRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.AsnRangeRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Rir = conv.Int64Ptr(model.RirID)
	requestDTO.Start = conv.Int64Ptr(model.Start)
	requestDTO.End = conv.Int64Ptr(model.End)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenAsnRange refreshes the Terraform model of type "asn_range" from an API response.
func flattenAsnRange(ctx context.Context, responseDTO *netboxapi.AsnRangeResponseDTO, model *asnRangeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.RirID = conv.FromInt64Ptr(responseDTO.Rir)
	model.Start = conv.FromInt64Ptr(responseDTO.Start)
	model.End = conv.FromInt64Ptr(responseDTO.End)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.AsnCount = conv.FromInt64Ptr(responseDTO.AsnCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
