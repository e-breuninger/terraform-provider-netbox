//go:build acctest

// Cables: both forms of a side (object_type plus ids, and the per-type alias lists), a breakout,
// and the plan-time conflicts.
package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxCable_basic(t *testing.T) {
	testName := testAccGetTestName("cable")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "a" {
  name           = "%[1]s-a"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device" "b" {
  name           = "%[1]s-b"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "a0" {
  device_id = netbox_device.a.id
  name      = "eth0"
  type      = "1000base-t"
}
resource "netbox_device_interface" "a1" {
  device_id = netbox_device.a.id
  name      = "eth1"
  type      = "1000base-t"
}
resource "netbox_device_interface" "b0" {
  device_id = netbox_device.b.id
  name      = "eth0"
  type      = "1000base-t"
}
resource "netbox_device_interface" "b1" {
  device_id = netbox_device.b.id
  name      = "eth1"
  type      = "1000base-t"
}
resource "netbox_device_interface" "b2" {
  device_id = netbox_device.b.id
  name      = "eth2"
  type      = "1000base-t"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_cable_bundle" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// The alias form: one interface on each side.
				Config: deps + fmt.Sprintf(`
resource "netbox_cable" "test" {
  a_side      = { device_interface_ids = [netbox_device_interface.a0.id] }
  b_side      = { device_interface_ids = [netbox_device_interface.b0.id] }
  type        = "cat6"
  status      = "planned"
  label       = "%[1]s"
  color_hex   = "ff0000"
  length      = 3
  length_unit = "m"
  tenant_id   = netbox_tenant.test.id
  bundle_id   = netbox_cable_bundle.test.id
  description = "Acceptance test cable."
  comments    = "Created by acceptance test."
}
data "netbox_cable" "test" {
  label = netbox_cable.test.label
}
data "netbox_cables" "test" {
  filters = [
    { name = "tenant_id", value = netbox_tenant.test.id },
  ]
  depends_on = [netbox_cable.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cable.test", "a_side.object_type", "dcim.interface"),
					resource.TestCheckResourceAttr("netbox_cable.test", "a_side.ids.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_cable.test", "a_side.ids.0", "netbox_device_interface.a0", "id"),
					resource.TestCheckResourceAttrPair("netbox_cable.test", "b_side.device_interface_ids.0", "netbox_device_interface.b0", "id"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "a_side.front_port_ids"),
					resource.TestCheckResourceAttr("netbox_cable.test", "type", "cat6"),
					resource.TestCheckResourceAttr("netbox_cable.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_cable.test", "length", "3"),
					resource.TestCheckResourceAttr("netbox_cable.test", "length_unit", "m"),
					resource.TestCheckResourceAttrPair("netbox_cable.test", "bundle_id", "netbox_cable_bundle.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cable.test", "id", "netbox_cable.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cable.test", "a_side.device_interface_ids.0", "netbox_device_interface.a0", "id"),
					resource.TestCheckResourceAttr("data.netbox_cables.test", "cables.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_cable.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// The explicit form, re-terminated as a breakout: one interface to two, with the
				// matching profile.
				Config: deps + fmt.Sprintf(`
resource "netbox_cable" "test" {
  a_side = { object_type = "dcim.interface", ids = [netbox_device_interface.a1.id] }
  b_side = {
    object_type = "dcim.interface"
    ids         = [netbox_device_interface.b1.id, netbox_device_interface.b2.id]
  }
  profile = "breakout-1c2p-2c1p"
  label   = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cable.test", "b_side.ids.#", "2"),
					resource.TestCheckResourceAttr("netbox_cable.test", "b_side.device_interface_ids.#", "2"),
					resource.TestCheckResourceAttrPair("netbox_cable.test", "b_side.device_interface_ids.1", "netbox_device_interface.b2", "id"),
					resource.TestCheckResourceAttrPair("netbox_cable.test", "a_side.device_interface_ids.0", "netbox_device_interface.a1", "id"),
					resource.TestCheckResourceAttr("netbox_cable.test", "profile", "breakout-1c2p-2c1p"),
				),
			},
			{
				// Shrink: the terminations stay (a cable cannot lose them); status and profile are
				// optional+computed and keep NetBox's value (clearing a profile is rejected by
				// NetBox anyway), the rest clears.
				Config: deps + `
resource "netbox_cable" "test" {
  a_side = { device_interface_ids = [netbox_device_interface.a1.id] }
  b_side = { device_interface_ids = [netbox_device_interface.b1.id, netbox_device_interface.b2.id] }
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cable.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_cable.test", "profile", "breakout-1c2p-2c1p"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "color_hex"),
					resource.TestCheckResourceAttr("netbox_cable.test", "b_side.ids.#", "2"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "type"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "length"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "length_unit"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "bundle_id"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_cable.test", "comments"),
				),
			},
			{
				// An alias next to object_type/ids is a plan-time error.
				Config: deps + `
resource "netbox_cable" "test" {
  a_side = {
    object_type          = "dcim.interface"
    ids                  = [netbox_device_interface.a1.id]
    device_interface_ids = [netbox_device_interface.a1.id]
  }
  b_side = { device_interface_ids = [netbox_device_interface.b1.id] }
}`,
				ExpectError: regexp.MustCompile("Conflicting terminations"),
			},
			{
				// So is a side without terminations.
				Config: deps + `
resource "netbox_cable" "test" {
  a_side = { device_interface_ids = [netbox_device_interface.a1.id] }
  b_side = {}
}`,
				ExpectError: regexp.MustCompile("Missing terminations"),
			},
		},
	})
}
