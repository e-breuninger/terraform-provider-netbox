//go:build acctest

// Device components: bays, console/power/rear ports, module bays and modules, inventory items
// and their roles, and virtual chassis.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// componentDeps is the catalog a device component needs: the device catalog plus a device. The
// device type's subdevice_role is caller-chosen because NetBox only accepts device bays on a
// device whose type is a "parent".
func componentDeps(testName, subdeviceRole string) string {
	role := ""
	if subdeviceRole != "" {
		role = fmt.Sprintf("\n  subdevice_role  = %q", subdeviceRole)
	}
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"%[2]s
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}`, testName, role)
}

// TestAccNetboxDevicePorts_basic covers the port resources on one device: console, console server,
// power port, power outlet and rear port share the device/module/name/label/mark_connected shape,
// so one device exercises them all; the front port has its own test below.
func TestAccNetboxDevicePorts_basic(t *testing.T) {
	testName := testAccGetTestName("ports")
	deps := componentDeps(testName, "")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_console_port" "test" {
  device_id      = netbox_device.test.id
  name           = "%[1]s-con"
  type           = "rj-45"
  speed          = 9600
  label          = "console"
  mark_connected = true
  description    = "Acceptance test console port."
}
resource "netbox_device_console_server_port" "test" {
  device_id      = netbox_device.test.id
  name           = "%[1]s-csp"
  type           = "rj-45"
  speed          = 115200
  label          = "console-server"
  mark_connected = true
  description    = "Acceptance test console server port."
}
resource "netbox_device_power_port" "test" {
  device_id      = netbox_device.test.id
  name           = "%[1]s-pp"
  type           = "iec-60320-c14"
  maximum_draw   = 100
  allocated_draw = 50
  label          = "psu"
  mark_connected = true
  description    = "Acceptance test power port."
}
resource "netbox_device_power_outlet" "test" {
  device_id      = netbox_device.test.id
  name           = "%[1]s-po"
  type           = "iec-60320-c13"
  power_port_id  = netbox_device_power_port.test.id
  feed_leg       = "A"
  label          = "outlet"
  status         = "disabled"
  color_hex      = "ff0000"
  mark_connected = true
  description    = "Acceptance test power outlet."
}
resource "netbox_device_rear_port" "test" {
  device_id      = netbox_device.test.id
  name           = "%[1]s-rp"
  type           = "8p8c"
  positions      = 4
  color_hex      = "ff0000"
  label          = "rear"
  mark_connected = true
  description    = "Acceptance test rear port."
}
data "netbox_device_console_port" "test" {
  name       = netbox_device_console_port.test.name
  depends_on = [netbox_device_console_port.test]
}
data "netbox_device_rear_ports" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_rear_port.test]
}
data "netbox_device_console_ports" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_console_port.test]
}
data "netbox_device_console_server_port" "test" {
  id = netbox_device_console_server_port.test.id
}
data "netbox_device_console_server_ports" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_console_server_port.test]
}
data "netbox_device_power_port" "test" {
  id = netbox_device_power_port.test.id
}
data "netbox_device_power_ports" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_power_port.test]
}
data "netbox_device_power_outlet" "test" {
  id = netbox_device_power_outlet.test.id
}
data "netbox_device_power_outlets" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_power_outlet.test]
}
data "netbox_device_rear_port" "test" {
  id = netbox_device_rear_port.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_console_port.test", "type", "rj-45"),
					resource.TestCheckResourceAttr("netbox_device_console_port.test", "speed", "9600"),
					resource.TestCheckResourceAttr("netbox_device_console_port.test", "label", "console"),
					resource.TestCheckResourceAttr("netbox_device_console_port.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_device_console_server_port.test", "speed", "115200"),
					resource.TestCheckResourceAttr("netbox_device_console_server_port.test", "label", "console-server"),
					resource.TestCheckResourceAttr("netbox_device_console_server_port.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "maximum_draw", "100"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "allocated_draw", "50"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_device_power_outlet.test", "feed_leg", "A"),
					resource.TestCheckResourceAttr("netbox_device_power_outlet.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_device_rear_port.test", "label", "rear"),
					resource.TestCheckResourceAttr("netbox_device_rear_port.test", "mark_connected", "true"),
					resource.TestCheckResourceAttrPair("netbox_device_power_outlet.test", "power_port_id",
						"netbox_device_power_port.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_rear_port.test", "positions", "4"),
					resource.TestCheckResourceAttr("netbox_device_rear_port.test", "color_hex", "ff0000"),
					resource.TestCheckResourceAttrPair("data.netbox_device_console_port.test", "id",
						"netbox_device_console_port.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_ports.test", "device_rear_ports.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_console_ports.test", "device_console_ports.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_device_console_server_port.test", "id",
						"netbox_device_console_server_port.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_console_server_ports.test",
						"device_console_server_ports.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_device_power_port.test", "id",
						"netbox_device_power_port.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_power_ports.test", "device_power_ports.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_device_power_outlet.test", "id",
						"netbox_device_power_outlet.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_power_outlets.test", "device_power_outlets.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_device_rear_port.test", "id",
						"netbox_device_rear_port.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_console_port.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_device_console_server_port.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_device_power_port.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_device_power_outlet.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_device_rear_port.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: the optional scalars clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_console_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-con"
  type      = "rj-45"
}
resource "netbox_device_console_server_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-csp"
  type      = "rj-45"
}
resource "netbox_device_power_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-pp"
  type      = "iec-60320-c14"
}
resource "netbox_device_power_outlet" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-po"
  type      = "iec-60320-c13"
}
resource "netbox_device_rear_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-rp"
  type      = "8p8c"
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device_console_port.test", "speed"),
					resource.TestCheckNoResourceAttr("netbox_device_console_port.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_device_console_port.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_device_power_port.test", "maximum_draw"),
					resource.TestCheckNoResourceAttr("netbox_device_power_port.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_device_power_port.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_device_power_outlet.test", "power_port_id"),
					resource.TestCheckNoResourceAttr("netbox_device_power_outlet.test", "feed_leg"),
					resource.TestCheckNoResourceAttr("netbox_device_power_outlet.test", "label"),
					resource.TestCheckResourceAttr("netbox_device_power_outlet.test", "status", "disabled"),
					resource.TestCheckNoResourceAttr("netbox_device_power_outlet.test", "color_hex"),
					resource.TestCheckNoResourceAttr("netbox_device_power_outlet.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_device_console_server_port.test", "speed"),
					resource.TestCheckNoResourceAttr("netbox_device_console_server_port.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_device_console_server_port.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_device_power_port.test", "allocated_draw"),
					resource.TestCheckNoResourceAttr("netbox_device_rear_port.test", "color_hex"),
					resource.TestCheckResourceAttr("netbox_device_rear_port.test", "positions", "1"),
					resource.TestCheckNoResourceAttr("netbox_device_rear_port.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_device_rear_port.test", "description"),
				),
			},
		},
	})
}

func TestAccNetboxDeviceBay_basic(t *testing.T) {
	testName := testAccGetTestName("devbay")
	deps := componentDeps(testName, "parent") + fmt.Sprintf(`
