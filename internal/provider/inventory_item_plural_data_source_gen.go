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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"name":                  netboxapi.FilterStrings(params.SetName),
		"name__n":               netboxapi.FilterStrings(params.SetNamen),
		"name__ic":              netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":             netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":              netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":             netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":             netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":            netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":             netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":            netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":           netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":           netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":          netboxapi.FilterStrings(params.SetNameIregex),
		"device_id":             netboxapi.FilterInt64s(params.SetDeviceID),
		"device_id__n":          netboxapi.FilterInt64s(params.SetDeviceIDn),
		"label":                 netboxapi.FilterStrings(params.SetLabel),
		"label__n":              netboxapi.FilterStrings(params.SetLabeln),
		"label__ic":             netboxapi.FilterStrings(params.SetLabelIc),
		"label__nic":            netboxapi.FilterStrings(params.SetLabelNic),
		"label__ie":             netboxapi.FilterStrings(params.SetLabelIe),
		"label__nie":            netboxapi.FilterStrings(params.SetLabelNie),
		"label__isw":            netboxapi.FilterStrings(params.SetLabelIsw),
		"label__nisw":           netboxapi.FilterStrings(params.SetLabelNisw),
		"label__iew":            netboxapi.FilterStrings(params.SetLabelIew),
		"label__niew":           netboxapi.FilterStrings(params.SetLabelNiew),
		"label__empty":          netboxapi.FilterBool(params.SetLabelEmpty),
		"label__regex":          netboxapi.FilterStrings(params.SetLabelRegex),
		"label__iregex":         netboxapi.FilterStrings(params.SetLabelIregex),
		"role_id":               netboxapi.FilterInt64s(params.SetRoleID),
		"role_id__n":            netboxapi.FilterInt64s(params.SetRoleIDn),
		"manufacturer_id":       netboxapi.FilterInt64s(params.SetManufacturerID),
		"manufacturer_id__n":    netboxapi.FilterInt64s(params.SetManufacturerIDn),
		"serial":                netboxapi.FilterStrings(params.SetSerial),
		"serial__n":             netboxapi.FilterStrings(params.SetSerialn),
		"serial__ic":            netboxapi.FilterStrings(params.SetSerialIc),
		"serial__nic":           netboxapi.FilterStrings(params.SetSerialNic),
		"serial__ie":            netboxapi.FilterStrings(params.SetSerialIe),
		"serial__nie":           netboxapi.FilterStrings(params.SetSerialNie),
		"serial__isw":           netboxapi.FilterStrings(params.SetSerialIsw),
		"serial__nisw":          netboxapi.FilterStrings(params.SetSerialNisw),
		"serial__iew":           netboxapi.FilterStrings(params.SetSerialIew),
		"serial__niew":          netboxapi.FilterStrings(params.SetSerialNiew),
		"serial__empty":         netboxapi.FilterBool(params.SetSerialEmpty),
		"serial__regex":         netboxapi.FilterStrings(params.SetSerialRegex),
		"serial__iregex":        netboxapi.FilterStrings(params.SetSerialIregex),
		"asset_tag":             netboxapi.FilterStrings(params.SetAssetTag),
		"asset_tag__n":          netboxapi.FilterStrings(params.SetAssetTagn),
		"asset_tag__ic":         netboxapi.FilterStrings(params.SetAssetTagIc),
		"asset_tag__nic":        netboxapi.FilterStrings(params.SetAssetTagNic),
		"asset_tag__ie":         netboxapi.FilterStrings(params.SetAssetTagIe),
		"asset_tag__nie":        netboxapi.FilterStrings(params.SetAssetTagNie),
		"asset_tag__isw":        netboxapi.FilterStrings(params.SetAssetTagIsw),
		"asset_tag__nisw":       netboxapi.FilterStrings(params.SetAssetTagNisw),
		"asset_tag__iew":        netboxapi.FilterStrings(params.SetAssetTagIew),
		"asset_tag__niew":       netboxapi.FilterStrings(params.SetAssetTagNiew),
		"asset_tag__empty":      netboxapi.FilterBool(params.SetAssetTagEmpty),
		"asset_tag__regex":      netboxapi.FilterStrings(params.SetAssetTagRegex),
		"asset_tag__iregex":     netboxapi.FilterStrings(params.SetAssetTagIregex),
		"discovered":            netboxapi.FilterBool(params.SetDiscovered),
		"parent_id":             netboxapi.FilterInt64s(params.SetParentID),
		"parent_id__n":          netboxapi.FilterInt64s(params.SetParentIDn),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"component_id":          netboxapi.FilterInt64s(params.SetComponentID),
		"component_id__n":       netboxapi.FilterInt64s(params.SetComponentIDn),
		"component_id__lt":      netboxapi.FilterInt64s(params.SetComponentIDLt),
		"component_id__lte":     netboxapi.FilterInt64s(params.SetComponentIDLte),
		"component_id__gt":      netboxapi.FilterInt64s(params.SetComponentIDGt),
		"component_id__gte":     netboxapi.FilterInt64s(params.SetComponentIDGte),
		"component_id__empty":   netboxapi.FilterInt64s(params.SetComponentIDEmpty),
		"component_type":        netboxapi.FilterStrings(params.SetComponentType),
		"component_type__n":     netboxapi.FilterStrings(params.SetComponentTypen),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":           netboxapi.FilterStrings(params.SetDescription),
		"description__n":        netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":       netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":      netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":       netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":      netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":      netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":     netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":      netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":     netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":    netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":    netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":   netboxapi.FilterStrings(params.SetDescriptionIregex),
		"device":                netboxapi.FilterStrings(params.SetDevice),
		"device__n":             netboxapi.FilterStrings(params.SetDevicen),
		"device_role":           netboxapi.FilterStrings(params.SetDeviceRole),
		"device_role__n":        netboxapi.FilterStrings(params.SetDeviceRolen),
		"device_role_id":        netboxapi.FilterInt64s(params.SetDeviceRoleID),
		"device_role_id__n":     netboxapi.FilterInt64s(params.SetDeviceRoleIDn),
		"device_status":         netboxapi.FilterStrings(params.SetDeviceStatus),
		"device_status__n":      netboxapi.FilterStrings(params.SetDeviceStatusn),
		"device_status__ic":     netboxapi.FilterStrings(params.SetDeviceStatusIc),
		"device_status__nic":    netboxapi.FilterStrings(params.SetDeviceStatusNic),
		"device_status__ie":     netboxapi.FilterStrings(params.SetDeviceStatusIe),
		"device_status__nie":    netboxapi.FilterStrings(params.SetDeviceStatusNie),
		"device_status__isw":    netboxapi.FilterStrings(params.SetDeviceStatusIsw),
		"device_status__nisw":   netboxapi.FilterStrings(params.SetDeviceStatusNisw),
		"device_status__iew":    netboxapi.FilterStrings(params.SetDeviceStatusIew),
		"device_status__niew":   netboxapi.FilterStrings(params.SetDeviceStatusNiew),
		"device_status__empty":  netboxapi.FilterBool(params.SetDeviceStatusEmpty),
		"device_status__regex":  netboxapi.FilterStrings(params.SetDeviceStatusRegex),
		"device_status__iregex": netboxapi.FilterStrings(params.SetDeviceStatusIregex),
		"device_type":           netboxapi.FilterStrings(params.SetDeviceType),
		"device_type__n":        netboxapi.FilterStrings(params.SetDeviceTypen),
		"device_type_id":        netboxapi.FilterInt64s(params.SetDeviceTypeID),
		"device_type_id__n":     netboxapi.FilterInt64s(params.SetDeviceTypeIDn),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":              netboxapi.FilterStrings(params.SetLocation),
		"location__n":           netboxapi.FilterStrings(params.SetLocationn),
		"location_id":           netboxapi.FilterInt64s(params.SetLocationID),
		"location_id__n":        netboxapi.FilterInt64s(params.SetLocationIDn),
		"manufacturer":          netboxapi.FilterStrings(params.SetManufacturer),
		"manufacturer__n":       netboxapi.FilterStrings(params.SetManufacturern),
		"owner":                 netboxapi.FilterStrings(params.SetOwner),
		"owner__n":              netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":           netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":        netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":        netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":     netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"part_id":               netboxapi.FilterStrings(params.SetPartID),
		"part_id__n":            netboxapi.FilterStrings(params.SetPartIDn),
		"part_id__ic":           netboxapi.FilterStrings(params.SetPartIDIc),
		"part_id__nic":          netboxapi.FilterStrings(params.SetPartIDNic),
		"part_id__ie":           netboxapi.FilterStrings(params.SetPartIDIe),
		"part_id__nie":          netboxapi.FilterStrings(params.SetPartIDNie),
		"part_id__isw":          netboxapi.FilterStrings(params.SetPartIDIsw),
		"part_id__nisw":         netboxapi.FilterStrings(params.SetPartIDNisw),
		"part_id__iew":          netboxapi.FilterStrings(params.SetPartIDIew),
		"part_id__niew":         netboxapi.FilterStrings(params.SetPartIDNiew),
		"part_id__empty":        netboxapi.FilterBool(params.SetPartIDEmpty),
		"part_id__regex":        netboxapi.FilterStrings(params.SetPartIDRegex),
		"part_id__iregex":       netboxapi.FilterStrings(params.SetPartIDIregex),
		"q":                     netboxapi.FilterString(params.SetQ),
		"rack":                  netboxapi.FilterStrings(params.SetRack),
		"rack__n":               netboxapi.FilterStrings(params.SetRackn),
		"rack_id":               netboxapi.FilterInt64s(params.SetRackID),
		"rack_id__n":            netboxapi.FilterInt64s(params.SetRackIDn),
		"region":                netboxapi.FilterStrings(params.SetRegion),
		"region__n":             netboxapi.FilterStrings(params.SetRegionn),
		"region_id":             netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":          netboxapi.FilterStrings(params.SetRegionIDn),
		"role":                  netboxapi.FilterStrings(params.SetRole),
		"role__n":               netboxapi.FilterStrings(params.SetRolen),
		"site":                  netboxapi.FilterStrings(params.SetSite),
		"site__n":               netboxapi.FilterStrings(params.SetSiten),
		"site_group":            netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":         netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":         netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":      netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"site_id":               netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":            netboxapi.FilterInt64s(params.SetSiteIDn),
		"status":                netboxapi.FilterStrings(params.SetStatus),
		"status__n":             netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":            netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":           netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":            netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":           netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":           netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":          netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":           netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":          netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":         netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":         netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":        netboxapi.FilterStrings(params.SetStatusIregex),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":             netboxapi.FilterStrings(params.SetTenantn),
		"tenant_id":             netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":          netboxapi.FilterInt64s(params.SetTenantIDn),
		"virtual_chassis":       netboxapi.FilterStrings(params.SetVirtualChassis),
		"virtual_chassis__n":    netboxapi.FilterStrings(params.SetVirtualChassisn),
		"virtual_chassis_id":    netboxapi.FilterInt64s(params.SetVirtualChassisID),
		"virtual_chassis_id__n": netboxapi.FilterInt64s(params.SetVirtualChassisIDn),
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
