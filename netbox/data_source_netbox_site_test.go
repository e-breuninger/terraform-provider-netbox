package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxSiteDataSource_basic(t *testing.T) {

	testSlug := "site_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}
data "netbox_site" "test" {
  depends_on = [netbox_site.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
                                        resource.TestCheckResourceAttrPair("data.netbox_site.test", "status", "netbox_site.test", "active"),
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
				),
			},
		},
	})
}
