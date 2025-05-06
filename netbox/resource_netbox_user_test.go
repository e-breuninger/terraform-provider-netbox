package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxUser_basic(t *testing.T) {
	testSlug := "users"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_user" "test_basic" {
  username = "%s"
  password = "Abcdefghijkl1"
  active = true
  staff = true
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_user.test_basic", "username", testName),
					resource.TestCheckResourceAttr("netbox_user.test_basic", "active", "true"),
					resource.TestCheckResourceAttr("netbox_user.test_basic", "staff", "true"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_user" "test_basic" {
  username = "%s"
  password = "Abcdefghijkl1"
	email = "foo@bar.com"
	first_name = "Hannah"
	last_name = "Acker"
  active = true
  staff = true
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_user.test_basic", "username", testName),
					resource.TestCheckResourceAttr("netbox_user.test_basic", "email", "foo@bar.com"),
					resource.TestCheckResourceAttr("netbox_user.test_basic", "first_name", "Hannah"),
					resource.TestCheckResourceAttr("netbox_user.test_basic", "last_name", "Acker"),
					resource.TestCheckResourceAttr("netbox_user.test_basic", "active", "true"),
					resource.TestCheckResourceAttr("netbox_user.test_basic", "staff", "true"),
				),
			},
			{
				ResourceName:      "netbox_user.test_basic",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func TestAccNetboxUser_group(t *testing.T) {
	testSlug := "users"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_group" "test_group" {
	  name = "%[1]s-group"
}

resource "netbox_user" "test_group" {
	  username = "%[1]s"
	  password = "Abcdefghijkl1"
	  active = true
	  staff = true
	  group_ids = [netbox_group.test_group.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_user.test_group", "username", testName),
					resource.TestCheckResourceAttr("netbox_user.test_group", "active", "true"),
					resource.TestCheckResourceAttr("netbox_user.test_group", "staff", "true"),
					resource.TestCheckResourceAttr("netbox_user.test_group", "group_ids.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_user.test_group",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_user", &resource.Sweeper{
		Name:         "netbox_user",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := users.NewUsersUsersListParams()
			res, err := api.Users.UsersUsersList(params, nil)
			if err != nil {
				return err
			}
			for _, user := range res.GetPayload().Results {
				if strings.HasPrefix(*user.Username, testPrefix) {
					deleteParams := users.NewUsersUsersDeleteParams().WithID(user.ID)
					_, err := api.Users.UsersUsersDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a user")
				}
			}
			return nil
		},
	})
}
