// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualDeviceContextDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualDeviceContextDataSource)(nil)
)

// NewVirtualDeviceContextDataSource returns a new virtual_device_context data source.
func NewVirtualDeviceContextDataSource() datasource.DataSource {
	return &virtualDeviceContextDataSource{}
}

type virtualDeviceContextDataSource struct {
	client *netboxapi.Client
}

// virtualDeviceContextDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type virtualDeviceContextDataSourceModel struct {
	virtualDeviceContextResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	Device             types.Int64  `tfsdk:"device"`
	HasPrimaryIP       types.Bool   `tfsdk:"has_primary_ip"`
	InterfaceID        types.Int64  `tfsdk:"interface_id"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	PrimaryIp4         types.String `tfsdk:"primary_ip4"`
	PrimaryIp6         types.String `tfsdk:"primary_ip6"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	TenantSlug         types.String `tfsdk:"tenant_slug"`
	TenantGroupSlug    types.String `tfsdk:"tenant_group_slug"`
	TenantGroupID      types.String `tfsdk:"tenant_group_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *virtualDeviceContextDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_device_context"
}

func (d *virtualDeviceContextDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A virtual device context: a logical partition of a device with its own interfaces and addresses (dcim.virtualdevicecontext).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/dcim/virtualdevicecontext/):\n\n> A virtual device context (VDC) represents a logical partition within a physical device, to which interfaces from the parent device can be allocated. Each VDC effectively provides an isolated control plane, but relies on shared resources of the parent device. A VDC is somewhat similar to a virtual machine in that it effects isolation between various components, but stops short of delivering a fully virtualized environment.\n>\n> Each VDC must be assigned to a device upon creation, after which interfaces belonging to that device can be assigned to one or more of its VDCs. A VDC can have any number of interfaces assigned to it, and an interface can belong to any number of VDCs.\n>\n> **A VDC by Any Other Name:** Network vendors use differing names for this concept. Cisco uses the term VDC, whereas Juniper refers to it as a _Virtual Routing Instance_, and Fortinet uses _Virtual Domain_, for instance. While there may be some nuance among the vendors' unique implementations, the general concept remains the same for each.",
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
			"identifier": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Numeric identifier, unique within the device. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 32767),
				},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: active, planned, offline.",
				Validators: []validator.String{
					stringvalidator.OneOf("active", "planned", "offline"),
				},
			},
			"tenant_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the tenant.",
			},
			"primary_ip4_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the primary IPv4 address; it must be assigned to an interface of the device.",
			},
			"primary_ip6_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the primary IPv6 address; it must be assigned to an interface of the device.",
			},
			"interface_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Number of interfaces assigned to the context.",
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"device": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the device model.",
			},
			"has_primary_ip": schema.BoolAttribute{
				Optional:    true,
				Description: "Has a primary IP.",
			},
			"interface_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the interface.",
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
			"primary_ip4": schema.StringAttribute{
				Optional:    true,
				Description: "Primary IPv4 (address).",
			},
			"primary_ip6": schema.StringAttribute{
				Optional:    true,
				Description: "Primary IPv6 (address).",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"tenant_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant.",
			},
			"tenant_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the tenant group.",
			},
			"tenant_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the tenant group.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *virtualDeviceContextDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualDeviceContextDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualDeviceContextDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimVirtualDeviceContextsListParams()
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
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.TenantID.IsNull() {
		v := state.TenantID.ValueInt64()
		params.SetTenantID([]int64{v})
		hasInput = true
	}
	if !state.Identifier.IsNull() {
		v := state.Identifier.ValueInt64()
		params.SetIdentifier([]int64{v})
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
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.Device.IsNull() {
		v := state.Device.ValueInt64()
		params.SetDevice([]int64{v})
		hasInput = true
	}
	if !state.HasPrimaryIP.IsNull() {
		v := state.HasPrimaryIP.ValueBool()
		params.SetHasPrimaryIP(&v)
		hasInput = true
	}
	if !state.InterfaceID.IsNull() {
		v := state.InterfaceID.ValueInt64()
		params.SetInterfaceID([]int64{v})
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
	if !state.PrimaryIp4.IsNull() {
		v := state.PrimaryIp4.ValueString()
		params.SetPrimaryIp4([]string{v})
		hasInput = true
	}
	if !state.PrimaryIp4ID.IsNull() {
		v := state.PrimaryIp4ID.ValueInt64()
		params.SetPrimaryIp4ID([]int64{v})
		hasInput = true
	}
	if !state.PrimaryIp6.IsNull() {
		v := state.PrimaryIp6.ValueString()
		params.SetPrimaryIp6([]string{v})
		hasInput = true
	}
	if !state.PrimaryIp6ID.IsNull() {
		v := state.PrimaryIp6ID.ValueInt64()
		params.SetPrimaryIp6ID([]int64{v})
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
	if !state.TenantSlug.IsNull() {
		v := state.TenantSlug.ValueString()
		params.SetTenant([]string{v})
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
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_virtual_device_context.")
		return
	}
	res, err := d.client.Dcim.DcimVirtualDeviceContextsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_virtual_device_context", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_virtual_device_context",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVirtualDeviceContext(ctx, netboxapi.VirtualDeviceContextResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.virtualDeviceContextResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
