package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
)

// Companion hooks of netbox_circuit_termination: site_id, location_id, region_id, site_group_id
// and provider_network_id are aliases of the termination_type/_id pair (see object_ref.go). The
// termination is mandatory: NetBox refuses one without it.

func (model *circuitTerminationResourceModel) objectRef() objectRef {
	return objectRef{
		typeName: "termination_type", idName: "termination_id",
		typ: &model.TerminationType, id: &model.TerminationID,
		aliases: []objectRefAlias{
			{name: "site_id", contentType: "dcim.site", id: &model.SiteID},
			{name: "location_id", contentType: "dcim.location", id: &model.LocationID},
			{name: "region_id", contentType: "dcim.region", id: &model.RegionID},
			{name: "site_group_id", contentType: "dcim.sitegroup", id: &model.SiteGroupID},
			{name: "provider_network_id", contentType: "circuits.providernetwork", id: &model.ProviderNetworkID},
		},
		required: true,
	}
}

func (resource *circuitTerminationResource) modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan circuitTerminationResourceModel
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

func (resource *circuitTerminationResource) preCreate(ctx context.Context, model *circuitTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *circuitTerminationResource) preUpdate(ctx context.Context, model *circuitTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().toPair()
}

func (resource *circuitTerminationResource) postCreate(ctx context.Context, model *circuitTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *circuitTerminationResource) postRead(ctx context.Context, model *circuitTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}

func (resource *circuitTerminationResource) postUpdate(ctx context.Context, model *circuitTerminationResourceModel, diag *diag.Diagnostics) {
	model.objectRef().fromPair()
}
