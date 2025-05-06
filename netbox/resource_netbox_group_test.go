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

func TestAccNetboxGroup_basic(t *testing.T) {
	testSlug := "groups"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_group" "test_basic" {
  name = "%s"
	description = "This is my example resource"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_group.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_group.test_basic", "description", "This is my example resource"),
				),
			},
			{
				ResourceName:      "netbox_group.test_basic",
				ImportState:       true,
				ImportStateVerify: false,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_group", &resource.Sweeper{
		Name:         "netbox_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := users.NewUsersGroupsListParams()
			res, err := api.Users.UsersGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, group := range res.GetPayload().Results {
				if strings.HasPrefix(*group.Name, testPrefix) {
					deleteParams := users.NewUsersGroupsDeleteParams().WithID(group.ID)
					_, err := api.Users.UsersGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a group")
				}
			}
			return nil
		},
	})
}
