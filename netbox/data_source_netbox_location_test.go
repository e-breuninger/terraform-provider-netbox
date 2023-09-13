package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxLocationDataSource_basic(t *testing.T) {
	testSlug := "location_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  description = "my-description"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
}

data "netbox_location" "by_name" {
  name = netbox_location.test.name
}

data "netbox_location" "by_id" {
  id = netbox_location.test.id
}

data "netbox_location" "by_slug" {
  slug = netbox_location.test.slug
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_id", "id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_slug", "id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_location.by_name", "name", testName),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "description", "netbox_location.test", "description"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "site_id", "netbox_location.test", "site_id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "tenant_id", "netbox_location.test", "tenant_id"),
				),
			},
		},
	})
}
