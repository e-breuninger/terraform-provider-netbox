// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
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
	var vID []int64
	var vIDn []int64
	var vIDLt []int64
	var vIDLte []int64
	var vIDGt []int64
	var vIDGte []int64
	var vIDEmpty *bool
	var vLabel []string
	var vLabeln []string
	var vLabelIc []string
	var vLabelNic []string
	var vLabelIe []string
	var vLabelNie []string
	var vLabelIsw []string
	var vLabelNisw []string
	var vLabelIew []string
	var vLabelNiew []string
	var vLabelEmpty *bool
	var vLabelRegex []string
	var vLabelIregex []string
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
	var vType []string
	var vTypen []string
	var vTypeIc []string
	var vTypeNic []string
	var vTypeIe []string
	var vTypeNie []string
	var vTypeIsw []string
	var vTypeNisw []string
	var vTypeIew []string
	var vTypeNiew []string
	var vTypeEmpty *bool
	var vTypeRegex []string
	var vTypeIregex []string
	var vTenantID []int64
	var vTenantIDn []int64
	var vOwnerID []int64
	var vOwnerIDn []int64
	var vTag []string
	var vTagn []string
	var vTagAny []string
	var vSiteID []int64
	var vSite []string
	var vLocationID []int64
	var vRackID []int64
	var vDeviceID []int64
	var vDevice []string
	var vBundle []string
	var vBundlen []string
	var vBundleID []int64
	var vBundleIDn []int64
	var vCircuitterminationID []int64
	var vColor []string
	var vColorn []string
	var vColorIc []string
	var vColorNic []string
	var vColorIe []string
	var vColorNie []string
	var vColorIsw []string
	var vColorNisw []string
	var vColorIew []string
	var vColorNiew []string
	var vColorEmpty *bool
	var vColorRegex []string
	var vColorIregex []string
	var vConsoleportID []int64
	var vConsoleserverportID []int64
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
	var vFrontportID []int64
	var vInterfaceID []int64
	var vLastUpdated []string
	var vLastUpdatedn []string
	var vLastUpdatedLt []string
	var vLastUpdatedLte []string
	var vLastUpdatedGt []string
	var vLastUpdatedGte []string
	var vLastUpdatedEmpty []string
	var vLength []float64
	var vLengthn []float64
	var vLengthLt []float64
	var vLengthLte []float64
	var vLengthGt []float64
	var vLengthGte []float64
	var vLengthEmpty *bool
	var vLengthUnit *string
	var vLengthUnitn *string
	var vLengthUnitIc []string
	var vLengthUnitNic []string
	var vLengthUnitIe []string
	var vLengthUnitNie []string
	var vLengthUnitIsw []string
	var vLengthUnitNisw []string
	var vLengthUnitIew []string
	var vLengthUnitNiew []string
	var vLengthUnitEmpty *bool
	var vLengthUnitRegex []string
	var vLengthUnitIregex []string
	var vLocation []string
	var vOwner []string
	var vOwnern []string
	var vOwnerGroup []string
	var vOwnerGroupn []string
	var vOwnerGroupID []int64
	var vOwnerGroupIDn []int64
	var vPowerfeedID []int64
	var vPoweroutletID []int64
	var vPowerportID []int64
	var vProfile []string
	var vProfilen []string
	var vProfileIc []string
	var vProfileNic []string
	var vProfileIe []string
	var vProfileNie []string
	var vProfileIsw []string
	var vProfileNisw []string
	var vProfileIew []string
	var vProfileNiew []string
	var vProfileEmpty *bool
	var vProfileRegex []string
	var vProfileIregex []string
	var vQ *string
	var vRack []string
	var vRearportID []int64
	var vTagID []int64
	var vTagIDn []int64
	var vTagIDAny []int64
	var vTenant []string
	var vTenantn []string
	var vTenantGroup []string
	var vTenantGroupn []string
	var vTenantGroupID []string
	var vTenantGroupIDn []string
	var vTerminationaID []int64
	var vTerminationaType []string
	var vTerminationaTypen []string
	var vTerminationbID []int64
	var vTerminationbType []string
	var vTerminationbTypen []string
	var vUnterminated *bool
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
		case "label":
			v := value
			vLabel = append(vLabel, v)
		case "label__n":
			v := value
			vLabeln = append(vLabeln, v)
		case "label__ic":
			v := value
			vLabelIc = append(vLabelIc, v)
		case "label__nic":
			v := value
			vLabelNic = append(vLabelNic, v)
		case "label__ie":
			v := value
			vLabelIe = append(vLabelIe, v)
		case "label__nie":
			v := value
			vLabelNie = append(vLabelNie, v)
		case "label__isw":
			v := value
			vLabelIsw = append(vLabelIsw, v)
		case "label__nisw":
			v := value
			vLabelNisw = append(vLabelNisw, v)
		case "label__iew":
			v := value
			vLabelIew = append(vLabelIew, v)
		case "label__niew":
			v := value
			vLabelNiew = append(vLabelNiew, v)
		case "label__empty":
			if vLabelEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'label__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'label__empty' takes a boolean, got %q.", value))
				return
			}
			vLabelEmpty = &v
		case "label__regex":
			v := value
			vLabelRegex = append(vLabelRegex, v)
		case "label__iregex":
			v := value
			vLabelIregex = append(vLabelIregex, v)
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
		case "type":
			v := value
			vType = append(vType, v)
		case "type__n":
			v := value
			vTypen = append(vTypen, v)
		case "type__ic":
			v := value
			vTypeIc = append(vTypeIc, v)
		case "type__nic":
			v := value
			vTypeNic = append(vTypeNic, v)
		case "type__ie":
			v := value
			vTypeIe = append(vTypeIe, v)
		case "type__nie":
			v := value
			vTypeNie = append(vTypeNie, v)
		case "type__isw":
			v := value
			vTypeIsw = append(vTypeIsw, v)
		case "type__nisw":
			v := value
			vTypeNisw = append(vTypeNisw, v)
		case "type__iew":
			v := value
			vTypeIew = append(vTypeIew, v)
		case "type__niew":
			v := value
			vTypeNiew = append(vTypeNiew, v)
		case "type__empty":
			if vTypeEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'type__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'type__empty' takes a boolean, got %q.", value))
				return
			}
			vTypeEmpty = &v
		case "type__regex":
			v := value
			vTypeRegex = append(vTypeRegex, v)
		case "type__iregex":
			v := value
			vTypeIregex = append(vTypeIregex, v)
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
		case "site_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'site_id' takes an integer, got %q.", value))
				return
			}
			vSiteID = append(vSiteID, v)
		case "site":
			v := value
			vSite = append(vSite, v)
		case "location_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'location_id' takes an integer, got %q.", value))
				return
			}
			vLocationID = append(vLocationID, v)
		case "rack_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rack_id' takes an integer, got %q.", value))
				return
			}
			vRackID = append(vRackID, v)
		case "device_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'device_id' takes an integer, got %q.", value))
				return
			}
			vDeviceID = append(vDeviceID, v)
		case "device":
			v := value
			vDevice = append(vDevice, v)
		case "bundle":
			v := value
			vBundle = append(vBundle, v)
		case "bundle__n":
			v := value
			vBundlen = append(vBundlen, v)
		case "bundle_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'bundle_id' takes an integer, got %q.", value))
				return
			}
			vBundleID = append(vBundleID, v)
		case "bundle_id__n":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'bundle_id__n' takes an integer, got %q.", value))
				return
			}
			vBundleIDn = append(vBundleIDn, v)
		case "circuittermination_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'circuittermination_id' takes an integer, got %q.", value))
				return
			}
			vCircuitterminationID = append(vCircuitterminationID, v)
		case "color":
			v := value
			vColor = append(vColor, v)
		case "color__n":
			v := value
			vColorn = append(vColorn, v)
		case "color__ic":
			v := value
			vColorIc = append(vColorIc, v)
		case "color__nic":
			v := value
			vColorNic = append(vColorNic, v)
		case "color__ie":
			v := value
			vColorIe = append(vColorIe, v)
		case "color__nie":
			v := value
			vColorNie = append(vColorNie, v)
		case "color__isw":
			v := value
			vColorIsw = append(vColorIsw, v)
		case "color__nisw":
			v := value
			vColorNisw = append(vColorNisw, v)
		case "color__iew":
			v := value
			vColorIew = append(vColorIew, v)
		case "color__niew":
			v := value
			vColorNiew = append(vColorNiew, v)
		case "color__empty":
			if vColorEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'color__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'color__empty' takes a boolean, got %q.", value))
				return
			}
			vColorEmpty = &v
		case "color__regex":
			v := value
			vColorRegex = append(vColorRegex, v)
		case "color__iregex":
			v := value
			vColorIregex = append(vColorIregex, v)
		case "consoleport_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'consoleport_id' takes an integer, got %q.", value))
				return
			}
			vConsoleportID = append(vConsoleportID, v)
		case "consoleserverport_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'consoleserverport_id' takes an integer, got %q.", value))
				return
			}
			vConsoleserverportID = append(vConsoleserverportID, v)
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
		case "frontport_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'frontport_id' takes an integer, got %q.", value))
				return
			}
			vFrontportID = append(vFrontportID, v)
		case "interface_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'interface_id' takes an integer, got %q.", value))
				return
			}
			vInterfaceID = append(vInterfaceID, v)
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
		case "length":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length' takes a number, got %q.", value))
				return
			}
			vLength = append(vLength, v)
		case "length__n":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length__n' takes a number, got %q.", value))
				return
			}
			vLengthn = append(vLengthn, v)
		case "length__lt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length__lt' takes a number, got %q.", value))
				return
			}
			vLengthLt = append(vLengthLt, v)
		case "length__lte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length__lte' takes a number, got %q.", value))
				return
			}
			vLengthLte = append(vLengthLte, v)
		case "length__gt":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length__gt' takes a number, got %q.", value))
				return
			}
			vLengthGt = append(vLengthGt, v)
		case "length__gte":
			v, err := strconv.ParseFloat(value, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length__gte' takes a number, got %q.", value))
				return
			}
			vLengthGte = append(vLengthGte, v)
		case "length__empty":
			if vLengthEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'length__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length__empty' takes a boolean, got %q.", value))
				return
			}
			vLengthEmpty = &v
		case "length_unit":
			if vLengthUnit != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'length_unit' takes a single value.")
				return
			}
			v := value
			vLengthUnit = &v
		case "length_unit__n":
			if vLengthUnitn != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'length_unit__n' takes a single value.")
				return
			}
			v := value
			vLengthUnitn = &v
		case "length_unit__ic":
			v := value
			vLengthUnitIc = append(vLengthUnitIc, v)
		case "length_unit__nic":
			v := value
			vLengthUnitNic = append(vLengthUnitNic, v)
		case "length_unit__ie":
			v := value
			vLengthUnitIe = append(vLengthUnitIe, v)
		case "length_unit__nie":
			v := value
			vLengthUnitNie = append(vLengthUnitNie, v)
		case "length_unit__isw":
			v := value
			vLengthUnitIsw = append(vLengthUnitIsw, v)
		case "length_unit__nisw":
			v := value
			vLengthUnitNisw = append(vLengthUnitNisw, v)
		case "length_unit__iew":
			v := value
			vLengthUnitIew = append(vLengthUnitIew, v)
		case "length_unit__niew":
			v := value
			vLengthUnitNiew = append(vLengthUnitNiew, v)
		case "length_unit__empty":
			if vLengthUnitEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'length_unit__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'length_unit__empty' takes a boolean, got %q.", value))
				return
			}
			vLengthUnitEmpty = &v
		case "length_unit__regex":
			v := value
			vLengthUnitRegex = append(vLengthUnitRegex, v)
		case "length_unit__iregex":
			v := value
			vLengthUnitIregex = append(vLengthUnitIregex, v)
		case "location":
			v := value
			vLocation = append(vLocation, v)
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
		case "powerfeed_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'powerfeed_id' takes an integer, got %q.", value))
				return
			}
			vPowerfeedID = append(vPowerfeedID, v)
		case "poweroutlet_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'poweroutlet_id' takes an integer, got %q.", value))
				return
			}
			vPoweroutletID = append(vPoweroutletID, v)
		case "powerport_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'powerport_id' takes an integer, got %q.", value))
				return
			}
			vPowerportID = append(vPowerportID, v)
		case "profile":
			v := value
			vProfile = append(vProfile, v)
		case "profile__n":
			v := value
			vProfilen = append(vProfilen, v)
		case "profile__ic":
			v := value
			vProfileIc = append(vProfileIc, v)
		case "profile__nic":
			v := value
			vProfileNic = append(vProfileNic, v)
		case "profile__ie":
			v := value
			vProfileIe = append(vProfileIe, v)
		case "profile__nie":
			v := value
			vProfileNie = append(vProfileNie, v)
		case "profile__isw":
			v := value
			vProfileIsw = append(vProfileIsw, v)
		case "profile__nisw":
			v := value
			vProfileNisw = append(vProfileNisw, v)
		case "profile__iew":
			v := value
			vProfileIew = append(vProfileIew, v)
		case "profile__niew":
			v := value
			vProfileNiew = append(vProfileNiew, v)
		case "profile__empty":
			if vProfileEmpty != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'profile__empty' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'profile__empty' takes a boolean, got %q.", value))
				return
			}
			vProfileEmpty = &v
		case "profile__regex":
			v := value
			vProfileRegex = append(vProfileRegex, v)
		case "profile__iregex":
			v := value
			vProfileIregex = append(vProfileIregex, v)
		case "q":
			if vQ != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'q' takes a single value.")
				return
			}
			v := value
			vQ = &v
		case "rack":
			v := value
			vRack = append(vRack, v)
		case "rearport_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'rearport_id' takes an integer, got %q.", value))
				return
			}
			vRearportID = append(vRearportID, v)
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
		case "tenant":
			v := value
			vTenant = append(vTenant, v)
		case "tenant__n":
			v := value
			vTenantn = append(vTenantn, v)
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
		case "termination_a_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_a_id' takes an integer, got %q.", value))
				return
			}
			vTerminationaID = append(vTerminationaID, v)
		case "termination_a_type":
			v := value
			vTerminationaType = append(vTerminationaType, v)
		case "termination_a_type__n":
			v := value
			vTerminationaTypen = append(vTerminationaTypen, v)
		case "termination_b_id":
			v, err := strconv.ParseInt(value, 10, 64)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'termination_b_id' takes an integer, got %q.", value))
				return
			}
			vTerminationbID = append(vTerminationbID, v)
		case "termination_b_type":
			v := value
			vTerminationbType = append(vTerminationbType, v)
		case "termination_b_type__n":
			v := value
			vTerminationbTypen = append(vTerminationbTypen, v)
		case "unterminated":
			if vUnterminated != nil {
				resp.Diagnostics.AddError("Duplicate filter", "Filter 'unterminated' takes a single value.")
				return
			}
			v, err := strconv.ParseBool(value)
			if err != nil {
				resp.Diagnostics.AddError("Invalid filter value", fmt.Sprintf("Filter 'unterminated' takes a boolean, got %q.", value))
				return
			}
			vUnterminated = &v
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
	if len(vLabel) > 0 {
		params.SetLabel(vLabel)
	}
	if len(vLabeln) > 0 {
		params.SetLabeln(vLabeln)
	}
	if len(vLabelIc) > 0 {
		params.SetLabelIc(vLabelIc)
	}
	if len(vLabelNic) > 0 {
		params.SetLabelNic(vLabelNic)
	}
	if len(vLabelIe) > 0 {
		params.SetLabelIe(vLabelIe)
	}
	if len(vLabelNie) > 0 {
		params.SetLabelNie(vLabelNie)
	}
	if len(vLabelIsw) > 0 {
		params.SetLabelIsw(vLabelIsw)
	}
	if len(vLabelNisw) > 0 {
		params.SetLabelNisw(vLabelNisw)
	}
	if len(vLabelIew) > 0 {
		params.SetLabelIew(vLabelIew)
	}
	if len(vLabelNiew) > 0 {
		params.SetLabelNiew(vLabelNiew)
	}
	if vLabelEmpty != nil {
		params.SetLabelEmpty(vLabelEmpty)
	}
	if len(vLabelRegex) > 0 {
		params.SetLabelRegex(vLabelRegex)
	}
	if len(vLabelIregex) > 0 {
		params.SetLabelIregex(vLabelIregex)
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
	if len(vType) > 0 {
		params.SetType(vType)
	}
	if len(vTypen) > 0 {
		params.SetTypen(vTypen)
	}
	if len(vTypeIc) > 0 {
		params.SetTypeIc(vTypeIc)
	}
	if len(vTypeNic) > 0 {
		params.SetTypeNic(vTypeNic)
	}
	if len(vTypeIe) > 0 {
		params.SetTypeIe(vTypeIe)
	}
	if len(vTypeNie) > 0 {
		params.SetTypeNie(vTypeNie)
	}
	if len(vTypeIsw) > 0 {
		params.SetTypeIsw(vTypeIsw)
	}
	if len(vTypeNisw) > 0 {
		params.SetTypeNisw(vTypeNisw)
	}
	if len(vTypeIew) > 0 {
		params.SetTypeIew(vTypeIew)
	}
	if len(vTypeNiew) > 0 {
		params.SetTypeNiew(vTypeNiew)
	}
	if vTypeEmpty != nil {
		params.SetTypeEmpty(vTypeEmpty)
	}
	if len(vTypeRegex) > 0 {
		params.SetTypeRegex(vTypeRegex)
	}
	if len(vTypeIregex) > 0 {
		params.SetTypeIregex(vTypeIregex)
	}
	if len(vTenantID) > 0 {
		params.SetTenantID(vTenantID)
	}
	if len(vTenantIDn) > 0 {
		params.SetTenantIDn(vTenantIDn)
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
	if len(vSiteID) > 0 {
		params.SetSiteID(vSiteID)
	}
	if len(vSite) > 0 {
		params.SetSite(vSite)
	}
	if len(vLocationID) > 0 {
		params.SetLocationID(vLocationID)
	}
	if len(vRackID) > 0 {
		params.SetRackID(vRackID)
	}
	if len(vDeviceID) > 0 {
		params.SetDeviceID(vDeviceID)
	}
	if len(vDevice) > 0 {
		params.SetDevice(vDevice)
	}
	if len(vBundle) > 0 {
		params.SetBundle(vBundle)
	}
	if len(vBundlen) > 0 {
		params.SetBundlen(vBundlen)
	}
	if len(vBundleID) > 0 {
		params.SetBundleID(vBundleID)
	}
	if len(vBundleIDn) > 0 {
		params.SetBundleIDn(vBundleIDn)
	}
	if len(vCircuitterminationID) > 0 {
		params.SetCircuitterminationID(vCircuitterminationID)
	}
	if len(vColor) > 0 {
		params.SetColor(vColor)
	}
	if len(vColorn) > 0 {
		params.SetColorn(vColorn)
	}
	if len(vColorIc) > 0 {
		params.SetColorIc(vColorIc)
	}
	if len(vColorNic) > 0 {
		params.SetColorNic(vColorNic)
	}
	if len(vColorIe) > 0 {
		params.SetColorIe(vColorIe)
	}
	if len(vColorNie) > 0 {
		params.SetColorNie(vColorNie)
	}
	if len(vColorIsw) > 0 {
		params.SetColorIsw(vColorIsw)
	}
	if len(vColorNisw) > 0 {
		params.SetColorNisw(vColorNisw)
	}
	if len(vColorIew) > 0 {
		params.SetColorIew(vColorIew)
	}
	if len(vColorNiew) > 0 {
		params.SetColorNiew(vColorNiew)
	}
	if vColorEmpty != nil {
		params.SetColorEmpty(vColorEmpty)
	}
	if len(vColorRegex) > 0 {
		params.SetColorRegex(vColorRegex)
	}
	if len(vColorIregex) > 0 {
		params.SetColorIregex(vColorIregex)
	}
	if len(vConsoleportID) > 0 {
		params.SetConsoleportID(vConsoleportID)
	}
	if len(vConsoleserverportID) > 0 {
		params.SetConsoleserverportID(vConsoleserverportID)
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
	if len(vFrontportID) > 0 {
		params.SetFrontportID(vFrontportID)
	}
	if len(vInterfaceID) > 0 {
		params.SetInterfaceID(vInterfaceID)
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
	if len(vLength) > 0 {
		params.SetLength(vLength)
	}
	if len(vLengthn) > 0 {
		params.SetLengthn(vLengthn)
	}
	if len(vLengthLt) > 0 {
		params.SetLengthLt(vLengthLt)
	}
	if len(vLengthLte) > 0 {
		params.SetLengthLte(vLengthLte)
	}
	if len(vLengthGt) > 0 {
		params.SetLengthGt(vLengthGt)
	}
	if len(vLengthGte) > 0 {
		params.SetLengthGte(vLengthGte)
	}
	if vLengthEmpty != nil {
		params.SetLengthEmpty(vLengthEmpty)
	}
	if vLengthUnit != nil {
		params.SetLengthUnit(vLengthUnit)
	}
	if vLengthUnitn != nil {
		params.SetLengthUnitn(vLengthUnitn)
	}
	if len(vLengthUnitIc) > 0 {
		params.SetLengthUnitIc(vLengthUnitIc)
	}
	if len(vLengthUnitNic) > 0 {
		params.SetLengthUnitNic(vLengthUnitNic)
	}
	if len(vLengthUnitIe) > 0 {
		params.SetLengthUnitIe(vLengthUnitIe)
	}
	if len(vLengthUnitNie) > 0 {
		params.SetLengthUnitNie(vLengthUnitNie)
	}
	if len(vLengthUnitIsw) > 0 {
		params.SetLengthUnitIsw(vLengthUnitIsw)
	}
	if len(vLengthUnitNisw) > 0 {
		params.SetLengthUnitNisw(vLengthUnitNisw)
	}
	if len(vLengthUnitIew) > 0 {
		params.SetLengthUnitIew(vLengthUnitIew)
	}
	if len(vLengthUnitNiew) > 0 {
		params.SetLengthUnitNiew(vLengthUnitNiew)
	}
	if vLengthUnitEmpty != nil {
		params.SetLengthUnitEmpty(vLengthUnitEmpty)
	}
	if len(vLengthUnitRegex) > 0 {
		params.SetLengthUnitRegex(vLengthUnitRegex)
	}
	if len(vLengthUnitIregex) > 0 {
		params.SetLengthUnitIregex(vLengthUnitIregex)
	}
	if len(vLocation) > 0 {
		params.SetLocation(vLocation)
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
	if len(vPowerfeedID) > 0 {
		params.SetPowerfeedID(vPowerfeedID)
	}
	if len(vPoweroutletID) > 0 {
		params.SetPoweroutletID(vPoweroutletID)
	}
	if len(vPowerportID) > 0 {
		params.SetPowerportID(vPowerportID)
	}
	if len(vProfile) > 0 {
		params.SetProfile(vProfile)
	}
	if len(vProfilen) > 0 {
		params.SetProfilen(vProfilen)
	}
	if len(vProfileIc) > 0 {
		params.SetProfileIc(vProfileIc)
	}
	if len(vProfileNic) > 0 {
		params.SetProfileNic(vProfileNic)
	}
	if len(vProfileIe) > 0 {
		params.SetProfileIe(vProfileIe)
	}
	if len(vProfileNie) > 0 {
		params.SetProfileNie(vProfileNie)
	}
	if len(vProfileIsw) > 0 {
		params.SetProfileIsw(vProfileIsw)
	}
	if len(vProfileNisw) > 0 {
		params.SetProfileNisw(vProfileNisw)
	}
	if len(vProfileIew) > 0 {
		params.SetProfileIew(vProfileIew)
	}
	if len(vProfileNiew) > 0 {
		params.SetProfileNiew(vProfileNiew)
	}
	if vProfileEmpty != nil {
		params.SetProfileEmpty(vProfileEmpty)
	}
	if len(vProfileRegex) > 0 {
		params.SetProfileRegex(vProfileRegex)
	}
	if len(vProfileIregex) > 0 {
		params.SetProfileIregex(vProfileIregex)
	}
	if vQ != nil {
		params.SetQ(vQ)
	}
	if len(vRack) > 0 {
		params.SetRack(vRack)
	}
	if len(vRearportID) > 0 {
		params.SetRearportID(vRearportID)
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
	if len(vTenant) > 0 {
		params.SetTenant(vTenant)
	}
	if len(vTenantn) > 0 {
		params.SetTenantn(vTenantn)
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
	if len(vTerminationaID) > 0 {
		params.SetTerminationaID(vTerminationaID)
	}
	if len(vTerminationaType) > 0 {
		params.SetTerminationaType(vTerminationaType)
	}
	if len(vTerminationaTypen) > 0 {
		params.SetTerminationaTypen(vTerminationaTypen)
	}
	if len(vTerminationbID) > 0 {
		params.SetTerminationbID(vTerminationbID)
	}
	if len(vTerminationbType) > 0 {
		params.SetTerminationbType(vTerminationbType)
	}
	if len(vTerminationbTypen) > 0 {
		params.SetTerminationbTypen(vTerminationbTypen)
	}
	if vUnterminated != nil {
		params.SetUnterminated(vUnterminated)
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
