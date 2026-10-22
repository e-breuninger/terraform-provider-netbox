// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*tokenDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*tokenDataSource)(nil)
)

// NewTokenDataSource returns a new token data source.
func NewTokenDataSource() datasource.DataSource {
	return &tokenDataSource{}
}

type tokenDataSource struct {
	client *netboxapi.Client
}

func (d *tokenDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_token"
}

func (d *tokenDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:An API token of a user (users.token).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/users/token/):\n\n> A token is a secret credential associated with a [user](https://netboxlabs.com/docs/netbox/models/users/user/) which authenticates requests to NetBox's REST and GraphQL APIs. A user may hold multiple tokens; each can be independently expired, restricted, or revoked.\n>\n> Beginning with NetBox v4.5, two token versions are supported. v2 tokens (the default for newly-created tokens) are stored only as a salted HMAC digest, and the plaintext is shown to the user only once at creation time. Legacy v1 tokens store the plaintext directly; **their use is deprecated and support will be removed in NetBox v5.0.** See the [REST API authentication](https://netboxlabs.com/docs/netbox/integrations/rest-api/#authentication) documentation for the request header formats used by each version.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"user_id": schema.Int64Attribute{
				Optional:    true,
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
	}
}

func (d *tokenDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *tokenDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data tokenResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := users.NewUsersTokensListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.UserID.IsNull() {
		v := state.UserID.ValueInt64()
		params.SetUserID([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or user_id to look up a netbox_token.")
		return
	}
	res, err := d.client.Users.UsersTokensListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_token", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_token",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenToken(ctx, netboxapi.TokenResponseDTOFromGoNetbox(res.Payload.Results[0]), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
