package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixesDataSource_basic(t *testing.T) {
	testPrefixes := []string{"10.0.4.0/24", "10.0.5.0/24", "10.0.6.0/24", "10.0.7.0/24", "10.0.8.0/24", "10.0.9.0/24"}
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
  status  = "container"
  vrf_id  = netbox_vrf.test_vrf.id
  vlan_id = netbox_vlan.test_vlan2.id
}

resource "netbox_prefix" "without_vrf_and_vlan" {
  prefix = "%[4]s"
  status = "container"
}

resource "netbox_tenant" "test" {
  name = "%[1]s_tenant"
}

resource "netbox_prefix" "with_tenant_id" {
  prefix    = "%[5]s"
  status    = "container"
  tenant_id = netbox_tenant.test.id
}

resource "netbox_site" "test" {
  name     = "site-%[1]s"
  timezone = "Europe/Berlin"
}

resource "netbox_prefix" "with_site_id" {
  prefix  = "%[6]s"
  status  = "container"
  site_id = netbox_site.test.id
}

resource "netbox_site" "test2" {
  name     = "site2-%[1]s"
  timezone = "Europe/Berlin"
}

resource "netbox_prefix" "with_container" {
  prefix  = "%[9]s"
  status  = "container"
  site_id = netbox_site.test2.id
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
  name = "tag-with-no-associations"
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

data "netbox_prefixes" "by_status" {
  depends_on = [netbox_prefix.test_prefix1]
  filter {
    name  = "status"
    value = "active"
  }
}

data "netbox_prefixes" "no_results" {
  depends_on = [netbox_prefix.test_prefix1]
  filter {
    name  = "tag"
    value = netbox_tag.test_tag2.name
  }
}

data "netbox_prefixes" "find_prefix_without_vrf_and_vlan" {
  depends_on = [netbox_prefix.without_vrf_and_vlan]
  filter {
    name  = "prefix"
    value = netbox_prefix.without_vrf_and_vlan.prefix
  }
}

data "netbox_prefixes" "find_prefix_with_tenant_id" {
  depends_on = [netbox_prefix.with_tenant_id]
  filter {
    name  = "tenant_id"
    value = netbox_tenant.test.id
  }
}

data "netbox_prefixes" "find_prefix_with_site_id" {
  depends_on = [netbox_prefix.with_site_id]
  filter {
    name  = "site_id"
    value = netbox_site.test.id
  }
}

data "netbox_prefixes" "find_prefix_with_contains" {
  depends_on = [netbox_prefix.with_container]
  filter {
    name  = "contains"
    value = "10.0.9.50"
  }
}

`, testName, testPrefixes[0], testPrefixes[1], testPrefixes[2], testPrefixes[3], testPrefixes[4], testVlanVids[0], testVlanVids[1], testPrefixes[5]),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_vrf", "prefixes.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.by_vrf", "prefixes.1.vlan_vid", "netbox_vlan.test_vlan2", "vid"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.by_vid", "prefixes.0.vlan_vid", "netbox_vlan.test_vlan1", "vid"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_tag", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_tag", "prefixes.0.description", "my-description"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_status", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.by_status", "prefixes.0.description", "my-description"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.no_results", "prefixes.#", "0"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_tenant_id", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_tenant_id", "prefixes.0.prefix", "10.0.7.0/24"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_site_id", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_site_id", "prefixes.0.prefix", "10.0.8.0/24"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_contains", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.find_prefix_with_contains", "prefixes.0.prefix", "10.0.9.0/24"),
					resource.TestCheckResourceAttrSet("data.netbox_prefixes.find_prefix_with_contains", "prefixes.0.site_id"),
				),
			},
		},
	})
}
