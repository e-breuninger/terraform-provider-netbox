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
	SsidContains       types.String `tfsdk:"ssid_contains"`
	TagSlug            types.String `tfsdk:"tag_slug"`
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
				Computed: true,
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the wireless LAN group.",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "Operational status. One of: active, reserved, disabled, deprecated.",
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
				Computed:    true,
				Description: "Authentication type. One of: open, wep, wpa-personal, wpa-enterprise.",
			},
			"auth_cipher": schema.StringAttribute{
				Computed:    true,
				Description: "Authentication cipher. One of: auto, tkip, aes.",
			},
			"auth_psk": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "Pre-shared key.",
			},
			"scope_type": schema.StringAttribute{
				Computed:    true,
				Description: "Content type of the scope. Derived from site_id, location_id, region_id or site_group_id when one of those is set; set it together with scope_id otherwise. One of: dcim.site, dcim.location, dcim.region, dcim.sitegroup.",
			},
			"scope_id": schema.Int64Attribute{
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
			"ssid_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the SSID.",
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
	if !state.SsidContains.IsNull() {
		v := state.SsidContains.ValueString()
		params.SetSsidIc([]string{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or ssid and/or group_id and/or vlan_id and/or tenant_id and/or ssid_contains and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_wireless_lan.")
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
