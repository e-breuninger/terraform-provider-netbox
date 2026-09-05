package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxOwnerDataSource_basic(t *testing.T) {
	testSlug := "owner_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxOwnerDependencies(testName) + fmt.Sprintf(`
resource "netbox_owner" "test" {
  name           = "%[1]s"
  description    = "This is my owner"
  group_id       = netbox_owner_group.test.id
  user_group_ids = [netbox_group.test.id]
  user_ids       = [netbox_user.test.id]
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
data "netbox_owner" "test" {
  depends_on = []

  name = netbox_owner.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_owner.test", "id", "netbox_owner.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_owner.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_owner.test", "description", "This is my owner"),
					resource.TestCheckResourceAttrPair("data.netbox_owner.test", "group_id", "netbox_owner_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_owner.test", "user_group_ids.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_owner.test", "user_ids.#", "1"),
				),
			},
		},
	})
}
