package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxRirSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}`, testName)
}

const testAccNetboxRirNoResult = `
data "netbox_rir" "test" {
  name = "nonexistent"
}`

func testAccNetboxRirByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_rir" "test" {
  name = "%s"
}`, testName)
}

func testAccNetboxRirBySlug(testName string) string {
	return fmt.Sprintf(`
data "netbox_rir" "test" {
  slug = "%[1]s"
}`, testName)
}

func TestAccNetboxRirDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("rir_ds_basic")
	setUp := testAccNetboxRirSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rir.test", "name", testName),
				),
			},
			{
				Config:      setUp + testAccNetboxRirNoResult,
				ExpectError: regexp.MustCompile("no rir found matching filter"),
			},
			{
				Config: setUp + testAccNetboxRirByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_rir.test", "id", "netbox_rir.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxRirBySlug(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_rir.test", "id", "netbox_rir.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_rir.test", "name", testName),
				),
			},
		},
	})
}
