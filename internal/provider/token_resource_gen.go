// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/boolplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64default"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*tokenResource)(nil)
	_ resource.ResourceWithConfigure   = (*tokenResource)(nil)
	_ resource.ResourceWithImportState = (*tokenResource)(nil)
)

// NewTokenResource returns a new token resource.
func NewTokenResource() resource.Resource {
	return &tokenResource{}
}

type tokenResource struct {
	client *netboxapi.Client
}

func (r *tokenResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_token"
}

func (r *tokenResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:An API token of a user (users.token).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/users/token/):\n\n> A token is a secret credential associated with a [user](https://netboxlabs.com/docs/netbox/models/users/user/) which authenticates requests to NetBox's REST and GraphQL APIs. A user may hold multiple tokens; each can be independently expired, restricted, or revoked.\n>\n> Beginning with NetBox v4.5, two token versions are supported. v2 tokens (the default for newly-created tokens) are stored only as a salted HMAC digest, and the plaintext is shown to the user only once at creation time. Legacy v1 tokens store the plaintext directly; **their use is deprecated and support will be removed in NetBox v5.0.** See the [REST API authentication](https://netboxlabs.com/docs/netbox/integrations/rest-api/#authentication) documentation for the request header formats used by each version.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"user_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the user the token belongs to.",
			},
			"key": schema.StringAttribute{
				Computed:    true,
				Description: "Identification key NetBox generated for the token; not the secret.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"token": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "The secret. NetBox returns it in the create response only, so an imported token has none in state and a secret changed outside Terraform goes unnoticed. Changing it replaces the token.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.RequiresReplace(),
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the token can authenticate; disabling keeps the token but rejects its use.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"version": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Token format version (1 legacy plaintext, 2 hashed). Defaults to 2. Changing it replaces the token. One of: 1, 2.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
					int64planmodifier.UseStateForUnknown(),
				},
				Default: int64default.StaticInt64(2),
				Validators: []validator.Int64{
					int64validator.OneOf(1, 2),
				},
			},
			"pepper_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the pepper the secret was hashed with.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"write_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the token permits write operations.",
				PlanModifiers: []planmodifier.Bool{
					boolplanmodifier.UseStateForUnknown(),
				},
			},
			"expires": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Expiry as an RFC 3339 timestamp with milliseconds (2030-01-01T00:00:00.000Z); never expires when unset.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_used": schema.StringAttribute{
				Computed:    true,
				Description: "When the token was last used.",
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *tokenResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

func (r *tokenResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data tokenResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	priorToken := plan.Token
	requestDTO, diags := expandToken(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Users.UsersTokensCreateContext(ctx, users.NewUsersTokensCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_token", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenToken(ctx, netboxapi.TokenResponseDTOFromGoNetbox(res.Payload), plan)...)
	if plan.Token.IsNull() {
		plan.Token = priorToken
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *tokenResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data tokenResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	priorToken := state.Token
	res, err := r.client.Users.UsersTokensRetrieveContext(ctx, users.NewUsersTokensRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_token", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenToken(ctx, netboxapi.TokenResponseDTOFromGoNetbox(res.Payload), state)...)
	if state.Token.IsNull() {
		state.Token = priorToken
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *tokenResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data tokenResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	priorToken := plan.Token
	requestDTO, diags := expandToken(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	delete(payload, "token")
	delete(payload, "version")
	res, err := r.client.Users.UsersTokensPartialUpdateContext(ctx, users.NewUsersTokensPartialUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_token", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenToken(ctx, netboxapi.TokenResponseDTOFromGoNetbox(res.Payload), plan)...)
	if plan.Token.IsNull() {
		plan.Token = priorToken
	}

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *tokenResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data tokenResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Users.UsersTokensDestroyContext(ctx, users.NewUsersTokensDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_token", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *tokenResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Expected an integer id, got %q.", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
