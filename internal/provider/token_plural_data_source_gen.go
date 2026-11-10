// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vUserID []int64
	var vUserIDn []int64
	var vCreated *string
	var vCreatedGte *string
	var vCreatedLte *string
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
	var vEnabled *bool
	var vExpires *string
	var vExpiresGte *string
	var vExpiresLte *string
	var vKey []string
	var vKeyn []string
	var vKeyIc []string
	var vKeyNic []string
	var vKeyIe []string
	var vKeyNie []string
	var vKeyIsw []string
	var vKeyNisw []string
	var vKeyIew []string
	var vKeyNiew []string
	var vKeyEmpty *bool
	var vKeyRegex []string
	var vKeyIregex []string
	var vLastUsed *string
	var vLastUsedGte *string
	var vLastUsedLte *string
	var vPepperID []int64
	var vPepperIDn []int64
	var vPepperIDLt []int64
	var vPepperIDLte []int64
	var vPepperIDGt []int64
	var vPepperIDGte []int64
	var vPepperIDEmpty *bool
	var vQ *string
	var vUser []string
	var vUsern []string
	var vVersion *int64
	var vVersionn *int64
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
	var vWriteEnabled *bool
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
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
		case "user_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'user_id' takes an integer, got %q.", value))
				return
			}
			vUserID = append(vUserID, v)
		case "user_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'user_id__n' takes an integer, got %q.", value))
				return
			}
			vUserIDn = append(vUserIDn, v)
		case "created":
			if vCreated != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'created' takes a single value.")
				return
			}
			v := value
			vCreated = &v
		case "created__gte":
			if vCreatedGte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'created__gte' takes a single value.")
				return
			}
			v := value
			vCreatedGte = &v
		case "created__lte":
			if vCreatedLte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'created__lte' takes a single value.")
				return
			}
			v := value
			vCreatedLte = &v
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
		case "enabled":
			if vEnabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'enabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'enabled' takes a boolean, got %q.", value))
				return
			}
			vEnabled = &v
		case "expires":
			if vExpires != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'expires' takes a single value.")
				return
			}
			v := value
			vExpires = &v
		case "expires__gte":
			if vExpiresGte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'expires__gte' takes a single value.")
				return
			}
			v := value
			vExpiresGte = &v
		case "expires__lte":
			if vExpiresLte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'expires__lte' takes a single value.")
				return
			}
			v := value
			vExpiresLte = &v
		case "key":
			v := value
			vKey = append(vKey, v)
		case "key__n":
			v := value
			vKeyn = append(vKeyn, v)
		case "key__ic":
			v := value
			vKeyIc = append(vKeyIc, v)
		case "key__nic":
			v := value
			vKeyNic = append(vKeyNic, v)
		case "key__ie":
			v := value
			vKeyIe = append(vKeyIe, v)
		case "key__nie":
			v := value
			vKeyNie = append(vKeyNie, v)
		case "key__isw":
			v := value
			vKeyIsw = append(vKeyIsw, v)
		case "key__nisw":
			v := value
			vKeyNisw = append(vKeyNisw, v)
		case "key__iew":
			v := value
			vKeyIew = append(vKeyIew, v)
		case "key__niew":
			v := value
			vKeyNiew = append(vKeyNiew, v)
		case "key__empty":
			if vKeyEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'key__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'key__empty' takes a boolean, got %q.", value))
				return
			}
			vKeyEmpty = &v
		case "key__regex":
			v := value
			vKeyRegex = append(vKeyRegex, v)
		case "key__iregex":
			v := value
			vKeyIregex = append(vKeyIregex, v)
		case "last_used":
			if vLastUsed != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'last_used' takes a single value.")
				return
			}
			v := value
			vLastUsed = &v
		case "last_used__gte":
			if vLastUsedGte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'last_used__gte' takes a single value.")
				return
			}
			v := value
			vLastUsedGte = &v
		case "last_used__lte":
			if vLastUsedLte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'last_used__lte' takes a single value.")
				return
			}
			v := value
			vLastUsedLte = &v
		case "pepper_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pepper_id' takes an integer, got %q.", value))
				return
			}
			vPepperID = append(vPepperID, v)
		case "pepper_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pepper_id__n' takes an integer, got %q.", value))
				return
			}
			vPepperIDn = append(vPepperIDn, v)
		case "pepper_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pepper_id__lt' takes an integer, got %q.", value))
				return
			}
			vPepperIDLt = append(vPepperIDLt, v)
		case "pepper_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pepper_id__lte' takes an integer, got %q.", value))
				return
			}
			vPepperIDLte = append(vPepperIDLte, v)
		case "pepper_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pepper_id__gt' takes an integer, got %q.", value))
				return
			}
			vPepperIDGt = append(vPepperIDGt, v)
		case "pepper_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pepper_id__gte' takes an integer, got %q.", value))
				return
			}
			vPepperIDGte = append(vPepperIDGte, v)
		case "pepper_id__empty":
			if vPepperIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'pepper_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'pepper_id__empty' takes a boolean, got %q.", value))
				return
			}
			vPepperIDEmpty = &v
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "user":
			v := value
			vUser = append(vUser, v)
		case "user__n":
			v := value
			vUsern = append(vUsern, v)
		case "version":
			if vVersion != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'version' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version' takes an integer, got %q.", value))
				return
			}
			vVersion = &v
		case "version__n":
			if vVersionn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'version__n' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'version__n' takes an integer, got %q.", value))
				return
			}
			vVersionn = &v
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
		case "write_enabled":
			if vWriteEnabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'write_enabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'write_enabled' takes a boolean, got %q.", value))
				return
			}
			vWriteEnabled = &v
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
	if len(vUserID) > 0 {
		params.SetUserID(vUserID)
	}
	if len(vUserIDn) > 0 {
		params.SetUserIDn(vUserIDn)
	}
	if vCreated != nil {
		params.SetCreated(vCreated)
	}
	if vCreatedGte != nil {
		params.SetCreatedGte(vCreatedGte)
	}
	if vCreatedLte != nil {
		params.SetCreatedLte(vCreatedLte)
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
	if vEnabled != nil {
		params.SetEnabled(vEnabled)
	}
	if vExpires != nil {
		params.SetExpires(vExpires)
	}
	if vExpiresGte != nil {
		params.SetExpiresGte(vExpiresGte)
	}
	if vExpiresLte != nil {
		params.SetExpiresLte(vExpiresLte)
	}
	if len(vKey) > 0 {
		params.SetKey(vKey)
	}
	if len(vKeyn) > 0 {
		params.SetKeyn(vKeyn)
	}
	if len(vKeyIc) > 0 {
		params.SetKeyIc(vKeyIc)
	}
	if len(vKeyNic) > 0 {
		params.SetKeyNic(vKeyNic)
	}
	if len(vKeyIe) > 0 {
		params.SetKeyIe(vKeyIe)
	}
	if len(vKeyNie) > 0 {
		params.SetKeyNie(vKeyNie)
	}
	if len(vKeyIsw) > 0 {
		params.SetKeyIsw(vKeyIsw)
	}
	if len(vKeyNisw) > 0 {
		params.SetKeyNisw(vKeyNisw)
	}
	if len(vKeyIew) > 0 {
		params.SetKeyIew(vKeyIew)
	}
	if len(vKeyNiew) > 0 {
		params.SetKeyNiew(vKeyNiew)
	}
	if vKeyEmpty != nil {
		params.SetKeyEmpty(vKeyEmpty)
	}
	if len(vKeyRegex) > 0 {
		params.SetKeyRegex(vKeyRegex)
	}
	if len(vKeyIregex) > 0 {
		params.SetKeyIregex(vKeyIregex)
	}
	if vLastUsed != nil {
		params.SetLastUsed(vLastUsed)
	}
	if vLastUsedGte != nil {
		params.SetLastUsedGte(vLastUsedGte)
	}
	if vLastUsedLte != nil {
		params.SetLastUsedLte(vLastUsedLte)
	}
	if len(vPepperID) > 0 {
		params.SetPepperID(vPepperID)
	}
	if len(vPepperIDn) > 0 {
		params.SetPepperIDn(vPepperIDn)
	}
	if len(vPepperIDLt) > 0 {
		params.SetPepperIDLt(vPepperIDLt)
	}
	if len(vPepperIDLte) > 0 {
		params.SetPepperIDLte(vPepperIDLte)
	}
	if len(vPepperIDGt) > 0 {
		params.SetPepperIDGt(vPepperIDGt)
	}
	if len(vPepperIDGte) > 0 {
		params.SetPepperIDGte(vPepperIDGte)
	}
	if vPepperIDEmpty != nil {
		params.SetPepperIDEmpty(vPepperIDEmpty)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vUser) > 0 {
		params.SetUser(vUser)
	}
	if len(vUsern) > 0 {
		params.SetUsern(vUsern)
	}
	if vVersion != nil {
		params.SetVersion(vVersion)
	}
	if vVersionn != nil {
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
	if vWriteEnabled != nil {
		params.SetWriteEnabled(vWriteEnabled)
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
