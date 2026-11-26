package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var (
	_ resource.ResourceWithConfigValidators = (*tokenResource)(nil)
	_ resource.ResourceWithValidateConfig   = (*tokenResource)(nil)
	_ resource.ResourceWithModifyPlan       = (*tokenResource)(nil)
)

type writeOnlyTokenResourceModel struct {
	tokenResourceModel
	TokenWo        types.String `tfsdk:"token_wo"`
	TokenWoVersion types.Int64  `tfsdk:"token_wo_version"`
}

func (r *tokenResource) ConfigValidators(context.Context) []resource.ConfigValidator {
	return []resource.ConfigValidator{
		resourcevalidator.Conflicting(path.MatchRoot("token"), path.MatchRoot("token_wo")),
		resourcevalidator.RequiredTogether(path.MatchRoot("token_wo"), path.MatchRoot("token_wo_version")),
	}
}

func (r *tokenResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var config writeOnlyTokenResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	if !resp.Diagnostics.HasError() && !config.TokenWoVersion.IsNull() && config.Version.Equal(types.Int64Value(1)) {
		resp.Diagnostics.AddAttributeError(path.Root("version"), "Write-only token requires v2", "Set version to 2 or omit it when using token_wo.")
	}
}

func (r *tokenResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var version types.Int64
	resp.Diagnostics.Append(req.Plan.GetAttribute(ctx, path.Root("token_wo_version"), &version)...)
	if !resp.Diagnostics.HasError() && !version.IsNull() {
		// The legacy computed attribute must not carry a prior token into a write-only plan.
		resp.Diagnostics.Append(resp.Plan.SetAttribute(ctx, path.Root("token"), types.StringNull())...)
	}
}

func (model *writeOnlyTokenResourceModel) restoreToken(prior types.String) {
	if !model.TokenWoVersion.IsNull() {
		// NetBox can echo the secret in its create response. Drop it before writing state.
		model.Token = types.StringNull()
	} else if model.Token.IsNull() {
		model.Token = prior
	}
}
