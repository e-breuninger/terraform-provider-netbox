package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixDataSource_basic(t *testing.T) {
	testv4Prefix := "10.0.0.0/24"
	testv6Prefix := "2000::/64"
	testSlug := "prefix_ds_basic"
	testVlanVid := 4090
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name = "%[1]s_vrf"
}

resource "netbox_vlan" "test" {
  name = "%[1]s_vlan_test_id"
  vid  = %[4]d
}

resource "netbox_site" "test" {
  name = "%[1]s_site"
}

resource "netbox_prefix" "testv4" {
  prefix = "%[2]s"
  status = "active"
  vrf_id = netbox_vrf.test.id
  vlan_id = netbox_vlan.test.id
  site_id = netbox_site.test.id
  description = "%[1]s_description_test_idv4"
}

resource "netbox_prefix" "testv6" {
  prefix = "%[3]s"
  status = "active"
  vrf_id = netbox_vrf.test.id
  vlan_id = netbox_vlan.test.id
  site_id = netbox_site.test.id
  description = "%[1]s_description_test_idv6"
}

data "netbox_prefix" "by_description" {
  description = netbox_prefix.testv4.description
}

data "netbox_prefix" "by_cidr" {
  depends_on = [netbox_prefix.testv4]
  cidr = "%[2]s"
}

data "netbox_prefix" "by_vrf_id" {
  depends_on = [netbox_prefix.testv4]
  vrf_id = netbox_vrf.test.id
  family = 4
}

data "netbox_prefix" "by_vlan_id" {
  depends_on = [netbox_prefix.testv4]
  vlan_id = netbox_vlan.test.id
  family  = 4
}

data "netbox_prefix" "by_vlan_vid" {
  depends_on = [netbox_prefix.testv4]
  vlan_vid = %[4]d
  family   = 4
}

data "netbox_prefix" "by_prefix" {
  depends_on = [netbox_prefix.testv4]
  prefix = "%[2]s"
}

data "netbox_prefix" "by_site_id" {
  depends_on = [netbox_prefix.testv4]
  site_id = netbox_site.test.id
  family  = 4
}

data "netbox_prefix" "by_family" {
  depends_on = [netbox_prefix.testv6]
	family     = 6
}

`, testName, testv4Prefix, testv6Prefix, testVlanVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_prefix", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_description", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_cidr", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vrf_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_vid", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_site_id", "id", "netbox_prefix.testv4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_family", "id", "netbox_prefix.testv6", "id"),
				),
			},
		},
	})
}
