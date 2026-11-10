// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
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
	TagSlug            types.String `tfsdk:"tag_slug"`
	CommitRate         types.Int64  `tfsdk:"commit_rate"`
	Contact            types.Int64  `tfsdk:"contact"`
	ContactGroupSlug   types.String `tfsdk:"contact_group_slug"`
	ContactRole        types.Int64  `tfsdk:"contact_role"`
	LocationID         types.Int64  `tfsdk:"location_id"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	ProviderSlug       types.String `tfsdk:"provider_slug"`
	ProviderAccount    types.String `tfsdk:"provider_account"`
	ProviderAccountID  types.Int64  `tfsdk:"provider_account_id"`
	ProviderNetworkID  types.Int64  `tfsdk:"provider_network_id"`
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
	TerminationAID     types.Int64  `tfsdk:"termination_a_id"`
	TerminationZID     types.Int64  `tfsdk:"termination_z_id"`
	TypeSlug           types.String `tfsdk:"type_slug"`
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
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"install_date": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Installation date as YYYY-MM-DD.",
			},
			"termination_date": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Termination date as YYYY-MM-DD.",
			},
			"commit_rate_kbps": schema.Int64Attribute{
				Computed:    true,
				Description: "Committed rate in kbps. Must be a whole number.",
			},
			"distance": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Length of the circuit (requires distance_unit).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"distance_unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of distance. One of: km, m, mi, ft.",
				Validators: []validator.String{
					stringvalidator.OneOf("km", "m", "mi", "ft"),
				},
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
			"commit_rate": schema.Int64Attribute{
				Optional:    true,
				Description: "Commit rate to match.",
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
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the location.",
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
			"provider_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the provider.",
			},
			"provider_account": schema.StringAttribute{
				Optional:    true,
				Description: "Provider account (account).",
			},
			"provider_account_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the provider account.",
			},
			"provider_network_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the provider network.",
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
			"termination_a_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the termination a.",
			},
			"termination_z_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the termination a.",
			},
			"type_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the circuit type.",
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
	if !state.CommitRate.IsNull() {
		v := state.CommitRate.ValueInt64()
		params.SetCommitRate([]int64{v})
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
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.Distance.IsNull() {
		v := state.Distance.ValueFloat64()
		params.SetDistance([]float64{v})
		hasInput = true
	}
	if !state.DistanceUnit.IsNull() {
		v := state.DistanceUnit.ValueString()
		params.SetDistanceUnit(&v)
		hasInput = true
	}
	if !state.InstallDate.IsNull() {
		v := state.InstallDate.ValueString()
		params.SetInstallDate([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.LocationID.IsNull() {
		v := state.LocationID.ValueInt64()
		params.SetLocationID([]int64{v})
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
	if !state.ProviderSlug.IsNull() {
		v := state.ProviderSlug.ValueString()
		params.SetProvider([]string{v})
		hasInput = true
	}
	if !state.ProviderAccount.IsNull() {
		v := state.ProviderAccount.ValueString()
		params.SetProviderAccount([]string{v})
		hasInput = true
	}
	if !state.ProviderAccountID.IsNull() {
		v := state.ProviderAccountID.ValueInt64()
		params.SetProviderAccountID([]int64{v})
		hasInput = true
	}
	if !state.ProviderNetworkID.IsNull() {
		v := state.ProviderNetworkID.ValueInt64()
		params.SetProviderNetworkID([]int64{v})
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
	if !state.TerminationAID.IsNull() {
		v := state.TerminationAID.ValueInt64()
		params.SetTerminationaID([]int64{v})
		hasInput = true
	}
	if !state.TerminationDate.IsNull() {
		v := state.TerminationDate.ValueString()
		params.SetTerminationDate([]string{v})
		hasInput = true
	}
	if !state.TerminationZID.IsNull() {
		v := state.TerminationZID.ValueInt64()
		params.SetTerminationzID([]int64{v})
		hasInput = true
	}
	if !state.TypeSlug.IsNull() {
		v := state.TypeSlug.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_circuit.")
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
