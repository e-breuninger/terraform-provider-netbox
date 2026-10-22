// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// devicePowerPortResourceModel is the Terraform state/plan model of the "device_power_port" resource.
type devicePowerPortResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	DeviceID      types.Int64  `tfsdk:"device_id"`
	ModuleID      types.Int64  `tfsdk:"module_id"`
	Name          types.String `tfsdk:"name"`
	Type          types.String `tfsdk:"type"`
	MaximumDraw   types.Int64  `tfsdk:"maximum_draw"`
	AllocatedDraw types.Int64  `tfsdk:"allocated_draw"`
	Label         types.String `tfsdk:"label"`
	MarkConnected types.Bool   `tfsdk:"mark_connected"`
	Description   types.String `tfsdk:"description"`
	OwnerID       types.Int64  `tfsdk:"owner_id"`
	Created       types.String `tfsdk:"created"`
	LastUpdated   types.String `tfsdk:"last_updated"`
	URL           types.String `tfsdk:"url"`
	Tags          types.Set    `tfsdk:"tags"`
	TagsAll       types.Set    `tfsdk:"tags_all"`
	CustomFields  types.Map    `tfsdk:"custom_fields"`
}

// expandDevicePowerPort converts the Terraform model of type "device_power_port" into an API request.
func expandDevicePowerPort(ctx context.Context, model *devicePowerPortResourceModel) (*netboxapi.DevicePowerPortRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DevicePowerPortRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.Module = conv.Int64Ptr(model.ModuleID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.MaximumDraw = conv.Int64Ptr(model.MaximumDraw)
	requestDTO.AllocatedDraw = conv.Int64Ptr(model.AllocatedDraw)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.MarkConnected = conv.BoolPtr(model.MarkConnected)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDevicePowerPort refreshes the Terraform model of type "device_power_port" from an API response.
func flattenDevicePowerPort(ctx context.Context, responseDTO *netboxapi.DevicePowerPortResponseDTO, model *devicePowerPortResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.ModuleID = conv.FromInt64Ptr(responseDTO.Module)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.MaximumDraw = conv.FromInt64Ptr(responseDTO.MaximumDraw)
	model.AllocatedDraw = conv.FromInt64Ptr(responseDTO.AllocatedDraw)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.MarkConnected = conv.FromBoolPtr(responseDTO.MarkConnected)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
