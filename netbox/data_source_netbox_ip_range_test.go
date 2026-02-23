package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpRangeDataSource_basic(t *testing.T) {
	testv4StartIP := "10.0.0.101/24"
	testv4EndIP := "10.0.0.150/24"
	testv6StartIP := "2001:db8:1:1::/112"
	testv6EndIP := "2001:db8:1:1::ffff/112"
	testSlug := "ip_range_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name = "%[1]s_vrf"
}

resource "netbox_tenant" "test" {
  name = "%[1]s_tenant"
}

resource "netbox_ipam_role" "test" {
  name = "%[1]s_role"
}

resource "netbox_tag" "test" {
  name = "%[1]s_role"
}

resource "netbox_ip_range" "testv4" {
  start_address = "%[2]s"
  end_address   = "%[3]s"
  vrf_id        = netbox_vrf.test.id
  tenant_id     = netbox_tenant.test.id
  role_id       = netbox_ipam_role.test.id
  description   = "%[1]s_description_testv4"
  status        = "active"
  tags          = [netbox_tag.test.name]
}

resource "netbox_ip_range" "testv6" {
  start_address = "%[4]s"
  end_address   = "%[5]s"
  description   = "%[1]s_description_testv6"
  status        = "reserved"
}

data "netbox_ip_range" "by_contains" {
  depends_on = [netbox_ip_range.testv4]
  contains   = "%[2]s"
  family     = 4
}

data "netbox_ip_range" "by_family" {
  depends_on = [netbox_ip_range.testv6]
  family     = 6
}

data "netbox_ip_range" "by_vrf_id" {
  depends_on = [netbox_ip_range.testv4]
  vrf_id     = netbox_vrf.test.id
  family     = 4
}

data "netbox_ip_range" "by_tenant_id" {
  depends_on = [netbox_ip_range.testv4]
  tenant_id  = netbox_tenant.test.id
  family     = 4
}

data "netbox_ip_range" "by_status" {
  depends_on = [netbox_ip_range.testv6]
  status     = "reserved"
}

data "netbox_ip_range" "by_role_id" {
  depends_on = [netbox_ip_range.testv4]
  role_id    = netbox_ipam_role.test.id
  family     = 4
}

data "netbox_ip_range" "by_tag" {
  depends_on = [netbox_ip_range.testv4]
  tag = netbox_tag.test.name
}

data "netbox_ip_range" "by_description" {
  depends_on  = [netbox_ip_range.testv4]
  description = "%[1]s_description_testv4"
}`, testName, testv4StartIP, testv4EndIP, testv6StartIP, testv6EndIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_contains", "id", "netbox_ip_range.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_family", "id", "netbox_ip_range.testv6", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_vrf_id", "id", "netbox_ip_range.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_tenant_id", "id", "netbox_ip_range.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_status", "id", "netbox_ip_range.testv6", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_role_id", "id", "netbox_ip_range.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_description", "id", "netbox_ip_range.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.by_tag", "id", "netbox_ip_range.testv4", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
