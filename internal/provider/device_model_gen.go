// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// deviceResourceModel is the Terraform state/plan model of the "device" resource.
type deviceResourceModel struct {
	ID                     types.Int64          `tfsdk:"id"`
	Name                   types.String         `tfsdk:"name"`
	DeviceTypeID           types.Int64          `tfsdk:"device_type_id"`
	RoleID                 types.Int64          `tfsdk:"role_id"`
	SiteID                 types.Int64          `tfsdk:"site_id"`
	LocationID             types.Int64          `tfsdk:"location_id"`
	RackID                 types.Int64          `tfsdk:"rack_id"`
	RackPosition           types.Float64        `tfsdk:"rack_position"`
	RackFace               types.String         `tfsdk:"rack_face"`
	Status                 types.String         `tfsdk:"status"`
	Airflow                types.String         `tfsdk:"airflow"`
	TenantID               types.Int64          `tfsdk:"tenant_id"`
	PlatformID             types.Int64          `tfsdk:"platform_id"`
	ClusterID              types.Int64          `tfsdk:"cluster_id"`
	VirtualChassisID       types.Int64          `tfsdk:"virtual_chassis_id"`
	VirtualChassisPosition types.Int64          `tfsdk:"virtual_chassis_position"`
	VirtualChassisPriority types.Int64          `tfsdk:"virtual_chassis_priority"`
	ConfigTemplateID       types.Int64          `tfsdk:"config_template_id"`
	Serial                 types.String         `tfsdk:"serial"`
	AssetTag               types.String         `tfsdk:"asset_tag"`
	PrimaryIp4ID           types.Int64          `tfsdk:"primary_ip4_id"`
	PrimaryIp6ID           types.Int64          `tfsdk:"primary_ip6_id"`
	OobIPID                types.Int64          `tfsdk:"oob_ip_id"`
	LocalContextData       jsontypes.Normalized `tfsdk:"local_context_data"`
	Description            types.String         `tfsdk:"description"`
	Comments               types.String         `tfsdk:"comments"`
	OwnerID                types.Int64          `tfsdk:"owner_id"`
	Created                types.String         `tfsdk:"created"`
	LastUpdated            types.String         `tfsdk:"last_updated"`
	URL                    types.String         `tfsdk:"url"`
	ConsolePortCount       types.Int64          `tfsdk:"console_port_count"`
	ConsoleServerPortCount types.Int64          `tfsdk:"console_server_port_count"`
	PowerPortCount         types.Int64          `tfsdk:"power_port_count"`
	PowerOutletCount       types.Int64          `tfsdk:"power_outlet_count"`
	InterfaceCount         types.Int64          `tfsdk:"interface_count"`
	FrontPortCount         types.Int64          `tfsdk:"front_port_count"`
	RearPortCount          types.Int64          `tfsdk:"rear_port_count"`
	DeviceBayCount         types.Int64          `tfsdk:"device_bay_count"`
	ModuleBayCount         types.Int64          `tfsdk:"module_bay_count"`
	InventoryItemCount     types.Int64          `tfsdk:"inventory_item_count"`
	Tags                   types.Set            `tfsdk:"tags"`
	TagsAll                types.Set            `tfsdk:"tags_all"`
	CustomFields           types.Map            `tfsdk:"custom_fields"`
}

// expandDevice converts the Terraform model of type "device" into an API request.
func expandDevice(ctx context.Context, model *deviceResourceModel) (*netboxapi.DeviceRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.DeviceRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.DeviceType = conv.Int64Ptr(model.DeviceTypeID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.Site = conv.Int64Ptr(model.SiteID)
	requestDTO.Location = conv.Int64Ptr(model.LocationID)
	requestDTO.Rack = conv.Int64Ptr(model.RackID)
	requestDTO.Position = conv.Float64Ptr(model.RackPosition)
	requestDTO.Face = conv.StringPtr(model.RackFace)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.Airflow = conv.StringPtr(model.Airflow)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Platform = conv.Int64Ptr(model.PlatformID)
	requestDTO.Cluster = conv.Int64Ptr(model.ClusterID)
	requestDTO.VirtualChassis = conv.Int64Ptr(model.VirtualChassisID)
	requestDTO.VcPosition = conv.Int64Ptr(model.VirtualChassisPosition)
	requestDTO.VcPriority = conv.Int64Ptr(model.VirtualChassisPriority)
	requestDTO.ConfigTemplate = conv.Int64Ptr(model.ConfigTemplateID)
	requestDTO.Serial = conv.StringPtr(model.Serial)
	requestDTO.AssetTag = conv.StringPtr(model.AssetTag)
	requestDTO.LocalContextData = conv.StringPtr(model.LocalContextData.StringValue)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenDevice refreshes the Terraform model of type "device" from an API response.
func flattenDevice(ctx context.Context, responseDTO *netboxapi.DeviceResponseDTO, model *deviceResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.DeviceTypeID = conv.FromInt64Ptr(responseDTO.DeviceType)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.SiteID = conv.FromInt64Ptr(responseDTO.Site)
	model.LocationID = conv.FromInt64Ptr(responseDTO.Location)
	model.RackID = conv.FromInt64Ptr(responseDTO.Rack)
	model.RackPosition = conv.FromFloat64Ptr(responseDTO.Position)
	model.RackFace = conv.FromStringPtr(responseDTO.Face)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.Airflow = conv.FromStringPtr(responseDTO.Airflow)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.PlatformID = conv.FromInt64Ptr(responseDTO.Platform)
	model.ClusterID = conv.FromInt64Ptr(responseDTO.Cluster)
	model.VirtualChassisID = conv.FromInt64Ptr(responseDTO.VirtualChassis)
	model.VirtualChassisPosition = conv.FromInt64Ptr(responseDTO.VcPosition)
	model.VirtualChassisPriority = conv.FromInt64Ptr(responseDTO.VcPriority)
	model.ConfigTemplateID = conv.FromInt64Ptr(responseDTO.ConfigTemplate)
	model.Serial = conv.FromStringPtr(responseDTO.Serial)
	model.AssetTag = conv.FromStringPtr(responseDTO.AssetTag)
	model.PrimaryIp4ID = conv.FromInt64Ptr(responseDTO.PrimaryIp4)
	model.PrimaryIp6ID = conv.FromInt64Ptr(responseDTO.PrimaryIp6)
	model.OobIPID = conv.FromInt64Ptr(responseDTO.OobIP)
	model.LocalContextData = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.LocalContextData))
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.ConsolePortCount = conv.FromInt64Ptr(responseDTO.ConsolePortCount)
	model.ConsoleServerPortCount = conv.FromInt64Ptr(responseDTO.ConsoleServerPortCount)
	model.PowerPortCount = conv.FromInt64Ptr(responseDTO.PowerPortCount)
	model.PowerOutletCount = conv.FromInt64Ptr(responseDTO.PowerOutletCount)
	model.InterfaceCount = conv.FromInt64Ptr(responseDTO.InterfaceCount)
	model.FrontPortCount = conv.FromInt64Ptr(responseDTO.FrontPortCount)
	model.RearPortCount = conv.FromInt64Ptr(responseDTO.RearPortCount)
	model.DeviceBayCount = conv.FromInt64Ptr(responseDTO.DeviceBayCount)
	model.ModuleBayCount = conv.FromInt64Ptr(responseDTO.ModuleBayCount)
	model.InventoryItemCount = conv.FromInt64Ptr(responseDTO.InventoryItemCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
