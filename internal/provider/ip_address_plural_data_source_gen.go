// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                        netboxapi.FilterInt64s(params.SetID),
		"id__n":                     netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                    netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                   netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                    netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                   netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                 netboxapi.FilterBool(params.SetIDEmpty),
		"address":                   netboxapi.FilterStrings(params.SetAddress),
		"dns_name":                  netboxapi.FilterStrings(params.SetDNSName),
		"dns_name__n":               netboxapi.FilterStrings(params.SetDNSNamen),
		"dns_name__ic":              netboxapi.FilterStrings(params.SetDNSNameIc),
		"dns_name__nic":             netboxapi.FilterStrings(params.SetDNSNameNic),
		"dns_name__ie":              netboxapi.FilterStrings(params.SetDNSNameIe),
		"dns_name__nie":             netboxapi.FilterStrings(params.SetDNSNameNie),
		"dns_name__isw":             netboxapi.FilterStrings(params.SetDNSNameIsw),
		"dns_name__nisw":            netboxapi.FilterStrings(params.SetDNSNameNisw),
		"dns_name__iew":             netboxapi.FilterStrings(params.SetDNSNameIew),
		"dns_name__niew":            netboxapi.FilterStrings(params.SetDNSNameNiew),
		"dns_name__empty":           netboxapi.FilterBool(params.SetDNSNameEmpty),
		"dns_name__regex":           netboxapi.FilterStrings(params.SetDNSNameRegex),
		"dns_name__iregex":          netboxapi.FilterStrings(params.SetDNSNameIregex),
		"vrf_id":                    netboxapi.FilterInt64s(params.SetVrfID),
		"vrf_id__n":                 netboxapi.FilterInt64s(params.SetVrfIDn),
		"tenant_id":                 netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":              netboxapi.FilterInt64s(params.SetTenantIDn),
		"device":                    netboxapi.FilterStrings(params.SetDevice),
		"interface_id":              netboxapi.FilterInt64s(params.SetInterfaceID),
		"interface_id__n":           netboxapi.FilterInt64s(params.SetInterfaceIDn),
		"status":                    netboxapi.FilterStrings(params.SetStatus),
		"status__n":                 netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":                netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":               netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":                netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":               netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":               netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":              netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":               netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":              netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":             netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":             netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":            netboxapi.FilterStrings(params.SetStatusIregex),
		"role":                      netboxapi.FilterStrings(params.SetRole),
		"role__n":                   netboxapi.FilterStrings(params.SetRolen),
		"role__ic":                  netboxapi.FilterStrings(params.SetRoleIc),
		"role__nic":                 netboxapi.FilterStrings(params.SetRoleNic),
		"role__ie":                  netboxapi.FilterStrings(params.SetRoleIe),
		"role__nie":                 netboxapi.FilterStrings(params.SetRoleNie),
		"role__isw":                 netboxapi.FilterStrings(params.SetRoleIsw),
		"role__nisw":                netboxapi.FilterStrings(params.SetRoleNisw),
		"role__iew":                 netboxapi.FilterStrings(params.SetRoleIew),
		"role__niew":                netboxapi.FilterStrings(params.SetRoleNiew),
		"role__empty":               netboxapi.FilterBool(params.SetRoleEmpty),
		"role__regex":               netboxapi.FilterStrings(params.SetRoleRegex),
		"role__iregex":              netboxapi.FilterStrings(params.SetRoleIregex),
		"description":               netboxapi.FilterStrings(params.SetDescription),
		"description__n":            netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":           netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":          netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":           netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":          netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":          netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":         netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":          netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":         netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":        netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":        netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":       netboxapi.FilterStrings(params.SetDescriptionIregex),
		"vminterface_id":            netboxapi.FilterInt64s(params.SetVminterfaceID),
		"vminterface_id__n":         netboxapi.FilterInt64s(params.SetVminterfaceIDn),
		"device_id":                 netboxapi.FilterInt64s(params.SetDeviceID),
		"parent":                    netboxapi.FilterStrings(params.SetParent),
		"tenant":                    netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":                 netboxapi.FilterStrings(params.SetTenantn),
		"vrf":                       netboxapi.FilterStrings(params.SetVrf),
		"vrf__n":                    netboxapi.FilterStrings(params.SetVrfn),
		"owner_id":                  netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":               netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                       netboxapi.FilterStrings(params.SetTag),
		"tag__n":                    netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                  netboxapi.FilterStrings(params.SetTagAny),
		"assigned":                  netboxapi.FilterBool(params.SetAssigned),
		"assigned_object_id":        netboxapi.FilterInt64s(params.SetAssignedObjectID),
		"assigned_object_id__n":     netboxapi.FilterInt64s(params.SetAssignedObjectIDn),
		"assigned_object_id__lt":    netboxapi.FilterInt64s(params.SetAssignedObjectIDLt),
		"assigned_object_id__lte":   netboxapi.FilterInt64s(params.SetAssignedObjectIDLte),
		"assigned_object_id__gt":    netboxapi.FilterInt64s(params.SetAssignedObjectIDGt),
		"assigned_object_id__gte":   netboxapi.FilterInt64s(params.SetAssignedObjectIDGte),
		"assigned_object_id__empty": netboxapi.FilterBool(params.SetAssignedObjectIDEmpty),
		"assigned_object_type":      netboxapi.FilterStrings(params.SetAssignedObjectType),
		"assigned_object_type__n":   netboxapi.FilterStrings(params.SetAssignedObjectTypen),
		"assigned_to_interface":     netboxapi.FilterBool(params.SetAssignedToInterface),
		"contact":                   netboxapi.FilterInt64s(params.SetContact),
		"contact__n":                netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":             netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":          netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":              netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":           netboxapi.FilterInt64s(params.SetContactRolen),
		"created":                   netboxapi.FilterStrings(params.SetCreated),
		"created__n":                netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":               netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":              netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":               netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":              netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":            netboxapi.FilterStrings(params.SetCreatedEmpty),
		"family":                    netboxapi.FilterFloat64(params.SetFamily),
		"fhrpgroup_id":              netboxapi.FilterInt64s(params.SetFhrpgroupID),
		"fhrpgroup_id__n":           netboxapi.FilterInt64s(params.SetFhrpgroupIDn),
		"interface":                 netboxapi.FilterStrings(params.SetInterface),
		"interface__n":              netboxapi.FilterStrings(params.SetInterfacen),
		"last_updated":              netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":           netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":          netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":         netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":          netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":         netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":       netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"mask_length__gte":          netboxapi.FilterFloat64(params.SetMaskLengthGte),
		"mask_length__lte":          netboxapi.FilterFloat64(params.SetMaskLengthLte),
		"mask_length":               netboxapi.FilterInt64s(params.SetMaskLength),
		"nat_inside_id":             netboxapi.FilterInt64s(params.SetNatInsideID),
		"nat_inside_id__n":          netboxapi.FilterInt64s(params.SetNatInsideIDn),
		"owner":                     netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                  netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":               netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":            netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":            netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":         netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"present_in_vrf":            netboxapi.FilterString(params.SetPresentInVrf),
		"present_in_vrf_id":         netboxapi.FilterString(params.SetPresentInVrfID),
		"q":                         netboxapi.FilterString(params.SetQ),
		"service_id":                netboxapi.FilterInt64s(params.SetServiceID),
		"service_id__n":             netboxapi.FilterInt64s(params.SetServiceIDn),
		"tag_id":                    netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                 netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":               netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant_group":              netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":           netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":           netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":        netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"virtual_machine":           netboxapi.FilterStrings(params.SetVirtualMachine),
		"virtual_machine_id":        netboxapi.FilterInt64s(params.SetVirtualMachineID),
		"vminterface":               netboxapi.FilterStrings(params.SetVminterface),
		"vminterface__n":            netboxapi.FilterStrings(params.SetVminterfacen),
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
