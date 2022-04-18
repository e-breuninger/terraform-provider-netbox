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

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  device_role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "lag"
}

resource "netbox_ip_address" "test_v4" {
  ip_address = "1.1.1.3/32"
  status = "active"
  interface_id = netbox_device_interface.test.id
  interface_type = "dcim.interface"
}

resource "netbox_ip_address" "test_v6" {
  ip_address = "2000::3/128"
  status = "active"
  interface_id = netbox_device_interface.test.id
  interface_type = "dcim.interface"
}
`, testName)
}

func TestAccNetboxDevicePrimaryIP4_basic(t *testing.T) {

	testSlug := "pr_ip4_basic"
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
  ip_address_version = 4
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4", "ip_address_id", "netbox_ip_address.test_v4", "id"),

					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "device_role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName),
				),
			},
		},
	})
}

func TestAccNetboxDevicePrimaryIP6_basic(t *testing.T) {

	testSlug := "pr_ip6_basic"
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
					resource.TestCheckResourceAttrPair("netbox_device.test", "device_role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName),
				),
			},
		},
	})
}
