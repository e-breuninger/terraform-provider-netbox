package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxPrimaryIPFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id = netbox_site.test.id
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_gb = 256
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  platform_id = netbox_platform.test.id
  vcpus = "4"

  tags = [netbox_tag.test.name]
}

resource "netbox_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name = "%[1]s"
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

					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "memory_mb", "1024"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "4"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "disk_size_gb", "256"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "tags.0", testName),
				),
			},
		},
	})
}
