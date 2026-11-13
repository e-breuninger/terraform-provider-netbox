package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// netbox_virtual_chassis_master is the virtual chassis counterpart of netbox_primary_ip: NetBox
// only accepts a master that is already a member of the chassis, and a member references its
// chassis (device -> virtual_chassis -> master device would cycle), so the master is set by a
// separate resource once the members exist. netbox_virtual_chassis reports it as master_id.

func init() {
	companionResources = append(companionResources, NewVirtualChassisMasterResource)
}

var (
	_ resource.Resource                = (*virtualChassisMasterResource)(nil)
	_ resource.ResourceWithConfigure   = (*virtualChassisMasterResource)(nil)
	_ resource.ResourceWithImportState = (*virtualChassisMasterResource)(nil)
)

func NewVirtualChassisMasterResource() resource.Resource {
	return &virtualChassisMasterResource{}
}

type virtualChassisMasterResource struct {
	client *netboxapi.Client
}

type virtualChassisMasterResourceModel struct {
	ID               types.String `tfsdk:"id"`
	VirtualChassisID types.Int64  `tfsdk:"virtual_chassis_id"`
	DeviceID         types.Int64  `tfsdk:"device_id"`
}

func (resource *virtualChassisMasterResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_chassis_master"
}

func (resource *virtualChassisMasterResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Sets the master device of a virtual chassis. The device must be a member of the chassis (its virtual_chassis_id). Import with the virtual chassis id.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The virtual chassis id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"virtual_chassis_id": schema.Int64Attribute{
				Required:      true,
				Description:   "Id of the virtual chassis.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			// Replaced, not updated: NetBox refuses to delete the master device or take it out of
			// the chassis, so the old master has to be cleared before its device changes. A
			// replacement destroys this resource first; an in-place update would come last.
			"device_id": schema.Int64Attribute{
				Required:      true,
				Description:   "Id of the member device to make the chassis master.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
		},
	}
}

func (resource *virtualChassisMasterResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected resource configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	resource.client = client
}

// patchMaster sets the chassis master to the device with the given id, or clears it for nil.
func (resource *virtualChassisMasterResource) patchMaster(ctx context.Context, chassisID int64, deviceID *int64) error {
	params := dcim.NewDcimVirtualChassisPartialUpdateParams().WithID(chassisID)
	_, err := resource.client.Dcim.DcimVirtualChassisPartialUpdateContext(ctx, params, nil,
		netboxapi.WithBody(map[string]any{"master": deviceID}))
	return err
}

func (resource *virtualChassisMasterResource) set(ctx context.Context, plan *virtualChassisMasterResourceModel) error {
	deviceID := plan.DeviceID.ValueInt64()
	if err := resource.patchMaster(ctx, plan.VirtualChassisID.ValueInt64(), &deviceID); err != nil {
		return err
	}
	plan.ID = types.StringValue(strconv.FormatInt(plan.VirtualChassisID.ValueInt64(), 10))
	return nil
}

func (resource *virtualChassisMasterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan virtualChassisMasterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := resource.set(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error setting the master of the virtual chassis", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *virtualChassisMasterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state virtualChassisMasterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	id, err := strconv.ParseInt(state.ID.ValueString(), 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid netbox_virtual_chassis_master id", fmt.Sprintf("Expected the virtual chassis id, got %q.", state.ID.ValueString()))
		return
	}
	res, err := resource.client.Dcim.DcimVirtualChassisRetrieveContext(ctx, dcim.NewDcimVirtualChassisRetrieveParams().WithID(id), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading the virtual chassis", err.Error())
		return
	}
	// A master cleared outside Terraform means the resource no longer exists; the next plan
	// sets it again.
	if res.Payload.Master == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.VirtualChassisID = types.Int64Value(id)
	state.DeviceID = types.Int64Value(res.Payload.Master.ID)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (resource *virtualChassisMasterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan virtualChassisMasterResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	if err := resource.set(ctx, &plan); err != nil {
		resp.Diagnostics.AddError("Error setting the master of the virtual chassis", err.Error())
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *virtualChassisMasterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state virtualChassisMasterResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := resource.patchMaster(ctx, state.VirtualChassisID.ValueInt64(), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error clearing the master of the virtual chassis", err.Error())
	}
}

func (resource *virtualChassisMasterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("Expected the virtual chassis id, got %q.", req.ID))
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &virtualChassisMasterResourceModel{
		ID: types.StringValue(strconv.FormatInt(id, 10)), VirtualChassisID: types.Int64Value(id), DeviceID: types.Int64Null(),
	})...)
}
