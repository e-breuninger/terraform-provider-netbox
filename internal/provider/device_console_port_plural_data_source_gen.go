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
	_ datasource.DataSource              = (*deviceConsolePortsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceConsolePortsDataSource)(nil)
)

// NewDeviceConsolePortsDataSource returns a new device_console_ports data source, which
// lists device_console_port objects matching its filters.
func NewDeviceConsolePortsDataSource() datasource.DataSource {
	return &deviceConsolePortsDataSource{}
}

type deviceConsolePortsDataSource struct {
	client *netboxapi.Client
}

// deviceConsolePortsDataSourceModel is the data source model: the filters, the limit and the matching
// device_console_ports.
type deviceConsolePortsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"device_console_ports"`
}

// deviceConsolePortResourceAttrTypes is the attribute type map of deviceConsolePortResourceModel.
func deviceConsolePortResourceAttrTypes() map[string]attr.Type {
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

func (d *deviceConsolePortsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_console_ports"
}

func (d *deviceConsolePortsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists device_console_port objects matching the given filters.",
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
			"device_console_ports": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching device_console_ports.",
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
							Description: "Port type as its NetBox slug, e.g. de-9, rj-45, usb-a (any of NetBox's console port type choices).",
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

func (d *deviceConsolePortsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceConsolePortsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceConsolePortsDataSourceModel

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
	var responseDTOs []*netboxapi.DeviceConsolePortResponseDTO

	params := dcim.NewDcimConsolePortsListParams()
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
	var vDeviceID []int64
	var vDeviceIDn []int64
	var vType []string
	var vTypen []string
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
	var vLabel []string
	var vLabeln []string
	var vLabelIc []string
	var vLabelNic []string
	var vLabelIe []string
	var vLabelNie []string
	var vLabelIsw []string
	var vLabelNisw []string
	var vLabelIew []string
	var vLabelNiew []string
	var vLabelEmpty *bool
	var vLabelRegex []string
	var vLabelIregex []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
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
	var vDevice []string
	var vDevicen []string
	var vDeviceRole []string
	var vDeviceRolen []string
	var vDeviceRoleID []int64
	var vDeviceRoleIDn []int64
	var vDeviceStatus []string
	var vDeviceStatusn []string
	var vDeviceStatusIc []string
	var vDeviceStatusNic []string
	var vDeviceStatusIe []string
	var vDeviceStatusNie []string
	var vDeviceStatusIsw []string
	var vDeviceStatusNisw []string
	var vDeviceStatusIew []string
	var vDeviceStatusNiew []string
	var vDeviceStatusEmpty *bool
	var vDeviceStatusRegex []string
	var vDeviceStatusIregex []string
	var vDeviceType []string
	var vDeviceTypen []string
	var vDeviceTypeID []int64
	var vDeviceTypeIDn []int64
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLocation []string
	var vLocationn []string
	var vLocationID []int64
	var vLocationIDn []int64
	var vMarkConnected *bool
	var vModuleID []int64
	var vModuleIDn []int64
	var vOccupied *bool
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vQ *string
	var vRack []string
	var vRackn []string
	var vRackID []int64
	var vRackIDn []int64
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
	var vSpeed *int64
	var vSpeedn *int64
	var vSpeedIc []int64
	var vSpeedNic []int64
	var vSpeedIe []int64
	var vSpeedNie []int64
	var vSpeedIsw []int64
	var vSpeedNisw []int64
	var vSpeedIew []int64
	var vSpeedNiew []int64
	var vSpeedRegex []int64
	var vSpeedIregex []int64
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantID []int64
	var vTenantIDn []int64
	var vVirtualChassis []string
	var vVirtualChassisn []string
	var vVirtualChassisID []int64
	var vVirtualChassisIDn []int64
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
		case "device_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id' takes an integer, got %q.", value))
				return
			}
			vDeviceID = append(vDeviceID, v)
		case "device_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceIDn = append(vDeviceIDn, v)
		case "type":
			v := value
			vType = append(vType, v)
		case "type__n":
			v := value
			vTypen = append(vTypen, v)
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
		case "label":
			v := value
			vLabel = append(vLabel, v)
		case "label__n":
			v := value
			vLabeln = append(vLabeln, v)
		case "label__ic":
			v := value
			vLabelIc = append(vLabelIc, v)
		case "label__nic":
			v := value
			vLabelNic = append(vLabelNic, v)
		case "label__ie":
			v := value
			vLabelIe = append(vLabelIe, v)
		case "label__nie":
			v := value
			vLabelNie = append(vLabelNie, v)
		case "label__isw":
			v := value
			vLabelIsw = append(vLabelIsw, v)
		case "label__nisw":
			v := value
			vLabelNisw = append(vLabelNisw, v)
		case "label__iew":
			v := value
			vLabelIew = append(vLabelIew, v)
		case "label__niew":
			v := value
			vLabelNiew = append(vLabelNiew, v)
		case "label__empty":
			if vLabelEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'label__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'label__empty' takes a boolean, got %q.", value))
				return
			}
			vLabelEmpty = &v
		case "label__regex":
			v := value
			vLabelRegex = append(vLabelRegex, v)
		case "label__iregex":
			v := value
			vLabelIregex = append(vLabelIregex, v)
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
		case "device":
			v := value
			vDevice = append(vDevice, v)
		case "device__n":
			v := value
			vDevicen = append(vDevicen, v)
		case "device_role":
			v := value
			vDeviceRole = append(vDeviceRole, v)
		case "device_role__n":
			v := value
			vDeviceRolen = append(vDeviceRolen, v)
		case "device_role_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_role_id' takes an integer, got %q.", value))
				return
			}
			vDeviceRoleID = append(vDeviceRoleID, v)
		case "device_role_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_role_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceRoleIDn = append(vDeviceRoleIDn, v)
		case "device_status":
			v := value
			vDeviceStatus = append(vDeviceStatus, v)
		case "device_status__n":
			v := value
			vDeviceStatusn = append(vDeviceStatusn, v)
		case "device_status__ic":
			v := value
			vDeviceStatusIc = append(vDeviceStatusIc, v)
		case "device_status__nic":
			v := value
			vDeviceStatusNic = append(vDeviceStatusNic, v)
		case "device_status__ie":
			v := value
			vDeviceStatusIe = append(vDeviceStatusIe, v)
		case "device_status__nie":
			v := value
			vDeviceStatusNie = append(vDeviceStatusNie, v)
		case "device_status__isw":
			v := value
			vDeviceStatusIsw = append(vDeviceStatusIsw, v)
		case "device_status__nisw":
			v := value
			vDeviceStatusNisw = append(vDeviceStatusNisw, v)
		case "device_status__iew":
			v := value
			vDeviceStatusIew = append(vDeviceStatusIew, v)
		case "device_status__niew":
			v := value
			vDeviceStatusNiew = append(vDeviceStatusNiew, v)
		case "device_status__empty":
			if vDeviceStatusEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'device_status__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_status__empty' takes a boolean, got %q.", value))
				return
			}
			vDeviceStatusEmpty = &v
		case "device_status__regex":
			v := value
			vDeviceStatusRegex = append(vDeviceStatusRegex, v)
		case "device_status__iregex":
			v := value
			vDeviceStatusIregex = append(vDeviceStatusIregex, v)
		case "device_type":
			v := value
			vDeviceType = append(vDeviceType, v)
		case "device_type__n":
			v := value
			vDeviceTypen = append(vDeviceTypen, v)
		case "device_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_type_id' takes an integer, got %q.", value))
				return
			}
			vDeviceTypeID = append(vDeviceTypeID, v)
		case "device_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_type_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceTypeIDn = append(vDeviceTypeIDn, v)
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
		case "location":
			v := value
			vLocation = append(vLocation, v)
		case "location__n":
			v := value
			vLocationn = append(vLocationn, v)
		case "location_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id' takes an integer, got %q.", value))
				return
			}
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id__n' takes an integer, got %q.", value))
				return
			}
			vLocationIDn = append(vLocationIDn, v)
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
		case "module_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_id' takes an integer, got %q.", value))
				return
			}
			vModuleID = append(vModuleID, v)
		case "module_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_id__n' takes an integer, got %q.", value))
				return
			}
			vModuleIDn = append(vModuleIDn, v)
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
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "rack":
			v := value
			vRack = append(vRack, v)
		case "rack__n":
			v := value
			vRackn = append(vRackn, v)
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
		case "speed":
			if vSpeed != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'speed' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed' takes an integer, got %q.", value))
				return
			}
			vSpeed = &v
		case "speed__n":
			if vSpeedn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'speed__n' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__n' takes an integer, got %q.", value))
				return
			}
			vSpeedn = &v
		case "speed__ic":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__ic' takes an integer, got %q.", value))
				return
			}
			vSpeedIc = append(vSpeedIc, v)
		case "speed__nic":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__nic' takes an integer, got %q.", value))
				return
			}
			vSpeedNic = append(vSpeedNic, v)
		case "speed__ie":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__ie' takes an integer, got %q.", value))
				return
			}
			vSpeedIe = append(vSpeedIe, v)
		case "speed__nie":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__nie' takes an integer, got %q.", value))
				return
			}
			vSpeedNie = append(vSpeedNie, v)
		case "speed__isw":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__isw' takes an integer, got %q.", value))
				return
			}
			vSpeedIsw = append(vSpeedIsw, v)
		case "speed__nisw":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__nisw' takes an integer, got %q.", value))
				return
			}
			vSpeedNisw = append(vSpeedNisw, v)
		case "speed__iew":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__iew' takes an integer, got %q.", value))
				return
			}
			vSpeedIew = append(vSpeedIew, v)
		case "speed__niew":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__niew' takes an integer, got %q.", value))
				return
			}
			vSpeedNiew = append(vSpeedNiew, v)
		case "speed__regex":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__regex' takes an integer, got %q.", value))
				return
			}
			vSpeedRegex = append(vSpeedRegex, v)
		case "speed__iregex":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'speed__iregex' takes an integer, got %q.", value))
				return
			}
			vSpeedIregex = append(vSpeedIregex, v)
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
		case "virtual_chassis":
			v := value
			vVirtualChassis = append(vVirtualChassis, v)
		case "virtual_chassis__n":
			v := value
			vVirtualChassisn = append(vVirtualChassisn, v)
		case "virtual_chassis_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_chassis_id' takes an integer, got %q.", value))
				return
			}
			vVirtualChassisID = append(vVirtualChassisID, v)
		case "virtual_chassis_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_chassis_id__n' takes an integer, got %q.", value))
				return
			}
			vVirtualChassisIDn = append(vVirtualChassisIDn, v)
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
	if len(vDeviceID) > 0 {
		params.SetDeviceID(vDeviceID)
	}
	if len(vDeviceIDn) > 0 {
		params.SetDeviceIDn(vDeviceIDn)
	}
	if len(vType) > 0 {
		params.SetType(vType)
	}
	if len(vTypen) > 0 {
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
	if len(vLabel) > 0 {
		params.SetLabel(vLabel)
	}
	if len(vLabeln) > 0 {
		params.SetLabeln(vLabeln)
	}
	if len(vLabelIc) > 0 {
		params.SetLabelIc(vLabelIc)
	}
	if len(vLabelNic) > 0 {
		params.SetLabelNic(vLabelNic)
	}
	if len(vLabelIe) > 0 {
		params.SetLabelIe(vLabelIe)
	}
	if len(vLabelNie) > 0 {
		params.SetLabelNie(vLabelNie)
	}
	if len(vLabelIsw) > 0 {
		params.SetLabelIsw(vLabelIsw)
	}
	if len(vLabelNisw) > 0 {
		params.SetLabelNisw(vLabelNisw)
	}
	if len(vLabelIew) > 0 {
		params.SetLabelIew(vLabelIew)
	}
	if len(vLabelNiew) > 0 {
		params.SetLabelNiew(vLabelNiew)
	}
	if vLabelEmpty != nil {
		params.SetLabelEmpty(vLabelEmpty)
	}
	if len(vLabelRegex) > 0 {
		params.SetLabelRegex(vLabelRegex)
	}
	if len(vLabelIregex) > 0 {
		params.SetLabelIregex(vLabelIregex)
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
	if len(vDevice) > 0 {
		params.SetDevice(vDevice)
	}
	if len(vDevicen) > 0 {
		params.SetDevicen(vDevicen)
	}
	if len(vDeviceRole) > 0 {
		params.SetDeviceRole(vDeviceRole)
	}
	if len(vDeviceRolen) > 0 {
		params.SetDeviceRolen(vDeviceRolen)
	}
	if len(vDeviceRoleID) > 0 {
		params.SetDeviceRoleID(vDeviceRoleID)
	}
	if len(vDeviceRoleIDn) > 0 {
		params.SetDeviceRoleIDn(vDeviceRoleIDn)
	}
	if len(vDeviceStatus) > 0 {
		params.SetDeviceStatus(vDeviceStatus)
	}
	if len(vDeviceStatusn) > 0 {
		params.SetDeviceStatusn(vDeviceStatusn)
	}
	if len(vDeviceStatusIc) > 0 {
		params.SetDeviceStatusIc(vDeviceStatusIc)
	}
	if len(vDeviceStatusNic) > 0 {
		params.SetDeviceStatusNic(vDeviceStatusNic)
	}
	if len(vDeviceStatusIe) > 0 {
		params.SetDeviceStatusIe(vDeviceStatusIe)
	}
	if len(vDeviceStatusNie) > 0 {
		params.SetDeviceStatusNie(vDeviceStatusNie)
	}
	if len(vDeviceStatusIsw) > 0 {
		params.SetDeviceStatusIsw(vDeviceStatusIsw)
	}
	if len(vDeviceStatusNisw) > 0 {
		params.SetDeviceStatusNisw(vDeviceStatusNisw)
	}
	if len(vDeviceStatusIew) > 0 {
		params.SetDeviceStatusIew(vDeviceStatusIew)
	}
	if len(vDeviceStatusNiew) > 0 {
		params.SetDeviceStatusNiew(vDeviceStatusNiew)
	}
	if vDeviceStatusEmpty != nil {
		params.SetDeviceStatusEmpty(vDeviceStatusEmpty)
	}
	if len(vDeviceStatusRegex) > 0 {
		params.SetDeviceStatusRegex(vDeviceStatusRegex)
	}
	if len(vDeviceStatusIregex) > 0 {
		params.SetDeviceStatusIregex(vDeviceStatusIregex)
	}
	if len(vDeviceType) > 0 {
		params.SetDeviceType(vDeviceType)
	}
	if len(vDeviceTypen) > 0 {
		params.SetDeviceTypen(vDeviceTypen)
	}
	if len(vDeviceTypeID) > 0 {
		params.SetDeviceTypeID(vDeviceTypeID)
	}
	if len(vDeviceTypeIDn) > 0 {
		params.SetDeviceTypeIDn(vDeviceTypeIDn)
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
	if len(vLocation) > 0 {
		params.SetLocation(vLocation)
	}
	if len(vLocationn) > 0 {
		params.SetLocationn(vLocationn)
	}
	if len(vLocationID) > 0 {
		params.SetLocationID(vLocationID)
	}
	if len(vLocationIDn) > 0 {
		params.SetLocationIDn(vLocationIDn)
	}
	if vMarkConnected != nil {
		params.SetMarkConnected(vMarkConnected)
	}
	if len(vModuleID) > 0 {
		params.SetModuleID(vModuleID)
	}
	if len(vModuleIDn) > 0 {
		params.SetModuleIDn(vModuleIDn)
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
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRack) > 0 {
		params.SetRack(vRack)
	}
	if len(vRackn) > 0 {
		params.SetRackn(vRackn)
	}
	if len(vRackID) > 0 {
		params.SetRackID(vRackID)
	}
	if len(vRackIDn) > 0 {
		params.SetRackIDn(vRackIDn)
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
	if vSpeed != nil {
		params.SetSpeed(vSpeed)
	}
	if vSpeedn != nil {
		params.SetSpeedn(vSpeedn)
	}
	if len(vSpeedIc) > 0 {
		params.SetSpeedIc(vSpeedIc)
	}
	if len(vSpeedNic) > 0 {
		params.SetSpeedNic(vSpeedNic)
	}
	if len(vSpeedIe) > 0 {
		params.SetSpeedIe(vSpeedIe)
	}
	if len(vSpeedNie) > 0 {
		params.SetSpeedNie(vSpeedNie)
	}
	if len(vSpeedIsw) > 0 {
		params.SetSpeedIsw(vSpeedIsw)
	}
	if len(vSpeedNisw) > 0 {
		params.SetSpeedNisw(vSpeedNisw)
	}
	if len(vSpeedIew) > 0 {
		params.SetSpeedIew(vSpeedIew)
	}
	if len(vSpeedNiew) > 0 {
		params.SetSpeedNiew(vSpeedNiew)
	}
	if len(vSpeedRegex) > 0 {
		params.SetSpeedRegex(vSpeedRegex)
	}
	if len(vSpeedIregex) > 0 {
		params.SetSpeedIregex(vSpeedIregex)
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
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vVirtualChassis) > 0 {
		params.SetVirtualChassis(vVirtualChassis)
	}
	if len(vVirtualChassisn) > 0 {
		params.SetVirtualChassisn(vVirtualChassisn)
	}
	if len(vVirtualChassisID) > 0 {
		params.SetVirtualChassisID(vVirtualChassisID)
	}
	if len(vVirtualChassisIDn) > 0 {
		params.SetVirtualChassisIDn(vVirtualChassisIDn)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Dcim.DcimConsolePortsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_device_console_ports", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.DeviceConsolePortResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]deviceConsolePortResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model deviceConsolePortResourceModel
		resp.Diagnostics.Append(flattenDeviceConsolePort(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: deviceConsolePortResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
