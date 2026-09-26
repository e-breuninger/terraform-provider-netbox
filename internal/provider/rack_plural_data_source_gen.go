// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*racksDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*racksDataSource)(nil)
)

// NewRacksDataSource returns a new racks data source, which
// lists rack objects matching its filters.
func NewRacksDataSource() datasource.DataSource {
	return &racksDataSource{}
}

type racksDataSource struct {
	client *netboxapi.Client
}

// racksDataSourceModel is the data source model: the filters, the limit and the matching
// racks.
type racksDataSourceModel struct {
	Filters   types.Set    `tfsdk:"filters"`
	NameRegex types.String `tfsdk:"name_regex"`
	Limit     types.Int64  `tfsdk:"limit"`
	Items     types.List   `tfsdk:"racks"`
}

// rackResourceAttrTypes is the attribute type map of rackResourceModel.
func rackResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":               types.Int64Type,
		"name":             types.StringType,
		"site_id":          types.Int64Type,
		"location_id":      types.Int64Type,
		"tenant_id":        types.Int64Type,
		"role_id":          types.Int64Type,
		"rack_type_id":     types.Int64Type,
		"status":           types.StringType,
		"form_factor":      types.StringType,
		"width":            types.Int64Type,
		"u_height":         types.Int64Type,
		"desc_units":       types.BoolType,
		"serial":           types.StringType,
		"asset_tag":        types.StringType,
		"facility_id":      types.StringType,
		"outer_width":      types.Int64Type,
		"outer_depth":      types.Int64Type,
		"outer_unit":       types.StringType,
		"mounting_depth":   types.Int64Type,
		"weight":           types.Float64Type,
		"max_weight":       types.Int64Type,
		"weight_unit":      types.StringType,
		"description":      types.StringType,
		"comments":         types.StringType,
		"owner_id":         types.Int64Type,
		"created":          types.StringType,
		"last_updated":     types.StringType,
		"url":              types.StringType,
		"device_count":     types.Int64Type,
		"power_feed_count": types.Int64Type,
		"tags":             types.SetType{ElemType: types.StringType},
		"tags_all":         types.SetType{ElemType: types.StringType},
		"custom_fields":    types.MapType{ElemType: types.StringType},
	}
}

func (d *racksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_racks"
}

