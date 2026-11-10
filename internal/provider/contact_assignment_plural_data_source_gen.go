// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vObjectType []string
	var vObjectTypen []string
	var vObjectID []int64
	var vObjectIDn []int64
	var vObjectIDLt []int64
	var vObjectIDLte []int64
	var vObjectIDGt []int64
	var vObjectIDGte []int64
	var vObjectIDEmpty *bool
	var vContactID []int64
	var vContactIDn []int64
	var vRoleID []int64
	var vRoleIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vGroup []string
	var vGroupn []string
	var vGroupID []string
	var vGroupIDn []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vObjectTypeID *int64
	var vObjectTypeIDn *int64
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
	var vQ *string
	var vRole []string
	var vRolen []string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	customFieldQuery := url.Values{}
	for _, filter := range filters {
		name, value := filter.Name.ValueString(), filter.Value.ValueString()
		// Custom fields are queried as cf_<field name> with the field's own filter logic; NetBox ignores
		// names it has no custom field for.
		if strings.HasPrefix(name, "cf_") {
			customFieldQuery.Add(name, value)
			continue
		}
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
		case "object_type":
			v := value
			vObjectType = append(vObjectType, v)
		case "object_type__n":
			v := value
			vObjectTypen = append(vObjectTypen, v)
		case "object_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_id' takes an integer, got %q.", value))
				return
			}
			vObjectID = append(vObjectID, v)
		case "object_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_id__n' takes an integer, got %q.", value))
				return
			}
			vObjectIDn = append(vObjectIDn, v)
		case "object_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_id__lt' takes an integer, got %q.", value))
				return
			}
			vObjectIDLt = append(vObjectIDLt, v)
		case "object_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_id__lte' takes an integer, got %q.", value))
				return
			}
			vObjectIDLte = append(vObjectIDLte, v)
		case "object_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_id__gt' takes an integer, got %q.", value))
				return
			}
			vObjectIDGt = append(vObjectIDGt, v)
		case "object_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_id__gte' takes an integer, got %q.", value))
				return
			}
			vObjectIDGte = append(vObjectIDGte, v)
		case "object_id__empty":
			if vObjectIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'object_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_id__empty' takes a boolean, got %q.", value))
				return
			}
			vObjectIDEmpty = &v
		case "contact_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_id' takes an integer, got %q.", value))
				return
			}
			vContactID = append(vContactID, v)
		case "contact_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_id__n' takes an integer, got %q.", value))
				return
			}
			vContactIDn = append(vContactIDn, v)
		case "role_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id' takes an integer, got %q.", value))
				return
			}
			vRoleID = append(vRoleID, v)
		case "role_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role_id__n' takes an integer, got %q.", value))
				return
			}
			vRoleIDn = append(vRoleIDn, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
		case "created":
			v := value
			vCreated = append(vCreated, v)
		case "created__n":
			v := value
			vCreatedn = append(vCreatedn, v)
		case "created__lt":
			v := value
			vCreatedLt = append(vCreatedLt, v)
		case "created__lte":
			v := value
			vCreatedLte = append(vCreatedLte, v)
		case "created__gt":
			v := value
			vCreatedGt = append(vCreatedGt, v)
		case "created__gte":
			v := value
			vCreatedGte = append(vCreatedGte, v)
		case "created__empty":
			v := value
			vCreatedEmpty = append(vCreatedEmpty, v)
		case "group":
			v := value
			vGroup = append(vGroup, v)
		case "group__n":
			v := value
			vGroupn = append(vGroupn, v)
		case "group_id":
			v := value
			vGroupID = append(vGroupID, v)
		case "group_id__n":
			v := value
			vGroupIDn = append(vGroupIDn, v)
		case "last_updated":
			v := value
			vLastUpdated = append(vLastUpdated, v)
		case "last_updated__n":
			v := value
			vLastUpdatedn = append(vLastUpdatedn, v)
		case "last_updated__lt":
			v := value
			vLastUpdatedLt = append(vLastUpdatedLt, v)
		case "last_updated__lte":
			v := value
			vLastUpdatedLte = append(vLastUpdatedLte, v)
		case "last_updated__gt":
			v := value
			vLastUpdatedGt = append(vLastUpdatedGt, v)
		case "last_updated__gte":
			v := value
			vLastUpdatedGte = append(vLastUpdatedGte, v)
		case "last_updated__empty":
			v := value
			vLastUpdatedEmpty = append(vLastUpdatedEmpty, v)
		case "object_type_id":
			if vObjectTypeID != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'object_type_id' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id' takes an integer, got %q.", value))
				return
			}
			vObjectTypeID = &v
		case "object_type_id__n":
			if vObjectTypeIDn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'object_type_id__n' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id__n' takes an integer, got %q.", value))
				return
			}
			vObjectTypeIDn = &v
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
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "role":
			v := value
			vRole = append(vRole, v)
		case "role__n":
			v := value
			vRolen = append(vRolen, v)
		case "tag_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id' takes an integer, got %q.", value))
				return
			}
			vTagID = append(vTagID, v)
		case "tag_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__n' takes an integer, got %q.", value))
				return
			}
			vTagIDn = append(vTagIDn, v)
		case "tag_id__any":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tag_id__any' takes an integer, got %q.", value))
				return
			}
			vTagIDAny = append(vTagIDAny, v)
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
	if len(vObjectType) > 0 {
		params.SetObjectType(vObjectType)
	}
	if len(vObjectTypen) > 0 {
		params.SetObjectTypen(vObjectTypen)
	}
	if len(vObjectID) > 0 {
		params.SetObjectID(vObjectID)
	}
	if len(vObjectIDn) > 0 {
		params.SetObjectIDn(vObjectIDn)
	}
	if len(vObjectIDLt) > 0 {
		params.SetObjectIDLt(vObjectIDLt)
	}
	if len(vObjectIDLte) > 0 {
		params.SetObjectIDLte(vObjectIDLte)
	}
	if len(vObjectIDGt) > 0 {
		params.SetObjectIDGt(vObjectIDGt)
	}
	if len(vObjectIDGte) > 0 {
		params.SetObjectIDGte(vObjectIDGte)
	}
	if vObjectIDEmpty != nil {
		params.SetObjectIDEmpty(vObjectIDEmpty)
	}
	if len(vContactID) > 0 {
		params.SetContactID(vContactID)
	}
	if len(vContactIDn) > 0 {
		params.SetContactIDn(vContactIDn)
	}
	if len(vRoleID) > 0 {
		params.SetRoleID(vRoleID)
	}
	if len(vRoleIDn) > 0 {
		params.SetRoleIDn(vRoleIDn)
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
	if len(vCreated) > 0 {
		params.SetCreated(vCreated)
	}
	if len(vCreatedn) > 0 {
		params.SetCreatedn(vCreatedn)
	}
	if len(vCreatedLt) > 0 {
		params.SetCreatedLt(vCreatedLt)
	}
	if len(vCreatedLte) > 0 {
		params.SetCreatedLte(vCreatedLte)
	}
	if len(vCreatedGt) > 0 {
		params.SetCreatedGt(vCreatedGt)
	}
	if len(vCreatedGte) > 0 {
		params.SetCreatedGte(vCreatedGte)
	}
	if len(vCreatedEmpty) > 0 {
		params.SetCreatedEmpty(vCreatedEmpty)
	}
	if len(vGroup) > 0 {
		params.SetGroup(vGroup)
	}
	if len(vGroupn) > 0 {
		params.SetGroupn(vGroupn)
	}
	if len(vGroupID) > 0 {
		params.SetGroupID(vGroupID)
	}
	if len(vGroupIDn) > 0 {
		params.SetGroupIDn(vGroupIDn)
	}
	if len(vLastUpdated) > 0 {
		params.SetLastUpdated(vLastUpdated)
	}
	if len(vLastUpdatedn) > 0 {
		params.SetLastUpdatedn(vLastUpdatedn)
	}
	if len(vLastUpdatedLt) > 0 {
		params.SetLastUpdatedLt(vLastUpdatedLt)
	}
	if len(vLastUpdatedLte) > 0 {
		params.SetLastUpdatedLte(vLastUpdatedLte)
	}
	if len(vLastUpdatedGt) > 0 {
		params.SetLastUpdatedGt(vLastUpdatedGt)
	}
	if len(vLastUpdatedGte) > 0 {
		params.SetLastUpdatedGte(vLastUpdatedGte)
	}
	if len(vLastUpdatedEmpty) > 0 {
		params.SetLastUpdatedEmpty(vLastUpdatedEmpty)
	}
	if vObjectTypeID != nil {
		params.SetObjectTypeID(vObjectTypeID)
	}
	if vObjectTypeIDn != nil {
		params.SetObjectTypeIDn(vObjectTypeIDn)
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
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRole) > 0 {
		params.SetRole(vRole)
	}
	if len(vRolen) > 0 {
		params.SetRolen(vRolen)
	}
	if len(vTagID) > 0 {
		params.SetTagID(vTagID)
	}
	if len(vTagIDn) > 0 {
		params.SetTagIDn(vTagIDn)
	}
	if len(vTagIDAny) > 0 {
		params.SetTagIDAny(vTagIDAny)
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
