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
	_ datasource.DataSource              = (*fhrpGroupAssignmentDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*fhrpGroupAssignmentDataSource)(nil)
)

// NewFhrpGroupAssignmentDataSource returns a new fhrp_group_assignment data source.
func NewFhrpGroupAssignmentDataSource() datasource.DataSource {
	return &fhrpGroupAssignmentDataSource{}
}

type fhrpGroupAssignmentDataSource struct {
	client *netboxapi.Client
}

func (d *fhrpGroupAssignmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fhrp_group_assignment"
}

func (d *fhrpGroupAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Assignment of an FHRP group to an interface (ipam.fhrpgroupassignment).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/ipam/fhrpgroupassignment/):\n\n> Member device and VM interfaces can be assigned to [FHRP groups](https://netboxlabs.com/docs/netbox/models/ipam/fhrpgroup/) to indicate their participation in maintaining a common virtual IP address (VIP). For instance, three interfaces, each belonging to a different router, may each be assigned to the same FHRP group to serve a shared VIP. Each of these assignments would typically receive a different priority.\n>\n> Interfaces are assigned to FHRP groups under the interface detail view.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the FHRP group.",
			},
			"interface_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the interface the group is assigned to, e.g. dcim.interface or virtualization.vminterface.",
			},
			"interface_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the interface named by interface_type.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of this interface within the group (0-255).",
				Validators: []validator.Int64{
					int64validator.Between(0, 255),
				},
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
		},
	}
}

func (d *fhrpGroupAssignmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *fhrpGroupAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data fhrpGroupAssignmentResourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := ipam.NewIpamFhrpGroupAssignmentsListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.GroupID.IsNull() {
		v := state.GroupID.ValueInt64()
		params.SetGroupID([]int64{v})
		hasInput = true
	}
	if !state.InterfaceType.IsNull() {
		v := state.InterfaceType.ValueString()
		params.SetInterfaceType([]string{v})
		hasInput = true
	}
	if !state.InterfaceID.IsNull() {
		v := state.InterfaceID.ValueInt64()
		params.SetInterfaceID([]int64{v})
		hasInput = true
	}
	if !state.Priority.IsNull() {
		v := state.Priority.ValueInt64()
		params.SetPriority([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or group_id and/or interface_type and/or interface_id and/or priority to look up a netbox_fhrp_group_assignment.")
		return
	}
	res, err := d.client.Ipam.IpamFhrpGroupAssignmentsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_fhrp_group_assignment", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_fhrp_group_assignment",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenFhrpGroupAssignment(ctx, netboxapi.FhrpGroupAssignmentResponseDTOFromGoNetbox(res.Payload.Results[0]), state)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
