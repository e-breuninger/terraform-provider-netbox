// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                 = (*virtualMachineResource)(nil)
	_ resource.ResourceWithConfigure    = (*virtualMachineResource)(nil)
	_ resource.ResourceWithImportState  = (*virtualMachineResource)(nil)
	_ resource.ResourceWithUpgradeState = (*virtualMachineResource)(nil)
	_ resource.ResourceWithModifyPlan   = (*virtualMachineResource)(nil)
)

// NewVirtualMachineResource returns a new virtual_machine resource.
func NewVirtualMachineResource() resource.Resource {
	return &virtualMachineResource{}
}

type virtualMachineResource struct {
	client *netboxapi.Client
}

func (r *virtualMachineResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_machine"
}

// UpgradeState upgrades any prior schema version by keeping only the id; the refresh that
// follows a plan repopulates every other attribute from the API. That deliberately also accepts
// state this provider never wrote — e.g. a migration from another provider via
// `terraform state replace-provider` — whose attribute layout is unknown here.
func (r *virtualMachineResource) UpgradeState(ctx context.Context) map[int64]resource.StateUpgrader {
	up := resource.StateUpgrader{
		StateUpgrader: func(ctx context.Context, req resource.UpgradeStateRequest, resp *resource.UpgradeStateResponse) {
			var prior struct {
				ID any `json:"id"`
			}
			if req.RawState == nil || json.Unmarshal(req.RawState.JSON, &prior) != nil || prior.ID == nil {
				resp.Diagnostics.AddError("Unable to upgrade the prior state", "The prior state carries no id.")
				return
			}
			var id int64
			switch v := prior.ID.(type) {
			case float64:
				id = int64(v)
			case string:
				parsed, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					resp.Diagnostics.AddError("Unable to upgrade the prior state", fmt.Sprintf("The prior state's id %q is not a number.", v))
					return
				}
				id = parsed
			default:
				resp.Diagnostics.AddError("Unable to upgrade the prior state", fmt.Sprintf("The prior state's id has unexpected type %T.", v))
				return
			}
			var sr resource.SchemaResponse
			r.Schema(ctx, resource.SchemaRequest{}, &sr)
			obj := conv.IDOnlyObject(ctx, sr.Schema.Type().(types.ObjectType), "id", types.Int64Value(id))
			resp.Diagnostics.Append(resp.State.Set(ctx, obj)...)
		},
	}
	out := map[int64]resource.StateUpgrader{}
	for v := int64(0); v < 3; v++ {
		out[v] = up
	}
	return out
}

func (r *virtualMachineResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Version:     3,
		Description: ":meta:subcategory:Virtualization:A virtual machine (virtualization.virtualmachine).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#virtual-machines):\n\n> A virtual machine is a virtualized compute instance. These behave in NetBox very similarly to device objects, but without any physical attributes. For example, a VM may have interfaces assigned to it with IP addresses and VLANs, however its interfaces cannot be connected via cables (because they are virtual). Each VM may also define its compute, memory, and storage resources as well.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"name": schema.StringAttribute{
				Required: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: offline, active, planned, staged, failed, decommissioning, paused.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("offline", "active", "planned", "staged", "failed", "decommissioning", "paused"),
				},
			},
			"cluster_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster (a virtual machine needs a cluster or a site).",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site (a virtual machine needs a cluster or a site).",
			},
			"role_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the functional device role.",
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant.",
			},
			"platform_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the platform.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device hosting the VM.",
			},
			"primary_ip4_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the primary IPv4 address (managed by netbox_primary_ip).",
			},
			"primary_ip6_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the primary IPv6 address (managed by netbox_primary_ip).",
			},
			"vcpus": schema.Float64Attribute{
				Optional:    true,
				Description: "Number of virtual CPUs (fractions allowed).",
				Validators: []validator.Float64{
					float64validator.Between(0.01, 10000),
				},
			},
			"memory_mb": schema.Int64Attribute{
				Optional:    true,
				Description: "Memory in MB.",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"disk_size_mb": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Disk size in MB. Once the machine has virtual disks NetBox computes it as their sum and it can no longer be set; leave it unset then.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
				},
			},
			"local_context_data": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Optional:    true,
				Description: "Local config context data as JSON text (use jsonencode()).",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Optional: true,
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner the object is assigned to.",
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
			"interface_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of interfaces on the virtual machine.",
			},
			"virtual_disk_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of virtual disks on the virtual machine.",
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Description: "Slugs of the tags assigned to the object (the provider's default_tags are added on top, see tags_all).",
			},
			"tags_all": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of all tags on the object, including the provider's default_tags.",
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Optional:    true,
				Computed:    true,
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
				PlanModifiers: []planmodifier.Map{
					mapplanmodifier.UseStateForUnknown(),
				},
			},
		},
	}
}

