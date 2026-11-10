// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*powerFeedDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*powerFeedDataSource)(nil)
)

// NewPowerFeedDataSource returns a new power_feed data source.
func NewPowerFeedDataSource() datasource.DataSource {
	return &powerFeedDataSource{}
}

type powerFeedDataSource struct {
	client *netboxapi.Client
}

// powerFeedDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type powerFeedDataSourceModel struct {
	powerFeedResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	AvailablePower     types.Int64  `tfsdk:"available_power"`
	CableConnector     types.Int64  `tfsdk:"cable_connector"`
	CableEnd           types.String `tfsdk:"cable_end"`
	CableID            types.Int64  `tfsdk:"cable_id"`
	Cabled             types.Bool   `tfsdk:"cabled"`
	Connected          types.Bool   `tfsdk:"connected"`
	MaxUtilization     types.Int64  `tfsdk:"max_utilization"`
	Occupied           types.Bool   `tfsdk:"occupied"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	Q                  types.String `tfsdk:"q"`
	RegionSlug         types.String `tfsdk:"region_slug"`
	RegionID           types.String `tfsdk:"region_id"`
	SiteSlug           types.String `tfsdk:"site_slug"`
	SiteGroupSlug      types.String `tfsdk:"site_group_slug"`
	SiteGroupID        types.String `tfsdk:"site_group_id"`
	SiteID             types.Int64  `tfsdk:"site_id"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String `tfsdk:"tenant_group_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *powerFeedDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_power_feed"
}

func (d *powerFeedDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A power feed from a panel to a rack (dcim.powerfeed).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/powerfeed/):\n\n> A power feed represents the distribution of power from a power panel to a particular device, typically a power distribution unit (PDU). The power port (inlet) on a device can be connected via a cable to a power feed. A power feed may optionally be assigned to a rack to allow more easily tracking the distribution of power among racks.",
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
			"power_panel_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the power panel the feed starts at.",
			},
			"rack_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the rack the feed serves.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: offline, active, planned, failed.",
				Validators: []validator.String{
					stringvalidator.OneOf("offline", "active", "planned", "failed"),
				},
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the feed is the primary or the redundant one. One of: primary, redundant.",
				Validators: []validator.String{
					stringvalidator.OneOf("primary", "redundant"),
				},
			},
			"supply": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Supply type. One of: ac, dc.",
				Validators: []validator.String{
					stringvalidator.OneOf("ac", "dc"),
				},
			},
			"phase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Phase type. One of: single-phase, three-phase.",
				Validators: []validator.String{
					stringvalidator.OneOf("single-phase", "three-phase"),
				},
			},
			"voltage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Voltage of the feed. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(-32768, 32767),
				},
			},
			"amperage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Amperage of the feed. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(1, 32767),
				},
			},
			"max_utilization_percent": schema.Int64Attribute{
				Computed:    true,
				Description: "Maximum permitted draw as a percentage. Must be a whole number.",
			},
			"max_percent_utilization": schema.Int64Attribute{
				Computed:           true,
				DeprecationMessage: "Use max_utilization_percent instead.",
				Description:        "Deprecated alias of max_utilization_percent.",
			},
			"mark_connected": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat the feed as connected even without a cable.",
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
			"available_power": schema.Int64Attribute{
				Optional:    true,
				Description: "Available power to match.",
			},
			"cable_connector": schema.Int64Attribute{
				Optional:    true,
				Description: "Cable connector to match.",
			},
			"cable_end": schema.StringAttribute{
				Optional:    true,
				Description: "Cable end. One of: A, B.",
			},
			"cable_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cable.",
			},
			"cabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether cabled is set.",
			},
			"connected": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether connected is set.",
			},
			"max_utilization": schema.Int64Attribute{
				Optional:    true,
				Description: "Max utilization to match.",
			},
			"occupied": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether occupied is set.",
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
			"region_id": schema.StringAttribute{
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
			"site_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the site group.",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site.",
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

func (d *powerFeedDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *powerFeedDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data powerFeedDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimPowerFeedsListParams()
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
	if !state.PowerPanelID.IsNull() {
		v := state.PowerPanelID.ValueInt64()
		params.SetPowerPanelID([]int64{v})
		hasInput = true
	}
	if !state.RackID.IsNull() {
		v := state.RackID.ValueInt64()
		params.SetRackID([]int64{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType(&v)
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
	if !state.Amperage.IsNull() {
		v := state.Amperage.ValueInt64()
		params.SetAmperage([]int64{v})
		hasInput = true
	}
	if !state.AvailablePower.IsNull() {
		v := state.AvailablePower.ValueInt64()
		params.SetAvailablePower([]int64{v})
		hasInput = true
	}
	if !state.CableConnector.IsNull() {
		v := state.CableConnector.ValueInt64()
		params.SetCableConnector([]int64{v})
		hasInput = true
	}
	if !state.CableEnd.IsNull() {
		v := state.CableEnd.ValueString()
		params.SetCableEnd(&v)
		hasInput = true
	}
	if !state.CableID.IsNull() {
		v := state.CableID.ValueInt64()
		params.SetCableID([]int64{v})
		hasInput = true
	}
	if !state.Cabled.IsNull() {
		v := state.Cabled.ValueBool()
		params.SetCabled(&v)
		hasInput = true
	}
	if !state.Connected.IsNull() {
		v := state.Connected.ValueBool()
		params.SetConnected(&v)
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
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.MarkConnected.IsNull() {
		v := state.MarkConnected.ValueBool()
		params.SetMarkConnected(&v)
		hasInput = true
	}
	if !state.MaxUtilization.IsNull() {
		v := state.MaxUtilization.ValueInt64()
		params.SetMaxUtilization([]int64{v})
		hasInput = true
	}
	if !state.Occupied.IsNull() {
		v := state.Occupied.ValueBool()
		params.SetOccupied(&v)
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
	if !state.Phase.IsNull() {
		v := state.Phase.ValueString()
		params.SetPhase(&v)
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
		v := state.RegionID.ValueString()
		params.SetRegionID([]string{v})
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
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.Supply.IsNull() {
		v := state.Supply.ValueString()
		params.SetSupply(&v)
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
	if !state.Voltage.IsNull() {
		v := state.Voltage.ValueInt64()
		params.SetVoltage([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_power_feed.")
		return
	}
	res, err := d.client.Dcim.DcimPowerFeedsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_power_feed", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_power_feed",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenPowerFeed(ctx, netboxapi.PowerFeedResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.powerFeedResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
