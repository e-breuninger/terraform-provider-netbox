// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*vpnTunnelTerminationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vpnTunnelTerminationDataSource)(nil)
)

// NewVpnTunnelTerminationDataSource returns a new vpn_tunnel_termination data source.
func NewVpnTunnelTerminationDataSource() datasource.DataSource {
	return &vpnTunnelTerminationDataSource{}
}

type vpnTunnelTerminationDataSource struct {
	client *netboxapi.Client
}

// vpnTunnelTerminationDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type vpnTunnelTerminationDataSourceModel struct {
	vpnTunnelTerminationResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *vpnTunnelTerminationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vpn_tunnel_termination"
}

func (d *vpnTunnelTerminationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:A tunnel termination attaching an interface to a tunnel (vpn.tunneltermination).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/vpn-tunnels/):\n\n> NetBox can model private tunnels formed among virtual termination points across your network. Typical tunnel implementations include GRE, IP-in-IP, and IPSec. A tunnel may be terminated to two or more device or virtual machine interfaces. For convenient organization, tunnels may be assigned to user-defined groups.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"vpn_tunnel_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tunnel.",
			},
			"role": schema.StringAttribute{
				Computed:    true,
				Description: "Role of the termination in the tunnel topology. One of: peer, hub, spoke.",
			},
			"termination_type": schema.StringAttribute{
				Computed:    true,
				Description: "Content type of the terminating interface. Derived from device_interface_id or virtual_machine_interface_id when one of those is set; set it together with termination_id otherwise. One assignment is required. One of: dcim.interface, virtualization.vminterface.",
			},
			"termination_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the terminating interface (see termination_type).",
			},
			"device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating device interface (termination_type dcim.interface). Conflicts with virtual_machine_interface_id and with setting the termination_* pair directly.",
			},
			"virtual_machine_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating virtual machine interface (termination_type virtualization.vminterface). Conflicts with device_interface_id and with setting the termination_* pair directly.",
			},
			"outside_ip_address_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the outside IP address of the termination.",
			},
			"created": schema.StringAttribute{
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
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
				Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
			},
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *vpnTunnelTerminationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vpnTunnelTerminationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vpnTunnelTerminationDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := vpn.NewVpnTunnelTerminationsListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.VpnTunnelID.IsNull() {
		v := state.VpnTunnelID.ValueInt64()
		params.SetTunnelID([]int64{v})
		hasInput = true
	}
	if !state.DeviceInterfaceID.IsNull() {
		v := state.DeviceInterfaceID.ValueInt64()
		params.SetInterfaceID([]int64{v})
		hasInput = true
	}
	if !state.VirtualMachineInterfaceID.IsNull() {
		v := state.VirtualMachineInterfaceID.ValueInt64()
		params.SetVminterfaceID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or vpn_tunnel_id and/or device_interface_id and/or virtual_machine_interface_id and/or tag_slug and/or custom_field_filters to look up a netbox_vpn_tunnel_termination.")
		return
	}
	res, err := d.client.Vpn.VpnTunnelTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_vpn_tunnel_termination", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_vpn_tunnel_termination",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVpnTunnelTermination(ctx, netboxapi.VpnTunnelTerminationResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.vpnTunnelTerminationResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&vpnTunnelTerminationResource{client: d.client}).postReadHook(ctx, &state.vpnTunnelTerminationResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
