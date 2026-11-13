// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"name":                netboxapi.FilterStrings(params.SetName),
		"name__n":             netboxapi.FilterStrings(params.SetNamen),
		"name__ic":            netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":           netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":            netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":           netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":           netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":          netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":           netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":          netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":         netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":         netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":        netboxapi.FilterStrings(params.SetNameIregex),
		"owner_id":            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":         netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":            netboxapi.FilterStrings(params.SetTagAny),
		"address":             netboxapi.FilterStrings(params.SetAddress),
		"address__n":          netboxapi.FilterStrings(params.SetAddressn),
		"address__ic":         netboxapi.FilterStrings(params.SetAddressIc),
		"address__nic":        netboxapi.FilterStrings(params.SetAddressNic),
		"address__ie":         netboxapi.FilterStrings(params.SetAddressIe),
		"address__nie":        netboxapi.FilterStrings(params.SetAddressNie),
		"address__isw":        netboxapi.FilterStrings(params.SetAddressIsw),
		"address__nisw":       netboxapi.FilterStrings(params.SetAddressNisw),
		"address__iew":        netboxapi.FilterStrings(params.SetAddressIew),
		"address__niew":       netboxapi.FilterStrings(params.SetAddressNiew),
		"address__empty":      netboxapi.FilterBool(params.SetAddressEmpty),
		"address__regex":      netboxapi.FilterStrings(params.SetAddressRegex),
		"address__iregex":     netboxapi.FilterStrings(params.SetAddressIregex),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":         netboxapi.FilterStrings(params.SetDescription),
		"description__n":      netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":     netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":    netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":     netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":    netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":    netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":   netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":    netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":   netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":  netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":  netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex": netboxapi.FilterStrings(params.SetDescriptionIregex),
		"email":               netboxapi.FilterStrings(params.SetEmail),
		"email__n":            netboxapi.FilterStrings(params.SetEmailn),
		"email__ic":           netboxapi.FilterStrings(params.SetEmailIc),
		"email__nic":          netboxapi.FilterStrings(params.SetEmailNic),
		"email__ie":           netboxapi.FilterStrings(params.SetEmailIe),
		"email__nie":          netboxapi.FilterStrings(params.SetEmailNie),
		"email__isw":          netboxapi.FilterStrings(params.SetEmailIsw),
		"email__nisw":         netboxapi.FilterStrings(params.SetEmailNisw),
		"email__iew":          netboxapi.FilterStrings(params.SetEmailIew),
		"email__niew":         netboxapi.FilterStrings(params.SetEmailNiew),
		"email__empty":        netboxapi.FilterBool(params.SetEmailEmpty),
		"email__regex":        netboxapi.FilterStrings(params.SetEmailRegex),
		"email__iregex":       netboxapi.FilterStrings(params.SetEmailIregex),
		"group":               netboxapi.FilterStrings(params.SetGroup),
		"group__n":            netboxapi.FilterStrings(params.SetGroupn),
		"group_id":            netboxapi.FilterStrings(params.SetGroupID),
		"group_id__n":         netboxapi.FilterStrings(params.SetGroupIDn),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"link":                netboxapi.FilterStrings(params.SetLink),
		"link__n":             netboxapi.FilterStrings(params.SetLinkn),
		"link__ic":            netboxapi.FilterStrings(params.SetLinkIc),
		"link__nic":           netboxapi.FilterStrings(params.SetLinkNic),
		"link__ie":            netboxapi.FilterStrings(params.SetLinkIe),
		"link__nie":           netboxapi.FilterStrings(params.SetLinkNie),
		"link__isw":           netboxapi.FilterStrings(params.SetLinkIsw),
		"link__nisw":          netboxapi.FilterStrings(params.SetLinkNisw),
		"link__iew":           netboxapi.FilterStrings(params.SetLinkIew),
		"link__niew":          netboxapi.FilterStrings(params.SetLinkNiew),
		"link__empty":         netboxapi.FilterBool(params.SetLinkEmpty),
		"link__regex":         netboxapi.FilterStrings(params.SetLinkRegex),
		"link__iregex":        netboxapi.FilterStrings(params.SetLinkIregex),
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"phone":               netboxapi.FilterStrings(params.SetPhone),
		"phone__n":            netboxapi.FilterStrings(params.SetPhonen),
		"phone__ic":           netboxapi.FilterStrings(params.SetPhoneIc),
		"phone__nic":          netboxapi.FilterStrings(params.SetPhoneNic),
		"phone__ie":           netboxapi.FilterStrings(params.SetPhoneIe),
		"phone__nie":          netboxapi.FilterStrings(params.SetPhoneNie),
		"phone__isw":          netboxapi.FilterStrings(params.SetPhoneIsw),
		"phone__nisw":         netboxapi.FilterStrings(params.SetPhoneNisw),
		"phone__iew":          netboxapi.FilterStrings(params.SetPhoneIew),
		"phone__niew":         netboxapi.FilterStrings(params.SetPhoneNiew),
		"phone__empty":        netboxapi.FilterBool(params.SetPhoneEmpty),
		"phone__regex":        netboxapi.FilterStrings(params.SetPhoneRegex),
		"phone__iregex":       netboxapi.FilterStrings(params.SetPhoneIregex),
		"q":                   netboxapi.FilterString(params.SetQ),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
		"title":               netboxapi.FilterStrings(params.SetTitle),
		"title__n":            netboxapi.FilterStrings(params.SetTitlen),
		"title__ic":           netboxapi.FilterStrings(params.SetTitleIc),
		"title__nic":          netboxapi.FilterStrings(params.SetTitleNic),
		"title__ie":           netboxapi.FilterStrings(params.SetTitleIe),
		"title__nie":          netboxapi.FilterStrings(params.SetTitleNie),
		"title__isw":          netboxapi.FilterStrings(params.SetTitleIsw),
		"title__nisw":         netboxapi.FilterStrings(params.SetTitleNisw),
		"title__iew":          netboxapi.FilterStrings(params.SetTitleIew),
		"title__niew":         netboxapi.FilterStrings(params.SetTitleNiew),
		"title__empty":        netboxapi.FilterBool(params.SetTitleEmpty),
		"title__regex":        netboxapi.FilterStrings(params.SetTitleRegex),
		"title__iregex":       netboxapi.FilterStrings(params.SetTitleIregex),
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
