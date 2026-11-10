// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"
	"net/url"
	"regexp"
	"strconv"

	"github.com/e-breuninger/terraform-provider-netbox/internal/conv"
	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework-validators/float64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/int64validator"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*virtualMachineTypeDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*virtualMachineTypeDataSource)(nil)
)

// NewVirtualMachineTypeDataSource returns a new virtual_machine_type data source.
func NewVirtualMachineTypeDataSource() datasource.DataSource {
	return &virtualMachineTypeDataSource{}
}

type virtualMachineTypeDataSource struct {
	client *netboxapi.Client
}

// virtualMachineTypeDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type virtualMachineTypeDataSourceModel struct {
	virtualMachineTypeResourceModel
	TagSlug            types.String `tfsdk:"tag_slug"`
	DefaultPlatform    types.String `tfsdk:"default_platform"`
	OwnerName          types.String `tfsdk:"owner_name"`
	OwnerGroupName     types.String `tfsdk:"owner_group_name"`
	OwnerGroupID       types.Int64  `tfsdk:"owner_group_id"`
	Q                  types.String `tfsdk:"q"`
	TagID              types.Int64  `tfsdk:"tag_id"`
	CustomFieldFilters types.Map    `tfsdk:"custom_field_filters"`
}

func (d *virtualMachineTypeDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_virtual_machine_type"
}

func (d *virtualMachineTypeDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Virtualization:A NetBox virtual machine type (virtualization.virtualmachinetype): a reusable sizing template for virtual machines.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/virtualization/virtualmachinetype/):\n\n> A virtual machine type defines a reusable classification and default configuration for [virtual machines](https://netboxlabs.com/docs/netbox/models/virtualization/virtualmachine/).\n>\n> A type can optionally provide default values for a VM's [platform](https://netboxlabs.com/docs/netbox/models/dcim/platform/), vCPU allocation, and memory allocation. When a virtual machine is created with an assigned type, any unset values among these fields will inherit their defaults from the type. Changes made to a virtual machine type do **not** apply retroactively to existing virtual machines.",
		Attributes: map[string]schema.Attribute{
			"id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "NetBox id.",
			},
			"name": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
				},
			},
			"slug": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "URL-friendly unique shorthand; derived from name when not set.",
				Validators: []validator.String{
					stringvalidator.LengthBetween(1, 100),
					stringvalidator.RegexMatches(regexp.MustCompile(`^[-a-zA-Z0-9_]+$`), ""),
				},
			},
			"default_platform_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Id of the platform virtual machines of this type default to.",
			},
			"default_vcpus": schema.Float64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default number of virtual CPUs.",
				Validators: []validator.Float64{
					float64validator.Between(0.01, 10000),
				},
			},
			"default_memory": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Default memory in megabytes. Must be a whole number.",
				Validators: []validator.Int64{
					int64validator.Between(0, 2147483647),
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
			"virtual_machine_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of virtual machines of the type.",
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
			"default_platform": schema.StringAttribute{
				Optional:    true,
				Description: "Default platform to match.",
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

func (d *virtualMachineTypeDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *virtualMachineTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data virtualMachineTypeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := virtualization.NewVirtualizationVirtualMachineTypesListParams()
	hasInput := false
	queryParams := url.Values{}
	if !state.ID.IsNull() {
		v := state.ID.ValueInt64()
		params.SetID([]int64{v})
		hasInput = true
	}
	if !state.Name.IsNull() {
		v := state.Name.ValueString()
		params.SetName([]string{v})
		hasInput = true
	}
	if !state.Slug.IsNull() {
		v := state.Slug.ValueString()
		params.SetSlug([]string{v})
		hasInput = true
	}
	if !state.DefaultPlatformID.IsNull() {
		v := strconv.FormatInt(state.DefaultPlatformID.ValueInt64(), 10)
		params.SetDefaultPlatformID([]string{v})
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
	if !state.DefaultMemory.IsNull() {
		v := state.DefaultMemory.ValueInt64()
		params.SetDefaultMemory([]int64{v})
		hasInput = true
	}
	if !state.DefaultPlatform.IsNull() {
		v := state.DefaultPlatform.ValueString()
		params.SetDefaultPlatform([]string{v})
		hasInput = true
	}
	if !state.DefaultVcpus.IsNull() {
		v := state.DefaultVcpus.ValueFloat64()
		params.SetDefaultVcpus([]float64{v})
		hasInput = true
	}
	if !state.Description.IsNull() {
		v := state.Description.ValueString()
		params.SetDescription([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
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
	if !state.Q.IsNull() {
		v := state.Q.ValueString()
		params.SetQ(&v)
		hasInput = true
	}
	if !state.TagID.IsNull() {
		v := state.TagID.ValueInt64()
		params.SetTagID([]int64{v})
		hasInput = true
	}
	if !state.VirtualMachineCount.IsNull() {
		v := state.VirtualMachineCount.ValueInt64()
		params.SetVirtualMachineCount([]int64{v})
		hasInput = true
	}
	if !state.CustomFieldFilters.IsNull() {
		for name, value := range conv.MapTo[string](ctx, state.CustomFieldFilters, &resp.Diagnostics) {
			queryParams.Add("cf_"+name, value)
		}
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_virtual_machine_type.")
		return
	}
	res, err := d.client.Virtualization.VirtualizationVirtualMachineTypesListContext(ctx, params, nil, netboxapi.WithQuery(queryParams))
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_virtual_machine_type", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_virtual_machine_type",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenVirtualMachineType(ctx, netboxapi.VirtualMachineTypeResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.virtualMachineTypeResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
