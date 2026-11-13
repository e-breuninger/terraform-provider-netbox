// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

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
	_ datasource.DataSource              = (*servicesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*servicesDataSource)(nil)
)

// NewServicesDataSource returns a new services data source, which
// lists service objects matching its filters.
func NewServicesDataSource() datasource.DataSource {
	return &servicesDataSource{}
}

type servicesDataSource struct {
	client *netboxapi.Client
}

// servicesDataSourceModel is the data source model: the filters, the limit and the matching
// services.
type servicesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"services"`
}

// serviceResourceAttrTypes is the attribute type map of serviceResourceModel.
func serviceResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":                 types.Int64Type,
		"name":               types.StringType,
		"protocol":           types.StringType,
		"ports":              types.SetType{ElemType: types.Int64Type},
		"parent_object_type": types.StringType,
		"parent_object_id":   types.Int64Type,
		"device_id":          types.Int64Type,
		"virtual_machine_id": types.Int64Type,
		"ip_address_ids":     types.SetType{ElemType: types.Int64Type},
		"description":        types.StringType,
		"comments":           types.StringType,
		"owner_id":           types.Int64Type,
		"created":            types.StringType,
		"last_updated":       types.StringType,
		"url":                types.StringType,
		"tags":               types.SetType{ElemType: types.StringType},
		"tags_all":           types.SetType{ElemType: types.StringType},
		"custom_fields":      types.MapType{ElemType: types.StringType},
	}
}

func (d *servicesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_services"
}

func (d *servicesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):Lists service objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: contact, contact__n, contact_group, contact_group__n, contact_role, contact_role__n, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device_id, fhrpgroup, fhrpgroup_id, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, ip_address, ip_address__n, ip_address_id, ip_address_id__n, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, name, name__empty, name__ic, name__ie, name__iew, name__iregex, name__isw, name__n, name__nic, name__nie, name__niew, name__nisw, name__regex, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, parent_object_id, parent_object_id__empty, parent_object_id__gt, parent_object_id__gte, parent_object_id__lt, parent_object_id__lte, parent_object_id__n, parent_object_type, parent_object_type__n, port, port__empty, port__gt, port__gte, port__lt, port__lte, port__n, protocol, protocol__empty, protocol__ic, protocol__ie, protocol__iew, protocol__iregex, protocol__isw, protocol__n, protocol__nic, protocol__nie, protocol__niew, protocol__nisw, protocol__regex, q, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, virtual_machine, virtual_machine_id. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"services": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching services.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"name": schema.StringAttribute{
							Computed: true,
						},
						"protocol": schema.StringAttribute{
							Computed:    true,
							Description: "Transport protocol. One of: tcp, udp, sctp.",
						},
						"ports": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Port numbers the service listens on.",
						},
						"parent_object_type": schema.StringAttribute{
							Computed:    true,
							Description: "Content type of the object the service runs on. Derived from device_id or virtual_machine_id when one of those is set; set it together with parent_object_id to address the object directly. One of: dcim.device, virtualization.virtualmachine.",
						},
						"parent_object_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the object named by parent_object_type.",
						},
						"device_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the device the service runs on (parent_object_type dcim.device). Conflicts with virtual_machine_id and with setting the parent_object_* pair directly.",
						},
						"virtual_machine_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the virtual machine the service runs on (parent_object_type virtualization.virtualmachine). Conflicts with device_id and with setting the parent_object_* pair directly.",
						},
						"ip_address_ids": schema.SetAttribute{
							ElementType: types.Int64Type,
							Computed:    true,
							Description: "Ids of the IP addresses the service is bound to; all of the parent object's addresses when unset.",
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
							Description: "Custom field values by field name. Every value is a string; NetBox coerces numbers and booleans. An object field holds the related object's id, e.g. \"12\", a multiobject field a JSON list of ids in any order, e.g. jsonencode([12, 13]). A key removed from the map is cleared in NetBox (set {} to clear all).",
						},
					},
				},
			},
		},
	}
}

func (d *servicesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *servicesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data servicesDataSourceModel

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
	var responseDTOs []*netboxapi.ServiceResponseDTO

	params := ipam.NewIpamServicesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                      netboxapi.FilterInt64s(params.SetID),
		"id__n":                   netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                  netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":                 netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                  netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":                 netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":               netboxapi.FilterBool(params.SetIDEmpty),
		"name":                    netboxapi.FilterStrings(params.SetName),
		"name__n":                 netboxapi.FilterStrings(params.SetNamen),
		"name__ic":                netboxapi.FilterStrings(params.SetNameIc),
		"name__nic":               netboxapi.FilterStrings(params.SetNameNic),
		"name__ie":                netboxapi.FilterStrings(params.SetNameIe),
		"name__nie":               netboxapi.FilterStrings(params.SetNameNie),
		"name__isw":               netboxapi.FilterStrings(params.SetNameIsw),
		"name__nisw":              netboxapi.FilterStrings(params.SetNameNisw),
		"name__iew":               netboxapi.FilterStrings(params.SetNameIew),
		"name__niew":              netboxapi.FilterStrings(params.SetNameNiew),
		"name__empty":             netboxapi.FilterBool(params.SetNameEmpty),
		"name__regex":             netboxapi.FilterStrings(params.SetNameRegex),
		"name__iregex":            netboxapi.FilterStrings(params.SetNameIregex),
		"protocol":                netboxapi.FilterString(params.SetProtocol),
		"protocol__n":             netboxapi.FilterString(params.SetProtocoln),
		"protocol__ic":            netboxapi.FilterStrings(params.SetProtocolIc),
		"protocol__nic":           netboxapi.FilterStrings(params.SetProtocolNic),
		"protocol__ie":            netboxapi.FilterStrings(params.SetProtocolIe),
		"protocol__nie":           netboxapi.FilterStrings(params.SetProtocolNie),
		"protocol__isw":           netboxapi.FilterStrings(params.SetProtocolIsw),
		"protocol__nisw":          netboxapi.FilterStrings(params.SetProtocolNisw),
		"protocol__iew":           netboxapi.FilterStrings(params.SetProtocolIew),
		"protocol__niew":          netboxapi.FilterStrings(params.SetProtocolNiew),
		"protocol__empty":         netboxapi.FilterBool(params.SetProtocolEmpty),
		"protocol__regex":         netboxapi.FilterStrings(params.SetProtocolRegex),
		"protocol__iregex":        netboxapi.FilterStrings(params.SetProtocolIregex),
		"device_id":               netboxapi.FilterInt64s(params.SetDeviceID),
		"virtual_machine_id":      netboxapi.FilterInt64s(params.SetVirtualMachineID),
		"description":             netboxapi.FilterStrings(params.SetDescription),
		"description__n":          netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":         netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":        netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":         netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":        netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":        netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":       netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":        netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":       netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":      netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":      netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":     netboxapi.FilterStrings(params.SetDescriptionIregex),
		"owner_id":                netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":             netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                     netboxapi.FilterStrings(params.SetTag),
		"tag__n":                  netboxapi.FilterStrings(params.SetTagn),
		"tag__any":                netboxapi.FilterStrings(params.SetTagAny),
		"contact":                 netboxapi.FilterInt64s(params.SetContact),
		"contact__n":              netboxapi.FilterInt64s(params.SetContactn),
		"contact_group":           netboxapi.FilterStrings(params.SetContactGroup),
		"contact_group__n":        netboxapi.FilterStrings(params.SetContactGroupn),
		"contact_role":            netboxapi.FilterInt64s(params.SetContactRole),
		"contact_role__n":         netboxapi.FilterInt64s(params.SetContactRolen),
		"created":                 netboxapi.FilterStrings(params.SetCreated),
		"created__n":              netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":             netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":            netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":             netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":            netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":          netboxapi.FilterStrings(params.SetCreatedEmpty),
		"device":                  netboxapi.FilterStrings(params.SetDevice),
		"fhrpgroup":               netboxapi.FilterStrings(params.SetFhrpgroup),
		"fhrpgroup_id":            netboxapi.FilterInt64s(params.SetFhrpgroupID),
		"ip_address":              netboxapi.FilterStrings(params.SetIPAddress),
		"ip_address__n":           netboxapi.FilterStrings(params.SetIPAddressn),
		"ip_address_id":           netboxapi.FilterInt64s(params.SetIPAddressID),
		"ip_address_id__n":        netboxapi.FilterInt64s(params.SetIPAddressIDn),
		"last_updated":            netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":         netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":        netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":       netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":        netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":       netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":     netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"owner":                   netboxapi.FilterStrings(params.SetOwner),
		"owner__n":                netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":             netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":          netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":          netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":       netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"parent_object_id":        netboxapi.FilterInt64s(params.SetParentObjectID),
		"parent_object_id__n":     netboxapi.FilterInt64s(params.SetParentObjectIDn),
		"parent_object_id__lt":    netboxapi.FilterInt64s(params.SetParentObjectIDLt),
		"parent_object_id__lte":   netboxapi.FilterInt64s(params.SetParentObjectIDLte),
		"parent_object_id__gt":    netboxapi.FilterInt64s(params.SetParentObjectIDGt),
		"parent_object_id__gte":   netboxapi.FilterInt64s(params.SetParentObjectIDGte),
		"parent_object_id__empty": netboxapi.FilterBool(params.SetParentObjectIDEmpty),
		"parent_object_type":      netboxapi.FilterStrings(params.SetParentObjectType),
		"parent_object_type__n":   netboxapi.FilterStrings(params.SetParentObjectTypen),
		"port":                    netboxapi.FilterFloat64(params.SetPort),
		"port__n":                 netboxapi.FilterFloat64(params.SetPortn),
		"port__lt":                netboxapi.FilterFloat64(params.SetPortLt),
		"port__lte":               netboxapi.FilterFloat64(params.SetPortLte),
		"port__gt":                netboxapi.FilterFloat64(params.SetPortGt),
		"port__gte":               netboxapi.FilterFloat64(params.SetPortGte),
		"port__empty":             netboxapi.FilterFloat64(params.SetPortEmpty),
		"q":                       netboxapi.FilterString(params.SetQ),
		"tag_id":                  netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":               netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":             netboxapi.FilterInt64s(params.SetTagIDAny),
		"virtual_machine":         netboxapi.FilterStrings(params.SetVirtualMachine),
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
		res, err := d.client.Ipam.IpamServicesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_services", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.ServiceResponseDTOFromGoNetbox(goNetboxModel))
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
	resourceWithHooks := &serviceResource{client: d.client}
	items := make([]serviceResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model serviceResourceModel
		resp.Diagnostics.Append(flattenService(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: serviceResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
