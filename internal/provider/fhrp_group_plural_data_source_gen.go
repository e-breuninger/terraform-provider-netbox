// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vProtocol []string
	var vProtocoln []string
	var vProtocolIc []string
	var vProtocolNic []string
	var vProtocolIe []string
	var vProtocolNie []string
	var vProtocolIsw []string
	var vProtocolNisw []string
	var vProtocolIew []string
	var vProtocolNiew []string
	var vProtocolEmpty *bool
	var vProtocolRegex []string
	var vProtocolIregex []string
	var vGroupID []int64
	var vGroupIDn []int64
	var vGroupIDLt []int64
	var vGroupIDLte []int64
	var vGroupIDGt []int64
	var vGroupIDGte []int64
	var vGroupIDEmpty *bool
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
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAuthKey []string
	var vAuthKeyn []string
	var vAuthKeyIc []string
	var vAuthKeyNic []string
	var vAuthKeyIe []string
	var vAuthKeyNie []string
	var vAuthKeyIsw []string
	var vAuthKeyNisw []string
	var vAuthKeyIew []string
	var vAuthKeyNiew []string
	var vAuthKeyEmpty *bool
	var vAuthKeyRegex []string
	var vAuthKeyIregex []string
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
	var vRelatedIP []string
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
		case "protocol":
			v := value
			vProtocol = append(vProtocol, v)
		case "protocol__n":
			v := value
			vProtocoln = append(vProtocoln, v)
		case "protocol__ic":
			v := value
			vProtocolIc = append(vProtocolIc, v)
		case "protocol__nic":
			v := value
			vProtocolNic = append(vProtocolNic, v)
		case "protocol__ie":
			v := value
			vProtocolIe = append(vProtocolIe, v)
		case "protocol__nie":
			v := value
			vProtocolNie = append(vProtocolNie, v)
		case "protocol__isw":
			v := value
			vProtocolIsw = append(vProtocolIsw, v)
		case "protocol__nisw":
			v := value
			vProtocolNisw = append(vProtocolNisw, v)
		case "protocol__iew":
			v := value
			vProtocolIew = append(vProtocolIew, v)
		case "protocol__niew":
			v := value
			vProtocolNiew = append(vProtocolNiew, v)
		case "protocol__empty":
			if vProtocolEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'protocol__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'protocol__empty' takes a boolean, got %q.", value))
				return
			}
			vProtocolEmpty = &v
		case "protocol__regex":
			v := value
			vProtocolRegex = append(vProtocolRegex, v)
		case "protocol__iregex":
			v := value
			vProtocolIregex = append(vProtocolIregex, v)
		case "group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id' takes an integer, got %q.", value))
				return
			}
			vGroupID = append(vGroupID, v)
		case "group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__n' takes an integer, got %q.", value))
				return
			}
			vGroupIDn = append(vGroupIDn, v)
		case "group_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__lt' takes an integer, got %q.", value))
				return
			}
			vGroupIDLt = append(vGroupIDLt, v)
		case "group_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__lte' takes an integer, got %q.", value))
				return
			}
			vGroupIDLte = append(vGroupIDLte, v)
		case "group_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__gt' takes an integer, got %q.", value))
				return
			}
			vGroupIDGt = append(vGroupIDGt, v)
		case "group_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__gte' takes an integer, got %q.", value))
				return
			}
			vGroupIDGte = append(vGroupIDGte, v)
		case "group_id__empty":
			if vGroupIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'group_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__empty' takes a boolean, got %q.", value))
				return
			}
			vGroupIDEmpty = &v
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
		case "auth_key":
			v := value
			vAuthKey = append(vAuthKey, v)
		case "auth_key__n":
			v := value
			vAuthKeyn = append(vAuthKeyn, v)
		case "auth_key__ic":
			v := value
			vAuthKeyIc = append(vAuthKeyIc, v)
		case "auth_key__nic":
			v := value
			vAuthKeyNic = append(vAuthKeyNic, v)
		case "auth_key__ie":
			v := value
			vAuthKeyIe = append(vAuthKeyIe, v)
		case "auth_key__nie":
			v := value
			vAuthKeyNie = append(vAuthKeyNie, v)
		case "auth_key__isw":
			v := value
			vAuthKeyIsw = append(vAuthKeyIsw, v)
		case "auth_key__nisw":
			v := value
			vAuthKeyNisw = append(vAuthKeyNisw, v)
		case "auth_key__iew":
			v := value
			vAuthKeyIew = append(vAuthKeyIew, v)
		case "auth_key__niew":
			v := value
			vAuthKeyNiew = append(vAuthKeyNiew, v)
		case "auth_key__empty":
			if vAuthKeyEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'auth_key__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'auth_key__empty' takes a boolean, got %q.", value))
				return
			}
			vAuthKeyEmpty = &v
		case "auth_key__regex":
			v := value
			vAuthKeyRegex = append(vAuthKeyRegex, v)
		case "auth_key__iregex":
			v := value
			vAuthKeyIregex = append(vAuthKeyIregex, v)
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
		case "related_ip":
			v := value
			vRelatedIP = append(vRelatedIP, v)
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
	if len(vProtocol) > 0 {
		params.SetProtocol(vProtocol)
	}
	if len(vProtocoln) > 0 {
		params.SetProtocoln(vProtocoln)
	}
	if len(vProtocolIc) > 0 {
		params.SetProtocolIc(vProtocolIc)
	}
	if len(vProtocolNic) > 0 {
		params.SetProtocolNic(vProtocolNic)
	}
	if len(vProtocolIe) > 0 {
		params.SetProtocolIe(vProtocolIe)
	}
	if len(vProtocolNie) > 0 {
		params.SetProtocolNie(vProtocolNie)
	}
	if len(vProtocolIsw) > 0 {
		params.SetProtocolIsw(vProtocolIsw)
	}
	if len(vProtocolNisw) > 0 {
		params.SetProtocolNisw(vProtocolNisw)
	}
	if len(vProtocolIew) > 0 {
		params.SetProtocolIew(vProtocolIew)
	}
	if len(vProtocolNiew) > 0 {
		params.SetProtocolNiew(vProtocolNiew)
	}
	if vProtocolEmpty != nil {
		params.SetProtocolEmpty(vProtocolEmpty)
	}
	if len(vProtocolRegex) > 0 {
		params.SetProtocolRegex(vProtocolRegex)
	}
	if len(vProtocolIregex) > 0 {
		params.SetProtocolIregex(vProtocolIregex)
	}
	if len(vGroupID) > 0 {
		params.SetGroupID(vGroupID)
	}
	if len(vGroupIDn) > 0 {
		params.SetGroupIDn(vGroupIDn)
	}
	if len(vGroupIDLt) > 0 {
		params.SetGroupIDLt(vGroupIDLt)
	}
	if len(vGroupIDLte) > 0 {
		params.SetGroupIDLte(vGroupIDLte)
	}
	if len(vGroupIDGt) > 0 {
		params.SetGroupIDGt(vGroupIDGt)
	}
	if len(vGroupIDGte) > 0 {
		params.SetGroupIDGte(vGroupIDGte)
	}
	if vGroupIDEmpty != nil {
		params.SetGroupIDEmpty(vGroupIDEmpty)
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
	if len(vAuthKey) > 0 {
		params.SetAuthKey(vAuthKey)
	}
	if len(vAuthKeyn) > 0 {
		params.SetAuthKeyn(vAuthKeyn)
	}
	if len(vAuthKeyIc) > 0 {
		params.SetAuthKeyIc(vAuthKeyIc)
	}
	if len(vAuthKeyNic) > 0 {
		params.SetAuthKeyNic(vAuthKeyNic)
	}
	if len(vAuthKeyIe) > 0 {
		params.SetAuthKeyIe(vAuthKeyIe)
	}
	if len(vAuthKeyNie) > 0 {
		params.SetAuthKeyNie(vAuthKeyNie)
	}
	if len(vAuthKeyIsw) > 0 {
		params.SetAuthKeyIsw(vAuthKeyIsw)
	}
	if len(vAuthKeyNisw) > 0 {
		params.SetAuthKeyNisw(vAuthKeyNisw)
	}
	if len(vAuthKeyIew) > 0 {
		params.SetAuthKeyIew(vAuthKeyIew)
	}
	if len(vAuthKeyNiew) > 0 {
		params.SetAuthKeyNiew(vAuthKeyNiew)
	}
	if vAuthKeyEmpty != nil {
		params.SetAuthKeyEmpty(vAuthKeyEmpty)
	}
	if len(vAuthKeyRegex) > 0 {
		params.SetAuthKeyRegex(vAuthKeyRegex)
	}
	if len(vAuthKeyIregex) > 0 {
		params.SetAuthKeyIregex(vAuthKeyIregex)
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
	if len(vRelatedIP) > 0 {
		params.SetRelatedIP(vRelatedIP)
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
