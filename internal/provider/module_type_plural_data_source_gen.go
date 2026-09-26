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
				Description: "Query filters, sent as API list parameters. Supported names: id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, model, model__empty, model__ic, model__ie, model__iew, model__iregex, model__isw, model__n, model__nic, model__nie, model__niew, model__nisw, model__regex, owner_id, owner_id__n, tag, tag__any, tag__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
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

	fetchAll := false
	_ = fetchAll
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
	var vModelIc []string
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
		case "model":
			v := value
			vModel = append(vModel, v)
		case "model__n":
			v := value
			vModeln = append(vModeln, v)
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
		case "model__ic":
			v := value
			vModelIc = append(vModelIc, v)
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
	if len(vModel) > 0 {
		params.SetModel(vModel)
	}
	if len(vModeln) > 0 {
		params.SetModeln(vModeln)
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
	if len(vModelIc) > 0 {
		params.SetModelIc(vModelIc)
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
		res, err := d.client.Dcim.DcimModuleTypesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_module_types", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ModuleTypeResponseDTOFromGoNetbox(goNetboxModel))
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
