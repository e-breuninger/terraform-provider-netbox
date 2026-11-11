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
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"racks"`
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
		"group_id":         types.Int64Type,
		"rack_type_id":     types.Int64Type,
		"status":           types.StringType,
		"form_factor":      types.StringType,
		"airflow":          types.StringType,
		"width":            types.Int64Type,
		"u_height":         types.Int64Type,
		"starting_unit":    types.Int64Type,
		"desc_units":       types.BoolType,
		"serial":           types.StringType,
		"asset_tag":        types.StringType,
		"facility_id":      types.StringType,
		"outer_width":      types.Int64Type,
		"outer_height":     types.Int64Type,
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
				Description: "Query filters, sent as API list parameters. Supported names: airflow, airflow__empty, airflow__ic, airflow__ie, airflow__iew, airflow__iregex, airflow__isw, airflow__n, airflow__nic, airflow__nie, airflow__niew, airflow__nisw, airflow__regex, asset_tag, asset_tag__empty, asset_tag__ic, asset_tag__ie, asset_tag__iew, asset_tag__iregex, asset_tag__isw, asset_tag__n, asset_tag__nic, asset_tag__nie, asset_tag__niew, asset_tag__nisw, asset_tag__regex, contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, desc_units, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, facility_id, facility_id__empty, facility_id__ic, facility_id__ie, facility_id__iew, facility_id__iregex, facility_id__isw, facility_id__n, facility_id__nic, facility_id__nie, facility_id__niew, facility_id__nisw, facility_id__regex, form_factor, form_factor__empty, form_factor__ic, form_factor__ie, form_factor__iew, form_factor__iregex, form_factor__isw, form_factor__n, form_factor__nic, form_factor__nie, form_factor__niew, form_factor__nisw, form_factor__regex, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, max_weight, max_weight__empty, max_weight__gt, max_weight__gte, max_weight__lt, max_weight__lte, max_weight__n, mounting_depth, mounting_depth__empty, mounting_depth__gt, mounting_depth__gte, mounting_depth__lt, mounting_depth__lte, mounting_depth__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, outer_depth, outer_depth__empty, outer_depth__gt, outer_depth__gte, outer_depth__lt, outer_depth__lte, outer_depth__n, outer_height, outer_height__empty, outer_height__gt, outer_height__gte, outer_height__lt, outer_height__lte, outer_height__n, outer_unit, outer_unit__empty, outer_unit__ic, outer_unit__ie, outer_unit__iew, outer_unit__iregex, outer_unit__isw, outer_unit__n, outer_unit__nic, outer_unit__nie, outer_unit__niew, outer_unit__nisw, outer_unit__regex, outer_width, outer_width__empty, outer_width__gt, outer_width__gte, outer_width__lt, outer_width__lte, outer_width__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, rack_type, rack_type__n, rack_type_id, rack_type_id__n, region, region__n, region_id, region_id__n, role, role__n, role_id, role_id__n, serial, serial__empty, serial__ic, serial__ie, serial__iew, serial__iregex, serial__isw, serial__n, serial__nic, serial__nie, serial__niew, serial__nisw, serial__regex, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, starting_unit, starting_unit__empty, starting_unit__gt, starting_unit__gte, starting_unit__lt, starting_unit__lte, starting_unit__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, u_height, u_height__empty, u_height__gt, u_height__gte, u_height__lt, u_height__lte, u_height__n, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n, weight_unit, weight_unit__empty, weight_unit__ic, weight_unit__ie, weight_unit__iew, weight_unit__iregex, weight_unit__isw, weight_unit__n, weight_unit__nic, weight_unit__nie, weight_unit__niew, weight_unit__nisw, weight_unit__regex, width, width__ic, width__ie, width__iew, width__iregex, width__isw, width__n, width__nic, width__nie, width__niew, width__nisw, width__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
						"group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the rack group.",
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
						"airflow": schema.StringAttribute{
							Computed:    true,
							Description: "Airflow direction. One of: front-to-rear, rear-to-front.",
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
							Description: "Maximum depth of a mounted device in millimeters. Must be a whole number.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
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
	var responseDTOs []*netboxapi.RackResponseDTO

	params := dcim.NewDcimRacksListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"name":                  netboxapi.FilterStrings(params.SetName),
		"name__n":               netboxapi.FilterStrings(params.SetNamen),
		"name__ic":              netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":             netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":              netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":             netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":             netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":            netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":             netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":            netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":           netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":           netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":          netboxapi.FilterStrings(params.SetNameIregex),
		"site_id":               netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":            netboxapi.FilterInt64s(params.SetSiteIDn),
		"location_id":           netboxapi.FilterStrings(params.SetLocationID),
		"location_id__n":        netboxapi.FilterStrings(params.SetLocationIDn),
		"tenant_id":             netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":          netboxapi.FilterInt64s(params.SetTenantIDn),
		"role_id":               netboxapi.FilterInt64s(params.SetRoleID),
		"role_id__n":            netboxapi.FilterInt64s(params.SetRoleIDn),
		"rack_type_id":          netboxapi.FilterInt64s(params.SetRackTypeID),
		"rack_type_id__n":       netboxapi.FilterInt64s(params.SetRackTypeIDn),
		"status":                netboxapi.FilterStrings(params.SetStatus),
		"status__n":             netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":            netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":           netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":            netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":           netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":           netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":          netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":           netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":          netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":         netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":         netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":        netboxapi.FilterStrings(params.SetStatusIregex),
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
		"u_height":              netboxapi.FilterInt64s(params.SetUHeight),
		"u_height__n":           netboxapi.FilterInt64s(params.SetUHeightn),
		"u_height__lt":          netboxapi.FilterInt64s(params.SetUHeightLt),
		"u_height__lte":         netboxapi.FilterInt64s(params.SetUHeightLte),
		"u_height__gt":          netboxapi.FilterInt64s(params.SetUHeightGt),
		"u_height__gte":         netboxapi.FilterInt64s(params.SetUHeightGte),
		"u_height__empty":       netboxapi.FilterBool(params.SetUHeightEmpty),
		"desc_units":            netboxapi.FilterBool(params.SetDescUnits),
		"serial":                netboxapi.FilterStrings(params.SetSerial),
		"serial__n":             netboxapi.FilterStrings(params.SetSerialn),
		"serial__ic":            netboxapi.FilterStrings(params.SetSerialIc),
		"serial__nic":           netboxapi.FilterStrings(params.SetSerialNic),
		"serial__ie":            netboxapi.FilterStrings(params.SetSerialIe),
		"serial__nie":           netboxapi.FilterStrings(params.SetSerialNie),
		"serial__isw":           netboxapi.FilterStrings(params.SetSerialIsw),
		"serial__nisw":          netboxapi.FilterStrings(params.SetSerialNisw),
		"serial__iew":           netboxapi.FilterStrings(params.SetSerialIew),
		"serial__niew":          netboxapi.FilterStrings(params.SetSerialNiew),
		"serial__empty":         netboxapi.FilterBool(params.SetSerialEmpty),
		"serial__regex":         netboxapi.FilterStrings(params.SetSerialRegex),
		"serial__iregex":        netboxapi.FilterStrings(params.SetSerialIregex),
		"asset_tag":             netboxapi.FilterStrings(params.SetAssetTag),
		"asset_tag__n":          netboxapi.FilterStrings(params.SetAssetTagn),
		"asset_tag__ic":         netboxapi.FilterStrings(params.SetAssetTagIc),
		"asset_tag__nic":        netboxapi.FilterStrings(params.SetAssetTagNic),
		"asset_tag__ie":         netboxapi.FilterStrings(params.SetAssetTagIe),
		"asset_tag__nie":        netboxapi.FilterStrings(params.SetAssetTagNie),
		"asset_tag__isw":        netboxapi.FilterStrings(params.SetAssetTagIsw),
		"asset_tag__nisw":       netboxapi.FilterStrings(params.SetAssetTagNisw),
		"asset_tag__iew":        netboxapi.FilterStrings(params.SetAssetTagIew),
		"asset_tag__niew":       netboxapi.FilterStrings(params.SetAssetTagNiew),
		"asset_tag__empty":      netboxapi.FilterBool(params.SetAssetTagEmpty),
		"asset_tag__regex":      netboxapi.FilterStrings(params.SetAssetTagRegex),
		"asset_tag__iregex":     netboxapi.FilterStrings(params.SetAssetTagIregex),
		"facility_id":           netboxapi.FilterStrings(params.SetFacilityID),
		"facility_id__n":        netboxapi.FilterStrings(params.SetFacilityIDn),
		"facility_id__ic":       netboxapi.FilterStrings(params.SetFacilityIDIc),
		"facility_id__nic":      netboxapi.FilterStrings(params.SetFacilityIDNic),
		"facility_id__ie":       netboxapi.FilterStrings(params.SetFacilityIDIe),
		"facility_id__nie":      netboxapi.FilterStrings(params.SetFacilityIDNie),
		"facility_id__isw":      netboxapi.FilterStrings(params.SetFacilityIDIsw),
		"facility_id__nisw":     netboxapi.FilterStrings(params.SetFacilityIDNisw),
		"facility_id__iew":      netboxapi.FilterStrings(params.SetFacilityIDIew),
		"facility_id__niew":     netboxapi.FilterStrings(params.SetFacilityIDNiew),
		"facility_id__empty":    netboxapi.FilterBool(params.SetFacilityIDEmpty),
		"facility_id__regex":    netboxapi.FilterStrings(params.SetFacilityIDRegex),
		"facility_id__iregex":   netboxapi.FilterStrings(params.SetFacilityIDIregex),
		"outer_width":           netboxapi.FilterInt64s(params.SetOuterWidth),
		"outer_width__n":        netboxapi.FilterInt64s(params.SetOuterWidthn),
		"outer_width__lt":       netboxapi.FilterInt64s(params.SetOuterWidthLt),
		"outer_width__lte":      netboxapi.FilterInt64s(params.SetOuterWidthLte),
		"outer_width__gt":       netboxapi.FilterInt64s(params.SetOuterWidthGt),
		"outer_width__gte":      netboxapi.FilterInt64s(params.SetOuterWidthGte),
		"outer_width__empty":    netboxapi.FilterBool(params.SetOuterWidthEmpty),
		"outer_depth":           netboxapi.FilterInt64s(params.SetOuterDepth),
		"outer_depth__n":        netboxapi.FilterInt64s(params.SetOuterDepthn),
		"outer_depth__lt":       netboxapi.FilterInt64s(params.SetOuterDepthLt),
		"outer_depth__lte":      netboxapi.FilterInt64s(params.SetOuterDepthLte),
		"outer_depth__gt":       netboxapi.FilterInt64s(params.SetOuterDepthGt),
		"outer_depth__gte":      netboxapi.FilterInt64s(params.SetOuterDepthGte),
		"outer_depth__empty":    netboxapi.FilterBool(params.SetOuterDepthEmpty),
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
		"mounting_depth":        netboxapi.FilterInt64s(params.SetMountingDepth),
		"mounting_depth__n":     netboxapi.FilterInt64s(params.SetMountingDepthn),
		"mounting_depth__lt":    netboxapi.FilterInt64s(params.SetMountingDepthLt),
		"mounting_depth__lte":   netboxapi.FilterInt64s(params.SetMountingDepthLte),
		"mounting_depth__gt":    netboxapi.FilterInt64s(params.SetMountingDepthGt),
		"mounting_depth__gte":   netboxapi.FilterInt64s(params.SetMountingDepthGte),
		"mounting_depth__empty": netboxapi.FilterBool(params.SetMountingDepthEmpty),
		"weight":                netboxapi.FilterFloat64s(params.SetWeight),
		"weight__n":             netboxapi.FilterFloat64s(params.SetWeightn),
		"weight__lt":            netboxapi.FilterFloat64s(params.SetWeightLt),
		"weight__lte":           netboxapi.FilterFloat64s(params.SetWeightLte),
		"weight__gt":            netboxapi.FilterFloat64s(params.SetWeightGt),
		"weight__gte":           netboxapi.FilterFloat64s(params.SetWeightGte),
		"weight__empty":         netboxapi.FilterBool(params.SetWeightEmpty),
		"max_weight":            netboxapi.FilterInt64s(params.SetMaxWeight),
		"max_weight__n":         netboxapi.FilterInt64s(params.SetMaxWeightn),
		"max_weight__lt":        netboxapi.FilterInt64s(params.SetMaxWeightLt),
		"max_weight__lte":       netboxapi.FilterInt64s(params.SetMaxWeightLte),
		"max_weight__gt":        netboxapi.FilterInt64s(params.SetMaxWeightGt),
		"max_weight__gte":       netboxapi.FilterInt64s(params.SetMaxWeightGte),
		"max_weight__empty":     netboxapi.FilterBool(params.SetMaxWeightEmpty),
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
		"region_id":             netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":          netboxapi.FilterStrings(params.SetRegionIDn),
		"contact":               netboxapi.FilterInt64s(params.SetContact),
		"contact__n":            netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":         netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":      netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":          netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":       netboxapi.FilterInt64s(params.SetContactRolen),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"airflow":               netboxapi.FilterString(params.SetAirflow),
		"airflow__n":            netboxapi.FilterString(params.SetAirflown),
		"airflow__ic":           netboxapi.FilterStrings(params.SetAirflowIc),
		"airflow__nic":          netboxapi.FilterStrings(params.SetAirflowNic),
		"airflow__ie":           netboxapi.FilterStrings(params.SetAirflowIe),
		"airflow__nie":          netboxapi.FilterStrings(params.SetAirflowNie),
		"airflow__isw":          netboxapi.FilterStrings(params.SetAirflowIsw),
		"airflow__nisw":         netboxapi.FilterStrings(params.SetAirflowNisw),
		"airflow__iew":          netboxapi.FilterStrings(params.SetAirflowIew),
		"airflow__niew":         netboxapi.FilterStrings(params.SetAirflowNiew),
		"airflow__empty":        netboxapi.FilterBool(params.SetAirflowEmpty),
		"airflow__regex":        netboxapi.FilterStrings(params.SetAirflowRegex),
		"airflow__iregex":       netboxapi.FilterStrings(params.SetAirflowIregex),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
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
		"group":                 netboxapi.FilterStrings(params.SetGroup),
		"group__n":              netboxapi.FilterStrings(params.SetGroupn),
		"group_id":              netboxapi.FilterInt64s(params.SetGroupID),
		"group_id__n":           netboxapi.FilterInt64s(params.SetGroupIDn),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":              netboxapi.FilterStrings(params.SetLocation),
		"location__n":           netboxapi.FilterStrings(params.SetLocationn),
		"manufacturer":          netboxapi.FilterStrings(params.SetManufacturer),
		"manufacturer__n":       netboxapi.FilterStrings(params.SetManufacturern),
		"manufacturer_id":       netboxapi.FilterInt64s(params.SetManufacturerID),
		"manufacturer_id__n":    netboxapi.FilterInt64s(params.SetManufacturerIDn),
		"outer_height":          netboxapi.FilterInt64s(params.SetOuterHeight),
		"outer_height__n":       netboxapi.FilterInt64s(params.SetOuterHeightn),
		"outer_height__lt":      netboxapi.FilterInt64s(params.SetOuterHeightLt),
		"outer_height__lte":     netboxapi.FilterInt64s(params.SetOuterHeightLte),
		"outer_height__gt":      netboxapi.FilterInt64s(params.SetOuterHeightGt),
		"outer_height__gte":     netboxapi.FilterInt64s(params.SetOuterHeightGte),
		"outer_height__empty":   netboxapi.FilterBool(params.SetOuterHeightEmpty),
		"owner":                 netboxapi.FilterStrings(params.SetOwner),
		"owner__n":              netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":           netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":        netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":        netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":     netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                     netboxapi.FilterString(params.SetQ),
		"rack_type":             netboxapi.FilterStrings(params.SetRackType),
		"rack_type__n":          netboxapi.FilterStrings(params.SetRackTypen),
		"region":                netboxapi.FilterStrings(params.SetRegion),
		"region__n":             netboxapi.FilterStrings(params.SetRegionn),
		"role":                  netboxapi.FilterStrings(params.SetRole),
		"role__n":               netboxapi.FilterStrings(params.SetRolen),
		"site":                  netboxapi.FilterStrings(params.SetSite),
		"site__n":               netboxapi.FilterStrings(params.SetSiten),
		"site_group":            netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":         netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":         netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":      netboxapi.FilterStrings(params.SetSiteGroupIDn),
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
		"tenant":                netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":             netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":          netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":       netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":       netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":    netboxapi.FilterStrings(params.SetTenantGroupIDn),
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
		res, err := d.client.Dcim.DcimRacksListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_racks", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.RackResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
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
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: rackResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
