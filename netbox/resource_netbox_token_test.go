package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxToken_basic(t *testing.T) {
	testSlug := "users"
	testName := testAccGetTestName(testSlug)
	testToken := testAccGetTestToken()
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_user" "test" {
  username = "%s"
  password = "Abcdefghijkl1"
}

resource "netbox_token" "test_basic" {
  user_id       = netbox_user.test.id
  key           = "%s"
  allowed_ips   = ["2.4.8.16/32"]
  write_enabled = false
  description   = "Netbox Test Basic Token"
}`, testName, testToken),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_token.test_basic", "key", testToken),
					resource.TestCheckResourceAttr("netbox_token.test_basic", "allowed_ips.#", "1"),
					resource.TestCheckResourceAttr("netbox_token.test_basic", "allowed_ips.0", "2.4.8.16/32"),
					resource.TestCheckResourceAttr("netbox_token.test_basic", "write_enabled", "false"),
					resource.TestCheckResourceAttr("netbox_token.test_basic", "description", "Netbox Test Basic Token"),
				),
			},
			{
				ResourceName:      "netbox_token.test_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_token", &resource.Sweeper{
		Name:         "netbox_token",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
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
					log.Print("[DEBUG] Deleted a token")
				}
			}
			return nil
		},
	})
}
