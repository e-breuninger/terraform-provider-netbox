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
				Description: "Query filters, sent as API list parameters. Supported names: date_joined, date_joined__empty, date_joined__gt, date_joined__gte, date_joined__lt, date_joined__lte, date_joined__n, email, email__empty, email__ic, email__ie, email__iew, email__iregex, email__isw, email__n, email__nic, email__nie, email__niew, email__nisw, email__regex, first_name, first_name__empty, first_name__ic, first_name__ie, first_name__iew, first_name__iregex, first_name__isw, first_name__n, first_name__nic, first_name__nie, first_name__niew, first_name__nisw, first_name__regex, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, is_active, is_superuser, last_login, last_login__empty, last_login__gt, last_login__gte, last_login__lt, last_login__lte, last_login__n, last_name, last_name__empty, last_name__ic, last_name__ie, last_name__iew, last_name__iregex, last_name__isw, last_name__n, last_name__nic, last_name__nie, last_name__niew, last_name__nisw, last_name__regex, notification_group_id, notification_group_id__n, owner, owner__n, owner_id, owner_id__n, permission_id, permission_id__n, q, username, username__empty, username__ic, username__ie, username__iew, username__iregex, username__isw, username__n, username__nic, username__nie, username__niew, username__nisw, username__regex. Repeating a name sends that parameter once per value.",
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
	var vUsernameIc []string
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
	var vDateJoined []string
	var vDateJoinedn []string
	var vDateJoinedLt []string
	var vDateJoinedLte []string
	var vDateJoinedGt []string
	var vDateJoinedGte []string
	var vDateJoinedEmpty *bool
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
	var vFirstName []string
	var vFirstNamen []string
	var vFirstNameIc []string
	var vFirstNameNic []string
	var vFirstNameIe []string
	var vFirstNameNie []string
	var vFirstNameIsw []string
	var vFirstNameNisw []string
	var vFirstNameIew []string
	var vFirstNameNiew []string
	var vFirstNameEmpty *bool
	var vFirstNameRegex []string
	var vFirstNameIregex []string
	var vGroup []string
	var vGroupn []string
	var vGroupID []int64
	var vGroupIDn []int64
	var vIsActive *bool
	var vIsSuperuser *bool
	var vLastLogin []string
	var vLastLoginn []string
	var vLastLoginLt []string
	var vLastLoginLte []string
	var vLastLoginGt []string
	var vLastLoginGte []string
	var vLastLoginEmpty *bool
	var vLastName []string
	var vLastNamen []string
	var vLastNameIc []string
	var vLastNameNic []string
	var vLastNameIe []string
	var vLastNameNie []string
	var vLastNameIsw []string
	var vLastNameNisw []string
	var vLastNameIew []string
	var vLastNameNiew []string
	var vLastNameEmpty *bool
	var vLastNameRegex []string
	var vLastNameIregex []string
	var vNotificationGroupID []int64
	var vNotificationGroupIDn []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vPermissionID []int64
	var vPermissionIDn []int64
	var vQ *string
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
		case "username__ic":
			v := value
			vUsernameIc = append(vUsernameIc, v)
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
		case "date_joined":
			v := value
			vDateJoined = append(vDateJoined, v)
		case "date_joined__n":
			v := value
			vDateJoinedn = append(vDateJoinedn, v)
		case "date_joined__lt":
			v := value
			vDateJoinedLt = append(vDateJoinedLt, v)
		case "date_joined__lte":
			v := value
			vDateJoinedLte = append(vDateJoinedLte, v)
		case "date_joined__gt":
			v := value
			vDateJoinedGt = append(vDateJoinedGt, v)
		case "date_joined__gte":
			v := value
			vDateJoinedGte = append(vDateJoinedGte, v)
		case "date_joined__empty":
			if vDateJoinedEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'date_joined__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'date_joined__empty' takes a boolean, got %q.", value))
				return
			}
			vDateJoinedEmpty = &v
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
		case "first_name":
			v := value
			vFirstName = append(vFirstName, v)
		case "first_name__n":
			v := value
			vFirstNamen = append(vFirstNamen, v)
		case "first_name__ic":
			v := value
			vFirstNameIc = append(vFirstNameIc, v)
		case "first_name__nic":
			v := value
			vFirstNameNic = append(vFirstNameNic, v)
		case "first_name__ie":
			v := value
			vFirstNameIe = append(vFirstNameIe, v)
		case "first_name__nie":
			v := value
			vFirstNameNie = append(vFirstNameNie, v)
		case "first_name__isw":
			v := value
			vFirstNameIsw = append(vFirstNameIsw, v)
		case "first_name__nisw":
			v := value
			vFirstNameNisw = append(vFirstNameNisw, v)
		case "first_name__iew":
			v := value
			vFirstNameIew = append(vFirstNameIew, v)
		case "first_name__niew":
			v := value
			vFirstNameNiew = append(vFirstNameNiew, v)
		case "first_name__empty":
			if vFirstNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'first_name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'first_name__empty' takes a boolean, got %q.", value))
				return
			}
			vFirstNameEmpty = &v
		case "first_name__regex":
			v := value
			vFirstNameRegex = append(vFirstNameRegex, v)
		case "first_name__iregex":
			v := value
			vFirstNameIregex = append(vFirstNameIregex, v)
		case "group":
			v := value
			vGroup = append(vGroup, v)
		case "group__n":
			v := value
			vGroupn = append(vGroupn, v)
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
		case "is_active":
			if vIsActive != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'is_active' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'is_active' takes a boolean, got %q.", value))
				return
			}
			vIsActive = &v
		case "is_superuser":
			if vIsSuperuser != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'is_superuser' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'is_superuser' takes a boolean, got %q.", value))
				return
			}
			vIsSuperuser = &v
		case "last_login":
			v := value
			vLastLogin = append(vLastLogin, v)
		case "last_login__n":
			v := value
			vLastLoginn = append(vLastLoginn, v)
		case "last_login__lt":
			v := value
			vLastLoginLt = append(vLastLoginLt, v)
		case "last_login__lte":
			v := value
			vLastLoginLte = append(vLastLoginLte, v)
		case "last_login__gt":
			v := value
			vLastLoginGt = append(vLastLoginGt, v)
		case "last_login__gte":
			v := value
			vLastLoginGte = append(vLastLoginGte, v)
		case "last_login__empty":
			if vLastLoginEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'last_login__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'last_login__empty' takes a boolean, got %q.", value))
				return
			}
			vLastLoginEmpty = &v
		case "last_name":
			v := value
			vLastName = append(vLastName, v)
		case "last_name__n":
			v := value
			vLastNamen = append(vLastNamen, v)
		case "last_name__ic":
			v := value
			vLastNameIc = append(vLastNameIc, v)
		case "last_name__nic":
			v := value
			vLastNameNic = append(vLastNameNic, v)
		case "last_name__ie":
			v := value
			vLastNameIe = append(vLastNameIe, v)
		case "last_name__nie":
			v := value
			vLastNameNie = append(vLastNameNie, v)
		case "last_name__isw":
			v := value
			vLastNameIsw = append(vLastNameIsw, v)
		case "last_name__nisw":
			v := value
			vLastNameNisw = append(vLastNameNisw, v)
		case "last_name__iew":
			v := value
			vLastNameIew = append(vLastNameIew, v)
		case "last_name__niew":
			v := value
			vLastNameNiew = append(vLastNameNiew, v)
		case "last_name__empty":
			if vLastNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'last_name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'last_name__empty' takes a boolean, got %q.", value))
				return
			}
			vLastNameEmpty = &v
		case "last_name__regex":
			v := value
			vLastNameRegex = append(vLastNameRegex, v)
		case "last_name__iregex":
			v := value
			vLastNameIregex = append(vLastNameIregex, v)
		case "notification_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'notification_group_id' takes an integer, got %q.", value))
				return
			}
			vNotificationGroupID = append(vNotificationGroupID, v)
		case "notification_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'notification_group_id__n' takes an integer, got %q.", value))
				return
			}
			vNotificationGroupIDn = append(vNotificationGroupIDn, v)
		case "owner":
			v := value
			vOwner = append(vOwner, v)
		case "owner__n":
			v := value
			vOwnern = append(vOwnern, v)
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
		case "permission_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'permission_id' takes an integer, got %q.", value))
				return
			}
			vPermissionID = append(vPermissionID, v)
		case "permission_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'permission_id__n' takes an integer, got %q.", value))
				return
			}
			vPermissionIDn = append(vPermissionIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
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
	if len(vUsernameIc) > 0 {
		params.SetUsernameIc(vUsernameIc)
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
	if len(vDateJoined) > 0 {
		params.SetDateJoined(vDateJoined)
	}
	if len(vDateJoinedn) > 0 {
		params.SetDateJoinedn(vDateJoinedn)
	}
	if len(vDateJoinedLt) > 0 {
		params.SetDateJoinedLt(vDateJoinedLt)
	}
	if len(vDateJoinedLte) > 0 {
		params.SetDateJoinedLte(vDateJoinedLte)
	}
	if len(vDateJoinedGt) > 0 {
		params.SetDateJoinedGt(vDateJoinedGt)
	}
	if len(vDateJoinedGte) > 0 {
		params.SetDateJoinedGte(vDateJoinedGte)
	}
	if vDateJoinedEmpty != nil {
		params.SetDateJoinedEmpty(vDateJoinedEmpty)
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
	if len(vFirstName) > 0 {
		params.SetFirstName(vFirstName)
	}
	if len(vFirstNamen) > 0 {
		params.SetFirstNamen(vFirstNamen)
	}
	if len(vFirstNameIc) > 0 {
		params.SetFirstNameIc(vFirstNameIc)
	}
	if len(vFirstNameNic) > 0 {
		params.SetFirstNameNic(vFirstNameNic)
	}
	if len(vFirstNameIe) > 0 {
		params.SetFirstNameIe(vFirstNameIe)
	}
	if len(vFirstNameNie) > 0 {
		params.SetFirstNameNie(vFirstNameNie)
	}
	if len(vFirstNameIsw) > 0 {
		params.SetFirstNameIsw(vFirstNameIsw)
	}
	if len(vFirstNameNisw) > 0 {
		params.SetFirstNameNisw(vFirstNameNisw)
	}
	if len(vFirstNameIew) > 0 {
		params.SetFirstNameIew(vFirstNameIew)
	}
	if len(vFirstNameNiew) > 0 {
		params.SetFirstNameNiew(vFirstNameNiew)
	}
	if vFirstNameEmpty != nil {
		params.SetFirstNameEmpty(vFirstNameEmpty)
	}
	if len(vFirstNameRegex) > 0 {
		params.SetFirstNameRegex(vFirstNameRegex)
	}
	if len(vFirstNameIregex) > 0 {
		params.SetFirstNameIregex(vFirstNameIregex)
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
	if vIsActive != nil {
		params.SetIsActive(vIsActive)
	}
	if vIsSuperuser != nil {
		params.SetIsSuperuser(vIsSuperuser)
	}
	if len(vLastLogin) > 0 {
		params.SetLastLogin(vLastLogin)
	}
	if len(vLastLoginn) > 0 {
		params.SetLastLoginn(vLastLoginn)
	}
	if len(vLastLoginLt) > 0 {
		params.SetLastLoginLt(vLastLoginLt)
	}
	if len(vLastLoginLte) > 0 {
		params.SetLastLoginLte(vLastLoginLte)
	}
	if len(vLastLoginGt) > 0 {
		params.SetLastLoginGt(vLastLoginGt)
	}
	if len(vLastLoginGte) > 0 {
		params.SetLastLoginGte(vLastLoginGte)
	}
	if vLastLoginEmpty != nil {
		params.SetLastLoginEmpty(vLastLoginEmpty)
	}
	if len(vLastName) > 0 {
		params.SetLastName(vLastName)
	}
	if len(vLastNamen) > 0 {
		params.SetLastNamen(vLastNamen)
	}
	if len(vLastNameIc) > 0 {
		params.SetLastNameIc(vLastNameIc)
	}
	if len(vLastNameNic) > 0 {
		params.SetLastNameNic(vLastNameNic)
	}
	if len(vLastNameIe) > 0 {
		params.SetLastNameIe(vLastNameIe)
	}
	if len(vLastNameNie) > 0 {
		params.SetLastNameNie(vLastNameNie)
	}
	if len(vLastNameIsw) > 0 {
		params.SetLastNameIsw(vLastNameIsw)
	}
	if len(vLastNameNisw) > 0 {
		params.SetLastNameNisw(vLastNameNisw)
	}
	if len(vLastNameIew) > 0 {
		params.SetLastNameIew(vLastNameIew)
	}
	if len(vLastNameNiew) > 0 {
		params.SetLastNameNiew(vLastNameNiew)
	}
	if vLastNameEmpty != nil {
		params.SetLastNameEmpty(vLastNameEmpty)
	}
	if len(vLastNameRegex) > 0 {
		params.SetLastNameRegex(vLastNameRegex)
	}
	if len(vLastNameIregex) > 0 {
		params.SetLastNameIregex(vLastNameIregex)
	}
	if len(vNotificationGroupID) > 0 {
		params.SetNotificationGroupID(vNotificationGroupID)
	}
	if len(vNotificationGroupIDn) > 0 {
		params.SetNotificationGroupIDn(vNotificationGroupIDn)
	}
	if len(vOwner) > 0 {
		params.SetOwner(vOwner)
	}
	if len(vOwnern) > 0 {
		params.SetOwnern(vOwnern)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
	}
	if len(vPermissionID) > 0 {
		params.SetPermissionID(vPermissionID)
	}
	if len(vPermissionIDn) > 0 {
		params.SetPermissionIDn(vPermissionIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Users.UsersUsersListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_users", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.UserResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
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
