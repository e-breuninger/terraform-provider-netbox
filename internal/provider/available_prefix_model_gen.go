// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// availablePrefixResourceModel is the Terraform state/plan model of the "available_prefix" resource.
type availablePrefixResourceModel struct {
	ID             types.Int64  `tfsdk:"id"`
	ParentPrefixID types.Int64  `tfsdk:"parent_prefix_id"`
	PrefixLength   types.Int64  `tfsdk:"prefix_length"`
	Prefix         types.String `tfsdk:"prefix"`
	Status         types.String `tfsdk:"status"`
	RoleID         types.Int64  `tfsdk:"role_id"`
	TenantID       types.Int64  `tfsdk:"tenant_id"`
	VlanID         types.Int64  `tfsdk:"vlan_id"`
	IsPool         types.Bool   `tfsdk:"is_pool"`
	MarkUtilized   types.Bool   `tfsdk:"mark_utilized"`
	Description    types.String `tfsdk:"description"`
	VrfID          types.Int64  `tfsdk:"vrf_id"`
	ScopeType      types.String `tfsdk:"scope_type"`
	ScopeID        types.Int64  `tfsdk:"scope_id"`
	SiteID         types.Int64  `tfsdk:"site_id"`
	LocationID     types.Int64  `tfsdk:"location_id"`
	RegionID       types.Int64  `tfsdk:"region_id"`
	SiteGroupID    types.Int64  `tfsdk:"site_group_id"`
	OwnerID        types.Int64  `tfsdk:"owner_id"`
	Created        types.String `tfsdk:"created"`
	LastUpdated    types.String `tfsdk:"last_updated"`
	URL            types.String `tfsdk:"url"`
	Tags           types.Set    `tfsdk:"tags"`
	TagsAll        types.Set    `tfsdk:"tags_all"`
	CustomFields   types.Map    `tfsdk:"custom_fields"`
}

// expandAvailablePrefix converts the Terraform model of type "available_prefix" into an API request.
func expandAvailablePrefix(ctx context.Context, model *availablePrefixResourceModel) (*netboxapi.AvailablePrefixRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.AvailablePrefixRequestDTO{}
	requestDTO.ParentPrefixID = conv.Int64Ptr(model.ParentPrefixID)
	requestDTO.PrefixLength = conv.Int64Ptr(model.PrefixLength)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Vlan = conv.Int64Ptr(model.VlanID)
	requestDTO.IsPool = conv.BoolPtr(model.IsPool)
	requestDTO.MarkUtilized = conv.BoolPtr(model.MarkUtilized)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Vrf = conv.Int64Ptr(model.VrfID)
	requestDTO.ScopeType = conv.StringPtr(model.ScopeType)
	requestDTO.ScopeID = conv.Int64Ptr(model.ScopeID)
	requestDTO.SiteID = conv.Int64Ptr(model.SiteID)
	requestDTO.LocationID = conv.Int64Ptr(model.LocationID)
	requestDTO.RegionID = conv.Int64Ptr(model.RegionID)
	requestDTO.SiteGroupID = conv.Int64Ptr(model.SiteGroupID)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenAvailablePrefix refreshes the Terraform model of type "available_prefix" from an API response.
func flattenAvailablePrefix(ctx context.Context, responseDTO *netboxapi.AvailablePrefixResponseDTO, model *availablePrefixResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Prefix = conv.FromStringPtr(responseDTO.Prefix)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.VlanID = conv.FromInt64Ptr(responseDTO.Vlan)
	model.IsPool = conv.FromBoolPtr(responseDTO.IsPool)
	model.MarkUtilized = conv.FromBoolPtr(responseDTO.MarkUtilized)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.VrfID = conv.FromInt64Ptr(responseDTO.Vrf)
	model.ScopeType = conv.FromStringPtr(responseDTO.ScopeType)
	model.ScopeID = conv.FromInt64Ptr(responseDTO.ScopeID)
	model.SiteID = conv.FromInt64Ptr(responseDTO.SiteID)
	model.LocationID = conv.FromInt64Ptr(responseDTO.LocationID)
	model.RegionID = conv.FromInt64Ptr(responseDTO.RegionID)
	model.SiteGroupID = conv.FromInt64Ptr(responseDTO.SiteGroupID)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
