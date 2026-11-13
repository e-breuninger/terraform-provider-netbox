package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// Companion of netbox_token: NetBox 4.7 made a token's plaintext read-only. It generates every
// secret itself and ignores one the client sends ("API token plaintexts can no longer be specified
// by the client when creating a token via the REST API", 4.7.0 release notes), so a configured
// token came back different and Terraform failed the apply with "inconsistent values for
// sensitive attribute" after the token had been created. On 4.7 and later a configured token is
// refused at plan time instead, before anything is created; up to 4.6 it is sent as before.

var _ resource.ResourceWithModifyPlan = (*tokenResource)(nil)

// tokenChosenSecretMinVersion is the first NetBox release that ignores a client-chosen secret.
var tokenChosenSecretMinVersion = [2]int{4, 7}

func (r *tokenResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() || r.client == nil {
		return // a destroy, or the provider is not configured yet (validate)
	}
	var token types.String
	resp.Diagnostics.Append(req.Config.GetAttribute(ctx, path.Root("token"), &token)...)
	if resp.Diagnostics.HasError() || token.IsNull() || token.IsUnknown() {
		return
	}
	if !req.State.Raw.IsNull() {
		// An existing token whose secret is unchanged is not created again: nothing to refuse.
		var prior types.String
		resp.Diagnostics.Append(req.State.GetAttribute(ctx, path.Root("token"), &prior)...)
		if resp.Diagnostics.HasError() || prior.Equal(token) {
			return
		}
	}
	_, version, err := netboxapi.Version(ctx, r.client.NetBoxAPI)
	if err != nil || !versionAtLeast(version, tokenChosenSecretMinVersion) {
		// Unknown version: the provider's own version check already reported it; send as before.
		return
	}
	resp.Diagnostics.AddAttributeError(
		path.Root("token"),
		"NetBox generates the token's secret",
		fmt.Sprintf("NetBox %s generates every API token's secret and ignores one chosen by the client (since NetBox 4.7). "+
			"Remove token from this netbox_token: the secret NetBox generates is in the state after the apply.", version),
	)
}

// versionAtLeast reports whether a MAJOR.MINOR.PATCH version is at least min (major, minor).
// A version it cannot parse counts as older.
func versionAtLeast(version string, min [2]int) bool {
	parts := strings.SplitN(version, ".", 3)
	if len(parts) < 2 {
		return false
	}
	major, err1 := strconv.Atoi(parts[0])
	minor, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		return false
	}
	return major > min[0] || (major == min[0] && minor >= min[1])
}
