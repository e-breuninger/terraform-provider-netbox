// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vModel []string
	var vModeln []string
	var vModelIc []string
	var vModelNic []string
	var vModelIe []string
	var vModelNie []string
	var vModelIsw []string
	var vModelNisw []string
	var vModelIew []string
	var vModelNiew []string
	var vModelEmpty *bool
	var vModelRegex []string
	var vModelIregex []string
	var vSlug []string
	var vSlugn []string
	var vSlugIc []string
	var vSlugNic []string
	var vSlugIe []string
	var vSlugNie []string
	var vSlugIsw []string
	var vSlugNisw []string
	var vSlugIew []string
	var vSlugNiew []string
	var vSlugEmpty *bool
	var vSlugRegex []string
	var vSlugIregex []string
	var vManufacturerID []int64
	var vManufacturerIDn []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vDescUnits *bool
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
	var vFormFactor []string
	var vFormFactorn []string
	var vFormFactorIc []string
	var vFormFactorNic []string
	var vFormFactorIe []string
	var vFormFactorNie []string
	var vFormFactorIsw []string
	var vFormFactorNisw []string
	var vFormFactorIew []string
	var vFormFactorNiew []string
	var vFormFactorEmpty *bool
	var vFormFactorRegex []string
	var vFormFactorIregex []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vManufacturer []string
	var vManufacturern []string
	var vMaxWeight []int64
	var vMaxWeightn []int64
	var vMaxWeightLt []int64
	var vMaxWeightLte []int64
	var vMaxWeightGt []int64
	var vMaxWeightGte []int64
	var vMaxWeightEmpty *bool
	var vMountingDepth []int64
	var vMountingDepthn []int64
	var vMountingDepthLt []int64
	var vMountingDepthLte []int64
	var vMountingDepthGt []int64
	var vMountingDepthGte []int64
	var vMountingDepthEmpty *bool
	var vOuterDepth []int64
	var vOuterDepthn []int64
	var vOuterDepthLt []int64
	var vOuterDepthLte []int64
	var vOuterDepthGt []int64
	var vOuterDepthGte []int64
	var vOuterDepthEmpty *bool
	var vOuterHeight []int64
	var vOuterHeightn []int64
	var vOuterHeightLt []int64
	var vOuterHeightLte []int64
	var vOuterHeightGt []int64
	var vOuterHeightGte []int64
	var vOuterHeightEmpty *bool
	var vOuterUnit *string
	var vOuterUnitn *string
	var vOuterUnitIc []string
	var vOuterUnitNic []string
	var vOuterUnitIe []string
	var vOuterUnitNie []string
	var vOuterUnitIsw []string
	var vOuterUnitNisw []string
	var vOuterUnitIew []string
	var vOuterUnitNiew []string
	var vOuterUnitEmpty *bool
	var vOuterUnitRegex []string
	var vOuterUnitIregex []string
	var vOuterWidth []int64
	var vOuterWidthn []int64
	var vOuterWidthLt []int64
	var vOuterWidthLte []int64
	var vOuterWidthGt []int64
	var vOuterWidthGte []int64
	var vOuterWidthEmpty *bool
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vQ *string
	var vRackCount []int64
	var vRackCountn []int64
	var vRackCountLt []int64
	var vRackCountLte []int64
	var vRackCountGt []int64
	var vRackCountGte []int64
	var vRackCountEmpty *bool
	var vStartingUnit []int64
	var vStartingUnitn []int64
	var vStartingUnitLt []int64
	var vStartingUnitLte []int64
	var vStartingUnitGt []int64
	var vStartingUnitGte []int64
	var vStartingUnitEmpty *bool
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vUHeight []int64
	var vUHeightn []int64
	var vUHeightLt []int64
	var vUHeightLte []int64
	var vUHeightGt []int64
	var vUHeightGte []int64
	var vUHeightEmpty *bool
	var vWeight []float64
	var vWeightn []float64
	var vWeightLt []float64
	var vWeightLte []float64
	var vWeightGt []float64
	var vWeightGte []float64
	var vWeightEmpty *bool
	var vWeightUnit *string
	var vWeightUnitn *string
	var vWeightUnitIc []string
	var vWeightUnitNic []string
	var vWeightUnitIe []string
	var vWeightUnitNie []string
	var vWeightUnitIsw []string
	var vWeightUnitNisw []string
	var vWeightUnitIew []string
	var vWeightUnitNiew []string
	var vWeightUnitEmpty *bool
	var vWeightUnitRegex []string
	var vWeightUnitIregex []string
	var vWidth []int64
	var vWidthn []int64
	var vWidthIc []int64
	var vWidthNic []int64
	var vWidthIe []int64
	var vWidthNie []int64
	var vWidthIsw []int64
	var vWidthNisw []int64
	var vWidthIew []int64
	var vWidthNiew []int64
	var vWidthRegex []int64
	var vWidthIregex []int64
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
		case "model":
			v := value
			vModel = append(vModel, v)
		case "model__n":
			v := value
			vModeln = append(vModeln, v)
		case "model__ic":
			v := value
			vModelIc = append(vModelIc, v)
		case "model__nic":
			v := value
			vModelNic = append(vModelNic, v)
		case "model__ie":
			v := value
			vModelIe = append(vModelIe, v)
		case "model__nie":
			v := value
			vModelNie = append(vModelNie, v)
		case "model__isw":
			v := value
			vModelIsw = append(vModelIsw, v)
		case "model__nisw":
			v := value
			vModelNisw = append(vModelNisw, v)
		case "model__iew":
			v := value
			vModelIew = append(vModelIew, v)
		case "model__niew":
			v := value
			vModelNiew = append(vModelNiew, v)
		case "model__empty":
			if vModelEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'model__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'model__empty' takes a boolean, got %q.", value))
				return
			}
			vModelEmpty = &v
		case "model__regex":
			v := value
			vModelRegex = append(vModelRegex, v)
		case "model__iregex":
			v := value
			vModelIregex = append(vModelIregex, v)
		case "slug":
			v := value
			vSlug = append(vSlug, v)
		case "slug__n":
			v := value
			vSlugn = append(vSlugn, v)
		case "slug__ic":
			v := value
			vSlugIc = append(vSlugIc, v)
		case "slug__nic":
			v := value
			vSlugNic = append(vSlugNic, v)
		case "slug__ie":
			v := value
			vSlugIe = append(vSlugIe, v)
		case "slug__nie":
			v := value
			vSlugNie = append(vSlugNie, v)
		case "slug__isw":
			v := value
			vSlugIsw = append(vSlugIsw, v)
		case "slug__nisw":
			v := value
			vSlugNisw = append(vSlugNisw, v)
		case "slug__iew":
			v := value
			vSlugIew = append(vSlugIew, v)
		case "slug__niew":
			v := value
			vSlugNiew = append(vSlugNiew, v)
		case "slug__empty":
			if vSlugEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'slug__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'slug__empty' takes a boolean, got %q.", value))
				return
			}
			vSlugEmpty = &v
		case "slug__regex":
			v := value
			vSlugRegex = append(vSlugRegex, v)
		case "slug__iregex":
			v := value
			vSlugIregex = append(vSlugIregex, v)
		case "manufacturer_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'manufacturer_id' takes an integer, got %q.", value))
				return
			}
			vManufacturerID = append(vManufacturerID, v)
		case "manufacturer_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'manufacturer_id__n' takes an integer, got %q.", value))
				return
			}
			vManufacturerIDn = append(vManufacturerIDn, v)
		case "owner_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id' takes an integer, got %q.", value))
				return
			}
			vOwnerID = append(vOwnerID, v)
		case "owner_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerIDn = append(vOwnerIDn, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
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
		case "desc_units":
			if vDescUnits != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'desc_units' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'desc_units' takes a boolean, got %q.", value))
				return
			}
			vDescUnits = &v
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
		case "form_factor":
			v := value
			vFormFactor = append(vFormFactor, v)
		case "form_factor__n":
			v := value
			vFormFactorn = append(vFormFactorn, v)
		case "form_factor__ic":
			v := value
			vFormFactorIc = append(vFormFactorIc, v)
		case "form_factor__nic":
			v := value
			vFormFactorNic = append(vFormFactorNic, v)
		case "form_factor__ie":
			v := value
			vFormFactorIe = append(vFormFactorIe, v)
		case "form_factor__nie":
			v := value
			vFormFactorNie = append(vFormFactorNie, v)
		case "form_factor__isw":
			v := value
			vFormFactorIsw = append(vFormFactorIsw, v)
		case "form_factor__nisw":
			v := value
			vFormFactorNisw = append(vFormFactorNisw, v)
		case "form_factor__iew":
			v := value
			vFormFactorIew = append(vFormFactorIew, v)
		case "form_factor__niew":
			v := value
			vFormFactorNiew = append(vFormFactorNiew, v)
		case "form_factor__empty":
			if vFormFactorEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'form_factor__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'form_factor__empty' takes a boolean, got %q.", value))
				return
			}
			vFormFactorEmpty = &v
		case "form_factor__regex":
			v := value
			vFormFactorRegex = append(vFormFactorRegex, v)
		case "form_factor__iregex":
			v := value
			vFormFactorIregex = append(vFormFactorIregex, v)
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
		case "manufacturer":
			v := value
			vManufacturer = append(vManufacturer, v)
		case "manufacturer__n":
			v := value
			vManufacturern = append(vManufacturern, v)
		case "max_weight":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_weight' takes an integer, got %q.", value))
				return
			}
			vMaxWeight = append(vMaxWeight, v)
		case "max_weight__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_weight__n' takes an integer, got %q.", value))
				return
			}
			vMaxWeightn = append(vMaxWeightn, v)
		case "max_weight__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_weight__lt' takes an integer, got %q.", value))
				return
			}
			vMaxWeightLt = append(vMaxWeightLt, v)
		case "max_weight__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_weight__lte' takes an integer, got %q.", value))
				return
			}
			vMaxWeightLte = append(vMaxWeightLte, v)
		case "max_weight__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_weight__gt' takes an integer, got %q.", value))
				return
			}
			vMaxWeightGt = append(vMaxWeightGt, v)
		case "max_weight__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_weight__gte' takes an integer, got %q.", value))
				return
			}
			vMaxWeightGte = append(vMaxWeightGte, v)
		case "max_weight__empty":
			if vMaxWeightEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'max_weight__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_weight__empty' takes a boolean, got %q.", value))
				return
			}
			vMaxWeightEmpty = &v
		case "mounting_depth":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mounting_depth' takes an integer, got %q.", value))
				return
			}
			vMountingDepth = append(vMountingDepth, v)
		case "mounting_depth__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mounting_depth__n' takes an integer, got %q.", value))
				return
			}
			vMountingDepthn = append(vMountingDepthn, v)
		case "mounting_depth__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mounting_depth__lt' takes an integer, got %q.", value))
				return
			}
			vMountingDepthLt = append(vMountingDepthLt, v)
		case "mounting_depth__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mounting_depth__lte' takes an integer, got %q.", value))
				return
			}
			vMountingDepthLte = append(vMountingDepthLte, v)
		case "mounting_depth__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mounting_depth__gt' takes an integer, got %q.", value))
				return
			}
			vMountingDepthGt = append(vMountingDepthGt, v)
		case "mounting_depth__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mounting_depth__gte' takes an integer, got %q.", value))
				return
			}
			vMountingDepthGte = append(vMountingDepthGte, v)
		case "mounting_depth__empty":
			if vMountingDepthEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mounting_depth__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mounting_depth__empty' takes a boolean, got %q.", value))
				return
			}
			vMountingDepthEmpty = &v
		case "outer_depth":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_depth' takes an integer, got %q.", value))
				return
			}
			vOuterDepth = append(vOuterDepth, v)
		case "outer_depth__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_depth__n' takes an integer, got %q.", value))
				return
			}
			vOuterDepthn = append(vOuterDepthn, v)
		case "outer_depth__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_depth__lt' takes an integer, got %q.", value))
				return
			}
			vOuterDepthLt = append(vOuterDepthLt, v)
		case "outer_depth__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_depth__lte' takes an integer, got %q.", value))
				return
			}
			vOuterDepthLte = append(vOuterDepthLte, v)
		case "outer_depth__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_depth__gt' takes an integer, got %q.", value))
				return
			}
			vOuterDepthGt = append(vOuterDepthGt, v)
		case "outer_depth__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_depth__gte' takes an integer, got %q.", value))
				return
			}
			vOuterDepthGte = append(vOuterDepthGte, v)
		case "outer_depth__empty":
			if vOuterDepthEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'outer_depth__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_depth__empty' takes a boolean, got %q.", value))
				return
			}
			vOuterDepthEmpty = &v
		case "outer_height":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_height' takes an integer, got %q.", value))
				return
			}
			vOuterHeight = append(vOuterHeight, v)
		case "outer_height__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_height__n' takes an integer, got %q.", value))
				return
			}
			vOuterHeightn = append(vOuterHeightn, v)
		case "outer_height__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_height__lt' takes an integer, got %q.", value))
				return
			}
			vOuterHeightLt = append(vOuterHeightLt, v)
		case "outer_height__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_height__lte' takes an integer, got %q.", value))
				return
			}
			vOuterHeightLte = append(vOuterHeightLte, v)
		case "outer_height__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_height__gt' takes an integer, got %q.", value))
				return
			}
			vOuterHeightGt = append(vOuterHeightGt, v)
		case "outer_height__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_height__gte' takes an integer, got %q.", value))
				return
			}
			vOuterHeightGte = append(vOuterHeightGte, v)
		case "outer_height__empty":
			if vOuterHeightEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'outer_height__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_height__empty' takes a boolean, got %q.", value))
				return
			}
			vOuterHeightEmpty = &v
		case "outer_unit":
			if vOuterUnit != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'outer_unit' takes a single value.")
				return
			}
			v := value
			vOuterUnit = &v
		case "outer_unit__n":
			if vOuterUnitn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'outer_unit__n' takes a single value.")
				return
			}
			v := value
			vOuterUnitn = &v
		case "outer_unit__ic":
			v := value
			vOuterUnitIc = append(vOuterUnitIc, v)
		case "outer_unit__nic":
			v := value
			vOuterUnitNic = append(vOuterUnitNic, v)
		case "outer_unit__ie":
			v := value
			vOuterUnitIe = append(vOuterUnitIe, v)
		case "outer_unit__nie":
			v := value
			vOuterUnitNie = append(vOuterUnitNie, v)
		case "outer_unit__isw":
			v := value
			vOuterUnitIsw = append(vOuterUnitIsw, v)
		case "outer_unit__nisw":
			v := value
			vOuterUnitNisw = append(vOuterUnitNisw, v)
		case "outer_unit__iew":
			v := value
			vOuterUnitIew = append(vOuterUnitIew, v)
		case "outer_unit__niew":
			v := value
			vOuterUnitNiew = append(vOuterUnitNiew, v)
		case "outer_unit__empty":
			if vOuterUnitEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'outer_unit__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_unit__empty' takes a boolean, got %q.", value))
				return
			}
			vOuterUnitEmpty = &v
		case "outer_unit__regex":
			v := value
			vOuterUnitRegex = append(vOuterUnitRegex, v)
		case "outer_unit__iregex":
			v := value
			vOuterUnitIregex = append(vOuterUnitIregex, v)
		case "outer_width":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_width' takes an integer, got %q.", value))
				return
			}
			vOuterWidth = append(vOuterWidth, v)
		case "outer_width__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_width__n' takes an integer, got %q.", value))
				return
			}
			vOuterWidthn = append(vOuterWidthn, v)
		case "outer_width__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_width__lt' takes an integer, got %q.", value))
				return
			}
			vOuterWidthLt = append(vOuterWidthLt, v)
		case "outer_width__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_width__lte' takes an integer, got %q.", value))
				return
			}
			vOuterWidthLte = append(vOuterWidthLte, v)
		case "outer_width__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_width__gt' takes an integer, got %q.", value))
				return
			}
			vOuterWidthGt = append(vOuterWidthGt, v)
		case "outer_width__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_width__gte' takes an integer, got %q.", value))
				return
			}
			vOuterWidthGte = append(vOuterWidthGte, v)
		case "outer_width__empty":
			if vOuterWidthEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'outer_width__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'outer_width__empty' takes a boolean, got %q.", value))
				return
			}
			vOuterWidthEmpty = &v
		case "owner":
			v := value
			vOwner = append(vOwner, v)
		case "owner__n":
			v := value
			vOwnern = append(vOwnern, v)
		case "owner_group":
			v := value
			vOwnerGroup = append(vOwnerGroup, v)
		case "owner_group__n":
			v := value
			vOwnerGroupn = append(vOwnerGroupn, v)
		case "owner_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupID = append(vOwnerGroupID, v)
		case "owner_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupIDn = append(vOwnerGroupIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "rack_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_count' takes an integer, got %q.", value))
				return
			}
			vRackCount = append(vRackCount, v)
		case "rack_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_count__n' takes an integer, got %q.", value))
				return
			}
			vRackCountn = append(vRackCountn, v)
		case "rack_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_count__lt' takes an integer, got %q.", value))
				return
			}
			vRackCountLt = append(vRackCountLt, v)
		case "rack_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_count__lte' takes an integer, got %q.", value))
				return
			}
			vRackCountLte = append(vRackCountLte, v)
		case "rack_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_count__gt' takes an integer, got %q.", value))
				return
			}
			vRackCountGt = append(vRackCountGt, v)
		case "rack_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_count__gte' takes an integer, got %q.", value))
				return
			}
			vRackCountGte = append(vRackCountGte, v)
		case "rack_count__empty":
			if vRackCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'rack_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_count__empty' takes a boolean, got %q.", value))
				return
			}
			vRackCountEmpty = &v
		case "starting_unit":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'starting_unit' takes an integer, got %q.", value))
				return
			}
			vStartingUnit = append(vStartingUnit, v)
		case "starting_unit__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'starting_unit__n' takes an integer, got %q.", value))
				return
			}
			vStartingUnitn = append(vStartingUnitn, v)
		case "starting_unit__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'starting_unit__lt' takes an integer, got %q.", value))
				return
			}
			vStartingUnitLt = append(vStartingUnitLt, v)
		case "starting_unit__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'starting_unit__lte' takes an integer, got %q.", value))
				return
			}
			vStartingUnitLte = append(vStartingUnitLte, v)
		case "starting_unit__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'starting_unit__gt' takes an integer, got %q.", value))
				return
			}
			vStartingUnitGt = append(vStartingUnitGt, v)
		case "starting_unit__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'starting_unit__gte' takes an integer, got %q.", value))
				return
			}
			vStartingUnitGte = append(vStartingUnitGte, v)
		case "starting_unit__empty":
			if vStartingUnitEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'starting_unit__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'starting_unit__empty' takes a boolean, got %q.", value))
				return
			}
			vStartingUnitEmpty = &v
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
		case "u_height":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'u_height' takes an integer, got %q.", value))
				return
			}
			vUHeight = append(vUHeight, v)
		case "u_height__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'u_height__n' takes an integer, got %q.", value))
				return
			}
			vUHeightn = append(vUHeightn, v)
		case "u_height__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'u_height__lt' takes an integer, got %q.", value))
				return
			}
			vUHeightLt = append(vUHeightLt, v)
		case "u_height__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'u_height__lte' takes an integer, got %q.", value))
				return
			}
			vUHeightLte = append(vUHeightLte, v)
		case "u_height__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'u_height__gt' takes an integer, got %q.", value))
				return
			}
			vUHeightGt = append(vUHeightGt, v)
		case "u_height__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'u_height__gte' takes an integer, got %q.", value))
				return
			}
			vUHeightGte = append(vUHeightGte, v)
		case "u_height__empty":
			if vUHeightEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'u_height__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'u_height__empty' takes a boolean, got %q.", value))
				return
			}
			vUHeightEmpty = &v
		case "weight":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight' takes a number, got %q.", value))
				return
			}
			vWeight = append(vWeight, v)
		case "weight__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__n' takes a number, got %q.", value))
				return
			}
			vWeightn = append(vWeightn, v)
		case "weight__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__lt' takes a number, got %q.", value))
				return
			}
			vWeightLt = append(vWeightLt, v)
		case "weight__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__lte' takes a number, got %q.", value))
				return
			}
			vWeightLte = append(vWeightLte, v)
		case "weight__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__gt' takes a number, got %q.", value))
				return
			}
			vWeightGt = append(vWeightGt, v)
		case "weight__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__gte' takes a number, got %q.", value))
				return
			}
			vWeightGte = append(vWeightGte, v)
		case "weight__empty":
			if vWeightEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'weight__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight__empty' takes a boolean, got %q.", value))
				return
			}
			vWeightEmpty = &v
		case "weight_unit":
			if vWeightUnit != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'weight_unit' takes a single value.")
				return
			}
			v := value
			vWeightUnit = &v
		case "weight_unit__n":
			if vWeightUnitn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'weight_unit__n' takes a single value.")
				return
			}
			v := value
			vWeightUnitn = &v
		case "weight_unit__ic":
			v := value
			vWeightUnitIc = append(vWeightUnitIc, v)
		case "weight_unit__nic":
			v := value
			vWeightUnitNic = append(vWeightUnitNic, v)
		case "weight_unit__ie":
			v := value
			vWeightUnitIe = append(vWeightUnitIe, v)
		case "weight_unit__nie":
			v := value
			vWeightUnitNie = append(vWeightUnitNie, v)
		case "weight_unit__isw":
			v := value
			vWeightUnitIsw = append(vWeightUnitIsw, v)
		case "weight_unit__nisw":
			v := value
			vWeightUnitNisw = append(vWeightUnitNisw, v)
		case "weight_unit__iew":
			v := value
			vWeightUnitIew = append(vWeightUnitIew, v)
		case "weight_unit__niew":
			v := value
			vWeightUnitNiew = append(vWeightUnitNiew, v)
		case "weight_unit__empty":
			if vWeightUnitEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'weight_unit__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'weight_unit__empty' takes a boolean, got %q.", value))
				return
			}
			vWeightUnitEmpty = &v
		case "weight_unit__regex":
			v := value
			vWeightUnitRegex = append(vWeightUnitRegex, v)
		case "weight_unit__iregex":
			v := value
			vWeightUnitIregex = append(vWeightUnitIregex, v)
		case "width":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width' takes an integer, got %q.", value))
				return
			}
			vWidth = append(vWidth, v)
		case "width__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__n' takes an integer, got %q.", value))
				return
			}
			vWidthn = append(vWidthn, v)
		case "width__ic":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__ic' takes an integer, got %q.", value))
				return
			}
			vWidthIc = append(vWidthIc, v)
		case "width__nic":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__nic' takes an integer, got %q.", value))
				return
			}
			vWidthNic = append(vWidthNic, v)
		case "width__ie":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__ie' takes an integer, got %q.", value))
				return
			}
			vWidthIe = append(vWidthIe, v)
		case "width__nie":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__nie' takes an integer, got %q.", value))
				return
			}
			vWidthNie = append(vWidthNie, v)
		case "width__isw":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__isw' takes an integer, got %q.", value))
				return
			}
			vWidthIsw = append(vWidthIsw, v)
		case "width__nisw":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__nisw' takes an integer, got %q.", value))
				return
			}
			vWidthNisw = append(vWidthNisw, v)
		case "width__iew":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__iew' takes an integer, got %q.", value))
				return
			}
			vWidthIew = append(vWidthIew, v)
		case "width__niew":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__niew' takes an integer, got %q.", value))
				return
			}
			vWidthNiew = append(vWidthNiew, v)
		case "width__regex":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__regex' takes an integer, got %q.", value))
				return
			}
			vWidthRegex = append(vWidthRegex, v)
		case "width__iregex":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'width__iregex' takes an integer, got %q.", value))
				return
			}
			vWidthIregex = append(vWidthIregex, v)
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
	if len(vModel) > 0 {
		params.SetModel(vModel)
	}
	if len(vModeln) > 0 {
		params.SetModeln(vModeln)
	}
	if len(vModelIc) > 0 {
		params.SetModelIc(vModelIc)
	}
	if len(vModelNic) > 0 {
		params.SetModelNic(vModelNic)
	}
	if len(vModelIe) > 0 {
		params.SetModelIe(vModelIe)
	}
	if len(vModelNie) > 0 {
		params.SetModelNie(vModelNie)
	}
	if len(vModelIsw) > 0 {
		params.SetModelIsw(vModelIsw)
	}
	if len(vModelNisw) > 0 {
		params.SetModelNisw(vModelNisw)
	}
	if len(vModelIew) > 0 {
		params.SetModelIew(vModelIew)
	}
	if len(vModelNiew) > 0 {
		params.SetModelNiew(vModelNiew)
	}
	if vModelEmpty != nil {
		params.SetModelEmpty(vModelEmpty)
	}
	if len(vModelRegex) > 0 {
		params.SetModelRegex(vModelRegex)
	}
	if len(vModelIregex) > 0 {
		params.SetModelIregex(vModelIregex)
	}
	if len(vSlug) > 0 {
		params.SetSlug(vSlug)
	}
	if len(vSlugn) > 0 {
		params.SetSlugn(vSlugn)
	}
	if len(vSlugIc) > 0 {
		params.SetSlugIc(vSlugIc)
	}
	if len(vSlugNic) > 0 {
		params.SetSlugNic(vSlugNic)
	}
	if len(vSlugIe) > 0 {
		params.SetSlugIe(vSlugIe)
	}
	if len(vSlugNie) > 0 {
		params.SetSlugNie(vSlugNie)
	}
	if len(vSlugIsw) > 0 {
		params.SetSlugIsw(vSlugIsw)
	}
	if len(vSlugNisw) > 0 {
		params.SetSlugNisw(vSlugNisw)
	}
	if len(vSlugIew) > 0 {
		params.SetSlugIew(vSlugIew)
	}
	if len(vSlugNiew) > 0 {
		params.SetSlugNiew(vSlugNiew)
	}
	if vSlugEmpty != nil {
		params.SetSlugEmpty(vSlugEmpty)
	}
	if len(vSlugRegex) > 0 {
		params.SetSlugRegex(vSlugRegex)
	}
	if len(vSlugIregex) > 0 {
		params.SetSlugIregex(vSlugIregex)
	}
	if len(vManufacturerID) > 0 {
		params.SetManufacturerID(vManufacturerID)
	}
	if len(vManufacturerIDn) > 0 {
		params.SetManufacturerIDn(vManufacturerIDn)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
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
	if vDescUnits != nil {
		params.SetDescUnits(vDescUnits)
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
	if len(vFormFactor) > 0 {
		params.SetFormFactor(vFormFactor)
	}
	if len(vFormFactorn) > 0 {
		params.SetFormFactorn(vFormFactorn)
	}
	if len(vFormFactorIc) > 0 {
		params.SetFormFactorIc(vFormFactorIc)
	}
	if len(vFormFactorNic) > 0 {
		params.SetFormFactorNic(vFormFactorNic)
	}
	if len(vFormFactorIe) > 0 {
		params.SetFormFactorIe(vFormFactorIe)
	}
	if len(vFormFactorNie) > 0 {
		params.SetFormFactorNie(vFormFactorNie)
	}
	if len(vFormFactorIsw) > 0 {
		params.SetFormFactorIsw(vFormFactorIsw)
	}
	if len(vFormFactorNisw) > 0 {
		params.SetFormFactorNisw(vFormFactorNisw)
	}
	if len(vFormFactorIew) > 0 {
		params.SetFormFactorIew(vFormFactorIew)
	}
	if len(vFormFactorNiew) > 0 {
		params.SetFormFactorNiew(vFormFactorNiew)
	}
	if vFormFactorEmpty != nil {
		params.SetFormFactorEmpty(vFormFactorEmpty)
	}
	if len(vFormFactorRegex) > 0 {
		params.SetFormFactorRegex(vFormFactorRegex)
	}
	if len(vFormFactorIregex) > 0 {
		params.SetFormFactorIregex(vFormFactorIregex)
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
	if len(vManufacturer) > 0 {
		params.SetManufacturer(vManufacturer)
	}
	if len(vManufacturern) > 0 {
		params.SetManufacturern(vManufacturern)
	}
	if len(vMaxWeight) > 0 {
		params.SetMaxWeight(vMaxWeight)
	}
	if len(vMaxWeightn) > 0 {
		params.SetMaxWeightn(vMaxWeightn)
	}
	if len(vMaxWeightLt) > 0 {
		params.SetMaxWeightLt(vMaxWeightLt)
	}
	if len(vMaxWeightLte) > 0 {
		params.SetMaxWeightLte(vMaxWeightLte)
	}
	if len(vMaxWeightGt) > 0 {
		params.SetMaxWeightGt(vMaxWeightGt)
	}
	if len(vMaxWeightGte) > 0 {
		params.SetMaxWeightGte(vMaxWeightGte)
	}
	if vMaxWeightEmpty != nil {
		params.SetMaxWeightEmpty(vMaxWeightEmpty)
	}
	if len(vMountingDepth) > 0 {
		params.SetMountingDepth(vMountingDepth)
	}
	if len(vMountingDepthn) > 0 {
		params.SetMountingDepthn(vMountingDepthn)
	}
	if len(vMountingDepthLt) > 0 {
		params.SetMountingDepthLt(vMountingDepthLt)
	}
	if len(vMountingDepthLte) > 0 {
		params.SetMountingDepthLte(vMountingDepthLte)
	}
	if len(vMountingDepthGt) > 0 {
		params.SetMountingDepthGt(vMountingDepthGt)
	}
	if len(vMountingDepthGte) > 0 {
		params.SetMountingDepthGte(vMountingDepthGte)
	}
	if vMountingDepthEmpty != nil {
		params.SetMountingDepthEmpty(vMountingDepthEmpty)
	}
	if len(vOuterDepth) > 0 {
		params.SetOuterDepth(vOuterDepth)
	}
	if len(vOuterDepthn) > 0 {
		params.SetOuterDepthn(vOuterDepthn)
	}
	if len(vOuterDepthLt) > 0 {
		params.SetOuterDepthLt(vOuterDepthLt)
	}
	if len(vOuterDepthLte) > 0 {
		params.SetOuterDepthLte(vOuterDepthLte)
	}
	if len(vOuterDepthGt) > 0 {
		params.SetOuterDepthGt(vOuterDepthGt)
	}
	if len(vOuterDepthGte) > 0 {
		params.SetOuterDepthGte(vOuterDepthGte)
	}
	if vOuterDepthEmpty != nil {
		params.SetOuterDepthEmpty(vOuterDepthEmpty)
	}
	if len(vOuterHeight) > 0 {
		params.SetOuterHeight(vOuterHeight)
	}
	if len(vOuterHeightn) > 0 {
		params.SetOuterHeightn(vOuterHeightn)
	}
	if len(vOuterHeightLt) > 0 {
		params.SetOuterHeightLt(vOuterHeightLt)
	}
	if len(vOuterHeightLte) > 0 {
		params.SetOuterHeightLte(vOuterHeightLte)
	}
	if len(vOuterHeightGt) > 0 {
		params.SetOuterHeightGt(vOuterHeightGt)
	}
	if len(vOuterHeightGte) > 0 {
		params.SetOuterHeightGte(vOuterHeightGte)
	}
	if vOuterHeightEmpty != nil {
		params.SetOuterHeightEmpty(vOuterHeightEmpty)
	}
	if vOuterUnit != nil {
		params.SetOuterUnit(vOuterUnit)
	}
	if vOuterUnitn != nil {
		params.SetOuterUnitn(vOuterUnitn)
	}
	if len(vOuterUnitIc) > 0 {
		params.SetOuterUnitIc(vOuterUnitIc)
	}
	if len(vOuterUnitNic) > 0 {
		params.SetOuterUnitNic(vOuterUnitNic)
	}
	if len(vOuterUnitIe) > 0 {
		params.SetOuterUnitIe(vOuterUnitIe)
	}
	if len(vOuterUnitNie) > 0 {
		params.SetOuterUnitNie(vOuterUnitNie)
	}
	if len(vOuterUnitIsw) > 0 {
		params.SetOuterUnitIsw(vOuterUnitIsw)
	}
	if len(vOuterUnitNisw) > 0 {
		params.SetOuterUnitNisw(vOuterUnitNisw)
	}
	if len(vOuterUnitIew) > 0 {
		params.SetOuterUnitIew(vOuterUnitIew)
	}
	if len(vOuterUnitNiew) > 0 {
		params.SetOuterUnitNiew(vOuterUnitNiew)
	}
	if vOuterUnitEmpty != nil {
		params.SetOuterUnitEmpty(vOuterUnitEmpty)
	}
	if len(vOuterUnitRegex) > 0 {
		params.SetOuterUnitRegex(vOuterUnitRegex)
	}
	if len(vOuterUnitIregex) > 0 {
		params.SetOuterUnitIregex(vOuterUnitIregex)
	}
	if len(vOuterWidth) > 0 {
		params.SetOuterWidth(vOuterWidth)
	}
	if len(vOuterWidthn) > 0 {
		params.SetOuterWidthn(vOuterWidthn)
	}
	if len(vOuterWidthLt) > 0 {
		params.SetOuterWidthLt(vOuterWidthLt)
	}
	if len(vOuterWidthLte) > 0 {
		params.SetOuterWidthLte(vOuterWidthLte)
	}
	if len(vOuterWidthGt) > 0 {
		params.SetOuterWidthGt(vOuterWidthGt)
	}
	if len(vOuterWidthGte) > 0 {
		params.SetOuterWidthGte(vOuterWidthGte)
	}
	if vOuterWidthEmpty != nil {
		params.SetOuterWidthEmpty(vOuterWidthEmpty)
	}
	if len(vOwner) > 0 {
		params.SetOwner(vOwner)
	}
	if len(vOwnern) > 0 {
		params.SetOwnern(vOwnern)
	}
	if len(vOwnerGroup) > 0 {
		params.SetOwnerGroup(vOwnerGroup)
	}
	if len(vOwnerGroupn) > 0 {
		params.SetOwnerGroupn(vOwnerGroupn)
	}
	if len(vOwnerGroupID) > 0 {
		params.SetOwnerGroupID(vOwnerGroupID)
	}
	if len(vOwnerGroupIDn) > 0 {
		params.SetOwnerGroupIDn(vOwnerGroupIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRackCount) > 0 {
		params.SetRackCount(vRackCount)
	}
	if len(vRackCountn) > 0 {
		params.SetRackCountn(vRackCountn)
	}
	if len(vRackCountLt) > 0 {
		params.SetRackCountLt(vRackCountLt)
	}
	if len(vRackCountLte) > 0 {
		params.SetRackCountLte(vRackCountLte)
	}
	if len(vRackCountGt) > 0 {
		params.SetRackCountGt(vRackCountGt)
	}
	if len(vRackCountGte) > 0 {
		params.SetRackCountGte(vRackCountGte)
	}
	if vRackCountEmpty != nil {
		params.SetRackCountEmpty(vRackCountEmpty)
	}
	if len(vStartingUnit) > 0 {
		params.SetStartingUnit(vStartingUnit)
	}
	if len(vStartingUnitn) > 0 {
		params.SetStartingUnitn(vStartingUnitn)
	}
	if len(vStartingUnitLt) > 0 {
		params.SetStartingUnitLt(vStartingUnitLt)
	}
	if len(vStartingUnitLte) > 0 {
		params.SetStartingUnitLte(vStartingUnitLte)
	}
	if len(vStartingUnitGt) > 0 {
		params.SetStartingUnitGt(vStartingUnitGt)
	}
	if len(vStartingUnitGte) > 0 {
		params.SetStartingUnitGte(vStartingUnitGte)
	}
	if vStartingUnitEmpty != nil {
		params.SetStartingUnitEmpty(vStartingUnitEmpty)
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
	if len(vUHeight) > 0 {
		params.SetUHeight(vUHeight)
	}
	if len(vUHeightn) > 0 {
		params.SetUHeightn(vUHeightn)
	}
	if len(vUHeightLt) > 0 {
		params.SetUHeightLt(vUHeightLt)
	}
	if len(vUHeightLte) > 0 {
		params.SetUHeightLte(vUHeightLte)
	}
	if len(vUHeightGt) > 0 {
		params.SetUHeightGt(vUHeightGt)
	}
	if len(vUHeightGte) > 0 {
		params.SetUHeightGte(vUHeightGte)
	}
	if vUHeightEmpty != nil {
		params.SetUHeightEmpty(vUHeightEmpty)
	}
	if len(vWeight) > 0 {
		params.SetWeight(vWeight)
	}
	if len(vWeightn) > 0 {
		params.SetWeightn(vWeightn)
	}
	if len(vWeightLt) > 0 {
		params.SetWeightLt(vWeightLt)
	}
	if len(vWeightLte) > 0 {
		params.SetWeightLte(vWeightLte)
	}
	if len(vWeightGt) > 0 {
		params.SetWeightGt(vWeightGt)
	}
	if len(vWeightGte) > 0 {
		params.SetWeightGte(vWeightGte)
	}
	if vWeightEmpty != nil {
		params.SetWeightEmpty(vWeightEmpty)
	}
	if vWeightUnit != nil {
		params.SetWeightUnit(vWeightUnit)
	}
	if vWeightUnitn != nil {
		params.SetWeightUnitn(vWeightUnitn)
	}
	if len(vWeightUnitIc) > 0 {
		params.SetWeightUnitIc(vWeightUnitIc)
	}
	if len(vWeightUnitNic) > 0 {
		params.SetWeightUnitNic(vWeightUnitNic)
	}
	if len(vWeightUnitIe) > 0 {
		params.SetWeightUnitIe(vWeightUnitIe)
	}
	if len(vWeightUnitNie) > 0 {
		params.SetWeightUnitNie(vWeightUnitNie)
	}
	if len(vWeightUnitIsw) > 0 {
		params.SetWeightUnitIsw(vWeightUnitIsw)
	}
	if len(vWeightUnitNisw) > 0 {
		params.SetWeightUnitNisw(vWeightUnitNisw)
	}
	if len(vWeightUnitIew) > 0 {
		params.SetWeightUnitIew(vWeightUnitIew)
	}
	if len(vWeightUnitNiew) > 0 {
		params.SetWeightUnitNiew(vWeightUnitNiew)
	}
	if vWeightUnitEmpty != nil {
		params.SetWeightUnitEmpty(vWeightUnitEmpty)
	}
	if len(vWeightUnitRegex) > 0 {
		params.SetWeightUnitRegex(vWeightUnitRegex)
	}
	if len(vWeightUnitIregex) > 0 {
		params.SetWeightUnitIregex(vWeightUnitIregex)
	}
	if len(vWidth) > 0 {
		params.SetWidth(vWidth)
	}
	if len(vWidthn) > 0 {
		params.SetWidthn(vWidthn)
	}
	if len(vWidthIc) > 0 {
		params.SetWidthIc(vWidthIc)
	}
	if len(vWidthNic) > 0 {
		params.SetWidthNic(vWidthNic)
	}
	if len(vWidthIe) > 0 {
		params.SetWidthIe(vWidthIe)
	}
	if len(vWidthNie) > 0 {
		params.SetWidthNie(vWidthNie)
	}
	if len(vWidthIsw) > 0 {
		params.SetWidthIsw(vWidthIsw)
	}
	if len(vWidthNisw) > 0 {
		params.SetWidthNisw(vWidthNisw)
	}
	if len(vWidthIew) > 0 {
		params.SetWidthIew(vWidthIew)
	}
	if len(vWidthNiew) > 0 {
		params.SetWidthNiew(vWidthNiew)
	}
	if len(vWidthRegex) > 0 {
		params.SetWidthRegex(vWidthRegex)
	}
	if len(vWidthIregex) > 0 {
		params.SetWidthIregex(vWidthIregex)
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
