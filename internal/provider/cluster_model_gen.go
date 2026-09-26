// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// clusterResourceModel is the Terraform state/plan model of the "cluster" resource.
type clusterResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	ClusterTypeID       types.Int64  `tfsdk:"cluster_type_id"`
	ClusterGroupID      types.Int64  `tfsdk:"cluster_group_id"`
	TenantID            types.Int64  `tfsdk:"tenant_id"`
	Status              types.String `tfsdk:"status"`
	Description         types.String `tfsdk:"description"`
	ScopeType           types.String `tfsdk:"scope_type"`
	ScopeID             types.Int64  `tfsdk:"scope_id"`
	SiteID              types.Int64  `tfsdk:"site_id"`
	LocationID          types.Int64  `tfsdk:"location_id"`
	RegionID            types.Int64  `tfsdk:"region_id"`
	SiteGroupID         types.Int64  `tfsdk:"site_group_id"`
	Comments            types.String `tfsdk:"comments"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	Created             types.String `tfsdk:"created"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	URL                 types.String `tfsdk:"url"`
	DeviceCount         types.Int64  `tfsdk:"device_count"`
	VirtualMachineCount types.Int64  `tfsdk:"virtual_machine_count"`
	Tags                types.Set    `tfsdk:"tags"`
	TagsAll             types.Set    `tfsdk:"tags_all"`
	CustomFields        types.Map    `tfsdk:"custom_fields"`
}

// expandCluster converts the Terraform model of type "cluster" into an API request.
func expandCluster(ctx context.Context, model *clusterResourceModel) (*netboxapi.ClusterRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ClusterRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.Int64Ptr(model.ClusterTypeID)
	requestDTO.Group = conv.Int64Ptr(model.ClusterGroupID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.ScopeType = conv.StringPtr(model.ScopeType)
	requestDTO.ScopeID = conv.Int64Ptr(model.ScopeID)
	requestDTO.SiteID = conv.Int64Ptr(model.SiteID)
	requestDTO.LocationID = conv.Int64Ptr(model.LocationID)
	requestDTO.RegionID = conv.Int64Ptr(model.RegionID)
	requestDTO.SiteGroupID = conv.Int64Ptr(model.SiteGroupID)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenCluster refreshes the Terraform model of type "cluster" from an API response.
func flattenCluster(ctx context.Context, responseDTO *netboxapi.ClusterResponseDTO, model *clusterResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.ClusterTypeID = conv.FromInt64Ptr(responseDTO.Type)
	model.ClusterGroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.ScopeType = conv.FromStringPtr(responseDTO.ScopeType)
	model.ScopeID = conv.FromInt64Ptr(responseDTO.ScopeID)
	model.SiteID = conv.FromInt64Ptr(responseDTO.SiteID)
	model.LocationID = conv.FromInt64Ptr(responseDTO.LocationID)
	model.RegionID = conv.FromInt64Ptr(responseDTO.RegionID)
	model.SiteGroupID = conv.FromInt64Ptr(responseDTO.SiteGroupID)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.DeviceCount = conv.FromInt64Ptr(responseDTO.DeviceCount)
	model.VirtualMachineCount = conv.FromInt64Ptr(responseDTO.VirtualmachineCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
