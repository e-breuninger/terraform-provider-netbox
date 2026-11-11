// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*wirelessLinksDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*wirelessLinksDataSource)(nil)
)

// NewWirelessLinksDataSource returns a new wireless_links data source, which
// lists wireless_link objects matching its filters.
func NewWirelessLinksDataSource() datasource.DataSource {
	return &wirelessLinksDataSource{}
}

type wirelessLinksDataSource struct {
	client *netboxapi.Client
}

// wirelessLinksDataSourceModel is the data source model: the filters, the limit and the matching
// wireless_links.
type wirelessLinksDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"wireless_links"`
}

// wirelessLinkResourceAttrTypes is the attribute type map of wirelessLinkResourceModel.
func wirelessLinkResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             types.Int64Type,
		"interface_a_id": types.Int64Type,
		"interface_b_id": types.Int64Type,
		"ssid":           types.StringType,
		"status":         types.StringType,
		"tenant_id":      types.Int64Type,
		"auth_type":      types.StringType,
		"auth_cipher":    types.StringType,
		"auth_psk":       types.StringType,
		"distance":       types.Float64Type,
		"distance_unit":  types.StringType,
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

func (d *wirelessLinksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_links"
}

func (d *wirelessLinksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Wireless:Lists wireless_link objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: auth_cipher, auth_cipher__empty, auth_cipher__ic, auth_cipher__ie, auth_cipher__iew, auth_cipher__iregex, auth_cipher__isw, auth_cipher__n, auth_cipher__nic, auth_cipher__nie, auth_cipher__niew, auth_cipher__nisw, auth_cipher__regex, auth_psk, auth_psk__empty, auth_psk__ic, auth_psk__ie, auth_psk__iew, auth_psk__iregex, auth_psk__isw, auth_psk__n, auth_psk__nic, auth_psk__nie, auth_psk__niew, auth_psk__nisw, auth_psk__regex, auth_type, auth_type__empty, auth_type__ic, auth_type__ie, auth_type__iew, auth_type__iregex, auth_type__isw, auth_type__n, auth_type__nic, auth_type__nie, auth_type__niew, auth_type__nisw, auth_type__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, distance, distance__empty, distance__gt, distance__gte, distance__lt, distance__lte, distance__n, distance_unit, distance_unit__empty, distance_unit__ic, distance_unit__ie, distance_unit__iew, distance_unit__iregex, distance_unit__isw, distance_unit__n, distance_unit__nic, distance_unit__nie, distance_unit__niew, distance_unit__nisw, distance_unit__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interfacea_id, interfacea_id__n, interfaceb_id, interfaceb_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, ssid, ssid__empty, ssid__ic, ssid__ie, ssid__iew, ssid__iregex, ssid__isw, ssid__n, ssid__nic, ssid__nie, ssid__niew, ssid__nisw, ssid__regex, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"wireless_links": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching wireless_links.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"interface_a_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the first interface; it must be of a wireless type.",
						},
						"interface_b_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the second interface; it must be of a wireless type.",
						},
						"ssid": schema.StringAttribute{
							Computed:    true,
							Description: "Service set identifier.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Link status. One of: connected, planned, decommissioning.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"auth_type": schema.StringAttribute{
							Computed:    true,
							Description: "Authentication type. One of: open, wep, wpa-personal, wpa-enterprise.",
						},
						"auth_cipher": schema.StringAttribute{
							Computed:    true,
							Description: "Authentication cipher. One of: auto, tkip, aes.",
						},
						"auth_psk": schema.StringAttribute{
							Computed:    true,
							Sensitive:   true,
							Description: "Pre-shared key.",
						},
						"distance": schema.Float64Attribute{
							Computed:    true,
							Description: "Link distance; requires distance_unit.",
						},
						"distance_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of distance. One of: km, m, mi, ft.",
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

func (d *wirelessLinksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *wirelessLinksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data wirelessLinksDataSourceModel

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
	var responseDTOs []*netboxapi.WirelessLinkResponseDTO

	params := wireless.NewWirelessWirelessLinksListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"ssid":                  netboxapi.FilterStrings(params.SetSsid),
		"ssid__n":               netboxapi.FilterStrings(params.SetSsidn),
		"ssid__ic":              netboxapi.FilterStrings(params.SetSsidIc),
		"ssid__nic":             netboxapi.FilterStrings(params.SetSsidNic),
		"ssid__ie":              netboxapi.FilterStrings(params.SetSsidIe),
		"ssid__nie":             netboxapi.FilterStrings(params.SetSsidNie),
		"ssid__isw":             netboxapi.FilterStrings(params.SetSsidIsw),
		"ssid__nisw":            netboxapi.FilterStrings(params.SetSsidNisw),
		"ssid__iew":             netboxapi.FilterStrings(params.SetSsidIew),
		"ssid__niew":            netboxapi.FilterStrings(params.SetSsidNiew),
		"ssid__empty":           netboxapi.FilterBool(params.SetSsidEmpty),
		"ssid__regex":           netboxapi.FilterStrings(params.SetSsidRegex),
		"ssid__iregex":          netboxapi.FilterStrings(params.SetSsidIregex),
		"status":                netboxapi.FilterStrings(params.SetStatus),
		"status__n":             netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":            netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":           netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":            netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":           netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":           netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":          netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":           netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":          netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":         netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":         netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":        netboxapi.FilterStrings(params.SetStatusIregex),
		"tenant_id":             netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":          netboxapi.FilterInt64s(params.SetTenantIDn),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"interfacea_id":         netboxapi.FilterInt64s(params.SetInterfaceaID),
		"interfacea_id__n":      netboxapi.FilterInt64s(params.SetInterfaceaIDn),
		"interfaceb_id":         netboxapi.FilterInt64s(params.SetInterfacebID),
		"interfaceb_id__n":      netboxapi.FilterInt64s(params.SetInterfacebIDn),
		"auth_cipher":           netboxapi.FilterStrings(params.SetAuthCipher),
		"auth_cipher__n":        netboxapi.FilterStrings(params.SetAuthCiphern),
		"auth_cipher__ic":       netboxapi.FilterStrings(params.SetAuthCipherIc),
		"auth_cipher__nic":      netboxapi.FilterStrings(params.SetAuthCipherNic),
		"auth_cipher__ie":       netboxapi.FilterStrings(params.SetAuthCipherIe),
		"auth_cipher__nie":      netboxapi.FilterStrings(params.SetAuthCipherNie),
		"auth_cipher__isw":      netboxapi.FilterStrings(params.SetAuthCipherIsw),
		"auth_cipher__nisw":     netboxapi.FilterStrings(params.SetAuthCipherNisw),
		"auth_cipher__iew":      netboxapi.FilterStrings(params.SetAuthCipherIew),
		"auth_cipher__niew":     netboxapi.FilterStrings(params.SetAuthCipherNiew),
		"auth_cipher__empty":    netboxapi.FilterBool(params.SetAuthCipherEmpty),
		"auth_cipher__regex":    netboxapi.FilterStrings(params.SetAuthCipherRegex),
		"auth_cipher__iregex":   netboxapi.FilterStrings(params.SetAuthCipherIregex),
		"auth_psk":              netboxapi.FilterStrings(params.SetAuthPsk),
		"auth_psk__n":           netboxapi.FilterStrings(params.SetAuthPskn),
		"auth_psk__ic":          netboxapi.FilterStrings(params.SetAuthPskIc),
		"auth_psk__nic":         netboxapi.FilterStrings(params.SetAuthPskNic),
		"auth_psk__ie":          netboxapi.FilterStrings(params.SetAuthPskIe),
		"auth_psk__nie":         netboxapi.FilterStrings(params.SetAuthPskNie),
		"auth_psk__isw":         netboxapi.FilterStrings(params.SetAuthPskIsw),
		"auth_psk__nisw":        netboxapi.FilterStrings(params.SetAuthPskNisw),
		"auth_psk__iew":         netboxapi.FilterStrings(params.SetAuthPskIew),
		"auth_psk__niew":        netboxapi.FilterStrings(params.SetAuthPskNiew),
		"auth_psk__empty":       netboxapi.FilterBool(params.SetAuthPskEmpty),
		"auth_psk__regex":       netboxapi.FilterStrings(params.SetAuthPskRegex),
		"auth_psk__iregex":      netboxapi.FilterStrings(params.SetAuthPskIregex),
		"auth_type":             netboxapi.FilterStrings(params.SetAuthType),
		"auth_type__n":          netboxapi.FilterStrings(params.SetAuthTypen),
		"auth_type__ic":         netboxapi.FilterStrings(params.SetAuthTypeIc),
		"auth_type__nic":        netboxapi.FilterStrings(params.SetAuthTypeNic),
		"auth_type__ie":         netboxapi.FilterStrings(params.SetAuthTypeIe),
		"auth_type__nie":        netboxapi.FilterStrings(params.SetAuthTypeNie),
		"auth_type__isw":        netboxapi.FilterStrings(params.SetAuthTypeIsw),
		"auth_type__nisw":       netboxapi.FilterStrings(params.SetAuthTypeNisw),
		"auth_type__iew":        netboxapi.FilterStrings(params.SetAuthTypeIew),
		"auth_type__niew":       netboxapi.FilterStrings(params.SetAuthTypeNiew),
		"auth_type__empty":      netboxapi.FilterBool(params.SetAuthTypeEmpty),
		"auth_type__regex":      netboxapi.FilterStrings(params.SetAuthTypeRegex),
		"auth_type__iregex":     netboxapi.FilterStrings(params.SetAuthTypeIregex),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":           netboxapi.FilterStrings(params.SetDescription),
		"description__n":        netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":       netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":      netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":       netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":      netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":      netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":     netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":      netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":     netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":    netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":    netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":   netboxapi.FilterStrings(params.SetDescriptionIregex),
		"distance":              netboxapi.FilterFloat64s(params.SetDistance),
		"distance__n":           netboxapi.FilterFloat64s(params.SetDistancen),
		"distance__lt":          netboxapi.FilterFloat64s(params.SetDistanceLt),
		"distance__lte":         netboxapi.FilterFloat64s(params.SetDistanceLte),
		"distance__gt":          netboxapi.FilterFloat64s(params.SetDistanceGt),
		"distance__gte":         netboxapi.FilterFloat64s(params.SetDistanceGte),
		"distance__empty":       netboxapi.FilterBool(params.SetDistanceEmpty),
		"distance_unit":         netboxapi.FilterString(params.SetDistanceUnit),
		"distance_unit__n":      netboxapi.FilterString(params.SetDistanceUnitn),
		"distance_unit__ic":     netboxapi.FilterStrings(params.SetDistanceUnitIc),
		"distance_unit__nic":    netboxapi.FilterStrings(params.SetDistanceUnitNic),
		"distance_unit__ie":     netboxapi.FilterStrings(params.SetDistanceUnitIe),
		"distance_unit__nie":    netboxapi.FilterStrings(params.SetDistanceUnitNie),
		"distance_unit__isw":    netboxapi.FilterStrings(params.SetDistanceUnitIsw),
		"distance_unit__nisw":   netboxapi.FilterStrings(params.SetDistanceUnitNisw),
		"distance_unit__iew":    netboxapi.FilterStrings(params.SetDistanceUnitIew),
		"distance_unit__niew":   netboxapi.FilterStrings(params.SetDistanceUnitNiew),
		"distance_unit__empty":  netboxapi.FilterBool(params.SetDistanceUnitEmpty),
		"distance_unit__regex":  netboxapi.FilterStrings(params.SetDistanceUnitRegex),
		"distance_unit__iregex": netboxapi.FilterStrings(params.SetDistanceUnitIregex),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                 netboxapi.FilterStrings(params.SetOwner),
		"owner__n":              netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":           netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":        netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":        netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":     netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                     netboxapi.FilterString(params.SetQ),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":             netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":          netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":       netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":       netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":    netboxapi.FilterStrings(params.SetTenantGroupIDn),
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
		res, err := d.client.Wireless.WirelessWirelessLinksListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_wireless_links", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.WirelessLinkResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]wirelessLinkResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model wirelessLinkResourceModel
		resp.Diagnostics.Append(flattenWirelessLink(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: wirelessLinkResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
