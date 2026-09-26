// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// moduleTypeResourceModel is the Terraform state/plan model of the "module_type" resource.
type moduleTypeResourceModel struct {
	ID                             types.Int64   `tfsdk:"id"`
	ManufacturerID                 types.Int64   `tfsdk:"manufacturer_id"`
	Model                          types.String  `tfsdk:"model"`
	PartNumber                     types.String  `tfsdk:"part_number"`
	Weight                         types.Float64 `tfsdk:"weight"`
	WeightUnit                     types.String  `tfsdk:"weight_unit"`
	Description                    types.String  `tfsdk:"description"`
	Comments                       types.String  `tfsdk:"comments"`
	OwnerID                        types.Int64   `tfsdk:"owner_id"`
	Created                        types.String  `tfsdk:"created"`
	LastUpdated                    types.String  `tfsdk:"last_updated"`
	URL                            types.String  `tfsdk:"url"`
	ModuleCount                    types.Int64   `tfsdk:"module_count"`
	ConsolePortTemplateCount       types.Int64   `tfsdk:"console_port_template_count"`
	ConsoleServerPortTemplateCount types.Int64   `tfsdk:"console_server_port_template_count"`
	PowerPortTemplateCount         types.Int64   `tfsdk:"power_port_template_count"`
	PowerOutletTemplateCount       types.Int64   `tfsdk:"power_outlet_template_count"`
	InterfaceTemplateCount         types.Int64   `tfsdk:"interface_template_count"`
	FrontPortTemplateCount         types.Int64   `tfsdk:"front_port_template_count"`
	RearPortTemplateCount          types.Int64   `tfsdk:"rear_port_template_count"`
	ModuleBayTemplateCount         types.Int64   `tfsdk:"module_bay_template_count"`
	Tags                           types.Set     `tfsdk:"tags"`
	TagsAll                        types.Set     `tfsdk:"tags_all"`
	CustomFields                   types.Map     `tfsdk:"custom_fields"`
}

// expandModuleType converts the Terraform model of type "module_type" into an API request.
func expandModuleType(ctx context.Context, model *moduleTypeResourceModel) (*netboxapi.ModuleTypeRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.ModuleTypeRequestDTO{}
	requestDTO.Manufacturer = conv.Int64Ptr(model.ManufacturerID)
	requestDTO.Model = conv.StringPtr(model.Model)
	requestDTO.PartNumber = conv.StringPtr(model.PartNumber)
	requestDTO.Weight = conv.Float64Ptr(model.Weight)
	requestDTO.WeightUnit = conv.StringPtr(model.WeightUnit)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenModuleType refreshes the Terraform model of type "module_type" from an API response.
func flattenModuleType(ctx context.Context, responseDTO *netboxapi.ModuleTypeResponseDTO, model *moduleTypeResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.ManufacturerID = conv.FromInt64Ptr(responseDTO.Manufacturer)
	model.Model = conv.FromStringPtr(responseDTO.Model)
	model.PartNumber = conv.FromStringPtr(responseDTO.PartNumber)
	model.Weight = conv.FromFloat64Ptr(responseDTO.Weight)
	model.WeightUnit = conv.FromStringPtr(responseDTO.WeightUnit)
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.ModuleCount = conv.FromInt64Ptr(responseDTO.ModuleCount)
	model.ConsolePortTemplateCount = conv.FromInt64Ptr(responseDTO.ConsolePortTemplateCount)
	model.ConsoleServerPortTemplateCount = conv.FromInt64Ptr(responseDTO.ConsoleServerPortTemplateCount)
	model.PowerPortTemplateCount = conv.FromInt64Ptr(responseDTO.PowerPortTemplateCount)
	model.PowerOutletTemplateCount = conv.FromInt64Ptr(responseDTO.PowerOutletTemplateCount)
	model.InterfaceTemplateCount = conv.FromInt64Ptr(responseDTO.InterfaceTemplateCount)
	model.FrontPortTemplateCount = conv.FromInt64Ptr(responseDTO.FrontPortTemplateCount)
	model.RearPortTemplateCount = conv.FromInt64Ptr(responseDTO.RearPortTemplateCount)
	model.ModuleBayTemplateCount = conv.FromInt64Ptr(responseDTO.ModuleBayTemplateCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
