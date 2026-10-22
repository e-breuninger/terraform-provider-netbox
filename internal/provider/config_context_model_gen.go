// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// configContextResourceModel is the Terraform state/plan model of the "config_context" resource.
type configContextResourceModel struct {
	ID              types.Int64          `tfsdk:"id"`
	Name            types.String         `tfsdk:"name"`
	Weight          types.Int64          `tfsdk:"weight"`
	Description     types.String         `tfsdk:"description"`
	IsActive        types.Bool           `tfsdk:"is_active"`
	Data            jsontypes.Normalized `tfsdk:"data"`
	ProfileID       types.Int64          `tfsdk:"profile_id"`
	RegionIds       types.Set            `tfsdk:"region_ids"`
	Regions         types.Set            `tfsdk:"regions"`
	SiteGroupIds    types.Set            `tfsdk:"site_group_ids"`
	SiteGroups      types.Set            `tfsdk:"site_groups"`
	SiteIds         types.Set            `tfsdk:"site_ids"`
	Sites           types.Set            `tfsdk:"sites"`
	LocationIds     types.Set            `tfsdk:"location_ids"`
	Locations       types.Set            `tfsdk:"locations"`
	DeviceTypeIds   types.Set            `tfsdk:"device_type_ids"`
	DeviceTypes     types.Set            `tfsdk:"device_types"`
	DeviceRoleIds   types.Set            `tfsdk:"device_role_ids"`
	Roles           types.Set            `tfsdk:"roles"`
	PlatformIds     types.Set            `tfsdk:"platform_ids"`
	Platforms       types.Set            `tfsdk:"platforms"`
	ClusterTypeIds  types.Set            `tfsdk:"cluster_type_ids"`
	ClusterTypes    types.Set            `tfsdk:"cluster_types"`
	ClusterGroupIds types.Set            `tfsdk:"cluster_group_ids"`
	ClusterGroups   types.Set            `tfsdk:"cluster_groups"`
	ClusterIds      types.Set            `tfsdk:"cluster_ids"`
	Clusters        types.Set            `tfsdk:"clusters"`
	TenantGroupIds  types.Set            `tfsdk:"tenant_group_ids"`
	TenantGroups    types.Set            `tfsdk:"tenant_groups"`
	TenantIds       types.Set            `tfsdk:"tenant_ids"`
	Tenants         types.Set            `tfsdk:"tenants"`
	TagSlugs        types.Set            `tfsdk:"tag_slugs"`
	OwnerID         types.Int64          `tfsdk:"owner_id"`
	Created         types.String         `tfsdk:"created"`
	LastUpdated     types.String         `tfsdk:"last_updated"`
	URL             types.String         `tfsdk:"url"`
}

// expandConfigContext converts the Terraform model of type "config_context" into an API request.
func expandConfigContext(ctx context.Context, model *configContextResourceModel) (*netboxapi.ConfigContextRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ConfigContextRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Weight = conv.Int64Ptr(model.Weight)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.IsActive = conv.BoolPtr(model.IsActive)
	requestDTO.Data = conv.StringPtr(model.Data.StringValue)
	requestDTO.Profile = conv.Int64Ptr(model.ProfileID)
	requestDTO.Regions = conv.SetTo[int64](ctx, model.RegionIds, &diags)
	requestDTO.SiteGroups = conv.SetTo[int64](ctx, model.SiteGroupIds, &diags)
	requestDTO.Sites = conv.SetTo[int64](ctx, model.SiteIds, &diags)
	requestDTO.Locations = conv.SetTo[int64](ctx, model.LocationIds, &diags)
	requestDTO.DeviceTypes = conv.SetTo[int64](ctx, model.DeviceTypeIds, &diags)
	requestDTO.Roles = conv.SetTo[int64](ctx, model.DeviceRoleIds, &diags)
	requestDTO.Platforms = conv.SetTo[int64](ctx, model.PlatformIds, &diags)
	requestDTO.ClusterTypes = conv.SetTo[int64](ctx, model.ClusterTypeIds, &diags)
	requestDTO.ClusterGroups = conv.SetTo[int64](ctx, model.ClusterGroupIds, &diags)
	requestDTO.Clusters = conv.SetTo[int64](ctx, model.ClusterIds, &diags)
	requestDTO.TenantGroups = conv.SetTo[int64](ctx, model.TenantGroupIds, &diags)
	requestDTO.Tenants = conv.SetTo[int64](ctx, model.TenantIds, &diags)
	requestDTO.Tags = conv.SetTo[string](ctx, model.TagSlugs, &diags)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	return requestDTO, diags
}

// flattenConfigContext refreshes the Terraform model of type "config_context" from an API response.
func flattenConfigContext(ctx context.Context, responseDTO *netboxapi.ConfigContextResponseDTO, model *configContextResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Weight = conv.FromInt64Ptr(responseDTO.Weight)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.IsActive = conv.FromBoolPtr(responseDTO.IsActive)
	model.Data = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.Data))
	model.ProfileID = conv.FromInt64Ptr(responseDTO.Profile)
	model.RegionIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Regions, model.RegionIds.IsNull() || model.RegionIds.IsUnknown(), &diags)
	model.SiteGroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.SiteGroups, model.SiteGroupIds.IsNull() || model.SiteGroupIds.IsUnknown(), &diags)
	model.SiteIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Sites, model.SiteIds.IsNull() || model.SiteIds.IsUnknown(), &diags)
	model.LocationIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Locations, model.LocationIds.IsNull() || model.LocationIds.IsUnknown(), &diags)
	model.DeviceTypeIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.DeviceTypes, model.DeviceTypeIds.IsNull() || model.DeviceTypeIds.IsUnknown(), &diags)
	model.DeviceRoleIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Roles, model.DeviceRoleIds.IsNull() || model.DeviceRoleIds.IsUnknown(), &diags)
	model.PlatformIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Platforms, model.PlatformIds.IsNull() || model.PlatformIds.IsUnknown(), &diags)
	model.ClusterTypeIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.ClusterTypes, model.ClusterTypeIds.IsNull() || model.ClusterTypeIds.IsUnknown(), &diags)
	model.ClusterGroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.ClusterGroups, model.ClusterGroupIds.IsNull() || model.ClusterGroupIds.IsUnknown(), &diags)
	model.ClusterIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Clusters, model.ClusterIds.IsNull() || model.ClusterIds.IsUnknown(), &diags)
	model.TenantGroupIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.TenantGroups, model.TenantGroupIds.IsNull() || model.TenantGroupIds.IsUnknown(), &diags)
	model.TenantIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Tenants, model.TenantIds.IsNull() || model.TenantIds.IsUnknown(), &diags)
	model.TagSlugs = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.TagSlugs.IsNull() || model.TagSlugs.IsUnknown(), &diags)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Regions = model.RegionIds
	model.SiteGroups = model.SiteGroupIds
	model.Sites = model.SiteIds
	model.Locations = model.LocationIds
	model.DeviceTypes = model.DeviceTypeIds
	model.Roles = model.DeviceRoleIds
	model.Platforms = model.PlatformIds
	model.ClusterTypes = model.ClusterTypeIds
	model.ClusterGroups = model.ClusterGroupIds
	model.Clusters = model.ClusterIds
	model.TenantGroups = model.TenantGroupIds
	model.Tenants = model.TenantIds
	return diags
}
