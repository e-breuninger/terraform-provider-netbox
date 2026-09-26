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
	_ datasource.DataSource              = (*modulesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*modulesDataSource)(nil)
)

// NewModulesDataSource returns a new modules data source, which
// lists module objects matching its filters.
func NewModulesDataSource() datasource.DataSource {
	return &modulesDataSource{}
}

type modulesDataSource struct {
	client *netboxapi.Client
}

// modulesDataSourceModel is the data source model: the filters, the limit and the matching
// modules.
type modulesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"modules"`
}

// moduleResourceAttrTypes is the attribute type map of moduleResourceModel.
func moduleResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             types.Int64Type,
		"device_id":      types.Int64Type,
		"module_bay_id":  types.Int64Type,
		"module_type_id": types.Int64Type,
		"status":         types.StringType,
		"serial":         types.StringType,
		"asset_tag":      types.StringType,
		"description":    types.StringType,
		"comments":       types.StringType,
		"owner_id":       types.Int64Type,
		"created":        types.StringType,
		"last_updated":   types.StringType,
		"url":            types.StringType,
		"tags":           types.SetType{ElemType: types.StringType},
		"tags_all":       types.SetType{ElemType: types.StringType},
		"custom_fields":  types.MapType{ElemType: types.StringType},
	}
}

func (d *modulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_modules"
}

func (d *modulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists module objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: asset_tag, asset_tag__empty, asset_tag__ic, asset_tag__ie, asset_tag__iew, asset_tag__iregex, asset_tag__isw, asset_tag__n, asset_tag__nic, asset_tag__nie, asset_tag__niew, asset_tag__nisw, asset_tag__regex, device_id, device_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, module_bay_id, module_bay_id__n, module_type_id, module_type_id__n, owner_id, owner_id__n, serial, serial__empty, serial__ic, serial__ie, serial__iew, serial__iregex, serial__isw, serial__n, serial__nic, serial__nie, serial__niew, serial__nisw, serial__regex, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"modules": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching modules.",
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
						"module_bay_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the module bay the module is installed in.",
						},
						"module_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the module type.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: offline, active, planned, staged, failed, decommissioning.",
						},
						"serial": schema.StringAttribute{
							Computed:    true,
							Description: "Serial number.",
						},
						"asset_tag": schema.StringAttribute{
							Computed:    true,
							Description: "Unique asset tag.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *modulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *modulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data modulesDataSourceModel

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
	var responseDTOs []*netboxapi.ModuleResponseDTO

	params := dcim.NewDcimModulesListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vDeviceID []int64
	var vDeviceIDn []int64
	var vModuleTypeID []int64
	var vModuleTypeIDn []int64
	var vModuleBayID []string
	var vModuleBayIDn []string
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
		case "module_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_type_id' takes an integer, got %q.", value))
				return
			}
			vModuleTypeID = append(vModuleTypeID, v)
		case "module_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'module_type_id__n' takes an integer, got %q.", value))
				return
			}
			vModuleTypeIDn = append(vModuleTypeIDn, v)
		case "module_bay_id":
			v := value
			vModuleBayID = append(vModuleBayID, v)
		case "module_bay_id__n":
			v := value
			vModuleBayIDn = append(vModuleBayIDn, v)
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
	if len(vDeviceID) > 0 {
		params.SetDeviceID(vDeviceID)
	}
	if len(vDeviceIDn) > 0 {
		params.SetDeviceIDn(vDeviceIDn)
	}
	if len(vModuleTypeID) > 0 {
		params.SetModuleTypeID(vModuleTypeID)
	}
	if len(vModuleTypeIDn) > 0 {
		params.SetModuleTypeIDn(vModuleTypeIDn)
	}
	if len(vModuleBayID) > 0 {
		params.SetModuleBayID(vModuleBayID)
	}
	if len(vModuleBayIDn) > 0 {
		params.SetModuleBayIDn(vModuleBayIDn)
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
		res, err := d.client.Dcim.DcimModulesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_modules", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ModuleResponseDTOFromGoNetbox(goNetboxModel))
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
	items := make([]moduleResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model moduleResourceModel
		resp.Diagnostics.Append(flattenModule(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: moduleResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
