// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualCircuitsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualCircuitsDataSource)(nil)
)

// NewVirtualCircuitsDataSource returns a new virtual_circuits data source, which
// lists virtual_circuit objects matching its filters.
func NewVirtualCircuitsDataSource() datasource.DataSource {
	return &virtualCircuitsDataSource{}
}

type virtualCircuitsDataSource struct {
	client *netboxapi.Client
}

// virtualCircuitsDataSourceModel is the data source model: the filters, the limit and the matching
// virtual_circuits.
type virtualCircuitsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"virtual_circuits"`
}

// virtualCircuitResourceAttrTypes is the attribute type map of virtualCircuitResourceModel.
func virtualCircuitResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                          types.Int64Type,
		"cid":                         types.StringType,
		"circuit_provider_network_id": types.Int64Type,
		"virtual_circuit_type_id":     types.Int64Type,
		"circuit_provider_account_id": types.Int64Type,
		"status":                      types.StringType,
		"tenant_id":                   types.Int64Type,
		"description":                 types.StringType,
		"comments":                    types.StringType,
		"owner_id":                    types.Int64Type,
		"created":                     types.StringType,
		"last_updated":                types.StringType,
		"url":                         types.StringType,
		"tags":                        types.SetType{ElemType: types.StringType},
		"tags_all":                    types.SetType{ElemType: types.StringType},
		"custom_fields":               types.MapType{ElemType: types.StringType},
	}
}

func (d *virtualCircuitsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_circuits"
}

func (d *virtualCircuitsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:Lists virtual_circuit objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: cid, cid__empty, cid__ic, cid__ie, cid__iew, cid__iregex, cid__isw, cid__n, cid__nic, cid__nie, cid__niew, cid__nisw, cid__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, provider, provider__n, provider_account, provider_account__n, provider_account_id, provider_account_id__n, provider_id, provider_id__n, provider_network_id, provider_network_id__n, q, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, type, type__n, type_id, type_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"virtual_circuits": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching virtual_circuits.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"cid": schema.StringAttribute{
							Computed:    true,
							Description: "Circuit ID: unique identifier of the virtual circuit with its provider.",
						},
						"circuit_provider_network_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the provider network the virtual circuit is delivered over.",
						},
						"virtual_circuit_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual circuit type.",
						},
						"circuit_provider_account_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the provider account the virtual circuit is billed to.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: planned, provisioning, active, offline, deprovisioning, decommissioned.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant the virtual circuit belongs to.",
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

func (d *virtualCircuitsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualCircuitsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualCircuitsDataSourceModel

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
	var responseDTOs []*netboxapi.VirtualCircuitResponseDTO

	params := circuits.NewCircuitsVirtualCircuitsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                     netboxapi.FilterInt64s(params.SetID),
		"id__n":                  netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                 netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                 netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":              netboxapi.FilterBool(params.SetIDEmpty),
		"cid":                    netboxapi.FilterStrings(params.SetCid),
		"cid__n":                 netboxapi.FilterStrings(params.SetCidn),
		"cid__ic":                netboxapi.FilterStrings(params.SetCidIc),
		"cid__nic":               netboxapi.FilterStrings(params.SetCidNic),
		"cid__ie":                netboxapi.FilterStrings(params.SetCidIe),
		"cid__nie":               netboxapi.FilterStrings(params.SetCidNie),
		"cid__isw":               netboxapi.FilterStrings(params.SetCidIsw),
		"cid__nisw":              netboxapi.FilterStrings(params.SetCidNisw),
		"cid__iew":               netboxapi.FilterStrings(params.SetCidIew),
		"cid__niew":              netboxapi.FilterStrings(params.SetCidNiew),
		"cid__empty":             netboxapi.FilterBool(params.SetCidEmpty),
		"cid__regex":             netboxapi.FilterStrings(params.SetCidRegex),
		"cid__iregex":            netboxapi.FilterStrings(params.SetCidIregex),
		"status":                 netboxapi.FilterStrings(params.SetStatus),
		"status__n":              netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":             netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":            netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":             netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":            netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":            netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":           netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":            netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":           netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":          netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":          netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":         netboxapi.FilterStrings(params.SetStatusIregex),
		"provider_network_id":    netboxapi.FilterInt64s(params.SetProviderNetworkID),
		"provider_network_id__n": netboxapi.FilterInt64s(params.SetProviderNetworkIDn),
		"provider_account_id":    netboxapi.FilterInt64s(params.SetProviderAccountID),
		"provider_account_id__n": netboxapi.FilterInt64s(params.SetProviderAccountIDn),
		"owner_id":               netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":            netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                    netboxapi.FilterStrings(params.SetTag),
		"tag__n":                 netboxapi.FilterStrings(params.SetTagn),
		"tag__any":               netboxapi.FilterStrings(params.SetTagAny),
		"created":                netboxapi.FilterStrings(params.SetCreated),
		"created__n":             netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":            netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":           netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":            netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":           netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":         netboxapi.FilterStrings(params.SetCreatedEmpty),
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
		"last_updated":           netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":        netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":       netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":      netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":       netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":      netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":    netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                  netboxapi.FilterStrings(params.SetOwner),
		"owner__n":               netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":            netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":         netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":         netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":      netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"provider":               netboxapi.FilterStrings(params.SetProvider),
		"provider__n":            netboxapi.FilterStrings(params.SetProvidern),
		"provider_account":       netboxapi.FilterStrings(params.SetProviderAccount),
		"provider_account__n":    netboxapi.FilterStrings(params.SetProviderAccountn),
		"provider_id":            netboxapi.FilterInt64s(params.SetProviderID),
		"provider_id__n":         netboxapi.FilterInt64s(params.SetProviderIDn),
		"q":                      netboxapi.FilterString(params.SetQ),
		"tag_id":                 netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":              netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":            netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                 netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":              netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":           netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":        netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":        netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":     netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"tenant_id":              netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":           netboxapi.FilterInt64s(params.SetTenantIDn),
		"type":                   netboxapi.FilterStrings(params.SetType),
		"type__n":                netboxapi.FilterStrings(params.SetTypen),
		"type_id":                netboxapi.FilterInt64s(params.SetTypeID),
		"type_id__n":             netboxapi.FilterInt64s(params.SetTypeIDn),
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
		res, err := d.client.Circuits.CircuitsVirtualCircuitsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_virtual_circuits", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VirtualCircuitResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]virtualCircuitResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model virtualCircuitResourceModel
		resp.Diagnostics.Append(flattenVirtualCircuit(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: virtualCircuitResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
