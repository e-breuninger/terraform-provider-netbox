// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*configContextDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*configContextDataSource)(nil)
)

// NewConfigContextDataSource returns a new config_context data source.
func NewConfigContextDataSource() datasource.DataSource {
	return &configContextDataSource{}
}

type configContextDataSource struct {
	client *netboxapi.Client
}

// configContextDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type configContextDataSourceModel struct {
	configContextResourceModel
	AutoSyncEnabled  types.Bool   `tfsdk:"auto_sync_enabled"`
	ClusterGroupSlug types.String `tfsdk:"cluster_group_slug"`
	ClusterGroupID   types.Int64  `tfsdk:"cluster_group_id"`
	ClusterID        types.Int64  `tfsdk:"cluster_id"`
	ClusterTypeSlug  types.String `tfsdk:"cluster_type_slug"`
	ClusterTypeID    types.Int64  `tfsdk:"cluster_type_id"`
	DataFileID       types.Int64  `tfsdk:"data_file_id"`
	DataSourceID     types.Int64  `tfsdk:"data_source_id"`
	DataSynced       types.String `tfsdk:"data_synced"`
	DeviceRoleSlug   types.String `tfsdk:"device_role_slug"`
	DeviceRoleID     types.Int64  `tfsdk:"device_role_id"`
	DeviceTypeID     types.Int64  `tfsdk:"device_type_id"`
	LocationSlug     types.String `tfsdk:"location_slug"`
	LocationID       types.Int64  `tfsdk:"location_id"`
	OwnerName        types.String `tfsdk:"owner_name"`
	OwnerGroupName   types.String `tfsdk:"owner_group_name"`
	OwnerGroupID     types.Int64  `tfsdk:"owner_group_id"`
	PlatformSlug     types.String `tfsdk:"platform_slug"`
	PlatformID       types.Int64  `tfsdk:"platform_id"`
	ProfileName      types.String `tfsdk:"profile_name"`
	Q                types.String `tfsdk:"q"`
	RegionSlug       types.String `tfsdk:"region_slug"`
	RegionID         types.Int64  `tfsdk:"region_id"`
	SiteSlug         types.String `tfsdk:"site_slug"`
	SiteGroupSlug    types.String `tfsdk:"site_group_slug"`
	SiteGroupID      types.Int64  `tfsdk:"site_group_id"`
	SiteID           types.Int64  `tfsdk:"site_id"`
	TagSlug          types.String `tfsdk:"tag_slug"`
	TagID            types.Int64  `tfsdk:"tag_id"`
	TenantSlug       types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug  types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID    types.Int64  `tfsdk:"tenant_group_id"`
	TenantID         types.Int64  `tfsdk:"tenant_id"`
}

func (d *configContextDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_context"
}

func (d *configContextDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox config context (extras.configcontext): JSON data merged into the rendered context of the devices and virtual machines its scope selects.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/extras/configcontext/):\n\n> Context data is made available to devices and/or virtual machines based on their relationships to other objects in NetBox. For example, context data can be associated only with devices assigned to a particular site, or only to virtual machines in a certain cluster.",
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
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Merge order; higher weights are applied later and win. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"is_active": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the context is applied.",
			},
			"data": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "The context data as JSON text (use jsonencode()).",
			},
			"profile_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the config context profile whose schema the data must satisfy.",
			},
			"region_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the regions the context applies to; empty means all.",
			},
			"regions": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use region_ids instead.",
				Description:        "Deprecated alias of region_ids.",
			},
			"site_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the site groups the context applies to; empty means all.",
			},
			"site_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use site_group_ids instead.",
				Description:        "Deprecated alias of site_group_ids.",
			},
			"site_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the sites the context applies to; empty means all.",
			},
			"sites": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use site_ids instead.",
				Description:        "Deprecated alias of site_ids.",
			},
			"location_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the locations the context applies to; empty means all.",
			},
			"locations": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use location_ids instead.",
				Description:        "Deprecated alias of location_ids.",
			},
			"device_type_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the device types the context applies to; empty means all.",
			},
			"device_types": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use device_type_ids instead.",
				Description:        "Deprecated alias of device_type_ids.",
			},
			"device_role_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the device roles the context applies to; empty means all.",
			},
			"roles": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use device_role_ids instead.",
				Description:        "Deprecated alias of device_role_ids.",
			},
			"platform_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the platforms the context applies to; empty means all.",
			},
			"platforms": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use platform_ids instead.",
				Description:        "Deprecated alias of platform_ids.",
			},
			"cluster_type_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the cluster types the context applies to; empty means all.",
			},
			"cluster_types": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use cluster_type_ids instead.",
				Description:        "Deprecated alias of cluster_type_ids.",
			},
			"cluster_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the cluster groups the context applies to; empty means all.",
			},
			"cluster_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use cluster_group_ids instead.",
				Description:        "Deprecated alias of cluster_group_ids.",
			},
			"cluster_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the clusters the context applies to; empty means all.",
			},
			"clusters": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use cluster_ids instead.",
				Description:        "Deprecated alias of cluster_ids.",
			},
			"tenant_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the tenant groups the context applies to; empty means all.",
			},
			"tenant_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use tenant_group_ids instead.",
				Description:        "Deprecated alias of tenant_group_ids.",
			},
			"tenant_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the tenants the context applies to; empty means all.",
			},
			"tenants": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use tenant_ids instead.",
				Description:        "Deprecated alias of tenant_ids.",
			},
			"tag_slugs": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of the tags an object must carry for the context to apply; empty means any. This is a scope, not the context's own tags.",
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
			"auto_sync_enabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether auto sync enabled is set.",
			},
			"cluster_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the cluster group.",
			},
			"cluster_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster group.",
			},
			"cluster_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster.",
			},
			"cluster_type_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the cluster type.",
			},
			"cluster_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster type.",
			},
			"data_file_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the data file.",
			},
			"data_source_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the data source.",
			},
			"data_synced": schema.StringAttribute{
				Optional:    true,
				Description: "Data synced to match.",
			},
			"device_role_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the role.",
			},
			"device_role_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the role.",
			},
			"device_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device type.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the location.",
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
			"platform_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the platform.",
			},
			"platform_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the platform.",
			},
			"profile_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the profile.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
			},
			"region_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"site_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site.",
			},
			"site_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site group.",
			},
			"site_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site group.",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
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
			"tenant_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant group.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant.",
			},
		},
	}
}

