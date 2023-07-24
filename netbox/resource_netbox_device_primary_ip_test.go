package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxDevicePrimaryIPFullDependencies(testName string) string {
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

resource "netbox_location" "test" {
  name = "%[1]s"
  site_id =netbox_site.test.id
}

resource "netbox_rack" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
  status = "reserved"
  width = 19
  u_height = 48
  tenant_id = netbox_tenant.test.id
  location_id = netbox_location.test.id
}

resource "netbox_device" "test" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  tenant_id = netbox_tenant.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "thisisacomment"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 10

  tags = [netbox_tag.test.name]
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "1000base-t"
}

resource "netbox_ip_address" "test_v4" {
  ip_address = "1.1.1.1/32"
  status = "active"
  interface_id = netbox_device_interface.test.id
  object_type = "dcim.interface"
}

resource "netbox_ip_address" "test_v6" {
  ip_address = "2000::1/128"
  status = "active"
  interface_id = netbox_device_interface.test.id
  object_type = "dcim.interface"
}
`, testName)
}

func TestAccNetboxDevicePrimaryIP4_basic(t *testing.T) {
	testSlug := "pr_ip_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + `
resource "netbox_device_primary_ip" "test_v4" {
  device_id = netbox_device.test.id
  ip_address_id = netbox_ip_address.test_v4.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4", "ip_address_id", "netbox_ip_address.test_v4", "id"),

					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "planned"),
				),
			},
		},
	})
}

func TestAccNetboxDevicePrimaryIP6_basic(t *testing.T) {
	testSlug := "pr_ip_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + `
resource "netbox_device_primary_ip" "test_v6" {
  device_id = netbox_device.test.id
  ip_address_id = netbox_ip_address.test_v6.id
  ip_address_version = 6
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v6", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v6", "ip_address_id", "netbox_ip_address.test_v6", "id"),

					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "planned"),
				),
			},
		},
	})
}
