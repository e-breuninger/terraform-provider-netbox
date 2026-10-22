// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// aggregateResourceModel is the Terraform state/plan model of the "aggregate" resource.
type aggregateResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Prefix       conv.CIDR    `tfsdk:"prefix"`
	RirID        types.Int64  `tfsdk:"rir_id"`
	TenantID     types.Int64  `tfsdk:"tenant_id"`
	DateAdded    types.String `tfsdk:"date_added"`
	Family       types.Int64  `tfsdk:"family"`
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

// expandAggregate converts the Terraform model of type "aggregate" into an API request.
func expandAggregate(ctx context.Context, model *aggregateResourceModel) (*netboxapi.AggregateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.AggregateRequestDTO{}
	requestDTO.Prefix = conv.StringPtr(model.Prefix.StringValue)
	requestDTO.Rir = conv.Int64Ptr(model.RirID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.DateAdded = conv.StringPtr(model.DateAdded)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenAggregate refreshes the Terraform model of type "aggregate" from an API response.
func flattenAggregate(ctx context.Context, responseDTO *netboxapi.AggregateResponseDTO, model *aggregateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Prefix = conv.NewCIDRPointerValue(conv.EmptyStringAsNil(responseDTO.Prefix))
	model.RirID = conv.FromInt64Ptr(responseDTO.Rir)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.DateAdded = conv.FromStringPtr(responseDTO.DateAdded)
	model.Family = conv.FromInt64Ptr(responseDTO.Family)
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
