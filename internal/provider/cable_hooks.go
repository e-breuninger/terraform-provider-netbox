package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Companion hooks of netbox_cable: the *_ids children of a_side and b_side are per-type aliases of
// the side's object_type and ids (see cable_side.go).

func (model *cableResourceModel) sideA() cableSide {
	return cableSide{name: "a_side", obj: &model.ASide, attrTypes: cableASideAttrTypes()}
}

func (model *cableResourceModel) sideB() cableSide {
	return cableSide{name: "b_side", obj: &model.BSide, attrTypes: cableBSideAttrTypes()}
}

func (resource *cableResource) modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan cableResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.sideA().plan(config.sideA(), &resp.Diagnostics)
	plan.sideB().plan(config.sideB(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (resource *cableResource) postCreate(ctx context.Context, model *cableResourceModel, diag *diag.Diagnostics) {
	model.sideA().fromWire()
	model.sideB().fromWire()
}

func (resource *cableResource) postRead(ctx context.Context, model *cableResourceModel, diag *diag.Diagnostics) {
	model.sideA().fromWire()
	model.sideB().fromWire()
}

func (resource *cableResource) postUpdate(ctx context.Context, model *cableResourceModel, diag *diag.Diagnostics) {
	model.sideA().fromWire()
	model.sideB().fromWire()
}
