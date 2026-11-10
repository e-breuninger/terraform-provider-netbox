// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

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
	_ datasource.DataSource              = (*serviceDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*serviceDataSource)(nil)
)

// NewServiceDataSource returns a new service data source.
func NewServiceDataSource() datasource.DataSource {
	return &serviceDataSource{}
}

type serviceDataSource struct {
	client *netboxapi.Client
}

// serviceDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type serviceDataSourceModel struct {
	serviceResourceModel
	TagSlug            types.String  `tfsdk:"tag_slug"`
	Contact            types.Int64   `tfsdk:"contact"`
	ContactGroupSlug   types.String  `tfsdk:"contact_group_slug"`
	ContactRole        types.Int64   `tfsdk:"contact_role"`
	DeviceName         types.String  `tfsdk:"device_name"`
	Fhrpgroup          types.String  `tfsdk:"fhrpgroup"`
	FhrpgroupID        types.Int64   `tfsdk:"fhrpgroup_id"`
	IPAddress          types.String  `tfsdk:"ip_address"`
	IPAddressID        types.Int64   `tfsdk:"ip_address_id"`
	OwnerName          types.String  `tfsdk:"owner_name"`
	OwnerGroupName     types.String  `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64   `tfsdk:"owner_group_id"`
	Port               types.Float64 `tfsdk:"port"`
	Q                  types.String  `tfsdk:"q"`
	TagID              types.Int64   `tfsdk:"tag_id"`
	VirtualMachineName types.String  `tfsdk:"virtual_machine_name"`
	CustomFieldFilters types.Map     `tfsdk:"custom_field_filters"`
}

func (d *serviceDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_service"
}

func (d *serviceDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):A service running on a device or virtual machine (ipam.service).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/service/):\n\n> An application service represents a layer seven application available on a device or virtual machine. For example, a service might be created in NetBox to represent an HTTP server running on TCP/8000. Each service may optionally be further bound to one or more specific interfaces assigned to the selected device or virtual machine.\n>\n> To aid in the efficient creation of application services, users may opt to first create an [application service template](https://netboxlabs.com/docs/netbox/models/ipam/servicetemplate/) from which service definitions can be quickly replicated.\n>\n> **Changed in NetBox v4.4:** Previously, application services were referred to simply as \"services\". The name has been changed in the UI to better reflect their intended use. There is no change to the name of the model or in any programmatic NetBox APIs.",
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
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Transport protocol. One of: tcp, udp, sctp.",
				Validators: []validator.String{
					stringvalidator.OneOf("tcp", "udp", "sctp"),
				},
			},
			"ports": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Port numbers the service listens on.",
			},
			"parent_object_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the object the service runs on. Derived from device_id or virtual_machine_id when one of those is set; set it together with parent_object_id to address the object directly. One of: dcim.device, virtualization.virtualmachine.",
				Validators: []validator.String{
					stringvalidator.OneOf("dcim.device", "virtualization.virtualmachine"),
				},
			},
			"parent_object_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the object named by parent_object_type.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device the service runs on (parent_object_type dcim.device). Conflicts with virtual_machine_id and with setting the parent_object_* pair directly.",
			},
			"virtual_machine_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the virtual machine the service runs on (parent_object_type virtualization.virtualmachine). Conflicts with device_id and with setting the parent_object_* pair directly.",
			},
			"ip_address_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the IP addresses the service is bound to; all of the parent object's addresses when unset.",
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
			"device_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the device.",
			},
			"fhrpgroup": schema.StringAttribute{
				Optional:    true,
				Description: "Fhrpgroup to match.",
			},
			"fhrpgroup_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the fhrpgroup.",
			},
			"ip_address": schema.StringAttribute{
				Optional:    true,
				Description: "IP address.",
			},
			"ip_address_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the IP address.",
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
			"port": schema.Float64Attribute{
				Optional:    true,
				Description: "Port to match.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"virtual_machine_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the virtual machine.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *serviceDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *serviceDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data serviceDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamServicesListParams()
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
	if !state.Protocol.IsNull() {
		v := state.Protocol.ValueString()
		params.SetProtocol(&v)
		hasInput = true
	}
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.VirtualMachineID.IsNull() {
		v := state.VirtualMachineID.ValueInt64()
		params.SetVirtualMachineID([]int64{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
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
	if !state.DeviceName.IsNull() {
		v := state.DeviceName.ValueString()
		params.SetDevice([]string{v})
		hasInput = true
	}
	if !state.Fhrpgroup.IsNull() {
		v := state.Fhrpgroup.ValueString()
		params.SetFhrpgroup([]string{v})
		hasInput = true
	}
	if !state.FhrpgroupID.IsNull() {
		v := state.FhrpgroupID.ValueInt64()
		params.SetFhrpgroupID([]int64{v})
		hasInput = true
	}
	if !state.IPAddress.IsNull() {
		v := state.IPAddress.ValueString()
		params.SetIPAddress([]string{v})
		hasInput = true
	}
	if !state.IPAddressID.IsNull() {
		v := state.IPAddressID.ValueInt64()
		params.SetIPAddressID([]int64{v})
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
	if !state.ParentObjectID.IsNull() {
		v := state.ParentObjectID.ValueInt64()
		params.SetParentObjectID([]int64{v})
		hasInput = true
	}
	if !state.ParentObjectType.IsNull() {
		v := state.ParentObjectType.ValueString()
		params.SetParentObjectType([]string{v})
		hasInput = true
	}
	if !state.Port.IsNull() {
		v := state.Port.ValueFloat64()
		params.SetPort(&v)
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
	if !state.VirtualMachineName.IsNull() {
		v := state.VirtualMachineName.ValueString()
		params.SetVirtualMachine([]string{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_service.")
		return
	}
	res, err := d.client.Ipam.IpamServicesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_service", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_service",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenService(ctx, netboxapi.ServiceResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.serviceResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) derives attributes from what the API
	// returned (aliases of a polymorphic pair, and the like); a data source reads the same object,
	// so it runs the hook on the embedded resource model.
	(&serviceResource{client: d.client}).postReadHook(ctx, &state.serviceResourceModel, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
