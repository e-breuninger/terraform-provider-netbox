// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*vlanTranslationRuleDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlanTranslationRuleDataSource)(nil)
)

// NewVlanTranslationRuleDataSource returns a new vlan_translation_rule data source.
func NewVlanTranslationRuleDataSource() datasource.DataSource {
	return &vlanTranslationRuleDataSource{}
}

type vlanTranslationRuleDataSource struct {
	client *netboxapi.Client
}

func (d *vlanTranslationRuleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_translation_rule"
}

func (d *vlanTranslationRuleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A single VLAN id mapping inside a translation policy (ipam.vlantranslationrule).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/vlantranslationrule/):\n\n> A VLAN translation rule represents a one-to-one mapping of a local VLAN ID (VID) to a remote VID. Many rules can belong to a single policy.\n>\n> See [VLAN translation policies](https://netboxlabs.com/docs/netbox/models/ipam/vlantranslationpolicy/) for an overview of the VLAN Translation feature.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"vlan_translation_policy_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the translation policy the rule belongs to.",
			},
			"local_vid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN id on the local side (1-4094).",
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"remote_vid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "VLAN id on the remote side (1-4094).",
				Validators: []validator.Int64{
					int64validator.Between(1, 4094),
				},
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *vlanTranslationRuleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vlanTranslationRuleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlanTranslationRuleResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamVlanTranslationRulesListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.VlanTranslationPolicyID.IsNull() {
		v := state.VlanTranslationPolicyID.ValueInt64()
		params.SetPolicyID([]int64{v})
		hasInput = true
	}
	if !state.LocalVid.IsNull() {
		v := state.LocalVid.ValueInt64()
		params.SetLocalVid([]int64{v})
		hasInput = true
	}
	if !state.RemoteVid.IsNull() {
		v := state.RemoteVid.ValueInt64()
		params.SetRemoteVid([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or vlan_translation_policy_id and/or local_vid and/or remote_vid to look up a netbox_vlan_translation_rule.")
		return
	}
	res, err := d.client.Ipam.IpamVlanTranslationRulesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_vlan_translation_rule", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_vlan_translation_rule",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVlanTranslationRule(ctx, netboxapi.VlanTranslationRuleResponseDTOFromGoNetbox(res.Payload.Results[0]), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
