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

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_gb = 256
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  platform_id = netbox_platform.test.id
  vcpus = "4"
  status = "planned"
  device_id = netbox_device.test.id
  local_context_data = jsonencode({"context_string"="context_value"})

  tags = [netbox_tag.test.name]
}

resource "netbox_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name = "%[1]s"
}

resource "netbox_ip_address" "test_v4" {
  ip_address = "1.1.1.1/32"
  status = "active"
  virtual_machine_interface_id = netbox_interface.test.id
}

resource "netbox_ip_address" "test_v6" {
  ip_address = "2000::1/128"
  status = "active"
  virtual_machine_interface_id = netbox_interface.test.id
}
`, testName)
}

func TestAccNetboxPrimaryIP4_basic(t *testing.T) {
	testSlug := "pr_ip_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxPrimaryIPFullDependencies(testName) + `
resource "netbox_primary_ip" "test_v4" {
  virtual_machine_id = netbox_virtual_machine.test.id
  ip_address_id = netbox_ip_address.test_v4.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test_v4", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test_v4", "ip_address_id", "netbox_ip_address.test_v4", "id"),

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
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "local_context_data", "{\"context_string\":\"context_value\"}"),
				),
			},
		},
	})
}

func TestAccNetboxPrimaryIP6_basic(t *testing.T) {
	testSlug := "pr_ip_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxPrimaryIPFullDependencies(testName) + `
resource "netbox_primary_ip" "test_v6" {
  virtual_machine_id = netbox_virtual_machine.test.id
  ip_address_id = netbox_ip_address.test_v6.id
  ip_address_version = 6
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test_v6", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test_v6", "ip_address_id", "netbox_ip_address.test_v6", "id"),

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
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "status", "planned"),
				),
			},
		},
	})
}
