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
	_ datasource.DataSource              = (*deviceConsoleServerPortsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceConsoleServerPortsDataSource)(nil)
)

// NewDeviceConsoleServerPortsDataSource returns a new device_console_server_ports data source, which
// lists device_console_server_port objects matching its filters.
func NewDeviceConsoleServerPortsDataSource() datasource.DataSource {
	return &deviceConsoleServerPortsDataSource{}
}

type deviceConsoleServerPortsDataSource struct {
	client *netboxapi.Client
}

// deviceConsoleServerPortsDataSourceModel is the data source model: the filters, the limit and the matching
// device_console_server_ports.
type deviceConsoleServerPortsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"device_console_server_ports"`
}

// deviceConsoleServerPortResourceAttrTypes is the attribute type map of deviceConsoleServerPortResourceModel.
func deviceConsoleServerPortResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             types.Int64Type,
		"device_id":      types.Int64Type,
		"module_id":      types.Int64Type,
		"name":           types.StringType,
		"type":           types.StringType,
		"speed":          types.Int64Type,
		"label":          types.StringType,
		"mark_connected": types.BoolType,
		"description":    types.StringType,
		"owner_id":       types.Int64Type,
		"created":        types.StringType,
		"last_updated":   types.StringType,
		"url":            types.StringType,
		"tags":           types.SetType{ElemType: types.StringType},
		"tags_all":       types.SetType{ElemType: types.StringType},
		"custom_fields":  types.MapType{ElemType: types.StringType},
	}
}

func (d *deviceConsoleServerPortsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_console_server_ports"
}

func (d *deviceConsoleServerPortsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists device_console_server_port objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: cable_connector, cable_connector__empty, cable_connector__gt, cable_connector__gte, cable_connector__lt, cable_connector__lte, cable_connector__n, cable_end, cable_end__empty, cable_end__ic, cable_end__ie, cable_end__iew, cable_end__iregex, cable_end__isw, cable_end__n, cable_end__nic, cable_end__nie, cable_end__niew, cable_end__nisw, cable_end__regex, cable_id, cable_id__n, cabled, connected, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device__n, device_id, device_id__n, device_role, device_role__n, device_role_id, device_role_id__n, device_status, device_status__empty, device_status__ic, device_status__ie, device_status__iew, device_status__iregex, device_status__isw, device_status__n, device_status__nic, device_status__nie, device_status__niew, device_status__nisw, device_status__regex, device_type, device_type__n, device_type_id, device_type_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, mark_connected, module_id, module_id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, occupied, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, rack, rack__n, rack_id, rack_id__n, region, region__n, region_id, region_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, speed, speed__ic, speed__ie, speed__iew, speed__iregex, speed__isw, speed__n, speed__nic, speed__nie, speed__niew, speed__nisw, speed__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_id, tenant_id__n, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex, virtual_chassis, virtual_chassis__n, virtual_chassis_id, virtual_chassis_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"device_console_server_ports": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching device_console_server_ports.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"device_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device.",
						},
						"module_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the installed module this component belongs to.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Port type as its NetBox slug, e.g. de-9, rj-45, usb-a (any of NetBox's console server port type choices).",
						},
						"speed": schema.Int64Attribute{
							Computed:    true,
							Description: "Port speed in bps. One of: 1200, 2400, 4800, 9600, 19200, 38400, 57600, 115200.",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Physical label.",
						},
						"mark_connected": schema.BoolAttribute{
							Computed:    true,
							Description: "Treat as if a cable is connected.",
						},
						"description": schema.StringAttribute{
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

func (d *deviceConsoleServerPortsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceConsoleServerPortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceConsoleServerPortsDataSourceModel

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
	var responseDTOs []*netboxapi.DeviceConsoleServerPortResponseDTO

	params := dcim.NewDcimConsoleServerPortsListParams()
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
		"device_id":              netboxapi.FilterInt64s(params.SetDeviceID),
		"device_id__n":           netboxapi.FilterInt64s(params.SetDeviceIDn),
		"type":                   netboxapi.FilterStrings(params.SetType),
		"type__n":                netboxapi.FilterStrings(params.SetTypen),
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
		"label":                  netboxapi.FilterStrings(params.SetLabel),
		"label__n":               netboxapi.FilterStrings(params.SetLabeln),
		"label__ic":              netboxapi.FilterStrings(params.SetLabelIc),
		"label__nic":             netboxapi.FilterStrings(params.SetLabelNic),
		"label__ie":              netboxapi.FilterStrings(params.SetLabelIe),
		"label__nie":             netboxapi.FilterStrings(params.SetLabelNie),
		"label__isw":             netboxapi.FilterStrings(params.SetLabelIsw),
		"label__nisw":            netboxapi.FilterStrings(params.SetLabelNisw),
		"label__iew":             netboxapi.FilterStrings(params.SetLabelIew),
		"label__niew":            netboxapi.FilterStrings(params.SetLabelNiew),
		"label__empty":           netboxapi.FilterBool(params.SetLabelEmpty),
		"label__regex":           netboxapi.FilterStrings(params.SetLabelRegex),
		"label__iregex":          netboxapi.FilterStrings(params.SetLabelIregex),
		"owner_id":               netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":            netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                    netboxapi.FilterStrings(params.SetTag),
		"tag__n":                 netboxapi.FilterStrings(params.SetTagn),
		"tag__any":               netboxapi.FilterStrings(params.SetTagAny),
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
		"device":                 netboxapi.FilterStrings(params.SetDevice),
		"device__n":              netboxapi.FilterStrings(params.SetDevicen),
		"device_role":            netboxapi.FilterStrings(params.SetDeviceRole),
		"device_role__n":         netboxapi.FilterStrings(params.SetDeviceRolen),
		"device_role_id":         netboxapi.FilterInt64s(params.SetDeviceRoleID),
		"device_role_id__n":      netboxapi.FilterInt64s(params.SetDeviceRoleIDn),
		"device_status":          netboxapi.FilterStrings(params.SetDeviceStatus),
		"device_status__n":       netboxapi.FilterStrings(params.SetDeviceStatusn),
		"device_status__ic":      netboxapi.FilterStrings(params.SetDeviceStatusIc),
		"device_status__nic":     netboxapi.FilterStrings(params.SetDeviceStatusNic),
		"device_status__ie":      netboxapi.FilterStrings(params.SetDeviceStatusIe),
		"device_status__nie":     netboxapi.FilterStrings(params.SetDeviceStatusNie),
		"device_status__isw":     netboxapi.FilterStrings(params.SetDeviceStatusIsw),
		"device_status__nisw":    netboxapi.FilterStrings(params.SetDeviceStatusNisw),
		"device_status__iew":     netboxapi.FilterStrings(params.SetDeviceStatusIew),
		"device_status__niew":    netboxapi.FilterStrings(params.SetDeviceStatusNiew),
		"device_status__empty":   netboxapi.FilterBool(params.SetDeviceStatusEmpty),
		"device_status__regex":   netboxapi.FilterStrings(params.SetDeviceStatusRegex),
		"device_status__iregex":  netboxapi.FilterStrings(params.SetDeviceStatusIregex),
		"device_type":            netboxapi.FilterStrings(params.SetDeviceType),
		"device_type__n":         netboxapi.FilterStrings(params.SetDeviceTypen),
		"device_type_id":         netboxapi.FilterInt64s(params.SetDeviceTypeID),
		"device_type_id__n":      netboxapi.FilterInt64s(params.SetDeviceTypeIDn),
		"last_updated":           netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":        netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":       netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":      netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":       netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":      netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":    netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":               netboxapi.FilterStrings(params.SetLocation),
		"location__n":            netboxapi.FilterStrings(params.SetLocationn),
		"location_id":            netboxapi.FilterInt64s(params.SetLocationID),
		"location_id__n":         netboxapi.FilterInt64s(params.SetLocationIDn),
		"mark_connected":         netboxapi.FilterBool(params.SetMarkConnected),
		"module_id":              netboxapi.FilterInt64s(params.SetModuleID),
		"module_id__n":           netboxapi.FilterInt64s(params.SetModuleIDn),
		"occupied":               netboxapi.FilterBool(params.SetOccupied),
		"owner":                  netboxapi.FilterStrings(params.SetOwner),
		"owner__n":               netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":            netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":         netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":         netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":      netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                      netboxapi.FilterString(params.SetQ),
		"rack":                   netboxapi.FilterStrings(params.SetRack),
		"rack__n":                netboxapi.FilterStrings(params.SetRackn),
		"rack_id":                netboxapi.FilterInt64s(params.SetRackID),
		"rack_id__n":             netboxapi.FilterInt64s(params.SetRackIDn),
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
		"speed":                  netboxapi.FilterInt64(params.SetSpeed),
		"speed__n":               netboxapi.FilterInt64(params.SetSpeedn),
		"speed__ic":              netboxapi.FilterInt64s(params.SetSpeedIc),
		"speed__nic":             netboxapi.FilterInt64s(params.SetSpeedNic),
		"speed__ie":              netboxapi.FilterInt64s(params.SetSpeedIe),
		"speed__nie":             netboxapi.FilterInt64s(params.SetSpeedNie),
		"speed__isw":             netboxapi.FilterInt64s(params.SetSpeedIsw),
		"speed__nisw":            netboxapi.FilterInt64s(params.SetSpeedNisw),
		"speed__iew":             netboxapi.FilterInt64s(params.SetSpeedIew),
		"speed__niew":            netboxapi.FilterInt64s(params.SetSpeedNiew),
		"speed__regex":           netboxapi.FilterInt64s(params.SetSpeedRegex),
		"speed__iregex":          netboxapi.FilterInt64s(params.SetSpeedIregex),
		"tag_id":                 netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":              netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":            netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                 netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":              netboxapi.FilterStrings(params.SetTenantn),
		"tenant_id":              netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":           netboxapi.FilterInt64s(params.SetTenantIDn),
		"virtual_chassis":        netboxapi.FilterStrings(params.SetVirtualChassis),
		"virtual_chassis__n":     netboxapi.FilterStrings(params.SetVirtualChassisn),
		"virtual_chassis_id":     netboxapi.FilterInt64s(params.SetVirtualChassisID),
		"virtual_chassis_id__n":  netboxapi.FilterInt64s(params.SetVirtualChassisIDn),
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
		res, err := d.client.Dcim.DcimConsoleServerPortsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_device_console_server_ports", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.DeviceConsoleServerPortResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]deviceConsoleServerPortResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model deviceConsoleServerPortResourceModel
		resp.Diagnostics.Append(flattenDeviceConsoleServerPort(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: deviceConsoleServerPortResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
