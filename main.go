package main

import (
	"github.com/e-breuninger/terraform-provider-netbox/netbox"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/plugin"
)

func main() {
	plugin.Serve(&plugin.ServeOpts{
		ProviderFunc: func() *schema.Provider {
			return netbox.Provider()
		},
	})
}
