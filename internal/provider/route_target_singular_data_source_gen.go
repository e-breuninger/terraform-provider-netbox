// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*routeTargetDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*routeTargetDataSource)(nil)
)

// NewRouteTargetDataSource returns a new route_target data source.
func NewRouteTargetDataSource() datasource.DataSource {
	return &routeTargetDataSource{}
}

type routeTargetDataSource struct {
	client *netboxapi.Client
}

// routeTargetDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type routeTargetDataSourceModel struct {
	routeTargetResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	ExportingL2vpn     types.Int64  `tfsdk:"exporting_l2vpn"`
	ExportingL2vpnID   types.Int64  `tfsdk:"exporting_l2vpn_id"`
	ExportingVrf       types.String `tfsdk:"exporting_vrf"`
	ExportingVrfID     types.Int64  `tfsdk:"exporting_vrf_id"`
	ImportingL2vpn     types.Int64  `tfsdk:"importing_l2vpn"`
	ImportingL2vpnID   types.Int64  `tfsdk:"importing_l2vpn_id"`
	ImportingVrf       types.String `tfsdk:"importing_vrf"`
	ImportingVrfID     types.Int64  `tfsdk:"importing_vrf_id"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String `tfsdk:"tenant_group_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *routeTargetDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_route_target"
}

func (d *routeTargetDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox route target (ipam.routetarget).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/ipam/routetarget/):\n\n> A route target is a particular type of extended BGP community used to control the redistribution of routes among VRF tables in a network. Route targets can be assigned to individual VRFs in NetBox as import or export targets (or both) to model this exchange in an L3VPN. Each route target must be given a unique name, which should be in a format prescribed by RFC 4364, similar to a VR route distinguisher.",
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
					stringvalidator.LengthBetween(1, 21),
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
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
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
			"exporting_l2vpn": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the exporting L2VPN (identifier).",
			},
			"exporting_l2vpn_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the exporting L2VPN.",
			},
			"exporting_vrf": schema.StringAttribute{
				Optional:    true,
				Description: "Export VRF (RD).",
			},
			"exporting_vrf_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the exporting VRF.",
			},
			"importing_l2vpn": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the importing L2VPN (identifier).",
			},
			"importing_l2vpn_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the importing L2VPN.",
			},
			"importing_vrf": schema.StringAttribute{
				Optional:    true,
				Description: "Import VRF (RD).",
			},
			"importing_vrf_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the importing VRF.",
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

func (d *routeTargetDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *routeTargetDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data routeTargetDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamRouteTargetsListParams()
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
	if !state.ExportingL2vpn.IsNull() {
		v := state.ExportingL2vpn.ValueInt64()
		params.SetExportingL2vpn([]int64{v})
		hasInput = true
	}
	if !state.ExportingL2vpnID.IsNull() {
		v := state.ExportingL2vpnID.ValueInt64()
		params.SetExportingL2vpnID([]int64{v})
		hasInput = true
	}
	if !state.ExportingVrf.IsNull() {
		v := state.ExportingVrf.ValueString()
		params.SetExportingVrf([]string{v})
		hasInput = true
	}
	if !state.ExportingVrfID.IsNull() {
		v := state.ExportingVrfID.ValueInt64()
		params.SetExportingVrfID([]int64{v})
		hasInput = true
	}
	if !state.ImportingL2vpn.IsNull() {
		v := state.ImportingL2vpn.ValueInt64()
		params.SetImportingL2vpn([]int64{v})
		hasInput = true
	}
	if !state.ImportingL2vpnID.IsNull() {
		v := state.ImportingL2vpnID.ValueInt64()
		params.SetImportingL2vpnID([]int64{v})
		hasInput = true
	}
	if !state.ImportingVrf.IsNull() {
		v := state.ImportingVrf.ValueString()
		params.SetImportingVrf([]string{v})
		hasInput = true
	}
	if !state.ImportingVrfID.IsNull() {
		v := state.ImportingVrfID.ValueInt64()
		params.SetImportingVrfID([]int64{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_route_target.")
		return
	}
	res, err := d.client.Ipam.IpamRouteTargetsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_route_target", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_route_target",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenRouteTarget(ctx, netboxapi.RouteTargetResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.routeTargetResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
