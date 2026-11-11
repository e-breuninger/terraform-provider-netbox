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
	_ datasource.DataSource              = (*vlansDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlansDataSource)(nil)
)

// NewVlansDataSource returns a new vlans data source, which
// lists vlan objects matching its filters.
func NewVlansDataSource() datasource.DataSource {
	return &vlansDataSource{}
}

type vlansDataSource struct {
	client *netboxapi.Client
}

// vlansDataSourceModel is the data source model: the filters, the limit and the matching
// vlans.
type vlansDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"vlans"`
}

// vlanResourceAttrTypes is the attribute type map of vlanResourceModel.
func vlanResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"name":          types.StringType,
		"vid":           types.Int64Type,
		"status":        types.StringType,
		"site_id":       types.Int64Type,
		"group_id":      types.Int64Type,
		"tenant_id":     types.Int64Type,
		"role_id":       types.Int64Type,
		"qinq_role":     types.StringType,
		"qinq_svlan_id": types.Int64Type,
		"description":   types.StringType,
		"comments":      types.StringType,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"prefix_count":  types.Int64Type,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *vlansDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlans"
}

func (d *vlansDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists vlan objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: available_at_site, available_at_site_group, available_on_device, available_on_virtualmachine, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_id, l2vpn, l2vpn__n, l2vpn_id, l2vpn_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, qinq_role, qinq_role__empty, qinq_role__ic, qinq_role__ie, qinq_role__iew, qinq_role__iregex, qinq_role__isw, qinq_role__n, qinq_role__nic, qinq_role__nie, qinq_role__niew, qinq_role__nisw, qinq_role__regex, qinq_svlan_id, qinq_svlan_id__n, qinq_svlan_vid, qinq_svlan_vid__empty, qinq_svlan_vid__gt, qinq_svlan_vid__gte, qinq_svlan_vid__lt, qinq_svlan_vid__lte, qinq_svlan_vid__n, region, region__n, region_id, region_id__n, role, role__n, role_id, role_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, vid, vid__empty, vid__gt, vid__gte, vid__lt, vid__lte, vid__n, vminterface_id. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"vlans": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching vlans.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"vid": schema.Int64Attribute{
							Computed:    true,
							Description: "VLAN id (1-4094). Must be a whole number.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: active, reserved, deprecated.",
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site.",
						},
						"group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VLAN group.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the role.",
						},
						"qinq_role": schema.StringAttribute{
							Computed:    true,
							Description: "Q-in-Q role: a service VLAN (svlan) carrying customer VLANs, or a customer VLAN (cvlan) inside one. One of: svlan, cvlan.",
						},
						"qinq_svlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the service VLAN this customer VLAN is carried in (qinq_role cvlan).",
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
						"prefix_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of prefixes on the VLAN.",
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

func (d *vlansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vlansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlansDataSourceModel

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
	var responseDTOs []*netboxapi.VlanResponseDTO

	params := ipam.NewIpamVlansListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                          netboxapi.FilterInt64s(params.SetID),
		"id__n":                       netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                      netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                     netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                      netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                     netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                   netboxapi.FilterBool(params.SetIDEmpty),
		"name":                        netboxapi.FilterStrings(params.SetName),
		"name__n":                     netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                    netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                   netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                    netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                   netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                   netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                  netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                   netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                  netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":                 netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":                 netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":                netboxapi.FilterStrings(params.SetNameIregex),
		"vid":                         netboxapi.FilterInt64s(params.SetVid),
		"vid__n":                      netboxapi.FilterInt64s(params.SetVidn),
		"vid__lt":                     netboxapi.FilterInt64s(params.SetVidLt),
		"vid__lte":                    netboxapi.FilterInt64s(params.SetVidLte),
		"vid__gt":                     netboxapi.FilterInt64s(params.SetVidGt),
		"vid__gte":                    netboxapi.FilterInt64s(params.SetVidGte),
		"vid__empty":                  netboxapi.FilterBool(params.SetVidEmpty),
		"status":                      netboxapi.FilterStrings(params.SetStatus),
		"status__n":                   netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":                  netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":                 netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":                  netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":                 netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":                 netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":                netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":                 netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":                netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":               netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":               netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":              netboxapi.FilterStrings(params.SetStatusIregex),
		"site_id":                     netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":                  netboxapi.FilterInt64s(params.SetSiteIDn),
		"group_id":                    netboxapi.FilterInt64s(params.SetGroupID),
		"group_id__n":                 netboxapi.FilterInt64s(params.SetGroupIDn),
		"tenant_id":                   netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":                netboxapi.FilterInt64s(params.SetTenantIDn),
		"group":                       netboxapi.FilterStrings(params.SetGroup),
		"group__n":                    netboxapi.FilterStrings(params.SetGroupn),
		"tenant":                      netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":                   netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":                netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":             netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":             netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":          netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"owner_id":                    netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                 netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                         netboxapi.FilterStrings(params.SetTag),
		"tag__n":                      netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                    netboxapi.FilterStrings(params.SetTagAny),
		"available_at_site":           netboxapi.FilterString(params.SetAvailableAtSite),
		"available_at_site_group":     netboxapi.FilterString(params.SetAvailableAtSiteGroup),
		"available_on_device":         netboxapi.FilterString(params.SetAvailableOnDevice),
		"available_on_virtualmachine": netboxapi.FilterString(params.SetAvailableOnVirtualmachine),
		"created":                     netboxapi.FilterStrings(params.SetCreated),
		"created__n":                  netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                 netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":                netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                 netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":                netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":              netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":                 netboxapi.FilterStrings(params.SetDescription),
		"description__n":              netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":             netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":            netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":             netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":            netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":            netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":           netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":            netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":           netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":          netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":          netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":         netboxapi.FilterStrings(params.SetDescriptionIregex),
		"interface_id":                netboxapi.FilterInt64(params.SetInterfaceID),
		"l2vpn":                       netboxapi.FilterInt64s(params.SetL2vpn),
		"l2vpn__n":                    netboxapi.FilterInt64s(params.SetL2vpnn),
		"l2vpn_id":                    netboxapi.FilterInt64s(params.SetL2vpnID),
		"l2vpn_id__n":                 netboxapi.FilterInt64s(params.SetL2vpnIDn),
		"last_updated":                netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":             netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":            netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":           netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":            netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":           netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":         netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                       netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                    netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                 netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":              netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":              netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":           netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                           netboxapi.FilterString(params.SetQ),
		"qinq_role":                   netboxapi.FilterStrings(params.SetQinqRole),
		"qinq_role__n":                netboxapi.FilterStrings(params.SetQinqRolen),
		"qinq_role__ic":               netboxapi.FilterStrings(params.SetQinqRoleIc),
		"qinq_role__nic":              netboxapi.FilterStrings(params.SetQinqRoleNic),
		"qinq_role__ie":               netboxapi.FilterStrings(params.SetQinqRoleIe),
		"qinq_role__nie":              netboxapi.FilterStrings(params.SetQinqRoleNie),
		"qinq_role__isw":              netboxapi.FilterStrings(params.SetQinqRoleIsw),
		"qinq_role__nisw":             netboxapi.FilterStrings(params.SetQinqRoleNisw),
		"qinq_role__iew":              netboxapi.FilterStrings(params.SetQinqRoleIew),
		"qinq_role__niew":             netboxapi.FilterStrings(params.SetQinqRoleNiew),
		"qinq_role__empty":            netboxapi.FilterBool(params.SetQinqRoleEmpty),
		"qinq_role__regex":            netboxapi.FilterStrings(params.SetQinqRoleRegex),
		"qinq_role__iregex":           netboxapi.FilterStrings(params.SetQinqRoleIregex),
		"qinq_svlan_id":               netboxapi.FilterInt64s(params.SetQinqSvlanID),
		"qinq_svlan_id__n":            netboxapi.FilterInt64s(params.SetQinqSvlanIDn),
		"qinq_svlan_vid":              netboxapi.FilterInt64s(params.SetQinqSvlanVid),
		"qinq_svlan_vid__n":           netboxapi.FilterInt64s(params.SetQinqSvlanVidn),
		"qinq_svlan_vid__lt":          netboxapi.FilterInt64s(params.SetQinqSvlanVidLt),
		"qinq_svlan_vid__lte":         netboxapi.FilterInt64s(params.SetQinqSvlanVidLte),
		"qinq_svlan_vid__gt":          netboxapi.FilterInt64s(params.SetQinqSvlanVidGt),
		"qinq_svlan_vid__gte":         netboxapi.FilterInt64s(params.SetQinqSvlanVidGte),
		"qinq_svlan_vid__empty":       netboxapi.FilterInt64s(params.SetQinqSvlanVidEmpty),
		"region":                      netboxapi.FilterStrings(params.SetRegion),
		"region__n":                   netboxapi.FilterStrings(params.SetRegionn),
		"region_id":                   netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":                netboxapi.FilterStrings(params.SetRegionIDn),
		"role":                        netboxapi.FilterStrings(params.SetRole),
		"role__n":                     netboxapi.FilterStrings(params.SetRolen),
		"role_id":                     netboxapi.FilterInt64s(params.SetRoleID),
		"role_id__n":                  netboxapi.FilterInt64s(params.SetRoleIDn),
		"site":                        netboxapi.FilterStrings(params.SetSite),
		"site__n":                     netboxapi.FilterStrings(params.SetSiten),
		"site_group":                  netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":               netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":               netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":            netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"tag_id":                      netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                   netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                 netboxapi.FilterInt64s(params.SetTagIDAny),
		"vminterface_id":              netboxapi.FilterInt64(params.SetVminterfaceID),
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
		res, err := d.client.Ipam.IpamVlansListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_vlans", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VlanResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]vlanResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model vlanResourceModel
		resp.Diagnostics.Append(flattenVlan(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: vlanResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
