// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualMachineTypesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualMachineTypesDataSource)(nil)
)

// NewVirtualMachineTypesDataSource returns a new virtual_machine_types data source, which
// lists virtual_machine_type objects matching its filters.
func NewVirtualMachineTypesDataSource() datasource.DataSource {
	return &virtualMachineTypesDataSource{}
}

type virtualMachineTypesDataSource struct {
	client *netboxapi.Client
}

// virtualMachineTypesDataSourceModel is the data source model: the filters, the limit and the matching
// virtual_machine_types.
type virtualMachineTypesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"virtual_machine_types"`
}

// virtualMachineTypeResourceAttrTypes is the attribute type map of virtualMachineTypeResourceModel.
func virtualMachineTypeResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                    types.Int64Type,
		"name":                  types.StringType,
		"slug":                  types.StringType,
		"default_platform_id":   types.Int64Type,
		"default_vcpus":         types.Float64Type,
		"default_memory":        types.Int64Type,
		"description":           types.StringType,
		"comments":              types.StringType,
		"owner_id":              types.Int64Type,
		"created":               types.StringType,
		"last_updated":          types.StringType,
		"url":                   types.StringType,
		"virtual_machine_count": types.Int64Type,
		"tags":                  types.SetType{ElemType: types.StringType},
		"tags_all":              types.SetType{ElemType: types.StringType},
		"custom_fields":         types.MapType{ElemType: types.StringType},
	}
}

func (d *virtualMachineTypesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_machine_types"
}

func (d *virtualMachineTypesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Virtualization:Lists virtual_machine_type objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, default_memory, default_memory__empty, default_memory__gt, default_memory__gte, default_memory__lt, default_memory__lte, default_memory__n, default_platform, default_platform__n, default_platform_id, default_platform_id__n, default_vcpus, default_vcpus__empty, default_vcpus__gt, default_vcpus__gte, default_vcpus__lt, default_vcpus__lte, default_vcpus__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, slug, slug__empty, slug__ic, slug__ie, slug__iew, slug__iregex, slug__isw, slug__n, slug__nic, slug__nie, slug__niew, slug__nisw, slug__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, virtual_machine_count, virtual_machine_count__empty, virtual_machine_count__gt, virtual_machine_count__gte, virtual_machine_count__lt, virtual_machine_count__lte, virtual_machine_count__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"virtual_machine_types": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching virtual_machine_types.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"slug": schema.StringAttribute{
							Computed:    true,
							Description: "URL-friendly unique shorthand; derived from name when not set.",
						},
						"default_platform_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the platform virtual machines of this type default to.",
						},
						"default_vcpus": schema.Float64Attribute{
							Computed:    true,
							Description: "Default number of virtual CPUs.",
						},
						"default_memory": schema.Int64Attribute{
							Computed:    true,
							Description: "Default memory in megabytes. Must be a whole number.",
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
						"virtual_machine_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of virtual machines of the type.",
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

func (d *virtualMachineTypesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualMachineTypesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualMachineTypesDataSourceModel

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
	var responseDTOs []*netboxapi.VirtualMachineTypeResponseDTO

	params := virtualization.NewVirtualizationVirtualMachineTypesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                           netboxapi.FilterInt64s(params.SetID),
		"id__n":                        netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                       netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                      netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                       netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                      netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                    netboxapi.FilterBool(params.SetIDEmpty),
		"name":                         netboxapi.FilterStrings(params.SetName),
		"name__n":                      netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                     netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                    netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                     netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                    netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                    netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                   netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                    netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                   netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":                  netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":                  netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":                 netboxapi.FilterStrings(params.SetNameIregex),
		"slug":                         netboxapi.FilterStrings(params.SetSlug),
		"slug__n":                      netboxapi.FilterStrings(params.SetSlugn),
		"slug__ic":                     netboxapi.FilterStrings(params.SetSlugIc),
		"slug__nic":                    netboxapi.FilterStrings(params.SetSlugNic),
		"slug__ie":                     netboxapi.FilterStrings(params.SetSlugIe),
		"slug__nie":                    netboxapi.FilterStrings(params.SetSlugNie),
		"slug__isw":                    netboxapi.FilterStrings(params.SetSlugIsw),
		"slug__nisw":                   netboxapi.FilterStrings(params.SetSlugNisw),
		"slug__iew":                    netboxapi.FilterStrings(params.SetSlugIew),
		"slug__niew":                   netboxapi.FilterStrings(params.SetSlugNiew),
		"slug__empty":                  netboxapi.FilterBool(params.SetSlugEmpty),
		"slug__regex":                  netboxapi.FilterStrings(params.SetSlugRegex),
		"slug__iregex":                 netboxapi.FilterStrings(params.SetSlugIregex),
		"default_platform_id":          netboxapi.FilterStrings(params.SetDefaultPlatformID),
		"default_platform_id__n":       netboxapi.FilterStrings(params.SetDefaultPlatformIDn),
		"owner_id":                     netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                  netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                          netboxapi.FilterStrings(params.SetTag),
		"tag__n":                       netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                     netboxapi.FilterStrings(params.SetTagAny),
		"created":                      netboxapi.FilterStrings(params.SetCreated),
		"created__n":                   netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                  netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":                 netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                  netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":                 netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":               netboxapi.FilterStrings(params.SetCreatedEmpty),
		"default_memory":               netboxapi.FilterInt64s(params.SetDefaultMemory),
		"default_memory__n":            netboxapi.FilterInt64s(params.SetDefaultMemoryn),
		"default_memory__lt":           netboxapi.FilterInt64s(params.SetDefaultMemoryLt),
		"default_memory__lte":          netboxapi.FilterInt64s(params.SetDefaultMemoryLte),
		"default_memory__gt":           netboxapi.FilterInt64s(params.SetDefaultMemoryGt),
		"default_memory__gte":          netboxapi.FilterInt64s(params.SetDefaultMemoryGte),
		"default_memory__empty":        netboxapi.FilterBool(params.SetDefaultMemoryEmpty),
		"default_platform":             netboxapi.FilterStrings(params.SetDefaultPlatform),
		"default_platform__n":          netboxapi.FilterStrings(params.SetDefaultPlatformn),
		"default_vcpus":                netboxapi.FilterFloat64s(params.SetDefaultVcpus),
		"default_vcpus__n":             netboxapi.FilterFloat64s(params.SetDefaultVcpusn),
		"default_vcpus__lt":            netboxapi.FilterFloat64s(params.SetDefaultVcpusLt),
		"default_vcpus__lte":           netboxapi.FilterFloat64s(params.SetDefaultVcpusLte),
		"default_vcpus__gt":            netboxapi.FilterFloat64s(params.SetDefaultVcpusGt),
		"default_vcpus__gte":           netboxapi.FilterFloat64s(params.SetDefaultVcpusGte),
		"default_vcpus__empty":         netboxapi.FilterBool(params.SetDefaultVcpusEmpty),
		"description":                  netboxapi.FilterStrings(params.SetDescription),
		"description__n":               netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":              netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":             netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":              netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":             netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":             netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":            netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":             netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":            netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":           netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":           netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":          netboxapi.FilterStrings(params.SetDescriptionIregex),
		"last_updated":                 netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":              netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":             netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":            netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":             netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":            netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":          netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                        netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                     netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                  netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":               netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":               netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":            netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                            netboxapi.FilterString(params.SetQ),
		"tag_id":                       netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                    netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                  netboxapi.FilterInt64s(params.SetTagIDAny),
		"virtual_machine_count":        netboxapi.FilterInt64s(params.SetVirtualMachineCount),
		"virtual_machine_count__n":     netboxapi.FilterInt64s(params.SetVirtualMachineCountn),
		"virtual_machine_count__lt":    netboxapi.FilterInt64s(params.SetVirtualMachineCountLt),
		"virtual_machine_count__lte":   netboxapi.FilterInt64s(params.SetVirtualMachineCountLte),
		"virtual_machine_count__gt":    netboxapi.FilterInt64s(params.SetVirtualMachineCountGt),
		"virtual_machine_count__gte":   netboxapi.FilterInt64s(params.SetVirtualMachineCountGte),
		"virtual_machine_count__empty": netboxapi.FilterBool(params.SetVirtualMachineCountEmpty),
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
		res, err := d.client.Virtualization.VirtualizationVirtualMachineTypesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_virtual_machine_types", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VirtualMachineTypeResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]virtualMachineTypeResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model virtualMachineTypeResourceModel
		resp.Diagnostics.Append(flattenVirtualMachineType(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: virtualMachineTypeResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
