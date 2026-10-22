// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// deviceBayResourceModel is the Terraform state/plan model of the "device_bay" resource.
type deviceBayResourceModel struct {
	ID                types.Int64  `tfsdk:"id"`
	DeviceID          types.Int64  `tfsdk:"device_id"`
	Name              types.String `tfsdk:"name"`
	Label             types.String `tfsdk:"label"`
	Enabled           types.Bool   `tfsdk:"enabled"`
	InstalledDeviceID types.Int64  `tfsdk:"installed_device_id"`
	Description       types.String `tfsdk:"description"`
	OwnerID           types.Int64  `tfsdk:"owner_id"`
	Created           types.String `tfsdk:"created"`
	LastUpdated       types.String `tfsdk:"last_updated"`
	URL               types.String `tfsdk:"url"`
	Tags              types.Set    `tfsdk:"tags"`
	TagsAll           types.Set    `tfsdk:"tags_all"`
	CustomFields      types.Map    `tfsdk:"custom_fields"`
}

// expandDeviceBay converts the Terraform model of type "device_bay" into an API request.
func expandDeviceBay(ctx context.Context, model *deviceBayResourceModel) (*netboxapi.DeviceBayRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceBayRequestDTO{}
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Label = conv.StringPtr(model.Label)
	requestDTO.Enabled = conv.BoolPtr(model.Enabled)
	requestDTO.InstalledDevice = conv.Int64Ptr(model.InstalledDeviceID)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDeviceBay refreshes the Terraform model of type "device_bay" from an API response.
func flattenDeviceBay(ctx context.Context, responseDTO *netboxapi.DeviceBayResponseDTO, model *deviceBayResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Label = conv.FromStringPtr(responseDTO.Label)
	model.Enabled = conv.FromBoolPtr(responseDTO.Enabled)
	model.InstalledDeviceID = conv.FromInt64Ptr(responseDTO.InstalledDevice)
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
