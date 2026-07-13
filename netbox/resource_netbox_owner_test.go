package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxOwner_basic(t *testing.T) {
	testSlug := "owners"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_owner" "test_basic" {
  name = "%s"
	description = "This is my example resource"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_owner.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_owner.test_basic", "description", "This is my example resource"),
				),
			},
			{
				ResourceName:      "netbox_owner.test_basic",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func TestAccNetboxOwner_full(t *testing.T) {
	testSlug := "owners"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_owner_group" "test_full" {
  name = "%[1]s-group"
}

resource "netbox_group" "test_full" {
  name = "%[1]s-usergroup"
}

resource "netbox_user" "test_full" {
  username = "%[1]s"
  password = "Abcdefghijkl1"
  active   = true
}

resource "netbox_owner" "test_full" {
  name           = "%[1]s"
  group_id       = netbox_owner_group.test_full.id
  user_group_ids = [netbox_group.test_full.id]
  user_ids       = [netbox_user.test_full.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_owner.test_full", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_owner.test_full", "group_id", "netbox_owner_group.test_full", "id"),
					resource.TestCheckResourceAttr("netbox_owner.test_full", "user_group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_owner.test_full", "user_ids.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_owner.test_full",
				ImportState:       true,
				ImportStateVerify: false,
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
