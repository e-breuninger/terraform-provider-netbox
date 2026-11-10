// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*l2vpnTerminationsDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*l2vpnTerminationsDataSource)(nil)
)

// NewL2vpnTerminationsDataSource returns a new l2vpn_terminations data source, which
// lists l2vpn_termination objects matching its filters.
func NewL2vpnTerminationsDataSource() datasource.DataSource {
	return &l2vpnTerminationsDataSource{}
}

type l2vpnTerminationsDataSource struct {
	client *netboxapi.Client
}

// l2vpnTerminationsDataSourceModel is the data source model: the filters, the limit and the matching
// l2vpn_terminations.
type l2vpnTerminationsDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"l2vpn_terminations"`
}

// l2vpnTerminationResourceAttrTypes is the attribute type map of l2vpnTerminationResourceModel.
func l2vpnTerminationResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                           types.Int64Type,
		"l2vpn_id":                     types.Int64Type,
		"assigned_object_type":         types.StringType,
		"assigned_object_id":           types.Int64Type,
		"device_interface_id":          types.Int64Type,
		"virtual_machine_interface_id": types.Int64Type,
		"vlan_id":                      types.Int64Type,
		"created":                      types.StringType,
		"last_updated":                 types.StringType,
		"url":                          types.StringType,
		"tags":                         types.SetType{ElemType: types.StringType},
		"tags_all":                     types.SetType{ElemType: types.StringType},
		"custom_fields":                types.MapType{ElemType: types.StringType},
	}
}

func (d *l2vpnTerminationsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_l2vpn_terminations"
}

func (d *l2vpnTerminationsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:VPN Tunnels:Lists l2vpn_termination objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: assigned_object_id, assigned_object_id__empty, assigned_object_id__gt, assigned_object_id__gte, assigned_object_id__lt, assigned_object_id__lte, assigned_object_id__n, assigned_object_type, assigned_object_type__n, assigned_object_type_id, assigned_object_type_id__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, device, device__n, device_id, device_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface, interface__n, interface_id, interface_id__n, l2vpn, l2vpn__n, l2vpn_id, l2vpn_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, q, region, region_id, site, site_id, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, virtual_machine, virtual_machine__n, virtual_machine_id, virtual_machine_id__n, vlan, vlan__n, vlan_id, vlan_id__n, vlan_vid, vlan_vid__empty, vlan_vid__gt, vlan_vid__gte, vlan_vid__lt, vlan_vid__lte, vlan_vid__n, vminterface, vminterface__n, vminterface_id, vminterface_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"l2vpn_terminations": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching l2vpn_terminations.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"l2vpn_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the L2VPN.",
						},
						"assigned_object_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the terminating object. Derived from device_interface_id, virtual_machine_interface_id or vlan_id when one of those is set; set it together with assigned_object_id otherwise. One assignment is required. One of: dcim.interface, virtualization.vminterface, ipam.vlan.",
						},
						"assigned_object_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating object (see assigned_object_type).",
						},
						"device_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating device interface (assigned_object_type dcim.interface). Conflicts with the other aliases and with setting the assigned_object_* pair directly.",
						},
						"virtual_machine_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating virtual machine interface (assigned_object_type virtualization.vminterface). Conflicts with the other aliases and with setting the assigned_object_* pair directly.",
						},
						"vlan_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the terminating VLAN (assigned_object_type ipam.vlan). Conflicts with the other aliases and with setting the assigned_object_* pair directly.",
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

func (d *l2vpnTerminationsDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *l2vpnTerminationsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data l2vpnTerminationsDataSourceModel

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
	var responseDTOs []*netboxapi.L2vpnTerminationResponseDTO

	params := vpn.NewVpnL2vpnTerminationsListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vL2vpnID []int64
	var vL2vpnIDn []int64
	var vVlanID []int64
	var vVlanIDn []int64
	var vInterfaceID []int64
	var vInterfaceIDn []int64
	var vVminterfaceID []int64
	var vVminterfaceIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAssignedObjectID []int64
	var vAssignedObjectIDn []int64
	var vAssignedObjectIDLt []int64
	var vAssignedObjectIDLte []int64
	var vAssignedObjectIDGt []int64
	var vAssignedObjectIDGte []int64
	var vAssignedObjectIDEmpty *bool
	var vAssignedObjectType []string
	var vAssignedObjectTypen []string
	var vAssignedObjectTypeID []int64
	var vAssignedObjectTypeIDn []int64
	var vCreated []string
	var vCreatedn []string
	var vCreatedLt []string
	var vCreatedLte []string
	var vCreatedGt []string
	var vCreatedGte []string
	var vCreatedEmpty []string
	var vDevice []string
	var vDevicen []string
	var vDeviceID []int64
	var vDeviceIDn []int64
	var vInterface []string
	var vInterfacen []string
	var vL2vpn []string
	var vL2vpnn []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vQ *string
	var vRegion []string
	var vRegionID []int64
	var vSite []string
	var vSiteID []int64
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vVirtualMachine []string
	var vVirtualMachinen []string
	var vVirtualMachineID []int64
	var vVirtualMachineIDn []int64
	var vVlan []string
	var vVlann []string
	var vVlanVid *int64
	var vVlanVidn *int64
	var vVlanVidLt *int64
	var vVlanVidLte *int64
	var vVlanVidGt *int64
	var vVlanVidGte *int64
	var vVlanVidEmpty *int64
	var vVminterface []string
	var vVminterfacen []string
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
		case "vlan_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_id' takes an integer, got %q.", value))
				return
			}
			vVlanID = append(vVlanID, v)
		case "vlan_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_id__n' takes an integer, got %q.", value))
				return
			}
			vVlanIDn = append(vVlanIDn, v)
		case "interface_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_id' takes an integer, got %q.", value))
				return
			}
			vInterfaceID = append(vInterfaceID, v)
		case "interface_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_id__n' takes an integer, got %q.", value))
				return
			}
			vInterfaceIDn = append(vInterfaceIDn, v)
		case "vminterface_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vminterface_id' takes an integer, got %q.", value))
				return
			}
			vVminterfaceID = append(vVminterfaceID, v)
		case "vminterface_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vminterface_id__n' takes an integer, got %q.", value))
				return
			}
			vVminterfaceIDn = append(vVminterfaceIDn, v)
		case "tag":
			v := value
			vTag = append(vTag, v)
		case "tag__n":
			v := value
			vTagn = append(vTagn, v)
		case "tag__any":
			v := value
			vTagAny = append(vTagAny, v)
		case "assigned_object_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_id' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectID = append(vAssignedObjectID, v)
		case "assigned_object_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_id__n' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectIDn = append(vAssignedObjectIDn, v)
		case "assigned_object_id__lt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_id__lt' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectIDLt = append(vAssignedObjectIDLt, v)
		case "assigned_object_id__lte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_id__lte' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectIDLte = append(vAssignedObjectIDLte, v)
		case "assigned_object_id__gt":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_id__gt' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectIDGt = append(vAssignedObjectIDGt, v)
		case "assigned_object_id__gte":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_id__gte' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectIDGte = append(vAssignedObjectIDGte, v)
		case "assigned_object_id__empty":
			if vAssignedObjectIDEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'assigned_object_id__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_id__empty' takes a boolean, got %q.", value))
				return
			}
			vAssignedObjectIDEmpty = &v
		case "assigned_object_type":
			v := value
			vAssignedObjectType = append(vAssignedObjectType, v)
		case "assigned_object_type__n":
			v := value
			vAssignedObjectTypen = append(vAssignedObjectTypen, v)
		case "assigned_object_type_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_type_id' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectTypeID = append(vAssignedObjectTypeID, v)
		case "assigned_object_type_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_object_type_id__n' takes an integer, got %q.", value))
				return
			}
			vAssignedObjectTypeIDn = append(vAssignedObjectTypeIDn, v)
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
		case "device":
			v := value
			vDevice = append(vDevice, v)
		case "device__n":
			v := value
			vDevicen = append(vDevicen, v)
		case "device_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id' takes an integer, got %q.", value))
				return
			}
			vDeviceID = append(vDeviceID, v)
		case "device_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id__n' takes an integer, got %q.", value))
				return
			}
			vDeviceIDn = append(vDeviceIDn, v)
		case "interface":
			v := value
			vInterface = append(vInterface, v)
		case "interface__n":
			v := value
			vInterfacen = append(vInterfacen, v)
		case "l2vpn":
			v := value
			vL2vpn = append(vL2vpn, v)
		case "l2vpn__n":
			v := value
			vL2vpnn = append(vL2vpnn, v)
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
		case "region_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'region_id' takes an integer, got %q.", value))
				return
			}
			vRegionID = append(vRegionID, v)
		case "site":
			v := value
			vSite = append(vSite, v)
		case "site_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id' takes an integer, got %q.", value))
				return
			}
			vSiteID = append(vSiteID, v)
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
		case "virtual_machine":
			v := value
			vVirtualMachine = append(vVirtualMachine, v)
		case "virtual_machine__n":
			v := value
			vVirtualMachinen = append(vVirtualMachinen, v)
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
		case "vlan":
			v := value
			vVlan = append(vVlan, v)
		case "vlan__n":
			v := value
			vVlann = append(vVlann, v)
		case "vlan_vid":
			if vVlanVid != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid' takes an integer, got %q.", value))
				return
			}
			vVlanVid = &v
		case "vlan_vid__n":
			if vVlanVidn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__n' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__n' takes an integer, got %q.", value))
				return
			}
			vVlanVidn = &v
		case "vlan_vid__lt":
			if vVlanVidLt != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__lt' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__lt' takes an integer, got %q.", value))
				return
			}
			vVlanVidLt = &v
		case "vlan_vid__lte":
			if vVlanVidLte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__lte' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__lte' takes an integer, got %q.", value))
				return
			}
			vVlanVidLte = &v
		case "vlan_vid__gt":
			if vVlanVidGt != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__gt' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__gt' takes an integer, got %q.", value))
				return
			}
			vVlanVidGt = &v
		case "vlan_vid__gte":
			if vVlanVidGte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__gte' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__gte' takes an integer, got %q.", value))
				return
			}
			vVlanVidGte = &v
		case "vlan_vid__empty":
			if vVlanVidEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'vlan_vid__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'vlan_vid__empty' takes an integer, got %q.", value))
				return
			}
			vVlanVidEmpty = &v
		case "vminterface":
			v := value
			vVminterface = append(vVminterface, v)
		case "vminterface__n":
			v := value
			vVminterfacen = append(vVminterfacen, v)
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
	if len(vL2vpnID) > 0 {
		params.SetL2vpnID(vL2vpnID)
	}
	if len(vL2vpnIDn) > 0 {
		params.SetL2vpnIDn(vL2vpnIDn)
	}
	if len(vVlanID) > 0 {
		params.SetVlanID(vVlanID)
	}
	if len(vVlanIDn) > 0 {
		params.SetVlanIDn(vVlanIDn)
	}
	if len(vInterfaceID) > 0 {
		params.SetInterfaceID(vInterfaceID)
	}
	if len(vInterfaceIDn) > 0 {
		params.SetInterfaceIDn(vInterfaceIDn)
	}
	if len(vVminterfaceID) > 0 {
		params.SetVminterfaceID(vVminterfaceID)
	}
	if len(vVminterfaceIDn) > 0 {
		params.SetVminterfaceIDn(vVminterfaceIDn)
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
	if len(vAssignedObjectID) > 0 {
		params.SetAssignedObjectID(vAssignedObjectID)
	}
	if len(vAssignedObjectIDn) > 0 {
		params.SetAssignedObjectIDn(vAssignedObjectIDn)
	}
	if len(vAssignedObjectIDLt) > 0 {
		params.SetAssignedObjectIDLt(vAssignedObjectIDLt)
	}
	if len(vAssignedObjectIDLte) > 0 {
		params.SetAssignedObjectIDLte(vAssignedObjectIDLte)
	}
	if len(vAssignedObjectIDGt) > 0 {
		params.SetAssignedObjectIDGt(vAssignedObjectIDGt)
	}
	if len(vAssignedObjectIDGte) > 0 {
		params.SetAssignedObjectIDGte(vAssignedObjectIDGte)
	}
	if vAssignedObjectIDEmpty != nil {
		params.SetAssignedObjectIDEmpty(vAssignedObjectIDEmpty)
	}
	if len(vAssignedObjectType) > 0 {
		params.SetAssignedObjectType(vAssignedObjectType)
	}
	if len(vAssignedObjectTypen) > 0 {
		params.SetAssignedObjectTypen(vAssignedObjectTypen)
	}
	if len(vAssignedObjectTypeID) > 0 {
		params.SetAssignedObjectTypeID(vAssignedObjectTypeID)
	}
	if len(vAssignedObjectTypeIDn) > 0 {
		params.SetAssignedObjectTypeIDn(vAssignedObjectTypeIDn)
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
	if len(vDevice) > 0 {
		params.SetDevice(vDevice)
	}
	if len(vDevicen) > 0 {
		params.SetDevicen(vDevicen)
	}
	if len(vDeviceID) > 0 {
		params.SetDeviceID(vDeviceID)
	}
	if len(vDeviceIDn) > 0 {
		params.SetDeviceIDn(vDeviceIDn)
	}
	if len(vInterface) > 0 {
		params.SetInterface(vInterface)
	}
	if len(vInterfacen) > 0 {
		params.SetInterfacen(vInterfacen)
	}
	if len(vL2vpn) > 0 {
		params.SetL2vpn(vL2vpn)
	}
	if len(vL2vpnn) > 0 {
		params.SetL2vpnn(vL2vpnn)
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
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRegion) > 0 {
		params.SetRegion(vRegion)
	}
	if len(vRegionID) > 0 {
		params.SetRegionID(vRegionID)
	}
	if len(vSite) > 0 {
		params.SetSite(vSite)
	}
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
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
	if len(vVirtualMachine) > 0 {
		params.SetVirtualMachine(vVirtualMachine)
	}
	if len(vVirtualMachinen) > 0 {
		params.SetVirtualMachinen(vVirtualMachinen)
	}
	if len(vVirtualMachineID) > 0 {
		params.SetVirtualMachineID(vVirtualMachineID)
	}
	if len(vVirtualMachineIDn) > 0 {
		params.SetVirtualMachineIDn(vVirtualMachineIDn)
	}
	if len(vVlan) > 0 {
		params.SetVlan(vVlan)
	}
	if len(vVlann) > 0 {
		params.SetVlann(vVlann)
	}
	if vVlanVid != nil {
		params.SetVlanVid(vVlanVid)
	}
	if vVlanVidn != nil {
		params.SetVlanVidn(vVlanVidn)
	}
	if vVlanVidLt != nil {
		params.SetVlanVidLt(vVlanVidLt)
	}
	if vVlanVidLte != nil {
		params.SetVlanVidLte(vVlanVidLte)
	}
	if vVlanVidGt != nil {
		params.SetVlanVidGt(vVlanVidGt)
	}
	if vVlanVidGte != nil {
		params.SetVlanVidGte(vVlanVidGte)
	}
	if vVlanVidEmpty != nil {
		params.SetVlanVidEmpty(vVlanVidEmpty)
	}
	if len(vVminterface) > 0 {
		params.SetVminterface(vVminterface)
	}
	if len(vVminterfacen) > 0 {
		params.SetVminterfacen(vVminterfacen)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Vpn.VpnL2vpnTerminationsListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_l2vpn_terminations", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.L2vpnTerminationResponseDTOFromGoNetbox(goNetboxModel))
		}
		offset += int64(len(res.Payload.Results))
		if res.Payload.Next == nil || *res.Payload.Next == "" {
			break
		}
	}

	if resp.Diagnostics.HasError() {
		return
	}
	// The resource's post_read companion hook (spec hooks) runs on every item, as on the resource
	// itself (see the singular data source).
	resourceWithHooks := &l2vpnTerminationResource{client: d.client}
	items := make([]l2vpnTerminationResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model l2vpnTerminationResourceModel
		resp.Diagnostics.Append(flattenL2vpnTermination(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: l2vpnTerminationResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
