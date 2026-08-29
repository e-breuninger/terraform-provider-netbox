package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRackGroupDataSource_basic(t *testing.T) {
	testSlug := "rack_group_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_rack_group" "test" {
  name        = "%[1]s"
  slug        = "%[1]s"
  description = "This is my rack group"
  comments    = "Rack group comments"
  tags        = [netbox_tag.test.name]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_rack_group" "test" {
  depends_on = []

  name = netbox_rack_group.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_rack_group.test", "id", "netbox_rack_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_rack_group.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_rack_group.test", "slug", testName),
					resource.TestCheckResourceAttr("data.netbox_rack_group.test", "description", "This is my rack group"),
					resource.TestCheckResourceAttr("data.netbox_rack_group.test", "comments", "Rack group comments"),
					resource.TestCheckResourceAttr("data.netbox_rack_group.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_rack_group.test", "tags.0", testName),
				),
			},
		},
	})
}
