// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vCircuitID []int64
	var vCircuitIDn []int64
	var vTermSide *string
	var vTermSiden *string
	var vTermSideIc []string
	var vTermSideNic []string
	var vTermSideIe []string
	var vTermSideNie []string
	var vTermSideIsw []string
	var vTermSideNisw []string
	var vTermSideIew []string
	var vTermSideNiew []string
	var vTermSideEmpty *bool
	var vTermSideRegex []string
	var vTermSideIregex []string
	var vSiteID []int64
	var vSiteIDn []int64
	var vProviderNetworkID []int64
	var vProviderNetworkIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vCableConnector []int64
	var vCableConnectorn []int64
	var vCableConnectorLt []int64
	var vCableConnectorLte []int64
	var vCableConnectorGt []int64
	var vCableConnectorGte []int64
	var vCableConnectorEmpty *bool
	var vCableEnd *string
	var vCableEndn *string
	var vCableEndIc []string
	var vCableEndNic []string
	var vCableEndIe []string
	var vCableEndNie []string
	var vCableEndIsw []string
	var vCableEndNisw []string
	var vCableEndIew []string
	var vCableEndNiew []string
	var vCableEndEmpty *bool
	var vCableEndRegex []string
	var vCableEndIregex []string
	var vCableID []int64
	var vCableIDn []int64
	var vCabled *bool
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vDescription []string
	var vDescriptionn []string
	var vDescriptionIc []string
	var vDescriptionNic []string
	var vDescriptionIe []string
	var vDescriptionNie []string
	var vDescriptionIsw []string
	var vDescriptionNisw []string
	var vDescriptionIew []string
	var vDescriptionNiew []string
	var vDescriptionEmpty *bool
	var vDescriptionRegex []string
	var vDescriptionIregex []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLocation []string
	var vLocationn []string
	var vLocationID []string
	var vLocationIDn []string
	var vMarkConnected *bool
	var vOccupied *bool
	var vPortSpeed []int64
	var vPortSpeedn []int64
	var vPortSpeedLt []int64
	var vPortSpeedLte []int64
	var vPortSpeedGt []int64
	var vPortSpeedGte []int64
	var vPortSpeedEmpty *bool
	var vPpInfo []string
	var vPpInfon []string
	var vPpInfoIc []string
	var vPpInfoNic []string
	var vPpInfoIe []string
	var vPpInfoNie []string
	var vPpInfoIsw []string
	var vPpInfoNisw []string
	var vPpInfoIew []string
	var vPpInfoNiew []string
	var vPpInfoEmpty *bool
	var vPpInfoRegex []string
	var vPpInfoIregex []string
	var vProvider []string
	var vProvidern []string
	var vProviderID []int64
	var vProviderIDn []int64
	var vQ *string
	var vRegion []string
	var vRegionn []string
	var vRegionID []string
	var vRegionIDn []string
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []string
	var vSiteGroupIDn []string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTerminationID []int64
	var vTerminationIDn []int64
	var vTerminationIDLt []int64
	var vTerminationIDLte []int64
	var vTerminationIDGt []int64
	var vTerminationIDGte []int64
	var vTerminationIDEmpty *bool
	var vTerminationType []string
	var vTerminationTypen []string
	var vUpstreamSpeed []int64
	var vUpstreamSpeedn []int64
	var vUpstreamSpeedLt []int64
	var vUpstreamSpeedLte []int64
	var vUpstreamSpeedGt []int64
	var vUpstreamSpeedGte []int64
	var vUpstreamSpeedEmpty *bool
	var vXconnectID []string
	var vXconnectIDn []string
	var vXconnectIDIc []string
	var vXconnectIDNic []string
	var vXconnectIDIe []string
	var vXconnectIDNie []string
	var vXconnectIDIsw []string
	var vXconnectIDNisw []string
	var vXconnectIDIew []string
	var vXconnectIDNiew []string
	var vXconnectIDEmpty *bool
	var vXconnectIDRegex []string
	var vXconnectIDIregex []string
	customFieldQuery := url.Values{}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		// Custom fields are queried as cf_<field name> with the field's own filter logic; NetBox ignores
		// names it has no custom field for.
		if strings.HasPrefix(name, "cf_") {
			customFieldQuery.Add(name, value)
			continue
		}
		switch name {
		case "id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id' takes an integer, got %q.", value))
				return
			}
			vID = append(vID, v)
		case "id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__n' takes an integer, got %q.", value))
				return
			}
			vIDn = append(vIDn, v)
		case "id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lt' takes an integer, got %q.", value))
				return
			}
			vIDLt = append(vIDLt, v)
		case "id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lte' takes an integer, got %q.", value))
				return
			}
			vIDLte = append(vIDLte, v)
		case "id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gt' takes an integer, got %q.", value))
				return
			}
			vIDGt = append(vIDGt, v)
		case "id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gte' takes an integer, got %q.", value))
				return
			}
			vIDGte = append(vIDGte, v)
		case "id__empty":
			if vIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__empty' takes a boolean, got %q.", value))
				return
			}
			vIDEmpty = &v
		case "circuit_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'circuit_id' takes an integer, got %q.", value))
				return
			}
			vCircuitID = append(vCircuitID, v)
		case "circuit_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'circuit_id__n' takes an integer, got %q.", value))
				return
			}
			vCircuitIDn = append(vCircuitIDn, v)
		case "term_side":
			if vTermSide != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'term_side' takes a single value.")
				return
			}
			v := value
			vTermSide = &v
		case "term_side__n":
			if vTermSiden != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'term_side__n' takes a single value.")
				return
			}
			v := value
			vTermSiden = &v
		case "term_side__ic":
			v := value
			vTermSideIc = append(vTermSideIc, v)
		case "term_side__nic":
			v := value
			vTermSideNic = append(vTermSideNic, v)
		case "term_side__ie":
			v := value
			vTermSideIe = append(vTermSideIe, v)
		case "term_side__nie":
			v := value
			vTermSideNie = append(vTermSideNie, v)
		case "term_side__isw":
			v := value
			vTermSideIsw = append(vTermSideIsw, v)
		case "term_side__nisw":
			v := value
			vTermSideNisw = append(vTermSideNisw, v)
		case "term_side__iew":
			v := value
			vTermSideIew = append(vTermSideIew, v)
		case "term_side__niew":
			v := value
			vTermSideNiew = append(vTermSideNiew, v)
		case "term_side__empty":
			if vTermSideEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'term_side__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'term_side__empty' takes a boolean, got %q.", value))
				return
			}
			vTermSideEmpty = &v
		case "term_side__regex":
			v := value
			vTermSideRegex = append(vTermSideRegex, v)
		case "term_side__iregex":
			v := value
			vTermSideIregex = append(vTermSideIregex, v)
		case "site_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id' takes an integer, got %q.", value))
				return
			}
			vSiteID = append(vSiteID, v)
		case "site_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id__n' takes an integer, got %q.", value))
				return
			}
			vSiteIDn = append(vSiteIDn, v)
		case "provider_network_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_network_id' takes an integer, got %q.", value))
				return
			}
			vProviderNetworkID = append(vProviderNetworkID, v)
		case "provider_network_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_network_id__n' takes an integer, got %q.", value))
				return
			}
			vProviderNetworkIDn = append(vProviderNetworkIDn, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
		case "cable_connector":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector' takes an integer, got %q.", value))
				return
			}
			vCableConnector = append(vCableConnector, v)
		case "cable_connector__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__n' takes an integer, got %q.", value))
				return
			}
			vCableConnectorn = append(vCableConnectorn, v)
		case "cable_connector__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__lt' takes an integer, got %q.", value))
				return
			}
			vCableConnectorLt = append(vCableConnectorLt, v)
		case "cable_connector__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__lte' takes an integer, got %q.", value))
				return
			}
			vCableConnectorLte = append(vCableConnectorLte, v)
		case "cable_connector__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__gt' takes an integer, got %q.", value))
				return
			}
			vCableConnectorGt = append(vCableConnectorGt, v)
		case "cable_connector__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__gte' takes an integer, got %q.", value))
				return
			}
			vCableConnectorGte = append(vCableConnectorGte, v)
		case "cable_connector__empty":
			if vCableConnectorEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_connector__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__empty' takes a boolean, got %q.", value))
				return
			}
			vCableConnectorEmpty = &v
		case "cable_end":
			if vCableEnd != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end' takes a single value.")
				return
			}
			v := value
			vCableEnd = &v
		case "cable_end__n":
			if vCableEndn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end__n' takes a single value.")
				return
			}
			v := value
			vCableEndn = &v
		case "cable_end__ic":
			v := value
			vCableEndIc = append(vCableEndIc, v)
		case "cable_end__nic":
			v := value
			vCableEndNic = append(vCableEndNic, v)
		case "cable_end__ie":
			v := value
			vCableEndIe = append(vCableEndIe, v)
		case "cable_end__nie":
			v := value
			vCableEndNie = append(vCableEndNie, v)
		case "cable_end__isw":
			v := value
			vCableEndIsw = append(vCableEndIsw, v)
		case "cable_end__nisw":
			v := value
			vCableEndNisw = append(vCableEndNisw, v)
		case "cable_end__iew":
			v := value
			vCableEndIew = append(vCableEndIew, v)
		case "cable_end__niew":
			v := value
			vCableEndNiew = append(vCableEndNiew, v)
		case "cable_end__empty":
			if vCableEndEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_end__empty' takes a boolean, got %q.", value))
				return
			}
			vCableEndEmpty = &v
		case "cable_end__regex":
			v := value
			vCableEndRegex = append(vCableEndRegex, v)
		case "cable_end__iregex":
			v := value
			vCableEndIregex = append(vCableEndIregex, v)
		case "cable_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_id' takes an integer, got %q.", value))
				return
			}
			vCableID = append(vCableID, v)
		case "cable_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_id__n' takes an integer, got %q.", value))
				return
			}
			vCableIDn = append(vCableIDn, v)
		case "cabled":
			if vCabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cabled' takes a boolean, got %q.", value))
				return
			}
			vCabled = &v
		case "created":
			v := value
			vCreated = append(vCreated, v)
		case "created__n":
			v := value
			vCreatedn = append(vCreatedn, v)
		case "created__lt":
			v := value
			vCreatedLt = append(vCreatedLt, v)
		case "created__lte":
			v := value
			vCreatedLte = append(vCreatedLte, v)
		case "created__gt":
			v := value
			vCreatedGt = append(vCreatedGt, v)
		case "created__gte":
			v := value
			vCreatedGte = append(vCreatedGte, v)
		case "created__empty":
			v := value
			vCreatedEmpty = append(vCreatedEmpty, v)
		case "description":
			v := value
			vDescription = append(vDescription, v)
		case "description__n":
			v := value
			vDescriptionn = append(vDescriptionn, v)
		case "description__ic":
			v := value
			vDescriptionIc = append(vDescriptionIc, v)
		case "description__nic":
			v := value
			vDescriptionNic = append(vDescriptionNic, v)
		case "description__ie":
			v := value
			vDescriptionIe = append(vDescriptionIe, v)
		case "description__nie":
			v := value
			vDescriptionNie = append(vDescriptionNie, v)
		case "description__isw":
			v := value
			vDescriptionIsw = append(vDescriptionIsw, v)
		case "description__nisw":
			v := value
			vDescriptionNisw = append(vDescriptionNisw, v)
		case "description__iew":
			v := value
			vDescriptionIew = append(vDescriptionIew, v)
		case "description__niew":
			v := value
			vDescriptionNiew = append(vDescriptionNiew, v)
		case "description__empty":
			if vDescriptionEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'description__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'description__empty' takes a boolean, got %q.", value))
				return
			}
			vDescriptionEmpty = &v
		case "description__regex":
			v := value
			vDescriptionRegex = append(vDescriptionRegex, v)
		case "description__iregex":
			v := value
			vDescriptionIregex = append(vDescriptionIregex, v)
		case "last_updated":
			v := value
			vLastUpdated = append(vLastUpdated, v)
		case "last_updated__n":
			v := value
			vLastUpdatedn = append(vLastUpdatedn, v)
		case "last_updated__lt":
			v := value
			vLastUpdatedLt = append(vLastUpdatedLt, v)
		case "last_updated__lte":
			v := value
			vLastUpdatedLte = append(vLastUpdatedLte, v)
		case "last_updated__gt":
			v := value
			vLastUpdatedGt = append(vLastUpdatedGt, v)
		case "last_updated__gte":
			v := value
			vLastUpdatedGte = append(vLastUpdatedGte, v)
		case "last_updated__empty":
			v := value
			vLastUpdatedEmpty = append(vLastUpdatedEmpty, v)
		case "location":
			v := value
			vLocation = append(vLocation, v)
		case "location__n":
			v := value
			vLocationn = append(vLocationn, v)
		case "location_id":
			v := value
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v := value
			vLocationIDn = append(vLocationIDn, v)
		case "mark_connected":
			if vMarkConnected != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mark_connected' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mark_connected' takes a boolean, got %q.", value))
				return
			}
			vMarkConnected = &v
		case "occupied":
			if vOccupied != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'occupied' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'occupied' takes a boolean, got %q.", value))
				return
			}
			vOccupied = &v
		case "port_speed":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'port_speed' takes an integer, got %q.", value))
				return
			}
			vPortSpeed = append(vPortSpeed, v)
		case "port_speed__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'port_speed__n' takes an integer, got %q.", value))
				return
			}
			vPortSpeedn = append(vPortSpeedn, v)
		case "port_speed__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'port_speed__lt' takes an integer, got %q.", value))
				return
			}
			vPortSpeedLt = append(vPortSpeedLt, v)
		case "port_speed__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'port_speed__lte' takes an integer, got %q.", value))
				return
			}
			vPortSpeedLte = append(vPortSpeedLte, v)
		case "port_speed__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'port_speed__gt' takes an integer, got %q.", value))
				return
			}
			vPortSpeedGt = append(vPortSpeedGt, v)
		case "port_speed__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'port_speed__gte' takes an integer, got %q.", value))
				return
			}
			vPortSpeedGte = append(vPortSpeedGte, v)
		case "port_speed__empty":
			if vPortSpeedEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'port_speed__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'port_speed__empty' takes a boolean, got %q.", value))
				return
			}
			vPortSpeedEmpty = &v
		case "pp_info":
			v := value
			vPpInfo = append(vPpInfo, v)
		case "pp_info__n":
			v := value
			vPpInfon = append(vPpInfon, v)
		case "pp_info__ic":
			v := value
			vPpInfoIc = append(vPpInfoIc, v)
		case "pp_info__nic":
			v := value
			vPpInfoNic = append(vPpInfoNic, v)
		case "pp_info__ie":
			v := value
			vPpInfoIe = append(vPpInfoIe, v)
		case "pp_info__nie":
			v := value
			vPpInfoNie = append(vPpInfoNie, v)
		case "pp_info__isw":
			v := value
			vPpInfoIsw = append(vPpInfoIsw, v)
		case "pp_info__nisw":
			v := value
			vPpInfoNisw = append(vPpInfoNisw, v)
		case "pp_info__iew":
			v := value
			vPpInfoIew = append(vPpInfoIew, v)
		case "pp_info__niew":
			v := value
			vPpInfoNiew = append(vPpInfoNiew, v)
		case "pp_info__empty":
			if vPpInfoEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'pp_info__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pp_info__empty' takes a boolean, got %q.", value))
				return
			}
			vPpInfoEmpty = &v
		case "pp_info__regex":
			v := value
			vPpInfoRegex = append(vPpInfoRegex, v)
		case "pp_info__iregex":
			v := value
			vPpInfoIregex = append(vPpInfoIregex, v)
		case "provider":
			v := value
			vProvider = append(vProvider, v)
		case "provider__n":
			v := value
			vProvidern = append(vProvidern, v)
		case "provider_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_id' takes an integer, got %q.", value))
				return
			}
			vProviderID = append(vProviderID, v)
		case "provider_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_id__n' takes an integer, got %q.", value))
				return
			}
			vProviderIDn = append(vProviderIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "region":
			v := value
			vRegion = append(vRegion, v)
		case "region__n":
			v := value
			vRegionn = append(vRegionn, v)
		case "region_id":
			v := value
			vRegionID = append(vRegionID, v)
		case "region_id__n":
			v := value
			vRegionIDn = append(vRegionIDn, v)
		case "site":
			v := value
			vSite = append(vSite, v)
		case "site__n":
			v := value
			vSiten = append(vSiten, v)
		case "site_group":
			v := value
			vSiteGroup = append(vSiteGroup, v)
		case "site_group__n":
			v := value
			vSiteGroupn = append(vSiteGroupn, v)
		case "site_group_id":
			v := value
			vSiteGroupID = append(vSiteGroupID, v)
		case "site_group_id__n":
			v := value
			vSiteGroupIDn = append(vSiteGroupIDn, v)
		case "tag_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id' takes an integer, got %q.", value))
				return
			}
			vTagID = append(vTagID, v)
		case "tag_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__n' takes an integer, got %q.", value))
				return
			}
			vTagIDn = append(vTagIDn, v)
		case "tag_id__any":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__any' takes an integer, got %q.", value))
				return
			}
			vTagIDAny = append(vTagIDAny, v)
		case "termination_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id' takes an integer, got %q.", value))
				return
			}
			vTerminationID = append(vTerminationID, v)
		case "termination_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__n' takes an integer, got %q.", value))
				return
			}
			vTerminationIDn = append(vTerminationIDn, v)
		case "termination_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__lt' takes an integer, got %q.", value))
				return
			}
			vTerminationIDLt = append(vTerminationIDLt, v)
		case "termination_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__lte' takes an integer, got %q.", value))
				return
			}
			vTerminationIDLte = append(vTerminationIDLte, v)
		case "termination_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__gt' takes an integer, got %q.", value))
				return
			}
			vTerminationIDGt = append(vTerminationIDGt, v)
		case "termination_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__gte' takes an integer, got %q.", value))
				return
			}
			vTerminationIDGte = append(vTerminationIDGte, v)
		case "termination_id__empty":
			if vTerminationIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'termination_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_id__empty' takes a boolean, got %q.", value))
				return
			}
			vTerminationIDEmpty = &v
		case "termination_type":
			v := value
			vTerminationType = append(vTerminationType, v)
		case "termination_type__n":
			v := value
			vTerminationTypen = append(vTerminationTypen, v)
		case "upstream_speed":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'upstream_speed' takes an integer, got %q.", value))
				return
			}
			vUpstreamSpeed = append(vUpstreamSpeed, v)
		case "upstream_speed__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'upstream_speed__n' takes an integer, got %q.", value))
				return
			}
			vUpstreamSpeedn = append(vUpstreamSpeedn, v)
		case "upstream_speed__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'upstream_speed__lt' takes an integer, got %q.", value))
				return
			}
			vUpstreamSpeedLt = append(vUpstreamSpeedLt, v)
		case "upstream_speed__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'upstream_speed__lte' takes an integer, got %q.", value))
				return
			}
			vUpstreamSpeedLte = append(vUpstreamSpeedLte, v)
		case "upstream_speed__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'upstream_speed__gt' takes an integer, got %q.", value))
				return
			}
			vUpstreamSpeedGt = append(vUpstreamSpeedGt, v)
		case "upstream_speed__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'upstream_speed__gte' takes an integer, got %q.", value))
				return
			}
			vUpstreamSpeedGte = append(vUpstreamSpeedGte, v)
		case "upstream_speed__empty":
			if vUpstreamSpeedEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'upstream_speed__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'upstream_speed__empty' takes a boolean, got %q.", value))
				return
			}
			vUpstreamSpeedEmpty = &v
		case "xconnect_id":
			v := value
			vXconnectID = append(vXconnectID, v)
		case "xconnect_id__n":
			v := value
			vXconnectIDn = append(vXconnectIDn, v)
		case "xconnect_id__ic":
			v := value
			vXconnectIDIc = append(vXconnectIDIc, v)
		case "xconnect_id__nic":
			v := value
			vXconnectIDNic = append(vXconnectIDNic, v)
		case "xconnect_id__ie":
			v := value
			vXconnectIDIe = append(vXconnectIDIe, v)
		case "xconnect_id__nie":
			v := value
			vXconnectIDNie = append(vXconnectIDNie, v)
		case "xconnect_id__isw":
			v := value
			vXconnectIDIsw = append(vXconnectIDIsw, v)
		case "xconnect_id__nisw":
			v := value
			vXconnectIDNisw = append(vXconnectIDNisw, v)
		case "xconnect_id__iew":
			v := value
			vXconnectIDIew = append(vXconnectIDIew, v)
		case "xconnect_id__niew":
			v := value
			vXconnectIDNiew = append(vXconnectIDNiew, v)
		case "xconnect_id__empty":
			if vXconnectIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'xconnect_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'xconnect_id__empty' takes a boolean, got %q.", value))
				return
			}
			vXconnectIDEmpty = &v
		case "xconnect_id__regex":
			v := value
			vXconnectIDRegex = append(vXconnectIDRegex, v)
		case "xconnect_id__iregex":
			v := value
			vXconnectIDIregex = append(vXconnectIDIregex, v)
		default:
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
	}
	if len(vID) > 0 {
		params.SetID(vID)
	}
	if len(vIDn) > 0 {
		params.SetIDn(vIDn)
	}
	if len(vIDLt) > 0 {
		params.SetIDLt(vIDLt)
	}
	if len(vIDLte) > 0 {
		params.SetIDLte(vIDLte)
	}
	if len(vIDGt) > 0 {
		params.SetIDGt(vIDGt)
	}
	if len(vIDGte) > 0 {
		params.SetIDGte(vIDGte)
	}
	if vIDEmpty != nil {
		params.SetIDEmpty(vIDEmpty)
	}
	if len(vCircuitID) > 0 {
		params.SetCircuitID(vCircuitID)
	}
	if len(vCircuitIDn) > 0 {
		params.SetCircuitIDn(vCircuitIDn)
	}
	if vTermSide != nil {
		params.SetTermSide(vTermSide)
	}
	if vTermSiden != nil {
		params.SetTermSiden(vTermSiden)
	}
	if len(vTermSideIc) > 0 {
		params.SetTermSideIc(vTermSideIc)
	}
	if len(vTermSideNic) > 0 {
		params.SetTermSideNic(vTermSideNic)
	}
	if len(vTermSideIe) > 0 {
		params.SetTermSideIe(vTermSideIe)
	}
	if len(vTermSideNie) > 0 {
		params.SetTermSideNie(vTermSideNie)
	}
	if len(vTermSideIsw) > 0 {
		params.SetTermSideIsw(vTermSideIsw)
	}
	if len(vTermSideNisw) > 0 {
		params.SetTermSideNisw(vTermSideNisw)
	}
	if len(vTermSideIew) > 0 {
		params.SetTermSideIew(vTermSideIew)
	}
	if len(vTermSideNiew) > 0 {
		params.SetTermSideNiew(vTermSideNiew)
	}
	if vTermSideEmpty != nil {
		params.SetTermSideEmpty(vTermSideEmpty)
	}
	if len(vTermSideRegex) > 0 {
		params.SetTermSideRegex(vTermSideRegex)
	}
	if len(vTermSideIregex) > 0 {
		params.SetTermSideIregex(vTermSideIregex)
	}
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
	}
	if len(vProviderNetworkID) > 0 {
		params.SetProviderNetworkID(vProviderNetworkID)
	}
	if len(vProviderNetworkIDn) > 0 {
		params.SetProviderNetworkIDn(vProviderNetworkIDn)
	}
	if len(vTag) > 0 {
		params.SetTag(vTag)
	}
	if len(vTagn) > 0 {
		params.SetTagn(vTagn)
	}
	if len(vTagAny) > 0 {
		params.SetTagAny(vTagAny)
	}
	if len(vCableConnector) > 0 {
		params.SetCableConnector(vCableConnector)
	}
	if len(vCableConnectorn) > 0 {
		params.SetCableConnectorn(vCableConnectorn)
	}
	if len(vCableConnectorLt) > 0 {
		params.SetCableConnectorLt(vCableConnectorLt)
	}
	if len(vCableConnectorLte) > 0 {
		params.SetCableConnectorLte(vCableConnectorLte)
	}
	if len(vCableConnectorGt) > 0 {
		params.SetCableConnectorGt(vCableConnectorGt)
	}
	if len(vCableConnectorGte) > 0 {
		params.SetCableConnectorGte(vCableConnectorGte)
	}
	if vCableConnectorEmpty != nil {
		params.SetCableConnectorEmpty(vCableConnectorEmpty)
	}
	if vCableEnd != nil {
		params.SetCableEnd(vCableEnd)
	}
	if vCableEndn != nil {
		params.SetCableEndn(vCableEndn)
	}
	if len(vCableEndIc) > 0 {
		params.SetCableEndIc(vCableEndIc)
	}
	if len(vCableEndNic) > 0 {
		params.SetCableEndNic(vCableEndNic)
	}
	if len(vCableEndIe) > 0 {
		params.SetCableEndIe(vCableEndIe)
	}
	if len(vCableEndNie) > 0 {
		params.SetCableEndNie(vCableEndNie)
	}
	if len(vCableEndIsw) > 0 {
		params.SetCableEndIsw(vCableEndIsw)
	}
	if len(vCableEndNisw) > 0 {
		params.SetCableEndNisw(vCableEndNisw)
	}
	if len(vCableEndIew) > 0 {
		params.SetCableEndIew(vCableEndIew)
	}
	if len(vCableEndNiew) > 0 {
		params.SetCableEndNiew(vCableEndNiew)
	}
	if vCableEndEmpty != nil {
		params.SetCableEndEmpty(vCableEndEmpty)
	}
	if len(vCableEndRegex) > 0 {
		params.SetCableEndRegex(vCableEndRegex)
	}
	if len(vCableEndIregex) > 0 {
		params.SetCableEndIregex(vCableEndIregex)
	}
	if len(vCableID) > 0 {
		params.SetCableID(vCableID)
	}
	if len(vCableIDn) > 0 {
		params.SetCableIDn(vCableIDn)
	}
	if vCabled != nil {
		params.SetCabled(vCabled)
	}
	if len(vCreated) > 0 {
		params.SetCreated(vCreated)
	}
	if len(vCreatedn) > 0 {
		params.SetCreatedn(vCreatedn)
	}
	if len(vCreatedLt) > 0 {
		params.SetCreatedLt(vCreatedLt)
	}
	if len(vCreatedLte) > 0 {
		params.SetCreatedLte(vCreatedLte)
	}
	if len(vCreatedGt) > 0 {
		params.SetCreatedGt(vCreatedGt)
	}
	if len(vCreatedGte) > 0 {
		params.SetCreatedGte(vCreatedGte)
	}
	if len(vCreatedEmpty) > 0 {
		params.SetCreatedEmpty(vCreatedEmpty)
	}
	if len(vDescription) > 0 {
		params.SetDescription(vDescription)
	}
	if len(vDescriptionn) > 0 {
		params.SetDescriptionn(vDescriptionn)
	}
	if len(vDescriptionIc) > 0 {
		params.SetDescriptionIc(vDescriptionIc)
	}
	if len(vDescriptionNic) > 0 {
		params.SetDescriptionNic(vDescriptionNic)
	}
	if len(vDescriptionIe) > 0 {
		params.SetDescriptionIe(vDescriptionIe)
	}
	if len(vDescriptionNie) > 0 {
		params.SetDescriptionNie(vDescriptionNie)
	}
	if len(vDescriptionIsw) > 0 {
		params.SetDescriptionIsw(vDescriptionIsw)
	}
	if len(vDescriptionNisw) > 0 {
		params.SetDescriptionNisw(vDescriptionNisw)
	}
	if len(vDescriptionIew) > 0 {
		params.SetDescriptionIew(vDescriptionIew)
	}
	if len(vDescriptionNiew) > 0 {
		params.SetDescriptionNiew(vDescriptionNiew)
	}
	if vDescriptionEmpty != nil {
		params.SetDescriptionEmpty(vDescriptionEmpty)
	}
	if len(vDescriptionRegex) > 0 {
		params.SetDescriptionRegex(vDescriptionRegex)
	}
	if len(vDescriptionIregex) > 0 {
		params.SetDescriptionIregex(vDescriptionIregex)
	}
	if len(vLastUpdated) > 0 {
		params.SetLastUpdated(vLastUpdated)
	}
	if len(vLastUpdatedn) > 0 {
		params.SetLastUpdatedn(vLastUpdatedn)
	}
	if len(vLastUpdatedLt) > 0 {
		params.SetLastUpdatedLt(vLastUpdatedLt)
	}
	if len(vLastUpdatedLte) > 0 {
		params.SetLastUpdatedLte(vLastUpdatedLte)
	}
	if len(vLastUpdatedGt) > 0 {
		params.SetLastUpdatedGt(vLastUpdatedGt)
	}
	if len(vLastUpdatedGte) > 0 {
		params.SetLastUpdatedGte(vLastUpdatedGte)
	}
	if len(vLastUpdatedEmpty) > 0 {
		params.SetLastUpdatedEmpty(vLastUpdatedEmpty)
	}
	if len(vLocation) > 0 {
		params.SetLocation(vLocation)
	}
	if len(vLocationn) > 0 {
		params.SetLocationn(vLocationn)
	}
	if len(vLocationID) > 0 {
		params.SetLocationID(vLocationID)
	}
	if len(vLocationIDn) > 0 {
		params.SetLocationIDn(vLocationIDn)
	}
	if vMarkConnected != nil {
		params.SetMarkConnected(vMarkConnected)
	}
	if vOccupied != nil {
		params.SetOccupied(vOccupied)
	}
	if len(vPortSpeed) > 0 {
		params.SetPortSpeed(vPortSpeed)
	}
	if len(vPortSpeedn) > 0 {
		params.SetPortSpeedn(vPortSpeedn)
	}
	if len(vPortSpeedLt) > 0 {
		params.SetPortSpeedLt(vPortSpeedLt)
	}
	if len(vPortSpeedLte) > 0 {
		params.SetPortSpeedLte(vPortSpeedLte)
	}
	if len(vPortSpeedGt) > 0 {
		params.SetPortSpeedGt(vPortSpeedGt)
	}
	if len(vPortSpeedGte) > 0 {
		params.SetPortSpeedGte(vPortSpeedGte)
	}
	if vPortSpeedEmpty != nil {
		params.SetPortSpeedEmpty(vPortSpeedEmpty)
	}
	if len(vPpInfo) > 0 {
		params.SetPpInfo(vPpInfo)
	}
	if len(vPpInfon) > 0 {
		params.SetPpInfon(vPpInfon)
	}
	if len(vPpInfoIc) > 0 {
		params.SetPpInfoIc(vPpInfoIc)
	}
	if len(vPpInfoNic) > 0 {
		params.SetPpInfoNic(vPpInfoNic)
	}
	if len(vPpInfoIe) > 0 {
		params.SetPpInfoIe(vPpInfoIe)
	}
	if len(vPpInfoNie) > 0 {
		params.SetPpInfoNie(vPpInfoNie)
	}
	if len(vPpInfoIsw) > 0 {
		params.SetPpInfoIsw(vPpInfoIsw)
	}
	if len(vPpInfoNisw) > 0 {
		params.SetPpInfoNisw(vPpInfoNisw)
	}
	if len(vPpInfoIew) > 0 {
		params.SetPpInfoIew(vPpInfoIew)
	}
	if len(vPpInfoNiew) > 0 {
		params.SetPpInfoNiew(vPpInfoNiew)
	}
	if vPpInfoEmpty != nil {
		params.SetPpInfoEmpty(vPpInfoEmpty)
	}
	if len(vPpInfoRegex) > 0 {
		params.SetPpInfoRegex(vPpInfoRegex)
	}
	if len(vPpInfoIregex) > 0 {
		params.SetPpInfoIregex(vPpInfoIregex)
	}
	if len(vProvider) > 0 {
		params.SetProvider(vProvider)
	}
	if len(vProvidern) > 0 {
		params.SetProvidern(vProvidern)
	}
	if len(vProviderID) > 0 {
		params.SetProviderID(vProviderID)
	}
	if len(vProviderIDn) > 0 {
		params.SetProviderIDn(vProviderIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRegion) > 0 {
		params.SetRegion(vRegion)
	}
	if len(vRegionn) > 0 {
		params.SetRegionn(vRegionn)
	}
	if len(vRegionID) > 0 {
		params.SetRegionID(vRegionID)
	}
	if len(vRegionIDn) > 0 {
		params.SetRegionIDn(vRegionIDn)
	}
	if len(vSite) > 0 {
		params.SetSite(vSite)
	}
	if len(vSiten) > 0 {
		params.SetSiten(vSiten)
	}
	if len(vSiteGroup) > 0 {
		params.SetSiteGroup(vSiteGroup)
	}
	if len(vSiteGroupn) > 0 {
		params.SetSiteGroupn(vSiteGroupn)
	}
	if len(vSiteGroupID) > 0 {
		params.SetSiteGroupID(vSiteGroupID)
	}
	if len(vSiteGroupIDn) > 0 {
		params.SetSiteGroupIDn(vSiteGroupIDn)
	}
	if len(vTagID) > 0 {
		params.SetTagID(vTagID)
	}
	if len(vTagIDn) > 0 {
		params.SetTagIDn(vTagIDn)
	}
	if len(vTagIDAny) > 0 {
		params.SetTagIDAny(vTagIDAny)
	}
	if len(vTerminationID) > 0 {
		params.SetTerminationID(vTerminationID)
	}
	if len(vTerminationIDn) > 0 {
		params.SetTerminationIDn(vTerminationIDn)
	}
	if len(vTerminationIDLt) > 0 {
		params.SetTerminationIDLt(vTerminationIDLt)
	}
	if len(vTerminationIDLte) > 0 {
		params.SetTerminationIDLte(vTerminationIDLte)
	}
	if len(vTerminationIDGt) > 0 {
		params.SetTerminationIDGt(vTerminationIDGt)
	}
	if len(vTerminationIDGte) > 0 {
		params.SetTerminationIDGte(vTerminationIDGte)
	}
	if vTerminationIDEmpty != nil {
		params.SetTerminationIDEmpty(vTerminationIDEmpty)
	}
	if len(vTerminationType) > 0 {
		params.SetTerminationType(vTerminationType)
	}
	if len(vTerminationTypen) > 0 {
		params.SetTerminationTypen(vTerminationTypen)
	}
	if len(vUpstreamSpeed) > 0 {
		params.SetUpstreamSpeed(vUpstreamSpeed)
	}
	if len(vUpstreamSpeedn) > 0 {
		params.SetUpstreamSpeedn(vUpstreamSpeedn)
	}
	if len(vUpstreamSpeedLt) > 0 {
		params.SetUpstreamSpeedLt(vUpstreamSpeedLt)
	}
	if len(vUpstreamSpeedLte) > 0 {
		params.SetUpstreamSpeedLte(vUpstreamSpeedLte)
	}
	if len(vUpstreamSpeedGt) > 0 {
		params.SetUpstreamSpeedGt(vUpstreamSpeedGt)
	}
	if len(vUpstreamSpeedGte) > 0 {
		params.SetUpstreamSpeedGte(vUpstreamSpeedGte)
	}
	if vUpstreamSpeedEmpty != nil {
		params.SetUpstreamSpeedEmpty(vUpstreamSpeedEmpty)
	}
	if len(vXconnectID) > 0 {
		params.SetXconnectID(vXconnectID)
	}
	if len(vXconnectIDn) > 0 {
		params.SetXconnectIDn(vXconnectIDn)
	}
	if len(vXconnectIDIc) > 0 {
		params.SetXconnectIDIc(vXconnectIDIc)
	}
	if len(vXconnectIDNic) > 0 {
		params.SetXconnectIDNic(vXconnectIDNic)
	}
	if len(vXconnectIDIe) > 0 {
		params.SetXconnectIDIe(vXconnectIDIe)
	}
	if len(vXconnectIDNie) > 0 {
		params.SetXconnectIDNie(vXconnectIDNie)
	}
	if len(vXconnectIDIsw) > 0 {
		params.SetXconnectIDIsw(vXconnectIDIsw)
	}
	if len(vXconnectIDNisw) > 0 {
		params.SetXconnectIDNisw(vXconnectIDNisw)
	}
	if len(vXconnectIDIew) > 0 {
		params.SetXconnectIDIew(vXconnectIDIew)
	}
	if len(vXconnectIDNiew) > 0 {
		params.SetXconnectIDNiew(vXconnectIDNiew)
	}
	if vXconnectIDEmpty != nil {
		params.SetXconnectIDEmpty(vXconnectIDEmpty)
	}
	if len(vXconnectIDRegex) > 0 {
		params.SetXconnectIDRegex(vXconnectIDRegex)
	}
	if len(vXconnectIDIregex) > 0 {
		params.SetXconnectIDIregex(vXconnectIDIregex)
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
