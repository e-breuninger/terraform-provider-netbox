// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// powerFeedResourceModel is the Terraform state/plan model of the "power_feed" resource.
type powerFeedResourceModel struct {
	ID                    types.Int64  `tfsdk:"id"`
	Name                  types.String `tfsdk:"name"`
	PowerPanelID          types.Int64  `tfsdk:"power_panel_id"`
	RackID                types.Int64  `tfsdk:"rack_id"`
	Status                types.String `tfsdk:"status"`
	Type                  types.String `tfsdk:"type"`
	Supply                types.String `tfsdk:"supply"`
	Phase                 types.String `tfsdk:"phase"`
	Voltage               types.Int64  `tfsdk:"voltage"`
	Amperage              types.Int64  `tfsdk:"amperage"`
	MaxUtilizationPercent types.Int64  `tfsdk:"max_utilization_percent"`
	MaxPercentUtilization types.Int64  `tfsdk:"max_percent_utilization"`
	MarkConnected         types.Bool   `tfsdk:"mark_connected"`
	Description           types.String `tfsdk:"description"`
	Comments              types.String `tfsdk:"comments"`
	OwnerID               types.Int64  `tfsdk:"owner_id"`
	Created               types.String `tfsdk:"created"`
	LastUpdated           types.String `tfsdk:"last_updated"`
	URL                   types.String `tfsdk:"url"`
	Tags                  types.Set    `tfsdk:"tags"`
	TagsAll               types.Set    `tfsdk:"tags_all"`
	CustomFields          types.Map    `tfsdk:"custom_fields"`
}

// expandPowerFeed converts the Terraform model of type "power_feed" into an API request.
func expandPowerFeed(ctx context.Context, model *powerFeedResourceModel) (*netboxapi.PowerFeedRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.PowerFeedRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.PowerPanel = conv.Int64Ptr(model.PowerPanelID)
	requestDTO.Rack = conv.Int64Ptr(model.RackID)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Supply = conv.StringPtr(model.Supply)
	requestDTO.Phase = conv.StringPtr(model.Phase)
	requestDTO.Voltage = conv.Int64Ptr(model.Voltage)
	requestDTO.Amperage = conv.Int64Ptr(model.Amperage)
	requestDTO.MaxUtilization = conv.Int64Ptr(model.MaxUtilizationPercent)
	requestDTO.MarkConnected = conv.BoolPtr(model.MarkConnected)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenPowerFeed refreshes the Terraform model of type "power_feed" from an API response.
func flattenPowerFeed(ctx context.Context, responseDTO *netboxapi.PowerFeedResponseDTO, model *powerFeedResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.PowerPanelID = conv.FromInt64Ptr(responseDTO.PowerPanel)
	model.RackID = conv.FromInt64Ptr(responseDTO.Rack)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Supply = conv.FromStringPtr(responseDTO.Supply)
	model.Phase = conv.FromStringPtr(responseDTO.Phase)
	model.Voltage = conv.FromInt64Ptr(responseDTO.Voltage)
	model.Amperage = conv.FromInt64Ptr(responseDTO.Amperage)
	model.MaxUtilizationPercent = conv.FromInt64Ptr(responseDTO.MaxUtilization)
	model.MarkConnected = conv.FromBoolPtr(responseDTO.MarkConnected)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	model.MaxPercentUtilization = model.MaxUtilizationPercent
	return diags
}
