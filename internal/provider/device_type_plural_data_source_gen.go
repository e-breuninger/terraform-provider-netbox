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
	_ datasource.DataSource              = (*deviceTypesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceTypesDataSource)(nil)
)

// NewDeviceTypesDataSource returns a new device_types data source, which
// lists device_type objects matching its filters.
func NewDeviceTypesDataSource() datasource.DataSource {
	return &deviceTypesDataSource{}
}

type deviceTypesDataSource struct {
	client *netboxapi.Client
}

// deviceTypesDataSourceModel is the data source model: the filters, the limit and the matching
// device_types.
type deviceTypesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"device_types"`
}

// deviceTypeResourceAttrTypes is the attribute type map of deviceTypeResourceModel.
func deviceTypeResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                                 types.Int64Type,
		"manufacturer_id":                    types.Int64Type,
		"model":                              types.StringType,
		"slug":                               types.StringType,
		"part_number":                        types.StringType,
		"default_platform_id":                types.Int64Type,
		"u_height":                           types.Float64Type,
		"is_full_depth":                      types.BoolType,
		"exclude_from_utilization":           types.BoolType,
		"subdevice_role":                     types.StringType,
		"airflow":                            types.StringType,
		"weight":                             types.Float64Type,
		"weight_unit":                        types.StringType,
		"description":                        types.StringType,
		"comments":                           types.StringType,
		"owner_id":                           types.Int64Type,
		"created":                            types.StringType,
		"last_updated":                       types.StringType,
		"url":                                types.StringType,
		"device_count":                       types.Int64Type,
		"console_port_template_count":        types.Int64Type,
		"console_server_port_template_count": types.Int64Type,
		"power_port_template_count":          types.Int64Type,
		"power_outlet_template_count":        types.Int64Type,
		"interface_template_count":           types.Int64Type,
		"front_port_template_count":          types.Int64Type,
		"rear_port_template_count":           types.Int64Type,
		"device_bay_template_count":          types.Int64Type,
		"module_bay_template_count":          types.Int64Type,
		"inventory_item_template_count":      types.Int64Type,
		"tags":                               types.SetType{ElemType: types.StringType},
		"tags_all":                           types.SetType{ElemType: types.StringType},
		"custom_fields":                      types.MapType{ElemType: types.StringType},
	}
}

func (d *deviceTypesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_types"
}

func (d *deviceTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists device_type objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: airflow, airflow__empty, airflow__ic, airflow__ie, airflow__iew, airflow__iregex, airflow__isw, airflow__n, airflow__nic, airflow__nie, airflow__niew, airflow__nisw, airflow__regex, console_port_template_count, console_port_template_count__empty, console_port_template_count__gt, console_port_template_count__gte, console_port_template_count__lt, console_port_template_count__lte, console_port_template_count__n, console_ports, console_server_port_template_count, console_server_port_template_count__empty, console_server_port_template_count__gt, console_server_port_template_count__gte, console_server_port_template_count__lt, console_server_port_template_count__lte, console_server_port_template_count__n, console_server_ports, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, default_platform, default_platform__n, default_platform_id, default_platform_id__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device_bay_template_count, device_bay_template_count__empty, device_bay_template_count__gt, device_bay_template_count__gte, device_bay_template_count__lt, device_bay_template_count__lte, device_bay_template_count__n, device_bays, device_count, device_count__empty, device_count__gt, device_count__gte, device_count__lt, device_count__lte, device_count__n, exclude_from_utilization, front_port_template_count, front_port_template_count__empty, front_port_template_count__gt, front_port_template_count__gte, front_port_template_count__lt, front_port_template_count__lte, front_port_template_count__n, has_front_image, has_rear_image, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_template_count, interface_template_count__empty, interface_template_count__gt, interface_template_count__gte, interface_template_count__lt, interface_template_count__lte, interface_template_count__n, interfaces, inventory_item_template_count, inventory_item_template_count__empty, inventory_item_template_count__gt, inventory_item_template_count__gte, inventory_item_template_count__lt, inventory_item_template_count__lte, inventory_item_template_count__n, inventory_items, is_full_depth, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, model, model__empty, model__ic, model__ie, model__iew, model__iregex, model__isw, model__n, model__nic, model__nie, model__niew, model__nisw, model__regex, module_bay_template_count, module_bay_template_count__empty, module_bay_template_count__gt, module_bay_template_count__gte, module_bay_template_count__lt, module_bay_template_count__lte, module_bay_template_count__n, module_bays, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, part_number, part_number__empty, part_number__ic, part_number__ie, part_number__iew, part_number__iregex, part_number__isw, part_number__n, part_number__nic, part_number__nie, part_number__niew, part_number__nisw, part_number__regex, pass_through_ports, power_outlet_template_count, power_outlet_template_count__empty, power_outlet_template_count__gt, power_outlet_template_count__gte, power_outlet_template_count__lt, power_outlet_template_count__lte, power_outlet_template_count__n, power_outlets, power_port_template_count, power_port_template_count__empty, power_port_template_count__gt, power_port_template_count__gte, power_port_template_count__lt, power_port_template_count__lte, power_port_template_count__n, power_ports, q, rear_port_template_count, rear_port_template_count__empty, rear_port_template_count__gt, rear_port_template_count__gte, rear_port_template_count__lt, rear_port_template_count__lte, rear_port_template_count__n, slug, slug__empty, slug__ic, slug__ie, slug__iew, slug__iregex, slug__isw, slug__n, slug__nic, slug__nie, slug__niew, slug__nisw, slug__regex, subdevice_role, subdevice_role__empty, subdevice_role__ic, subdevice_role__ie, subdevice_role__iew, subdevice_role__iregex, subdevice_role__isw, subdevice_role__n, subdevice_role__nic, subdevice_role__nie, subdevice_role__niew, subdevice_role__nisw, subdevice_role__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, u_height, u_height__empty, u_height__gt, u_height__gte, u_height__lt, u_height__lte, u_height__n, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n, weight_unit, weight_unit__empty, weight_unit__ic, weight_unit__ie, weight_unit__iew, weight_unit__iregex, weight_unit__isw, weight_unit__n, weight_unit__nic, weight_unit__nie, weight_unit__niew, weight_unit__nisw, weight_unit__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"device_types": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching device_types.",
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
							Computed:    true,
							Description: "Model name (unique per manufacturer).",
						},
						"slug": schema.StringAttribute{
							Computed:    true,
							Description: "URL-friendly unique shorthand; derived from model when not set.",
						},
						"part_number": schema.StringAttribute{
							Computed:    true,
							Description: "Discrete part number.",
						},
						"default_platform_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the platform assigned to new devices of this type.",
						},
						"u_height": schema.Float64Attribute{
							Computed:    true,
							Description: "Height in rack units (0.5 steps).",
						},
						"is_full_depth": schema.BoolAttribute{
							Computed:    true,
							Description: "Device consumes both the front and rear rack faces.",
						},
						"exclude_from_utilization": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether devices of this type are left out of the rack utilization.",
						},
						"subdevice_role": schema.StringAttribute{
							Computed:    true,
							Description: "Parent devices house child devices in device bays. One of: parent, child.",
						},
						"airflow": schema.StringAttribute{
							Computed:    true,
							Description: "Airflow direction. One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive, mixed, rear-to-side, bottom-to-top, top-to-bottom.",
						},
						"weight": schema.Float64Attribute{
							Computed:    true,
							Description: "Weight of the device type (requires weight_unit).",
						},
						"weight_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of weight. One of: kg, g, lb, oz.",
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
							Description: "Number of devices of the device type.",
						},
						"console_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of console port templates of the device type.",
						},
						"console_server_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of console server port templates of the device type.",
						},
						"power_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of power port templates of the device type.",
						},
						"power_outlet_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of power outlet templates of the device type.",
						},
						"interface_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of interface templates of the device type.",
						},
						"front_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of front port templates of the device type.",
						},
						"rear_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of rear port templates of the device type.",
						},
						"device_bay_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of device bay templates of the device type.",
						},
						"module_bay_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of module bay templates of the device type.",
						},
						"inventory_item_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of inventory item templates of the device type.",
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

func (d *deviceTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceTypesDataSourceModel

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
	var responseDTOs []*netboxapi.DeviceTypeResponseDTO

	params := dcim.NewDcimDeviceTypesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                                     netboxapi.FilterInt64s(params.SetID),
		"id__n":                                  netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                                 netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                                netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                                 netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                                netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                              netboxapi.FilterBool(params.SetIDEmpty),
		"model":                                  netboxapi.FilterStrings(params.SetModel),
		"model__n":                               netboxapi.FilterStrings(params.SetModeln),
		"model__ic":                              netboxapi.FilterStrings(params.SetModelIc),
		"model__nic":                             netboxapi.FilterStrings(params.SetModelNic),
		"model__ie":                              netboxapi.FilterStrings(params.SetModelIe),
		"model__nie":                             netboxapi.FilterStrings(params.SetModelNie),
		"model__isw":                             netboxapi.FilterStrings(params.SetModelIsw),
		"model__nisw":                            netboxapi.FilterStrings(params.SetModelNisw),
		"model__iew":                             netboxapi.FilterStrings(params.SetModelIew),
		"model__niew":                            netboxapi.FilterStrings(params.SetModelNiew),
		"model__empty":                           netboxapi.FilterBool(params.SetModelEmpty),
		"model__regex":                           netboxapi.FilterStrings(params.SetModelRegex),
		"model__iregex":                          netboxapi.FilterStrings(params.SetModelIregex),
		"slug":                                   netboxapi.FilterStrings(params.SetSlug),
		"slug__n":                                netboxapi.FilterStrings(params.SetSlugn),
		"slug__ic":                               netboxapi.FilterStrings(params.SetSlugIc),
		"slug__nic":                              netboxapi.FilterStrings(params.SetSlugNic),
		"slug__ie":                               netboxapi.FilterStrings(params.SetSlugIe),
		"slug__nie":                              netboxapi.FilterStrings(params.SetSlugNie),
		"slug__isw":                              netboxapi.FilterStrings(params.SetSlugIsw),
		"slug__nisw":                             netboxapi.FilterStrings(params.SetSlugNisw),
		"slug__iew":                              netboxapi.FilterStrings(params.SetSlugIew),
		"slug__niew":                             netboxapi.FilterStrings(params.SetSlugNiew),
		"slug__empty":                            netboxapi.FilterBool(params.SetSlugEmpty),
		"slug__regex":                            netboxapi.FilterStrings(params.SetSlugRegex),
		"slug__iregex":                           netboxapi.FilterStrings(params.SetSlugIregex),
		"owner_id":                               netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                            netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                                    netboxapi.FilterStrings(params.SetTag),
		"tag__n":                                 netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                               netboxapi.FilterStrings(params.SetTagAny),
		"airflow":                                netboxapi.FilterString(params.SetAirflow),
		"airflow__n":                             netboxapi.FilterString(params.SetAirflown),
		"airflow__ic":                            netboxapi.FilterStrings(params.SetAirflowIc),
		"airflow__nic":                           netboxapi.FilterStrings(params.SetAirflowNic),
		"airflow__ie":                            netboxapi.FilterStrings(params.SetAirflowIe),
		"airflow__nie":                           netboxapi.FilterStrings(params.SetAirflowNie),
		"airflow__isw":                           netboxapi.FilterStrings(params.SetAirflowIsw),
		"airflow__nisw":                          netboxapi.FilterStrings(params.SetAirflowNisw),
		"airflow__iew":                           netboxapi.FilterStrings(params.SetAirflowIew),
		"airflow__niew":                          netboxapi.FilterStrings(params.SetAirflowNiew),
		"airflow__empty":                         netboxapi.FilterBool(params.SetAirflowEmpty),
		"airflow__regex":                         netboxapi.FilterStrings(params.SetAirflowRegex),
		"airflow__iregex":                        netboxapi.FilterStrings(params.SetAirflowIregex),
		"console_port_template_count":            netboxapi.FilterInt64s(params.SetConsolePortTemplateCount),
		"console_port_template_count__n":         netboxapi.FilterInt64s(params.SetConsolePortTemplateCountn),
		"console_port_template_count__lt":        netboxapi.FilterInt64s(params.SetConsolePortTemplateCountLt),
		"console_port_template_count__lte":       netboxapi.FilterInt64s(params.SetConsolePortTemplateCountLte),
		"console_port_template_count__gt":        netboxapi.FilterInt64s(params.SetConsolePortTemplateCountGt),
		"console_port_template_count__gte":       netboxapi.FilterInt64s(params.SetConsolePortTemplateCountGte),
		"console_port_template_count__empty":     netboxapi.FilterBool(params.SetConsolePortTemplateCountEmpty),
		"console_ports":                          netboxapi.FilterBool(params.SetConsolePorts),
		"console_server_port_template_count":     netboxapi.FilterInt64s(params.SetConsoleServerPortTemplateCount),
		"console_server_port_template_count__n":  netboxapi.FilterInt64s(params.SetConsoleServerPortTemplateCountn),
		"console_server_port_template_count__lt": netboxapi.FilterInt64s(params.SetConsoleServerPortTemplateCountLt),
		"console_server_port_template_count__lte":   netboxapi.FilterInt64s(params.SetConsoleServerPortTemplateCountLte),
		"console_server_port_template_count__gt":    netboxapi.FilterInt64s(params.SetConsoleServerPortTemplateCountGt),
		"console_server_port_template_count__gte":   netboxapi.FilterInt64s(params.SetConsoleServerPortTemplateCountGte),
		"console_server_port_template_count__empty": netboxapi.FilterBool(params.SetConsoleServerPortTemplateCountEmpty),
		"console_server_ports":                      netboxapi.FilterBool(params.SetConsoleServerPorts),
		"created":                                   netboxapi.FilterStrings(params.SetCreated),
		"created__n":                                netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                               netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":                              netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                               netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":                              netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":                            netboxapi.FilterStrings(params.SetCreatedEmpty),
		"default_platform":                          netboxapi.FilterStrings(params.SetDefaultPlatform),
		"default_platform__n":                       netboxapi.FilterStrings(params.SetDefaultPlatformn),
		"default_platform_id":                       netboxapi.FilterStrings(params.SetDefaultPlatformID),
		"default_platform_id__n":                    netboxapi.FilterStrings(params.SetDefaultPlatformIDn),
		"description":                               netboxapi.FilterStrings(params.SetDescription),
		"description__n":                            netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":                           netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":                          netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":                           netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":                          netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":                          netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":                         netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":                          netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":                         netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":                        netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":                        netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":                       netboxapi.FilterStrings(params.SetDescriptionIregex),
		"device_bay_template_count":                 netboxapi.FilterInt64s(params.SetDeviceBayTemplateCount),
		"device_bay_template_count__n":              netboxapi.FilterInt64s(params.SetDeviceBayTemplateCountn),
		"device_bay_template_count__lt":             netboxapi.FilterInt64s(params.SetDeviceBayTemplateCountLt),
		"device_bay_template_count__lte":            netboxapi.FilterInt64s(params.SetDeviceBayTemplateCountLte),
		"device_bay_template_count__gt":             netboxapi.FilterInt64s(params.SetDeviceBayTemplateCountGt),
		"device_bay_template_count__gte":            netboxapi.FilterInt64s(params.SetDeviceBayTemplateCountGte),
		"device_bay_template_count__empty":          netboxapi.FilterBool(params.SetDeviceBayTemplateCountEmpty),
		"device_bays":                               netboxapi.FilterBool(params.SetDeviceBays),
		"device_count":                              netboxapi.FilterInt64s(params.SetDeviceCount),
		"device_count__n":                           netboxapi.FilterInt64s(params.SetDeviceCountn),
		"device_count__lt":                          netboxapi.FilterInt64s(params.SetDeviceCountLt),
		"device_count__lte":                         netboxapi.FilterInt64s(params.SetDeviceCountLte),
		"device_count__gt":                          netboxapi.FilterInt64s(params.SetDeviceCountGt),
		"device_count__gte":                         netboxapi.FilterInt64s(params.SetDeviceCountGte),
		"device_count__empty":                       netboxapi.FilterBool(params.SetDeviceCountEmpty),
		"exclude_from_utilization":                  netboxapi.FilterBool(params.SetExcludeFromUtilization),
		"front_port_template_count":                 netboxapi.FilterInt64s(params.SetFrontPortTemplateCount),
		"front_port_template_count__n":              netboxapi.FilterInt64s(params.SetFrontPortTemplateCountn),
		"front_port_template_count__lt":             netboxapi.FilterInt64s(params.SetFrontPortTemplateCountLt),
		"front_port_template_count__lte":            netboxapi.FilterInt64s(params.SetFrontPortTemplateCountLte),
		"front_port_template_count__gt":             netboxapi.FilterInt64s(params.SetFrontPortTemplateCountGt),
		"front_port_template_count__gte":            netboxapi.FilterInt64s(params.SetFrontPortTemplateCountGte),
		"front_port_template_count__empty":          netboxapi.FilterBool(params.SetFrontPortTemplateCountEmpty),
		"has_front_image":                           netboxapi.FilterBool(params.SetHasFrontImage),
		"has_rear_image":                            netboxapi.FilterBool(params.SetHasRearImage),
		"interface_template_count":                  netboxapi.FilterInt64s(params.SetInterfaceTemplateCount),
		"interface_template_count__n":               netboxapi.FilterInt64s(params.SetInterfaceTemplateCountn),
		"interface_template_count__lt":              netboxapi.FilterInt64s(params.SetInterfaceTemplateCountLt),
		"interface_template_count__lte":             netboxapi.FilterInt64s(params.SetInterfaceTemplateCountLte),
		"interface_template_count__gt":              netboxapi.FilterInt64s(params.SetInterfaceTemplateCountGt),
		"interface_template_count__gte":             netboxapi.FilterInt64s(params.SetInterfaceTemplateCountGte),
		"interface_template_count__empty":           netboxapi.FilterBool(params.SetInterfaceTemplateCountEmpty),
		"interfaces":                                netboxapi.FilterBool(params.SetInterfaces),
		"inventory_item_template_count":             netboxapi.FilterInt64s(params.SetInventoryItemTemplateCount),
		"inventory_item_template_count__n":          netboxapi.FilterInt64s(params.SetInventoryItemTemplateCountn),
		"inventory_item_template_count__lt":         netboxapi.FilterInt64s(params.SetInventoryItemTemplateCountLt),
		"inventory_item_template_count__lte":        netboxapi.FilterInt64s(params.SetInventoryItemTemplateCountLte),
		"inventory_item_template_count__gt":         netboxapi.FilterInt64s(params.SetInventoryItemTemplateCountGt),
		"inventory_item_template_count__gte":        netboxapi.FilterInt64s(params.SetInventoryItemTemplateCountGte),
		"inventory_item_template_count__empty":      netboxapi.FilterBool(params.SetInventoryItemTemplateCountEmpty),
		"inventory_items":                           netboxapi.FilterBool(params.SetInventoryItems),
		"is_full_depth":                             netboxapi.FilterBool(params.SetIsFullDepth),
		"last_updated":                              netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":                           netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":                          netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":                         netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":                          netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":                         netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":                       netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"manufacturer":                              netboxapi.FilterStrings(params.SetManufacturer),
		"manufacturer__n":                           netboxapi.FilterStrings(params.SetManufacturern),
		"manufacturer_id":                           netboxapi.FilterInt64s(params.SetManufacturerID),
		"manufacturer_id__n":                        netboxapi.FilterInt64s(params.SetManufacturerIDn),
		"module_bay_template_count":                 netboxapi.FilterInt64s(params.SetModuleBayTemplateCount),
		"module_bay_template_count__n":              netboxapi.FilterInt64s(params.SetModuleBayTemplateCountn),
		"module_bay_template_count__lt":             netboxapi.FilterInt64s(params.SetModuleBayTemplateCountLt),
		"module_bay_template_count__lte":            netboxapi.FilterInt64s(params.SetModuleBayTemplateCountLte),
		"module_bay_template_count__gt":             netboxapi.FilterInt64s(params.SetModuleBayTemplateCountGt),
		"module_bay_template_count__gte":            netboxapi.FilterInt64s(params.SetModuleBayTemplateCountGte),
		"module_bay_template_count__empty":          netboxapi.FilterBool(params.SetModuleBayTemplateCountEmpty),
		"module_bays":                               netboxapi.FilterBool(params.SetModuleBays),
		"owner":                                     netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                                  netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                               netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":                            netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":                            netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":                         netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"part_number":                               netboxapi.FilterStrings(params.SetPartNumber),
		"part_number__n":                            netboxapi.FilterStrings(params.SetPartNumbern),
		"part_number__ic":                           netboxapi.FilterStrings(params.SetPartNumberIc),
		"part_number__nic":                          netboxapi.FilterStrings(params.SetPartNumberNic),
		"part_number__ie":                           netboxapi.FilterStrings(params.SetPartNumberIe),
		"part_number__nie":                          netboxapi.FilterStrings(params.SetPartNumberNie),
		"part_number__isw":                          netboxapi.FilterStrings(params.SetPartNumberIsw),
		"part_number__nisw":                         netboxapi.FilterStrings(params.SetPartNumberNisw),
		"part_number__iew":                          netboxapi.FilterStrings(params.SetPartNumberIew),
		"part_number__niew":                         netboxapi.FilterStrings(params.SetPartNumberNiew),
		"part_number__empty":                        netboxapi.FilterBool(params.SetPartNumberEmpty),
		"part_number__regex":                        netboxapi.FilterStrings(params.SetPartNumberRegex),
		"part_number__iregex":                       netboxapi.FilterStrings(params.SetPartNumberIregex),
		"pass_through_ports":                        netboxapi.FilterBool(params.SetPassThroughPorts),
		"power_outlet_template_count":               netboxapi.FilterInt64s(params.SetPowerOutletTemplateCount),
		"power_outlet_template_count__n":            netboxapi.FilterInt64s(params.SetPowerOutletTemplateCountn),
		"power_outlet_template_count__lt":           netboxapi.FilterInt64s(params.SetPowerOutletTemplateCountLt),
		"power_outlet_template_count__lte":          netboxapi.FilterInt64s(params.SetPowerOutletTemplateCountLte),
		"power_outlet_template_count__gt":           netboxapi.FilterInt64s(params.SetPowerOutletTemplateCountGt),
		"power_outlet_template_count__gte":          netboxapi.FilterInt64s(params.SetPowerOutletTemplateCountGte),
		"power_outlet_template_count__empty":        netboxapi.FilterBool(params.SetPowerOutletTemplateCountEmpty),
		"power_outlets":                             netboxapi.FilterBool(params.SetPowerOutlets),
		"power_port_template_count":                 netboxapi.FilterInt64s(params.SetPowerPortTemplateCount),
		"power_port_template_count__n":              netboxapi.FilterInt64s(params.SetPowerPortTemplateCountn),
		"power_port_template_count__lt":             netboxapi.FilterInt64s(params.SetPowerPortTemplateCountLt),
		"power_port_template_count__lte":            netboxapi.FilterInt64s(params.SetPowerPortTemplateCountLte),
		"power_port_template_count__gt":             netboxapi.FilterInt64s(params.SetPowerPortTemplateCountGt),
		"power_port_template_count__gte":            netboxapi.FilterInt64s(params.SetPowerPortTemplateCountGte),
		"power_port_template_count__empty":          netboxapi.FilterBool(params.SetPowerPortTemplateCountEmpty),
		"power_ports":                               netboxapi.FilterBool(params.SetPowerPorts),
		"q":                                         netboxapi.FilterString(params.SetQ),
		"rear_port_template_count":                  netboxapi.FilterInt64s(params.SetRearPortTemplateCount),
		"rear_port_template_count__n":               netboxapi.FilterInt64s(params.SetRearPortTemplateCountn),
		"rear_port_template_count__lt":              netboxapi.FilterInt64s(params.SetRearPortTemplateCountLt),
		"rear_port_template_count__lte":             netboxapi.FilterInt64s(params.SetRearPortTemplateCountLte),
		"rear_port_template_count__gt":              netboxapi.FilterInt64s(params.SetRearPortTemplateCountGt),
		"rear_port_template_count__gte":             netboxapi.FilterInt64s(params.SetRearPortTemplateCountGte),
		"rear_port_template_count__empty":           netboxapi.FilterBool(params.SetRearPortTemplateCountEmpty),
		"subdevice_role":                            netboxapi.FilterString(params.SetSubdeviceRole),
		"subdevice_role__n":                         netboxapi.FilterString(params.SetSubdeviceRolen),
		"subdevice_role__ic":                        netboxapi.FilterStrings(params.SetSubdeviceRoleIc),
		"subdevice_role__nic":                       netboxapi.FilterStrings(params.SetSubdeviceRoleNic),
		"subdevice_role__ie":                        netboxapi.FilterStrings(params.SetSubdeviceRoleIe),
		"subdevice_role__nie":                       netboxapi.FilterStrings(params.SetSubdeviceRoleNie),
		"subdevice_role__isw":                       netboxapi.FilterStrings(params.SetSubdeviceRoleIsw),
		"subdevice_role__nisw":                      netboxapi.FilterStrings(params.SetSubdeviceRoleNisw),
		"subdevice_role__iew":                       netboxapi.FilterStrings(params.SetSubdeviceRoleIew),
		"subdevice_role__niew":                      netboxapi.FilterStrings(params.SetSubdeviceRoleNiew),
		"subdevice_role__empty":                     netboxapi.FilterBool(params.SetSubdeviceRoleEmpty),
		"subdevice_role__regex":                     netboxapi.FilterStrings(params.SetSubdeviceRoleRegex),
		"subdevice_role__iregex":                    netboxapi.FilterStrings(params.SetSubdeviceRoleIregex),
		"tag_id":                                    netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                                 netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                               netboxapi.FilterInt64s(params.SetTagIDAny),
		"u_height":                                  netboxapi.FilterFloat64s(params.SetUHeight),
		"u_height__n":                               netboxapi.FilterFloat64s(params.SetUHeightn),
		"u_height__lt":                              netboxapi.FilterFloat64s(params.SetUHeightLt),
		"u_height__lte":                             netboxapi.FilterFloat64s(params.SetUHeightLte),
		"u_height__gt":                              netboxapi.FilterFloat64s(params.SetUHeightGt),
		"u_height__gte":                             netboxapi.FilterFloat64s(params.SetUHeightGte),
		"u_height__empty":                           netboxapi.FilterBool(params.SetUHeightEmpty),
		"weight":                                    netboxapi.FilterFloat64s(params.SetWeight),
		"weight__n":                                 netboxapi.FilterFloat64s(params.SetWeightn),
		"weight__lt":                                netboxapi.FilterFloat64s(params.SetWeightLt),
		"weight__lte":                               netboxapi.FilterFloat64s(params.SetWeightLte),
		"weight__gt":                                netboxapi.FilterFloat64s(params.SetWeightGt),
		"weight__gte":                               netboxapi.FilterFloat64s(params.SetWeightGte),
		"weight__empty":                             netboxapi.FilterBool(params.SetWeightEmpty),
		"weight_unit":                               netboxapi.FilterString(params.SetWeightUnit),
		"weight_unit__n":                            netboxapi.FilterString(params.SetWeightUnitn),
		"weight_unit__ic":                           netboxapi.FilterStrings(params.SetWeightUnitIc),
		"weight_unit__nic":                          netboxapi.FilterStrings(params.SetWeightUnitNic),
		"weight_unit__ie":                           netboxapi.FilterStrings(params.SetWeightUnitIe),
		"weight_unit__nie":                          netboxapi.FilterStrings(params.SetWeightUnitNie),
		"weight_unit__isw":                          netboxapi.FilterStrings(params.SetWeightUnitIsw),
		"weight_unit__nisw":                         netboxapi.FilterStrings(params.SetWeightUnitNisw),
		"weight_unit__iew":                          netboxapi.FilterStrings(params.SetWeightUnitIew),
		"weight_unit__niew":                         netboxapi.FilterStrings(params.SetWeightUnitNiew),
		"weight_unit__empty":                        netboxapi.FilterBool(params.SetWeightUnitEmpty),
		"weight_unit__regex":                        netboxapi.FilterStrings(params.SetWeightUnitRegex),
		"weight_unit__iregex":                       netboxapi.FilterStrings(params.SetWeightUnitIregex),
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
		res, err := d.client.Dcim.DcimDeviceTypesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_device_types", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.DeviceTypeResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]deviceTypeResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model deviceTypeResourceModel
		resp.Diagnostics.Append(flattenDeviceType(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: deviceTypeResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
