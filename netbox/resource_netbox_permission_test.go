package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPermission_basic(t *testing.T) {
	testSlug := "user_permissions"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_permission" "test_basic" {
  name = "%s"
  description = "This is a terraform test."
  enabled = true
  object_types = ["ipam.prefix"]
  actions = ["add", "change"]
  users = [1]
  constraints = jsonencode([{
    "status" = "active"
  }]
    )
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "description", "This is a terraform test."),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "object_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "object_types.0", "ipam.prefix"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "actions.#", "2"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "actions.0", "add"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "actions.1", "change"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "users.#", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "users.0", "1"),
					resource.TestCheckResourceAttr("netbox_permission.test_basic", "constraints", "[{\"status\":\"active\"}]"),
				),
			},
			{
				ResourceName:      "netbox_permission.test_basic",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_permission", &resource.Sweeper{
		Name:         "netbox_permission",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := users.NewUsersPermissionsListParams()
			res, err := api.Users.UsersPermissionsList(params, nil)
			if err != nil {
				return err
			}
			for _, perm := range res.GetPayload().Results {
				if strings.HasPrefix(*perm.Name, testPrefix) {
					deleteParams := users.NewUsersPermissionsDeleteParams().WithID(perm.ID)
					_, err := api.Users.UsersPermissionsDelete(deleteParams, nil)
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
