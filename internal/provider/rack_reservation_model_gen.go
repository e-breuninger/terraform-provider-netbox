// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rackReservationResourceModel is the Terraform state/plan model of the "rack_reservation" resource.
type rackReservationResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	RackID       types.Int64  `tfsdk:"rack_id"`
	Units        types.Set    `tfsdk:"units"`
	UserID       types.Int64  `tfsdk:"user_id"`
	TenantID     types.Int64  `tfsdk:"tenant_id"`
	Description  types.String `tfsdk:"description"`
	Comments     types.String `tfsdk:"comments"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	UnitCount    types.Int64  `tfsdk:"unit_count"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandRackReservation converts the Terraform model of type "rack_reservation" into an API request.
func expandRackReservation(ctx context.Context, model *rackReservationResourceModel) (*netboxapi.RackReservationRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.RackReservationRequestDTO{}
	requestDTO.Rack = conv.Int64Ptr(model.RackID)
	requestDTO.Units = conv.SetTo[int64](ctx, model.Units, &diags)
	requestDTO.User = conv.Int64Ptr(model.UserID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenRackReservation refreshes the Terraform model of type "rack_reservation" from an API response.
func flattenRackReservation(ctx context.Context, responseDTO *netboxapi.RackReservationResponseDTO, model *rackReservationResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.RackID = conv.FromInt64Ptr(responseDTO.Rack)
	model.Units = conv.SetFrom(ctx, types.Int64Type, responseDTO.Units, false, &diags)
	model.UserID = conv.FromInt64Ptr(responseDTO.User)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.UnitCount = conv.FromInt64Ptr(responseDTO.UnitCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
