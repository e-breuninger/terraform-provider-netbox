// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*moduleDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*moduleDataSource)(nil)
)

// NewModuleDataSource returns a new module data source.
func NewModuleDataSource() datasource.DataSource {
	return &moduleDataSource{}
}

type moduleDataSource struct {
	client *netboxapi.Client
}

// moduleDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type moduleDataSourceModel struct {
	moduleResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	DeviceName         types.String `tfsdk:"device_name"`
	LocationSlug       types.String `tfsdk:"location_slug"`
	LocationID         types.Int64  `tfsdk:"location_id"`
	ManufacturerSlug   types.String `tfsdk:"manufacturer_slug"`
	ManufacturerID     types.Int64  `tfsdk:"manufacturer_id"`
	ModuleTypeModel    types.String `tfsdk:"module_type_model"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	ProfileName        types.String `tfsdk:"profile_name"`
	ProfileID          types.Int64  `tfsdk:"profile_id"`
	Q                  types.String `tfsdk:"q"`
	RackName           types.String `tfsdk:"rack_name"`
	RackID             types.Int64  `tfsdk:"rack_id"`
	RegionSlug         types.String `tfsdk:"region_slug"`
	RegionID           types.String `tfsdk:"region_id"`
	SiteSlug           types.String `tfsdk:"site_slug"`
	SiteGroupSlug      types.String `tfsdk:"site_group_slug"`
	SiteGroupID        types.String `tfsdk:"site_group_id"`
	SiteID             types.Int64  `tfsdk:"site_id"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *moduleDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_module"
}

func (d *moduleDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Data Center Inventory Management (DCIM):A module installed in a device module bay (dcim.module).\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/dcim/module/):\n\n> A module is a field-replaceable hardware component installed within a device which houses its own child components. The most common example is a chassis-based router or switch.\n\nSimilar to devices, modules are instantiated from module types, and any components associated with the module type are automatically instantiated on the new model. Each module must be installed within a module bay on a device, and each module bay may have only one module installed in it.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"device_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the device.",
			},
			"module_bay_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the module bay the module is installed in.",
			},
			"module_type_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the module type.",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operational status. One of: offline, active, planned, staged, failed, decommissioning.",
				Validators: []validator.String{
					stringvalidator.OneOf("offline", "active", "planned", "staged", "failed", "decommissioning"),
				},
			},
			"serial": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Serial number.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"asset_tag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Unique asset tag.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"replicate_components": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether the components of the module type are created on the device along with the module. Only applied when the module is created.",
			},
			"adopt_components": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether existing components of the device that match the module type's templates are adopted by the module. Only applied when the module is created.",
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
			"device_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the device.",
			},
			"location_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the location.",
			},
			"location_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the location.",
			},
			"manufacturer_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the manufacturer.",
			},
			"manufacturer_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the manufacturer.",
			},
			"module_type_model": schema.StringAttribute{
				Optional:    true,
				Description: "Model of the module type.",
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
			"profile_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the profile.",
			},
			"profile_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the profile.",
			},
			"q": schema.StringAttribute{
				Optional:    true,
				Description: "Search term, matched against the fields NetBox searches for this object type.",
			},
			"rack_name": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the rack.",
			},
			"rack_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the rack.",
			},
			"region_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the region.",
			},
			"region_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the region.",
			},
			"site_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site.",
			},
			"site_group_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of the site group.",
			},
			"site_group_id": schema.StringAttribute{
				Optional:    true,
				Description: "Id of the site group.",
			},
			"site_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the site.",
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

func (d *moduleDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *moduleDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data moduleDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := dcim.NewDcimModulesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.DeviceID.IsNull() {
		v := state.DeviceID.ValueInt64()
		params.SetDeviceID([]int64{v})
		hasInput = true
	}
	if !state.ModuleTypeID.IsNull() {
		v := state.ModuleTypeID.ValueInt64()
		params.SetModuleTypeID([]int64{v})
		hasInput = true
	}
	if !state.ModuleBayID.IsNull() {
		v := strconv.FormatInt(state.ModuleBayID.ValueInt64(), 10)
		params.SetModuleBayID([]string{v})
		hasInput = true
	}
	if !state.Status.IsNull() {
		v := state.Status.ValueString()
		params.SetStatus([]string{v})
		hasInput = true
	}
	if !state.Serial.IsNull() {
		v := state.Serial.ValueString()
		params.SetSerial([]string{v})
		hasInput = true
	}
	if !state.AssetTag.IsNull() {
		v := state.AssetTag.ValueString()
		params.SetAssetTag([]string{v})
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
	if !state.DeviceName.IsNull() {
		v := state.DeviceName.ValueString()
		params.SetDevice([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.LocationSlug.IsNull() {
		v := state.LocationSlug.ValueString()
		params.SetLocation([]string{v})
		hasInput = true
	}
	if !state.LocationID.IsNull() {
		v := state.LocationID.ValueInt64()
		params.SetLocationID([]int64{v})
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
	if !state.ModuleTypeModel.IsNull() {
		v := state.ModuleTypeModel.ValueString()
		params.SetModuleType([]string{v})
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
	if !state.RackName.IsNull() {
		v := state.RackName.ValueString()
		params.SetRack([]string{v})
		hasInput = true
	}
	if !state.RackID.IsNull() {
		v := state.RackID.ValueInt64()
		params.SetRackID([]int64{v})
		hasInput = true
	}
	if !state.RegionSlug.IsNull() {
		v := state.RegionSlug.ValueString()
		params.SetRegion([]string{v})
		hasInput = true
	}
	if !state.RegionID.IsNull() {
		v := state.RegionID.ValueString()
		params.SetRegionID([]string{v})
		hasInput = true
	}
	if !state.SiteSlug.IsNull() {
		v := state.SiteSlug.ValueString()
		params.SetSite([]string{v})
		hasInput = true
	}
	if !state.SiteGroupSlug.IsNull() {
		v := state.SiteGroupSlug.ValueString()
		params.SetSiteGroup([]string{v})
		hasInput = true
	}
	if !state.SiteGroupID.IsNull() {
		v := state.SiteGroupID.ValueString()
		params.SetSiteGroupID([]string{v})
		hasInput = true
	}
	if !state.SiteID.IsNull() {
		v := state.SiteID.ValueInt64()
		params.SetSiteID([]int64{v})
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_module.")
		return
	}
	res, err := d.client.Dcim.DcimModulesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_module", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_module",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenModule(ctx, netboxapi.ModuleResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.moduleResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
