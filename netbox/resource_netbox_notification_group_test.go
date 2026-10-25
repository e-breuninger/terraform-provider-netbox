package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxNotificationGroup_basic(t *testing.T) {
	testSlug := "notif_grp"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
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
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_notification_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_notification_group.test", "description", "test description"),
					resource.TestCheckResourceAttr("netbox_notification_group.test", "group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_notification_group.test", "user_ids.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_group" "test" {
  name = "%[1]s"
}

resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Abcdefghijkl1"
  active   = true
}

resource "netbox_notification_group" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_notification_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_notification_group.test", "group_ids.#", "0"),
					resource.TestCheckResourceAttr("netbox_notification_group.test", "user_ids.#", "0"),
				),
			},
			{
				ResourceName:      "netbox_notification_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_notification_group", &resource.Sweeper{
		Name:         "netbox_notification_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := extras.NewExtrasNotificationGroupsListParams()
			res, err := api.Extras.ExtrasNotificationGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, notificationGroup := range res.GetPayload().Results {
				if strings.HasPrefix(*notificationGroup.Name, testPrefix) {
					deleteParams := extras.NewExtrasNotificationGroupsDeleteParams().WithID(notificationGroup.ID)
					_, err := api.Extras.ExtrasNotificationGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a notification group")
				}
			}
			return nil
		},
	})
}
