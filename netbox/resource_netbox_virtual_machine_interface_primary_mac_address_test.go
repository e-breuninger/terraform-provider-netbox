package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVirtualMachineInterfacePrimaryMACAddressFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id         = netbox_site.test.id
}

resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.test.id
  site_id    = netbox_site.test.id
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_vlan" "test_untagged" {
  name = "%[1]s-untagged"
  vid  = 100
}

resource "netbox_vlan" "test_tagged_200" {
  name = "%[1]s-tagged-200"
  vid  = 200
}

resource "netbox_vlan" "test_tagged_300" {
  name = "%[1]s-tagged-300"
  vid  = 300
}

resource "netbox_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name        = "%[1]s"
  description = "test interface"
  enabled     = true
  mode        = "access"
  mtu         = 1500
  untagged_vlan = netbox_vlan.test_untagged.id
  tagged_vlans  = [
    netbox_vlan.test_tagged_200.id,
    netbox_vlan.test_tagged_300.id,
  ]
  tags = [netbox_tag.test.name]
}

resource "netbox_mac_address" "test" {
  mac_address  = "00:1A:2B:3C:4D:5E"
  virtual_machine_interface_id = netbox_interface.test.id
}

resource "netbox_virtual_machine_interface_primary_mac_address" "test" {
  interface_id = netbox_interface.test.id
  mac_address_id = netbox_mac_address.test.id
}
`, testName)
}

func TestAccNetboxVirtualMachineInterfacePrimaryMACAddress_basic(t *testing.T) {
	testSlug := "pr_mac_basic"
	testName := testAccGetTestName(testSlug)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualMachineInterfacePrimaryMACAddressFullDependencies(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface_primary_mac_address.test", "interface_id", "netbox_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "untagged_vlan", "netbox_vlan.test_untagged", "id"),
					resource.TestCheckResourceAttr("netbox_interface.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_interface.test", "description", "test interface"),
					resource.TestCheckResourceAttr("netbox_interface.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_interface.test", "mode", "access"),
					resource.TestCheckResourceAttr("netbox_interface.test", "mtu", "1500"),
					resource.TestCheckResourceAttr("netbox_interface.test", "tagged_vlans.#", "2"),
					resource.TestCheckResourceAttr("netbox_interface.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_interface.test", "tags.0", testName),
				),
			},
			// Second apply: no config change, just refresh state and check that the interface now has the primary mac address assigned
			{
				Config: testAccNetboxVirtualMachineInterfacePrimaryMACAddressFullDependencies(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface_primary_mac_address.test", "interface_id", "netbox_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "primary_mac_address_id", "netbox_virtual_machine_interface_primary_mac_address.test", "mac_address_id"),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "mac_address", "netbox_mac_address.test", "mac_address"),
				),
			},
		},
	})
}
