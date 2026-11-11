// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

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
	_ datasource.DataSource              = (*circuitsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*circuitsDataSource)(nil)
)

// NewCircuitsDataSource returns a new circuits data source, which
// lists circuit objects matching its filters.
func NewCircuitsDataSource() datasource.DataSource {
	return &circuitsDataSource{}
}

type circuitsDataSource struct {
	client *netboxapi.Client
}

// circuitsDataSourceModel is the data source model: the filters, the limit and the matching
// circuits.
type circuitsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"circuits"`
}

// circuitResourceAttrTypes is the attribute type map of circuitResourceModel.
func circuitResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                          types.Int64Type,
		"cid":                         types.StringType,
		"circuit_provider_id":         types.Int64Type,
		"circuit_provider_account_id": types.Int64Type,
		"circuit_type_id":             types.Int64Type,
		"status":                      types.StringType,
		"tenant_id":                   types.Int64Type,
		"install_date":                types.StringType,
		"termination_date":            types.StringType,
		"commit_rate_kbps":            types.Int64Type,
		"distance":                    types.Float64Type,
		"distance_unit":               types.StringType,
		"description":                 types.StringType,
		"comments":                    types.StringType,
		"owner_id":                    types.Int64Type,
		"created":                     types.StringType,
		"last_updated":                types.StringType,
		"url":                         types.StringType,
		"tags":                        types.SetType{ElemType: types.StringType},
		"tags_all":                    types.SetType{ElemType: types.StringType},
		"custom_fields":               types.MapType{ElemType: types.StringType},
	}
}

func (d *circuitsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_circuits"
}

func (d *circuitsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:Lists circuit objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: cid, cid__empty, cid__ic, cid__ie, cid__iew, cid__iregex, cid__isw, cid__n, cid__nic, cid__nie, cid__niew, cid__nisw, cid__regex, commit_rate, commit_rate__empty, commit_rate__gt, commit_rate__gte, commit_rate__lt, commit_rate__lte, commit_rate__n, contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, distance, distance__empty, distance__gt, distance__gte, distance__lt, distance__lte, distance__n, distance_unit, distance_unit__empty, distance_unit__ic, distance_unit__ie, distance_unit__iew, distance_unit__iregex, distance_unit__isw, distance_unit__n, distance_unit__nic, distance_unit__nie, distance_unit__niew, distance_unit__nisw, distance_unit__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, install_date, install_date__empty, install_date__gt, install_date__gte, install_date__lt, install_date__lte, install_date__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, location_id, location_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, provider, provider__n, provider_account, provider_account__n, provider_account_id, provider_account_id__n, provider_id, provider_id__n, provider_network_id, provider_network_id__n, q, region, region__n, region_id, region_id__n, site, site__n, site_group, site_group__n, site_group_id, site_group_id__n, site_id, site_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, termination_a_id, termination_a_id__n, termination_date, termination_date__empty, termination_date__gt, termination_date__gte, termination_date__lt, termination_date__lte, termination_date__n, termination_z_id, termination_z_id__n, type, type__n, type_id, type_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"circuits": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching circuits.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"cid": schema.StringAttribute{
							Computed:    true,
							Description: "Circuit id assigned by the provider; unique per provider.",
						},
						"circuit_provider_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the provider the circuit is leased from.",
						},
						"circuit_provider_account_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the provider account the circuit is billed to.",
						},
						"circuit_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the circuit type.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: planned, provisioning, active, offline, deprovisioning, decommissioned.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"install_date": schema.StringAttribute{
							Computed:    true,
							Description: "Installation date as YYYY-MM-DD.",
						},
						"termination_date": schema.StringAttribute{
							Computed:    true,
							Description: "Termination date as YYYY-MM-DD.",
						},
						"commit_rate_kbps": schema.Int64Attribute{
							Computed:    true,
							Description: "Committed rate in kbps. Must be a whole number.",
						},
						"distance": schema.Float64Attribute{
							Computed:    true,
							Description: "Length of the circuit (requires distance_unit).",
						},
						"distance_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of distance. One of: km, m, mi, ft.",
						},
						"description": schema.StringAttribute{
							Computed: true,
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

func (d *circuitsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *circuitsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data circuitsDataSourceModel

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
	var responseDTOs []*netboxapi.CircuitResponseDTO

	params := circuits.NewCircuitsCircuitsListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                      netboxapi.FilterInt64s(params.SetID),
		"id__n":                   netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                  netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                 netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                  netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                 netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":               netboxapi.FilterBool(params.SetIDEmpty),
		"cid":                     netboxapi.FilterStrings(params.SetCid),
		"cid__n":                  netboxapi.FilterStrings(params.SetCidn),
		"cid__ic":                 netboxapi.FilterStrings(params.SetCidIc),
		"cid__nic":                netboxapi.FilterStrings(params.SetCidNic),
		"cid__ie":                 netboxapi.FilterStrings(params.SetCidIe),
		"cid__nie":                netboxapi.FilterStrings(params.SetCidNie),
		"cid__isw":                netboxapi.FilterStrings(params.SetCidIsw),
		"cid__nisw":               netboxapi.FilterStrings(params.SetCidNisw),
		"cid__iew":                netboxapi.FilterStrings(params.SetCidIew),
		"cid__niew":               netboxapi.FilterStrings(params.SetCidNiew),
		"cid__empty":              netboxapi.FilterBool(params.SetCidEmpty),
		"cid__regex":              netboxapi.FilterStrings(params.SetCidRegex),
		"cid__iregex":             netboxapi.FilterStrings(params.SetCidIregex),
		"provider_id":             netboxapi.FilterInt64s(params.SetProviderID),
		"provider_id__n":          netboxapi.FilterInt64s(params.SetProviderIDn),
		"type_id":                 netboxapi.FilterInt64s(params.SetTypeID),
		"type_id__n":              netboxapi.FilterInt64s(params.SetTypeIDn),
		"status":                  netboxapi.FilterStrings(params.SetStatus),
		"status__n":               netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":              netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":             netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":              netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":             netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":             netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":            netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":             netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":            netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":           netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":           netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":          netboxapi.FilterStrings(params.SetStatusIregex),
		"owner_id":                netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":             netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                     netboxapi.FilterStrings(params.SetTag),
		"tag__n":                  netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                netboxapi.FilterStrings(params.SetTagAny),
		"commit_rate":             netboxapi.FilterInt64s(params.SetCommitRate),
		"commit_rate__n":          netboxapi.FilterInt64s(params.SetCommitRaten),
		"commit_rate__lt":         netboxapi.FilterInt64s(params.SetCommitRateLt),
		"commit_rate__lte":        netboxapi.FilterInt64s(params.SetCommitRateLte),
		"commit_rate__gt":         netboxapi.FilterInt64s(params.SetCommitRateGt),
		"commit_rate__gte":        netboxapi.FilterInt64s(params.SetCommitRateGte),
		"commit_rate__empty":      netboxapi.FilterBool(params.SetCommitRateEmpty),
		"contact":                 netboxapi.FilterInt64s(params.SetContact),
		"contact__n":              netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":           netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":        netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":            netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":         netboxapi.FilterInt64s(params.SetContactRolen),
		"created":                 netboxapi.FilterStrings(params.SetCreated),
		"created__n":              netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":             netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":            netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":             netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":            netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":          netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":             netboxapi.FilterStrings(params.SetDescription),
		"description__n":          netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":         netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":        netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":         netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":        netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":        netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":       netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":        netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":       netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":      netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":      netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":     netboxapi.FilterStrings(params.SetDescriptionIregex),
		"distance":                netboxapi.FilterFloat64s(params.SetDistance),
		"distance__n":             netboxapi.FilterFloat64s(params.SetDistancen),
		"distance__lt":            netboxapi.FilterFloat64s(params.SetDistanceLt),
		"distance__lte":           netboxapi.FilterFloat64s(params.SetDistanceLte),
		"distance__gt":            netboxapi.FilterFloat64s(params.SetDistanceGt),
		"distance__gte":           netboxapi.FilterFloat64s(params.SetDistanceGte),
		"distance__empty":         netboxapi.FilterBool(params.SetDistanceEmpty),
		"distance_unit":           netboxapi.FilterString(params.SetDistanceUnit),
		"distance_unit__n":        netboxapi.FilterString(params.SetDistanceUnitn),
		"distance_unit__ic":       netboxapi.FilterStrings(params.SetDistanceUnitIc),
		"distance_unit__nic":      netboxapi.FilterStrings(params.SetDistanceUnitNic),
		"distance_unit__ie":       netboxapi.FilterStrings(params.SetDistanceUnitIe),
		"distance_unit__nie":      netboxapi.FilterStrings(params.SetDistanceUnitNie),
		"distance_unit__isw":      netboxapi.FilterStrings(params.SetDistanceUnitIsw),
		"distance_unit__nisw":     netboxapi.FilterStrings(params.SetDistanceUnitNisw),
		"distance_unit__iew":      netboxapi.FilterStrings(params.SetDistanceUnitIew),
		"distance_unit__niew":     netboxapi.FilterStrings(params.SetDistanceUnitNiew),
		"distance_unit__empty":    netboxapi.FilterBool(params.SetDistanceUnitEmpty),
		"distance_unit__regex":    netboxapi.FilterStrings(params.SetDistanceUnitRegex),
		"distance_unit__iregex":   netboxapi.FilterStrings(params.SetDistanceUnitIregex),
		"install_date":            netboxapi.FilterStrings(params.SetInstallDate),
		"install_date__n":         netboxapi.FilterStrings(params.SetInstallDaten),
		"install_date__lt":        netboxapi.FilterStrings(params.SetInstallDateLt),
		"install_date__lte":       netboxapi.FilterStrings(params.SetInstallDateLte),
		"install_date__gt":        netboxapi.FilterStrings(params.SetInstallDateGt),
		"install_date__gte":       netboxapi.FilterStrings(params.SetInstallDateGte),
		"install_date__empty":     netboxapi.FilterBool(params.SetInstallDateEmpty),
		"last_updated":            netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":         netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":        netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":       netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":        netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":       netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":     netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"location_id":             netboxapi.FilterInt64s(params.SetLocationID),
		"location_id__n":          netboxapi.FilterInt64s(params.SetLocationIDn),
		"owner":                   netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":             netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":          netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":          netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":       netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"provider":                netboxapi.FilterStrings(params.SetProvider),
		"provider__n":             netboxapi.FilterStrings(params.SetProvidern),
		"provider_account":        netboxapi.FilterStrings(params.SetProviderAccount),
		"provider_account__n":     netboxapi.FilterStrings(params.SetProviderAccountn),
		"provider_account_id":     netboxapi.FilterInt64s(params.SetProviderAccountID),
		"provider_account_id__n":  netboxapi.FilterInt64s(params.SetProviderAccountIDn),
		"provider_network_id":     netboxapi.FilterInt64s(params.SetProviderNetworkID),
		"provider_network_id__n":  netboxapi.FilterInt64s(params.SetProviderNetworkIDn),
		"q":                       netboxapi.FilterString(params.SetQ),
		"region":                  netboxapi.FilterStrings(params.SetRegion),
		"region__n":               netboxapi.FilterStrings(params.SetRegionn),
		"region_id":               netboxapi.FilterStrings(params.SetRegionID),
		"region_id__n":            netboxapi.FilterStrings(params.SetRegionIDn),
		"site":                    netboxapi.FilterStrings(params.SetSite),
		"site__n":                 netboxapi.FilterStrings(params.SetSiten),
		"site_group":              netboxapi.FilterStrings(params.SetSiteGroup),
		"site_group__n":           netboxapi.FilterStrings(params.SetSiteGroupn),
		"site_group_id":           netboxapi.FilterStrings(params.SetSiteGroupID),
		"site_group_id__n":        netboxapi.FilterStrings(params.SetSiteGroupIDn),
		"site_id":                 netboxapi.FilterInt64s(params.SetSiteID),
		"site_id__n":              netboxapi.FilterInt64s(params.SetSiteIDn),
		"tag_id":                  netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":               netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":             netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                  netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":               netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":            netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":         netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":         netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":      netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"tenant_id":               netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":            netboxapi.FilterInt64s(params.SetTenantIDn),
		"termination_a_id":        netboxapi.FilterInt64s(params.SetTerminationaID),
		"termination_a_id__n":     netboxapi.FilterInt64s(params.SetTerminationaIDn),
		"termination_date":        netboxapi.FilterStrings(params.SetTerminationDate),
		"termination_date__n":     netboxapi.FilterStrings(params.SetTerminationDaten),
		"termination_date__lt":    netboxapi.FilterStrings(params.SetTerminationDateLt),
		"termination_date__lte":   netboxapi.FilterStrings(params.SetTerminationDateLte),
		"termination_date__gt":    netboxapi.FilterStrings(params.SetTerminationDateGt),
		"termination_date__gte":   netboxapi.FilterStrings(params.SetTerminationDateGte),
		"termination_date__empty": netboxapi.FilterBool(params.SetTerminationDateEmpty),
		"termination_z_id":        netboxapi.FilterInt64s(params.SetTerminationzID),
		"termination_z_id__n":     netboxapi.FilterInt64s(params.SetTerminationzIDn),
		"type":                    netboxapi.FilterStrings(params.SetType),
		"type__n":                 netboxapi.FilterStrings(params.SetTypen),
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
		res, err := d.client.Circuits.CircuitsCircuitsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_circuits", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CircuitResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]circuitResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model circuitResourceModel
		resp.Diagnostics.Append(flattenCircuit(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: circuitResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
