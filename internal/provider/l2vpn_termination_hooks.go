package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Companion hooks of netbox_l2vpn_termination: device_interface_id, virtual_machine_interface_id
// and vlan_id are aliases of the assigned_object_type/_id pair (see object_ref.go).

func (model *l2vpnTerminationResourceModel) objectRef() objectRef {
	return objectRef{
		typeName: "assigned_object_type", idName: "assigned_object_id",
		typ: &model.AssignedObjectType, id: &model.AssignedObjectID,
		aliases: []objectRefAlias{
			{name: "device_interface_id", contentType: "dcim.interface", id: &model.DeviceInterfaceID},
			{name: "virtual_machine_interface_id", contentType: "virtualization.vminterface", id: &model.VirtualMachineInterfaceID},
			{name: "vlan_id", contentType: "ipam.vlan", id: &model.VlanID},
		},
	}
}

func (resource *l2vpnTerminationResource) modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan l2vpnTerminationResourceModel
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

func (resource *l2vpnTerminationResource) preCreate(ctx context.Context, model *l2vpnTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *l2vpnTerminationResource) preUpdate(ctx context.Context, model *l2vpnTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *l2vpnTerminationResource) postCreate(ctx context.Context, model *l2vpnTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *l2vpnTerminationResource) postRead(ctx context.Context, model *l2vpnTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *l2vpnTerminationResource) postUpdate(ctx context.Context, model *l2vpnTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}
