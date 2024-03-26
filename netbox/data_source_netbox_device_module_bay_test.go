package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxDeviceModuleBaySetUp(testName string) string {
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

resource "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  label = "test_label"
  #position = "1"
  description = "test_description"
}

resource "netbox_device_module_bay" "test_2" {
  device_id = netbox_device.test_2.id
  name = "%[1]s"
  label = "test_label"
  #position = "1"
  description = "test_description"
}`, testName)
}

const testAccNetboxDeviceModuleBayNoResult = `
data "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name = "_does_not_exist_"
}`

func testAccNetboxDeviceModuleBayByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
}`, testName)
}

func testAccNetboxDeviceModuleBayByLabel() string {
	return `
data "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  label = "test_label"
}`
}

func TestAccNetboxDeviceModuleBayDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("module_bay_ds_basic")
	setUp := testAccNetboxDeviceModuleBaySetUp(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      setUp + testAccNetboxDeviceModuleBayNoResult,
				ExpectError: regexp.MustCompile("expected one"),
			},
			{
				Config: setUp + testAccNetboxDeviceModuleBayByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_module_bay.test", "id", "netbox_device_module_bay.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "description", "test_description"),
					resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "label", "test_label"),
					//resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "position", "1"),
				),
			},
			{
				Config: setUp + testAccNetboxDeviceModuleBayByLabel(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_module_bay.test", "id", "netbox_device_module_bay.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "description", "test_description"),
					resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "label", "test_label"),
					//resource.TestCheckResourceAttr("data.netbox_device_module_bay.test", "position", "1"),
				),
			},
		},
	})
}
