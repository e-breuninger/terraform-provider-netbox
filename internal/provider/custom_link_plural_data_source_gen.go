// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*customLinksDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customLinksDataSource)(nil)
)

// NewCustomLinksDataSource returns a new custom_links data source, which
// lists custom_link objects matching its filters.
func NewCustomLinksDataSource() datasource.DataSource {
	return &customLinksDataSource{}
}

type customLinksDataSource struct {
	client *netboxapi.Client
}

// customLinksDataSourceModel is the data source model: the filters, the limit and the matching
// custom_links.
type customLinksDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"custom_links"`
}

// customLinkResourceAttrTypes is the attribute type map of customLinkResourceModel.
func customLinkResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":           types.Int64Type,
		"name":         types.StringType,
		"object_types": types.SetType{ElemType: types.StringType},
		"enabled":      types.BoolType,
		"link_text":    types.StringType,
		"link_url":     types.StringType,
		"weight":       types.Int64Type,
		"group_name":   types.StringType,
		"button_class": types.StringType,
		"new_window":   types.BoolType,
		"owner_id":     types.Int64Type,
		"created":      types.StringType,
		"last_updated": types.StringType,
		"url":          types.StringType,
	}
}

func (d *customLinksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_links"
}

func (d *customLinksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists custom_link objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: button_class, button_class__empty, button_class__ic, button_class__ie, button_class__iew, button_class__iregex, button_class__isw, button_class__n, button_class__nic, button_class__nie, button_class__niew, button_class__nisw, button_class__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, enabled, group_name, group_name__empty, group_name__ic, group_name__ie, group_name__iew, group_name__iregex, group_name__isw, group_name__n, group_name__nic, group_name__nie, group_name__niew, group_name__nisw, group_name__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, link_text, link_text__ic, link_text__ie, link_text__iew, link_text__iregex, link_text__isw, link_text__n, link_text__nic, link_text__nie, link_text__niew, link_text__nisw, link_text__regex, link_url, link_url__ic, link_url__ie, link_url__iew, link_url__iregex, link_url__isw, link_url__n, link_url__nic, link_url__nie, link_url__niew, link_url__nisw, link_url__regex, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, new_window, object_type, object_type__ic, object_type__ie, object_type__iew, object_type__iregex, object_type__isw, object_type__n, object_type__nic, object_type__nie, object_type__niew, object_type__nisw, object_type__regex, object_type_id, object_type_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, weight, weight__empty, weight__gt, weight__gte, weight__lt, weight__lte, weight__n. Repeating a name sends that parameter once per value.",
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
			"custom_links": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching custom_links.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"object_types": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Object types the link is shown on, as app-labelled content types (dcim.device).",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the link is rendered.",
						},
						"link_text": schema.StringAttribute{
							Computed:    true,
							Description: "Jinja2 template for the link text; the object is available as {{ object }}.",
						},
						"link_url": schema.StringAttribute{
							Computed:    true,
							Description: "Jinja2 template for the link URL; the object is available as {{ object }}.",
						},
						"weight": schema.Int64Attribute{
							Computed:    true,
							Description: "Order within the group. Must be a whole number.",
						},
						"group_name": schema.StringAttribute{
							Computed:    true,
							Description: "Links with the same group name are rendered as a dropdown.",
						},
						"button_class": schema.StringAttribute{
							Computed:    true,
							Description: "Button style. One of: default, blue, indigo, purple, pink, red, orange, yellow, green, teal, cyan, gray, black, white, ghost-dark.",
						},
						"new_window": schema.BoolAttribute{
							Computed:    true,
							Description: "Open the link in a new browser window.",
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
					},
				},
			},
		},
	}
}

func (d *customLinksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *customLinksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customLinksDataSourceModel

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
	var responseDTOs []*netboxapi.CustomLinkResponseDTO

	params := extras.NewExtrasCustomLinksListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                   netboxapi.FilterInt64s(params.SetID),
		"id__n":                netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":               netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":              netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":               netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":              netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":            netboxapi.FilterBool(params.SetIDEmpty),
		"name":                 netboxapi.FilterStrings(params.SetName),
		"name__n":              netboxapi.FilterStrings(params.SetNamen),
		"name__ic":             netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":            netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":             netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":            netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":            netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":           netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":            netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":           netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":          netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":          netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":         netboxapi.FilterStrings(params.SetNameIregex),
		"enabled":              netboxapi.FilterBool(params.SetEnabled),
		"owner_id":             netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":          netboxapi.FilterInt64s(params.SetOwnerIDn),
		"button_class":         netboxapi.FilterString(params.SetButtonClass),
		"button_class__n":      netboxapi.FilterString(params.SetButtonClassn),
		"button_class__ic":     netboxapi.FilterStrings(params.SetButtonClassIc),
		"button_class__nic":    netboxapi.FilterStrings(params.SetButtonClassNic),
		"button_class__ie":     netboxapi.FilterStrings(params.SetButtonClassIe),
		"button_class__nie":    netboxapi.FilterStrings(params.SetButtonClassNie),
		"button_class__isw":    netboxapi.FilterStrings(params.SetButtonClassIsw),
		"button_class__nisw":   netboxapi.FilterStrings(params.SetButtonClassNisw),
		"button_class__iew":    netboxapi.FilterStrings(params.SetButtonClassIew),
		"button_class__niew":   netboxapi.FilterStrings(params.SetButtonClassNiew),
		"button_class__empty":  netboxapi.FilterBool(params.SetButtonClassEmpty),
		"button_class__regex":  netboxapi.FilterStrings(params.SetButtonClassRegex),
		"button_class__iregex": netboxapi.FilterStrings(params.SetButtonClassIregex),
		"created":              netboxapi.FilterStrings(params.SetCreated),
		"created__n":           netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":          netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":         netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":          netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":         netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":       netboxapi.FilterStrings(params.SetCreatedEmpty),
		"group_name":           netboxapi.FilterStrings(params.SetGroupName),
		"group_name__n":        netboxapi.FilterStrings(params.SetGroupNamen),
		"group_name__ic":       netboxapi.FilterStrings(params.SetGroupNameIc),
		"group_name__nic":      netboxapi.FilterStrings(params.SetGroupNameNic),
		"group_name__ie":       netboxapi.FilterStrings(params.SetGroupNameIe),
		"group_name__nie":      netboxapi.FilterStrings(params.SetGroupNameNie),
		"group_name__isw":      netboxapi.FilterStrings(params.SetGroupNameIsw),
		"group_name__nisw":     netboxapi.FilterStrings(params.SetGroupNameNisw),
		"group_name__iew":      netboxapi.FilterStrings(params.SetGroupNameIew),
		"group_name__niew":     netboxapi.FilterStrings(params.SetGroupNameNiew),
		"group_name__empty":    netboxapi.FilterBool(params.SetGroupNameEmpty),
		"group_name__regex":    netboxapi.FilterStrings(params.SetGroupNameRegex),
		"group_name__iregex":   netboxapi.FilterStrings(params.SetGroupNameIregex),
		"last_updated":         netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":      netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":     netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":    netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":     netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":    netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":  netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"link_text":            netboxapi.FilterString(params.SetLinkText),
		"link_text__n":         netboxapi.FilterString(params.SetLinkTextn),
		"link_text__ic":        netboxapi.FilterString(params.SetLinkTextIc),
		"link_text__nic":       netboxapi.FilterString(params.SetLinkTextNic),
		"link_text__ie":        netboxapi.FilterString(params.SetLinkTextIe),
		"link_text__nie":       netboxapi.FilterString(params.SetLinkTextNie),
		"link_text__isw":       netboxapi.FilterString(params.SetLinkTextIsw),
		"link_text__nisw":      netboxapi.FilterString(params.SetLinkTextNisw),
		"link_text__iew":       netboxapi.FilterString(params.SetLinkTextIew),
		"link_text__niew":      netboxapi.FilterString(params.SetLinkTextNiew),
		"link_text__regex":     netboxapi.FilterString(params.SetLinkTextRegex),
		"link_text__iregex":    netboxapi.FilterString(params.SetLinkTextIregex),
		"link_url":             netboxapi.FilterString(params.SetLinkURL),
		"link_url__n":          netboxapi.FilterString(params.SetLinkURLn),
		"link_url__ic":         netboxapi.FilterString(params.SetLinkURLIc),
		"link_url__nic":        netboxapi.FilterString(params.SetLinkURLNic),
		"link_url__ie":         netboxapi.FilterString(params.SetLinkURLIe),
		"link_url__nie":        netboxapi.FilterString(params.SetLinkURLNie),
		"link_url__isw":        netboxapi.FilterString(params.SetLinkURLIsw),
		"link_url__nisw":       netboxapi.FilterString(params.SetLinkURLNisw),
		"link_url__iew":        netboxapi.FilterString(params.SetLinkURLIew),
		"link_url__niew":       netboxapi.FilterString(params.SetLinkURLNiew),
		"link_url__regex":      netboxapi.FilterString(params.SetLinkURLRegex),
		"link_url__iregex":     netboxapi.FilterString(params.SetLinkURLIregex),
		"new_window":           netboxapi.FilterBool(params.SetNewWindow),
		"object_type":          netboxapi.FilterStrings(params.SetObjectType),
		"object_type__n":       netboxapi.FilterStrings(params.SetObjectTypen),
		"object_type__ic":      netboxapi.FilterStrings(params.SetObjectTypeIc),
		"object_type__nic":     netboxapi.FilterStrings(params.SetObjectTypeNic),
		"object_type__ie":      netboxapi.FilterStrings(params.SetObjectTypeIe),
		"object_type__nie":     netboxapi.FilterStrings(params.SetObjectTypeNie),
		"object_type__isw":     netboxapi.FilterStrings(params.SetObjectTypeIsw),
		"object_type__nisw":    netboxapi.FilterStrings(params.SetObjectTypeNisw),
		"object_type__iew":     netboxapi.FilterStrings(params.SetObjectTypeIew),
		"object_type__niew":    netboxapi.FilterStrings(params.SetObjectTypeNiew),
		"object_type__regex":   netboxapi.FilterStrings(params.SetObjectTypeRegex),
		"object_type__iregex":  netboxapi.FilterStrings(params.SetObjectTypeIregex),
		"object_type_id":       netboxapi.FilterInt64s(params.SetObjectTypeID),
		"object_type_id__n":    netboxapi.FilterInt64s(params.SetObjectTypeIDn),
		"owner":                netboxapi.FilterStrings(params.SetOwner),
		"owner__n":             netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":          netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":       netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":       netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":    netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                    netboxapi.FilterString(params.SetQ),
		"weight":               netboxapi.FilterInt64s(params.SetWeight),
		"weight__n":            netboxapi.FilterInt64s(params.SetWeightn),
		"weight__lt":           netboxapi.FilterInt64s(params.SetWeightLt),
		"weight__lte":          netboxapi.FilterInt64s(params.SetWeightLte),
		"weight__gt":           netboxapi.FilterInt64s(params.SetWeightGt),
		"weight__gte":          netboxapi.FilterInt64s(params.SetWeightGte),
		"weight__empty":        netboxapi.FilterBool(params.SetWeightEmpty),
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
		res, err := d.client.Extras.ExtrasCustomLinksListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_custom_links", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CustomLinkResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]customLinkResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model customLinkResourceModel
		resp.Diagnostics.Append(flattenCustomLink(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: customLinkResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
