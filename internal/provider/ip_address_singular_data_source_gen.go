// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
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
	_ datasource.DataSource              = (*ipAddressDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ipAddressDataSource)(nil)
)

// NewIPAddressDataSource returns a new ip_address data source.
func NewIPAddressDataSource() datasource.DataSource {
	return &ipAddressDataSource{}
}

type ipAddressDataSource struct {
	client *netboxapi.Client
}

// ipAddressDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type ipAddressDataSourceModel struct {
	ipAddressResourceModel
	Device             types.String `tfsdk:"device"`
	InterfaceID        types.Int64  `tfsdk:"interface_id"`
	DeviceID           types.Int64  `tfsdk:"device_id"`
	ParentPrefix       types.String `tfsdk:"parent_prefix"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	VrfRd              types.String `tfsdk:"vrf_rd"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *ipAddressDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_address"
}

func (d *ipAddressDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A NetBox IP address (ipam.ipaddress).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/features/ipam/#ip-addresses):\n\n> An IP address comprises a single host address (either IPv4 or IPv6) and its subnet mask. Its mask should match exactly how the IP address is configured on an interface in the real world.\n>\n> Like a prefix, an IP address can optionally be assigned to a VRF (otherwise, it will appear in the \"global\" table). IP addresses are automatically arranged under parent prefixes within their respective VRFs according to the IP hierarchy.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id of the IP address.",
			},
			"ip_address": schema.StringAttribute{
				CustomType:  conv.IPAddressType{},
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address with prefix length (e.g. 10.0.0.1/24).",
				Validators: []validator.String{
					stringvalidator.LengthAtLeast(1),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: active, reserved, deprecated, dhcp, slaac.",
				Validators: []validator.String{
					stringvalidator.OneOf("active", "reserved", "deprecated", "dhcp", "slaac"),
				},
			},
			"role": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Functional role. One of: loopback, secondary, anycast, vip, vrrp, hsrp, glbp, carp.",
				Validators: []validator.String{
					stringvalidator.OneOf("loopback", "secondary", "anycast", "vip", "vrrp", "hsrp", "glbp", "carp"),
				},
			},
			"dns_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hostname or FQDN (not case-sensitive).",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(255),
					stringvalidator.RegexMatches(regexp.MustCompile(`^([0-9A-Za-z_-]+|\*)(\.[0-9A-Za-z_-]+)*\.?$`), ""),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"vrf_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the VRF.",
			},
			"nat_inside_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the IP address this address is the NAT (outside) IP for.",
			},
			"nat_inside_address_id": schema.Int64Attribute{
				Computed:           true,
				DeprecationMessage: "Use nat_inside_id instead.",
				Description:        "Deprecated alias of nat_inside_id.",
			},
			"nat_outside_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the addresses that name this one as their nat_inside_id.",
			},
			"assigned_object_type": schema.StringAttribute{
				Computed:    true,
				Description: "Content type of the object the address is assigned to. Derived from device_interface_id or virtual_machine_interface_id when one of those is set; set it together with assigned_object_id for other object types. One of: dcim.interface, virtualization.vminterface, ipam.fhrpgroup.",
			},
			"assigned_object_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the object the address is assigned to (see assigned_object_type).",
			},
			"device_interface_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the device interface the address is assigned to (assigned_object_type dcim.interface). Conflicts with virtual_machine_interface_id and with setting the assigned_object_* pair directly.",
			},
			"virtual_machine_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the virtual machine interface the address is assigned to (assigned_object_type virtualization.vminterface). Conflicts with device_interface_id and with setting the assigned_object_* pair directly.",
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
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
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
			"device": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the device the address is assigned to.",
			},
			"interface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the interface the address is assigned to.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device the address is assigned to.",
			},
			"parent_prefix": schema.StringAttribute{
				Optional:    true,
				Description: "Prefix the address lies within.",
			},
			"tenant_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant.",
			},
			"vrf_rd": schema.StringAttribute{
				Optional:    true,
				Description: "Route distinguisher of the VRF.",
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

func (d *ipAddressDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ipAddressDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ipAddressDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamIPAddressesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.IPAddress.IsNull() {
		v := state.IPAddress.ValueString()
		params.SetAddress([]string{v})
		hasInput = true
	}
	if !state.DNSName.IsNull() {
		v := state.DNSName.ValueString()
		params.SetDNSName([]string{v})
		hasInput = true
	}
	if !state.VrfID.IsNull() {
		v := state.VrfID.ValueInt64()
		params.SetVrfID([]int64{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.Device.IsNull() {
		v := state.Device.ValueString()
		params.SetDevice([]string{v})
		hasInput = true
	}
	if !state.InterfaceID.IsNull() {
		v := state.InterfaceID.ValueInt64()
		params.SetInterfaceID([]int64{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.Role.IsNull() {
		v := state.Role.ValueString()
		params.SetRole([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.VirtualMachineInterfaceID.IsNull() {
		v := state.VirtualMachineInterfaceID.ValueInt64()
		params.SetVminterfaceID([]int64{v})
		hasInput = true
	}
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.ParentPrefix.IsNull() {
		v := state.ParentPrefix.ValueString()
		params.SetParent([]string{v})
		hasInput = true
	}
	if !state.TenantSlug.IsNull() {
		v := state.TenantSlug.ValueString()
		params.SetTenant([]string{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or ip_address and/or dns_name and/or vrf_id and/or tenant_id and/or device and/or interface_id and/or status and/or role and/or description and/or virtual_machine_interface_id and/or device_id and/or parent_prefix and/or tenant_slug and/or vrf_rd and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_ip_address.")
		return
	}
	res, err := d.client.Ipam.IpamIPAddressesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_ip_address", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_ip_address",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenIPAddress(ctx, netboxapi.IPAddressResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.ipAddressResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&ipAddressResource{client: d.client}).postRead(ctx, &state.ipAddressResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
