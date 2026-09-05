package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxOwnerDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_group" "test" {
  name = "%[1]s"
}

resource "netbox_user" "test" {
  username = "%[1]s"
  password = "Abcdefghijkl1"
  active   = true
}

resource "netbox_owner_group" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxOwner_basic(t *testing.T) {
	testSlug := "owner"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxOwnerDependencies(testName) + fmt.Sprintf(`
resource "netbox_owner" "test" {
  name           = "%[1]s"
  description    = "This is my owner"
  group_id       = netbox_owner_group.test.id
  user_group_ids = [netbox_group.test.id]
  user_ids       = [netbox_user.test.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_owner.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_owner.test", "description", "This is my owner"),
					resource.TestCheckResourceAttrPair("netbox_owner.test", "group_id", "netbox_owner_group.test", "id"),
					resource.TestCheckResourceAttr("netbox_owner.test", "user_group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_owner.test", "user_ids.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_owner.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_owner", &resource.Sweeper{
		Name:         "netbox_owner",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := users.NewUsersOwnersListParams()
			res, err := api.Users.UsersOwnersList(params, nil)
			if err != nil {
				return err
			}
			for _, owner := range res.GetPayload().Results {
				if strings.HasPrefix(*owner.Name, testPrefix) {
					deleteParams := users.NewUsersOwnersDeleteParams().WithID(owner.ID)
					_, err := api.Users.UsersOwnersDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an owner")
				}
			}
			return nil
		},
	})
}