func (r *virtualMachineResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *virtualMachineResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var plan virtualMachineResourceModel
		resp.Diagnostics.Append(req.Plan.Get(ctx, &plan)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if r.client != nil && !plan.Tags.IsUnknown() {
			configured := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
			plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, netboxapi.MergeTags(configured, r.client.DefaultTags), false, &resp.Diagnostics)
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}()
	if resp.Diagnostics.HasError() {
		return
	}
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state virtualMachineResourceModel
		resp.Diagnostics.Append(req.Config.Get(ctx, &config)...)
		resp.Diagnostics.Append(resp.Plan.Get(ctx, &plan)...)
		resp.Diagnostics.Append(req.State.Get(ctx, &state)...)
		if resp.Diagnostics.HasError() {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = state.LastUpdated
		}
		if config.InterfaceCount.IsNull() {
			plan.InterfaceCount = state.InterfaceCount
		}
		if config.VirtualDiskCount.IsNull() {
			plan.VirtualDiskCount = state.VirtualDiskCount
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
		if resp.Diagnostics.HasError() || resp.Plan.Raw.Equal(req.State.Raw) {
			return
		}
		if config.LastUpdated.IsNull() {
			plan.LastUpdated = types.StringUnknown()
		}
		if config.InterfaceCount.IsNull() {
			plan.InterfaceCount = types.Int64Unknown()
		}
		if config.VirtualDiskCount.IsNull() {
			plan.VirtualDiskCount = types.Int64Unknown()
		}
		resp.Diagnostics.Append(resp.Plan.Set(ctx, &plan)...)
	}
}

func (r *virtualMachineResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data virtualMachineResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandVirtualMachine(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// beforeRequest
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_virtual_machine custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)

	// Send request
	res, err := r.client.Virtualization.VirtualizationVirtualMachinesCreateContext(ctx, virtualization.NewVirtualizationVirtualMachinesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_virtual_machine", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenVirtualMachine(ctx, netboxapi.VirtualMachineResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *virtualMachineResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data virtualMachineResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Virtualization.VirtualizationVirtualMachinesRetrieveContext(ctx, virtualization.NewVirtualizationVirtualMachinesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_virtual_machine", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenVirtualMachine(ctx, netboxapi.VirtualMachineResponseDTOFromGoNetbox(res.Payload), state)...)
	allTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	state.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	state.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	state.CustomFields = keepPriorCustomFields(priorCustomFields, state.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *virtualMachineResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data virtualMachineResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandVirtualMachine(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior virtualMachineResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_virtual_machine custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	res, err := r.client.Virtualization.VirtualizationVirtualMachinesUpdateContext(ctx, virtualization.NewVirtualizationVirtualMachinesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_virtual_machine", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenVirtualMachine(ctx, netboxapi.VirtualMachineResponseDTOFromGoNetbox(res.Payload), plan)...)
	allTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	if allTags == nil {
		allTags = []string{}
	}
	plan.TagsAll = conv.SetFrom[string](ctx, types.StringType, allTags, false, &resp.Diagnostics)
	plan.Tags = conv.SetFrom[string](ctx, types.StringType, netboxapi.ConfiguredTags(allTags, configuredTags, r.client.DefaultTags), tagsNull, &resp.Diagnostics)
	plan.CustomFields = keepPriorCustomFields(priorCustomFields, plan.CustomFields)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *virtualMachineResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data virtualMachineResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Virtualization.VirtualizationVirtualMachinesDestroyContext(ctx, virtualization.NewVirtualizationVirtualMachinesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_virtual_machine", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *virtualMachineResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
