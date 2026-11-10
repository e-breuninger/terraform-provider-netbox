// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*permissionsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*permissionsDataSource)(nil)
)

// NewPermissionsDataSource returns a new permissions data source, which
// lists permission objects matching its filters.
func NewPermissionsDataSource() datasource.DataSource {
	return &permissionsDataSource{}
}

type permissionsDataSource struct {
	client *netboxapi.Client
}

// permissionsDataSourceModel is the data source model: the filters, the limit and the matching
// permissions.
type permissionsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"permissions"`
}

// permissionResourceAttrTypes is the attribute type map of permissionResourceModel.
func permissionResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":           types.Int64Type,
		"name":         types.StringType,
		"description":  types.StringType,
		"enabled":      types.BoolType,
		"object_types": types.SetType{ElemType: types.StringType},
		"actions":      types.SetType{ElemType: types.StringType},
		"group_ids":    types.SetType{ElemType: types.Int64Type},
		"groups":       types.SetType{ElemType: types.Int64Type},
		"user_ids":     types.SetType{ElemType: types.Int64Type},
		"users":        types.SetType{ElemType: types.Int64Type},
		"constraints":  jsontypes.NormalizedType{},
		"url":          types.StringType,
	}
}

func (d *permissionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_permissions"
}

func (d *permissionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:Lists permission objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: can_add, can_change, can_delete, can_view, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, enabled, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, object_type, object_type__ic, object_type__ie, object_type__iew, object_type__iregex, object_type__isw, object_type__n, object_type__nic, object_type__nie, object_type__niew, object_type__nisw, object_type__regex, object_type_id, object_type_id__n, object_types, object_types__n, q, user, user__n, user_id, user_id__n. Repeating a name sends that parameter once per value.",
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
			"permissions": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching permissions.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the permission is in effect.",
						},
						"object_types": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Object types the permission applies to, as app-labelled content types (dcim.device).",
						},
						"actions": schema.SetAttribute{
							ElementType: types.StringType,
							Computed:    true,
							Description: "Permitted actions: view, add, change, delete or a custom action name.",
						},
						"group_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the groups the permission is granted to.",
						},
						"groups": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use group_ids instead.",
							Description:        "Deprecated alias of group_ids.",
						},
						"user_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the users the permission is granted to.",
						},
						"users": schema.SetAttribute{
							ElementType:        types.Int64Type,
							Computed:           true,
							DeprecationMessage: "Use user_ids instead.",
							Description:        "Deprecated alias of user_ids.",
						},
						"constraints": schema.StringAttribute{
							CustomType:  jsontypes.NormalizedType{},
							Computed:    true,
							Description: "Queryset filter as JSON text, an object or a list of objects (use jsonencode()).",
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

func (d *permissionsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *permissionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data permissionsDataSourceModel

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
	var responseDTOs []*netboxapi.PermissionResponseDTO

	params := users.NewUsersPermissionsListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vName []string
	var vNamen []string
	var vNameIc []string
	var vNameNic []string
	var vNameIe []string
	var vNameNie []string
	var vNameIsw []string
	var vNameNisw []string
	var vNameIew []string
	var vNameNiew []string
	var vNameEmpty *bool
	var vNameRegex []string
	var vNameIregex []string
	var vCanAdd *bool
	var vCanChange *bool
	var vCanDelete *bool
	var vCanView *bool
	var vDescription []string
	var vDescriptionn []string
	var vDescriptionIc []string
	var vDescriptionNic []string
	var vDescriptionIe []string
	var vDescriptionNie []string
	var vDescriptionIsw []string
	var vDescriptionNisw []string
	var vDescriptionIew []string
	var vDescriptionNiew []string
	var vDescriptionEmpty *bool
	var vDescriptionRegex []string
	var vDescriptionIregex []string
	var vEnabled *bool
	var vGroup []string
	var vGroupn []string
	var vGroupID []int64
	var vGroupIDn []int64
	var vObjectType []string
	var vObjectTypen []string
	var vObjectTypeIc []string
	var vObjectTypeNic []string
	var vObjectTypeIe []string
	var vObjectTypeNie []string
	var vObjectTypeIsw []string
	var vObjectTypeNisw []string
	var vObjectTypeIew []string
	var vObjectTypeNiew []string
	var vObjectTypeRegex []string
	var vObjectTypeIregex []string
	var vObjectTypeID []int64
	var vObjectTypeIDn []int64
	var vObjectTypes []int64
	var vObjectTypesn []int64
	var vQ *string
	var vUser []string
	var vUsern []string
	var vUserID []int64
	var vUserIDn []int64
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
		case "name":
			v := value
			vName = append(vName, v)
		case "name__n":
			v := value
			vNamen = append(vNamen, v)
		case "name__ic":
			v := value
			vNameIc = append(vNameIc, v)
		case "name__nic":
			v := value
			vNameNic = append(vNameNic, v)
		case "name__ie":
			v := value
			vNameIe = append(vNameIe, v)
		case "name__nie":
			v := value
			vNameNie = append(vNameNie, v)
		case "name__isw":
			v := value
			vNameIsw = append(vNameIsw, v)
		case "name__nisw":
			v := value
			vNameNisw = append(vNameNisw, v)
		case "name__iew":
			v := value
			vNameIew = append(vNameIew, v)
		case "name__niew":
			v := value
			vNameNiew = append(vNameNiew, v)
		case "name__empty":
			if vNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'name__empty' takes a boolean, got %q.", value))
				return
			}
			vNameEmpty = &v
		case "name__regex":
			v := value
			vNameRegex = append(vNameRegex, v)
		case "name__iregex":
			v := value
			vNameIregex = append(vNameIregex, v)
		case "can_add":
			if vCanAdd != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'can_add' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'can_add' takes a boolean, got %q.", value))
				return
			}
			vCanAdd = &v
		case "can_change":
			if vCanChange != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'can_change' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'can_change' takes a boolean, got %q.", value))
				return
			}
			vCanChange = &v
		case "can_delete":
			if vCanDelete != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'can_delete' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'can_delete' takes a boolean, got %q.", value))
				return
			}
			vCanDelete = &v
		case "can_view":
			if vCanView != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'can_view' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'can_view' takes a boolean, got %q.", value))
				return
			}
			vCanView = &v
		case "description":
			v := value
			vDescription = append(vDescription, v)
		case "description__n":
			v := value
			vDescriptionn = append(vDescriptionn, v)
		case "description__ic":
			v := value
			vDescriptionIc = append(vDescriptionIc, v)
		case "description__nic":
			v := value
			vDescriptionNic = append(vDescriptionNic, v)
		case "description__ie":
			v := value
			vDescriptionIe = append(vDescriptionIe, v)
		case "description__nie":
			v := value
			vDescriptionNie = append(vDescriptionNie, v)
		case "description__isw":
			v := value
			vDescriptionIsw = append(vDescriptionIsw, v)
		case "description__nisw":
			v := value
			vDescriptionNisw = append(vDescriptionNisw, v)
		case "description__iew":
			v := value
			vDescriptionIew = append(vDescriptionIew, v)
		case "description__niew":
			v := value
			vDescriptionNiew = append(vDescriptionNiew, v)
		case "description__empty":
			if vDescriptionEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'description__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'description__empty' takes a boolean, got %q.", value))
				return
			}
			vDescriptionEmpty = &v
		case "description__regex":
			v := value
			vDescriptionRegex = append(vDescriptionRegex, v)
		case "description__iregex":
			v := value
			vDescriptionIregex = append(vDescriptionIregex, v)
		case "enabled":
			if vEnabled != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'enabled' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'enabled' takes a boolean, got %q.", value))
				return
			}
			vEnabled = &v
		case "group":
			v := value
			vGroup = append(vGroup, v)
		case "group__n":
			v := value
			vGroupn = append(vGroupn, v)
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
		case "object_type":
			v := value
			vObjectType = append(vObjectType, v)
		case "object_type__n":
			v := value
			vObjectTypen = append(vObjectTypen, v)
		case "object_type__ic":
			v := value
			vObjectTypeIc = append(vObjectTypeIc, v)
		case "object_type__nic":
			v := value
			vObjectTypeNic = append(vObjectTypeNic, v)
		case "object_type__ie":
			v := value
			vObjectTypeIe = append(vObjectTypeIe, v)
		case "object_type__nie":
			v := value
			vObjectTypeNie = append(vObjectTypeNie, v)
		case "object_type__isw":
			v := value
			vObjectTypeIsw = append(vObjectTypeIsw, v)
		case "object_type__nisw":
			v := value
			vObjectTypeNisw = append(vObjectTypeNisw, v)
		case "object_type__iew":
			v := value
			vObjectTypeIew = append(vObjectTypeIew, v)
		case "object_type__niew":
			v := value
			vObjectTypeNiew = append(vObjectTypeNiew, v)
		case "object_type__regex":
			v := value
			vObjectTypeRegex = append(vObjectTypeRegex, v)
		case "object_type__iregex":
			v := value
			vObjectTypeIregex = append(vObjectTypeIregex, v)
		case "object_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id' takes an integer, got %q.", value))
				return
			}
			vObjectTypeID = append(vObjectTypeID, v)
		case "object_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_type_id__n' takes an integer, got %q.", value))
				return
			}
			vObjectTypeIDn = append(vObjectTypeIDn, v)
		case "object_types":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_types' takes an integer, got %q.", value))
				return
			}
			vObjectTypes = append(vObjectTypes, v)
		case "object_types__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'object_types__n' takes an integer, got %q.", value))
				return
			}
			vObjectTypesn = append(vObjectTypesn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "user":
			v := value
			vUser = append(vUser, v)
		case "user__n":
			v := value
			vUsern = append(vUsern, v)
		case "user_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'user_id' takes an integer, got %q.", value))
				return
			}
			vUserID = append(vUserID, v)
		case "user_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'user_id__n' takes an integer, got %q.", value))
				return
			}
			vUserIDn = append(vUserIDn, v)
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
	if len(vName) > 0 {
		params.SetName(vName)
	}
	if len(vNamen) > 0 {
		params.SetNamen(vNamen)
	}
	if len(vNameIc) > 0 {
		params.SetNameIc(vNameIc)
	}
	if len(vNameNic) > 0 {
		params.SetNameNic(vNameNic)
	}
	if len(vNameIe) > 0 {
		params.SetNameIe(vNameIe)
	}
	if len(vNameNie) > 0 {
		params.SetNameNie(vNameNie)
	}
	if len(vNameIsw) > 0 {
		params.SetNameIsw(vNameIsw)
	}
	if len(vNameNisw) > 0 {
		params.SetNameNisw(vNameNisw)
	}
	if len(vNameIew) > 0 {
		params.SetNameIew(vNameIew)
	}
	if len(vNameNiew) > 0 {
		params.SetNameNiew(vNameNiew)
	}
	if vNameEmpty != nil {
		params.SetNameEmpty(vNameEmpty)
	}
	if len(vNameRegex) > 0 {
		params.SetNameRegex(vNameRegex)
	}
	if len(vNameIregex) > 0 {
		params.SetNameIregex(vNameIregex)
	}
	if vCanAdd != nil {
		params.SetCanAdd(vCanAdd)
	}
	if vCanChange != nil {
		params.SetCanChange(vCanChange)
	}
	if vCanDelete != nil {
		params.SetCanDelete(vCanDelete)
	}
	if vCanView != nil {
		params.SetCanView(vCanView)
	}
	if len(vDescription) > 0 {
		params.SetDescription(vDescription)
	}
	if len(vDescriptionn) > 0 {
		params.SetDescriptionn(vDescriptionn)
	}
	if len(vDescriptionIc) > 0 {
		params.SetDescriptionIc(vDescriptionIc)
	}
	if len(vDescriptionNic) > 0 {
		params.SetDescriptionNic(vDescriptionNic)
	}
	if len(vDescriptionIe) > 0 {
		params.SetDescriptionIe(vDescriptionIe)
	}
	if len(vDescriptionNie) > 0 {
		params.SetDescriptionNie(vDescriptionNie)
	}
	if len(vDescriptionIsw) > 0 {
		params.SetDescriptionIsw(vDescriptionIsw)
	}
	if len(vDescriptionNisw) > 0 {
		params.SetDescriptionNisw(vDescriptionNisw)
	}
	if len(vDescriptionIew) > 0 {
		params.SetDescriptionIew(vDescriptionIew)
	}
	if len(vDescriptionNiew) > 0 {
		params.SetDescriptionNiew(vDescriptionNiew)
	}
	if vDescriptionEmpty != nil {
		params.SetDescriptionEmpty(vDescriptionEmpty)
	}
	if len(vDescriptionRegex) > 0 {
		params.SetDescriptionRegex(vDescriptionRegex)
	}
	if len(vDescriptionIregex) > 0 {
		params.SetDescriptionIregex(vDescriptionIregex)
	}
	if vEnabled != nil {
		params.SetEnabled(vEnabled)
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
	if len(vObjectType) > 0 {
		params.SetObjectType(vObjectType)
	}
	if len(vObjectTypen) > 0 {
		params.SetObjectTypen(vObjectTypen)
	}
	if len(vObjectTypeIc) > 0 {
		params.SetObjectTypeIc(vObjectTypeIc)
	}
	if len(vObjectTypeNic) > 0 {
		params.SetObjectTypeNic(vObjectTypeNic)
	}
	if len(vObjectTypeIe) > 0 {
		params.SetObjectTypeIe(vObjectTypeIe)
	}
	if len(vObjectTypeNie) > 0 {
		params.SetObjectTypeNie(vObjectTypeNie)
	}
	if len(vObjectTypeIsw) > 0 {
		params.SetObjectTypeIsw(vObjectTypeIsw)
	}
	if len(vObjectTypeNisw) > 0 {
		params.SetObjectTypeNisw(vObjectTypeNisw)
	}
	if len(vObjectTypeIew) > 0 {
		params.SetObjectTypeIew(vObjectTypeIew)
	}
	if len(vObjectTypeNiew) > 0 {
		params.SetObjectTypeNiew(vObjectTypeNiew)
	}
	if len(vObjectTypeRegex) > 0 {
		params.SetObjectTypeRegex(vObjectTypeRegex)
	}
	if len(vObjectTypeIregex) > 0 {
		params.SetObjectTypeIregex(vObjectTypeIregex)
	}
	if len(vObjectTypeID) > 0 {
		params.SetObjectTypeID(vObjectTypeID)
	}
	if len(vObjectTypeIDn) > 0 {
		params.SetObjectTypeIDn(vObjectTypeIDn)
	}
	if len(vObjectTypes) > 0 {
		params.SetObjectTypes(vObjectTypes)
	}
	if len(vObjectTypesn) > 0 {
		params.SetObjectTypesn(vObjectTypesn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vUser) > 0 {
		params.SetUser(vUser)
	}
	if len(vUsern) > 0 {
		params.SetUsern(vUsern)
	}
	if len(vUserID) > 0 {
		params.SetUserID(vUserID)
	}
	if len(vUserIDn) > 0 {
		params.SetUserIDn(vUserIDn)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Users.UsersPermissionsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_permissions", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.PermissionResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]permissionResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model permissionResourceModel
		resp.Diagnostics.Append(flattenPermission(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: permissionResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
