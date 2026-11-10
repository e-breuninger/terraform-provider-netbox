// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*circuitTerminationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*circuitTerminationDataSource)(nil)
)

// NewCircuitTerminationDataSource returns a new circuit_termination data source.
func NewCircuitTerminationDataSource() datasource.DataSource {
	return &circuitTerminationDataSource{}
}

type circuitTerminationDataSource struct {
	client *netboxapi.Client
}

// circuitTerminationDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type circuitTerminationDataSourceModel struct {
	circuitTerminationResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	CableConnector     types.Int64  `tfsdk:"cable_connector"`
	CableEnd           types.String `tfsdk:"cable_end"`
	CableID            types.Int64  `tfsdk:"cable_id"`
	Cabled             types.Bool   `tfsdk:"cabled"`
	LocationSlug       types.String `tfsdk:"location_slug"`
	Occupied           types.Bool   `tfsdk:"occupied"`
	ProviderSlug       types.String `tfsdk:"provider_slug"`
	ProviderID         types.Int64  `tfsdk:"provider_id"`
	Q                  types.String `tfsdk:"q"`
	RegionSlug         types.String `tfsdk:"region_slug"`
	SiteSlug           types.String `tfsdk:"site_slug"`
	SiteGroupSlug      types.String `tfsdk:"site_group_slug"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *circuitTerminationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_circuit_termination"
}

func (d *circuitTerminationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:A termination attaching one side of a circuit to a site or provider network (circuits.circuittermination).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuit-terminations):\n\n> The association of a circuit with a particular site and/or device is modeled separately as a circuit termination. A circuit may have up to two terminations, labeled A and Z. A single-termination circuit can be used when you don't know (or care) about the far end of a circuit (for example, an Internet access circuit which connects to a transit provider). A dual-termination circuit is useful for tracking circuits which connect two sites.\n>\n> Each circuit termination is attached to either a site or to a provider network. Site terminations may optionally be connected via a cable to a specific device interface or port within that site. Each termination must be assigned a port speed, and can optionally be assigned an upstream speed if it differs from the downstream speed (a common scenario with e.g. DOCSIS cable modems). Fields are also available to track cross-connect and patch panel details.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"circuit_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the circuit.",
			},
			"term_side": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Side of the circuit this terminates. One of: A, Z.",
				Validators: []validator.String{
					stringvalidator.OneOf("A", "Z"),
				},
			},
			"termination_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the termination. Derived from site_id, location_id, region_id, site_group_id or provider_network_id when one of those is set; set it together with termination_id otherwise. One termination is required. One of: dcim.site, dcim.location, dcim.region, dcim.sitegroup, circuits.providernetwork.",
				Validators: []validator.String{
					stringvalidator.OneOf("dcim.site", "dcim.location", "dcim.region", "dcim.sitegroup", "circuits.providernetwork"),
				},
			},
			"termination_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating object (see termination_type).",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating site (termination_type dcim.site). Conflicts with the other aliases and with setting the termination_* pair directly.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating location (termination_type dcim.location). Conflicts with the other aliases and with setting the termination_* pair directly.",
			},
			"region_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating region (termination_type dcim.region). Conflicts with the other aliases and with setting the termination_* pair directly.",
			},
			"site_group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating site group (termination_type dcim.sitegroup). Conflicts with the other aliases and with setting the termination_* pair directly.",
			},
			"provider_network_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating provider network (termination_type circuits.providernetwork). Conflicts with the other aliases and with setting the termination_* pair directly.",
			},
			"port_speed_kbps": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Physical circuit speed in kbps. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"port_speed": schema.Int64Attribute{
				Computed:           true,
				DeprecationMessage: "Use port_speed_kbps instead.",
				Description:        "Deprecated alias of port_speed_kbps.",
			},
			"upstream_speed_kbps": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Upstream speed in kbps if different from the port speed. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"upstream_speed": schema.Int64Attribute{
				Computed:           true,
				DeprecationMessage: "Use upstream_speed_kbps instead.",
				Description:        "Deprecated alias of upstream_speed_kbps.",
			},
			"xconnect_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Cross-connect id assigned by the provider.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"pp_info": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Patch panel and port information.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(100),
				},
			},
			"mark_connected": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Treat the termination as connected even without a cable.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
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
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"occupied": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether occupied is set.",
			},
			"provider_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the provider.",
			},
			"provider_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the provider.",
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
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *circuitTerminationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *circuitTerminationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data circuitTerminationDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := circuits.NewCircuitsCircuitTerminationsListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.CircuitID.IsNull() {
		v := state.CircuitID.ValueInt64()
		params.SetCircuitID([]int64{v})
		hasInput = true
	}
	if !state.TermSide.IsNull() {
		v := state.TermSide.ValueString()
		params.SetTermSide(&v)
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.ProviderNetworkID.IsNull() {
		v := state.ProviderNetworkID.ValueInt64()
		params.SetProviderNetworkID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
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
	if !state.MarkConnected.IsNull() {
		v := state.MarkConnected.ValueBool()
		params.SetMarkConnected(&v)
		hasInput = true
	}
	if !state.Occupied.IsNull() {
		v := state.Occupied.ValueBool()
		params.SetOccupied(&v)
		hasInput = true
	}
	if !state.PortSpeedKbps.IsNull() {
		v := state.PortSpeedKbps.ValueInt64()
		params.SetPortSpeed([]int64{v})
		hasInput = true
	}
	if !state.PpInfo.IsNull() {
		v := state.PpInfo.ValueString()
		params.SetPpInfo([]string{v})
		hasInput = true
	}
	if !state.ProviderSlug.IsNull() {
		v := state.ProviderSlug.ValueString()
		params.SetProvider([]string{v})
		hasInput = true
	}
	if !state.ProviderID.IsNull() {
		v := state.ProviderID.ValueInt64()
		params.SetProviderID([]int64{v})
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
	if !state.TerminationID.IsNull() {
		v := state.TerminationID.ValueInt64()
		params.SetTerminationID([]int64{v})
		hasInput = true
	}
	if !state.TerminationType.IsNull() {
		v := state.TerminationType.ValueString()
		params.SetTerminationType([]string{v})
		hasInput = true
	}
	if !state.UpstreamSpeedKbps.IsNull() {
		v := state.UpstreamSpeedKbps.ValueInt64()
		params.SetUpstreamSpeed([]int64{v})
		hasInput = true
	}
	if !state.XconnectID.IsNull() {
		v := state.XconnectID.ValueString()
		params.SetXconnectID([]string{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_circuit_termination.")
		return
	}
	res, err := d.client.Circuits.CircuitsCircuitTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_circuit_termination", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_circuit_termination",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCircuitTermination(ctx, netboxapi.CircuitTerminationResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.circuitTerminationResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&circuitTerminationResource{client: d.client}).postReadHook(ctx, &state.circuitTerminationResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
