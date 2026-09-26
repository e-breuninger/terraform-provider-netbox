// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"regexp"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/resourcevalidator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/mapplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringdefault"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                     = (*availableIPAddressResource)(nil)
	_ resource.ResourceWithConfigure        = (*availableIPAddressResource)(nil)
	_ resource.ResourceWithImportState      = (*availableIPAddressResource)(nil)
	_ resource.ResourceWithModifyPlan       = (*availableIPAddressResource)(nil)
	_ resource.ResourceWithConfigValidators = (*availableIPAddressResource)(nil)
)

// NewAvailableIPAddressResource returns a new available_ip_address resource.
func NewAvailableIPAddressResource() resource.Resource {
	return &availableIPAddressResource{}
}

type availableIPAddressResource struct {
	client *netboxapi.Client
}

func (r *availableIPAddressResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_available_ip_address"
}

func (r *availableIPAddressResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox IP address (ipam.ipaddress) allocated from the next free address of a prefix or IP range. After creation it is an ordinary IP address; the allocation inputs replace the resource when changed.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/ipaddress/):\n\n> An IP address object in NetBox comprises a single host address (either IPv4 or IPv6) and its subnet mask, and represents an IP address as configured on a network interface. IP addresses can be assigned to [device](https://netboxlabs.com/docs/netbox/models/dcim/device/) and [virtual machine](https://netboxlabs.com/docs/netbox/models/virtualization/virtualmachine/) interfaces, as well as to [FHRP groups](https://netboxlabs.com/docs/netbox/models/ipam/fhrpgroup/). Further, each device and virtual machine may have one of its interface IPs designated as its primary IP per address family (one for IPv4 and one for IPv6).\n>\n> **Tip:** When primary IPs are set for both IPv4 and IPv6, NetBox will prefer IPv6. This can be changed by setting the `PREFER_IPV4` configuration parameter.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Computed:    true,
				Description: "NetBox id of the IP address.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"prefix_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the prefix to allocate the next available address from. Exactly one of prefix_id and ip_range_id must be set.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIf(func(ctx context.Context, req planmodifier.Int64Request, resp *int64planmodifier.RequiresReplaceIfFuncResponse) {
						// Write-only: null in state means the value is unknown (imported), not different.
						resp.RequiresReplace = !req.StateValue.IsNull() && !req.PlanValue.IsUnknown()
					}, "Replaces the resource when the value changes between two known values; a value missing from state (after import) is adopted instead.", "Replaces the resource when the value changes between two known values; a value missing from state (after import) is adopted instead."),
				},
			},
			"ip_range_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the IP range to allocate the next available address from. Exactly one of prefix_id and ip_range_id must be set.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.RequiresReplaceIf(func(ctx context.Context, req planmodifier.Int64Request, resp *int64planmodifier.RequiresReplaceIfFuncResponse) {
						// Write-only: null in state means the value is unknown (imported), not different.
						resp.RequiresReplace = !req.StateValue.IsNull() && !req.PlanValue.IsUnknown()
					}, "Replaces the resource when the value changes between two known values; a value missing from state (after import) is adopted instead.", "Replaces the resource when the value changes between two known values; a value missing from state (after import) is adopted instead."),
				},
			},
			"ip_address": schema.StringAttribute{
				Computed:    true,
				Description: "The allocated IPv4 or IPv6 address with prefix length (e.g. 10.0.0.1/24).",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: active, reserved, deprecated, dhcp, slaac.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Default: stringdefault.StaticString("active"),
				Validators: []validator.String{
					stringvalidator.OneOf("active", "reserved", "deprecated", "dhcp", "slaac"),
				},
			},
			"role": schema.StringAttribute{
				Optional:    true,
				Description: "Functional role. One of: loopback, secondary, anycast, vip, vrrp, hsrp, glbp, carp.",
				Validators: []validator.String{
					stringvalidator.OneOf("loopback", "secondary", "anycast", "vip", "vrrp", "hsrp", "glbp", "carp"),
				},
			},
			"dns_name": schema.StringAttribute{
				Optional:    true,
				Description: "Hostname or FQDN (not case-sensitive).",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
					stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9A-Za-z_-]+|\*)(\.[0-9A-Za-z_-]+)*\.?$`), ""),
				},
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
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the tenant.",
			},
			"vrf_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VRF.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"nat_inside_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the IP address this address is the NAT (outside) IP for.",
			},
			"assigned_object_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the object the address is assigned to. Derived from device_interface_id or virtual_machine_interface_id when one of those is set; set it together with assigned_object_id for other object types. One of: dcim.interface, virtualization.vminterface, ipam.fhrpgroup.",
				PlanModifiers: []planmodifier.String{
					stringplanmodifier.UseStateForUnknown(),
				},
				Validators: []validator.String{
					stringvalidator.OneOf("dcim.interface", "virtualization.vminterface", "ipam.fhrpgroup"),
				},
			},
			"assigned_object_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the object the address is assigned to (see assigned_object_type).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device interface the address is assigned to (assigned_object_type dcim.interface). Conflicts with virtual_machine_interface_id and with setting the assigned_object_* pair directly.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"virtual_machine_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the virtual machine interface the address is assigned to (assigned_object_type virtualization.vminterface). Conflicts with device_interface_id and with setting the assigned_object_* pair directly.",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
			},
			"assigned_object": schema.SingleNestedAttribute{
				Attributes: map[string]schema.Attribute{
					"id": schema.Int64Attribute{
						Computed: true,
					},
					"name": schema.StringAttribute{
						Computed: true,
					},
					"device": schema.SingleNestedAttribute{
						Attributes: map[string]schema.Attribute{
							"id": schema.Int64Attribute{
								Computed: true,
							},
							"name": schema.StringAttribute{
								Computed: true,
							},
						},
						Computed: true,
					},
				},
				Computed:    true,
				Description: "The interface the address is assigned to. device is null for VM interfaces.",
			},
			"family": schema.Int64Attribute{
				Computed:    true,
				Description: "IP family (4 or 6).",
				PlanModifiers: []planmodifier.Int64{
					int64planmodifier.UseStateForUnknown(),
				},
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

// ConfigValidators enforces the spec's cross-attribute rules (exactly_one_of / conflicts_with /
// required_with) on the configuration, and the ones the companion file's config_validators hook
// (spec hooks) contributes.
func (r *availableIPAddressResource) ConfigValidators(ctx context.Context) []resource.ConfigValidator {
	out := []resource.ConfigValidator{
		resourcevalidator.ExactlyOneOf(path.MatchRoot("prefix_id"), path.MatchRoot("ip_range_id")),
	}
	return out
}

func (r *availableIPAddressResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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
func (r *availableIPAddressResource) ModifyPlan(ctx context.Context, req resource.ModifyPlanRequest, resp *resource.ModifyPlanResponse) {
	// The backend fragment runs in a closure so its bare returns end only the fragment; the
	// companion hook and the stages below still run.
	func() {
		if req.Plan.Raw.IsNull() {
			return
		}
		var plan availableIPAddressResourceModel
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
	// The companion file's modify_plan hook (spec hooks), after the backend fragment.
	r.modifyPlan(ctx, req, resp)
	// Volatile stage, last: the framework decided before any stage above whether the volatile
	// attributes are unknown, from the proposed state. Decide again from the final plan: a plan
	// that changes nothing else keeps the state values (no update runs), any other plan makes
	// them unknown. A configured value always wins.
	if !resp.Diagnostics.HasError() && !req.State.Raw.IsNull() && !resp.Plan.Raw.IsNull() {
		var config, plan, state availableIPAddressResourceModel
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

func (r *availableIPAddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data availableIPAddressResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// Manual create (spec operations.create = "manual"): the companion file defines create.
	r.create(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *availableIPAddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data availableIPAddressResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	configuredTags := conv.SetTo[string](ctx, state.Tags, &resp.Diagnostics)
	tagsNull := state.Tags.IsNull()
	priorCustomFields := state.CustomFields
	res, err := r.client.Ipam.IpamIPAddressesRetrieveContext(ctx, ipam.NewIpamIPAddressesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_available_ip_address", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenAvailableIPAddress(ctx, netboxapi.AvailableIPAddressResponseDTOFromGoNetbox(res.Payload), state)...)
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
	// The companion file's post_read hook (spec hooks).
	r.postRead(ctx, state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *availableIPAddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data availableIPAddressResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data
	// The companion file's pre_update hook (spec hooks).
	r.preUpdate(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	configuredTags := conv.SetTo[string](ctx, plan.Tags, &resp.Diagnostics)
	tagsNull := plan.Tags.IsNull()
	priorCustomFields := plan.CustomFields
	requestDTO, diags := expandAvailableIPAddress(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	var prior availableIPAddressResourceModel
	resp.Diagnostics.Append(req.State.Get(ctx, &prior)...)
	netboxapi.ClearRemovedCustomFields(payload, conv.MapTo[string](ctx, prior.CustomFields, &resp.Diagnostics))
	if err := netboxapi.TypedCustomFields(ctx, r.client, payload); err != nil {
		resp.Diagnostics.AddError("Error encoding netbox_available_ip_address custom fields", err.Error())
		return
	}
	netboxapi.AddDefaultTags(payload, r.client.DefaultTags)
	payload["address"] = plan.IPAddress.ValueString()
	res, err := r.client.Ipam.IpamIPAddressesPartialUpdateContext(ctx, ipam.NewIpamIPAddressesPartialUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_available_ip_address", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenAvailableIPAddress(ctx, netboxapi.AvailableIPAddressResponseDTOFromGoNetbox(res.Payload), plan)...)
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
	// The companion file's post_update hook (spec hooks).
	r.postUpdate(ctx, plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *availableIPAddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data availableIPAddressResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Ipam.IpamIPAddressesDestroyContext(ctx, ipam.NewIpamIPAddressesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_available_ip_address", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *availableIPAddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
