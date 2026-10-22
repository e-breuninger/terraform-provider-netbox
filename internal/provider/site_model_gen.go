// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// siteResourceModel is the Terraform state/plan model of the "site" resource.
type siteResourceModel struct {
	ID                  types.Int64   `tfsdk:"id"`
	Name                types.String  `tfsdk:"name"`
	Slug                types.String  `tfsdk:"slug"`
	Status              types.String  `tfsdk:"status"`
	Description         types.String  `tfsdk:"description"`
	Facility            types.String  `tfsdk:"facility"`
	PhysicalAddress     types.String  `tfsdk:"physical_address"`
	ShippingAddress     types.String  `tfsdk:"shipping_address"`
	Comments            types.String  `tfsdk:"comments"`
	Timezone            types.String  `tfsdk:"timezone"`
	Latitude            types.Float64 `tfsdk:"latitude"`
	Longitude           types.Float64 `tfsdk:"longitude"`
	RegionID            types.Int64   `tfsdk:"region_id"`
	TenantID            types.Int64   `tfsdk:"tenant_id"`
	GroupID             types.Int64   `tfsdk:"group_id"`
	AsnIds              types.Set     `tfsdk:"asn_ids"`
	OwnerID             types.Int64   `tfsdk:"owner_id"`
	Created             types.String  `tfsdk:"created"`
	LastUpdated         types.String  `tfsdk:"last_updated"`
	URL                 types.String  `tfsdk:"url"`
	CircuitCount        types.Int64   `tfsdk:"circuit_count"`
	DeviceCount         types.Int64   `tfsdk:"device_count"`
	PrefixCount         types.Int64   `tfsdk:"prefix_count"`
	RackCount           types.Int64   `tfsdk:"rack_count"`
	VirtualMachineCount types.Int64   `tfsdk:"virtual_machine_count"`
	VlanCount           types.Int64   `tfsdk:"vlan_count"`
	Tags                types.Set     `tfsdk:"tags"`
	TagsAll             types.Set     `tfsdk:"tags_all"`
	CustomFields        types.Map     `tfsdk:"custom_fields"`
}

// expandSite converts the Terraform model of type "site" into an API request.
func expandSite(ctx context.Context, model *siteResourceModel) (*netboxapi.SiteRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.SiteRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Facility = conv.StringPtr(model.Facility)
	requestDTO.PhysicalAddress = conv.StringPtr(model.PhysicalAddress)
	requestDTO.ShippingAddress = conv.StringPtr(model.ShippingAddress)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.TimeZone = conv.StringPtr(model.Timezone)
	requestDTO.Latitude = conv.Float64Ptr(model.Latitude)
	requestDTO.Longitude = conv.Float64Ptr(model.Longitude)
	requestDTO.Region = conv.Int64Ptr(model.RegionID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Group = conv.Int64Ptr(model.GroupID)
	requestDTO.Asns = conv.SetTo[int64](ctx, model.AsnIds, &diags)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenSite refreshes the Terraform model of type "site" from an API response.
func flattenSite(ctx context.Context, responseDTO *netboxapi.SiteResponseDTO, model *siteResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Facility = conv.FromStringPtr(responseDTO.Facility)
	model.PhysicalAddress = conv.FromStringPtr(responseDTO.PhysicalAddress)
	model.ShippingAddress = conv.FromStringPtr(responseDTO.ShippingAddress)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.Timezone = conv.FromStringPtr(responseDTO.TimeZone)
	model.Latitude = conv.FromFloat64Ptr(responseDTO.Latitude)
	model.Longitude = conv.FromFloat64Ptr(responseDTO.Longitude)
	model.RegionID = conv.FromInt64Ptr(responseDTO.Region)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.GroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.AsnIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Asns, model.AsnIds.IsNull() || model.AsnIds.IsUnknown(), &diags)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.CircuitCount = conv.FromInt64Ptr(responseDTO.CircuitCount)
	model.DeviceCount = conv.FromInt64Ptr(responseDTO.DeviceCount)
	model.PrefixCount = conv.FromInt64Ptr(responseDTO.PrefixCount)
	model.RackCount = conv.FromInt64Ptr(responseDTO.RackCount)
	model.VirtualMachineCount = conv.FromInt64Ptr(responseDTO.VirtualmachineCount)
	model.VlanCount = conv.FromInt64Ptr(responseDTO.VlanCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
