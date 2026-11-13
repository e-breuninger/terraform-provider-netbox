// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*tokensDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*tokensDataSource)(nil)
)

// NewTokensDataSource returns a new tokens data source, which
// lists token objects matching its filters.
func NewTokensDataSource() datasource.DataSource {
	return &tokensDataSource{}
}

type tokensDataSource struct {
	client *netboxapi.Client
}

// tokensDataSourceModel is the data source model: the filters, the limit and the matching
// tokens.
type tokensDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"tokens"`
}

// tokenResourceAttrTypes is the attribute type map of tokenResourceModel.
func tokenResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"user_id":       types.Int64Type,
		"key":           types.StringType,
		"token":         types.StringType,
		"enabled":       types.BoolType,
		"version":       types.Int64Type,
		"pepper_id":     types.Int64Type,
		"description":   types.StringType,
		"write_enabled": types.BoolType,
		"expires":       types.StringType,
		"last_used":     types.StringType,
		"created":       types.StringType,
		"url":           types.StringType,
	}
}

func (d *tokensDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tokens"
}

func (d *tokensDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:Lists token objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__gte, created__lte, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, enabled, expires, expires__gte, expires__lte, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, key, key__empty, key__ic, key__ie, key__iew, key__iregex, key__isw, key__n, key__nic, key__nie, key__niew, key__nisw, key__regex, last_used, last_used__gte, last_used__lte, pepper_id, pepper_id__empty, pepper_id__gt, pepper_id__gte, pepper_id__lt, pepper_id__lte, pepper_id__n, q, user, user__n, user_id, user_id__n, version, version__ic, version__ie, version__iew, version__iregex, version__isw, version__n, version__nic, version__nie, version__niew, version__nisw, version__regex, write_enabled. Repeating a name sends that parameter once per value.",
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
			"tokens": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching tokens.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"user_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the user the token belongs to.",
						},
						"key": schema.StringAttribute{
							Computed:    true,
							Description: "Identification key NetBox generated for the token; not the secret.",
						},
						"token": schema.StringAttribute{
							Computed:    true,
							Sensitive:   true,
							Description: "The secret. NetBox returns it in the create response only, so an imported token has none in state and a secret changed outside Terraform goes unnoticed. Changing it replaces the token.",
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the token can authenticate; disabling keeps the token but rejects its use.",
						},
						"version": schema.Int64Attribute{
							Computed:    true,
							Description: "Token format version (1 legacy plaintext, 2 hashed). Defaults to 2. Changing it replaces the token. One of: 1, 2.",
						},
						"pepper_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the pepper the secret was hashed with.",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"write_enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the token permits write operations.",
						},
						"expires": schema.StringAttribute{
							Computed:    true,
							Description: "Expiry as an RFC 3339 timestamp with milliseconds (2030-01-01T00:00:00.000Z); never expires when unset.",
						},
						"last_used": schema.StringAttribute{
							Computed:    true,
							Description: "When the token was last used.",
						},
						"created": schema.StringAttribute{
							Computed: true,
						},
						"url": schema.StringAttribute{
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func (d *tokensDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *tokensDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data tokensDataSourceModel

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
	var responseDTOs []*netboxapi.TokenResponseDTO

	params := users.NewUsersTokensListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"user_id":             netboxapi.FilterInt64s(params.SetUserID),
		"user_id__n":          netboxapi.FilterInt64s(params.SetUserIDn),
		"created":             netboxapi.FilterString(params.SetCreated),
		"created__gte":        netboxapi.FilterString(params.SetCreatedGte),
		"created__lte":        netboxapi.FilterString(params.SetCreatedLte),
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
		"enabled":             netboxapi.FilterBool(params.SetEnabled),
		"expires":             netboxapi.FilterString(params.SetExpires),
		"expires__gte":        netboxapi.FilterString(params.SetExpiresGte),
		"expires__lte":        netboxapi.FilterString(params.SetExpiresLte),
		"key":                 netboxapi.FilterStrings(params.SetKey),
		"key__n":              netboxapi.FilterStrings(params.SetKeyn),
		"key__ic":             netboxapi.FilterStrings(params.SetKeyIc),
		"key__nic":            netboxapi.FilterStrings(params.SetKeyNic),
		"key__ie":             netboxapi.FilterStrings(params.SetKeyIe),
		"key__nie":            netboxapi.FilterStrings(params.SetKeyNie),
		"key__isw":            netboxapi.FilterStrings(params.SetKeyIsw),
		"key__nisw":           netboxapi.FilterStrings(params.SetKeyNisw),
		"key__iew":            netboxapi.FilterStrings(params.SetKeyIew),
		"key__niew":           netboxapi.FilterStrings(params.SetKeyNiew),
		"key__empty":          netboxapi.FilterBool(params.SetKeyEmpty),
		"key__regex":          netboxapi.FilterStrings(params.SetKeyRegex),
		"key__iregex":         netboxapi.FilterStrings(params.SetKeyIregex),
		"last_used":           netboxapi.FilterString(params.SetLastUsed),
		"last_used__gte":      netboxapi.FilterString(params.SetLastUsedGte),
		"last_used__lte":      netboxapi.FilterString(params.SetLastUsedLte),
		"pepper_id":           netboxapi.FilterInt64s(params.SetPepperID),
		"pepper_id__n":        netboxapi.FilterInt64s(params.SetPepperIDn),
		"pepper_id__lt":       netboxapi.FilterInt64s(params.SetPepperIDLt),
		"pepper_id__lte":      netboxapi.FilterInt64s(params.SetPepperIDLte),
		"pepper_id__gt":       netboxapi.FilterInt64s(params.SetPepperIDGt),
		"pepper_id__gte":      netboxapi.FilterInt64s(params.SetPepperIDGte),
		"pepper_id__empty":    netboxapi.FilterBool(params.SetPepperIDEmpty),
		"q":                   netboxapi.FilterString(params.SetQ),
		"user":                netboxapi.FilterStrings(params.SetUser),
		"user__n":             netboxapi.FilterStrings(params.SetUsern),
		"version":             netboxapi.FilterInt64(params.SetVersion),
		"version__n":          netboxapi.FilterInt64(params.SetVersionn),
		"version__ic":         netboxapi.FilterInt64s(params.SetVersionIc),
		"version__nic":        netboxapi.FilterInt64s(params.SetVersionNic),
		"version__ie":         netboxapi.FilterInt64s(params.SetVersionIe),
		"version__nie":        netboxapi.FilterInt64s(params.SetVersionNie),
		"version__isw":        netboxapi.FilterInt64s(params.SetVersionIsw),
		"version__nisw":       netboxapi.FilterInt64s(params.SetVersionNisw),
		"version__iew":        netboxapi.FilterInt64s(params.SetVersionIew),
		"version__niew":       netboxapi.FilterInt64s(params.SetVersionNiew),
		"version__regex":      netboxapi.FilterInt64s(params.SetVersionRegex),
		"version__iregex":     netboxapi.FilterInt64s(params.SetVersionIregex),
		"write_enabled":       netboxapi.FilterBool(params.SetWriteEnabled),
	}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
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
		res, err := d.client.Users.UsersTokensListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_tokens", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.TokenResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]tokenResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model tokenResourceModel
		resp.Diagnostics.Append(flattenToken(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: tokenResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
