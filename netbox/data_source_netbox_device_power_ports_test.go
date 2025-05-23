package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDevicePowerportsDataSource_basic(t *testing.T) {
	testSlug := "dev_powerports_ds_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxDevicePowerportsDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + fmt.Sprintf(`
data "netbox_device_power_ports" "by_name" {
  filter {
    name = "name"
    value  = "%[1]s"
  }
}

data "netbox_device_power_ports" "by_device_id" {
  filter {
    name = "device_id"
    value  = netbox_device.test.id
  }
}

data "netbox_device_power_ports" "by_tag" {
  filter {
    name = "tag"
    value  = "%[1]s"
  }
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_power_ports.by_name", "power_ports.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_power_ports.by_name", "power_ports.0.name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_power_ports.by_name", "power_ports.0.type", "iec-60309-3p-n-e-9h"),
					resource.TestCheckResourceAttrPair("data.netbox_device_power_ports.by_name", "power_ports.0.device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_power_ports.by_device_id", "power_ports.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_device_power_ports.by_tag", "power_ports.#", "2"),
				),
			},
		},
	})
}

func testAccNetboxDevicePowerportsDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device_power_port" "test" {
  name = "%[1]s"
  device_id = netbox_device.test.id
  type = "iec-60309-3p-n-e-9h"
  tags = ["%[1]s"]
}

resource "netbox_device_power_port" "test2" {
  name = "%[1]s_two"
  device_id = netbox_device.test.id
  tags = ["%[1]s"]
}

`, testName)
}
