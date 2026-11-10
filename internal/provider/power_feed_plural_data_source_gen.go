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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vName []string
	var vNamen []string
	var vNameIc []string
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
	var vPowerPanelID []int64
	var vPowerPanelIDn []int64
	var vRackID []int64
	var vRackIDn []int64
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
	var vType *string
	var vTypen *string
	var vTypeIc []string
	var vTypeNic []string
	var vTypeIe []string
	var vTypeNie []string
	var vTypeIsw []string
	var vTypeNisw []string
	var vTypeIew []string
	var vTypeNiew []string
	var vTypeEmpty *bool
	var vTypeRegex []string
	var vTypeIregex []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAmperage []int64
	var vAmperagen []int64
	var vAmperageLt []int64
	var vAmperageLte []int64
	var vAmperageGt []int64
	var vAmperageGte []int64
	var vAmperageEmpty *bool
	var vAvailablePower []int64
	var vAvailablePowern []int64
	var vAvailablePowerLt []int64
	var vAvailablePowerLte []int64
	var vAvailablePowerGt []int64
	var vAvailablePowerGte []int64
	var vAvailablePowerEmpty *bool
	var vCableConnector []int64
	var vCableConnectorn []int64
	var vCableConnectorLt []int64
	var vCableConnectorLte []int64
	var vCableConnectorGt []int64
	var vCableConnectorGte []int64
	var vCableConnectorEmpty *bool
	var vCableEnd *string
	var vCableEndn *string
	var vCableEndIc []string
	var vCableEndNic []string
	var vCableEndIe []string
	var vCableEndNie []string
	var vCableEndIsw []string
	var vCableEndNisw []string
	var vCableEndIew []string
	var vCableEndNiew []string
	var vCableEndEmpty *bool
	var vCableEndRegex []string
	var vCableEndIregex []string
	var vCableID []int64
	var vCableIDn []int64
	var vCabled *bool
	var vConnected *bool
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
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
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vMarkConnected *bool
	var vMaxUtilization []int64
	var vMaxUtilizationn []int64
	var vMaxUtilizationLt []int64
	var vMaxUtilizationLte []int64
	var vMaxUtilizationGt []int64
	var vMaxUtilizationGte []int64
	var vMaxUtilizationEmpty *bool
	var vOccupied *bool
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPhase *string
	var vPhasen *string
	var vPhaseIc []string
	var vPhaseNic []string
	var vPhaseIe []string
	var vPhaseNie []string
	var vPhaseIsw []string
	var vPhaseNisw []string
	var vPhaseIew []string
	var vPhaseNiew []string
	var vPhaseEmpty *bool
	var vPhaseRegex []string
	var vPhaseIregex []string
	var vQ *string
	var vRegion []string
	var vRegionn []string
	var vRegionID []string
	var vRegionIDn []string
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []string
	var vSiteGroupIDn []string
	var vSiteID []int64
	var vSiteIDn []int64
	var vSupply *string
	var vSupplyn *string
	var vSupplyIc []string
	var vSupplyNic []string
	var vSupplyIe []string
	var vSupplyNie []string
	var vSupplyIsw []string
	var vSupplyNisw []string
	var vSupplyIew []string
	var vSupplyNiew []string
	var vSupplyEmpty *bool
	var vSupplyRegex []string
	var vSupplyIregex []string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
	var vTenantID []int64
	var vTenantIDn []int64
	var vVoltage []int64
	var vVoltagen []int64
	var vVoltageLt []int64
	var vVoltageLte []int64
	var vVoltageGt []int64
	var vVoltageGte []int64
	var vVoltageEmpty *bool
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
		case "name__ic":
			v := value
			vNameIc = append(vNameIc, v)
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
		case "power_panel_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_panel_id' takes an integer, got %q.", value))
				return
			}
			vPowerPanelID = append(vPowerPanelID, v)
		case "power_panel_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_panel_id__n' takes an integer, got %q.", value))
				return
			}
			vPowerPanelIDn = append(vPowerPanelIDn, v)
		case "rack_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_id' takes an integer, got %q.", value))
				return
			}
			vRackID = append(vRackID, v)
		case "rack_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_id__n' takes an integer, got %q.", value))
				return
			}
			vRackIDn = append(vRackIDn, v)
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
		case "type":
			if vType != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'type' takes a single value.")
				return
			}
			v := value
			vType = &v
		case "type__n":
			if vTypen != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'type__n' takes a single value.")
				return
			}
			v := value
			vTypen = &v
		case "type__ic":
			v := value
			vTypeIc = append(vTypeIc, v)
		case "type__nic":
			v := value
			vTypeNic = append(vTypeNic, v)
		case "type__ie":
			v := value
			vTypeIe = append(vTypeIe, v)
		case "type__nie":
			v := value
			vTypeNie = append(vTypeNie, v)
		case "type__isw":
			v := value
			vTypeIsw = append(vTypeIsw, v)
		case "type__nisw":
			v := value
			vTypeNisw = append(vTypeNisw, v)
		case "type__iew":
			v := value
			vTypeIew = append(vTypeIew, v)
		case "type__niew":
			v := value
			vTypeNiew = append(vTypeNiew, v)
		case "type__empty":
			if vTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'type__empty' takes a boolean, got %q.", value))
				return
			}
			vTypeEmpty = &v
		case "type__regex":
			v := value
			vTypeRegex = append(vTypeRegex, v)
		case "type__iregex":
			v := value
			vTypeIregex = append(vTypeIregex, v)
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
		case "amperage":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'amperage' takes an integer, got %q.", value))
				return
			}
			vAmperage = append(vAmperage, v)
		case "amperage__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'amperage__n' takes an integer, got %q.", value))
				return
			}
			vAmperagen = append(vAmperagen, v)
		case "amperage__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'amperage__lt' takes an integer, got %q.", value))
				return
			}
			vAmperageLt = append(vAmperageLt, v)
		case "amperage__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'amperage__lte' takes an integer, got %q.", value))
				return
			}
			vAmperageLte = append(vAmperageLte, v)
		case "amperage__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'amperage__gt' takes an integer, got %q.", value))
				return
			}
			vAmperageGt = append(vAmperageGt, v)
		case "amperage__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'amperage__gte' takes an integer, got %q.", value))
				return
			}
			vAmperageGte = append(vAmperageGte, v)
		case "amperage__empty":
			if vAmperageEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'amperage__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'amperage__empty' takes a boolean, got %q.", value))
				return
			}
			vAmperageEmpty = &v
		case "available_power":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'available_power' takes an integer, got %q.", value))
				return
			}
			vAvailablePower = append(vAvailablePower, v)
		case "available_power__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'available_power__n' takes an integer, got %q.", value))
				return
			}
			vAvailablePowern = append(vAvailablePowern, v)
		case "available_power__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'available_power__lt' takes an integer, got %q.", value))
				return
			}
			vAvailablePowerLt = append(vAvailablePowerLt, v)
		case "available_power__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'available_power__lte' takes an integer, got %q.", value))
				return
			}
			vAvailablePowerLte = append(vAvailablePowerLte, v)
		case "available_power__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'available_power__gt' takes an integer, got %q.", value))
				return
			}
			vAvailablePowerGt = append(vAvailablePowerGt, v)
		case "available_power__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'available_power__gte' takes an integer, got %q.", value))
				return
			}
			vAvailablePowerGte = append(vAvailablePowerGte, v)
		case "available_power__empty":
			if vAvailablePowerEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'available_power__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'available_power__empty' takes a boolean, got %q.", value))
				return
			}
			vAvailablePowerEmpty = &v
		case "cable_connector":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector' takes an integer, got %q.", value))
				return
			}
			vCableConnector = append(vCableConnector, v)
		case "cable_connector__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__n' takes an integer, got %q.", value))
				return
			}
			vCableConnectorn = append(vCableConnectorn, v)
		case "cable_connector__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__lt' takes an integer, got %q.", value))
				return
			}
			vCableConnectorLt = append(vCableConnectorLt, v)
		case "cable_connector__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__lte' takes an integer, got %q.", value))
				return
			}
			vCableConnectorLte = append(vCableConnectorLte, v)
		case "cable_connector__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__gt' takes an integer, got %q.", value))
				return
			}
			vCableConnectorGt = append(vCableConnectorGt, v)
		case "cable_connector__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__gte' takes an integer, got %q.", value))
				return
			}
			vCableConnectorGte = append(vCableConnectorGte, v)
		case "cable_connector__empty":
			if vCableConnectorEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_connector__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_connector__empty' takes a boolean, got %q.", value))
				return
			}
			vCableConnectorEmpty = &v
		case "cable_end":
			if vCableEnd != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end' takes a single value.")
				return
			}
			v := value
			vCableEnd = &v
		case "cable_end__n":
			if vCableEndn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end__n' takes a single value.")
				return
			}
			v := value
			vCableEndn = &v
		case "cable_end__ic":
			v := value
			vCableEndIc = append(vCableEndIc, v)
		case "cable_end__nic":
			v := value
			vCableEndNic = append(vCableEndNic, v)
		case "cable_end__ie":
			v := value
			vCableEndIe = append(vCableEndIe, v)
		case "cable_end__nie":
			v := value
			vCableEndNie = append(vCableEndNie, v)
		case "cable_end__isw":
			v := value
			vCableEndIsw = append(vCableEndIsw, v)
		case "cable_end__nisw":
			v := value
			vCableEndNisw = append(vCableEndNisw, v)
		case "cable_end__iew":
			v := value
			vCableEndIew = append(vCableEndIew, v)
		case "cable_end__niew":
			v := value
			vCableEndNiew = append(vCableEndNiew, v)
		case "cable_end__empty":
			if vCableEndEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cable_end__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_end__empty' takes a boolean, got %q.", value))
				return
			}
			vCableEndEmpty = &v
		case "cable_end__regex":
			v := value
			vCableEndRegex = append(vCableEndRegex, v)
		case "cable_end__iregex":
			v := value
			vCableEndIregex = append(vCableEndIregex, v)
		case "cable_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_id' takes an integer, got %q.", value))
				return
			}
			vCableID = append(vCableID, v)
		case "cable_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cable_id__n' takes an integer, got %q.", value))
				return
			}
			vCableIDn = append(vCableIDn, v)
		case "cabled":
			if vCabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cabled' takes a boolean, got %q.", value))
				return
			}
			vCabled = &v
		case "connected":
			if vConnected != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'connected' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'connected' takes a boolean, got %q.", value))
				return
			}
			vConnected = &v
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
		case "mark_connected":
			if vMarkConnected != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mark_connected' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mark_connected' takes a boolean, got %q.", value))
				return
			}
			vMarkConnected = &v
		case "max_utilization":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_utilization' takes an integer, got %q.", value))
				return
			}
			vMaxUtilization = append(vMaxUtilization, v)
		case "max_utilization__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_utilization__n' takes an integer, got %q.", value))
				return
			}
			vMaxUtilizationn = append(vMaxUtilizationn, v)
		case "max_utilization__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_utilization__lt' takes an integer, got %q.", value))
				return
			}
			vMaxUtilizationLt = append(vMaxUtilizationLt, v)
		case "max_utilization__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_utilization__lte' takes an integer, got %q.", value))
				return
			}
			vMaxUtilizationLte = append(vMaxUtilizationLte, v)
		case "max_utilization__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_utilization__gt' takes an integer, got %q.", value))
				return
			}
			vMaxUtilizationGt = append(vMaxUtilizationGt, v)
		case "max_utilization__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_utilization__gte' takes an integer, got %q.", value))
				return
			}
			vMaxUtilizationGte = append(vMaxUtilizationGte, v)
		case "max_utilization__empty":
			if vMaxUtilizationEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'max_utilization__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'max_utilization__empty' takes a boolean, got %q.", value))
				return
			}
			vMaxUtilizationEmpty = &v
		case "occupied":
			if vOccupied != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'occupied' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'occupied' takes a boolean, got %q.", value))
				return
			}
			vOccupied = &v
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
		case "phase":
			if vPhase != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'phase' takes a single value.")
				return
			}
			v := value
			vPhase = &v
		case "phase__n":
			if vPhasen != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'phase__n' takes a single value.")
				return
			}
			v := value
			vPhasen = &v
		case "phase__ic":
			v := value
			vPhaseIc = append(vPhaseIc, v)
		case "phase__nic":
			v := value
			vPhaseNic = append(vPhaseNic, v)
		case "phase__ie":
			v := value
			vPhaseIe = append(vPhaseIe, v)
		case "phase__nie":
			v := value
			vPhaseNie = append(vPhaseNie, v)
		case "phase__isw":
			v := value
			vPhaseIsw = append(vPhaseIsw, v)
		case "phase__nisw":
			v := value
			vPhaseNisw = append(vPhaseNisw, v)
		case "phase__iew":
			v := value
			vPhaseIew = append(vPhaseIew, v)
		case "phase__niew":
			v := value
			vPhaseNiew = append(vPhaseNiew, v)
		case "phase__empty":
			if vPhaseEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'phase__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'phase__empty' takes a boolean, got %q.", value))
				return
			}
			vPhaseEmpty = &v
		case "phase__regex":
			v := value
			vPhaseRegex = append(vPhaseRegex, v)
		case "phase__iregex":
			v := value
			vPhaseIregex = append(vPhaseIregex, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "region":
			v := value
			vRegion = append(vRegion, v)
		case "region__n":
			v := value
			vRegionn = append(vRegionn, v)
		case "region_id":
			v := value
			vRegionID = append(vRegionID, v)
		case "region_id__n":
			v := value
			vRegionIDn = append(vRegionIDn, v)
		case "site":
			v := value
			vSite = append(vSite, v)
		case "site__n":
			v := value
			vSiten = append(vSiten, v)
		case "site_group":
			v := value
			vSiteGroup = append(vSiteGroup, v)
		case "site_group__n":
			v := value
			vSiteGroupn = append(vSiteGroupn, v)
		case "site_group_id":
			v := value
			vSiteGroupID = append(vSiteGroupID, v)
		case "site_group_id__n":
			v := value
			vSiteGroupIDn = append(vSiteGroupIDn, v)
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
		case "supply":
			if vSupply != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'supply' takes a single value.")
				return
			}
			v := value
			vSupply = &v
		case "supply__n":
			if vSupplyn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'supply__n' takes a single value.")
				return
			}
			v := value
			vSupplyn = &v
		case "supply__ic":
			v := value
			vSupplyIc = append(vSupplyIc, v)
		case "supply__nic":
			v := value
			vSupplyNic = append(vSupplyNic, v)
		case "supply__ie":
			v := value
			vSupplyIe = append(vSupplyIe, v)
		case "supply__nie":
			v := value
			vSupplyNie = append(vSupplyNie, v)
		case "supply__isw":
			v := value
			vSupplyIsw = append(vSupplyIsw, v)
		case "supply__nisw":
			v := value
			vSupplyNisw = append(vSupplyNisw, v)
		case "supply__iew":
			v := value
			vSupplyIew = append(vSupplyIew, v)
		case "supply__niew":
			v := value
			vSupplyNiew = append(vSupplyNiew, v)
		case "supply__empty":
			if vSupplyEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'supply__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'supply__empty' takes a boolean, got %q.", value))
				return
			}
			vSupplyEmpty = &v
		case "supply__regex":
			v := value
			vSupplyRegex = append(vSupplyRegex, v)
		case "supply__iregex":
			v := value
			vSupplyIregex = append(vSupplyIregex, v)
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
		case "tenant":
			v := value
			vTenant = append(vTenant, v)
		case "tenant__n":
			v := value
			vTenantn = append(vTenantn, v)
		case "tenant_group":
			v := value
			vTenantGroup = append(vTenantGroup, v)
		case "tenant_group__n":
			v := value
			vTenantGroupn = append(vTenantGroupn, v)
		case "tenant_group_id":
			v := value
			vTenantGroupID = append(vTenantGroupID, v)
		case "tenant_group_id__n":
			v := value
			vTenantGroupIDn = append(vTenantGroupIDn, v)
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
		case "voltage":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'voltage' takes an integer, got %q.", value))
				return
			}
			vVoltage = append(vVoltage, v)
		case "voltage__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'voltage__n' takes an integer, got %q.", value))
				return
			}
			vVoltagen = append(vVoltagen, v)
		case "voltage__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'voltage__lt' takes an integer, got %q.", value))
				return
			}
			vVoltageLt = append(vVoltageLt, v)
		case "voltage__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'voltage__lte' takes an integer, got %q.", value))
				return
			}
			vVoltageLte = append(vVoltageLte, v)
		case "voltage__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'voltage__gt' takes an integer, got %q.", value))
				return
			}
			vVoltageGt = append(vVoltageGt, v)
		case "voltage__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'voltage__gte' takes an integer, got %q.", value))
				return
			}
			vVoltageGte = append(vVoltageGte, v)
		case "voltage__empty":
			if vVoltageEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'voltage__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'voltage__empty' takes a boolean, got %q.", value))
				return
			}
			vVoltageEmpty = &v
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
	if len(vNameIc) > 0 {
		params.SetNameIc(vNameIc)
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
	if len(vPowerPanelID) > 0 {
		params.SetPowerPanelID(vPowerPanelID)
	}
	if len(vPowerPanelIDn) > 0 {
		params.SetPowerPanelIDn(vPowerPanelIDn)
	}
	if len(vRackID) > 0 {
		params.SetRackID(vRackID)
	}
	if len(vRackIDn) > 0 {
		params.SetRackIDn(vRackIDn)
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
	if vType != nil {
		params.SetType(vType)
	}
	if vTypen != nil {
		params.SetTypen(vTypen)
	}
	if len(vTypeIc) > 0 {
		params.SetTypeIc(vTypeIc)
	}
	if len(vTypeNic) > 0 {
		params.SetTypeNic(vTypeNic)
	}
	if len(vTypeIe) > 0 {
		params.SetTypeIe(vTypeIe)
	}
	if len(vTypeNie) > 0 {
		params.SetTypeNie(vTypeNie)
	}
	if len(vTypeIsw) > 0 {
		params.SetTypeIsw(vTypeIsw)
	}
	if len(vTypeNisw) > 0 {
		params.SetTypeNisw(vTypeNisw)
	}
	if len(vTypeIew) > 0 {
		params.SetTypeIew(vTypeIew)
	}
	if len(vTypeNiew) > 0 {
		params.SetTypeNiew(vTypeNiew)
	}
	if vTypeEmpty != nil {
		params.SetTypeEmpty(vTypeEmpty)
	}
	if len(vTypeRegex) > 0 {
		params.SetTypeRegex(vTypeRegex)
	}
	if len(vTypeIregex) > 0 {
		params.SetTypeIregex(vTypeIregex)
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
	if len(vAmperage) > 0 {
		params.SetAmperage(vAmperage)
	}
	if len(vAmperagen) > 0 {
		params.SetAmperagen(vAmperagen)
	}
	if len(vAmperageLt) > 0 {
		params.SetAmperageLt(vAmperageLt)
	}
	if len(vAmperageLte) > 0 {
		params.SetAmperageLte(vAmperageLte)
	}
	if len(vAmperageGt) > 0 {
		params.SetAmperageGt(vAmperageGt)
	}
	if len(vAmperageGte) > 0 {
		params.SetAmperageGte(vAmperageGte)
	}
	if vAmperageEmpty != nil {
		params.SetAmperageEmpty(vAmperageEmpty)
	}
	if len(vAvailablePower) > 0 {
		params.SetAvailablePower(vAvailablePower)
	}
	if len(vAvailablePowern) > 0 {
		params.SetAvailablePowern(vAvailablePowern)
	}
	if len(vAvailablePowerLt) > 0 {
		params.SetAvailablePowerLt(vAvailablePowerLt)
	}
	if len(vAvailablePowerLte) > 0 {
		params.SetAvailablePowerLte(vAvailablePowerLte)
	}
	if len(vAvailablePowerGt) > 0 {
		params.SetAvailablePowerGt(vAvailablePowerGt)
	}
	if len(vAvailablePowerGte) > 0 {
		params.SetAvailablePowerGte(vAvailablePowerGte)
	}
	if vAvailablePowerEmpty != nil {
		params.SetAvailablePowerEmpty(vAvailablePowerEmpty)
	}
	if len(vCableConnector) > 0 {
		params.SetCableConnector(vCableConnector)
	}
	if len(vCableConnectorn) > 0 {
		params.SetCableConnectorn(vCableConnectorn)
	}
	if len(vCableConnectorLt) > 0 {
		params.SetCableConnectorLt(vCableConnectorLt)
	}
	if len(vCableConnectorLte) > 0 {
		params.SetCableConnectorLte(vCableConnectorLte)
	}
	if len(vCableConnectorGt) > 0 {
		params.SetCableConnectorGt(vCableConnectorGt)
	}
	if len(vCableConnectorGte) > 0 {
		params.SetCableConnectorGte(vCableConnectorGte)
	}
	if vCableConnectorEmpty != nil {
		params.SetCableConnectorEmpty(vCableConnectorEmpty)
	}
	if vCableEnd != nil {
		params.SetCableEnd(vCableEnd)
	}
	if vCableEndn != nil {
		params.SetCableEndn(vCableEndn)
	}
	if len(vCableEndIc) > 0 {
		params.SetCableEndIc(vCableEndIc)
	}
	if len(vCableEndNic) > 0 {
		params.SetCableEndNic(vCableEndNic)
	}
	if len(vCableEndIe) > 0 {
		params.SetCableEndIe(vCableEndIe)
	}
	if len(vCableEndNie) > 0 {
		params.SetCableEndNie(vCableEndNie)
	}
	if len(vCableEndIsw) > 0 {
		params.SetCableEndIsw(vCableEndIsw)
	}
	if len(vCableEndNisw) > 0 {
		params.SetCableEndNisw(vCableEndNisw)
	}
	if len(vCableEndIew) > 0 {
		params.SetCableEndIew(vCableEndIew)
	}
	if len(vCableEndNiew) > 0 {
		params.SetCableEndNiew(vCableEndNiew)
	}
	if vCableEndEmpty != nil {
		params.SetCableEndEmpty(vCableEndEmpty)
	}
	if len(vCableEndRegex) > 0 {
		params.SetCableEndRegex(vCableEndRegex)
	}
	if len(vCableEndIregex) > 0 {
		params.SetCableEndIregex(vCableEndIregex)
	}
	if len(vCableID) > 0 {
		params.SetCableID(vCableID)
	}
	if len(vCableIDn) > 0 {
		params.SetCableIDn(vCableIDn)
	}
	if vCabled != nil {
		params.SetCabled(vCabled)
	}
	if vConnected != nil {
		params.SetConnected(vConnected)
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
	if vMarkConnected != nil {
		params.SetMarkConnected(vMarkConnected)
	}
	if len(vMaxUtilization) > 0 {
		params.SetMaxUtilization(vMaxUtilization)
	}
	if len(vMaxUtilizationn) > 0 {
		params.SetMaxUtilizationn(vMaxUtilizationn)
	}
	if len(vMaxUtilizationLt) > 0 {
		params.SetMaxUtilizationLt(vMaxUtilizationLt)
	}
	if len(vMaxUtilizationLte) > 0 {
		params.SetMaxUtilizationLte(vMaxUtilizationLte)
	}
	if len(vMaxUtilizationGt) > 0 {
		params.SetMaxUtilizationGt(vMaxUtilizationGt)
	}
	if len(vMaxUtilizationGte) > 0 {
		params.SetMaxUtilizationGte(vMaxUtilizationGte)
	}
	if vMaxUtilizationEmpty != nil {
		params.SetMaxUtilizationEmpty(vMaxUtilizationEmpty)
	}
	if vOccupied != nil {
		params.SetOccupied(vOccupied)
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
	if vPhase != nil {
		params.SetPhase(vPhase)
	}
	if vPhasen != nil {
		params.SetPhasen(vPhasen)
	}
	if len(vPhaseIc) > 0 {
		params.SetPhaseIc(vPhaseIc)
	}
	if len(vPhaseNic) > 0 {
		params.SetPhaseNic(vPhaseNic)
	}
	if len(vPhaseIe) > 0 {
		params.SetPhaseIe(vPhaseIe)
	}
	if len(vPhaseNie) > 0 {
		params.SetPhaseNie(vPhaseNie)
	}
	if len(vPhaseIsw) > 0 {
		params.SetPhaseIsw(vPhaseIsw)
	}
	if len(vPhaseNisw) > 0 {
		params.SetPhaseNisw(vPhaseNisw)
	}
	if len(vPhaseIew) > 0 {
		params.SetPhaseIew(vPhaseIew)
	}
	if len(vPhaseNiew) > 0 {
		params.SetPhaseNiew(vPhaseNiew)
	}
	if vPhaseEmpty != nil {
		params.SetPhaseEmpty(vPhaseEmpty)
	}
	if len(vPhaseRegex) > 0 {
		params.SetPhaseRegex(vPhaseRegex)
	}
	if len(vPhaseIregex) > 0 {
		params.SetPhaseIregex(vPhaseIregex)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRegion) > 0 {
		params.SetRegion(vRegion)
	}
	if len(vRegionn) > 0 {
		params.SetRegionn(vRegionn)
	}
	if len(vRegionID) > 0 {
		params.SetRegionID(vRegionID)
	}
	if len(vRegionIDn) > 0 {
		params.SetRegionIDn(vRegionIDn)
	}
	if len(vSite) > 0 {
		params.SetSite(vSite)
	}
	if len(vSiten) > 0 {
		params.SetSiten(vSiten)
	}
	if len(vSiteGroup) > 0 {
		params.SetSiteGroup(vSiteGroup)
	}
	if len(vSiteGroupn) > 0 {
		params.SetSiteGroupn(vSiteGroupn)
	}
	if len(vSiteGroupID) > 0 {
		params.SetSiteGroupID(vSiteGroupID)
	}
	if len(vSiteGroupIDn) > 0 {
		params.SetSiteGroupIDn(vSiteGroupIDn)
	}
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
	}
	if vSupply != nil {
		params.SetSupply(vSupply)
	}
	if vSupplyn != nil {
		params.SetSupplyn(vSupplyn)
	}
	if len(vSupplyIc) > 0 {
		params.SetSupplyIc(vSupplyIc)
	}
	if len(vSupplyNic) > 0 {
		params.SetSupplyNic(vSupplyNic)
	}
	if len(vSupplyIe) > 0 {
		params.SetSupplyIe(vSupplyIe)
	}
	if len(vSupplyNie) > 0 {
		params.SetSupplyNie(vSupplyNie)
	}
	if len(vSupplyIsw) > 0 {
		params.SetSupplyIsw(vSupplyIsw)
	}
	if len(vSupplyNisw) > 0 {
		params.SetSupplyNisw(vSupplyNisw)
	}
	if len(vSupplyIew) > 0 {
		params.SetSupplyIew(vSupplyIew)
	}
	if len(vSupplyNiew) > 0 {
		params.SetSupplyNiew(vSupplyNiew)
	}
	if vSupplyEmpty != nil {
		params.SetSupplyEmpty(vSupplyEmpty)
	}
	if len(vSupplyRegex) > 0 {
		params.SetSupplyRegex(vSupplyRegex)
	}
	if len(vSupplyIregex) > 0 {
		params.SetSupplyIregex(vSupplyIregex)
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
	if len(vTenant) > 0 {
		params.SetTenant(vTenant)
	}
	if len(vTenantn) > 0 {
		params.SetTenantn(vTenantn)
	}
	if len(vTenantGroup) > 0 {
		params.SetTenantGroup(vTenantGroup)
	}
	if len(vTenantGroupn) > 0 {
		params.SetTenantGroupn(vTenantGroupn)
	}
	if len(vTenantGroupID) > 0 {
		params.SetTenantGroupID(vTenantGroupID)
	}
	if len(vTenantGroupIDn) > 0 {
		params.SetTenantGroupIDn(vTenantGroupIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vVoltage) > 0 {
		params.SetVoltage(vVoltage)
	}
	if len(vVoltagen) > 0 {
		params.SetVoltagen(vVoltagen)
	}
	if len(vVoltageLt) > 0 {
		params.SetVoltageLt(vVoltageLt)
	}
	if len(vVoltageLte) > 0 {
		params.SetVoltageLte(vVoltageLte)
	}
	if len(vVoltageGt) > 0 {
		params.SetVoltageGt(vVoltageGt)
	}
	if len(vVoltageGte) > 0 {
		params.SetVoltageGte(vVoltageGte)
	}
	if vVoltageEmpty != nil {
		params.SetVoltageEmpty(vVoltageEmpty)
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
