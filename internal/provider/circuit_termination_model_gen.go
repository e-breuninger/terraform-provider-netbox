// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// circuitTerminationResourceModel is the Terraform state/plan model of the "circuit_termination" resource.
type circuitTerminationResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	CircuitID         types.Int64  `tfsdk:"circuit_id"`
	TermSide          types.String `tfsdk:"term_side"`
	TerminationType   types.String `tfsdk:"termination_type"`
	TerminationID     types.Int64  `tfsdk:"termination_id"`
	SiteID            types.Int64  `tfsdk:"site_id"`
	LocationID        types.Int64  `tfsdk:"location_id"`
	RegionID          types.Int64  `tfsdk:"region_id"`
	SiteGroupID       types.Int64  `tfsdk:"site_group_id"`
	ProviderNetworkID types.Int64  `tfsdk:"provider_network_id"`
	PortSpeedKbps     types.Int64  `tfsdk:"port_speed_kbps"`
	PortSpeed         types.Int64  `tfsdk:"port_speed"`
	UpstreamSpeedKbps types.Int64  `tfsdk:"upstream_speed_kbps"`
	UpstreamSpeed     types.Int64  `tfsdk:"upstream_speed"`
	XconnectID        types.String `tfsdk:"xconnect_id"`
	PpInfo            types.String `tfsdk:"pp_info"`
	MarkConnected     types.Bool   `tfsdk:"mark_connected"`
	Description       types.String `tfsdk:"description"`
	Created           types.String `tfsdk:"created"`
	LastUpdated       types.String `tfsdk:"last_updated"`
	URL               types.String `tfsdk:"url"`
	Tags              types.Set    `tfsdk:"tags"`
	TagsAll           types.Set    `tfsdk:"tags_all"`
	CustomFields      types.Map    `tfsdk:"custom_fields"`
}

// expandCircuitTermination converts the Terraform model of type "circuit_termination" into an API request.
func expandCircuitTermination(ctx context.Context, model *circuitTerminationResourceModel) (*netboxapi.CircuitTerminationRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CircuitTerminationRequestDTO{}
	requestDTO.Circuit = conv.Int64Ptr(model.CircuitID)
	requestDTO.TermSide = conv.StringPtr(model.TermSide)
	requestDTO.TerminationType = conv.StringPtr(model.TerminationType)
	requestDTO.TerminationID = conv.Int64Ptr(model.TerminationID)
	requestDTO.SiteID = conv.Int64Ptr(model.SiteID)
	requestDTO.LocationID = conv.Int64Ptr(model.LocationID)
	requestDTO.RegionID = conv.Int64Ptr(model.RegionID)
	requestDTO.SiteGroupID = conv.Int64Ptr(model.SiteGroupID)
	requestDTO.ProviderNetworkID = conv.Int64Ptr(model.ProviderNetworkID)
	requestDTO.PortSpeed = conv.Int64Ptr(model.PortSpeedKbps)
	requestDTO.UpstreamSpeed = conv.Int64Ptr(model.UpstreamSpeedKbps)
	requestDTO.XconnectID = conv.StringPtr(model.XconnectID)
	requestDTO.PpInfo = conv.StringPtr(model.PpInfo)
	requestDTO.MarkConnected = conv.BoolPtr(model.MarkConnected)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenCircuitTermination refreshes the Terraform model of type "circuit_termination" from an API response.
func flattenCircuitTermination(ctx context.Context, responseDTO *netboxapi.CircuitTerminationResponseDTO, model *circuitTerminationResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.CircuitID = conv.FromInt64Ptr(responseDTO.Circuit)
	model.TermSide = conv.FromStringPtr(responseDTO.TermSide)
	model.TerminationType = conv.FromStringPtr(responseDTO.TerminationType)
	model.TerminationID = conv.FromInt64Ptr(responseDTO.TerminationID)
	model.SiteID = conv.FromInt64Ptr(responseDTO.SiteID)
	model.LocationID = conv.FromInt64Ptr(responseDTO.LocationID)
	model.RegionID = conv.FromInt64Ptr(responseDTO.RegionID)
	model.SiteGroupID = conv.FromInt64Ptr(responseDTO.SiteGroupID)
	model.ProviderNetworkID = conv.FromInt64Ptr(responseDTO.ProviderNetworkID)
	model.PortSpeedKbps = conv.FromInt64Ptr(responseDTO.PortSpeed)
	model.UpstreamSpeedKbps = conv.FromInt64Ptr(responseDTO.UpstreamSpeed)
	model.XconnectID = conv.FromStringPtr(responseDTO.XconnectID)
	model.PpInfo = conv.FromStringPtr(responseDTO.PpInfo)
	model.MarkConnected = conv.FromBoolPtr(responseDTO.MarkConnected)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	model.PortSpeed = model.PortSpeedKbps
	model.UpstreamSpeed = model.UpstreamSpeedKbps
	return diags
}
