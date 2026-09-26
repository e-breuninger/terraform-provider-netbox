// Code generated automatically. DO NOT EDIT.

package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// supportedVersions lists the NetBox versions the provider was tested against:
// the spec's backend.options.supported_versions, in its order. The provider's
// startup version check warns when the running NetBox is not among them, and
// the supported_versions data source serves the list.
var supportedVersions = []string{"4.6.8", "4.6.9", "4.6.10"}

// The data source has no NetBox object behind it, so it registers like a
// hand-written companion instead of through the generated DataSources list.
func init() {
	companionDataSources = append(companionDataSources, NewSupportedVersionsDataSource)
}

var _ datasource.DataSource = (*supportedVersionsDataSource)(nil)

// NewSupportedVersionsDataSource returns a new supported_versions data source,
// which lists the NetBox versions the provider was tested against.
func NewSupportedVersionsDataSource() datasource.DataSource {
	return &supportedVersionsDataSource{}
}

type supportedVersionsDataSource struct{}

type supportedVersionsDataSourceModel struct {
	Versions types.List `tfsdk:"versions"`
}

func (d *supportedVersionsDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_supported_versions"
}

func (d *supportedVersionsDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "The NetBox versions this provider release was tested against. Any other version only gets a \"Possibly unsupported Netbox version\" warning at startup, so compare a planned NetBox upgrade against this list before applying it.",
		Attributes: map[string]schema.Attribute{
			"versions": schema.ListAttribute{
				ElementType: types.StringType,
				Computed:    true,
				Description: "The supported NetBox versions as MAJOR.MINOR.PATCH.",
			},
		},
	}
}

func (d *supportedVersionsDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	versions, diags := types.ListValueFrom(ctx, types.StringType, supportedVersions)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, supportedVersionsDataSourceModel{Versions: versions})...)
}
