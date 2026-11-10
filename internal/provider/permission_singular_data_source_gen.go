// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*permissionDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*permissionDataSource)(nil)
)

// NewPermissionDataSource returns a new permission data source.
func NewPermissionDataSource() datasource.DataSource {
	return &permissionDataSource{}
}

type permissionDataSource struct {
	client *netboxapi.Client
}

// permissionDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type permissionDataSourceModel struct {
	permissionResourceModel
	CanAdd         types.Bool   `tfsdk:"can_add"`
	CanChange      types.Bool   `tfsdk:"can_change"`
	CanDelete      types.Bool   `tfsdk:"can_delete"`
	CanView        types.Bool   `tfsdk:"can_view"`
	GroupName      types.String `tfsdk:"group_name"`
	GroupID        types.Int64  `tfsdk:"group_id"`
	ObjectType     types.String `tfsdk:"object_type"`
	ObjectTypeID   types.Int64  `tfsdk:"object_type_id"`
	ObjectTypeName types.Int64  `tfsdk:"object_type_name"`
	Q              types.String `tfsdk:"q"`
	UserName       types.String `tfsdk:"user_name"`
	UserID         types.Int64  `tfsdk:"user_id"`
}

func (d *permissionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_permission"
}

func (d *permissionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:An object permission granting actions on object types to users and groups (users.objectpermission).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/users/objectpermission/):\n\n> An object permission grants the ability to perform one or more actions (e.g. view, add, change, delete) against a defined set of object types, and may be restricted to a subset of objects matching a configured filter. Permissions are assigned to [users](https://netboxlabs.com/docs/netbox/models/users/user/) and/or [groups](https://netboxlabs.com/docs/netbox/models/users/group/); a user's effective permissions are the union of those assigned directly and those inherited via group membership.\n>\n> See the [permissions documentation](https://netboxlabs.com/docs/netbox/administration/permissions/) for a detailed walkthrough of how permissions are evaluated.",
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
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the permission is in effect.",
			},
			"object_types": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Object types the permission applies to, as app-labelled content types (dcim.device).",
			},
			"actions": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Permitted actions: view, add, change, delete or a custom action name.",
			},
			"group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the groups the permission is granted to.",
			},
			"groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use group_ids instead.",
				Description:        "Deprecated alias of group_ids.",
			},
			"user_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the users the permission is granted to.",
			},
			"users": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use user_ids instead.",
				Description:        "Deprecated alias of user_ids.",
			},
			"constraints": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Queryset filter as JSON text, an object or a list of objects (use jsonencode()).",
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"can_add": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether can add is set.",
			},
			"can_change": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether can change is set.",
			},
			"can_delete": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether can delete is set.",
			},
			"can_view": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether can view is set.",
			},
			"group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the group.",
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the group.",
			},
			"object_type": schema.StringAttribute{
				Optional:    true,
				Description: "Object type to match.",
			},
			"object_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the object type.",
			},
			"object_type_name": schema.Int64Attribute{
				Optional:    true,
				Description: "Name of the object types.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"user_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the user.",
			},
			"user_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the user.",
			},
		},
	}
}

func (d *permissionDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *permissionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data permissionDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := users.NewUsersPermissionsListParams()
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
	if !state.CanAdd.IsNull() {
		v := state.CanAdd.ValueBool()
		params.SetCanAdd(&v)
		hasInput = true
	}
	if !state.CanChange.IsNull() {
		v := state.CanChange.ValueBool()
		params.SetCanChange(&v)
		hasInput = true
	}
	if !state.CanDelete.IsNull() {
		v := state.CanDelete.ValueBool()
		params.SetCanDelete(&v)
		hasInput = true
	}
	if !state.CanView.IsNull() {
		v := state.CanView.ValueBool()
		params.SetCanView(&v)
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.Enabled.IsNull() {
		v := state.Enabled.ValueBool()
		params.SetEnabled(&v)
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
	if !state.ObjectType.IsNull() {
		v := state.ObjectType.ValueString()
		params.SetObjectType([]string{v})
		hasInput = true
	}
	if !state.ObjectTypeID.IsNull() {
		v := state.ObjectTypeID.ValueInt64()
		params.SetObjectTypeID([]int64{v})
		hasInput = true
	}
	if !state.ObjectTypeName.IsNull() {
		v := state.ObjectTypeName.ValueInt64()
		params.SetObjectTypes([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.UserName.IsNull() {
		v := state.UserName.ValueString()
		params.SetUser([]string{v})
		hasInput = true
	}
	if !state.UserID.IsNull() {
		v := state.UserID.ValueInt64()
		params.SetUserID([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_permission.")
		return
	}
	res, err := d.client.Users.UsersPermissionsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_permission", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_permission",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenPermission(ctx, netboxapi.PermissionResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.permissionResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
