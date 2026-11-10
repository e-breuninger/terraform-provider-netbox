// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
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
	_ datasource.DataSource              = (*ipAddressesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*ipAddressesDataSource)(nil)
)

// NewIPAddressesDataSource returns a new ip_addresses data source, which
// lists ip_address objects matching its filters.
func NewIPAddressesDataSource() datasource.DataSource {
	return &ipAddressesDataSource{}
}

type ipAddressesDataSource struct {
	client *netboxapi.Client
}

// ipAddressesDataSourceModel is the data source model: the filters, the limit and the matching
// ip_addresses.
type ipAddressesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"ip_addresses"`
}

// ipAddressResourceAttrTypes is the attribute type map of ipAddressResourceModel.
func ipAddressResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                           types.Int64Type,
		"ip_address":                   conv.IPAddressType{},
		"status":                       types.StringType,
		"role":                         types.StringType,
		"dns_name":                     types.StringType,
		"description":                  types.StringType,
		"comments":                     types.StringType,
		"tenant_id":                    types.Int64Type,
		"vrf_id":                       types.Int64Type,
		"nat_inside_id":                types.Int64Type,
		"nat_inside_address_id":        types.Int64Type,
		"nat_outside_ids":              types.SetType{ElemType: types.Int64Type},
		"assigned_object_type":         types.StringType,
		"assigned_object_id":           types.Int64Type,
		"device_interface_id":          types.Int64Type,
		"virtual_machine_interface_id": types.Int64Type,
		"assigned_object":              types.ObjectType{AttrTypes: ipAddressAssignedObjectAttrTypes()},
		"family":                       types.Int64Type,
		"owner_id":                     types.Int64Type,
		"created":                      types.StringType,
		"last_updated":                 types.StringType,
		"url":                          types.StringType,
		"tags":                         types.SetType{ElemType: types.StringType},
		"tags_all":                     types.SetType{ElemType: types.StringType},
		"custom_fields":                types.MapType{ElemType: types.StringType},
	}
}

func (d *ipAddressesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_addresses"
}

func (d *ipAddressesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists ip_address objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: address, assigned, assigned_object_id, assigned_object_id__empty, assigned_object_id__gt, assigned_object_id__gte, assigned_object_id__lt, assigned_object_id__lte, assigned_object_id__n, assigned_object_type, assigned_object_type__n, assigned_to_interface, contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device_id, dns_name, dns_name__empty, dns_name__ic, dns_name__ie, dns_name__iew, dns_name__iregex, dns_name__isw, dns_name__n, dns_name__nic, dns_name__nie, dns_name__niew, dns_name__nisw, dns_name__regex, family, fhrpgroup_id, fhrpgroup_id__n, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface, interface__n, interface_id, interface_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mask_length, mask_length__gte, mask_length__lte, nat_inside_id, nat_inside_id__n, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, parent, present_in_vrf, present_in_vrf_id, q, role, role__empty, role__ic, role__ie, role__iew, role__iregex, role__isw, role__n, role__nic, role__nie, role__niew, role__nisw, role__regex, service_id, service_id__n, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, virtual_machine, virtual_machine_id, vminterface, vminterface__n, vminterface_id, vminterface_id__n, vrf, vrf__n, vrf_id, vrf_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"ip_addresses": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching ip_addresses.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id of the IP address.",
						},
						"ip_address": schema.StringAttribute{
							CustomType:  conv.IPAddressType{},
							Computed:    true,
							Description: "IPv4 or IPv6 address with prefix length (e.g. 10.0.0.1/24).",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Operational status. One of: active, reserved, deprecated, dhcp, slaac.",
						},
						"role": schema.StringAttribute{
							Computed:    true,
							Description: "Functional role. One of: loopback, secondary, anycast, vip, vrrp, hsrp, glbp, carp.",
						},
						"dns_name": schema.StringAttribute{
							Computed:    true,
							Description: "Hostname or FQDN (not case-sensitive).",
						},
						"description": schema.StringAttribute{
							Computed: true,
						},
						"comments": schema.StringAttribute{
							Computed: true,
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"vrf_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VRF.",
						},
						"nat_inside_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the IP address this address is the NAT (outside) IP for.",
						},
						"nat_inside_address_id": schema.Int64Attribute{
							Computed:           true,
							DeprecationMessage: "Use nat_inside_id instead.",
							Description:        "Deprecated alias of nat_inside_id.",
						},
						"nat_outside_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the addresses that name this one as their nat_inside_id.",
						},
						"assigned_object_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the object the address is assigned to. Derived from device_interface_id or virtual_machine_interface_id when one of those is set; set it together with assigned_object_id for other object types. One of: dcim.interface, virtualization.vminterface, ipam.fhrpgroup.",
						},
						"assigned_object_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the object the address is assigned to (see assigned_object_type).",
						},
						"device_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device interface the address is assigned to (assigned_object_type dcim.interface). Conflicts with virtual_machine_interface_id and with setting the assigned_object_* pair directly.",
						},
						"virtual_machine_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual machine interface the address is assigned to (assigned_object_type virtualization.vminterface). Conflicts with device_interface_id and with setting the assigned_object_* pair directly.",
						},
						"assigned_object": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"id": schema.Int64Attribute{
									Computed: true,
								},
								"name": schema.StringAttribute{
									Computed: true,
								},
								"device": schema.SingleNestedAttribute{
									Attributes: map[string]schema.Attribute{
										"id": schema.Int64Attribute{
											Computed: true,
										},
										"name": schema.StringAttribute{
											Computed: true,
										},
									},
									Computed: true,
								},
							},
							Computed:    true,
							Description: "The interface the address is assigned to. device is null for VM interfaces.",
						},
						"family": schema.Int64Attribute{
							Computed:    true,
							Description: "IP family (4 or 6).",
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

func (d *ipAddressesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *ipAddressesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data ipAddressesDataSourceModel

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
	var responseDTOs []*netboxapi.IPAddressResponseDTO

	params := ipam.NewIpamIPAddressesListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vAddress []string
	var vDNSName []string
	var vDNSNamen []string
	var vDNSNameIc []string
	var vDNSNameNic []string
	var vDNSNameIe []string
	var vDNSNameNie []string
	var vDNSNameIsw []string
	var vDNSNameNisw []string
	var vDNSNameIew []string
	var vDNSNameNiew []string
	var vDNSNameEmpty *bool
	var vDNSNameRegex []string
	var vDNSNameIregex []string
	var vVrfID []int64
	var vVrfIDn []int64
	var vTenantID []int64
	var vTenantIDn []int64
	var vDevice []string
	var vInterfaceID []int64
	var vInterfaceIDn []int64
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
	var vRole []string
	var vRolen []string
	var vRoleIc []string
	var vRoleNic []string
	var vRoleIe []string
	var vRoleNie []string
	var vRoleIsw []string
	var vRoleNisw []string
	var vRoleIew []string
	var vRoleNiew []string
	var vRoleEmpty *bool
	var vRoleRegex []string
	var vRoleIregex []string
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
	var vVminterfaceID []int64
	var vVminterfaceIDn []int64
	var vDeviceID []int64
	var vParent []string
	var vTenant []string
	var vTenantn []string
	var vVrf []string
	var vVrfn []string
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vAssigned *bool
	var vAssignedObjectID []int64
	var vAssignedObjectIDn []int64
	var vAssignedObjectIDLt []int64
	var vAssignedObjectIDLte []int64
	var vAssignedObjectIDGt []int64
	var vAssignedObjectIDGte []int64
	var vAssignedObjectIDEmpty *bool
	var vAssignedObjectType []string
	var vAssignedObjectTypen []string
	var vAssignedToInterface *bool
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
	var vFamily *float64
	var vFhrpgroupID []int64
	var vFhrpgroupIDn []int64
	var vInterface []string
	var vInterfacen []string
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vMaskLengthGte *float64
	var vMaskLengthLte *float64
	var vMaskLength []int64
	var vNatInsideID []int64
	var vNatInsideIDn []int64
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPresentInVrf *string
	var vPresentInVrfID *string
	var vQ *string
	var vServiceID []int64
	var vServiceIDn []int64
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
	var vVirtualMachine []string
	var vVirtualMachineID []int64
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
		case "address":
			v := value
			vAddress = append(vAddress, v)
		case "dns_name":
			v := value
			vDNSName = append(vDNSName, v)
		case "dns_name__n":
			v := value
			vDNSNamen = append(vDNSNamen, v)
		case "dns_name__ic":
			v := value
			vDNSNameIc = append(vDNSNameIc, v)
		case "dns_name__nic":
			v := value
			vDNSNameNic = append(vDNSNameNic, v)
		case "dns_name__ie":
			v := value
			vDNSNameIe = append(vDNSNameIe, v)
		case "dns_name__nie":
			v := value
			vDNSNameNie = append(vDNSNameNie, v)
		case "dns_name__isw":
			v := value
			vDNSNameIsw = append(vDNSNameIsw, v)
		case "dns_name__nisw":
			v := value
			vDNSNameNisw = append(vDNSNameNisw, v)
		case "dns_name__iew":
			v := value
			vDNSNameIew = append(vDNSNameIew, v)
		case "dns_name__niew":
			v := value
			vDNSNameNiew = append(vDNSNameNiew, v)
		case "dns_name__empty":
			if vDNSNameEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'dns_name__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'dns_name__empty' takes a boolean, got %q.", value))
				return
			}
			vDNSNameEmpty = &v
		case "dns_name__regex":
			v := value
			vDNSNameRegex = append(vDNSNameRegex, v)
		case "dns_name__iregex":
			v := value
			vDNSNameIregex = append(vDNSNameIregex, v)
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
		case "device":
			v := value
			vDevice = append(vDevice, v)
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
		case "role":
			v := value
			vRole = append(vRole, v)
		case "role__n":
			v := value
			vRolen = append(vRolen, v)
		case "role__ic":
			v := value
			vRoleIc = append(vRoleIc, v)
		case "role__nic":
			v := value
			vRoleNic = append(vRoleNic, v)
		case "role__ie":
			v := value
			vRoleIe = append(vRoleIe, v)
		case "role__nie":
			v := value
			vRoleNie = append(vRoleNie, v)
		case "role__isw":
			v := value
			vRoleIsw = append(vRoleIsw, v)
		case "role__nisw":
			v := value
			vRoleNisw = append(vRoleNisw, v)
		case "role__iew":
			v := value
			vRoleIew = append(vRoleIew, v)
		case "role__niew":
			v := value
			vRoleNiew = append(vRoleNiew, v)
		case "role__empty":
			if vRoleEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'role__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'role__empty' takes a boolean, got %q.", value))
				return
			}
			vRoleEmpty = &v
		case "role__regex":
			v := value
			vRoleRegex = append(vRoleRegex, v)
		case "role__iregex":
			v := value
			vRoleIregex = append(vRoleIregex, v)
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
		case "device_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id' takes an integer, got %q.", value))
				return
			}
			vDeviceID = append(vDeviceID, v)
		case "parent":
			v := value
			vParent = append(vParent, v)
		case "tenant":
			v := value
			vTenant = append(vTenant, v)
		case "tenant__n":
			v := value
			vTenantn = append(vTenantn, v)
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
		case "assigned":
			if vAssigned != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'assigned' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned' takes a boolean, got %q.", value))
				return
			}
			vAssigned = &v
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
		case "assigned_to_interface":
			if vAssignedToInterface != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'assigned_to_interface' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'assigned_to_interface' takes a boolean, got %q.", value))
				return
			}
			vAssignedToInterface = &v
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
		case "family":
			if vFamily != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'family' takes a single value.")
				return
			}
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'family' takes a number, got %q.", value))
				return
			}
			vFamily = &v
		case "fhrpgroup_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'fhrpgroup_id' takes an integer, got %q.", value))
				return
			}
			vFhrpgroupID = append(vFhrpgroupID, v)
		case "fhrpgroup_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'fhrpgroup_id__n' takes an integer, got %q.", value))
				return
			}
			vFhrpgroupIDn = append(vFhrpgroupIDn, v)
		case "interface":
			v := value
			vInterface = append(vInterface, v)
		case "interface__n":
			v := value
			vInterfacen = append(vInterfacen, v)
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
		case "mask_length__gte":
			if vMaskLengthGte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mask_length__gte' takes a single value.")
				return
			}
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mask_length__gte' takes a number, got %q.", value))
				return
			}
			vMaskLengthGte = &v
		case "mask_length__lte":
			if vMaskLengthLte != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'mask_length__lte' takes a single value.")
				return
			}
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mask_length__lte' takes a number, got %q.", value))
				return
			}
			vMaskLengthLte = &v
		case "mask_length":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'mask_length' takes an integer, got %q.", value))
				return
			}
			vMaskLength = append(vMaskLength, v)
		case "nat_inside_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'nat_inside_id' takes an integer, got %q.", value))
				return
			}
			vNatInsideID = append(vNatInsideID, v)
		case "nat_inside_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'nat_inside_id__n' takes an integer, got %q.", value))
				return
			}
			vNatInsideIDn = append(vNatInsideIDn, v)
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
		case "present_in_vrf":
			if vPresentInVrf != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'present_in_vrf' takes a single value.")
				return
			}
			v := value
			vPresentInVrf = &v
		case "present_in_vrf_id":
			if vPresentInVrfID != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'present_in_vrf_id' takes a single value.")
				return
			}
			v := value
			vPresentInVrfID = &v
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "service_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'service_id' takes an integer, got %q.", value))
				return
			}
			vServiceID = append(vServiceID, v)
		case "service_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'service_id__n' takes an integer, got %q.", value))
				return
			}
			vServiceIDn = append(vServiceIDn, v)
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
		case "virtual_machine":
			v := value
			vVirtualMachine = append(vVirtualMachine, v)
		case "virtual_machine_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'virtual_machine_id' takes an integer, got %q.", value))
				return
			}
			vVirtualMachineID = append(vVirtualMachineID, v)
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
	if len(vAddress) > 0 {
		params.SetAddress(vAddress)
	}
	if len(vDNSName) > 0 {
		params.SetDNSName(vDNSName)
	}
	if len(vDNSNamen) > 0 {
		params.SetDNSNamen(vDNSNamen)
	}
	if len(vDNSNameIc) > 0 {
		params.SetDNSNameIc(vDNSNameIc)
	}
	if len(vDNSNameNic) > 0 {
		params.SetDNSNameNic(vDNSNameNic)
	}
	if len(vDNSNameIe) > 0 {
		params.SetDNSNameIe(vDNSNameIe)
	}
	if len(vDNSNameNie) > 0 {
		params.SetDNSNameNie(vDNSNameNie)
	}
	if len(vDNSNameIsw) > 0 {
		params.SetDNSNameIsw(vDNSNameIsw)
	}
	if len(vDNSNameNisw) > 0 {
		params.SetDNSNameNisw(vDNSNameNisw)
	}
	if len(vDNSNameIew) > 0 {
		params.SetDNSNameIew(vDNSNameIew)
	}
	if len(vDNSNameNiew) > 0 {
		params.SetDNSNameNiew(vDNSNameNiew)
	}
	if vDNSNameEmpty != nil {
		params.SetDNSNameEmpty(vDNSNameEmpty)
	}
	if len(vDNSNameRegex) > 0 {
		params.SetDNSNameRegex(vDNSNameRegex)
	}
	if len(vDNSNameIregex) > 0 {
		params.SetDNSNameIregex(vDNSNameIregex)
	}
	if len(vVrfID) > 0 {
		params.SetVrfID(vVrfID)
	}
	if len(vVrfIDn) > 0 {
		params.SetVrfIDn(vVrfIDn)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
	}
	if len(vDevice) > 0 {
		params.SetDevice(vDevice)
	}
	if len(vInterfaceID) > 0 {
		params.SetInterfaceID(vInterfaceID)
	}
	if len(vInterfaceIDn) > 0 {
		params.SetInterfaceIDn(vInterfaceIDn)
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
	if len(vRole) > 0 {
		params.SetRole(vRole)
	}
	if len(vRolen) > 0 {
		params.SetRolen(vRolen)
	}
	if len(vRoleIc) > 0 {
		params.SetRoleIc(vRoleIc)
	}
	if len(vRoleNic) > 0 {
		params.SetRoleNic(vRoleNic)
	}
	if len(vRoleIe) > 0 {
		params.SetRoleIe(vRoleIe)
	}
	if len(vRoleNie) > 0 {
		params.SetRoleNie(vRoleNie)
	}
	if len(vRoleIsw) > 0 {
		params.SetRoleIsw(vRoleIsw)
	}
	if len(vRoleNisw) > 0 {
		params.SetRoleNisw(vRoleNisw)
	}
	if len(vRoleIew) > 0 {
		params.SetRoleIew(vRoleIew)
	}
	if len(vRoleNiew) > 0 {
		params.SetRoleNiew(vRoleNiew)
	}
	if vRoleEmpty != nil {
		params.SetRoleEmpty(vRoleEmpty)
	}
	if len(vRoleRegex) > 0 {
		params.SetRoleRegex(vRoleRegex)
	}
	if len(vRoleIregex) > 0 {
		params.SetRoleIregex(vRoleIregex)
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
	if len(vVminterfaceID) > 0 {
		params.SetVminterfaceID(vVminterfaceID)
	}
	if len(vVminterfaceIDn) > 0 {
		params.SetVminterfaceIDn(vVminterfaceIDn)
	}
	if len(vDeviceID) > 0 {
		params.SetDeviceID(vDeviceID)
	}
	if len(vParent) > 0 {
		params.SetParent(vParent)
	}
	if len(vTenant) > 0 {
		params.SetTenant(vTenant)
	}
	if len(vTenantn) > 0 {
		params.SetTenantn(vTenantn)
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
	if vAssigned != nil {
		params.SetAssigned(vAssigned)
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
	if vAssignedToInterface != nil {
		params.SetAssignedToInterface(vAssignedToInterface)
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
	if vFamily != nil {
		params.SetFamily(vFamily)
	}
	if len(vFhrpgroupID) > 0 {
		params.SetFhrpgroupID(vFhrpgroupID)
	}
	if len(vFhrpgroupIDn) > 0 {
		params.SetFhrpgroupIDn(vFhrpgroupIDn)
	}
	if len(vInterface) > 0 {
		params.SetInterface(vInterface)
	}
	if len(vInterfacen) > 0 {
		params.SetInterfacen(vInterfacen)
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
	if vMaskLengthGte != nil {
		params.SetMaskLengthGte(vMaskLengthGte)
	}
	if vMaskLengthLte != nil {
		params.SetMaskLengthLte(vMaskLengthLte)
	}
	if len(vMaskLength) > 0 {
		params.SetMaskLength(vMaskLength)
	}
	if len(vNatInsideID) > 0 {
		params.SetNatInsideID(vNatInsideID)
	}
	if len(vNatInsideIDn) > 0 {
		params.SetNatInsideIDn(vNatInsideIDn)
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
	if vPresentInVrf != nil {
		params.SetPresentInVrf(vPresentInVrf)
	}
	if vPresentInVrfID != nil {
		params.SetPresentInVrfID(vPresentInVrfID)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vServiceID) > 0 {
		params.SetServiceID(vServiceID)
	}
	if len(vServiceIDn) > 0 {
		params.SetServiceIDn(vServiceIDn)
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
	if len(vVirtualMachine) > 0 {
		params.SetVirtualMachine(vVirtualMachine)
	}
	if len(vVirtualMachineID) > 0 {
		params.SetVirtualMachineID(vVirtualMachineID)
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
		res, err := d.client.Ipam.IpamIPAddressesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_ip_addresses", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.IPAddressResponseDTOFromGoNetbox(goNetboxModel))
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
	resourceWithHooks := &ipAddressResource{client: d.client}
	items := make([]ipAddressResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model ipAddressResourceModel
		resp.Diagnostics.Append(flattenIPAddress(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: ipAddressResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
