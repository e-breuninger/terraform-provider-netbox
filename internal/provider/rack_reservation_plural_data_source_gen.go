// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*rackReservationsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*rackReservationsDataSource)(nil)
)

// NewRackReservationsDataSource returns a new rack_reservations data source, which
// lists rack_reservation objects matching its filters.
func NewRackReservationsDataSource() datasource.DataSource {
	return &rackReservationsDataSource{}
}

type rackReservationsDataSource struct {
	client *netboxapi.Client
}

// rackReservationsDataSourceModel is the data source model: the filters, the limit and the matching
// rack_reservations.
type rackReservationsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"rack_reservations"`
}

// rackReservationResourceAttrTypes is the attribute type map of rackReservationResourceModel.
func rackReservationResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"rack_id":       types.Int64Type,
		"units":         types.SetType{ElemType: types.Int64Type},
		"user_id":       types.Int64Type,
		"tenant_id":     types.Int64Type,
		"status":        types.StringType,
		"description":   types.StringType,
		"comments":      types.StringType,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"unit_count":    types.Int64Type,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *rackReservationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_rack_reservations"
}

func (d *rackReservationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists rack_reservation objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, group, group__n, group_id, group_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location, location__n, location_id, location_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, q, rack_id, rack_id__n, region, region__n, region_id, region_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, unit, unit__empty, unit__gt, unit__gte, unit__lt, unit__lte, unit__n, unit_count_max, unit_count_min, user, user__n, user_id, user_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"rack_reservations": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching rack_reservations.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"rack_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the rack.",
						},
						"units": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Rack unit numbers being reserved.",
						},
						"user_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the NetBox user holding the reservation.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Status of the reservation. One of: pending, active, stale.",
						},
						"description": schema.StringAttribute{
							Computed:    true,
							Description: "Purpose of the reservation.",
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"owner_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the owner the object is assigned to.",
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
						"unit_count": schema.Int64Attribute{
							Computed:    true,
							Description: "Number of rack units reserved.",
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

func (d *rackReservationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *rackReservationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data rackReservationsDataSourceModel

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
	var responseDTOs []*netboxapi.RackReservationResponseDTO

	params := dcim.NewDcimRackReservationsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                  netboxapi.FilterInt64s(params.SetID),
		"id__n":               netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":              netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":             netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":              netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":             netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":           netboxapi.FilterBool(params.SetIDEmpty),
		"owner_id":            netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":         netboxapi.FilterInt64s(params.SetOwnerIDn),
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
		"description":         netboxapi.FilterStrings(params.SetDescription),
		"description__n":      netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":     netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":    netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":     netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":    netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":    netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":   netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":    netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":   netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":  netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":  netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex": netboxapi.FilterStrings(params.SetDescriptionIregex),
		"group":               netboxapi.FilterStrings(params.SetGroup),
		"group__n":            netboxapi.FilterStrings(params.SetGroupn),
		"group_id":            netboxapi.FilterInt64s(params.SetGroupID),
		"group_id__n":         netboxapi.FilterInt64s(params.SetGroupIDn),
		"last_updated":        netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":     netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":    netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":   netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":    netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":   netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty": netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location":            netboxapi.FilterStrings(params.SetLocation),
		"location__n":         netboxapi.FilterStrings(params.SetLocationn),
		"location_id":         netboxapi.FilterStrings(params.SetLocationID),
		"location_id__n":      netboxapi.FilterStrings(params.SetLocationIDn),
		"owner":               netboxapi.FilterStrings(params.SetOwner),
		"owner__n":            netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":         netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":      netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":      netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":   netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"q":                   netboxapi.FilterString(params.SetQ),
		"rack_id":             netboxapi.FilterInt64s(params.SetRackID),
		"rack_id__n":          netboxapi.FilterInt64s(params.SetRackIDn),
		"region":              netboxapi.FilterStrings(params.SetRegion),
		"region__n":           netboxapi.FilterStrings(params.SetRegionn),
		"region_id":           netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":        netboxapi.FilterStrings(params.SetRegionIDn),
		"site":                netboxapi.FilterStrings(params.SetSite),
		"site__n":             netboxapi.FilterStrings(params.SetSiten),
		"site_group":          netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":       netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":       netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":    netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"site_id":             netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":          netboxapi.FilterInt64s(params.SetSiteIDn),
		"status":              netboxapi.FilterStrings(params.SetStatus),
		"status__n":           netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":          netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":         netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":          netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":         netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":         netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":        netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":         netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":        netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":       netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":       netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":      netboxapi.FilterStrings(params.SetStatusIregex),
		"tag_id":              netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":           netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":         netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":              netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":           netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":        netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":     netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":     netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":  netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"tenant_id":           netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":        netboxapi.FilterInt64s(params.SetTenantIDn),
		"unit":                netboxapi.FilterFloat64(params.SetUnit),
		"unit__n":             netboxapi.FilterFloat64(params.SetUnitn),
		"unit__lt":            netboxapi.FilterFloat64(params.SetUnitLt),
		"unit__lte":           netboxapi.FilterFloat64(params.SetUnitLte),
		"unit__gt":            netboxapi.FilterFloat64(params.SetUnitGt),
		"unit__gte":           netboxapi.FilterFloat64(params.SetUnitGte),
		"unit__empty":         netboxapi.FilterFloat64(params.SetUnitEmpty),
		"unit_count_max":      netboxapi.FilterFloat64(params.SetUnitCountMax),
		"unit_count_min":      netboxapi.FilterFloat64(params.SetUnitCountMin),
		"user":                netboxapi.FilterStrings(params.SetUser),
		"user__n":             netboxapi.FilterStrings(params.SetUsern),
		"user_id":             netboxapi.FilterInt64s(params.SetUserID),
		"user_id__n":          netboxapi.FilterInt64s(params.SetUserIDn),
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
		res, err := d.client.Dcim.DcimRackReservationsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_rack_reservations", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.RackReservationResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]rackReservationResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model rackReservationResourceModel
		resp.Diagnostics.Append(flattenRackReservation(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: rackReservationResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
