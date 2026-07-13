package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDevicePowerOutletsDataSource_basic(t *testing.T) {
	testSlug := "dev_poweroutlets_ds_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxDevicePowerOutletsDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + fmt.Sprintf(`
data "netbox_device_power_outlets" "by_name" {
  depends_on = [
    netbox_device_power_outlet.test,
    netbox_device_power_outlet.test2,
  ]

  filter {
    name  = "name"
    value = "%[1]s"
  }
}

data "netbox_device_power_outlets" "by_device_id" {
  depends_on = [
    netbox_device_power_outlet.test,
    netbox_device_power_outlet.test2,
  ]

  filter {
    name  = "device_id"
    value = netbox_device.test.id
  }
}

data "netbox_device_power_outlets" "by_tag" {
  depends_on = [
    netbox_device_power_outlet.test,
    netbox_device_power_outlet.test2,
  ]

  filter {
    name  = "tag"
    value = netbox_tag.test.name
  }
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_power_outlets.by_name", "power_outlets.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_power_outlets.by_name", "power_outlets.0.name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_power_outlets.by_name", "power_outlets.0.type", "iec-60320-c5"),
					resource.TestCheckResourceAttr("data.netbox_device_power_outlets.by_name", "power_outlets.0.feed_leg", "A"),
					resource.TestCheckResourceAttrPair("data.netbox_device_power_outlets.by_name", "power_outlets.0.device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_device_power_outlets.by_name", "power_outlets.0.power_port_id", "netbox_device_power_port.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_power_outlets.by_device_id", "power_outlets.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_device_power_outlets.by_tag", "power_outlets.#", "2"),
				),
			},
		},
	})
}

func testAccNetboxDevicePowerOutletsDataSourceDependencies(testName string) string {
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
}

resource "netbox_device_power_outlet" "test" {
  name = "%[1]s"
  device_id = netbox_device.test.id
  power_port_id = netbox_device_power_port.test.id
  type = "iec-60320-c5"
  feed_leg = "A"
  tags = [netbox_tag.test.name]
}

resource "netbox_device_power_outlet" "test2" {
  name = "%[1]s_two"
  device_id = netbox_device.test.id
  tags = [netbox_tag.test.name]
}
`, testName)
}