resource "netbox_device_type" "child" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s-child"
  subdevice_role  = "child"
  u_height        = 0
}
resource "netbox_device" "child" {
  name           = "%[1]s-child"
  device_type_id = netbox_device_type.child.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_bay" "test" {
  device_id           = netbox_device.test.id
  name                = "%[1]s-bay"
  label               = "bay 1"
  installed_device_id = netbox_device.child.id
  description         = "Acceptance test device bay."
}
resource "netbox_device_bay" "disabled" {
  device_id = netbox_device.test.id
  name      = "%[1]s-disabled"
  enabled   = false
}
data "netbox_device_bays" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_bay.test, netbox_device_bay.disabled]
}
data "netbox_device_bay" "test" {
  id = netbox_device_bay.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_bay.test", "name", testName+"-bay"),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "label", "bay 1"),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "enabled", "true"),
					// NetBox refuses to install a device in a disabled bay, so another bay is the disabled one.
					resource.TestCheckResourceAttr("netbox_device_bay.disabled", "enabled", "false"),
					resource.TestCheckResourceAttrPair("netbox_device_bay.test", "installed_device_id",
						"netbox_device.child", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_bays.test", "device_bays.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_device_bay.test", "id", "netbox_device_bay.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_bay.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_bay" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-bay"
}
resource "netbox_device_bay" "disabled" {
  device_id = netbox_device.test.id
  name      = "%[1]s-disabled"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device_bay.test", "installed_device_id"),
					resource.TestCheckNoResourceAttr("netbox_device_bay.test", "label"),
					resource.TestCheckResourceAttr("netbox_device_bay.disabled", "enabled", "true"),
					resource.TestCheckNoResourceAttr("netbox_device_bay.test", "description"),
				),
			},
		},
	})
}

