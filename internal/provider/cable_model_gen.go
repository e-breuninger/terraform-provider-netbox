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

// cableResourceModel is the Terraform state/plan model of the "cable" resource.
type cableResourceModel struct {
	ID           types.Int64   `tfsdk:"id"`
	ASide        types.Object  `tfsdk:"a_side"`
	BSide        types.Object  `tfsdk:"b_side"`
	Type         types.String  `tfsdk:"type"`
	Status       types.String  `tfsdk:"status"`
	Profile      types.String  `tfsdk:"profile"`
	TenantID     types.Int64   `tfsdk:"tenant_id"`
	BundleID     types.Int64   `tfsdk:"bundle_id"`
	Label        types.String  `tfsdk:"label"`
	ColorHex     types.String  `tfsdk:"color_hex"`
	Length       types.Float64 `tfsdk:"length"`
	LengthUnit   types.String  `tfsdk:"length_unit"`
	Description  types.String  `tfsdk:"description"`
	Comments     types.String  `tfsdk:"comments"`
	OwnerID      types.Int64   `tfsdk:"owner_id"`
	Created      types.String  `tfsdk:"created"`
	LastUpdated  types.String  `tfsdk:"last_updated"`
	URL          types.String  `tfsdk:"url"`
	Tags         types.Set     `tfsdk:"tags"`
	TagsAll      types.Set     `tfsdk:"tags_all"`
	CustomFields types.Map     `tfsdk:"custom_fields"`
}

// expandCable converts the Terraform model of type "cable" into an API request.
func expandCable(ctx context.Context, model *cableResourceModel) (*netboxapi.CableRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CableRequestDTO{}
	if !model.ASide.IsNull() && !model.ASide.IsUnknown() {
		var o cableASideModel
		diags.Append(model.ASide.As(ctx, &o, conv.ObjectAsOptions)...)
		v, d := expandCableASide(ctx, &o)
		diags.Append(d...)
		requestDTO.ATerminations = v
	}
	if !model.BSide.IsNull() && !model.BSide.IsUnknown() {
		var o cableBSideModel
		diags.Append(model.BSide.As(ctx, &o, conv.ObjectAsOptions)...)
		v, d := expandCableBSide(ctx, &o)
		diags.Append(d...)
		requestDTO.BTerminations = v
	}
	requestDTO.Type = conv.StringPtr(model.Type)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Profile = conv.StringPtr(model.Profile)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Bundle = conv.Int64Ptr(model.BundleID)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Color = conv.StringPtr(model.ColorHex)
	requestDTO.Length = conv.Float64Ptr(model.Length)
	requestDTO.LengthUnit = conv.StringPtr(model.LengthUnit)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenCable refreshes the Terraform model of type "cable" from an API response.
func flattenCable(ctx context.Context, responseDTO *netboxapi.CableResponseDTO, model *cableResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	if responseDTO.ATerminations == nil {
		model.ASide = types.ObjectNull(cableASideAttrTypes())
	} else {
		var o cableASideModel
		diags.Append(flattenCableASide(ctx, responseDTO.ATerminations, &o)...)
		v, d := types.ObjectValueFrom(ctx, cableASideAttrTypes(), o)
		diags.Append(d...)
		model.ASide = v
	}
	if responseDTO.BTerminations == nil {
		model.BSide = types.ObjectNull(cableBSideAttrTypes())
	} else {
		var o cableBSideModel
		diags.Append(flattenCableBSide(ctx, responseDTO.BTerminations, &o)...)
		v, d := types.ObjectValueFrom(ctx, cableBSideAttrTypes(), o)
		diags.Append(d...)
		model.BSide = v
	}
	model.Type = conv.FromStringPtr(responseDTO.Type)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Profile = conv.FromStringPtr(responseDTO.Profile)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.BundleID = conv.FromInt64Ptr(responseDTO.Bundle)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.ColorHex = conv.FromStringPtr(responseDTO.Color)
	model.Length = conv.FromFloat64Ptr(responseDTO.Length)
	model.LengthUnit = conv.FromStringPtr(responseDTO.LengthUnit)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}

// cableASideModel is the model of the a_side object.
type cableASideModel struct {
	ObjectType            types.String `tfsdk:"object_type"`
	Ids                   types.List   `tfsdk:"ids"`
	DeviceInterfaceIds    types.List   `tfsdk:"device_interface_ids"`
	FrontPortIds          types.List   `tfsdk:"front_port_ids"`
	RearPortIds           types.List   `tfsdk:"rear_port_ids"`
	ConsolePortIds        types.List   `tfsdk:"console_port_ids"`
	ConsoleServerPortIds  types.List   `tfsdk:"console_server_port_ids"`
	PowerPortIds          types.List   `tfsdk:"power_port_ids"`
	PowerOutletIds        types.List   `tfsdk:"power_outlet_ids"`
	PowerFeedIds          types.List   `tfsdk:"power_feed_ids"`
	CircuitTerminationIds types.List   `tfsdk:"circuit_termination_ids"`
}

// cableASideAttrTypes is the attribute type map of cableASideModel.
func cableASideAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"object_type":             types.StringType,
		"ids":                     types.ListType{ElemType: types.Int64Type},
		"device_interface_ids":    types.ListType{ElemType: types.Int64Type},
		"front_port_ids":          types.ListType{ElemType: types.Int64Type},
		"rear_port_ids":           types.ListType{ElemType: types.Int64Type},
		"console_port_ids":        types.ListType{ElemType: types.Int64Type},
		"console_server_port_ids": types.ListType{ElemType: types.Int64Type},
		"power_port_ids":          types.ListType{ElemType: types.Int64Type},
		"power_outlet_ids":        types.ListType{ElemType: types.Int64Type},
		"power_feed_ids":          types.ListType{ElemType: types.Int64Type},
		"circuit_termination_ids": types.ListType{ElemType: types.Int64Type},
	}
}

