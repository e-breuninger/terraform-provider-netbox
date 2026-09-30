// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*ikeProposalsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ikeProposalsDataSource)(nil)
)

// NewIkeProposalsDataSource returns a new ike_proposals data source, which
// lists ike_proposal objects matching its filters.
func NewIkeProposalsDataSource() datasource.DataSource {
	return &ikeProposalsDataSource{}
}

type ikeProposalsDataSource struct {
	client *netboxapi.Client
}

// ikeProposalsDataSourceModel is the data source model: the filters, the limit and the matching
// ike_proposals.
type ikeProposalsDataSourceModel struct {
	Filters   types.Set    `tfsdk:"filters"`
	NameRegex types.String `tfsdk:"name_regex"`
	Limit     types.Int64  `tfsdk:"limit"`
	Items     types.List   `tfsdk:"ike_proposals"`
}

// ikeProposalResourceAttrTypes is the attribute type map of ikeProposalResourceModel.
func ikeProposalResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                       types.Int64Type,
		"name":                     types.StringType,
		"authentication_method":    types.StringType,
		"encryption_algorithm":     types.StringType,
		"authentication_algorithm": types.StringType,
		"group":                    types.Int64Type,
		"sa_lifetime":              types.Int64Type,
		"description":              types.StringType,
		"comments":                 types.StringType,
		"owner_id":                 types.Int64Type,
		"created":                  types.StringType,
		"last_updated":             types.StringType,
		"url":                      types.StringType,
		"tags":                     types.SetType{ElemType: types.StringType},
		"tags_all":                 types.SetType{ElemType: types.StringType},
		"custom_fields":            types.MapType{ElemType: types.StringType},
	}
}

func (d *ikeProposalsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ike_proposals"
}