func TestAccNetboxModule_basic(t *testing.T) {
	testName := testAccGetTestName("module")
	deps := componentDeps(testName, "") + fmt.Sprintf(`
resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_module_bay" "test" {
  device_id   = netbox_device.test.id
  name        = "%[1]s-mb"
  position    = "1"
  label       = "slot 1"
  description = "Acceptance test module bay."
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status         = "active"
  serial         = "SN-%[1]s"
  asset_tag      = "AT-%[1]s"
  description    = "Acceptance test module."
  comments       = "Created by acceptance test."

  replicate_components = false
  adopt_components     = true
}
resource "netbox_device_module_bay" "nested" {
  device_id = netbox_device.test.id
  module_id = netbox_module.test.id
  name      = "%[1]s-nested"
  enabled   = false
}
resource "netbox_device_interface" "on_module" {
  device_id = netbox_device.test.id
  module_id = netbox_module.test.id
  name      = "%[1]s-eth0"
  type      = "1000base-t"
}
resource "netbox_device_console_port" "on_module" {
  device_id = netbox_device.test.id
  module_id = netbox_module.test.id
  name      = "%[1]s-con"
  type      = "rj-45"
}
resource "netbox_device_console_server_port" "on_module" {
  device_id = netbox_device.test.id
  module_id = netbox_module.test.id
  name      = "%[1]s-csp"
  type      = "rj-45"
}
resource "netbox_device_power_port" "on_module" {
  device_id = netbox_device.test.id
  module_id = netbox_module.test.id
  name      = "%[1]s-pp"
  type      = "iec-60320-c14"
}
resource "netbox_device_power_outlet" "on_module" {
  device_id = netbox_device.test.id
  module_id = netbox_module.test.id
  name      = "%[1]s-po"
  type      = "iec-60320-c13"
}
resource "netbox_device_rear_port" "on_module" {
  device_id = netbox_device.test.id
  module_id = netbox_module.test.id
  name      = "%[1]s-rp"
  type      = "8p8c"
  positions = 2
}
resource "netbox_device_front_port" "on_module" {
  device_id          = netbox_device.test.id
  module_id          = netbox_module.test.id
  name               = "%[1]s-fp"
  type               = "8p8c"
  rear_port_id       = netbox_device_rear_port.on_module.id
  rear_port_position = 1
}
data "netbox_modules" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_module.test]
}
data "netbox_module" "test" {
  id = netbox_module.test.id
}
data "netbox_device_module_bay" "test" {
  id = netbox_device_module_bay.test.id
}
data "netbox_device_module_bays" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_module_bay.test, netbox_device_module_bay.nested]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_module_bay.test", "position", "1"),
					resource.TestCheckResourceAttr("netbox_device_module_bay.test", "label", "slot 1"),
					// NetBox refuses to install a module in a disabled bay, so the nested bay is the disabled one.
					resource.TestCheckResourceAttr("netbox_device_module_bay.nested", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_module.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_module.test", "serial", "SN-"+testName),
					resource.TestCheckResourceAttr("netbox_module.test", "asset_tag", "AT-"+testName),
					resource.TestCheckResourceAttr("netbox_module.test", "replicate_components", "false"),
					resource.TestCheckResourceAttr("netbox_module.test", "adopt_components", "true"),
					resource.TestCheckResourceAttrPair("netbox_module.test", "module_bay_id",
						"netbox_device_module_bay.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_module_bay.nested", "module_id",
						"netbox_module.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_modules.test", "modules.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_module.test", "id", "netbox_module.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_device_module_bay.test", "id",
						"netbox_device_module_bay.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_module_bays.test", "device_module_bays.#", "2"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.on_module", "module_id",
						"netbox_module.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_console_port.on_module", "module_id",
						"netbox_module.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_console_server_port.on_module", "module_id",
						"netbox_module.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_power_port.on_module", "module_id",
						"netbox_module.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_power_outlet.on_module", "module_id",
						"netbox_module.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_rear_port.on_module", "module_id",
						"netbox_module.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_front_port.on_module", "module_id",
						"netbox_module.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_module.test",
				ImportState:       true,
				ImportStateVerify: true,
				// Inputs of the create request that NetBox never reports back.
				ImportStateVerifyIgnore: []string{"replicate_components", "adopt_components"},
			},
			{
				ResourceName:      "netbox_device_module_bay.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: the module bay and the module clear everything optional, and every
				// component moves off the module in place (module clears to null).
				Config: componentDeps(testName, "") + fmt.Sprintf(`
resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-mb"
}
resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
}
resource "netbox_device_module_bay" "nested" {
  device_id = netbox_device.test.id
  name      = "%[1]s-nested"
}
resource "netbox_device_interface" "on_module" {
  device_id = netbox_device.test.id
  name      = "%[1]s-eth0"
  type      = "1000base-t"
}
resource "netbox_device_console_port" "on_module" {
  device_id = netbox_device.test.id
  name      = "%[1]s-con"
  type      = "rj-45"
}
resource "netbox_device_console_server_port" "on_module" {
  device_id = netbox_device.test.id
  name      = "%[1]s-csp"
  type      = "rj-45"
}
resource "netbox_device_power_port" "on_module" {
  device_id = netbox_device.test.id
  name      = "%[1]s-pp"
  type      = "iec-60320-c14"
}
resource "netbox_device_power_outlet" "on_module" {
  device_id = netbox_device.test.id
  name      = "%[1]s-po"
  type      = "iec-60320-c13"
}
resource "netbox_device_rear_port" "on_module" {
  device_id = netbox_device.test.id
  name      = "%[1]s-rp"
  type      = "8p8c"
  positions = 2
}
resource "netbox_device_front_port" "on_module" {
  device_id          = netbox_device.test.id
  name               = "%[1]s-fp"
  type               = "8p8c"
  rear_port_id       = netbox_device_rear_port.on_module.id
  rear_port_position = 1
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device_module_bay.test", "position"),
					resource.TestCheckNoResourceAttr("netbox_device_module_bay.test", "label"),
					resource.TestCheckResourceAttr("netbox_device_module_bay.nested", "enabled", "true"),
					resource.TestCheckNoResourceAttr("netbox_device_module_bay.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_module.test", "serial"),
					resource.TestCheckNoResourceAttr("netbox_module.test", "asset_tag"),
					resource.TestCheckNoResourceAttr("netbox_module.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_module.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_module.test", "replicate_components"),
					resource.TestCheckNoResourceAttr("netbox_module.test", "adopt_components"),
					resource.TestCheckNoResourceAttr("netbox_device_module_bay.nested", "module_id"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.on_module", "module_id"),
					resource.TestCheckNoResourceAttr("netbox_device_console_port.on_module", "module_id"),
					resource.TestCheckNoResourceAttr("netbox_device_console_server_port.on_module", "module_id"),
					resource.TestCheckNoResourceAttr("netbox_device_power_port.on_module", "module_id"),
					resource.TestCheckNoResourceAttr("netbox_device_power_outlet.on_module", "module_id"),
					resource.TestCheckNoResourceAttr("netbox_device_rear_port.on_module", "module_id"),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.on_module", "module_id"),
				),
			},
		},
	})
}

