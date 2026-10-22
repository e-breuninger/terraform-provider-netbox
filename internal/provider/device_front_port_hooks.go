package provider

import (
	"context"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// Companion hooks of netbox_device_front_port: rear_port_id and rear_port_position are the flat form of
// the rear_ports set, and updates go through portMappings.beforeUpdate (see port_mappings.go).

func (model *deviceFrontPortResourceModel) portMappings() portMappings {
	return portMappings{
		positions: &model.Positions, set: &model.RearPorts,
		id: &model.RearPortID, pos: &model.RearPortPosition,
		attrTypes: deviceFrontPortRearPortsAttrTypes(),
	}
}

func (resource *deviceFrontPortResource) modifyPlanHook(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	if req.Plan.Raw.IsNull() {
		return
	}
	var config, plan deviceFrontPortResourceModel
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

func (resource *deviceFrontPortResource) preCreateHook(ctx context.Context, model *deviceFrontPortResourceModel, diag *diag.Diagnostics) {
	model.portMappings().toSet()
}

func (resource *deviceFrontPortResource) preUpdateHook(ctx context.Context, model *deviceFrontPortResourceModel, diag *diag.Diagnostics) {
	portMappings := model.portMappings()
	portMappings.toSet()
	id := model.ID.ValueInt64()
	res, err := resource.client.Dcim.DcimFrontPortsRetrieveContext(ctx, dcim.NewDcimFrontPortsRetrieveParams().WithID(id), nil)
	if err != nil {
		diag.AddError("Error reading netbox_device_front_port", err.Error())
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
	params := dcim.NewDcimFrontPortsPartialUpdateParams().WithID(id)
	if _, err := resource.client.Dcim.DcimFrontPortsPartialUpdateContext(ctx, params, nil, netboxapi.WithBody(clearMappingsBody())); err != nil {
		diag.AddError("Error clearing the rear port mappings of netbox_device_front_port", err.Error())
	}
}

func (resource *deviceFrontPortResource) postCreateHook(ctx context.Context, model *deviceFrontPortResourceModel, diag *diag.Diagnostics) {
	model.portMappings().fromSet()
}

func (resource *deviceFrontPortResource) postReadHook(ctx context.Context, model *deviceFrontPortResourceModel, diag *diag.Diagnostics) {
	model.portMappings().fromSet()
}

func (resource *deviceFrontPortResource) postUpdateHook(ctx context.Context, model *deviceFrontPortResourceModel, diag *diag.Diagnostics) {
	model.portMappings().fromSet()
}
