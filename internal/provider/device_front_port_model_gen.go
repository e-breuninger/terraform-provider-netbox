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

// deviceFrontPortResourceModel is the Terraform state/plan model of the "device_front_port" resource.
type deviceFrontPortResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	DeviceID         types.Int64  `tfsdk:"device_id"`
	ModuleID         types.Int64  `tfsdk:"module_id"`
	Name             types.String `tfsdk:"name"`
	Type             types.String `tfsdk:"type"`
	Positions        types.Int64  `tfsdk:"positions"`
	RearPorts        types.Set    `tfsdk:"rear_ports"`
	RearPortID       types.Int64  `tfsdk:"rear_port_id"`
	RearPortPosition types.Int64  `tfsdk:"rear_port_position"`
	ColorHex         types.String `tfsdk:"color_hex"`
	Label            types.String `tfsdk:"label"`
	MarkConnected    types.Bool   `tfsdk:"mark_connected"`
	Description      types.String `tfsdk:"description"`
	OwnerID          types.Int64  `tfsdk:"owner_id"`
	Created          types.String `tfsdk:"created"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	URL              types.String `tfsdk:"url"`
	Tags             types.Set    `tfsdk:"tags"`
	TagsAll          types.Set    `tfsdk:"tags_all"`
	CustomFields     types.Map    `tfsdk:"custom_fields"`
}

// expandDeviceFrontPort converts the Terraform model of type "device_front_port" into an API request.
func expandDeviceFrontPort(ctx context.Context, model *deviceFrontPortResourceModel) (*netboxapi.DeviceFrontPortRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceFrontPortRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.Module = conv.Int64Ptr(model.ModuleID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Positions = conv.Int64Ptr(model.Positions)
	if !model.RearPorts.IsNull() && !model.RearPorts.IsUnknown() {
		var items []deviceFrontPortRearPortsModel
		diags.Append(model.RearPorts.ElementsAs(ctx, &items, false)...)
		requestDTO.RearPorts = make([]*netboxapi.DeviceFrontPortRequestDTORearPorts, 0, len(items))
		for i := range items {
			v, d := expandDeviceFrontPortRearPorts(ctx, &items[i])
			diags.Append(d...)
			requestDTO.RearPorts = append(requestDTO.RearPorts, v)
		}
	}
	requestDTO.RearPortID = conv.Int64Ptr(model.RearPortID)
	requestDTO.RearPortPosition = conv.Int64Ptr(model.RearPortPosition)
	requestDTO.Color = conv.StringPtr(model.ColorHex)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.MarkConnected = conv.BoolPtr(model.MarkConnected)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDeviceFrontPort refreshes the Terraform model of type "device_front_port" from an API response.
func flattenDeviceFrontPort(ctx context.Context, responseDTO *netboxapi.DeviceFrontPortResponseDTO, model *deviceFrontPortResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.ModuleID = conv.FromInt64Ptr(responseDTO.Module)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Positions = conv.FromInt64Ptr(responseDTO.Positions)
	if responseDTO.RearPorts == nil {
		model.RearPorts = types.SetNull(types.ObjectType{AttrTypes: deviceFrontPortRearPortsAttrTypes()})
	} else {
		items := make([]deviceFrontPortRearPortsModel, 0, len(responseDTO.RearPorts))
		for _, element := range responseDTO.RearPorts {
			var o deviceFrontPortRearPortsModel
			diags.Append(flattenDeviceFrontPortRearPorts(ctx, element, &o)...)
			items = append(items, o)
		}
		v, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: deviceFrontPortRearPortsAttrTypes()}, items)
		diags.Append(d...)
		model.RearPorts = v
	}
	model.RearPortID = conv.FromInt64Ptr(responseDTO.RearPortID)
	model.RearPortPosition = conv.FromInt64Ptr(responseDTO.RearPortPosition)
	model.ColorHex = conv.FromStringPtr(responseDTO.Color)
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

// deviceFrontPortRearPortsModel is the model of the rear_ports object.
type deviceFrontPortRearPortsModel struct {
	Position         types.Int64 `tfsdk:"position"`
	RearPortID       types.Int64 `tfsdk:"rear_port_id"`
	RearPortPosition types.Int64 `tfsdk:"rear_port_position"`
}

// deviceFrontPortRearPortsAttrTypes is the attribute type map of deviceFrontPortRearPortsModel.
func deviceFrontPortRearPortsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"position":           types.Int64Type,
		"rear_port_id":       types.Int64Type,
		"rear_port_position": types.Int64Type,
	}
}

func expandDeviceFrontPortRearPorts(ctx context.Context, model *deviceFrontPortRearPortsModel) (*netboxapi.DeviceFrontPortRequestDTORearPorts, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceFrontPortRequestDTORearPorts{}
	requestDTO.Position = conv.Int64Ptr(model.Position)
	requestDTO.RearPort = conv.Int64Ptr(model.RearPortID)
	requestDTO.RearPortPosition = conv.Int64Ptr(model.RearPortPosition)
	return requestDTO, diags
}

func flattenDeviceFrontPortRearPorts(ctx context.Context, responseDTO *netboxapi.DeviceFrontPortResponseDTORearPorts, model *deviceFrontPortRearPortsModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.DeviceFrontPortResponseDTORearPorts{}
	}
	model.Position = conv.FromInt64Ptr(responseDTO.Position)
	model.RearPortID = conv.FromInt64Ptr(responseDTO.RearPort)
	model.RearPortPosition = conv.FromInt64Ptr(responseDTO.RearPortPosition)
	return diags
}
