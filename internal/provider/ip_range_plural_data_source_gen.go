// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*ipRangesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ipRangesDataSource)(nil)
)

// NewIPRangesDataSource returns a new ip_ranges data source, which
// lists ip_range objects matching its filters.
func NewIPRangesDataSource() datasource.DataSource {
	return &ipRangesDataSource{}
}

type ipRangesDataSource struct {
	client *netboxapi.Client
}

// ipRangesDataSourceModel is the data source model: the filters, the limit and the matching
// ip_ranges.
type ipRangesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"ip_ranges"`
}

// ipRangeResourceAttrTypes is the attribute type map of ipRangeResourceModel.
func ipRangeResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             types.Int64Type,
		"start_address":  conv.IPAddressType{},
		"end_address":    conv.IPAddressType{},
		"vrf_id":         types.Int64Type,
		"tenant_id":      types.Int64Type,
		"role_id":        types.Int64Type,
		"status":         types.StringType,
		"mark_utilized":  types.BoolType,
		"mark_populated": types.BoolType,
		"size":           types.Int64Type,
		"family":         types.Int64Type,
		"description":    types.StringType,
		"comments":       types.StringType,
		"owner_id":       types.Int64Type,
		"created":        types.StringType,
		"last_updated":   types.StringType,
		"url":            types.StringType,
		"tags":           types.SetType{ElemType: types.StringType},
		"tags_all":       types.SetType{ElemType: types.StringType},
		"custom_fields":  types.MapType{ElemType: types.StringType},
	}
}

func (d *ipRangesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ranges"
}

func (d *ipRangesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists ip_range objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, contains, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, end_address, family, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mark_populated, mark_utilized, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, parent, q, role, role__n, role_id, role_id__n, size, size__empty, size__gt, size__gte, size__lt, size__lte, size__n, start_address, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, vrf, vrf__n, vrf_id, vrf_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"ip_ranges": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching ip_ranges.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"start_address": schema.StringAttribute{
							CustomType:  conv.IPAddressType{},
							Computed:    true,
							Description: "First address of the range with its mask, e.g. 10.0.0.1/24.",
						},
						"end_address": schema.StringAttribute{
							CustomType:  conv.IPAddressType{},
							Computed:    true,
							Description: "Last address of the range with its mask, e.g. 10.0.0.50/24.",
						},
						"vrf_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VRF.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the IPAM role.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: active, reserved, deprecated.",
						},
						"mark_utilized": schema.BoolAttribute{
							Computed:    true,
							Description: "Treat the range as fully utilized.",
						},
						"mark_populated": schema.BoolAttribute{
							Computed:    true,
							Description: "Treat the range as fully populated.",
						},
						"size": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of addresses in the range.",
						},
						"family": schema.Int64Attribute{
							Computed:    true,
							Description: "Address family: 4 or 6.",
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

func (d *ipRangesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ipRangesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ipRangesDataSourceModel

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
	var responseDTOs []*netboxapi.IPRangeResponseDTO

	params := ipam.NewIpamIPRangesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"start_address":       netboxapi.FilterStrings(params.SetStartAddress),
		"end_address":         netboxapi.FilterStrings(params.SetEndAddress),
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
		"role_id":             netboxapi.FilterInt64s(params.SetRoleID),
		"role_id__n":          netboxapi.FilterInt64s(params.SetRoleIDn),
		"tenant_id":           netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":        netboxapi.FilterInt64s(params.SetTenantIDn),
		"vrf_id":              netboxapi.FilterInt64s(params.SetVrfID),
		"vrf_id__n":           netboxapi.FilterInt64s(params.SetVrfIDn),
		"mark_utilized":       netboxapi.FilterBool(params.SetMarkUtilized),
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
		"contains":            netboxapi.FilterString(params.SetContains),
		"role":                netboxapi.FilterStrings(params.SetRole),
		"role__n":             netboxapi.FilterStrings(params.SetRolen),
		"tenant":              netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":           netboxapi.FilterStrings(params.SetTenantn),
		"vrf":                 netboxapi.FilterStrings(params.SetVrf),
		"vrf__n":              netboxapi.FilterStrings(params.SetVrfn),
		"owner_id":            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":         netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":            netboxapi.FilterStrings(params.SetTagAny),
		"contact":             netboxapi.FilterInt64s(params.SetContact),
		"contact__n":          netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":       netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":    netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":        netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":     netboxapi.FilterInt64s(params.SetContactRolen),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"family":              netboxapi.FilterFloat64(params.SetFamily),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"mark_populated":      netboxapi.FilterBool(params.SetMarkPopulated),
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"parent":              netboxapi.FilterStrings(params.SetParent),
		"q":                   netboxapi.FilterString(params.SetQ),
		"size":                netboxapi.FilterInt64s(params.SetSize),
		"size__n":             netboxapi.FilterInt64s(params.SetSizen),
		"size__lt":            netboxapi.FilterInt64s(params.SetSizeLt),
		"size__lte":           netboxapi.FilterInt64s(params.SetSizeLte),
		"size__gt":            netboxapi.FilterInt64s(params.SetSizeGt),
		"size__gte":           netboxapi.FilterInt64s(params.SetSizeGte),
		"size__empty":         netboxapi.FilterBool(params.SetSizeEmpty),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant_group":        netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":     netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":     netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":  netboxapi.FilterStrings(params.SetTenantGroupIDn),
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
		res, err := d.client.Ipam.IpamIPRangesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_ip_ranges", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.IPRangeResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]ipRangeResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model ipRangeResourceModel
		resp.Diagnostics.Append(flattenIPRange(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ipRangeResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
