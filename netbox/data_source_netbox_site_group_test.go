package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxSiteGroupDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("sitegrp_ds_basic")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "test" {
  name = "%[1]s"
  description = "foo"
}

data "netbox_site_group" "by_name" {
  depends_on = [netbox_site_group.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_site_group.by_name", "id", "netbox_site_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_site_group.by_name", "name", "netbox_site_group.test", "name"),
					resource.TestCheckResourceAttr("data.netbox_site_group.by_name", "description", "foo"),
				),
			},
		},
	})
}
