// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

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
	_ datasource.DataSource              = (*cablesDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*cablesDataSource)(nil)
)

// NewCablesDataSource returns a new cables data source, which
// lists cable objects matching its filters.
func NewCablesDataSource() datasource.DataSource {
	return &cablesDataSource{}
}

type cablesDataSource struct {
	client *netboxapi.Client
}

// cablesDataSourceModel is the data source model: the filters, the limit and the matching
// cables.
type cablesDataSourceModel struct {
	Filters types.Set   `tfsdk:"filters"`
	Limit   types.Int64 `tfsdk:"limit"`
	Items   types.List  `tfsdk:"cables"`
}

// cableResourceAttrTypes is the attribute type map of cableResourceModel.
func cableResourceAttrTypes() map[string]attr.Type {
	return map[string]attr.Type{
		"id":            types.Int64Type,
		"a_side":        types.ObjectType{AttrTypes: cableASideAttrTypes()},
		"b_side":        types.ObjectType{AttrTypes: cableBSideAttrTypes()},
		"type":          types.StringType,
		"status":        types.StringType,
		"profile":       types.StringType,
		"tenant_id":     types.Int64Type,
		"bundle_id":     types.Int64Type,
		"label":         types.StringType,
		"color_hex":     types.StringType,
		"length":        types.Float64Type,
		"length_unit":   types.StringType,
		"description":   types.StringType,
		"comments":      types.StringType,
		"owner_id":      types.Int64Type,
		"created":       types.StringType,
		"last_updated":  types.StringType,
		"url":           types.StringType,
		"tags":          types.SetType{ElemType: types.StringType},
		"tags_all":      types.SetType{ElemType: types.StringType},
		"custom_fields": types.MapType{ElemType: types.StringType},
	}
}

func (d *cablesDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cables"
}

