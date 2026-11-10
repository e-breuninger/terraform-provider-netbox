// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*vlanDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlanDataSource)(nil)
)

// NewVlanDataSource returns a new vlan data source.
func NewVlanDataSource() datasource.DataSource {
	return &vlanDataSource{}
}

type vlanDataSource struct {
	client *netboxapi.Client
}

// vlanDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type vlanDataSourceModel struct {
	vlanResourceModel
	GroupSlug                 types.String `tfsdk:"group_slug"`
	TenantSlug                types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug           types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID             types.Int64  `tfsdk:"tenant_group_id"`
	TagSlug                   types.String `tfsdk:"tag_slug"`
	AvailableAtSite           types.String `tfsdk:"available_at_site"`
	AvailableAtSiteGroup      types.String `tfsdk:"available_at_site_group"`
	AvailableOnDevice         types.String `tfsdk:"available_on_device"`
	AvailableOnVirtualmachine types.String `tfsdk:"available_on_virtualmachine"`
	InterfaceID               types.Int64  `tfsdk:"interface_id"`
	L2vpn                     types.Int64  `tfsdk:"l2vpn"`
	L2vpnID                   types.Int64  `tfsdk:"l2vpn_id"`
	OwnerName                 types.String `tfsdk:"owner_name"`
	OwnerGroupName            types.String `tfsdk:"owner_group_name"`
	OwnerGroupID              types.Int64  `tfsdk:"owner_group_id"`
	Q                         types.String `tfsdk:"q"`
	QinqSvlanVid              types.Int64  `tfsdk:"qinq_svlan_vid"`
	RegionSlug                types.String `tfsdk:"region_slug"`
	RegionID                  types.String `tfsdk:"region_id"`
	RoleSlug                  types.String `tfsdk:"role_slug"`
	SiteSlug                  types.String `tfsdk:"site_slug"`
	SiteGroupSlug             types.String `tfsdk:"site_group_slug"`
	SiteGroupID               types.String `tfsdk:"site_group_id"`
	TagID                     types.Int64  `tfsdk:"tag_id"`
	VminterfaceID             types.Int64  `tfsdk:"vminterface_id"`
	CustomFieldFilters        types.Map    `tfsdk:"custom_field_filters"`
}

func (d *vlanDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan"
}

func (d *vlanDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox VLAN (ipam.vlan).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/vlan/):\n\n> A Virtual LAN (VLAN) represents an isolated layer two domain, identified by a name and a numeric ID (1-4094) as defined in [IEEE 802.1Q](https://en.wikipedia.org/wiki/IEEE_802.1Q). VLANs are arranged into [VLAN groups](https://netboxlabs.com/docs/netbox/models/ipam/vlangroup/) to define scope and to enforce uniqueness.",
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
			"vid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN id (1-4094). Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: active, reserved, deprecated.",
				Validators: []validator.String{
					stringvalidator.OneOf("active", "reserved", "deprecated"),
				},
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the site.",
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VLAN group.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the role.",
			},
			"qinq_role": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Q-in-Q role: a service VLAN (svlan) carrying customer VLANs, or a customer VLAN (cvlan) inside one. One of: svlan, cvlan.",
				Validators: []validator.String{
					stringvalidator.OneOf("svlan", "cvlan"),
				},
			},
			"qinq_svlan_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the service VLAN this customer VLAN is carried in (qinq_role cvlan).",
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
			"prefix_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of prefixes on the VLAN.",
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
			"group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the VLAN group.",
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"available_at_site": schema.StringAttribute{
				Optional:    true,
				Description: "Available at site to match.",
			},
			"available_at_site_group": schema.StringAttribute{
				Optional:    true,
				Description: "Available at site group to match.",
			},
			"available_on_device": schema.StringAttribute{
				Optional:    true,
				Description: "Available on device to match.",
			},
			"available_on_virtualmachine": schema.StringAttribute{
				Optional:    true,
				Description: "Available on virtualmachine to match.",
			},
			"interface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the interface.",
			},
			"l2vpn": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the L2VPN.",
			},
			"l2vpn_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the L2VPN.",
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
			"qinq_svlan_vid": schema.Int64Attribute{
				Optional:    true,
				Description: "Qinq svlan vid to match.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
			},
			"region_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"role_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the role.",
			},
			"site_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site.",
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
			"vminterface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the vminterface.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *vlanDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vlanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlanDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamVlansListParams()
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
	if !state.Vid.IsNull() {
		v := state.Vid.ValueInt64()
		params.SetVid([]int64{v})
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
	if !state.GroupID.IsNull() {
		v := state.GroupID.ValueInt64()
		params.SetGroupID([]int64{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.GroupSlug.IsNull() {
		v := state.GroupSlug.ValueString()
		params.SetGroup([]string{v})
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
		v := strconv.FormatInt(state.TenantGroupID.ValueInt64(), 10)
		params.SetTenantGroupID([]string{v})
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
	if !state.AvailableAtSite.IsNull() {
		v := state.AvailableAtSite.ValueString()
		params.SetAvailableAtSite(&v)
		hasInput = true
	}
	if !state.AvailableAtSiteGroup.IsNull() {
		v := state.AvailableAtSiteGroup.ValueString()
		params.SetAvailableAtSiteGroup(&v)
		hasInput = true
	}
	if !state.AvailableOnDevice.IsNull() {
		v := state.AvailableOnDevice.ValueString()
		params.SetAvailableOnDevice(&v)
		hasInput = true
	}
	if !state.AvailableOnVirtualmachine.IsNull() {
		v := state.AvailableOnVirtualmachine.ValueString()
		params.SetAvailableOnVirtualmachine(&v)
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
	if !state.InterfaceID.IsNull() {
		v := state.InterfaceID.ValueInt64()
		params.SetInterfaceID(&v)
		hasInput = true
	}
	if !state.L2vpn.IsNull() {
		v := state.L2vpn.ValueInt64()
		params.SetL2vpn([]int64{v})
		hasInput = true
	}
	if !state.L2vpnID.IsNull() {
		v := state.L2vpnID.ValueInt64()
		params.SetL2vpnID([]int64{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
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
	if !state.QinqRole.IsNull() {
		v := state.QinqRole.ValueString()
		params.SetQinqRole([]string{v})
		hasInput = true
	}
	if !state.QinqSvlanID.IsNull() {
		v := state.QinqSvlanID.ValueInt64()
		params.SetQinqSvlanID([]int64{v})
		hasInput = true
	}
	if !state.QinqSvlanVid.IsNull() {
		v := state.QinqSvlanVid.ValueInt64()
		params.SetQinqSvlanVid([]int64{v})
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueString()
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.RoleSlug.IsNull() {
		v := state.RoleSlug.ValueString()
		params.SetRole([]string{v})
		hasInput = true
	}
	if !state.RoleID.IsNull() {
		v := state.RoleID.ValueInt64()
		params.SetRoleID([]int64{v})
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
		v := state.SiteGroupID.ValueString()
		params.SetSiteGroupID([]string{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.VminterfaceID.IsNull() {
		v := state.VminterfaceID.ValueInt64()
		params.SetVminterfaceID(&v)
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_vlan.")
		return
	}
	res, err := d.client.Ipam.IpamVlansListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_vlan", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_vlan",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVlan(ctx, netboxapi.VlanResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.vlanResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
