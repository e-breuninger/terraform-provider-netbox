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
	_ datasource.DataSource              = (*customFieldChoiceSetsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*customFieldChoiceSetsDataSource)(nil)
)

// NewCustomFieldChoiceSetsDataSource returns a new custom_field_choice_sets data source, which
// lists custom_field_choice_set objects matching its filters.
func NewCustomFieldChoiceSetsDataSource() datasource.DataSource {
	return &customFieldChoiceSetsDataSource{}
}

type customFieldChoiceSetsDataSource struct {
	client *netboxapi.Client
}

// customFieldChoiceSetsDataSourceModel is the data source model: the filters, the limit and the matching
// custom_field_choice_sets.
type customFieldChoiceSetsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"custom_field_choice_sets"`
}

// customFieldChoiceSetResourceAttrTypes is the attribute type map of customFieldChoiceSetResourceModel.
func customFieldChoiceSetResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                   types.Int64Type,
		"name":                 types.StringType,
		"description":          types.StringType,
		"base_choices":         types.StringType,
		"extra_choices":        types.ListType{ElemType: types.ObjectType{AttrTypes: customFieldChoiceSetExtraChoicesAttrTypes()}},
		"choice_colors":        types.MapType{ElemType: types.StringType},
		"order_alphabetically": types.BoolType,
		"choices_count":        types.Int64Type,
		"owner_id":             types.Int64Type,
		"created":              types.StringType,
		"last_updated":         types.StringType,
		"url":                  types.StringType,
	}
}

func (d *customFieldChoiceSetsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_custom_field_choice_sets"
}

func (d *customFieldChoiceSetsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists custom_field_choice_set objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: base_choices, base_choices__empty, base_choices__ic, base_choices__ie, base_choices__iew, base_choices__iregex, base_choices__isw, base_choices__n, base_choices__nic, base_choices__nie, base_choices__niew, base_choices__nisw, base_choices__regex, choice, choice_colors, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, order_alphabetically, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q. Repeating a name sends that parameter once per value.",
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
			"custom_field_choice_sets": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching custom_field_choice_sets.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id of the choice set.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"base_choices": schema.StringAttribute{
							Computed:    true,
							Description: "Predefined set of choices to start from. One of: IATA, ISO_3166, UN_LOCODE.",
						},
						"extra_choices": schema.ListNestedAttribute{
							NestedObject: schema.NestedAttributeObject{
								Attributes: map[string]schema.Attribute{
									"value": schema.StringAttribute{
										Computed: true,
									},
									"label": schema.StringAttribute{
										Computed: true,
									},
								},
							},
							Computed:    true,
							Description: "Additional choices as value/label pairs.",
						},
						"choice_colors": schema.MapAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Colors of choices, by choice value. One of blue, indigo, purple, pink, red, orange, yellow, green, teal, cyan, gray, black or white each.",
						},
						"order_alphabetically": schema.BoolAttribute{
							Computed:    true,
							Description: "Sort choices alphabetically instead of by definition order.",
						},
						"choices_count": schema.Int64Attribute{
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
					},
				},
			},
		},
	}
}

func (d *customFieldChoiceSetsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *customFieldChoiceSetsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data customFieldChoiceSetsDataSourceModel

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
	var responseDTOs []*netboxapi.CustomFieldChoiceSetResponseDTO

	params := extras.NewExtrasCustomFieldChoiceSetsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                   netboxapi.FilterInt64s(params.SetID),
		"id__n":                netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":               netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":              netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":               netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":              netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":            netboxapi.FilterBool(params.SetIDEmpty),
		"owner_id":             netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":          netboxapi.FilterInt64s(params.SetOwnerIDn),
		"base_choices":         netboxapi.FilterString(params.SetBaseChoices),
		"base_choices__n":      netboxapi.FilterString(params.SetBaseChoicesn),
		"base_choices__ic":     netboxapi.FilterStrings(params.SetBaseChoicesIc),
		"base_choices__nic":    netboxapi.FilterStrings(params.SetBaseChoicesNic),
		"base_choices__ie":     netboxapi.FilterStrings(params.SetBaseChoicesIe),
		"base_choices__nie":    netboxapi.FilterStrings(params.SetBaseChoicesNie),
		"base_choices__isw":    netboxapi.FilterStrings(params.SetBaseChoicesIsw),
		"base_choices__nisw":   netboxapi.FilterStrings(params.SetBaseChoicesNisw),
		"base_choices__iew":    netboxapi.FilterStrings(params.SetBaseChoicesIew),
		"base_choices__niew":   netboxapi.FilterStrings(params.SetBaseChoicesNiew),
		"base_choices__empty":  netboxapi.FilterBool(params.SetBaseChoicesEmpty),
		"base_choices__regex":  netboxapi.FilterStrings(params.SetBaseChoicesRegex),
		"base_choices__iregex": netboxapi.FilterStrings(params.SetBaseChoicesIregex),
		"choice":               netboxapi.FilterStrings(params.SetChoice),
		"choice_colors":        netboxapi.FilterStrings(params.SetChoiceColors),
		"created":              netboxapi.FilterStrings(params.SetCreated),
		"created__n":           netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":          netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":         netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":          netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":         netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":       netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":          netboxapi.FilterStrings(params.SetDescription),
		"description__n":       netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":      netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":     netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":      netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":     netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":     netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":    netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":     netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":    netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":   netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":   netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":  netboxapi.FilterStrings(params.SetDescriptionIregex),
		"last_updated":         netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":      netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":     netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":    netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":     netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":    netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":  netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
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
		"order_alphabetically": netboxapi.FilterBool(params.SetOrderAlphabetically),
		"owner":                netboxapi.FilterStrings(params.SetOwner),
		"owner__n":             netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":          netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":       netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":       netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":    netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                    netboxapi.FilterString(params.SetQ),
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
		res, err := d.client.Extras.ExtrasCustomFieldChoiceSetsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_custom_field_choice_sets", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CustomFieldChoiceSetResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]customFieldChoiceSetResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model customFieldChoiceSetResourceModel
		resp.Diagnostics.Append(flattenCustomFieldChoiceSet(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: customFieldChoiceSetResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
