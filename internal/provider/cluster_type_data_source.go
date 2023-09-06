package provider

import (
	"context"
	"fmt"
	"strconv"

	netboxclient "github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// Ensure the implementation satisfies the expected interfaces.
var (
	_ datasource.DataSource              = &clusterTypeDataSource{}
	_ datasource.DataSourceWithConfigure = &clusterTypeDataSource{}
)

// NewClusterTypeDataSource is a helper function to simplify the provider implementation.
func NewClusterTypeDataSource() datasource.DataSource {
	return &clusterTypeDataSource{}
}

// clusterTypeDataSourceModel maps the data source schema data.
type clusterTypeDataSourceModel struct {
	ClusterTypeID types.Int64  `tfsdk:"cluster_type_id"`
	Name          types.String `tfsdk:"name"`
	ID            types.String `tfsdk:"id"`
}

// clusterTypeDataSource is the data source implementation.
type clusterTypeDataSource struct {
	client *netboxclient.NetBoxAPI
}

// Metadata returns the data source type name.
func (d *clusterTypeDataSource) Metadata(_ context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_cluster_type"
}

// Schema defines the schema for the data source.
func (d *clusterTypeDataSource) Schema(_ context.Context, _ datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Attributes: map[string]schema.Attribute{
			"cluster_type_id": schema.Int64Attribute{
				Computed: true,
			},
			"name": schema.StringAttribute{
				Required: true,
			},
			"id": schema.StringAttribute{
				MarkdownDescription: "Example identifier",
				Computed:            true,
			},
		},
	}
}

// Read refreshes the Terraform state with the latest data.
func (d *clusterTypeDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var data clusterTypeDataSourceModel

	// Read Terraform configuration data into the model
	resp.Diagnostics.Append(req.Config.Get(ctx, &data)...)

	if resp.Diagnostics.HasError() {
		return
	}

	params := virtualization.NewVirtualizationClusterTypesListParams()
	params.Namen = data.Name.ValueStringPointer()
	limit := int64(2) // Limit of 2 is enough
	params.Limit = &limit

	res, err := d.client.Virtualization.VirtualizationClusterTypesList(params, nil)
	if err != nil {
		resp.Diagnostics.AddError("Client Error", fmt.Sprintf("Unable to read cluster type, got error: %s", err))
		return
	}

	if *res.GetPayload().Count > int64(1) {
		resp.Diagnostics.AddError("Client Error", "more than one result, specify a more narrow filter")
		return
	}
	if *res.GetPayload().Count == int64(0) {
		resp.Diagnostics.AddError("Client Error", "no result")
		return
	}

	result := res.GetPayload().Results[0]
	// data.Id = types.StringValue(result.ID)
	data.ID = types.StringValue(strconv.FormatInt(result.ID, 10))
	data.ClusterTypeID = types.Int64Value(result.ID)

	// Save data into Terraform state
	resp.Diagnostics.Append(resp.State.Set(ctx, &data)...)
}

// Configure adds the provider configured client to the data source.
func (d *clusterTypeDataSource) Configure(_ context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	client, ok := req.ProviderData.(*netboxclient.NetBoxAPI)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Data Source Configure Type",
			fmt.Sprintf("Expected *netboxclient.NetBoxAPI, got: %T. Please report this issue to the provider developers.", req.ProviderData),
		)

		return
	}

	d.client = client
}
