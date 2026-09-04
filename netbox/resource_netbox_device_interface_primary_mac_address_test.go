package netbox

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

func testAccNetboxDeviceInterfacePrimaryMACAddressUpdatableInterface(testName, description string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
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
  name        = "%[1]s-if0"
  device_id   = netbox_device.test.id
  type        = "1000base-t"
  description = "%[2]s"
}

resource "netbox_mac_address" "test" {
  mac_address         = "00:1A:2B:3C:4D:5E"
  device_interface_id = netbox_device_interface.test.id
}

resource "netbox_device_interface_primary_mac_address" "test" {
  interface_id   = netbox_device_interface.test.id
  mac_address_id = netbox_mac_address.test.id
}
`, testName, description)
}

// testAccCheckDeviceInterfacePrimaryMACAddressSet asserts NetBox itself still
// reports the expected primary MAC on the interface. It deliberately does not
// go through Terraform state: netbox_device_interface's Update does not re-read
// after PATCHing, so a primary MAC cleared as a side effect of that PATCH stays
// in state until the next refresh.
func testAccCheckDeviceInterfacePrimaryMACAddressSet(interfaceResourceName, macResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		api := testAccProvider.Meta().(*providerState)

		interfaceState, ok := s.RootModule().Resources[interfaceResourceName]
		if !ok {
			return fmt.Errorf("%s not found in state", interfaceResourceName)
		}
		macState, ok := s.RootModule().Resources[macResourceName]
		if !ok {
			return fmt.Errorf("%s not found in state", macResourceName)
		}

		interfaceID, err := strconv.ParseInt(interfaceState.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		res, err := api.Dcim.DcimInterfacesRead(dcim.NewDcimInterfacesReadParams().WithID(interfaceID), nil)
		if err != nil {
			return err
		}

		primaryMAC := res.GetPayload().PrimaryMacAddress
		if primaryMAC == nil {
			return fmt.Errorf("interface %d has no primary MAC address in Netbox, expected MAC address %s", interfaceID, macState.Primary.ID)
		}
		if got, want := strconv.FormatInt(primaryMAC.ID, 10), macState.Primary.ID; got != want {
			return fmt.Errorf("interface %d has primary MAC address %s in Netbox, expected %s", interfaceID, got, want)
		}
		return nil
	}
}

// TestAccNetboxDeviceInterfacePrimaryMACAddress_survivesInterfaceUpdate covers
// the primary MAC surviving an unrelated in-place update of the interface it is
// attached to. WritableInterface.PrimaryMacAddress has no `omitempty`, so any
// PATCH that leaves it nil sends `"primary_mac_address": null` and Netbox drops
// the primary MAC as collateral damage.
func TestAccNetboxDeviceInterfacePrimaryMACAddress_survivesInterfaceUpdate(t *testing.T) {
	testSlug := "pr_device_mac_upd"
	testName := testAccGetTestName(testSlug)

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDeviceInterfacePrimaryMACAddressUpdatableInterface(testName, "before update"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_interface_primary_mac_address.test", "mac_address_id", "netbox_mac_address.test", "id"),
					testAccCheckDeviceInterfacePrimaryMACAddressSet("netbox_device_interface.test", "netbox_mac_address.test"),
				),
			},
			{
				// Only the interface's description changes; nothing about the
				// primary MAC does. Besides the explicit check, the test
				// framework's own post-apply plan must stay empty: once Netbox
				// has dropped the primary MAC, the binding resource's Read
				// clears its id and the next plan wants to re-create it.
				Config: testAccNetboxDeviceInterfacePrimaryMACAddressUpdatableInterface(testName, "after update"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "description", "after update"),
					testAccCheckDeviceInterfacePrimaryMACAddressSet("netbox_device_interface.test", "netbox_mac_address.test"),
				),
			},
		},
	})
}

// TestAccNetboxDeviceInterfacePrimaryMACAddress_survivesDeletedMAC is the other
// half of carrying the primary MAC through an update: the value comes from
// state, so a MAC address deleted out of band must not be sent back. Netbox
// rejects the dangling id with "Related object not found", which would fail the
// update of an interface that has nothing to do with MAC addresses.
func TestAccNetboxDeviceInterfacePrimaryMACAddress_survivesDeletedMAC(t *testing.T) {
	testSlug := "pr_device_mac_gone"
	testName := testAccGetTestName(testSlug)

	var macID int64

	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDeviceInterfacePrimaryMACAddressUpdatableInterface(testName, "before update"),
				Check: func(s *terraform.State) (err error) {
					rs, ok := s.RootModule().Resources["netbox_mac_address.test"]
					if !ok {
						return fmt.Errorf("netbox_mac_address.test not found in state")
					}
					macID, err = strconv.ParseInt(rs.Primary.ID, 10, 64)
					return err
				},
			},
			{
				// Delete the MAC address behind Terraform's back: its id stays in
				// state, and the interface loses its primary MAC.
				PreConfig: func() {
					api := testAccProvider.Meta().(*providerState)
					_, err := api.Dcim.DcimMacAddressesDelete(dcim.NewDcimMacAddressesDeleteParams().WithID(macID), nil)
					if err != nil {
						t.Fatalf("deleting MAC address %d: %s", macID, err)
					}
				},
				// The description change forces an interface update in the same
				// apply that re-creates the MAC and the binding.
				Config: testAccNetboxDeviceInterfacePrimaryMACAddressUpdatableInterface(testName, "after update"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "description", "after update"),
					testAccCheckDeviceInterfacePrimaryMACAddressSet("netbox_device_interface.test", "netbox_mac_address.test"),
				),
			},
		},
	})
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
