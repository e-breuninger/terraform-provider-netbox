// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*moduleTypeDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*moduleTypeDataSource)(nil)
)

// NewModuleTypeDataSource returns a new module_type data source.
func NewModuleTypeDataSource() datasource.DataSource {
	return &moduleTypeDataSource{}
}

type moduleTypeDataSource struct {
	client *netboxapi.Client
}

// moduleTypeDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type moduleTypeDataSourceModel struct {
	moduleTypeResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	ConsolePorts       types.Bool   `tfsdk:"console_ports"`
	ConsoleServerPorts types.Bool   `tfsdk:"console_server_ports"`
	Interfaces         types.Bool   `tfsdk:"interfaces"`
	ManufacturerSlug   types.String `tfsdk:"manufacturer_slug"`
	ModuleBays         types.Bool   `tfsdk:"module_bays"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	PassThroughPorts   types.Bool   `tfsdk:"pass_through_ports"`
	PowerOutlets       types.Bool   `tfsdk:"power_outlets"`
	PowerPorts         types.Bool   `tfsdk:"power_ports"`
	ProfileName        types.String `tfsdk:"profile_name"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *moduleTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_module_type"
}

func (d *moduleTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A module type (dcim.moduletype): a hardware module model made by a manufacturer.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/moduletype/):\n\n> A module type represents a specific make and model of hardware component which is installable within a device's module bay and has its own child components. For example, consider a chassis-based switch or router with a number of field-replaceable line cards. Each line card has its own model number and includes a certain set of components such as interfaces. Each module type may have a manufacturer, model number, and part number assigned to it.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"manufacturer_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the manufacturer.",
			},
			"model": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Model name (unique per manufacturer).",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"part_number": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Discrete part number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"profile_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the module type profile.",
			},
			"attributes": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Attribute values of the module type as JSON text, as the profile's schema defines them (use jsonencode()).",
			},
			"airflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Airflow direction. One of: front-to-rear, rear-to-front, left-to-right, right-to-left, side-to-rear, passive.",
				Validators: []validator.String{
					stringvalidator.OneOf("front-to-rear", "rear-to-front", "left-to-right", "right-to-left", "side-to-rear", "passive"),
				},
			},
			"weight": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Weight of the module type (requires weight_unit).",
				Validators: []validator.Float64{
					float64validator.Between(-1000000, 1000000),
				},
			},
			"weight_unit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unit of weight. One of: kg, g, lb, oz.",
				Validators: []validator.String{
					stringvalidator.OneOf("kg", "g", "lb", "oz"),
				},
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"comments": schema.StringAttribute{
				Computed: true,
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the owner the object is assigned to.",
			},
			"created": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"last_updated": schema.StringAttribute{
				Optional: true,
				Computed: true,
			},
			"url": schema.StringAttribute{
				Computed: true,
			},
			"module_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of modules of the module type.",
			},
			"console_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of console port templates of the module type.",
			},
			"console_server_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of console server port templates of the module type.",
			},
			"power_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of power port templates of the module type.",
			},
			"power_outlet_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of power outlet templates of the module type.",
			},
			"interface_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of interface templates of the module type.",
			},
			"front_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of front port templates of the module type.",
			},
			"rear_port_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of rear port templates of the module type.",
			},
			"module_bay_template_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of module bay templates of the module type.",
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"console_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has console ports.",
			},
			"console_server_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has console server ports.",
			},
			"interfaces": schema.BoolAttribute{
				Optional:    true,
				Description: "Has interfaces.",
			},
			"manufacturer_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the manufacturer.",
			},
			"module_bays": schema.BoolAttribute{
				Optional:    true,
				Description: "Has module bays.",
			},
			"owner_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner.",
			},
			"owner_group_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the owner group.",
			},
			"owner_group_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the owner group.",
			},
			"pass_through_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has pass-through ports.",
			},
			"power_outlets": schema.BoolAttribute{
				Optional:    true,
				Description: "Has power outlets.",
			},
			"power_ports": schema.BoolAttribute{
				Optional:    true,
				Description: "Has power ports.",
			},
			"profile_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the profile.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"tag_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of a tag the object carries.",
			},
			"custom_field_filters": schema.MapAttribute{
				Optional:    true,
				ElementType: types.StringType,
				Description: "Custom field values to match, by field name, e.g. { tier = \"gold\" }: each entry filters as cf_<field name> with the field's own filter logic, loose (case-insensitive substring) or exact.",
			},
		},
	}
}

