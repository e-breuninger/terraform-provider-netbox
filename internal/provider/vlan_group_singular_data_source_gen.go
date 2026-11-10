// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*vlanGroupDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlanGroupDataSource)(nil)
)

// NewVlanGroupDataSource returns a new vlan_group data source.
func NewVlanGroupDataSource() datasource.DataSource {
	return &vlanGroupDataSource{}
}

type vlanGroupDataSource struct {
	client *netboxapi.Client
}

// vlanGroupDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type vlanGroupDataSourceModel struct {
	vlanGroupResourceModel
	TagSlug            types.String  `tfsdk:"tag_slug"`
	Cluster            types.Int64   `tfsdk:"cluster"`
	ClusterGroup       types.Int64   `tfsdk:"cluster_group"`
	ContainsVid        types.Float64 `tfsdk:"contains_vid"`
	Location           types.Int64   `tfsdk:"location"`
	OwnerName          types.String  `tfsdk:"owner_name"`
	OwnerGroupName     types.String  `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64   `tfsdk:"owner_group_id"`
	Q                  types.String  `tfsdk:"q"`
	Rack               types.Int64   `tfsdk:"rack"`
	RackGroup          types.Float64 `tfsdk:"rack_group"`
	Region             types.Int64   `tfsdk:"region"`
	Site               types.Int64   `tfsdk:"site"`
	SiteGroup          types.Int64   `tfsdk:"site_group"`
	TagID              types.Int64   `tfsdk:"tag_id"`
	TenantSlug         types.String  `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String  `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String  `tfsdk:"tenant_group_id"`
	TotalVlanIds       types.Int64   `tfsdk:"total_vlan_ids"`
	CustomFieldFilters types.Map     `tfsdk:"custom_field_filters"`
}

func (d *vlanGroupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_group"
}

func (d *vlanGroupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A VLAN group (ipam.vlan-group).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/vlangroup/):\n\n> VLAN groups can be used to organize [VLANs](https://netboxlabs.com/docs/netbox/models/ipam/vlan/) within NetBox. Each VLAN group can be scoped to a particular [region](https://netboxlabs.com/docs/netbox/models/dcim/region/), [site group](https://netboxlabs.com/docs/netbox/models/dcim/sitegroup/), [site](https://netboxlabs.com/docs/netbox/models/dcim/sitegroup/), [location](https://netboxlabs.com/docs/netbox/models/dcim/location/), [rack](https://netboxlabs.com/docs/netbox/models/dcim/rack/), [cluster group](https://netboxlabs.com/docs/netbox/models/virtualization/clustergroup/), or [cluster](https://netboxlabs.com/docs/netbox/models/virtualization/cluster/). Member VLANs will be available for assignment to devices and/or virtual machines within the specified scope.\n>\n> Groups can also be used to enforce uniqueness: Each VLAN within a group must have a unique ID and name. VLANs which are not assigned to a group may have overlapping names and IDs (including VLANs which belong to a common site). For example, two VLANs with ID 123 may be created, but they cannot both be assigned to the same group.",
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
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from name when not set.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
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
			"vid_ranges": schema.ListNestedAttribute{
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"start": schema.Int64Attribute{
							Computed:    true,
							Description: "First VLAN id of the range. Must be a whole number.",
						},
						"end": schema.Int64Attribute{
							Computed:    true,
							Description: "Last VLAN id of the range. Must be a whole number.",
						},
					},
				},
				Computed:    true,
				Description: "VLAN id ranges of this group as {start, end} pairs.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"scope_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the scope. Derived from site_id, location_id, region_id or site_group_id when one of those is set; set it together with scope_id otherwise. One of: dcim.site, dcim.location, dcim.region, dcim.sitegroup.",
				Validators: []validator.String{
					stringvalidator.OneOf("dcim.site", "dcim.location", "dcim.region", "dcim.sitegroup"),
				},
			},
			"scope_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the scope object (see scope_type).",
			},
			"site_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the site the object is scoped to (scope_type dcim.site). Conflicts with the other scope aliases and with setting the scope_* pair directly.",
			},
			"location_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the location the object is scoped to (scope_type dcim.location).",
			},
			"region_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the region the object is scoped to (scope_type dcim.region).",
			},
			"site_group_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the site group the object is scoped to (scope_type dcim.sitegroup).",
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
			"vlan_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of VLANs in the group.",
			},
			"utilization": schema.StringAttribute{
				Computed:    true,
				Description: "Share of the VID ranges in use, as NetBox reports it, e.g. \"12.50%\".",
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"cluster": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster.",
			},
			"cluster_group": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster group.",
			},
			"contains_vid": schema.Float64Attribute{
				Optional:    true,
				Description: "Id of the contains vid.",
			},
			"location": schema.Int64Attribute{
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
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"rack": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the rack.",
			},
			"rack_group": schema.Float64Attribute{
				Optional:    true,
				Description: "Id of the rack group.",
			},
			"region": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"site": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site.",
			},
			"site_group": schema.Int64Attribute{
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
			"total_vlan_ids": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the total vlan ids.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *vlanGroupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vlanGroupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlanGroupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamVlanGroupsListParams()
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
	if !state.Slug.IsNull() {
		v := state.Slug.ValueString()
		params.SetSlug([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.ScopeType.IsNull() {
		v := state.ScopeType.ValueString()
		params.SetScopeType([]string{v})
		hasInput = true
	}
	if !state.ScopeID.IsNull() {
		v := state.ScopeID.ValueInt64()
		params.SetScopeID([]int64{v})
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
		v := state.Cluster.ValueInt64()
		params.SetCluster(&v)
		hasInput = true
	}
	if !state.ClusterGroup.IsNull() {
		v := state.ClusterGroup.ValueInt64()
		params.SetClusterGroup(&v)
		hasInput = true
	}
	if !state.ContainsVid.IsNull() {
		v := state.ContainsVid.ValueFloat64()
		params.SetContainsVid(&v)
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.Location.IsNull() {
		v := state.Location.ValueInt64()
		params.SetLocation(&v)
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
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.Rack.IsNull() {
		v := state.Rack.ValueInt64()
		params.SetRack(&v)
		hasInput = true
	}
	if !state.RackGroup.IsNull() {
		v := state.RackGroup.ValueFloat64()
		params.SetRackGroup(&v)
		hasInput = true
	}
	if !state.Region.IsNull() {
		v := state.Region.ValueInt64()
		params.SetRegion(&v)
		hasInput = true
	}
	if !state.Site.IsNull() {
		v := state.Site.ValueInt64()
		params.SetSite(&v)
		hasInput = true
	}
	if !state.SiteGroup.IsNull() {
		v := state.SiteGroup.ValueInt64()
		params.SetSiteGroup(&v)
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
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.TotalVlanIds.IsNull() {
		v := state.TotalVlanIds.ValueInt64()
		params.SetTotalVlanIds([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_vlan_group.")
		return
	}
	res, err := d.client.Ipam.IpamVlanGroupsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_vlan_group", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_vlan_group",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVlanGroup(ctx, netboxapi.VlanGroupResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.vlanGroupResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&vlanGroupResource{client: d.client}).postReadHook(ctx, &state.vlanGroupResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
