// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*deviceInterfaceDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*deviceInterfaceDataSource)(nil)
)

// NewDeviceInterfaceDataSource returns a new device_interface data source.
func NewDeviceInterfaceDataSource() datasource.DataSource {
	return &deviceInterfaceDataSource{}
}

type deviceInterfaceDataSource struct {
	client *netboxapi.Client
}

// deviceInterfaceDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type deviceInterfaceDataSourceModel struct {
	deviceInterfaceResourceModel
	MacAddress         types.String `tfsdk:"mac_address"`
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *deviceInterfaceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_device_interface"
}

func (d *deviceInterfaceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A device interface (dcim.interface).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/interface/):\n\n> Interfaces in NetBox represent network interfaces used to exchange data with connected devices. On modern networks, these are most commonly Ethernet, but other types are supported as well. IP addresses and VLANs can be assigned to interfaces.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 64),
				},
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface type as its NetBox slug, e.g. virtual, lag, bridge, 1000base-t, 10gbase-x-sfpp (any of NetBox's interface type choices).",
			},
			"label": schema.StringAttribute{
				Computed:    true,
				Description: "Physical label.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the interface is enabled.",
			},
			"mgmt_only": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the interface is used for out-of-band management only.",
			},
			"mgmtonly": schema.BoolAttribute{
				Computed:           true,
				DeprecationMessage: "Use mgmt_only instead.",
				Description:        "Deprecated alias of mgmt_only.",
			},
			"mark_connected": schema.BoolAttribute{
				Computed:    true,
				Description: "Treat as if a cable is connected.",
			},
			"mtu": schema.Int64Attribute{
				Computed: true,
			},
			"speed": schema.Int64Attribute{
				Computed:    true,
				Description: "Speed in kbps.",
			},
			"duplex": schema.StringAttribute{
				Computed:    true,
				Description: "One of: half, full, auto.",
			},
			"wwn": schema.StringAttribute{
				Computed:    true,
				Description: "64-bit World Wide Name.",
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
			"untagged_vlan": schema.Int64Attribute{
				Computed:           true,
				DeprecationMessage: "Use untagged_vlan_id instead.",
				Description:        "Deprecated alias of untagged_vlan_id.",
			},
			"tagged_vlan_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the tagged VLANs.",
			},
			"tagged_vlans": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use tagged_vlan_ids instead.",
				Description:        "Deprecated alias of tagged_vlan_ids.",
			},
			"module_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the installed module this interface belongs to.",
			},
			"lag_device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the parent LAG interface.",
			},
			"parent_device_interface_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the parent interface.",
			},
			"bridge_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the bridge interface.",
			},
			"vrf_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the VRF.",
			},
			"vdc_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the virtual device contexts.",
			},
			"description": schema.StringAttribute{
				Computed: true,
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
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Description: "MAC address assigned to the interface.",
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

func (d *deviceInterfaceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *deviceInterfaceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data deviceInterfaceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimInterfacesListParams()
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
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.Type.IsNull() {
		v := state.Type.ValueString()
		params.SetType([]string{v})
		hasInput = true
	}
	if !state.Mode.IsNull() {
		v := state.Mode.ValueString()
		params.SetMode([]string{v})
		hasInput = true
	}
	if !state.Enabled.IsNull() {
		v := state.Enabled.ValueBool()
		params.SetEnabled(&v)
		hasInput = true
	}
	if !state.LagDeviceInterfaceID.IsNull() {
		v := state.LagDeviceInterfaceID.ValueInt64()
		params.SetLagID([]int64{v})
		hasInput = true
	}
	if !state.MacAddress.IsNull() {
		v := state.MacAddress.ValueString()
		params.SetMacAddress([]string{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or device_id and/or type and/or mode and/or enabled and/or lag_device_interface_id and/or mac_address and/or owner_id and/or tag_slug and/or custom_field_filters to look up a netbox_device_interface.")
		return
	}
	res, err := d.client.Dcim.DcimInterfacesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_device_interface", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_device_interface",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenDeviceInterface(ctx, netboxapi.DeviceInterfaceResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.deviceInterfaceResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
