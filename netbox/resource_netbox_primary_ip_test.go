package netbox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func testAccNetboxPrimaryIPFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name = "%[1]s"
  type = "virtual"
}

resource "netbox_ip_address" "test" {
  ip_address = "1.1.1.1/32"
  status = "active"
  interface_id = netbox_interface.test.id
}
`, testName)
}

func TestAccNetboxPrimaryIP_basic(t *testing.T) {

	testSlug := "pr_ip_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxPrimaryIPFullDependencies(testName) + `
resource "netbox_primary_ip" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  ip_address_id = netbox_ip_address.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test", "ip_address_id", "netbox_ip_address.test", "id"),
				),
			},
		},
	})
}
