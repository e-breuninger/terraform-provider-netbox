// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// deviceTypeResourceModel is the Terraform state/plan model of the "device_type" resource.
type deviceTypeResourceModel struct {
	ID                             types.Int64   `tfsdk:"id"`
	ManufacturerID                 types.Int64   `tfsdk:"manufacturer_id"`
	Model                          types.String  `tfsdk:"model"`
	Slug                           types.String  `tfsdk:"slug"`
	PartNumber                     types.String  `tfsdk:"part_number"`
	UHeight                        types.Float64 `tfsdk:"u_height"`
	IsFullDepth                    types.Bool    `tfsdk:"is_full_depth"`
	SubdeviceRole                  types.String  `tfsdk:"subdevice_role"`
	Airflow                        types.String  `tfsdk:"airflow"`
	Weight                         types.Float64 `tfsdk:"weight"`
	WeightUnit                     types.String  `tfsdk:"weight_unit"`
	Description                    types.String  `tfsdk:"description"`
	Comments                       types.String  `tfsdk:"comments"`
	OwnerID                        types.Int64   `tfsdk:"owner_id"`
	Created                        types.String  `tfsdk:"created"`
	LastUpdated                    types.String  `tfsdk:"last_updated"`
	URL                            types.String  `tfsdk:"url"`
	DeviceCount                    types.Int64   `tfsdk:"device_count"`
	ConsolePortTemplateCount       types.Int64   `tfsdk:"console_port_template_count"`
	ConsoleServerPortTemplateCount types.Int64   `tfsdk:"console_server_port_template_count"`
	PowerPortTemplateCount         types.Int64   `tfsdk:"power_port_template_count"`
	PowerOutletTemplateCount       types.Int64   `tfsdk:"power_outlet_template_count"`
	InterfaceTemplateCount         types.Int64   `tfsdk:"interface_template_count"`
	FrontPortTemplateCount         types.Int64   `tfsdk:"front_port_template_count"`
	RearPortTemplateCount          types.Int64   `tfsdk:"rear_port_template_count"`
	DeviceBayTemplateCount         types.Int64   `tfsdk:"device_bay_template_count"`
	ModuleBayTemplateCount         types.Int64   `tfsdk:"module_bay_template_count"`
	InventoryItemTemplateCount     types.Int64   `tfsdk:"inventory_item_template_count"`
	Tags                           types.Set     `tfsdk:"tags"`
	TagsAll                        types.Set     `tfsdk:"tags_all"`
	CustomFields                   types.Map     `tfsdk:"custom_fields"`
}

// expandDeviceType converts the Terraform model of type "device_type" into an API request.
func expandDeviceType(ctx context.Context, model *deviceTypeResourceModel) (*netboxapi.DeviceTypeRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceTypeRequestDTO{}
	requestDTO.Manufacturer = conv.Int64Ptr(model.ManufacturerID)
	requestDTO.Model = conv.StringPtr(model.Model)
	requestDTO.Slug = conv.StringPtr(model.Slug)
	requestDTO.PartNumber = conv.StringPtr(model.PartNumber)
	requestDTO.UHeight = conv.Float64Ptr(model.UHeight)
	requestDTO.IsFullDepth = conv.BoolPtr(model.IsFullDepth)
	requestDTO.SubdeviceRole = conv.StringPtr(model.SubdeviceRole)
	requestDTO.Airflow = conv.StringPtr(model.Airflow)
	requestDTO.Weight = conv.Float64Ptr(model.Weight)
	requestDTO.WeightUnit = conv.StringPtr(model.WeightUnit)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDeviceType refreshes the Terraform model of type "device_type" from an API response.
func flattenDeviceType(ctx context.Context, responseDTO *netboxapi.DeviceTypeResponseDTO, model *deviceTypeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.ManufacturerID = conv.FromInt64Ptr(responseDTO.Manufacturer)
	model.Model = conv.FromStringPtr(responseDTO.Model)
	model.Slug = conv.FromStringPtr(responseDTO.Slug)
	model.PartNumber = conv.FromStringPtr(responseDTO.PartNumber)
	model.UHeight = conv.FromFloat64Ptr(responseDTO.UHeight)
	model.IsFullDepth = conv.FromBoolPtr(responseDTO.IsFullDepth)
	model.SubdeviceRole = conv.FromStringPtr(responseDTO.SubdeviceRole)
	model.Airflow = conv.FromStringPtr(responseDTO.Airflow)
	model.Weight = conv.FromFloat64Ptr(responseDTO.Weight)
	model.WeightUnit = conv.FromStringPtr(responseDTO.WeightUnit)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.DeviceCount = conv.FromInt64Ptr(responseDTO.DeviceCount)
	model.ConsolePortTemplateCount = conv.FromInt64Ptr(responseDTO.ConsolePortTemplateCount)
	model.ConsoleServerPortTemplateCount = conv.FromInt64Ptr(responseDTO.ConsoleServerPortTemplateCount)
	model.PowerPortTemplateCount = conv.FromInt64Ptr(responseDTO.PowerPortTemplateCount)
	model.PowerOutletTemplateCount = conv.FromInt64Ptr(responseDTO.PowerOutletTemplateCount)
	model.InterfaceTemplateCount = conv.FromInt64Ptr(responseDTO.InterfaceTemplateCount)
	model.FrontPortTemplateCount = conv.FromInt64Ptr(responseDTO.FrontPortTemplateCount)
	model.RearPortTemplateCount = conv.FromInt64Ptr(responseDTO.RearPortTemplateCount)
	model.DeviceBayTemplateCount = conv.FromInt64Ptr(responseDTO.DeviceBayTemplateCount)
	model.ModuleBayTemplateCount = conv.FromInt64Ptr(responseDTO.ModuleBayTemplateCount)
	model.InventoryItemTemplateCount = conv.FromInt64Ptr(responseDTO.InventoryItemTemplateCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
