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
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
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
	Filters   types.Set    `tfsdk:"filters"`
	NameRegex types.String `tfsdk:"name_regex"`
	Limit     types.Int64  `tfsdk:"limit"`
	Items     types.List   `tfsdk:"fhrp_groups"`
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
				Description: "Query filters, sent as API list parameters. Supported names: auth_type, auth_type__empty, auth_type__ic, auth_type__ie, auth_type__iew, auth_type__iregex, auth_type__isw, auth_type__n, auth_type__nic, auth_type__nie, auth_type__niew, auth_type__nisw, auth_type__regex, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner_id, owner_id__n, protocol, protocol__empty, protocol__ic, protocol__ie, protocol__iew, protocol__iregex, protocol__isw, protocol__n, protocol__nic, protocol__nie, protocol__niew, protocol__nisw, protocol__regex, tag, tag__any, tag__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
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
	var vNameIc []string
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
		case "name__ic":
			v := value
			vNameIc = append(vNameIc, v)
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
	if len(vNameIc) > 0 {
		params.SetNameIc(vNameIc)
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
		res, err := d.client.Ipam.IpamFhrpGroupsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_fhrp_groups", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.FhrpGroupResponseDTOFromGoNetbox(goNetboxModel))
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
	items := make([]fhrpGroupResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model fhrpGroupResourceModel
		resp.Diagnostics.Append(flattenFhrpGroup(ctx, responseDTO, &model)...)
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
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: fhrpGroupResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
