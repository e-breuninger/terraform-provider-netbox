package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Companion hooks of netbox_mac_address: device_interface_id and virtual_machine_interface_id are
// aliases of the assigned_object_type/_id pair (see object_ref.go).

func (model *macAddressResourceModel) objectRef() objectRef {
	return objectRef{
		typeName: "assigned_object_type", idName: "assigned_object_id",
		typ: &model.AssignedObjectType, id: &model.AssignedObjectID,
		aliases: []objectRefAlias{
			{name: "device_interface_id", contentType: "dcim.interface", id: &model.DeviceInterfaceID},
			{name: "virtual_machine_interface_id", contentType: "virtualization.vminterface", id: &model.VirtualMachineInterfaceID},
		},
	}
}

func (resource *macAddressResource) modifyPlanHook(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan macAddressResourceModel
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

func (resource *macAddressResource) preCreateHook(ctx context.Context, model *macAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *macAddressResource) preUpdateHook(ctx context.Context, model *macAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *macAddressResource) postCreateHook(ctx context.Context, model *macAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *macAddressResource) postReadHook(ctx context.Context, model *macAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *macAddressResource) postUpdateHook(ctx context.Context, model *macAddressResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}
