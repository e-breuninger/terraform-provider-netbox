package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Companion hooks of netbox_vpn_tunnel_termination: device_interface_id and
// virtual_machine_interface_id are aliases of the termination_type/_id pair (see object_ref.go).
// The assignment is mandatory: NetBox refuses a termination without one.

func (model *vpnTunnelTerminationResourceModel) objectRef() objectRef {
	return objectRef{
		typeName: "termination_type", idName: "termination_id",
		typ: &model.TerminationType, id: &model.TerminationID,
		aliases: []objectRefAlias{
			{name: "device_interface_id", contentType: "dcim.interface", id: &model.DeviceInterfaceID},
			{name: "virtual_machine_interface_id", contentType: "virtualization.vminterface", id: &model.VirtualMachineInterfaceID},
		},
		required: true,
	}
}

func (resource *vpnTunnelTerminationResource) modifyPlanHook(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan vpnTunnelTerminationResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.objectRef().plan(config.objectRef(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (resource *vpnTunnelTerminationResource) preCreateHook(ctx context.Context, model *vpnTunnelTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *vpnTunnelTerminationResource) preUpdateHook(ctx context.Context, model *vpnTunnelTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *vpnTunnelTerminationResource) postCreateHook(ctx context.Context, model *vpnTunnelTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *vpnTunnelTerminationResource) postReadHook(ctx context.Context, model *vpnTunnelTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *vpnTunnelTerminationResource) postUpdateHook(ctx context.Context, model *vpnTunnelTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}
