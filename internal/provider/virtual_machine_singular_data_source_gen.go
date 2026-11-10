// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualMachineDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualMachineDataSource)(nil)
)

// NewVirtualMachineDataSource returns a new virtual_machine data source.
func NewVirtualMachineDataSource() datasource.DataSource {
	return &virtualMachineDataSource{}
}

type virtualMachineDataSource struct {
	client *netboxapi.Client
}

// virtualMachineDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type virtualMachineDataSourceModel struct {
	virtualMachineResourceModel
	ClusterGroupSlug       types.String `tfsdk:"cluster_group_slug"`
	DeviceName             types.String `tfsdk:"device_name"`
	PlatformSlug           types.String `tfsdk:"platform_slug"`
	RegionSlug             types.String `tfsdk:"region_slug"`
	RoleSlug               types.String `tfsdk:"role_slug"`
	SiteSlug               types.String `tfsdk:"site_slug"`
	TagSlug                types.String `tfsdk:"tag_slug"`
	Cluster                types.String `tfsdk:"cluster"`
	ClusterGroupID         types.Int64  `tfsdk:"cluster_group_id"`
	ClusterTypeSlug        types.String `tfsdk:"cluster_type_slug"`
	ClusterTypeID          types.Int64  `tfsdk:"cluster_type_id"`
	Contact                types.Int64  `tfsdk:"contact"`
	ContactGroupSlug       types.String `tfsdk:"contact_group_slug"`
	ContactRole            types.Int64  `tfsdk:"contact_role"`
	Disk                   types.Int64  `tfsdk:"disk"`
	HasPrimaryIP           types.Bool   `tfsdk:"has_primary_ip"`
	HasLocalContextData    types.Bool   `tfsdk:"has_local_context_data"`
	MacAddress             types.String `tfsdk:"mac_address"`
	Memory                 types.Int64  `tfsdk:"memory"`
	OwnerName              types.String `tfsdk:"owner_name"`
	OwnerGroupName         types.String `tfsdk:"owner_group_name"`
	OwnerGroupID           types.Int64  `tfsdk:"owner_group_id"`
	PrimaryIp4             types.String `tfsdk:"primary_ip4"`
	PrimaryIp6             types.String `tfsdk:"primary_ip6"`
	Q                      types.String `tfsdk:"q"`
	RegionID               types.String `tfsdk:"region_id"`
	SiteGroupSlug          types.String `tfsdk:"site_group_slug"`
	SiteGroupID            types.String `tfsdk:"site_group_id"`
	TagID                  types.Int64  `tfsdk:"tag_id"`
	TenantSlug             types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug        types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID          types.String `tfsdk:"tenant_group_id"`
	VirtualMachineTypeSlug types.String `tfsdk:"virtual_machine_type_slug"`
	CustomFieldFilters     types.Map    `tfsdk:"custom_field_filters"`
}

func (d *virtualMachineDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_machine"
}

