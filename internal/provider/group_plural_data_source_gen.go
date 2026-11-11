// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*groupsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*groupsDataSource)(nil)
)

// NewGroupsDataSource returns a new groups data source, which
// lists group objects matching its filters.
func NewGroupsDataSource() datasource.DataSource {
	return &groupsDataSource{}
}

type groupsDataSource struct {
	client *netboxapi.Client
}

// groupsDataSourceModel is the data source model: the filters, the limit and the matching
// groups.
type groupsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"groups"`
}

// groupResourceAttrTypes is the attribute type map of groupResourceModel.
func groupResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":          types.Int64Type,
		"name":        types.StringType,
		"description": types.StringType,
		"user_count":  types.Int64Type,
		"url":         types.StringType,
	}
}

func (d *groupsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_groups"
}

func (d *groupsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Authentication:Lists group objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, notification_group_id, notification_group_id__n, owner, owner__n, owner_id, owner_id__n, permission_id, permission_id__n, q, user_id, user_id__n. Repeating a name sends that parameter once per value.",
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
			"groups": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching groups.",
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
						"user_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of users in the group.",
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

func (d *groupsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *groupsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data groupsDataSourceModel

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
	var responseDTOs []*netboxapi.GroupResponseDTO

	params := users.NewUsersGroupsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                       netboxapi.FilterInt64s(params.SetID),
		"id__n":                    netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                   netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                  netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                   netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                  netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                netboxapi.FilterBool(params.SetIDEmpty),
		"name":                     netboxapi.FilterStrings(params.SetName),
		"name__n":                  netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                 netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                 netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":               netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":               netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":              netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":              netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":             netboxapi.FilterStrings(params.SetNameIregex),
		"description":              netboxapi.FilterStrings(params.SetDescription),
		"description__n":           netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":          netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":         netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":          netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":         netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":         netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":        netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":         netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":        netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":       netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":       netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":      netboxapi.FilterStrings(params.SetDescriptionIregex),
		"notification_group_id":    netboxapi.FilterInt64s(params.SetNotificationGroupID),
		"notification_group_id__n": netboxapi.FilterInt64s(params.SetNotificationGroupIDn),
		"owner":                    netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                 netboxapi.FilterStrings(params.SetOwnern),
		"owner_id":                 netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":              netboxapi.FilterInt64s(params.SetOwnerIDn),
		"permission_id":            netboxapi.FilterInt64s(params.SetPermissionID),
		"permission_id__n":         netboxapi.FilterInt64s(params.SetPermissionIDn),
		"q":                        netboxapi.FilterString(params.SetQ),
		"user_id":                  netboxapi.FilterInt64s(params.SetUserID),
		"user_id__n":               netboxapi.FilterInt64s(params.SetUserIDn),
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
		res, err := d.client.Users.UsersGroupsListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_groups", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.GroupResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]groupResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model groupResourceModel
		resp.Diagnostics.Append(flattenGroup(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: groupResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
