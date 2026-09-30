// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"strconv"

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
	_ datasource.DataSource              = (*vlanTranslationRulesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*vlanTranslationRulesDataSource)(nil)
)

// NewVlanTranslationRulesDataSource returns a new vlan_translation_rules data source, which
// lists vlan_translation_rule objects matching its filters.
func NewVlanTranslationRulesDataSource() datasource.DataSource {
	return &vlanTranslationRulesDataSource{}
}

type vlanTranslationRulesDataSource struct {
	client *netboxapi.Client
}

// vlanTranslationRulesDataSourceModel is the data source model: the filters, the limit and the matching
// vlan_translation_rules.
type vlanTranslationRulesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"vlan_translation_rules"`
}

// vlanTranslationRuleResourceAttrTypes is the attribute type map of vlanTranslationRuleResourceModel.
func vlanTranslationRuleResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                         types.Int64Type,
		"vlan_translation_policy_id": types.Int64Type,
		"local_vid":                  types.Int64Type,
		"remote_vid":                 types.Int64Type,
		"description":                types.StringType,
		"url":                        types.StringType,
	}
}

func (d *vlanTranslationRulesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_vlan_translation_rules"
}

func (d *vlanTranslationRulesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists vlan_translation_rule objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, local_vid, local_vid__empty, local_vid__gt, local_vid__gte, local_vid__lt, local_vid__lte, local_vid__n, policy_id, policy_id__n, remote_vid, remote_vid__empty, remote_vid__gt, remote_vid__gte, remote_vid__lt, remote_vid__lte, remote_vid__n. Repeating a name sends that parameter once per value.",
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
			"vlan_translation_rules": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching vlan_translation_rules.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"vlan_translation_policy_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the translation policy the rule belongs to.",
						},
						"local_vid": schema.Int64Attribute{
							Computed:    true,
							Description: "VLAN id on the local side (1-4094).",
						},
						"remote_vid": schema.Int64Attribute{
							Computed:    true,
							Description: "VLAN id on the remote side (1-4094).",
						},
						"description": schema.StringAttribute{
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

func (d *vlanTranslationRulesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *vlanTranslationRulesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data vlanTranslationRulesDataSourceModel

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
	var responseDTOs []*netboxapi.VlanTranslationRuleResponseDTO

	params := ipam.NewIpamVlanTranslationRulesListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vPolicyID []int64
	var vPolicyIDn []int64
	var vLocalVid []int64
	var vLocalVidn []int64
	var vLocalVidLt []int64
	var vLocalVidLte []int64
	var vLocalVidGt []int64
	var vLocalVidGte []int64
	var vLocalVidEmpty *bool
	var vRemoteVid []int64
	var vRemoteVidn []int64
	var vRemoteVidLt []int64
	var vRemoteVidLte []int64
	var vRemoteVidGt []int64
	var vRemoteVidGte []int64
	var vRemoteVidEmpty *bool
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
		case "policy_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'policy_id' takes an integer, got %q.", value))
				return
			}
			vPolicyID = append(vPolicyID, v)
		case "policy_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'policy_id__n' takes an integer, got %q.", value))
				return
			}
			vPolicyIDn = append(vPolicyIDn, v)
		case "local_vid":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_vid' takes an integer, got %q.", value))
				return
			}
			vLocalVid = append(vLocalVid, v)
		case "local_vid__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_vid__n' takes an integer, got %q.", value))
				return
			}
			vLocalVidn = append(vLocalVidn, v)
		case "local_vid__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_vid__lt' takes an integer, got %q.", value))
				return
			}
			vLocalVidLt = append(vLocalVidLt, v)
		case "local_vid__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_vid__lte' takes an integer, got %q.", value))
				return
			}
			vLocalVidLte = append(vLocalVidLte, v)
		case "local_vid__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_vid__gt' takes an integer, got %q.", value))
				return
			}
			vLocalVidGt = append(vLocalVidGt, v)
		case "local_vid__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_vid__gte' takes an integer, got %q.", value))
				return
			}
			vLocalVidGte = append(vLocalVidGte, v)
		case "local_vid__empty":
			if vLocalVidEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'local_vid__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'local_vid__empty' takes a boolean, got %q.", value))
				return
			}
			vLocalVidEmpty = &v
		case "remote_vid":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'remote_vid' takes an integer, got %q.", value))
				return
			}
			vRemoteVid = append(vRemoteVid, v)
		case "remote_vid__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'remote_vid__n' takes an integer, got %q.", value))
				return
			}
			vRemoteVidn = append(vRemoteVidn, v)
		case "remote_vid__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'remote_vid__lt' takes an integer, got %q.", value))
				return
			}
			vRemoteVidLt = append(vRemoteVidLt, v)
		case "remote_vid__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'remote_vid__lte' takes an integer, got %q.", value))
				return
			}
			vRemoteVidLte = append(vRemoteVidLte, v)
		case "remote_vid__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'remote_vid__gt' takes an integer, got %q.", value))
				return
			}
			vRemoteVidGt = append(vRemoteVidGt, v)
		case "remote_vid__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'remote_vid__gte' takes an integer, got %q.", value))
				return
			}
			vRemoteVidGte = append(vRemoteVidGte, v)
		case "remote_vid__empty":
			if vRemoteVidEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'remote_vid__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'remote_vid__empty' takes a boolean, got %q.", value))
				return
			}
			vRemoteVidEmpty = &v
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
	if len(vPolicyID) > 0 {
		params.SetPolicyID(vPolicyID)
	}
	if len(vPolicyIDn) > 0 {
		params.SetPolicyIDn(vPolicyIDn)
	}
	if len(vLocalVid) > 0 {
		params.SetLocalVid(vLocalVid)
	}
	if len(vLocalVidn) > 0 {
		params.SetLocalVidn(vLocalVidn)
	}
	if len(vLocalVidLt) > 0 {
		params.SetLocalVidLt(vLocalVidLt)
	}
	if len(vLocalVidLte) > 0 {
		params.SetLocalVidLte(vLocalVidLte)
	}
	if len(vLocalVidGt) > 0 {
		params.SetLocalVidGt(vLocalVidGt)
	}
	if len(vLocalVidGte) > 0 {
		params.SetLocalVidGte(vLocalVidGte)
	}
	if vLocalVidEmpty != nil {
		params.SetLocalVidEmpty(vLocalVidEmpty)
	}
	if len(vRemoteVid) > 0 {
		params.SetRemoteVid(vRemoteVid)
	}
	if len(vRemoteVidn) > 0 {
		params.SetRemoteVidn(vRemoteVidn)
	}
	if len(vRemoteVidLt) > 0 {
		params.SetRemoteVidLt(vRemoteVidLt)
	}
	if len(vRemoteVidLte) > 0 {
		params.SetRemoteVidLte(vRemoteVidLte)
	}
	if len(vRemoteVidGt) > 0 {
		params.SetRemoteVidGt(vRemoteVidGt)
	}
	if len(vRemoteVidGte) > 0 {
		params.SetRemoteVidGte(vRemoteVidGte)
	}
	if vRemoteVidEmpty != nil {
		params.SetRemoteVidEmpty(vRemoteVidEmpty)
	}
	pageSize := limit
	if fetchAll {
		pageSize = 1000
	}
	params.SetLimit(&pageSize)
	for offset := int64(0); ; {
		res, err := d.client.Ipam.IpamVlanTranslationRulesListContext(ctx, params, nil)
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_vlan_translation_rules", err.Error())
			return
		}
		if res.Payload == nil {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VlanTranslationRuleResponseDTOFromGoNetbox(goNetboxModel))
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
	items := make([]vlanTranslationRuleResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model vlanTranslationRuleResourceModel
		resp.Diagnostics.Append(flattenVlanTranslationRule(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: vlanTranslationRuleResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
