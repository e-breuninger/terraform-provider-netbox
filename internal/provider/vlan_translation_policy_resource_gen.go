// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/int64planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ resource.Resource                = (*vlanTranslationPolicyResource)(nil)
	_ resource.ResourceWithConfigure   = (*vlanTranslationPolicyResource)(nil)
	_ resource.ResourceWithImportState = (*vlanTranslationPolicyResource)(nil)
)

// NewVlanTranslationPolicyResource returns a new vlan_translation_policy resource.
func NewVlanTranslationPolicyResource() resource.Resource {
	return &vlanTranslationPolicyResource{}
}

type vlanTranslationPolicyResource struct {
	client *netboxapi.Client
}

func (r *vlanTranslationPolicyResource) Metadata(ctx context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_translation_policy"
}

func (r *vlanTranslationPolicyResource) Schema(ctx context.Context, req resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox VLAN translation policy (ipam.vlantranslationpolicy): a named set of VLAN translation rules.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/vlantranslationpolicy/):\n\n> VLAN translation is a feature that consists of VLAN translation policies and [VLAN translation rules](https://netboxlabs.com/docs/netbox/models/ipam/vlantranslationrule/). Many rules can belong to a policy, and each rule defines a mapping of a local to remote VLAN ID (VID). A policy can then be assigned to an [Interface](https://netboxlabs.com/docs/netbox/models/dcim/interface/) or [VMInterface](https://netboxlabs.com/docs/netbox/models/virtualization/vminterface/), and all VLAN translation rules associated with that policy will be visible in the interface details.\n>\n> There are uniqueness constraints on `(policy, local_vid)` and on `(policy, remote_vid)` in the `VLANTranslationRule` model. Thus, you cannot have multiple rules linked to the same policy that have the same local VID or the same remote VID. A set of policies and rules might look like this:\n>\n> Policy 1:\n> - Rule: 100 -> 200\n> - Rule: 101 -> 201\n>\n> Policy 2:\n> - Rule: 100 -> 300\n> - Rule: 101 -> 301\n>\n> However this is not allowed:\n>\n> Policy 3:\n> - Rule: 100 -> 200\n> - Rule: 100 -> 300",
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
					stringvalidator.LengthBetween(1, 100),
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
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner the object is assigned to.",
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

func (r *vlanTranslationPolicyResource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
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

func (r *vlanTranslationPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var data vlanTranslationPolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	// convert to DTO
	requestDTO, diags := expandVlanTranslationPolicy(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	// DTO to payload
	payload := requestDTO.Payload()

	// Send request
	res, err := r.client.Ipam.IpamVlanTranslationPoliciesCreateContext(ctx, ipam.NewIpamVlanTranslationPoliciesCreateParams(), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error creating netbox_vlan_translation_policy", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenVlanTranslationPolicy(ctx, netboxapi.VlanTranslationPolicyResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *vlanTranslationPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var data vlanTranslationPolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	res, err := r.client.Ipam.IpamVlanTranslationPoliciesRetrieveContext(ctx, ipam.NewIpamVlanTranslationPoliciesRetrieveParams().WithID(state.ID.ValueInt64()), nil)
	if netboxapi.IsNotFound(err) {
		resp.State.RemoveResource(ctx)
		return
	}
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_vlan_translation_policy", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenVlanTranslationPolicy(ctx, netboxapi.VlanTranslationPolicyResponseDTOFromGoNetbox(res.Payload), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}

func (r *vlanTranslationPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var data vlanTranslationPolicyResourceModel

	// Read Terraform plan data into the model
	resp.Diagnostics.Append(req.Plan.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	plan := &data

	requestDTO, diags := expandVlanTranslationPolicy(ctx, plan)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	payload := requestDTO.Payload()
	res, err := r.client.Ipam.IpamVlanTranslationPoliciesUpdateContext(ctx, ipam.NewIpamVlanTranslationPoliciesUpdateParams().WithID(plan.ID.ValueInt64()), nil, netboxapi.WithBody(payload))
	if err != nil {
		resp.Diagnostics.AddError("Error updating netbox_vlan_translation_policy", err.Error())
		return
	}
	resp.Diagnostics.Append(flattenVlanTranslationPolicy(ctx, netboxapi.VlanTranslationPolicyResponseDTOFromGoNetbox(res.Payload), plan)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, plan)...)
}

func (r *vlanTranslationPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var data vlanTranslationPolicyResourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.State.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	_, err := r.client.Ipam.IpamVlanTranslationPoliciesDestroyContext(ctx, ipam.NewIpamVlanTranslationPoliciesDestroyParams().WithID(state.ID.ValueInt64()), nil)
	if err != nil && !netboxapi.IsNotFound(err) {
		resp.Diagnostics.AddError("Error deleting netbox_vlan_translation_policy", err.Error())
		return
	}

	if resp.Diagnostics.HasError() {
		return
	}
}

func (r *vlanTranslationPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
