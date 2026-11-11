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
	_ datasource.DataSource              = (*powerFeedsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*powerFeedsDataSource)(nil)
)

// NewPowerFeedsDataSource returns a new power_feeds data source, which
// lists power_feed objects matching its filters.
func NewPowerFeedsDataSource() datasource.DataSource {
	return &powerFeedsDataSource{}
}

type powerFeedsDataSource struct {
	client *netboxapi.Client
}

// powerFeedsDataSourceModel is the data source model: the filters, the limit and the matching
// power_feeds.
type powerFeedsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"power_feeds"`
}

// powerFeedResourceAttrTypes is the attribute type map of powerFeedResourceModel.
func powerFeedResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                      types.Int64Type,
		"name":                    types.StringType,
		"power_panel_id":          types.Int64Type,
		"rack_id":                 types.Int64Type,
		"tenant_id":               types.Int64Type,
		"status":                  types.StringType,
		"type":                    types.StringType,
		"supply":                  types.StringType,
		"phase":                   types.StringType,
		"voltage":                 types.Int64Type,
		"amperage":                types.Int64Type,
		"max_utilization_percent": types.Int64Type,
		"max_percent_utilization": types.Int64Type,
		"mark_connected":          types.BoolType,
		"description":             types.StringType,
		"comments":                types.StringType,
		"owner_id":                types.Int64Type,
		"created":                 types.StringType,
		"last_updated":            types.StringType,
		"url":                     types.StringType,
		"tags":                    types.SetType{ElemType: types.StringType},
		"tags_all":                types.SetType{ElemType: types.StringType},
		"custom_fields":           types.MapType{ElemType: types.StringType},
	}
}

func (d *powerFeedsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_power_feeds"
}

func (d *powerFeedsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists power_feed objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: amperage, amperage__empty, amperage__gt, amperage__gte, amperage__lt, amperage__lte, amperage__n, available_power, available_power__empty, available_power__gt, available_power__gte, available_power__lt, available_power__lte, available_power__n, cable_connector, cable_connector__empty, cable_connector__gt, cable_connector__gte, cable_connector__lt, cable_connector__lte, cable_connector__n, cable_end, cable_end__empty, cable_end__ic, cable_end__ie, cable_end__iew, cable_end__iregex, cable_end__isw, cable_end__n, cable_end__nic, cable_end__nie, cable_end__niew, cable_end__nisw, cable_end__regex, cable_id, cable_id__n, cabled, connected, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mark_connected, max_utilization, max_utilization__empty, max_utilization__gt, max_utilization__gte, max_utilization__lt, max_utilization__lte, max_utilization__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, occupied, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, phase, phase__empty, phase__ic, phase__ie, phase__iew, phase__iregex, phase__isw, phase__n, phase__nic, phase__nie, phase__niew, phase__nisw, phase__regex, power_panel_id, power_panel_id__n, q, rack_id, rack_id__n, region, region__n, region_id, region_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, supply, supply__empty, supply__ic, supply__ie, supply__iew, supply__iregex, supply__isw, supply__n, supply__nic, supply__nie, supply__niew, supply__nisw, supply__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex, voltage, voltage__empty, voltage__gt, voltage__gte, voltage__lt, voltage__lte, voltage__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"power_feeds": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching power_feeds.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"power_panel_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the power panel the feed starts at.",
						},
						"rack_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the rack the feed serves.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: offline, active, planned, failed.",
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Whether the feed is the primary or the redundant one. One of: primary, redundant.",
						},
						"supply": schema.StringAttribute{
							Computed:    true,
							Description: "Supply type. One of: ac, dc.",
						},
						"phase": schema.StringAttribute{
							Computed:    true,
							Description: "Phase type. One of: single-phase, three-phase.",
						},
						"voltage": schema.Int64Attribute{
							Computed:    true,
							Description: "Voltage of the feed. Must be a whole number.",
						},
						"amperage": schema.Int64Attribute{
							Computed:    true,
							Description: "Amperage of the feed. Must be a whole number.",
						},
						"max_utilization_percent": schema.Int64Attribute{
							Computed:    true,
							Description: "Maximum permitted draw as a percentage. Must be a whole number.",
						},
						"max_percent_utilization": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use max_utilization_percent instead.",
							Description:        "Deprecated alias of max_utilization_percent.",
						},
						"mark_connected": schema.BoolAttribute{
							Computed:    true,
							Description: "Treat the feed as connected even without a cable.",
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

func (d *powerFeedsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *powerFeedsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data powerFeedsDataSourceModel

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
	var responseDTOs []*netboxapi.PowerFeedResponseDTO

	params := dcim.NewDcimPowerFeedsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                     netboxapi.FilterInt64s(params.SetID),
		"id__n":                  netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                 netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                 netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":              netboxapi.FilterBool(params.SetIDEmpty),
		"name":                   netboxapi.FilterStrings(params.SetName),
		"name__n":                netboxapi.FilterStrings(params.SetNamen),
		"name__ic":               netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":              netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":               netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":              netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":              netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":             netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":              netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":             netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":            netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":            netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":           netboxapi.FilterStrings(params.SetNameIregex),
		"power_panel_id":         netboxapi.FilterInt64s(params.SetPowerPanelID),
		"power_panel_id__n":      netboxapi.FilterInt64s(params.SetPowerPanelIDn),
		"rack_id":                netboxapi.FilterInt64s(params.SetRackID),
		"rack_id__n":             netboxapi.FilterInt64s(params.SetRackIDn),
		"status":                 netboxapi.FilterStrings(params.SetStatus),
		"status__n":              netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":             netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":            netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":             netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":            netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":            netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":           netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":            netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":           netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":          netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":          netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":         netboxapi.FilterStrings(params.SetStatusIregex),
		"type":                   netboxapi.FilterString(params.SetType),
		"type__n":                netboxapi.FilterString(params.SetTypen),
		"type__ic":               netboxapi.FilterStrings(params.SetTypeIc),
		"type__nic":              netboxapi.FilterStrings(params.SetTypeNic),
		"type__ie":               netboxapi.FilterStrings(params.SetTypeIe),
		"type__nie":              netboxapi.FilterStrings(params.SetTypeNie),
		"type__isw":              netboxapi.FilterStrings(params.SetTypeIsw),
		"type__nisw":             netboxapi.FilterStrings(params.SetTypeNisw),
		"type__iew":              netboxapi.FilterStrings(params.SetTypeIew),
		"type__niew":             netboxapi.FilterStrings(params.SetTypeNiew),
		"type__empty":            netboxapi.FilterBool(params.SetTypeEmpty),
		"type__regex":            netboxapi.FilterStrings(params.SetTypeRegex),
		"type__iregex":           netboxapi.FilterStrings(params.SetTypeIregex),
		"owner_id":               netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":            netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                    netboxapi.FilterStrings(params.SetTag),
		"tag__n":                 netboxapi.FilterStrings(params.SetTagn),
		"tag__any":               netboxapi.FilterStrings(params.SetTagAny),
		"amperage":               netboxapi.FilterInt64s(params.SetAmperage),
		"amperage__n":            netboxapi.FilterInt64s(params.SetAmperagen),
		"amperage__lt":           netboxapi.FilterInt64s(params.SetAmperageLt),
		"amperage__lte":          netboxapi.FilterInt64s(params.SetAmperageLte),
		"amperage__gt":           netboxapi.FilterInt64s(params.SetAmperageGt),
		"amperage__gte":          netboxapi.FilterInt64s(params.SetAmperageGte),
		"amperage__empty":        netboxapi.FilterBool(params.SetAmperageEmpty),
		"available_power":        netboxapi.FilterInt64s(params.SetAvailablePower),
		"available_power__n":     netboxapi.FilterInt64s(params.SetAvailablePowern),
		"available_power__lt":    netboxapi.FilterInt64s(params.SetAvailablePowerLt),
		"available_power__lte":   netboxapi.FilterInt64s(params.SetAvailablePowerLte),
		"available_power__gt":    netboxapi.FilterInt64s(params.SetAvailablePowerGt),
		"available_power__gte":   netboxapi.FilterInt64s(params.SetAvailablePowerGte),
		"available_power__empty": netboxapi.FilterBool(params.SetAvailablePowerEmpty),
		"cable_connector":        netboxapi.FilterInt64s(params.SetCableConnector),
		"cable_connector__n":     netboxapi.FilterInt64s(params.SetCableConnectorn),
		"cable_connector__lt":    netboxapi.FilterInt64s(params.SetCableConnectorLt),
		"cable_connector__lte":   netboxapi.FilterInt64s(params.SetCableConnectorLte),
		"cable_connector__gt":    netboxapi.FilterInt64s(params.SetCableConnectorGt),
		"cable_connector__gte":   netboxapi.FilterInt64s(params.SetCableConnectorGte),
		"cable_connector__empty": netboxapi.FilterBool(params.SetCableConnectorEmpty),
		"cable_end":              netboxapi.FilterString(params.SetCableEnd),
		"cable_end__n":           netboxapi.FilterString(params.SetCableEndn),
		"cable_end__ic":          netboxapi.FilterStrings(params.SetCableEndIc),
		"cable_end__nic":         netboxapi.FilterStrings(params.SetCableEndNic),
		"cable_end__ie":          netboxapi.FilterStrings(params.SetCableEndIe),
		"cable_end__nie":         netboxapi.FilterStrings(params.SetCableEndNie),
		"cable_end__isw":         netboxapi.FilterStrings(params.SetCableEndIsw),
		"cable_end__nisw":        netboxapi.FilterStrings(params.SetCableEndNisw),
		"cable_end__iew":         netboxapi.FilterStrings(params.SetCableEndIew),
		"cable_end__niew":        netboxapi.FilterStrings(params.SetCableEndNiew),
		"cable_end__empty":       netboxapi.FilterBool(params.SetCableEndEmpty),
		"cable_end__regex":       netboxapi.FilterStrings(params.SetCableEndRegex),
		"cable_end__iregex":      netboxapi.FilterStrings(params.SetCableEndIregex),
		"cable_id":               netboxapi.FilterInt64s(params.SetCableID),
		"cable_id__n":            netboxapi.FilterInt64s(params.SetCableIDn),
		"cabled":                 netboxapi.FilterBool(params.SetCabled),
		"connected":              netboxapi.FilterBool(params.SetConnected),
		"created":                netboxapi.FilterStrings(params.SetCreated),
		"created__n":             netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":            netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":           netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":            netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":           netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":         netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":            netboxapi.FilterStrings(params.SetDescription),
		"description__n":         netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":        netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":       netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":        netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":       netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":       netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":      netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":       netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":      netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":     netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":     netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":    netboxapi.FilterStrings(params.SetDescriptionIregex),
		"last_updated":           netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":        netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":       netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":      netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":       netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":      netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":    netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"mark_connected":         netboxapi.FilterBool(params.SetMarkConnected),
		"max_utilization":        netboxapi.FilterInt64s(params.SetMaxUtilization),
		"max_utilization__n":     netboxapi.FilterInt64s(params.SetMaxUtilizationn),
		"max_utilization__lt":    netboxapi.FilterInt64s(params.SetMaxUtilizationLt),
		"max_utilization__lte":   netboxapi.FilterInt64s(params.SetMaxUtilizationLte),
		"max_utilization__gt":    netboxapi.FilterInt64s(params.SetMaxUtilizationGt),
		"max_utilization__gte":   netboxapi.FilterInt64s(params.SetMaxUtilizationGte),
		"max_utilization__empty": netboxapi.FilterBool(params.SetMaxUtilizationEmpty),
		"occupied":               netboxapi.FilterBool(params.SetOccupied),
		"owner":                  netboxapi.FilterStrings(params.SetOwner),
		"owner__n":               netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":            netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":         netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":         netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":      netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"phase":                  netboxapi.FilterString(params.SetPhase),
		"phase__n":               netboxapi.FilterString(params.SetPhasen),
		"phase__ic":              netboxapi.FilterStrings(params.SetPhaseIc),
		"phase__nic":             netboxapi.FilterStrings(params.SetPhaseNic),
		"phase__ie":              netboxapi.FilterStrings(params.SetPhaseIe),
		"phase__nie":             netboxapi.FilterStrings(params.SetPhaseNie),
		"phase__isw":             netboxapi.FilterStrings(params.SetPhaseIsw),
		"phase__nisw":            netboxapi.FilterStrings(params.SetPhaseNisw),
		"phase__iew":             netboxapi.FilterStrings(params.SetPhaseIew),
		"phase__niew":            netboxapi.FilterStrings(params.SetPhaseNiew),
		"phase__empty":           netboxapi.FilterBool(params.SetPhaseEmpty),
		"phase__regex":           netboxapi.FilterStrings(params.SetPhaseRegex),
		"phase__iregex":          netboxapi.FilterStrings(params.SetPhaseIregex),
		"q":                      netboxapi.FilterString(params.SetQ),
		"region":                 netboxapi.FilterStrings(params.SetRegion),
		"region__n":              netboxapi.FilterStrings(params.SetRegionn),
		"region_id":              netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":           netboxapi.FilterStrings(params.SetRegionIDn),
		"site":                   netboxapi.FilterStrings(params.SetSite),
		"site__n":                netboxapi.FilterStrings(params.SetSiten),
		"site_group":             netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":          netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":          netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":       netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"site_id":                netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":             netboxapi.FilterInt64s(params.SetSiteIDn),
		"supply":                 netboxapi.FilterString(params.SetSupply),
		"supply__n":              netboxapi.FilterString(params.SetSupplyn),
		"supply__ic":             netboxapi.FilterStrings(params.SetSupplyIc),
		"supply__nic":            netboxapi.FilterStrings(params.SetSupplyNic),
		"supply__ie":             netboxapi.FilterStrings(params.SetSupplyIe),
		"supply__nie":            netboxapi.FilterStrings(params.SetSupplyNie),
		"supply__isw":            netboxapi.FilterStrings(params.SetSupplyIsw),
		"supply__nisw":           netboxapi.FilterStrings(params.SetSupplyNisw),
		"supply__iew":            netboxapi.FilterStrings(params.SetSupplyIew),
		"supply__niew":           netboxapi.FilterStrings(params.SetSupplyNiew),
		"supply__empty":          netboxapi.FilterBool(params.SetSupplyEmpty),
		"supply__regex":          netboxapi.FilterStrings(params.SetSupplyRegex),
		"supply__iregex":         netboxapi.FilterStrings(params.SetSupplyIregex),
		"tag_id":                 netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":              netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":            netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                 netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":              netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":           netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":        netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":        netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":     netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"tenant_id":              netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":           netboxapi.FilterInt64s(params.SetTenantIDn),
		"voltage":                netboxapi.FilterInt64s(params.SetVoltage),
		"voltage__n":             netboxapi.FilterInt64s(params.SetVoltagen),
		"voltage__lt":            netboxapi.FilterInt64s(params.SetVoltageLt),
		"voltage__lte":           netboxapi.FilterInt64s(params.SetVoltageLte),
		"voltage__gt":            netboxapi.FilterInt64s(params.SetVoltageGt),
		"voltage__gte":           netboxapi.FilterInt64s(params.SetVoltageGte),
		"voltage__empty":         netboxapi.FilterBool(params.SetVoltageEmpty),
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
		res, err := d.client.Dcim.DcimPowerFeedsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_power_feeds", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.PowerFeedResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]powerFeedResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model powerFeedResourceModel
		resp.Diagnostics.Append(flattenPowerFeed(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: powerFeedResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
