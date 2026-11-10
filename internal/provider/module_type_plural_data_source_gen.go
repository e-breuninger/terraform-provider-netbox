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
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*moduleTypesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*moduleTypesDataSource)(nil)
)

// NewModuleTypesDataSource returns a new module_types data source, which
// lists module_type objects matching its filters.
func NewModuleTypesDataSource() datasource.DataSource {
	return &moduleTypesDataSource{}
}

type moduleTypesDataSource struct {
	client *netboxapi.Client
}

// moduleTypesDataSourceModel is the data source model: the filters, the limit and the matching
// module_types.
type moduleTypesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"module_types"`
}

// moduleTypeResourceAttrTypes is the attribute type map of moduleTypeResourceModel.
func moduleTypeResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                                 types.Int64Type,
		"manufacturer_id":                    types.Int64Type,
		"model":                              types.StringType,
		"part_number":                        types.StringType,
		"profile_id":                         types.Int64Type,
		"attributes":                         jsontypes.NormalizedType{},
		"airflow":                            types.StringType,
		"weight":                             types.Float64Type,
		"weight_unit":                        types.StringType,
		"description":                        types.StringType,
		"comments":                           types.StringType,
		"owner_id":                           types.Int64Type,
		"created":                            types.StringType,
		"last_updated":                       types.StringType,
		"url":                                types.StringType,
		"module_count":                       types.Int64Type,
		"console_port_template_count":        types.Int64Type,
		"console_server_port_template_count": types.Int64Type,
		"power_port_template_count":          types.Int64Type,
		"power_outlet_template_count":        types.Int64Type,
		"interface_template_count":           types.Int64Type,
		"front_port_template_count":          types.Int64Type,
		"rear_port_template_count":           types.Int64Type,
		"module_bay_template_count":          types.Int64Type,
		"tags":                               types.SetType{ElemType: types.StringType},
		"tags_all":                           types.SetType{ElemType: types.StringType},
		"custom_fields":                      types.MapType{ElemType: types.StringType},
	}
}

func (d *moduleTypesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_module_types"
}

func (d *moduleTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists module_type objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: airflow, airflow__empty, airflow__ic, airflow__ie, airflow__iew, airflow__iregex, airflow__isw, airflow__n, airflow__nic, airflow__nie, airflow__niew, airflow__nisw, airflow__regex, console_port_template_count, console_port_template_count__empty, console_port_template_count__gt, console_port_template_count__gte, console_port_template_count__lt, console_port_template_count__lte, console_port_template_count__n, console_ports, console_server_port_template_count, console_server_port_template_count__empty, console_server_port_template_count__gt, console_server_port_template_count__gte, console_server_port_template_count__lt, console_server_port_template_count__lte, console_server_port_template_count__n, console_server_ports, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, front_port_template_count, front_port_template_count__empty, front_port_template_count__gt, front_port_template_count__gte, front_port_template_count__lt, front_port_template_count__lte, front_port_template_count__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_template_count, interface_template_count__empty, interface_template_count__gt, interface_template_count__gte, interface_template_count__lt, interface_template_count__lte, interface_template_count__n, interfaces, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, model, model__empty, model__ic, model__ie, model__iew, model__iregex, model__isw, model__n, model__nic, model__nie, model__niew, model__nisw, model__regex, module_bay_template_count, module_bay_template_count__empty, module_bay_template_count__gt, module_bay_template_count__gte, module_bay_template_count__lt, module_bay_template_count__lte, module_bay_template_count__n, module_bays, module_count, module_count__empty, module_count__gt, module_count__gte, module_count__lt, module_count__lte, module_count__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, part_number, part_number__empty, part_number__ic, part_number__ie, part_number__iew, part_number__iregex, part_number__isw, part_number__n, part_number__nic, part_number__nie, part_number__niew, part_number__nisw, part_number__regex, pass_through_ports, power_outlet_template_count, power_outlet_template_count__empty, power_outlet_template_count__gt, power_outlet_template_count__gte, power_outlet_template_count__lt, power_outlet_template_count__lte, power_outlet_template_count__n, power_outlets, power_port_template_count, power_port_template_count__empty, power_port_template_count__gt, power_port_template_count__gte, power_port_template_count__lt, power_port_template_count__lte, power_port_template_count__n, power_ports, profile, profile__n, profile_id, profile_id__n, q, rear_port_template_count, rear_port_template_count__empty, rear_port_template_count__gt, rear_port_template_count__gte, rear_port_template_count__lt, rear_port_template_count__lte, rear_port_template_count__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n, weight_unit, weight_unit__empty, weight_unit__ic, weight_unit__ie, weight_unit__iew, weight_unit__iregex, weight_unit__isw, weight_unit__n, weight_unit__nic, weight_unit__nie, weight_unit__niew, weight_unit__nisw, weight_unit__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"module_types": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching module_types.",
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
						"part_number": schema.StringAttribute{
							Computed:    true,
							Description: "Discrete part number.",
						},
						"profile_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the module type profile.",
						},
						"attributes": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Attribute values of the module type as JSON text, as the profile's schema defines them (use jsonencode()).",
						},
						"airflow": schema.StringAttribute{
							Computed:    true,
							Description: "Airflow direction. One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive.",
						},
						"weight": schema.Float64Attribute{
							Computed:    true,
							Description: "Weight of the module type (requires weight_unit).",
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
						"module_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of modules of the module type.",
						},
						"console_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of console port templates of the module type.",
						},
						"console_server_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of console server port templates of the module type.",
						},
						"power_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of power port templates of the module type.",
						},
						"power_outlet_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of power outlet templates of the module type.",
						},
						"interface_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of interface templates of the module type.",
						},
						"front_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of front port templates of the module type.",
						},
						"rear_port_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of rear port templates of the module type.",
						},
						"module_bay_template_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of module bay templates of the module type.",
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

func (d *moduleTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *moduleTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data moduleTypesDataSourceModel

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
	var responseDTOs []*netboxapi.ModuleTypeResponseDTO

	params := dcim.NewDcimModuleTypesListParams()
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
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAirflow *string
	var vAirflown *string
	var vAirflowIc []string
	var vAirflowNic []string
	var vAirflowIe []string
	var vAirflowNie []string
	var vAirflowIsw []string
	var vAirflowNisw []string
	var vAirflowIew []string
	var vAirflowNiew []string
	var vAirflowEmpty *bool
	var vAirflowRegex []string
	var vAirflowIregex []string
	var vConsolePortTemplateCount []int64
	var vConsolePortTemplateCountn []int64
	var vConsolePortTemplateCountLt []int64
	var vConsolePortTemplateCountLte []int64
	var vConsolePortTemplateCountGt []int64
	var vConsolePortTemplateCountGte []int64
	var vConsolePortTemplateCountEmpty *bool
	var vConsolePorts *bool
	var vConsoleServerPortTemplateCount []int64
	var vConsoleServerPortTemplateCountn []int64
	var vConsoleServerPortTemplateCountLt []int64
	var vConsoleServerPortTemplateCountLte []int64
	var vConsoleServerPortTemplateCountGt []int64
	var vConsoleServerPortTemplateCountGte []int64
	var vConsoleServerPortTemplateCountEmpty *bool
	var vConsoleServerPorts *bool
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
	var vFrontPortTemplateCount []int64
	var vFrontPortTemplateCountn []int64
	var vFrontPortTemplateCountLt []int64
	var vFrontPortTemplateCountLte []int64
	var vFrontPortTemplateCountGt []int64
	var vFrontPortTemplateCountGte []int64
	var vFrontPortTemplateCountEmpty *bool
	var vInterfaceTemplateCount []int64
	var vInterfaceTemplateCountn []int64
	var vInterfaceTemplateCountLt []int64
	var vInterfaceTemplateCountLte []int64
	var vInterfaceTemplateCountGt []int64
	var vInterfaceTemplateCountGte []int64
	var vInterfaceTemplateCountEmpty *bool
	var vInterfaces *bool
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vManufacturer []string
	var vManufacturern []string
	var vManufacturerID []int64
	var vManufacturerIDn []int64
	var vModuleBayTemplateCount []int64
	var vModuleBayTemplateCountn []int64
	var vModuleBayTemplateCountLt []int64
	var vModuleBayTemplateCountLte []int64
	var vModuleBayTemplateCountGt []int64
	var vModuleBayTemplateCountGte []int64
	var vModuleBayTemplateCountEmpty *bool
	var vModuleBays *bool
	var vModuleCount []int64
	var vModuleCountn []int64
	var vModuleCountLt []int64
	var vModuleCountLte []int64
	var vModuleCountGt []int64
	var vModuleCountGte []int64
	var vModuleCountEmpty *bool
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPartNumber []string
	var vPartNumbern []string
	var vPartNumberIc []string
	var vPartNumberNic []string
	var vPartNumberIe []string
	var vPartNumberNie []string
	var vPartNumberIsw []string
	var vPartNumberNisw []string
	var vPartNumberIew []string
	var vPartNumberNiew []string
	var vPartNumberEmpty *bool
	var vPartNumberRegex []string
	var vPartNumberIregex []string
	var vPassThroughPorts *bool
	var vPowerOutletTemplateCount []int64
	var vPowerOutletTemplateCountn []int64
	var vPowerOutletTemplateCountLt []int64
	var vPowerOutletTemplateCountLte []int64
	var vPowerOutletTemplateCountGt []int64
	var vPowerOutletTemplateCountGte []int64
	var vPowerOutletTemplateCountEmpty *bool
	var vPowerOutlets *bool
	var vPowerPortTemplateCount []int64
	var vPowerPortTemplateCountn []int64
	var vPowerPortTemplateCountLt []int64
	var vPowerPortTemplateCountLte []int64
	var vPowerPortTemplateCountGt []int64
	var vPowerPortTemplateCountGte []int64
	var vPowerPortTemplateCountEmpty *bool
	var vPowerPorts *bool
	var vProfile []string
	var vProfilen []string
	var vProfileID []int64
	var vProfileIDn []int64
	var vQ *string
	var vRearPortTemplateCount []int64
	var vRearPortTemplateCountn []int64
	var vRearPortTemplateCountLt []int64
	var vRearPortTemplateCountLte []int64
	var vRearPortTemplateCountGt []int64
	var vRearPortTemplateCountGte []int64
	var vRearPortTemplateCountEmpty *bool
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
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
		case "airflow":
			if vAirflow != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'airflow' takes a single value.")
				return
			}
			v := value
			vAirflow = &v
		case "airflow__n":
			if vAirflown != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'airflow__n' takes a single value.")
				return
			}
			v := value
			vAirflown = &v
		case "airflow__ic":
			v := value
			vAirflowIc = append(vAirflowIc, v)
		case "airflow__nic":
			v := value
			vAirflowNic = append(vAirflowNic, v)
		case "airflow__ie":
			v := value
			vAirflowIe = append(vAirflowIe, v)
		case "airflow__nie":
			v := value
			vAirflowNie = append(vAirflowNie, v)
		case "airflow__isw":
			v := value
			vAirflowIsw = append(vAirflowIsw, v)
		case "airflow__nisw":
			v := value
			vAirflowNisw = append(vAirflowNisw, v)
		case "airflow__iew":
			v := value
			vAirflowIew = append(vAirflowIew, v)
		case "airflow__niew":
			v := value
			vAirflowNiew = append(vAirflowNiew, v)
		case "airflow__empty":
			if vAirflowEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'airflow__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'airflow__empty' takes a boolean, got %q.", value))
				return
			}
			vAirflowEmpty = &v
		case "airflow__regex":
			v := value
			vAirflowRegex = append(vAirflowRegex, v)
		case "airflow__iregex":
			v := value
			vAirflowIregex = append(vAirflowIregex, v)
		case "console_port_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_template_count' takes an integer, got %q.", value))
				return
			}
			vConsolePortTemplateCount = append(vConsolePortTemplateCount, v)
		case "console_port_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_template_count__n' takes an integer, got %q.", value))
				return
			}
			vConsolePortTemplateCountn = append(vConsolePortTemplateCountn, v)
		case "console_port_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vConsolePortTemplateCountLt = append(vConsolePortTemplateCountLt, v)
		case "console_port_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vConsolePortTemplateCountLte = append(vConsolePortTemplateCountLte, v)
		case "console_port_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vConsolePortTemplateCountGt = append(vConsolePortTemplateCountGt, v)
		case "console_port_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vConsolePortTemplateCountGte = append(vConsolePortTemplateCountGte, v)
		case "console_port_template_count__empty":
			if vConsolePortTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_port_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_port_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vConsolePortTemplateCountEmpty = &v
		case "console_ports":
			if vConsolePorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_ports' takes a boolean, got %q.", value))
				return
			}
			vConsolePorts = &v
		case "console_server_port_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_template_count' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortTemplateCount = append(vConsoleServerPortTemplateCount, v)
		case "console_server_port_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_template_count__n' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortTemplateCountn = append(vConsoleServerPortTemplateCountn, v)
		case "console_server_port_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortTemplateCountLt = append(vConsoleServerPortTemplateCountLt, v)
		case "console_server_port_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortTemplateCountLte = append(vConsoleServerPortTemplateCountLte, v)
		case "console_server_port_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortTemplateCountGt = append(vConsoleServerPortTemplateCountGt, v)
		case "console_server_port_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vConsoleServerPortTemplateCountGte = append(vConsoleServerPortTemplateCountGte, v)
		case "console_server_port_template_count__empty":
			if vConsoleServerPortTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_server_port_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_port_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vConsoleServerPortTemplateCountEmpty = &v
		case "console_server_ports":
			if vConsoleServerPorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'console_server_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'console_server_ports' takes a boolean, got %q.", value))
				return
			}
			vConsoleServerPorts = &v
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
		case "front_port_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_template_count' takes an integer, got %q.", value))
				return
			}
			vFrontPortTemplateCount = append(vFrontPortTemplateCount, v)
		case "front_port_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_template_count__n' takes an integer, got %q.", value))
				return
			}
			vFrontPortTemplateCountn = append(vFrontPortTemplateCountn, v)
		case "front_port_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vFrontPortTemplateCountLt = append(vFrontPortTemplateCountLt, v)
		case "front_port_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vFrontPortTemplateCountLte = append(vFrontPortTemplateCountLte, v)
		case "front_port_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vFrontPortTemplateCountGt = append(vFrontPortTemplateCountGt, v)
		case "front_port_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vFrontPortTemplateCountGte = append(vFrontPortTemplateCountGte, v)
		case "front_port_template_count__empty":
			if vFrontPortTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'front_port_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'front_port_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vFrontPortTemplateCountEmpty = &v
		case "interface_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_template_count' takes an integer, got %q.", value))
				return
			}
			vInterfaceTemplateCount = append(vInterfaceTemplateCount, v)
		case "interface_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_template_count__n' takes an integer, got %q.", value))
				return
			}
			vInterfaceTemplateCountn = append(vInterfaceTemplateCountn, v)
		case "interface_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vInterfaceTemplateCountLt = append(vInterfaceTemplateCountLt, v)
		case "interface_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vInterfaceTemplateCountLte = append(vInterfaceTemplateCountLte, v)
		case "interface_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vInterfaceTemplateCountGt = append(vInterfaceTemplateCountGt, v)
		case "interface_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vInterfaceTemplateCountGte = append(vInterfaceTemplateCountGte, v)
		case "interface_template_count__empty":
			if vInterfaceTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'interface_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vInterfaceTemplateCountEmpty = &v
		case "interfaces":
			if vInterfaces != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'interfaces' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interfaces' takes a boolean, got %q.", value))
				return
			}
			vInterfaces = &v
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
		case "module_bay_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_template_count' takes an integer, got %q.", value))
				return
			}
			vModuleBayTemplateCount = append(vModuleBayTemplateCount, v)
		case "module_bay_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_template_count__n' takes an integer, got %q.", value))
				return
			}
			vModuleBayTemplateCountn = append(vModuleBayTemplateCountn, v)
		case "module_bay_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vModuleBayTemplateCountLt = append(vModuleBayTemplateCountLt, v)
		case "module_bay_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vModuleBayTemplateCountLte = append(vModuleBayTemplateCountLte, v)
		case "module_bay_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vModuleBayTemplateCountGt = append(vModuleBayTemplateCountGt, v)
		case "module_bay_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vModuleBayTemplateCountGte = append(vModuleBayTemplateCountGte, v)
		case "module_bay_template_count__empty":
			if vModuleBayTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'module_bay_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bay_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vModuleBayTemplateCountEmpty = &v
		case "module_bays":
			if vModuleBays != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'module_bays' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_bays' takes a boolean, got %q.", value))
				return
			}
			vModuleBays = &v
		case "module_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_count' takes an integer, got %q.", value))
				return
			}
			vModuleCount = append(vModuleCount, v)
		case "module_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_count__n' takes an integer, got %q.", value))
				return
			}
			vModuleCountn = append(vModuleCountn, v)
		case "module_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_count__lt' takes an integer, got %q.", value))
				return
			}
			vModuleCountLt = append(vModuleCountLt, v)
		case "module_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_count__lte' takes an integer, got %q.", value))
				return
			}
			vModuleCountLte = append(vModuleCountLte, v)
		case "module_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_count__gt' takes an integer, got %q.", value))
				return
			}
			vModuleCountGt = append(vModuleCountGt, v)
		case "module_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_count__gte' takes an integer, got %q.", value))
				return
			}
			vModuleCountGte = append(vModuleCountGte, v)
		case "module_count__empty":
			if vModuleCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'module_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_count__empty' takes a boolean, got %q.", value))
				return
			}
			vModuleCountEmpty = &v
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
		case "part_number":
			v := value
			vPartNumber = append(vPartNumber, v)
		case "part_number__n":
			v := value
			vPartNumbern = append(vPartNumbern, v)
		case "part_number__ic":
			v := value
			vPartNumberIc = append(vPartNumberIc, v)
		case "part_number__nic":
			v := value
			vPartNumberNic = append(vPartNumberNic, v)
		case "part_number__ie":
			v := value
			vPartNumberIe = append(vPartNumberIe, v)
		case "part_number__nie":
			v := value
			vPartNumberNie = append(vPartNumberNie, v)
		case "part_number__isw":
			v := value
			vPartNumberIsw = append(vPartNumberIsw, v)
		case "part_number__nisw":
			v := value
			vPartNumberNisw = append(vPartNumberNisw, v)
		case "part_number__iew":
			v := value
			vPartNumberIew = append(vPartNumberIew, v)
		case "part_number__niew":
			v := value
			vPartNumberNiew = append(vPartNumberNiew, v)
		case "part_number__empty":
			if vPartNumberEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'part_number__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'part_number__empty' takes a boolean, got %q.", value))
				return
			}
			vPartNumberEmpty = &v
		case "part_number__regex":
			v := value
			vPartNumberRegex = append(vPartNumberRegex, v)
		case "part_number__iregex":
			v := value
			vPartNumberIregex = append(vPartNumberIregex, v)
		case "pass_through_ports":
			if vPassThroughPorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'pass_through_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pass_through_ports' takes a boolean, got %q.", value))
				return
			}
			vPassThroughPorts = &v
		case "power_outlet_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_template_count' takes an integer, got %q.", value))
				return
			}
			vPowerOutletTemplateCount = append(vPowerOutletTemplateCount, v)
		case "power_outlet_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_template_count__n' takes an integer, got %q.", value))
				return
			}
			vPowerOutletTemplateCountn = append(vPowerOutletTemplateCountn, v)
		case "power_outlet_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vPowerOutletTemplateCountLt = append(vPowerOutletTemplateCountLt, v)
		case "power_outlet_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vPowerOutletTemplateCountLte = append(vPowerOutletTemplateCountLte, v)
		case "power_outlet_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vPowerOutletTemplateCountGt = append(vPowerOutletTemplateCountGt, v)
		case "power_outlet_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vPowerOutletTemplateCountGte = append(vPowerOutletTemplateCountGte, v)
		case "power_outlet_template_count__empty":
			if vPowerOutletTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_outlet_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlet_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vPowerOutletTemplateCountEmpty = &v
		case "power_outlets":
			if vPowerOutlets != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_outlets' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_outlets' takes a boolean, got %q.", value))
				return
			}
			vPowerOutlets = &v
		case "power_port_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_template_count' takes an integer, got %q.", value))
				return
			}
			vPowerPortTemplateCount = append(vPowerPortTemplateCount, v)
		case "power_port_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_template_count__n' takes an integer, got %q.", value))
				return
			}
			vPowerPortTemplateCountn = append(vPowerPortTemplateCountn, v)
		case "power_port_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vPowerPortTemplateCountLt = append(vPowerPortTemplateCountLt, v)
		case "power_port_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vPowerPortTemplateCountLte = append(vPowerPortTemplateCountLte, v)
		case "power_port_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vPowerPortTemplateCountGt = append(vPowerPortTemplateCountGt, v)
		case "power_port_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vPowerPortTemplateCountGte = append(vPowerPortTemplateCountGte, v)
		case "power_port_template_count__empty":
			if vPowerPortTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_port_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_port_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vPowerPortTemplateCountEmpty = &v
		case "power_ports":
			if vPowerPorts != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'power_ports' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'power_ports' takes a boolean, got %q.", value))
				return
			}
			vPowerPorts = &v
		case "profile":
			v := value
			vProfile = append(vProfile, v)
		case "profile__n":
			v := value
			vProfilen = append(vProfilen, v)
		case "profile_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'profile_id' takes an integer, got %q.", value))
				return
			}
			vProfileID = append(vProfileID, v)
		case "profile_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'profile_id__n' takes an integer, got %q.", value))
				return
			}
			vProfileIDn = append(vProfileIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "rear_port_template_count":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_template_count' takes an integer, got %q.", value))
				return
			}
			vRearPortTemplateCount = append(vRearPortTemplateCount, v)
		case "rear_port_template_count__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_template_count__n' takes an integer, got %q.", value))
				return
			}
			vRearPortTemplateCountn = append(vRearPortTemplateCountn, v)
		case "rear_port_template_count__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_template_count__lt' takes an integer, got %q.", value))
				return
			}
			vRearPortTemplateCountLt = append(vRearPortTemplateCountLt, v)
		case "rear_port_template_count__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_template_count__lte' takes an integer, got %q.", value))
				return
			}
			vRearPortTemplateCountLte = append(vRearPortTemplateCountLte, v)
		case "rear_port_template_count__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_template_count__gt' takes an integer, got %q.", value))
				return
			}
			vRearPortTemplateCountGt = append(vRearPortTemplateCountGt, v)
		case "rear_port_template_count__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_template_count__gte' takes an integer, got %q.", value))
				return
			}
			vRearPortTemplateCountGte = append(vRearPortTemplateCountGte, v)
		case "rear_port_template_count__empty":
			if vRearPortTemplateCountEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'rear_port_template_count__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rear_port_template_count__empty' takes a boolean, got %q.", value))
				return
			}
			vRearPortTemplateCountEmpty = &v
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
	if vAirflow != nil {
		params.SetAirflow(vAirflow)
	}
	if vAirflown != nil {
		params.SetAirflown(vAirflown)
	}
	if len(vAirflowIc) > 0 {
		params.SetAirflowIc(vAirflowIc)
	}
	if len(vAirflowNic) > 0 {
		params.SetAirflowNic(vAirflowNic)
	}
	if len(vAirflowIe) > 0 {
		params.SetAirflowIe(vAirflowIe)
	}
	if len(vAirflowNie) > 0 {
		params.SetAirflowNie(vAirflowNie)
	}
	if len(vAirflowIsw) > 0 {
		params.SetAirflowIsw(vAirflowIsw)
	}
	if len(vAirflowNisw) > 0 {
		params.SetAirflowNisw(vAirflowNisw)
	}
	if len(vAirflowIew) > 0 {
		params.SetAirflowIew(vAirflowIew)
	}
	if len(vAirflowNiew) > 0 {
		params.SetAirflowNiew(vAirflowNiew)
	}
	if vAirflowEmpty != nil {
		params.SetAirflowEmpty(vAirflowEmpty)
	}
	if len(vAirflowRegex) > 0 {
		params.SetAirflowRegex(vAirflowRegex)
	}
	if len(vAirflowIregex) > 0 {
		params.SetAirflowIregex(vAirflowIregex)
	}
	if len(vConsolePortTemplateCount) > 0 {
		params.SetConsolePortTemplateCount(vConsolePortTemplateCount)
	}
	if len(vConsolePortTemplateCountn) > 0 {
		params.SetConsolePortTemplateCountn(vConsolePortTemplateCountn)
	}
	if len(vConsolePortTemplateCountLt) > 0 {
		params.SetConsolePortTemplateCountLt(vConsolePortTemplateCountLt)
	}
	if len(vConsolePortTemplateCountLte) > 0 {
		params.SetConsolePortTemplateCountLte(vConsolePortTemplateCountLte)
	}
	if len(vConsolePortTemplateCountGt) > 0 {
		params.SetConsolePortTemplateCountGt(vConsolePortTemplateCountGt)
	}
	if len(vConsolePortTemplateCountGte) > 0 {
		params.SetConsolePortTemplateCountGte(vConsolePortTemplateCountGte)
	}
	if vConsolePortTemplateCountEmpty != nil {
		params.SetConsolePortTemplateCountEmpty(vConsolePortTemplateCountEmpty)
	}
	if vConsolePorts != nil {
		params.SetConsolePorts(vConsolePorts)
	}
	if len(vConsoleServerPortTemplateCount) > 0 {
		params.SetConsoleServerPortTemplateCount(vConsoleServerPortTemplateCount)
	}
	if len(vConsoleServerPortTemplateCountn) > 0 {
		params.SetConsoleServerPortTemplateCountn(vConsoleServerPortTemplateCountn)
	}
	if len(vConsoleServerPortTemplateCountLt) > 0 {
		params.SetConsoleServerPortTemplateCountLt(vConsoleServerPortTemplateCountLt)
	}
	if len(vConsoleServerPortTemplateCountLte) > 0 {
		params.SetConsoleServerPortTemplateCountLte(vConsoleServerPortTemplateCountLte)
	}
	if len(vConsoleServerPortTemplateCountGt) > 0 {
		params.SetConsoleServerPortTemplateCountGt(vConsoleServerPortTemplateCountGt)
	}
	if len(vConsoleServerPortTemplateCountGte) > 0 {
		params.SetConsoleServerPortTemplateCountGte(vConsoleServerPortTemplateCountGte)
	}
	if vConsoleServerPortTemplateCountEmpty != nil {
		params.SetConsoleServerPortTemplateCountEmpty(vConsoleServerPortTemplateCountEmpty)
	}
	if vConsoleServerPorts != nil {
		params.SetConsoleServerPorts(vConsoleServerPorts)
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
	if len(vFrontPortTemplateCount) > 0 {
		params.SetFrontPortTemplateCount(vFrontPortTemplateCount)
	}
	if len(vFrontPortTemplateCountn) > 0 {
		params.SetFrontPortTemplateCountn(vFrontPortTemplateCountn)
	}
	if len(vFrontPortTemplateCountLt) > 0 {
		params.SetFrontPortTemplateCountLt(vFrontPortTemplateCountLt)
	}
	if len(vFrontPortTemplateCountLte) > 0 {
		params.SetFrontPortTemplateCountLte(vFrontPortTemplateCountLte)
	}
	if len(vFrontPortTemplateCountGt) > 0 {
		params.SetFrontPortTemplateCountGt(vFrontPortTemplateCountGt)
	}
	if len(vFrontPortTemplateCountGte) > 0 {
		params.SetFrontPortTemplateCountGte(vFrontPortTemplateCountGte)
	}
	if vFrontPortTemplateCountEmpty != nil {
		params.SetFrontPortTemplateCountEmpty(vFrontPortTemplateCountEmpty)
	}
	if len(vInterfaceTemplateCount) > 0 {
		params.SetInterfaceTemplateCount(vInterfaceTemplateCount)
	}
	if len(vInterfaceTemplateCountn) > 0 {
		params.SetInterfaceTemplateCountn(vInterfaceTemplateCountn)
	}
	if len(vInterfaceTemplateCountLt) > 0 {
		params.SetInterfaceTemplateCountLt(vInterfaceTemplateCountLt)
	}
	if len(vInterfaceTemplateCountLte) > 0 {
		params.SetInterfaceTemplateCountLte(vInterfaceTemplateCountLte)
	}
	if len(vInterfaceTemplateCountGt) > 0 {
		params.SetInterfaceTemplateCountGt(vInterfaceTemplateCountGt)
	}
	if len(vInterfaceTemplateCountGte) > 0 {
		params.SetInterfaceTemplateCountGte(vInterfaceTemplateCountGte)
	}
	if vInterfaceTemplateCountEmpty != nil {
		params.SetInterfaceTemplateCountEmpty(vInterfaceTemplateCountEmpty)
	}
	if vInterfaces != nil {
		params.SetInterfaces(vInterfaces)
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
	if len(vManufacturerID) > 0 {
		params.SetManufacturerID(vManufacturerID)
	}
	if len(vManufacturerIDn) > 0 {
		params.SetManufacturerIDn(vManufacturerIDn)
	}
	if len(vModuleBayTemplateCount) > 0 {
		params.SetModuleBayTemplateCount(vModuleBayTemplateCount)
	}
	if len(vModuleBayTemplateCountn) > 0 {
		params.SetModuleBayTemplateCountn(vModuleBayTemplateCountn)
	}
	if len(vModuleBayTemplateCountLt) > 0 {
		params.SetModuleBayTemplateCountLt(vModuleBayTemplateCountLt)
	}
	if len(vModuleBayTemplateCountLte) > 0 {
		params.SetModuleBayTemplateCountLte(vModuleBayTemplateCountLte)
	}
	if len(vModuleBayTemplateCountGt) > 0 {
		params.SetModuleBayTemplateCountGt(vModuleBayTemplateCountGt)
	}
	if len(vModuleBayTemplateCountGte) > 0 {
		params.SetModuleBayTemplateCountGte(vModuleBayTemplateCountGte)
	}
	if vModuleBayTemplateCountEmpty != nil {
		params.SetModuleBayTemplateCountEmpty(vModuleBayTemplateCountEmpty)
	}
	if vModuleBays != nil {
		params.SetModuleBays(vModuleBays)
	}
	if len(vModuleCount) > 0 {
		params.SetModuleCount(vModuleCount)
	}
	if len(vModuleCountn) > 0 {
		params.SetModuleCountn(vModuleCountn)
	}
	if len(vModuleCountLt) > 0 {
		params.SetModuleCountLt(vModuleCountLt)
	}
	if len(vModuleCountLte) > 0 {
		params.SetModuleCountLte(vModuleCountLte)
	}
	if len(vModuleCountGt) > 0 {
		params.SetModuleCountGt(vModuleCountGt)
	}
	if len(vModuleCountGte) > 0 {
		params.SetModuleCountGte(vModuleCountGte)
	}
	if vModuleCountEmpty != nil {
		params.SetModuleCountEmpty(vModuleCountEmpty)
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
	if len(vPartNumber) > 0 {
		params.SetPartNumber(vPartNumber)
	}
	if len(vPartNumbern) > 0 {
		params.SetPartNumbern(vPartNumbern)
	}
	if len(vPartNumberIc) > 0 {
		params.SetPartNumberIc(vPartNumberIc)
	}
	if len(vPartNumberNic) > 0 {
		params.SetPartNumberNic(vPartNumberNic)
	}
	if len(vPartNumberIe) > 0 {
		params.SetPartNumberIe(vPartNumberIe)
	}
	if len(vPartNumberNie) > 0 {
		params.SetPartNumberNie(vPartNumberNie)
	}
	if len(vPartNumberIsw) > 0 {
		params.SetPartNumberIsw(vPartNumberIsw)
	}
	if len(vPartNumberNisw) > 0 {
		params.SetPartNumberNisw(vPartNumberNisw)
	}
	if len(vPartNumberIew) > 0 {
		params.SetPartNumberIew(vPartNumberIew)
	}
	if len(vPartNumberNiew) > 0 {
		params.SetPartNumberNiew(vPartNumberNiew)
	}
	if vPartNumberEmpty != nil {
		params.SetPartNumberEmpty(vPartNumberEmpty)
	}
	if len(vPartNumberRegex) > 0 {
		params.SetPartNumberRegex(vPartNumberRegex)
	}
	if len(vPartNumberIregex) > 0 {
		params.SetPartNumberIregex(vPartNumberIregex)
	}
	if vPassThroughPorts != nil {
		params.SetPassThroughPorts(vPassThroughPorts)
	}
	if len(vPowerOutletTemplateCount) > 0 {
		params.SetPowerOutletTemplateCount(vPowerOutletTemplateCount)
	}
	if len(vPowerOutletTemplateCountn) > 0 {
		params.SetPowerOutletTemplateCountn(vPowerOutletTemplateCountn)
	}
	if len(vPowerOutletTemplateCountLt) > 0 {
		params.SetPowerOutletTemplateCountLt(vPowerOutletTemplateCountLt)
	}
	if len(vPowerOutletTemplateCountLte) > 0 {
		params.SetPowerOutletTemplateCountLte(vPowerOutletTemplateCountLte)
	}
	if len(vPowerOutletTemplateCountGt) > 0 {
		params.SetPowerOutletTemplateCountGt(vPowerOutletTemplateCountGt)
	}
	if len(vPowerOutletTemplateCountGte) > 0 {
		params.SetPowerOutletTemplateCountGte(vPowerOutletTemplateCountGte)
	}
	if vPowerOutletTemplateCountEmpty != nil {
		params.SetPowerOutletTemplateCountEmpty(vPowerOutletTemplateCountEmpty)
	}
	if vPowerOutlets != nil {
		params.SetPowerOutlets(vPowerOutlets)
	}
	if len(vPowerPortTemplateCount) > 0 {
		params.SetPowerPortTemplateCount(vPowerPortTemplateCount)
	}
	if len(vPowerPortTemplateCountn) > 0 {
		params.SetPowerPortTemplateCountn(vPowerPortTemplateCountn)
	}
	if len(vPowerPortTemplateCountLt) > 0 {
		params.SetPowerPortTemplateCountLt(vPowerPortTemplateCountLt)
	}
	if len(vPowerPortTemplateCountLte) > 0 {
		params.SetPowerPortTemplateCountLte(vPowerPortTemplateCountLte)
	}
	if len(vPowerPortTemplateCountGt) > 0 {
		params.SetPowerPortTemplateCountGt(vPowerPortTemplateCountGt)
	}
	if len(vPowerPortTemplateCountGte) > 0 {
		params.SetPowerPortTemplateCountGte(vPowerPortTemplateCountGte)
	}
	if vPowerPortTemplateCountEmpty != nil {
		params.SetPowerPortTemplateCountEmpty(vPowerPortTemplateCountEmpty)
	}
	if vPowerPorts != nil {
		params.SetPowerPorts(vPowerPorts)
	}
	if len(vProfile) > 0 {
		params.SetProfile(vProfile)
	}
	if len(vProfilen) > 0 {
		params.SetProfilen(vProfilen)
	}
	if len(vProfileID) > 0 {
		params.SetProfileID(vProfileID)
	}
	if len(vProfileIDn) > 0 {
		params.SetProfileIDn(vProfileIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRearPortTemplateCount) > 0 {
		params.SetRearPortTemplateCount(vRearPortTemplateCount)
	}
	if len(vRearPortTemplateCountn) > 0 {
		params.SetRearPortTemplateCountn(vRearPortTemplateCountn)
	}
	if len(vRearPortTemplateCountLt) > 0 {
		params.SetRearPortTemplateCountLt(vRearPortTemplateCountLt)
	}
	if len(vRearPortTemplateCountLte) > 0 {
		params.SetRearPortTemplateCountLte(vRearPortTemplateCountLte)
	}
	if len(vRearPortTemplateCountGt) > 0 {
		params.SetRearPortTemplateCountGt(vRearPortTemplateCountGt)
	}
	if len(vRearPortTemplateCountGte) > 0 {
		params.SetRearPortTemplateCountGte(vRearPortTemplateCountGte)
	}
	if vRearPortTemplateCountEmpty != nil {
		params.SetRearPortTemplateCountEmpty(vRearPortTemplateCountEmpty)
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
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Dcim.DcimModuleTypesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_module_types", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ModuleTypeResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]moduleTypeResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model moduleTypeResourceModel
		resp.Diagnostics.Append(flattenModuleType(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: moduleTypeResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
