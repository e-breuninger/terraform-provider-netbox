//go:build acctest

// The primary MAC address link resources: an interface's primary_mac_address is set after the
// address exists, breaking the interface -> mac_address -> interface cycle. No sweeper: the link
// is a field on the interface, which goes with its device or virtual machine.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func randomMAC() string {
	return fmt.Sprintf("02:42:%02x:%02x:%02x:%02x",
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255),
		acctest.RandIntRange(0, 255), acctest.RandIntRange(0, 255))
}

func TestAccNetboxDeviceInterfacePrimaryMACAddress_basic(t *testing.T) {
	testName := testAccGetTestName("ifprimarymac")
	mac1, mac2 := randomMAC(), randomMAC()
	deps := componentDeps(testName, "") + fmt.Sprintf(`
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-eth0"
  type      = "virtual"
}
resource "netbox_mac_address" "one" {
  mac_address         = "%[2]s"
  device_interface_id = netbox_device_interface.test.id
}
resource "netbox_mac_address" "two" {
  mac_address         = "%[3]s"
  device_interface_id = netbox_device_interface.test.id
}`, testName, mac1, mac2)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_device_interface_primary_mac_address" "test" {
  device_interface_id = netbox_device_interface.test.id
  mac_address_id      = netbox_mac_address.one.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.one", "id"),
				),
			},
			{
				// Move the primary to the other address in place.
				Config: deps + `
resource "netbox_device_interface_primary_mac_address" "test" {
  device_interface_id = netbox_device_interface.test.id
  mac_address_id      = netbox_mac_address.two.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.two", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_interface_primary_mac_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVirtualMachineInterfacePrimaryMACAddress_basic(t *testing.T) {
	testName := testAccGetTestName("vmifprimarymac")
	mac := randomMAC()
	deps := fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.test.id
}
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%[1]s"
}
resource "netbox_mac_address" "test" {
  mac_address                  = "%[2]s"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}`, testName, mac)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_virtual_machine_interface_primary_mac_address" "test" {
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
  mac_address_id               = netbox_mac_address.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface_primary_mac_address.test", "id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_virtual_machine_interface_primary_mac_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
