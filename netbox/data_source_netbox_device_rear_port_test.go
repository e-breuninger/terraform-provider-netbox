package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxDeviceRearPortSetUp(testName string) string {
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
}`, testName)
}

const testAccNetboxDeviceRearPortNoResult = `
data "netbox_device_rear_port" "test" {
  device_id = netbox_device.test.id
  name = "_does_not_exist_"
  depends_on = [
      netbox_device_rear_port.test
    ]
}`

func testAccNetboxDeviceRearPortByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_device_rear_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  depends_on = [
      netbox_device_rear_port.test
    ]
}`, testName)
}

func TestAccNetboxDeviceRearPortDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("module_bay_ds_basic")
	setUp := testAccNetboxDeviceRearPortSetUp(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      setUp + testAccNetboxDeviceRearPortNoResult,
				ExpectError: regexp.MustCompile("expected one"),
			},
			{
				Config: setUp + testAccNetboxDeviceRearPortByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_rear_port.test", "id", "netbox_device_rear_port.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_port.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_rear_port.test", "description", "test_description"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_port.test", "type", "8p8c"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_port.test", "positions", "2"),
					resource.TestCheckResourceAttr("data.netbox_device_rear_port.test", "mark_connected", "true"),
				),
			},
		},
	})
}
