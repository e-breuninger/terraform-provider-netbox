// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vVirtualMachineID []int64
	var vVirtualMachineIDn []int64
	var vClusterID []int64
	var vClusterIDn []int64
	var vMacAddress []string
	var vMacAddressn []string
	var vMacAddressIc []string
	var vMacAddressNic []string
	var vMacAddressIe []string
	var vMacAddressNie []string
	var vMacAddressIsw []string
	var vMacAddressNisw []string
	var vMacAddressIew []string
	var vMacAddressNiew []string
	var vMacAddressRegex []string
	var vMacAddressIregex []string
	var vEnabled *bool
	var vMode []string
	var vModen []string
	var vModeIc []string
	var vModeNic []string
	var vModeIe []string
	var vModeNie []string
	var vModeIsw []string
	var vModeNisw []string
	var vModeIew []string
	var vModeNiew []string
	var vModeEmpty *bool
	var vModeRegex []string
	var vModeIregex []string
	var vMtu []int64
	var vMtun []int64
	var vMtuLt []int64
	var vMtuLte []int64
	var vMtuGt []int64
	var vMtuGte []int64
	var vMtuEmpty *bool
	var vVrfID []int64
	var vVrfIDn []int64
	var vParentID []int64
	var vParentIDn []int64
	var vBridgeID []int64
	var vBridgeIDn []int64
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
	var vVirtualMachine []string
	var vVirtualMachinen []string
	var vCluster []string
	var vClustern []string
	var vVrf []string
	var vVrfn []string
	var vOwnerID []int64
	var vOwnerIDn []int64
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
	var vL2vpn []int64
	var vL2vpnn []int64
	var vL2vpnID []int64
	var vL2vpnIDn []int64
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPrimaryMacAddress []string
	var vPrimaryMacAddressn []string
	var vPrimaryMacAddressID []int64
	var vPrimaryMacAddressIDn []int64
	var vQ *string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vVlan *string
	var vVlanID *string
	var vVlanTranslationPolicy []string
	var vVlanTranslationPolicyn []string
	var vVlanTranslationPolicyID []int64
	var vVlanTranslationPolicyIDn []int64
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
		case "virtual_machine_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_machine_id' takes an integer, got %q.", value))
				return
			}
			vVirtualMachineID = append(vVirtualMachineID, v)
		case "virtual_machine_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_machine_id__n' takes an integer, got %q.", value))
				return
			}
			vVirtualMachineIDn = append(vVirtualMachineIDn, v)
		case "cluster_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_id' takes an integer, got %q.", value))
				return
			}
			vClusterID = append(vClusterID, v)
		case "cluster_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'cluster_id__n' takes an integer, got %q.", value))
				return
			}
			vClusterIDn = append(vClusterIDn, v)
		case "mac_address":
			v := value
			vMacAddress = append(vMacAddress, v)
		case "mac_address__n":
			v := value
			vMacAddressn = append(vMacAddressn, v)
		case "mac_address__ic":
			v := value
			vMacAddressIc = append(vMacAddressIc, v)
		case "mac_address__nic":
			v := value
			vMacAddressNic = append(vMacAddressNic, v)
		case "mac_address__ie":
			v := value
			vMacAddressIe = append(vMacAddressIe, v)
		case "mac_address__nie":
			v := value
			vMacAddressNie = append(vMacAddressNie, v)
		case "mac_address__isw":
			v := value
			vMacAddressIsw = append(vMacAddressIsw, v)
		case "mac_address__nisw":
			v := value
			vMacAddressNisw = append(vMacAddressNisw, v)
		case "mac_address__iew":
			v := value
			vMacAddressIew = append(vMacAddressIew, v)
		case "mac_address__niew":
			v := value
			vMacAddressNiew = append(vMacAddressNiew, v)
		case "mac_address__regex":
			v := value
			vMacAddressRegex = append(vMacAddressRegex, v)
		case "mac_address__iregex":
			v := value
			vMacAddressIregex = append(vMacAddressIregex, v)
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
		case "mode":
			v := value
			vMode = append(vMode, v)
		case "mode__n":
			v := value
			vModen = append(vModen, v)
		case "mode__ic":
			v := value
			vModeIc = append(vModeIc, v)
		case "mode__nic":
			v := value
			vModeNic = append(vModeNic, v)
		case "mode__ie":
			v := value
			vModeIe = append(vModeIe, v)
		case "mode__nie":
			v := value
			vModeNie = append(vModeNie, v)
		case "mode__isw":
			v := value
			vModeIsw = append(vModeIsw, v)
		case "mode__nisw":
			v := value
			vModeNisw = append(vModeNisw, v)
		case "mode__iew":
			v := value
			vModeIew = append(vModeIew, v)
		case "mode__niew":
			v := value
			vModeNiew = append(vModeNiew, v)
		case "mode__empty":
			if vModeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mode__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mode__empty' takes a boolean, got %q.", value))
				return
			}
			vModeEmpty = &v
		case "mode__regex":
			v := value
			vModeRegex = append(vModeRegex, v)
		case "mode__iregex":
			v := value
			vModeIregex = append(vModeIregex, v)
		case "mtu":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu' takes an integer, got %q.", value))
				return
			}
			vMtu = append(vMtu, v)
		case "mtu__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__n' takes an integer, got %q.", value))
				return
			}
			vMtun = append(vMtun, v)
		case "mtu__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__lt' takes an integer, got %q.", value))
				return
			}
			vMtuLt = append(vMtuLt, v)
		case "mtu__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__lte' takes an integer, got %q.", value))
				return
			}
			vMtuLte = append(vMtuLte, v)
		case "mtu__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__gt' takes an integer, got %q.", value))
				return
			}
			vMtuGt = append(vMtuGt, v)
		case "mtu__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__gte' takes an integer, got %q.", value))
				return
			}
			vMtuGte = append(vMtuGte, v)
		case "mtu__empty":
			if vMtuEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mtu__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mtu__empty' takes a boolean, got %q.", value))
				return
			}
			vMtuEmpty = &v
		case "vrf_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vrf_id' takes an integer, got %q.", value))
				return
			}
			vVrfID = append(vVrfID, v)
		case "vrf_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vrf_id__n' takes an integer, got %q.", value))
				return
			}
			vVrfIDn = append(vVrfIDn, v)
		case "parent_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_id' takes an integer, got %q.", value))
				return
			}
			vParentID = append(vParentID, v)
		case "parent_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'parent_id__n' takes an integer, got %q.", value))
				return
			}
			vParentIDn = append(vParentIDn, v)
		case "bridge_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'bridge_id' takes an integer, got %q.", value))
				return
			}
			vBridgeID = append(vBridgeID, v)
		case "bridge_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'bridge_id__n' takes an integer, got %q.", value))
				return
			}
			vBridgeIDn = append(vBridgeIDn, v)
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
		case "virtual_machine":
			v := value
			vVirtualMachine = append(vVirtualMachine, v)
		case "virtual_machine__n":
			v := value
			vVirtualMachinen = append(vVirtualMachinen, v)
		case "cluster":
			v := value
			vCluster = append(vCluster, v)
		case "cluster__n":
			v := value
			vClustern = append(vClustern, v)
		case "vrf":
			v := value
			vVrf = append(vVrf, v)
		case "vrf__n":
			v := value
			vVrfn = append(vVrfn, v)
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
		case "l2vpn":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn' takes an integer, got %q.", value))
				return
			}
			vL2vpn = append(vL2vpn, v)
		case "l2vpn__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn__n' takes an integer, got %q.", value))
				return
			}
			vL2vpnn = append(vL2vpnn, v)
		case "l2vpn_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn_id' takes an integer, got %q.", value))
				return
			}
			vL2vpnID = append(vL2vpnID, v)
		case "l2vpn_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'l2vpn_id__n' takes an integer, got %q.", value))
				return
			}
			vL2vpnIDn = append(vL2vpnIDn, v)
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
		case "primary_mac_address":
			v := value
			vPrimaryMacAddress = append(vPrimaryMacAddress, v)
		case "primary_mac_address__n":
			v := value
			vPrimaryMacAddressn = append(vPrimaryMacAddressn, v)
		case "primary_mac_address_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_mac_address_id' takes an integer, got %q.", value))
				return
			}
			vPrimaryMacAddressID = append(vPrimaryMacAddressID, v)
		case "primary_mac_address_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary_mac_address_id__n' takes an integer, got %q.", value))
				return
			}
			vPrimaryMacAddressIDn = append(vPrimaryMacAddressIDn, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
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
		case "vlan":
			if vVlan != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan' takes a single value.")
				return
			}
			v := value
			vVlan = &v
		case "vlan_id":
			if vVlanID != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_id' takes a single value.")
				return
			}
			v := value
			vVlanID = &v
		case "vlan_translation_policy":
			v := value
			vVlanTranslationPolicy = append(vVlanTranslationPolicy, v)
		case "vlan_translation_policy__n":
			v := value
			vVlanTranslationPolicyn = append(vVlanTranslationPolicyn, v)
		case "vlan_translation_policy_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_translation_policy_id' takes an integer, got %q.", value))
				return
			}
			vVlanTranslationPolicyID = append(vVlanTranslationPolicyID, v)
		case "vlan_translation_policy_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_translation_policy_id__n' takes an integer, got %q.", value))
				return
			}
			vVlanTranslationPolicyIDn = append(vVlanTranslationPolicyIDn, v)
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
	if len(vVirtualMachineID) > 0 {
		params.SetVirtualMachineID(vVirtualMachineID)
	}
	if len(vVirtualMachineIDn) > 0 {
		params.SetVirtualMachineIDn(vVirtualMachineIDn)
	}
	if len(vClusterID) > 0 {
		params.SetClusterID(vClusterID)
	}
	if len(vClusterIDn) > 0 {
		params.SetClusterIDn(vClusterIDn)
	}
	if len(vMacAddress) > 0 {
		params.SetMacAddress(vMacAddress)
	}
	if len(vMacAddressn) > 0 {
		params.SetMacAddressn(vMacAddressn)
	}
	if len(vMacAddressIc) > 0 {
		params.SetMacAddressIc(vMacAddressIc)
	}
	if len(vMacAddressNic) > 0 {
		params.SetMacAddressNic(vMacAddressNic)
	}
	if len(vMacAddressIe) > 0 {
		params.SetMacAddressIe(vMacAddressIe)
	}
	if len(vMacAddressNie) > 0 {
		params.SetMacAddressNie(vMacAddressNie)
	}
	if len(vMacAddressIsw) > 0 {
		params.SetMacAddressIsw(vMacAddressIsw)
	}
	if len(vMacAddressNisw) > 0 {
		params.SetMacAddressNisw(vMacAddressNisw)
	}
	if len(vMacAddressIew) > 0 {
		params.SetMacAddressIew(vMacAddressIew)
	}
	if len(vMacAddressNiew) > 0 {
		params.SetMacAddressNiew(vMacAddressNiew)
	}
	if len(vMacAddressRegex) > 0 {
		params.SetMacAddressRegex(vMacAddressRegex)
	}
	if len(vMacAddressIregex) > 0 {
		params.SetMacAddressIregex(vMacAddressIregex)
	}
	if vEnabled != nil {
		params.SetEnabled(vEnabled)
	}
	if len(vMode) > 0 {
		params.SetMode(vMode)
	}
	if len(vModen) > 0 {
		params.SetModen(vModen)
	}
	if len(vModeIc) > 0 {
		params.SetModeIc(vModeIc)
	}
	if len(vModeNic) > 0 {
		params.SetModeNic(vModeNic)
	}
	if len(vModeIe) > 0 {
		params.SetModeIe(vModeIe)
	}
	if len(vModeNie) > 0 {
		params.SetModeNie(vModeNie)
	}
	if len(vModeIsw) > 0 {
		params.SetModeIsw(vModeIsw)
	}
	if len(vModeNisw) > 0 {
		params.SetModeNisw(vModeNisw)
	}
	if len(vModeIew) > 0 {
		params.SetModeIew(vModeIew)
	}
	if len(vModeNiew) > 0 {
		params.SetModeNiew(vModeNiew)
	}
	if vModeEmpty != nil {
		params.SetModeEmpty(vModeEmpty)
	}
	if len(vModeRegex) > 0 {
		params.SetModeRegex(vModeRegex)
	}
	if len(vModeIregex) > 0 {
		params.SetModeIregex(vModeIregex)
	}
	if len(vMtu) > 0 {
		params.SetMtu(vMtu)
	}
	if len(vMtun) > 0 {
		params.SetMtun(vMtun)
	}
	if len(vMtuLt) > 0 {
		params.SetMtuLt(vMtuLt)
	}
	if len(vMtuLte) > 0 {
		params.SetMtuLte(vMtuLte)
	}
	if len(vMtuGt) > 0 {
		params.SetMtuGt(vMtuGt)
	}
	if len(vMtuGte) > 0 {
		params.SetMtuGte(vMtuGte)
	}
	if vMtuEmpty != nil {
		params.SetMtuEmpty(vMtuEmpty)
	}
	if len(vVrfID) > 0 {
		params.SetVrfID(vVrfID)
	}
	if len(vVrfIDn) > 0 {
		params.SetVrfIDn(vVrfIDn)
	}
	if len(vParentID) > 0 {
		params.SetParentID(vParentID)
	}
	if len(vParentIDn) > 0 {
		params.SetParentIDn(vParentIDn)
	}
	if len(vBridgeID) > 0 {
		params.SetBridgeID(vBridgeID)
	}
	if len(vBridgeIDn) > 0 {
		params.SetBridgeIDn(vBridgeIDn)
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
	if len(vVirtualMachine) > 0 {
		params.SetVirtualMachine(vVirtualMachine)
	}
	if len(vVirtualMachinen) > 0 {
		params.SetVirtualMachinen(vVirtualMachinen)
	}
	if len(vCluster) > 0 {
		params.SetCluster(vCluster)
	}
	if len(vClustern) > 0 {
		params.SetClustern(vClustern)
	}
	if len(vVrf) > 0 {
		params.SetVrf(vVrf)
	}
	if len(vVrfn) > 0 {
		params.SetVrfn(vVrfn)
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
	if len(vL2vpn) > 0 {
		params.SetL2vpn(vL2vpn)
	}
	if len(vL2vpnn) > 0 {
		params.SetL2vpnn(vL2vpnn)
	}
	if len(vL2vpnID) > 0 {
		params.SetL2vpnID(vL2vpnID)
	}
	if len(vL2vpnIDn) > 0 {
		params.SetL2vpnIDn(vL2vpnIDn)
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
	if len(vPrimaryMacAddress) > 0 {
		params.SetPrimaryMacAddress(vPrimaryMacAddress)
	}
	if len(vPrimaryMacAddressn) > 0 {
		params.SetPrimaryMacAddressn(vPrimaryMacAddressn)
	}
	if len(vPrimaryMacAddressID) > 0 {
		params.SetPrimaryMacAddressID(vPrimaryMacAddressID)
	}
	if len(vPrimaryMacAddressIDn) > 0 {
		params.SetPrimaryMacAddressIDn(vPrimaryMacAddressIDn)
	}
	if vQ != nil {
		params.SetQ(vQ)
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
	if vVlan != nil {
		params.SetVlan(vVlan)
	}
	if vVlanID != nil {
		params.SetVlanID(vVlanID)
	}
	if len(vVlanTranslationPolicy) > 0 {
		params.SetVlanTranslationPolicy(vVlanTranslationPolicy)
	}
	if len(vVlanTranslationPolicyn) > 0 {
		params.SetVlanTranslationPolicyn(vVlanTranslationPolicyn)
	}
	if len(vVlanTranslationPolicyID) > 0 {
		params.SetVlanTranslationPolicyID(vVlanTranslationPolicyID)
	}
	if len(vVlanTranslationPolicyIDn) > 0 {
		params.SetVlanTranslationPolicyIDn(vVlanTranslationPolicyIDn)
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
