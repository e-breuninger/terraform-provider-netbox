// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

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
	_ datasource.DataSource              = (*inventoryItemTemplatesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*inventoryItemTemplatesDataSource)(nil)
)

// NewInventoryItemTemplatesDataSource returns a new inventory_item_templates data source, which
// lists inventory_item_template objects matching its filters.
func NewInventoryItemTemplatesDataSource() datasource.DataSource {
	return &inventoryItemTemplatesDataSource{}
}

type inventoryItemTemplatesDataSource struct {
	client *netboxapi.Client
}

// inventoryItemTemplatesDataSourceModel is the data source model: the filters, the limit and the matching
// inventory_item_templates.
type inventoryItemTemplatesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"inventory_item_templates"`
}

// inventoryItemTemplateResourceAttrTypes is the attribute type map of inventoryItemTemplateResourceModel.
func inventoryItemTemplateResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":              types.Int64Type,
		"device_type_id":  types.Int64Type,
		"name":            types.StringType,
		"parent_id":       types.Int64Type,
		"role_id":         types.Int64Type,
		"manufacturer_id": types.Int64Type,
		"part_id":         types.StringType,
		"label":           types.StringType,
		"component_type":  types.StringType,
		"component_id":    types.Int64Type,
		"description":     types.StringType,
		"created":         types.StringType,
		"last_updated":    types.StringType,
		"url":             types.StringType,
	}
}

func (d *inventoryItemTemplatesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_item_templates"
}

func (d *inventoryItemTemplatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists inventory_item_template objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: component_id, component_id__empty, component_id__gt, component_id__gte, component_id__lt, component_id__lte, component_id__n, component_type, component_type__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device_type_id, device_type_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, manufacturer, manufacturer__n, manufacturer_id, manufacturer_id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, parent_id, parent_id__n, part_id, part_id__empty, part_id__ic, part_id__ie, part_id__iew, part_id__iregex, part_id__isw, part_id__n, part_id__nic, part_id__nie, part_id__niew, part_id__nisw, part_id__regex, q, role, role__n, role_id, role_id__n. Repeating a name sends that parameter once per value.",
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
			"inventory_item_templates": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching inventory_item_templates.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"device_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device type. NetBox refuses to move a template to another type.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"parent_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the parent inventory item template.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the inventory item role.",
						},
						"manufacturer_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the manufacturer.",
						},
						"part_id": schema.StringAttribute{
							Computed:    true,
							Description: "Manufacturer part number.",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Physical label.",
						},
						"component_type": schema.StringAttribute{
							Computed:    true,
							Description: "Type of the component template the item is bound to, e.g. dcim.interfacetemplate. Requires component_id.",
						},
						"component_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the component template named by component_type.",
						},
						"description": schema.StringAttribute{
							Computed: true,
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
					},
				},
			},
		},
	}
}

func (d *inventoryItemTemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *inventoryItemTemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data inventoryItemTemplatesDataSourceModel

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
	var responseDTOs []*netboxapi.InventoryItemTemplateResponseDTO

	params := dcim.NewDcimInventoryItemTemplatesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"name":                netboxapi.FilterStrings(params.SetName),
		"name__n":             netboxapi.FilterStrings(params.SetNamen),
		"name__ic":            netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":           netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":            netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":           netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":           netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":          netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":           netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":          netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":         netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":         netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":        netboxapi.FilterStrings(params.SetNameIregex),
		"device_type_id":      netboxapi.FilterInt64s(params.SetDeviceTypeID),
		"device_type_id__n":   netboxapi.FilterInt64s(params.SetDeviceTypeIDn),
		"label":               netboxapi.FilterStrings(params.SetLabel),
		"label__n":            netboxapi.FilterStrings(params.SetLabeln),
		"label__ic":           netboxapi.FilterStrings(params.SetLabelIc),
		"label__nic":          netboxapi.FilterStrings(params.SetLabelNic),
		"label__ie":           netboxapi.FilterStrings(params.SetLabelIe),
		"label__nie":          netboxapi.FilterStrings(params.SetLabelNie),
		"label__isw":          netboxapi.FilterStrings(params.SetLabelIsw),
		"label__nisw":         netboxapi.FilterStrings(params.SetLabelNisw),
		"label__iew":          netboxapi.FilterStrings(params.SetLabelIew),
		"label__niew":         netboxapi.FilterStrings(params.SetLabelNiew),
		"label__empty":        netboxapi.FilterBool(params.SetLabelEmpty),
		"label__regex":        netboxapi.FilterStrings(params.SetLabelRegex),
		"label__iregex":       netboxapi.FilterStrings(params.SetLabelIregex),
		"role_id":             netboxapi.FilterInt64s(params.SetRoleID),
		"role_id__n":          netboxapi.FilterInt64s(params.SetRoleIDn),
		"manufacturer_id":     netboxapi.FilterInt64s(params.SetManufacturerID),
		"manufacturer_id__n":  netboxapi.FilterInt64s(params.SetManufacturerIDn),
		"parent_id":           netboxapi.FilterInt64s(params.SetParentID),
		"parent_id__n":        netboxapi.FilterInt64s(params.SetParentIDn),
		"component_id":        netboxapi.FilterInt64s(params.SetComponentID),
		"component_id__n":     netboxapi.FilterInt64s(params.SetComponentIDn),
		"component_id__lt":    netboxapi.FilterInt64s(params.SetComponentIDLt),
		"component_id__lte":   netboxapi.FilterInt64s(params.SetComponentIDLte),
		"component_id__gt":    netboxapi.FilterInt64s(params.SetComponentIDGt),
		"component_id__gte":   netboxapi.FilterInt64s(params.SetComponentIDGte),
		"component_id__empty": netboxapi.FilterInt64s(params.SetComponentIDEmpty),
		"component_type":      netboxapi.FilterStrings(params.SetComponentType),
		"component_type__n":   netboxapi.FilterStrings(params.SetComponentTypen),
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
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"manufacturer":        netboxapi.FilterStrings(params.SetManufacturer),
		"manufacturer__n":     netboxapi.FilterStrings(params.SetManufacturern),
		"part_id":             netboxapi.FilterStrings(params.SetPartID),
		"part_id__n":          netboxapi.FilterStrings(params.SetPartIDn),
		"part_id__ic":         netboxapi.FilterStrings(params.SetPartIDIc),
		"part_id__nic":        netboxapi.FilterStrings(params.SetPartIDNic),
		"part_id__ie":         netboxapi.FilterStrings(params.SetPartIDIe),
		"part_id__nie":        netboxapi.FilterStrings(params.SetPartIDNie),
		"part_id__isw":        netboxapi.FilterStrings(params.SetPartIDIsw),
		"part_id__nisw":       netboxapi.FilterStrings(params.SetPartIDNisw),
		"part_id__iew":        netboxapi.FilterStrings(params.SetPartIDIew),
		"part_id__niew":       netboxapi.FilterStrings(params.SetPartIDNiew),
		"part_id__empty":      netboxapi.FilterBool(params.SetPartIDEmpty),
		"part_id__regex":      netboxapi.FilterStrings(params.SetPartIDRegex),
		"part_id__iregex":     netboxapi.FilterStrings(params.SetPartIDIregex),
		"q":                   netboxapi.FilterString(params.SetQ),
		"role":                netboxapi.FilterStrings(params.SetRole),
		"role__n":             netboxapi.FilterStrings(params.SetRolen),
	}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
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
		res, err := d.client.Dcim.DcimInventoryItemTemplatesListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_inventory_item_templates", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.InventoryItemTemplateResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]inventoryItemTemplateResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model inventoryItemTemplateResourceModel
		resp.Diagnostics.Append(flattenInventoryItemTemplate(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: inventoryItemTemplateResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
