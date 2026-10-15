//go:build acctest

// The type of a device's console, console server, power port and power outlet
// is optional in NetBox, as it is on their templates: a port can be created
// without one, given one later, and have it cleared again.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxOptionalPortType(t *testing.T) {
	testName := testAccGetTestName("porttype")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}`, testName)
	ports := func(consoleType, powerPortType, powerOutletType string) string {
		return deps + fmt.Sprintf(`
resource "netbox_device_console_port" "test" {
  device_id = netbox_device.test.id
  name      = "console"
  type      = %[1]s
}
resource "netbox_device_console_server_port" "test" {
  device_id = netbox_device.test.id
  name      = "console-server"
  type      = %[1]s
}
resource "netbox_device_power_port" "test" {
  device_id = netbox_device.test.id
  name      = "PSU1"
  type      = %[2]s
}
resource "netbox_device_power_outlet" "test" {
  device_id = netbox_device.test.id
  name      = "Outlet1"
  type      = %[3]s
}`, consoleType, powerPortType, powerOutletType)
	}
	noType := resource.ComposeTestCheckFunc(
		resource.TestCheckNoResourceAttr("netbox_device_console_port.test", "type"),
		resource.TestCheckNoResourceAttr("netbox_device_console_server_port.test", "type"),
		resource.TestCheckNoResourceAttr("netbox_device_power_port.test", "type"),
		resource.TestCheckNoResourceAttr("netbox_device_power_outlet.test", "type"),
	)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{ // created without a type
				Config: ports("null", "null", "null"),
				Check:  noType,
			},
			{
				Config: ports(`"rj-45"`, `"iec-60320-c14"`, `"iec-60320-c13"`),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_console_port.test", "type", "rj-45"),
					resource.TestCheckResourceAttr("netbox_device_console_server_port.test", "type", "rj-45"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "type", "iec-60320-c14"),
					resource.TestCheckResourceAttr("netbox_device_power_outlet.test", "type", "iec-60320-c13"),
				),
			},
			{ // removing the type clears it
				Config: ports("null", "null", "null"),
				Check:  noType,
			},
		},
	})
}
