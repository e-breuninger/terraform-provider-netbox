package provider

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"sync"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// netbox_primary_ip is a hand-written companion resource: NetBox only accepts a primary IP
// that is assigned to one of the device's (or VM's) own interfaces, so device -> interface ->
// ip_address -> device.primary_ip is a cycle in Terraform. This resource breaks it by PATCHing
// primary_ip4 or primary_ip6 (by the address family) on the device or virtual machine after the
// address exists. The device's primary_ip4_id/primary_ip6_id are computed and reflect it on the
// next refresh.

func init() {
	companionResources = append(companionResources, NewPrimaryIPResource)
}

var (
	_ resource.Resource                = (*primaryIPResource)(nil)
	_ resource.ResourceWithConfigure   = (*primaryIPResource)(nil)
	_ resource.ResourceWithImportState = (*primaryIPResource)(nil)
)

func NewPrimaryIPResource() resource.Resource {
	return &primaryIPResource{}
}

type primaryIPResource struct {
	client *netboxapi.Client
}

type primaryIPResourceModel struct {
	ID               types.String `tfsdk:"id"`
	DeviceID         types.Int64  `tfsdk:"device_id"`
	VirtualMachineID types.Int64  `tfsdk:"virtual_machine_id"`
	IPAddressID      types.Int64  `tfsdk:"ip_address_id"`
	Family           types.Int64  `tfsdk:"ip_address_version"`
}

func (resource *primaryIPResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_primary_ip"
}

func (resource *primaryIPResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Virtualization:Sets the primary IPv4 or IPv6 address of a device or virtual machine. The address must be assigned to one of its interfaces; the family is taken from the address. Import as device/<id> or vm/<id>.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "device/<device id> or vm/<virtual machine id>.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device. Conflicts with virtual_machine_id.",
				Validators: []validator.Int64{
					int64validator.ExactlyOneOf(path.MatchRoot("device_id"), path.MatchRoot("virtual_machine_id")),
				},
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"virtual_machine_id": schema.Int64Attribute{
				Optional:      true,
				Description:   "Id of the virtual machine. Conflicts with device_id.",
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"ip_address_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the IP address to make primary. Changing it to an address of the other family clears the previous primary.",
			},
			"ip_address_version": schema.Int64Attribute{
				Computed:    true,
				Description: "IP version of the primary address (4 or 6).",
			},
		},
	}
}

