// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

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
	_ datasource.DataSource              = (*groupDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*groupDataSource)(nil)
)

// NewGroupDataSource returns a new group data source.
func NewGroupDataSource() datasource.DataSource {
	return &groupDataSource{}
}

type groupDataSource struct {
	client *netboxapi.Client
}

// groupDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type groupDataSourceModel struct {
	groupResourceModel
	NotificationGroupID types.Int64  `tfsdk:"notification_group_id"`
	OwnerName           types.String `tfsdk:"owner_name"`
	OwnerID             types.Int64  `tfsdk:"owner_id"`
	PermissionID        types.Int64  `tfsdk:"permission_id"`
	Q                   types.String `tfsdk:"q"`
	UserID              types.Int64  `tfsdk:"user_id"`
}

func (d *groupDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_group"
}

func (d *groupDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:A user group (users.group).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/users/group/):\n\n> A group is a collection of [users](https://netboxlabs.com/docs/netbox/models/users/user/) which share a common set of permissions. Assigning [object permissions](https://netboxlabs.com/docs/netbox/models/users/objectpermission/) to a group, rather than to individual users, simplifies the administration of permissions for related users (e.g. members of a particular team).\n>\n> A user inherits the union of all permissions assigned to each group of which they are a member, in addition to any permissions assigned directly to the user.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 150),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"user_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of users in the group.",
			},
			"url": schema.StringAttribute{
				Computed: true,
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
			"user_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the user.",
			},
		},
	}
}

func (d *groupDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *groupDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data groupDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := users.NewUsersGroupsListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
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
	if !state.UserID.IsNull() {
		v := state.UserID.ValueInt64()
		params.SetUserID([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_group.")
		return
	}
	res, err := d.client.Users.UsersGroupsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_group", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_group",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenGroup(ctx, netboxapi.GroupResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.groupResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
