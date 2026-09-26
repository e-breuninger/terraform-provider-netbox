// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"
	"fmt"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-framework-jsontypes/jsontypes"
	"github.com/hashicorp/terraform-plugin-framework-validators/stringvalidator"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure provider defined types fully satisfy framework interfaces.
var (
	_ datasource.DataSource              = (*configContextDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*configContextDataSource)(nil)
)

// NewConfigContextDataSource returns a new config_context data source.
func NewConfigContextDataSource() datasource.DataSource {
	return &configContextDataSource{}
}

type configContextDataSource struct {
	client *netboxapi.Client
}

// configContextDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type configContextDataSourceModel struct {
	configContextResourceModel
	NameContains types.String `tfsdk:"name_contains"`
}

func (d *configContextDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_context"
}

func (d *configContextDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox config context (extras.configcontext): JSON data merged into the rendered context of the devices and virtual machines its scope selects.\n\nFrom the [official documentation](https://docs.netbox.dev/en/stable/models/extras/configcontext/):\n\n> Context data is made available to devices and/or virtual machines based on their relationships to other objects in NetBox. For example, context data can be associated only with devices assigned to a particular site, or only to virtual machines in a certain cluster.",
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
			"weight": schema.Int64Attribute{
				Computed:    true,
				Description: "Merge order; higher weights are applied later and win.",
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"is_active": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the context is applied.",
			},
			"data": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "The context data as JSON text (use jsonencode()).",
			},
			"profile_id": schema.Int64Attribute{
				Computed:    true,
				Description: "Id of the config context profile whose schema the data must satisfy.",
			},
			"region_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the regions the context applies to; empty means all.",
			},
			"regions": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use region_ids instead.",
				Description:        "Deprecated alias of region_ids.",
			},
			"site_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the site groups the context applies to; empty means all.",
			},
			"site_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use site_group_ids instead.",
				Description:        "Deprecated alias of site_group_ids.",
			},
			"site_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the sites the context applies to; empty means all.",
			},
			"sites": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use site_ids instead.",
				Description:        "Deprecated alias of site_ids.",
			},
			"location_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the locations the context applies to; empty means all.",
			},
			"locations": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use location_ids instead.",
				Description:        "Deprecated alias of location_ids.",
			},
			"device_type_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the device types the context applies to; empty means all.",
			},
			"device_types": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use device_type_ids instead.",
				Description:        "Deprecated alias of device_type_ids.",
			},
			"device_role_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the device roles the context applies to; empty means all.",
			},
			"roles": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use device_role_ids instead.",
				Description:        "Deprecated alias of device_role_ids.",
			},
			"platform_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the platforms the context applies to; empty means all.",
			},
			"platforms": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use platform_ids instead.",
				Description:        "Deprecated alias of platform_ids.",
			},
			"cluster_type_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the cluster types the context applies to; empty means all.",
			},
			"cluster_types": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use cluster_type_ids instead.",
				Description:        "Deprecated alias of cluster_type_ids.",
			},
			"cluster_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the cluster groups the context applies to; empty means all.",
			},
			"cluster_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use cluster_group_ids instead.",
				Description:        "Deprecated alias of cluster_group_ids.",
			},
			"cluster_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the clusters the context applies to; empty means all.",
			},
			"clusters": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use cluster_ids instead.",
				Description:        "Deprecated alias of cluster_ids.",
			},
			"tenant_group_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the tenant groups the context applies to; empty means all.",
			},
			"tenant_groups": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use tenant_group_ids instead.",
				Description:        "Deprecated alias of tenant_group_ids.",
			},
			"tenant_ids": schema.SetAttribute{
				ElementType: types.Int64Type,
				Computed:    true,
				Description: "Ids of the tenants the context applies to; empty means all.",
			},
			"tenants": schema.SetAttribute{
				ElementType:        types.Int64Type,
				Computed:           true,
				DeprecationMessage: "Use tenant_ids instead.",
				Description:        "Deprecated alias of tenant_ids.",
			},
			"tag_slugs": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Slugs of the tags an object must carry for the context to apply; empty means any. This is a scope, not the context's own tags.",
			},
			"owner_id": schema.Int64Attribute{
				Optional:    true,
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
			"name_contains": schema.StringAttribute{
				Optional:    true,
				Description: "Case-insensitive substring of the name.",
			},
		},
	}
}

func (d *configContextDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *configContextDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data configContextDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasConfigContextsListParams()
	hasInput := false
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
	if !state.IsActive.IsNull() {
		v := state.IsActive.ValueBool()
		params.SetIsActive(&v)
		hasInput = true
	}
	if !state.NameContains.IsNull() {
		v := state.NameContains.ValueString()
		params.SetNameIc([]string{v})
		hasInput = true
	}
	if !state.OwnerID.IsNull() {
		v := state.OwnerID.ValueInt64()
		params.SetOwnerID([]int64{v})
		hasInput = true
	}
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set id and/or name and/or is_active and/or name_contains and/or owner_id to look up a netbox_config_context.")
		return
	}
	res, err := d.client.Extras.ExtrasConfigContextsListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_config_context", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_config_context",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenConfigContext(ctx, netboxapi.ConfigContextResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.configContextResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
