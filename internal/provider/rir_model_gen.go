// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rirResourceModel is the Terraform state/plan model of the "rir" resource.
type rirResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	Name           types.String `tfsdk:"name"`
	Slug           types.String `tfsdk:"slug"`
	Description    types.String `tfsdk:"description"`
	IsPrivate      types.Bool   `tfsdk:"is_private"`
	OwnerID        types.Int64  `tfsdk:"owner_id"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
	AggregateCount types.Int64  `tfsdk:"aggregate_count"`
	Tags           types.Set    `tfsdk:"tags"`
	TagsAll        types.Set    `tfsdk:"tags_all"`
	CustomFields   types.Map    `tfsdk:"custom_fields"`
}

// expandRir converts the Terraform model of type "rir" into an API request.
func expandRir(ctx context.Context, model *rirResourceModel) (*netboxapi.RirRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.RirRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.IsPrivate = conv.BoolPtr(model.IsPrivate)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenRir refreshes the Terraform model of type "rir" from an API response.
func flattenRir(ctx context.Context, responseDTO *netboxapi.RirResponseDTO, model *rirResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.IsPrivate = conv.FromBoolPtr(responseDTO.IsPrivate)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.AggregateCount = conv.FromInt64Ptr(responseDTO.AggregateCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
