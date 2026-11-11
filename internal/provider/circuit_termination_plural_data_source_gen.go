// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*circuitTerminationsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*circuitTerminationsDataSource)(nil)
)

// NewCircuitTerminationsDataSource returns a new circuit_terminations data source, which
// lists circuit_termination objects matching its filters.
func NewCircuitTerminationsDataSource() datasource.DataSource {
	return &circuitTerminationsDataSource{}
}

type circuitTerminationsDataSource struct {
	client *netboxapi.Client
}

// circuitTerminationsDataSourceModel is the data source model: the filters, the limit and the matching
// circuit_terminations.
type circuitTerminationsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"circuit_terminations"`
}

// circuitTerminationResourceAttrTypes is the attribute type map of circuitTerminationResourceModel.
func circuitTerminationResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                  types.Int64Type,
		"circuit_id":          types.Int64Type,
		"term_side":           types.StringType,
		"termination_type":    types.StringType,
		"termination_id":      types.Int64Type,
		"site_id":             types.Int64Type,
		"location_id":         types.Int64Type,
		"region_id":           types.Int64Type,
		"site_group_id":       types.Int64Type,
		"provider_network_id": types.Int64Type,
		"port_speed_kbps":     types.Int64Type,
		"port_speed":          types.Int64Type,
		"upstream_speed_kbps": types.Int64Type,
		"upstream_speed":      types.Int64Type,
		"xconnect_id":         types.StringType,
		"pp_info":             types.StringType,
		"mark_connected":      types.BoolType,
		"description":         types.StringType,
		"created":             types.StringType,
		"last_updated":        types.StringType,
		"url":                 types.StringType,
		"tags":                types.SetType{ElemType: types.StringType},
		"tags_all":            types.SetType{ElemType: types.StringType},
		"custom_fields":       types.MapType{ElemType: types.StringType},
	}
}

func (d *circuitTerminationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_circuit_terminations"
}

func (d *circuitTerminationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:Lists circuit_termination objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: cable_connector, cable_connector__empty, cable_connector__gt, cable_connector__gte, cable_connector__lt, cable_connector__lte, cable_connector__n, cable_end, cable_end__empty, cable_end__ic, cable_end__ie, cable_end__iew, cable_end__iregex, cable_end__isw, cable_end__n, cable_end__nic, cable_end__nie, cable_end__niew, cable_end__nisw, cable_end__regex, cable_id, cable_id__n, cabled, circuit_id, circuit_id__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, mark_connected, occupied, port_speed, port_speed__empty, port_speed__gt, port_speed__gte, port_speed__lt, port_speed__lte, port_speed__n, pp_info, pp_info__empty, pp_info__ic, pp_info__ie, pp_info__iew, pp_info__iregex, pp_info__isw, pp_info__n, pp_info__nic, pp_info__nie, pp_info__niew, pp_info__nisw, pp_info__regex, provider, provider__n, provider_id, provider_id__n, provider_network_id, provider_network_id__n, q, region, region__n, region_id, region_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, term_side, term_side__empty, term_side__ic, term_side__ie, term_side__iew, term_side__iregex, term_side__isw, term_side__n, term_side__nic, term_side__nie, term_side__niew, term_side__nisw, term_side__regex, termination_id, termination_id__empty, termination_id__gt, termination_id__gte, termination_id__lt, termination_id__lte, termination_id__n, termination_type, termination_type__n, upstream_speed, upstream_speed__empty, upstream_speed__gt, upstream_speed__gte, upstream_speed__lt, upstream_speed__lte, upstream_speed__n, xconnect_id, xconnect_id__empty, xconnect_id__ic, xconnect_id__ie, xconnect_id__iew, xconnect_id__iregex, xconnect_id__isw, xconnect_id__n, xconnect_id__nic, xconnect_id__nie, xconnect_id__niew, xconnect_id__nisw, xconnect_id__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"circuit_terminations": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching circuit_terminations.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"circuit_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the circuit.",
						},
						"term_side": schema.StringAttribute{
							Computed:    true,
							Description: "Side of the circuit this terminates. One of: A, Z.",
						},
						"termination_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the termination. Derived from site_id, location_id, region_id, site_group_id or provider_network_id when one of those is set; set it together with termination_id otherwise. One termination is required. One of: dcim.site, dcim.location, dcim.region, dcim.sitegroup, circuits.providernetwork.",
						},
						"termination_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating object (see termination_type).",
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating site (termination_type dcim.site). Conflicts with the other aliases and with setting the termination_* pair directly.",
						},
						"location_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating location (termination_type dcim.location). Conflicts with the other aliases and with setting the termination_* pair directly.",
						},
						"region_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating region (termination_type dcim.region). Conflicts with the other aliases and with setting the termination_* pair directly.",
						},
						"site_group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating site group (termination_type dcim.sitegroup). Conflicts with the other aliases and with setting the termination_* pair directly.",
						},
						"provider_network_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating provider network (termination_type circuits.providernetwork). Conflicts with the other aliases and with setting the termination_* pair directly.",
						},
						"port_speed_kbps": schema.Int64Attribute{
							Computed:    true,
							Description: "Physical circuit speed in kbps. Must be a whole number.",
						},
						"port_speed": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use port_speed_kbps instead.",
							Description:        "Deprecated alias of port_speed_kbps.",
						},
						"upstream_speed_kbps": schema.Int64Attribute{
							Computed:    true,
							Description: "Upstream speed in kbps if different from the port speed. Must be a whole number.",
						},
						"upstream_speed": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use upstream_speed_kbps instead.",
							Description:        "Deprecated alias of upstream_speed_kbps.",
						},
						"xconnect_id": schema.StringAttribute{
							Computed:    true,
							Description: "Cross-connect id assigned by the provider.",
						},
						"pp_info": schema.StringAttribute{
							Computed:    true,
							Description: "Patch panel and port information.",
						},
						"mark_connected": schema.BoolAttribute{
							Computed:    true,
							Description: "Treat the termination as connected even without a cable.",
						},
						"description": schema.StringAttribute{
							Computed: true,
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *circuitTerminationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *circuitTerminationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data circuitTerminationsDataSourceModel

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
	var responseDTOs []*netboxapi.CircuitTerminationResponseDTO

	params := circuits.NewCircuitsCircuitTerminationsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                     netboxapi.FilterInt64s(params.SetID),
		"id__n":                  netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                 netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                 netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":              netboxapi.FilterBool(params.SetIDEmpty),
		"circuit_id":             netboxapi.FilterInt64s(params.SetCircuitID),
		"circuit_id__n":          netboxapi.FilterInt64s(params.SetCircuitIDn),
		"term_side":              netboxapi.FilterString(params.SetTermSide),
		"term_side__n":           netboxapi.FilterString(params.SetTermSiden),
		"term_side__ic":          netboxapi.FilterStrings(params.SetTermSideIc),
		"term_side__nic":         netboxapi.FilterStrings(params.SetTermSideNic),
		"term_side__ie":          netboxapi.FilterStrings(params.SetTermSideIe),
		"term_side__nie":         netboxapi.FilterStrings(params.SetTermSideNie),
		"term_side__isw":         netboxapi.FilterStrings(params.SetTermSideIsw),
		"term_side__nisw":        netboxapi.FilterStrings(params.SetTermSideNisw),
		"term_side__iew":         netboxapi.FilterStrings(params.SetTermSideIew),
		"term_side__niew":        netboxapi.FilterStrings(params.SetTermSideNiew),
		"term_side__empty":       netboxapi.FilterBool(params.SetTermSideEmpty),
		"term_side__regex":       netboxapi.FilterStrings(params.SetTermSideRegex),
		"term_side__iregex":      netboxapi.FilterStrings(params.SetTermSideIregex),
		"site_id":                netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":             netboxapi.FilterInt64s(params.SetSiteIDn),
		"provider_network_id":    netboxapi.FilterInt64s(params.SetProviderNetworkID),
		"provider_network_id__n": netboxapi.FilterInt64s(params.SetProviderNetworkIDn),
		"tag":                    netboxapi.FilterStrings(params.SetTag),
		"tag__n":                 netboxapi.FilterStrings(params.SetTagn),
		"tag__any":               netboxapi.FilterStrings(params.SetTagAny),
		"cable_connector":        netboxapi.FilterInt64s(params.SetCableConnector),
		"cable_connector__n":     netboxapi.FilterInt64s(params.SetCableConnectorn),
		"cable_connector__lt":    netboxapi.FilterInt64s(params.SetCableConnectorLt),
		"cable_connector__lte":   netboxapi.FilterInt64s(params.SetCableConnectorLte),
		"cable_connector__gt":    netboxapi.FilterInt64s(params.SetCableConnectorGt),
		"cable_connector__gte":   netboxapi.FilterInt64s(params.SetCableConnectorGte),
		"cable_connector__empty": netboxapi.FilterBool(params.SetCableConnectorEmpty),
		"cable_end":              netboxapi.FilterString(params.SetCableEnd),
		"cable_end__n":           netboxapi.FilterString(params.SetCableEndn),
		"cable_end__ic":          netboxapi.FilterStrings(params.SetCableEndIc),
		"cable_end__nic":         netboxapi.FilterStrings(params.SetCableEndNic),
		"cable_end__ie":          netboxapi.FilterStrings(params.SetCableEndIe),
		"cable_end__nie":         netboxapi.FilterStrings(params.SetCableEndNie),
		"cable_end__isw":         netboxapi.FilterStrings(params.SetCableEndIsw),
		"cable_end__nisw":        netboxapi.FilterStrings(params.SetCableEndNisw),
		"cable_end__iew":         netboxapi.FilterStrings(params.SetCableEndIew),
		"cable_end__niew":        netboxapi.FilterStrings(params.SetCableEndNiew),
		"cable_end__empty":       netboxapi.FilterBool(params.SetCableEndEmpty),
		"cable_end__regex":       netboxapi.FilterStrings(params.SetCableEndRegex),
		"cable_end__iregex":      netboxapi.FilterStrings(params.SetCableEndIregex),
		"cable_id":               netboxapi.FilterInt64s(params.SetCableID),
		"cable_id__n":            netboxapi.FilterInt64s(params.SetCableIDn),
		"cabled":                 netboxapi.FilterBool(params.SetCabled),
		"created":                netboxapi.FilterStrings(params.SetCreated),
		"created__n":             netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":            netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":           netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":            netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":           netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":         netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":            netboxapi.FilterStrings(params.SetDescription),
		"description__n":         netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":        netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":       netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":        netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":       netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":       netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":      netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":       netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":      netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":     netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":     netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":    netboxapi.FilterStrings(params.SetDescriptionIregex),
		"last_updated":           netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":        netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":       netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":      netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":       netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":      netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":    netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":               netboxapi.FilterStrings(params.SetLocation),
		"location__n":            netboxapi.FilterStrings(params.SetLocationn),
		"location_id":            netboxapi.FilterStrings(params.SetLocationID),
		"location_id__n":         netboxapi.FilterStrings(params.SetLocationIDn),
		"mark_connected":         netboxapi.FilterBool(params.SetMarkConnected),
		"occupied":               netboxapi.FilterBool(params.SetOccupied),
		"port_speed":             netboxapi.FilterInt64s(params.SetPortSpeed),
		"port_speed__n":          netboxapi.FilterInt64s(params.SetPortSpeedn),
		"port_speed__lt":         netboxapi.FilterInt64s(params.SetPortSpeedLt),
		"port_speed__lte":        netboxapi.FilterInt64s(params.SetPortSpeedLte),
		"port_speed__gt":         netboxapi.FilterInt64s(params.SetPortSpeedGt),
		"port_speed__gte":        netboxapi.FilterInt64s(params.SetPortSpeedGte),
		"port_speed__empty":      netboxapi.FilterBool(params.SetPortSpeedEmpty),
		"pp_info":                netboxapi.FilterStrings(params.SetPpInfo),
		"pp_info__n":             netboxapi.FilterStrings(params.SetPpInfon),
		"pp_info__ic":            netboxapi.FilterStrings(params.SetPpInfoIc),
		"pp_info__nic":           netboxapi.FilterStrings(params.SetPpInfoNic),
		"pp_info__ie":            netboxapi.FilterStrings(params.SetPpInfoIe),
		"pp_info__nie":           netboxapi.FilterStrings(params.SetPpInfoNie),
		"pp_info__isw":           netboxapi.FilterStrings(params.SetPpInfoIsw),
		"pp_info__nisw":          netboxapi.FilterStrings(params.SetPpInfoNisw),
		"pp_info__iew":           netboxapi.FilterStrings(params.SetPpInfoIew),
		"pp_info__niew":          netboxapi.FilterStrings(params.SetPpInfoNiew),
		"pp_info__empty":         netboxapi.FilterBool(params.SetPpInfoEmpty),
		"pp_info__regex":         netboxapi.FilterStrings(params.SetPpInfoRegex),
		"pp_info__iregex":        netboxapi.FilterStrings(params.SetPpInfoIregex),
		"provider":               netboxapi.FilterStrings(params.SetProvider),
		"provider__n":            netboxapi.FilterStrings(params.SetProvidern),
		"provider_id":            netboxapi.FilterInt64s(params.SetProviderID),
		"provider_id__n":         netboxapi.FilterInt64s(params.SetProviderIDn),
		"q":                      netboxapi.FilterString(params.SetQ),
		"region":                 netboxapi.FilterStrings(params.SetRegion),
		"region__n":              netboxapi.FilterStrings(params.SetRegionn),
		"region_id":              netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":           netboxapi.FilterStrings(params.SetRegionIDn),
		"site":                   netboxapi.FilterStrings(params.SetSite),
		"site__n":                netboxapi.FilterStrings(params.SetSiten),
		"site_group":             netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":          netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":          netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":       netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"tag_id":                 netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":              netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":            netboxapi.FilterInt64s(params.SetTagIDAny),
		"termination_id":         netboxapi.FilterInt64s(params.SetTerminationID),
		"termination_id__n":      netboxapi.FilterInt64s(params.SetTerminationIDn),
		"termination_id__lt":     netboxapi.FilterInt64s(params.SetTerminationIDLt),
		"termination_id__lte":    netboxapi.FilterInt64s(params.SetTerminationIDLte),
		"termination_id__gt":     netboxapi.FilterInt64s(params.SetTerminationIDGt),
		"termination_id__gte":    netboxapi.FilterInt64s(params.SetTerminationIDGte),
		"termination_id__empty":  netboxapi.FilterBool(params.SetTerminationIDEmpty),
		"termination_type":       netboxapi.FilterStrings(params.SetTerminationType),
		"termination_type__n":    netboxapi.FilterStrings(params.SetTerminationTypen),
		"upstream_speed":         netboxapi.FilterInt64s(params.SetUpstreamSpeed),
		"upstream_speed__n":      netboxapi.FilterInt64s(params.SetUpstreamSpeedn),
		"upstream_speed__lt":     netboxapi.FilterInt64s(params.SetUpstreamSpeedLt),
		"upstream_speed__lte":    netboxapi.FilterInt64s(params.SetUpstreamSpeedLte),
		"upstream_speed__gt":     netboxapi.FilterInt64s(params.SetUpstreamSpeedGt),
		"upstream_speed__gte":    netboxapi.FilterInt64s(params.SetUpstreamSpeedGte),
		"upstream_speed__empty":  netboxapi.FilterBool(params.SetUpstreamSpeedEmpty),
		"xconnect_id":            netboxapi.FilterStrings(params.SetXconnectID),
		"xconnect_id__n":         netboxapi.FilterStrings(params.SetXconnectIDn),
		"xconnect_id__ic":        netboxapi.FilterStrings(params.SetXconnectIDIc),
		"xconnect_id__nic":       netboxapi.FilterStrings(params.SetXconnectIDNic),
		"xconnect_id__ie":        netboxapi.FilterStrings(params.SetXconnectIDIe),
		"xconnect_id__nie":       netboxapi.FilterStrings(params.SetXconnectIDNie),
		"xconnect_id__isw":       netboxapi.FilterStrings(params.SetXconnectIDIsw),
		"xconnect_id__nisw":      netboxapi.FilterStrings(params.SetXconnectIDNisw),
		"xconnect_id__iew":       netboxapi.FilterStrings(params.SetXconnectIDIew),
		"xconnect_id__niew":      netboxapi.FilterStrings(params.SetXconnectIDNiew),
		"xconnect_id__empty":     netboxapi.FilterBool(params.SetXconnectIDEmpty),
		"xconnect_id__regex":     netboxapi.FilterStrings(params.SetXconnectIDRegex),
		"xconnect_id__iregex":    netboxapi.FilterStrings(params.SetXconnectIDIregex),
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
		res, err := d.client.Circuits.CircuitsCircuitTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_circuit_terminations", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CircuitTerminationResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) runs on every item, as on the resource
	// itself (see the singular data source).
	resourceWithHooks := &circuitTerminationResource{client: d.client}
	items := make([]circuitTerminationResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model circuitTerminationResourceModel
		resp.Diagnostics.Append(flattenCircuitTermination(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: circuitTerminationResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
