// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                        netboxapi.FilterInt64s(params.SetID),
		"id__n":                     netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                    netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                   netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                    netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                   netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                 netboxapi.FilterBool(params.SetIDEmpty),
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
		"created":                   netboxapi.FilterStrings(params.SetCreated),
		"created__n":                netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":               netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":              netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":               netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":              netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":            netboxapi.FilterStrings(params.SetCreatedEmpty),
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
		"device":                    netboxapi.FilterStrings(params.SetDevice),
		"device_id":                 netboxapi.FilterInt64s(params.SetDeviceID),
		"interface":                 netboxapi.FilterStrings(params.SetInterface),
		"interface__n":              netboxapi.FilterStrings(params.SetInterfacen),
		"interface_id":              netboxapi.FilterInt64s(params.SetInterfaceID),
		"interface_id__n":           netboxapi.FilterInt64s(params.SetInterfaceIDn),
		"last_updated":              netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":           netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":          netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":         netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":          netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":         netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":       netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"mac_address":               netboxapi.FilterStrings(params.SetMacAddress),
		"mac_address__n":            netboxapi.FilterStrings(params.SetMacAddressn),
		"mac_address__ic":           netboxapi.FilterStrings(params.SetMacAddressIc),
		"mac_address__nic":          netboxapi.FilterStrings(params.SetMacAddressNic),
		"mac_address__ie":           netboxapi.FilterStrings(params.SetMacAddressIe),
		"mac_address__nie":          netboxapi.FilterStrings(params.SetMacAddressNie),
		"mac_address__isw":          netboxapi.FilterStrings(params.SetMacAddressIsw),
		"mac_address__nisw":         netboxapi.FilterStrings(params.SetMacAddressNisw),
		"mac_address__iew":          netboxapi.FilterStrings(params.SetMacAddressIew),
		"mac_address__niew":         netboxapi.FilterStrings(params.SetMacAddressNiew),
		"mac_address__regex":        netboxapi.FilterStrings(params.SetMacAddressRegex),
		"mac_address__iregex":       netboxapi.FilterStrings(params.SetMacAddressIregex),
		"owner":                     netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                  netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":               netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":            netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":            netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":         netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"primary":                   netboxapi.FilterBool(params.SetPrimary),
		"q":                         netboxapi.FilterString(params.SetQ),
		"tag_id":                    netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                 netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":               netboxapi.FilterInt64s(params.SetTagIDAny),
		"virtual_machine":           netboxapi.FilterStrings(params.SetVirtualMachine),
		"virtual_machine_id":        netboxapi.FilterInt64s(params.SetVirtualMachineID),
		"vminterface":               netboxapi.FilterStrings(params.SetVminterface),
		"vminterface__n":            netboxapi.FilterStrings(params.SetVminterfacen),
		"vminterface_id":            netboxapi.FilterInt64s(params.SetVminterfaceID),
		"vminterface_id__n":         netboxapi.FilterInt64s(params.SetVminterfaceIDn),
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
