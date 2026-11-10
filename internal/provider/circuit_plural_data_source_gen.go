// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vCid []string
	var vCidn []string
	var vCidIc []string
	var vCidNic []string
	var vCidIe []string
	var vCidNie []string
	var vCidIsw []string
	var vCidNisw []string
	var vCidIew []string
	var vCidNiew []string
	var vCidEmpty *bool
	var vCidRegex []string
	var vCidIregex []string
	var vProviderID []int64
	var vProviderIDn []int64
	var vTypeID []int64
	var vTypeIDn []int64
	var vStatus []string
	var vStatusn []string
	var vStatusIc []string
	var vStatusNic []string
	var vStatusIe []string
	var vStatusNie []string
	var vStatusIsw []string
	var vStatusNisw []string
	var vStatusIew []string
	var vStatusNiew []string
	var vStatusEmpty *bool
	var vStatusRegex []string
	var vStatusIregex []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vCommitRate []int64
	var vCommitRaten []int64
	var vCommitRateLt []int64
	var vCommitRateLte []int64
	var vCommitRateGt []int64
	var vCommitRateGte []int64
	var vCommitRateEmpty *bool
	var vContact []int64
	var vContactn []int64
	var vContactGroup []string
	var vContactGroupn []string
	var vContactRole []int64
	var vContactRolen []int64
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
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
	var vDistance []float64
	var vDistancen []float64
	var vDistanceLt []float64
	var vDistanceLte []float64
	var vDistanceGt []float64
	var vDistanceGte []float64
	var vDistanceEmpty *bool
	var vDistanceUnit *string
	var vDistanceUnitn *string
	var vDistanceUnitIc []string
	var vDistanceUnitNic []string
	var vDistanceUnitIe []string
	var vDistanceUnitNie []string
	var vDistanceUnitIsw []string
	var vDistanceUnitNisw []string
	var vDistanceUnitIew []string
	var vDistanceUnitNiew []string
	var vDistanceUnitEmpty *bool
	var vDistanceUnitRegex []string
	var vDistanceUnitIregex []string
	var vInstallDate []string
	var vInstallDaten []string
	var vInstallDateLt []string
	var vInstallDateLte []string
	var vInstallDateGt []string
	var vInstallDateGte []string
	var vInstallDateEmpty *bool
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLocationID []int64
	var vLocationIDn []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vProvider []string
	var vProvidern []string
	var vProviderAccount []string
	var vProviderAccountn []string
	var vProviderAccountID []int64
	var vProviderAccountIDn []int64
	var vProviderNetworkID []int64
	var vProviderNetworkIDn []int64
	var vQ *string
	var vRegion []string
	var vRegionn []string
	var vRegionID []string
	var vRegionIDn []string
	var vSite []string
	var vSiten []string
	var vSiteGroup []string
	var vSiteGroupn []string
	var vSiteGroupID []string
	var vSiteGroupIDn []string
	var vSiteID []int64
	var vSiteIDn []int64
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
	var vTenantID []int64
	var vTenantIDn []int64
	var vTerminationaID []int64
	var vTerminationaIDn []int64
	var vTerminationDate []string
	var vTerminationDaten []string
	var vTerminationDateLt []string
	var vTerminationDateLte []string
	var vTerminationDateGt []string
	var vTerminationDateGte []string
	var vTerminationDateEmpty *bool
	var vTerminationzID []int64
	var vTerminationzIDn []int64
	var vType []string
	var vTypen []string
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
		case "cid":
			v := value
			vCid = append(vCid, v)
		case "cid__n":
			v := value
			vCidn = append(vCidn, v)
		case "cid__ic":
			v := value
			vCidIc = append(vCidIc, v)
		case "cid__nic":
			v := value
			vCidNic = append(vCidNic, v)
		case "cid__ie":
			v := value
			vCidIe = append(vCidIe, v)
		case "cid__nie":
			v := value
			vCidNie = append(vCidNie, v)
		case "cid__isw":
			v := value
			vCidIsw = append(vCidIsw, v)
		case "cid__nisw":
			v := value
			vCidNisw = append(vCidNisw, v)
		case "cid__iew":
			v := value
			vCidIew = append(vCidIew, v)
		case "cid__niew":
			v := value
			vCidNiew = append(vCidNiew, v)
		case "cid__empty":
			if vCidEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'cid__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cid__empty' takes a boolean, got %q.", value))
				return
			}
			vCidEmpty = &v
		case "cid__regex":
			v := value
			vCidRegex = append(vCidRegex, v)
		case "cid__iregex":
			v := value
			vCidIregex = append(vCidIregex, v)
		case "provider_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_id' takes an integer, got %q.", value))
				return
			}
			vProviderID = append(vProviderID, v)
		case "provider_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_id__n' takes an integer, got %q.", value))
				return
			}
			vProviderIDn = append(vProviderIDn, v)
		case "type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'type_id' takes an integer, got %q.", value))
				return
			}
			vTypeID = append(vTypeID, v)
		case "type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'type_id__n' takes an integer, got %q.", value))
				return
			}
			vTypeIDn = append(vTypeIDn, v)
		case "status":
			v := value
			vStatus = append(vStatus, v)
		case "status__n":
			v := value
			vStatusn = append(vStatusn, v)
		case "status__ic":
			v := value
			vStatusIc = append(vStatusIc, v)
		case "status__nic":
			v := value
			vStatusNic = append(vStatusNic, v)
		case "status__ie":
			v := value
			vStatusIe = append(vStatusIe, v)
		case "status__nie":
			v := value
			vStatusNie = append(vStatusNie, v)
		case "status__isw":
			v := value
			vStatusIsw = append(vStatusIsw, v)
		case "status__nisw":
			v := value
			vStatusNisw = append(vStatusNisw, v)
		case "status__iew":
			v := value
			vStatusIew = append(vStatusIew, v)
		case "status__niew":
			v := value
			vStatusNiew = append(vStatusNiew, v)
		case "status__empty":
			if vStatusEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'status__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'status__empty' takes a boolean, got %q.", value))
				return
			}
			vStatusEmpty = &v
		case "status__regex":
			v := value
			vStatusRegex = append(vStatusRegex, v)
		case "status__iregex":
			v := value
			vStatusIregex = append(vStatusIregex, v)
		case "owner_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id' takes an integer, got %q.", value))
				return
			}
			vOwnerID = append(vOwnerID, v)
		case "owner_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerIDn = append(vOwnerIDn, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
		case "commit_rate":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'commit_rate' takes an integer, got %q.", value))
				return
			}
			vCommitRate = append(vCommitRate, v)
		case "commit_rate__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'commit_rate__n' takes an integer, got %q.", value))
				return
			}
			vCommitRaten = append(vCommitRaten, v)
		case "commit_rate__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'commit_rate__lt' takes an integer, got %q.", value))
				return
			}
			vCommitRateLt = append(vCommitRateLt, v)
		case "commit_rate__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'commit_rate__lte' takes an integer, got %q.", value))
				return
			}
			vCommitRateLte = append(vCommitRateLte, v)
		case "commit_rate__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'commit_rate__gt' takes an integer, got %q.", value))
				return
			}
			vCommitRateGt = append(vCommitRateGt, v)
		case "commit_rate__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'commit_rate__gte' takes an integer, got %q.", value))
				return
			}
			vCommitRateGte = append(vCommitRateGte, v)
		case "commit_rate__empty":
			if vCommitRateEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'commit_rate__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'commit_rate__empty' takes a boolean, got %q.", value))
				return
			}
			vCommitRateEmpty = &v
		case "contact":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact' takes an integer, got %q.", value))
				return
			}
			vContact = append(vContact, v)
		case "contact__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact__n' takes an integer, got %q.", value))
				return
			}
			vContactn = append(vContactn, v)
		case "contact_group":
			v := value
			vContactGroup = append(vContactGroup, v)
		case "contact_group__n":
			v := value
			vContactGroupn = append(vContactGroupn, v)
		case "contact_role":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_role' takes an integer, got %q.", value))
				return
			}
			vContactRole = append(vContactRole, v)
		case "contact_role__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'contact_role__n' takes an integer, got %q.", value))
				return
			}
			vContactRolen = append(vContactRolen, v)
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
		case "distance":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance' takes a number, got %q.", value))
				return
			}
			vDistance = append(vDistance, v)
		case "distance__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__n' takes a number, got %q.", value))
				return
			}
			vDistancen = append(vDistancen, v)
		case "distance__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__lt' takes a number, got %q.", value))
				return
			}
			vDistanceLt = append(vDistanceLt, v)
		case "distance__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__lte' takes a number, got %q.", value))
				return
			}
			vDistanceLte = append(vDistanceLte, v)
		case "distance__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__gt' takes a number, got %q.", value))
				return
			}
			vDistanceGt = append(vDistanceGt, v)
		case "distance__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__gte' takes a number, got %q.", value))
				return
			}
			vDistanceGte = append(vDistanceGte, v)
		case "distance__empty":
			if vDistanceEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance__empty' takes a boolean, got %q.", value))
				return
			}
			vDistanceEmpty = &v
		case "distance_unit":
			if vDistanceUnit != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance_unit' takes a single value.")
				return
			}
			v := value
			vDistanceUnit = &v
		case "distance_unit__n":
			if vDistanceUnitn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance_unit__n' takes a single value.")
				return
			}
			v := value
			vDistanceUnitn = &v
		case "distance_unit__ic":
			v := value
			vDistanceUnitIc = append(vDistanceUnitIc, v)
		case "distance_unit__nic":
			v := value
			vDistanceUnitNic = append(vDistanceUnitNic, v)
		case "distance_unit__ie":
			v := value
			vDistanceUnitIe = append(vDistanceUnitIe, v)
		case "distance_unit__nie":
			v := value
			vDistanceUnitNie = append(vDistanceUnitNie, v)
		case "distance_unit__isw":
			v := value
			vDistanceUnitIsw = append(vDistanceUnitIsw, v)
		case "distance_unit__nisw":
			v := value
			vDistanceUnitNisw = append(vDistanceUnitNisw, v)
		case "distance_unit__iew":
			v := value
			vDistanceUnitIew = append(vDistanceUnitIew, v)
		case "distance_unit__niew":
			v := value
			vDistanceUnitNiew = append(vDistanceUnitNiew, v)
		case "distance_unit__empty":
			if vDistanceUnitEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'distance_unit__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'distance_unit__empty' takes a boolean, got %q.", value))
				return
			}
			vDistanceUnitEmpty = &v
		case "distance_unit__regex":
			v := value
			vDistanceUnitRegex = append(vDistanceUnitRegex, v)
		case "distance_unit__iregex":
			v := value
			vDistanceUnitIregex = append(vDistanceUnitIregex, v)
		case "install_date":
			v := value
			vInstallDate = append(vInstallDate, v)
		case "install_date__n":
			v := value
			vInstallDaten = append(vInstallDaten, v)
		case "install_date__lt":
			v := value
			vInstallDateLt = append(vInstallDateLt, v)
		case "install_date__lte":
			v := value
			vInstallDateLte = append(vInstallDateLte, v)
		case "install_date__gt":
			v := value
			vInstallDateGt = append(vInstallDateGt, v)
		case "install_date__gte":
			v := value
			vInstallDateGte = append(vInstallDateGte, v)
		case "install_date__empty":
			if vInstallDateEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'install_date__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'install_date__empty' takes a boolean, got %q.", value))
				return
			}
			vInstallDateEmpty = &v
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
		case "location_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id' takes an integer, got %q.", value))
				return
			}
			vLocationID = append(vLocationID, v)
		case "location_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id__n' takes an integer, got %q.", value))
				return
			}
			vLocationIDn = append(vLocationIDn, v)
		case "owner":
			v := value
			vOwner = append(vOwner, v)
		case "owner__n":
			v := value
			vOwnern = append(vOwnern, v)
		case "owner_group":
			v := value
			vOwnerGroup = append(vOwnerGroup, v)
		case "owner_group__n":
			v := value
			vOwnerGroupn = append(vOwnerGroupn, v)
		case "owner_group_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupID = append(vOwnerGroupID, v)
		case "owner_group_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'owner_group_id__n' takes an integer, got %q.", value))
				return
			}
			vOwnerGroupIDn = append(vOwnerGroupIDn, v)
		case "provider":
			v := value
			vProvider = append(vProvider, v)
		case "provider__n":
			v := value
			vProvidern = append(vProvidern, v)
		case "provider_account":
			v := value
			vProviderAccount = append(vProviderAccount, v)
		case "provider_account__n":
			v := value
			vProviderAccountn = append(vProviderAccountn, v)
		case "provider_account_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_account_id' takes an integer, got %q.", value))
				return
			}
			vProviderAccountID = append(vProviderAccountID, v)
		case "provider_account_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_account_id__n' takes an integer, got %q.", value))
				return
			}
			vProviderAccountIDn = append(vProviderAccountIDn, v)
		case "provider_network_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_network_id' takes an integer, got %q.", value))
				return
			}
			vProviderNetworkID = append(vProviderNetworkID, v)
		case "provider_network_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'provider_network_id__n' takes an integer, got %q.", value))
				return
			}
			vProviderNetworkIDn = append(vProviderNetworkIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "region":
			v := value
			vRegion = append(vRegion, v)
		case "region__n":
			v := value
			vRegionn = append(vRegionn, v)
		case "region_id":
			v := value
			vRegionID = append(vRegionID, v)
		case "region_id__n":
			v := value
			vRegionIDn = append(vRegionIDn, v)
		case "site":
			v := value
			vSite = append(vSite, v)
		case "site__n":
			v := value
			vSiten = append(vSiten, v)
		case "site_group":
			v := value
			vSiteGroup = append(vSiteGroup, v)
		case "site_group__n":
			v := value
			vSiteGroupn = append(vSiteGroupn, v)
		case "site_group_id":
			v := value
			vSiteGroupID = append(vSiteGroupID, v)
		case "site_group_id__n":
			v := value
			vSiteGroupIDn = append(vSiteGroupIDn, v)
		case "site_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id' takes an integer, got %q.", value))
				return
			}
			vSiteID = append(vSiteID, v)
		case "site_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id__n' takes an integer, got %q.", value))
				return
			}
			vSiteIDn = append(vSiteIDn, v)
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
		case "tenant":
			v := value
			vTenant = append(vTenant, v)
		case "tenant__n":
			v := value
			vTenantn = append(vTenantn, v)
		case "tenant_group":
			v := value
			vTenantGroup = append(vTenantGroup, v)
		case "tenant_group__n":
			v := value
			vTenantGroupn = append(vTenantGroupn, v)
		case "tenant_group_id":
			v := value
			vTenantGroupID = append(vTenantGroupID, v)
		case "tenant_group_id__n":
			v := value
			vTenantGroupIDn = append(vTenantGroupIDn, v)
		case "tenant_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id' takes an integer, got %q.", value))
				return
			}
			vTenantID = append(vTenantID, v)
		case "tenant_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'tenant_id__n' takes an integer, got %q.", value))
				return
			}
			vTenantIDn = append(vTenantIDn, v)
		case "termination_a_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_a_id' takes an integer, got %q.", value))
				return
			}
			vTerminationaID = append(vTerminationaID, v)
		case "termination_a_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_a_id__n' takes an integer, got %q.", value))
				return
			}
			vTerminationaIDn = append(vTerminationaIDn, v)
		case "termination_date":
			v := value
			vTerminationDate = append(vTerminationDate, v)
		case "termination_date__n":
			v := value
			vTerminationDaten = append(vTerminationDaten, v)
		case "termination_date__lt":
			v := value
			vTerminationDateLt = append(vTerminationDateLt, v)
		case "termination_date__lte":
			v := value
			vTerminationDateLte = append(vTerminationDateLte, v)
		case "termination_date__gt":
			v := value
			vTerminationDateGt = append(vTerminationDateGt, v)
		case "termination_date__gte":
			v := value
			vTerminationDateGte = append(vTerminationDateGte, v)
		case "termination_date__empty":
			if vTerminationDateEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'termination_date__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_date__empty' takes a boolean, got %q.", value))
				return
			}
			vTerminationDateEmpty = &v
		case "termination_z_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_z_id' takes an integer, got %q.", value))
				return
			}
			vTerminationzID = append(vTerminationzID, v)
		case "termination_z_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_z_id__n' takes an integer, got %q.", value))
				return
			}
			vTerminationzIDn = append(vTerminationzIDn, v)
		case "type":
			v := value
			vType = append(vType, v)
		case "type__n":
			v := value
			vTypen = append(vTypen, v)
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
	if len(vCid) > 0 {
		params.SetCid(vCid)
	}
	if len(vCidn) > 0 {
		params.SetCidn(vCidn)
	}
	if len(vCidIc) > 0 {
		params.SetCidIc(vCidIc)
	}
	if len(vCidNic) > 0 {
		params.SetCidNic(vCidNic)
	}
	if len(vCidIe) > 0 {
		params.SetCidIe(vCidIe)
	}
	if len(vCidNie) > 0 {
		params.SetCidNie(vCidNie)
	}
	if len(vCidIsw) > 0 {
		params.SetCidIsw(vCidIsw)
	}
	if len(vCidNisw) > 0 {
		params.SetCidNisw(vCidNisw)
	}
	if len(vCidIew) > 0 {
		params.SetCidIew(vCidIew)
	}
	if len(vCidNiew) > 0 {
		params.SetCidNiew(vCidNiew)
	}
	if vCidEmpty != nil {
		params.SetCidEmpty(vCidEmpty)
	}
	if len(vCidRegex) > 0 {
		params.SetCidRegex(vCidRegex)
	}
	if len(vCidIregex) > 0 {
		params.SetCidIregex(vCidIregex)
	}
	if len(vProviderID) > 0 {
		params.SetProviderID(vProviderID)
	}
	if len(vProviderIDn) > 0 {
		params.SetProviderIDn(vProviderIDn)
	}
	if len(vTypeID) > 0 {
		params.SetTypeID(vTypeID)
	}
	if len(vTypeIDn) > 0 {
		params.SetTypeIDn(vTypeIDn)
	}
	if len(vStatus) > 0 {
		params.SetStatus(vStatus)
	}
	if len(vStatusn) > 0 {
		params.SetStatusn(vStatusn)
	}
	if len(vStatusIc) > 0 {
		params.SetStatusIc(vStatusIc)
	}
	if len(vStatusNic) > 0 {
		params.SetStatusNic(vStatusNic)
	}
	if len(vStatusIe) > 0 {
		params.SetStatusIe(vStatusIe)
	}
	if len(vStatusNie) > 0 {
		params.SetStatusNie(vStatusNie)
	}
	if len(vStatusIsw) > 0 {
		params.SetStatusIsw(vStatusIsw)
	}
	if len(vStatusNisw) > 0 {
		params.SetStatusNisw(vStatusNisw)
	}
	if len(vStatusIew) > 0 {
		params.SetStatusIew(vStatusIew)
	}
	if len(vStatusNiew) > 0 {
		params.SetStatusNiew(vStatusNiew)
	}
	if vStatusEmpty != nil {
		params.SetStatusEmpty(vStatusEmpty)
	}
	if len(vStatusRegex) > 0 {
		params.SetStatusRegex(vStatusRegex)
	}
	if len(vStatusIregex) > 0 {
		params.SetStatusIregex(vStatusIregex)
	}
	if len(vOwnerID) > 0 {
		params.SetOwnerID(vOwnerID)
	}
	if len(vOwnerIDn) > 0 {
		params.SetOwnerIDn(vOwnerIDn)
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
	if len(vCommitRate) > 0 {
		params.SetCommitRate(vCommitRate)
	}
	if len(vCommitRaten) > 0 {
		params.SetCommitRaten(vCommitRaten)
	}
	if len(vCommitRateLt) > 0 {
		params.SetCommitRateLt(vCommitRateLt)
	}
	if len(vCommitRateLte) > 0 {
		params.SetCommitRateLte(vCommitRateLte)
	}
	if len(vCommitRateGt) > 0 {
		params.SetCommitRateGt(vCommitRateGt)
	}
	if len(vCommitRateGte) > 0 {
		params.SetCommitRateGte(vCommitRateGte)
	}
	if vCommitRateEmpty != nil {
		params.SetCommitRateEmpty(vCommitRateEmpty)
	}
	if len(vContact) > 0 {
		params.SetContact(vContact)
	}
	if len(vContactn) > 0 {
		params.SetContactn(vContactn)
	}
	if len(vContactGroup) > 0 {
		params.SetContactGroup(vContactGroup)
	}
	if len(vContactGroupn) > 0 {
		params.SetContactGroupn(vContactGroupn)
	}
	if len(vContactRole) > 0 {
		params.SetContactRole(vContactRole)
	}
	if len(vContactRolen) > 0 {
		params.SetContactRolen(vContactRolen)
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
	if len(vDistance) > 0 {
		params.SetDistance(vDistance)
	}
	if len(vDistancen) > 0 {
		params.SetDistancen(vDistancen)
	}
	if len(vDistanceLt) > 0 {
		params.SetDistanceLt(vDistanceLt)
	}
	if len(vDistanceLte) > 0 {
		params.SetDistanceLte(vDistanceLte)
	}
	if len(vDistanceGt) > 0 {
		params.SetDistanceGt(vDistanceGt)
	}
	if len(vDistanceGte) > 0 {
		params.SetDistanceGte(vDistanceGte)
	}
	if vDistanceEmpty != nil {
		params.SetDistanceEmpty(vDistanceEmpty)
	}
	if vDistanceUnit != nil {
		params.SetDistanceUnit(vDistanceUnit)
	}
	if vDistanceUnitn != nil {
		params.SetDistanceUnitn(vDistanceUnitn)
	}
	if len(vDistanceUnitIc) > 0 {
		params.SetDistanceUnitIc(vDistanceUnitIc)
	}
	if len(vDistanceUnitNic) > 0 {
		params.SetDistanceUnitNic(vDistanceUnitNic)
	}
	if len(vDistanceUnitIe) > 0 {
		params.SetDistanceUnitIe(vDistanceUnitIe)
	}
	if len(vDistanceUnitNie) > 0 {
		params.SetDistanceUnitNie(vDistanceUnitNie)
	}
	if len(vDistanceUnitIsw) > 0 {
		params.SetDistanceUnitIsw(vDistanceUnitIsw)
	}
	if len(vDistanceUnitNisw) > 0 {
		params.SetDistanceUnitNisw(vDistanceUnitNisw)
	}
	if len(vDistanceUnitIew) > 0 {
		params.SetDistanceUnitIew(vDistanceUnitIew)
	}
	if len(vDistanceUnitNiew) > 0 {
		params.SetDistanceUnitNiew(vDistanceUnitNiew)
	}
	if vDistanceUnitEmpty != nil {
		params.SetDistanceUnitEmpty(vDistanceUnitEmpty)
	}
	if len(vDistanceUnitRegex) > 0 {
		params.SetDistanceUnitRegex(vDistanceUnitRegex)
	}
	if len(vDistanceUnitIregex) > 0 {
		params.SetDistanceUnitIregex(vDistanceUnitIregex)
	}
	if len(vInstallDate) > 0 {
		params.SetInstallDate(vInstallDate)
	}
	if len(vInstallDaten) > 0 {
		params.SetInstallDaten(vInstallDaten)
	}
	if len(vInstallDateLt) > 0 {
		params.SetInstallDateLt(vInstallDateLt)
	}
	if len(vInstallDateLte) > 0 {
		params.SetInstallDateLte(vInstallDateLte)
	}
	if len(vInstallDateGt) > 0 {
		params.SetInstallDateGt(vInstallDateGt)
	}
	if len(vInstallDateGte) > 0 {
		params.SetInstallDateGte(vInstallDateGte)
	}
	if vInstallDateEmpty != nil {
		params.SetInstallDateEmpty(vInstallDateEmpty)
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
	if len(vLocationID) > 0 {
		params.SetLocationID(vLocationID)
	}
	if len(vLocationIDn) > 0 {
		params.SetLocationIDn(vLocationIDn)
	}
	if len(vOwner) > 0 {
		params.SetOwner(vOwner)
	}
	if len(vOwnern) > 0 {
		params.SetOwnern(vOwnern)
	}
	if len(vOwnerGroup) > 0 {
		params.SetOwnerGroup(vOwnerGroup)
	}
	if len(vOwnerGroupn) > 0 {
		params.SetOwnerGroupn(vOwnerGroupn)
	}
	if len(vOwnerGroupID) > 0 {
		params.SetOwnerGroupID(vOwnerGroupID)
	}
	if len(vOwnerGroupIDn) > 0 {
		params.SetOwnerGroupIDn(vOwnerGroupIDn)
	}
	if len(vProvider) > 0 {
		params.SetProvider(vProvider)
	}
	if len(vProvidern) > 0 {
		params.SetProvidern(vProvidern)
	}
	if len(vProviderAccount) > 0 {
		params.SetProviderAccount(vProviderAccount)
	}
	if len(vProviderAccountn) > 0 {
		params.SetProviderAccountn(vProviderAccountn)
	}
	if len(vProviderAccountID) > 0 {
		params.SetProviderAccountID(vProviderAccountID)
	}
	if len(vProviderAccountIDn) > 0 {
		params.SetProviderAccountIDn(vProviderAccountIDn)
	}
	if len(vProviderNetworkID) > 0 {
		params.SetProviderNetworkID(vProviderNetworkID)
	}
	if len(vProviderNetworkIDn) > 0 {
		params.SetProviderNetworkIDn(vProviderNetworkIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRegion) > 0 {
		params.SetRegion(vRegion)
	}
	if len(vRegionn) > 0 {
		params.SetRegionn(vRegionn)
	}
	if len(vRegionID) > 0 {
		params.SetRegionID(vRegionID)
	}
	if len(vRegionIDn) > 0 {
		params.SetRegionIDn(vRegionIDn)
	}
	if len(vSite) > 0 {
		params.SetSite(vSite)
	}
	if len(vSiten) > 0 {
		params.SetSiten(vSiten)
	}
	if len(vSiteGroup) > 0 {
		params.SetSiteGroup(vSiteGroup)
	}
	if len(vSiteGroupn) > 0 {
		params.SetSiteGroupn(vSiteGroupn)
	}
	if len(vSiteGroupID) > 0 {
		params.SetSiteGroupID(vSiteGroupID)
	}
	if len(vSiteGroupIDn) > 0 {
		params.SetSiteGroupIDn(vSiteGroupIDn)
	}
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSiteIDn) > 0 {
		params.SetSiteIDn(vSiteIDn)
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
	if len(vTenant) > 0 {
		params.SetTenant(vTenant)
	}
	if len(vTenantn) > 0 {
		params.SetTenantn(vTenantn)
	}
	if len(vTenantGroup) > 0 {
		params.SetTenantGroup(vTenantGroup)
	}
	if len(vTenantGroupn) > 0 {
		params.SetTenantGroupn(vTenantGroupn)
	}
	if len(vTenantGroupID) > 0 {
		params.SetTenantGroupID(vTenantGroupID)
	}
	if len(vTenantGroupIDn) > 0 {
		params.SetTenantGroupIDn(vTenantGroupIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vTerminationaID) > 0 {
		params.SetTerminationaID(vTerminationaID)
	}
	if len(vTerminationaIDn) > 0 {
		params.SetTerminationaIDn(vTerminationaIDn)
	}
	if len(vTerminationDate) > 0 {
		params.SetTerminationDate(vTerminationDate)
	}
	if len(vTerminationDaten) > 0 {
		params.SetTerminationDaten(vTerminationDaten)
	}
	if len(vTerminationDateLt) > 0 {
		params.SetTerminationDateLt(vTerminationDateLt)
	}
	if len(vTerminationDateLte) > 0 {
		params.SetTerminationDateLte(vTerminationDateLte)
	}
	if len(vTerminationDateGt) > 0 {
		params.SetTerminationDateGt(vTerminationDateGt)
	}
	if len(vTerminationDateGte) > 0 {
		params.SetTerminationDateGte(vTerminationDateGte)
	}
	if vTerminationDateEmpty != nil {
		params.SetTerminationDateEmpty(vTerminationDateEmpty)
	}
	if len(vTerminationzID) > 0 {
		params.SetTerminationzID(vTerminationzID)
	}
	if len(vTerminationzIDn) > 0 {
		params.SetTerminationzIDn(vTerminationzIDn)
	}
	if len(vType) > 0 {
		params.SetType(vType)
	}
	if len(vTypen) > 0 {
		params.SetTypen(vTypen)
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
