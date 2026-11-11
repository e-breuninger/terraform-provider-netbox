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
	_ datasource.DataSource              = (*powerOutletTemplatesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*powerOutletTemplatesDataSource)(nil)
)

// NewPowerOutletTemplatesDataSource returns a new power_outlet_templates data source, which
// lists power_outlet_template objects matching its filters.
func NewPowerOutletTemplatesDataSource() datasource.DataSource {
	return &powerOutletTemplatesDataSource{}
}

type powerOutletTemplatesDataSource struct {
	client *netboxapi.Client
}

// powerOutletTemplatesDataSourceModel is the data source model: the filters, the limit and the matching
// power_outlet_templates.
type powerOutletTemplatesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"power_outlet_templates"`
}

// powerOutletTemplateResourceAttrTypes is the attribute type map of powerOutletTemplateResourceModel.
func powerOutletTemplateResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                     types.Int64Type,
		"device_type_id":         types.Int64Type,
		"module_type_id":         types.Int64Type,
		"name":                   types.StringType,
		"type":                   types.StringType,
		"power_port_template_id": types.Int64Type,
		"power_port_id":          types.Int64Type,
		"feed_leg":               types.StringType,
		"color_hex":              types.StringType,
		"label":                  types.StringType,
		"description":            types.StringType,
		"created":                types.StringType,
		"last_updated":           types.StringType,
		"url":                    types.StringType,
	}
}

func (d *powerOutletTemplatesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_power_outlet_templates"
}

func (d *powerOutletTemplatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists power_outlet_template objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: color, color__empty, color__ic, color__ie, color__iew, color__iregex, color__isw, color__n, color__nic, color__nie, color__niew, color__nisw, color__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device_type_id, device_type_id__n, feed_leg, feed_leg__empty, feed_leg__ic, feed_leg__ie, feed_leg__iew, feed_leg__iregex, feed_leg__isw, feed_leg__n, feed_leg__nic, feed_leg__nie, feed_leg__niew, feed_leg__nisw, feed_leg__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, module_type_id, module_type_id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, power_port_id, power_port_id__n, q, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex. Repeating a name sends that parameter once per value.",
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
			"power_outlet_templates": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching power_outlet_templates.",
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
							Description: "Outlet type as its NetBox slug, e.g. iec-60320-c13, nema-5-15r, CS8364C (any of NetBox's power outlet type choices).",
						},
						"power_port_template_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the power port template this outlet draws from.",
						},
						"power_port_id": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use power_port_template_id instead.",
							Description:        "Deprecated alias of power_port_template_id.",
						},
						"feed_leg": schema.StringAttribute{
							Computed:    true,
							Description: "Phase of the upstream feed this outlet is mapped to. One of: A, B, C.",
						},
						"color_hex": schema.StringAttribute{
							Computed:    true,
							Description: "RGB color in hex (e.g. ff0000).",
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

func (d *powerOutletTemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *powerOutletTemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data powerOutletTemplatesDataSourceModel

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
	var responseDTOs []*netboxapi.PowerOutletTemplateResponseDTO

	params := dcim.NewDcimPowerOutletTemplatesListParams()
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
		"type":                netboxapi.FilterString(params.SetType),
		"type__n":             netboxapi.FilterString(params.SetTypen),
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
		"feed_leg":            netboxapi.FilterStrings(params.SetFeedLeg),
		"feed_leg__n":         netboxapi.FilterStrings(params.SetFeedLegn),
		"feed_leg__ic":        netboxapi.FilterStrings(params.SetFeedLegIc),
		"feed_leg__nic":       netboxapi.FilterStrings(params.SetFeedLegNic),
		"feed_leg__ie":        netboxapi.FilterStrings(params.SetFeedLegIe),
		"feed_leg__nie":       netboxapi.FilterStrings(params.SetFeedLegNie),
		"feed_leg__isw":       netboxapi.FilterStrings(params.SetFeedLegIsw),
		"feed_leg__nisw":      netboxapi.FilterStrings(params.SetFeedLegNisw),
		"feed_leg__iew":       netboxapi.FilterStrings(params.SetFeedLegIew),
		"feed_leg__niew":      netboxapi.FilterStrings(params.SetFeedLegNiew),
		"feed_leg__empty":     netboxapi.FilterBool(params.SetFeedLegEmpty),
		"feed_leg__regex":     netboxapi.FilterStrings(params.SetFeedLegRegex),
		"feed_leg__iregex":    netboxapi.FilterStrings(params.SetFeedLegIregex),
		"module_type_id":      netboxapi.FilterInt64s(params.SetModuleTypeID),
		"module_type_id__n":   netboxapi.FilterInt64s(params.SetModuleTypeIDn),
		"color":               netboxapi.FilterStrings(params.SetColor),
		"color__n":            netboxapi.FilterStrings(params.SetColorn),
		"color__ic":           netboxapi.FilterStrings(params.SetColorIc),
		"color__nic":          netboxapi.FilterStrings(params.SetColorNic),
		"color__ie":           netboxapi.FilterStrings(params.SetColorIe),
		"color__nie":          netboxapi.FilterStrings(params.SetColorNie),
		"color__isw":          netboxapi.FilterStrings(params.SetColorIsw),
		"color__nisw":         netboxapi.FilterStrings(params.SetColorNisw),
		"color__iew":          netboxapi.FilterStrings(params.SetColorIew),
		"color__niew":         netboxapi.FilterStrings(params.SetColorNiew),
		"color__empty":        netboxapi.FilterBool(params.SetColorEmpty),
		"color__regex":        netboxapi.FilterStrings(params.SetColorRegex),
		"color__iregex":       netboxapi.FilterStrings(params.SetColorIregex),
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
		"power_port_id":       netboxapi.FilterInt64s(params.SetPowerPortID),
		"power_port_id__n":    netboxapi.FilterInt64s(params.SetPowerPortIDn),
		"q":                   netboxapi.FilterString(params.SetQ),
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
		res, err := d.client.Dcim.DcimPowerOutletTemplatesListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_power_outlet_templates", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.PowerOutletTemplateResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]powerOutletTemplateResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model powerOutletTemplateResourceModel
		resp.Diagnostics.Append(flattenPowerOutletTemplate(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: powerOutletTemplateResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
