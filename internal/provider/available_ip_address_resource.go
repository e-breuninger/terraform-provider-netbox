package provider

import (
	"context"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// Companion of netbox_available_ip_address (spec operations.create = "manual"): create allocates
// the next free address of a prefix or IP range through NetBox's available-ips endpoint, then
// configures the address with the same PATCH the generated Update sends. The allocation endpoint
// only accepts the VRF; everything else lands in that second call. Read, update and delete are the
// generated ordinary ip-addresses operations.
//
// The assigned-object aliases (device_interface_id, virtual_machine_interface_id) are handled like
// in ip_address_hooks.go; create applies the pair conversion itself because the pre/post hooks
// only wrap generated bodies.
//
// The prefix_id/ip_range_id choice is enforced by the spec's exactly_one_of, not here.

func (resource *availableIPAddressResource) create(ctx context.Context, plan *availableIPAddressResourceModel, diags *diag.Diagnostics) {
	// Allocate. The body is a list; NetBox answers with a list of the addresses it created.
	item := map[string]any{}
	if !plan.VrfID.IsNull() && !plan.VrfID.IsUnknown() {
		item["vrf"] = plan.VrfID.ValueInt64()
	}
	body := []map[string]any{item}
	var allocated []*models.IPAddress
	switch {
	case !plan.PrefixID.IsNull():
		res, err := resource.client.Ipam.IpamPrefixesAvailableIpsCreateContext(ctx,
			ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(plan.PrefixID.ValueInt64()),
			nil, netboxapi.WithBody(body))
		if err != nil {
			diags.AddError(fmt.Sprintf("Error allocating an IP address from prefix %d", plan.PrefixID.ValueInt64()), err.Error())
			return
		}
		allocated = res.Payload
	default:
		res, err := resource.client.Ipam.IpamIPRangesAvailableIpsCreateContext(ctx,
			ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(plan.IPRangeID.ValueInt64()),
			nil, netboxapi.WithBody(body))
		if err != nil {
			diags.AddError(fmt.Sprintf("Error allocating an IP address from IP range %d", plan.IPRangeID.ValueInt64()), err.Error())
			return
		}
		allocated = res.Payload
	}
	if len(allocated) == 0 || allocated[0].Address == nil {
		diags.AddError("Error allocating an IP address", "NetBox returned no available address.")
		return
	}
	plan.ID = types.Int64Value(allocated[0].ID)
	plan.IPAddress = types.StringValue(*allocated[0].Address)
	// vrf_id unset means the VRF of the prefix or range, which the allocation reports; the configure
	// PATCH sends it back rather than null, which would move the address into the global table.
	if plan.VrfID.IsNull() || plan.VrfID.IsUnknown() {
		plan.VrfID = types.Int64Null()
		if allocated[0].Vrf != nil {
			plan.VrfID = types.Int64Value(allocated[0].Vrf.ID)
		}
	}

	// Configure: PATCH the allocated address with the planned attributes.
	plan.objectRef().toPair()
	configuredTags := conv.SetTo[string](ctx, plan.Tags, diags)
	tagsNull := plan.Tags.IsNull()
	req, diag := expandAvailableIPAddress(ctx, plan)
	diags.Append(diag...)
	if diags.HasError() {
		return
	}
	payload := req.Payload()
	if err := netboxapi.TypedCustomFields(ctx, resource.client, payload); err != nil {
		diags.AddError("Error encoding netbox_available_ip_address custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, resource.client.DefaultTags)
	payload["address"] = plan.IPAddress.ValueString()
	res, err := resource.client.Ipam.IpamIPAddressesPartialUpdateContext(ctx,
		ipam.NewIpamIPAddressesPartialUpdateParams().WithID(plan.ID.ValueInt64()),
		nil, netboxapi.WithBody(payload))
	if err != nil {
		// A failed create leaves no state behind, so free the allocation again rather than leak it.
		resource.rollback(ctx, plan, diags)
		diags.AddError(fmt.Sprintf("Error configuring the allocated IP address %s", plan.IPAddress.ValueString()), err.Error())
		return
	}
	diags.Append(flattenAvailableIPAddress(ctx, netboxapi.AvailableIPAddressResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, diags)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, diags)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, resource.client.DefaultTags), tagsNull, diags)
	plan.objectRef().fromPair()
}

// rollback deletes the address create allocated when configuring it failed.
func (resource *availableIPAddressResource) rollback(ctx context.Context, plan *availableIPAddressResourceModel, diags *diag.Diagnostics) {
	_, err := resource.client.Ipam.IpamIPAddressesDestroyContext(ctx, ipam.NewIpamIPAddressesDestroyParams().WithID(plan.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		diags.AddWarning(fmt.Sprintf("The allocated IP address %s (id %d) could not be released", plan.IPAddress.ValueString(), plan.ID.ValueInt64()), err.Error())
	}
}

func (model *availableIPAddressResourceModel) objectRef() objectRef {
	return objectRef{
		typeName: "assigned_object_type", idName: "assigned_object_id",
		typ: &model.AssignedObjectType, id: &model.AssignedObjectID,
		aliases: []objectRefAlias{
			{name: "device_interface_id", contentType: "dcim.interface", id: &model.DeviceInterfaceID},
			{name: "virtual_machine_interface_id", contentType: "virtualization.vminterface", id: &model.VirtualMachineInterfaceID},
		},
	}
}

func (resource *availableIPAddressResource) modifyPlanHook(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan availableIPAddressResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.objectRef().plan(config.objectRef(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	if !req.State.Raw.IsNull() {
		var state availableIPAddressResourceModel
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if plan.objectRef().changed(state.objectRef()) {
			plan.AssignedObject = types.ObjectUnknown(availableIPAddressAssignedObjectAttrTypes())
		}
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (resource *availableIPAddressResource) preUpdateHook(ctx context.Context, model *availableIPAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *availableIPAddressResource) postReadHook(ctx context.Context, model *availableIPAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *availableIPAddressResource) postUpdateHook(ctx context.Context, model *availableIPAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}
