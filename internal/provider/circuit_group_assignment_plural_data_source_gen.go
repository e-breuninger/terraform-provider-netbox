// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

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
				Description: "Query filters, sent as API list parameters. Supported names: group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, member_id, member_id__n, member_type, member_type__n, priority, priority__empty, priority__ic, priority__ie, priority__iew, priority__iregex, priority__isw, priority__n, priority__nic, priority__nie, priority__niew, priority__nisw, priority__regex, tag, tag__any, tag__n. Repeating a name sends that parameter once per value.",
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

	fetchAll := false
	_ = fetchAll
	var responseDTOs []*netboxapi.CircuitGroupAssignmentResponseDTO

	params := circuits.NewCircuitsCircuitGroupAssignmentsListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vGroupID []int64
	var vGroupIDn []int64
	var vMemberType []string
	var vMemberTypen []string
	var vMemberID []int64
	var vMemberIDn []int64
	var vPriority *string
	var vPriorityn *string
	var vPriorityIc []string
	var vPriorityNic []string
	var vPriorityIe []string
	var vPriorityNie []string
	var vPriorityIsw []string
	var vPriorityNisw []string
	var vPriorityIew []string
	var vPriorityNiew []string
	var vPriorityEmpty *bool
	var vPriorityRegex []string
	var vPriorityIregex []string
	var vTag []string
	var vTagn []string
	var vTagAny []string
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		switch name {
		case "id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id' takes an integer, got %q.", value))
				return
			}
			vID = append(vID, v)
		case "id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__n' takes an integer, got %q.", value))
				return
			}
			vIDn = append(vIDn, v)
		case "id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lt' takes an integer, got %q.", value))
				return
			}
			vIDLt = append(vIDLt, v)
		case "id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__lte' takes an integer, got %q.", value))
				return
			}
			vIDLte = append(vIDLte, v)
		case "id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gt' takes an integer, got %q.", value))
				return
			}
			vIDGt = append(vIDGt, v)
		case "id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__gte' takes an integer, got %q.", value))
				return
			}
			vIDGte = append(vIDGte, v)
		case "id__empty":
			if vIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'id__empty' takes a boolean, got %q.", value))
				return
			}
			vIDEmpty = &v
		case "group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id' takes an integer, got %q.", value))
				return
			}
			vGroupID = append(vGroupID, v)
		case "group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'group_id__n' takes an integer, got %q.", value))
				return
			}
			vGroupIDn = append(vGroupIDn, v)
		case "member_type":
			v := value
			vMemberType = append(vMemberType, v)
		case "member_type__n":
			v := value
			vMemberTypen = append(vMemberTypen, v)
		case "member_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'member_id' takes an integer, got %q.", value))
				return
			}
			vMemberID = append(vMemberID, v)
		case "member_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'member_id__n' takes an integer, got %q.", value))
				return
			}
			vMemberIDn = append(vMemberIDn, v)
		case "priority":
			if vPriority != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'priority' takes a single value.")
				return
			}
			v := value
			vPriority = &v
		case "priority__n":
			if vPriorityn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'priority__n' takes a single value.")
				return
			}
			v := value
			vPriorityn = &v
		case "priority__ic":
			v := value
			vPriorityIc = append(vPriorityIc, v)
		case "priority__nic":
			v := value
			vPriorityNic = append(vPriorityNic, v)
		case "priority__ie":
			v := value
			vPriorityIe = append(vPriorityIe, v)
		case "priority__nie":
			v := value
			vPriorityNie = append(vPriorityNie, v)
		case "priority__isw":
			v := value
			vPriorityIsw = append(vPriorityIsw, v)
		case "priority__nisw":
			v := value
			vPriorityNisw = append(vPriorityNisw, v)
		case "priority__iew":
			v := value
			vPriorityIew = append(vPriorityIew, v)
		case "priority__niew":
			v := value
			vPriorityNiew = append(vPriorityNiew, v)
		case "priority__empty":
			if vPriorityEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'priority__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'priority__empty' takes a boolean, got %q.", value))
				return
			}
			vPriorityEmpty = &v
		case "priority__regex":
			v := value
			vPriorityRegex = append(vPriorityRegex, v)
		case "priority__iregex":
			v := value
			vPriorityIregex = append(vPriorityIregex, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
		default:
			resp.Diagnostics.AddError("Unsupported filter", fmt.Sprintf("'%s' is not a supported filter parameter", name))
			return
		}
	}
	if len(vID) > 0 {
		params.SetID(vID)
	}
	if len(vIDn) > 0 {
		params.SetIDn(vIDn)
	}
	if len(vIDLt) > 0 {
		params.SetIDLt(vIDLt)
	}
	if len(vIDLte) > 0 {
		params.SetIDLte(vIDLte)
	}
	if len(vIDGt) > 0 {
		params.SetIDGt(vIDGt)
	}
	if len(vIDGte) > 0 {
		params.SetIDGte(vIDGte)
	}
	if vIDEmpty != nil {
		params.SetIDEmpty(vIDEmpty)
	}
	if len(vGroupID) > 0 {
		params.SetGroupID(vGroupID)
	}
	if len(vGroupIDn) > 0 {
		params.SetGroupIDn(vGroupIDn)
	}
	if len(vMemberType) > 0 {
		params.SetMemberType(vMemberType)
	}
	if len(vMemberTypen) > 0 {
		params.SetMemberTypen(vMemberTypen)
	}
	if len(vMemberID) > 0 {
		params.SetMemberID(vMemberID)
	}
	if len(vMemberIDn) > 0 {
		params.SetMemberIDn(vMemberIDn)
	}
	if vPriority != nil {
		params.SetPriority(vPriority)
	}
	if vPriorityn != nil {
		params.SetPriorityn(vPriorityn)
	}
	if len(vPriorityIc) > 0 {
		params.SetPriorityIc(vPriorityIc)
	}
	if len(vPriorityNic) > 0 {
		params.SetPriorityNic(vPriorityNic)
	}
	if len(vPriorityIe) > 0 {
		params.SetPriorityIe(vPriorityIe)
	}
	if len(vPriorityNie) > 0 {
		params.SetPriorityNie(vPriorityNie)
	}
	if len(vPriorityIsw) > 0 {
		params.SetPriorityIsw(vPriorityIsw)
	}
	if len(vPriorityNisw) > 0 {
		params.SetPriorityNisw(vPriorityNisw)
	}
	if len(vPriorityIew) > 0 {
		params.SetPriorityIew(vPriorityIew)
	}
	if len(vPriorityNiew) > 0 {
		params.SetPriorityNiew(vPriorityNiew)
	}
	if vPriorityEmpty != nil {
		params.SetPriorityEmpty(vPriorityEmpty)
	}
	if len(vPriorityRegex) > 0 {
		params.SetPriorityRegex(vPriorityRegex)
	}
	if len(vPriorityIregex) > 0 {
		params.SetPriorityIregex(vPriorityIregex)
	}
	if len(vTag) > 0 {
		params.SetTag(vTag)
	}
	if len(vTagn) > 0 {
		params.SetTagn(vTagn)
	}
	if len(vTagAny) > 0 {
		params.SetTagAny(vTagAny)
	}
	pageSize := limit
	if fetchAll {
		pageSize = 1000
	}
	params.SetLimit(&pageSize)
	for offset := int64(0); ; {
		res, err := d.client.Circuits.CircuitsCircuitGroupAssignmentsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_circuit_group_assignments", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CircuitGroupAssignmentResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if !fetchAll || len(res.Payload.Results) == 0 || res.Payload.Count == nil || offset >= *res.Payload.Count {
			break
		}
		params.SetOffset(&offset)
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
