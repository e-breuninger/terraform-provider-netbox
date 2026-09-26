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
				Description: "Query filters, sent as API list parameters. Supported names: circuit_id, circuit_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, provider_network_id, provider_network_id__n, site_id, site_id__n, tag, tag__any, tag__n, term_side, term_side__empty, term_side__ic, term_side__ie, term_side__iew, term_side__iregex, term_side__isw, term_side__n, term_side__nic, term_side__nie, term_side__niew, term_side__nisw, term_side__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
							Description: "Physical circuit speed in kbps.",
						},
						"port_speed": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use port_speed_kbps instead.",
							Description:        "Deprecated alias of port_speed_kbps.",
						},
						"upstream_speed_kbps": schema.Int64Attribute{
							Computed:    true,
							Description: "Upstream speed in kbps if different from the port speed.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
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

	fetchAll := false
	_ = fetchAll
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
	pageSize := limit
	if fetchAll {
		pageSize = 1000
	}
	params.SetLimit(&pageSize)
	for offset := int64(0); ; {
		res, err := d.client.Circuits.CircuitsCircuitTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_circuit_terminations", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CircuitTerminationResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if !fetchAll || len(res.Payload.Results) == 0 || res.Payload.Count == nil || offset >= *res.Payload.Count {
			break
		}
		params.SetOffset(&offset)
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
		resourceWithHooks.postRead(ctx, &model, &resp.Diagnostics)
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
