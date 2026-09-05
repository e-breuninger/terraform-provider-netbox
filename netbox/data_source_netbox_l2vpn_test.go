package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxL2VPNDataSource_basic(t *testing.T) {
	testSlug := "l2vpn_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxL2VPNDependencies(testName) + fmt.Sprintf(`
resource "netbox_l2vpn" "test" {
  name           = "%[1]s"
  slug           = "%[1]s"
  type           = "vpws"
  identifier     = 1234
  tenant_id      = netbox_tenant.test.id
  description    = "test description"
  comments       = "test comments"
  import_targets = [netbox_route_target.import.id]
  export_targets = [netbox_route_target.export.id]
  tags           = [netbox_tag.test.name]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_l2vpn" "test" {
  depends_on = []

  name = netbox_l2vpn.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_l2vpn.test", "id", "netbox_l2vpn.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "slug", testName),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "type", "vpws"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "identifier", "1234"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "description", "test description"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "comments", "test comments"),
					resource.TestCheckResourceAttrPair("data.netbox_l2vpn.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "import_targets.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "export_targets.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn.test", "tags.0", testName),
				),
			},
		},
	})
}
