package main

import (
	"context"

	"github.com/e-breuninger/terraform-provider-netbox/pkg/client"
)

func main() {
	c, _ := client.NewClient("https://demo.netbox.dev")

	region := client.WritableSiteRequest_Region{}
	region.FromForeignID(1)

	tenant := client.WritableSiteRequest_Tenant{}
	tenant.FromForeignID(2)

	siteRequest := client.WritableSiteRequest{
		Name:   "Test Site",
		Region: &region,
		Slug:   "test-site",
		Tenant: &tenant,
	}

	_, _ = c.DcimSitesCreate(context.TODO(), siteRequest)
}
