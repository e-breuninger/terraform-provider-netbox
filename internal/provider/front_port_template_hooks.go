package provider

import (
	"context"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// Companion hooks of netbox_front_port_template: rear_port_id and rear_port_position are the flat form of
// the rear_ports set, and updates go through portMappings.beforeUpdate (see port_mappings.go).

func (model *frontPortTemplateResourceModel) portMappings() portMappings {
	return portMappings{
		positions: &model.Positions, set: &model.RearPorts,
		id: &model.RearPortID, pos: &model.RearPortPosition,
		attrTypes: frontPortTemplateRearPortsAttrTypes(),
	}
}

func (resource *frontPortTemplateResource) modifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan frontPortTemplateResourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
	resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan.portMappings().plan(config.portMappings(), &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
}

func (resource *frontPortTemplateResource) preCreate(ctx context.Context, model *frontPortTemplateResourceModel, diag *diag.Diagnostics) {
	model.portMappings().toSet()
}

func (resource *frontPortTemplateResource) preUpdate(ctx context.Context, model *frontPortTemplateResourceModel, diag *diag.Diagnostics) {
	portMappings := model.portMappings()
	portMappings.toSet()
	id := model.ID.ValueInt64()
	res, err := resource.client.Dcim.DcimFrontPortTemplatesRetrieveContext(ctx, dcim.NewDcimFrontPortTemplatesRetrieveParams().WithID(id), nil)
	if err != nil {
		diag.AddError("Error reading netbox_front_port_template", err.Error())
		return
	}
	current := make([]portMapping, 0, len(res.Payload.RearPorts))
	for _, rearPort := range res.Payload.RearPorts {
		if rearPort != nil && rearPort.Position != nil && rearPort.RearPort != nil {
			current = append(current, portMapping{*rearPort.Position, *rearPort.RearPort, rearPort.RearPortPosition})
		}
	}
	if !portMappings.beforeUpdate(current) {
		return
	}
	params := dcim.NewDcimFrontPortTemplatesPartialUpdateParams().WithID(id)
	if _, err := resource.client.Dcim.DcimFrontPortTemplatesPartialUpdateContext(ctx, params, nil, netboxapi.WithBody(clearMappingsBody())); err != nil {
		diag.AddError("Error clearing the rear port mappings of netbox_front_port_template", err.Error())
	}
}

func (resource *frontPortTemplateResource) postCreate(ctx context.Context, model *frontPortTemplateResourceModel, diag *diag.Diagnostics) {
	model.portMappings().fromSet()
}

func (resource *frontPortTemplateResource) postRead(ctx context.Context, model *frontPortTemplateResourceModel, diag *diag.Diagnostics) {
	model.portMappings().fromSet()
}

func (resource *frontPortTemplateResource) postUpdate(ctx context.Context, model *frontPortTemplateResourceModel, diag *diag.Diagnostics) {
	model.portMappings().fromSet()
}
