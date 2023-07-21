package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRole_basic(t *testing.T) {
	testSlug := "role_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ipam_role" "test_basic" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipam_role.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipam_role.test_basic", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_ipam_role.test_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxRole_extended(t *testing.T) {
	testSlug := "role_extended"
	testName := testAccGetTestName(testSlug)
	testWeight := "55"
	testDescription := "Test description"
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ipam_role" "role_extended" {
  name = "%[1]s"
  slug = "%[2]s"
  weight = "%[3]s"
  description = "%[4]s"

}`, testName, randomSlug, testWeight, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipam_role.role_extended", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipam_role.role_extended", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_ipam_role.role_extended", "weight", testWeight),
					resource.TestCheckResourceAttr("netbox_ipam_role.role_extended", "description", testDescription),
				),
			},
			{
				ResourceName:      "netbox_ipam_role.role_extended",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ipam_role", &resource.Sweeper{
		Name:         "netbox_ipam_role",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamRolesListParams()
			res, err := api.Ipam.IpamRolesList(params, nil)
			if err != nil {
				return err
			}
			for _, role := range res.GetPayload().Results {
				if strings.HasPrefix(*role.Name, testPrefix) {
					deleteParams := ipam.NewIpamRolesDeleteParams().WithID(role.ID)
					_, err := api.Ipam.IpamRolesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a role")
				}
			}
			return nil
		},
	})
}
