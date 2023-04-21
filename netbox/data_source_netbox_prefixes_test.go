package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixesDataSource_basic(t *testing.T) {

	testPrefixes := []string{"10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24"}
	testSlug := "prefixes_ds_basic"
	testVlanVids := []int{4093, 4094}
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test_prefix1" {
  prefix      = "%[2]s"
  status      = "active"
  description = "my-description"
  vrf_id      = netbox_vrf.test_vrf.id
  vlan_id     = netbox_vlan.test_vlan1.id
  tags        = [netbox_tag.test_tag1.slug]
}

resource "netbox_prefix" "test_prefix2" {
  prefix  = "%[3]s"
  status  = "active"
  vrf_id  = netbox_vrf.test_vrf.id
  vlan_id = netbox_vlan.test_vlan2.id
}

resource "netbox_prefix" "without_vrf_and_vlan" {
  prefix = "%[4]s"
  status = "active"
}

resource "netbox_vrf" "test_vrf" {
  name = "%[1]s_test_vrf"
}

resource "netbox_vlan" "test_vlan1" {
  name = "%[1]s_vlan1"
  vid  = %[5]d
}

resource "netbox_vlan" "test_vlan2" {
  name = "%[1]s_vlan2"
  vid  = %[6]d
}

resource "netbox_tag" "test_tag1" {
  name = "%[1]s"
}

resource "netbox_tag" "test_tag2" {
  name = "tag-with-no-associtions"
}

data "netbox_prefixes" "by_vrf" {
  depends_on = [netbox_prefix.test_prefix1, netbox_prefix.test_prefix2]
  filter {
    name  = "vrf_id"
    value = netbox_vrf.test_vrf.id
  }
}

data "netbox_prefixes" "by_vid" {
  depends_on = [netbox_prefix.test_prefix1, netbox_prefix.test_prefix2]
  filter {
    name  = "vlan_vid"
    value = "%[5]d"
  }
}

data "netbox_prefixes" "by_tag" {
  depends_on = [netbox_prefix.test_prefix1]
  filter {
    name  = "tag"
    value = "%[1]s"
  }
}

data "netbox_prefixes" "no_results" {
  depends_on = [netbox_prefix.test_prefix1]
  filter {
    name  = "tag"
    value = "tag-with-no-associtions"
  }
}

data "netbox_prefixes" "find_prefix_without_vrf_and_vlan" {
  depends_on = [netbox_prefix.without_vrf_and_vlan]
  filter {
    name  = "prefix"
    value = netbox_prefix.without_vrf_and_vlan.prefix
  }
}
`, testName, testPrefixes[0], testPrefixes[1], testPrefixes[2], testVlanVids[0], testVlanVids[1]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_vrf", "prefixes.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.by_vrf", "prefixes.1.vlan_vid", "netbox_vlan.test_vlan2", "vid"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.by_vid", "prefixes.0.vlan_vid", "netbox_vlan.test_vlan1", "vid"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_tag", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_tag", "prefixes.0.description", "my-description"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.no_results", "prefixes.#", "0"),
				),
			},
		},
	})
}
