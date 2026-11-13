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
	_ datasource.DataSource              = (*ikePoliciesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ikePoliciesDataSource)(nil)
)

// NewIkePoliciesDataSource returns a new ike_policies data source, which
// lists ike_policy objects matching its filters.
func NewIkePoliciesDataSource() datasource.DataSource {
	return &ikePoliciesDataSource{}
}

type ikePoliciesDataSource struct {
	client *netboxapi.Client
}

// ikePoliciesDataSourceModel is the data source model: the filters, the limit and the matching
// ike_policies.
type ikePoliciesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"ike_policies"`
}

// ikePolicyResourceAttrTypes is the attribute type map of ikePolicyResourceModel.
func ikePolicyResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":               types.Int64Type,
		"name":             types.StringType,
		"version":          types.Int64Type,
		"mode":             types.StringType,
		"ike_proposal_ids": types.SetType{ElemType: types.Int64Type},
		"preshared_key":    types.StringType,
		"description":      types.StringType,
		"comments":         types.StringType,
		"owner_id":         types.Int64Type,
		"created":          types.StringType,
		"last_updated":     types.StringType,
		"url":              types.StringType,
		"tags":             types.SetType{ElemType: types.StringType},
		"tags_all":         types.SetType{ElemType: types.StringType},
		"custom_fields":    types.MapType{ElemType: types.StringType},
	}
}

func (d *ikePoliciesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ike_policies"
}

func (d *ikePoliciesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:Lists ike_policy objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, ike_proposal, ike_proposal__n, ike_proposal_id, ike_proposal_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mode, mode__empty, mode__ic, mode__ie, mode__iew, mode__iregex, mode__isw, mode__n, mode__nic, mode__nie, mode__niew, mode__nisw, mode__regex, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, preshared_key, preshared_key__ic, preshared_key__ie, preshared_key__iew, preshared_key__iregex, preshared_key__isw, preshared_key__n, preshared_key__nic, preshared_key__nie, preshared_key__niew, preshared_key__nisw, preshared_key__regex, q, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, version, version__ic, version__ie, version__iew, version__iregex, version__isw, version__n, version__nic, version__nie, version__niew, version__nisw, version__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"ike_policies": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching ike_policies.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"version": schema.Int64Attribute{
							Computed:    true,
							Description: "IKE version. One of: 1, 2.",
						},
						"mode": schema.StringAttribute{
							Computed:    true,
							Description: "IKEv1 negotiation mode; IKEv2 has none. NetBox does not allow clearing it once set, so a policy that has a mode cannot be moved to version 2; leaving it unset keeps the current value. One of: aggressive, main.",
						},
						"ike_proposal_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the IKE proposals the policy offers.",
						},
						"preshared_key": schema.StringAttribute{
							Computed:    true,
							Description: "Pre-shared key, when the proposals authenticate with one.",
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

func (d *ikePoliciesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ikePoliciesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ikePoliciesDataSourceModel

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
	var responseDTOs []*netboxapi.IkePolicyResponseDTO

	params := vpn.NewVpnIkePoliciesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"name":                  netboxapi.FilterStrings(params.SetName),
		"name__n":               netboxapi.FilterStrings(params.SetNamen),
		"name__ic":              netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":             netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":              netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":             netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":             netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":            netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":             netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":            netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":           netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":           netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":          netboxapi.FilterStrings(params.SetNameIregex),
		"mode":                  netboxapi.FilterStrings(params.SetMode),
		"mode__n":               netboxapi.FilterStrings(params.SetModen),
		"mode__ic":              netboxapi.FilterStrings(params.SetModeIc),
		"mode__nic":             netboxapi.FilterStrings(params.SetModeNic),
		"mode__ie":              netboxapi.FilterStrings(params.SetModeIe),
		"mode__nie":             netboxapi.FilterStrings(params.SetModeNie),
		"mode__isw":             netboxapi.FilterStrings(params.SetModeIsw),
		"mode__nisw":            netboxapi.FilterStrings(params.SetModeNisw),
		"mode__iew":             netboxapi.FilterStrings(params.SetModeIew),
		"mode__niew":            netboxapi.FilterStrings(params.SetModeNiew),
		"mode__empty":           netboxapi.FilterBool(params.SetModeEmpty),
		"mode__regex":           netboxapi.FilterStrings(params.SetModeRegex),
		"mode__iregex":          netboxapi.FilterStrings(params.SetModeIregex),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
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
		"ike_proposal":          netboxapi.FilterStrings(params.SetIkeProposal),
		"ike_proposal__n":       netboxapi.FilterStrings(params.SetIkeProposaln),
		"ike_proposal_id":       netboxapi.FilterInt64s(params.SetIkeProposalID),
		"ike_proposal_id__n":    netboxapi.FilterInt64s(params.SetIkeProposalIDn),
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
		"preshared_key":         netboxapi.FilterString(params.SetPresharedKey),
		"preshared_key__n":      netboxapi.FilterString(params.SetPresharedKeyn),
		"preshared_key__ic":     netboxapi.FilterString(params.SetPresharedKeyIc),
		"preshared_key__nic":    netboxapi.FilterString(params.SetPresharedKeyNic),
		"preshared_key__ie":     netboxapi.FilterString(params.SetPresharedKeyIe),
		"preshared_key__nie":    netboxapi.FilterString(params.SetPresharedKeyNie),
		"preshared_key__isw":    netboxapi.FilterString(params.SetPresharedKeyIsw),
		"preshared_key__nisw":   netboxapi.FilterString(params.SetPresharedKeyNisw),
		"preshared_key__iew":    netboxapi.FilterString(params.SetPresharedKeyIew),
		"preshared_key__niew":   netboxapi.FilterString(params.SetPresharedKeyNiew),
		"preshared_key__regex":  netboxapi.FilterString(params.SetPresharedKeyRegex),
		"preshared_key__iregex": netboxapi.FilterString(params.SetPresharedKeyIregex),
		"q":                     netboxapi.FilterString(params.SetQ),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"version":               netboxapi.FilterInt64s(params.SetVersion),
		"version__n":            netboxapi.FilterInt64s(params.SetVersionn),
		"version__ic":           netboxapi.FilterInt64s(params.SetVersionIc),
		"version__nic":          netboxapi.FilterInt64s(params.SetVersionNic),
		"version__ie":           netboxapi.FilterInt64s(params.SetVersionIe),
		"version__nie":          netboxapi.FilterInt64s(params.SetVersionNie),
		"version__isw":          netboxapi.FilterInt64s(params.SetVersionIsw),
		"version__nisw":         netboxapi.FilterInt64s(params.SetVersionNisw),
		"version__iew":          netboxapi.FilterInt64s(params.SetVersionIew),
		"version__niew":         netboxapi.FilterInt64s(params.SetVersionNiew),
		"version__regex":        netboxapi.FilterInt64s(params.SetVersionRegex),
		"version__iregex":       netboxapi.FilterInt64s(params.SetVersionIregex),
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
		res, err := d.client.Vpn.VpnIkePoliciesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_ike_policies", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.IkePolicyResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]ikePolicyResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model ikePolicyResourceModel
		resp.Diagnostics.Append(flattenIkePolicy(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ikePolicyResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
