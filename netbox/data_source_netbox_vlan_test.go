package netbox

import (
	"fmt"
	"regexp"
	"testing"

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
				ExpectError: regexp.MustCompile("expected one device type, but got 0"),
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
				ExpectError: regexp.MustCompile("expected one device type, but got 2"),
			},
			{
				Config:      setUp + extendedSetUp + testAccNetboxVlanDataByVid(testVid),
				ExpectError: regexp.MustCompile("expected one device type, but got 2"),
			},
		},
	})
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
