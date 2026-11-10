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
	_ datasource.DataSource              = (*virtualCircuitDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualCircuitDataSource)(nil)
)

// NewVirtualCircuitDataSource returns a new virtual_circuit data source.
func NewVirtualCircuitDataSource() datasource.DataSource {
	return &virtualCircuitDataSource{}
}

type virtualCircuitDataSource struct {
	client *netboxapi.Client
}

// virtualCircuitDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type virtualCircuitDataSourceModel struct {
	virtualCircuitResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	ProviderSlug       types.String `tfsdk:"provider_slug"`
	ProviderAccount    types.String `tfsdk:"provider_account"`
	ProviderID         types.Int64  `tfsdk:"provider_id"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String `tfsdk:"tenant_group_id"`
	TypeSlug           types.String `tfsdk:"type_slug"`
	TypeID             types.Int64  `tfsdk:"type_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *virtualCircuitDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_circuit"
}

func (d *virtualCircuitDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:A NetBox virtual circuit (circuits.virtualcircuit): a layer 2 or 3 service delivered over a provider network.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/circuits/virtualcircuit/):\n\n> A virtual circuit can connect two or more interfaces atop a set of decoupled physical connections. For example, it's very common to form a virtual connection between two virtual interfaces, each of which is bound to a physical interface on its respective device and physically connected to a [provider network](https://netboxlabs.com/docs/netbox/models/circuits/providernetwork/) via an independent [physical circuit](https://netboxlabs.com/docs/netbox/models/circuits/circuit/).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"cid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Circuit ID: unique identifier of the virtual circuit with its provider.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"circuit_provider_network_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the provider network the virtual circuit is delivered over.",
			},
			"virtual_circuit_type_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the virtual circuit type.",
			},
			"circuit_provider_account_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the provider account the virtual circuit is billed to.",
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
				Description: "Id of the tenant the virtual circuit belongs to.",
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
			"provider_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the provider.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
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
			"type_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the virtual circuit type.",
			},
			"type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual circuit type.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *virtualCircuitDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualCircuitDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualCircuitDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := circuits.NewCircuitsVirtualCircuitsListParams()
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
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.CircuitProviderNetworkID.IsNull() {
		v := state.CircuitProviderNetworkID.ValueInt64()
		params.SetProviderNetworkID([]int64{v})
		hasInput = true
	}
	if !state.CircuitProviderAccountID.IsNull() {
		v := state.CircuitProviderAccountID.ValueInt64()
		params.SetProviderAccountID([]int64{v})
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
	if !state.TypeSlug.IsNull() {
		v := state.TypeSlug.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.TypeID.IsNull() {
		v := state.TypeID.ValueInt64()
		params.SetTypeID([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_virtual_circuit.")
		return
	}
	res, err := d.client.Circuits.CircuitsVirtualCircuitsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_virtual_circuit", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_virtual_circuit",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVirtualCircuit(ctx, netboxapi.VirtualCircuitResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.virtualCircuitResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
