package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVlanGroupDataSource_basic(t *testing.T) {
	testSlug := "test_group"
	anotherSlug := "not_test_group"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxVlanGroupSetUp(testSlug, testName)
	extendedSetUp := testAccNetboxVlanGroupSetUpMore(testSlug, anotherSlug, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
			},
			{
				Config:      setUp + testAccNetboxVlanGroupDataNoResult,
				ExpectError: regexp.MustCompile("no result"),
			},
			{
				Config: setUp + testAccNetboxVlanGroupDataByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan_group.test", "id", "netbox_vlan_group.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanGroupDataBySlug(testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan_group.test", "id", "netbox_vlan_group.test", "id"),
				),
			},
			{
				Config: setUp + extendedSetUp + testAccNetboxVlanGroupDataByNameAndScope(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan_group.test", "id", "netbox_vlan_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "slug", testSlug),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "description", "Test"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "min_vid", "20"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "max_vid", "200"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_group.test", "scope_id", "netbox_site.test", "id"),
				),
			},
			{
				Config: setUp + extendedSetUp + testAccNetboxVlanGroupDataBySlugAndScope(testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan_group.test", "id", "netbox_vlan_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "slug", testSlug),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "description", "Test"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "min_vid", "20"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "max_vid", "200"),
					resource.TestCheckResourceAttr("data.netbox_vlan_group.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_group.test", "scope_id", "netbox_site.test", "id"),
				),
			},
			{
				Config:      setUp + extendedSetUp + testAccNetboxVlanGroupDataByName(testName),
				ExpectError: regexp.MustCompile("more than one result, specify a more narrow filter"),
			},
			{
				Config:      setUp + extendedSetUp + testAccNetboxVlanGroupDataBySlug(testSlug),
				ExpectError: regexp.MustCompile("more than one result, specify a more narrow filter"),
			},
		},
	})
}

func testAccNetboxVlanGroupSetUp(testSlug, testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
	name = "%[2]s"
}

resource "netbox_vlan_group" "test" {
	slug        = "%[1]s"
	name        = "%[2]s"
	description = "Test"
	min_vid     = 20
	max_vid     = 200
	scope_type  = "dcim.site"
	scope_id    = netbox_site.test.id
	tags        = []
}
`, testSlug, testName)
}

func testAccNetboxVlanGroupSetUpMore(testSlug, anotherSlug, testName string) string {
	return fmt.Sprintf(`
resource "netbox_vlan_group" "same_name" {
	slug    = "%[1]s"
	name    = "%[3]s"
	min_vid = 20
	max_vid = 200
}

resource "netbox_vlan_group" "not_same" {
	slug    = "%[2]s"
	name    = "%[3]s_unique"
	min_vid = 20
	max_vid = 200
}
`, testSlug, anotherSlug, testName)
}

const testAccNetboxVlanGroupDataNoResult = `
data "netbox_vlan_group" "no_result" {
	name = "_no_result_"
}`

func testAccNetboxVlanGroupDataByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_vlan_group" "test" {
	name = "%[1]s"
}`, testName)
}

func testAccNetboxVlanGroupDataBySlug(testSlug string) string {
	return fmt.Sprintf(`
data "netbox_vlan_group" "test" {
	slug = "%[1]s"
}`, testSlug)
}

func testAccNetboxVlanGroupDataByNameAndScope(testName string) string {
	return fmt.Sprintf(`
data "netbox_vlan_group" "test" {
	name       = "%[1]s"
	scope_type = "dcim.site"
	scope_id   = netbox_site.test.id
}`, testName)
}

func testAccNetboxVlanGroupDataBySlugAndScope(testSlug string) string {
	return fmt.Sprintf(`
data "netbox_vlan_group" "test" {
	slug       = "%[1]s"
	scope_type = "dcim.site"
	scope_id   = netbox_site.test.id
}`, testSlug)
}
