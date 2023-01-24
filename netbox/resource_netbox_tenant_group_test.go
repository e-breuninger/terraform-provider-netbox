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

func TestAccNetboxTenantGroup_basic(t *testing.T) {

	testSlug := "t_grp_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_tenant_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxTenantGroup_defaultSlug(t *testing.T) {

	testSlug := "tenant_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_tenant_group", &resource.Sweeper{
		Name:         "netbox_tenant_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := tenancy.NewTenancyTenantGroupsListParams()
			res, err := api.Tenancy.TenancyTenantGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, tenant := range res.GetPayload().Results {
				if strings.HasPrefix(*tenant.Name, testPrefix) {
					deleteParams := tenancy.NewTenancyTenantGroupsDeleteParams().WithID(tenant.ID)
					_, err := api.Tenancy.TenancyTenantGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a tenant group")
				}
			}
			return nil
		},
	})
}
