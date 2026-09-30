// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*rackDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*rackDataSource)(nil)
)

// NewRackDataSource returns a new rack data source.
func NewRackDataSource() datasource.DataSource {
	return &rackDataSource{}
}

type rackDataSource struct {
	client *netboxapi.Client
}

// rackDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type rackDataSourceModel struct {
	rackResourceModel
	NameContains       types.String `tfsdk:"name_contains"`
	RegionID           types.Int64  `tfsdk:"region_id"`
	ContactID          types.Int64  `tfsdk:"contact_id"`
	ContactGroupID     types.Int64  `tfsdk:"contact_group_id"`
	ContactRoleID      types.Int64  `tfsdk:"contact_role_id"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *rackDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rack"
}

func (d *rackDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A NetBox rack (dcim.rack).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/rack/):\n\n> The rack model represents a physical two- or four-post equipment rack in which devices can be installed. Each rack must be assigned to a site, and may optionally be assigned to a location within that site. Racks can also be organized by user-defined functional roles. The name and facility ID of each rack within a location must be unique.\n\nRack height is measured in rack units (U); racks are commonly between 42U and 48U tall, but NetBox allows you to define racks of arbitrary height. A toggle is provided to indicate whether rack units are in ascending (from the ground up) or descending order.\n\nEach rack is assigned a name and (optionally) a separate facility ID. This is helpful when leasing space in a data center your organization does not own: The facility will often assign a seemingly arbitrary ID to a rack (for example, \"M204.313\") whereas internally you refer to is simply as \"R113.\" A unique serial number and asset tag may also be associated with each rack.",
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
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the site the rack belongs to.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the location within the site.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the rack role.",
			},
			"group_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the rack group.",
			},
			"rack_type_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the rack type. NetBox copies the type's form factor, width, height and outer dimensions onto the rack; set the same values here too, or the next plan shows them as drift.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: reserved, available, planned, active, deprecated.",
				Validators: []validator.String{
					stringvalidator.OneOf("reserved", "available", "planned", "active", "deprecated"),
				},
			},
			"form_factor": schema.StringAttribute{
				Computed:    true,
				Description: "Rack form factor. One of: 2-post-frame, 4-post-frame, 4-post-cabinet, wall-frame, wall-frame-vertical, wall-cabinet, wall-cabinet-vertical.",
			},
			"airflow": schema.StringAttribute{
				Computed:    true,
				Description: "Airflow direction. One of: front-to-rear, rear-to-front.",
			},
			"width": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Rail-to-rail width in inches. One of: 10, 19, 21, 23.",
				Validators: []validator.Int64{
					int64validator.OneOf(10, 19, 21, 23),
				},
			},
			"u_height": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Height in rack units.",
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"starting_unit": schema.Int64Attribute{
				Computed:    true,
				Description: "Lowest unit number.",
			},
			"desc_units": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units are numbered top-to-bottom.",
			},
			"serial": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Serial number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"asset_tag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique asset tag.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"facility_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Locally-assigned facility identifier.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"outer_width": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Outer width (requires outer_unit).",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"outer_height": schema.Int64Attribute{
				Computed:    true,
				Description: "Outer height (requires outer_unit).",
			},
			"outer_depth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Outer depth (requires outer_unit).",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"outer_unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of the outer dimensions. One of: mm, in.",
				Validators: []validator.String{
					stringvalidator.OneOf("mm", "in"),
				},
			},
			"mounting_depth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum depth of a mounted device in millimeters.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"weight": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of the rack (requires weight_unit).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"max_weight": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum load capacity (requires weight_unit).",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"weight_unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of weight and max_weight. One of: kg, g, lb, oz.",
				Validators: []validator.String{
					stringvalidator.OneOf("kg", "g", "lb", "oz"),
				},
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
			"device_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of devices in the rack.",
			},
			"power_feed_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of power feeds in the rack.",
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
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
			},
			"region_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the region of the site.",
			},
			"contact_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a contact assigned to the rack.",
			},
			"contact_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a contact group with a contact assigned to the rack.",
			},
			"contact_role_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the role a contact is assigned to the rack with.",
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

func (d *rackDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *rackDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data rackDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimRacksListParams()
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
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.LocationID.IsNull() {
		v := strconv.FormatInt(state.LocationID.ValueInt64(), 10)
		params.SetLocationID([]string{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.RoleID.IsNull() {
		v := state.RoleID.ValueInt64()
		params.SetRoleID([]int64{v})
		hasInput = true
	}
	if !state.RackTypeID.IsNull() {
		v := state.RackTypeID.ValueInt64()
		params.SetRackTypeID([]int64{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.Width.IsNull() {
		v := state.Width.ValueInt64()
		params.SetWidth([]int64{v})
		hasInput = true
	}
	if !state.UHeight.IsNull() {
		v := state.UHeight.ValueInt64()
		params.SetUHeight([]int64{v})
		hasInput = true
	}
	if !state.DescUnits.IsNull() {
		v := state.DescUnits.ValueBool()
		params.SetDescUnits(&v)
		hasInput = true
	}
	if !state.Serial.IsNull() {
		v := state.Serial.ValueString()
		params.SetSerial([]string{v})
		hasInput = true
	}
	if !state.AssetTag.IsNull() {
		v := state.AssetTag.ValueString()
		params.SetAssetTag([]string{v})
		hasInput = true
	}
	if !state.FacilityID.IsNull() {
		v := state.FacilityID.ValueString()
		params.SetFacilityID([]string{v})
		hasInput = true
	}
	if !state.OuterWidth.IsNull() {
		v := state.OuterWidth.ValueInt64()
		params.SetOuterWidth([]int64{v})
		hasInput = true
	}
	if !state.OuterDepth.IsNull() {
		v := state.OuterDepth.ValueInt64()
		params.SetOuterDepth([]int64{v})
		hasInput = true
	}
	if !state.OuterUnit.IsNull() {
		v := state.OuterUnit.ValueString()
		params.SetOuterUnit(&v)
		hasInput = true
	}
	if !state.MountingDepth.IsNull() {
		v := state.MountingDepth.ValueInt64()
		params.SetMountingDepth([]int64{v})
		hasInput = true
	}
	if !state.Weight.IsNull() {
		v := state.Weight.ValueFloat64()
		params.SetWeight([]float64{v})
		hasInput = true
	}
	if !state.MaxWeight.IsNull() {
		v := state.MaxWeight.ValueInt64()
		params.SetMaxWeight([]int64{v})
		hasInput = true
	}
	if !state.WeightUnit.IsNull() {
		v := state.WeightUnit.ValueString()
		params.SetWeightUnit(&v)
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := strconv.FormatInt(state.RegionID.ValueInt64(), 10)
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.ContactID.IsNull() {
		v := state.ContactID.ValueInt64()
		params.SetContact([]int64{v})
		hasInput = true
	}
	if !state.ContactGroupID.IsNull() {
		v := strconv.FormatInt(state.ContactGroupID.ValueInt64(), 10)
		params.SetContactGroup([]string{v})
		hasInput = true
	}
	if !state.ContactRoleID.IsNull() {
		v := state.ContactRoleID.ValueInt64()
		params.SetContactRole([]int64{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or name_contains and/or site_id and/or location_id and/or tenant_id and/or role_id and/or rack_type_id and/or status and/or width and/or u_height and/or desc_units and/or serial and/or asset_tag and/or facility_id and/or outer_width and/or outer_depth and/or outer_unit and/or mounting_depth and/or weight and/or max_weight and/or weight_unit and/or region_id and/or contact_id and/or contact_group_id and/or contact_role_id and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_rack.")
		return
	}
	res, err := d.client.Dcim.DcimRacksListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_rack", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_rack",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenRack(ctx, netboxapi.RackResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.rackResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
