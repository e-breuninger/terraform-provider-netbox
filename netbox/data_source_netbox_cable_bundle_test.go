package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCableBundleDataSource_basic(t *testing.T) {
	testSlug := "cable_bundle_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxCableBundleDataSourceDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxCableBundleDataSource(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_cable_bundle.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_cable_bundle.test", "description", "This is my cable bundle"),
					resource.TestCheckResourceAttr("data.netbox_cable_bundle.test", "comments", "Cable Bundle Comments"),
					resource.TestCheckResourceAttr("data.netbox_cable_bundle.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_cable_bundle.test", "tags.0", testName+"-tag"),
				),
			},
		},
	})
}

func testAccNetboxCableBundleDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
	name	= "%[1]s-tag"
}

resource "netbox_cable_bundle" "test" {
	name		= "%[1]s"
	description	= "This is my cable bundle"
	comments	= "Cable Bundle Comments"
	tags		= [ netbox_tag.test.name ]
}
`, testName)
}

func testAccNetboxCableBundleDataSource(testName string) string {
	return fmt.Sprintf(`
data "netbox_cable_bundle" "test" {
	depends_on	=	[]

	name		= "%[1]s"
}
`, testName)
}
