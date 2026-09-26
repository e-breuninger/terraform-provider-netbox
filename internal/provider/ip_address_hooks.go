package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Companion hooks of netbox_ip_address: device_interface_id and virtual_machine_interface_id are
// aliases of the assigned_object_type/_id pair (see object_ref.go).

func (model *ipAddressResourceModel) objectRef() objectRef {
	return objectRef{
		typeName: "assigned_object_type", idName: "assigned_object_id",
		typ: &model.AssignedObjectType, id: &model.AssignedObjectID,
		aliases: []objectRefAlias{
			{name: "device_interface_id", contentType: "dcim.interface", id: &model.DeviceInterfaceID},
			{name: "virtual_machine_interface_id", contentType: "virtualization.vminterface", id: &model.VirtualMachineInterfaceID},
		},
	}
}

func (resource *ipAddressResource) modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan ipAddressResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.objectRef().plan(config.objectRef(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	// assigned_object follows the pair; when the pair changes it is unknown until NetBox answers.
	if !req.State.Raw.IsNull() {
		var state ipAddressResourceModel
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if plan.objectRef().changed(state.objectRef()) {
			plan.AssignedObject = types.ObjectUnknown(ipAddressAssignedObjectAttrTypes())
		}
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (resource *ipAddressResource) preCreate(ctx context.Context, model *ipAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *ipAddressResource) preUpdate(ctx context.Context, model *ipAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *ipAddressResource) postCreate(ctx context.Context, model *ipAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *ipAddressResource) postRead(ctx context.Context, model *ipAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *ipAddressResource) postUpdate(ctx context.Context, model *ipAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}
