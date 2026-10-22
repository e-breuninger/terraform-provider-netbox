package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// netbox_device_interface_primary_mac_address and netbox_virtual_machine_interface_primary_mac_address
// are hand-written companion resources, for the same reason as netbox_primary_ip: NetBox only
// accepts a primary MAC address that is assigned to the interface itself, so interface -> mac_address
// -> interface.primary_mac_address is a cycle in Terraform. Each breaks it by PATCHing
// primary_mac_address on the interface after the address exists. One implementation serves both;
// the kind picks the interface endpoint and the name of the interface attribute.

func init() {
	companionResources = append(companionResources,
		func() resource.Resource { return &interfacePrimaryMACResource{kind: primaryMACDevice} },
		func() resource.Resource { return &interfacePrimaryMACResource{kind: primaryMACVM} },
	)
}

var (
	_ resource.Resource                = (*interfacePrimaryMACResource)(nil)
	_ resource.ResourceWithConfigure   = (*interfacePrimaryMACResource)(nil)
	_ resource.ResourceWithImportState = (*interfacePrimaryMACResource)(nil)
)

// primaryMACKind is the flavour of interface a primary MAC resource manages.
type primaryMACKind struct {
	typeSuffix    string // resource type name after the provider prefix
	interfaceAttr string // name of the interface id attribute
	noun          string // for messages
	subcategory   string // registry docs subcategory, e.g. "Virtualization"
}

var (
	primaryMACDevice = primaryMACKind{"_device_interface_primary_mac_address", "device_interface_id", "device interface", "Data Center Inventory Management (DCIM)"}
	primaryMACVM     = primaryMACKind{"_virtual_machine_interface_primary_mac_address", "virtual_machine_interface_id", "virtual machine interface", "Virtualization"}
)

type interfacePrimaryMACResource struct {
	kind   primaryMACKind
	client *netboxapi.Client
}

// attrSource is the part of a plan or state this resource reads: tfsdk.Plan and tfsdk.State both
// satisfy it. The model is read and written attribute by attribute rather than through tfsdk tags,
// because the interface attribute's name differs per kind.
type attrSource interface {
	GetAttribute(context.Context, path.Path, any) diag.Diagnostics
}

type attrSink interface {
	SetAttribute(context.Context, path.Path, any) diag.Diagnostics
}

type interfacePrimaryMACModel struct {
	id          types.String
	interfaceID types.Int64
	macID       types.Int64
}

func (resource *interfacePrimaryMACResource) readModel(ctx context.Context, from attrSource) (interfacePrimaryMACModel, diag.Diagnostics) {
	var model interfacePrimaryMACModel
	var diags diag.Diagnostics
	diags.Append(from.GetAttribute(ctx, path.Root("id"), &model.id)...)
	diags.Append(from.GetAttribute(ctx, path.Root(resource.kind.interfaceAttr), &model.interfaceID)...)
	diags.Append(from.GetAttribute(ctx, path.Root("mac_address_id"), &model.macID)...)
	return model, diags
}

func (resource *interfacePrimaryMACResource) writeModel(ctx context.Context, model interfacePrimaryMACModel, sink attrSink) diag.Diagnostics {
	var diags diag.Diagnostics
	diags.Append(sink.SetAttribute(ctx, path.Root("id"), model.id)...)
	diags.Append(sink.SetAttribute(ctx, path.Root(resource.kind.interfaceAttr), model.interfaceID)...)
	diags.Append(sink.SetAttribute(ctx, path.Root("mac_address_id"), model.macID)...)
	return diags
}

func (resource *interfacePrimaryMACResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + resource.kind.typeSuffix
}

func (resource *interfacePrimaryMACResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: fmt.Sprintf(":meta:subcategory:%s:Sets the primary MAC address of a %s. The address must be assigned to that interface (netbox_mac_address with %s). Import with the interface id.", resource.kind.subcategory, resource.kind.noun, resource.kind.interfaceAttr),
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "The interface id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			resource.kind.interfaceAttr: schema.Int64Attribute{
				Required:      true,
				Description:   fmt.Sprintf("Id of the %s.", resource.kind.noun),
				PlanModifiers: []planmodifier.Int64{int64planmodifier.RequiresReplace()},
			},
			"mac_address_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the MAC address to make primary; it must be assigned to the interface.",
			},
		},
	}
}

