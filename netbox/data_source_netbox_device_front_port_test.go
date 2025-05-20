package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxDeviceFrontPortSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device" "test_2" {
 name = "%[1]s_2"
 device_type_id = netbox_device_type.test.id
 role_id = netbox_device_role.test.id
 site_id = netbox_site.test.id
}

resource "netbox_device_rear_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  description = "test_description"
  type           = "8p8c"
  positions      = 2
  mark_connected = true
}

resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "8p8c"
  rear_port_id = netbox_device_rear_port.test.id
  rear_port_position = 1

  mark_connected = true
  label = "%[1]s_label"
  #color_hex = "123456"
  description = "test_description"
  #tags = ["%[1]sa"]
}`, testName)
}

const testAccNetboxDeviceFrontPortNoResult = `
data "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name = "_does_not_exist_"
  depends_on = [
      netbox_device_front_port.test
    ]
}`

func testAccNetboxDeviceFrontPortByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  depends_on = [
      netbox_device_front_port.test
    ]
}`, testName)
}

func TestAccNetboxDeviceFrontPortDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("device_front_port_ds_basic")
	setUp := testAccNetboxDeviceFrontPortSetUp(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      setUp + testAccNetboxDeviceFrontPortNoResult,
				ExpectError: regexp.MustCompile("expected one"),
			},
			{
				Config: setUp + testAccNetboxDeviceFrontPortByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_front_port.test", "id", "netbox_device_front_port.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "type", "8p8c"),
					resource.TestCheckResourceAttrPair("data.netbox_device_front_port.test", "rear_port_id", "netbox_device_front_port.test", "rear_port_id"),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "rear_port_position", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "module_id", "0"),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "label", testName+"_label"),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "color_hex", ""),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "description", "test_description"),
					resource.TestCheckResourceAttr("data.netbox_device_front_port.test", "mark_connected", "true"),
				),
			},
		},
	})
}
