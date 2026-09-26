// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualCircuitTerminationDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualCircuitTerminationDataSource)(nil)
)

// NewVirtualCircuitTerminationDataSource returns a new virtual_circuit_termination data source.
func NewVirtualCircuitTerminationDataSource() datasource.DataSource {
	return &virtualCircuitTerminationDataSource{}
}

type virtualCircuitTerminationDataSource struct {
	client *netboxapi.Client
}

// virtualCircuitTerminationDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type virtualCircuitTerminationDataSourceModel struct {
	virtualCircuitTerminationResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *virtualCircuitTerminationDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_circuit_termination"
}

func (d *virtualCircuitTerminationDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:Termination of a virtual circuit on an interface (circuits.virtualcircuittermination).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/circuits/virtualcircuittermination/):\n\n> This model represents the connection of a virtual [interface](https://netboxlabs.com/docs/netbox/models/dcim/interface/) to a [virtual circuit](https://netboxlabs.com/docs/netbox/models/circuits/virtualcircuit/).",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"virtual_circuit_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the terminated virtual circuit.",
			},
			"device_interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device interface the virtual circuit terminates on.",
			},
			"role": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Role of this termination in the virtual circuit. Leaving it unset keeps the current value. One of: peer, hub, spoke.",
				Validators: []validator.String{
					stringvalidator.OneOf("peer", "hub", "spoke"),
				},
			},
			"description": schema.StringAttribute{
				Computed: true,
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

func (d *virtualCircuitTerminationDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualCircuitTerminationDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualCircuitTerminationDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := circuits.NewCircuitsVirtualCircuitTerminationsListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.VirtualCircuitID.IsNull() {
		v := state.VirtualCircuitID.ValueInt64()
		params.SetVirtualCircuitID([]int64{v})
		hasInput = true
	}
	if !state.DeviceInterfaceID.IsNull() {
		v := state.DeviceInterfaceID.ValueInt64()
		params.SetInterfaceID([]int64{v})
		hasInput = true
	}
	if !state.Role.IsNull() {
		v := state.Role.ValueString()
		params.SetRole([]string{v})
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
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or virtual_circuit_id and/or device_interface_id and/or role and/or tag_slug and/or custom_field_filters to look up a netbox_virtual_circuit_termination.")
		return
	}
	res, err := d.client.Circuits.CircuitsVirtualCircuitTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_virtual_circuit_termination", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_virtual_circuit_termination",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVirtualCircuitTermination(ctx, netboxapi.VirtualCircuitTerminationResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.virtualCircuitTerminationResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
