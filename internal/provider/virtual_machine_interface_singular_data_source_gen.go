// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualMachineInterfaceDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualMachineInterfaceDataSource)(nil)
)

// NewVirtualMachineInterfaceDataSource returns a new virtual_machine_interface data source.
func NewVirtualMachineInterfaceDataSource() datasource.DataSource {
	return &virtualMachineInterfaceDataSource{}
}

type virtualMachineInterfaceDataSource struct {
	client *netboxapi.Client
}

// virtualMachineInterfaceDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type virtualMachineInterfaceDataSourceModel struct {
	virtualMachineInterfaceResourceModel
	ClusterID             types.Int64  `tfsdk:"cluster_id"`
	MacAddress            types.String `tfsdk:"mac_address"`
	VirtualMachineName    types.String `tfsdk:"virtual_machine_name"`
	ClusterName           types.String `tfsdk:"cluster_name"`
	VrfRd                 types.String `tfsdk:"vrf_rd"`
	TagSlug               types.String `tfsdk:"tag_slug"`
	L2vpn                 types.Int64  `tfsdk:"l2vpn"`
	L2vpnID               types.Int64  `tfsdk:"l2vpn_id"`
	OwnerName             types.String `tfsdk:"owner_name"`
	OwnerGroupName        types.String `tfsdk:"owner_group_name"`
	OwnerGroupID          types.Int64  `tfsdk:"owner_group_id"`
	PrimaryMacAddress     types.String `tfsdk:"primary_mac_address"`
	PrimaryMacAddressID   types.Int64  `tfsdk:"primary_mac_address_id"`
	Q                     types.String `tfsdk:"q"`
	TagID                 types.Int64  `tfsdk:"tag_id"`
	Vlan                  types.String `tfsdk:"vlan"`
	VlanID                types.String `tfsdk:"vlan_id"`
	VlanTranslationPolicy types.String `tfsdk:"vlan_translation_policy"`
	CustomFieldFilters    types.Map    `tfsdk:"custom_field_filters"`
}

func (d *virtualMachineInterfaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_machine_interface"
}

