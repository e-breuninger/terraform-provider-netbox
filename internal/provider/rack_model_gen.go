// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// rackResourceModel is the Terraform state/plan model of the "rack" resource.
type rackResourceModel struct {
	ID             types.Int64   `tfsdk:"id"`
	Name           types.String  `tfsdk:"name"`
	SiteID         types.Int64   `tfsdk:"site_id"`
	LocationID     types.Int64   `tfsdk:"location_id"`
	TenantID       types.Int64   `tfsdk:"tenant_id"`
	RoleID         types.Int64   `tfsdk:"role_id"`
	GroupID        types.Int64   `tfsdk:"group_id"`
	RackTypeID     types.Int64   `tfsdk:"rack_type_id"`
	Status         types.String  `tfsdk:"status"`
	FormFactor     types.String  `tfsdk:"form_factor"`
	Airflow        types.String  `tfsdk:"airflow"`
	Width          types.Int64   `tfsdk:"width"`
	UHeight        types.Int64   `tfsdk:"u_height"`
	StartingUnit   types.Int64   `tfsdk:"starting_unit"`
	DescUnits      types.Bool    `tfsdk:"desc_units"`
	Serial         types.String  `tfsdk:"serial"`
	AssetTag       types.String  `tfsdk:"asset_tag"`
	FacilityID     types.String  `tfsdk:"facility_id"`
	OuterWidth     types.Int64   `tfsdk:"outer_width"`
	OuterHeight    types.Int64   `tfsdk:"outer_height"`
	OuterDepth     types.Int64   `tfsdk:"outer_depth"`
	OuterUnit      types.String  `tfsdk:"outer_unit"`
	MountingDepth  types.Int64   `tfsdk:"mounting_depth"`
	Weight         types.Float64 `tfsdk:"weight"`
	MaxWeight      types.Int64   `tfsdk:"max_weight"`
	WeightUnit     types.String  `tfsdk:"weight_unit"`
	Description    types.String  `tfsdk:"description"`
	Comments       types.String  `tfsdk:"comments"`
	OwnerID        types.Int64   `tfsdk:"owner_id"`
	Created        types.String  `tfsdk:"created"`
	LastUpdated    types.String  `tfsdk:"last_updated"`
	URL            types.String  `tfsdk:"url"`
	DeviceCount    types.Int64   `tfsdk:"device_count"`
	PowerFeedCount types.Int64   `tfsdk:"power_feed_count"`
	Tags           types.Set     `tfsdk:"tags"`
	TagsAll        types.Set     `tfsdk:"tags_all"`
	CustomFields   types.Map     `tfsdk:"custom_fields"`
}

// expandRack converts the Terraform model of type "rack" into an API request.
func expandRack(ctx context.Context, model *rackResourceModel) (*netboxapi.RackRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.RackRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Site = conv.Int64Ptr(model.SiteID)
	requestDTO.Location = conv.Int64Ptr(model.LocationID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Group = conv.Int64Ptr(model.GroupID)
	requestDTO.RackType = conv.Int64Ptr(model.RackTypeID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.FormFactor = conv.StringPtr(model.FormFactor)
	requestDTO.Airflow = conv.StringPtr(model.Airflow)
	requestDTO.Width = conv.Int64Ptr(model.Width)
	requestDTO.UHeight = conv.Int64Ptr(model.UHeight)
	requestDTO.StartingUnit = conv.Int64Ptr(model.StartingUnit)
	requestDTO.DescUnits = conv.BoolPtr(model.DescUnits)
	requestDTO.Serial = conv.StringPtr(model.Serial)
	requestDTO.AssetTag = conv.StringPtr(model.AssetTag)
	requestDTO.FacilityID = conv.StringPtr(model.FacilityID)
	requestDTO.OuterWidth = conv.Int64Ptr(model.OuterWidth)
	requestDTO.OuterHeight = conv.Int64Ptr(model.OuterHeight)
	requestDTO.OuterDepth = conv.Int64Ptr(model.OuterDepth)
	requestDTO.OuterUnit = conv.StringPtr(model.OuterUnit)
	requestDTO.MountingDepth = conv.Int64Ptr(model.MountingDepth)
	requestDTO.Weight = conv.Float64Ptr(model.Weight)
	requestDTO.MaxWeight = conv.Int64Ptr(model.MaxWeight)
	requestDTO.WeightUnit = conv.StringPtr(model.WeightUnit)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenRack refreshes the Terraform model of type "rack" from an API response.
func flattenRack(ctx context.Context, responseDTO *netboxapi.RackResponseDTO, model *rackResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.SiteID = conv.FromInt64Ptr(responseDTO.Site)
	model.LocationID = conv.FromInt64Ptr(responseDTO.Location)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.GroupID = conv.FromInt64Ptr(responseDTO.Group)
	model.RackTypeID = conv.FromInt64Ptr(responseDTO.RackType)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.FormFactor = conv.FromStringPtr(responseDTO.FormFactor)
	model.Airflow = conv.FromStringPtr(responseDTO.Airflow)
	model.Width = conv.FromInt64Ptr(responseDTO.Width)
	model.UHeight = conv.FromInt64Ptr(responseDTO.UHeight)
	model.StartingUnit = conv.FromInt64Ptr(responseDTO.StartingUnit)
	model.DescUnits = conv.FromBoolPtr(responseDTO.DescUnits)
	model.Serial = conv.FromStringPtr(responseDTO.Serial)
	model.AssetTag = conv.FromStringPtr(responseDTO.AssetTag)
	model.FacilityID = conv.FromStringPtr(responseDTO.FacilityID)
	model.OuterWidth = conv.FromInt64Ptr(responseDTO.OuterWidth)
	model.OuterHeight = conv.FromInt64Ptr(responseDTO.OuterHeight)
	model.OuterDepth = conv.FromInt64Ptr(responseDTO.OuterDepth)
	model.OuterUnit = conv.FromStringPtr(responseDTO.OuterUnit)
	model.MountingDepth = conv.FromInt64Ptr(responseDTO.MountingDepth)
	model.Weight = conv.FromFloat64Ptr(responseDTO.Weight)
	model.MaxWeight = conv.FromInt64Ptr(responseDTO.MaxWeight)
	model.WeightUnit = conv.FromStringPtr(responseDTO.WeightUnit)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.DeviceCount = conv.FromInt64Ptr(responseDTO.DeviceCount)
	model.PowerFeedCount = conv.FromInt64Ptr(responseDTO.PowerfeedCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
