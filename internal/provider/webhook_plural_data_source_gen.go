// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*webhooksDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*webhooksDataSource)(nil)
)

// NewWebhooksDataSource returns a new webhooks data source, which
// lists webhook objects matching its filters.
func NewWebhooksDataSource() datasource.DataSource {
	return &webhooksDataSource{}
}

type webhooksDataSource struct {
	client *netboxapi.Client
}

// webhooksDataSourceModel is the data source model: the filters, the limit and the matching
// webhooks.
type webhooksDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"webhooks"`
}

// webhookResourceAttrTypes is the attribute type map of webhookResourceModel.
func webhookResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                 types.Int64Type,
		"name":               types.StringType,
		"payload_url":        types.StringType,
		"http_method":        types.StringType,
		"http_content_type":  types.StringType,
		"additional_headers": types.StringType,
		"body_template":      types.StringType,
		"secret":             types.StringType,
		"ssl_verification":   types.BoolType,
		"ca_file_path":       types.StringType,
		"description":        types.StringType,
		"owner_id":           types.Int64Type,
		"created":            types.StringType,
		"last_updated":       types.StringType,
		"url":                types.StringType,
		"tags":               types.SetType{ElemType: types.StringType},
		"tags_all":           types.SetType{ElemType: types.StringType},
		"custom_fields":      types.MapType{ElemType: types.StringType},
	}
}

func (d *webhooksDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_webhooks"
}

func (d *webhooksDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:Lists webhook objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: ca_file_path, ca_file_path__empty, ca_file_path__ic, ca_file_path__ie, ca_file_path__iew, ca_file_path__iregex, ca_file_path__isw, ca_file_path__n, ca_file_path__nic, ca_file_path__nie, ca_file_path__niew, ca_file_path__nisw, ca_file_path__regex, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, http_content_type, http_content_type__empty, http_content_type__ic, http_content_type__ie, http_content_type__iew, http_content_type__iregex, http_content_type__isw, http_content_type__n, http_content_type__nic, http_content_type__nie, http_content_type__niew, http_content_type__nisw, http_content_type__regex, http_method, http_method__empty, http_method__ic, http_method__ie, http_method__iew, http_method__iregex, http_method__isw, http_method__n, http_method__nic, http_method__nie, http_method__niew, http_method__nisw, http_method__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, payload_url, q, secret, secret__empty, secret__ic, secret__ie, secret__iew, secret__iregex, secret__isw, secret__n, secret__nic, secret__nie, secret__niew, secret__nisw, secret__regex, ssl_verification, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"webhooks": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching webhooks.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"payload_url": schema.StringAttribute{
							Computed:    true,
							Description: "URL the request is sent to; may be a Jinja2 template.",
						},
						"http_method": schema.StringAttribute{
							Computed:    true,
							Description: "HTTP method. One of: GET, POST, PUT, PATCH, DELETE.",
						},
						"http_content_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content-Type header.",
						},
						"additional_headers": schema.StringAttribute{
							Computed:    true,
							Description: "Extra request headers, one Name: Value per line; may be a Jinja2 template.",
						},
						"body_template": schema.StringAttribute{
							Computed:    true,
							Description: "Jinja2 template for the request body.",
						},
						"secret": schema.StringAttribute{
							Computed:    true,
							Sensitive:   true,
							Description: "Secret used to sign the request in the X-Hook-Signature header.",
						},
						"ssl_verification": schema.BoolAttribute{
							Computed:    true,
							Description: "Verify the endpoint's TLS certificate.",
						},
						"ca_file_path": schema.StringAttribute{
							Computed:    true,
							Description: "Path on the NetBox server to a CA certificate file to verify against. NetBox rejects it while ssl_verification is false.",
						},
						"description": schema.StringAttribute{
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

func (d *webhooksDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *webhooksDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data webhooksDataSourceModel

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
	var responseDTOs []*netboxapi.WebhookResponseDTO

	params := extras.NewExtrasWebhooksListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                        netboxapi.FilterInt64s(params.SetID),
		"id__n":                     netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                    netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                   netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                    netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                   netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                 netboxapi.FilterBool(params.SetIDEmpty),
		"name":                      netboxapi.FilterStrings(params.SetName),
		"name__n":                   netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                  netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                 netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                  netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                 netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                 netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                 netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":               netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":               netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":              netboxapi.FilterStrings(params.SetNameIregex),
		"payload_url":               netboxapi.FilterStrings(params.SetPayloadURL),
		"owner_id":                  netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":               netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                       netboxapi.FilterStrings(params.SetTag),
		"tag__n":                    netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                  netboxapi.FilterStrings(params.SetTagAny),
		"ca_file_path":              netboxapi.FilterStrings(params.SetCaFilePath),
		"ca_file_path__n":           netboxapi.FilterStrings(params.SetCaFilePathn),
		"ca_file_path__ic":          netboxapi.FilterStrings(params.SetCaFilePathIc),
		"ca_file_path__nic":         netboxapi.FilterStrings(params.SetCaFilePathNic),
		"ca_file_path__ie":          netboxapi.FilterStrings(params.SetCaFilePathIe),
		"ca_file_path__nie":         netboxapi.FilterStrings(params.SetCaFilePathNie),
		"ca_file_path__isw":         netboxapi.FilterStrings(params.SetCaFilePathIsw),
		"ca_file_path__nisw":        netboxapi.FilterStrings(params.SetCaFilePathNisw),
		"ca_file_path__iew":         netboxapi.FilterStrings(params.SetCaFilePathIew),
		"ca_file_path__niew":        netboxapi.FilterStrings(params.SetCaFilePathNiew),
		"ca_file_path__empty":       netboxapi.FilterBool(params.SetCaFilePathEmpty),
		"ca_file_path__regex":       netboxapi.FilterStrings(params.SetCaFilePathRegex),
		"ca_file_path__iregex":      netboxapi.FilterStrings(params.SetCaFilePathIregex),
		"created":                   netboxapi.FilterStrings(params.SetCreated),
		"created__n":                netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":               netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":              netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":               netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":              netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":            netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":               netboxapi.FilterStrings(params.SetDescription),
		"description__n":            netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":           netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":          netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":           netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":          netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":          netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":         netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":          netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":         netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":        netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":        netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":       netboxapi.FilterStrings(params.SetDescriptionIregex),
		"http_content_type":         netboxapi.FilterStrings(params.SetHTTPContentType),
		"http_content_type__n":      netboxapi.FilterStrings(params.SetHTTPContentTypen),
		"http_content_type__ic":     netboxapi.FilterStrings(params.SetHTTPContentTypeIc),
		"http_content_type__nic":    netboxapi.FilterStrings(params.SetHTTPContentTypeNic),
		"http_content_type__ie":     netboxapi.FilterStrings(params.SetHTTPContentTypeIe),
		"http_content_type__nie":    netboxapi.FilterStrings(params.SetHTTPContentTypeNie),
		"http_content_type__isw":    netboxapi.FilterStrings(params.SetHTTPContentTypeIsw),
		"http_content_type__nisw":   netboxapi.FilterStrings(params.SetHTTPContentTypeNisw),
		"http_content_type__iew":    netboxapi.FilterStrings(params.SetHTTPContentTypeIew),
		"http_content_type__niew":   netboxapi.FilterStrings(params.SetHTTPContentTypeNiew),
		"http_content_type__empty":  netboxapi.FilterBool(params.SetHTTPContentTypeEmpty),
		"http_content_type__regex":  netboxapi.FilterStrings(params.SetHTTPContentTypeRegex),
		"http_content_type__iregex": netboxapi.FilterStrings(params.SetHTTPContentTypeIregex),
		"http_method":               netboxapi.FilterStrings(params.SetHTTPMethod),
		"http_method__n":            netboxapi.FilterStrings(params.SetHTTPMethodn),
		"http_method__ic":           netboxapi.FilterStrings(params.SetHTTPMethodIc),
		"http_method__nic":          netboxapi.FilterStrings(params.SetHTTPMethodNic),
		"http_method__ie":           netboxapi.FilterStrings(params.SetHTTPMethodIe),
		"http_method__nie":          netboxapi.FilterStrings(params.SetHTTPMethodNie),
		"http_method__isw":          netboxapi.FilterStrings(params.SetHTTPMethodIsw),
		"http_method__nisw":         netboxapi.FilterStrings(params.SetHTTPMethodNisw),
		"http_method__iew":          netboxapi.FilterStrings(params.SetHTTPMethodIew),
		"http_method__niew":         netboxapi.FilterStrings(params.SetHTTPMethodNiew),
		"http_method__empty":        netboxapi.FilterBool(params.SetHTTPMethodEmpty),
		"http_method__regex":        netboxapi.FilterStrings(params.SetHTTPMethodRegex),
		"http_method__iregex":       netboxapi.FilterStrings(params.SetHTTPMethodIregex),
		"last_updated":              netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":           netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":          netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":         netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":          netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":         netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":       netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                     netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                  netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":               netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":            netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":            netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":         netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                         netboxapi.FilterString(params.SetQ),
		"secret":                    netboxapi.FilterStrings(params.SetSecret),
		"secret__n":                 netboxapi.FilterStrings(params.SetSecretn),
		"secret__ic":                netboxapi.FilterStrings(params.SetSecretIc),
		"secret__nic":               netboxapi.FilterStrings(params.SetSecretNic),
		"secret__ie":                netboxapi.FilterStrings(params.SetSecretIe),
		"secret__nie":               netboxapi.FilterStrings(params.SetSecretNie),
		"secret__isw":               netboxapi.FilterStrings(params.SetSecretIsw),
		"secret__nisw":              netboxapi.FilterStrings(params.SetSecretNisw),
		"secret__iew":               netboxapi.FilterStrings(params.SetSecretIew),
		"secret__niew":              netboxapi.FilterStrings(params.SetSecretNiew),
		"secret__empty":             netboxapi.FilterBool(params.SetSecretEmpty),
		"secret__regex":             netboxapi.FilterStrings(params.SetSecretRegex),
		"secret__iregex":            netboxapi.FilterStrings(params.SetSecretIregex),
		"ssl_verification":          netboxapi.FilterBool(params.SetSslVerification),
		"tag_id":                    netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                 netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":               netboxapi.FilterInt64s(params.SetTagIDAny),
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
		res, err := d.client.Extras.ExtrasWebhooksListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_webhooks", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.WebhookResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]webhookResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model webhookResourceModel
		resp.Diagnostics.Append(flattenWebhook(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: webhookResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