func (d *moduleTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *moduleTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data moduleTypeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimModuleTypesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Model.IsNull() {
		v := state.Model.ValueString()
		params.SetModel([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !state.TagSlug.IsNull() {
		v := state.TagSlug.ValueString()
		params.SetTag([]string{v})
		hasInput = true
	}
	if !state.Airflow.IsNull() {
		v := state.Airflow.ValueString()
		params.SetAirflow(&v)
		hasInput = true
	}
	if !state.ConsolePortTemplateCount.IsNull() {
		v := state.ConsolePortTemplateCount.ValueInt64()
		params.SetConsolePortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.ConsolePorts.IsNull() {
		v := state.ConsolePorts.ValueBool()
		params.SetConsolePorts(&v)
		hasInput = true
	}
	if !state.ConsoleServerPortTemplateCount.IsNull() {
		v := state.ConsoleServerPortTemplateCount.ValueInt64()
		params.SetConsoleServerPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.ConsoleServerPorts.IsNull() {
		v := state.ConsoleServerPorts.ValueBool()
		params.SetConsoleServerPorts(&v)
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.FrontPortTemplateCount.IsNull() {
		v := state.FrontPortTemplateCount.ValueInt64()
		params.SetFrontPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.InterfaceTemplateCount.IsNull() {
		v := state.InterfaceTemplateCount.ValueInt64()
		params.SetInterfaceTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.Interfaces.IsNull() {
		v := state.Interfaces.ValueBool()
		params.SetInterfaces(&v)
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.ManufacturerSlug.IsNull() {
		v := state.ManufacturerSlug.ValueString()
		params.SetManufacturer([]string{v})
		hasInput = true
	}
	if !state.ManufacturerID.IsNull() {
		v := state.ManufacturerID.ValueInt64()
		params.SetManufacturerID([]int64{v})
		hasInput = true
	}
	if !state.ModuleBayTemplateCount.IsNull() {
		v := state.ModuleBayTemplateCount.ValueInt64()
		params.SetModuleBayTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.ModuleBays.IsNull() {
		v := state.ModuleBays.ValueBool()
		params.SetModuleBays(&v)
		hasInput = true
	}
	if !state.ModuleCount.IsNull() {
		v := state.ModuleCount.ValueInt64()
		params.SetModuleCount([]int64{v})
		hasInput = true
	}
	if !state.OwnerName.IsNull() {
		v := state.OwnerName.ValueString()
		params.SetOwner([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupName.IsNull() {
		v := state.OwnerGroupName.ValueString()
		params.SetOwnerGroup([]string{v})
		hasInput = true
	}
	if !state.OwnerGroupID.IsNull() {
		v := state.OwnerGroupID.ValueInt64()
		params.SetOwnerGroupID([]int64{v})
		hasInput = true
	}
	if !state.PartNumber.IsNull() {
		v := state.PartNumber.ValueString()
		params.SetPartNumber([]string{v})
		hasInput = true
	}
	if !state.PassThroughPorts.IsNull() {
		v := state.PassThroughPorts.ValueBool()
		params.SetPassThroughPorts(&v)
		hasInput = true
	}
	if !state.PowerOutletTemplateCount.IsNull() {
		v := state.PowerOutletTemplateCount.ValueInt64()
		params.SetPowerOutletTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.PowerOutlets.IsNull() {
		v := state.PowerOutlets.ValueBool()
		params.SetPowerOutlets(&v)
		hasInput = true
	}
	if !state.PowerPortTemplateCount.IsNull() {
		v := state.PowerPortTemplateCount.ValueInt64()
		params.SetPowerPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.PowerPorts.IsNull() {
		v := state.PowerPorts.ValueBool()
		params.SetPowerPorts(&v)
		hasInput = true
	}
	if !state.ProfileName.IsNull() {
		v := state.ProfileName.ValueString()
		params.SetProfile([]string{v})
		hasInput = true
	}
	if !state.ProfileID.IsNull() {
		v := state.ProfileID.ValueInt64()
		params.SetProfileID([]int64{v})
		hasInput = true
	}
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.RearPortTemplateCount.IsNull() {
		v := state.RearPortTemplateCount.ValueInt64()
		params.SetRearPortTemplateCount([]int64{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.Weight.IsNull() {
		v := state.Weight.ValueFloat64()
		params.SetWeight([]float64{v})
		hasInput = true
	}
	if !state.WeightUnit.IsNull() {
		v := state.WeightUnit.ValueString()
		params.SetWeightUnit(&v)
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_module_type.")
		return
	}
	res, err := d.client.Dcim.DcimModuleTypesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_module_type", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_module_type",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenModuleType(ctx, netboxapi.ModuleTypeResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.moduleTypeResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
