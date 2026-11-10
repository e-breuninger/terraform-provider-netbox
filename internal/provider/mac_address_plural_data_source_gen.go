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
	_ datasource.DataSource              = (*macAddressesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*macAddressesDataSource)(nil)
)

// NewMacAddressesDataSource returns a new mac_addresses data source, which
// lists mac_address objects matching its filters.
func NewMacAddressesDataSource() datasource.DataSource {
	return &macAddressesDataSource{}
}

type macAddressesDataSource struct {
	client *netboxapi.Client
}

// macAddressesDataSourceModel is the data source model: the filters, the limit and the matching
// mac_addresses.
type macAddressesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"mac_addresses"`
}

// macAddressResourceAttrTypes is the attribute type map of macAddressResourceModel.
func macAddressResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                           types.Int64Type,
		"mac_address":                  conv.MACAddressType{},
		"assigned_object_type":         types.StringType,
		"assigned_object_id":           types.Int64Type,
		"device_interface_id":          types.Int64Type,
		"virtual_machine_interface_id": types.Int64Type,
		"description":                  types.StringType,
		"comments":                     types.StringType,
		"owner_id":                     types.Int64Type,
		"last_updated":                 types.StringType,
		"url":                          types.StringType,
		"tags":                         types.SetType{ElemType: types.StringType},
		"tags_all":                     types.SetType{ElemType: types.StringType},
		"custom_fields":                types.MapType{ElemType: types.StringType},
	}
}

func (d *macAddressesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mac_addresses"
}

func (d *macAddressesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists mac_address objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: assigned, assigned_object_id, assigned_object_id__empty, assigned_object_id__gt, assigned_object_id__gte, assigned_object_id__lt, assigned_object_id__lte, assigned_object_id__n, assigned_object_type, assigned_object_type__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device_id, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface, interface__n, interface_id, interface_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, mac_address, mac_address__ic, mac_address__ie, mac_address__iew, mac_address__iregex, mac_address__isw, mac_address__n, mac_address__nic, mac_address__nie, mac_address__niew, mac_address__nisw, mac_address__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, primary, q, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, virtual_machine, virtual_machine_id, vminterface, vminterface__n, vminterface_id, vminterface_id__n. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"mac_addresses": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching mac_addresses.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"mac_address": schema.StringAttribute{
							CustomType:  conv.MACAddressType{},
							Computed:    true,
							Description: "The MAC address (EUI-48) in colon, dash or dot notation, any case; the configured spelling is kept in state (NetBox stores it colon-separated and uppercased).",
						},
						"assigned_object_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the interface the address is assigned to. Derived from device_interface_id or virtual_machine_interface_id when one of those is set. One of: dcim.interface, virtualization.vminterface.",
						},
						"assigned_object_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the interface the address is assigned to (see assigned_object_type).",
						},
						"device_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device interface the address is assigned to (assigned_object_type dcim.interface). Conflicts with virtual_machine_interface_id and with setting the assigned_object_* pair directly.",
						},
						"virtual_machine_interface_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual machine interface the address is assigned to (assigned_object_type virtualization.vminterface). Conflicts with device_interface_id and with setting the assigned_object_* pair directly.",
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

func (d *macAddressesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *macAddressesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data macAddressesDataSourceModel

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
	var responseDTOs []*netboxapi.MacAddressResponseDTO

	params := dcim.NewDcimMacAddressesListParams()
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
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
	var vDevice []string
	var vDeviceID []int64
	var vInterface []string
	var vInterfacen []string
	var vInterfaceID []int64
	var vInterfaceIDn []int64
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
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
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPrimary *bool
	var vQ *string
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vVirtualMachine []string
	var vVirtualMachineID []int64
	var vVminterface []string
	var vVminterfacen []string
	var vVminterfaceID []int64
	var vVminterfaceIDn []int64
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
		case "device":
			v := value
			vDevice = append(vDevice, v)
		case "device_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id' takes an integer, got %q.", value))
				return
			}
			vDeviceID = append(vDeviceID, v)
		case "interface":
			v := value
			vInterface = append(vInterface, v)
		case "interface__n":
			v := value
			vInterfacen = append(vInterfacen, v)
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
		case "primary":
			if vPrimary != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'primary' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'primary' takes a boolean, got %q.", value))
				return
			}
			vPrimary = &v
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
	if len(vDevice) > 0 {
		params.SetDevice(vDevice)
	}
	if len(vDeviceID) > 0 {
		params.SetDeviceID(vDeviceID)
	}
	if len(vInterface) > 0 {
		params.SetInterface(vInterface)
	}
	if len(vInterfacen) > 0 {
		params.SetInterfacen(vInterfacen)
	}
	if len(vInterfaceID) > 0 {
		params.SetInterfaceID(vInterfaceID)
	}
	if len(vInterfaceIDn) > 0 {
		params.SetInterfaceIDn(vInterfaceIDn)
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
	if vPrimary != nil {
		params.SetPrimary(vPrimary)
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
	if len(vVminterfaceID) > 0 {
		params.SetVminterfaceID(vVminterfaceID)
	}
	if len(vVminterfaceIDn) > 0 {
		params.SetVminterfaceIDn(vVminterfaceIDn)
	}
	// NetBox caps a page at its MAX_PAGE_SIZE, so the objects are read page by page.
	offset := int64(0)
	for int64(len(responseDTOs)) < limit {
		pageLimit := limit - int64(len(responseDTOs))
		params.SetLimit(&pageLimit)
		params.SetOffset(&offset)
		res, err := d.client.Dcim.DcimMacAddressesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_mac_addresses", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.MacAddressResponseDTOFromGoNetbox(goNetboxModel))
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
	resourceWithHooks := &macAddressResource{client: d.client}
	items := make([]macAddressResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model macAddressResourceModel
		resp.Diagnostics.Append(flattenMacAddress(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: macAddressResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
