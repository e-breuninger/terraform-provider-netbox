// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vPayloadURL []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vCaFilePath []string
	var vCaFilePathn []string
	var vCaFilePathIc []string
	var vCaFilePathNic []string
	var vCaFilePathIe []string
	var vCaFilePathNie []string
	var vCaFilePathIsw []string
	var vCaFilePathNisw []string
	var vCaFilePathIew []string
	var vCaFilePathNiew []string
	var vCaFilePathEmpty *bool
	var vCaFilePathRegex []string
	var vCaFilePathIregex []string
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
	var vHTTPContentType []string
	var vHTTPContentTypen []string
	var vHTTPContentTypeIc []string
	var vHTTPContentTypeNic []string
	var vHTTPContentTypeIe []string
	var vHTTPContentTypeNie []string
	var vHTTPContentTypeIsw []string
	var vHTTPContentTypeNisw []string
	var vHTTPContentTypeIew []string
	var vHTTPContentTypeNiew []string
	var vHTTPContentTypeEmpty *bool
	var vHTTPContentTypeRegex []string
	var vHTTPContentTypeIregex []string
	var vHTTPMethod []string
	var vHTTPMethodn []string
	var vHTTPMethodIc []string
	var vHTTPMethodNic []string
	var vHTTPMethodIe []string
	var vHTTPMethodNie []string
	var vHTTPMethodIsw []string
	var vHTTPMethodNisw []string
	var vHTTPMethodIew []string
	var vHTTPMethodNiew []string
	var vHTTPMethodEmpty *bool
	var vHTTPMethodRegex []string
	var vHTTPMethodIregex []string
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
	var vSecret []string
	var vSecretn []string
	var vSecretIc []string
	var vSecretNic []string
	var vSecretIe []string
	var vSecretNie []string
	var vSecretIsw []string
	var vSecretNisw []string
	var vSecretIew []string
	var vSecretNiew []string
	var vSecretEmpty *bool
	var vSecretRegex []string
	var vSecretIregex []string
	var vSslVerification *bool
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
		case "payload_url":
			v := value
			vPayloadURL = append(vPayloadURL, v)
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
		case "ca_file_path":
			v := value
			vCaFilePath = append(vCaFilePath, v)
		case "ca_file_path__n":
			v := value
			vCaFilePathn = append(vCaFilePathn, v)
		case "ca_file_path__ic":
			v := value
			vCaFilePathIc = append(vCaFilePathIc, v)
		case "ca_file_path__nic":
			v := value
			vCaFilePathNic = append(vCaFilePathNic, v)
		case "ca_file_path__ie":
			v := value
			vCaFilePathIe = append(vCaFilePathIe, v)
		case "ca_file_path__nie":
			v := value
			vCaFilePathNie = append(vCaFilePathNie, v)
		case "ca_file_path__isw":
			v := value
			vCaFilePathIsw = append(vCaFilePathIsw, v)
		case "ca_file_path__nisw":
			v := value
			vCaFilePathNisw = append(vCaFilePathNisw, v)
		case "ca_file_path__iew":
			v := value
			vCaFilePathIew = append(vCaFilePathIew, v)
		case "ca_file_path__niew":
			v := value
			vCaFilePathNiew = append(vCaFilePathNiew, v)
		case "ca_file_path__empty":
			if vCaFilePathEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ca_file_path__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ca_file_path__empty' takes a boolean, got %q.", value))
				return
			}
			vCaFilePathEmpty = &v
		case "ca_file_path__regex":
			v := value
			vCaFilePathRegex = append(vCaFilePathRegex, v)
		case "ca_file_path__iregex":
			v := value
			vCaFilePathIregex = append(vCaFilePathIregex, v)
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
		case "http_content_type":
			v := value
			vHTTPContentType = append(vHTTPContentType, v)
		case "http_content_type__n":
			v := value
			vHTTPContentTypen = append(vHTTPContentTypen, v)
		case "http_content_type__ic":
			v := value
			vHTTPContentTypeIc = append(vHTTPContentTypeIc, v)
		case "http_content_type__nic":
			v := value
			vHTTPContentTypeNic = append(vHTTPContentTypeNic, v)
		case "http_content_type__ie":
			v := value
			vHTTPContentTypeIe = append(vHTTPContentTypeIe, v)
		case "http_content_type__nie":
			v := value
			vHTTPContentTypeNie = append(vHTTPContentTypeNie, v)
		case "http_content_type__isw":
			v := value
			vHTTPContentTypeIsw = append(vHTTPContentTypeIsw, v)
		case "http_content_type__nisw":
			v := value
			vHTTPContentTypeNisw = append(vHTTPContentTypeNisw, v)
		case "http_content_type__iew":
			v := value
			vHTTPContentTypeIew = append(vHTTPContentTypeIew, v)
		case "http_content_type__niew":
			v := value
			vHTTPContentTypeNiew = append(vHTTPContentTypeNiew, v)
		case "http_content_type__empty":
			if vHTTPContentTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'http_content_type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'http_content_type__empty' takes a boolean, got %q.", value))
				return
			}
			vHTTPContentTypeEmpty = &v
		case "http_content_type__regex":
			v := value
			vHTTPContentTypeRegex = append(vHTTPContentTypeRegex, v)
		case "http_content_type__iregex":
			v := value
			vHTTPContentTypeIregex = append(vHTTPContentTypeIregex, v)
		case "http_method":
			v := value
			vHTTPMethod = append(vHTTPMethod, v)
		case "http_method__n":
			v := value
			vHTTPMethodn = append(vHTTPMethodn, v)
		case "http_method__ic":
			v := value
			vHTTPMethodIc = append(vHTTPMethodIc, v)
		case "http_method__nic":
			v := value
			vHTTPMethodNic = append(vHTTPMethodNic, v)
		case "http_method__ie":
			v := value
			vHTTPMethodIe = append(vHTTPMethodIe, v)
		case "http_method__nie":
			v := value
			vHTTPMethodNie = append(vHTTPMethodNie, v)
		case "http_method__isw":
			v := value
			vHTTPMethodIsw = append(vHTTPMethodIsw, v)
		case "http_method__nisw":
			v := value
			vHTTPMethodNisw = append(vHTTPMethodNisw, v)
		case "http_method__iew":
			v := value
			vHTTPMethodIew = append(vHTTPMethodIew, v)
		case "http_method__niew":
			v := value
			vHTTPMethodNiew = append(vHTTPMethodNiew, v)
		case "http_method__empty":
			if vHTTPMethodEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'http_method__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'http_method__empty' takes a boolean, got %q.", value))
				return
			}
			vHTTPMethodEmpty = &v
		case "http_method__regex":
			v := value
			vHTTPMethodRegex = append(vHTTPMethodRegex, v)
		case "http_method__iregex":
			v := value
			vHTTPMethodIregex = append(vHTTPMethodIregex, v)
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
		case "secret":
			v := value
			vSecret = append(vSecret, v)
		case "secret__n":
			v := value
			vSecretn = append(vSecretn, v)
		case "secret__ic":
			v := value
			vSecretIc = append(vSecretIc, v)
		case "secret__nic":
			v := value
			vSecretNic = append(vSecretNic, v)
		case "secret__ie":
			v := value
			vSecretIe = append(vSecretIe, v)
		case "secret__nie":
			v := value
			vSecretNie = append(vSecretNie, v)
		case "secret__isw":
			v := value
			vSecretIsw = append(vSecretIsw, v)
		case "secret__nisw":
			v := value
			vSecretNisw = append(vSecretNisw, v)
		case "secret__iew":
			v := value
			vSecretIew = append(vSecretIew, v)
		case "secret__niew":
			v := value
			vSecretNiew = append(vSecretNiew, v)
		case "secret__empty":
			if vSecretEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'secret__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'secret__empty' takes a boolean, got %q.", value))
				return
			}
			vSecretEmpty = &v
		case "secret__regex":
			v := value
			vSecretRegex = append(vSecretRegex, v)
		case "secret__iregex":
			v := value
			vSecretIregex = append(vSecretIregex, v)
		case "ssl_verification":
			if vSslVerification != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'ssl_verification' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'ssl_verification' takes a boolean, got %q.", value))
				return
			}
			vSslVerification = &v
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
	if len(vPayloadURL) > 0 {
		params.SetPayloadURL(vPayloadURL)
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
	if len(vCaFilePath) > 0 {
		params.SetCaFilePath(vCaFilePath)
	}
	if len(vCaFilePathn) > 0 {
		params.SetCaFilePathn(vCaFilePathn)
	}
	if len(vCaFilePathIc) > 0 {
		params.SetCaFilePathIc(vCaFilePathIc)
	}
	if len(vCaFilePathNic) > 0 {
		params.SetCaFilePathNic(vCaFilePathNic)
	}
	if len(vCaFilePathIe) > 0 {
		params.SetCaFilePathIe(vCaFilePathIe)
	}
	if len(vCaFilePathNie) > 0 {
		params.SetCaFilePathNie(vCaFilePathNie)
	}
	if len(vCaFilePathIsw) > 0 {
		params.SetCaFilePathIsw(vCaFilePathIsw)
	}
	if len(vCaFilePathNisw) > 0 {
		params.SetCaFilePathNisw(vCaFilePathNisw)
	}
	if len(vCaFilePathIew) > 0 {
		params.SetCaFilePathIew(vCaFilePathIew)
	}
	if len(vCaFilePathNiew) > 0 {
		params.SetCaFilePathNiew(vCaFilePathNiew)
	}
	if vCaFilePathEmpty != nil {
		params.SetCaFilePathEmpty(vCaFilePathEmpty)
	}
	if len(vCaFilePathRegex) > 0 {
		params.SetCaFilePathRegex(vCaFilePathRegex)
	}
	if len(vCaFilePathIregex) > 0 {
		params.SetCaFilePathIregex(vCaFilePathIregex)
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
	if len(vHTTPContentType) > 0 {
		params.SetHTTPContentType(vHTTPContentType)
	}
	if len(vHTTPContentTypen) > 0 {
		params.SetHTTPContentTypen(vHTTPContentTypen)
	}
	if len(vHTTPContentTypeIc) > 0 {
		params.SetHTTPContentTypeIc(vHTTPContentTypeIc)
	}
	if len(vHTTPContentTypeNic) > 0 {
		params.SetHTTPContentTypeNic(vHTTPContentTypeNic)
	}
	if len(vHTTPContentTypeIe) > 0 {
		params.SetHTTPContentTypeIe(vHTTPContentTypeIe)
	}
	if len(vHTTPContentTypeNie) > 0 {
		params.SetHTTPContentTypeNie(vHTTPContentTypeNie)
	}
	if len(vHTTPContentTypeIsw) > 0 {
		params.SetHTTPContentTypeIsw(vHTTPContentTypeIsw)
	}
	if len(vHTTPContentTypeNisw) > 0 {
		params.SetHTTPContentTypeNisw(vHTTPContentTypeNisw)
	}
	if len(vHTTPContentTypeIew) > 0 {
		params.SetHTTPContentTypeIew(vHTTPContentTypeIew)
	}
	if len(vHTTPContentTypeNiew) > 0 {
		params.SetHTTPContentTypeNiew(vHTTPContentTypeNiew)
	}
	if vHTTPContentTypeEmpty != nil {
		params.SetHTTPContentTypeEmpty(vHTTPContentTypeEmpty)
	}
	if len(vHTTPContentTypeRegex) > 0 {
		params.SetHTTPContentTypeRegex(vHTTPContentTypeRegex)
	}
	if len(vHTTPContentTypeIregex) > 0 {
		params.SetHTTPContentTypeIregex(vHTTPContentTypeIregex)
	}
	if len(vHTTPMethod) > 0 {
		params.SetHTTPMethod(vHTTPMethod)
	}
	if len(vHTTPMethodn) > 0 {
		params.SetHTTPMethodn(vHTTPMethodn)
	}
	if len(vHTTPMethodIc) > 0 {
		params.SetHTTPMethodIc(vHTTPMethodIc)
	}
	if len(vHTTPMethodNic) > 0 {
		params.SetHTTPMethodNic(vHTTPMethodNic)
	}
	if len(vHTTPMethodIe) > 0 {
		params.SetHTTPMethodIe(vHTTPMethodIe)
	}
	if len(vHTTPMethodNie) > 0 {
		params.SetHTTPMethodNie(vHTTPMethodNie)
	}
	if len(vHTTPMethodIsw) > 0 {
		params.SetHTTPMethodIsw(vHTTPMethodIsw)
	}
	if len(vHTTPMethodNisw) > 0 {
		params.SetHTTPMethodNisw(vHTTPMethodNisw)
	}
	if len(vHTTPMethodIew) > 0 {
		params.SetHTTPMethodIew(vHTTPMethodIew)
	}
	if len(vHTTPMethodNiew) > 0 {
		params.SetHTTPMethodNiew(vHTTPMethodNiew)
	}
	if vHTTPMethodEmpty != nil {
		params.SetHTTPMethodEmpty(vHTTPMethodEmpty)
	}
	if len(vHTTPMethodRegex) > 0 {
		params.SetHTTPMethodRegex(vHTTPMethodRegex)
	}
	if len(vHTTPMethodIregex) > 0 {
		params.SetHTTPMethodIregex(vHTTPMethodIregex)
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
	if len(vSecret) > 0 {
		params.SetSecret(vSecret)
	}
	if len(vSecretn) > 0 {
		params.SetSecretn(vSecretn)
	}
	if len(vSecretIc) > 0 {
		params.SetSecretIc(vSecretIc)
	}
	if len(vSecretNic) > 0 {
		params.SetSecretNic(vSecretNic)
	}
	if len(vSecretIe) > 0 {
		params.SetSecretIe(vSecretIe)
	}
	if len(vSecretNie) > 0 {
		params.SetSecretNie(vSecretNie)
	}
	if len(vSecretIsw) > 0 {
		params.SetSecretIsw(vSecretIsw)
	}
	if len(vSecretNisw) > 0 {
		params.SetSecretNisw(vSecretNisw)
	}
	if len(vSecretIew) > 0 {
		params.SetSecretIew(vSecretIew)
	}
	if len(vSecretNiew) > 0 {
		params.SetSecretNiew(vSecretNiew)
	}
	if vSecretEmpty != nil {
		params.SetSecretEmpty(vSecretEmpty)
	}
	if len(vSecretRegex) > 0 {
		params.SetSecretRegex(vSecretRegex)
	}
	if len(vSecretIregex) > 0 {
		params.SetSecretIregex(vSecretIregex)
	}
	if vSslVerification != nil {
		params.SetSslVerification(vSslVerification)
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
