// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// vlanGroupResourceModel is the Terraform state/plan model of the "vlan_group" resource.
type vlanGroupResourceModel struct {
	ID           types.Int64  `tfsdk:"id"`
	Name         types.String `tfsdk:"name"`
	Slug         types.String `tfsdk:"slug"`
	Description  types.String `tfsdk:"description"`
	Comments     types.String `tfsdk:"comments"`
	VidRanges    types.List   `tfsdk:"vid_ranges"`
	TenantID     types.Int64  `tfsdk:"tenant_id"`
	ScopeType    types.String `tfsdk:"scope_type"`
	ScopeID      types.Int64  `tfsdk:"scope_id"`
	SiteID       types.Int64  `tfsdk:"site_id"`
	LocationID   types.Int64  `tfsdk:"location_id"`
	RegionID     types.Int64  `tfsdk:"region_id"`
	SiteGroupID  types.Int64  `tfsdk:"site_group_id"`
	OwnerID      types.Int64  `tfsdk:"owner_id"`
	Created      types.String `tfsdk:"created"`
	LastUpdated  types.String `tfsdk:"last_updated"`
	URL          types.String `tfsdk:"url"`
	VlanCount    types.Int64  `tfsdk:"vlan_count"`
	Utilization  types.String `tfsdk:"utilization"`
	Tags         types.Set    `tfsdk:"tags"`
	TagsAll      types.Set    `tfsdk:"tags_all"`
	CustomFields types.Map    `tfsdk:"custom_fields"`
}

// expandVlanGroup converts the Terraform model of type "vlan_group" into an API request.
func expandVlanGroup(ctx context.Context, model *vlanGroupResourceModel) (*netboxapi.VlanGroupRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VlanGroupRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	if !model.VidRanges.IsNull() && !model.VidRanges.IsUnknown() {
		var items []vlanGroupVidRangesModel
		diags.Append(model.VidRanges.ElementsAs(ctx, &items, false)...)
		requestDTO.VidRanges = make([]*netboxapi.VlanGroupRequestDTOVidRanges, 0, len(items))
		for i := range items {
			v, d := expandVlanGroupVidRanges(ctx, &items[i])
			diags.Append(d...)
			requestDTO.VidRanges = append(requestDTO.VidRanges, v)
		}
	}
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
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

// flattenVlanGroup refreshes the Terraform model of type "vlan_group" from an API response.
func flattenVlanGroup(ctx context.Context, responseDTO *netboxapi.VlanGroupResponseDTO, model *vlanGroupResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	if responseDTO.VidRanges == nil {
		model.VidRanges = types.ListNull(types.ObjectType{AttrTypes: vlanGroupVidRangesAttrTypes()})
	} else {
		items := make([]vlanGroupVidRangesModel, 0, len(responseDTO.VidRanges))
		for _, element := range responseDTO.VidRanges {
			var o vlanGroupVidRangesModel
			diags.Append(flattenVlanGroupVidRanges(ctx, element, &o)...)
			items = append(items, o)
		}
		v, d := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: vlanGroupVidRangesAttrTypes()}, items)
		diags.Append(d...)
		model.VidRanges = v
	}
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
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
	model.VlanCount = conv.FromInt64Ptr(responseDTO.VlanCount)
	model.Utilization = conv.FromStringPtr(responseDTO.Utilization)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}

// vlanGroupVidRangesModel is the model of the vid_ranges object.
type vlanGroupVidRangesModel struct {
	Start types.Int64 `tfsdk:"start"`
	End   types.Int64 `tfsdk:"end"`
}

// vlanGroupVidRangesAttrTypes is the attribute type map of vlanGroupVidRangesModel.
func vlanGroupVidRangesAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"start": types.Int64Type,
		"end":   types.Int64Type,
	}
}

func expandVlanGroupVidRanges(ctx context.Context, model *vlanGroupVidRangesModel) (*netboxapi.VlanGroupRequestDTOVidRanges, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VlanGroupRequestDTOVidRanges{}
	requestDTO.Start = conv.Int64Ptr(model.Start)
	requestDTO.End = conv.Int64Ptr(model.End)
	return requestDTO, diags
}

func flattenVlanGroupVidRanges(ctx context.Context, responseDTO *netboxapi.VlanGroupResponseDTOVidRanges, model *vlanGroupVidRangesModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.VlanGroupResponseDTOVidRanges{}
	}
	model.Start = conv.FromInt64Ptr(responseDTO.Start)
	model.End = conv.FromInt64Ptr(responseDTO.End)
	return diags
}