func (d *virtualMachineInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Virtualization:A virtual machine interface (virtualization.vminterface).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/virtualization/#interfaces):\n\n> Virtual machine interfaces behave similarly to device interfaces, and can be assigned to VRFs, and may have IP addresses, VLANs, and services attached to them. However, given their virtual nature, they lack properties pertaining to physical attributes. For example, VM interfaces do not have a physical type and cannot have cables attached to them.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"virtual_machine_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the virtual machine.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the interface is enabled.",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU in bytes. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(1, 65536),
				},
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "802.1Q tagging mode. One of: access, tagged, tagged-all, q-in-q.",
				Validators: []validator.String{
					stringvalidator.OneOf("access", "tagged", "tagged-all", "q-in-q"),
				},
			},
			"untagged_vlan_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the untagged VLAN.",
			},
			"tagged_vlan_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the tagged VLANs.",
			},
			"qinq_svlan_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the Q-in-Q service VLAN.",
			},
			"vlan_translation_policy_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VLAN translation policy.",
			},
			"vrf_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VRF.",
			},
			"parent_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the parent interface.",
			},
			"bridge_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the bridged interface.",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"tags": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of the tags assigned to the object (the provider's default_tags are added on top, see tags_all).",
			},
			"tags_all": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of all tags on the object, including the provider's default_tags.",
			},
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"cluster_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the cluster of the virtual machine.",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Description: "MAC address assigned to the interface.",
			},
			"virtual_machine_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the virtual machine.",
			},
			"cluster_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the cluster of the virtual machine.",
			},
			"vrf_rd": schema.StringAttribute{
				Optional:    true,
				Description: "Route distinguisher of the VRF.",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"l2vpn": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the L2VPN.",
			},
			"l2vpn_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the L2VPN.",
			},
			"owner_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner.",
			},
			"owner_group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner group.",
			},
			"owner_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner group.",
			},
			"primary_mac_address": schema.StringAttribute{
				Optional:    true,
				Description: "Primary MAC address.",
			},
			"primary_mac_address_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the primary MAC address.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"vlan": schema.StringAttribute{
				Optional:    true,
				Description: "Assigned VID.",
			},
			"vlan_id": schema.StringAttribute{
				Optional:    true,
				Description: "Assigned VLAN.",
			},
			"vlan_translation_policy": schema.StringAttribute{
				Optional:    true,
				Description: "VLAN Translation Policy.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *virtualMachineInterfaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualMachineInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualMachineInterfaceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := virtualization.NewVirtualizationInterfacesListParams()
	hasInput := false
	queryParams := url.Values{}
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
	if !state.VirtualMachineID.IsNull() {
		v := state.VirtualMachineID.ValueInt64()
		params.SetVirtualMachineID([]int64{v})
		hasInput = true
	}
	if !state.ClusterID.IsNull() {
		v := state.ClusterID.ValueInt64()
		params.SetClusterID([]int64{v})
		hasInput = true
	}
	if !state.MacAddress.IsNull() {
		v := state.MacAddress.ValueString()
		params.SetMacAddress([]string{v})
		hasInput = true
	}
	if !state.Enabled.IsNull() {
		v := state.Enabled.ValueBool()
		params.SetEnabled(&v)
		hasInput = true
	}
	if !state.Mode.IsNull() {
		v := state.Mode.ValueString()
		params.SetMode([]string{v})
		hasInput = true
	}
	if !state.Mtu.IsNull() {
		v := state.Mtu.ValueInt64()
		params.SetMtu([]int64{v})
		hasInput = true
	}
	if !state.VrfID.IsNull() {
		v := state.VrfID.ValueInt64()
		params.SetVrfID([]int64{v})
		hasInput = true
	}
	if !state.ParentID.IsNull() {
		v := state.ParentID.ValueInt64()
		params.SetParentID([]int64{v})
		hasInput = true
	}
	if !state.BridgeID.IsNull() {
		v := state.BridgeID.ValueInt64()
		params.SetBridgeID([]int64{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.VirtualMachineName.IsNull() {
		v := state.VirtualMachineName.ValueString()
		params.SetVirtualMachine([]string{v})
		hasInput = true
	}
	if !state.ClusterName.IsNull() {
		v := state.ClusterName.ValueString()
		params.SetCluster([]string{v})
		hasInput = true
	}
	if !state.VrfRd.IsNull() {
		v := state.VrfRd.ValueString()
		params.SetVrf([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.L2vpn.IsNull() {
		v := state.L2vpn.ValueInt64()
		params.SetL2vpn([]int64{v})
		hasInput = true
	}
	if !state.L2vpnID.IsNull() {
		v := state.L2vpnID.ValueInt64()
		params.SetL2vpnID([]int64{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.OwnerName.IsNull() {
		v := state.OwnerName.ValueString()
		params.SetOwner([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupName.IsNull() {
		v := state.OwnerGroupName.ValueString()
		params.SetOwnerGroup([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupID.IsNull() {
		v := state.OwnerGroupID.ValueInt64()
		params.SetOwnerGroupID([]int64{v})
		hasInput = true
	}
	if !state.PrimaryMacAddress.IsNull() {
		v := state.PrimaryMacAddress.ValueString()
		params.SetPrimaryMacAddress([]string{v})
		hasInput = true
	}
	if !state.PrimaryMacAddressID.IsNull() {
		v := state.PrimaryMacAddressID.ValueInt64()
		params.SetPrimaryMacAddressID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.Vlan.IsNull() {
		v := state.Vlan.ValueString()
		params.SetVlan(&v)
		hasInput = true
	}
	if !state.VlanID.IsNull() {
		v := state.VlanID.ValueString()
		params.SetVlanID(&v)
		hasInput = true
	}
	if !state.VlanTranslationPolicy.IsNull() {
		v := state.VlanTranslationPolicy.ValueString()
		params.SetVlanTranslationPolicy([]string{v})
		hasInput = true
	}
	if !state.VlanTranslationPolicyID.IsNull() {
		v := state.VlanTranslationPolicyID.ValueInt64()
		params.SetVlanTranslationPolicyID([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_virtual_machine_interface.")
		return
	}
	res, err := d.client.Virtualization.VirtualizationInterfacesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_virtual_machine_interface", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_virtual_machine_interface",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVirtualMachineInterface(ctx, netboxapi.VirtualMachineInterfaceResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.virtualMachineInterfaceResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
