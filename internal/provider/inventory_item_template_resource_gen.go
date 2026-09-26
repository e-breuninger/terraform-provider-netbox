// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*inventoryItemTemplateResource)(nil)
	_ resource.ResourceWithConfigure   = (*inventoryItemTemplateResource)(nil)
	_ resource.ResourceWithImportState = (*inventoryItemTemplateResource)(nil)
	_ resource.ResourceWithModifyPlan  = (*inventoryItemTemplateResource)(nil)
)

// NewInventoryItemTemplateResource returns a new inventory_item_template resource.
func NewInventoryItemTemplateResource() resource.Resource {
	return &inventoryItemTemplateResource{}
}

type inventoryItemTemplateResource struct {
	client *netboxapi.Client
}

func (r *inventoryItemTemplateResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_inventory_item_template"
}

func (r *inventoryItemTemplateResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):An inventory item template of a device type (dcim.inventoryitemtemplate).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/inventoryitemtemplate/):\n\n> **Deprecation Warning:** Beginning in NetBox v4.3, the use of inventory items has been deprecated. They are planned for removal in a future NetBox release. Users are strongly encouraged to begin using [modules](https://netboxlabs.com/docs/netbox/models/dcim/module/) and [module types](https://netboxlabs.com/docs/netbox/models/dcim/moduletype/) in place of inventory items. Modules provide enhanced functionality and can be configured with user-defined attributes.\n>\n> A template for an inventory item that will be automatically created when instantiating a new device. All attributes of this object will be copied to the new inventory item, including the associations with a parent item and assigned component, if any. See the [inventory item](https://netboxlabs.com/docs/netbox/models/dcim/inventoryitem/) documentation for more detail.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"device_type_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the device type. NetBox refuses to move a template to another type.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplace(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"parent_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the parent inventory item template.",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the inventory item role.",
			},
			"manufacturer_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the manufacturer.",
			},
			"part_id": schema.StringAttribute{
				Optional:    true,
				Description: "Manufacturer part number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"label": schema.StringAttribute{
				Optional:    true,
				Description: "Physical label.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(64),
				},
			},
			"component_type": schema.StringAttribute{
				Optional:    true,
				Description: "Type of the component template the item is bound to, e.g. dcim.interfacetemplate. Requires component_id.",
			},
			"component_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the component template named by component_type.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"created": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"last_updated": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *inventoryItemTemplateResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected resource configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	r.client = client
}

// ModifyPlan runs on every plan, after Terraform builds the proposed new state. It is the
// resource-level hook for changing the plan itself: filling computed attributes already known at
// plan time (avoiding "known after apply" noise), or vetoing the plan with diagnostics. It is
// generated when something needs it, each part only where the spec asks for it: a backend plan
// fragment, the companion file's modify_plan hook, the sync that keeps attributes marked as
// deprecated equal to their canonical ones and the stage that decides the volatile attributes.
func (r *inventoryItemTemplateResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state inventoryItemTemplateResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *inventoryItemTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data inventoryItemTemplateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	requestDTO, diags := expandInventoryItemTemplate(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Dcim.DcimInventoryItemTemplatesCreateContext(ctx, dcim.NewDcimInventoryItemTemplatesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_inventory_item_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenInventoryItemTemplate(ctx, netboxapi.InventoryItemTemplateResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *inventoryItemTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data inventoryItemTemplateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	res, err := r.client.Dcim.DcimInventoryItemTemplatesRetrieveContext(ctx, dcim.NewDcimInventoryItemTemplatesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_inventory_item_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenInventoryItemTemplate(ctx, netboxapi.InventoryItemTemplateResponseDTOFromGoNetbox(res.Payload), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *inventoryItemTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data inventoryItemTemplateResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	requestDTO, diags := expandInventoryItemTemplate(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	res, err := r.client.Dcim.DcimInventoryItemTemplatesUpdateContext(ctx, dcim.NewDcimInventoryItemTemplatesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_inventory_item_template", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenInventoryItemTemplate(ctx, netboxapi.InventoryItemTemplateResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *inventoryItemTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data inventoryItemTemplateResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Dcim.DcimInventoryItemTemplatesDestroyContext(ctx, dcim.NewDcimInventoryItemTemplatesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_inventory_item_template", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *inventoryItemTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	id, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		resp.Diagnostics.AddError(
			"Invalid import ID",
			fmt.Sprintf("Expected an integer id, got %q.", req.ID),
		)
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), id)...)
}
