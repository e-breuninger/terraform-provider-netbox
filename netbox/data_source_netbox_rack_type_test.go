package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRackTypeDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("ds_rk_tp")
	setUp := testAccNetboxRackTypeDataSourceSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_type.test", "slug", testName),
				),
			},
			{
				Config:      setUp + testAccNetboxRackTypeDataNoResult,
				ExpectError: regexp.MustCompile("no rack type found matching filter"),
			},
			{
				Config: setUp + testAccNetboxRackTypeDataByModel(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_rack_type.test", "slug", testName),
				),
			},
			{
				Config: setUp + testAccNetboxRackTypeDataBySlug(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_rack_type.test", "model", testName),
					resource.TestCheckResourceAttrPair("data.netbox_rack_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_rack_type.test", "form_factor", "4-post-frame"),
					resource.TestCheckResourceAttr("data.netbox_rack_type.test", "width", "19"),
					resource.TestCheckResourceAttr("data.netbox_rack_type.test", "u_height", "48"),
					resource.TestCheckResourceAttr("data.netbox_rack_type.test", "starting_unit", "1"),
					resource.TestCheckResourceAttrSet("data.netbox_rack_type.test", "id"),
				),
			},
		},
	})
}

func testAccNetboxRackTypeDataSourceSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_rack_type" "test" {
  model           = "%[1]s"
  slug            = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
  form_factor     = "4-post-frame"
  width           = 19
  u_height        = 48
  starting_unit   = 1
}`, testName)
}

const testAccNetboxRackTypeDataNoResult = `
data "netbox_rack_type" "no_result" {
  model = "_no_result_"
}`

func testAccNetboxRackTypeDataByModel(testName string) string {
	return fmt.Sprintf(`
data "netbox_rack_type" "test" {
  model = "%[1]s"
}`, testName)
}

func testAccNetboxRackTypeDataBySlug(testName string) string {
	return fmt.Sprintf(`
data "netbox_rack_type" "test" {
  slug = "%[1]s"
}`, testName)
}
