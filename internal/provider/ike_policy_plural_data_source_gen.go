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
	var vMode []string
	var vModen []string
	var vModeIc []string
	var vModeNic []string
	var vModeIe []string
	var vModeNie []string
	var vModeIsw []string
	var vModeNisw []string
	var vModeIew []string
	var vModeNiew []string
	var vModeEmpty *bool
	var vModeRegex []string
	var vModeIregex []string
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
	var vIkeProposal []string
	var vIkeProposaln []string
	var vIkeProposalID []int64
	var vIkeProposalIDn []int64
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
	var vPresharedKey *string
	var vPresharedKeyn *string
	var vPresharedKeyIc *string
	var vPresharedKeyNic *string
	var vPresharedKeyIe *string
	var vPresharedKeyNie *string
	var vPresharedKeyIsw *string
	var vPresharedKeyNisw *string
	var vPresharedKeyIew *string
	var vPresharedKeyNiew *string
	var vPresharedKeyRegex *string
	var vPresharedKeyIregex *string
	var vQ *string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vVersion []int64
	var vVersionn []int64
	var vVersionIc []int64
	var vVersionNic []int64
	var vVersionIe []int64
	var vVersionNie []int64
	var vVersionIsw []int64
	var vVersionNisw []int64
	var vVersionIew []int64
	var vVersionNiew []int64
	var vVersionRegex []int64
	var vVersionIregex []int64
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
		case "mode":
			v := value
			vMode = append(vMode, v)
		case "mode__n":
			v := value
			vModen = append(vModen, v)
		case "mode__ic":
			v := value
			vModeIc = append(vModeIc, v)
		case "mode__nic":
			v := value
			vModeNic = append(vModeNic, v)
		case "mode__ie":
			v := value
			vModeIe = append(vModeIe, v)
		case "mode__nie":
			v := value
			vModeNie = append(vModeNie, v)
		case "mode__isw":
			v := value
			vModeIsw = append(vModeIsw, v)
		case "mode__nisw":
			v := value
			vModeNisw = append(vModeNisw, v)
		case "mode__iew":
			v := value
			vModeIew = append(vModeIew, v)
		case "mode__niew":
			v := value
			vModeNiew = append(vModeNiew, v)
		case "mode__empty":
			if vModeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mode__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mode__empty' takes a boolean, got %q.", value))
				return
			}
			vModeEmpty = &v
		case "mode__regex":
			v := value
			vModeRegex = append(vModeRegex, v)
		case "mode__iregex":
			v := value
			vModeIregex = append(vModeIregex, v)
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
		case "ike_proposal":
			v := value
			vIkeProposal = append(vIkeProposal, v)
		case "ike_proposal__n":
			v := value
			vIkeProposaln = append(vIkeProposaln, v)
		case "ike_proposal_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ike_proposal_id' takes an integer, got %q.", value))
				return
			}
			vIkeProposalID = append(vIkeProposalID, v)
		case "ike_proposal_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ike_proposal_id__n' takes an integer, got %q.", value))
				return
			}
			vIkeProposalIDn = append(vIkeProposalIDn, v)
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
		case "preshared_key":
			if vPresharedKey != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key' takes a single value.")
				return
			}
			v := value
			vPresharedKey = &v
		case "preshared_key__n":
			if vPresharedKeyn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__n' takes a single value.")
				return
			}
			v := value
			vPresharedKeyn = &v
		case "preshared_key__ic":
			if vPresharedKeyIc != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__ic' takes a single value.")
				return
			}
			v := value
			vPresharedKeyIc = &v
		case "preshared_key__nic":
			if vPresharedKeyNic != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__nic' takes a single value.")
				return
			}
			v := value
			vPresharedKeyNic = &v
		case "preshared_key__ie":
			if vPresharedKeyIe != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__ie' takes a single value.")
				return
			}
			v := value
			vPresharedKeyIe = &v
		case "preshared_key__nie":
			if vPresharedKeyNie != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__nie' takes a single value.")
				return
			}
			v := value
			vPresharedKeyNie = &v
		case "preshared_key__isw":
			if vPresharedKeyIsw != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__isw' takes a single value.")
				return
			}
			v := value
			vPresharedKeyIsw = &v
		case "preshared_key__nisw":
			if vPresharedKeyNisw != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__nisw' takes a single value.")
				return
			}
			v := value
			vPresharedKeyNisw = &v
		case "preshared_key__iew":
			if vPresharedKeyIew != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__iew' takes a single value.")
				return
			}
			v := value
			vPresharedKeyIew = &v
		case "preshared_key__niew":
			if vPresharedKeyNiew != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__niew' takes a single value.")
				return
			}
			v := value
			vPresharedKeyNiew = &v
		case "preshared_key__regex":
			if vPresharedKeyRegex != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__regex' takes a single value.")
				return
			}
			v := value
			vPresharedKeyRegex = &v
		case "preshared_key__iregex":
			if vPresharedKeyIregex != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'preshared_key__iregex' takes a single value.")
				return
			}
			v := value
			vPresharedKeyIregex = &v
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
		case "version":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version' takes an integer, got %q.", value))
				return
			}
			vVersion = append(vVersion, v)
		case "version__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__n' takes an integer, got %q.", value))
				return
			}
			vVersionn = append(vVersionn, v)
		case "version__ic":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__ic' takes an integer, got %q.", value))
				return
			}
			vVersionIc = append(vVersionIc, v)
		case "version__nic":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__nic' takes an integer, got %q.", value))
				return
			}
			vVersionNic = append(vVersionNic, v)
		case "version__ie":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__ie' takes an integer, got %q.", value))
				return
			}
			vVersionIe = append(vVersionIe, v)
		case "version__nie":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__nie' takes an integer, got %q.", value))
				return
			}
			vVersionNie = append(vVersionNie, v)
		case "version__isw":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__isw' takes an integer, got %q.", value))
				return
			}
			vVersionIsw = append(vVersionIsw, v)
		case "version__nisw":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__nisw' takes an integer, got %q.", value))
				return
			}
			vVersionNisw = append(vVersionNisw, v)
		case "version__iew":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__iew' takes an integer, got %q.", value))
				return
			}
			vVersionIew = append(vVersionIew, v)
		case "version__niew":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__niew' takes an integer, got %q.", value))
				return
			}
			vVersionNiew = append(vVersionNiew, v)
		case "version__regex":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__regex' takes an integer, got %q.", value))
				return
			}
			vVersionRegex = append(vVersionRegex, v)
		case "version__iregex":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__iregex' takes an integer, got %q.", value))
				return
			}
			vVersionIregex = append(vVersionIregex, v)
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
	if len(vMode) > 0 {
		params.SetMode(vMode)
	}
	if len(vModen) > 0 {
		params.SetModen(vModen)
	}
	if len(vModeIc) > 0 {
		params.SetModeIc(vModeIc)
	}
	if len(vModeNic) > 0 {
		params.SetModeNic(vModeNic)
	}
	if len(vModeIe) > 0 {
		params.SetModeIe(vModeIe)
	}
	if len(vModeNie) > 0 {
		params.SetModeNie(vModeNie)
	}
	if len(vModeIsw) > 0 {
		params.SetModeIsw(vModeIsw)
	}
	if len(vModeNisw) > 0 {
		params.SetModeNisw(vModeNisw)
	}
	if len(vModeIew) > 0 {
		params.SetModeIew(vModeIew)
	}
	if len(vModeNiew) > 0 {
		params.SetModeNiew(vModeNiew)
	}
	if vModeEmpty != nil {
		params.SetModeEmpty(vModeEmpty)
	}
	if len(vModeRegex) > 0 {
		params.SetModeRegex(vModeRegex)
	}
	if len(vModeIregex) > 0 {
		params.SetModeIregex(vModeIregex)
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
	if len(vIkeProposal) > 0 {
		params.SetIkeProposal(vIkeProposal)
	}
	if len(vIkeProposaln) > 0 {
		params.SetIkeProposaln(vIkeProposaln)
	}
	if len(vIkeProposalID) > 0 {
		params.SetIkeProposalID(vIkeProposalID)
	}
	if len(vIkeProposalIDn) > 0 {
		params.SetIkeProposalIDn(vIkeProposalIDn)
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
	if vPresharedKey != nil {
		params.SetPresharedKey(vPresharedKey)
	}
	if vPresharedKeyn != nil {
		params.SetPresharedKeyn(vPresharedKeyn)
	}
	if vPresharedKeyIc != nil {
		params.SetPresharedKeyIc(vPresharedKeyIc)
	}
	if vPresharedKeyNic != nil {
		params.SetPresharedKeyNic(vPresharedKeyNic)
	}
	if vPresharedKeyIe != nil {
		params.SetPresharedKeyIe(vPresharedKeyIe)
	}
	if vPresharedKeyNie != nil {
		params.SetPresharedKeyNie(vPresharedKeyNie)
	}
	if vPresharedKeyIsw != nil {
		params.SetPresharedKeyIsw(vPresharedKeyIsw)
	}
	if vPresharedKeyNisw != nil {
		params.SetPresharedKeyNisw(vPresharedKeyNisw)
	}
	if vPresharedKeyIew != nil {
		params.SetPresharedKeyIew(vPresharedKeyIew)
	}
	if vPresharedKeyNiew != nil {
		params.SetPresharedKeyNiew(vPresharedKeyNiew)
	}
	if vPresharedKeyRegex != nil {
		params.SetPresharedKeyRegex(vPresharedKeyRegex)
	}
	if vPresharedKeyIregex != nil {
		params.SetPresharedKeyIregex(vPresharedKeyIregex)
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
	if len(vVersion) > 0 {
		params.SetVersion(vVersion)
	}
	if len(vVersionn) > 0 {
		params.SetVersionn(vVersionn)
	}
	if len(vVersionIc) > 0 {
		params.SetVersionIc(vVersionIc)
	}
	if len(vVersionNic) > 0 {
		params.SetVersionNic(vVersionNic)
	}
	if len(vVersionIe) > 0 {
		params.SetVersionIe(vVersionIe)
	}
	if len(vVersionNie) > 0 {
		params.SetVersionNie(vVersionNie)
	}
	if len(vVersionIsw) > 0 {
		params.SetVersionIsw(vVersionIsw)
	}
	if len(vVersionNisw) > 0 {
		params.SetVersionNisw(vVersionNisw)
	}
	if len(vVersionIew) > 0 {
		params.SetVersionIew(vVersionIew)
	}
	if len(vVersionNiew) > 0 {
		params.SetVersionNiew(vVersionNiew)
	}
	if len(vVersionRegex) > 0 {
		params.SetVersionRegex(vVersionRegex)
	}
	if len(vVersionIregex) > 0 {
		params.SetVersionIregex(vVersionIregex)
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