func (resource *primaryIPResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// primaryTarget is the device or VM whose primary IP is managed.
type primaryTarget struct {
	kind string // "device" or "vm"
	id   int64
}

func (primaryTarget primaryTarget) String() string {
	return primaryTarget.kind + "/" + strconv.FormatInt(primaryTarget.id, 10)
}

func parsePrimaryTarget(s string) (primaryTarget, error) {
	kind, rest, ok := strings.Cut(s, "/")
	if !ok || (kind != "device" && kind != "vm") {
		return primaryTarget{}, fmt.Errorf("expected device/<id> or vm/<id>, got %q", s)
	}
	id, err := strconv.ParseInt(rest, 10, 64)
	if err != nil {
		return primaryTarget{}, fmt.Errorf("expected device/<id> or vm/<id>, got %q", s)
	}
	return primaryTarget{kind: kind, id: id}, nil
}

func (model *primaryIPResourceModel) target() primaryTarget {
	if !model.DeviceID.IsNull() {
		return primaryTarget{kind: "device", id: model.DeviceID.ValueInt64()}
	}
	return primaryTarget{kind: "vm", id: model.VirtualMachineID.ValueInt64()}
}

func (model *primaryIPResourceModel) setTarget(primaryTarget primaryTarget) {
	model.ID = types.StringValue(primaryTarget.String())
	model.DeviceID, model.VirtualMachineID = types.Int64Null(), types.Int64Null()
	if primaryTarget.kind == "device" {
		model.DeviceID = types.Int64Value(primaryTarget.id)
	} else {
		model.VirtualMachineID = types.Int64Value(primaryTarget.id)
	}
}

// readPrimaries returns the target's current primary_ip4 and primary_ip6 ids.
func (resource *primaryIPResource) readPrimaries(ctx context.Context, primaryTarget primaryTarget) (ip4, ip6 *int64, err error) {
	id := func(briefIPAddress *models.BriefIPAddress) *int64 {
		if briefIPAddress == nil {
			return nil
		}
		v := briefIPAddress.ID
		return &v
	}
	if primaryTarget.kind == "device" {
		res, err := resource.client.Dcim.DcimDevicesRetrieveContext(ctx, dcim.NewDcimDevicesRetrieveParams().WithID(primaryTarget.id), nil)
		if err != nil {
			return nil, nil, err
		}
		return id(res.Payload.PrimaryIp4), id(res.Payload.PrimaryIp6), nil
	}
	res, err := resource.client.Virtualization.VirtualizationVirtualMachinesRetrieveContext(ctx, virtualization.NewVirtualizationVirtualMachinesRetrieveParams().WithID(primaryTarget.id), nil)
	if err != nil {
		return nil, nil, err
	}
	return id(res.Payload.PrimaryIp4), id(res.Payload.PrimaryIp6), nil
}

// patchPrimaries PATCHes only the given primary_ip4/primary_ip6 fields (nil clears); go-netbox's
// writable models have required fields without omitempty, so the body is written directly.
func (resource *primaryIPResource) patchPrimaries(ctx context.Context, primaryTarget primaryTarget, fields map[string]*int64) error {
	body := map[string]any{}
	for field, value := range fields {
		if value == nil {
			body[field] = nil
		} else {
			body[field] = *value
		}
	}
	if primaryTarget.kind == "device" {
		return patchDevice(ctx, resource.client, primaryTarget.id, body)
	}
	params := virtualization.NewVirtualizationVirtualMachinesPartialUpdateParams().WithID(primaryTarget.id)
	_, err := resource.client.Virtualization.VirtualizationVirtualMachinesPartialUpdateContext(ctx, params, nil, netboxapi.WithBody(body))
	return err
}

// deviceLocks serializes PATCHes per device id: NetBox applies a PATCH as read-modify-write, so
// companions (primary IP, OOB IP, device tag) writing the same device in parallel would let the
// last one overwrite the other's field with its stale null.
var deviceLocks sync.Map

// patchDevice PATCHes exactly the given wire fields on a device: go-netbox's writable model has
// required fields without omitempty, so the body is written directly.
func patchDevice(ctx context.Context, client *netboxapi.Client, id int64, body map[string]any) error {
	mutex, _ := deviceLocks.LoadOrStore(id, &sync.Mutex{})
	mutex.(*sync.Mutex).Lock()
	defer mutex.(*sync.Mutex).Unlock()
	params := dcim.NewDcimDevicesPartialUpdateParams().WithID(id)
	_, err := client.Dcim.DcimDevicesPartialUpdateContext(ctx, params, nil, netboxapi.WithBody(body))
	return err
}

func (resource *primaryIPResource) ipFamily(ctx context.Context, ipID int64) (int64, error) {
	res, err := resource.client.Ipam.IpamIPAddressesRetrieveContext(ctx, ipam.NewIpamIPAddressesRetrieveParams().WithID(ipID), nil)
	if err != nil {
		return 0, err
	}
	if res.Payload.Family == nil || res.Payload.Family.Value == 0 {
		return 0, fmt.Errorf("ip address %d has no family", ipID)
	}
	return res.Payload.Family.Value, nil
}

func familyField(family int64) string {
	if family == 6 {
		return "primary_ip6"
	}
	return "primary_ip4"
}

func (resource *primaryIPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan primaryIPResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	if resp.Diagnostics.HasError() {
		return
	}
	primaryTarget := plan.target()
	ipID := plan.IPAddressID.ValueInt64()
	family, err := resource.ipFamily(ctx, ipID)
	if err != nil {
		resp.Diagnostics.AddError("Error reading the IP address for netbox_primary_ip", err.Error())
		return
	}
	if err := resource.patchPrimaries(ctx, primaryTarget, map[string]*int64{familyField(family): &ipID}); err != nil {
		resp.Diagnostics.AddError("Error setting the primary IP of "+primaryTarget.String(), err.Error())
		return
	}
	plan.setTarget(primaryTarget)
	plan.Family = types.Int64Value(family)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *primaryIPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state primaryIPResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	primaryTarget, err := parsePrimaryTarget(state.ID.ValueString())
	if err != nil {
		resp.Diagnostics.AddError("Invalid netbox_primary_ip id", err.Error())
		return
	}
	ip4, ip6, err := resource.readPrimaries(ctx, primaryTarget)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading "+primaryTarget.String(), err.Error())
		return
	}
	// The family in state picks which primary this resource owns; after an import (family unknown)
	// whichever is set wins, IPv4 first.
	var current *int64
	family := state.Family.ValueInt64()
	switch {
	case family == 4:
		current = ip4
	case family == 6:
		current = ip6
	case ip4 != nil:
		current, family = ip4, 4
	case ip6 != nil:
		current, family = ip6, 6
	}
	if current == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.setTarget(primaryTarget)
	state.IPAddressID = types.Int64Value(*current)
	state.Family = types.Int64Value(family)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (resource *primaryIPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state primaryIPResourceModel
	resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	primaryTarget := plan.target()
	ipID := plan.IPAddressID.ValueInt64()
	family, err := resource.ipFamily(ctx, ipID)
	if err != nil {
		resp.Diagnostics.AddError("Error reading the IP address for netbox_primary_ip", err.Error())
		return
	}
	fields := map[string]*int64{familyField(family): &ipID}
	if prior := state.Family.ValueInt64(); prior != 0 && prior != family {
		fields[familyField(prior)] = nil
	}
	if err := resource.patchPrimaries(ctx, primaryTarget, fields); err != nil {
		resp.Diagnostics.AddError("Error setting the primary IP of "+primaryTarget.String(), err.Error())
		return
	}
	plan.setTarget(primaryTarget)
	plan.Family = types.Int64Value(family)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (resource *primaryIPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state primaryIPResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
	if resp.Diagnostics.HasError() {
		return
	}
	primaryTarget := state.target()
	err := resource.patchPrimaries(ctx, primaryTarget, map[string]*int64{familyField(state.Family.ValueInt64()): nil})
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error clearing the primary IP of "+primaryTarget.String(), err.Error())
	}
}

func (resource *primaryIPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	primaryTarget, err := parsePrimaryTarget(req.ID)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", err.Error())
		return
	}
	var model primaryIPResourceModel
	model.setTarget(primaryTarget)
	model.IPAddressID = types.Int64Null()
	model.Family = types.Int64Null()
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
