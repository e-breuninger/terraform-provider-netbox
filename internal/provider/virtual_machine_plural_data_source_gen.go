// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualMachinesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualMachinesDataSource)(nil)
)

// NewVirtualMachinesDataSource returns a new virtual_machines data source, which
// lists virtual_machine objects matching its filters.
func NewVirtualMachinesDataSource() datasource.DataSource {
	return &virtualMachinesDataSource{}
}

type virtualMachinesDataSource struct {
	client *netboxapi.Client
}

// virtualMachinesDataSourceModel is the data source model: the filters, the limit and the matching
// virtual_machines.
type virtualMachinesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"virtual_machines"`
}

// virtualMachineResourceAttrTypes is the attribute type map of virtualMachineResourceModel.
func virtualMachineResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                      types.Int64Type,
		"name":                    types.StringType,
		"status":                  types.StringType,
		"start_on_boot":           types.StringType,
		"cluster_id":              types.Int64Type,
		"site_id":                 types.Int64Type,
		"role_id":                 types.Int64Type,
		"virtual_machine_type_id": types.Int64Type,
		"tenant_id":               types.Int64Type,
		"platform_id":             types.Int64Type,
		"device_id":               types.Int64Type,
		"config_template_id":      types.Int64Type,
		"serial":                  types.StringType,
		"primary_ip4_id":          types.Int64Type,
		"primary_ip6_id":          types.Int64Type,
		"vcpus":                   types.Float64Type,
		"memory_mb":               types.Int64Type,
		"disk_size_mb":            types.Int64Type,
		"local_context_data":      jsontypes.NormalizedType{},
		"config_context":          jsontypes.NormalizedType{},
		"description":             types.StringType,
		"comments":                types.StringType,
		"owner_id":                types.Int64Type,
		"created":                 types.StringType,
		"last_updated":            types.StringType,
		"url":                     types.StringType,
		"interface_count":         types.Int64Type,
		"virtual_disk_count":      types.Int64Type,
		"tags":                    types.SetType{ElemType: types.StringType},
		"tags_all":                types.SetType{ElemType: types.StringType},
		"custom_fields":           types.MapType{ElemType: types.StringType},
	}
}

func (d *virtualMachinesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_machines"
}

func (d *virtualMachinesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Virtualization:Lists virtual_machine objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: cluster, cluster__n, cluster_group, cluster_group__n, cluster_group_id, cluster_group_id__n, cluster_id, cluster_id__n, cluster_type, cluster_type__n, cluster_type_id, cluster_type_id__n, config_template_id, config_template_id__n, contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device__n, device_id, device_id__n, disk, disk__empty, disk__gt, disk__gte, disk__lt, disk__lte, disk__n, has_primary_ip, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_count, interface_count__empty, interface_count__gt, interface_count__gte, interface_count__lt, interface_count__lte, interface_count__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, local_context_data, mac_address, mac_address__ic, mac_address__ie, mac_address__iew, mac_address__iregex, mac_address__isw, mac_address__n, mac_address__nic, mac_address__nie, mac_address__niew, mac_address__nisw, mac_address__regex, memory, memory__empty, memory__gt, memory__gte, memory__lt, memory__lte, memory__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, platform, platform__n, platform_id, platform_id__n, primary_ip4, primary_ip4__n, primary_ip4_id, primary_ip4_id__n, primary_ip6, primary_ip6__n, primary_ip6_id, primary_ip6_id__n, q, region, region__n, region_id, region_id__n, role, role__n, role_id, role_id__n, serial, serial__empty, serial__ic, serial__ie, serial__iew, serial__iregex, serial__isw, serial__n, serial__nic, serial__nie, serial__niew, serial__nisw, serial__regex, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, start_on_boot, start_on_boot__empty, start_on_boot__ic, start_on_boot__ie, start_on_boot__iew, start_on_boot__iregex, start_on_boot__isw, start_on_boot__n, start_on_boot__nic, start_on_boot__nie, start_on_boot__niew, start_on_boot__nisw, start_on_boot__regex, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, vcpus, vcpus__empty, vcpus__gt, vcpus__gte, vcpus__lt, vcpus__lte, vcpus__n, virtual_disk_count, virtual_disk_count__empty, virtual_disk_count__gt, virtual_disk_count__gte, virtual_disk_count__lt, virtual_disk_count__lte, virtual_disk_count__n, virtual_machine_type, virtual_machine_type__n, virtual_machine_type_id, virtual_machine_type_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "Name of the query filter (the API list parameter).",
						},
						"value": schema.StringAttribute{
							Required:    true,
							Description: "Value to filter by.",
						},
					},
				},
			},
			"limit": schema.Int64Attribute{
				Optional:    true,
				Description: "The maximum number of objects to return from the API lookup. Defaults to 1000.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"virtual_machines": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching virtual_machines.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: offline, active, planned, staged, failed, decommissioning, paused.",
						},
						"start_on_boot": schema.StringAttribute{
							Computed:    true,
							Description: "Power state the virtual machine takes when its host boots. One of: on, off, laststate.",
						},
						"cluster_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the cluster (a virtual machine needs a cluster or a site).",
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site (a virtual machine needs a cluster or a site).",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the functional device role.",
						},
						"virtual_machine_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual machine type.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"platform_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the platform.",
						},
						"device_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device hosting the VM.",
						},
						"config_template_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the config template.",
						},
						"serial": schema.StringAttribute{
							Computed:    true,
							Description: "Serial number.",
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
							Computed: true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"owner_id": schema.Int64Attribute{
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *virtualMachinesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualMachinesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualMachinesDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	limit := int64(1000)
	if !state.Limit.IsNull() && !state.Limit.IsUnknown() {
		limit = state.Limit.ValueInt64()
	}
	_ = limit

	// The filters entries, in no particular order. Entries sharing a name are distinct values of
	// the same API parameter.
	var filters []struct {
		Name  types.String `tfsdk:"name"`
		Value types.String `tfsdk:"value"`
	}
	if !state.Filters.IsNull() {
		resp.Diagnostics.Append(state.Filters.ElementsAs(ctx, &filters, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	_ = filters
	var responseDTOs []*netboxapi.VirtualMachineResponseDTO

	params := virtualization.NewVirtualizationVirtualMachinesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                         netboxapi.FilterInt64s(params.SetID),
		"id__n":                      netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                     netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                    netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                     netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                    netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                  netboxapi.FilterBool(params.SetIDEmpty),
		"name":                       netboxapi.FilterStrings(params.SetName),
		"name__n":                    netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                   netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                  netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                   netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                  netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                  netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                 netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                  netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                 netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":                netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":                netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":               netboxapi.FilterStrings(params.SetNameIregex),
		"cluster_id":                 netboxapi.FilterInt64s(params.SetClusterID),
		"cluster_id__n":              netboxapi.FilterInt64s(params.SetClusterIDn),
		"status":                     netboxapi.FilterStrings(params.SetStatus),
		"status__n":                  netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":                 netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":                netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":                 netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":                netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":                netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":               netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":                netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":               netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":              netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":              netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":             netboxapi.FilterStrings(params.SetStatusIregex),
		"site_id":                    netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":                 netboxapi.FilterInt64s(params.SetSiteIDn),
		"role_id":                    netboxapi.FilterStrings(params.SetRoleID),
		"role_id__n":                 netboxapi.FilterStrings(params.SetRoleIDn),
		"tenant_id":                  netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":               netboxapi.FilterInt64s(params.SetTenantIDn),
		"platform_id":                netboxapi.FilterStrings(params.SetPlatformID),
		"platform_id__n":             netboxapi.FilterStrings(params.SetPlatformIDn),
		"device_id":                  netboxapi.FilterInt64s(params.SetDeviceID),
		"device_id__n":               netboxapi.FilterInt64s(params.SetDeviceIDn),
		"cluster_group":              netboxapi.FilterStrings(params.SetClusterGroup),
		"cluster_group__n":           netboxapi.FilterStrings(params.SetClusterGroupn),
		"device":                     netboxapi.FilterStrings(params.SetDevice),
		"device__n":                  netboxapi.FilterStrings(params.SetDevicen),
		"platform":                   netboxapi.FilterStrings(params.SetPlatform),
		"platform__n":                netboxapi.FilterStrings(params.SetPlatformn),
		"region":                     netboxapi.FilterStrings(params.SetRegion),
		"region__n":                  netboxapi.FilterStrings(params.SetRegionn),
		"role":                       netboxapi.FilterStrings(params.SetRole),
		"role__n":                    netboxapi.FilterStrings(params.SetRolen),
		"site":                       netboxapi.FilterStrings(params.SetSite),
		"site__n":                    netboxapi.FilterStrings(params.SetSiten),
		"owner_id":                   netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                        netboxapi.FilterStrings(params.SetTag),
		"tag__n":                     netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                   netboxapi.FilterStrings(params.SetTagAny),
		"cluster":                    netboxapi.FilterStrings(params.SetCluster),
		"cluster__n":                 netboxapi.FilterStrings(params.SetClustern),
		"cluster_group_id":           netboxapi.FilterInt64s(params.SetClusterGroupID),
		"cluster_group_id__n":        netboxapi.FilterInt64s(params.SetClusterGroupIDn),
		"cluster_type":               netboxapi.FilterStrings(params.SetClusterType),
		"cluster_type__n":            netboxapi.FilterStrings(params.SetClusterTypen),
		"cluster_type_id":            netboxapi.FilterInt64s(params.SetClusterTypeID),
		"cluster_type_id__n":         netboxapi.FilterInt64s(params.SetClusterTypeIDn),
		"config_template_id":         netboxapi.FilterInt64s(params.SetConfigTemplateID),
		"config_template_id__n":      netboxapi.FilterInt64s(params.SetConfigTemplateIDn),
		"contact":                    netboxapi.FilterInt64s(params.SetContact),
		"contact__n":                 netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":              netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":           netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":               netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":            netboxapi.FilterInt64s(params.SetContactRolen),
		"created":                    netboxapi.FilterStrings(params.SetCreated),
		"created__n":                 netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":               netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":               netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":             netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":                netboxapi.FilterStrings(params.SetDescription),
		"description__n":             netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":            netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":           netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":            netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":           netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":           netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":          netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":           netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":          netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":         netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":         netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":        netboxapi.FilterStrings(params.SetDescriptionIregex),
		"disk":                       netboxapi.FilterInt64s(params.SetDisk),
		"disk__n":                    netboxapi.FilterInt64s(params.SetDiskn),
		"disk__lt":                   netboxapi.FilterInt64s(params.SetDiskLt),
		"disk__lte":                  netboxapi.FilterInt64s(params.SetDiskLte),
		"disk__gt":                   netboxapi.FilterInt64s(params.SetDiskGt),
		"disk__gte":                  netboxapi.FilterInt64s(params.SetDiskGte),
		"disk__empty":                netboxapi.FilterBool(params.SetDiskEmpty),
		"has_primary_ip":             netboxapi.FilterBool(params.SetHasPrimaryIP),
		"interface_count":            netboxapi.FilterInt64s(params.SetInterfaceCount),
		"interface_count__n":         netboxapi.FilterInt64s(params.SetInterfaceCountn),
		"interface_count__lt":        netboxapi.FilterInt64s(params.SetInterfaceCountLt),
		"interface_count__lte":       netboxapi.FilterInt64s(params.SetInterfaceCountLte),
		"interface_count__gt":        netboxapi.FilterInt64s(params.SetInterfaceCountGt),
		"interface_count__gte":       netboxapi.FilterInt64s(params.SetInterfaceCountGte),
		"interface_count__empty":     netboxapi.FilterBool(params.SetInterfaceCountEmpty),
		"last_updated":               netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":            netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":           netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":          netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":           netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":          netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":        netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"local_context_data":         netboxapi.FilterBool(params.SetLocalContextData),
		"mac_address":                netboxapi.FilterStrings(params.SetMacAddress),
		"mac_address__n":             netboxapi.FilterStrings(params.SetMacAddressn),
		"mac_address__ic":            netboxapi.FilterStrings(params.SetMacAddressIc),
		"mac_address__nic":           netboxapi.FilterStrings(params.SetMacAddressNic),
		"mac_address__ie":            netboxapi.FilterStrings(params.SetMacAddressIe),
		"mac_address__nie":           netboxapi.FilterStrings(params.SetMacAddressNie),
		"mac_address__isw":           netboxapi.FilterStrings(params.SetMacAddressIsw),
		"mac_address__nisw":          netboxapi.FilterStrings(params.SetMacAddressNisw),
		"mac_address__iew":           netboxapi.FilterStrings(params.SetMacAddressIew),
		"mac_address__niew":          netboxapi.FilterStrings(params.SetMacAddressNiew),
		"mac_address__regex":         netboxapi.FilterStrings(params.SetMacAddressRegex),
		"mac_address__iregex":        netboxapi.FilterStrings(params.SetMacAddressIregex),
		"memory":                     netboxapi.FilterInt64s(params.SetMemory),
		"memory__n":                  netboxapi.FilterInt64s(params.SetMemoryn),
		"memory__lt":                 netboxapi.FilterInt64s(params.SetMemoryLt),
		"memory__lte":                netboxapi.FilterInt64s(params.SetMemoryLte),
		"memory__gt":                 netboxapi.FilterInt64s(params.SetMemoryGt),
		"memory__gte":                netboxapi.FilterInt64s(params.SetMemoryGte),
		"memory__empty":              netboxapi.FilterBool(params.SetMemoryEmpty),
		"owner":                      netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                   netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":             netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":             netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":          netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"primary_ip4":                netboxapi.FilterStrings(params.SetPrimaryIp4),
		"primary_ip4__n":             netboxapi.FilterStrings(params.SetPrimaryIp4n),
		"primary_ip4_id":             netboxapi.FilterInt64s(params.SetPrimaryIp4ID),
		"primary_ip4_id__n":          netboxapi.FilterInt64s(params.SetPrimaryIp4IDn),
		"primary_ip6":                netboxapi.FilterStrings(params.SetPrimaryIp6),
		"primary_ip6__n":             netboxapi.FilterStrings(params.SetPrimaryIp6n),
		"primary_ip6_id":             netboxapi.FilterInt64s(params.SetPrimaryIp6ID),
		"primary_ip6_id__n":          netboxapi.FilterInt64s(params.SetPrimaryIp6IDn),
		"q":                          netboxapi.FilterString(params.SetQ),
		"region_id":                  netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":               netboxapi.FilterStrings(params.SetRegionIDn),
		"serial":                     netboxapi.FilterStrings(params.SetSerial),
		"serial__n":                  netboxapi.FilterStrings(params.SetSerialn),
		"serial__ic":                 netboxapi.FilterStrings(params.SetSerialIc),
		"serial__nic":                netboxapi.FilterStrings(params.SetSerialNic),
		"serial__ie":                 netboxapi.FilterStrings(params.SetSerialIe),
		"serial__nie":                netboxapi.FilterStrings(params.SetSerialNie),
		"serial__isw":                netboxapi.FilterStrings(params.SetSerialIsw),
		"serial__nisw":               netboxapi.FilterStrings(params.SetSerialNisw),
		"serial__iew":                netboxapi.FilterStrings(params.SetSerialIew),
		"serial__niew":               netboxapi.FilterStrings(params.SetSerialNiew),
		"serial__empty":              netboxapi.FilterBool(params.SetSerialEmpty),
		"serial__regex":              netboxapi.FilterStrings(params.SetSerialRegex),
		"serial__iregex":             netboxapi.FilterStrings(params.SetSerialIregex),
		"site_group":                 netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":              netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":              netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":           netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"start_on_boot":              netboxapi.FilterStrings(params.SetStartOnBoot),
		"start_on_boot__n":           netboxapi.FilterStrings(params.SetStartOnBootn),
		"start_on_boot__ic":          netboxapi.FilterStrings(params.SetStartOnBootIc),
		"start_on_boot__nic":         netboxapi.FilterStrings(params.SetStartOnBootNic),
		"start_on_boot__ie":          netboxapi.FilterStrings(params.SetStartOnBootIe),
		"start_on_boot__nie":         netboxapi.FilterStrings(params.SetStartOnBootNie),
		"start_on_boot__isw":         netboxapi.FilterStrings(params.SetStartOnBootIsw),
		"start_on_boot__nisw":        netboxapi.FilterStrings(params.SetStartOnBootNisw),
		"start_on_boot__iew":         netboxapi.FilterStrings(params.SetStartOnBootIew),
		"start_on_boot__niew":        netboxapi.FilterStrings(params.SetStartOnBootNiew),
		"start_on_boot__empty":       netboxapi.FilterBool(params.SetStartOnBootEmpty),
		"start_on_boot__regex":       netboxapi.FilterStrings(params.SetStartOnBootRegex),
		"start_on_boot__iregex":      netboxapi.FilterStrings(params.SetStartOnBootIregex),
		"tag_id":                     netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                  netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                     netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":                  netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":               netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":            netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":            netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":         netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"vcpus":                      netboxapi.FilterFloat64s(params.SetVcpus),
		"vcpus__n":                   netboxapi.FilterFloat64s(params.SetVcpusn),
		"vcpus__lt":                  netboxapi.FilterFloat64s(params.SetVcpusLt),
		"vcpus__lte":                 netboxapi.FilterFloat64s(params.SetVcpusLte),
		"vcpus__gt":                  netboxapi.FilterFloat64s(params.SetVcpusGt),
		"vcpus__gte":                 netboxapi.FilterFloat64s(params.SetVcpusGte),
		"vcpus__empty":               netboxapi.FilterBool(params.SetVcpusEmpty),
		"virtual_disk_count":         netboxapi.FilterInt64s(params.SetVirtualDiskCount),
		"virtual_disk_count__n":      netboxapi.FilterInt64s(params.SetVirtualDiskCountn),
		"virtual_disk_count__lt":     netboxapi.FilterInt64s(params.SetVirtualDiskCountLt),
		"virtual_disk_count__lte":    netboxapi.FilterInt64s(params.SetVirtualDiskCountLte),
		"virtual_disk_count__gt":     netboxapi.FilterInt64s(params.SetVirtualDiskCountGt),
		"virtual_disk_count__gte":    netboxapi.FilterInt64s(params.SetVirtualDiskCountGte),
		"virtual_disk_count__empty":  netboxapi.FilterBool(params.SetVirtualDiskCountEmpty),
		"virtual_machine_type":       netboxapi.FilterStrings(params.SetVirtualMachineType),
		"virtual_machine_type__n":    netboxapi.FilterStrings(params.SetVirtualMachineTypen),
		"virtual_machine_type_id":    netboxapi.FilterInt64s(params.SetVirtualMachineTypeID),
		"virtual_machine_type_id__n": netboxapi.FilterInt64s(params.SetVirtualMachineTypeIDn),
	}
	customFieldQuery := url.Values{}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		// Custom fields are queried as cf_<field name> with the field's own filter logic; NetBox ignores
		// names it has no custom field for.
		if strings.HasPrefix(name, "cf_") {
			customFieldQuery.Add(name, value)
			continue
		}
		listFilter, ok := listFilters[name]
		if !ok {
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
		if problem := listFilter.Add(name, value); problem != nil {
			resp.Diagnostics.AddError(problem.Summary, problem.Detail)
			return
		}
	}
	for _, listFilter := range listFilters {
		listFilter.Apply()
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Virtualization.VirtualizationVirtualMachinesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_virtual_machines", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VirtualMachineResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]virtualMachineResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model virtualMachineResourceModel
		resp.Diagnostics.Append(flattenVirtualMachine(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: virtualMachineResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
