// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*rackReservationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*rackReservationDataSource)(nil)
)

// NewRackReservationDataSource returns a new rack_reservation data source.
func NewRackReservationDataSource() datasource.DataSource {
	return &rackReservationDataSource{}
}

type rackReservationDataSource struct {
	client *netboxapi.Client
}

// rackReservationDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type rackReservationDataSourceModel struct {
	rackReservationResourceModel
	TagSlug            types.String  `tfsdk:"tag_slug"`
	GroupSlug          types.String  `tfsdk:"group_slug"`
	GroupID            types.Int64   `tfsdk:"group_id"`
	LocationSlug       types.String  `tfsdk:"location_slug"`
	LocationID         types.String  `tfsdk:"location_id"`
	OwnerName          types.String  `tfsdk:"owner_name"`
	OwnerGroupName     types.String  `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64   `tfsdk:"owner_group_id"`
	Q                  types.String  `tfsdk:"q"`
	RegionSlug         types.String  `tfsdk:"region_slug"`
	RegionID           types.String  `tfsdk:"region_id"`
	SiteSlug           types.String  `tfsdk:"site_slug"`
	SiteGroupSlug      types.String  `tfsdk:"site_group_slug"`
	SiteGroupID        types.String  `tfsdk:"site_group_id"`
	SiteID             types.Int64   `tfsdk:"site_id"`
	TagID              types.Int64   `tfsdk:"tag_id"`
	TenantSlug         types.String  `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String  `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String  `tfsdk:"tenant_group_id"`
	Unit               types.Float64 `tfsdk:"unit"`
	UnitCountMax       types.Float64 `tfsdk:"unit_count_max"`
	UnitCountMin       types.Float64 `tfsdk:"unit_count_min"`
	UserName           types.String  `tfsdk:"user_name"`
	CustomFieldFilters types.Map     `tfsdk:"custom_field_filters"`
}

func (d *rackReservationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rack_reservation"
}

func (d *rackReservationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A reservation of rack units (dcim.rackreservation).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rackreservation/):\n\n> Users can reserve specific units within a rack for future use. An arbitrary set of units within a rack can be associated with a single reservation, but reservations cannot span multiple racks. A description is required for each reservation, reservations may optionally be associated with a specific tenant.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"rack_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the rack.",
			},
			"units": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Rack unit numbers being reserved.",
			},
			"user_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the NetBox user holding the reservation.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Status of the reservation. One of: pending, active, stale.",
				Validators: []validator.String{
					stringvalidator.OneOf("pending", "active", "stale"),
				},
			},
			"description": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Purpose of the reservation.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 200),
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
			"unit_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of rack units reserved.",
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
			"group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the group.",
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the group.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"location_id": schema.StringAttribute{
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
			"unit": schema.Float64Attribute{
				Optional:    true,
				Description: "Unit to match.",
			},
			"unit_count_max": schema.Float64Attribute{
				Optional:    true,
				Description: "Maximum unit count.",
			},
			"unit_count_min": schema.Float64Attribute{
				Optional:    true,
				Description: "Minimum unit count.",
			},
			"user_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the user.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *rackReservationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *rackReservationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data rackReservationDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimRackReservationsListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
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
	if !state.GroupSlug.IsNull() {
		v := state.GroupSlug.ValueString()
		params.SetGroup([]string{v})
		hasInput = true
	}
	if !state.GroupID.IsNull() {
		v := state.GroupID.ValueInt64()
		params.SetGroupID([]int64{v})
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
		v := state.LocationID.ValueString()
		params.SetLocationID([]string{v})
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
	if !state.RackID.IsNull() {
		v := state.RackID.ValueInt64()
		params.SetRackID([]int64{v})
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
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
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
	if !state.Unit.IsNull() {
		v := state.Unit.ValueFloat64()
		params.SetUnit(&v)
		hasInput = true
	}
	if !state.UnitCountMax.IsNull() {
		v := state.UnitCountMax.ValueFloat64()
		params.SetUnitCountMax(&v)
		hasInput = true
	}
	if !state.UnitCountMin.IsNull() {
		v := state.UnitCountMin.ValueFloat64()
		params.SetUnitCountMin(&v)
		hasInput = true
	}
	if !state.UserName.IsNull() {
		v := state.UserName.ValueString()
		params.SetUser([]string{v})
		hasInput = true
	}
	if !state.UserID.IsNull() {
		v := state.UserID.ValueInt64()
		params.SetUserID([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_rack_reservation.")
		return
	}
	res, err := d.client.Dcim.DcimRackReservationsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_rack_reservation", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_rack_reservation",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenRackReservation(ctx, netboxapi.RackReservationResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.rackReservationResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