func TestAccNetboxInventoryItem_basic(t *testing.T) {
	testName := testAccGetTestName("invitem")
	deps := componentDeps(testName, "") + fmt.Sprintf(`
resource "netbox_inventory_item_role" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  color_hex   = "00ff00"
  description = "Acceptance test inventory item role."
  comments    = "Acceptance test comments."
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-eth0"
  type      = "virtual"
}
resource "netbox_inventory_item" "parent" {
  device_id = netbox_device.test.id
  name      = "%[1]s-chassis"
}`, testName, getSlug(testName))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_inventory_item" "test" {
  device_id       = netbox_device.test.id
  name            = "%[1]s"
  label           = "slot A"
  status          = "planned"
  parent_id       = netbox_inventory_item.parent.id
  role_id         = netbox_inventory_item_role.test.id
  manufacturer_id = netbox_manufacturer.test.id
  part_id         = "PN-%[1]s"
  serial          = "SN-%[1]s"
  asset_tag       = "AT-%[1]s"
  discovered      = true
  component_type  = "dcim.interface"
  component_id    = netbox_device_interface.test.id
  description     = "Acceptance test inventory item."
}
data "netbox_inventory_items" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_inventory_item.test]
}
data "netbox_inventory_item_roles" "test" {
  filters = [
    { name = "name", value = netbox_inventory_item_role.test.name },
  ]
  depends_on = [netbox_inventory_item_role.test]
}
data "netbox_inventory_item" "test" {
  id = netbox_inventory_item.test.id
}
data "netbox_inventory_item_role" "test" {
  id = netbox_inventory_item_role.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "color_hex", "00ff00"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "slug", testName),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "part_id", "PN-"+testName),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "serial", "SN-"+testName),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "discovered", "true"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "parent_id",
						"netbox_inventory_item.parent", "id"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "component_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "component_id",
						"netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item.test", "role_id",
						"netbox_inventory_item_role.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_inventory_items.test", "inventory_items.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_inventory_item_roles.test", "inventory_item_roles.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_inventory_item.test", "id",
						"netbox_inventory_item.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_inventory_item_role.test", "id",
						"netbox_inventory_item_role.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_inventory_item.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_inventory_item_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: the item clears every optional, and the role clears its description
				// (slug is computed and keeps its value, color_hex falls back to its default).
				Config: componentDeps(testName, "") + fmt.Sprintf(`
resource "netbox_inventory_item_role" "test" {
  name = "%[1]s"
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-eth0"
  type      = "virtual"
}
resource "netbox_inventory_item" "parent" {
  device_id = netbox_device.test.id
  name      = "%[1]s-chassis"
}
resource "netbox_inventory_item" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "parent_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "manufacturer_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "component_type"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "component_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "part_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "serial"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "asset_tag"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "label"),
					resource.TestCheckResourceAttr("netbox_inventory_item.test", "status", "planned"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_role.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_role.test", "comments"),
					resource.TestCheckResourceAttr("netbox_inventory_item_role.test", "color_hex", "9e9e9e"),
				),
			},
		},
	})
}

func TestAccNetboxVirtualChassis_basic(t *testing.T) {
	testName := testAccGetTestName("vc")
	deps := componentDeps(testName, "")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_chassis" "test" {
  name        = "%[1]s"
  domain      = "%[1]s.example.com"
  description = "Acceptance test virtual chassis."
  comments    = "Created by acceptance test."
}
data "netbox_virtual_chassis" "test" {
  name       = netbox_virtual_chassis.test.name
  depends_on = [netbox_virtual_chassis.test]
}
data "netbox_virtual_chassis_list" "test" {
  filters = [
    { name = "name", value = netbox_virtual_chassis.test.name },
  ]
  depends_on = [netbox_virtual_chassis.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_chassis.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_chassis.test", "domain", testName+".example.com"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_chassis.test", "id",
						"netbox_virtual_chassis.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_chassis_list.test", "virtual_chassis_list.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_chassis_list.test", "virtual_chassis_list.0.id",
						"netbox_virtual_chassis.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_virtual_chassis.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_chassis" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_virtual_chassis.test", "domain"),
					resource.TestCheckNoResourceAttr("netbox_virtual_chassis.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_virtual_chassis.test", "comments"),
				),
			},
		},
	})
}

// TestAccNetboxDeviceFrontPort_basic covers the front port's two mapping forms: the flat
// rear_port_id/rear_port_position pair of a single-position port and the rear_ports set of a
// multi-position port, including re-pointing one slot and unmapping. positions is lowered in a step
// of its own because NetBox checks it against the mappings it still holds.
func TestAccNetboxDeviceFrontPort_basic(t *testing.T) {
	testName := testAccGetTestName("frontport")
	deps := componentDeps(testName, "") + fmt.Sprintf(`
resource "netbox_device_rear_port" "a" {
  device_id = netbox_device.test.id
  name      = "%[1]s-rp-a"
  type      = "8p8c"
  positions = 4
}
resource "netbox_device_rear_port" "b" {
  device_id = netbox_device.test.id
  name      = "%[1]s-rp-b"
  type      = "8p8c"
  positions = 2
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// The flat pair: a single-position port on one rear port position.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id          = netbox_device.test.id
  name               = "%[1]s-fp"
  type               = "8p8c"
  rear_port_id       = netbox_device_rear_port.a.id
  rear_port_position = 2
  color_hex          = "ff0000"
  label              = "front"
  mark_connected     = true
  description        = "Acceptance test front port."
}
data "netbox_device_front_ports" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_front_port.test]
}
data "netbox_device_front_port" "test" {
  id = netbox_device_front_port.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "positions", "1"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_port_position", "2"),
					resource.TestCheckResourceAttrPair("netbox_device_front_port.test", "rear_port_id",
						"netbox_device_rear_port.a", "id"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_ports.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("netbox_device_front_port.test", "rear_ports.*",
						map[string]string{"position": "1", "rear_port_position": "2"}),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "color_hex", "ff0000"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "label", "front"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("data.netbox_device_front_ports.test", "device_front_ports.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_device_front_ports.test", "device_front_ports.0.rear_port_id", "netbox_device_rear_port.a", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_front_ports.test", "device_front_ports.0.rear_port_position", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_device_front_port.test", "id",
						"netbox_device_front_port.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_front_port.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// The set: two positions on two rear ports. The pair is not derivable and nulls.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-fp"
  type      = "8p8c"
  positions = 2
  rear_ports = [
    { position = 1, rear_port_id = netbox_device_rear_port.a.id, rear_port_position = 3 },
    { position = 2, rear_port_id = netbox_device_rear_port.b.id, rear_port_position = 1 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "positions", "2"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_ports.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("netbox_device_front_port.test", "rear_ports.*",
						map[string]string{"position": "2", "rear_port_position": "1"}),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.test", "rear_port_id"),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.test", "rear_port_position"),
				),
			},
			{
				// Re-point one slot: NetBox reconciles by position and keeps the other mapping.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-fp"
  type      = "8p8c"
  positions = 2
  rear_ports = [
    { position = 1, rear_port_id = netbox_device_rear_port.a.id, rear_port_position = 3 },
    { position = 2, rear_port_id = netbox_device_rear_port.b.id, rear_port_position = 2 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_ports.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("netbox_device_front_port.test", "rear_ports.*",
						map[string]string{"position": "2", "rear_port_position": "2"}),
				),
			},
			{
				// An update that keeps the set: rear_ports leaves the request body (NetBox 4.6 rejects a
				// resend of an existing mapping) and the mappings stay as they are.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-fp"
  type      = "8p8c"
  positions = 2
  label     = "relabelled"
  rear_ports = [
    { position = 1, rear_port_id = netbox_device_rear_port.a.id, rear_port_position = 3 },
    { position = 2, rear_port_id = netbox_device_rear_port.b.id, rear_port_position = 2 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "label", "relabelled"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_ports.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("netbox_device_front_port.test", "rear_ports.*",
						map[string]string{"position": "1", "rear_port_position": "3"}),
				),
			},
			{
				// Shrink: the mappings go, positions stays for now (NetBox checks a lower value against
				// the mappings still in the database), the optional scalars clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-fp"
  type      = "8p8c"
  positions = 2
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_ports.#", "0"),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.test", "rear_port_id"),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.test", "rear_port_position"),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.test", "color_hex"),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_device_front_port.test", "description"),
				),
			},
			{
				// Now positions can fall back to its default.
				Config: deps + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-fp"
  type      = "8p8c"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "positions", "1"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_ports.#", "0"),
				),
			},
		},
	})
}
