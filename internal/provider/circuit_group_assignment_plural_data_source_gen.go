// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*circuitGroupAssignmentsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*circuitGroupAssignmentsDataSource)(nil)
)

// NewCircuitGroupAssignmentsDataSource returns a new circuit_group_assignments data source, which
// lists circuit_group_assignment objects matching its filters.
func NewCircuitGroupAssignmentsDataSource() datasource.DataSource {
	return &circuitGroupAssignmentsDataSource{}
}

type circuitGroupAssignmentsDataSource struct {
	client *netboxapi.Client
}

// circuitGroupAssignmentsDataSourceModel is the data source model: the filters, the limit and the matching
// circuit_group_assignments.
type circuitGroupAssignmentsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"circuit_group_assignments"`
}

// circuitGroupAssignmentResourceAttrTypes is the attribute type map of circuitGroupAssignmentResourceModel.
func circuitGroupAssignmentResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":               types.Int64Type,
		"circuit_group_id": types.Int64Type,
		"member_type":      types.StringType,
		"member_id":        types.Int64Type,
		"priority":         types.StringType,
		"created":          types.StringType,
		"last_updated":     types.StringType,
		"url":              types.StringType,
		"tags":             types.SetType{ElemType: types.StringType},
		"tags_all":         types.SetType{ElemType: types.StringType},
	}
}

func (d *circuitGroupAssignmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_circuit_group_assignments"
}

func (d *circuitGroupAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:Lists circuit_group_assignment objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: circuit, circuit_id, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, member_id, member_id__empty, member_id__gt, member_id__gte, member_id__lt, member_id__lte, member_id__n, member_type, member_type__n, member_type_id, member_type_id__n, priority, priority__empty, priority__ic, priority__ie, priority__iew, priority__iregex, priority__isw, priority__n, priority__nic, priority__nie, priority__niew, priority__nisw, priority__regex, provider, provider_id, q, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, virtual_circuit, virtual_circuit_id. Repeating a name sends that parameter once per value.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"name": schema.StringAttribute{
							Required:    true,
							Description: "Name of the query filter (the API list parameter).",
						},
						"value": schema.StringAttribute{
							Required:    true,
							Description: "Value to filter by.",
						},
					},
				},
			},
			"limit": schema.Int64Attribute{
				Optional:    true,
				Description: "The maximum number of objects to return from the API lookup. Defaults to 1000.",
				Validators: []validator.Int64{
					int64validator.AtLeast(1),
				},
			},
			"circuit_group_assignments": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching circuit_group_assignments.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"circuit_group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the circuit group the member is assigned to.",
						},
						"member_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the assigned member. One of: circuits.circuit, circuits.virtualcircuit.",
						},
						"member_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the assigned circuit or virtual circuit.",
						},
						"priority": schema.StringAttribute{
							Computed:    true,
							Description: "Priority of this member within the group. One of: primary, secondary, tertiary, inactive.",
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
					},
				},
			},
		},
	}
}

func (d *circuitGroupAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *circuitGroupAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data circuitGroupAssignmentsDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	limit := int64(1000)
	if !state.Limit.IsNull() && !state.Limit.IsUnknown() {
		limit = state.Limit.ValueInt64()
	}
	_ = limit

	// The filters entries, in no particular order. Entries sharing a name are distinct values of
	// the same API parameter.
	var filters []struct {
		Name  types.String `tfsdk:"name"`
		Value types.String `tfsdk:"value"`
	}
	if !state.Filters.IsNull() {
		resp.Diagnostics.Append(state.Filters.ElementsAs(ctx, &filters, false)...)
		if resp.Diagnostics.HasError() {
			return
		}
	}
	_ = filters
	var responseDTOs []*netboxapi.CircuitGroupAssignmentResponseDTO

	params := circuits.NewCircuitsCircuitGroupAssignmentsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"group_id":            netboxapi.FilterInt64s(params.SetGroupID),
		"group_id__n":         netboxapi.FilterInt64s(params.SetGroupIDn),
		"member_type":         netboxapi.FilterStrings(params.SetMemberType),
		"member_type__n":      netboxapi.FilterStrings(params.SetMemberTypen),
		"member_id":           netboxapi.FilterInt64s(params.SetMemberID),
		"member_id__n":        netboxapi.FilterInt64s(params.SetMemberIDn),
		"member_id__lt":       netboxapi.FilterInt64s(params.SetMemberIDLt),
		"member_id__lte":      netboxapi.FilterInt64s(params.SetMemberIDLte),
		"member_id__gt":       netboxapi.FilterInt64s(params.SetMemberIDGt),
		"member_id__gte":      netboxapi.FilterInt64s(params.SetMemberIDGte),
		"member_id__empty":    netboxapi.FilterBool(params.SetMemberIDEmpty),
		"priority":            netboxapi.FilterString(params.SetPriority),
		"priority__n":         netboxapi.FilterString(params.SetPriorityn),
		"priority__ic":        netboxapi.FilterStrings(params.SetPriorityIc),
		"priority__nic":       netboxapi.FilterStrings(params.SetPriorityNic),
		"priority__ie":        netboxapi.FilterStrings(params.SetPriorityIe),
		"priority__nie":       netboxapi.FilterStrings(params.SetPriorityNie),
		"priority__isw":       netboxapi.FilterStrings(params.SetPriorityIsw),
		"priority__nisw":      netboxapi.FilterStrings(params.SetPriorityNisw),
		"priority__iew":       netboxapi.FilterStrings(params.SetPriorityIew),
		"priority__niew":      netboxapi.FilterStrings(params.SetPriorityNiew),
		"priority__empty":     netboxapi.FilterBool(params.SetPriorityEmpty),
		"priority__regex":     netboxapi.FilterStrings(params.SetPriorityRegex),
		"priority__iregex":    netboxapi.FilterStrings(params.SetPriorityIregex),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":            netboxapi.FilterStrings(params.SetTagAny),
		"circuit":             netboxapi.FilterStrings(params.SetCircuit),
		"circuit_id":          netboxapi.FilterInt64s(params.SetCircuitID),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"group":               netboxapi.FilterStrings(params.SetGroup),
		"group__n":            netboxapi.FilterStrings(params.SetGroupn),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"member_type_id":      netboxapi.FilterInt64s(params.SetMemberTypeID),
		"member_type_id__n":   netboxapi.FilterInt64s(params.SetMemberTypeIDn),
		"provider":            netboxapi.FilterStrings(params.SetProvider),
		"provider_id":         netboxapi.FilterInt64s(params.SetProviderID),
		"q":                   netboxapi.FilterString(params.SetQ),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
		"virtual_circuit":     netboxapi.FilterStrings(params.SetVirtualCircuit),
		"virtual_circuit_id":  netboxapi.FilterInt64s(params.SetVirtualCircuitID),
	}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		listFilter, ok := listFilters[name]
		if !ok {
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
		if problem := listFilter.Add(name, value); problem != nil {
			resp.Diagnostics.AddError(problem.Summary, problem.Detail)
			return
		}
	}
	for _, listFilter := range listFilters {
		listFilter.Apply()
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Circuits.CircuitsCircuitGroupAssignmentsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_circuit_group_assignments", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CircuitGroupAssignmentResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]circuitGroupAssignmentResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model circuitGroupAssignmentResourceModel
		resp.Diagnostics.Append(flattenCircuitGroupAssignment(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: circuitGroupAssignmentResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