func (d *cablesDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):Lists cable objects matching the given filters.",
		Attributes: map[string]schema.Attribute{
			"filters": schema.SetNestedAttribute{
				Optional:    true,
				Description: "Query filters, sent as API list parameters. Supported names: bundle, bundle__n, bundle_id, bundle_id__n, circuittermination_id, color, color__empty, color__ic, color__ie, color__iew, color__iregex, color__isw, color__n, color__nic, color__nie, color__niew, color__nisw, color__regex, consoleport_id, consoleserverport_id, created, created__empty, created__gt, created__gte, created__lt, created__lte, created__n, description, description__empty, description__ic, description__ie, description__iew, description__iregex, description__isw, description__n, description__nic, description__nie, description__niew, description__nisw, description__regex, device, device_id, frontport_id, id, id__empty, id__gt, id__gte, id__lt, id__lte, id__n, interface_id, label, label__empty, label__ic, label__ie, label__iew, label__iregex, label__isw, label__n, label__nic, label__nie, label__niew, label__nisw, label__regex, last_updated, last_updated__empty, last_updated__gt, last_updated__gte, last_updated__lt, last_updated__lte, last_updated__n, length, length__empty, length__gt, length__gte, length__lt, length__lte, length__n, length_unit, length_unit__empty, length_unit__ic, length_unit__ie, length_unit__iew, length_unit__iregex, length_unit__isw, length_unit__n, length_unit__nic, length_unit__nie, length_unit__niew, length_unit__nisw, length_unit__regex, location, location_id, owner, owner__n, owner_group, owner_group__n, owner_group_id, owner_group_id__n, owner_id, owner_id__n, powerfeed_id, poweroutlet_id, powerport_id, profile, profile__empty, profile__ic, profile__ie, profile__iew, profile__iregex, profile__isw, profile__n, profile__nic, profile__nie, profile__niew, profile__nisw, profile__regex, q, rack, rack_id, rearport_id, site, site_id, status, status__empty, status__ic, status__ie, status__iew, status__iregex, status__isw, status__n, status__nic, status__nie, status__niew, status__nisw, status__regex, tag, tag__any, tag__n, tag_id, tag_id__any, tag_id__n, tenant, tenant__n, tenant_group, tenant_group__n, tenant_group_id, tenant_group_id__n, tenant_id, tenant_id__n, termination_a_id, termination_a_type, termination_a_type__n, termination_b_id, termination_b_type, termination_b_type__n, type, type__empty, type__ic, type__ie, type__iew, type__iregex, type__isw, type__n, type__nic, type__nie, type__niew, type__nisw, type__regex, unterminated. Repeating a name sends that parameter once per value. Custom fields filter as cf_<field name>, e.g. cf_tier, with the field's own filter logic: loose is a case-insensitive substring match, exact an exact match. NetBox ignores names it has no custom field for, and such an entry does not narrow the result.",
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
			"cables": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The matching cables.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"id": schema.Int64Attribute{
							Computed:    true,
							Description: "NetBox id.",
						},
						"a_side": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"object_type": schema.StringAttribute{
									Computed:    true,
									Description: "Content type of the terminating objects. Derived from the *_ids list that is set; set it together with ids otherwise. One of: dcim.interface, dcim.frontport, dcim.rearport, dcim.consoleport, dcim.consoleserverport, dcim.powerport, dcim.poweroutlet, dcim.powerfeed, circuits.circuittermination.",
								},
								"ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the terminating objects, in order (see object_type).",
								},
								"device_interface_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the device interfaces this side terminates on (object_type dcim.interface). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"front_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the front ports this side terminates on (object_type dcim.frontport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"rear_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the rear ports this side terminates on (object_type dcim.rearport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"console_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the console ports this side terminates on (object_type dcim.consoleport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"console_server_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the console server ports this side terminates on (object_type dcim.consoleserverport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"power_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the power ports this side terminates on (object_type dcim.powerport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"power_outlet_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the power outlets this side terminates on (object_type dcim.poweroutlet). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"power_feed_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the power feeds this side terminates on (object_type dcim.powerfeed). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"circuit_termination_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the circuit terminations this side terminates on (object_type circuits.circuittermination). Conflicts with the other *_ids lists and with object_type/ids.",
								},
							},
							Computed:    true,
							Description: "What the A side terminates on: objects of one type (NetBox's rule), one or several for a breakout, in order (with a profile the position is the connector). Give object_type and ids, or exactly one of the typed *_ids lists, which is an alias of the same thing.",
						},
						"b_side": schema.SingleNestedAttribute{
							Attributes: map[string]schema.Attribute{
								"object_type": schema.StringAttribute{
									Computed:    true,
									Description: "Content type of the terminating objects. Derived from the *_ids list that is set; set it together with ids otherwise. One of: dcim.interface, dcim.frontport, dcim.rearport, dcim.consoleport, dcim.consoleserverport, dcim.powerport, dcim.poweroutlet, dcim.powerfeed, circuits.circuittermination.",
								},
								"ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the terminating objects, in order (see object_type).",
								},
								"device_interface_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the device interfaces this side terminates on (object_type dcim.interface). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"front_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the front ports this side terminates on (object_type dcim.frontport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"rear_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the rear ports this side terminates on (object_type dcim.rearport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"console_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the console ports this side terminates on (object_type dcim.consoleport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"console_server_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the console server ports this side terminates on (object_type dcim.consoleserverport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"power_port_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the power ports this side terminates on (object_type dcim.powerport). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"power_outlet_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the power outlets this side terminates on (object_type dcim.poweroutlet). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"power_feed_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the power feeds this side terminates on (object_type dcim.powerfeed). Conflicts with the other *_ids lists and with object_type/ids.",
								},
								"circuit_termination_ids": schema.ListAttribute{
									ElementType: types.Int64Type,
									Computed:    true,
									Description: "Ids of the circuit terminations this side terminates on (object_type circuits.circuittermination). Conflicts with the other *_ids lists and with object_type/ids.",
								},
							},
							Computed:    true,
							Description: "What the B side terminates on: objects of one type (NetBox's rule), one or several for a breakout, in order (with a profile the position is the connector). Give object_type and ids, or exactly one of the typed *_ids lists, which is an alias of the same thing.",
						},
						"type": schema.StringAttribute{
							Computed:    true,
							Description: "Cable type as its NetBox slug, e.g. cat6, smf-os2, power (any of NetBox's cable type choices).",
						},
						"status": schema.StringAttribute{
							Computed:    true,
							Description: "Cable status. One of: connected, planned, decommissioning.",
						},
						"profile": schema.StringAttribute{
							Computed:    true,
							Description: "Cable profile (connectors and positions per side); with one set, each termination's position in the list is its connector. NetBox rejects clearing it, so unset keeps the current value. One of: single-1c1p, single-1c2p, single-1c4p, single-1c6p, single-1c8p, single-1c12p, single-1c16p, trunk-2c1p, trunk-2c2p, trunk-2c4p, trunk-2c4p-shuffle, trunk-2c6p, trunk-2c8p, trunk-2c12p, trunk-4c1p, trunk-4c2p, trunk-4c4p, trunk-4c4p-shuffle, trunk-4c6p, trunk-4c8p, trunk-8c4p, breakout-1c2p-2c1p, breakout-1c4p-4c1p, breakout-1c6p-6c1p, breakout-1c8p-8c1p, breakout-2c4p-8c1p-shuffle.",
						},
						"tenant_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the tenant.",
						},
						"bundle_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the cable bundle.",
						},
						"label": schema.StringAttribute{
							Computed:    true,
							Description: "Physical label.",
						},
						"color_hex": schema.StringAttribute{
							Computed:    true,
							Description: "RGB color in hex (e.g. ff0000).",
						},
						"length": schema.Float64Attribute{
							Computed:    true,
							Description: "Cable length; requires length_unit.",
						},
						"length_unit": schema.StringAttribute{
							Computed:    true,
							Description: "Unit of length. One of: km, m, cm, mi, ft, in.",
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

func (d *cablesDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *cablesDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data cablesDataSourceModel

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
	var responseDTOs []*netboxapi.CableResponseDTO

	params := dcim.NewDcimCablesListParams()
	listFilters := map[string]netboxapi.ListFilter{
		"id":                    netboxapi.FilterInt64s(params.SetID),
		"id__n":                 netboxapi.FilterInt64s(params.SetIDn),
		"id__lt":                netboxapi.FilterInt64s(params.SetIDLt),
		"id__lte":               netboxapi.FilterInt64s(params.SetIDLte),
		"id__gt":                netboxapi.FilterInt64s(params.SetIDGt),
		"id__gte":               netboxapi.FilterInt64s(params.SetIDGte),
		"id__empty":             netboxapi.FilterBool(params.SetIDEmpty),
		"label":                 netboxapi.FilterStrings(params.SetLabel),
		"label__n":              netboxapi.FilterStrings(params.SetLabeln),
		"label__ic":             netboxapi.FilterStrings(params.SetLabelIc),
		"label__nic":            netboxapi.FilterStrings(params.SetLabelNic),
		"label__ie":             netboxapi.FilterStrings(params.SetLabelIe),
		"label__nie":            netboxapi.FilterStrings(params.SetLabelNie),
		"label__isw":            netboxapi.FilterStrings(params.SetLabelIsw),
		"label__nisw":           netboxapi.FilterStrings(params.SetLabelNisw),
		"label__iew":            netboxapi.FilterStrings(params.SetLabelIew),
		"label__niew":           netboxapi.FilterStrings(params.SetLabelNiew),
		"label__empty":          netboxapi.FilterBool(params.SetLabelEmpty),
		"label__regex":          netboxapi.FilterStrings(params.SetLabelRegex),
		"label__iregex":         netboxapi.FilterStrings(params.SetLabelIregex),
		"status":                netboxapi.FilterStrings(params.SetStatus),
		"status__n":             netboxapi.FilterStrings(params.SetStatusn),
		"status__ic":            netboxapi.FilterStrings(params.SetStatusIc),
		"status__nic":           netboxapi.FilterStrings(params.SetStatusNic),
		"status__ie":            netboxapi.FilterStrings(params.SetStatusIe),
		"status__nie":           netboxapi.FilterStrings(params.SetStatusNie),
		"status__isw":           netboxapi.FilterStrings(params.SetStatusIsw),
		"status__nisw":          netboxapi.FilterStrings(params.SetStatusNisw),
		"status__iew":           netboxapi.FilterStrings(params.SetStatusIew),
		"status__niew":          netboxapi.FilterStrings(params.SetStatusNiew),
		"status__empty":         netboxapi.FilterBool(params.SetStatusEmpty),
		"status__regex":         netboxapi.FilterStrings(params.SetStatusRegex),
		"status__iregex":        netboxapi.FilterStrings(params.SetStatusIregex),
		"type":                  netboxapi.FilterStrings(params.SetType),
		"type__n":               netboxapi.FilterStrings(params.SetTypen),
		"type__ic":              netboxapi.FilterStrings(params.SetTypeIc),
		"type__nic":             netboxapi.FilterStrings(params.SetTypeNic),
		"type__ie":              netboxapi.FilterStrings(params.SetTypeIe),
		"type__nie":             netboxapi.FilterStrings(params.SetTypeNie),
		"type__isw":             netboxapi.FilterStrings(params.SetTypeIsw),
		"type__nisw":            netboxapi.FilterStrings(params.SetTypeNisw),
		"type__iew":             netboxapi.FilterStrings(params.SetTypeIew),
		"type__niew":            netboxapi.FilterStrings(params.SetTypeNiew),
		"type__empty":           netboxapi.FilterBool(params.SetTypeEmpty),
		"type__regex":           netboxapi.FilterStrings(params.SetTypeRegex),
		"type__iregex":          netboxapi.FilterStrings(params.SetTypeIregex),
		"tenant_id":             netboxapi.FilterInt64s(params.SetTenantID),
		"tenant_id__n":          netboxapi.FilterInt64s(params.SetTenantIDn),
		"owner_id":              netboxapi.FilterInt64s(params.SetOwnerID),
		"owner_id__n":           netboxapi.FilterInt64s(params.SetOwnerIDn),
		"tag":                   netboxapi.FilterStrings(params.SetTag),
		"tag__n":                netboxapi.FilterStrings(params.SetTagn),
		"tag__any":              netboxapi.FilterStrings(params.SetTagAny),
		"site_id":               netboxapi.FilterInt64s(params.SetSiteID),
		"site":                  netboxapi.FilterStrings(params.SetSite),
		"location_id":           netboxapi.FilterInt64s(params.SetLocationID),
		"rack_id":               netboxapi.FilterInt64s(params.SetRackID),
		"device_id":             netboxapi.FilterInt64s(params.SetDeviceID),
		"device":                netboxapi.FilterStrings(params.SetDevice),
		"bundle":                netboxapi.FilterStrings(params.SetBundle),
		"bundle__n":             netboxapi.FilterStrings(params.SetBundlen),
		"bundle_id":             netboxapi.FilterInt64s(params.SetBundleID),
		"bundle_id__n":          netboxapi.FilterInt64s(params.SetBundleIDn),
		"circuittermination_id": netboxapi.FilterInt64s(params.SetCircuitterminationID),
		"color":                 netboxapi.FilterStrings(params.SetColor),
		"color__n":              netboxapi.FilterStrings(params.SetColorn),
		"color__ic":             netboxapi.FilterStrings(params.SetColorIc),
		"color__nic":            netboxapi.FilterStrings(params.SetColorNic),
		"color__ie":             netboxapi.FilterStrings(params.SetColorIe),
		"color__nie":            netboxapi.FilterStrings(params.SetColorNie),
		"color__isw":            netboxapi.FilterStrings(params.SetColorIsw),
		"color__nisw":           netboxapi.FilterStrings(params.SetColorNisw),
		"color__iew":            netboxapi.FilterStrings(params.SetColorIew),
		"color__niew":           netboxapi.FilterStrings(params.SetColorNiew),
		"color__empty":          netboxapi.FilterBool(params.SetColorEmpty),
		"color__regex":          netboxapi.FilterStrings(params.SetColorRegex),
		"color__iregex":         netboxapi.FilterStrings(params.SetColorIregex),
		"consoleport_id":        netboxapi.FilterInt64s(params.SetConsoleportID),
		"consoleserverport_id":  netboxapi.FilterInt64s(params.SetConsoleserverportID),
		"created":               netboxapi.FilterStrings(params.SetCreated),
		"created__n":            netboxapi.FilterStrings(params.SetCreatedn),
		"created__lt":           netboxapi.FilterStrings(params.SetCreatedLt),
		"created__lte":          netboxapi.FilterStrings(params.SetCreatedLte),
		"created__gt":           netboxapi.FilterStrings(params.SetCreatedGt),
		"created__gte":          netboxapi.FilterStrings(params.SetCreatedGte),
		"created__empty":        netboxapi.FilterStrings(params.SetCreatedEmpty),
		"description":           netboxapi.FilterStrings(params.SetDescription),
		"description__n":        netboxapi.FilterStrings(params.SetDescriptionn),
		"description__ic":       netboxapi.FilterStrings(params.SetDescriptionIc),
		"description__nic":      netboxapi.FilterStrings(params.SetDescriptionNic),
		"description__ie":       netboxapi.FilterStrings(params.SetDescriptionIe),
		"description__nie":      netboxapi.FilterStrings(params.SetDescriptionNie),
		"description__isw":      netboxapi.FilterStrings(params.SetDescriptionIsw),
		"description__nisw":     netboxapi.FilterStrings(params.SetDescriptionNisw),
		"description__iew":      netboxapi.FilterStrings(params.SetDescriptionIew),
		"description__niew":     netboxapi.FilterStrings(params.SetDescriptionNiew),
		"description__empty":    netboxapi.FilterBool(params.SetDescriptionEmpty),
		"description__regex":    netboxapi.FilterStrings(params.SetDescriptionRegex),
		"description__iregex":   netboxapi.FilterStrings(params.SetDescriptionIregex),
		"frontport_id":          netboxapi.FilterInt64s(params.SetFrontportID),
		"interface_id":          netboxapi.FilterInt64s(params.SetInterfaceID),
		"last_updated":          netboxapi.FilterStrings(params.SetLastUpdated),
		"last_updated__n":       netboxapi.FilterStrings(params.SetLastUpdatedn),
		"last_updated__lt":      netboxapi.FilterStrings(params.SetLastUpdatedLt),
		"last_updated__lte":     netboxapi.FilterStrings(params.SetLastUpdatedLte),
		"last_updated__gt":      netboxapi.FilterStrings(params.SetLastUpdatedGt),
		"last_updated__gte":     netboxapi.FilterStrings(params.SetLastUpdatedGte),
		"last_updated__empty":   netboxapi.FilterStrings(params.SetLastUpdatedEmpty),
		"length":                netboxapi.FilterFloat64s(params.SetLength),
		"length__n":             netboxapi.FilterFloat64s(params.SetLengthn),
		"length__lt":            netboxapi.FilterFloat64s(params.SetLengthLt),
		"length__lte":           netboxapi.FilterFloat64s(params.SetLengthLte),
		"length__gt":            netboxapi.FilterFloat64s(params.SetLengthGt),
		"length__gte":           netboxapi.FilterFloat64s(params.SetLengthGte),
		"length__empty":         netboxapi.FilterBool(params.SetLengthEmpty),
		"length_unit":           netboxapi.FilterString(params.SetLengthUnit),
		"length_unit__n":        netboxapi.FilterString(params.SetLengthUnitn),
		"length_unit__ic":       netboxapi.FilterStrings(params.SetLengthUnitIc),
		"length_unit__nic":      netboxapi.FilterStrings(params.SetLengthUnitNic),
		"length_unit__ie":       netboxapi.FilterStrings(params.SetLengthUnitIe),
		"length_unit__nie":      netboxapi.FilterStrings(params.SetLengthUnitNie),
		"length_unit__isw":      netboxapi.FilterStrings(params.SetLengthUnitIsw),
		"length_unit__nisw":     netboxapi.FilterStrings(params.SetLengthUnitNisw),
		"length_unit__iew":      netboxapi.FilterStrings(params.SetLengthUnitIew),
		"length_unit__niew":     netboxapi.FilterStrings(params.SetLengthUnitNiew),
		"length_unit__empty":    netboxapi.FilterBool(params.SetLengthUnitEmpty),
		"length_unit__regex":    netboxapi.FilterStrings(params.SetLengthUnitRegex),
		"length_unit__iregex":   netboxapi.FilterStrings(params.SetLengthUnitIregex),
		"location":              netboxapi.FilterStrings(params.SetLocation),
		"owner":                 netboxapi.FilterStrings(params.SetOwner),
		"owner__n":              netboxapi.FilterStrings(params.SetOwnern),
		"owner_group":           netboxapi.FilterStrings(params.SetOwnerGroup),
		"owner_group__n":        netboxapi.FilterStrings(params.SetOwnerGroupn),
		"owner_group_id":        netboxapi.FilterInt64s(params.SetOwnerGroupID),
		"owner_group_id__n":     netboxapi.FilterInt64s(params.SetOwnerGroupIDn),
		"powerfeed_id":          netboxapi.FilterInt64s(params.SetPowerfeedID),
		"poweroutlet_id":        netboxapi.FilterInt64s(params.SetPoweroutletID),
		"powerport_id":          netboxapi.FilterInt64s(params.SetPowerportID),
		"profile":               netboxapi.FilterStrings(params.SetProfile),
		"profile__n":            netboxapi.FilterStrings(params.SetProfilen),
		"profile__ic":           netboxapi.FilterStrings(params.SetProfileIc),
		"profile__nic":          netboxapi.FilterStrings(params.SetProfileNic),
		"profile__ie":           netboxapi.FilterStrings(params.SetProfileIe),
		"profile__nie":          netboxapi.FilterStrings(params.SetProfileNie),
		"profile__isw":          netboxapi.FilterStrings(params.SetProfileIsw),
		"profile__nisw":         netboxapi.FilterStrings(params.SetProfileNisw),
		"profile__iew":          netboxapi.FilterStrings(params.SetProfileIew),
		"profile__niew":         netboxapi.FilterStrings(params.SetProfileNiew),
		"profile__empty":        netboxapi.FilterBool(params.SetProfileEmpty),
		"profile__regex":        netboxapi.FilterStrings(params.SetProfileRegex),
		"profile__iregex":       netboxapi.FilterStrings(params.SetProfileIregex),
		"q":                     netboxapi.FilterString(params.SetQ),
		"rack":                  netboxapi.FilterStrings(params.SetRack),
		"rearport_id":           netboxapi.FilterInt64s(params.SetRearportID),
		"tag_id":                netboxapi.FilterInt64s(params.SetTagID),
		"tag_id__n":             netboxapi.FilterInt64s(params.SetTagIDn),
		"tag_id__any":           netboxapi.FilterInt64s(params.SetTagIDAny),
		"tenant":                netboxapi.FilterStrings(params.SetTenant),
		"tenant__n":             netboxapi.FilterStrings(params.SetTenantn),
		"tenant_group":          netboxapi.FilterStrings(params.SetTenantGroup),
		"tenant_group__n":       netboxapi.FilterStrings(params.SetTenantGroupn),
		"tenant_group_id":       netboxapi.FilterStrings(params.SetTenantGroupID),
		"tenant_group_id__n":    netboxapi.FilterStrings(params.SetTenantGroupIDn),
		"termination_a_id":      netboxapi.FilterInt64s(params.SetTerminationaID),
		"termination_a_type":    netboxapi.FilterStrings(params.SetTerminationaType),
		"termination_a_type__n": netboxapi.FilterStrings(params.SetTerminationaTypen),
		"termination_b_id":      netboxapi.FilterInt64s(params.SetTerminationbID),
		"termination_b_type":    netboxapi.FilterStrings(params.SetTerminationbType),
		"termination_b_type__n": netboxapi.FilterStrings(params.SetTerminationbTypen),
		"unterminated":          netboxapi.FilterBool(params.SetUnterminated),
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
		res, err := d.client.Dcim.DcimCablesListContext(ctx, params, nil, netboxapi.WithQuery(customFieldQuery))
		if err != nil {
			resp.Diagnostics.AddError("Error listing netbox_cables", err.Error())
			return
		}
		if res.Payload == nil || len(res.Payload.Results) == 0 {
			break
		}
		for _, goNetboxModel := range res.Payload.Results {
			responseDTOs = append(responseDTOs, netboxapi.CableResponseDTOFromGoNetbox(goNetboxModel))
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
	resourceWithHooks := &cableResource{client: d.client}
	items := make([]cableResourceModel, 0, len(responseDTOs))
	for _, responseDTO := range responseDTOs {
		var model cableResourceModel
		resp.Diagnostics.Append(flattenCable(ctx, responseDTO, &model)...)
		resourceWithHooks.postReadHook(ctx, &model, &resp.Diagnostics)
		items = append(items, model)
	}
	if resp.Diagnostics.HasError() {
		return
	}
	if int64(len(items)) > limit {
		items = items[:limit]
	}
	list, diags := types.ListValueFrom(ctx, types.ObjectType{AttrTypes: cableResourceAttrTypes()}, items)
	resp.Diagnostics.Append(diags...)
	state.Items = list

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
