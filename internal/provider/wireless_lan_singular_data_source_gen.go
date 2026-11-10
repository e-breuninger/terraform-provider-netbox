// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*wirelessLanDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*wirelessLanDataSource)(nil)
)

// NewWirelessLanDataSource returns a new wireless_lan data source.
func NewWirelessLanDataSource() datasource.DataSource {
	return &wirelessLanDataSource{}
}

type wirelessLanDataSource struct {
	client *netboxapi.Client
}

// wirelessLanDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type wirelessLanDataSourceModel struct {
	wirelessLanResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	Group              types.String `tfsdk:"group"`
	InterfaceID        types.Int64  `tfsdk:"interface_id"`
	LocationSlug       types.String `tfsdk:"location_slug"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	Q                  types.String `tfsdk:"q"`
	RegionSlug         types.String `tfsdk:"region_slug"`
	SiteSlug           types.String `tfsdk:"site_slug"`
	SiteGroupSlug      types.String `tfsdk:"site_group_slug"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String `tfsdk:"tenant_group_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *wirelessLanDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_lan"
}

func (d *wirelessLanDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Wireless:A NetBox wireless LAN (wireless.wirelesslan): an SSID with its authentication and VLAN.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/wireless/wirelesslan/):\n\n> A wireless LAN is a set of interfaces connected via a common wireless channel, identified by its SSID and authentication parameters. Wireless [interfaces](https://netboxlabs.com/docs/netbox/models/dcim/interface/) can be associated with wireless LANs to model multi-acess wireless segments.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Service set identifier.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 32),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the wireless LAN group.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: active, reserved, disabled, deprecated.",
				Validators: []validator.String{
					stringvalidator.OneOf("active", "reserved", "disabled", "deprecated"),
				},
			},
			"vlan_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VLAN the wireless LAN is bridged to.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"auth_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication type. One of: open, wep, wpa-personal, wpa-enterprise.",
				Validators: []validator.String{
					stringvalidator.OneOf("open", "wep", "wpa-personal", "wpa-enterprise"),
				},
			},
			"auth_cipher": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Authentication cipher. One of: auto, tkip, aes.",
				Validators: []validator.String{
					stringvalidator.OneOf("auto", "tkip", "aes"),
				},
			},
			"auth_psk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Pre-shared key.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
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
			"group": schema.StringAttribute{
				Optional:    true,
				Description: "Group to match.",
			},
			"interface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the interface.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
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
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
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
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *wirelessLanDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *wirelessLanDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data wirelessLanDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := wireless.NewWirelessWirelessLansListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Ssid.IsNull() {
		v := state.Ssid.ValueString()
		params.SetSsid([]string{v})
		hasInput = true
	}
	if !state.GroupID.IsNull() {
		v := strconv.FormatInt(state.GroupID.ValueInt64(), 10)
		params.SetGroupID([]string{v})
		hasInput = true
	}
	if !state.VlanID.IsNull() {
		v := state.VlanID.ValueInt64()
		params.SetVlanID([]int64{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
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
	if !state.AuthCipher.IsNull() {
		v := state.AuthCipher.ValueString()
		params.SetAuthCipher([]string{v})
		hasInput = true
	}
	if !state.AuthPsk.IsNull() {
		v := state.AuthPsk.ValueString()
		params.SetAuthPsk([]string{v})
		hasInput = true
	}
	if !state.AuthType.IsNull() {
		v := state.AuthType.ValueString()
		params.SetAuthType([]string{v})
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
	if !state.Group.IsNull() {
		v := state.Group.ValueString()
		params.SetGroup([]string{v})
		hasInput = true
	}
	if !state.InterfaceID.IsNull() {
		v := state.InterfaceID.ValueInt64()
		params.SetInterfaceID([]int64{v})
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
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := strconv.FormatInt(state.RegionID.ValueInt64(), 10)
		params.SetRegionID([]string{v})
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
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_wireless_lan.")
		return
	}
	res, err := d.client.Wireless.WirelessWirelessLansListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_wireless_lan", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_wireless_lan",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenWirelessLan(ctx, netboxapi.WirelessLanResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.wirelessLanResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&wirelessLanResource{client: d.client}).postReadHook(ctx, &state.wirelessLanResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
