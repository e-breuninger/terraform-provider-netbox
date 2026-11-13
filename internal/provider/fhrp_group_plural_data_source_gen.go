// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

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
	_ datasource.DataSource              = (*fhrpGroupsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*fhrpGroupsDataSource)(nil)
)

// NewFhrpGroupsDataSource returns a new fhrp_groups data source, which
// lists fhrp_group objects matching its filters.
func NewFhrpGroupsDataSource() datasource.DataSource {
	return &fhrpGroupsDataSource{}
}

type fhrpGroupsDataSource struct {
	client *netboxapi.Client
}

// fhrpGroupsDataSourceModel is the data source model: the filters, the limit and the matching
// fhrp_groups.
type fhrpGroupsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"fhrp_groups"`
}

// fhrpGroupResourceAttrTypes is the attribute type map of fhrpGroupResourceModel.
func fhrpGroupResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"name":          types.StringType,
		"protocol":      types.StringType,
		"group_id":      types.Int64Type,
		"auth_type":     types.StringType,
		"auth_key":      types.StringType,
		"description":   types.StringType,
		"comments":      types.StringType,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *fhrpGroupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fhrp_groups"
}

func (d *fhrpGroupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists fhrp_group objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: auth_key, auth_key__empty, auth_key__ic, auth_key__ie, auth_key__iew, auth_key__iregex, auth_key__isw, auth_key__n, auth_key__nic, auth_key__nie, auth_key__niew, auth_key__nisw, auth_key__regex, auth_type, auth_type__empty, auth_type__ic, auth_type__ie, auth_type__iew, auth_type__iregex, auth_type__isw, auth_type__n, auth_type__nic, auth_type__nie, auth_type__niew, auth_type__nisw, auth_type__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, group_id, group_id__empty, group_id__gt, group_id__gte, group_id__lt, group_id__lte, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, protocol, protocol__empty, protocol__ic, protocol__ie, protocol__iew, protocol__iregex, protocol__isw, protocol__n, protocol__nic, protocol__nie, protocol__niew, protocol__nisw, protocol__regex, q, related_ip, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"fhrp_groups": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching fhrp_groups.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"protocol": schema.StringAttribute{
							Computed:    true,
							Description: "Redundancy protocol. One of: vrrp2, vrrp3, carp, clusterxl, hsrp, glbp, other.",
						},
						"group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Protocol group number (the VRRP/HSRP group id, not a reference to another object).",
						},
						"auth_type": schema.StringAttribute{
							Computed:    true,
							Description: "Authentication type. One of: plaintext, md5.",
						},
						"auth_key": schema.StringAttribute{
							Computed:    true,
							Sensitive:   true,
							Description: "Authentication key.",
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

func (d *fhrpGroupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *fhrpGroupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data fhrpGroupsDataSourceModel

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
	var responseDTOs []*netboxapi.FhrpGroupResponseDTO

	params := ipam.NewIpamFhrpGroupsListParams()
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
		"protocol":            netboxapi.FilterStrings(params.SetProtocol),
		"protocol__n":         netboxapi.FilterStrings(params.SetProtocoln),
		"protocol__ic":        netboxapi.FilterStrings(params.SetProtocolIc),
		"protocol__nic":       netboxapi.FilterStrings(params.SetProtocolNic),
		"protocol__ie":        netboxapi.FilterStrings(params.SetProtocolIe),
		"protocol__nie":       netboxapi.FilterStrings(params.SetProtocolNie),
		"protocol__isw":       netboxapi.FilterStrings(params.SetProtocolIsw),
		"protocol__nisw":      netboxapi.FilterStrings(params.SetProtocolNisw),
		"protocol__iew":       netboxapi.FilterStrings(params.SetProtocolIew),
		"protocol__niew":      netboxapi.FilterStrings(params.SetProtocolNiew),
		"protocol__empty":     netboxapi.FilterBool(params.SetProtocolEmpty),
		"protocol__regex":     netboxapi.FilterStrings(params.SetProtocolRegex),
		"protocol__iregex":    netboxapi.FilterStrings(params.SetProtocolIregex),
		"group_id":            netboxapi.FilterInt64s(params.SetGroupID),
		"group_id__n":         netboxapi.FilterInt64s(params.SetGroupIDn),
		"group_id__lt":        netboxapi.FilterInt64s(params.SetGroupIDLt),
		"group_id__lte":       netboxapi.FilterInt64s(params.SetGroupIDLte),
		"group_id__gt":        netboxapi.FilterInt64s(params.SetGroupIDGt),
		"group_id__gte":       netboxapi.FilterInt64s(params.SetGroupIDGte),
		"group_id__empty":     netboxapi.FilterBool(params.SetGroupIDEmpty),
		"auth_type":           netboxapi.FilterStrings(params.SetAuthType),
		"auth_type__n":        netboxapi.FilterStrings(params.SetAuthTypen),
		"auth_type__ic":       netboxapi.FilterStrings(params.SetAuthTypeIc),
		"auth_type__nic":      netboxapi.FilterStrings(params.SetAuthTypeNic),
		"auth_type__ie":       netboxapi.FilterStrings(params.SetAuthTypeIe),
		"auth_type__nie":      netboxapi.FilterStrings(params.SetAuthTypeNie),
		"auth_type__isw":      netboxapi.FilterStrings(params.SetAuthTypeIsw),
		"auth_type__nisw":     netboxapi.FilterStrings(params.SetAuthTypeNisw),
		"auth_type__iew":      netboxapi.FilterStrings(params.SetAuthTypeIew),
		"auth_type__niew":     netboxapi.FilterStrings(params.SetAuthTypeNiew),
		"auth_type__empty":    netboxapi.FilterBool(params.SetAuthTypeEmpty),
		"auth_type__regex":    netboxapi.FilterStrings(params.SetAuthTypeRegex),
		"auth_type__iregex":   netboxapi.FilterStrings(params.SetAuthTypeIregex),
		"owner_id":            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":         netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":            netboxapi.FilterStrings(params.SetTagAny),
		"auth_key":            netboxapi.FilterStrings(params.SetAuthKey),
		"auth_key__n":         netboxapi.FilterStrings(params.SetAuthKeyn),
		"auth_key__ic":        netboxapi.FilterStrings(params.SetAuthKeyIc),
		"auth_key__nic":       netboxapi.FilterStrings(params.SetAuthKeyNic),
		"auth_key__ie":        netboxapi.FilterStrings(params.SetAuthKeyIe),
		"auth_key__nie":       netboxapi.FilterStrings(params.SetAuthKeyNie),
		"auth_key__isw":       netboxapi.FilterStrings(params.SetAuthKeyIsw),
		"auth_key__nisw":      netboxapi.FilterStrings(params.SetAuthKeyNisw),
		"auth_key__iew":       netboxapi.FilterStrings(params.SetAuthKeyIew),
		"auth_key__niew":      netboxapi.FilterStrings(params.SetAuthKeyNiew),
		"auth_key__empty":     netboxapi.FilterBool(params.SetAuthKeyEmpty),
		"auth_key__regex":     netboxapi.FilterStrings(params.SetAuthKeyRegex),
		"auth_key__iregex":    netboxapi.FilterStrings(params.SetAuthKeyIregex),
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
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                   netboxapi.FilterString(params.SetQ),
		"related_ip":          netboxapi.FilterStrings(params.SetRelatedIP),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
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
		res, err := d.client.Ipam.IpamFhrpGroupsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_fhrp_groups", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.FhrpGroupResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]fhrpGroupResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model fhrpGroupResourceModel
		resp.Diagnostics.Append(flattenFhrpGroup(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: fhrpGroupResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
