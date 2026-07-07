package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVlanDataSource_basic(t *testing.T) {
	testVid := 4092
	testName := testAccGetTestName("vlan")
	setUp := testAccNetboxVlanSetUp(testVid, testName)
	extendedSetUp := testAccNetboxVlanSetUpMore(testVid, testVid-1, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
			},
			{
				Config:      setUp + testAccNetboxVlanDataNoResult,
				ExpectError: regexp.MustCompile("no vlan found matching filter"),
			},
			{
				Config: setUp + testAccNetboxVlanDataByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "id", "netbox_vlan.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanDataByVid(testVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "id", "netbox_vlan.test", "id"),
				),
			},

			{
				Config: setUp + extendedSetUp + testAccNetboxVlanDataByNameAndRole(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "id", "netbox_vlan.test", "id"),
				),
			},
			{
				Config: setUp + extendedSetUp + testAccNetboxVlanDataByVidAndTenant(testVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_vlan.test", "status", "active"),
					resource.TestCheckResourceAttr("data.netbox_vlan.test", "description", "Test"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "role", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "site", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "tenant", "netbox_tenant.test", "id"),
				),
			},
			{
				Config:      setUp + extendedSetUp + testAccNetboxVlanDataByName(testName),
				ExpectError: regexp.MustCompile("more than one vlan returned, specify a more narrow filter"),
			},
			{
				Config:      setUp + extendedSetUp + testAccNetboxVlanDataByVid(testVid),
				ExpectError: regexp.MustCompile("more than one vlan returned, specify a more narrow filter"),
			},
		},
	})
}

func TestAccNetboxVlanDataSource_customFields(t *testing.T) {
	testName := testAccGetTestName("vlan_ds_custom_fields")
	testField := fmt.Sprintf("vlan_ds_cf_%s", acctest.RandStringFromCharSet(10, "abcdefghijklmnopqrstuvwxyz"))
	testVid := acctest.RandIntRange(1000, 4000)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVlanDataSourceWithCustomFields(testName, testField, testVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan.by_vid", "id", "netbox_vlan.test_custom_fields", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan.by_vid", fmt.Sprintf("custom_fields.%s", testField), "match"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan.by_custom_fields", "id", "netbox_vlan.test_custom_fields", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan.by_custom_fields", fmt.Sprintf("custom_fields.%s", testField), "match"),
				),
			},
		},
	})
}

func testAccNetboxVlanDataSourceWithCustomFields(testName string, testField string, testVid int) string {
	return fmt.Sprintf(`
resource "netbox_custom_field" "test" {
	name          = "%[2]s"
	type          = "text"
	content_types = ["ipam.vlan"]
}

resource "netbox_vlan" "test_custom_fields" {
	name = "%[1]s"
	vid  = %[3]d
	tags = []

	custom_fields = {
		(netbox_custom_field.test.name) = "match"
	}
}

data "netbox_vlan" "by_vid" {
	depends_on = [netbox_vlan.test_custom_fields]
	vid        = netbox_vlan.test_custom_fields.vid
}

data "netbox_vlan" "by_custom_fields" {
	depends_on = [netbox_vlan.test_custom_fields]

	custom_fields = {
		(netbox_custom_field.test.name) = "match"
	}
}
`, testName, testField, testVid)
}

func testAccNetboxVlanSetUp(testVid int, testName string) string {
	return fmt.Sprintf(`
resource "netbox_ipam_role" "test" {
	name = "%[2]s"
}

resource "netbox_site" "test" {
	name = "%[2]s"
}

resource "netbox_tenant" "test" {
	name = "%[2]s"
}

resource "netbox_vlan" "test" {
	vid         = %[1]d
	name        = "%[2]s"
	description = "Test"
	role_id     = netbox_ipam_role.test.id
	site_id     = netbox_site.test.id
	status      = "active"
	tags        = []
	tenant_id   = netbox_tenant.test.id
}
`, testVid, testName)
}

func testAccNetboxVlanSetUpMore(testVid int, anotherVid int, testName string) string {
	return fmt.Sprintf(`
resource "netbox_vlan" "same_name" {
	vid  = %[1]d
	name = "%[3]s"
}

resource "netbox_vlan" "not_same" {
	vid  = %[2]d
	name = "%[3]s_unique"
}
`, testVid, anotherVid, testName)
}

const testAccNetboxVlanDataNoResult = `
data "netbox_vlan" "no_result" {
	name = "_no_result_"
}`

func testAccNetboxVlanDataByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_vlan" "test" {
	name = "%[1]s"
}`, testName)
}

func testAccNetboxVlanDataByVid(testVid int) string {
	return fmt.Sprintf(`
data "netbox_vlan" "test" {
	vid = "%[1]d"
}`, testVid)
}

func testAccNetboxVlanDataByNameAndRole(testName string) string {
	return fmt.Sprintf(`
data "netbox_vlan" "test" {
	name = "%[1]s"
	role = netbox_ipam_role.test.id
}`, testName)
}

func testAccNetboxVlanDataByVidAndTenant(testVid int) string {
	return fmt.Sprintf(`
data "netbox_vlan" "test" {
	vid    = "%[1]d"
	tenant = netbox_tenant.test.id
}`, testVid)
}