func expandCableASide(ctx context.Context, model *cableASideModel) (*netboxapi.CableRequestDTOATerminations, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CableRequestDTOATerminations{}
	requestDTO.ObjectType = conv.StringPtr(model.ObjectType)
	requestDTO.Ids = conv.ListTo[int64](ctx, model.Ids, &diags)
	requestDTO.DeviceInterfaceIds = conv.ListTo[int64](ctx, model.DeviceInterfaceIds, &diags)
	requestDTO.FrontPortIds = conv.ListTo[int64](ctx, model.FrontPortIds, &diags)
	requestDTO.RearPortIds = conv.ListTo[int64](ctx, model.RearPortIds, &diags)
	requestDTO.ConsolePortIds = conv.ListTo[int64](ctx, model.ConsolePortIds, &diags)
	requestDTO.ConsoleServerPortIds = conv.ListTo[int64](ctx, model.ConsoleServerPortIds, &diags)
	requestDTO.PowerPortIds = conv.ListTo[int64](ctx, model.PowerPortIds, &diags)
	requestDTO.PowerOutletIds = conv.ListTo[int64](ctx, model.PowerOutletIds, &diags)
	requestDTO.PowerFeedIds = conv.ListTo[int64](ctx, model.PowerFeedIds, &diags)
	requestDTO.CircuitTerminationIds = conv.ListTo[int64](ctx, model.CircuitTerminationIds, &diags)
	return requestDTO, diags
}

func flattenCableASide(ctx context.Context, responseDTO *netboxapi.CableResponseDTOATerminations, model *cableASideModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.CableResponseDTOATerminations{}
	}
	model.ObjectType = conv.FromStringPtr(responseDTO.ObjectType)
	model.Ids = conv.ListFrom(ctx, types.Int64Type, responseDTO.Ids, false, &diags)
	model.DeviceInterfaceIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.DeviceInterfaceIds, false, &diags)
	model.FrontPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.FrontPortIds, false, &diags)
	model.RearPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.RearPortIds, false, &diags)
	model.ConsolePortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.ConsolePortIds, false, &diags)
	model.ConsoleServerPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.ConsoleServerPortIds, false, &diags)
	model.PowerPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.PowerPortIds, false, &diags)
	model.PowerOutletIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.PowerOutletIds, false, &diags)
	model.PowerFeedIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.PowerFeedIds, false, &diags)
	model.CircuitTerminationIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.CircuitTerminationIds, false, &diags)
	return diags
}

