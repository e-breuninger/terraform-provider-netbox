package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRegionDataSource_basic(t *testing.T) {
	testSlug := "region_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_region" "test" {
  name = "%[1]s"
}
data "netbox_region" "test" {
  depends_on = [netbox_region.test]
  filter {
	name = "%[1]s"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_region.test", "id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_region.test", "slug", "netbox_region.test", "slug"),
				),
			},
		},
	})
}

func TestAccNetboxRegionDataSource_parent(t *testing.T) {
	testSlug := "region_ds_parent"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test-child" {
  name             = "%[1]s-child"
  parent_region_id = netbox_region.test.id
}
data "netbox_region" "test-child" {
  depends_on = [netbox_region.test-child]
  filter {
	name = "%[1]s-child"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_region.test-child", "id", "netbox_region.test-child", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_region.test-child", "parent_region_id", "netbox_region.test", "id"),
				),
			},
		},
	})
}
