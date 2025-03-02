package provider

import (
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/netbox-community/go-netbox/v4"
)

func testClient(client *netbox.APIClient) diag.Diagnostics {
	var diags = diag.Diagnostics{}
	//Validate that the APIClient exist.
	if client == nil {
		diags.AddError(
			"Create: Unconfigured API Client",
			"Expected configured API Client. Please report this issue to the provider developers.",
		)
		return nil
	}
	return diags
}
