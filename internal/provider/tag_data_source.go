package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
	"github.com/netbox-community/go-netbox/v4"
)

var _ datasource.DataSource = (*tagDataSource)(nil)

func NewTagDataSource() datasource.DataSource {
	return &tagDataSource{}
}

type tagDataSource struct {
	provider *netboxProvider
}

type tagDataSourceModel struct {
	Color       types.String `tfsdk:"color"`
	Description types.String `tfsdk:"description"`
	Id          types.Int32  `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	ObjectTypes types.List   `tfsdk:"object_types"`
	Slug        types.String `tfsdk:"slug"`
	ColorHex    types.String `tfsdk:"color_hex"`
}

func (d *tagDataSource) readAPI(ctx context.Context, data *tagDataSourceModel, tag *netbox.Tag) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	data.Id = types.Int32Value(tag.Id)
	data.Name = types.StringValue(tag.Name)
	data.Slug = types.StringValue(tag.Slug)
	data.Color = types.StringPointerValue(tag.Color)
	data.Description = types.StringPointerValue(tag.Description)
	listObjectTypes, err := types.ListValueFrom(ctx, types.StringType, tag.ObjectTypes)
	if err != nil {
		return err
	}
	data.ObjectTypes = listObjectTypes

	return diags
}

func (d *tagDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (d *tagDataSource) Configure(ctx context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(*netboxProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *netbox.apiCLient, got: %T, Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.provider = provider
}

func (d *tagDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"color": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.Int32Attribute{
				Required:            true,
				Description:         "A unique integer value identifying this tag.",
				MarkdownDescription: "A unique integer value identifying this tag.",
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"object_types": schema.ListAttribute{
				Computed: true,
				Optional: true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
			},
		},
	}
}

func (d *tagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data tagDataSourceModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(d.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	tag, httpCode, err := d.provider.client.ExtrasAPI.ExtrasTagsRetrieve(ctx, data.Id.ValueInt32()).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to retrieve Webhook value.",
				err.Error(),
			)
			return
		}
	}

	errors := d.readAPI(ctx, &data, tag)

	if errors != nil {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