func (d *ikeProposalsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:Lists ike_proposal objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: authentication_method, authentication_method__empty, authentication_method__ic, authentication_method__ie, authentication_method__iew, authentication_method__iregex, authentication_method__isw, authentication_method__n, authentication_method__nic, authentication_method__nie, authentication_method__niew, authentication_method__nisw, authentication_method__regex, encryption_algorithm, encryption_algorithm__empty, encryption_algorithm__ic, encryption_algorithm__ie, encryption_algorithm__iew, encryption_algorithm__iregex, encryption_algorithm__isw, encryption_algorithm__n, encryption_algorithm__nic, encryption_algorithm__nie, encryption_algorithm__niew, encryption_algorithm__nisw, encryption_algorithm__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner_id, owner_id__n, tag, tag__any, tag__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"name_regex": schema.StringAttribute{
				Optional:    true,
				Description: "Go regular expression the name must match. Applied after the API query, which then fetches every object matching the filters; limit applies to the matches.",
				Validators: []validator.String{
					conv.ValidRegexp(),
				},
			},
			"limit": schema.Int64Attribute{
				Optional:    true,
				Description: "The maximum number of objects to return from the API lookup. Defaults to 1000.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"ike_proposals": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching ike_proposals.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"authentication_method": schema.StringAttribute{
							Computed:    true,
							Description: "One of: preshared-keys, certificates, rsa-signatures, dsa-signatures.",
						},
						"encryption_algorithm": schema.StringAttribute{
							Computed:    true,
							Description: "One of: aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, 3des-cbc, des-cbc.",
						},
						"authentication_algorithm": schema.StringAttribute{
							Computed:    true,
							Description: "Not needed for GCM encryption algorithms. NetBox does not allow clearing it once set; leaving it unset keeps the current value. One of: hmac-sha1, hmac-sha256, hmac-sha384, hmac-sha512, hmac-md5.",
						},
						"group": schema.Int64Attribute{
							Computed:    true,
							Description: "Diffie-Hellman group number. One of: 1, 2, 5, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34.",
						},
						"sa_lifetime": schema.Int64Attribute{
							Computed:    true,
							Description: "Security association lifetime in seconds.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *ikeProposalsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ikeProposalsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ikeProposalsDataSourceModel

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

	// name_regex is applied here, not by the API: the fragment fetches every match (fetchAll),
	// the names are matched below, and only then does limit apply.
	var nameRegex *regexp.Regexp
	fetchAll := false
	if !state.NameRegex.IsNull() && !state.NameRegex.IsUnknown() {
		re, err := regexp.Compile(state.NameRegex.ValueString())
		if err != nil {
			resp.Diagnostics.AddAttributeError(path.Root("name_regex"), "Invalid regular expression", err.Error())
			return
		}
		nameRegex, fetchAll = re, true
	}
	_ = fetchAll
	var responseDTOs []*netboxapi.IkeProposalResponseDTO

	params := vpn.NewVpnIkeProposalsListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vName []string
	var vNamen []string
	var vNameIc []string
	var vNameNic []string
	var vNameIe []string
	var vNameNie []string
	var vNameIsw []string
	var vNameNisw []string
	var vNameIew []string
	var vNameNiew []string
	var vNameEmpty *bool
	var vNameRegex []string
	var vNameIregex []string
	var vAuthenticationMethod []string
	var vAuthenticationMethodn []string
	var vAuthenticationMethodIc []string
	var vAuthenticationMethodNic []string
	var vAuthenticationMethodIe []string
	var vAuthenticationMethodNie []string
	var vAuthenticationMethodIsw []string
	var vAuthenticationMethodNisw []string
	var vAuthenticationMethodIew []string
	var vAuthenticationMethodNiew []string
	var vAuthenticationMethodEmpty *bool
	var vAuthenticationMethodRegex []string
	var vAuthenticationMethodIregex []string
	var vEncryptionAlgorithm []string
	var vEncryptionAlgorithmn []string
	var vEncryptionAlgorithmIc []string
	var vEncryptionAlgorithmNic []string
	var vEncryptionAlgorithmIe []string
	var vEncryptionAlgorithmNie []string
	var vEncryptionAlgorithmIsw []string
	var vEncryptionAlgorithmNisw []string
	var vEncryptionAlgorithmIew []string
	var vEncryptionAlgorithmNiew []string
	var vEncryptionAlgorithmEmpty *bool
	var vEncryptionAlgorithmRegex []string
	var vEncryptionAlgorithmIregex []string
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
		case "name":
			v := value
			vName = append(vName, v)
		case "name__n":
			v := value
			vNamen = append(vNamen, v)
		case "name__ic":
			v := value
			vNameIc = append(vNameIc, v)
		case "name__nic":
			v := value
			vNameNic = append(vNameNic, v)
		case "name__ie":
			v := value
			vNameIe = append(vNameIe, v)
		case "name__nie":
			v := value
			vNameNie = append(vNameNie, v)
		case "name__isw":
			v := value
			vNameIsw = append(vNameIsw, v)
		case "name__nisw":
			v := value
			vNameNisw = append(vNameNisw, v)
		case "name__iew":
			v := value
			vNameIew = append(vNameIew, v)
		case "name__niew":
			v := value
			vNameNiew = append(vNameNiew, v)
		case "name__empty":
			if vNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'name__empty' takes a boolean, got %q.", value))
				return
			}
			vNameEmpty = &v
		case "name__regex":
			v := value
			vNameRegex = append(vNameRegex, v)
		case "name__iregex":
			v := value
			vNameIregex = append(vNameIregex, v)
		case "authentication_method":
			v := value
			vAuthenticationMethod = append(vAuthenticationMethod, v)
		case "authentication_method__n":
			v := value
			vAuthenticationMethodn = append(vAuthenticationMethodn, v)
		case "authentication_method__ic":
			v := value
			vAuthenticationMethodIc = append(vAuthenticationMethodIc, v)
		case "authentication_method__nic":
			v := value
			vAuthenticationMethodNic = append(vAuthenticationMethodNic, v)
		case "authentication_method__ie":
			v := value
			vAuthenticationMethodIe = append(vAuthenticationMethodIe, v)
		case "authentication_method__nie":
			v := value
			vAuthenticationMethodNie = append(vAuthenticationMethodNie, v)
		case "authentication_method__isw":
			v := value
			vAuthenticationMethodIsw = append(vAuthenticationMethodIsw, v)
		case "authentication_method__nisw":
			v := value
			vAuthenticationMethodNisw = append(vAuthenticationMethodNisw, v)
		case "authentication_method__iew":
			v := value
			vAuthenticationMethodIew = append(vAuthenticationMethodIew, v)
		case "authentication_method__niew":
			v := value
			vAuthenticationMethodNiew = append(vAuthenticationMethodNiew, v)
		case "authentication_method__empty":
			if vAuthenticationMethodEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'authentication_method__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'authentication_method__empty' takes a boolean, got %q.", value))
				return
			}
			vAuthenticationMethodEmpty = &v
		case "authentication_method__regex":
			v := value
			vAuthenticationMethodRegex = append(vAuthenticationMethodRegex, v)
		case "authentication_method__iregex":
			v := value
			vAuthenticationMethodIregex = append(vAuthenticationMethodIregex, v)
		case "encryption_algorithm":
			v := value
			vEncryptionAlgorithm = append(vEncryptionAlgorithm, v)
		case "encryption_algorithm__n":
			v := value
			vEncryptionAlgorithmn = append(vEncryptionAlgorithmn, v)
		case "encryption_algorithm__ic":
			v := value
			vEncryptionAlgorithmIc = append(vEncryptionAlgorithmIc, v)
		case "encryption_algorithm__nic":
			v := value
			vEncryptionAlgorithmNic = append(vEncryptionAlgorithmNic, v)
		case "encryption_algorithm__ie":
			v := value
			vEncryptionAlgorithmIe = append(vEncryptionAlgorithmIe, v)
		case "encryption_algorithm__nie":
			v := value
			vEncryptionAlgorithmNie = append(vEncryptionAlgorithmNie, v)
		case "encryption_algorithm__isw":
			v := value
			vEncryptionAlgorithmIsw = append(vEncryptionAlgorithmIsw, v)
		case "encryption_algorithm__nisw":
			v := value
			vEncryptionAlgorithmNisw = append(vEncryptionAlgorithmNisw, v)
		case "encryption_algorithm__iew":
			v := value
			vEncryptionAlgorithmIew = append(vEncryptionAlgorithmIew, v)
		case "encryption_algorithm__niew":
			v := value
			vEncryptionAlgorithmNiew = append(vEncryptionAlgorithmNiew, v)
		case "encryption_algorithm__empty":
			if vEncryptionAlgorithmEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'encryption_algorithm__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'encryption_algorithm__empty' takes a boolean, got %q.", value))
				return
			}
			vEncryptionAlgorithmEmpty = &v
		case "encryption_algorithm__regex":
			v := value
			vEncryptionAlgorithmRegex = append(vEncryptionAlgorithmRegex, v)
		case "encryption_algorithm__iregex":
			v := value
			vEncryptionAlgorithmIregex = append(vEncryptionAlgorithmIregex, v)
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
	if len(vName) > 0 {
		params.SetName(vName)
	}
	if len(vNamen) > 0 {
		params.SetNamen(vNamen)
	}
	if len(vNameIc) > 0 {
		params.SetNameIc(vNameIc)
	}
	if len(vNameNic) > 0 {
		params.SetNameNic(vNameNic)
	}
	if len(vNameIe) > 0 {
		params.SetNameIe(vNameIe)
	}
	if len(vNameNie) > 0 {
		params.SetNameNie(vNameNie)
	}
	if len(vNameIsw) > 0 {
		params.SetNameIsw(vNameIsw)
	}
	if len(vNameNisw) > 0 {
		params.SetNameNisw(vNameNisw)
	}
	if len(vNameIew) > 0 {
		params.SetNameIew(vNameIew)
	}
	if len(vNameNiew) > 0 {
		params.SetNameNiew(vNameNiew)
	}
	if vNameEmpty != nil {
		params.SetNameEmpty(vNameEmpty)
	}
	if len(vNameRegex) > 0 {
		params.SetNameRegex(vNameRegex)
	}
	if len(vNameIregex) > 0 {
		params.SetNameIregex(vNameIregex)
	}
	if len(vAuthenticationMethod) > 0 {
		params.SetAuthenticationMethod(vAuthenticationMethod)
	}
	if len(vAuthenticationMethodn) > 0 {
		params.SetAuthenticationMethodn(vAuthenticationMethodn)
	}
	if len(vAuthenticationMethodIc) > 0 {
		params.SetAuthenticationMethodIc(vAuthenticationMethodIc)
	}
	if len(vAuthenticationMethodNic) > 0 {
		params.SetAuthenticationMethodNic(vAuthenticationMethodNic)
	}
	if len(vAuthenticationMethodIe) > 0 {
		params.SetAuthenticationMethodIe(vAuthenticationMethodIe)
	}
	if len(vAuthenticationMethodNie) > 0 {
		params.SetAuthenticationMethodNie(vAuthenticationMethodNie)
	}
	if len(vAuthenticationMethodIsw) > 0 {
		params.SetAuthenticationMethodIsw(vAuthenticationMethodIsw)
	}
	if len(vAuthenticationMethodNisw) > 0 {
		params.SetAuthenticationMethodNisw(vAuthenticationMethodNisw)
	}
	if len(vAuthenticationMethodIew) > 0 {
		params.SetAuthenticationMethodIew(vAuthenticationMethodIew)
	}
	if len(vAuthenticationMethodNiew) > 0 {
		params.SetAuthenticationMethodNiew(vAuthenticationMethodNiew)
	}
	if vAuthenticationMethodEmpty != nil {
		params.SetAuthenticationMethodEmpty(vAuthenticationMethodEmpty)
	}
	if len(vAuthenticationMethodRegex) > 0 {
		params.SetAuthenticationMethodRegex(vAuthenticationMethodRegex)
	}
	if len(vAuthenticationMethodIregex) > 0 {
		params.SetAuthenticationMethodIregex(vAuthenticationMethodIregex)
	}
	if len(vEncryptionAlgorithm) > 0 {
		params.SetEncryptionAlgorithm(vEncryptionAlgorithm)
	}
	if len(vEncryptionAlgorithmn) > 0 {
		params.SetEncryptionAlgorithmn(vEncryptionAlgorithmn)
	}
	if len(vEncryptionAlgorithmIc) > 0 {
		params.SetEncryptionAlgorithmIc(vEncryptionAlgorithmIc)
	}
	if len(vEncryptionAlgorithmNic) > 0 {
		params.SetEncryptionAlgorithmNic(vEncryptionAlgorithmNic)
	}
	if len(vEncryptionAlgorithmIe) > 0 {
		params.SetEncryptionAlgorithmIe(vEncryptionAlgorithmIe)
	}
	if len(vEncryptionAlgorithmNie) > 0 {
		params.SetEncryptionAlgorithmNie(vEncryptionAlgorithmNie)
	}
	if len(vEncryptionAlgorithmIsw) > 0 {
		params.SetEncryptionAlgorithmIsw(vEncryptionAlgorithmIsw)
	}
	if len(vEncryptionAlgorithmNisw) > 0 {
		params.SetEncryptionAlgorithmNisw(vEncryptionAlgorithmNisw)
	}
	if len(vEncryptionAlgorithmIew) > 0 {
		params.SetEncryptionAlgorithmIew(vEncryptionAlgorithmIew)
	}
	if len(vEncryptionAlgorithmNiew) > 0 {
		params.SetEncryptionAlgorithmNiew(vEncryptionAlgorithmNiew)
	}
	if vEncryptionAlgorithmEmpty != nil {
		params.SetEncryptionAlgorithmEmpty(vEncryptionAlgorithmEmpty)
	}
	if len(vEncryptionAlgorithmRegex) > 0 {
		params.SetEncryptionAlgorithmRegex(vEncryptionAlgorithmRegex)
	}
	if len(vEncryptionAlgorithmIregex) > 0 {
		params.SetEncryptionAlgorithmIregex(vEncryptionAlgorithmIregex)
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
		res, err := d.client.Vpn.VpnIkeProposalsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_ike_proposals", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.IkeProposalResponseDTOFromGoNetbox(goNetboxModel))
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
	items := make([]ikeProposalResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model ikeProposalResourceModel
		resp.Diagnostics.Append(flattenIkeProposal(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if nameRegex != nil {
		kept := items[:0]
		for _, item := range items {
			if nameRegex.MatchString(item.Name.ValueString()) {
				kept = append(kept, item)
			}
		}
		items = kept
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ikeProposalResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
