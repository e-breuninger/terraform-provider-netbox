package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Companion hooks of netbox_service: device_id and virtual_machine_id are aliases of the
// parent_object_type/_id pair (see object_ref.go). NetBox 4.6 attaches a service through that
// pair and no longer accepts the device/virtual_machine fields go-netbox still carries; the aliases
// restore the shape users expect. The assignment is required — NetBox rejects a service without a
// parent — so an empty one is a plan-time error rather than a cleared assignment.

func (model *serviceResourceModel) objectRef() objectRef {
	return objectRef{
		typeName: "parent_object_type", idName: "parent_object_id",
		typ: &model.ParentObjectType, id: &model.ParentObjectID,
		aliases: []objectRefAlias{
			{name: "device_id", contentType: "dcim.device", id: &model.DeviceID},
			{name: "virtual_machine_id", contentType: "virtualization.virtualmachine", id: &model.VirtualMachineID},
		},
		required: true,
	}
}

func (resource *serviceResource) modifyPlanHook(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan serviceResourceModel
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

func (resource *serviceResource) preCreateHook(ctx context.Context, model *serviceResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *serviceResource) preUpdateHook(ctx context.Context, model *serviceResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *serviceResource) postCreateHook(ctx context.Context, model *serviceResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *serviceResource) postReadHook(ctx context.Context, model *serviceResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *serviceResource) postUpdateHook(ctx context.Context, model *serviceResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}