func (d *virtualMachineDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Virtualization:A virtual machine (virtualization.virtualmachine).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#virtual-machines):\n\n> A virtual machine is a virtualized compute instance. These behave in NetBox very similarly to device objects, but without any physical attributes. For example, a VM may have interfaces assigned to it with IP addresses and VLANs, however its interfaces cannot be connected via cables (because they are virtual). Each VM may also define its compute, memory, and storage resources as well.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: offline, active, planned, staged, failed, decommissioning, paused.",
				Validators: []validator.String{
					stringvalidator.OneOf("offline", "active", "planned", "staged", "failed", "decommissioning", "paused"),
				},
			},
			"start_on_boot": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Power state the virtual machine takes when its host boots. One of: on, off, laststate.",
				Validators: []validator.String{
					stringvalidator.OneOf("on", "off", "laststate"),
				},
			},
			"cluster_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the cluster (a virtual machine needs a cluster or a site).",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the site (a virtual machine needs a cluster or a site).",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the functional device role.",
			},
			"virtual_machine_type_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the virtual machine type.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"platform_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the platform.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device hosting the VM.",
			},
			"config_template_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the config template.",
			},
			"serial": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Serial number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"primary_ip4_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the primary IPv4 address (managed by netbox_primary_ip).",
			},
			"primary_ip6_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the primary IPv6 address (managed by netbox_primary_ip).",
			},
			"vcpus": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of virtual CPUs (fractions allowed).",
				Validators: []validator.Float64{
					float64validator.Between(0.01, 10000),
				},
			},
			"memory_mb": schema.Int64Attribute{
				Computed:    true,
				Description: "Memory in MB. Must be a whole number.",
			},
			"disk_size_mb": schema.Int64Attribute{
				Computed:    true,
				Description: "Disk size in MB. Once the machine has virtual disks NetBox computes it as their sum and it can no longer be set; leave it unset then. Must be a whole number.",
			},
			"local_context_data": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Local config context data as JSON text (use jsonencode()).",
			},
			"config_context": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Rendered config context (local_context_data merged with the context assigned through cluster, platform, role, site, tenant and tags) as JSON text, decode with jsondecode().",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"interface_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of interfaces on the virtual machine.",
			},
			"virtual_disk_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of virtual disks on the virtual machine.",
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of the tags assigned to the object (the provider's default_tags are added on top, see tags_all).",
			},
			"tags_all": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of all tags on the object, including the provider's default_tags.",
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"cluster_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the cluster group.",
			},
			"device_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the device hosting the virtual machine.",
			},
			"platform_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the platform.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region of the site.",
			},
			"role_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the role.",
			},
			"site_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"cluster": schema.StringAttribute{
				Optional:    true,
				Description: "Cluster.",
			},
			"cluster_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster group.",
			},
			"cluster_type_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the cluster type.",
			},
			"cluster_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster type.",
			},
			"contact": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact.",
			},
			"contact_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the contact group.",
			},
			"contact_role": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact role.",
			},
			"disk": schema.Int64Attribute{
				Optional:    true,
				Description: "Disk to match.",
			},
			"has_primary_ip": schema.BoolAttribute{
				Optional:    true,
				Description: "Has a primary IP.",
			},
			"has_local_context_data": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether the object has local config context data.",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Description: "Mac address to match.",
			},
			"memory": schema.Int64Attribute{
				Optional:    true,
				Description: "Memory to match.",
			},
			"owner_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner.",
			},
			"owner_group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner group.",
			},
			"owner_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner group.",
			},
			"primary_ip4": schema.StringAttribute{
				Optional:    true,
				Description: "Primary IPv4 (address).",
			},
			"primary_ip6": schema.StringAttribute{
				Optional:    true,
				Description: "Primary IPv6 (address).",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"region_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"site_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site group.",
			},
			"site_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the site group.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"tenant_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant.",
			},
			"tenant_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant group.",
			},
			"tenant_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the tenant group.",
			},
			"virtual_machine_type_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the virtual machine type.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *virtualMachineDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected data source configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = c
}

