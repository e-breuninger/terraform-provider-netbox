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
		"id":                   types.Int64Type,
		"device_id":            types.Int64Type,
		"module_bay_id":        types.Int64Type,
		"module_type_id":       types.Int64Type,
		"status":               types.StringType,
		"serial":               types.StringType,
		"asset_tag":            types.StringType,
		"replicate_components": types.BoolType,
		"adopt_components":     types.BoolType,
		"description":          types.StringType,
		"comments":             types.StringType,
		"owner_id":             types.Int64Type,
		"created":              types.StringType,
		"last_updated":         types.StringType,
		"url":                  types.StringType,
		"tags":                 types.SetType{ElemType: types.StringType},
		"tags_all":             types.SetType{ElemType: types.StringType},
		"custom_fields":        types.MapType{ElemType: types.StringType},
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
				Description: "Query filters, sent as API list parameters. Supported names: asset_tag, asset_tag__empty, asset_tag__ic, asset_tag__ie, asset_tag__iew, asset_tag__iregex, asset_tag__isw, asset_tag__n, asset_tag__nic, asset_tag__nie, asset_tag__niew, asset_tag__nisw, asset_tag__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device__n, device_id, device_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, module_bay_id, module_bay_id__n, module_type, module_type__n, module_type_id, module_type_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, profile, profile__n, profile_id, profile_id__n, q, rack, rack__n, rack_id, rack_id__n, region, region__n, region_id, region_id__n, serial, serial__empty, serial__ic, serial__ie, serial__iew, serial__iregex, serial__isw, serial__n, serial__nic, serial__nie, serial__niew, serial__nisw, serial__regex, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
						"replicate_components": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the components of the module type are created on the device along with the module. Only applied when the module is created.",
						},
						"adopt_components": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether existing components of the device that match the module type's templates are adopted by the module. Only applied when the module is created.",
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
	var vManufacturer []string
	var vManufacturern []string
	var vManufacturerID []int64
	var vManufacturerIDn []int64
	var vModuleType []string
	var vModuleTypen []string
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vProfile []string
	var vProfilen []string
	var vProfileID []int64
	var vProfileIDn []int64
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
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
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
		case "module_type":
			v := value
			vModuleType = append(vModuleType, v)
		case "module_type__n":
			v := value
			vModuleTypen = append(vModuleTypen, v)
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
	if len(vModuleType) > 0 {
		params.SetModuleType(vModuleType)
	}
	if len(vModuleTypen) > 0 {
		params.SetModuleTypen(vModuleTypen)
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
	if len(vTagID) > 0 {
		params.SetTagID(vTagID)
	}
	if len(vTagIDn) > 0 {
		params.SetTagIDn(vTagIDn)
	}
	if len(vTagIDAny) > 0 {
		params.SetTagIDAny(vTagIDAny)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Dcim.DcimModulesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_modules", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ModuleResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]moduleResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model moduleResourceModel
		model.ReplicateComponents = types.BoolNull()
		model.AdoptComponents = types.BoolNull()
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
