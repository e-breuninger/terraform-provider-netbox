// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// serviceResourceModel is the Terraform state/plan model of the "service" resource.
type serviceResourceModel struct {
	ID               types.Int64  `tfsdk:"id"`
	Name             types.String `tfsdk:"name"`
	Protocol         types.String `tfsdk:"protocol"`
	Ports            types.Set    `tfsdk:"ports"`
	ParentObjectType types.String `tfsdk:"parent_object_type"`
	ParentObjectID   types.Int64  `tfsdk:"parent_object_id"`
	DeviceID         types.Int64  `tfsdk:"device_id"`
	VirtualMachineID types.Int64  `tfsdk:"virtual_machine_id"`
	IPAddressIds     types.Set    `tfsdk:"ip_address_ids"`
	Description      types.String `tfsdk:"description"`
	Comments         types.String `tfsdk:"comments"`
	OwnerID          types.Int64  `tfsdk:"owner_id"`
	Created          types.String `tfsdk:"created"`
	LastUpdated      types.String `tfsdk:"last_updated"`
	URL              types.String `tfsdk:"url"`
	Tags             types.Set    `tfsdk:"tags"`
	TagsAll          types.Set    `tfsdk:"tags_all"`
	CustomFields     types.Map    `tfsdk:"custom_fields"`
}

// expandService converts the Terraform model of type "service" into an API request.
func expandService(ctx context.Context, model *serviceResourceModel) (*netboxapi.ServiceRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ServiceRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Protocol = conv.StringPtr(model.Protocol)
	requestDTO.Ports = conv.SetTo[int64](ctx, model.Ports, &diags)
	requestDTO.ParentObjectType = conv.StringPtr(model.ParentObjectType)
	requestDTO.ParentObjectID = conv.Int64Ptr(model.ParentObjectID)
	requestDTO.DeviceID = conv.Int64Ptr(model.DeviceID)
	requestDTO.VirtualMachineID = conv.Int64Ptr(model.VirtualMachineID)
	requestDTO.Ipaddresses = conv.SetTo[int64](ctx, model.IPAddressIds, &diags)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenService refreshes the Terraform model of type "service" from an API response.
func flattenService(ctx context.Context, responseDTO *netboxapi.ServiceResponseDTO, model *serviceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Protocol = conv.FromStringPtr(responseDTO.Protocol)
	model.Ports = conv.SetFrom(ctx, types.Int64Type, responseDTO.Ports, false, &diags)
	model.ParentObjectType = conv.FromStringPtr(responseDTO.ParentObjectType)
	model.ParentObjectID = conv.FromInt64Ptr(responseDTO.ParentObjectID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.DeviceID)
	model.VirtualMachineID = conv.FromInt64Ptr(responseDTO.VirtualMachineID)
	model.IPAddressIds = conv.SetFrom(ctx, types.Int64Type, responseDTO.Ipaddresses, model.IPAddressIds.IsNull() || model.IPAddressIds.IsUnknown(), &diags)
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
