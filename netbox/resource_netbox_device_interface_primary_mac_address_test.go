package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxDeviceInterfacePrimaryMACAddressFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_vlan" "untagged" {
  name = "%[1]s-untagged"
  vid  = 100
}

resource "netbox_vlan" "tagged_200" {
  name = "%[1]s-tagged-200"
  vid  = 200
}

resource "netbox_vlan" "tagged_300" {
  name = "%[1]s-tagged-300"
  vid  = 300
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s-manufacturer"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s-model"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_interface" "test" {
  name      = "%[1]s-if0"
  device_id = netbox_device.test.id
  type      = "1000base-t"
  description = "primary MAC test interface"
  enabled     = true
  mode = "tagged"
  mtu  = 1500
  untagged_vlan = netbox_vlan.untagged.id
  tagged_vlans  = [
    netbox_vlan.tagged_200.id,
    netbox_vlan.tagged_300.id,
  ]
  tags = [netbox_tag.test.name]
}

resource "netbox_mac_address" "test" {
  mac_address         = "00:1A:2B:3C:4D:5E"
  device_interface_id = netbox_device_interface.test.id
}

resource "netbox_device_interface_primary_mac_address" "test" {
  interface_id = netbox_device_interface.test.id
  mac_address_id      = netbox_mac_address.test.id
}
`, testName)
}

func TestAccNetboxDeviceInterfacePrimaryMACAddress_basic(t *testing.T) {
	testSlug := "pr_device_mac_basic"
	testName := testAccGetTestName(testSlug)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDeviceInterfacePrimaryMACAddressFullDependencies(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "untagged_vlan", "netbox_vlan.untagged", "id"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "name", testName+"-if0"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "type", "1000base-t"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "description", "primary MAC test interface"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mode", "tagged"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mtu", "1500"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "tagged_vlans.#", "2"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "tags.0", testName),
				),
			},
			// Second apply: no config change, just refresh state and check that the interface now has the primary mac address assigned
			{
				Config: testAccNetboxDeviceInterfacePrimaryMACAddressFullDependencies(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "primary_mac_address_id", "netbox_device_interface_primary_mac_address.test", "mac_address_id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "mac_address", "netbox_mac_address.test", "mac_address"),
				),
			},
		},
	})
}
