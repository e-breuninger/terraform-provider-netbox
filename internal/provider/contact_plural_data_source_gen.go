// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*contactsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*contactsDataSource)(nil)
)

// NewContactsDataSource returns a new contacts data source, which
// lists contact objects matching its filters.
func NewContactsDataSource() datasource.DataSource {
	return &contactsDataSource{}
}

type contactsDataSource struct {
	client *netboxapi.Client
}

// contactsDataSourceModel is the data source model: the filters, the limit and the matching
// contacts.
type contactsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"contacts"`
}

// contactResourceAttrTypes is the attribute type map of contactResourceModel.
func contactResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"name":          types.StringType,
		"title":         types.StringType,
		"phone":         types.StringType,
		"email":         types.StringType,
		"address":       types.StringType,
		"link":          types.StringType,
		"group_ids":     types.SetType{ElemType: types.Int64Type},
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

func (d *contactsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contacts"
}

func (d *contactsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Tenancy:Lists contact objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: address, address__empty, address__ic, address__ie, address__iew, address__iregex, address__isw, address__n, address__nic, address__nie, address__niew, address__nisw, address__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, email, email__empty, email__ic, email__ie, email__iew, email__iregex, email__isw, email__n, email__nic, email__nie, email__niew, email__nisw, email__regex, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, link, link__empty, link__ic, link__ie, link__iew, link__iregex, link__isw, link__n, link__nic, link__nie, link__niew, link__nisw, link__regex, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, phone, phone__empty, phone__ic, phone__ie, phone__iew, phone__iregex, phone__isw, phone__n, phone__nic, phone__nie, phone__niew, phone__nisw, phone__regex, q, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, title, title__empty, title__ic, title__ie, title__iew, title__iregex, title__isw, title__n, title__nic, title__nie, title__niew, title__nisw, title__regex. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"contacts": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching contacts.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"title": schema.StringAttribute{
							Computed: true,
						},
						"phone": schema.StringAttribute{
							Computed: true,
						},
						"email": schema.StringAttribute{
							Computed:    true,
							Description: "Email address.",
						},
						"address": schema.StringAttribute{
							Computed: true,
						},
						"link": schema.StringAttribute{
							Computed:    true,
							Description: "URL related to the contact.",
						},
						"group_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of contact groups this contact belongs to.",
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

func (d *contactsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *contactsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data contactsDataSourceModel

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
	var responseDTOs []*netboxapi.ContactResponseDTO

	params := tenancy.NewTenancyContactsListParams()
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
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAddress []string
	var vAddressn []string
	var vAddressIc []string
	var vAddressNic []string
	var vAddressIe []string
	var vAddressNie []string
	var vAddressIsw []string
	var vAddressNisw []string
	var vAddressIew []string
	var vAddressNiew []string
	var vAddressEmpty *bool
	var vAddressRegex []string
	var vAddressIregex []string
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
	var vEmail []string
	var vEmailn []string
	var vEmailIc []string
	var vEmailNic []string
	var vEmailIe []string
	var vEmailNie []string
	var vEmailIsw []string
	var vEmailNisw []string
	var vEmailIew []string
	var vEmailNiew []string
	var vEmailEmpty *bool
	var vEmailRegex []string
	var vEmailIregex []string
	var vGroup []string
	var vGroupn []string
	var vGroupID []string
	var vGroupIDn []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLink []string
	var vLinkn []string
	var vLinkIc []string
	var vLinkNic []string
	var vLinkIe []string
	var vLinkNie []string
	var vLinkIsw []string
	var vLinkNisw []string
	var vLinkIew []string
	var vLinkNiew []string
	var vLinkEmpty *bool
	var vLinkRegex []string
	var vLinkIregex []string
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPhone []string
	var vPhonen []string
	var vPhoneIc []string
	var vPhoneNic []string
	var vPhoneIe []string
	var vPhoneNie []string
	var vPhoneIsw []string
	var vPhoneNisw []string
	var vPhoneIew []string
	var vPhoneNiew []string
	var vPhoneEmpty *bool
	var vPhoneRegex []string
	var vPhoneIregex []string
	var vQ *string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTitle []string
	var vTitlen []string
	var vTitleIc []string
	var vTitleNic []string
	var vTitleIe []string
	var vTitleNie []string
	var vTitleIsw []string
	var vTitleNisw []string
	var vTitleIew []string
	var vTitleNiew []string
	var vTitleEmpty *bool
	var vTitleRegex []string
	var vTitleIregex []string
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
		case "address":
			v := value
			vAddress = append(vAddress, v)
		case "address__n":
			v := value
			vAddressn = append(vAddressn, v)
		case "address__ic":
			v := value
			vAddressIc = append(vAddressIc, v)
		case "address__nic":
			v := value
			vAddressNic = append(vAddressNic, v)
		case "address__ie":
			v := value
			vAddressIe = append(vAddressIe, v)
		case "address__nie":
			v := value
			vAddressNie = append(vAddressNie, v)
		case "address__isw":
			v := value
			vAddressIsw = append(vAddressIsw, v)
		case "address__nisw":
			v := value
			vAddressNisw = append(vAddressNisw, v)
		case "address__iew":
			v := value
			vAddressIew = append(vAddressIew, v)
		case "address__niew":
			v := value
			vAddressNiew = append(vAddressNiew, v)
		case "address__empty":
			if vAddressEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'address__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'address__empty' takes a boolean, got %q.", value))
				return
			}
			vAddressEmpty = &v
		case "address__regex":
			v := value
			vAddressRegex = append(vAddressRegex, v)
		case "address__iregex":
			v := value
			vAddressIregex = append(vAddressIregex, v)
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
		case "email":
			v := value
			vEmail = append(vEmail, v)
		case "email__n":
			v := value
			vEmailn = append(vEmailn, v)
		case "email__ic":
			v := value
			vEmailIc = append(vEmailIc, v)
		case "email__nic":
			v := value
			vEmailNic = append(vEmailNic, v)
		case "email__ie":
			v := value
			vEmailIe = append(vEmailIe, v)
		case "email__nie":
			v := value
			vEmailNie = append(vEmailNie, v)
		case "email__isw":
			v := value
			vEmailIsw = append(vEmailIsw, v)
		case "email__nisw":
			v := value
			vEmailNisw = append(vEmailNisw, v)
		case "email__iew":
			v := value
			vEmailIew = append(vEmailIew, v)
		case "email__niew":
			v := value
			vEmailNiew = append(vEmailNiew, v)
		case "email__empty":
			if vEmailEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'email__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'email__empty' takes a boolean, got %q.", value))
				return
			}
			vEmailEmpty = &v
		case "email__regex":
			v := value
			vEmailRegex = append(vEmailRegex, v)
		case "email__iregex":
			v := value
			vEmailIregex = append(vEmailIregex, v)
		case "group":
			v := value
			vGroup = append(vGroup, v)
		case "group__n":
			v := value
			vGroupn = append(vGroupn, v)
		case "group_id":
			v := value
			vGroupID = append(vGroupID, v)
		case "group_id__n":
			v := value
			vGroupIDn = append(vGroupIDn, v)
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
		case "link":
			v := value
			vLink = append(vLink, v)
		case "link__n":
			v := value
			vLinkn = append(vLinkn, v)
		case "link__ic":
			v := value
			vLinkIc = append(vLinkIc, v)
		case "link__nic":
			v := value
			vLinkNic = append(vLinkNic, v)
		case "link__ie":
			v := value
			vLinkIe = append(vLinkIe, v)
		case "link__nie":
			v := value
			vLinkNie = append(vLinkNie, v)
		case "link__isw":
			v := value
			vLinkIsw = append(vLinkIsw, v)
		case "link__nisw":
			v := value
			vLinkNisw = append(vLinkNisw, v)
		case "link__iew":
			v := value
			vLinkIew = append(vLinkIew, v)
		case "link__niew":
			v := value
			vLinkNiew = append(vLinkNiew, v)
		case "link__empty":
			if vLinkEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'link__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'link__empty' takes a boolean, got %q.", value))
				return
			}
			vLinkEmpty = &v
		case "link__regex":
			v := value
			vLinkRegex = append(vLinkRegex, v)
		case "link__iregex":
			v := value
			vLinkIregex = append(vLinkIregex, v)
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
		case "phone":
			v := value
			vPhone = append(vPhone, v)
		case "phone__n":
			v := value
			vPhonen = append(vPhonen, v)
		case "phone__ic":
			v := value
			vPhoneIc = append(vPhoneIc, v)
		case "phone__nic":
			v := value
			vPhoneNic = append(vPhoneNic, v)
		case "phone__ie":
			v := value
			vPhoneIe = append(vPhoneIe, v)
		case "phone__nie":
			v := value
			vPhoneNie = append(vPhoneNie, v)
		case "phone__isw":
			v := value
			vPhoneIsw = append(vPhoneIsw, v)
		case "phone__nisw":
			v := value
			vPhoneNisw = append(vPhoneNisw, v)
		case "phone__iew":
			v := value
			vPhoneIew = append(vPhoneIew, v)
		case "phone__niew":
			v := value
			vPhoneNiew = append(vPhoneNiew, v)
		case "phone__empty":
			if vPhoneEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'phone__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'phone__empty' takes a boolean, got %q.", value))
				return
			}
			vPhoneEmpty = &v
		case "phone__regex":
			v := value
			vPhoneRegex = append(vPhoneRegex, v)
		case "phone__iregex":
			v := value
			vPhoneIregex = append(vPhoneIregex, v)
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
		case "title":
			v := value
			vTitle = append(vTitle, v)
		case "title__n":
			v := value
			vTitlen = append(vTitlen, v)
		case "title__ic":
			v := value
			vTitleIc = append(vTitleIc, v)
		case "title__nic":
			v := value
			vTitleNic = append(vTitleNic, v)
		case "title__ie":
			v := value
			vTitleIe = append(vTitleIe, v)
		case "title__nie":
			v := value
			vTitleNie = append(vTitleNie, v)
		case "title__isw":
			v := value
			vTitleIsw = append(vTitleIsw, v)
		case "title__nisw":
			v := value
			vTitleNisw = append(vTitleNisw, v)
		case "title__iew":
			v := value
			vTitleIew = append(vTitleIew, v)
		case "title__niew":
			v := value
			vTitleNiew = append(vTitleNiew, v)
		case "title__empty":
			if vTitleEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'title__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'title__empty' takes a boolean, got %q.", value))
				return
			}
			vTitleEmpty = &v
		case "title__regex":
			v := value
			vTitleRegex = append(vTitleRegex, v)
		case "title__iregex":
			v := value
			vTitleIregex = append(vTitleIregex, v)
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
	if len(vAddress) > 0 {
		params.SetAddress(vAddress)
	}
	if len(vAddressn) > 0 {
		params.SetAddressn(vAddressn)
	}
	if len(vAddressIc) > 0 {
		params.SetAddressIc(vAddressIc)
	}
	if len(vAddressNic) > 0 {
		params.SetAddressNic(vAddressNic)
	}
	if len(vAddressIe) > 0 {
		params.SetAddressIe(vAddressIe)
	}
	if len(vAddressNie) > 0 {
		params.SetAddressNie(vAddressNie)
	}
	if len(vAddressIsw) > 0 {
		params.SetAddressIsw(vAddressIsw)
	}
	if len(vAddressNisw) > 0 {
		params.SetAddressNisw(vAddressNisw)
	}
	if len(vAddressIew) > 0 {
		params.SetAddressIew(vAddressIew)
	}
	if len(vAddressNiew) > 0 {
		params.SetAddressNiew(vAddressNiew)
	}
	if vAddressEmpty != nil {
		params.SetAddressEmpty(vAddressEmpty)
	}
	if len(vAddressRegex) > 0 {
		params.SetAddressRegex(vAddressRegex)
	}
	if len(vAddressIregex) > 0 {
		params.SetAddressIregex(vAddressIregex)
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
	if len(vEmail) > 0 {
		params.SetEmail(vEmail)
	}
	if len(vEmailn) > 0 {
		params.SetEmailn(vEmailn)
	}
	if len(vEmailIc) > 0 {
		params.SetEmailIc(vEmailIc)
	}
	if len(vEmailNic) > 0 {
		params.SetEmailNic(vEmailNic)
	}
	if len(vEmailIe) > 0 {
		params.SetEmailIe(vEmailIe)
	}
	if len(vEmailNie) > 0 {
		params.SetEmailNie(vEmailNie)
	}
	if len(vEmailIsw) > 0 {
		params.SetEmailIsw(vEmailIsw)
	}
	if len(vEmailNisw) > 0 {
		params.SetEmailNisw(vEmailNisw)
	}
	if len(vEmailIew) > 0 {
		params.SetEmailIew(vEmailIew)
	}
	if len(vEmailNiew) > 0 {
		params.SetEmailNiew(vEmailNiew)
	}
	if vEmailEmpty != nil {
		params.SetEmailEmpty(vEmailEmpty)
	}
	if len(vEmailRegex) > 0 {
		params.SetEmailRegex(vEmailRegex)
	}
	if len(vEmailIregex) > 0 {
		params.SetEmailIregex(vEmailIregex)
	}
	if len(vGroup) > 0 {
		params.SetGroup(vGroup)
	}
	if len(vGroupn) > 0 {
		params.SetGroupn(vGroupn)
	}
	if len(vGroupID) > 0 {
		params.SetGroupID(vGroupID)
	}
	if len(vGroupIDn) > 0 {
		params.SetGroupIDn(vGroupIDn)
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
	if len(vLink) > 0 {
		params.SetLink(vLink)
	}
	if len(vLinkn) > 0 {
		params.SetLinkn(vLinkn)
	}
	if len(vLinkIc) > 0 {
		params.SetLinkIc(vLinkIc)
	}
	if len(vLinkNic) > 0 {
		params.SetLinkNic(vLinkNic)
	}
	if len(vLinkIe) > 0 {
		params.SetLinkIe(vLinkIe)
	}
	if len(vLinkNie) > 0 {
		params.SetLinkNie(vLinkNie)
	}
	if len(vLinkIsw) > 0 {
		params.SetLinkIsw(vLinkIsw)
	}
	if len(vLinkNisw) > 0 {
		params.SetLinkNisw(vLinkNisw)
	}
	if len(vLinkIew) > 0 {
		params.SetLinkIew(vLinkIew)
	}
	if len(vLinkNiew) > 0 {
		params.SetLinkNiew(vLinkNiew)
	}
	if vLinkEmpty != nil {
		params.SetLinkEmpty(vLinkEmpty)
	}
	if len(vLinkRegex) > 0 {
		params.SetLinkRegex(vLinkRegex)
	}
	if len(vLinkIregex) > 0 {
		params.SetLinkIregex(vLinkIregex)
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
	if len(vPhone) > 0 {
		params.SetPhone(vPhone)
	}
	if len(vPhonen) > 0 {
		params.SetPhonen(vPhonen)
	}
	if len(vPhoneIc) > 0 {
		params.SetPhoneIc(vPhoneIc)
	}
	if len(vPhoneNic) > 0 {
		params.SetPhoneNic(vPhoneNic)
	}
	if len(vPhoneIe) > 0 {
		params.SetPhoneIe(vPhoneIe)
	}
	if len(vPhoneNie) > 0 {
		params.SetPhoneNie(vPhoneNie)
	}
	if len(vPhoneIsw) > 0 {
		params.SetPhoneIsw(vPhoneIsw)
	}
	if len(vPhoneNisw) > 0 {
		params.SetPhoneNisw(vPhoneNisw)
	}
	if len(vPhoneIew) > 0 {
		params.SetPhoneIew(vPhoneIew)
	}
	if len(vPhoneNiew) > 0 {
		params.SetPhoneNiew(vPhoneNiew)
	}
	if vPhoneEmpty != nil {
		params.SetPhoneEmpty(vPhoneEmpty)
	}
	if len(vPhoneRegex) > 0 {
		params.SetPhoneRegex(vPhoneRegex)
	}
	if len(vPhoneIregex) > 0 {
		params.SetPhoneIregex(vPhoneIregex)
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
	if len(vTitle) > 0 {
		params.SetTitle(vTitle)
	}
	if len(vTitlen) > 0 {
		params.SetTitlen(vTitlen)
	}
	if len(vTitleIc) > 0 {
		params.SetTitleIc(vTitleIc)
	}
	if len(vTitleNic) > 0 {
		params.SetTitleNic(vTitleNic)
	}
	if len(vTitleIe) > 0 {
		params.SetTitleIe(vTitleIe)
	}
	if len(vTitleNie) > 0 {
		params.SetTitleNie(vTitleNie)
	}
	if len(vTitleIsw) > 0 {
		params.SetTitleIsw(vTitleIsw)
	}
	if len(vTitleNisw) > 0 {
		params.SetTitleNisw(vTitleNisw)
	}
	if len(vTitleIew) > 0 {
		params.SetTitleIew(vTitleIew)
	}
	if len(vTitleNiew) > 0 {
		params.SetTitleNiew(vTitleNiew)
	}
	if vTitleEmpty != nil {
		params.SetTitleEmpty(vTitleEmpty)
	}
	if len(vTitleRegex) > 0 {
		params.SetTitleRegex(vTitleRegex)
	}
	if len(vTitleIregex) > 0 {
		params.SetTitleIregex(vTitleIregex)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Tenancy.TenancyContactsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_contacts", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ContactResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]contactResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model contactResourceModel
		resp.Diagnostics.Append(flattenContact(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: contactResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