func (d *racksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists rack objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: asset_tag, asset_tag__empty, asset_tag__ic, asset_tag__ie, asset_tag__iew, asset_tag__iregex, asset_tag__isw, asset_tag__n, asset_tag__nic, asset_tag__nie, asset_tag__niew, asset_tag__nisw, asset_tag__regex, contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, desc_units, facility_id, facility_id__empty, facility_id__ic, facility_id__ie, facility_id__iew, facility_id__iregex, facility_id__isw, facility_id__n, facility_id__nic, facility_id__nie, facility_id__niew, facility_id__nisw, facility_id__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, location_id, location_id__n, max_weight, max_weight__empty, max_weight__gt, max_weight__gte, max_weight__lt, max_weight__lte, max_weight__n, mounting_depth, mounting_depth__empty, mounting_depth__gt, mounting_depth__gte, mounting_depth__lt, mounting_depth__lte, mounting_depth__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, outer_depth, outer_depth__empty, outer_depth__gt, outer_depth__gte, outer_depth__lt, outer_depth__lte, outer_depth__n, outer_unit, outer_unit__empty, outer_unit__ic, outer_unit__ie, outer_unit__iew, outer_unit__iregex, outer_unit__isw, outer_unit__n, outer_unit__nic, outer_unit__nie, outer_unit__niew, outer_unit__nisw, outer_unit__regex, outer_width, outer_width__empty, outer_width__gt, outer_width__gte, outer_width__lt, outer_width__lte, outer_width__n, owner_id, owner_id__n, rack_type_id, rack_type_id__n, region_id, region_id__n, role_id, role_id__n, serial, serial__empty, serial__ic, serial__ie, serial__iew, serial__iregex, serial__isw, serial__n, serial__nic, serial__nie, serial__niew, serial__nisw, serial__regex, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tenant_id, tenant_id__n, u_height, u_height__empty, u_height__gt, u_height__gte, u_height__lt, u_height__lte, u_height__n, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n, weight_unit, weight_unit__empty, weight_unit__ic, weight_unit__ie, weight_unit__iew, weight_unit__iregex, weight_unit__isw, weight_unit__n, weight_unit__nic, weight_unit__nie, weight_unit__niew, weight_unit__nisw, weight_unit__regex, width, width__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"name_regex": schema.StringAttribute{
				Optional:    true,
				Description: "Go regular expression the name must match. Applied after the API query, which then fetches every object matching the filters; limit applies to the matches.",
				Validators: []validator.String{
					conv.ValidRegexp(),
				},
			},
			"limit": schema.Int64Attribute{
				Optional:    true,
				Description: "The maximum number of objects to return from the API lookup. Defaults to 1000.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"racks": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching racks.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site the rack belongs to.",
						},
						"location_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the location within the site.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the rack role.",
						},
						"rack_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the rack type. NetBox copies the type's form factor, width, height and outer dimensions onto the rack; set the same values here too, or the next plan shows them as drift.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: reserved, available, planned, active, deprecated.",
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
							Description: "Height in rack units.",
						},
						"desc_units": schema.BoolAttribute{
							Computed:    true,
							Description: "Units are numbered top-to-bottom.",
						},
						"serial": schema.StringAttribute{
							Computed:    true,
							Description: "Serial number.",
						},
						"asset_tag": schema.StringAttribute{
							Computed:    true,
							Description: "Unique asset tag.",
						},
						"facility_id": schema.StringAttribute{
							Computed:    true,
							Description: "Locally-assigned facility identifier.",
						},
						"outer_width": schema.Int64Attribute{
							Computed:    true,
							Description: "Outer width (requires outer_unit).",
						},
						"outer_depth": schema.Int64Attribute{
							Computed:    true,
							Description: "Outer depth (requires outer_unit).",
						},
						"outer_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of the outer dimensions. One of: mm, in.",
						},
						"mounting_depth": schema.Int64Attribute{
							Computed:    true,
							Description: "Maximum depth of a mounted device in millimeters.",
						},
						"weight": schema.Float64Attribute{
							Computed:    true,
							Description: "Weight of the rack (requires weight_unit).",
						},
						"max_weight": schema.Int64Attribute{
							Computed:    true,
							Description: "Maximum load capacity (requires weight_unit).",
						},
						"weight_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of weight and max_weight. One of: kg, g, lb, oz.",
						},
						"description": schema.StringAttribute{
							Computed: true,
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
					},
				},
			},
		},
	}
}

