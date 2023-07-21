package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxContactRole_basic(t *testing.T) {
	testSlug := "contactrole"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact_role" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact_role.test", "name", testName),
				),
			},
			{
				ResourceName:      "netbox_contact_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_contact_role", &resource.Sweeper{
		Name:         "netbox_contact_role",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := tenancy.NewTenancyContactRolesListParams()
			res, err := api.Tenancy.TenancyContactRolesList(params, nil)
			if err != nil {
				return err
			}
			for _, contactrole := range res.GetPayload().Results {
				if strings.HasPrefix(*contactrole.Name, testPrefix) {
					deleteParams := tenancy.NewTenancyContactRolesDeleteParams().WithID(contactrole.ID)
					_, err := api.Tenancy.TenancyContactRolesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a contact role")
				}
			}
			return nil
		},
	})
}
