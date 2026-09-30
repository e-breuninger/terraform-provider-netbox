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
	_ datasource.DataSource              = (*usersDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*usersDataSource)(nil)
)

// NewUsersDataSource returns a new users data source, which
// lists user objects matching its filters.
func NewUsersDataSource() datasource.DataSource {
	return &usersDataSource{}
}

type usersDataSource struct {
	client *netboxapi.Client
}

// usersDataSourceModel is the data source model: the filters, the limit and the matching
// users.
type usersDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"users"`
}

// userResourceAttrTypes is the attribute type map of userResourceModel.
func userResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":          types.Int64Type,
		"username":    types.StringType,
		"password":    types.StringType,
		"first_name":  types.StringType,
		"last_name":   types.StringType,
		"email":       types.StringType,
		"is_active":   types.BoolType,
		"active":      types.BoolType,
		"group_ids":   types.SetType{ElemType: types.Int64Type},
		"date_joined": types.StringType,
		"last_login":  types.StringType,
		"url":         types.StringType,
	}
}

func (d *usersDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_users"
}

func (d *usersDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:Lists user objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, username, username__empty, username__ic, username__ie, username__iew, username__iregex, username__isw, username__n, username__nic, username__nie, username__niew, username__nisw, username__regex. Repeating a name sends that parameter once per value.",
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
			"users": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching users.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"username": schema.StringAttribute{
							Computed: true,
						},
						"password": schema.StringAttribute{
							Computed:    true,
							Sensitive:   true,
							Description: "Password of the user. Never returned by NetBox, so changes made outside Terraform go unnoticed.",
						},
						"first_name": schema.StringAttribute{
							Computed: true,
						},
						"last_name": schema.StringAttribute{
							Computed: true,
						},
						"email": schema.StringAttribute{
							Computed:    true,
							Description: "Email address.",
						},
						"is_active": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the user may log in.",
						},
						"active": schema.BoolAttribute{
							Computed:           true,
							DeprecationMessage: "Use is_active instead.",
							Description:        "Deprecated alias of is_active.",
						},
						"group_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the groups the user belongs to.",
						},
						"date_joined": schema.StringAttribute{
							Computed:    true,
							Description: "When the account was created.",
						},
						"last_login": schema.StringAttribute{
							Computed:    true,
							Description: "When the user last logged in; null until the first login.",
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

func (d *usersDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *usersDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data usersDataSourceModel

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

	fetchAll := false
	_ = fetchAll
	var responseDTOs []*netboxapi.UserResponseDTO

	params := users.NewUsersUsersListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vUsername []string
	var vUsernamen []string
	var vUsernameNic []string
	var vUsernameIe []string
	var vUsernameNie []string
	var vUsernameIsw []string
	var vUsernameNisw []string
	var vUsernameIew []string
	var vUsernameNiew []string
	var vUsernameEmpty *bool
	var vUsernameRegex []string
	var vUsernameIregex []string
	var vUsernameIc []string
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
		case "username":
			v := value
			vUsername = append(vUsername, v)
		case "username__n":
			v := value
			vUsernamen = append(vUsernamen, v)
		case "username__nic":
			v := value
			vUsernameNic = append(vUsernameNic, v)
		case "username__ie":
			v := value
			vUsernameIe = append(vUsernameIe, v)
		case "username__nie":
			v := value
			vUsernameNie = append(vUsernameNie, v)
		case "username__isw":
			v := value
			vUsernameIsw = append(vUsernameIsw, v)
		case "username__nisw":
			v := value
			vUsernameNisw = append(vUsernameNisw, v)
		case "username__iew":
			v := value
			vUsernameIew = append(vUsernameIew, v)
		case "username__niew":
			v := value
			vUsernameNiew = append(vUsernameNiew, v)
		case "username__empty":
			if vUsernameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'username__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'username__empty' takes a boolean, got %q.", value))
				return
			}
			vUsernameEmpty = &v
		case "username__regex":
			v := value
			vUsernameRegex = append(vUsernameRegex, v)
		case "username__iregex":
			v := value
			vUsernameIregex = append(vUsernameIregex, v)
		case "username__ic":
			v := value
			vUsernameIc = append(vUsernameIc, v)
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
	if len(vUsername) > 0 {
		params.SetUsername(vUsername)
	}
	if len(vUsernamen) > 0 {
		params.SetUsernamen(vUsernamen)
	}
	if len(vUsernameNic) > 0 {
		params.SetUsernameNic(vUsernameNic)
	}
	if len(vUsernameIe) > 0 {
		params.SetUsernameIe(vUsernameIe)
	}
	if len(vUsernameNie) > 0 {
		params.SetUsernameNie(vUsernameNie)
	}
	if len(vUsernameIsw) > 0 {
		params.SetUsernameIsw(vUsernameIsw)
	}
	if len(vUsernameNisw) > 0 {
		params.SetUsernameNisw(vUsernameNisw)
	}
	if len(vUsernameIew) > 0 {
		params.SetUsernameIew(vUsernameIew)
	}
	if len(vUsernameNiew) > 0 {
		params.SetUsernameNiew(vUsernameNiew)
	}
	if vUsernameEmpty != nil {
		params.SetUsernameEmpty(vUsernameEmpty)
	}
	if len(vUsernameRegex) > 0 {
		params.SetUsernameRegex(vUsernameRegex)
	}
	if len(vUsernameIregex) > 0 {
		params.SetUsernameIregex(vUsernameIregex)
	}
	if len(vUsernameIc) > 0 {
		params.SetUsernameIc(vUsernameIc)
	}
	pageSize := limit
	if fetchAll {
		pageSize = 1000
	}
	params.SetLimit(&pageSize)
	for offset := int64(0); ; {
		res, err := d.client.Users.UsersUsersListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_users", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.UserResponseDTOFromGoNetbox(goNetboxModel))
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
	items := make([]userResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model userResourceModel
		model.Password = types.StringNull()
		resp.Diagnostics.Append(flattenUser(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: userResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
