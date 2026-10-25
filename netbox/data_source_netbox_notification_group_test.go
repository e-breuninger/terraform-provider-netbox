package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxNotificationGroupDataSource_basic(t *testing.T) {
	testSlug := "notif_grp_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_group" "test" {
  name = "%[1]s"
}

resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Abcdefghijkl1"
  active   = true
}

resource "netbox_notification_group" "test" {
  name        = "%[1]s"
  description = "test description"
  group_ids   = [netbox_group.test.id]
  user_ids    = [netbox_user.test.id]
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
data "netbox_notification_group" "test" {
  depends_on = []

  name = netbox_notification_group.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_notification_group.test", "id", "netbox_notification_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_notification_group.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_notification_group.test", "description", "test description"),
					resource.TestCheckResourceAttr("data.netbox_notification_group.test", "group_ids.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_notification_group.test", "user_ids.#", "1"),
				),
			},
		},
	})
}
