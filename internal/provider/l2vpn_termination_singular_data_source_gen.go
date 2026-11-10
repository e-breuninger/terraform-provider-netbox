// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*l2vpnTerminationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*l2vpnTerminationDataSource)(nil)
)

// NewL2vpnTerminationDataSource returns a new l2vpn_termination data source.
func NewL2vpnTerminationDataSource() datasource.DataSource {
	return &l2vpnTerminationDataSource{}
}

type l2vpnTerminationDataSource struct {
	client *netboxapi.Client
}

// l2vpnTerminationDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type l2vpnTerminationDataSourceModel struct {
	l2vpnTerminationResourceModel
	TagSlug              types.String `tfsdk:"tag_slug"`
	AssignedObjectTypeID types.Int64  `tfsdk:"assigned_object_type_id"`
	DeviceName           types.String `tfsdk:"device_name"`
	DeviceID             types.Int64  `tfsdk:"device_id"`
	InterfaceName        types.String `tfsdk:"interface_name"`
	L2vpnSlug            types.String `tfsdk:"l2vpn_slug"`
	Q                    types.String `tfsdk:"q"`
	RegionSlug           types.String `tfsdk:"region_slug"`
	RegionID             types.Int64  `tfsdk:"region_id"`
	SiteSlug             types.String `tfsdk:"site_slug"`
	SiteID               types.Int64  `tfsdk:"site_id"`
	TagID                types.Int64  `tfsdk:"tag_id"`
	VirtualMachineName   types.String `tfsdk:"virtual_machine_name"`
	VirtualMachineID     types.Int64  `tfsdk:"virtual_machine_id"`
	VlanName             types.String `tfsdk:"vlan_name"`
	VlanVid              types.Int64  `tfsdk:"vlan_vid"`
	VminterfaceName      types.String `tfsdk:"vminterface_name"`
	CustomFieldFilters   types.Map    `tfsdk:"custom_field_filters"`
}

func (d *l2vpnTerminationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2vpn_termination"
}

func (d *l2vpnTerminationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:An L2VPN termination attaching an interface or VLAN to an L2VPN (vpn.l2vpntermination).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/vpn/l2vpntermination/):\n\n> A L2VPN termination is the attachment of an [L2VPN](https://netboxlabs.com/docs/netbox/models/vpn/l2vpn/) to an [interface](https://netboxlabs.com/docs/netbox/models/dcim/interface/) or [VLAN](https://netboxlabs.com/docs/netbox/models/ipam/vlan/). Note that the L2VPNs of the following types may have only two terminations assigned to them:\n>\n> * VPWS\n> * EPL\n> * EP-LAN\n> * EP-TREE",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"l2vpn_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the L2VPN.",
			},
			"assigned_object_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the terminating object. Derived from device_interface_id, virtual_machine_interface_id or vlan_id when one of those is set; set it together with assigned_object_id otherwise. One assignment is required. One of: dcim.interface, virtualization.vminterface, ipam.vlan.",
				Validators: []validator.String{
					stringvalidator.OneOf("dcim.interface", "virtualization.vminterface", "ipam.vlan"),
				},
			},
			"assigned_object_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating object (see assigned_object_type).",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating device interface (assigned_object_type dcim.interface). Conflicts with the other aliases and with setting the assigned_object_* pair directly.",
			},
			"virtual_machine_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating virtual machine interface (assigned_object_type virtualization.vminterface). Conflicts with the other aliases and with setting the assigned_object_* pair directly.",
			},
			"vlan_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminating VLAN (assigned_object_type ipam.vlan). Conflicts with the other aliases and with setting the assigned_object_* pair directly.",
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"assigned_object_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the assigned object type.",
			},
			"device_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the device.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device.",
			},
			"interface_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the interface.",
			},
			"l2vpn_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the L2VPN.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
			},
			"region_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"site_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site.",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"virtual_machine_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the virtual machine.",
			},
			"virtual_machine_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the virtual machine.",
			},
			"vlan_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the VLAN.",
			},
			"vlan_vid": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the VLAN number (1-4094).",
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

func (d *l2vpnTerminationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *l2vpnTerminationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data l2vpnTerminationDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := vpn.NewVpnL2vpnTerminationsListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.L2vpnID.IsNull() {
		v := state.L2vpnID.ValueInt64()
		params.SetL2vpnID([]int64{v})
		hasInput = true
	}
	if !state.VlanID.IsNull() {
		v := state.VlanID.ValueInt64()
		params.SetVlanID([]int64{v})
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
	if !state.AssignedObjectTypeID.IsNull() {
		v := state.AssignedObjectTypeID.ValueInt64()
		params.SetAssignedObjectTypeID([]int64{v})
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.DeviceName.IsNull() {
		v := state.DeviceName.ValueString()
		params.SetDevice([]string{v})
		hasInput = true
	}
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.InterfaceName.IsNull() {
		v := state.InterfaceName.ValueString()
		params.SetInterface([]string{v})
		hasInput = true
	}
	if !state.L2vpnSlug.IsNull() {
		v := state.L2vpnSlug.ValueString()
		params.SetL2vpn([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueInt64()
		params.SetRegionID([]int64{v})
		hasInput = true
	}
	if !state.SiteSlug.IsNull() {
		v := state.SiteSlug.ValueString()
		params.SetSite([]string{v})
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
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
	if !state.VlanName.IsNull() {
		v := state.VlanName.ValueString()
		params.SetVlan([]string{v})
		hasInput = true
	}
	if !state.VlanVid.IsNull() {
		v := state.VlanVid.ValueInt64()
		params.SetVlanVid(&v)
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_l2vpn_termination.")
		return
	}
	res, err := d.client.Vpn.VpnL2vpnTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_l2vpn_termination", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_l2vpn_termination",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenL2vpnTermination(ctx, netboxapi.L2vpnTerminationResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.l2vpnTerminationResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&l2vpnTerminationResource{client: d.client}).postReadHook(ctx, &state.l2vpnTerminationResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
