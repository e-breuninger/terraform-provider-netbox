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
	_ datasource.DataSource              = (*inventoryItemsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*inventoryItemsDataSource)(nil)
)

// NewInventoryItemsDataSource returns a new inventory_items data source, which
// lists inventory_item objects matching its filters.
func NewInventoryItemsDataSource() datasource.DataSource {
	return &inventoryItemsDataSource{}
}

type inventoryItemsDataSource struct {
	client *netboxapi.Client
}

// inventoryItemsDataSourceModel is the data source model: the filters, the limit and the matching
// inventory_items.
type inventoryItemsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"inventory_items"`
}

// inventoryItemResourceAttrTypes is the attribute type map of inventoryItemResourceModel.
func inventoryItemResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":              types.Int64Type,
		"device_id":       types.Int64Type,
		"name":            types.StringType,
		"parent_id":       types.Int64Type,
		"role_id":         types.Int64Type,
		"status":          types.StringType,
		"manufacturer_id": types.Int64Type,
		"part_id":         types.StringType,
		"serial":          types.StringType,
		"asset_tag":       types.StringType,
		"discovered":      types.BoolType,
		"component_type":  types.StringType,
		"component_id":    types.Int64Type,
		"label":           types.StringType,
		"description":     types.StringType,
		"owner_id":        types.Int64Type,
		"created":         types.StringType,
		"last_updated":    types.StringType,
		"url":             types.StringType,
		"tags":            types.SetType{ElemType: types.StringType},
		"tags_all":        types.SetType{ElemType: types.StringType},
		"custom_fields":   types.MapType{ElemType: types.StringType},
	}
}

func (d *inventoryItemsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_items"
}

func (d *inventoryItemsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists inventory_item objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: asset_tag, asset_tag__empty, asset_tag__ic, asset_tag__ie, asset_tag__iew, asset_tag__iregex, asset_tag__isw, asset_tag__n, asset_tag__nic, asset_tag__nie, asset_tag__niew, asset_tag__nisw, asset_tag__regex, component_id, component_id__empty, component_id__gt, component_id__gte, component_id__lt, component_id__lte, component_id__n, component_type, component_type__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device__n, device_id, device_id__n, device_role, device_role__n, device_role_id, device_role_id__n, device_status, device_status__empty, device_status__ic, device_status__ie, device_status__iew, device_status__iregex, device_status__isw, device_status__n, device_status__nic, device_status__nie, device_status__niew, device_status__nisw, device_status__regex, device_type, device_type__n, device_type_id, device_type_id__n, discovered, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, parent_id, parent_id__n, part_id, part_id__empty, part_id__ic, part_id__ie, part_id__iew, part_id__iregex, part_id__isw, part_id__n, part_id__nic, part_id__nie, part_id__niew, part_id__nisw, part_id__regex, q, rack, rack__n, rack_id, rack_id__n, region, region__n, region_id, region_id__n, role, role__n, role_id, role_id__n, serial, serial__empty, serial__ic, serial__ie, serial__iew, serial__iregex, serial__isw, serial__n, serial__nic, serial__nie, serial__niew, serial__nisw, serial__regex, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_id, tenant_id__n, virtual_chassis, virtual_chassis__n, virtual_chassis_id, virtual_chassis_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"inventory_items": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching inventory_items.",
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
						"name": schema.StringAttribute{
							Computed: true,
						},
						"parent_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the parent inventory item.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the inventory item role.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: offline, active, planned, staged, failed, decommissioning.",
						},
						"manufacturer_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the manufacturer.",
						},
						"part_id": schema.StringAttribute{
							Computed:    true,
							Description: "Manufacturer part number.",
						},
						"serial": schema.StringAttribute{
							Computed:    true,
							Description: "Serial number.",
						},
						"asset_tag": schema.StringAttribute{
							Computed:    true,
							Description: "Unique asset tag.",
						},
						"discovered": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the item was discovered automatically.",
						},
						"component_type": schema.StringAttribute{
							Computed:    true,
							Description: "Type of the device component the item is bound to, e.g. dcim.interface. Requires component_id.",
						},
						"component_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device component named by component_type.",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Physical label.",
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

func (d *inventoryItemsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *inventoryItemsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data inventoryItemsDataSourceModel

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
	var responseDTOs []*netboxapi.InventoryItemResponseDTO

	params := dcim.NewDcimInventoryItemsListParams()
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
	var vRoleID []int64
	var vRoleIDn []int64
	var vManufacturerID []int64
	var vManufacturerIDn []int64
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
	var vDiscovered *bool
	var vParentID []int64
	var vParentIDn []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vComponentID []int64
	var vComponentIDn []int64
	var vComponentIDLt []int64
	var vComponentIDLte []int64
	var vComponentIDGt []int64
	var vComponentIDGte []int64
	var vComponentIDEmpty []int64
	var vComponentType []string
	var vComponentTypen []string
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
	var vManufacturer []string
	var vManufacturern []string
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPartID []string
	var vPartIDn []string
	var vPartIDIc []string
	var vPartIDNic []string
	var vPartIDIe []string
	var vPartIDNie []string
	var vPartIDIsw []string
	var vPartIDNisw []string
	var vPartIDIew []string
	var vPartIDNiew []string
	var vPartIDEmpty *bool
	var vPartIDRegex []string
	var vPartIDIregex []string
	var vQ *string
	var vRack []string
	var vRackn []string
	var vRackID []int64
	var vRackIDn []int64
	var vRegion []string
	var vRegionn []string
	var vRegionID []string
	var vRegionIDn []string
	var vRole []string
	var vRolen []string
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []string
	var vSiteGroupIDn []string
	var vSiteID []int64
	var vSiteIDn []int64
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
		case "role_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id' takes an integer, got %q.", value))
				return
			}
			vRoleID = append(vRoleID, v)
		case "role_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id__n' takes an integer, got %q.", value))
				return
			}
			vRoleIDn = append(vRoleIDn, v)
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
		case "discovered":
			if vDiscovered != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'discovered' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'discovered' takes a boolean, got %q.", value))
				return
			}
			vDiscovered = &v
		case "parent_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_id' takes an integer, got %q.", value))
				return
			}
			vParentID = append(vParentID, v)
		case "parent_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_id__n' takes an integer, got %q.", value))
				return
			}
			vParentIDn = append(vParentIDn, v)
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
		case "component_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'component_id' takes an integer, got %q.", value))
				return
			}
			vComponentID = append(vComponentID, v)
		case "component_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'component_id__n' takes an integer, got %q.", value))
				return
			}
			vComponentIDn = append(vComponentIDn, v)
		case "component_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'component_id__lt' takes an integer, got %q.", value))
				return
			}
			vComponentIDLt = append(vComponentIDLt, v)
		case "component_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'component_id__lte' takes an integer, got %q.", value))
				return
			}
			vComponentIDLte = append(vComponentIDLte, v)
		case "component_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'component_id__gt' takes an integer, got %q.", value))
				return
			}
			vComponentIDGt = append(vComponentIDGt, v)
		case "component_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'component_id__gte' takes an integer, got %q.", value))
				return
			}
			vComponentIDGte = append(vComponentIDGte, v)
		case "component_id__empty":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'component_id__empty' takes an integer, got %q.", value))
				return
			}
			vComponentIDEmpty = append(vComponentIDEmpty, v)
		case "component_type":
			v := value
			vComponentType = append(vComponentType, v)
		case "component_type__n":
			v := value
			vComponentTypen = append(vComponentTypen, v)
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
		case "manufacturer":
			v := value
			vManufacturer = append(vManufacturer, v)
		case "manufacturer__n":
			v := value
			vManufacturern = append(vManufacturern, v)
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
		case "part_id":
			v := value
			vPartID = append(vPartID, v)
		case "part_id__n":
			v := value
			vPartIDn = append(vPartIDn, v)
		case "part_id__ic":
			v := value
			vPartIDIc = append(vPartIDIc, v)
		case "part_id__nic":
			v := value
			vPartIDNic = append(vPartIDNic, v)
		case "part_id__ie":
			v := value
			vPartIDIe = append(vPartIDIe, v)
		case "part_id__nie":
			v := value
			vPartIDNie = append(vPartIDNie, v)
		case "part_id__isw":
			v := value
			vPartIDIsw = append(vPartIDIsw, v)
		case "part_id__nisw":
			v := value
			vPartIDNisw = append(vPartIDNisw, v)
		case "part_id__iew":
			v := value
			vPartIDIew = append(vPartIDIew, v)
		case "part_id__niew":
			v := value
			vPartIDNiew = append(vPartIDNiew, v)
		case "part_id__empty":
			if vPartIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'part_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'part_id__empty' takes a boolean, got %q.", value))
				return
			}
			vPartIDEmpty = &v
		case "part_id__regex":
			v := value
			vPartIDRegex = append(vPartIDRegex, v)
		case "part_id__iregex":
			v := value
			vPartIDIregex = append(vPartIDIregex, v)
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
		case "role":
			v := value
			vRole = append(vRole, v)
		case "role__n":
			v := value
			vRolen = append(vRolen, v)
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
	if len(vRoleID) > 0 {
		params.SetRoleID(vRoleID)
	}
	if len(vRoleIDn) > 0 {
		params.SetRoleIDn(vRoleIDn)
	}
	if len(vManufacturerID) > 0 {
		params.SetManufacturerID(vManufacturerID)
	}
	if len(vManufacturerIDn) > 0 {
		params.SetManufacturerIDn(vManufacturerIDn)
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
	if vDiscovered != nil {
		params.SetDiscovered(vDiscovered)
	}
	if len(vParentID) > 0 {
		params.SetParentID(vParentID)
	}
	if len(vParentIDn) > 0 {
		params.SetParentIDn(vParentIDn)
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
	if len(vComponentID) > 0 {
		params.SetComponentID(vComponentID)
	}
	if len(vComponentIDn) > 0 {
		params.SetComponentIDn(vComponentIDn)
	}
	if len(vComponentIDLt) > 0 {
		params.SetComponentIDLt(vComponentIDLt)
	}
	if len(vComponentIDLte) > 0 {
		params.SetComponentIDLte(vComponentIDLte)
	}
	if len(vComponentIDGt) > 0 {
		params.SetComponentIDGt(vComponentIDGt)
	}
	if len(vComponentIDGte) > 0 {
		params.SetComponentIDGte(vComponentIDGte)
	}
	if len(vComponentIDEmpty) > 0 {
		params.SetComponentIDEmpty(vComponentIDEmpty)
	}
	if len(vComponentType) > 0 {
		params.SetComponentType(vComponentType)
	}
	if len(vComponentTypen) > 0 {
		params.SetComponentTypen(vComponentTypen)
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
	if len(vManufacturer) > 0 {
		params.SetManufacturer(vManufacturer)
	}
	if len(vManufacturern) > 0 {
		params.SetManufacturern(vManufacturern)
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
	if len(vPartID) > 0 {
		params.SetPartID(vPartID)
	}
	if len(vPartIDn) > 0 {
		params.SetPartIDn(vPartIDn)
	}
	if len(vPartIDIc) > 0 {
		params.SetPartIDIc(vPartIDIc)
	}
	if len(vPartIDNic) > 0 {
		params.SetPartIDNic(vPartIDNic)
	}
	if len(vPartIDIe) > 0 {
		params.SetPartIDIe(vPartIDIe)
	}
	if len(vPartIDNie) > 0 {
		params.SetPartIDNie(vPartIDNie)
	}
	if len(vPartIDIsw) > 0 {
		params.SetPartIDIsw(vPartIDIsw)
	}
	if len(vPartIDNisw) > 0 {
		params.SetPartIDNisw(vPartIDNisw)
	}
	if len(vPartIDIew) > 0 {
		params.SetPartIDIew(vPartIDIew)
	}
	if len(vPartIDNiew) > 0 {
		params.SetPartIDNiew(vPartIDNiew)
	}
	if vPartIDEmpty != nil {
		params.SetPartIDEmpty(vPartIDEmpty)
	}
	if len(vPartIDRegex) > 0 {
		params.SetPartIDRegex(vPartIDRegex)
	}
	if len(vPartIDIregex) > 0 {
		params.SetPartIDIregex(vPartIDIregex)
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
	if len(vRole) > 0 {
		params.SetRole(vRole)
	}
	if len(vRolen) > 0 {
		params.SetRolen(vRolen)
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
		res, err := d.client.Dcim.DcimInventoryItemsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_inventory_items", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.InventoryItemResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]inventoryItemResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model inventoryItemResourceModel
		resp.Diagnostics.Append(flattenInventoryItem(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: inventoryItemResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
