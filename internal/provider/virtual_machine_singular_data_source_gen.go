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
	NameContains       types.String `tfsdk:"name_contains"`
	ClusterGroupSlug   types.String `tfsdk:"cluster_group_slug"`
	DeviceName         types.String `tfsdk:"device_name"`
	PlatformSlug       types.String `tfsdk:"platform_slug"`
	RegionSlug         types.String `tfsdk:"region_slug"`
	RoleSlug           types.String `tfsdk:"role_slug"`
	SiteSlug           types.String `tfsdk:"site_slug"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
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
			"primary_ip4_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the primary IPv4 address (managed by netbox_primary_ip).",
			},
			"primary_ip6_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the primary IPv6 address (managed by netbox_primary_ip).",
			},
			"vcpus": schema.Float64Attribute{
				Computed:    true,
				Description: "Number of virtual CPUs (fractions allowed).",
			},
			"memory_mb": schema.Int64Attribute{
				Computed:    true,
				Description: "Memory in MB.",
			},
			"disk_size_mb": schema.Int64Attribute{
				Computed:    true,
				Description: "Disk size in MB. Once the machine has virtual disks NetBox computes it as their sum and it can no longer be set; leave it unset then.",
			},
			"local_context_data": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Local config context data as JSON text (use jsonencode()).",
			},
			"description": schema.StringAttribute{
				Computed: true,
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
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"interface_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of interfaces on the virtual machine.",
			},
			"virtual_disk_count": schema.Int64Attribute{
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
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
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
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or name_contains and/or cluster_id and/or status and/or site_id and/or role_id and/or tenant_id and/or platform_id and/or device_id and/or cluster_group_slug and/or device_name and/or platform_slug and/or region_slug and/or role_slug and/or site_slug and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_virtual_machine.")
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
