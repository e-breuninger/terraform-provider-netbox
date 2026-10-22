package provider

import (
	"context"
	"fmt"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/datasource/schema"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
)

// netbox_available_prefix is a hand-written companion data source: it lists the free child
// prefixes of a prefix through NetBox's available-prefixes endpoint, which has no object of its own
// for the spec to describe. Reading allocates nothing; the netbox_available_prefix resource does.

func init() {
	companionDataSources = append(companionDataSources, NewAvailablePrefixDataSource)
}

var (
	_ datasource.DataSource              = (*availablePrefixDataSource)(nil)
	_ datasource.DataSourceWithConfigure = (*availablePrefixDataSource)(nil)
)

func NewAvailablePrefixDataSource() datasource.DataSource {
	return &availablePrefixDataSource{}
}

type availablePrefixDataSource struct {
	client *netboxapi.Client
}

type availablePrefixDataSourceModel struct {
	PrefixID          types.Int64 `tfsdk:"prefix_id"`
	AvailablePrefixes types.List  `tfsdk:"available_prefixes"`
}

// availablePrefixEntry is one free block of available_prefixes.
type availablePrefixEntry struct {
	Family types.Int64  `tfsdk:"family"`
	Prefix types.String `tfsdk:"prefix"`
	VrfID  types.Int64  `tfsdk:"vrf_id"`
}

var availablePrefixEntryType = types.ObjectType{AttrTypes: map[string]attr.Type{
	"family": types.Int64Type,
	"prefix": types.StringType,
	"vrf_id": types.Int64Type,
}}

func (dataSource *availablePrefixDataSource) Metadata(ctx context.Context, req datasource.MetadataRequest, resp *datasource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_available_prefix"
}

func (dataSource *availablePrefixDataSource) Schema(ctx context.Context, req datasource.SchemaRequest, resp *datasource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: ":meta:subcategory:IP Address Management (IPAM):The free child prefixes of a NetBox prefix (ipam.prefix): one entry per contiguous free block, as NetBox's available-prefixes endpoint lists them. Reading allocates nothing; the netbox_available_prefix resource allocates.",
		Attributes: map[string]schema.Attribute{
			"prefix_id": schema.Int64Attribute{
				Required:    true,
				Description: "Id of the prefix whose free space is listed.",
			},
			"available_prefixes": schema.ListNestedAttribute{
				Computed:    true,
				Description: "The free blocks, in address order.",
				NestedObject: schema.NestedAttributeObject{
					Attributes: map[string]schema.Attribute{
						"family": schema.Int64Attribute{
							Computed:    true,
							Description: "IP family of the block, 4 or 6.",
						},
						"prefix": schema.StringAttribute{
							Computed:    true,
							Description: "The free block in CIDR notation, e.g. 10.0.1.0/24.",
						},
						"vrf_id": schema.Int64Attribute{
							Computed:    true,
							Description: "Id of the VRF the block belongs to; unset in the global table.",
						},
					},
				},
			},
		},
	}
}

func (dataSource *availablePrefixDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}
	client, ok := req.ProviderData.(*netboxapi.Client)
	if !ok {
		resp.Diagnostics.AddError("Unexpected data source configure type",
			fmt.Sprintf("Expected *netboxapi.Client, got: %T. Please report this issue to the provider developers.", req.ProviderData))
		return
	}
	dataSource.client = client
}

func (dataSource *availablePrefixDataSource) Read(ctx context.Context, req datasource.ReadRequest, resp *datasource.ReadResponse) {
	var model availablePrefixDataSourceModel
	resp.Diagnostics.Append(req.Config.Get(ctx, &model)...)
	if resp.Diagnostics.HasError() {
		return
	}
	prefixID := model.PrefixID.ValueInt64()
	res, err := dataSource.client.Ipam.IpamPrefixesAvailablePrefixesListContext(ctx,
		ipam.NewIpamPrefixesAvailablePrefixesListParams().WithID(prefixID), nil)
	if err != nil {
		resp.Diagnostics.AddError(fmt.Sprintf("Error listing the available prefixes of prefix %d", prefixID), err.Error())
		return
	}
	entries := make([]availablePrefixEntry, 0, len(res.Payload))
	for _, availablePrefix := range res.Payload {
		entry := availablePrefixEntry{
			Family: types.Int64Value(availablePrefix.Family),
			Prefix: types.StringValue(availablePrefix.Prefix),
			VrfID:  types.Int64Null(),
		}
		if availablePrefix.Vrf != nil {
			entry.VrfID = types.Int64Value(availablePrefix.Vrf.ID)
		}
		entries = append(entries, entry)
	}
	availablePrefixes, diags := types.ListValueFrom(ctx, availablePrefixEntryType, entries)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}
	model.AvailablePrefixes = availablePrefixes
	resp.Diagnostics.Append(resp.State.Set(ctx, &model)...)
}
