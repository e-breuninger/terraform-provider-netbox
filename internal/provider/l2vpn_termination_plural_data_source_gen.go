// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
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
	listFilters := map[string]netboxapi.ListFilter{
		"id":                         netboxapi.FilterInt64s(params.SetID),
		"id__n":                      netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                     netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                    netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                     netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                    netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":                  netboxapi.FilterBool(params.SetIDEmpty),
		"l2vpn_id":                   netboxapi.FilterInt64s(params.SetL2vpnID),
		"l2vpn_id__n":                netboxapi.FilterInt64s(params.SetL2vpnIDn),
		"vlan_id":                    netboxapi.FilterInt64s(params.SetVlanID),
		"vlan_id__n":                 netboxapi.FilterInt64s(params.SetVlanIDn),
		"interface_id":               netboxapi.FilterInt64s(params.SetInterfaceID),
		"interface_id__n":            netboxapi.FilterInt64s(params.SetInterfaceIDn),
		"vminterface_id":             netboxapi.FilterInt64s(params.SetVminterfaceID),
		"vminterface_id__n":          netboxapi.FilterInt64s(params.SetVminterfaceIDn),
		"tag":                        netboxapi.FilterStrings(params.SetTag),
		"tag__n":                     netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                   netboxapi.FilterStrings(params.SetTagAny),
		"assigned_object_id":         netboxapi.FilterInt64s(params.SetAssignedObjectID),
		"assigned_object_id__n":      netboxapi.FilterInt64s(params.SetAssignedObjectIDn),
		"assigned_object_id__lt":     netboxapi.FilterInt64s(params.SetAssignedObjectIDLt),
		"assigned_object_id__lte":    netboxapi.FilterInt64s(params.SetAssignedObjectIDLte),
		"assigned_object_id__gt":     netboxapi.FilterInt64s(params.SetAssignedObjectIDGt),
		"assigned_object_id__gte":    netboxapi.FilterInt64s(params.SetAssignedObjectIDGte),
		"assigned_object_id__empty":  netboxapi.FilterBool(params.SetAssignedObjectIDEmpty),
		"assigned_object_type":       netboxapi.FilterStrings(params.SetAssignedObjectType),
		"assigned_object_type__n":    netboxapi.FilterStrings(params.SetAssignedObjectTypen),
		"assigned_object_type_id":    netboxapi.FilterInt64s(params.SetAssignedObjectTypeID),
		"assigned_object_type_id__n": netboxapi.FilterInt64s(params.SetAssignedObjectTypeIDn),
		"created":                    netboxapi.FilterStrings(params.SetCreated),
		"created__n":                 netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":                netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":               netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":                netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":               netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":             netboxapi.FilterStrings(params.SetCreatedEmpty),
		"device":                     netboxapi.FilterStrings(params.SetDevice),
		"device__n":                  netboxapi.FilterStrings(params.SetDevicen),
		"device_id":                  netboxapi.FilterInt64s(params.SetDeviceID),
		"device_id__n":               netboxapi.FilterInt64s(params.SetDeviceIDn),
		"interface":                  netboxapi.FilterStrings(params.SetInterface),
		"interface__n":               netboxapi.FilterStrings(params.SetInterfacen),
		"l2vpn":                      netboxapi.FilterStrings(params.SetL2vpn),
		"l2vpn__n":                   netboxapi.FilterStrings(params.SetL2vpnn),
		"last_updated":               netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":            netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":           netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":          netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":           netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":          netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":        netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"q":                          netboxapi.FilterString(params.SetQ),
		"region":                     netboxapi.FilterStrings(params.SetRegion),
		"region_id":                  netboxapi.FilterInt64s(params.SetRegionID),
		"site":                       netboxapi.FilterStrings(params.SetSite),
		"site_id":                    netboxapi.FilterInt64s(params.SetSiteID),
		"tag_id":                     netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":                  netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":                netboxapi.FilterInt64s(params.SetTagIDAny),
		"virtual_machine":            netboxapi.FilterStrings(params.SetVirtualMachine),
		"virtual_machine__n":         netboxapi.FilterStrings(params.SetVirtualMachinen),
		"virtual_machine_id":         netboxapi.FilterInt64s(params.SetVirtualMachineID),
		"virtual_machine_id__n":      netboxapi.FilterInt64s(params.SetVirtualMachineIDn),
		"vlan":                       netboxapi.FilterStrings(params.SetVlan),
		"vlan__n":                    netboxapi.FilterStrings(params.SetVlann),
		"vlan_vid":                   netboxapi.FilterInt64(params.SetVlanVid),
		"vlan_vid__n":                netboxapi.FilterInt64(params.SetVlanVidn),
		"vlan_vid__lt":               netboxapi.FilterInt64(params.SetVlanVidLt),
		"vlan_vid__lte":              netboxapi.FilterInt64(params.SetVlanVidLte),
		"vlan_vid__gt":               netboxapi.FilterInt64(params.SetVlanVidGt),
		"vlan_vid__gte":              netboxapi.FilterInt64(params.SetVlanVidGte),
		"vlan_vid__empty":            netboxapi.FilterInt64(params.SetVlanVidEmpty),
		"vminterface":                netboxapi.FilterStrings(params.SetVminterface),
		"vminterface__n":             netboxapi.FilterStrings(params.SetVminterfacen),
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