// cableBSideModel is the model of the b_side object.
type cableBSideModel struct {
	ObjectType            types.String `tfsdk:"object_type"`
	Ids                   types.List   `tfsdk:"ids"`
	DeviceInterfaceIds    types.List   `tfsdk:"device_interface_ids"`
	FrontPortIds          types.List   `tfsdk:"front_port_ids"`
	RearPortIds           types.List   `tfsdk:"rear_port_ids"`
	ConsolePortIds        types.List   `tfsdk:"console_port_ids"`
	ConsoleServerPortIds  types.List   `tfsdk:"console_server_port_ids"`
	PowerPortIds          types.List   `tfsdk:"power_port_ids"`
	PowerOutletIds        types.List   `tfsdk:"power_outlet_ids"`
	PowerFeedIds          types.List   `tfsdk:"power_feed_ids"`
	CircuitTerminationIds types.List   `tfsdk:"circuit_termination_ids"`
}

// cableBSideAttrTypes is the attribute type map of cableBSideModel.
func cableBSideAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"object_type":             types.StringType,
		"ids":                     types.ListType{ElemType: types.Int64Type},
		"device_interface_ids":    types.ListType{ElemType: types.Int64Type},
		"front_port_ids":          types.ListType{ElemType: types.Int64Type},
		"rear_port_ids":           types.ListType{ElemType: types.Int64Type},
		"console_port_ids":        types.ListType{ElemType: types.Int64Type},
		"console_server_port_ids": types.ListType{ElemType: types.Int64Type},
		"power_port_ids":          types.ListType{ElemType: types.Int64Type},
		"power_outlet_ids":        types.ListType{ElemType: types.Int64Type},
		"power_feed_ids":          types.ListType{ElemType: types.Int64Type},
		"circuit_termination_ids": types.ListType{ElemType: types.Int64Type},
	}
}

func expandCableBSide(ctx context.Context, model *cableBSideModel) (*netboxapi.CableRequestDTOBTerminations, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.CableRequestDTOBTerminations{}
	requestDTO.ObjectType = conv.StringPtr(model.ObjectType)
	requestDTO.Ids = conv.ListTo[int64](ctx, model.Ids, &diags)
	requestDTO.DeviceInterfaceIds = conv.ListTo[int64](ctx, model.DeviceInterfaceIds, &diags)
	requestDTO.FrontPortIds = conv.ListTo[int64](ctx, model.FrontPortIds, &diags)
	requestDTO.RearPortIds = conv.ListTo[int64](ctx, model.RearPortIds, &diags)
	requestDTO.ConsolePortIds = conv.ListTo[int64](ctx, model.ConsolePortIds, &diags)
	requestDTO.ConsoleServerPortIds = conv.ListTo[int64](ctx, model.ConsoleServerPortIds, &diags)
	requestDTO.PowerPortIds = conv.ListTo[int64](ctx, model.PowerPortIds, &diags)
	requestDTO.PowerOutletIds = conv.ListTo[int64](ctx, model.PowerOutletIds, &diags)
	requestDTO.PowerFeedIds = conv.ListTo[int64](ctx, model.PowerFeedIds, &diags)
	requestDTO.CircuitTerminationIds = conv.ListTo[int64](ctx, model.CircuitTerminationIds, &diags)
	return requestDTO, diags
}

func flattenCableBSide(ctx context.Context, responseDTO *netboxapi.CableResponseDTOBTerminations, model *cableBSideModel) diag.Diagnostics {
	var diags diag.Diagnostics
	if responseDTO == nil {
		responseDTO = &netboxapi.CableResponseDTOBTerminations{}
	}
	model.ObjectType = conv.FromStringPtr(responseDTO.ObjectType)
	model.Ids = conv.ListFrom(ctx, types.Int64Type, responseDTO.Ids, false, &diags)
	model.DeviceInterfaceIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.DeviceInterfaceIds, false, &diags)
	model.FrontPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.FrontPortIds, false, &diags)
	model.RearPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.RearPortIds, false, &diags)
	model.ConsolePortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.ConsolePortIds, false, &diags)
	model.ConsoleServerPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.ConsoleServerPortIds, false, &diags)
	model.PowerPortIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.PowerPortIds, false, &diags)
	model.PowerOutletIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.PowerOutletIds, false, &diags)
	model.PowerFeedIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.PowerFeedIds, false, &diags)
	model.CircuitTerminationIds = conv.ListFrom(ctx, types.Int64Type, responseDTO.CircuitTerminationIds, false, &diags)
	return diags
}
