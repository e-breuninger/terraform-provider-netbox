// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*circuitDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*circuitDataSource)(nil)
)

// NewCircuitDataSource returns a new circuit data source.
func NewCircuitDataSource() datasource.DataSource {
	return &circuitDataSource{}
}

type circuitDataSource struct {
	client *netboxapi.Client
}

// circuitDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type circuitDataSourceModel struct {
	circuitResourceModel
	CidContains        types.String `tfsdk:"cid_contains"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *circuitDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_circuit"
}

func (d *circuitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:A circuit leased from a provider (circuits.circuit).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/circuits/#circuits_1):\n\n> A communications circuit represents a single physical link connecting exactly two endpoints, commonly referred to as its A and Z terminations. A circuit in NetBox may have zero, one, or two terminations defined. It is common to have only one termination defined when you don't necessarily care about the details of the provider side of the circuit, e.g. for Internet access circuits. Both terminations would likely be modeled for circuits which connect one customer site to another.\n>\n> Each circuit is associated with a provider and a user-defined type. For example, you might have Internet access circuits delivered to each site by one provider, and private MPLS circuits delivered by another. Each circuit must be assigned a circuit ID, each of which must be unique per provider.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"cid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Circuit id assigned by the provider; unique per provider.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"circuit_provider_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the provider the circuit is leased from.",
			},
			"circuit_provider_account_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the provider account the circuit is billed to.",
			},
			"circuit_type_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the circuit type.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: planned, provisioning, active, offline, deprovisioning, decommissioned.",
				Validators: []validator.String{
					stringvalidator.OneOf("planned", "provisioning", "active", "offline", "deprovisioning", "decommissioned"),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"install_date": schema.StringAttribute{
				Computed:    true,
				Description: "Installation date as YYYY-MM-DD.",
			},
			"termination_date": schema.StringAttribute{
				Computed:    true,
				Description: "Termination date as YYYY-MM-DD.",
			},
			"commit_rate_kbps": schema.Int64Attribute{
				Computed:    true,
				Description: "Committed rate in kbps.",
			},
			"distance": schema.Float64Attribute{
				Computed:    true,
				Description: "Length of the circuit (requires distance_unit).",
			},
			"distance_unit": schema.StringAttribute{
				Computed:    true,
				Description: "Unit of distance. One of: km, m, mi, ft.",
			},
			"description": schema.StringAttribute{
				Computed: true,
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
			"cid_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the circuit id.",
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

func (d *circuitDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *circuitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data circuitDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := circuits.NewCircuitsCircuitsListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Cid.IsNull() {
		v := state.Cid.ValueString()
		params.SetCid([]string{v})
		hasInput = true
	}
	if !state.CidContains.IsNull() {
		v := state.CidContains.ValueString()
		params.SetCidIc([]string{v})
		hasInput = true
	}
	if !state.CircuitProviderID.IsNull() {
		v := state.CircuitProviderID.ValueInt64()
		params.SetProviderID([]int64{v})
		hasInput = true
	}
	if !state.CircuitTypeID.IsNull() {
		v := state.CircuitTypeID.ValueInt64()
		params.SetTypeID([]int64{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or cid and/or cid_contains and/or circuit_provider_id and/or circuit_type_id and/or status and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_circuit.")
		return
	}
	res, err := d.client.Circuits.CircuitsCircuitsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_circuit", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_circuit",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCircuit(ctx, netboxapi.CircuitResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.circuitResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
