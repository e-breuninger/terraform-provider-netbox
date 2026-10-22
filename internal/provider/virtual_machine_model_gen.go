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

// virtualMachineResourceModel is the Terraform state/plan model of the "virtual_machine" resource.
type virtualMachineResourceModel struct {
	ID                   types.Int64          `tfsdk:"id"`
	Name                 types.String         `tfsdk:"name"`
	Status               types.String         `tfsdk:"status"`
	StartOnBoot          types.String         `tfsdk:"start_on_boot"`
	ClusterID            types.Int64          `tfsdk:"cluster_id"`
	SiteID               types.Int64          `tfsdk:"site_id"`
	RoleID               types.Int64          `tfsdk:"role_id"`
	VirtualMachineTypeID types.Int64          `tfsdk:"virtual_machine_type_id"`
	TenantID             types.Int64          `tfsdk:"tenant_id"`
	PlatformID           types.Int64          `tfsdk:"platform_id"`
	DeviceID             types.Int64          `tfsdk:"device_id"`
	ConfigTemplateID     types.Int64          `tfsdk:"config_template_id"`
	Serial               types.String         `tfsdk:"serial"`
	PrimaryIp4ID         types.Int64          `tfsdk:"primary_ip4_id"`
	PrimaryIp6ID         types.Int64          `tfsdk:"primary_ip6_id"`
	Vcpus                types.Float64        `tfsdk:"vcpus"`
	MemoryMb             types.Int64          `tfsdk:"memory_mb"`
	DiskSizeMb           types.Int64          `tfsdk:"disk_size_mb"`
	LocalContextData     jsontypes.Normalized `tfsdk:"local_context_data"`
	Description          types.String         `tfsdk:"description"`
	Comments             types.String         `tfsdk:"comments"`
	OwnerID              types.Int64          `tfsdk:"owner_id"`
	Created              types.String         `tfsdk:"created"`
	LastUpdated          types.String         `tfsdk:"last_updated"`
	URL                  types.String         `tfsdk:"url"`
	InterfaceCount       types.Int64          `tfsdk:"interface_count"`
	VirtualDiskCount     types.Int64          `tfsdk:"virtual_disk_count"`
	Tags                 types.Set            `tfsdk:"tags"`
	TagsAll              types.Set            `tfsdk:"tags_all"`
	CustomFields         types.Map            `tfsdk:"custom_fields"`
}

// expandVirtualMachine converts the Terraform model of type "virtual_machine" into an API request.
func expandVirtualMachine(ctx context.Context, model *virtualMachineResourceModel) (*netboxapi.VirtualMachineRequestDTO, diag.Diagnostics) {
	var diags diag.Diagnostics
	requestDTO := &netboxapi.VirtualMachineRequestDTO{}
	requestDTO.Name = conv.StringPtr(model.Name)
	requestDTO.Status = conv.StringPtr(model.Status)
	requestDTO.StartOnBoot = conv.StringPtr(model.StartOnBoot)
	requestDTO.Cluster = conv.Int64Ptr(model.ClusterID)
	requestDTO.Site = conv.Int64Ptr(model.SiteID)
	requestDTO.Role = conv.Int64Ptr(model.RoleID)
	requestDTO.VirtualMachineType = conv.Int64Ptr(model.VirtualMachineTypeID)
	requestDTO.Tenant = conv.Int64Ptr(model.TenantID)
	requestDTO.Platform = conv.Int64Ptr(model.PlatformID)
	requestDTO.Device = conv.Int64Ptr(model.DeviceID)
	requestDTO.ConfigTemplate = conv.Int64Ptr(model.ConfigTemplateID)
	requestDTO.Serial = conv.StringPtr(model.Serial)
	requestDTO.Vcpus = conv.Float64Ptr(model.Vcpus)
	requestDTO.Memory = conv.Int64Ptr(model.MemoryMb)
	requestDTO.Disk = conv.Int64Ptr(model.DiskSizeMb)
	requestDTO.LocalContextData = conv.StringPtr(model.LocalContextData.StringValue)
	requestDTO.Description = conv.StringPtr(model.Description)
	requestDTO.Comments = conv.StringPtr(model.Comments)
	requestDTO.Owner = conv.Int64Ptr(model.OwnerID)
	requestDTO.Tags = conv.SetTo[string](ctx, model.Tags, &diags)
	requestDTO.CustomFields = conv.MapTo[string](ctx, model.CustomFields, &diags)
	return requestDTO, diags
}

// flattenVirtualMachine refreshes the Terraform model of type "virtual_machine" from an API response.
func flattenVirtualMachine(ctx context.Context, responseDTO *netboxapi.VirtualMachineResponseDTO, model *virtualMachineResourceModel) diag.Diagnostics {
	var diags diag.Diagnostics
	model.ID = conv.FromInt64Ptr(responseDTO.ID)
	model.Name = conv.FromStringPtr(responseDTO.Name)
	model.Status = conv.FromStringPtr(responseDTO.Status)
	model.StartOnBoot = conv.FromStringPtr(responseDTO.StartOnBoot)
	model.ClusterID = conv.FromInt64Ptr(responseDTO.Cluster)
	model.SiteID = conv.FromInt64Ptr(responseDTO.Site)
	model.RoleID = conv.FromInt64Ptr(responseDTO.Role)
	model.VirtualMachineTypeID = conv.FromInt64Ptr(responseDTO.VirtualMachineType)
	model.TenantID = conv.FromInt64Ptr(responseDTO.Tenant)
	model.PlatformID = conv.FromInt64Ptr(responseDTO.Platform)
	model.DeviceID = conv.FromInt64Ptr(responseDTO.Device)
	model.ConfigTemplateID = conv.FromInt64Ptr(responseDTO.ConfigTemplate)
	model.Serial = conv.FromStringPtr(responseDTO.Serial)
	model.PrimaryIp4ID = conv.FromInt64Ptr(responseDTO.PrimaryIp4)
	model.PrimaryIp6ID = conv.FromInt64Ptr(responseDTO.PrimaryIp6)
	model.Vcpus = conv.FromFloat64Ptr(responseDTO.Vcpus)
	model.MemoryMb = conv.FromInt64Ptr(responseDTO.Memory)
	model.DiskSizeMb = conv.FromInt64Ptr(responseDTO.Disk)
	model.LocalContextData = jsontypes.NewNormalizedPointerValue(conv.EmptyStringAsNil(responseDTO.LocalContextData))
	model.Description = conv.FromStringPtr(responseDTO.Description)
	model.Comments = conv.FromStringPtr(responseDTO.Comments)
	model.OwnerID = conv.FromInt64Ptr(responseDTO.Owner)
	model.Created = conv.FromStringPtr(responseDTO.Created)
	model.LastUpdated = conv.FromStringPtr(responseDTO.LastUpdated)
	model.URL = conv.FromStringPtr(responseDTO.URL)
	model.InterfaceCount = conv.FromInt64Ptr(responseDTO.InterfaceCount)
	model.VirtualDiskCount = conv.FromInt64Ptr(responseDTO.VirtualDiskCount)
	model.Tags = conv.SetFrom(ctx, types.StringType, responseDTO.Tags, model.Tags.IsNull() || model.Tags.IsUnknown(), &diags)
	model.TagsAll = conv.SetFrom(ctx, types.StringType, responseDTO.TagsAll, false, &diags)
	model.CustomFields = conv.MapFrom(ctx, types.StringType, responseDTO.CustomFields, false, &diags)
	return diags
}