func (d *racksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *racksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data racksDataSourceModel

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

	// name_regex is applied here, not by the API: the fragment fetches every match (fetchAll),
	// the names are matched below, and only then does limit apply.
	var nameRegex *regexp.Regexp
	fetchAll := false
	if !state.NameRegex.IsNull() && !state.NameRegex.IsUnknown() {
		re, err := regexp.Compile(state.NameRegex.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("name_regex"), "Invalid regular expression", err.Error())
			return
		}
		nameRegex, fetchAll = re, true
	}
	_ = fetchAll
	var responseDTOs []*netboxapi.RackResponseDTO

	params := dcim.NewDcimRacksListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vName []string
	var vNamen []string
	var vNameNic []string
	var vNameIe []string
	var vNameNie []string
	var vNameIsw []string
	var vNameNisw []string
	var vNameIew []string
	var vNameNiew []string
	var vNameEmpty *bool
	var vNameRegex []string
	var vNameIregex []string
	var vNameIc []string
	var vSiteID []int64
	var vSiteIDn []int64
	var vLocationID []string
	var vLocationIDn []string
	var vTenantID []int64
	var vTenantIDn []int64
	var vRoleID []int64
	var vRoleIDn []int64
	var vRackTypeID []int64
	var vRackTypeIDn []int64
	var vStatus []string
	var vStatusn []string
	var vStatusIc []string
	var vStatusNic []string
	var vStatusIe []string
	var vStatusNie []string
	var vStatusIsw []string
	var vStatusNisw []string
	var vStatusIew []string
	var vStatusNiew []string
	var vStatusEmpty *bool
	var vStatusRegex []string
	var vStatusIregex []string
	var vWidth []int64
	var vWidthn []int64
	var vUHeight []int64
	var vUHeightn []int64
	var vUHeightLt []int64
	var vUHeightLte []int64
	var vUHeightGt []int64
	var vUHeightGte []int64
	var vUHeightEmpty *bool
	var vDescUnits *bool
	var vSerial []string
	var vSerialn []string
	var vSerialIc []string
	var vSerialNic []string
	var vSerialIe []string
	var vSerialNie []string
	var vSerialIsw []string
	var vSerialNisw []string
	var vSerialIew []string
	var vSerialNiew []string
	var vSerialEmpty *bool
	var vSerialRegex []string
	var vSerialIregex []string
	var vAssetTag []string
	var vAssetTagn []string
	var vAssetTagIc []string
	var vAssetTagNic []string
	var vAssetTagIe []string
	var vAssetTagNie []string
	var vAssetTagIsw []string
	var vAssetTagNisw []string
	var vAssetTagIew []string
	var vAssetTagNiew []string
	var vAssetTagEmpty *bool
	var vAssetTagRegex []string
	var vAssetTagIregex []string
	var vFacilityID []string
	var vFacilityIDn []string
	var vFacilityIDIc []string
	var vFacilityIDNic []string
	var vFacilityIDIe []string
	var vFacilityIDNie []string
	var vFacilityIDIsw []string
	var vFacilityIDNisw []string
	var vFacilityIDIew []string
	var vFacilityIDNiew []string
	var vFacilityIDEmpty *bool
	var vFacilityIDRegex []string
	var vFacilityIDIregex []string
	var vOuterWidth []int64
	var vOuterWidthn []int64
	var vOuterWidthLt []int64
	var vOuterWidthLte []int64
	var vOuterWidthGt []int64
	var vOuterWidthGte []int64
	var vOuterWidthEmpty *bool
	var vOuterDepth []int64
	var vOuterDepthn []int64
	var vOuterDepthLt []int64
	var vOuterDepthLte []int64
	var vOuterDepthGt []int64
	var vOuterDepthGte []int64
	var vOuterDepthEmpty *bool
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
	var vMountingDepth []int64
	var vMountingDepthn []int64
	var vMountingDepthLt []int64
	var vMountingDepthLte []int64
	var vMountingDepthGt []int64
	var vMountingDepthGte []int64
	var vMountingDepthEmpty *bool
	var vWeight []float64
	var vWeightn []float64
	var vWeightLt []float64
	var vWeightLte []float64
	var vWeightGt []float64
	var vWeightGte []float64
	var vWeightEmpty *bool
	var vMaxWeight []int64
	var vMaxWeightn []int64
	var vMaxWeightLt []int64
	var vMaxWeightLte []int64
	var vMaxWeightGt []int64
	var vMaxWeightGte []int64
	var vMaxWeightEmpty *bool
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
	var vRegionID []string
	var vRegionIDn []string
	var vContact []int64
	var vContactn []int64
	var vContactGroup []string
	var vContactGroupn []string
	var vContactRole []int64
	var vContactRolen []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
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
		case "name":
			v := value
			vName = append(vName, v)
		case "name__n":
			v := value
			vNamen = append(vNamen, v)
		case "name__nic":
			v := value
			vNameNic = append(vNameNic, v)
		case "name__ie":
			v := value
			vNameIe = append(vNameIe, v)
		case "name__nie":
			v := value
			vNameNie = append(vNameNie, v)
		case "name__isw":
			v := value
			vNameIsw = append(vNameIsw, v)
		case "name__nisw":
			v := value
			vNameNisw = append(vNameNisw, v)
		case "name__iew":
			v := value
			vNameIew = append(vNameIew, v)
		case "name__niew":
			v := value
			vNameNiew = append(vNameNiew, v)
		case "name__empty":
			if vNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'name__empty' takes a boolean, got %q.", value))
				return
			}
			vNameEmpty = &v
		case "name__regex":
			v := value
			vNameRegex = append(vNameRegex, v)
		case "name__iregex":
			v := value
			vNameIregex = append(vNameIregex, v)
		case "name__ic":
			v := value
			vNameIc = append(vNameIc, v)
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
		case "location_id":
			v := value
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v := value
			vLocationIDn = append(vLocationIDn, v)
		case "tenant_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id' takes an integer, got %q.", value))
				return
			}
			vTenantID = append(vTenantID, v)
		case "tenant_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id__n' takes an integer, got %q.", value))
				return
			}
			vTenantIDn = append(vTenantIDn, v)
		case "role_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id' takes an integer, got %q.", value))
				return
			}
			vRoleID = append(vRoleID, v)
		case "role_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id__n' takes an integer, got %q.", value))
				return
			}
			vRoleIDn = append(vRoleIDn, v)
		case "rack_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_type_id' takes an integer, got %q.", value))
				return
			}
			vRackTypeID = append(vRackTypeID, v)
		case "rack_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_type_id__n' takes an integer, got %q.", value))
				return
			}
			vRackTypeIDn = append(vRackTypeIDn, v)
		case "status":
			v := value
			vStatus = append(vStatus, v)
		case "status__n":
			v := value
			vStatusn = append(vStatusn, v)
		case "status__ic":
			v := value
			vStatusIc = append(vStatusIc, v)
		case "status__nic":
			v := value
			vStatusNic = append(vStatusNic, v)
		case "status__ie":
			v := value
			vStatusIe = append(vStatusIe, v)
		case "status__nie":
			v := value
			vStatusNie = append(vStatusNie, v)
		case "status__isw":
			v := value
			vStatusIsw = append(vStatusIsw, v)
		case "status__nisw":
			v := value
			vStatusNisw = append(vStatusNisw, v)
		case "status__iew":
			v := value
			vStatusIew = append(vStatusIew, v)
		case "status__niew":
			v := value
			vStatusNiew = append(vStatusNiew, v)
		case "status__empty":
			if vStatusEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'status__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'status__empty' takes a boolean, got %q.", value))
				return
			}
			vStatusEmpty = &v
		case "status__regex":
			v := value
			vStatusRegex = append(vStatusRegex, v)
		case "status__iregex":
			v := value
			vStatusIregex = append(vStatusIregex, v)
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
		case "serial":
			v := value
			vSerial = append(vSerial, v)
		case "serial__n":
			v := value
			vSerialn = append(vSerialn, v)
		case "serial__ic":
			v := value
			vSerialIc = append(vSerialIc, v)
		case "serial__nic":
			v := value
			vSerialNic = append(vSerialNic, v)
		case "serial__ie":
			v := value
			vSerialIe = append(vSerialIe, v)
		case "serial__nie":
			v := value
			vSerialNie = append(vSerialNie, v)
		case "serial__isw":
			v := value
			vSerialIsw = append(vSerialIsw, v)
		case "serial__nisw":
			v := value
			vSerialNisw = append(vSerialNisw, v)
		case "serial__iew":
			v := value
			vSerialIew = append(vSerialIew, v)
		case "serial__niew":
			v := value
			vSerialNiew = append(vSerialNiew, v)
		case "serial__empty":
			if vSerialEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'serial__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'serial__empty' takes a boolean, got %q.", value))
				return
			}
			vSerialEmpty = &v
		case "serial__regex":
			v := value
			vSerialRegex = append(vSerialRegex, v)
		case "serial__iregex":
			v := value
			vSerialIregex = append(vSerialIregex, v)
		case "asset_tag":
			v := value
			vAssetTag = append(vAssetTag, v)
		case "asset_tag__n":
			v := value
			vAssetTagn = append(vAssetTagn, v)
		case "asset_tag__ic":
			v := value
			vAssetTagIc = append(vAssetTagIc, v)
		case "asset_tag__nic":
			v := value
			vAssetTagNic = append(vAssetTagNic, v)
		case "asset_tag__ie":
			v := value
			vAssetTagIe = append(vAssetTagIe, v)
		case "asset_tag__nie":
			v := value
			vAssetTagNie = append(vAssetTagNie, v)
		case "asset_tag__isw":
			v := value
			vAssetTagIsw = append(vAssetTagIsw, v)
		case "asset_tag__nisw":
			v := value
			vAssetTagNisw = append(vAssetTagNisw, v)
		case "asset_tag__iew":
			v := value
			vAssetTagIew = append(vAssetTagIew, v)
		case "asset_tag__niew":
			v := value
			vAssetTagNiew = append(vAssetTagNiew, v)
		case "asset_tag__empty":
			if vAssetTagEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'asset_tag__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'asset_tag__empty' takes a boolean, got %q.", value))
				return
			}
			vAssetTagEmpty = &v
		case "asset_tag__regex":
			v := value
			vAssetTagRegex = append(vAssetTagRegex, v)
		case "asset_tag__iregex":
			v := value
			vAssetTagIregex = append(vAssetTagIregex, v)
		case "facility_id":
			v := value
			vFacilityID = append(vFacilityID, v)
		case "facility_id__n":
			v := value
			vFacilityIDn = append(vFacilityIDn, v)
		case "facility_id__ic":
			v := value
			vFacilityIDIc = append(vFacilityIDIc, v)
		case "facility_id__nic":
			v := value
			vFacilityIDNic = append(vFacilityIDNic, v)
		case "facility_id__ie":
			v := value
			vFacilityIDIe = append(vFacilityIDIe, v)
		case "facility_id__nie":
			v := value
			vFacilityIDNie = append(vFacilityIDNie, v)
		case "facility_id__isw":
			v := value
			vFacilityIDIsw = append(vFacilityIDIsw, v)
		case "facility_id__nisw":
			v := value
			vFacilityIDNisw = append(vFacilityIDNisw, v)
		case "facility_id__iew":
			v := value
			vFacilityIDIew = append(vFacilityIDIew, v)
		case "facility_id__niew":
			v := value
			vFacilityIDNiew = append(vFacilityIDNiew, v)
		case "facility_id__empty":
			if vFacilityIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'facility_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'facility_id__empty' takes a boolean, got %q.", value))
				return
			}
			vFacilityIDEmpty = &v
		case "facility_id__regex":
			v := value
			vFacilityIDRegex = append(vFacilityIDRegex, v)
		case "facility_id__iregex":
			v := value
			vFacilityIDIregex = append(vFacilityIDIregex, v)
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
		case "region_id":
			v := value
			vRegionID = append(vRegionID, v)
		case "region_id__n":
			v := value
			vRegionIDn = append(vRegionIDn, v)
		case "contact":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact' takes an integer, got %q.", value))
				return
			}
			vContact = append(vContact, v)
		case "contact__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact__n' takes an integer, got %q.", value))
				return
			}
			vContactn = append(vContactn, v)
		case "contact_group":
			v := value
			vContactGroup = append(vContactGroup, v)
		case "contact_group__n":
			v := value
			vContactGroupn = append(vContactGroupn, v)
		case "contact_role":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_role' takes an integer, got %q.", value))
				return
			}
			vContactRole = append(vContactRole, v)
		case "contact_role__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_role__n' takes an integer, got %q.", value))
				return
			}
			vContactRolen = append(vContactRolen, v)
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
	if len(vName) > 0 {
		params.SetName(vName)
	}
	if len(vNamen) > 0 {
		params.SetNamen(vNamen)
	}
	if len(vNameNic) > 0 {
		params.SetNameNic(vNameNic)
	}
	if len(vNameIe) > 0 {
		params.SetNameIe(vNameIe)
	}
	if len(vNameNie) > 0 {
		params.SetNameNie(vNameNie)
	}
	if len(vNameIsw) > 0 {
		params.SetNameIsw(vNameIsw)
	}
	if len(vNameNisw) > 0 {
		params.SetNameNisw(vNameNisw)
	}
	if len(vNameIew) > 0 {
		params.SetNameIew(vNameIew)
	}
	if len(vNameNiew) > 0 {
		params.SetNameNiew(vNameNiew)
	}
	if vNameEmpty != nil {
		params.SetNameEmpty(vNameEmpty)
	}
	if len(vNameRegex) > 0 {
		params.SetNameRegex(vNameRegex)
	}
	if len(vNameIregex) > 0 {
		params.SetNameIregex(vNameIregex)
	}
	if len(vNameIc) > 0 {
		params.SetNameIc(vNameIc)
	}
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
	}
	if len(vLocationID) > 0 {
		params.SetLocationID(vLocationID)
	}
	if len(vLocationIDn) > 0 {
		params.SetLocationIDn(vLocationIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vRoleID) > 0 {
		params.SetRoleID(vRoleID)
	}
	if len(vRoleIDn) > 0 {
		params.SetRoleIDn(vRoleIDn)
	}
	if len(vRackTypeID) > 0 {
		params.SetRackTypeID(vRackTypeID)
	}
	if len(vRackTypeIDn) > 0 {
		params.SetRackTypeIDn(vRackTypeIDn)
	}
	if len(vStatus) > 0 {
		params.SetStatus(vStatus)
	}
	if len(vStatusn) > 0 {
		params.SetStatusn(vStatusn)
	}
	if len(vStatusIc) > 0 {
		params.SetStatusIc(vStatusIc)
	}
	if len(vStatusNic) > 0 {
		params.SetStatusNic(vStatusNic)
	}
	if len(vStatusIe) > 0 {
		params.SetStatusIe(vStatusIe)
	}
	if len(vStatusNie) > 0 {
		params.SetStatusNie(vStatusNie)
	}
	if len(vStatusIsw) > 0 {
		params.SetStatusIsw(vStatusIsw)
	}
	if len(vStatusNisw) > 0 {
		params.SetStatusNisw(vStatusNisw)
	}
	if len(vStatusIew) > 0 {
		params.SetStatusIew(vStatusIew)
	}
	if len(vStatusNiew) > 0 {
		params.SetStatusNiew(vStatusNiew)
	}
	if vStatusEmpty != nil {
		params.SetStatusEmpty(vStatusEmpty)
	}
	if len(vStatusRegex) > 0 {
		params.SetStatusRegex(vStatusRegex)
	}
	if len(vStatusIregex) > 0 {
		params.SetStatusIregex(vStatusIregex)
	}
	if len(vWidth) > 0 {
		params.SetWidth(vWidth)
	}
	if len(vWidthn) > 0 {
		params.SetWidthn(vWidthn)
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
	if vDescUnits != nil {
		params.SetDescUnits(vDescUnits)
	}
	if len(vSerial) > 0 {
		params.SetSerial(vSerial)
	}
	if len(vSerialn) > 0 {
		params.SetSerialn(vSerialn)
	}
	if len(vSerialIc) > 0 {
		params.SetSerialIc(vSerialIc)
	}
	if len(vSerialNic) > 0 {
		params.SetSerialNic(vSerialNic)
	}
	if len(vSerialIe) > 0 {
		params.SetSerialIe(vSerialIe)
	}
	if len(vSerialNie) > 0 {
		params.SetSerialNie(vSerialNie)
	}
	if len(vSerialIsw) > 0 {
		params.SetSerialIsw(vSerialIsw)
	}
	if len(vSerialNisw) > 0 {
		params.SetSerialNisw(vSerialNisw)
	}
	if len(vSerialIew) > 0 {
		params.SetSerialIew(vSerialIew)
	}
	if len(vSerialNiew) > 0 {
		params.SetSerialNiew(vSerialNiew)
	}
	if vSerialEmpty != nil {
		params.SetSerialEmpty(vSerialEmpty)
	}
	if len(vSerialRegex) > 0 {
		params.SetSerialRegex(vSerialRegex)
	}
	if len(vSerialIregex) > 0 {
		params.SetSerialIregex(vSerialIregex)
	}
	if len(vAssetTag) > 0 {
		params.SetAssetTag(vAssetTag)
	}
	if len(vAssetTagn) > 0 {
		params.SetAssetTagn(vAssetTagn)
	}
	if len(vAssetTagIc) > 0 {
		params.SetAssetTagIc(vAssetTagIc)
	}
	if len(vAssetTagNic) > 0 {
		params.SetAssetTagNic(vAssetTagNic)
	}
	if len(vAssetTagIe) > 0 {
		params.SetAssetTagIe(vAssetTagIe)
	}
	if len(vAssetTagNie) > 0 {
		params.SetAssetTagNie(vAssetTagNie)
	}
	if len(vAssetTagIsw) > 0 {
		params.SetAssetTagIsw(vAssetTagIsw)
	}
	if len(vAssetTagNisw) > 0 {
		params.SetAssetTagNisw(vAssetTagNisw)
	}
	if len(vAssetTagIew) > 0 {
		params.SetAssetTagIew(vAssetTagIew)
	}
	if len(vAssetTagNiew) > 0 {
		params.SetAssetTagNiew(vAssetTagNiew)
	}
	if vAssetTagEmpty != nil {
		params.SetAssetTagEmpty(vAssetTagEmpty)
	}
	if len(vAssetTagRegex) > 0 {
		params.SetAssetTagRegex(vAssetTagRegex)
	}
	if len(vAssetTagIregex) > 0 {
		params.SetAssetTagIregex(vAssetTagIregex)
	}
	if len(vFacilityID) > 0 {
		params.SetFacilityID(vFacilityID)
	}
	if len(vFacilityIDn) > 0 {
		params.SetFacilityIDn(vFacilityIDn)
	}
	if len(vFacilityIDIc) > 0 {
		params.SetFacilityIDIc(vFacilityIDIc)
	}
	if len(vFacilityIDNic) > 0 {
		params.SetFacilityIDNic(vFacilityIDNic)
	}
	if len(vFacilityIDIe) > 0 {
		params.SetFacilityIDIe(vFacilityIDIe)
	}
	if len(vFacilityIDNie) > 0 {
		params.SetFacilityIDNie(vFacilityIDNie)
	}
	if len(vFacilityIDIsw) > 0 {
		params.SetFacilityIDIsw(vFacilityIDIsw)
	}
	if len(vFacilityIDNisw) > 0 {
		params.SetFacilityIDNisw(vFacilityIDNisw)
	}
	if len(vFacilityIDIew) > 0 {
		params.SetFacilityIDIew(vFacilityIDIew)
	}
	if len(vFacilityIDNiew) > 0 {
		params.SetFacilityIDNiew(vFacilityIDNiew)
	}
	if vFacilityIDEmpty != nil {
		params.SetFacilityIDEmpty(vFacilityIDEmpty)
	}
	if len(vFacilityIDRegex) > 0 {
		params.SetFacilityIDRegex(vFacilityIDRegex)
	}
	if len(vFacilityIDIregex) > 0 {
		params.SetFacilityIDIregex(vFacilityIDIregex)
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
	if len(vRegionID) > 0 {
		params.SetRegionID(vRegionID)
	}
	if len(vRegionIDn) > 0 {
		params.SetRegionIDn(vRegionIDn)
	}
	if len(vContact) > 0 {
		params.SetContact(vContact)
	}
	if len(vContactn) > 0 {
		params.SetContactn(vContactn)
	}
	if len(vContactGroup) > 0 {
		params.SetContactGroup(vContactGroup)
	}
	if len(vContactGroupn) > 0 {
		params.SetContactGroupn(vContactGroupn)
	}
	if len(vContactRole) > 0 {
		params.SetContactRole(vContactRole)
	}
	if len(vContactRolen) > 0 {
		params.SetContactRolen(vContactRolen)
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
	pageSize := limit
	if fetchAll {
		pageSize = 1000
	}
	params.SetLimit(&pageSize)
	for offset := int64(0); ; {
		res, err := d.client.Dcim.DcimRacksListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_racks", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.RackResponseDTOFromGoNetbox(goNetboxModel))
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
	items := make([]rackResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model rackResourceModel
		resp.Diagnostics.Append(flattenRack(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if nameRegex != nil {
		kept := items[:0]
		for _, item := range items {
			if nameRegex.MatchString(item.Name.ValueString()) {
				kept = append(kept, item)
			}
		}
		items = kept
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: rackResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
