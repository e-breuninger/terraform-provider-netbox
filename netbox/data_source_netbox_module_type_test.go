package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxModuleTypeSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test" {
	name = "%[1]sa"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s"
  part_number = "test_pn"
  description = "test_description"
  comments = "test_comments"

  weight = 1
  weight_unit = "kg"
  tags = ["%[1]sa"]
}

resource "netbox_module_type" "test_2" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s_2"
}`, testName)
}

const testAccNetboxModuleTypeNoResult = `
data "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "_does_not_exist_"
}`

func testAccNetboxModuleType(testName string) string {
	return fmt.Sprintf(`
data "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s"
  depends_on = [
      netbox_module_type.test
    ]
}`, testName)
}

func testAccNetboxModuleType2(testName string) string {
	return fmt.Sprintf(`
data "netbox_module_type" "test_2" {
 manufacturer_id = netbox_manufacturer.test.id
 model = "%[1]s_2"
 depends_on = [
      netbox_module_type.test_2
    ]
}`, testName)
}

func TestAccNetboxModuleTypeDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("module_type_ds_basic")
	setUp := testAccNetboxModuleTypeSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config:      setUp + testAccNetboxModuleTypeNoResult,
				ExpectError: regexp.MustCompile("expected one"),
			},
			{
				Config: setUp + testAccNetboxModuleType(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_module_type.test", "id", "netbox_module_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_module_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_module_type.test", "model", testName),
					resource.TestCheckResourceAttr("data.netbox_module_type.test", "part_number", "test_pn"),
					resource.TestCheckResourceAttr("data.netbox_module_type.test", "description", "test_description"),
					resource.TestCheckResourceAttr("data.netbox_module_type.test", "comments", "test_comments"),
					resource.TestCheckResourceAttr("data.netbox_module_type.test", "weight", "1"),
					resource.TestCheckResourceAttr("data.netbox_module_type.test", "weight_unit", "kg"),
				),
			},
			{
				Config: setUp + testAccNetboxModuleType2(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_module_type.test_2", "id", "netbox_module_type.test_2", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_module_type.test_2", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_module_type.test_2", "model", fmt.Sprintf("%s_2", testName)),
					resource.TestCheckResourceAttr("data.netbox_module_type.test_2", "part_number", ""),
					resource.TestCheckResourceAttr("data.netbox_module_type.test_2", "description", ""),
					resource.TestCheckResourceAttr("data.netbox_module_type.test_2", "comments", ""),
					resource.TestCheckResourceAttr("data.netbox_module_type.test_2", "weight", "0"),
					resource.TestCheckResourceAttr("data.netbox_module_type.test_2", "weight_unit", ""),
				),
			},
		},
	})
}
