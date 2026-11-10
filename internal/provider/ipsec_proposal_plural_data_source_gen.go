// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*ipsecProposalsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ipsecProposalsDataSource)(nil)
)

// NewIpsecProposalsDataSource returns a new ipsec_proposals data source, which
// lists ipsec_proposal objects matching its filters.
func NewIpsecProposalsDataSource() datasource.DataSource {
	return &ipsecProposalsDataSource{}
}

type ipsecProposalsDataSource struct {
	client *netboxapi.Client
}

// ipsecProposalsDataSourceModel is the data source model: the filters, the limit and the matching
// ipsec_proposals.
type ipsecProposalsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"ipsec_proposals"`
}

// ipsecProposalResourceAttrTypes is the attribute type map of ipsecProposalResourceModel.
func ipsecProposalResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                       types.Int64Type,
		"name":                     types.StringType,
		"encryption_algorithm":     types.StringType,
		"authentication_algorithm": types.StringType,
		"sa_lifetime_seconds":      types.Int64Type,
		"sa_lifetime_data":         types.Int64Type,
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

func (d *ipsecProposalsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipsec_proposals"
}

func (d *ipsecProposalsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:Lists ipsec_proposal objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: authentication_algorithm, authentication_algorithm__empty, authentication_algorithm__ic, authentication_algorithm__ie, authentication_algorithm__iew, authentication_algorithm__iregex, authentication_algorithm__isw, authentication_algorithm__n, authentication_algorithm__nic, authentication_algorithm__nie, authentication_algorithm__niew, authentication_algorithm__nisw, authentication_algorithm__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, encryption_algorithm, encryption_algorithm__empty, encryption_algorithm__ic, encryption_algorithm__ie, encryption_algorithm__iew, encryption_algorithm__iregex, encryption_algorithm__isw, encryption_algorithm__n, encryption_algorithm__nic, encryption_algorithm__nie, encryption_algorithm__niew, encryption_algorithm__nisw, encryption_algorithm__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, ipsec_policy, ipsec_policy__n, ipsec_policy_id, ipsec_policy_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, sa_lifetime_data, sa_lifetime_data__empty, sa_lifetime_data__gt, sa_lifetime_data__gte, sa_lifetime_data__lt, sa_lifetime_data__lte, sa_lifetime_data__n, sa_lifetime_seconds, sa_lifetime_seconds__empty, sa_lifetime_seconds__gt, sa_lifetime_seconds__gte, sa_lifetime_seconds__lt, sa_lifetime_seconds__lte, sa_lifetime_seconds__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"ipsec_proposals": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching ipsec_proposals.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"encryption_algorithm": schema.StringAttribute{
							Computed:    true,
							Description: "Encryption algorithm. NetBox requires an encryption algorithm, an authentication algorithm, or both, and does not allow clearing either once set; leaving it unset keeps the current value. One of: aes-128-cbc, aes-128-gcm, aes-192-cbc, aes-192-gcm, aes-256-cbc, aes-256-gcm, 3des-cbc, des-cbc.",
						},
						"authentication_algorithm": schema.StringAttribute{
							Computed:    true,
							Description: "Authentication algorithm; not needed for GCM encryption algorithms. NetBox does not allow clearing it once set; leaving it unset keeps the current value. One of: hmac-sha1, hmac-sha256, hmac-sha384, hmac-sha512, hmac-md5.",
						},
						"sa_lifetime_seconds": schema.Int64Attribute{
							Computed:    true,
							Description: "Security association lifetime in seconds. Must be a whole number.",
						},
						"sa_lifetime_data": schema.Int64Attribute{
							Computed:    true,
							Description: "Security association lifetime in kilobytes. Must be a whole number.",
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

func (d *ipsecProposalsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ipsecProposalsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ipsecProposalsDataSourceModel

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
	var responseDTOs []*netboxapi.IpsecProposalResponseDTO

	params := vpn.NewVpnIpsecProposalsListParams()
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
	var vAuthenticationAlgorithm []string
	var vAuthenticationAlgorithmn []string
	var vAuthenticationAlgorithmIc []string
	var vAuthenticationAlgorithmNic []string
	var vAuthenticationAlgorithmIe []string
	var vAuthenticationAlgorithmNie []string
	var vAuthenticationAlgorithmIsw []string
	var vAuthenticationAlgorithmNisw []string
	var vAuthenticationAlgorithmIew []string
	var vAuthenticationAlgorithmNiew []string
	var vAuthenticationAlgorithmEmpty *bool
	var vAuthenticationAlgorithmRegex []string
	var vAuthenticationAlgorithmIregex []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
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
	var vIpsecPolicy []string
	var vIpsecPolicyn []string
	var vIpsecPolicyID []int64
	var vIpsecPolicyIDn []int64
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
	var vSaLifetimeData []int64
	var vSaLifetimeDatan []int64
	var vSaLifetimeDataLt []int64
	var vSaLifetimeDataLte []int64
	var vSaLifetimeDataGt []int64
	var vSaLifetimeDataGte []int64
	var vSaLifetimeDataEmpty *bool
	var vSaLifetimeSeconds []int64
	var vSaLifetimeSecondsn []int64
	var vSaLifetimeSecondsLt []int64
	var vSaLifetimeSecondsLte []int64
	var vSaLifetimeSecondsGt []int64
	var vSaLifetimeSecondsGte []int64
	var vSaLifetimeSecondsEmpty *bool
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
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
		case "authentication_algorithm":
			v := value
			vAuthenticationAlgorithm = append(vAuthenticationAlgorithm, v)
		case "authentication_algorithm__n":
			v := value
			vAuthenticationAlgorithmn = append(vAuthenticationAlgorithmn, v)
		case "authentication_algorithm__ic":
			v := value
			vAuthenticationAlgorithmIc = append(vAuthenticationAlgorithmIc, v)
		case "authentication_algorithm__nic":
			v := value
			vAuthenticationAlgorithmNic = append(vAuthenticationAlgorithmNic, v)
		case "authentication_algorithm__ie":
			v := value
			vAuthenticationAlgorithmIe = append(vAuthenticationAlgorithmIe, v)
		case "authentication_algorithm__nie":
			v := value
			vAuthenticationAlgorithmNie = append(vAuthenticationAlgorithmNie, v)
		case "authentication_algorithm__isw":
			v := value
			vAuthenticationAlgorithmIsw = append(vAuthenticationAlgorithmIsw, v)
		case "authentication_algorithm__nisw":
			v := value
			vAuthenticationAlgorithmNisw = append(vAuthenticationAlgorithmNisw, v)
		case "authentication_algorithm__iew":
			v := value
			vAuthenticationAlgorithmIew = append(vAuthenticationAlgorithmIew, v)
		case "authentication_algorithm__niew":
			v := value
			vAuthenticationAlgorithmNiew = append(vAuthenticationAlgorithmNiew, v)
		case "authentication_algorithm__empty":
			if vAuthenticationAlgorithmEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'authentication_algorithm__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'authentication_algorithm__empty' takes a boolean, got %q.", value))
				return
			}
			vAuthenticationAlgorithmEmpty = &v
		case "authentication_algorithm__regex":
			v := value
			vAuthenticationAlgorithmRegex = append(vAuthenticationAlgorithmRegex, v)
		case "authentication_algorithm__iregex":
			v := value
			vAuthenticationAlgorithmIregex = append(vAuthenticationAlgorithmIregex, v)
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
		case "ipsec_policy":
			v := value
			vIpsecPolicy = append(vIpsecPolicy, v)
		case "ipsec_policy__n":
			v := value
			vIpsecPolicyn = append(vIpsecPolicyn, v)
		case "ipsec_policy_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ipsec_policy_id' takes an integer, got %q.", value))
				return
			}
			vIpsecPolicyID = append(vIpsecPolicyID, v)
		case "ipsec_policy_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ipsec_policy_id__n' takes an integer, got %q.", value))
				return
			}
			vIpsecPolicyIDn = append(vIpsecPolicyIDn, v)
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
		case "sa_lifetime_data":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_data' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeData = append(vSaLifetimeData, v)
		case "sa_lifetime_data__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_data__n' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeDatan = append(vSaLifetimeDatan, v)
		case "sa_lifetime_data__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_data__lt' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeDataLt = append(vSaLifetimeDataLt, v)
		case "sa_lifetime_data__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_data__lte' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeDataLte = append(vSaLifetimeDataLte, v)
		case "sa_lifetime_data__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_data__gt' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeDataGt = append(vSaLifetimeDataGt, v)
		case "sa_lifetime_data__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_data__gte' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeDataGte = append(vSaLifetimeDataGte, v)
		case "sa_lifetime_data__empty":
			if vSaLifetimeDataEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'sa_lifetime_data__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_data__empty' takes a boolean, got %q.", value))
				return
			}
			vSaLifetimeDataEmpty = &v
		case "sa_lifetime_seconds":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_seconds' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeSeconds = append(vSaLifetimeSeconds, v)
		case "sa_lifetime_seconds__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_seconds__n' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeSecondsn = append(vSaLifetimeSecondsn, v)
		case "sa_lifetime_seconds__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_seconds__lt' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeSecondsLt = append(vSaLifetimeSecondsLt, v)
		case "sa_lifetime_seconds__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_seconds__lte' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeSecondsLte = append(vSaLifetimeSecondsLte, v)
		case "sa_lifetime_seconds__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_seconds__gt' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeSecondsGt = append(vSaLifetimeSecondsGt, v)
		case "sa_lifetime_seconds__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_seconds__gte' takes an integer, got %q.", value))
				return
			}
			vSaLifetimeSecondsGte = append(vSaLifetimeSecondsGte, v)
		case "sa_lifetime_seconds__empty":
			if vSaLifetimeSecondsEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'sa_lifetime_seconds__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'sa_lifetime_seconds__empty' takes a boolean, got %q.", value))
				return
			}
			vSaLifetimeSecondsEmpty = &v
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
	if len(vAuthenticationAlgorithm) > 0 {
		params.SetAuthenticationAlgorithm(vAuthenticationAlgorithm)
	}
	if len(vAuthenticationAlgorithmn) > 0 {
		params.SetAuthenticationAlgorithmn(vAuthenticationAlgorithmn)
	}
	if len(vAuthenticationAlgorithmIc) > 0 {
		params.SetAuthenticationAlgorithmIc(vAuthenticationAlgorithmIc)
	}
	if len(vAuthenticationAlgorithmNic) > 0 {
		params.SetAuthenticationAlgorithmNic(vAuthenticationAlgorithmNic)
	}
	if len(vAuthenticationAlgorithmIe) > 0 {
		params.SetAuthenticationAlgorithmIe(vAuthenticationAlgorithmIe)
	}
	if len(vAuthenticationAlgorithmNie) > 0 {
		params.SetAuthenticationAlgorithmNie(vAuthenticationAlgorithmNie)
	}
	if len(vAuthenticationAlgorithmIsw) > 0 {
		params.SetAuthenticationAlgorithmIsw(vAuthenticationAlgorithmIsw)
	}
	if len(vAuthenticationAlgorithmNisw) > 0 {
		params.SetAuthenticationAlgorithmNisw(vAuthenticationAlgorithmNisw)
	}
	if len(vAuthenticationAlgorithmIew) > 0 {
		params.SetAuthenticationAlgorithmIew(vAuthenticationAlgorithmIew)
	}
	if len(vAuthenticationAlgorithmNiew) > 0 {
		params.SetAuthenticationAlgorithmNiew(vAuthenticationAlgorithmNiew)
	}
	if vAuthenticationAlgorithmEmpty != nil {
		params.SetAuthenticationAlgorithmEmpty(vAuthenticationAlgorithmEmpty)
	}
	if len(vAuthenticationAlgorithmRegex) > 0 {
		params.SetAuthenticationAlgorithmRegex(vAuthenticationAlgorithmRegex)
	}
	if len(vAuthenticationAlgorithmIregex) > 0 {
		params.SetAuthenticationAlgorithmIregex(vAuthenticationAlgorithmIregex)
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
	if len(vIpsecPolicy) > 0 {
		params.SetIpsecPolicy(vIpsecPolicy)
	}
	if len(vIpsecPolicyn) > 0 {
		params.SetIpsecPolicyn(vIpsecPolicyn)
	}
	if len(vIpsecPolicyID) > 0 {
		params.SetIpsecPolicyID(vIpsecPolicyID)
	}
	if len(vIpsecPolicyIDn) > 0 {
		params.SetIpsecPolicyIDn(vIpsecPolicyIDn)
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
	if len(vSaLifetimeData) > 0 {
		params.SetSaLifetimeData(vSaLifetimeData)
	}
	if len(vSaLifetimeDatan) > 0 {
		params.SetSaLifetimeDatan(vSaLifetimeDatan)
	}
	if len(vSaLifetimeDataLt) > 0 {
		params.SetSaLifetimeDataLt(vSaLifetimeDataLt)
	}
	if len(vSaLifetimeDataLte) > 0 {
		params.SetSaLifetimeDataLte(vSaLifetimeDataLte)
	}
	if len(vSaLifetimeDataGt) > 0 {
		params.SetSaLifetimeDataGt(vSaLifetimeDataGt)
	}
	if len(vSaLifetimeDataGte) > 0 {
		params.SetSaLifetimeDataGte(vSaLifetimeDataGte)
	}
	if vSaLifetimeDataEmpty != nil {
		params.SetSaLifetimeDataEmpty(vSaLifetimeDataEmpty)
	}
	if len(vSaLifetimeSeconds) > 0 {
		params.SetSaLifetimeSeconds(vSaLifetimeSeconds)
	}
	if len(vSaLifetimeSecondsn) > 0 {
		params.SetSaLifetimeSecondsn(vSaLifetimeSecondsn)
	}
	if len(vSaLifetimeSecondsLt) > 0 {
		params.SetSaLifetimeSecondsLt(vSaLifetimeSecondsLt)
	}
	if len(vSaLifetimeSecondsLte) > 0 {
		params.SetSaLifetimeSecondsLte(vSaLifetimeSecondsLte)
	}
	if len(vSaLifetimeSecondsGt) > 0 {
		params.SetSaLifetimeSecondsGt(vSaLifetimeSecondsGt)
	}
	if len(vSaLifetimeSecondsGte) > 0 {
		params.SetSaLifetimeSecondsGte(vSaLifetimeSecondsGte)
	}
	if vSaLifetimeSecondsEmpty != nil {
		params.SetSaLifetimeSecondsEmpty(vSaLifetimeSecondsEmpty)
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
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Vpn.VpnIpsecProposalsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_ipsec_proposals", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.IpsecProposalResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]ipsecProposalResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model ipsecProposalResourceModel
		resp.Diagnostics.Append(flattenIpsecProposal(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ipsecProposalResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
