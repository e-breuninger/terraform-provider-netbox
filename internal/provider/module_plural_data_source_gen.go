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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"device_id":           netboxapi.FilterInt64s(params.SetDeviceID),
		"device_id__n":        netboxapi.FilterInt64s(params.SetDeviceIDn),
		"module_type_id":      netboxapi.FilterInt64s(params.SetModuleTypeID),
		"module_type_id__n":   netboxapi.FilterInt64s(params.SetModuleTypeIDn),
		"module_bay_id":       netboxapi.FilterStrings(params.SetModuleBayID),
		"module_bay_id__n":    netboxapi.FilterStrings(params.SetModuleBayIDn),
		"status":              netboxapi.FilterStrings(params.SetStatus),
		"status__n":           netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":          netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":         netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":          netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":         netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":         netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":        netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":         netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":        netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":       netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":       netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":      netboxapi.FilterStrings(params.SetStatusIregex),
		"serial":              netboxapi.FilterStrings(params.SetSerial),
		"serial__n":           netboxapi.FilterStrings(params.SetSerialn),
		"serial__ic":          netboxapi.FilterStrings(params.SetSerialIc),
		"serial__nic":         netboxapi.FilterStrings(params.SetSerialNic),
		"serial__ie":          netboxapi.FilterStrings(params.SetSerialIe),
		"serial__nie":         netboxapi.FilterStrings(params.SetSerialNie),
		"serial__isw":         netboxapi.FilterStrings(params.SetSerialIsw),
		"serial__nisw":        netboxapi.FilterStrings(params.SetSerialNisw),
		"serial__iew":         netboxapi.FilterStrings(params.SetSerialIew),
		"serial__niew":        netboxapi.FilterStrings(params.SetSerialNiew),
		"serial__empty":       netboxapi.FilterBool(params.SetSerialEmpty),
		"serial__regex":       netboxapi.FilterStrings(params.SetSerialRegex),
		"serial__iregex":      netboxapi.FilterStrings(params.SetSerialIregex),
		"asset_tag":           netboxapi.FilterStrings(params.SetAssetTag),
		"asset_tag__n":        netboxapi.FilterStrings(params.SetAssetTagn),
		"asset_tag__ic":       netboxapi.FilterStrings(params.SetAssetTagIc),
		"asset_tag__nic":      netboxapi.FilterStrings(params.SetAssetTagNic),
		"asset_tag__ie":       netboxapi.FilterStrings(params.SetAssetTagIe),
		"asset_tag__nie":      netboxapi.FilterStrings(params.SetAssetTagNie),
		"asset_tag__isw":      netboxapi.FilterStrings(params.SetAssetTagIsw),
		"asset_tag__nisw":     netboxapi.FilterStrings(params.SetAssetTagNisw),
		"asset_tag__iew":      netboxapi.FilterStrings(params.SetAssetTagIew),
		"asset_tag__niew":     netboxapi.FilterStrings(params.SetAssetTagNiew),
		"asset_tag__empty":    netboxapi.FilterBool(params.SetAssetTagEmpty),
		"asset_tag__regex":    netboxapi.FilterStrings(params.SetAssetTagRegex),
		"asset_tag__iregex":   netboxapi.FilterStrings(params.SetAssetTagIregex),
		"owner_id":            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":         netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":            netboxapi.FilterStrings(params.SetTagAny),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":         netboxapi.FilterStrings(params.SetDescription),
		"description__n":      netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":     netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":    netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":     netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":    netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":    netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":   netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":    netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":   netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":  netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":  netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex": netboxapi.FilterStrings(params.SetDescriptionIregex),
		"device":              netboxapi.FilterStrings(params.SetDevice),
		"device__n":           netboxapi.FilterStrings(params.SetDevicen),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":            netboxapi.FilterStrings(params.SetLocation),
		"location__n":         netboxapi.FilterStrings(params.SetLocationn),
		"location_id":         netboxapi.FilterInt64s(params.SetLocationID),
		"location_id__n":      netboxapi.FilterInt64s(params.SetLocationIDn),
		"manufacturer":        netboxapi.FilterStrings(params.SetManufacturer),
		"manufacturer__n":     netboxapi.FilterStrings(params.SetManufacturern),
		"manufacturer_id":     netboxapi.FilterInt64s(params.SetManufacturerID),
		"manufacturer_id__n":  netboxapi.FilterInt64s(params.SetManufacturerIDn),
		"module_type":         netboxapi.FilterStrings(params.SetModuleType),
		"module_type__n":      netboxapi.FilterStrings(params.SetModuleTypen),
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"profile":             netboxapi.FilterStrings(params.SetProfile),
		"profile__n":          netboxapi.FilterStrings(params.SetProfilen),
		"profile_id":          netboxapi.FilterInt64s(params.SetProfileID),
		"profile_id__n":       netboxapi.FilterInt64s(params.SetProfileIDn),
		"q":                   netboxapi.FilterString(params.SetQ),
		"rack":                netboxapi.FilterStrings(params.SetRack),
		"rack__n":             netboxapi.FilterStrings(params.SetRackn),
		"rack_id":             netboxapi.FilterInt64s(params.SetRackID),
		"rack_id__n":          netboxapi.FilterInt64s(params.SetRackIDn),
		"region":              netboxapi.FilterStrings(params.SetRegion),
		"region__n":           netboxapi.FilterStrings(params.SetRegionn),
		"region_id":           netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":        netboxapi.FilterStrings(params.SetRegionIDn),
		"site":                netboxapi.FilterStrings(params.SetSite),
		"site__n":             netboxapi.FilterStrings(params.SetSiten),
		"site_group":          netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":       netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":       netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":    netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"site_id":             netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":          netboxapi.FilterInt64s(params.SetSiteIDn),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
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
