package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixesDataSource_basic(t *testing.T) {

	testPrefixes := []string{"10.0.4.0/24", "10.0.5.0/24"}
	testSlug := "prefixes_ds_basic"
	testVlanVids := []int{4093, 4094}
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test_prefix1" {
  prefix = "%[2]s"
  status = "active"
  vrf_id = netbox_vrf.test_vrf.id
  vlan_id = netbox_vlan.test_vlan1.id
}

resource "netbox_prefix" "test_prefix2" {
  prefix = "%[3]s"
  status = "active"
  vrf_id = netbox_vrf.test_vrf.id
  vlan_id = netbox_vlan.test_vlan2.id
}

resource "netbox_vrf" "test_vrf" {
  name = "%[1]s_test_vrf"
}

resource "netbox_vlan" "test_vlan1" {
  name = "%[1]s_vlan1"
  vid  = %[4]d
}

resource "netbox_vlan" "test_vlan2" {
  name = "%[1]s_vlan2"
  vid  = %[5]d
}

data "netbox_prefixes" "by_vrf" {
  depends_on = [netbox_prefix.test_prefix1]
  filter{
	name  = "vrf_id"
    value = netbox_vrf.test_vrf.id
  }
}
`, testName, testPrefixes[0], testPrefixes[1], testVlanVids[0], testVlanVids[1]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_vrf", "prefixes.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.by_vrf", "prefixes.1.vlan_vid", "netbox_vlan.test_vlan2", "vid"),
				),
			},
		},
	})
}
