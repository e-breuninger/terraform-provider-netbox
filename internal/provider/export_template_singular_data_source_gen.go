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
	_ datasource.DataSource              = (*exportTemplateDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*exportTemplateDataSource)(nil)
)

// NewExportTemplateDataSource returns a new export_template data source.
func NewExportTemplateDataSource() datasource.DataSource {
	return &exportTemplateDataSource{}
}

type exportTemplateDataSource struct {
	client *netboxapi.Client
}

// exportTemplateDataSourceModel is the data source model: the resource model plus lookup-only inputs.
type exportTemplateDataSourceModel struct {
	exportTemplateResourceModel
	AutoSyncEnabled types.Bool   `tfsdk:"auto_sync_enabled"`
	DataFileID      types.Int64  `tfsdk:"data_file_id"`
	DataSourceID    types.Int64  `tfsdk:"data_source_id"`
	DataSynced      types.String `tfsdk:"data_synced"`
	ObjectType      types.String `tfsdk:"object_type"`
	ObjectTypeID    types.Int64  `tfsdk:"object_type_id"`
	OwnerName       types.String `tfsdk:"owner_name"`
	OwnerGroupName  types.String `tfsdk:"owner_group_name"`
	OwnerGroupID    types.Int64  `tfsdk:"owner_group_id"`
	Q               types.String `tfsdk:"q"`
}

func (d *exportTemplateDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_export_template"
}

func (d *exportTemplateDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:Extras:A Jinja2 export template for the object types it applies to (extras.exporttemplate).\n\nFrom the [official documentation](https://netboxlabs.com/docs/netbox/models/extras/exporttemplate/):\n\n> Export templates are used to render arbitrary data from a set of NetBox objects. For example, you might want to automatically generate a network monitoring service configuration from a list of device objects. See the [export templates documentation](https://netboxlabs.com/docs/netbox/customization/export-templates/) for more information.",
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
			"object_types": schema.SetAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "Object types the template can export, as app-labelled content types (dcim.device).",
			},
			"template_code": schema.StringAttribute{
				Computed:    true,
				Description: "Jinja2 template code; the queryset is available as {{ queryset }}.",
			},
			"mime_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MIME type of the rendered output, e.g. text/csv.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(50),
				},
			},
			"file_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File name of the rendered output, without extension.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
			},
			"file_extension": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Extension appended to the file name, e.g. csv.",
				Validators: []validator.String{
					stringvalidator.LengthAtMost(15),
				},
			},
			"as_attachment": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Download the rendered output as a file instead of displaying it.",
			},
			"environment_params": schema.StringAttribute{
				CustomType:  jsontypes.NormalizedType{},
				Computed:    true,
				Description: "Parameters for the Jinja2 environment as JSON text (use jsonencode()).",
			},
			"description": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Validators: []validator.String{
					stringvalidator.LengthAtMost(200),
				},
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
			"object_type": schema.StringAttribute{
				Optional:    true,
				Description: "Object type to match.",
			},
			"object_type_id": schema.Int64Attribute{
				Optional:    true,
				Description: "Id of the object type.",
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
		},
	}
}

func (d *exportTemplateDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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

func (d *exportTemplateDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data exportTemplateDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	if resp.Diagnostics.HasError() {
		return
	}
	state := &data

	params := extras.NewExtrasExportTemplatesListParams()
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
	if !state.AsAttachment.IsNull() {
		v := state.AsAttachment.ValueBool()
		params.SetAsAttachment(&v)
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
	if !state.FileExtension.IsNull() {
		v := state.FileExtension.ValueString()
		params.SetFileExtension([]string{v})
		hasInput = true
	}
	if !state.FileName.IsNull() {
		v := state.FileName.ValueString()
		params.SetFileName([]string{v})
		hasInput = true
	}
	if !state.LastUpdated.IsNull() {
		v := state.LastUpdated.ValueString()
		params.SetLastUpdated([]string{v})
		hasInput = true
	}
	if !state.MimeType.IsNull() {
		v := state.MimeType.ValueString()
		params.SetMimeType([]string{v})
		hasInput = true
	}
	if !state.ObjectType.IsNull() {
		v := state.ObjectType.ValueString()
		params.SetObjectType([]string{v})
		hasInput = true
	}
	if !state.ObjectTypeID.IsNull() {
		v := state.ObjectTypeID.ValueInt64()
		params.SetObjectTypeID([]int64{v})
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
	if !hasInput {
		resp.Diagnostics.AddError("Missing lookup attribute", "Set at least one lookup attribute to look up a netbox_export_template.")
		return
	}
	res, err := d.client.Extras.ExtrasExportTemplatesListContext(ctx, params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Error reading netbox_export_template", err.Error())
		return
	}
	if res.Payload == nil || res.Payload.Count == nil || *res.Payload.Count != 1 {
		n := int64(0)
		if res.Payload != nil && res.Payload.Count != nil {
			n = *res.Payload.Count
		}
		resp.Diagnostics.AddError(
			"Lookup did not match exactly one netbox_export_template",
			fmt.Sprintf("The query matched %d objects; refine the lookup so exactly one matches.", n),
		)
		return
	}
	resp.Diagnostics.Append(flattenExportTemplate(ctx, netboxapi.ExportTemplateResponseDTOFromGoNetbox(res.Payload.Results[0]), &state.exportTemplateResourceModel)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, state)...)
}
