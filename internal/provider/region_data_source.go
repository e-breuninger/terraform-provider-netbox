package provider

import (
	"context"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TODO: Support filters
var _ datasource.DataSource = (*regionDataSource)(nil)

func NewRegionDataSource() datasource.DataSource {
	return &regionDataSource{}
}

type regionDataSource struct {
	NetboxDataSource
}

func (d *regionDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_region"
}

func (d *regionDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: `:meta:subcategory:Data Center Inventory Management (DCIM):`,
		Attributes: map[string]schema.Attribute{
			"custom_fields": schema.MapAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.Int32Attribute{
				Required:            true,
				Description:         "A unique integer value identifying this region.",
				MarkdownDescription: "A unique integer value identifying this region.",
			},
			"name": schema.StringAttribute{
				Computed: true,
			},
			"parent": schema.Int32Attribute{
				Computed: true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
			},
			"tags": schema.ListAttribute{
				ElementType: types.Int32Type,
				Computed:    true,
			},
		},
	}
}

func (d *regionDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.RegionTerraformModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(d.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	region, httpCode, err := d.provider.client.DcimAPI.DcimRegionsRetrieve(ctx, data.Id.ValueInt32()).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to retrieve region value.",
				err.Error(),
			)
			return
		}
	}

	errors := data.ReadAPI(ctx, region)
	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}
	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
