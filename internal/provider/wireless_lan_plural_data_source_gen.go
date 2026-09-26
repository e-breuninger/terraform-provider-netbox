// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	_ datasource.DataSource              = (*wirelessLansDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*wirelessLansDataSource)(nil)
)

// NewWirelessLansDataSource returns a new wireless_lans data source, which
// lists wireless_lan objects matching its filters.
func NewWirelessLansDataSource() datasource.DataSource {
	return &wirelessLansDataSource{}
}

type wirelessLansDataSource struct {
	client *netboxapi.Client
}

// wirelessLansDataSourceModel is the data source model: the filters, the limit and the matching
// wireless_lans.
type wirelessLansDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"wireless_lans"`
}

// wirelessLanResourceAttrTypes is the attribute type map of wirelessLanResourceModel.
func wirelessLanResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"ssid":          types.StringType,
		"description":   types.StringType,
		"group_id":      types.Int64Type,
		"status":        types.StringType,
		"vlan_id":       types.Int64Type,
		"tenant_id":     types.Int64Type,
		"auth_type":     types.StringType,
		"auth_cipher":   types.StringType,
		"auth_psk":      types.StringType,
		"scope_type":    types.StringType,
		"scope_id":      types.Int64Type,
		"site_id":       types.Int64Type,
		"location_id":   types.Int64Type,
		"region_id":     types.Int64Type,
		"site_group_id": types.Int64Type,
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

func (d *wirelessLansDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_wireless_lans"
}

func (d *wirelessLansDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Wireless:Lists wireless_lan objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, owner_id, owner_id__n, ssid, ssid__empty, ssid__ic, ssid__ie, ssid__iew, ssid__iregex, ssid__isw, ssid__n, ssid__nic, ssid__nie, ssid__niew, ssid__nisw, ssid__regex, tag, tag__any, tag__n, tenant_id, tenant_id__n, vlan_id, vlan_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"wireless_lans": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching wireless_lans.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"ssid": schema.StringAttribute{
							Computed:    true,
							Description: "Service set identifier.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the wireless LAN group.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: active, reserved, disabled, deprecated.",
						},
						"vlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VLAN the wireless LAN is bridged to.",
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
						"scope_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the scope. Derived from site_id, location_id, region_id or site_group_id when one of those is set; set it together with scope_id otherwise. One of: dcim.site, dcim.location, dcim.region, dcim.sitegroup.",
						},
						"scope_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the scope object (see scope_type).",
						},
						"site_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site the object is scoped to (scope_type dcim.site). Conflicts with the other scope aliases and with setting the scope_* pair directly.",
						},
						"location_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the location the object is scoped to (scope_type dcim.location).",
						},
						"region_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the region the object is scoped to (scope_type dcim.region).",
						},
						"site_group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the site group the object is scoped to (scope_type dcim.sitegroup).",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *wirelessLansDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *wirelessLansDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data wirelessLansDataSourceModel

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

	fetchAll := false
	_ = fetchAll
	var responseDTOs []*netboxapi.WirelessLanResponseDTO

	params := wireless.NewWirelessWirelessLansListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vSsid []string
	var vSsidn []string
	var vSsidNic []string
	var vSsidIe []string
	var vSsidNie []string
	var vSsidIsw []string
	var vSsidNisw []string
	var vSsidIew []string
	var vSsidNiew []string
	var vSsidEmpty *bool
	var vSsidRegex []string
	var vSsidIregex []string
	var vGroupID []string
	var vGroupIDn []string
	var vVlanID []int64
	var vVlanIDn []int64
	var vTenantID []int64
	var vTenantIDn []int64
	var vSsidIc []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	customFieldQuery := url.Values{}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		// Custom fields are queried as cf_<field name> with the field's own filter logic; NetBox ignores
		// names it has no custom field for.
		if strings.HasPrefix(name, "cf_") {
			customFieldQuery.Add(name, value)
			continue
		}
		switch name {
		case "id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id' takes an integer, got %q.", value))
				return
			}
			vID = append(vID, v)
		case "id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__n' takes an integer, got %q.", value))
				return
			}
			vIDn = append(vIDn, v)
		case "id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lt' takes an integer, got %q.", value))
				return
			}
			vIDLt = append(vIDLt, v)
		case "id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lte' takes an integer, got %q.", value))
				return
			}
			vIDLte = append(vIDLte, v)
		case "id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gt' takes an integer, got %q.", value))
				return
			}
			vIDGt = append(vIDGt, v)
		case "id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gte' takes an integer, got %q.", value))
				return
			}
			vIDGte = append(vIDGte, v)
		case "id__empty":
			if vIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__empty' takes a boolean, got %q.", value))
				return
			}
			vIDEmpty = &v
		case "ssid":
			v := value
			vSsid = append(vSsid, v)
		case "ssid__n":
			v := value
			vSsidn = append(vSsidn, v)
		case "ssid__nic":
			v := value
			vSsidNic = append(vSsidNic, v)
		case "ssid__ie":
			v := value
			vSsidIe = append(vSsidIe, v)
		case "ssid__nie":
			v := value
			vSsidNie = append(vSsidNie, v)
		case "ssid__isw":
			v := value
			vSsidIsw = append(vSsidIsw, v)
		case "ssid__nisw":
			v := value
			vSsidNisw = append(vSsidNisw, v)
		case "ssid__iew":
			v := value
			vSsidIew = append(vSsidIew, v)
		case "ssid__niew":
			v := value
			vSsidNiew = append(vSsidNiew, v)
		case "ssid__empty":
			if vSsidEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ssid__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ssid__empty' takes a boolean, got %q.", value))
				return
			}
			vSsidEmpty = &v
		case "ssid__regex":
			v := value
			vSsidRegex = append(vSsidRegex, v)
		case "ssid__iregex":
			v := value
			vSsidIregex = append(vSsidIregex, v)
		case "group_id":
			v := value
			vGroupID = append(vGroupID, v)
		case "group_id__n":
			v := value
			vGroupIDn = append(vGroupIDn, v)
		case "vlan_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_id' takes an integer, got %q.", value))
				return
			}
			vVlanID = append(vVlanID, v)
		case "vlan_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_id__n' takes an integer, got %q.", value))
				return
			}
			vVlanIDn = append(vVlanIDn, v)
		case "tenant_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id' takes an integer, got %q.", value))
				return
			}
			vTenantID = append(vTenantID, v)
		case "tenant_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id__n' takes an integer, got %q.", value))
				return
			}
			vTenantIDn = append(vTenantIDn, v)
		case "ssid__ic":
			v := value
			vSsidIc = append(vSsidIc, v)
		case "owner_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id' takes an integer, got %q.", value))
				return
			}
			vOwnerID = append(vOwnerID, v)
		case "owner_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerIDn = append(vOwnerIDn, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
		default:
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
	}
	if len(vID) > 0 {
		params.SetID(vID)
	}
	if len(vIDn) > 0 {
		params.SetIDn(vIDn)
	}
	if len(vIDLt) > 0 {
		params.SetIDLt(vIDLt)
	}
	if len(vIDLte) > 0 {
		params.SetIDLte(vIDLte)
	}
	if len(vIDGt) > 0 {
		params.SetIDGt(vIDGt)
	}
	if len(vIDGte) > 0 {
		params.SetIDGte(vIDGte)
	}
	if vIDEmpty != nil {
		params.SetIDEmpty(vIDEmpty)
	}
	if len(vSsid) > 0 {
		params.SetSsid(vSsid)
	}
	if len(vSsidn) > 0 {
		params.SetSsidn(vSsidn)
	}
	if len(vSsidNic) > 0 {
		params.SetSsidNic(vSsidNic)
	}
	if len(vSsidIe) > 0 {
		params.SetSsidIe(vSsidIe)
	}
	if len(vSsidNie) > 0 {
		params.SetSsidNie(vSsidNie)
	}
	if len(vSsidIsw) > 0 {
		params.SetSsidIsw(vSsidIsw)
	}
	if len(vSsidNisw) > 0 {
		params.SetSsidNisw(vSsidNisw)
	}
	if len(vSsidIew) > 0 {
		params.SetSsidIew(vSsidIew)
	}
	if len(vSsidNiew) > 0 {
		params.SetSsidNiew(vSsidNiew)
	}
	if vSsidEmpty != nil {
		params.SetSsidEmpty(vSsidEmpty)
	}
	if len(vSsidRegex) > 0 {
		params.SetSsidRegex(vSsidRegex)
	}
	if len(vSsidIregex) > 0 {
		params.SetSsidIregex(vSsidIregex)
	}
	if len(vGroupID) > 0 {
		params.SetGroupID(vGroupID)
	}
	if len(vGroupIDn) > 0 {
		params.SetGroupIDn(vGroupIDn)
	}
	if len(vVlanID) > 0 {
		params.SetVlanID(vVlanID)
	}
	if len(vVlanIDn) > 0 {
		params.SetVlanIDn(vVlanIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vSsidIc) > 0 {
		params.SetSsidIc(vSsidIc)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
	}
	if len(vTag) > 0 {
		params.SetTag(vTag)
	}
	if len(vTagn) > 0 {
		params.SetTagn(vTagn)
	}
	if len(vTagAny) > 0 {
		params.SetTagAny(vTagAny)
	}
	pageSize := limit
	if fetchAll {
		pageSize = 1000
	}
	params.SetLimit(&pageSize)
	for offset := int64(0); ; {
		res, err := d.client.Wireless.WirelessWirelessLansListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_wireless_lans", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.WirelessLanResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if !fetchAll || len(res.Payload.Results) == 0 || res.Payload.Count == nil || offset >= *res.Payload.Count {
			break
		}
		params.SetOffset(&offset)
	}

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) runs on every item, as on the resource
	// itself (see the singular data source).
	resourceWithHooks := &wirelessLanResource{client: d.client}
	items := make([]wirelessLanResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model wirelessLanResourceModel
		resp.Diagnostics.Append(flattenWirelessLan(ctx, responseDTO, &model)...)
		resourceWithHooks.postRead(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: wirelessLanResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
