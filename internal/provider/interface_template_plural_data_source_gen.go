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
	_ datasource.DataSource              = (*interfaceTemplatesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*interfaceTemplatesDataSource)(nil)
)

// NewInterfaceTemplatesDataSource returns a new interface_templates data source, which
// lists interface_template objects matching its filters.
func NewInterfaceTemplatesDataSource() datasource.DataSource {
	return &interfaceTemplatesDataSource{}
}

type interfaceTemplatesDataSource struct {
	client *netboxapi.Client
}

// interfaceTemplatesDataSourceModel is the data source model: the filters, the limit and the matching
// interface_templates.
type interfaceTemplatesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"interface_templates"`
}

// interfaceTemplateResourceAttrTypes is the attribute type map of interfaceTemplateResourceModel.
func interfaceTemplateResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             types.Int64Type,
		"device_type_id": types.Int64Type,
		"module_type_id": types.Int64Type,
		"name":           types.StringType,
		"type":           types.StringType,
		"enabled":        types.BoolType,
		"mgmt_only":      types.BoolType,
		"bridge_id":      types.Int64Type,
		"poe_mode":       types.StringType,
		"poe_type":       types.StringType,
		"rf_role":        types.StringType,
		"label":          types.StringType,
		"description":    types.StringType,
		"created":        types.StringType,
		"last_updated":   types.StringType,
		"url":            types.StringType,
	}
}

func (d *interfaceTemplatesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_templates"
}

func (d *interfaceTemplatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists interface_template objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: bridge_id, bridge_id__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device_type_id, device_type_id__n, enabled, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mgmt_only, module_type_id, module_type_id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, poe_mode, poe_mode__empty, poe_mode__ic, poe_mode__ie, poe_mode__iew, poe_mode__iregex, poe_mode__isw, poe_mode__n, poe_mode__nic, poe_mode__nie, poe_mode__niew, poe_mode__nisw, poe_mode__regex, poe_type, poe_type__empty, poe_type__ic, poe_type__ie, poe_type__iew, poe_type__iregex, poe_type__isw, poe_type__n, poe_type__nic, poe_type__nie, poe_type__niew, poe_type__nisw, poe_type__regex, q, rf_role, rf_role__empty, rf_role__ic, rf_role__ie, rf_role__iew, rf_role__iregex, rf_role__isw, rf_role__n, rf_role__nic, rf_role__nie, rf_role__niew, rf_role__nisw, rf_role__regex, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex. Repeating a name sends that parameter once per value.",
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
			"interface_templates": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching interface_templates.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"device_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device type. Set exactly one of device_type_id and module_type_id; NetBox rejects both and neither.",
						},
						"module_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the module type. Set exactly one of device_type_id and module_type_id; NetBox rejects both and neither.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Interface type as its NetBox slug, e.g. virtual, lag, bridge, 1000base-t, 10gbase-x-sfpp (any of NetBox's interface type choices).",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the interface is enabled.",
						},
						"mgmt_only": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the interface is used for out-of-band management only.",
						},
						"bridge_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the bridge interface template.",
						},
						"poe_mode": schema.StringAttribute{
							Computed:    true,
							Description: "Power over Ethernet mode. One of: pd, pse.",
						},
						"poe_type": schema.StringAttribute{
							Computed:    true,
							Description: "Power over Ethernet type as its NetBox slug, e.g. type1-ieee802.3af, type3-ieee802.3bt, passive-24v-2pair (any of NetBox's PoE type choices).",
						},
						"rf_role": schema.StringAttribute{
							Computed:    true,
							Description: "Wireless role. One of: ap, station.",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Physical label.",
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

func (d *interfaceTemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *interfaceTemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data interfaceTemplatesDataSourceModel

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
	var responseDTOs []*netboxapi.InterfaceTemplateResponseDTO

	params := dcim.NewDcimInterfaceTemplatesListParams()
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
		"type":                netboxapi.FilterStrings(params.SetType),
		"type__n":             netboxapi.FilterStrings(params.SetTypen),
		"type__ic":            netboxapi.FilterStrings(params.SetTypeIc),
		"type__nic":           netboxapi.FilterStrings(params.SetTypeNic),
		"type__ie":            netboxapi.FilterStrings(params.SetTypeIe),
		"type__nie":           netboxapi.FilterStrings(params.SetTypeNie),
		"type__isw":           netboxapi.FilterStrings(params.SetTypeIsw),
		"type__nisw":          netboxapi.FilterStrings(params.SetTypeNisw),
		"type__iew":           netboxapi.FilterStrings(params.SetTypeIew),
		"type__niew":          netboxapi.FilterStrings(params.SetTypeNiew),
		"type__empty":         netboxapi.FilterBool(params.SetTypeEmpty),
		"type__regex":         netboxapi.FilterStrings(params.SetTypeRegex),
		"type__iregex":        netboxapi.FilterStrings(params.SetTypeIregex),
		"mgmt_only":           netboxapi.FilterBool(params.SetMgmtOnly),
		"module_type_id":      netboxapi.FilterInt64s(params.SetModuleTypeID),
		"module_type_id__n":   netboxapi.FilterInt64s(params.SetModuleTypeIDn),
		"bridge_id":           netboxapi.FilterInt64s(params.SetBridgeID),
		"bridge_id__n":        netboxapi.FilterInt64s(params.SetBridgeIDn),
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
		"device_type_id":      netboxapi.FilterInt64s(params.SetDeviceTypeID),
		"device_type_id__n":   netboxapi.FilterInt64s(params.SetDeviceTypeIDn),
		"enabled":             netboxapi.FilterBool(params.SetEnabled),
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
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"poe_mode":            netboxapi.FilterStrings(params.SetPoeMode),
		"poe_mode__n":         netboxapi.FilterStrings(params.SetPoeModen),
		"poe_mode__ic":        netboxapi.FilterStrings(params.SetPoeModeIc),
		"poe_mode__nic":       netboxapi.FilterStrings(params.SetPoeModeNic),
		"poe_mode__ie":        netboxapi.FilterStrings(params.SetPoeModeIe),
		"poe_mode__nie":       netboxapi.FilterStrings(params.SetPoeModeNie),
		"poe_mode__isw":       netboxapi.FilterStrings(params.SetPoeModeIsw),
		"poe_mode__nisw":      netboxapi.FilterStrings(params.SetPoeModeNisw),
		"poe_mode__iew":       netboxapi.FilterStrings(params.SetPoeModeIew),
		"poe_mode__niew":      netboxapi.FilterStrings(params.SetPoeModeNiew),
		"poe_mode__empty":     netboxapi.FilterBool(params.SetPoeModeEmpty),
		"poe_mode__regex":     netboxapi.FilterStrings(params.SetPoeModeRegex),
		"poe_mode__iregex":    netboxapi.FilterStrings(params.SetPoeModeIregex),
		"poe_type":            netboxapi.FilterStrings(params.SetPoeType),
		"poe_type__n":         netboxapi.FilterStrings(params.SetPoeTypen),
		"poe_type__ic":        netboxapi.FilterStrings(params.SetPoeTypeIc),
		"poe_type__nic":       netboxapi.FilterStrings(params.SetPoeTypeNic),
		"poe_type__ie":        netboxapi.FilterStrings(params.SetPoeTypeIe),
		"poe_type__nie":       netboxapi.FilterStrings(params.SetPoeTypeNie),
		"poe_type__isw":       netboxapi.FilterStrings(params.SetPoeTypeIsw),
		"poe_type__nisw":      netboxapi.FilterStrings(params.SetPoeTypeNisw),
		"poe_type__iew":       netboxapi.FilterStrings(params.SetPoeTypeIew),
		"poe_type__niew":      netboxapi.FilterStrings(params.SetPoeTypeNiew),
		"poe_type__empty":     netboxapi.FilterBool(params.SetPoeTypeEmpty),
		"poe_type__regex":     netboxapi.FilterStrings(params.SetPoeTypeRegex),
		"poe_type__iregex":    netboxapi.FilterStrings(params.SetPoeTypeIregex),
		"q":                   netboxapi.FilterString(params.SetQ),
		"rf_role":             netboxapi.FilterStrings(params.SetRfRole),
		"rf_role__n":          netboxapi.FilterStrings(params.SetRfRolen),
		"rf_role__ic":         netboxapi.FilterStrings(params.SetRfRoleIc),
		"rf_role__nic":        netboxapi.FilterStrings(params.SetRfRoleNic),
		"rf_role__ie":         netboxapi.FilterStrings(params.SetRfRoleIe),
		"rf_role__nie":        netboxapi.FilterStrings(params.SetRfRoleNie),
		"rf_role__isw":        netboxapi.FilterStrings(params.SetRfRoleIsw),
		"rf_role__nisw":       netboxapi.FilterStrings(params.SetRfRoleNisw),
		"rf_role__iew":        netboxapi.FilterStrings(params.SetRfRoleIew),
		"rf_role__niew":       netboxapi.FilterStrings(params.SetRfRoleNiew),
		"rf_role__empty":      netboxapi.FilterBool(params.SetRfRoleEmpty),
		"rf_role__regex":      netboxapi.FilterStrings(params.SetRfRoleRegex),
		"rf_role__iregex":     netboxapi.FilterStrings(params.SetRfRoleIregex),
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
		res, err := d.client.Dcim.DcimInterfaceTemplatesListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_interface_templates", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.InterfaceTemplateResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]interfaceTemplateResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model interfaceTemplateResourceModel
		resp.Diagnostics.Append(flattenInterfaceTemplate(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: interfaceTemplateResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