func (d *virtualMachineDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualMachineDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := virtualization.NewVirtualizationVirtualMachinesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.ClusterID.IsNull() {
		v := state.ClusterID.ValueInt64()
		params.SetClusterID([]int64{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.RoleID.IsNull() {
		v := strconv.FormatInt(state.RoleID.ValueInt64(), 10)
		params.SetRoleID([]string{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.PlatformID.IsNull() {
		v := strconv.FormatInt(state.PlatformID.ValueInt64(), 10)
		params.SetPlatformID([]string{v})
		hasInput = true
	}
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.ClusterGroupSlug.IsNull() {
		v := state.ClusterGroupSlug.ValueString()
		params.SetClusterGroup([]string{v})
		hasInput = true
	}
	if !state.DeviceName.IsNull() {
		v := state.DeviceName.ValueString()
		params.SetDevice([]string{v})
		hasInput = true
	}
	if !state.PlatformSlug.IsNull() {
		v := state.PlatformSlug.ValueString()
		params.SetPlatform([]string{v})
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RoleSlug.IsNull() {
		v := state.RoleSlug.ValueString()
		params.SetRole([]string{v})
		hasInput = true
	}
	if !state.SiteSlug.IsNull() {
		v := state.SiteSlug.ValueString()
		params.SetSite([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
		hasInput = true
	}
	if !state.Cluster.IsNull() {
		v := state.Cluster.ValueString()
		params.SetCluster([]string{v})
		hasInput = true
	}
	if !state.ClusterGroupID.IsNull() {
		v := state.ClusterGroupID.ValueInt64()
		params.SetClusterGroupID([]int64{v})
		hasInput = true
	}
	if !state.ClusterTypeSlug.IsNull() {
		v := state.ClusterTypeSlug.ValueString()
		params.SetClusterType([]string{v})
		hasInput = true
	}
	if !state.ClusterTypeID.IsNull() {
		v := state.ClusterTypeID.ValueInt64()
		params.SetClusterTypeID([]int64{v})
		hasInput = true
	}
	if !state.ConfigTemplateID.IsNull() {
		v := state.ConfigTemplateID.ValueInt64()
		params.SetConfigTemplateID([]int64{v})
		hasInput = true
	}
	if !state.Contact.IsNull() {
		v := state.Contact.ValueInt64()
		params.SetContact([]int64{v})
		hasInput = true
	}
	if !state.ContactGroupSlug.IsNull() {
		v := state.ContactGroupSlug.ValueString()
		params.SetContactGroup([]string{v})
		hasInput = true
	}
	if !state.ContactRole.IsNull() {
		v := state.ContactRole.ValueInt64()
		params.SetContactRole([]int64{v})
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.Disk.IsNull() {
		v := state.Disk.ValueInt64()
		params.SetDisk([]int64{v})
		hasInput = true
	}
	if !state.HasPrimaryIP.IsNull() {
		v := state.HasPrimaryIP.ValueBool()
		params.SetHasPrimaryIP(&v)
		hasInput = true
	}
	if !state.InterfaceCount.IsNull() {
		v := state.InterfaceCount.ValueInt64()
		params.SetInterfaceCount([]int64{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.HasLocalContextData.IsNull() {
		v := state.HasLocalContextData.ValueBool()
		params.SetLocalContextData(&v)
		hasInput = true
	}
	if !state.MacAddress.IsNull() {
		v := state.MacAddress.ValueString()
		params.SetMacAddress([]string{v})
		hasInput = true
	}
	if !state.Memory.IsNull() {
		v := state.Memory.ValueInt64()
		params.SetMemory([]int64{v})
		hasInput = true
	}
	if !state.OwnerName.IsNull() {
		v := state.OwnerName.ValueString()
		params.SetOwner([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupName.IsNull() {
		v := state.OwnerGroupName.ValueString()
		params.SetOwnerGroup([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupID.IsNull() {
		v := state.OwnerGroupID.ValueInt64()
		params.SetOwnerGroupID([]int64{v})
		hasInput = true
	}
	if !state.PrimaryIp4.IsNull() {
		v := state.PrimaryIp4.ValueString()
		params.SetPrimaryIp4([]string{v})
		hasInput = true
	}
	if !state.PrimaryIp4ID.IsNull() {
		v := state.PrimaryIp4ID.ValueInt64()
		params.SetPrimaryIp4ID([]int64{v})
		hasInput = true
	}
	if !state.PrimaryIp6.IsNull() {
		v := state.PrimaryIp6.ValueString()
		params.SetPrimaryIp6([]string{v})
		hasInput = true
	}
	if !state.PrimaryIp6ID.IsNull() {
		v := state.PrimaryIp6ID.ValueInt64()
		params.SetPrimaryIp6ID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueString()
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.Serial.IsNull() {
		v := state.Serial.ValueString()
		params.SetSerial([]string{v})
		hasInput = true
	}
	if !state.SiteGroupSlug.IsNull() {
		v := state.SiteGroupSlug.ValueString()
		params.SetSiteGroup([]string{v})
		hasInput = true
	}
	if !state.SiteGroupID.IsNull() {
		v := state.SiteGroupID.ValueString()
		params.SetSiteGroupID([]string{v})
		hasInput = true
	}
	if !state.StartOnBoot.IsNull() {
		v := state.StartOnBoot.ValueString()
		params.SetStartOnBoot([]string{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.TenantSlug.IsNull() {
		v := state.TenantSlug.ValueString()
		params.SetTenant([]string{v})
		hasInput = true
	}
	if !state.TenantGroupSlug.IsNull() {
		v := state.TenantGroupSlug.ValueString()
		params.SetTenantGroup([]string{v})
		hasInput = true
	}
	if !state.TenantGroupID.IsNull() {
		v := state.TenantGroupID.ValueString()
		params.SetTenantGroupID([]string{v})
		hasInput = true
	}
	if !state.Vcpus.IsNull() {
		v := state.Vcpus.ValueFloat64()
		params.SetVcpus([]float64{v})
		hasInput = true
	}
	if !state.VirtualDiskCount.IsNull() {
		v := state.VirtualDiskCount.ValueInt64()
		params.SetVirtualDiskCount([]int64{v})
		hasInput = true
	}
	if !state.VirtualMachineTypeSlug.IsNull() {
		v := state.VirtualMachineTypeSlug.ValueString()
		params.SetVirtualMachineType([]string{v})
		hasInput = true
	}
	if !state.VirtualMachineTypeID.IsNull() {
		v := state.VirtualMachineTypeID.ValueInt64()
		params.SetVirtualMachineTypeID([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_virtual_machine.")
		return
	}
	res, err := d.client.Virtualization.VirtualizationVirtualMachinesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_virtual_machine", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_virtual_machine",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVirtualMachine(ctx, netboxapi.VirtualMachineResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.virtualMachineResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
