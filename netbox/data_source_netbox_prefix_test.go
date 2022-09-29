package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixDataSource_basic(t *testing.T) {

	testPrefixes := []string{"10.0.0.0/24", "10.0.1.0/24", "10.0.2.0/24", "10.0.3.0/24"}
	testSlug := "prefix_ds_basic"
	testVlanVids := []int{4093, 4094}
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "by_cidr" {
  prefix = "%[2]s"
  status = "active"
}

resource "netbox_vrf" "test" {
  name = "%[1]s"
}

resource "netbox_vlan" "test_id" {
  name = "%[4]s"
  vid  = %[6]d
}

resource "netbox_vlan" "test_vid" {
  name = "%[5]s"
  vid  = %[7]d
}

resource "netbox_prefix" "by_vrf" {
  prefix = "%[3]s"
  status = "active"
  vrf_id = netbox_vrf.test.id
}

resource "netbox_prefix" "by_vlan_id" {
  prefix  = "%[4]s"
  status  = "active"
  vlan_id = netbox_vlan.test_id.id
}

resource "netbox_prefix" "by_vlan_vid" {
  prefix   = "%[5]s"
  status   = "active"
  vlan_id = netbox_vlan.test_vid.id
}

data "netbox_prefix" "by_cidr" {
  depends_on = [netbox_prefix.by_cidr]
  cidr = "%[2]s"
}

data "netbox_prefix" "by_vrf_id" {
  depends_on = [netbox_prefix.by_vrf]
  vrf_id = netbox_vrf.test.id
}

data "netbox_prefix" "by_vlan_id" {
  depends_on = [netbox_prefix.by_vlan_id]
  vlan_id = netbox_vlan.test_id.id
}

data "netbox_prefix" "by_vlan_vid" {
  depends_on = [netbox_prefix.by_vlan_vid]
  vid = %[7]d
}
`, testName, testPrefixes[0], testPrefixes[1], testPrefixes[2], testPrefixes[3], testVlanVids[0], testVlanVids[1]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_cidr", "id", "netbox_prefix.by_cidr", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vrf_id", "id", "netbox_prefix.by_vrf", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_id", "id", "netbox_prefix.by_vlan_id", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.by_vlan_vid", "id", "netbox_prefix.by_vlan_vid", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