func (d *configContextDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *configContextDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data configContextDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasConfigContextsListParams()
	hasInput := false
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
	if !state.IsActive.IsNull() {
		v := state.IsActive.ValueBool()
		params.SetIsActive(&v)
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.AutoSyncEnabled.IsNull() {
		v := state.AutoSyncEnabled.ValueBool()
		params.SetAutoSyncEnabled(&v)
		hasInput = true
	}
	if !state.ClusterGroupSlug.IsNull() {
		v := state.ClusterGroupSlug.ValueString()
		params.SetClusterGroup([]string{v})
		hasInput = true
	}
	if !state.ClusterGroupID.IsNull() {
		v := state.ClusterGroupID.ValueInt64()
		params.SetClusterGroupID([]int64{v})
		hasInput = true
	}
	if !state.ClusterID.IsNull() {
		v := state.ClusterID.ValueInt64()
		params.SetClusterID([]int64{v})
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
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.DataFileID.IsNull() {
		v := state.DataFileID.ValueInt64()
		params.SetDataFileID([]int64{v})
		hasInput = true
	}
	if !state.DataSourceID.IsNull() {
		v := state.DataSourceID.ValueInt64()
		params.SetDataSourceID([]int64{v})
		hasInput = true
	}
	if !state.DataSynced.IsNull() {
		v := state.DataSynced.ValueString()
		params.SetDataSynced([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.DeviceRoleSlug.IsNull() {
		v := state.DeviceRoleSlug.ValueString()
		params.SetDeviceRole([]string{v})
		hasInput = true
	}
	if !state.DeviceRoleID.IsNull() {
		v := state.DeviceRoleID.ValueInt64()
		params.SetDeviceRoleID([]int64{v})
		hasInput = true
	}
	if !state.DeviceTypeID.IsNull() {
		v := state.DeviceTypeID.ValueInt64()
		params.SetDeviceTypeID([]int64{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.LocationSlug.IsNull() {
		v := state.LocationSlug.ValueString()
		params.SetLocation([]string{v})
		hasInput = true
	}
	if !state.LocationID.IsNull() {
		v := state.LocationID.ValueInt64()
		params.SetLocationID([]int64{v})
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
	if !state.PlatformSlug.IsNull() {
		v := state.PlatformSlug.ValueString()
		params.SetPlatform([]string{v})
		hasInput = true
	}
	if !state.PlatformID.IsNull() {
		v := state.PlatformID.ValueInt64()
		params.SetPlatformID([]int64{v})
		hasInput = true
	}
	if !state.ProfileName.IsNull() {
		v := state.ProfileName.ValueString()
		params.SetProfile([]string{v})
		hasInput = true
	}
	if !state.ProfileID.IsNull() {
		v := state.ProfileID.ValueInt64()
		params.SetProfileID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueInt64()
		params.SetRegionID([]int64{v})
		hasInput = true
	}
	if !state.SiteSlug.IsNull() {
		v := state.SiteSlug.ValueString()
		params.SetSite([]string{v})
		hasInput = true
	}
	if !state.SiteGroupSlug.IsNull() {
		v := state.SiteGroupSlug.ValueString()
		params.SetSiteGroup([]string{v})
		hasInput = true
	}
	if !state.SiteGroupID.IsNull() {
		v := state.SiteGroupID.ValueInt64()
		params.SetSiteGroupID([]int64{v})
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
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
		v := state.TenantGroupID.ValueInt64()
		params.SetTenantGroupID([]int64{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.Weight.IsNull() {
		v := state.Weight.ValueInt64()
		params.SetWeight([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_config_context.")
		return
	}
	res, err := d.client.Extras.ExtrasConfigContextsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_config_context", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_config_context",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenConfigContext(ctx, netboxapi.ConfigContextResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.configContextResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
