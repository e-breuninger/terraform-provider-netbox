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
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
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
	Device              types.String  `tfsdk:"device"`
	InterfaceID         types.Int64   `tfsdk:"interface_id"`
	DeviceID            types.Int64   `tfsdk:"device_id"`
	ParentPrefix        types.String  `tfsdk:"parent_prefix"`
	TenantSlug          types.String  `tfsdk:"tenant_slug"`
	VrfRd               types.String  `tfsdk:"vrf_rd"`
	TagSlug             types.String  `tfsdk:"tag_slug"`
	Assigned            types.Bool    `tfsdk:"assigned"`
	AssignedToInterface types.Bool    `tfsdk:"assigned_to_interface"`
	Contact             types.Int64   `tfsdk:"contact"`
	ContactGroupSlug    types.String  `tfsdk:"contact_group_slug"`
	ContactRole         types.Int64   `tfsdk:"contact_role"`
	FhrpgroupID         types.Int64   `tfsdk:"fhrpgroup_id"`
	InterfaceName       types.String  `tfsdk:"interface_name"`
	MaskLengthGte       types.Float64 `tfsdk:"mask_length__gte"`
	MaskLengthLte       types.Float64 `tfsdk:"mask_length__lte"`
	MaskLength          types.Int64   `tfsdk:"mask_length"`
	OwnerName           types.String  `tfsdk:"owner_name"`
	OwnerGroupName      types.String  `tfsdk:"owner_group_name"`
	OwnerGroupID        types.Int64   `tfsdk:"owner_group_id"`
	PresentInVrf        types.String  `tfsdk:"present_in_vrf"`
	PresentInVrfID      types.String  `tfsdk:"present_in_vrf_id"`
	Q                   types.String  `tfsdk:"q"`
	ServiceID           types.Int64   `tfsdk:"service_id"`
	TagID               types.Int64   `tfsdk:"tag_id"`
	TenantGroupSlug     types.String  `tfsdk:"tenant_group_slug"`
	TenantGroupID       types.String  `tfsdk:"tenant_group_id"`
	VirtualMachineName  types.String  `tfsdk:"virtual_machine_name"`
	VirtualMachineID    types.Int64   `tfsdk:"virtual_machine_id"`
	VminterfaceName     types.String  `tfsdk:"vminterface_name"`
	CustomFieldFilters  types.Map     `tfsdk:"custom_field_filters"`
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
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "Content type of the object the address is assigned to. Derived from device_interface_id or virtual_machine_interface_id when one of those is set; set it together with assigned_object_id for other object types. One of: dcim.interface, virtualization.vminterface, ipam.fhrpgroup.",
				Validators: []validator.String{
					stringvalidator.OneOf("dcim.interface", "virtualization.vminterface", "ipam.fhrpgroup"),
				},
			},
			"assigned_object_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the object the address is assigned to (see assigned_object_type).",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
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
				Optional:    true,
				Computed:    true,
				Description: "IP family (4 or 6).",
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
			"assigned": schema.BoolAttribute{
				Optional:    true,
				Description: "Is assigned.",
			},
			"assigned_to_interface": schema.BoolAttribute{
				Optional:    true,
				Description: "Is assigned to an interface.",
			},
			"contact": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact.",
			},
			"contact_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the contact group.",
			},
			"contact_role": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the contact role.",
			},
			"fhrpgroup_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the FHRP group.",
			},
			"interface_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the interface.",
			},
			"mask_length__gte": schema.Float64Attribute{
				Optional:    true,
				Description: "Mask length is at least this.",
			},
			"mask_length__lte": schema.Float64Attribute{
				Optional:    true,
				Description: "Mask length is at most this.",
			},
			"mask_length": schema.Int64Attribute{
				Optional:    true,
				Description: "Mask length to match.",
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
			"present_in_vrf": schema.StringAttribute{
				Optional:    true,
				Description: "Present in vrf to match.",
			},
			"present_in_vrf_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the present in vrf.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"service_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the application service.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"tenant_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant group.",
			},
			"tenant_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the tenant group.",
			},
			"virtual_machine_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the virtual machine.",
			},
			"virtual_machine_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual machine.",
			},
			"vminterface_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the VM interface.",
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
	if !state.Assigned.IsNull() {
		v := state.Assigned.ValueBool()
		params.SetAssigned(&v)
		hasInput = true
	}
	if !state.AssignedObjectID.IsNull() {
		v := state.AssignedObjectID.ValueInt64()
		params.SetAssignedObjectID([]int64{v})
		hasInput = true
	}
	if !state.AssignedObjectType.IsNull() {
		v := state.AssignedObjectType.ValueString()
		params.SetAssignedObjectType([]string{v})
		hasInput = true
	}
	if !state.AssignedToInterface.IsNull() {
		v := state.AssignedToInterface.ValueBool()
		params.SetAssignedToInterface(&v)
		hasInput = true
	}
	if !state.Contact.IsNull() {
		v := state.Contact.ValueInt64()
		params.SetContact([]int64{v})
		hasInput = true
	}
	if !state.ContactGroupSlug.IsNull() {
		v := state.ContactGroupSlug.ValueString()
		params.SetContactGroup([]string{v})
		hasInput = true
	}
	if !state.ContactRole.IsNull() {
		v := state.ContactRole.ValueInt64()
		params.SetContactRole([]int64{v})
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.Family.IsNull() {
		v := float64(state.Family.ValueInt64())
		params.SetFamily(&v)
		hasInput = true
	}
	if !state.FhrpgroupID.IsNull() {
		v := state.FhrpgroupID.ValueInt64()
		params.SetFhrpgroupID([]int64{v})
		hasInput = true
	}
	if !state.InterfaceName.IsNull() {
		v := state.InterfaceName.ValueString()
		params.SetInterface([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.MaskLengthGte.IsNull() {
		v := state.MaskLengthGte.ValueFloat64()
		params.SetMaskLengthGte(&v)
		hasInput = true
	}
	if !state.MaskLengthLte.IsNull() {
		v := state.MaskLengthLte.ValueFloat64()
		params.SetMaskLengthLte(&v)
		hasInput = true
	}
	if !state.MaskLength.IsNull() {
		v := state.MaskLength.ValueInt64()
		params.SetMaskLength([]int64{v})
		hasInput = true
	}
	if !state.NatInsideID.IsNull() {
		v := state.NatInsideID.ValueInt64()
		params.SetNatInsideID([]int64{v})
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
	if !state.PresentInVrf.IsNull() {
		v := state.PresentInVrf.ValueString()
		params.SetPresentInVrf(&v)
		hasInput = true
	}
	if !state.PresentInVrfID.IsNull() {
		v := state.PresentInVrfID.ValueString()
		params.SetPresentInVrfID(&v)
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.ServiceID.IsNull() {
		v := state.ServiceID.ValueInt64()
		params.SetServiceID([]int64{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.TenantGroupSlug.IsNull() {
		v := state.TenantGroupSlug.ValueString()
		params.SetTenantGroup([]string{v})
		hasInput = true
	}
	if !state.TenantGroupID.IsNull() {
		v := state.TenantGroupID.ValueString()
		params.SetTenantGroupID([]string{v})
		hasInput = true
	}
	if !state.VirtualMachineName.IsNull() {
		v := state.VirtualMachineName.ValueString()
		params.SetVirtualMachine([]string{v})
		hasInput = true
	}
	if !state.VirtualMachineID.IsNull() {
		v := state.VirtualMachineID.ValueInt64()
		params.SetVirtualMachineID([]int64{v})
		hasInput = true
	}
	if !state.VminterfaceName.IsNull() {
		v := state.VminterfaceName.ValueString()
		params.SetVminterface([]string{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_ip_address.")
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
	(&ipAddressResource{client: d.client}).postReadHook(ctx, &state.ipAddressResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
