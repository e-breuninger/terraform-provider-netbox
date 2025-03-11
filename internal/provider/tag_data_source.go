package provider

import (
	"context"
	"fmt"
	"github.com/e-breuninger/terraform-provider-netbox/internal/provider/models"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

var _ datasource.DataSource = (*tagDataSource)(nil)

func NewTagDataSource() datasource.DataSource {
	return &tagDataSource{}
}

type tagDataSource struct {
	NetboxDataSource
}

func (d *tagDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tag"
}

func (d *tagDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
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
		Description: "Tags are user-defined labels which can be applied to a variety of objects within NetBox. They can be used to establish dimensions of organization beyond the relationships built into NetBox. For example, you might create a tag to identify a particular ownership or condition across several types of objects.",
		Attributes: map[string]schema.Attribute{
			"color_hex": schema.StringAttribute{
				Computed: true,
			},
			"description": schema.StringAttribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"object_types": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
			},
			"slug": schema.StringAttribute{
				Computed: true,
			},
			"id": schema.Int32Attribute{
				Computed:            true,
				Description:         "A unique integer value identifying this tag.",
				MarkdownDescription: "A unique integer value identifying this tag.",
			},
		},
	}
}

func (d *tagDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data models.TagTerraformModel

	// Read Terraform prior state data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)
	resp.Diagnostics.Append(testClient(d.provider.client)...)

	if resp.Diagnostics.HasError() {
		return
	}

	// Read API call logic
	nameList := make([]string, 1)
	nameList = append(nameList, data.Name.ValueString())
	paginatedTagList, httpCode, err := d.provider.client.ExtrasAPI.ExtrasTagsList(ctx).Name(nameList).Execute()

	if err != nil {
		if httpCode != nil && httpCode.StatusCode == 404 {
			resp.State.RemoveResource(ctx)
			return
		} else {
			resp.Diagnostics.AddError(
				"Unable to retrieve Tag value.",
				err.Error(),
			)
			return
		}
	}

	if paginatedTagList.Count > 1 {
		resp.Diagnostics.AddError("more than one tag returned, specify a more narrow filter", "")
		return
	} else if paginatedTagList.Count == 0 {
		resp.Diagnostics.AddError("no tag found matching filter.", "")
		return
	}

	errors := data.ReadAPI(ctx, &paginatedTagList.Results[0])

	if errors.HasError() {
		resp.Diagnostics.Append(errors...)
		return
	}

	// Save updated data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}
