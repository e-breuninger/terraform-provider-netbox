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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                       netboxapi.FilterInt64s(params.SetID),
		"id__n":                    netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                   netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                  netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                   netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                  netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                netboxapi.FilterBool(params.SetIDEmpty),
		"username":                 netboxapi.FilterStrings(params.SetUsername),
		"username__n":              netboxapi.FilterStrings(params.SetUsernamen),
		"username__ic":             netboxapi.FilterStrings(params.SetUsernameIc),
		"username__nic":            netboxapi.FilterStrings(params.SetUsernameNic),
		"username__ie":             netboxapi.FilterStrings(params.SetUsernameIe),
		"username__nie":            netboxapi.FilterStrings(params.SetUsernameNie),
		"username__isw":            netboxapi.FilterStrings(params.SetUsernameIsw),
		"username__nisw":           netboxapi.FilterStrings(params.SetUsernameNisw),
		"username__iew":            netboxapi.FilterStrings(params.SetUsernameIew),
		"username__niew":           netboxapi.FilterStrings(params.SetUsernameNiew),
		"username__empty":          netboxapi.FilterBool(params.SetUsernameEmpty),
		"username__regex":          netboxapi.FilterStrings(params.SetUsernameRegex),
		"username__iregex":         netboxapi.FilterStrings(params.SetUsernameIregex),
		"date_joined":              netboxapi.FilterStrings(params.SetDateJoined),
		"date_joined__n":           netboxapi.FilterStrings(params.SetDateJoinedn),
		"date_joined__lt":          netboxapi.FilterStrings(params.SetDateJoinedLt),
		"date_joined__lte":         netboxapi.FilterStrings(params.SetDateJoinedLte),
		"date_joined__gt":          netboxapi.FilterStrings(params.SetDateJoinedGt),
		"date_joined__gte":         netboxapi.FilterStrings(params.SetDateJoinedGte),
		"date_joined__empty":       netboxapi.FilterBool(params.SetDateJoinedEmpty),
		"email":                    netboxapi.FilterStrings(params.SetEmail),
		"email__n":                 netboxapi.FilterStrings(params.SetEmailn),
		"email__ic":                netboxapi.FilterStrings(params.SetEmailIc),
		"email__nic":               netboxapi.FilterStrings(params.SetEmailNic),
		"email__ie":                netboxapi.FilterStrings(params.SetEmailIe),
		"email__nie":               netboxapi.FilterStrings(params.SetEmailNie),
		"email__isw":               netboxapi.FilterStrings(params.SetEmailIsw),
		"email__nisw":              netboxapi.FilterStrings(params.SetEmailNisw),
		"email__iew":               netboxapi.FilterStrings(params.SetEmailIew),
		"email__niew":              netboxapi.FilterStrings(params.SetEmailNiew),
		"email__empty":             netboxapi.FilterBool(params.SetEmailEmpty),
		"email__regex":             netboxapi.FilterStrings(params.SetEmailRegex),
		"email__iregex":            netboxapi.FilterStrings(params.SetEmailIregex),
		"first_name":               netboxapi.FilterStrings(params.SetFirstName),
		"first_name__n":            netboxapi.FilterStrings(params.SetFirstNamen),
		"first_name__ic":           netboxapi.FilterStrings(params.SetFirstNameIc),
		"first_name__nic":          netboxapi.FilterStrings(params.SetFirstNameNic),
		"first_name__ie":           netboxapi.FilterStrings(params.SetFirstNameIe),
		"first_name__nie":          netboxapi.FilterStrings(params.SetFirstNameNie),
		"first_name__isw":          netboxapi.FilterStrings(params.SetFirstNameIsw),
		"first_name__nisw":         netboxapi.FilterStrings(params.SetFirstNameNisw),
		"first_name__iew":          netboxapi.FilterStrings(params.SetFirstNameIew),
		"first_name__niew":         netboxapi.FilterStrings(params.SetFirstNameNiew),
		"first_name__empty":        netboxapi.FilterBool(params.SetFirstNameEmpty),
		"first_name__regex":        netboxapi.FilterStrings(params.SetFirstNameRegex),
		"first_name__iregex":       netboxapi.FilterStrings(params.SetFirstNameIregex),
		"group":                    netboxapi.FilterStrings(params.SetGroup),
		"group__n":                 netboxapi.FilterStrings(params.SetGroupn),
		"group_id":                 netboxapi.FilterInt64s(params.SetGroupID),
		"group_id__n":              netboxapi.FilterInt64s(params.SetGroupIDn),
		"is_active":                netboxapi.FilterBool(params.SetIsActive),
		"is_superuser":             netboxapi.FilterBool(params.SetIsSuperuser),
		"last_login":               netboxapi.FilterStrings(params.SetLastLogin),
		"last_login__n":            netboxapi.FilterStrings(params.SetLastLoginn),
		"last_login__lt":           netboxapi.FilterStrings(params.SetLastLoginLt),
		"last_login__lte":          netboxapi.FilterStrings(params.SetLastLoginLte),
		"last_login__gt":           netboxapi.FilterStrings(params.SetLastLoginGt),
		"last_login__gte":          netboxapi.FilterStrings(params.SetLastLoginGte),
		"last_login__empty":        netboxapi.FilterBool(params.SetLastLoginEmpty),
		"last_name":                netboxapi.FilterStrings(params.SetLastName),
		"last_name__n":             netboxapi.FilterStrings(params.SetLastNamen),
		"last_name__ic":            netboxapi.FilterStrings(params.SetLastNameIc),
		"last_name__nic":           netboxapi.FilterStrings(params.SetLastNameNic),
		"last_name__ie":            netboxapi.FilterStrings(params.SetLastNameIe),
		"last_name__nie":           netboxapi.FilterStrings(params.SetLastNameNie),
		"last_name__isw":           netboxapi.FilterStrings(params.SetLastNameIsw),
		"last_name__nisw":          netboxapi.FilterStrings(params.SetLastNameNisw),
		"last_name__iew":           netboxapi.FilterStrings(params.SetLastNameIew),
		"last_name__niew":          netboxapi.FilterStrings(params.SetLastNameNiew),
		"last_name__empty":         netboxapi.FilterBool(params.SetLastNameEmpty),
		"last_name__regex":         netboxapi.FilterStrings(params.SetLastNameRegex),
		"last_name__iregex":        netboxapi.FilterStrings(params.SetLastNameIregex),
		"notification_group_id":    netboxapi.FilterInt64s(params.SetNotificationGroupID),
		"notification_group_id__n": netboxapi.FilterInt64s(params.SetNotificationGroupIDn),
		"owner":                    netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                 netboxapi.FilterStrings(params.SetOwnern),
		"owner_id":                 netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":              netboxapi.FilterInt64s(params.SetOwnerIDn),
		"permission_id":            netboxapi.FilterInt64s(params.SetPermissionID),
		"permission_id__n":         netboxapi.FilterInt64s(params.SetPermissionIDn),
		"q":                        netboxapi.FilterString(params.SetQ),
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
