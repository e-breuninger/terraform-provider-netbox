package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceTypeDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("ds_dv_tp")
	setUp := testAccNetboxDeviceTypeSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "slug", testName),
				),
			},
			{
				Config:      setUp + testAccNetboxDeviceTypeDataNoResult,
				ExpectError: regexp.MustCompile("no device type found matching filter"),
			},
			{
				Config: setUp + testAccNetboxDeviceTypeDataByModel(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_type.test", "slug", testName),
				),
			},
			{
				Config: setUp + testAccNetboxDeviceTypeDataCombo(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_type.test", "slug", testName),
				),
			},
			{
				Config: setUp + testAccNetboxDeviceTypeDataBySlug(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_type.test", "model", testName),
					resource.TestCheckResourceAttr("data.netbox_device_type.test", "part_number", testName),
					resource.TestCheckResourceAttrPair("data.netbox_device_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttrSet("data.netbox_device_type.test", "is_full_depth"),
					resource.TestCheckResourceAttrSet("data.netbox_device_type.test", "u_height"),
				),
			},
		},
	})
}

func testAccNetboxDeviceTypeSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}
resource "netbox_device_type" "test" {
	manufacturer_id = netbox_manufacturer.test.id
	model = "%[1]s"
	part_number = "%[1]s"
}`, testName)
}

const testAccNetboxDeviceTypeDataNoResult = `
data "netbox_device_type" "no_result" {
	model = "_no_result_"
}`

func testAccNetboxDeviceTypeDataByModel(testName string) string {
	return fmt.Sprintf(`
data "netbox_device_type" "test" {
	model = "%[1]s"
}`, testName)
}

func testAccNetboxDeviceTypeDataCombo(testName string) string {
	return fmt.Sprintf(`
data "netbox_device_type" "test" {
	manufacturer = "%[1]s"
	part_number = "%[1]s"
}`, testName)
}

func testAccNetboxDeviceTypeDataBySlug(testName string) string {
	return fmt.Sprintf(`
data "netbox_device_type" "test" {
	slug = "%[1]s"
}`, testName)
}
