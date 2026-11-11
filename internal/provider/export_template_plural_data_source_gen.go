// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*exportTemplatesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*exportTemplatesDataSource)(nil)
)

// NewExportTemplatesDataSource returns a new export_templates data source, which
// lists export_template objects matching its filters.
func NewExportTemplatesDataSource() datasource.DataSource {
	return &exportTemplatesDataSource{}
}

type exportTemplatesDataSource struct {
	client *netboxapi.Client
}

// exportTemplatesDataSourceModel is the data source model: the filters, the limit and the matching
// export_templates.
type exportTemplatesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"export_templates"`
}

// exportTemplateResourceAttrTypes is the attribute type map of exportTemplateResourceModel.
func exportTemplateResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                 types.Int64Type,
		"name":               types.StringType,
		"object_types":       types.SetType{ElemType: types.StringType},
		"template_code":      types.StringType,
		"mime_type":          types.StringType,
		"file_name":          types.StringType,
		"file_extension":     types.StringType,
		"as_attachment":      types.BoolType,
		"environment_params": jsontypes.NormalizedType{},
		"description":        types.StringType,
		"owner_id":           types.Int64Type,
		"created":            types.StringType,
		"last_updated":       types.StringType,
		"url":                types.StringType,
	}
}

func (d *exportTemplatesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_export_templates"
}

func (d *exportTemplatesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists export_template objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: as_attachment, auto_sync_enabled, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, data_file_id, data_file_id__n, data_source_id, data_source_id__n, data_synced, data_synced__empty, data_synced__gt, data_synced__gte, data_synced__lt, data_synced__lte, data_synced__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, file_extension, file_extension__empty, file_extension__ic, file_extension__ie, file_extension__iew, file_extension__iregex, file_extension__isw, file_extension__n, file_extension__nic, file_extension__nie, file_extension__niew, file_extension__nisw, file_extension__regex, file_name, file_name__empty, file_name__ic, file_name__ie, file_name__iew, file_name__iregex, file_name__isw, file_name__n, file_name__nic, file_name__nie, file_name__niew, file_name__nisw, file_name__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mime_type, mime_type__empty, mime_type__ic, mime_type__ie, mime_type__iew, mime_type__iregex, mime_type__isw, mime_type__n, mime_type__nic, mime_type__nie, mime_type__niew, mime_type__nisw, mime_type__regex, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, object_type, object_type__ic, object_type__ie, object_type__iew, object_type__iregex, object_type__isw, object_type__n, object_type__nic, object_type__nie, object_type__niew, object_type__nisw, object_type__regex, object_type_id, object_type_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q. Repeating a name sends that parameter once per value.",
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
			"export_templates": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching export_templates.",
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
							Description: "Object types the template can export, as app-labelled content types (dcim.device).",
						},
						"template_code": schema.StringAttribute{
							Computed:    true,
							Description: "Jinja2 template code; the queryset is available as {{ queryset }}.",
						},
						"mime_type": schema.StringAttribute{
							Computed:    true,
							Description: "MIME type of the rendered output, e.g. text/csv.",
						},
						"file_name": schema.StringAttribute{
							Computed:    true,
							Description: "File name of the rendered output, without extension.",
						},
						"file_extension": schema.StringAttribute{
							Computed:    true,
							Description: "Extension appended to the file name, e.g. csv.",
						},
						"as_attachment": schema.BoolAttribute{
							Computed:    true,
							Description: "Download the rendered output as a file instead of displaying it.",
						},
						"environment_params": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Parameters for the Jinja2 environment as JSON text (use jsonencode()).",
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
					},
				},
			},
		},
	}
}

func (d *exportTemplatesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *exportTemplatesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data exportTemplatesDataSourceModel

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
	var responseDTOs []*netboxapi.ExportTemplateResponseDTO

	params := extras.NewExtrasExportTemplatesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                     netboxapi.FilterInt64s(params.SetID),
		"id__n":                  netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                 netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                 netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":              netboxapi.FilterBool(params.SetIDEmpty),
		"name":                   netboxapi.FilterStrings(params.SetName),
		"name__n":                netboxapi.FilterStrings(params.SetNamen),
		"name__ic":               netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":              netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":               netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":              netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":              netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":             netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":              netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":             netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":            netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":            netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":           netboxapi.FilterStrings(params.SetNameIregex),
		"owner_id":               netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":            netboxapi.FilterInt64s(params.SetOwnerIDn),
		"as_attachment":          netboxapi.FilterBool(params.SetAsAttachment),
		"auto_sync_enabled":      netboxapi.FilterBool(params.SetAutoSyncEnabled),
		"created":                netboxapi.FilterStrings(params.SetCreated),
		"created__n":             netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":            netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":           netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":            netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":           netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":         netboxapi.FilterStrings(params.SetCreatedEmpty),
		"data_file_id":           netboxapi.FilterInt64s(params.SetDataFileID),
		"data_file_id__n":        netboxapi.FilterInt64s(params.SetDataFileIDn),
		"data_source_id":         netboxapi.FilterInt64s(params.SetDataSourceID),
		"data_source_id__n":      netboxapi.FilterInt64s(params.SetDataSourceIDn),
		"data_synced":            netboxapi.FilterStrings(params.SetDataSynced),
		"data_synced__n":         netboxapi.FilterStrings(params.SetDataSyncedn),
		"data_synced__lt":        netboxapi.FilterStrings(params.SetDataSyncedLt),
		"data_synced__lte":       netboxapi.FilterStrings(params.SetDataSyncedLte),
		"data_synced__gt":        netboxapi.FilterStrings(params.SetDataSyncedGt),
		"data_synced__gte":       netboxapi.FilterStrings(params.SetDataSyncedGte),
		"data_synced__empty":     netboxapi.FilterBool(params.SetDataSyncedEmpty),
		"description":            netboxapi.FilterStrings(params.SetDescription),
		"description__n":         netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":        netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":       netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":        netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":       netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":       netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":      netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":       netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":      netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":     netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":     netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":    netboxapi.FilterStrings(params.SetDescriptionIregex),
		"file_extension":         netboxapi.FilterStrings(params.SetFileExtension),
		"file_extension__n":      netboxapi.FilterStrings(params.SetFileExtensionn),
		"file_extension__ic":     netboxapi.FilterStrings(params.SetFileExtensionIc),
		"file_extension__nic":    netboxapi.FilterStrings(params.SetFileExtensionNic),
		"file_extension__ie":     netboxapi.FilterStrings(params.SetFileExtensionIe),
		"file_extension__nie":    netboxapi.FilterStrings(params.SetFileExtensionNie),
		"file_extension__isw":    netboxapi.FilterStrings(params.SetFileExtensionIsw),
		"file_extension__nisw":   netboxapi.FilterStrings(params.SetFileExtensionNisw),
		"file_extension__iew":    netboxapi.FilterStrings(params.SetFileExtensionIew),
		"file_extension__niew":   netboxapi.FilterStrings(params.SetFileExtensionNiew),
		"file_extension__empty":  netboxapi.FilterBool(params.SetFileExtensionEmpty),
		"file_extension__regex":  netboxapi.FilterStrings(params.SetFileExtensionRegex),
		"file_extension__iregex": netboxapi.FilterStrings(params.SetFileExtensionIregex),
		"file_name":              netboxapi.FilterStrings(params.SetFileName),
		"file_name__n":           netboxapi.FilterStrings(params.SetFileNamen),
		"file_name__ic":          netboxapi.FilterStrings(params.SetFileNameIc),
		"file_name__nic":         netboxapi.FilterStrings(params.SetFileNameNic),
		"file_name__ie":          netboxapi.FilterStrings(params.SetFileNameIe),
		"file_name__nie":         netboxapi.FilterStrings(params.SetFileNameNie),
		"file_name__isw":         netboxapi.FilterStrings(params.SetFileNameIsw),
		"file_name__nisw":        netboxapi.FilterStrings(params.SetFileNameNisw),
		"file_name__iew":         netboxapi.FilterStrings(params.SetFileNameIew),
		"file_name__niew":        netboxapi.FilterStrings(params.SetFileNameNiew),
		"file_name__empty":       netboxapi.FilterBool(params.SetFileNameEmpty),
		"file_name__regex":       netboxapi.FilterStrings(params.SetFileNameRegex),
		"file_name__iregex":      netboxapi.FilterStrings(params.SetFileNameIregex),
		"last_updated":           netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":        netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":       netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":      netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":       netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":      netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":    netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"mime_type":              netboxapi.FilterStrings(params.SetMimeType),
		"mime_type__n":           netboxapi.FilterStrings(params.SetMimeTypen),
		"mime_type__ic":          netboxapi.FilterStrings(params.SetMimeTypeIc),
		"mime_type__nic":         netboxapi.FilterStrings(params.SetMimeTypeNic),
		"mime_type__ie":          netboxapi.FilterStrings(params.SetMimeTypeIe),
		"mime_type__nie":         netboxapi.FilterStrings(params.SetMimeTypeNie),
		"mime_type__isw":         netboxapi.FilterStrings(params.SetMimeTypeIsw),
		"mime_type__nisw":        netboxapi.FilterStrings(params.SetMimeTypeNisw),
		"mime_type__iew":         netboxapi.FilterStrings(params.SetMimeTypeIew),
		"mime_type__niew":        netboxapi.FilterStrings(params.SetMimeTypeNiew),
		"mime_type__empty":       netboxapi.FilterBool(params.SetMimeTypeEmpty),
		"mime_type__regex":       netboxapi.FilterStrings(params.SetMimeTypeRegex),
		"mime_type__iregex":      netboxapi.FilterStrings(params.SetMimeTypeIregex),
		"object_type":            netboxapi.FilterStrings(params.SetObjectType),
		"object_type__n":         netboxapi.FilterStrings(params.SetObjectTypen),
		"object_type__ic":        netboxapi.FilterStrings(params.SetObjectTypeIc),
		"object_type__nic":       netboxapi.FilterStrings(params.SetObjectTypeNic),
		"object_type__ie":        netboxapi.FilterStrings(params.SetObjectTypeIe),
		"object_type__nie":       netboxapi.FilterStrings(params.SetObjectTypeNie),
		"object_type__isw":       netboxapi.FilterStrings(params.SetObjectTypeIsw),
		"object_type__nisw":      netboxapi.FilterStrings(params.SetObjectTypeNisw),
		"object_type__iew":       netboxapi.FilterStrings(params.SetObjectTypeIew),
		"object_type__niew":      netboxapi.FilterStrings(params.SetObjectTypeNiew),
		"object_type__regex":     netboxapi.FilterStrings(params.SetObjectTypeRegex),
		"object_type__iregex":    netboxapi.FilterStrings(params.SetObjectTypeIregex),
		"object_type_id":         netboxapi.FilterInt64s(params.SetObjectTypeID),
		"object_type_id__n":      netboxapi.FilterInt64s(params.SetObjectTypeIDn),
		"owner":                  netboxapi.FilterStrings(params.SetOwner),
		"owner__n":               netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":            netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":         netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":         netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":      netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                      netboxapi.FilterString(params.SetQ),
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
		res, err := d.client.Extras.ExtrasExportTemplatesListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_export_templates", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ExportTemplateResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]exportTemplateResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model exportTemplateResourceModel
		resp.Diagnostics.Append(flattenExportTemplate(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: exportTemplateResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
