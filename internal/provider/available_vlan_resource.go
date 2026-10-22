package provider

import (
	"context"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// Companion of netbox_available_vlan (spec operations.create = "manual"): create allocates the
// next free VLAN id of a group through NetBox's available-vlans endpoint. Unlike the IP and
// prefix allocation endpoints this one takes the whole object, so a single call creates and
// configures the VLAN. Read, update and delete are the generated ordinary vlans operations.

func (resource *availableVlanResource) create(ctx context.Context, plan *availableVlanResourceModel, diags *diag.Diagnostics) {
	configuredTags := conv.SetTo[string](ctx, plan.Tags, diags)
	tagsNull := plan.Tags.IsNull()
	req, diag := expandAvailableVlan(ctx, plan)
	diags.Append(diag...)
	if diags.HasError() {
		return
	}
	payload := req.Payload()
	// The group is the path parameter, so it does not belong in the body of the allocation call.
	delete(payload, "group")
	if err := netboxapi.TypedCustomFields(ctx, resource.client, payload); err != nil {
		diags.AddError("Error encoding netbox_available_vlan custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, resource.client.DefaultTags)
	res, err := resource.client.Ipam.IpamVlanGroupsAvailableVlansCreateContext(ctx,
		ipam.NewIpamVlanGroupsAvailableVlansCreateParams().WithID(plan.GroupID.ValueInt64()),
		nil, netboxapi.WithBody([]map[string]any{payload}))
	if err != nil {
		diags.AddError(fmt.Sprintf("Error allocating a VLAN from VLAN group %d", plan.GroupID.ValueInt64()), err.Error())
		return
	}
	// 4.6.9 answers with a list, one entry per VLAN asked for; this asks for one.
	if len(res.Payload) == 0 || res.Payload[0] == nil {
		diags.AddError("Error allocating a VLAN", "NetBox returned no VLAN.")
		return
	}
	diags.Append(flattenAvailableVlan(ctx, netboxapi.AvailableVlanResponseDTOFromGoNetbox(res.Payload[0]), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, diags)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, diags)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, resource.client.DefaultTags), tagsNull, diags)
}
