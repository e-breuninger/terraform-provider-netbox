// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*contactAssignmentsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*contactAssignmentsDataSource)(nil)
)

// NewContactAssignmentsDataSource returns a new contact_assignments data source, which
// lists contact_assignment objects matching its filters.
func NewContactAssignmentsDataSource() datasource.DataSource {
	return &contactAssignmentsDataSource{}
}

type contactAssignmentsDataSource struct {
	client *netboxapi.Client
}

// contactAssignmentsDataSourceModel is the data source model: the filters, the limit and the matching
// contact_assignments.
type contactAssignmentsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"contact_assignments"`
}

// contactAssignmentResourceAttrTypes is the attribute type map of contactAssignmentResourceModel.
func contactAssignmentResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"object_type":   types.StringType,
		"object_id":     types.Int64Type,
		"contact_id":    types.Int64Type,
		"role_id":       types.Int64Type,
		"priority":      types.StringType,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *contactAssignmentsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_contact_assignments"
}

func (d *contactAssignmentsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Tenancy:Lists contact_assignment objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: contact_id, contact_id__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, object_id, object_id__empty, object_id__gt, object_id__gte, object_id__lt, object_id__lte, object_id__n, object_type, object_type__n, object_type_id, object_type_id__n, priority, priority__empty, priority__ic, priority__ie, priority__iew, priority__iregex, priority__isw, priority__n, priority__nic, priority__nie, priority__niew, priority__nisw, priority__regex, q, role, role__n, role_id, role_id__n, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"contact_assignments": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching contact_assignments.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"object_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the object the contact is assigned to (dcim.site, dcim.device, tenancy.tenant, ...).",
						},
						"object_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the object the contact is assigned to.",
						},
						"contact_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the contact.",
						},
						"role_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the contact role. Once set it cannot be removed, only changed: NetBox rejects a null. If you drop it from the config, keep a depends_on on the role so it is destroyed after the assignment.",
						},
						"priority": schema.StringAttribute{
							Computed:    true,
							Description: "Priority of the contact for this object. One of: primary, secondary, tertiary, inactive.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *contactAssignmentsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *contactAssignmentsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data contactAssignmentsDataSourceModel

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
	var responseDTOs []*netboxapi.ContactAssignmentResponseDTO

	params := tenancy.NewTenancyContactAssignmentsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"object_type":         netboxapi.FilterStrings(params.SetObjectType),
		"object_type__n":      netboxapi.FilterStrings(params.SetObjectTypen),
		"object_id":           netboxapi.FilterInt64s(params.SetObjectID),
		"object_id__n":        netboxapi.FilterInt64s(params.SetObjectIDn),
		"object_id__lt":       netboxapi.FilterInt64s(params.SetObjectIDLt),
		"object_id__lte":      netboxapi.FilterInt64s(params.SetObjectIDLte),
		"object_id__gt":       netboxapi.FilterInt64s(params.SetObjectIDGt),
		"object_id__gte":      netboxapi.FilterInt64s(params.SetObjectIDGte),
		"object_id__empty":    netboxapi.FilterBool(params.SetObjectIDEmpty),
		"contact_id":          netboxapi.FilterInt64s(params.SetContactID),
		"contact_id__n":       netboxapi.FilterInt64s(params.SetContactIDn),
		"role_id":             netboxapi.FilterInt64s(params.SetRoleID),
		"role_id__n":          netboxapi.FilterInt64s(params.SetRoleIDn),
		"tag":                 netboxapi.FilterStrings(params.SetTag),
		"tag__n":              netboxapi.FilterStrings(params.SetTagn),
		"tag__any":            netboxapi.FilterStrings(params.SetTagAny),
		"created":             netboxapi.FilterStrings(params.SetCreated),
		"created__n":          netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":         netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":        netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":         netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":        netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":      netboxapi.FilterStrings(params.SetCreatedEmpty),
		"group":               netboxapi.FilterStrings(params.SetGroup),
		"group__n":            netboxapi.FilterStrings(params.SetGroupn),
		"group_id":            netboxapi.FilterStrings(params.SetGroupID),
		"group_id__n":         netboxapi.FilterStrings(params.SetGroupIDn),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"object_type_id":      netboxapi.FilterInt64(params.SetObjectTypeID),
		"object_type_id__n":   netboxapi.FilterInt64(params.SetObjectTypeIDn),
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
		"q":                   netboxapi.FilterString(params.SetQ),
		"role":                netboxapi.FilterStrings(params.SetRole),
		"role__n":             netboxapi.FilterStrings(params.SetRolen),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
	}
	customFieldQuery := url.Values{}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		// Custom fields are queried as cf_<field name> with the field's own filter logic; NetBox ignores
		// names it has no custom field for.
		if strings.HasPrefix(name, "cf_") {
			customFieldQuery.Add(name, value)
			continue
		}
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
		res, err := d.client.Tenancy.TenancyContactAssignmentsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_contact_assignments", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ContactAssignmentResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]contactAssignmentResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model contactAssignmentResourceModel
		resp.Diagnostics.Append(flattenContactAssignment(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: contactAssignmentResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
