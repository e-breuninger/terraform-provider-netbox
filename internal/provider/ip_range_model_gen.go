// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// ipRangeResourceModel is the Terraform state/plan model of the "ip_range" resource.
type ipRangeResourceModel struct {
	ID            types.Int64    `tfsdk:"id"`
	StartAddress  conv.IPAddress `tfsdk:"start_address"`
	EndAddress    conv.IPAddress `tfsdk:"end_address"`
	VrfID         types.Int64    `tfsdk:"vrf_id"`
	TenantID      types.Int64    `tfsdk:"tenant_id"`
	RoleID        types.Int64    `tfsdk:"role_id"`
	Status        types.String   `tfsdk:"status"`
	MarkUtilized  types.Bool     `tfsdk:"mark_utilized"`
	MarkPopulated types.Bool     `tfsdk:"mark_populated"`
	Size          types.Int64    `tfsdk:"size"`
	Family        types.Int64    `tfsdk:"family"`
	Description   types.String   `tfsdk:"description"`
	Comments      types.String   `tfsdk:"comments"`
	OwnerID       types.Int64    `tfsdk:"owner_id"`
	Created       types.String   `tfsdk:"created"`
	LastUpdated   types.String   `tfsdk:"last_updated"`
	URL           types.String   `tfsdk:"url"`
	Tags          types.Set      `tfsdk:"tags"`
	TagsAll       types.Set      `tfsdk:"tags_all"`
	CustomFields  types.Map      `tfsdk:"custom_fields"`
}

// expandIPRange converts the Terraform model of type "ip_range" into an API request.
func expandIPRange(ctx context.Context, model *ipRangeResourceModel) (*netboxapi.IPRangeRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.IPRangeRequestDTO{}
	requestDTO.StartAddress = conv.StringPtr(model.StartAddress.StringValue)
	requestDTO.EndAddress = conv.StringPtr(model.EndAddress.StringValue)
	requestDTO.Vrf = conv.Int64Ptr(model.VrfID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.MarkUtilized = conv.BoolPtr(model.MarkUtilized)
	requestDTO.MarkPopulated = conv.BoolPtr(model.MarkPopulated)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenIPRange refreshes the Terraform model of type "ip_range" from an API response.
func flattenIPRange(ctx context.Context, responseDTO *netboxapi.IPRangeResponseDTO, model *ipRangeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.StartAddress = conv.NewIPAddressPointerValue(conv.EmptyStringAsNil(responseDTO.StartAddress))
	model.EndAddress = conv.NewIPAddressPointerValue(conv.EmptyStringAsNil(responseDTO.EndAddress))
	model.VrfID = conv.FromInt64Ptr(responseDTO.Vrf)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.MarkUtilized = conv.FromBoolPtr(responseDTO.MarkUtilized)
	model.MarkPopulated = conv.FromBoolPtr(responseDTO.MarkPopulated)
	model.Size = conv.FromInt64Ptr(responseDTO.Size)
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
