// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*circuitGroupAssignmentDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*circuitGroupAssignmentDataSource)(nil)
)

// NewCircuitGroupAssignmentDataSource returns a new circuit_group_assignment data source.
func NewCircuitGroupAssignmentDataSource() datasource.DataSource {
	return &circuitGroupAssignmentDataSource{}
}

type circuitGroupAssignmentDataSource struct {
	client *netboxapi.Client
}

// circuitGroupAssignmentDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type circuitGroupAssignmentDataSourceModel struct {
	circuitGroupAssignmentResourceModel
	TagSlug types.String `tfsdk:"tag_slug"`
}

func (d *circuitGroupAssignmentDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_circuit_group_assignment"
}

func (d *circuitGroupAssignmentDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:Membership of a circuit or virtual circuit in a circuit group (circuits.circuitgroupassignment).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/circuits/circuitgroupassignment/):\n\n> Circuits can be assigned to [circuit groups](https://netboxlabs.com/docs/netbox/models/circuits/circuitgroup/) for correlation purposes. For instance, three circuits, each belonging to a different provider, may each be assigned to the same circuit group. Each assignment may optionally include a priority designation.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"circuit_group_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the circuit group the member is assigned to.",
			},
			"member_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Content type of the assigned member. One of: circuits.circuit, circuits.virtualcircuit.",
				Validators: []validator.String{
					stringvalidator.OneOf("circuits.circuit", "circuits.virtualcircuit"),
				},
			},
			"member_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the assigned circuit or virtual circuit.",
				Validators: []validator.Int64{
					int64validator.AtLeast(0),
				},
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Priority of this member within the group. One of: primary, secondary, tertiary, inactive.",
				Validators: []validator.String{
					stringvalidator.OneOf("primary", "secondary", "tertiary", "inactive"),
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
		},
	}
}

func (d *circuitGroupAssignmentDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *circuitGroupAssignmentDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data circuitGroupAssignmentDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := circuits.NewCircuitsCircuitGroupAssignmentsListParams()
	hasInput := false
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.CircuitGroupID.IsNull() {
		v := state.CircuitGroupID.ValueInt64()
		params.SetGroupID([]int64{v})
		hasInput = true
	}
	if !state.MemberType.IsNull() {
		v := state.MemberType.ValueString()
		params.SetMemberType([]string{v})
		hasInput = true
	}
	if !state.MemberID.IsNull() {
		v := state.MemberID.ValueInt64()
		params.SetMemberID([]int64{v})
		hasInput = true
	}
	if !state.Priority.IsNull() {
		v := state.Priority.ValueString()
		params.SetPriority(&v)
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or circuit_group_id and/or member_type and/or member_id and/or priority and/or tag_slug to look up a netbox_circuit_group_assignment.")
		return
	}
	res, err := d.client.Circuits.CircuitsCircuitGroupAssignmentsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_circuit_group_assignment", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_circuit_group_assignment",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenCircuitGroupAssignment(ctx, netboxapi.CircuitGroupAssignmentResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.circuitGroupAssignmentResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
