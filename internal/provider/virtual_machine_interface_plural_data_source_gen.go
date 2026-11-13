// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualMachineInterfacesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualMachineInterfacesDataSource)(nil)
)

// NewVirtualMachineInterfacesDataSource returns a new virtual_machine_interfaces data source, which
// lists virtual_machine_interface objects matching its filters.
func NewVirtualMachineInterfacesDataSource() datasource.DataSource {
	return &virtualMachineInterfacesDataSource{}
}

type virtualMachineInterfacesDataSource struct {
	client *netboxapi.Client
}

// virtualMachineInterfacesDataSourceModel is the data source model: the filters, the limit and the matching
// virtual_machine_interfaces.
type virtualMachineInterfacesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"virtual_machine_interfaces"`
}

// virtualMachineInterfaceResourceAttrTypes is the attribute type map of virtualMachineInterfaceResourceModel.
func virtualMachineInterfaceResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                         types.Int64Type,
		"virtual_machine_id":         types.Int64Type,
		"name":                       types.StringType,
		"enabled":                    types.BoolType,
		"mtu":                        types.Int64Type,
		"mode":                       types.StringType,
		"untagged_vlan_id":           types.Int64Type,
		"tagged_vlan_ids":            types.SetType{ElemType: types.Int64Type},
		"qinq_svlan_id":              types.Int64Type,
		"vlan_translation_policy_id": types.Int64Type,
		"vrf_id":                     types.Int64Type,
		"parent_id":                  types.Int64Type,
		"bridge_id":                  types.Int64Type,
		"description":                types.StringType,
		"owner_id":                   types.Int64Type,
		"created":                    types.StringType,
		"last_updated":               types.StringType,
		"url":                        types.StringType,
		"tags":                       types.SetType{ElemType: types.StringType},
		"tags_all":                   types.SetType{ElemType: types.StringType},
		"custom_fields":              types.MapType{ElemType: types.StringType},
	}
}

func (d *virtualMachineInterfacesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_machine_interfaces"
}

func (d *virtualMachineInterfacesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Virtualization:Lists virtual_machine_interface objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: bridge_id, bridge_id__n, cluster, cluster__n, cluster_id, cluster_id__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, enabled, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, l2vpn, l2vpn__n, l2vpn_id, l2vpn_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mac_address, mac_address__ic, mac_address__ie, mac_address__iew, mac_address__iregex, mac_address__isw, mac_address__n, mac_address__nic, mac_address__nie, mac_address__niew, mac_address__nisw, mac_address__regex, mode, mode__empty, mode__ic, mode__ie, mode__iew, mode__iregex, mode__isw, mode__n, mode__nic, mode__nie, mode__niew, mode__nisw, mode__regex, mtu, mtu__empty, mtu__gt, mtu__gte, mtu__lt, mtu__lte, mtu__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, parent_id, parent_id__n, primary_mac_address, primary_mac_address__n, primary_mac_address_id, primary_mac_address_id__n, q, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, virtual_machine, virtual_machine__n, virtual_machine_id, virtual_machine_id__n, vlan, vlan_id, vlan_translation_policy, vlan_translation_policy__n, vlan_translation_policy_id, vlan_translation_policy_id__n, vrf, vrf__n, vrf_id, vrf_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"virtual_machine_interfaces": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching virtual_machine_interfaces.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"virtual_machine_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual machine.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"enabled": schema.BoolAttribute{
							Computed:    true,
							Description: "Whether the interface is enabled.",
						},
						"mtu": schema.Int64Attribute{
							Computed:    true,
							Description: "MTU in bytes. Must be a whole number.",
						},
						"mode": schema.StringAttribute{
							Computed:    true,
							Description: "802.1Q tagging mode. One of: access, tagged, tagged-all, q-in-q.",
						},
						"untagged_vlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the untagged VLAN.",
						},
						"tagged_vlan_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the tagged VLANs.",
						},
						"qinq_svlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the Q-in-Q service VLAN.",
						},
						"vlan_translation_policy_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VLAN translation policy.",
						},
						"vrf_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VRF.",
						},
						"parent_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the parent interface.",
						},
						"bridge_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the bridged interface.",
						},
						"description": schema.StringAttribute{
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

func (d *virtualMachineInterfacesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualMachineInterfacesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualMachineInterfacesDataSourceModel

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
	var responseDTOs []*netboxapi.VirtualMachineInterfaceResponseDTO

	params := virtualization.NewVirtualizationInterfacesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                            netboxapi.FilterInt64s(params.SetID),
		"id__n":                         netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                        netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                       netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                        netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                       netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                     netboxapi.FilterBool(params.SetIDEmpty),
		"name":                          netboxapi.FilterStrings(params.SetName),
		"name__n":                       netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                      netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":                     netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                      netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":                     netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":                     netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":                    netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":                     netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":                    netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":                   netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":                   netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":                  netboxapi.FilterStrings(params.SetNameIregex),
		"virtual_machine_id":            netboxapi.FilterInt64s(params.SetVirtualMachineID),
		"virtual_machine_id__n":         netboxapi.FilterInt64s(params.SetVirtualMachineIDn),
		"cluster_id":                    netboxapi.FilterInt64s(params.SetClusterID),
		"cluster_id__n":                 netboxapi.FilterInt64s(params.SetClusterIDn),
		"mac_address":                   netboxapi.FilterStrings(params.SetMacAddress),
		"mac_address__n":                netboxapi.FilterStrings(params.SetMacAddressn),
		"mac_address__ic":               netboxapi.FilterStrings(params.SetMacAddressIc),
		"mac_address__nic":              netboxapi.FilterStrings(params.SetMacAddressNic),
		"mac_address__ie":               netboxapi.FilterStrings(params.SetMacAddressIe),
		"mac_address__nie":              netboxapi.FilterStrings(params.SetMacAddressNie),
		"mac_address__isw":              netboxapi.FilterStrings(params.SetMacAddressIsw),
		"mac_address__nisw":             netboxapi.FilterStrings(params.SetMacAddressNisw),
		"mac_address__iew":              netboxapi.FilterStrings(params.SetMacAddressIew),
		"mac_address__niew":             netboxapi.FilterStrings(params.SetMacAddressNiew),
		"mac_address__regex":            netboxapi.FilterStrings(params.SetMacAddressRegex),
		"mac_address__iregex":           netboxapi.FilterStrings(params.SetMacAddressIregex),
		"enabled":                       netboxapi.FilterBool(params.SetEnabled),
		"mode":                          netboxapi.FilterStrings(params.SetMode),
		"mode__n":                       netboxapi.FilterStrings(params.SetModen),
		"mode__ic":                      netboxapi.FilterStrings(params.SetModeIc),
		"mode__nic":                     netboxapi.FilterStrings(params.SetModeNic),
		"mode__ie":                      netboxapi.FilterStrings(params.SetModeIe),
		"mode__nie":                     netboxapi.FilterStrings(params.SetModeNie),
		"mode__isw":                     netboxapi.FilterStrings(params.SetModeIsw),
		"mode__nisw":                    netboxapi.FilterStrings(params.SetModeNisw),
		"mode__iew":                     netboxapi.FilterStrings(params.SetModeIew),
		"mode__niew":                    netboxapi.FilterStrings(params.SetModeNiew),
		"mode__empty":                   netboxapi.FilterBool(params.SetModeEmpty),
		"mode__regex":                   netboxapi.FilterStrings(params.SetModeRegex),
		"mode__iregex":                  netboxapi.FilterStrings(params.SetModeIregex),
		"mtu":                           netboxapi.FilterInt64s(params.SetMtu),
		"mtu__n":                        netboxapi.FilterInt64s(params.SetMtun),
		"mtu__lt":                       netboxapi.FilterInt64s(params.SetMtuLt),
		"mtu__lte":                      netboxapi.FilterInt64s(params.SetMtuLte),
		"mtu__gt":                       netboxapi.FilterInt64s(params.SetMtuGt),
		"mtu__gte":                      netboxapi.FilterInt64s(params.SetMtuGte),
		"mtu__empty":                    netboxapi.FilterBool(params.SetMtuEmpty),
		"vrf_id":                        netboxapi.FilterInt64s(params.SetVrfID),
		"vrf_id__n":                     netboxapi.FilterInt64s(params.SetVrfIDn),
		"parent_id":                     netboxapi.FilterInt64s(params.SetParentID),
		"parent_id__n":                  netboxapi.FilterInt64s(params.SetParentIDn),
		"bridge_id":                     netboxapi.FilterInt64s(params.SetBridgeID),
		"bridge_id__n":                  netboxapi.FilterInt64s(params.SetBridgeIDn),
		"description":                   netboxapi.FilterStrings(params.SetDescription),
		"description__n":                netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":               netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":              netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":               netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":              netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":              netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":             netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":              netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":             netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":            netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":            netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":           netboxapi.FilterStrings(params.SetDescriptionIregex),
		"virtual_machine":               netboxapi.FilterStrings(params.SetVirtualMachine),
		"virtual_machine__n":            netboxapi.FilterStrings(params.SetVirtualMachinen),
		"cluster":                       netboxapi.FilterStrings(params.SetCluster),
		"cluster__n":                    netboxapi.FilterStrings(params.SetClustern),
		"vrf":                           netboxapi.FilterStrings(params.SetVrf),
		"vrf__n":                        netboxapi.FilterStrings(params.SetVrfn),
		"owner_id":                      netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":                   netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                           netboxapi.FilterStrings(params.SetTag),
		"tag__n":                        netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                      netboxapi.FilterStrings(params.SetTagAny),
		"created":                       netboxapi.FilterStrings(params.SetCreated),
		"created__n":                    netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                   netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":                  netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                   netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":                  netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":                netboxapi.FilterStrings(params.SetCreatedEmpty),
		"l2vpn":                         netboxapi.FilterInt64s(params.SetL2vpn),
		"l2vpn__n":                      netboxapi.FilterInt64s(params.SetL2vpnn),
		"l2vpn_id":                      netboxapi.FilterInt64s(params.SetL2vpnID),
		"l2vpn_id__n":                   netboxapi.FilterInt64s(params.SetL2vpnIDn),
		"last_updated":                  netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":               netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":              netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":             netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":              netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":             netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":           netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                         netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                      netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":                   netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":                netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":                netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":             netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"primary_mac_address":           netboxapi.FilterStrings(params.SetPrimaryMacAddress),
		"primary_mac_address__n":        netboxapi.FilterStrings(params.SetPrimaryMacAddressn),
		"primary_mac_address_id":        netboxapi.FilterInt64s(params.SetPrimaryMacAddressID),
		"primary_mac_address_id__n":     netboxapi.FilterInt64s(params.SetPrimaryMacAddressIDn),
		"q":                             netboxapi.FilterString(params.SetQ),
		"tag_id":                        netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                     netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                   netboxapi.FilterInt64s(params.SetTagIDAny),
		"vlan":                          netboxapi.FilterString(params.SetVlan),
		"vlan_id":                       netboxapi.FilterString(params.SetVlanID),
		"vlan_translation_policy":       netboxapi.FilterStrings(params.SetVlanTranslationPolicy),
		"vlan_translation_policy__n":    netboxapi.FilterStrings(params.SetVlanTranslationPolicyn),
		"vlan_translation_policy_id":    netboxapi.FilterInt64s(params.SetVlanTranslationPolicyID),
		"vlan_translation_policy_id__n": netboxapi.FilterInt64s(params.SetVlanTranslationPolicyIDn),
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
		res, err := d.client.Virtualization.VirtualizationInterfacesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_virtual_machine_interfaces", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.VirtualMachineInterfaceResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	items := make([]virtualMachineInterfaceResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model virtualMachineInterfaceResourceModel
		resp.Diagnostics.Append(flattenVirtualMachineInterface(ctx, responseDTO, &model)...)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: virtualMachineInterfaceResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
