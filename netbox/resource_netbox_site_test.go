package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxSite_basic(t *testing.T) {

	testSlug := "site_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%s"
  slug = "%s"
  status = "active"
  description = "%[1]s"
  facility = "%[1]s"
  asn = 1337
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_site.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_site.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "facility", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "asn", "1337"),
				),
			},
			{
				ResourceName:      "netbox_site.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxSite_defaultSlug(t *testing.T) {

	testSlug := "site_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%s"
  status = "active"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "slug", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "status", "active"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_site", &resource.Sweeper{
		Name:         "netbox_site",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimSitesListParams()
			res, err := api.Dcim.DcimSitesList(params, nil)
			if err != nil {
				return err
			}
			for _, site := range res.GetPayload().Results {
				if strings.HasPrefix(*site.Name, testPrefix) {
					deleteParams := dcim.NewDcimSitesDeleteParams().WithID(site.ID)
					_, err := api.Dcim.DcimSitesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a site")
				}
			}
			return nil
		},
	})
}
