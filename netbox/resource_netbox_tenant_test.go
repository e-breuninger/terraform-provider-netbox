package netbox

import (
	"fmt"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"log"
	"strings"
	"testing"
)

func TestAccNetboxTenant_basic(t *testing.T) {

	testSlug := "tenant_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant.test", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_tenant.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxTenant_defaultSlug(t *testing.T) {

	testSlug := "tenant_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant.test", "slug", testName),
				),
			},
		},
	})
}

func TestAccNetboxTenant_tags(t *testing.T) {

	testSlug := "tenant_tags"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test_tags" {
  name = "%s"
  tags = ["boo"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test_tags", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant.test_tags", "tags.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test_tags" {
  name = "%s"
  tags = ["boo", "foo"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test_tags", "tags.#", "2"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test_tags" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test_tags", "tags.#", "0"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_tenant", &resource.Sweeper{
		Name:         "netbox_tenant",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBox)
			params := tenancy.NewTenancyTenantsListParams()
			res, err := api.Tenancy.TenancyTenantsList(params, nil)
			if err != nil {
				return err
			}
			for _, tenant := range res.GetPayload().Results {
				if strings.HasPrefix(*tenant.Name, testPrefix) {
					deleteParams := tenancy.NewTenancyTenantsDeleteParams().WithID(tenant.ID)
					_, err := api.Tenancy.TenancyTenantsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a tenant")
				}
			}
			return nil
		},
	})
}
