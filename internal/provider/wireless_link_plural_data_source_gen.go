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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vSsid []string
	var vSsidn []string
	var vSsidIc []string
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
	var vStatus []string
	var vStatusn []string
	var vStatusIc []string
	var vStatusNic []string
	var vStatusIe []string
	var vStatusNie []string
	var vStatusIsw []string
	var vStatusNisw []string
	var vStatusIew []string
	var vStatusNiew []string
	var vStatusEmpty *bool
	var vStatusRegex []string
	var vStatusIregex []string
	var vTenantID []int64
	var vTenantIDn []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vInterfaceaID []int64
	var vInterfaceaIDn []int64
	var vInterfacebID []int64
	var vInterfacebIDn []int64
	var vAuthCipher []string
	var vAuthCiphern []string
	var vAuthCipherIc []string
	var vAuthCipherNic []string
	var vAuthCipherIe []string
	var vAuthCipherNie []string
	var vAuthCipherIsw []string
	var vAuthCipherNisw []string
	var vAuthCipherIew []string
	var vAuthCipherNiew []string
	var vAuthCipherEmpty *bool
	var vAuthCipherRegex []string
	var vAuthCipherIregex []string
	var vAuthPsk []string
	var vAuthPskn []string
	var vAuthPskIc []string
	var vAuthPskNic []string
	var vAuthPskIe []string
	var vAuthPskNie []string
	var vAuthPskIsw []string
	var vAuthPskNisw []string
	var vAuthPskIew []string
	var vAuthPskNiew []string
	var vAuthPskEmpty *bool
	var vAuthPskRegex []string
	var vAuthPskIregex []string
	var vAuthType []string
	var vAuthTypen []string
	var vAuthTypeIc []string
	var vAuthTypeNic []string
	var vAuthTypeIe []string
	var vAuthTypeNie []string
	var vAuthTypeIsw []string
	var vAuthTypeNisw []string
	var vAuthTypeIew []string
	var vAuthTypeNiew []string
	var vAuthTypeEmpty *bool
	var vAuthTypeRegex []string
	var vAuthTypeIregex []string
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vDescription []string
	var vDescriptionn []string
	var vDescriptionIc []string
	var vDescriptionNic []string
	var vDescriptionIe []string
	var vDescriptionNie []string
	var vDescriptionIsw []string
	var vDescriptionNisw []string
	var vDescriptionIew []string
	var vDescriptionNiew []string
	var vDescriptionEmpty *bool
	var vDescriptionRegex []string
	var vDescriptionIregex []string
	var vDistance []float64
	var vDistancen []float64
	var vDistanceLt []float64
	var vDistanceLte []float64
	var vDistanceGt []float64
	var vDistanceGte []float64
	var vDistanceEmpty *bool
	var vDistanceUnit *string
	var vDistanceUnitn *string
	var vDistanceUnitIc []string
	var vDistanceUnitNic []string
	var vDistanceUnitIe []string
	var vDistanceUnitNie []string
	var vDistanceUnitIsw []string
	var vDistanceUnitNisw []string
	var vDistanceUnitIew []string
	var vDistanceUnitNiew []string
	var vDistanceUnitEmpty *bool
	var vDistanceUnitRegex []string
	var vDistanceUnitIregex []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vQ *string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
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
		case "ssid__ic":
			v := value
			vSsidIc = append(vSsidIc, v)
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
		case "status":
			v := value
			vStatus = append(vStatus, v)
		case "status__n":
			v := value
			vStatusn = append(vStatusn, v)
		case "status__ic":
			v := value
			vStatusIc = append(vStatusIc, v)
		case "status__nic":
			v := value
			vStatusNic = append(vStatusNic, v)
		case "status__ie":
			v := value
			vStatusIe = append(vStatusIe, v)
		case "status__nie":
			v := value
			vStatusNie = append(vStatusNie, v)
		case "status__isw":
			v := value
			vStatusIsw = append(vStatusIsw, v)
		case "status__nisw":
			v := value
			vStatusNisw = append(vStatusNisw, v)
		case "status__iew":
			v := value
			vStatusIew = append(vStatusIew, v)
		case "status__niew":
			v := value
			vStatusNiew = append(vStatusNiew, v)
		case "status__empty":
			if vStatusEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'status__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'status__empty' takes a boolean, got %q.", value))
				return
			}
			vStatusEmpty = &v
		case "status__regex":
			v := value
			vStatusRegex = append(vStatusRegex, v)
		case "status__iregex":
			v := value
			vStatusIregex = append(vStatusIregex, v)
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
		case "interfacea_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interfacea_id' takes an integer, got %q.", value))
				return
			}
			vInterfaceaID = append(vInterfaceaID, v)
		case "interfacea_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interfacea_id__n' takes an integer, got %q.", value))
				return
			}
			vInterfaceaIDn = append(vInterfaceaIDn, v)
		case "interfaceb_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interfaceb_id' takes an integer, got %q.", value))
				return
			}
			vInterfacebID = append(vInterfacebID, v)
		case "interfaceb_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interfaceb_id__n' takes an integer, got %q.", value))
				return
			}
			vInterfacebIDn = append(vInterfacebIDn, v)
		case "auth_cipher":
			v := value
			vAuthCipher = append(vAuthCipher, v)
		case "auth_cipher__n":
			v := value
			vAuthCiphern = append(vAuthCiphern, v)
		case "auth_cipher__ic":
			v := value
			vAuthCipherIc = append(vAuthCipherIc, v)
		case "auth_cipher__nic":
			v := value
			vAuthCipherNic = append(vAuthCipherNic, v)
		case "auth_cipher__ie":
			v := value
			vAuthCipherIe = append(vAuthCipherIe, v)
		case "auth_cipher__nie":
			v := value
			vAuthCipherNie = append(vAuthCipherNie, v)
		case "auth_cipher__isw":
			v := value
			vAuthCipherIsw = append(vAuthCipherIsw, v)
		case "auth_cipher__nisw":
			v := value
			vAuthCipherNisw = append(vAuthCipherNisw, v)
		case "auth_cipher__iew":
			v := value
			vAuthCipherIew = append(vAuthCipherIew, v)
		case "auth_cipher__niew":
			v := value
			vAuthCipherNiew = append(vAuthCipherNiew, v)
		case "auth_cipher__empty":
			if vAuthCipherEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'auth_cipher__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'auth_cipher__empty' takes a boolean, got %q.", value))
				return
			}
			vAuthCipherEmpty = &v
		case "auth_cipher__regex":
			v := value
			vAuthCipherRegex = append(vAuthCipherRegex, v)
		case "auth_cipher__iregex":
			v := value
			vAuthCipherIregex = append(vAuthCipherIregex, v)
		case "auth_psk":
			v := value
			vAuthPsk = append(vAuthPsk, v)
		case "auth_psk__n":
			v := value
			vAuthPskn = append(vAuthPskn, v)
		case "auth_psk__ic":
			v := value
			vAuthPskIc = append(vAuthPskIc, v)
		case "auth_psk__nic":
			v := value
			vAuthPskNic = append(vAuthPskNic, v)
		case "auth_psk__ie":
			v := value
			vAuthPskIe = append(vAuthPskIe, v)
		case "auth_psk__nie":
			v := value
			vAuthPskNie = append(vAuthPskNie, v)
		case "auth_psk__isw":
			v := value
			vAuthPskIsw = append(vAuthPskIsw, v)
		case "auth_psk__nisw":
			v := value
			vAuthPskNisw = append(vAuthPskNisw, v)
		case "auth_psk__iew":
			v := value
			vAuthPskIew = append(vAuthPskIew, v)
		case "auth_psk__niew":
			v := value
			vAuthPskNiew = append(vAuthPskNiew, v)
		case "auth_psk__empty":
			if vAuthPskEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'auth_psk__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'auth_psk__empty' takes a boolean, got %q.", value))
				return
			}
			vAuthPskEmpty = &v
		case "auth_psk__regex":
			v := value
			vAuthPskRegex = append(vAuthPskRegex, v)
		case "auth_psk__iregex":
			v := value
			vAuthPskIregex = append(vAuthPskIregex, v)
		case "auth_type":
			v := value
			vAuthType = append(vAuthType, v)
		case "auth_type__n":
			v := value
			vAuthTypen = append(vAuthTypen, v)
		case "auth_type__ic":
			v := value
			vAuthTypeIc = append(vAuthTypeIc, v)
		case "auth_type__nic":
			v := value
			vAuthTypeNic = append(vAuthTypeNic, v)
		case "auth_type__ie":
			v := value
			vAuthTypeIe = append(vAuthTypeIe, v)
		case "auth_type__nie":
			v := value
			vAuthTypeNie = append(vAuthTypeNie, v)
		case "auth_type__isw":
			v := value
			vAuthTypeIsw = append(vAuthTypeIsw, v)
		case "auth_type__nisw":
			v := value
			vAuthTypeNisw = append(vAuthTypeNisw, v)
		case "auth_type__iew":
			v := value
			vAuthTypeIew = append(vAuthTypeIew, v)
		case "auth_type__niew":
			v := value
			vAuthTypeNiew = append(vAuthTypeNiew, v)
		case "auth_type__empty":
			if vAuthTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'auth_type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'auth_type__empty' takes a boolean, got %q.", value))
				return
			}
			vAuthTypeEmpty = &v
		case "auth_type__regex":
			v := value
			vAuthTypeRegex = append(vAuthTypeRegex, v)
		case "auth_type__iregex":
			v := value
			vAuthTypeIregex = append(vAuthTypeIregex, v)
		case "created":
			v := value
			vCreated = append(vCreated, v)
		case "created__n":
			v := value
			vCreatedn = append(vCreatedn, v)
		case "created__lt":
			v := value
			vCreatedLt = append(vCreatedLt, v)
		case "created__lte":
			v := value
			vCreatedLte = append(vCreatedLte, v)
		case "created__gt":
			v := value
			vCreatedGt = append(vCreatedGt, v)
		case "created__gte":
			v := value
			vCreatedGte = append(vCreatedGte, v)
		case "created__empty":
			v := value
			vCreatedEmpty = append(vCreatedEmpty, v)
		case "description":
			v := value
			vDescription = append(vDescription, v)
		case "description__n":
			v := value
			vDescriptionn = append(vDescriptionn, v)
		case "description__ic":
			v := value
			vDescriptionIc = append(vDescriptionIc, v)
		case "description__nic":
			v := value
			vDescriptionNic = append(vDescriptionNic, v)
		case "description__ie":
			v := value
			vDescriptionIe = append(vDescriptionIe, v)
		case "description__nie":
			v := value
			vDescriptionNie = append(vDescriptionNie, v)
		case "description__isw":
			v := value
			vDescriptionIsw = append(vDescriptionIsw, v)
		case "description__nisw":
			v := value
			vDescriptionNisw = append(vDescriptionNisw, v)
		case "description__iew":
			v := value
			vDescriptionIew = append(vDescriptionIew, v)
		case "description__niew":
			v := value
			vDescriptionNiew = append(vDescriptionNiew, v)
		case "description__empty":
			if vDescriptionEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'description__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'description__empty' takes a boolean, got %q.", value))
				return
			}
			vDescriptionEmpty = &v
		case "description__regex":
			v := value
			vDescriptionRegex = append(vDescriptionRegex, v)
		case "description__iregex":
			v := value
			vDescriptionIregex = append(vDescriptionIregex, v)
		case "distance":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance' takes a number, got %q.", value))
				return
			}
			vDistance = append(vDistance, v)
		case "distance__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__n' takes a number, got %q.", value))
				return
			}
			vDistancen = append(vDistancen, v)
		case "distance__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__lt' takes a number, got %q.", value))
				return
			}
			vDistanceLt = append(vDistanceLt, v)
		case "distance__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__lte' takes a number, got %q.", value))
				return
			}
			vDistanceLte = append(vDistanceLte, v)
		case "distance__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__gt' takes a number, got %q.", value))
				return
			}
			vDistanceGt = append(vDistanceGt, v)
		case "distance__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__gte' takes a number, got %q.", value))
				return
			}
			vDistanceGte = append(vDistanceGte, v)
		case "distance__empty":
			if vDistanceEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__empty' takes a boolean, got %q.", value))
				return
			}
			vDistanceEmpty = &v
		case "distance_unit":
			if vDistanceUnit != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance_unit' takes a single value.")
				return
			}
			v := value
			vDistanceUnit = &v
		case "distance_unit__n":
			if vDistanceUnitn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance_unit__n' takes a single value.")
				return
			}
			v := value
			vDistanceUnitn = &v
		case "distance_unit__ic":
			v := value
			vDistanceUnitIc = append(vDistanceUnitIc, v)
		case "distance_unit__nic":
			v := value
			vDistanceUnitNic = append(vDistanceUnitNic, v)
		case "distance_unit__ie":
			v := value
			vDistanceUnitIe = append(vDistanceUnitIe, v)
		case "distance_unit__nie":
			v := value
			vDistanceUnitNie = append(vDistanceUnitNie, v)
		case "distance_unit__isw":
			v := value
			vDistanceUnitIsw = append(vDistanceUnitIsw, v)
		case "distance_unit__nisw":
			v := value
			vDistanceUnitNisw = append(vDistanceUnitNisw, v)
		case "distance_unit__iew":
			v := value
			vDistanceUnitIew = append(vDistanceUnitIew, v)
		case "distance_unit__niew":
			v := value
			vDistanceUnitNiew = append(vDistanceUnitNiew, v)
		case "distance_unit__empty":
			if vDistanceUnitEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance_unit__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance_unit__empty' takes a boolean, got %q.", value))
				return
			}
			vDistanceUnitEmpty = &v
		case "distance_unit__regex":
			v := value
			vDistanceUnitRegex = append(vDistanceUnitRegex, v)
		case "distance_unit__iregex":
			v := value
			vDistanceUnitIregex = append(vDistanceUnitIregex, v)
		case "last_updated":
			v := value
			vLastUpdated = append(vLastUpdated, v)
		case "last_updated__n":
			v := value
			vLastUpdatedn = append(vLastUpdatedn, v)
		case "last_updated__lt":
			v := value
			vLastUpdatedLt = append(vLastUpdatedLt, v)
		case "last_updated__lte":
			v := value
			vLastUpdatedLte = append(vLastUpdatedLte, v)
		case "last_updated__gt":
			v := value
			vLastUpdatedGt = append(vLastUpdatedGt, v)
		case "last_updated__gte":
			v := value
			vLastUpdatedGte = append(vLastUpdatedGte, v)
		case "last_updated__empty":
			v := value
			vLastUpdatedEmpty = append(vLastUpdatedEmpty, v)
		case "owner":
			v := value
			vOwner = append(vOwner, v)
		case "owner__n":
			v := value
			vOwnern = append(vOwnern, v)
		case "owner_group":
			v := value
			vOwnerGroup = append(vOwnerGroup, v)
		case "owner_group__n":
			v := value
			vOwnerGroupn = append(vOwnerGroupn, v)
		case "owner_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupID = append(vOwnerGroupID, v)
		case "owner_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupIDn = append(vOwnerGroupIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "tag_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id' takes an integer, got %q.", value))
				return
			}
			vTagID = append(vTagID, v)
		case "tag_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__n' takes an integer, got %q.", value))
				return
			}
			vTagIDn = append(vTagIDn, v)
		case "tag_id__any":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__any' takes an integer, got %q.", value))
				return
			}
			vTagIDAny = append(vTagIDAny, v)
		case "tenant":
			v := value
			vTenant = append(vTenant, v)
		case "tenant__n":
			v := value
			vTenantn = append(vTenantn, v)
		case "tenant_group":
			v := value
			vTenantGroup = append(vTenantGroup, v)
		case "tenant_group__n":
			v := value
			vTenantGroupn = append(vTenantGroupn, v)
		case "tenant_group_id":
			v := value
			vTenantGroupID = append(vTenantGroupID, v)
		case "tenant_group_id__n":
			v := value
			vTenantGroupIDn = append(vTenantGroupIDn, v)
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
	if len(vSsidIc) > 0 {
		params.SetSsidIc(vSsidIc)
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
	if len(vStatus) > 0 {
		params.SetStatus(vStatus)
	}
	if len(vStatusn) > 0 {
		params.SetStatusn(vStatusn)
	}
	if len(vStatusIc) > 0 {
		params.SetStatusIc(vStatusIc)
	}
	if len(vStatusNic) > 0 {
		params.SetStatusNic(vStatusNic)
	}
	if len(vStatusIe) > 0 {
		params.SetStatusIe(vStatusIe)
	}
	if len(vStatusNie) > 0 {
		params.SetStatusNie(vStatusNie)
	}
	if len(vStatusIsw) > 0 {
		params.SetStatusIsw(vStatusIsw)
	}
	if len(vStatusNisw) > 0 {
		params.SetStatusNisw(vStatusNisw)
	}
	if len(vStatusIew) > 0 {
		params.SetStatusIew(vStatusIew)
	}
	if len(vStatusNiew) > 0 {
		params.SetStatusNiew(vStatusNiew)
	}
	if vStatusEmpty != nil {
		params.SetStatusEmpty(vStatusEmpty)
	}
	if len(vStatusRegex) > 0 {
		params.SetStatusRegex(vStatusRegex)
	}
	if len(vStatusIregex) > 0 {
		params.SetStatusIregex(vStatusIregex)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
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
	if len(vInterfaceaID) > 0 {
		params.SetInterfaceaID(vInterfaceaID)
	}
	if len(vInterfaceaIDn) > 0 {
		params.SetInterfaceaIDn(vInterfaceaIDn)
	}
	if len(vInterfacebID) > 0 {
		params.SetInterfacebID(vInterfacebID)
	}
	if len(vInterfacebIDn) > 0 {
		params.SetInterfacebIDn(vInterfacebIDn)
	}
	if len(vAuthCipher) > 0 {
		params.SetAuthCipher(vAuthCipher)
	}
	if len(vAuthCiphern) > 0 {
		params.SetAuthCiphern(vAuthCiphern)
	}
	if len(vAuthCipherIc) > 0 {
		params.SetAuthCipherIc(vAuthCipherIc)
	}
	if len(vAuthCipherNic) > 0 {
		params.SetAuthCipherNic(vAuthCipherNic)
	}
	if len(vAuthCipherIe) > 0 {
		params.SetAuthCipherIe(vAuthCipherIe)
	}
	if len(vAuthCipherNie) > 0 {
		params.SetAuthCipherNie(vAuthCipherNie)
	}
	if len(vAuthCipherIsw) > 0 {
		params.SetAuthCipherIsw(vAuthCipherIsw)
	}
	if len(vAuthCipherNisw) > 0 {
		params.SetAuthCipherNisw(vAuthCipherNisw)
	}
	if len(vAuthCipherIew) > 0 {
		params.SetAuthCipherIew(vAuthCipherIew)
	}
	if len(vAuthCipherNiew) > 0 {
		params.SetAuthCipherNiew(vAuthCipherNiew)
	}
	if vAuthCipherEmpty != nil {
		params.SetAuthCipherEmpty(vAuthCipherEmpty)
	}
	if len(vAuthCipherRegex) > 0 {
		params.SetAuthCipherRegex(vAuthCipherRegex)
	}
	if len(vAuthCipherIregex) > 0 {
		params.SetAuthCipherIregex(vAuthCipherIregex)
	}
	if len(vAuthPsk) > 0 {
		params.SetAuthPsk(vAuthPsk)
	}
	if len(vAuthPskn) > 0 {
		params.SetAuthPskn(vAuthPskn)
	}
	if len(vAuthPskIc) > 0 {
		params.SetAuthPskIc(vAuthPskIc)
	}
	if len(vAuthPskNic) > 0 {
		params.SetAuthPskNic(vAuthPskNic)
	}
	if len(vAuthPskIe) > 0 {
		params.SetAuthPskIe(vAuthPskIe)
	}
	if len(vAuthPskNie) > 0 {
		params.SetAuthPskNie(vAuthPskNie)
	}
	if len(vAuthPskIsw) > 0 {
		params.SetAuthPskIsw(vAuthPskIsw)
	}
	if len(vAuthPskNisw) > 0 {
		params.SetAuthPskNisw(vAuthPskNisw)
	}
	if len(vAuthPskIew) > 0 {
		params.SetAuthPskIew(vAuthPskIew)
	}
	if len(vAuthPskNiew) > 0 {
		params.SetAuthPskNiew(vAuthPskNiew)
	}
	if vAuthPskEmpty != nil {
		params.SetAuthPskEmpty(vAuthPskEmpty)
	}
	if len(vAuthPskRegex) > 0 {
		params.SetAuthPskRegex(vAuthPskRegex)
	}
	if len(vAuthPskIregex) > 0 {
		params.SetAuthPskIregex(vAuthPskIregex)
	}
	if len(vAuthType) > 0 {
		params.SetAuthType(vAuthType)
	}
	if len(vAuthTypen) > 0 {
		params.SetAuthTypen(vAuthTypen)
	}
	if len(vAuthTypeIc) > 0 {
		params.SetAuthTypeIc(vAuthTypeIc)
	}
	if len(vAuthTypeNic) > 0 {
		params.SetAuthTypeNic(vAuthTypeNic)
	}
	if len(vAuthTypeIe) > 0 {
		params.SetAuthTypeIe(vAuthTypeIe)
	}
	if len(vAuthTypeNie) > 0 {
		params.SetAuthTypeNie(vAuthTypeNie)
	}
	if len(vAuthTypeIsw) > 0 {
		params.SetAuthTypeIsw(vAuthTypeIsw)
	}
	if len(vAuthTypeNisw) > 0 {
		params.SetAuthTypeNisw(vAuthTypeNisw)
	}
	if len(vAuthTypeIew) > 0 {
		params.SetAuthTypeIew(vAuthTypeIew)
	}
	if len(vAuthTypeNiew) > 0 {
		params.SetAuthTypeNiew(vAuthTypeNiew)
	}
	if vAuthTypeEmpty != nil {
		params.SetAuthTypeEmpty(vAuthTypeEmpty)
	}
	if len(vAuthTypeRegex) > 0 {
		params.SetAuthTypeRegex(vAuthTypeRegex)
	}
	if len(vAuthTypeIregex) > 0 {
		params.SetAuthTypeIregex(vAuthTypeIregex)
	}
	if len(vCreated) > 0 {
		params.SetCreated(vCreated)
	}
	if len(vCreatedn) > 0 {
		params.SetCreatedn(vCreatedn)
	}
	if len(vCreatedLt) > 0 {
		params.SetCreatedLt(vCreatedLt)
	}
	if len(vCreatedLte) > 0 {
		params.SetCreatedLte(vCreatedLte)
	}
	if len(vCreatedGt) > 0 {
		params.SetCreatedGt(vCreatedGt)
	}
	if len(vCreatedGte) > 0 {
		params.SetCreatedGte(vCreatedGte)
	}
	if len(vCreatedEmpty) > 0 {
		params.SetCreatedEmpty(vCreatedEmpty)
	}
	if len(vDescription) > 0 {
		params.SetDescription(vDescription)
	}
	if len(vDescriptionn) > 0 {
		params.SetDescriptionn(vDescriptionn)
	}
	if len(vDescriptionIc) > 0 {
		params.SetDescriptionIc(vDescriptionIc)
	}
	if len(vDescriptionNic) > 0 {
		params.SetDescriptionNic(vDescriptionNic)
	}
	if len(vDescriptionIe) > 0 {
		params.SetDescriptionIe(vDescriptionIe)
	}
	if len(vDescriptionNie) > 0 {
		params.SetDescriptionNie(vDescriptionNie)
	}
	if len(vDescriptionIsw) > 0 {
		params.SetDescriptionIsw(vDescriptionIsw)
	}
	if len(vDescriptionNisw) > 0 {
		params.SetDescriptionNisw(vDescriptionNisw)
	}
	if len(vDescriptionIew) > 0 {
		params.SetDescriptionIew(vDescriptionIew)
	}
	if len(vDescriptionNiew) > 0 {
		params.SetDescriptionNiew(vDescriptionNiew)
	}
	if vDescriptionEmpty != nil {
		params.SetDescriptionEmpty(vDescriptionEmpty)
	}
	if len(vDescriptionRegex) > 0 {
		params.SetDescriptionRegex(vDescriptionRegex)
	}
	if len(vDescriptionIregex) > 0 {
		params.SetDescriptionIregex(vDescriptionIregex)
	}
	if len(vDistance) > 0 {
		params.SetDistance(vDistance)
	}
	if len(vDistancen) > 0 {
		params.SetDistancen(vDistancen)
	}
	if len(vDistanceLt) > 0 {
		params.SetDistanceLt(vDistanceLt)
	}
	if len(vDistanceLte) > 0 {
		params.SetDistanceLte(vDistanceLte)
	}
	if len(vDistanceGt) > 0 {
		params.SetDistanceGt(vDistanceGt)
	}
	if len(vDistanceGte) > 0 {
		params.SetDistanceGte(vDistanceGte)
	}
	if vDistanceEmpty != nil {
		params.SetDistanceEmpty(vDistanceEmpty)
	}
	if vDistanceUnit != nil {
		params.SetDistanceUnit(vDistanceUnit)
	}
	if vDistanceUnitn != nil {
		params.SetDistanceUnitn(vDistanceUnitn)
	}
	if len(vDistanceUnitIc) > 0 {
		params.SetDistanceUnitIc(vDistanceUnitIc)
	}
	if len(vDistanceUnitNic) > 0 {
		params.SetDistanceUnitNic(vDistanceUnitNic)
	}
	if len(vDistanceUnitIe) > 0 {
		params.SetDistanceUnitIe(vDistanceUnitIe)
	}
	if len(vDistanceUnitNie) > 0 {
		params.SetDistanceUnitNie(vDistanceUnitNie)
	}
	if len(vDistanceUnitIsw) > 0 {
		params.SetDistanceUnitIsw(vDistanceUnitIsw)
	}
	if len(vDistanceUnitNisw) > 0 {
		params.SetDistanceUnitNisw(vDistanceUnitNisw)
	}
	if len(vDistanceUnitIew) > 0 {
		params.SetDistanceUnitIew(vDistanceUnitIew)
	}
	if len(vDistanceUnitNiew) > 0 {
		params.SetDistanceUnitNiew(vDistanceUnitNiew)
	}
	if vDistanceUnitEmpty != nil {
		params.SetDistanceUnitEmpty(vDistanceUnitEmpty)
	}
	if len(vDistanceUnitRegex) > 0 {
		params.SetDistanceUnitRegex(vDistanceUnitRegex)
	}
	if len(vDistanceUnitIregex) > 0 {
		params.SetDistanceUnitIregex(vDistanceUnitIregex)
	}
	if len(vLastUpdated) > 0 {
		params.SetLastUpdated(vLastUpdated)
	}
	if len(vLastUpdatedn) > 0 {
		params.SetLastUpdatedn(vLastUpdatedn)
	}
	if len(vLastUpdatedLt) > 0 {
		params.SetLastUpdatedLt(vLastUpdatedLt)
	}
	if len(vLastUpdatedLte) > 0 {
		params.SetLastUpdatedLte(vLastUpdatedLte)
	}
	if len(vLastUpdatedGt) > 0 {
		params.SetLastUpdatedGt(vLastUpdatedGt)
	}
	if len(vLastUpdatedGte) > 0 {
		params.SetLastUpdatedGte(vLastUpdatedGte)
	}
	if len(vLastUpdatedEmpty) > 0 {
		params.SetLastUpdatedEmpty(vLastUpdatedEmpty)
	}
	if len(vOwner) > 0 {
		params.SetOwner(vOwner)
	}
	if len(vOwnern) > 0 {
		params.SetOwnern(vOwnern)
	}
	if len(vOwnerGroup) > 0 {
		params.SetOwnerGroup(vOwnerGroup)
	}
	if len(vOwnerGroupn) > 0 {
		params.SetOwnerGroupn(vOwnerGroupn)
	}
	if len(vOwnerGroupID) > 0 {
		params.SetOwnerGroupID(vOwnerGroupID)
	}
	if len(vOwnerGroupIDn) > 0 {
		params.SetOwnerGroupIDn(vOwnerGroupIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vTagID) > 0 {
		params.SetTagID(vTagID)
	}
	if len(vTagIDn) > 0 {
		params.SetTagIDn(vTagIDn)
	}
	if len(vTagIDAny) > 0 {
		params.SetTagIDAny(vTagIDAny)
	}
	if len(vTenant) > 0 {
		params.SetTenant(vTenant)
	}
	if len(vTenantn) > 0 {
		params.SetTenantn(vTenantn)
	}
	if len(vTenantGroup) > 0 {
		params.SetTenantGroup(vTenantGroup)
	}
	if len(vTenantGroupn) > 0 {
		params.SetTenantGroupn(vTenantGroupn)
	}
	if len(vTenantGroupID) > 0 {
		params.SetTenantGroupID(vTenantGroupID)
	}
	if len(vTenantGroupIDn) > 0 {
		params.SetTenantGroupIDn(vTenantGroupIDn)
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
