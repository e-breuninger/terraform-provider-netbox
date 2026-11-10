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
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*prefixDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*prefixDataSource)(nil)
)

// NewPrefixDataSource returns a new prefix data source.
func NewPrefixDataSource() datasource.DataSource {
	return &prefixDataSource{}
}

type prefixDataSource struct {
	client *netboxapi.Client
}

// prefixDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type prefixDataSourceModel struct {
	prefixResourceModel
	PrefixContains     types.String  `tfsdk:"prefix_contains"`
	VlanVid            types.Int64   `tfsdk:"vlan_vid"`
	TagSlug            types.String  `tfsdk:"tag_slug"`
	Children           types.Int64   `tfsdk:"children"`
	Contact            types.Int64   `tfsdk:"contact"`
	ContactGroupSlug   types.String  `tfsdk:"contact_group_slug"`
	ContactRole        types.Int64   `tfsdk:"contact_role"`
	Depth              types.Int64   `tfsdk:"depth"`
	LocationSlug       types.String  `tfsdk:"location_slug"`
	MaskLengthGte      types.Float64 `tfsdk:"mask_length__gte"`
	MaskLengthLte      types.Float64 `tfsdk:"mask_length__lte"`
	MaskLength         types.Int64   `tfsdk:"mask_length"`
	OwnerName          types.String  `tfsdk:"owner_name"`
	OwnerGroupName     types.String  `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64   `tfsdk:"owner_group_id"`
	PresentInVrf       types.String  `tfsdk:"present_in_vrf"`
	PresentInVrfID     types.String  `tfsdk:"present_in_vrf_id"`
	Q                  types.String  `tfsdk:"q"`
	RegionSlug         types.String  `tfsdk:"region_slug"`
	RoleSlug           types.String  `tfsdk:"role_slug"`
	SiteSlug           types.String  `tfsdk:"site_slug"`
	SiteGroupSlug      types.String  `tfsdk:"site_group_slug"`
	TagID              types.Int64   `tfsdk:"tag_id"`
	TenantSlug         types.String  `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String  `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String  `tfsdk:"tenant_group_id"`
	VlanGroupSlug      types.String  `tfsdk:"vlan_group_slug"`
	VlanGroupID        types.Int64   `tfsdk:"vlan_group_id"`
	Vrf                types.String  `tfsdk:"vrf"`
	Within             types.String  `tfsdk:"within"`
	WithinInclude      types.String  `tfsdk:"within_include"`
	CustomFieldFilters types.Map     `tfsdk:"custom_field_filters"`
}

func (d *prefixDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_prefix"
}

func (d *prefixDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox prefix (ipam.prefix).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#prefixes):\n\n> A prefix is an IPv4 or IPv6 network and mask expressed in CIDR notation (e.g. 192.0.2.0/24). A prefix entails only the \"network portion\" of an IP address: All bits in the address not covered by the mask must be zero. (In other words, a prefix cannot be a specific IP address.)\n>\n> Prefixes are automatically organized by their parent aggregates. Additionally, each prefix can be assigned to a particular site and virtual routing and forwarding instance (VRF). Each VRF represents a separate IP space or routing table. All prefixes not assigned to a VRF are considered to be in the \"global\" table.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"prefix": schema.StringAttribute{
				CustomType:  conv.CIDRType{},
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 network with mask, e.g. 10.0.0.0/24.",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: container, active, reserved, deprecated.",
				Validators: []validator.String{
					stringvalidator.OneOf("container", "active", "reserved", "deprecated"),
				},
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the role.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"vlan_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VLAN.",
			},
			"vrf_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VRF.",
			},
			"is_pool": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "All addresses within this prefix are usable.",
			},
			"mark_utilized": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Report space as fully utilized.",
			},
			"family": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Address family: 4 or 6.",
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
				Optional:    true,
				Computed:    true,
				Description: "Id of the site the object is scoped to (scope_type dcim.site). Conflicts with the other scope aliases and with setting the scope_* pair directly.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the location the object is scoped to (scope_type dcim.location).",
			},
			"region_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the region the object is scoped to (scope_type dcim.region).",
			},
			"site_group_id": schema.Int64Attribute{
				Optional:    true,
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
			"prefix_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Prefixes containing this prefix or IP.",
			},
			"vlan_vid": schema.Int64Attribute{
				Optional:    true,
				Description: "VLAN number (vid) of the VLAN the prefix is assigned to.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"children": schema.Int64Attribute{
				Optional:    true,
				Description: "Children to match.",
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
			"depth": schema.Int64Attribute{
				Optional:    true,
				Description: "Depth to match.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"mask_length__gte": schema.Float64Attribute{
				Optional:    true,
				Description: "Mask length is at least this.",
			},
			"mask_length__lte": schema.Float64Attribute{
				Optional:    true,
				Description: "Mask length is at most this.",
			},
			"mask_length": schema.Int64Attribute{
				Optional:    true,
				Description: "Mask length to match.",
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
			"present_in_vrf": schema.StringAttribute{
				Optional:    true,
				Description: "Present in vrf to match.",
			},
			"present_in_vrf_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the present in vrf.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
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
			"vlan_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the VLAN group.",
			},
			"vlan_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the VLAN group.",
			},
			"vrf": schema.StringAttribute{
				Optional:    true,
				Description: "VRF (RD).",
			},
			"within": schema.StringAttribute{
				Optional:    true,
				Description: "Within prefix.",
			},
			"within_include": schema.StringAttribute{
				Optional:    true,
				Description: "Within and including prefix.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *prefixDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *prefixDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data prefixDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamPrefixesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Prefix.IsNull() {
		v := state.Prefix.ValueString()
		params.SetPrefix([]string{v})
		hasInput = true
	}
	if !state.VrfID.IsNull() {
		v := state.VrfID.ValueInt64()
		params.SetVrfID([]int64{v})
		hasInput = true
	}
	if !state.PrefixContains.IsNull() {
		v := state.PrefixContains.ValueString()
		params.SetContains(&v)
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.RoleID.IsNull() {
		v := state.RoleID.ValueInt64()
		params.SetRoleID([]int64{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.VlanID.IsNull() {
		v := state.VlanID.ValueInt64()
		params.SetVlanID([]int64{v})
		hasInput = true
	}
	if !state.VlanVid.IsNull() {
		v := state.VlanVid.ValueInt64()
		params.SetVlanVid(&v)
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := strconv.FormatInt(state.RegionID.ValueInt64(), 10)
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
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
	if !state.Children.IsNull() {
		v := state.Children.ValueInt64()
		params.SetChildren([]int64{v})
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
	if !state.Depth.IsNull() {
		v := state.Depth.ValueInt64()
		params.SetDepth([]int64{v})
		hasInput = true
	}
	if !state.Family.IsNull() {
		v := float64(state.Family.ValueInt64())
		params.SetFamily(&v)
		hasInput = true
	}
	if !state.IsPool.IsNull() {
		v := state.IsPool.ValueBool()
		params.SetIsPool(&v)
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
		v := strconv.FormatInt(state.LocationID.ValueInt64(), 10)
		params.SetLocationID([]string{v})
		hasInput = true
	}
	if !state.MarkUtilized.IsNull() {
		v := state.MarkUtilized.ValueBool()
		params.SetMarkUtilized(&v)
		hasInput = true
	}
	if !state.MaskLengthGte.IsNull() {
		v := state.MaskLengthGte.ValueFloat64()
		params.SetMaskLengthGte(&v)
		hasInput = true
	}
	if !state.MaskLengthLte.IsNull() {
		v := state.MaskLengthLte.ValueFloat64()
		params.SetMaskLengthLte(&v)
		hasInput = true
	}
	if !state.MaskLength.IsNull() {
		v := state.MaskLength.ValueInt64()
		params.SetMaskLength([]int64{v})
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
	if !state.PresentInVrf.IsNull() {
		v := state.PresentInVrf.ValueString()
		params.SetPresentInVrf(&v)
		hasInput = true
	}
	if !state.PresentInVrfID.IsNull() {
		v := state.PresentInVrfID.ValueString()
		params.SetPresentInVrfID(&v)
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
	if !state.RoleSlug.IsNull() {
		v := state.RoleSlug.ValueString()
		params.SetRole([]string{v})
		hasInput = true
	}
	if !state.ScopeID.IsNull() {
		v := state.ScopeID.ValueInt64()
		params.SetScopeID([]int64{v})
		hasInput = true
	}
	if !state.ScopeType.IsNull() {
		v := state.ScopeType.ValueString()
		params.SetScopeType([]string{v})
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
		v := strconv.FormatInt(state.SiteGroupID.ValueInt64(), 10)
		params.SetSiteGroupID([]string{v})
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
	if !state.VlanGroupSlug.IsNull() {
		v := state.VlanGroupSlug.ValueString()
		params.SetVlanGroup([]string{v})
		hasInput = true
	}
	if !state.VlanGroupID.IsNull() {
		v := state.VlanGroupID.ValueInt64()
		params.SetVlanGroupID([]int64{v})
		hasInput = true
	}
	if !state.Vrf.IsNull() {
		v := state.Vrf.ValueString()
		params.SetVrf([]string{v})
		hasInput = true
	}
	if !state.Within.IsNull() {
		v := state.Within.ValueString()
		params.SetWithin(&v)
		hasInput = true
	}
	if !state.WithinInclude.IsNull() {
		v := state.WithinInclude.ValueString()
		params.SetWithinInclude(&v)
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_prefix.")
		return
	}
	res, err := d.client.Ipam.IpamPrefixesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_prefix", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_prefix",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenPrefix(ctx, netboxapi.PrefixResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.prefixResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&prefixResource{client: d.client}).postReadHook(ctx, &state.prefixResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
