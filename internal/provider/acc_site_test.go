//go:build acctest

package provider_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxSite_basic(t *testing.T) {
	testSlug := "site_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
  slug = "%[2]s"
  status = "planned"
  description = "%[1]s"
  comments = "%[1]s"
  facility = "%[1]s"
  physical_address = "%[1]s"
  shipping_address = "%[1]s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_site.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_site.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "comments", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "facility", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "physical_address", testName),
					resource.TestCheckResourceAttr("netbox_site.test", "shipping_address", testName),
					resource.TestCheckResourceAttrSet("netbox_site.test", "created"),
					resource.TestCheckResourceAttrSet("netbox_site.test", "last_updated"),
					resource.TestCheckResourceAttrSet("netbox_site.test", "url"),
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

func TestAccNetboxSite_defaultSlugAndStatus(t *testing.T) {
	testSlug := "site_defaults"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_site.test", "status", "active"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "tags"),
				),
			},
			// Renaming follows through to the derived slug.
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%s renamed"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "slug", getSlug(testName+" renamed")),
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
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
	name         = "%[1]s"
	type         = "text"
	object_types = ["dcim.site"]
}
resource "netbox_site" "test" {
  name          = "%[2]s"
  status        = "decommissioning"
  latitude      = 12.123456
  longitude     = -13.123456
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
			// Removing the key clears the value in NetBox (the plan converges); the attribute is computed,
			// so {} rather than omission clears it.
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
	name         = "%[1]s"
	type         = "text"
	object_types = ["dcim.site"]
}
resource "netbox_site" "test" {
  name          = "%[2]s"
  custom_fields = {}
  depends_on    = [netbox_custom_field.test]
}`, testField, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "custom_fields.%", "0"),
				),
			},
		},
	})
}

func TestAccNetboxSite_fieldUpdate(t *testing.T) {
	testSlug := "site_field_update"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
	name        = "%[1]s"
	description = "Test site description"
	comments = "Test comment"
	physical_address = "Physical address"
	shipping_address = "Shipping address"
	facility      = "Facility"
	timezone      = "Europe/Berlin"
	latitude      = 12.123456
	longitude     = -13.123456
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "description", "Test site description"),
					resource.TestCheckResourceAttr("netbox_site.test", "comments", "Test comment"),
					resource.TestCheckResourceAttr("netbox_site.test", "physical_address", "Physical address"),
					resource.TestCheckResourceAttr("netbox_site.test", "shipping_address", "Shipping address"),
					resource.TestCheckResourceAttr("netbox_site.test", "facility", "Facility"),
					resource.TestCheckResourceAttr("netbox_site.test", "timezone", "Europe/Berlin"),
					resource.TestCheckResourceAttr("netbox_site.test", "latitude", "12.123456"),
					resource.TestCheckResourceAttr("netbox_site.test", "longitude", "-13.123456"),
				)},
			// Unset optional attributes are cleared in NetBox and read back as null.
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
	name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_site.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "physical_address"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "shipping_address"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "facility"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "timezone"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "latitude"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "longitude"),
				),
			},
		},
	})
}

func TestAccNetboxSite_references(t *testing.T) {
	testSlug := "site_refs"
	testName := testAccGetTestName(testSlug)
	asn := acctest.RandIntRange(64512, 65500)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "test" {
  asn    = %[2]d
  rir_id = netbox_rir.test.id
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name      = "%[1]s"
  region_id = netbox_region.test.id
  group_id  = netbox_site_group.test.id
  tenant_id = netbox_tenant.test.id
  asn_ids   = [netbox_asn.test.id]
  tags      = [netbox_tag.test.slug]
}`, testName, asn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_site.test", "region_id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_site.test", "group_id", "netbox_site_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_site.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_site.test", "tags.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_site.test", "tags.0", "netbox_tag.test", "slug"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "test" {
  asn    = %[2]d
  rir_id = netbox_rir.test.id
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}`, testName, asn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_site.test", "region_id"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "group_id"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "asn_ids"),
					resource.TestCheckNoResourceAttr("netbox_site.test", "tags"),
				),
			},
		},
	})
}

// Data sources.

func testAccNetboxSiteSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  description = "Test"
  region_id = netbox_region.test.id
  tenant_id = netbox_tenant.test.id
  timezone = "Europe/Berlin"
  facility = "Facility"
  physical_address = "Platz d. Republik 1, 10557 Berlin, Germany"
}`, testName)
}

func TestAccNetboxSiteDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("site_ds_basic")
	setUp := testAccNetboxSiteSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check:  resource.TestCheckResourceAttr("netbox_site.test", "slug", getSlug(testName)),
			},
			{
				Config: setUp + `
data "netbox_site" "test" {
  name = "_does_not_exist_"
}`,
				ExpectError: regexp.MustCompile("Lookup did not match exactly one netbox_site"),
			},
			{
				Config: setUp + `
data "netbox_site" "test" {
}`,
				ExpectError: regexp.MustCompile("Missing lookup attribute"),
			},
			{
				Config: setUp + fmt.Sprintf(`
data "netbox_site" "test" {
  name = "%s"
  depends_on = [netbox_site.test]
}`, testName),
				Check: resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
			},
			{
				Config: setUp + fmt.Sprintf(`
data "netbox_site" "test" {
  slug = "%s"
  depends_on = [netbox_site.test]
}`, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_site.test", "description", "Test"),
					resource.TestCheckResourceAttr("data.netbox_site.test", "timezone", "Europe/Berlin"),
					resource.TestCheckResourceAttr("data.netbox_site.test", "status", "active"),
					resource.TestCheckResourceAttr("data.netbox_site.test", "physical_address", "Platz d. Republik 1, 10557 Berlin, Germany"),
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "region_id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_site.test", "tenant_id", "netbox_tenant.test", "id"),
				),
			},
			{
				Config: setUp + `
data "netbox_site" "test" {
  id = netbox_site.test.id
}`,
				Check: resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
			},
			{
				Config: setUp + `
data "netbox_site" "test" {
  facility = netbox_site.test.facility
}`,
				Check: resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
			},
			{
				Config: setUp + fmt.Sprintf(`
data "netbox_site" "test" {
  name_contains = "%s"
  depends_on = [netbox_site.test]
}`, strings.ToUpper(testName[len(testName)-6:])),
				Check: resource.TestCheckResourceAttrPair("data.netbox_site.test", "id", "netbox_site.test", "id"),
			},
		},
	})
}

func TestAccNetboxSitesDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("sites_ds")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// The regex is validated at plan time. First, so the post-test destroy runs
				// against the valid config of the last step.
				Config: `
data "netbox_sites" "bad" {
  name_regex = "("
}`,
				ExpectError: regexp.MustCompile("Invalid regular expression"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "a" {
  name   = "%[1]s-a"
  status = "planned"
}
resource "netbox_site" "b" {
  name   = "%[1]s-b"
  status = "planned"
}
data "netbox_sites" "by_prefix" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_site.a, netbox_site.b]
}
data "netbox_sites" "limited" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  limit         = 1
  depends_on    = [netbox_site.a, netbox_site.b]
}
data "netbox_sites" "none" {
  filters = [
    { name = "name", value = "_does_not_exist_" },
  ]
}
data "netbox_sites" "regex" {
  name_regex = "^%[1]s-[a]$"
  depends_on = [netbox_site.a, netbox_site.b]
}
data "netbox_sites" "regex_limited" {
  name_regex = "^%[1]s-"
  limit      = 1
  depends_on = [netbox_site.a, netbox_site.b]
}
data "netbox_sites" "lookups" {
  filters = [
    { name = "name__isw", value = "%[1]s-" },
    { name = "name__n", value = "%[1]s-a" },
    { name = "id__empty", value = "false" },
  ]
  depends_on = [netbox_site.a, netbox_site.b]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_sites.by_prefix", "sites.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_sites.by_prefix", "sites.0.status", "planned"),
					resource.TestCheckResourceAttr("data.netbox_sites.limited", "sites.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_sites.none", "sites.#", "0"),
					// name_regex matches client-side after fetching every site; limit applies to
					// the matches, not to what the API returned.
					resource.TestCheckResourceAttr("data.netbox_sites.regex", "sites.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_sites.regex", "sites.0.id", "netbox_site.a", "id"),
					resource.TestCheckResourceAttr("data.netbox_sites.regex_limited", "sites.#", "1"),
					// Lookup variants: starts-with matches both, negation drops a, the boolean
					// __empty parses.
					resource.TestCheckResourceAttr("data.netbox_sites.lookups", "sites.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_sites.lookups", "sites.0.id", "netbox_site.b", "id"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_site",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimSitesList(dcim.NewDcimSitesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimSitesDestroy(dcim.NewDcimSitesDestroyParams().WithID(id), nil)
			return err
		})
}
