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

// Companion of netbox_available_asn (spec operations.create = "manual"): create allocates the next
// free AS number of an ASN range through NetBox's available-asns endpoint. Like the available-vlans
// endpoint and unlike available-ips, this one takes the whole object, so a single call both
// allocates and configures. Read, update and delete are the generated ordinary asns operations.
//
// The range is the path parameter, and the AS number and its RIR are what NetBox picks from that
// range, so none of the three is in the body: asn_range_id is a derived attribute the request DTO
// never writes, and asn and rir are computed, which keeps them out of the request body. They come
// back on updates instead, through the resend the spec asks for.

func (resource *availableAsnResource) create(ctx context.Context, plan *availableAsnResourceModel, diags *diag.Diagnostics) {
	configuredTags := conv.SetTo[string](ctx, plan.Tags, diags)
	tagsNull := plan.Tags.IsNull()
	req, diag := expandAvailableAsn(ctx, plan)
	diags.Append(diag...)
	if diags.HasError() {
		return
	}
	payload := req.Payload()
	if err := netboxapi.TypedCustomFields(ctx, resource.client, payload); err != nil {
		diags.AddError("Error encoding netbox_available_asn custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, resource.client.DefaultTags)
	// An unset optional string is sent as "" so an update clears it, but there is nothing to clear
	// on create and NetBox's allocation serializer rejects a blank description outright. Leave the
	// ones the plan does not set out of the allocation body.
	if plan.Description.IsNull() {
		delete(payload, "description")
	}
	if plan.Comments.IsNull() {
		delete(payload, "comments")
	}

	rangeID := plan.AsnRangeID.ValueInt64()
	// The body is a list, and NetBox answers with a list of the ASNs it created.
	res, err := resource.client.Ipam.IpamAsnRangesAvailableAsnsCreateContext(ctx,
		ipam.NewIpamAsnRangesAvailableAsnsCreateParams().WithID(rangeID),
		nil, netboxapi.WithBody([]map[string]any{payload}))
	if err != nil {
		diags.AddError(fmt.Sprintf("Error allocating an ASN from ASN range %d", rangeID), err.Error())
		return
	}
	if len(res.Payload) == 0 {
		diags.AddError(fmt.Sprintf("Error allocating an ASN from ASN range %d", rangeID),
			"NetBox returned no ASN; the range may be exhausted.")
		return
	}
	diags.Append(flattenAvailableAsn(ctx, netboxapi.AvailableAsnResponseDTOFromGoNetbox(res.Payload[0]), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, diags)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, diags)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, resource.client.DefaultTags), tagsNull, diags)
}
