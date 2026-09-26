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
	_ datasource.DataSource              = (*ownerDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ownerDataSource)(nil)
)

// NewOwnerDataSource returns a new owner data source.
func NewOwnerDataSource() datasource.DataSource {
	return &ownerDataSource{}
}

type ownerDataSource struct {
	client *netboxapi.Client
}

// ownerDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type ownerDataSourceModel struct {
	ownerResourceModel
	NameContains types.String `tfsdk:"name_contains"`
}

func (d *ownerDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_owner"
}

func (d *ownerDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:A NetBox owner (users.owner): the users and groups an object can be assigned to as its owner.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/users/owner/):\n\n> An owner is a set of users and/or groups who are responsible for the administration of certain resources within NetBox. The users and groups assigned to an owner are referred to as its members. Owner assignments are useful for indicating which parties are responsible for the administration of a particular object.\n>\n> Most objects within NetBox can be assigned an owner, although this is not required.",
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
			"owner_group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner group the owner belongs to.",
			},
			"user_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the users that make up the owner.",
			},
			"user_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the user groups that make up the owner.",
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
			},
		},
	}
}

func (d *ownerDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ownerDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ownerDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := users.NewUsersOwnersListParams()
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
	if !state.OwnerGroupID.IsNull() {
		v := state.OwnerGroupID.ValueInt64()
		params.SetGroupID([]int64{v})
		hasInput = true
	}
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or owner_group_id and/or name_contains to look up a netbox_owner.")
		return
	}
	res, err := d.client.Users.UsersOwnersListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_owner", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_owner",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenOwner(ctx, netboxapi.OwnerResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.ownerResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
