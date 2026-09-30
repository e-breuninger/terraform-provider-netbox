//go:build acctest

// The site<->ASN relation is one set of links that either end may declare: netbox_site.asn_ids or
// netbox_asn.site_ids. An end that leaves its attribute unset never sends it, so the other end's
// links survive its updates; an empty set clears them.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccNetboxSiteASN_fromASN: the ASN declares the link and the site says nothing. The link
// survives an update of the site, and the site can be updated in the apply that destroys the ASN.
func TestAccNetboxSiteASN_fromASN(t *testing.T) {
	testName := testAccGetTestName("site_asn_from_asn")
	asn := acctest.RandIntRange(4200100000, 4200190000)
	site := func(description string, dependsOn string) string {
		return fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name        = "%[1]s"
  description = "%[2]s"
}
data "netbox_site" "test" {
  id         = netbox_site.test.id
  depends_on = [%[3]s]
}`, testName, description, dependsOn)
	}
	attached := fmt.Sprintf(`
resource "netbox_asn" "test" {
  asn         = %[2]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
  site_ids    = [netbox_site.test.id]
}`, testName, asn)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: site("one", "netbox_asn.test") + attached,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_asn.test", "site_ids.#", "1"),
					resource.TestCheckTypeSetElemAttrPair("netbox_asn.test", "site_ids.*", "netbox_site.test", "id"),
					resource.TestCheckTypeSetElemAttrPair("data.netbox_site.test", "asn_ids.*", "netbox_asn.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_asn.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// The site changes and still says nothing about its ASNs: the link stays.
				Config: site("two", "netbox_asn.test") + attached,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "description", "two"),
					resource.TestCheckResourceAttr("netbox_asn.test", "site_ids.#", "1"),
					resource.TestCheckTypeSetElemAttrPair("data.netbox_site.test", "asn_ids.*", "netbox_asn.test", "id"),
				),
			},
			{
				// The ASN goes away while the site changes again: the site must not send the
				// ASN it last saw.
				Config: site("three", "netbox_site.test"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "description", "three"),
					resource.TestCheckNoResourceAttr("data.netbox_site.test", "asn_ids"),
				),
			},
		},
	})
}

// TestAccNetboxSiteASN_fromSite: the site declares the link. Dropping asn_ids leaves the link
// alone; an empty set removes it, and the ASN reads each state back.
func TestAccNetboxSiteASN_fromSite(t *testing.T) {
	testName := testAccGetTestName("site_asn_from_site")
	asn := acctest.RandIntRange(4200200000, 4200290000)
	deps := fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "test" {
  asn         = %[2]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
}
data "netbox_asn" "test" {
  id         = netbox_asn.test.id
  depends_on = [netbox_site.test]
}`, testName, asn)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_site" "test" {
  name    = "%[1]s"
  asn_ids = [netbox_asn.test.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "1"),
					resource.TestCheckTypeSetElemAttrPair("data.netbox_asn.test", "site_ids.*", "netbox_site.test", "id"),
				),
			},
			{
				// Unset: the site stops declaring the link and leaves it in place.
				Config: deps + fmt.Sprintf(`
resource "netbox_site" "test" {
  name        = "%[1]s"
  description = "unset"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "1"),
					resource.TestCheckTypeSetElemAttrPair("data.netbox_asn.test", "site_ids.*", "netbox_site.test", "id"),
				),
			},
			{
				// An empty set removes the link.
				Config: deps + fmt.Sprintf(`
resource "netbox_site" "test" {
  name    = "%[1]s"
  asn_ids = []
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "0"),
					resource.TestCheckNoResourceAttr("data.netbox_asn.test", "site_ids"),
				),
			},
		},
	})
}

// TestAccNetboxSiteASN_bothEnds: one link is declared by its site and another by its ASN, on
// different sites and ASNs. Neither resource touches the other's link.
func TestAccNetboxSiteASN_bothEnds(t *testing.T) {
	testName := testAccGetTestName("site_asn_both")
	asn := acctest.RandIntRange(4200300000, 4200390000)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "declared_by_site" {
  asn         = %[2]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
}
resource "netbox_site" "declares" {
  name    = "%[1]s-declares"
  asn_ids = [netbox_asn.declared_by_site.id]
}
resource "netbox_site" "silent" {
  name = "%[1]s-silent"
}
resource "netbox_asn" "declares" {
  asn         = %[3]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
  site_ids    = [netbox_site.silent.id]
}`, testName, asn, asn+1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site.declares", "asn_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_asn.declares", "site_ids.#", "1"),
				),
			},
		},
	})
}

// TestAccNetboxSiteASN_conflictingEnds: a site and an ASN both declare the ASNs of the same site,
// with different contents. Each apply of one undoes the other, so the plan never comes up empty.
func TestAccNetboxSiteASN_conflictingEnds(t *testing.T) {
	testName := testAccGetTestName("site_asn_conflict")
	asn := acctest.RandIntRange(4200400000, 4200490000)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "listed" {
  asn         = %[2]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
}
resource "netbox_site" "test" {
  name    = "%[1]s"
  asn_ids = [netbox_asn.listed.id]
}
resource "netbox_asn" "unlisted" {
  asn         = %[3]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
  site_ids    = [netbox_site.test.id]
}`, testName, asn, asn+1),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
