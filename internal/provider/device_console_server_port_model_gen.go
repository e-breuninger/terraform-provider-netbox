// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// deviceConsoleServerPortResourceModel is the Terraform state/plan model of the "device_console_server_port" resource.
type deviceConsoleServerPortResourceModel struct {
	ID            types.Int64  `tfsdk:"id"`
	DeviceID      types.Int64  `tfsdk:"device_id"`
	ModuleID      types.Int64  `tfsdk:"module_id"`
	Name          types.String `tfsdk:"name"`
	Type          types.String `tfsdk:"type"`
	Speed         types.Int64  `tfsdk:"speed"`
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

// expandDeviceConsoleServerPort converts the Terraform model of type "device_console_server_port" into an API request.
func expandDeviceConsoleServerPort(ctx context.Context, model *deviceConsoleServerPortResourceModel) (*netboxapi.DeviceConsoleServerPortRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceConsoleServerPortRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.Module = conv.Int64Ptr(model.ModuleID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Speed = conv.Int64Ptr(model.Speed)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.MarkConnected = conv.BoolPtr(model.MarkConnected)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDeviceConsoleServerPort refreshes the Terraform model of type "device_console_server_port" from an API response.
func flattenDeviceConsoleServerPort(ctx context.Context, responseDTO *netboxapi.DeviceConsoleServerPortResponseDTO, model *deviceConsoleServerPortResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.ModuleID = conv.FromInt64Ptr(responseDTO.Module)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Speed = conv.FromInt64Ptr(responseDTO.Speed)
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
