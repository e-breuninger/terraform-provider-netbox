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
	_ datasource.DataSource              = (*configContextProfileDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*configContextProfileDataSource)(nil)
)

// NewConfigContextProfileDataSource returns a new config_context_profile data source.
func NewConfigContextProfileDataSource() datasource.DataSource {
	return &configContextProfileDataSource{}
}

type configContextProfileDataSource struct {
	client *netboxapi.Client
}

// configContextProfileDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type configContextProfileDataSourceModel struct {
	configContextProfileResourceModel
	TagSlug         types.String `tfsdk:"tag_slug"`
	AutoSyncEnabled types.Bool   `tfsdk:"auto_sync_enabled"`
	DataFileID      types.Int64  `tfsdk:"data_file_id"`
	DataSourceID    types.Int64  `tfsdk:"data_source_id"`
	DataSynced      types.String `tfsdk:"data_synced"`
	OwnerName       types.String `tfsdk:"owner_name"`
	OwnerGroupName  types.String `tfsdk:"owner_group_name"`
	OwnerGroupID    types.Int64  `tfsdk:"owner_group_id"`
	Q               types.String `tfsdk:"q"`
	TagID           types.Int64  `tfsdk:"tag_id"`
}

func (d *configContextProfileDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_config_context_profile"
}

func (d *configContextProfileDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A NetBox config context profile (extras.configcontextprofile): a JSON schema for the config contexts that use it.\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/extras/configcontextprofile/):\n\n> Profiles can be used to organize [configuration contexts](https://netboxlabs.com/docs/netbox/models/extras/configcontext/) and to enforce a desired structure for their data. The latter is achieved by defining a [JSON schema](https://json-schema.org/) to which all config contexts assigned to the profile must comply. Any context whose data fails validation against the profile's schema cannot be saved.",
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
			"schema": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "JSON schema the config contexts with this profile must satisfy (use jsonencode()).",
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
			"tag_slug": schema.StringAttribute{
				Optional:    true,
				Description: "Slug of a tag the object carries.",
			},
			"auto_sync_enabled": schema.BoolAttribute{
				Optional:    true,
				Description: "Whether auto sync enabled is set.",
			},
			"data_file_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the data file.",
			},
			"data_source_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the data source.",
			},
			"data_synced": schema.StringAttribute{
				Optional:    true,
				Description: "Data synced to match.",
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
		},
	}
}

func (d *configContextProfileDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *configContextProfileDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data configContextProfileDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasConfigContextProfilesListParams()
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
	if !state.AutoSyncEnabled.IsNull() {
		v := state.AutoSyncEnabled.ValueBool()
		params.SetAutoSyncEnabled(&v)
		hasInput = true
	}
	if !state.Created.IsNull() {
		v := state.Created.ValueString()
		params.SetCreated([]string{v})
		hasInput = true
	}
	if !state.DataFileID.IsNull() {
		v := state.DataFileID.ValueInt64()
		params.SetDataFileID([]int64{v})
		hasInput = true
	}
	if !state.DataSourceID.IsNull() {
		v := state.DataSourceID.ValueInt64()
		params.SetDataSourceID([]int64{v})
		hasInput = true
	}
	if !state.DataSynced.IsNull() {
		v := state.DataSynced.ValueString()
		params.SetDataSynced([]string{v})
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
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_config_context_profile.")
		return
	}
	res, err := d.client.Extras.ExtrasConfigContextProfilesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_config_context_profile", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_config_context_profile",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenConfigContextProfile(ctx, netboxapi.ConfigContextProfileResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.configContextProfileResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
