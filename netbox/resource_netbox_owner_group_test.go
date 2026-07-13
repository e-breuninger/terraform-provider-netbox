package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/users"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxOwnerGroup_basic(t *testing.T) {
	testSlug := "owner_groups"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_owner_group" "test_basic" {
  name = "%s"
	description = "This is my example resource"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_owner_group.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_owner_group.test_basic", "description", "This is my example resource"),
				),
			},
			{
				ResourceName:      "netbox_owner_group.test_basic",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_owner_group", &resource.Sweeper{
		Name:         "netbox_owner_group",
		Dependencies: []string{"netbox_owner"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := users.NewUsersOwnerGroupsListParams()
			res, err := api.Users.UsersOwnerGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, ownerGroup := range res.GetPayload().Results {
				if strings.HasPrefix(*ownerGroup.Name, testPrefix) {
					deleteParams := users.NewUsersOwnerGroupsDeleteParams().WithID(ownerGroup.ID)
					_, err := api.Users.UsersOwnerGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an owner group")
				}
			}
			return nil
		},
	})
}
