// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// virtualCircuitTerminationResourceModel is the Terraform state/plan model of the "virtual_circuit_termination" resource.
type virtualCircuitTerminationResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	VirtualCircuitID  types.Int64  `tfsdk:"virtual_circuit_id"`
	DeviceInterfaceID types.Int64  `tfsdk:"device_interface_id"`
	Role              types.String `tfsdk:"role"`
	Description       types.String `tfsdk:"description"`
	Created           types.String `tfsdk:"created"`
	LastUpdated       types.String `tfsdk:"last_updated"`
	URL               types.String `tfsdk:"url"`
	Tags              types.Set    `tfsdk:"tags"`
	TagsAll           types.Set    `tfsdk:"tags_all"`
	CustomFields      types.Map    `tfsdk:"custom_fields"`
}

// expandVirtualCircuitTermination converts the Terraform model of type "virtual_circuit_termination" into an API request.
func expandVirtualCircuitTermination(ctx context.Context, model *virtualCircuitTerminationResourceModel) (*netboxapi.VirtualCircuitTerminationRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualCircuitTerminationRequestDTO{}
	requestDTO.VirtualCircuit = conv.Int64Ptr(model.VirtualCircuitID)
	requestDTO.Interface = conv.Int64Ptr(model.DeviceInterfaceID)
	requestDTO.Role = conv.StringPtr(model.Role)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualCircuitTermination refreshes the Terraform model of type "virtual_circuit_termination" from an API response.
func flattenVirtualCircuitTermination(ctx context.Context, responseDTO *netboxapi.VirtualCircuitTerminationResponseDTO, model *virtualCircuitTerminationResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.VirtualCircuitID = conv.FromInt64Ptr(responseDTO.VirtualCircuit)
	model.DeviceInterfaceID = conv.FromInt64Ptr(responseDTO.Interface)
	model.Role = conv.FromStringPtr(responseDTO.Role)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
