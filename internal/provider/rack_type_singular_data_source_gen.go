// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

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
	_ datasource.DataSource              = (*rackTypeDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*rackTypeDataSource)(nil)
)

// NewRackTypeDataSource returns a new rack_type data source.
func NewRackTypeDataSource() datasource.DataSource {
	return &rackTypeDataSource{}
}

type rackTypeDataSource struct {
	client *netboxapi.Client
}

// rackTypeDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type rackTypeDataSourceModel struct {
	rackTypeResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	ManufacturerSlug   types.String `tfsdk:"manufacturer_slug"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *rackTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rack_type"
}

func (d *rackTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A NetBox rack type (dcim.racktype): a rack model made by a manufacturer, whose physical attributes racks inherit.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/en/stable/models/dcim/racktype/):\n\n> A rack type defines the physical characteristics of a particular model of rack.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"manufacturer_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the manufacturer.",
			},
			"model": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from model when not set.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"form_factor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rack form factor. One of: 2-post-frame, 4-post-frame, 4-post-cabinet, wall-frame, wall-frame-vertical, wall-cabinet, wall-cabinet-vertical.",
				Validators: []validator.String{
					stringvalidator.OneOf("2-post-frame", "4-post-frame", "4-post-cabinet", "wall-frame", "wall-frame-vertical", "wall-cabinet", "wall-cabinet-vertical"),
				},
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
				Description: "Height in rack units. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(1, 100),
				},
			},
			"starting_unit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Lowest unit number. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(1, 32767),
				},
			},
			"desc_units": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Units are numbered top-to-bottom.",
			},
			"outer_width": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Outer width (requires outer_unit). Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"outer_height": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Outer height (requires outer_unit). Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"outer_depth": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Outer depth (requires outer_unit). Must be a whole number.",
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
				Description: "Maximum depth of a mounted device, in millimetres. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"mounting_depth_mm": schema.Int64Attribute{
				Computed:           true,
				DeprecationMessage: "Use mounting_depth instead.",
				Description:        "Deprecated alias of mounting_depth.",
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
				Description: "Maximum load capacity (requires weight_unit). Must be a whole number.",
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
			"rack_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of racks of the rack type.",
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
			"manufacturer_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the manufacturer.",
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
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *rackTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *rackTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data rackTypeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimRackTypesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Model.IsNull() {
		v := state.Model.ValueString()
		params.SetModel([]string{v})
		hasInput = true
	}
	if !state.Slug.IsNull() {
		v := state.Slug.ValueString()
		params.SetSlug([]string{v})
		hasInput = true
	}
	if !state.ManufacturerID.IsNull() {
		v := state.ManufacturerID.ValueInt64()
		params.SetManufacturerID([]int64{v})
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
	if !state.DescUnits.IsNull() {
		v := state.DescUnits.ValueBool()
		params.SetDescUnits(&v)
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.FormFactor.IsNull() {
		v := state.FormFactor.ValueString()
		params.SetFormFactor([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.ManufacturerSlug.IsNull() {
		v := state.ManufacturerSlug.ValueString()
		params.SetManufacturer([]string{v})
		hasInput = true
	}
	if !state.MaxWeight.IsNull() {
		v := state.MaxWeight.ValueInt64()
		params.SetMaxWeight([]int64{v})
		hasInput = true
	}
	if !state.MountingDepth.IsNull() {
		v := state.MountingDepth.ValueInt64()
		params.SetMountingDepth([]int64{v})
		hasInput = true
	}
	if !state.OuterDepth.IsNull() {
		v := state.OuterDepth.ValueInt64()
		params.SetOuterDepth([]int64{v})
		hasInput = true
	}
	if !state.OuterHeight.IsNull() {
		v := state.OuterHeight.ValueInt64()
		params.SetOuterHeight([]int64{v})
		hasInput = true
	}
	if !state.OuterUnit.IsNull() {
		v := state.OuterUnit.ValueString()
		params.SetOuterUnit(&v)
		hasInput = true
	}
	if !state.OuterWidth.IsNull() {
		v := state.OuterWidth.ValueInt64()
		params.SetOuterWidth([]int64{v})
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
	if !state.RackCount.IsNull() {
		v := state.RackCount.ValueInt64()
		params.SetRackCount([]int64{v})
		hasInput = true
	}
	if !state.StartingUnit.IsNull() {
		v := state.StartingUnit.ValueInt64()
		params.SetStartingUnit([]int64{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.UHeight.IsNull() {
		v := state.UHeight.ValueInt64()
		params.SetUHeight([]int64{v})
		hasInput = true
	}
	if !state.Weight.IsNull() {
		v := state.Weight.ValueFloat64()
		params.SetWeight([]float64{v})
		hasInput = true
	}
	if !state.WeightUnit.IsNull() {
		v := state.WeightUnit.ValueString()
		params.SetWeightUnit(&v)
		hasInput = true
	}
	if !state.Width.IsNull() {
		v := state.Width.ValueInt64()
		params.SetWidth([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_rack_type.")
		return
	}
	res, err := d.client.Dcim.DcimRackTypesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_rack_type", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_rack_type",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenRackType(ctx, netboxapi.RackTypeResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.rackTypeResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
