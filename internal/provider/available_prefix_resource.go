package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// Companion of netbox_available_prefix (spec operations create and import = "manual"): create
// allocates the next free child prefix of the requested length through NetBox's
// available-prefixes endpoint, then configures it with the same PATCH the generated Update
// sends (the allocation endpoint only takes the length). Read, update and delete are the generated
// ordinary prefixes operations.
//
// Import takes "parent_prefix_id/prefix_id/prefix_length" because the two allocation inputs are
// required and cannot be recovered from the API.

func (resource *availablePrefixResource) create(ctx context.Context, plan *availablePrefixResourceModel, diags *diag.Diagnostics) {
	length := plan.PrefixLength.ValueInt64()
	if length < 0 || length > 128 {
		diags.AddAttributeError(path.Root("prefix_length"), "Invalid prefix length", fmt.Sprintf("Expected 0-128, got %d.", length))
		return
	}
	res, err := resource.client.Ipam.IpamPrefixesAvailablePrefixesCreateContext(ctx,
		ipam.NewIpamPrefixesAvailablePrefixesCreateParams().WithID(plan.ParentPrefixID.ValueInt64()).WithData([]*models.PrefixLength{{PrefixLength: &length}}),
		nil)
	if err != nil {
		diags.AddError(fmt.Sprintf("Error allocating a /%d prefix from prefix %d", length, plan.ParentPrefixID.ValueInt64()), err.Error())
		return
	}
	// 4.6.9 answers with a list, one entry per prefix asked for; this asks for one.
	if len(res.Payload) == 0 || res.Payload[0] == nil || res.Payload[0].Prefix == nil {
		diags.AddError("Error allocating a prefix", "NetBox returned no available prefix.")
		return
	}
	plan.ID = types.Int64Value(res.Payload[0].ID)
	plan.Prefix = types.StringValue(*res.Payload[0].Prefix)
	// vrf_id unset means the parent's VRF, which the allocation reports; the configure PATCH sends
	// it back rather than null, which would move the prefix into the global table.
	if plan.VrfID.IsNull() || plan.VrfID.IsUnknown() {
		plan.VrfID = types.Int64Null()
		if res.Payload[0].Vrf != nil {
			plan.VrfID = types.Int64Value(res.Payload[0].Vrf.ID)
		}
	}

	configuredTags := conv.SetTo[string](ctx, plan.Tags, diags)
	tagsNull := plan.Tags.IsNull()
	req, diag := expandAvailablePrefix(ctx, plan)
	diags.Append(diag...)
	if diags.HasError() {
		return
	}
	payload := req.Payload()
	if err := netboxapi.TypedCustomFields(ctx, resource.client, payload); err != nil {
		diags.AddError("Error encoding netbox_available_prefix custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, resource.client.DefaultTags)
	payload["prefix"] = plan.Prefix.ValueString()
	upd, err := resource.client.Ipam.IpamPrefixesPartialUpdateContext(ctx,
		ipam.NewIpamPrefixesPartialUpdateParams().WithID(plan.ID.ValueInt64()),
		nil, netboxapi.WithBody(payload))
	if err != nil {
		// A failed create leaves no state behind, so free the allocation again rather than leak it.
		if _, derr := resource.client.Ipam.IpamPrefixesDestroyContext(ctx, ipam.NewIpamPrefixesDestroyParams().WithID(plan.ID.ValueInt64()), nil); derr != nil && !netboxapi.IsNotFound(derr) {
			diags.AddWarning(fmt.Sprintf("The allocated prefix %s (id %d) could not be released", plan.Prefix.ValueString(), plan.ID.ValueInt64()), derr.Error())
		}
		diags.AddError(fmt.Sprintf("Error configuring the allocated prefix %s", plan.Prefix.ValueString()), err.Error())
		return
	}
	diags.Append(flattenAvailablePrefix(ctx, netboxapi.AvailablePrefixResponseDTOFromGoNetbox(upd.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, diags)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, diags)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, resource.client.DefaultTags), tagsNull, diags)
	// The manual create bypasses postCreateHook, so derive the scope aliases here (see
	// available_prefix_hooks.go).
	plan.scopeRef().fromPair()
}

func (resource *availablePrefixResource) importState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	parts := strings.Split(req.ID, "/")
	var ids [3]int64
	if len(parts) == 3 {
		for i, part := range parts {
			v, err := strconv.ParseInt(part, 10, 64)
			if err != nil {
				parts = nil
				break
			}
			ids[i] = v
		}
	}
	if len(parts) != 3 {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Expected \"parent_prefix_id/prefix_id/prefix_length\" (three integers), got %q.", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("parent_prefix_id"), ids[0])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), ids[1])...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("prefix_length"), ids[2])...)
}
