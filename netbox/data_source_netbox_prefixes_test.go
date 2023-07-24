package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixesDataSource_basic(t *testing.T) {

	testPrefixes := []string{"10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24", "10.0.7.0/24", "10.0.8.0/25"}
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

resource "netbox_prefix" "test_prefix_two_tags_and_length_25" {
  prefix      = "%[6]s"
  status      = "active"
  description = "multiple-tag-prefix"
  tags        = [netbox_tag.test_tag3.slug, netbox_tag.test_tag4.slug]
}

resource "netbox_prefix" "without_vrf_and_vlan" {
  prefix = "%[5]s"
  status = "active"
}

resource "netbox_site" "test" {
  name = "site-%[1]s"
  timezone = "Europe/Berlin"
}

resource "netbox_prefix" "with_site_id" {
  prefix  = "%[5]s"
  status  = "active"
  site_id = netbox_site.test.id
}

resource "netbox_vrf" "test_vrf" {
  name = "%[1]s_test_vrf"
}

resource "netbox_vlan" "test_vlan1" {
  name = "%[1]s_vlan1"
  vid  = %[7]d
}

resource "netbox_vlan" "test_vlan2" {
  name = "%[1]s_vlan2"
  vid  = %[8]d
}

resource "netbox_tag" "test_tag1" {
  name = "%[1]s"
}

resource "netbox_tag" "test_tag2" {
  name = "tag-with-no-associtions"
}

resource "netbox_tag" "test_tag3" {
  name = "%[1]s-tag-3"
}

resource "netbox_tag" "test_tag4" {
  name = "%[1]s-tag-4"
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
    value = "%[7]d"
  }
}

data "netbox_prefixes" "by_tag" {
  depends_on = [netbox_prefix.test_prefix1]
  filter {
    name  = "tag"
    value = "%[1]s"
  }
}

data "netbox_prefixes" "by_mask_length" {
  depends_on = [netbox_prefix.test_prefix_two_tags_and_length_25]
  filter  {
    name = "mask_length"
    value = "25"
  }
}

data "netbox_prefixes" "by_mask_length_and_tag" {
  depends_on = [netbox_prefix.test_prefix1]
  filter  {
    name = "mask_length"
    value = "24"
  }
  filter {
    name  = "tag"
    value = "%[1]s"
  }
}

data "netbox_prefixes" "by_multiple_tags" {
  depends_on = [netbox_prefix.test_prefix_two_tags_and_length_25]
  filter {
    name  = "tag"
    value = netbox_tag.test_tag3.slug
  }
  filter {
    name  = "tag"
    value = netbox_tag.test_tag4.slug
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
data "netbox_prefixes" "find_prefix_with_site_id" {
  depends_on = [netbox_prefix.with_site_id]
  filter {
    name  = "site_id"
    value = netbox_site.test.id
  }
}
`, testName, testPrefixes[0], testPrefixes[1], testPrefixes[2], testPrefixes[3], testPrefixes[4], testVlanVids[0], testVlanVids[1]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_vrf", "prefixes.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.by_vrf", "prefixes.1.vlan_vid", "netbox_vlan.test_vlan2", "vid"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.by_vid", "prefixes.0.vlan_vid", "netbox_vlan.test_vlan1", "vid"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_tag", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_tag", "prefixes.0.description", "my-description"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.no_results", "prefixes.#", "0"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_site_id", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_site_id", "prefixes.0.prefix", "10.0.7.0/24"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_mask_length", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_mask_length", "prefixes.0.prefix", "10.0.8.0/25"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_multiple_tags", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_multiple_tags", "prefixes.0.prefix", "10.0.8.0/25"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_mask_length_and_tag", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_mask_length_and_tag", "prefixes.0.prefix", "10.0.4.0/24"),
				),
			},
		},
	})
}