func (resource *interfacePrimaryMACResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

// readPrimary returns the interface's current primary MAC address id, nil when it has none.
func (resource *interfacePrimaryMACResource) readPrimary(ctx context.Context, interfaceID int64) (*int64, error) {
	id := func(briefMACAddress *models.BriefMACAddress) *int64 {
		if briefMACAddress == nil {
			return nil
		}
		v := briefMACAddress.ID
		return &v
	}
	if resource.kind == primaryMACDevice {
		res, err := resource.client.Dcim.DcimInterfacesRetrieveContext(ctx, dcim.NewDcimInterfacesRetrieveParams().WithID(interfaceID), nil)
		if err != nil {
			return nil, err
		}
		return id(res.Payload.PrimaryMacAddress), nil
	}
	res, err := resource.client.Virtualization.VirtualizationInterfacesRetrieveContext(ctx, virtualization.NewVirtualizationInterfacesRetrieveParams().WithID(interfaceID), nil)
	if err != nil {
		return nil, err
	}
	return id(res.Payload.PrimaryMacAddress), nil
}

// patchPrimary PATCHes only primary_mac_address (nil clears); go-netbox's writable interface models
// have required fields without omitempty, so the body is written directly.
func (resource *interfacePrimaryMACResource) patchPrimary(ctx context.Context, interfaceID int64, macID *int64) error {
	body := map[string]any{"primary_mac_address": nil}
	if macID != nil {
		body["primary_mac_address"] = *macID
	}
	if resource.kind == primaryMACDevice {
		params := dcim.NewDcimInterfacesPartialUpdateParams().WithID(interfaceID)
		_, err := resource.client.Dcim.DcimInterfacesPartialUpdateContext(ctx, params, nil, netboxapi.WithBody(body))
		return err
	}
	params := virtualization.NewVirtualizationInterfacesPartialUpdateParams().WithID(interfaceID)
	_, err := resource.client.Virtualization.VirtualizationInterfacesPartialUpdateContext(ctx, params, nil, netboxapi.WithBody(body))
	return err
}

// apply sets the planned MAC as primary and fills in the id.
func (resource *interfacePrimaryMACResource) apply(ctx context.Context, model interfacePrimaryMACModel) (interfacePrimaryMACModel, error) {
	macID := model.macID.ValueInt64()
	if err := resource.patchPrimary(ctx, model.interfaceID.ValueInt64(), &macID); err != nil {
		return model, err
	}
	model.id = types.StringValue(strconv.FormatInt(model.interfaceID.ValueInt64(), 10))
	return model, nil
}

func (resource *interfacePrimaryMACResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	plan, diags := resource.readModel(ctx, req.Plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan, err := resource.apply(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Error setting the primary MAC address of "+resource.kind.noun+" "+plan.interfaceID.String(), err.Error())
		return
	}
	resp.Diagnostics.Append(resource.writeModel(ctx, plan, &resp.State)...)
}

func (resource *interfacePrimaryMACResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	state, diags := resource.readModel(ctx, req.State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	interfaceID, err := strconv.ParseInt(state.id.ValueString(), 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid id", fmt.Sprintf("expected an interface id, got %q", state.id.ValueString()))
		return
	}
	current, err := resource.readPrimary(ctx, interfaceID)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading "+resource.kind.noun+" "+state.id.ValueString(), err.Error())
		return
	}
	if current == nil {
		resp.State.RemoveResource(ctx)
		return
	}
	state.interfaceID = types.Int64Value(interfaceID)
	state.macID = types.Int64Value(*current)
	resp.Diagnostics.Append(resource.writeModel(ctx, state, &resp.State)...)
}

func (resource *interfacePrimaryMACResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	plan, diags := resource.readModel(ctx, req.Plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan, err := resource.apply(ctx, plan)
	if err != nil {
		resp.Diagnostics.AddError("Error setting the primary MAC address of "+resource.kind.noun+" "+plan.interfaceID.String(), err.Error())
		return
	}
	resp.Diagnostics.Append(resource.writeModel(ctx, plan, &resp.State)...)
}

func (resource *interfacePrimaryMACResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	state, diags := resource.readModel(ctx, req.State)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	err := resource.patchPrimary(ctx, state.interfaceID.ValueInt64(), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error clearing the primary MAC address of "+resource.kind.noun+" "+state.interfaceID.String(), err.Error())
	}
}

func (resource *interfacePrimaryMACResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	interfaceID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError("Invalid import ID", fmt.Sprintf("expected the %s id, got %q", resource.kind.noun, req.ID))
		return
	}
	model := interfacePrimaryMACModel{
		id:          types.StringValue(req.ID),
		interfaceID: types.Int64Value(interfaceID),
		macID:       types.Int64Null(),
	}
	resp.Diagnostics.Append(resource.writeModel(ctx, model, &resp.State)...)
}
