//go:build acctest

// Plural data sources page through results instead of returning only the first
// page. NetBox caps a page at MAX_PAGE_SIZE (1000 by default): against a NetBox
// with a smaller MAX_PAGE_SIZE than the objects below, a single request used to
// return just that many.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxPluralDataSourcePagination(t *testing.T) {
	testName := testAccGetTestName("paging")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  count = 5
  name  = "%[1]s-${count.index}"
}
data "netbox_tags" "all" {
  filters    = [{ name = "name__isw", value = "%[1]s-" }]
  depends_on = [netbox_tag.test]
}
data "netbox_tags" "limited" {
  filters    = [{ name = "name__isw", value = "%[1]s-" }]
  limit      = 4
  depends_on = [netbox_tag.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tags.all", "tags.#", "5"),
					resource.TestCheckResourceAttr("data.netbox_tags.limited", "tags.#", "4"),
				),
			},
		},
	})
}
