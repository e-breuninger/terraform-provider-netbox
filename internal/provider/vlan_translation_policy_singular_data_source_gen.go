// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*vlanTranslationPolicyDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlanTranslationPolicyDataSource)(nil)
)

// NewVlanTranslationPolicyDataSource returns a new vlan_translation_policy data source.
func NewVlanTranslationPolicyDataSource() datasource.DataSource {
	return &vlanTranslationPolicyDataSource{}
}

type vlanTranslationPolicyDataSource struct {
	client *netboxapi.Client
}

// vlanTranslationPolicyDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type vlanTranslationPolicyDataSourceModel struct {
	vlanTranslationPolicyResourceModel
	NameContains types.String `tfsdk:"name_contains"`
}

func (d *vlanTranslationPolicyDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_translation_policy"
}

func (d *vlanTranslationPolicyDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox VLAN translation policy (ipam.vlantranslationpolicy): a named set of VLAN translation rules.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/vlantranslationpolicy/):\n\n> VLAN translation is a feature that consists of VLAN translation policies and [VLAN translation rules](https://netboxlabs.com/docs/netbox/models/ipam/vlantranslationrule/). Many rules can belong to a policy, and each rule defines a mapping of a local to remote VLAN ID (VID). A policy can then be assigned to an [Interface](https://netboxlabs.com/docs/netbox/models/dcim/interface/) or [VMInterface](https://netboxlabs.com/docs/netbox/models/virtualization/vminterface/), and all VLAN translation rules associated with that policy will be visible in the interface details.\n>\n> There are uniqueness constraints on `(policy, local_vid)` and on `(policy, remote_vid)` in the `VLANTranslationRule` model. Thus, you cannot have multiple rules linked to the same policy that have the same local VID or the same remote VID. A set of policies and rules might look like this:\n>\n> Policy 1:\n> - Rule: 100 -> 200\n> - Rule: 101 -> 201\n>\n> Policy 2:\n> - Rule: 100 -> 300\n> - Rule: 101 -> 301\n>\n> However this is not allowed:\n>\n> Policy 3:\n> - Rule: 100 -> 200\n> - Rule: 100 -> 300",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
			},
		},
	}
}

func (d *vlanTranslationPolicyDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	c, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected data source configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}
	d.client = c
}

func (d *vlanTranslationPolicyDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlanTranslationPolicyDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamVlanTranslationPoliciesListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or name_contains and/or owner_id to look up a netbox_vlan_translation_policy.")
		return
	}
	res, err := d.client.Ipam.IpamVlanTranslationPoliciesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_vlan_translation_policy", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_vlan_translation_policy",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVlanTranslationPolicy(ctx, netboxapi.VlanTranslationPolicyResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.vlanTranslationPolicyResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
