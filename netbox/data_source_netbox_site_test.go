package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxSiteSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}

resource "netbox_asn" "test" {
  asn = 234
  rir_id = netbox_rir.test.id
}

resource "netbox_region" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  asn_ids = [netbox_asn.test.id]
  description = "Test"
  region_id = netbox_region.test.id
  tenant_id = netbox_tenant.test.id
  timezone = "Europe/Berlin"
}`, testName)
}

const testAccNetboxSiteNoResult = `
data "netbox_site" "test" {
  name = "_does_not_exist_"
}`

func testAccNetboxSiteByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_site" "test" {
  name = "%[1]s"
}`, testName)
}

func testAccNetboxSiteBySlug(testName string) string {
	return fmt.Sprintf(`
data "netbox_site" "test" {
  slug = "%[1]s"
}`, testName)
}

func TestAccNetboxSiteDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("site_ds_basic")
	setUp := testAccNetboxSiteSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "slug", testName),
				),
			},
			{
				Config:      setUp + testAccNetboxSiteNoResult,
				ExpectError: regexp.MustCompile("expected one site, but got 0"),
			},
			{
				Config: setUp + testAccNetboxSiteByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxSiteBySlug(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_site.test", "asn_ids.0", "netbox_asn.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_site.test", "description", "Test"),
					resource.TestCheckResourceAttr("data.netbox_site.test", "time_zone", "Europe/Berlin"),
					resource.TestCheckResourceAttr("data.netbox_site.test", "status", "active"),
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "region_id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "tenant_id", "netbox_tenant.test", "id"),
				),
			},
		},
	})
}
