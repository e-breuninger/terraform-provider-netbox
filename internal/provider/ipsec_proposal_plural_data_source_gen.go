// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                               netboxapi.FilterInt64s(params.SetID),
		"id__n":                            netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                           netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                          netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                           netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                          netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                        netboxapi.FilterBool(params.SetIDEmpty),
		"name":                             netboxapi.FilterStrings(params.SetName),
		"name__n":                          netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                         netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                        netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                         netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                        netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                        netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                       netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                        netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                       netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":                      netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":                      netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":                     netboxapi.FilterStrings(params.SetNameIregex),
		"encryption_algorithm":             netboxapi.FilterStrings(params.SetEncryptionAlgorithm),
		"encryption_algorithm__n":          netboxapi.FilterStrings(params.SetEncryptionAlgorithmn),
		"encryption_algorithm__ic":         netboxapi.FilterStrings(params.SetEncryptionAlgorithmIc),
		"encryption_algorithm__nic":        netboxapi.FilterStrings(params.SetEncryptionAlgorithmNic),
		"encryption_algorithm__ie":         netboxapi.FilterStrings(params.SetEncryptionAlgorithmIe),
		"encryption_algorithm__nie":        netboxapi.FilterStrings(params.SetEncryptionAlgorithmNie),
		"encryption_algorithm__isw":        netboxapi.FilterStrings(params.SetEncryptionAlgorithmIsw),
		"encryption_algorithm__nisw":       netboxapi.FilterStrings(params.SetEncryptionAlgorithmNisw),
		"encryption_algorithm__iew":        netboxapi.FilterStrings(params.SetEncryptionAlgorithmIew),
		"encryption_algorithm__niew":       netboxapi.FilterStrings(params.SetEncryptionAlgorithmNiew),
		"encryption_algorithm__empty":      netboxapi.FilterBool(params.SetEncryptionAlgorithmEmpty),
		"encryption_algorithm__regex":      netboxapi.FilterStrings(params.SetEncryptionAlgorithmRegex),
		"encryption_algorithm__iregex":     netboxapi.FilterStrings(params.SetEncryptionAlgorithmIregex),
		"authentication_algorithm":         netboxapi.FilterStrings(params.SetAuthenticationAlgorithm),
		"authentication_algorithm__n":      netboxapi.FilterStrings(params.SetAuthenticationAlgorithmn),
		"authentication_algorithm__ic":     netboxapi.FilterStrings(params.SetAuthenticationAlgorithmIc),
		"authentication_algorithm__nic":    netboxapi.FilterStrings(params.SetAuthenticationAlgorithmNic),
		"authentication_algorithm__ie":     netboxapi.FilterStrings(params.SetAuthenticationAlgorithmIe),
		"authentication_algorithm__nie":    netboxapi.FilterStrings(params.SetAuthenticationAlgorithmNie),
		"authentication_algorithm__isw":    netboxapi.FilterStrings(params.SetAuthenticationAlgorithmIsw),
		"authentication_algorithm__nisw":   netboxapi.FilterStrings(params.SetAuthenticationAlgorithmNisw),
		"authentication_algorithm__iew":    netboxapi.FilterStrings(params.SetAuthenticationAlgorithmIew),
		"authentication_algorithm__niew":   netboxapi.FilterStrings(params.SetAuthenticationAlgorithmNiew),
		"authentication_algorithm__empty":  netboxapi.FilterBool(params.SetAuthenticationAlgorithmEmpty),
		"authentication_algorithm__regex":  netboxapi.FilterStrings(params.SetAuthenticationAlgorithmRegex),
		"authentication_algorithm__iregex": netboxapi.FilterStrings(params.SetAuthenticationAlgorithmIregex),
		"owner_id":                         netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                      netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                              netboxapi.FilterStrings(params.SetTag),
		"tag__n":                           netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                         netboxapi.FilterStrings(params.SetTagAny),
		"created":                          netboxapi.FilterStrings(params.SetCreated),
		"created__n":                       netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                      netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":                     netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                      netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":                     netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":                   netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":                      netboxapi.FilterStrings(params.SetDescription),
		"description__n":                   netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":                  netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":                 netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":                  netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":                 netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":                 netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":                netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":                 netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":                netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":               netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":               netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":              netboxapi.FilterStrings(params.SetDescriptionIregex),
		"ipsec_policy":                     netboxapi.FilterStrings(params.SetIpsecPolicy),
		"ipsec_policy__n":                  netboxapi.FilterStrings(params.SetIpsecPolicyn),
		"ipsec_policy_id":                  netboxapi.FilterInt64s(params.SetIpsecPolicyID),
		"ipsec_policy_id__n":               netboxapi.FilterInt64s(params.SetIpsecPolicyIDn),
		"last_updated":                     netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":                  netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":                 netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":                netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":                 netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":                netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":              netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                            netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                         netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                      netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":                   netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":                   netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":                netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                                netboxapi.FilterString(params.SetQ),
		"sa_lifetime_data":                 netboxapi.FilterInt64s(params.SetSaLifetimeData),
		"sa_lifetime_data__n":              netboxapi.FilterInt64s(params.SetSaLifetimeDatan),
		"sa_lifetime_data__lt":             netboxapi.FilterInt64s(params.SetSaLifetimeDataLt),
		"sa_lifetime_data__lte":            netboxapi.FilterInt64s(params.SetSaLifetimeDataLte),
		"sa_lifetime_data__gt":             netboxapi.FilterInt64s(params.SetSaLifetimeDataGt),
		"sa_lifetime_data__gte":            netboxapi.FilterInt64s(params.SetSaLifetimeDataGte),
		"sa_lifetime_data__empty":          netboxapi.FilterBool(params.SetSaLifetimeDataEmpty),
		"sa_lifetime_seconds":              netboxapi.FilterInt64s(params.SetSaLifetimeSeconds),
		"sa_lifetime_seconds__n":           netboxapi.FilterInt64s(params.SetSaLifetimeSecondsn),
		"sa_lifetime_seconds__lt":          netboxapi.FilterInt64s(params.SetSaLifetimeSecondsLt),
		"sa_lifetime_seconds__lte":         netboxapi.FilterInt64s(params.SetSaLifetimeSecondsLte),
		"sa_lifetime_seconds__gt":          netboxapi.FilterInt64s(params.SetSaLifetimeSecondsGt),
		"sa_lifetime_seconds__gte":         netboxapi.FilterInt64s(params.SetSaLifetimeSecondsGte),
		"sa_lifetime_seconds__empty":       netboxapi.FilterBool(params.SetSaLifetimeSecondsEmpty),
		"tag_id":                           netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                        netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                      netboxapi.FilterInt64s(params.SetTagIDAny),
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
