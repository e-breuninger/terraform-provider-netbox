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
	_ datasource.DataSource              = (*virtualCircuitsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualCircuitsDataSource)(nil)
)

// NewVirtualCircuitsDataSource returns a new virtual_circuits data source, which
// lists virtual_circuit objects matching its filters.
func NewVirtualCircuitsDataSource() datasource.DataSource {
	return &virtualCircuitsDataSource{}
}

type virtualCircuitsDataSource struct {
	client *netboxapi.Client
}

// virtualCircuitsDataSourceModel is the data source model: the filters, the limit and the matching
// virtual_circuits.
type virtualCircuitsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"virtual_circuits"`
}

// virtualCircuitResourceAttrTypes is the attribute type map of virtualCircuitResourceModel.
func virtualCircuitResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                          types.Int64Type,
		"cid":                         types.StringType,
		"circuit_provider_network_id": types.Int64Type,
		"virtual_circuit_type_id":     types.Int64Type,
		"circuit_provider_account_id": types.Int64Type,
		"status":                      types.StringType,
		"tenant_id":                   types.Int64Type,
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

func (d *virtualCircuitsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_circuits"
}

func (d *virtualCircuitsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Circuits:Lists virtual_circuit objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: cid, cid__empty, cid__ic, cid__ie, cid__iew, cid__iregex, cid__isw, cid__n, cid__nic, cid__nie, cid__niew, cid__nisw, cid__regex, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, owner_id, owner_id__n, provider_account_id, provider_account_id__n, provider_network_id, provider_network_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"virtual_circuits": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching virtual_circuits.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"cid": schema.StringAttribute{
							Computed:    true,
							Description: "Circuit ID: unique identifier of the virtual circuit with its provider.",
						},
						"circuit_provider_network_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the provider network the virtual circuit is delivered over.",
						},
						"virtual_circuit_type_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual circuit type.",
						},
						"circuit_provider_account_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the provider account the virtual circuit is billed to.",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: planned, provisioning, active, offline, deprovisioning, decommissioned.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant the virtual circuit belongs to.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *virtualCircuitsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualCircuitsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualCircuitsDataSourceModel

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
	var responseDTOs []*netboxapi.VirtualCircuitResponseDTO

	params := circuits.NewCircuitsVirtualCircuitsListParams()
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
	var vProviderNetworkID []int64
	var vProviderNetworkIDn []int64
	var vProviderAccountID []int64
	var vProviderAccountIDn []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
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
	if len(vProviderNetworkID) > 0 {
		params.SetProviderNetworkID(vProviderNetworkID)
	}
	if len(vProviderNetworkIDn) > 0 {
		params.SetProviderNetworkIDn(vProviderNetworkIDn)
	}
	if len(vProviderAccountID) > 0 {
		params.SetProviderAccountID(vProviderAccountID)
	}
	if len(vProviderAccountIDn) > 0 {
		params.SetProviderAccountIDn(vProviderAccountIDn)
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
	pageSize := limit
	if fetchAll {
		pageSize = 1000
	}
	params.SetLimit(&pageSize)
	for offset := int64(0); ; {
		res, err := d.client.Circuits.CircuitsVirtualCircuitsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_virtual_circuits", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VirtualCircuitResponseDTOFromGoNetbox(goNetboxModel))
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
	items := make([]virtualCircuitResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model virtualCircuitResourceModel
		resp.Diagnostics.Append(flattenVirtualCircuit(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: virtualCircuitResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
