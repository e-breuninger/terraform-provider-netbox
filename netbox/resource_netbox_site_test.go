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
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
  tenant_id = netbox_tenant.test.id
  tags = ["%[1]s"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "slug", testName),
					resource.TestCheckResourceAttrPair("netbox_site.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_site.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_site.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_site.test", "tags.0", testName),
				),
			},
		},
	})
}

func TestAccNetboxSite_customFields(t *testing.T) {

	testSlug := "site_customFields"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
	name = "issue"
	type = "text"
	content_types = ["dcim.site"]
}
resource "netbox_site" "test" {
  name = "%s"
  slug = "slug"
  status = "active"
  description = "%[1]s"
  facility = "%[1]s"
  asn = 1337
  custom_fields = {"${netbox_custom_field.test.name}" = "76"}
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "custom_fields.issue", "76"),
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
