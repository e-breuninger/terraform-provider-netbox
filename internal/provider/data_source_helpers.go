package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-framework/datasource"
)

type NetboxDataSource struct {
	datasource.DataSource
	provider *netboxProvider
}

func (d *NetboxDataSource) Configure(ctx context.Context, req datasource.ConfigureRequest, resp *datasource.ConfigureResponse) {
	if req.ProviderData == nil {
		return
	}

	provider, ok := req.ProviderData.(*netboxProvider)
	if !ok {
		resp.Diagnostics.AddError(
			"Unexpected Resource Configure Type",
			fmt.Sprintf("Expected *netboxProvider, got: %T, Please report this issue to the provider developers.", req.ProviderData),
		)
		return
	}

	d.provider = provider
}
