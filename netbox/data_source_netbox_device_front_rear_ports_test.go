package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxDeviceFrontRearPortsDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_rear_port" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s"
  type      = "8p8c"
  positions = 1
  tags      = [netbox_tag.test.name]
}

resource "netbox_device_rear_port" "test2" {
  device_id = netbox_device.test.id
  name      = "%[1]s_two"
  type      = "8p8c"
  positions = 1
  tags      = [netbox_tag.test.name]
}

resource "netbox_device_front_port" "test" {
  device_id          = netbox_device.test.id
  name               = "%[1]s"
  type               = "8p8c"
  rear_port_id       = netbox_device_rear_port.test.id
  rear_port_position = 1
  tags               = [netbox_tag.test.name]
}

resource "netbox_device_front_port" "test2" {
  device_id          = netbox_device.test.id
  name               = "%[1]s_two"
  type               = "8p8c"
  rear_port_id       = netbox_device_rear_port.test2.id
  rear_port_position = 1
  tags               = [netbox_tag.test.name]
}
`, testName)
}

func TestAccNetboxDeviceFrontPortsDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("dev_frontports_ds_basic")
	dependencies := testAccNetboxDeviceFrontRearPortsDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + fmt.Sprintf(`
data "netbox_device_front_ports" "by_name" {
  depends_on = [netbox_device_front_port.test, netbox_device_front_port.test2]
  filter {
    name  = "name"
    value = "%[1]s"
  }
}

data "netbox_device_front_ports" "by_device_id" {
  depends_on = [netbox_device_front_port.test, netbox_device_front_port.test2]
  filter {
    name  = "device_id"
    value = netbox_device.test.id
  }
}

data "netbox_device_front_ports" "by_tag" {
  depends_on = [netbox_device_front_port.test, netbox_device_front_port.test2]
  filter {
    name  = "tag"
    value = netbox_tag.test.name
  }
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_front_ports.by_name", "front_ports.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_front_ports.by_name", "front_ports.0.name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_front_ports.by_name", "front_ports.0.type", "8p8c"),
					resource.TestCheckResourceAttrPair("data.netbox_device_front_ports.by_name", "front_ports.0.device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_front_ports.by_device_id", "front_ports.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_device_front_ports.by_tag", "front_ports.#", "2"),
				),
			},
		},
	})
}

func TestAccNetboxDeviceRearPortsDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("dev_rearports_ds_basic")
	dependencies := testAccNetboxDeviceFrontRearPortsDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + fmt.Sprintf(`
data "netbox_device_rear_ports" "by_name" {
  depends_on = [netbox_device_rear_port.test, netbox_device_rear_port.test2]
  filter {
    name  = "name"
    value = "%[1]s"
  }
}

data "netbox_device_rear_ports" "by_device_id" {
  depends_on = [netbox_device_rear_port.test, netbox_device_rear_port.test2]
  filter {
    name  = "device_id"
    value = netbox_device.test.id
  }
}

data "netbox_device_rear_ports" "by_tag" {
  depends_on = [netbox_device_rear_port.test, netbox_device_rear_port.test2]
  filter {
    name  = "tag"
    value = netbox_tag.test.name
  }
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_rear_ports.by_name", "rear_ports.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_ports.by_name", "rear_ports.0.name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_rear_ports.by_name", "rear_ports.0.type", "8p8c"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_ports.by_name", "rear_ports.0.positions", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_device_rear_ports.by_name", "rear_ports.0.device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_ports.by_device_id", "rear_ports.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_ports.by_tag", "rear_ports.#", "2"),
				),
			},
		},
	})
}
