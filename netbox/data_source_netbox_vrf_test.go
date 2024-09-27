package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVrfDataSource_basic(t *testing.T) {
	testSlug := "vrf_ds_basic"
	testName := testAccGetTestName(testSlug)
	simpleSetup := testAccNetboxVrfSetUp(testName)
	advancedSetup := testAccNetboxVrfAdvancedSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: simpleSetup + testAccNetboxVrfData(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vrf.test", "id", "netbox_vrf.test", "id"),
				),
			},
			{
				Config:      simpleSetup + advancedSetup + testAccNetboxVrfData(testName),
				ExpectError: regexp.MustCompile("more than one vrf returned, specify a more narrow filter"),
			},
			{
				Config:      advancedSetup + testAccNetboxVrfDataWithoutTenantID(testName),
				ExpectError: regexp.MustCompile("more than one vrf returned, specify a more narrow filter"),
			},
			{
				Config: simpleSetup + advancedSetup + testAccNetboxVrfDataWithTenantID(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vrf.vrf_a", "id", "netbox_vrf.vrf_a", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vrf.vrf_b", "id", "netbox_vrf.vrf_b", "id"),
				),
			},
		},
	})
}

func testAccNetboxVrfSetUp(testName string) string {
	return fmt.Sprintf(`
resource"netbox_vrf" "test" {
  name = "%[1]s"
}`, testName)
}

func testAccNetboxVrfData(testName string) string {
	return fmt.Sprintf(`
data "netbox_vrf" "test" {
  depends_on = [netbox_vrf.test]
  name = "%[1]s"
}`, testName)
}

func testAccNetboxVrfAdvancedSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "tenant_a" {
	name = "%[1]s-a"
}

resource "netbox_tenant" "tenant_b" {
	name = "%[1]s-b"
}

resource "netbox_vrf" "vrf_a" {
	name      = "%[1]s"
	tenant_id = netbox_tenant.tenant_a.id
}

resource "netbox_vrf" "vrf_b" {
	name      = "%[1]s"
	tenant_id = netbox_tenant.tenant_b.id
}
`, testName)
}

func testAccNetboxVrfDataWithoutTenantID(testName string) string {
	return fmt.Sprintf(`
data "netbox_vrf" "vrf_a" {
	name       = "%[1]s"
}
`, testName)
}

func testAccNetboxVrfDataWithTenantID(testName string) string {
	return fmt.Sprintf(`
data "netbox_vrf" "vrf_a" {
	name       = "%[1]s"
	tenant_id  = netbox_tenant.tenant_a.id
}
data "netbox_vrf" "vrf_b" {
	name       = "%[1]s"
	tenant_id  = netbox_tenant.tenant_b.id
}
`, testName)
}
