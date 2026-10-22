//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// templateDeps is a device type and a module type to hang component templates off.
func templateDeps(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
  # NetBox only allows device bays on a parent device type.
  subdevice_role = "parent"
}
resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}`, testName)
}

func TestAccNetboxConsolePortTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("cptmpl")
	deps := templateDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_console_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "rj-45"
  label          = "console"
  description    = "Acceptance test template."
}
data "netbox_console_port_template" "test" {
  name = netbox_console_port_template.test.name
}
data "netbox_console_port_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_console_port_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_console_port_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_console_port_template.test", "type", "rj-45"),
					resource.TestCheckResourceAttrPair("netbox_console_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_console_port_template.test", "id", "netbox_console_port_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_console_port_templates.by_name", "console_port_templates.#", "1"),
				),
			},
			{
				// Shrink: the type moves to the module type and everything optional clears.
				Config: deps + fmt.Sprintf(`
resource "netbox_console_port_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_console_port_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_console_port_template.test", "device_type_id"),
					resource.TestCheckNoResourceAttr("netbox_console_port_template.test", "type"),
					resource.TestCheckNoResourceAttr("netbox_console_port_template.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_console_port_template.test", "description"),
				),
			},
			{
				// Back to the device type. The owner is immutable in NetBox, so both switches are
				// replaces (device_type_id and module_type_id are force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_console_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_console_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_console_port_template.test", "module_type_id"),
				),
			},
			{
				ResourceName:      "netbox_console_port_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxConsoleServerPortTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("cspttmpl")
	deps := templateDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_console_server_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "rj-45"
  label          = "console-server"
  description    = "Acceptance test template."
}
data "netbox_console_server_port_template" "test" {
  name = netbox_console_server_port_template.test.name
}
data "netbox_console_server_port_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_console_server_port_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_console_server_port_template.test", "type", "rj-45"),
					resource.TestCheckResourceAttrPair("data.netbox_console_server_port_template.test", "id", "netbox_console_server_port_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_console_server_port_templates.by_name", "console_server_port_templates.#", "1"),
				),
			},
			{
				// Shrink onto the module type: the optionals clear and the device ownership goes
				// with the replace (the owner is immutable in NetBox, both ids are force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_console_server_port_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_console_server_port_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_console_server_port_template.test", "device_type_id"),
					resource.TestCheckNoResourceAttr("netbox_console_server_port_template.test", "type"),
					resource.TestCheckNoResourceAttr("netbox_console_server_port_template.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_console_server_port_template.test", "description"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_console_server_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_console_server_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_console_server_port_template.test", "module_type_id"),
				),
			},
			{
				ResourceName:      "netbox_console_server_port_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxDeviceBayTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("dbtmpl")
	deps := templateDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_bay_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  label          = "bay"
  enabled        = false
  description    = "Acceptance test template."
}
data "netbox_device_bay_template" "test" {
  name = netbox_device_bay_template.test.name
}
data "netbox_device_bay_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_device_bay_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_bay_template.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device_bay_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_device_bay_template.test", "id", "netbox_device_bay_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_bay_templates.by_name", "device_bay_templates.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_bay_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device_bay_template.test", "label"),
					resource.TestCheckResourceAttr("netbox_device_bay_template.test", "enabled", "true"),
					resource.TestCheckNoResourceAttr("netbox_device_bay_template.test", "description"),
				),
			},
			{
				ResourceName:      "netbox_device_bay_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxInterfaceTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("iftmpl")
	deps := templateDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_interface_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "1000base-t"
  mgmt_only      = true
  poe_mode       = "pse"
  poe_type       = "type1-ieee802.3af"
  label          = "eth"
  enabled        = false
  description    = "Acceptance test template."
}
resource "netbox_interface_template" "module" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s-module"
  type           = "1000base-t"
}
resource "netbox_interface_template" "wireless" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-wireless"
  type           = "ieee802.11ax"
  rf_role        = "ap"
  bridge_id      = netbox_interface_template.test.id
}
data "netbox_interface_template" "test" {
  name = netbox_interface_template.test.name
}
data "netbox_interface_templates" "by_module_type" {
  filters = [
    { name = "module_type_id", value = netbox_module_type.test.id },
  ]
  depends_on     = [netbox_interface_template.module]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface_template.test", "type", "1000base-t"),
					resource.TestCheckResourceAttrPair("data.netbox_interface_templates.by_module_type", "interface_templates.0.id", "netbox_interface_template.module", "id"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "mgmt_only", "true"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "poe_mode", "pse"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "poe_type", "type1-ieee802.3af"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_interface_template.wireless", "rf_role", "ap"),
					resource.TestCheckResourceAttrPair("netbox_interface_template.wireless", "bridge_id", "netbox_interface_template.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_interface_template.test", "id", "netbox_interface_template.test", "id"),
				),
			},
			{
				// Shrink: the optionals clear, mgmt_only is computed and keeps its value.
				Config: deps + fmt.Sprintf(`
resource "netbox_interface_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "1000base-t"
}
resource "netbox_interface_template" "wireless" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-wireless"
  type           = "ieee802.11ax"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_interface_template.wireless", "rf_role"),
					resource.TestCheckNoResourceAttr("netbox_interface_template.wireless", "bridge_id"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "mgmt_only", "true"),
					resource.TestCheckNoResourceAttr("netbox_interface_template.test", "poe_mode"),
					resource.TestCheckNoResourceAttr("netbox_interface_template.test", "poe_type"),
					resource.TestCheckNoResourceAttr("netbox_interface_template.test", "label"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "enabled", "true"),
					resource.TestCheckNoResourceAttr("netbox_interface_template.test", "description"),
				),
			},
			{
				// Owner switch: to the module type and back, each a replace (both ids force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_interface_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
  type           = "1000base-t"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_interface_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_interface_template.test", "device_type_id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_interface_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "1000base-t"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_interface_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_interface_template.test", "module_type_id"),
				),
			},
			{
				ResourceName:      "netbox_interface_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxModuleBayTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("mbtmpl")
	deps := templateDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_module_bay_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  position       = "1"
  label          = "bay"
  enabled        = false
  description    = "Acceptance test template."
}
data "netbox_module_bay_template" "test" {
  name = netbox_module_bay_template.test.name
}
data "netbox_module_bay_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_module_bay_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "position", "1"),
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "enabled", "false"),
					resource.TestCheckResourceAttrPair("netbox_module_bay_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_module_bay_template.test", "id", "netbox_module_bay_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_module_bay_templates.by_name", "module_bay_templates.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_module_bay_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_module_bay_template.test", "position"),
					resource.TestCheckNoResourceAttr("netbox_module_bay_template.test", "label"),
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "enabled", "true"),
					resource.TestCheckNoResourceAttr("netbox_module_bay_template.test", "description"),
				),
			},
			{
				// Owner switch: onto the module type, a replace (both ids are force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_module_bay_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_module_bay_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_module_bay_template.test", "device_type_id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_module_bay_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_module_bay_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_module_bay_template.test", "module_type_id"),
				),
			},
			{
				ResourceName:      "netbox_module_bay_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPowerPortTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("pptmpl")
	deps := templateDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_power_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "iec-60320-c14"
  maximum_draw   = 600
  allocated_draw = 300
  label          = "psu"
  description    = "Acceptance test template."
}
data "netbox_power_port_template" "test" {
  name = netbox_power_port_template.test.name
}
data "netbox_power_port_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_power_port_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "type", "iec-60320-c14"),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "maximum_draw", "600"),
					resource.TestCheckResourceAttr("netbox_power_port_template.test", "allocated_draw", "300"),
					resource.TestCheckResourceAttrPair("data.netbox_power_port_template.test", "id", "netbox_power_port_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_power_port_templates.by_name", "power_port_templates.#", "1"),
				),
			},
			{
				// Shrink onto the module type: the optionals clear and the device ownership goes
				// with the replace (the owner is immutable in NetBox, both ids are force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_power_port_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_power_port_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_power_port_template.test", "device_type_id"),
					resource.TestCheckNoResourceAttr("netbox_power_port_template.test", "type"),
					resource.TestCheckNoResourceAttr("netbox_power_port_template.test", "maximum_draw"),
					resource.TestCheckNoResourceAttr("netbox_power_port_template.test", "allocated_draw"),
					resource.TestCheckNoResourceAttr("netbox_power_port_template.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_power_port_template.test", "description"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_power_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_power_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_power_port_template.test", "module_type_id"),
				),
			},
			{
				ResourceName:      "netbox_power_port_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPowerOutletTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("potmpl")
	deps := templateDeps(testName) + fmt.Sprintf(`
resource "netbox_power_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-psu"
  type           = "iec-60320-c14"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_power_outlet_template" "test" {
  device_type_id         = netbox_device_type.test.id
  name                   = "%[1]s"
  type                   = "iec-60320-c13"
  power_port_template_id = netbox_power_port_template.test.id
  feed_leg               = "A"
  label                  = "outlet"
  color_hex              = "ff0000"
  description            = "Acceptance test template."
}
data "netbox_power_outlet_template" "test" {
  name = netbox_power_outlet_template.test.name
}
data "netbox_power_outlet_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_power_outlet_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "type", "iec-60320-c13"),
					resource.TestCheckResourceAttr("netbox_power_outlet_template.test", "feed_leg", "A"),
					resource.TestCheckResourceAttrPair("netbox_power_outlet_template.test", "power_port_template_id", "netbox_power_port_template.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_power_outlet_template.test", "id", "netbox_power_outlet_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_power_outlet_templates.by_name", "power_outlet_templates.#", "1"),
				),
			},
			{
				// Shrink onto the module type: the optionals clear and the device ownership goes
				// with the replace (the owner is immutable in NetBox, both ids are force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_power_outlet_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_power_outlet_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "device_type_id"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "type"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "power_port_template_id"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "feed_leg"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "color_hex"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "description"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_power_outlet_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_power_outlet_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_power_outlet_template.test", "module_type_id"),
				),
			},
			{
				ResourceName:      "netbox_power_outlet_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxRearPortTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("rptmpl")
	deps := templateDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_rear_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
  positions      = 4
  color_hex      = "ff0000"
  label          = "rear"
  description    = "Acceptance test template."
}
data "netbox_rear_port_template" "test" {
  name = netbox_rear_port_template.test.name
}
data "netbox_rear_port_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_rear_port_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rear_port_template.test", "type", "8p8c"),
					resource.TestCheckResourceAttr("netbox_rear_port_template.test", "positions", "4"),
					resource.TestCheckResourceAttr("netbox_rear_port_template.test", "color_hex", "ff0000"),
					resource.TestCheckResourceAttrPair("data.netbox_rear_port_template.test", "id", "netbox_rear_port_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_rear_port_templates.by_name", "rear_port_templates.#", "1"),
				),
			},
			{
				// Shrink onto the module type: positions falls back to its default, everything
				// else clears, and the device ownership goes with the replace (the owner is
				// immutable in NetBox, both ids are force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_rear_port_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_rear_port_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_rear_port_template.test", "device_type_id"),
					resource.TestCheckResourceAttr("netbox_rear_port_template.test", "positions", "1"),
					resource.TestCheckNoResourceAttr("netbox_rear_port_template.test", "color_hex"),
					resource.TestCheckNoResourceAttr("netbox_rear_port_template.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_rear_port_template.test", "description"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_rear_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_rear_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_rear_port_template.test", "module_type_id"),
				),
			},
			{
				ResourceName:      "netbox_rear_port_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccNetboxFrontPortTemplate_basic mirrors TestAccNetboxDeviceFrontPort_basic on the device
// type: the flat pair, the rear_ports set, unmapping, then positions back to its default.
func TestAccNetboxFrontPortTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("fptmpl")
	deps := templateDeps(testName) + fmt.Sprintf(`
resource "netbox_rear_port_template" "a" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-rp-a"
  type           = "8p8c"
  positions      = 4
}
resource "netbox_rear_port_template" "b" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-rp-b"
  type           = "8p8c"
  positions      = 2
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_front_port_template" "test" {
  device_type_id     = netbox_device_type.test.id
  name               = "%[1]s"
  type               = "8p8c"
  rear_port_id       = netbox_rear_port_template.a.id
  rear_port_position = 2
  color_hex          = "ff0000"
  label              = "front"
  description        = "Acceptance test template."
}
data "netbox_front_port_template" "test" {
  name = netbox_front_port_template.test.name
}
data "netbox_front_port_templates" "by_name" {
  filters = [
    { name = "name", value = netbox_front_port_template.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "positions", "1"),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_port_position", "2"),
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "rear_port_id",
						"netbox_rear_port_template.a", "id"),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_ports.#", "1"),
					resource.TestCheckTypeSetElemNestedAttrs("netbox_front_port_template.test", "rear_ports.*",
						map[string]string{"position": "1", "rear_port_position": "2"}),
					resource.TestCheckResourceAttrPair("data.netbox_front_port_template.test", "id", "netbox_front_port_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_front_port_templates.by_name", "front_port_templates.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_front_port_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_front_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
  positions      = 2
  rear_ports = [
    { position = 1, rear_port_id = netbox_rear_port_template.a.id, rear_port_position = 3 },
    { position = 2, rear_port_id = netbox_rear_port_template.b.id, rear_port_position = 1 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_ports.#", "2"),
					resource.TestCheckTypeSetElemNestedAttrs("netbox_front_port_template.test", "rear_ports.*",
						map[string]string{"position": "2", "rear_port_position": "1"}),
					resource.TestCheckNoResourceAttr("netbox_front_port_template.test", "rear_port_id"),
				),
			},
			{
				// An update that keeps the set leaves rear_ports out of the request body.
				Config: deps + fmt.Sprintf(`
resource "netbox_front_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
  positions      = 2
  label          = "relabelled"
  rear_ports = [
    { position = 1, rear_port_id = netbox_rear_port_template.a.id, rear_port_position = 3 },
    { position = 2, rear_port_id = netbox_rear_port_template.b.id, rear_port_position = 1 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "label", "relabelled"),
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_ports.#", "2"),
				),
			},
			{
				// Shrink in two steps: the mappings first, positions in the next apply.
				Config: deps + fmt.Sprintf(`
resource "netbox_front_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
  positions      = 2
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_front_port_template.test", "rear_ports.#", "0"),
					resource.TestCheckNoResourceAttr("netbox_front_port_template.test", "rear_port_id"),
					resource.TestCheckNoResourceAttr("netbox_front_port_template.test", "color_hex"),
					resource.TestCheckNoResourceAttr("netbox_front_port_template.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_front_port_template.test", "description"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_front_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
}`, testName),
				Check: resource.TestCheckResourceAttr("netbox_front_port_template.test", "positions", "1"),
			},
			{
				// Owner switch: to the module type and back, each a replace (both ids force_new).
				Config: deps + fmt.Sprintf(`
resource "netbox_front_port_template" "test" {
  module_type_id = netbox_module_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "module_type_id", "netbox_module_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_front_port_template.test", "device_type_id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_front_port_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
  type           = "8p8c"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_front_port_template.test", "device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_front_port_template.test", "module_type_id"),
				),
			},
		},
	})
}

func TestAccNetboxInventoryItemTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("iitmpl")
	deps := templateDeps(testName) + fmt.Sprintf(`
resource "netbox_inventory_item_role" "test" {
  name = "%[1]s"
}
resource "netbox_interface_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-eth0"
  type           = "1000base-t"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_inventory_item_template" "parent" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-chassis"
}
resource "netbox_inventory_item_template" "test" {
  device_type_id  = netbox_device_type.test.id
  parent_id       = netbox_inventory_item_template.parent.id
  name            = "%[1]s"
  label           = "SFP"
  role_id         = netbox_inventory_item_role.test.id
  manufacturer_id = netbox_manufacturer.test.id
  part_id         = "SFP-10G-SR"
  component_type  = "dcim.interfacetemplate"
  component_id    = netbox_interface_template.test.id
  description     = "Acceptance test template."
}
data "netbox_inventory_item_template" "test" {
  name = netbox_inventory_item_template.test.name
}
data "netbox_inventory_item_templates" "test" {
  filters = [
    { name = "device_type_id", value = netbox_device_type.test.id },
  ]
  depends_on     = [netbox_inventory_item_template.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_inventory_item_template.test", "parent_id", "netbox_inventory_item_template.parent", "id"),
					resource.TestCheckResourceAttr("netbox_inventory_item_template.test", "part_id", "SFP-10G-SR"),
					resource.TestCheckResourceAttr("netbox_inventory_item_template.test", "component_type", "dcim.interfacetemplate"),
					resource.TestCheckResourceAttrPair("netbox_inventory_item_template.test", "component_id", "netbox_interface_template.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_inventory_item_template.test", "id", "netbox_inventory_item_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_inventory_item_templates.test", "inventory_item_templates.#", "2"),
				),
			},
			{
				ResourceName:      "netbox_inventory_item_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink. The parent stays: NetBox deletes a template's children with it.
				Config: deps + fmt.Sprintf(`
resource "netbox_inventory_item_template" "parent" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s-chassis"
}
resource "netbox_inventory_item_template" "test" {
  device_type_id = netbox_device_type.test.id
  name           = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "parent_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "manufacturer_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "part_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "component_type"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "component_id"),
					resource.TestCheckNoResourceAttr("netbox_inventory_item_template.test", "description"),
				),
			},
		},
	})
}
