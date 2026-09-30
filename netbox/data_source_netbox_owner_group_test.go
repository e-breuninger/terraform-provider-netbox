package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxOwnerGroupDataSource_basic(t *testing.T) {
	testSlug := "owner_group_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_owner_group" "test" {
  name        = "%s"
  description = "This is my owner group"
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
data "netbox_owner_group" "test" {
  depends_on = []

  name = netbox_owner_group.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_owner_group.test", "id", "netbox_owner_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_owner_group.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_owner_group.test", "description", "This is my owner group"),
				),
			},
		},
	})
}
