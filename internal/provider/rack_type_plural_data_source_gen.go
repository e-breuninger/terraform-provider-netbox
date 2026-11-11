// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*rackTypesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*rackTypesDataSource)(nil)
)

// NewRackTypesDataSource returns a new rack_types data source, which
// lists rack_type objects matching its filters.
func NewRackTypesDataSource() datasource.DataSource {
	return &rackTypesDataSource{}
}

type rackTypesDataSource struct {
	client *netboxapi.Client
}

// rackTypesDataSourceModel is the data source model: the filters, the limit and the matching
// rack_types.
type rackTypesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"rack_types"`
}

// rackTypeResourceAttrTypes is the attribute type map of rackTypeResourceModel.
func rackTypeResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                types.Int64Type,
		"manufacturer_id":   types.Int64Type,
		"model":             types.StringType,
		"slug":              types.StringType,
		"description":       types.StringType,
		"form_factor":       types.StringType,
		"width":             types.Int64Type,
		"u_height":          types.Int64Type,
		"starting_unit":     types.Int64Type,
		"desc_units":        types.BoolType,
		"outer_width":       types.Int64Type,
		"outer_height":      types.Int64Type,
		"outer_depth":       types.Int64Type,
		"outer_unit":        types.StringType,
		"mounting_depth":    types.Int64Type,
		"mounting_depth_mm": types.Int64Type,
		"weight":            types.Float64Type,
		"max_weight":        types.Int64Type,
		"weight_unit":       types.StringType,
		"comments":          types.StringType,
		"owner_id":          types.Int64Type,
		"created":           types.StringType,
		"last_updated":      types.StringType,
		"url":               types.StringType,
		"rack_count":        types.Int64Type,
		"tags":              types.SetType{ElemType: types.StringType},
		"tags_all":          types.SetType{ElemType: types.StringType},
		"custom_fields":     types.MapType{ElemType: types.StringType},
	}
}

func (d *rackTypesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rack_types"
}

func (d *rackTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists rack_type objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, desc_units, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, form_factor, form_factor__empty, form_factor__ic, form_factor__ie, form_factor__iew, form_factor__iregex, form_factor__isw, form_factor__n, form_factor__nic, form_factor__nie, form_factor__niew, form_factor__nisw, form_factor__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, max_weight, max_weight__empty, max_weight__gt, max_weight__gte, max_weight__lt, max_weight__lte, max_weight__n, model, model__empty, model__ic, model__ie, model__iew, model__iregex, model__isw, model__n, model__nic, model__nie, model__niew, model__nisw, model__regex, mounting_depth, mounting_depth__empty, mounting_depth__gt, mounting_depth__gte, mounting_depth__lt, mounting_depth__lte, mounting_depth__n, outer_depth, outer_depth__empty, outer_depth__gt, outer_depth__gte, outer_depth__lt, outer_depth__lte, outer_depth__n, outer_height, outer_height__empty, outer_height__gt, outer_height__gte, outer_height__lt, outer_height__lte, outer_height__n, outer_unit, outer_unit__empty, outer_unit__ic, outer_unit__ie, outer_unit__iew, outer_unit__iregex, outer_unit__isw, outer_unit__n, outer_unit__nic, outer_unit__nie, outer_unit__niew, outer_unit__nisw, outer_unit__regex, outer_width, outer_width__empty, outer_width__gt, outer_width__gte, outer_width__lt, outer_width__lte, outer_width__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, rack_count, rack_count__empty, rack_count__gt, rack_count__gte, rack_count__lt, rack_count__lte, rack_count__n, slug, slug__empty, slug__ic, slug__ie, slug__iew, slug__iregex, slug__isw, slug__n, slug__nic, slug__nie, slug__niew, slug__nisw, slug__regex, starting_unit, starting_unit__empty, starting_unit__gt, starting_unit__gte, starting_unit__lt, starting_unit__lte, starting_unit__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, u_height, u_height__empty, u_height__gt, u_height__gte, u_height__lt, u_height__lte, u_height__n, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n, weight_unit, weight_unit__empty, weight_unit__ic, weight_unit__ie, weight_unit__iew, weight_unit__iregex, weight_unit__isw, weight_unit__n, weight_unit__nic, weight_unit__nie, weight_unit__niew, weight_unit__nisw, weight_unit__regex, width, width__ic, width__ie, width__iew, width__iregex, width__isw, width__n, width__nic, width__nie, width__niew, width__nisw, width__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"rack_types": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching rack_types.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"manufacturer_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the manufacturer.",
						},
						"model": schema.StringAttribute{
							Computed: true,
						},
						"slug": schema.StringAttribute{
							Computed:    true,
							Description: "URL-friendly unique shorthand; derived from model when not set.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"form_factor": schema.StringAttribute{
							Computed:    true,
							Description: "Rack form factor. One of: 2-post-frame, 4-post-frame, 4-post-cabinet, wall-frame, wall-frame-vertical, wall-cabinet, wall-cabinet-vertical.",
						},
						"width": schema.Int64Attribute{
							Computed:    true,
							Description: "Rail-to-rail width in inches. One of: 10, 19, 21, 23.",
						},
						"u_height": schema.Int64Attribute{
							Computed:    true,
							Description: "Height in rack units. Must be a whole number.",
						},
						"starting_unit": schema.Int64Attribute{
							Computed:    true,
							Description: "Lowest unit number. Must be a whole number.",
						},
						"desc_units": schema.BoolAttribute{
							Computed:    true,
							Description: "Units are numbered top-to-bottom.",
						},
						"outer_width": schema.Int64Attribute{
							Computed:    true,
							Description: "Outer width (requires outer_unit). Must be a whole number.",
						},
						"outer_height": schema.Int64Attribute{
							Computed:    true,
							Description: "Outer height (requires outer_unit). Must be a whole number.",
						},
						"outer_depth": schema.Int64Attribute{
							Computed:    true,
							Description: "Outer depth (requires outer_unit). Must be a whole number.",
						},
						"outer_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of the outer dimensions. One of: mm, in.",
						},
						"mounting_depth": schema.Int64Attribute{
							Computed:    true,
							Description: "Maximum depth of a mounted device, in millimetres. Must be a whole number.",
						},
						"mounting_depth_mm": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use mounting_depth instead.",
							Description:        "Deprecated alias of mounting_depth.",
						},
						"weight": schema.Float64Attribute{
							Computed:    true,
							Description: "Weight of the rack (requires weight_unit).",
						},
						"max_weight": schema.Int64Attribute{
							Computed:    true,
							Description: "Maximum load capacity (requires weight_unit). Must be a whole number.",
						},
						"weight_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of weight and max_weight. One of: kg, g, lb, oz.",
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"owner_id": schema.Int64Attribute{
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
						"rack_count": schema.Int64Attribute{
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
					},
				},
			},
		},
	}
}

func (d *rackTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *rackTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data rackTypesDataSourceModel

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
	var responseDTOs []*netboxapi.RackTypeResponseDTO

	params := dcim.NewDcimRackTypesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"model":                 netboxapi.FilterStrings(params.SetModel),
		"model__n":              netboxapi.FilterStrings(params.SetModeln),
		"model__ic":             netboxapi.FilterStrings(params.SetModelIc),
		"model__nic":            netboxapi.FilterStrings(params.SetModelNic),
		"model__ie":             netboxapi.FilterStrings(params.SetModelIe),
		"model__nie":            netboxapi.FilterStrings(params.SetModelNie),
		"model__isw":            netboxapi.FilterStrings(params.SetModelIsw),
		"model__nisw":           netboxapi.FilterStrings(params.SetModelNisw),
		"model__iew":            netboxapi.FilterStrings(params.SetModelIew),
		"model__niew":           netboxapi.FilterStrings(params.SetModelNiew),
		"model__empty":          netboxapi.FilterBool(params.SetModelEmpty),
		"model__regex":          netboxapi.FilterStrings(params.SetModelRegex),
		"model__iregex":         netboxapi.FilterStrings(params.SetModelIregex),
		"slug":                  netboxapi.FilterStrings(params.SetSlug),
		"slug__n":               netboxapi.FilterStrings(params.SetSlugn),
		"slug__ic":              netboxapi.FilterStrings(params.SetSlugIc),
		"slug__nic":             netboxapi.FilterStrings(params.SetSlugNic),
		"slug__ie":              netboxapi.FilterStrings(params.SetSlugIe),
		"slug__nie":             netboxapi.FilterStrings(params.SetSlugNie),
		"slug__isw":             netboxapi.FilterStrings(params.SetSlugIsw),
		"slug__nisw":            netboxapi.FilterStrings(params.SetSlugNisw),
		"slug__iew":             netboxapi.FilterStrings(params.SetSlugIew),
		"slug__niew":            netboxapi.FilterStrings(params.SetSlugNiew),
		"slug__empty":           netboxapi.FilterBool(params.SetSlugEmpty),
		"slug__regex":           netboxapi.FilterStrings(params.SetSlugRegex),
		"slug__iregex":          netboxapi.FilterStrings(params.SetSlugIregex),
		"manufacturer_id":       netboxapi.FilterInt64s(params.SetManufacturerID),
		"manufacturer_id__n":    netboxapi.FilterInt64s(params.SetManufacturerIDn),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
		"desc_units":            netboxapi.FilterBool(params.SetDescUnits),
		"description":           netboxapi.FilterStrings(params.SetDescription),
		"description__n":        netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":       netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":      netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":       netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":      netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":      netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":     netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":      netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":     netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":    netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":    netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":   netboxapi.FilterStrings(params.SetDescriptionIregex),
		"form_factor":           netboxapi.FilterStrings(params.SetFormFactor),
		"form_factor__n":        netboxapi.FilterStrings(params.SetFormFactorn),
		"form_factor__ic":       netboxapi.FilterStrings(params.SetFormFactorIc),
		"form_factor__nic":      netboxapi.FilterStrings(params.SetFormFactorNic),
		"form_factor__ie":       netboxapi.FilterStrings(params.SetFormFactorIe),
		"form_factor__nie":      netboxapi.FilterStrings(params.SetFormFactorNie),
		"form_factor__isw":      netboxapi.FilterStrings(params.SetFormFactorIsw),
		"form_factor__nisw":     netboxapi.FilterStrings(params.SetFormFactorNisw),
		"form_factor__iew":      netboxapi.FilterStrings(params.SetFormFactorIew),
		"form_factor__niew":     netboxapi.FilterStrings(params.SetFormFactorNiew),
		"form_factor__empty":    netboxapi.FilterBool(params.SetFormFactorEmpty),
		"form_factor__regex":    netboxapi.FilterStrings(params.SetFormFactorRegex),
		"form_factor__iregex":   netboxapi.FilterStrings(params.SetFormFactorIregex),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"manufacturer":          netboxapi.FilterStrings(params.SetManufacturer),
		"manufacturer__n":       netboxapi.FilterStrings(params.SetManufacturern),
		"max_weight":            netboxapi.FilterInt64s(params.SetMaxWeight),
		"max_weight__n":         netboxapi.FilterInt64s(params.SetMaxWeightn),
		"max_weight__lt":        netboxapi.FilterInt64s(params.SetMaxWeightLt),
		"max_weight__lte":       netboxapi.FilterInt64s(params.SetMaxWeightLte),
		"max_weight__gt":        netboxapi.FilterInt64s(params.SetMaxWeightGt),
		"max_weight__gte":       netboxapi.FilterInt64s(params.SetMaxWeightGte),
		"max_weight__empty":     netboxapi.FilterBool(params.SetMaxWeightEmpty),
		"mounting_depth":        netboxapi.FilterInt64s(params.SetMountingDepth),
		"mounting_depth__n":     netboxapi.FilterInt64s(params.SetMountingDepthn),
		"mounting_depth__lt":    netboxapi.FilterInt64s(params.SetMountingDepthLt),
		"mounting_depth__lte":   netboxapi.FilterInt64s(params.SetMountingDepthLte),
		"mounting_depth__gt":    netboxapi.FilterInt64s(params.SetMountingDepthGt),
		"mounting_depth__gte":   netboxapi.FilterInt64s(params.SetMountingDepthGte),
		"mounting_depth__empty": netboxapi.FilterBool(params.SetMountingDepthEmpty),
		"outer_depth":           netboxapi.FilterInt64s(params.SetOuterDepth),
		"outer_depth__n":        netboxapi.FilterInt64s(params.SetOuterDepthn),
		"outer_depth__lt":       netboxapi.FilterInt64s(params.SetOuterDepthLt),
		"outer_depth__lte":      netboxapi.FilterInt64s(params.SetOuterDepthLte),
		"outer_depth__gt":       netboxapi.FilterInt64s(params.SetOuterDepthGt),
		"outer_depth__gte":      netboxapi.FilterInt64s(params.SetOuterDepthGte),
		"outer_depth__empty":    netboxapi.FilterBool(params.SetOuterDepthEmpty),
		"outer_height":          netboxapi.FilterInt64s(params.SetOuterHeight),
		"outer_height__n":       netboxapi.FilterInt64s(params.SetOuterHeightn),
		"outer_height__lt":      netboxapi.FilterInt64s(params.SetOuterHeightLt),
		"outer_height__lte":     netboxapi.FilterInt64s(params.SetOuterHeightLte),
		"outer_height__gt":      netboxapi.FilterInt64s(params.SetOuterHeightGt),
		"outer_height__gte":     netboxapi.FilterInt64s(params.SetOuterHeightGte),
		"outer_height__empty":   netboxapi.FilterBool(params.SetOuterHeightEmpty),
		"outer_unit":            netboxapi.FilterString(params.SetOuterUnit),
		"outer_unit__n":         netboxapi.FilterString(params.SetOuterUnitn),
		"outer_unit__ic":        netboxapi.FilterStrings(params.SetOuterUnitIc),
		"outer_unit__nic":       netboxapi.FilterStrings(params.SetOuterUnitNic),
		"outer_unit__ie":        netboxapi.FilterStrings(params.SetOuterUnitIe),
		"outer_unit__nie":       netboxapi.FilterStrings(params.SetOuterUnitNie),
		"outer_unit__isw":       netboxapi.FilterStrings(params.SetOuterUnitIsw),
		"outer_unit__nisw":      netboxapi.FilterStrings(params.SetOuterUnitNisw),
		"outer_unit__iew":       netboxapi.FilterStrings(params.SetOuterUnitIew),
		"outer_unit__niew":      netboxapi.FilterStrings(params.SetOuterUnitNiew),
		"outer_unit__empty":     netboxapi.FilterBool(params.SetOuterUnitEmpty),
		"outer_unit__regex":     netboxapi.FilterStrings(params.SetOuterUnitRegex),
		"outer_unit__iregex":    netboxapi.FilterStrings(params.SetOuterUnitIregex),
		"outer_width":           netboxapi.FilterInt64s(params.SetOuterWidth),
		"outer_width__n":        netboxapi.FilterInt64s(params.SetOuterWidthn),
		"outer_width__lt":       netboxapi.FilterInt64s(params.SetOuterWidthLt),
		"outer_width__lte":      netboxapi.FilterInt64s(params.SetOuterWidthLte),
		"outer_width__gt":       netboxapi.FilterInt64s(params.SetOuterWidthGt),
		"outer_width__gte":      netboxapi.FilterInt64s(params.SetOuterWidthGte),
		"outer_width__empty":    netboxapi.FilterBool(params.SetOuterWidthEmpty),
		"owner":                 netboxapi.FilterStrings(params.SetOwner),
		"owner__n":              netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":           netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":        netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":        netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":     netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                     netboxapi.FilterString(params.SetQ),
		"rack_count":            netboxapi.FilterInt64s(params.SetRackCount),
		"rack_count__n":         netboxapi.FilterInt64s(params.SetRackCountn),
		"rack_count__lt":        netboxapi.FilterInt64s(params.SetRackCountLt),
		"rack_count__lte":       netboxapi.FilterInt64s(params.SetRackCountLte),
		"rack_count__gt":        netboxapi.FilterInt64s(params.SetRackCountGt),
		"rack_count__gte":       netboxapi.FilterInt64s(params.SetRackCountGte),
		"rack_count__empty":     netboxapi.FilterBool(params.SetRackCountEmpty),
		"starting_unit":         netboxapi.FilterInt64s(params.SetStartingUnit),
		"starting_unit__n":      netboxapi.FilterInt64s(params.SetStartingUnitn),
		"starting_unit__lt":     netboxapi.FilterInt64s(params.SetStartingUnitLt),
		"starting_unit__lte":    netboxapi.FilterInt64s(params.SetStartingUnitLte),
		"starting_unit__gt":     netboxapi.FilterInt64s(params.SetStartingUnitGt),
		"starting_unit__gte":    netboxapi.FilterInt64s(params.SetStartingUnitGte),
		"starting_unit__empty":  netboxapi.FilterBool(params.SetStartingUnitEmpty),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"u_height":              netboxapi.FilterInt64s(params.SetUHeight),
		"u_height__n":           netboxapi.FilterInt64s(params.SetUHeightn),
		"u_height__lt":          netboxapi.FilterInt64s(params.SetUHeightLt),
		"u_height__lte":         netboxapi.FilterInt64s(params.SetUHeightLte),
		"u_height__gt":          netboxapi.FilterInt64s(params.SetUHeightGt),
		"u_height__gte":         netboxapi.FilterInt64s(params.SetUHeightGte),
		"u_height__empty":       netboxapi.FilterBool(params.SetUHeightEmpty),
		"weight":                netboxapi.FilterFloat64s(params.SetWeight),
		"weight__n":             netboxapi.FilterFloat64s(params.SetWeightn),
		"weight__lt":            netboxapi.FilterFloat64s(params.SetWeightLt),
		"weight__lte":           netboxapi.FilterFloat64s(params.SetWeightLte),
		"weight__gt":            netboxapi.FilterFloat64s(params.SetWeightGt),
		"weight__gte":           netboxapi.FilterFloat64s(params.SetWeightGte),
		"weight__empty":         netboxapi.FilterBool(params.SetWeightEmpty),
		"weight_unit":           netboxapi.FilterString(params.SetWeightUnit),
		"weight_unit__n":        netboxapi.FilterString(params.SetWeightUnitn),
		"weight_unit__ic":       netboxapi.FilterStrings(params.SetWeightUnitIc),
		"weight_unit__nic":      netboxapi.FilterStrings(params.SetWeightUnitNic),
		"weight_unit__ie":       netboxapi.FilterStrings(params.SetWeightUnitIe),
		"weight_unit__nie":      netboxapi.FilterStrings(params.SetWeightUnitNie),
		"weight_unit__isw":      netboxapi.FilterStrings(params.SetWeightUnitIsw),
		"weight_unit__nisw":     netboxapi.FilterStrings(params.SetWeightUnitNisw),
		"weight_unit__iew":      netboxapi.FilterStrings(params.SetWeightUnitIew),
		"weight_unit__niew":     netboxapi.FilterStrings(params.SetWeightUnitNiew),
		"weight_unit__empty":    netboxapi.FilterBool(params.SetWeightUnitEmpty),
		"weight_unit__regex":    netboxapi.FilterStrings(params.SetWeightUnitRegex),
		"weight_unit__iregex":   netboxapi.FilterStrings(params.SetWeightUnitIregex),
		"width":                 netboxapi.FilterInt64s(params.SetWidth),
		"width__n":              netboxapi.FilterInt64s(params.SetWidthn),
		"width__ic":             netboxapi.FilterInt64s(params.SetWidthIc),
		"width__nic":            netboxapi.FilterInt64s(params.SetWidthNic),
		"width__ie":             netboxapi.FilterInt64s(params.SetWidthIe),
		"width__nie":            netboxapi.FilterInt64s(params.SetWidthNie),
		"width__isw":            netboxapi.FilterInt64s(params.SetWidthIsw),
		"width__nisw":           netboxapi.FilterInt64s(params.SetWidthNisw),
		"width__iew":            netboxapi.FilterInt64s(params.SetWidthIew),
		"width__niew":           netboxapi.FilterInt64s(params.SetWidthNiew),
		"width__regex":          netboxapi.FilterInt64s(params.SetWidthRegex),
		"width__iregex":         netboxapi.FilterInt64s(params.SetWidthIregex),
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
		res, err := d.client.Dcim.DcimRackTypesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_rack_types", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.RackTypeResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]rackTypeResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model rackTypeResourceModel
		resp.Diagnostics.Append(flattenRackType(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: rackTypeResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
