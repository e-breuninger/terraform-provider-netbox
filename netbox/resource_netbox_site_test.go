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
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_rir" "test" {
  name = "%[1]s"
}

resource "netbox_asn" "test" {
  asn = 1338
  rir_id = netbox_rir.test.id
}

resource "netbox_site" "test" {
  name = "%[1]s"
  slug = "%[2]s"
  status = "planned"
  description = "%[1]s"
  facility = "%[1]s"
  physical_address = "%[1]s"
  shipping_address = "%[1]s"
  asn_ids = [netbox_asn.test.id]
  group_id = netbox_site_group.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_site.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_site.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "facility", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "physical_address", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "shipping_address", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_site.test", "asn_ids.0", "netbox_asn.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_site.test", "group_id", "netbox_site_group.test", "id"),
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
  tenant_id = netbox_tenant.test.id
  tags = ["%[1]s"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "slug", getSlug(testName)),
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
	testSlug := "site_detail"
	testName := testAccGetTestName(testSlug)
	testField := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
	name          = "%[1]s"
	type          = "text"
	content_types = ["dcim.site"]
}
resource "netbox_site" "test" {
  name          = "%[2]s"
  status        = "decommissioning"
  latitude      = "12.123456"
  longitude     = "-13.123456"
  timezone      = "Africa/Johannesburg"
  custom_fields = {"${netbox_custom_field.test.name}" = "81"}
}`, testField, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "status", "decommissioning"),
					resource.TestCheckResourceAttr("netbox_site.test", "custom_fields."+testField, "81"),
					resource.TestCheckResourceAttr("netbox_site.test", "timezone", "Africa/Johannesburg"),
					resource.TestCheckResourceAttr("netbox_site.test", "latitude", "12.123456"),
					resource.TestCheckResourceAttr("netbox_site.test", "longitude", "-13.123456"),
				),
			},
		},
	})
}

func TestAccNetboxSite_fieldUpdate(t *testing.T) {
	testSlug := "site_field_update"
	testName := testAccGetTestName(testSlug)
	testField := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
	name        = "%[2]s"
	description = "Test site description"
	physical_address = "Physical address"
	shipping_address = "Shipping address"

}`, testField, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "description", "Test site description"),
					resource.TestCheckResourceAttr("netbox_site.test", "physical_address", "Physical address"),
					resource.TestCheckResourceAttr("netbox_site.test", "shipping_address", "Shipping address"),
				)},
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
	name = "%[2]s"
}`, testField, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_site.test", "physical_address", ""),
					resource.TestCheckResourceAttr("netbox_site.test", "shipping_address", ""),
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
