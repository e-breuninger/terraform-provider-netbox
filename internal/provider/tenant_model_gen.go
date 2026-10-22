// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// tenantResourceModel is the Terraform state/plan model of the "tenant" resource.
type tenantResourceModel struct {
	ID                  types.Int64  `tfsdk:"id"`
	Name                types.String `tfsdk:"name"`
	Slug                types.String `tfsdk:"slug"`
	Description         types.String `tfsdk:"description"`
	Comments            types.String `tfsdk:"comments"`
	GroupID             types.Int64  `tfsdk:"group_id"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	Created             types.String `tfsdk:"created"`
	LastUpdated         types.String `tfsdk:"last_updated"`
	URL                 types.String `tfsdk:"url"`
	CircuitCount        types.Int64  `tfsdk:"circuit_count"`
	DeviceCount         types.Int64  `tfsdk:"device_count"`
	IPAddressCount      types.Int64  `tfsdk:"ip_address_count"`
	PrefixCount         types.Int64  `tfsdk:"prefix_count"`
	RackCount           types.Int64  `tfsdk:"rack_count"`
	SiteCount           types.Int64  `tfsdk:"site_count"`
	VirtualMachineCount types.Int64  `tfsdk:"virtual_machine_count"`
	VlanCount           types.Int64  `tfsdk:"vlan_count"`
	VrfCount            types.Int64  `tfsdk:"vrf_count"`
	ClusterCount        types.Int64  `tfsdk:"cluster_count"`
	Tags                types.Set    `tfsdk:"tags"`
	TagsAll             types.Set    `tfsdk:"tags_all"`
	CustomFields        types.Map    `tfsdk:"custom_fields"`
}

// expandTenant converts the Terraform model of type "tenant" into an API request.
func expandTenant(ctx context.Context, model *tenantResourceModel) (*netboxapi.TenantRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.TenantRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Group = conv.Int64Ptr(model.GroupID)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenTenant refreshes the Terraform model of type "tenant" from an API response.
func flattenTenant(ctx context.Context, responseDTO *netboxapi.TenantResponseDTO, model *tenantResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.GroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.CircuitCount = conv.FromInt64Ptr(responseDTO.CircuitCount)
	model.DeviceCount = conv.FromInt64Ptr(responseDTO.DeviceCount)
	model.IPAddressCount = conv.FromInt64Ptr(responseDTO.IpaddressCount)
	model.PrefixCount = conv.FromInt64Ptr(responseDTO.PrefixCount)
	model.RackCount = conv.FromInt64Ptr(responseDTO.RackCount)
	model.SiteCount = conv.FromInt64Ptr(responseDTO.SiteCount)
	model.VirtualMachineCount = conv.FromInt64Ptr(responseDTO.VirtualmachineCount)
	model.VlanCount = conv.FromInt64Ptr(responseDTO.VlanCount)
	model.VrfCount = conv.FromInt64Ptr(responseDTO.VrfCount)
	model.ClusterCount = conv.FromInt64Ptr(responseDTO.ClusterCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
