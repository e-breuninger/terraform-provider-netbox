// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*fhrpGroupAssignmentsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*fhrpGroupAssignmentsDataSource)(nil)
)

// NewFhrpGroupAssignmentsDataSource returns a new fhrp_group_assignments data source, which
// lists fhrp_group_assignment objects matching its filters.
func NewFhrpGroupAssignmentsDataSource() datasource.DataSource {
	return &fhrpGroupAssignmentsDataSource{}
}

type fhrpGroupAssignmentsDataSource struct {
	client *netboxapi.Client
}

// fhrpGroupAssignmentsDataSourceModel is the data source model: the filters, the limit and the matching
// fhrp_group_assignments.
type fhrpGroupAssignmentsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"fhrp_group_assignments"`
}

// fhrpGroupAssignmentResourceAttrTypes is the attribute type map of fhrpGroupAssignmentResourceModel.
func fhrpGroupAssignmentResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":             types.Int64Type,
		"group_id":       types.Int64Type,
		"interface_type": types.StringType,
		"interface_id":   types.Int64Type,
		"priority":       types.Int64Type,
		"created":        types.StringType,
		"last_updated":   types.StringType,
		"url":            types.StringType,
	}
}

func (d *fhrpGroupAssignmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_fhrp_group_assignments"
}

func (d *fhrpGroupAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists fhrp_group_assignment objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, device, device_id, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_id, interface_id__empty, interface_id__gt, interface_id__gte, interface_id__lt, interface_id__lte, interface_id__n, interface_type, interface_type__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, priority, priority__empty, priority__gt, priority__gte, priority__lt, priority__lte, priority__n, virtual_machine, virtual_machine_id. Repeating a name sends that parameter once per value.",
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
			"fhrp_group_assignments": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching fhrp_group_assignments.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"group_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the FHRP group.",
						},
						"interface_type": schema.StringAttribute{
							Computed:    true,
							Description: "Type of the interface the group is assigned to, e.g. dcim.interface or virtualization.vminterface.",
						},
						"interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the interface named by interface_type.",
						},
						"priority": schema.Int64Attribute{
							Computed:    true,
							Description: "Priority of this interface within the group (0-255). Must be a whole number.",
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
				},
			},
		},
	}
}

func (d *fhrpGroupAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *fhrpGroupAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data fhrpGroupAssignmentsDataSourceModel

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
	var responseDTOs []*netboxapi.FhrpGroupAssignmentResponseDTO

	params := ipam.NewIpamFhrpGroupAssignmentsListParams()
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
		"interface_type":      netboxapi.FilterStrings(params.SetInterfaceType),
		"interface_type__n":   netboxapi.FilterStrings(params.SetInterfaceTypen),
		"interface_id":        netboxapi.FilterInt64s(params.SetInterfaceID),
		"interface_id__n":     netboxapi.FilterInt64s(params.SetInterfaceIDn),
		"interface_id__lt":    netboxapi.FilterInt64s(params.SetInterfaceIDLt),
		"interface_id__lte":   netboxapi.FilterInt64s(params.SetInterfaceIDLte),
		"interface_id__gt":    netboxapi.FilterInt64s(params.SetInterfaceIDGt),
		"interface_id__gte":   netboxapi.FilterInt64s(params.SetInterfaceIDGte),
		"interface_id__empty": netboxapi.FilterBool(params.SetInterfaceIDEmpty),
		"priority":            netboxapi.FilterInt64s(params.SetPriority),
		"priority__n":         netboxapi.FilterInt64s(params.SetPriorityn),
		"priority__lt":        netboxapi.FilterInt64s(params.SetPriorityLt),
		"priority__lte":       netboxapi.FilterInt64s(params.SetPriorityLte),
		"priority__gt":        netboxapi.FilterInt64s(params.SetPriorityGt),
		"priority__gte":       netboxapi.FilterInt64s(params.SetPriorityGte),
		"priority__empty":     netboxapi.FilterBool(params.SetPriorityEmpty),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"device":              netboxapi.FilterStrings(params.SetDevice),
		"device_id":           netboxapi.FilterInt64s(params.SetDeviceID),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"virtual_machine":     netboxapi.FilterStrings(params.SetVirtualMachine),
		"virtual_machine_id":  netboxapi.FilterInt64s(params.SetVirtualMachineID),
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
		res, err := d.client.Ipam.IpamFhrpGroupAssignmentsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_fhrp_group_assignments", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.FhrpGroupAssignmentResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]fhrpGroupAssignmentResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model fhrpGroupAssignmentResourceModel
		resp.Diagnostics.Append(flattenFhrpGroupAssignment(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: fhrpGroupAssignmentResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
