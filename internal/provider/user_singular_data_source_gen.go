// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*userDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*userDataSource)(nil)
)

// NewUserDataSource returns a new user data source.
func NewUserDataSource() datasource.DataSource {
	return &userDataSource{}
}

type userDataSource struct {
	client *netboxapi.Client
}

// userDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type userDataSourceModel struct {
	userResourceModel
	GroupName           types.String `tfsdk:"group_name"`
	GroupID             types.Int64  `tfsdk:"group_id"`
	IsSuperuser         types.Bool   `tfsdk:"is_superuser"`
	NotificationGroupID types.Int64  `tfsdk:"notification_group_id"`
	OwnerName           types.String `tfsdk:"owner_name"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	PermissionID        types.Int64  `tfsdk:"permission_id"`
	Q                   types.String `tfsdk:"q"`
}

func (d *userDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (d *userDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:A NetBox user (users.user).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/users/user/):\n\n> A user represents an individual account in NetBox. Users authenticate to access the application, and may be granted permissions either directly or through their assigned [groups](https://netboxlabs.com/docs/netbox/models/users/group/). Each user can hold one or more API [tokens](https://netboxlabs.com/docs/netbox/models/users/token/) for use with the REST and GraphQL APIs.\n>\n> NetBox extends Django's stock user model to support multiple API tokens per user, configurable [object permissions](https://netboxlabs.com/docs/netbox/models/users/objectpermission/), and integration with [remote authentication backends](https://netboxlabs.com/docs/netbox/administration/authentication/overview/).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"username": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 150),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[\w.@+-]+$`), ""),
				},
			},
			"password": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "Password of the user. Never returned by NetBox, so changes made outside Terraform go unnoticed.",
			},
			"first_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(150),
				},
			},
			"last_name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(150),
				},
			},
			"email": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Email address.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(254),
				},
			},
			"is_active": schema.BoolAttribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "When the account was created.",
			},
			"last_login": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When the user last logged in; null until the first login.",
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the group.",
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the group.",
			},
			"is_superuser": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether is superuser is set.",
			},
			"notification_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the notification group.",
			},
			"owner_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner.",
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner.",
			},
			"permission_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the permission.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
		},
	}
}

func (d *userDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *userDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data userDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := users.NewUsersUsersListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Username.IsNull() {
		v := state.Username.ValueString()
		params.SetUsername([]string{v})
		hasInput = true
	}
	if !state.DateJoined.IsNull() {
		v := state.DateJoined.ValueString()
		params.SetDateJoined([]string{v})
		hasInput = true
	}
	if !state.Email.IsNull() {
		v := state.Email.ValueString()
		params.SetEmail([]string{v})
		hasInput = true
	}
	if !state.FirstName.IsNull() {
		v := state.FirstName.ValueString()
		params.SetFirstName([]string{v})
		hasInput = true
	}
	if !state.GroupName.IsNull() {
		v := state.GroupName.ValueString()
		params.SetGroup([]string{v})
		hasInput = true
	}
	if !state.GroupID.IsNull() {
		v := state.GroupID.ValueInt64()
		params.SetGroupID([]int64{v})
		hasInput = true
	}
	if !state.IsActive.IsNull() {
		v := state.IsActive.ValueBool()
		params.SetIsActive(&v)
		hasInput = true
	}
	if !state.IsSuperuser.IsNull() {
		v := state.IsSuperuser.ValueBool()
		params.SetIsSuperuser(&v)
		hasInput = true
	}
	if !state.LastLogin.IsNull() {
		v := state.LastLogin.ValueString()
		params.SetLastLogin([]string{v})
		hasInput = true
	}
	if !state.LastName.IsNull() {
		v := state.LastName.ValueString()
		params.SetLastName([]string{v})
		hasInput = true
	}
	if !state.NotificationGroupID.IsNull() {
		v := state.NotificationGroupID.ValueInt64()
		params.SetNotificationGroupID([]int64{v})
		hasInput = true
	}
	if !state.OwnerName.IsNull() {
		v := state.OwnerName.ValueString()
		params.SetOwner([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.PermissionID.IsNull() {
		v := state.PermissionID.ValueInt64()
		params.SetPermissionID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_user.")
		return
	}
	res, err := d.client.Users.UsersUsersListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_user", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_user",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenUser(ctx, netboxapi.UserResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.userResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
