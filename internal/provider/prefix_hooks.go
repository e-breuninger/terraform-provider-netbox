package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Companion hooks of netbox_prefix: site_id, location_id, region_id and site_group_id are aliases
// of the scope_type/scope_id pair (see object_ref.go).

func (model *prefixResourceModel) scopeRef() objectRef {
	return objectRef{
		typeName: "scope_type", idName: "scope_id",
		typ: &model.ScopeType, id: &model.ScopeID,
		aliases: []objectRefAlias{
			{name: "site_id", contentType: "dcim.site", id: &model.SiteID},
			{name: "location_id", contentType: "dcim.location", id: &model.LocationID},
			{name: "region_id", contentType: "dcim.region", id: &model.RegionID},
			{name: "site_group_id", contentType: "dcim.sitegroup", id: &model.SiteGroupID},
		},
	}
}

func (resource *prefixResource) modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan prefixResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.scopeRef().plan(config.scopeRef(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (resource *prefixResource) preCreate(ctx context.Context, model *prefixResourceModel, diag *diag.Diagnostics) {
	model.scopeRef().toPair()
}

func (resource *prefixResource) preUpdate(ctx context.Context, model *prefixResourceModel, diag *diag.Diagnostics) {
	model.scopeRef().toPair()
}

func (resource *prefixResource) postCreate(ctx context.Context, model *prefixResourceModel, diag *diag.Diagnostics) {
	model.scopeRef().fromPair()
}

func (resource *prefixResource) postRead(ctx context.Context, model *prefixResourceModel, diag *diag.Diagnostics) {
	model.scopeRef().fromPair()
}

func (resource *prefixResource) postUpdate(ctx context.Context, model *prefixResourceModel, diag *diag.Diagnostics) {
	model.scopeRef().fromPair()
}
