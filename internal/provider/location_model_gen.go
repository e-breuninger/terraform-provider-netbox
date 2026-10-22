// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// locationResourceModel is the Terraform state/plan model of the "location" resource.
type locationResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Slug         types.String `tfsdk:"slug"`
	SiteID       types.Int64  `tfsdk:"site_id"`
	ParentID     types.Int64  `tfsdk:"parent_id"`
	Status       types.String `tfsdk:"status"`
	TenantID     types.Int64  `tfsdk:"tenant_id"`
	Facility     types.String `tfsdk:"facility"`
	Description  types.String `tfsdk:"description"`
	Comments     types.String `tfsdk:"comments"`
	Depth        types.Int64  `tfsdk:"depth"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	RackCount    types.Int64  `tfsdk:"rack_count"`
	DeviceCount  types.Int64  `tfsdk:"device_count"`
	PrefixCount  types.Int64  `tfsdk:"prefix_count"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandLocation converts the Terraform model of type "location" into an API request.
func expandLocation(ctx context.Context, model *locationResourceModel) (*netboxapi.LocationRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.LocationRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Site = conv.Int64Ptr(model.SiteID)
	requestDTO.Parent = conv.Int64Ptr(model.ParentID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Facility = conv.StringPtr(model.Facility)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenLocation refreshes the Terraform model of type "location" from an API response.
func flattenLocation(ctx context.Context, responseDTO *netboxapi.LocationResponseDTO, model *locationResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.SiteID = conv.FromInt64Ptr(responseDTO.Site)
	model.ParentID = conv.FromInt64Ptr(responseDTO.Parent)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.Facility = conv.FromStringPtr(responseDTO.Facility)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.Depth = conv.FromInt64Ptr(responseDTO.Depth)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.RackCount = conv.FromInt64Ptr(responseDTO.RackCount)
	model.DeviceCount = conv.FromInt64Ptr(responseDTO.DeviceCount)
	model.PrefixCount = conv.FromInt64Ptr(responseDTO.PrefixCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
