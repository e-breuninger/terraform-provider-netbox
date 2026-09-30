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

// frontPortTemplateResourceModel is the Terraform state/plan model of the "front_port_template" resource.
type frontPortTemplateResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	DeviceTypeID     types.Int64  `tfsdk:"device_type_id"`
	ModuleTypeID     types.Int64  `tfsdk:"module_type_id"`
	Name             types.String `tfsdk:"name"`
	Type             types.String `tfsdk:"type"`
	Positions        types.Int64  `tfsdk:"positions"`
	RearPorts        types.Set    `tfsdk:"rear_ports"`
	RearPortID       types.Int64  `tfsdk:"rear_port_id"`
	RearPortPosition types.Int64  `tfsdk:"rear_port_position"`
	ColorHex         types.String `tfsdk:"color_hex"`
	Label            types.String `tfsdk:"label"`
	Description      types.String `tfsdk:"description"`
	Created          types.String `tfsdk:"created"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	URL              types.String `tfsdk:"url"`
}

// expandFrontPortTemplate converts the Terraform model of type "front_port_template" into an API request.
func expandFrontPortTemplate(ctx context.Context, model *frontPortTemplateResourceModel) (*netboxapi.FrontPortTemplateRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.FrontPortTemplateRequestDTO{}
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.ModuleType = conv.Int64Ptr(model.ModuleTypeID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Positions = conv.Int64Ptr(model.Positions)
	if !model.RearPorts.IsNull() && !model.RearPorts.IsUnknown() {
		var items []frontPortTemplateRearPortsModel
		diags.Append(model.RearPorts.ElementsAs(ctx, &items, false)...)
		requestDTO.RearPorts = make([]*netboxapi.FrontPortTemplateRequestDTORearPorts, 0, len(items))
		for i := range items {
			v, d := expandFrontPortTemplateRearPorts(ctx, &items[i])
			diags.Append(d...)
			requestDTO.RearPorts = append(requestDTO.RearPorts, v)
		}
	}
	requestDTO.RearPortID = conv.Int64Ptr(model.RearPortID)
	requestDTO.RearPortPosition = conv.Int64Ptr(model.RearPortPosition)
	requestDTO.Color = conv.StringPtr(model.ColorHex)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Description = conv.StringPtr(model.Description)
	return requestDTO, diags
}

// flattenFrontPortTemplate refreshes the Terraform model of type "front_port_template" from an API response.
func flattenFrontPortTemplate(ctx context.Context, responseDTO *netboxapi.FrontPortTemplateResponseDTO, model *frontPortTemplateResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.ModuleTypeID = conv.FromInt64Ptr(responseDTO.ModuleType)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Positions = conv.FromInt64Ptr(responseDTO.Positions)
	if responseDTO.RearPorts == nil {
		model.RearPorts = types.SetNull(types.ObjectType{AttrTypes: frontPortTemplateRearPortsAttrTypes()})
	} else {
		items := make([]frontPortTemplateRearPortsModel, 0, len(responseDTO.RearPorts))
		for _, element := range responseDTO.RearPorts {
			var o frontPortTemplateRearPortsModel
			diags.Append(flattenFrontPortTemplateRearPorts(ctx, element, &o)...)
			items = append(items, o)
		}
		v, d := types.SetValueFrom(ctx, types.ObjectType{AttrTypes: frontPortTemplateRearPortsAttrTypes()}, items)
		diags.Append(d...)
		model.RearPorts = v
	}
	model.RearPortID = conv.FromInt64Ptr(responseDTO.RearPortID)
	model.RearPortPosition = conv.FromInt64Ptr(responseDTO.RearPortPosition)
	model.ColorHex = conv.FromStringPtr(responseDTO.Color)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	return diags
}

// frontPortTemplateRearPortsModel is the model of the rear_ports object.
type frontPortTemplateRearPortsModel struct {
	Position         types.Int64 `tfsdk:"position"`
	RearPortID       types.Int64 `tfsdk:"rear_port_id"`
	RearPortPosition types.Int64 `tfsdk:"rear_port_position"`
}

// frontPortTemplateRearPortsAttrTypes is the attribute type map of frontPortTemplateRearPortsModel.
func frontPortTemplateRearPortsAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"position":           types.Int64Type,
		"rear_port_id":       types.Int64Type,
		"rear_port_position": types.Int64Type,
	}
}

func expandFrontPortTemplateRearPorts(ctx context.Context, model *frontPortTemplateRearPortsModel) (*netboxapi.FrontPortTemplateRequestDTORearPorts, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.FrontPortTemplateRequestDTORearPorts{}
	requestDTO.Position = conv.Int64Ptr(model.Position)
	requestDTO.RearPort = conv.Int64Ptr(model.RearPortID)
	requestDTO.RearPortPosition = conv.Int64Ptr(model.RearPortPosition)
	return requestDTO, diags
}

func flattenFrontPortTemplateRearPorts(ctx context.Context, responseDTO *netboxapi.FrontPortTemplateResponseDTORearPorts, model *frontPortTemplateRearPortsModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.FrontPortTemplateResponseDTORearPorts{}
	}
	model.Position = conv.FromInt64Ptr(responseDTO.Position)
	model.RearPortID = conv.FromInt64Ptr(responseDTO.RearPort)
	model.RearPortPosition = conv.FromInt64Ptr(responseDTO.RearPortPosition)
	return diags
}
