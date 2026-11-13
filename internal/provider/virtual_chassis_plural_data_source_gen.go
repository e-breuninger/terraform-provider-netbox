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
	_ datasource.DataSource              = (*virtualChassisListDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualChassisListDataSource)(nil)
)

// NewVirtualChassisListDataSource returns a new virtual_chassis_list data source, which
// lists virtual_chassis objects matching its filters.
func NewVirtualChassisListDataSource() datasource.DataSource {
	return &virtualChassisListDataSource{}
}

type virtualChassisListDataSource struct {
	client *netboxapi.Client
}

// virtualChassisListDataSourceModel is the data source model: the filters, the limit and the matching
// virtual_chassis_list.
type virtualChassisListDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"virtual_chassis_list"`
}

// virtualChassisResourceAttrTypes is the attribute type map of virtualChassisResourceModel.
func virtualChassisResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"name":          types.StringType,
		"domain":        types.StringType,
		"master_id":     types.Int64Type,
		"description":   types.StringType,
		"comments":      types.StringType,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"member_count":  types.Int64Type,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *virtualChassisListDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_chassis_list"
}

func (d *virtualChassisListDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists virtual_chassis objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, domain, domain__empty, domain__ic, domain__ie, domain__iew, domain__iregex, domain__isw, domain__n, domain__nic, domain__nie, domain__niew, domain__nisw, domain__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, master, master__n, master_id, master_id__n, member_count, member_count__empty, member_count__gt, member_count__gte, member_count__lt, member_count__lte, member_count__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, region, region__n, region_id, region_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_id, tenant_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"virtual_chassis_list": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching virtual_chassis_list.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"domain": schema.StringAttribute{
							Computed: true,
						},
						"master_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device acting as chassis master (managed by netbox_virtual_chassis_master).",
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
						"member_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of members of the virtual chassis.",
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

func (d *virtualChassisListDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualChassisListDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualChassisListDataSourceModel

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
	var responseDTOs []*netboxapi.VirtualChassisResponseDTO

	params := dcim.NewDcimVirtualChassisListParams()
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
		"domain":              netboxapi.FilterStrings(params.SetDomain),
		"domain__n":           netboxapi.FilterStrings(params.SetDomainn),
		"domain__ic":          netboxapi.FilterStrings(params.SetDomainIc),
		"domain__nic":         netboxapi.FilterStrings(params.SetDomainNic),
		"domain__ie":          netboxapi.FilterStrings(params.SetDomainIe),
		"domain__nie":         netboxapi.FilterStrings(params.SetDomainNie),
		"domain__isw":         netboxapi.FilterStrings(params.SetDomainIsw),
		"domain__nisw":        netboxapi.FilterStrings(params.SetDomainNisw),
		"domain__iew":         netboxapi.FilterStrings(params.SetDomainIew),
		"domain__niew":        netboxapi.FilterStrings(params.SetDomainNiew),
		"domain__empty":       netboxapi.FilterBool(params.SetDomainEmpty),
		"domain__regex":       netboxapi.FilterStrings(params.SetDomainRegex),
		"domain__iregex":      netboxapi.FilterStrings(params.SetDomainIregex),
		"master_id":           netboxapi.FilterInt64s(params.SetMasterID),
		"master_id__n":        netboxapi.FilterInt64s(params.SetMasterIDn),
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
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"master":              netboxapi.FilterStrings(params.SetMaster),
		"master__n":           netboxapi.FilterStrings(params.SetMastern),
		"member_count":        netboxapi.FilterInt64s(params.SetMemberCount),
		"member_count__n":     netboxapi.FilterInt64s(params.SetMemberCountn),
		"member_count__lt":    netboxapi.FilterInt64s(params.SetMemberCountLt),
		"member_count__lte":   netboxapi.FilterInt64s(params.SetMemberCountLte),
		"member_count__gt":    netboxapi.FilterInt64s(params.SetMemberCountGt),
		"member_count__gte":   netboxapi.FilterInt64s(params.SetMemberCountGte),
		"member_count__empty": netboxapi.FilterBool(params.SetMemberCountEmpty),
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                   netboxapi.FilterString(params.SetQ),
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
		"tenant":              netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":           netboxapi.FilterStrings(params.SetTenantn),
		"tenant_id":           netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":        netboxapi.FilterInt64s(params.SetTenantIDn),
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
		res, err := d.client.Dcim.DcimVirtualChassisListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_virtual_chassis_list", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VirtualChassisResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]virtualChassisResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model virtualChassisResourceModel
		resp.Diagnostics.Append(flattenVirtualChassis(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: virtualChassisResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
