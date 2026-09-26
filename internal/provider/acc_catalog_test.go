//go:build acctest

// Catalog resources: flat NetBox objects with little to no required structure around them. The
// slugged name+description family shares one table-driven test; contact, vlan, prefix, cluster and
// route_target have their own.
package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// TestAccNetboxCatalog_slugged covers every catalog resource of the name+slug+description shape:
// create, derived slug, singular and plural data source, import.
func TestAccNetboxCatalog_slugged(t *testing.T) {
	cases := []struct{ resource, plural string }{
		{"manufacturer", "manufacturers"},
		{"platform", "platforms"},
		{"device_role", "device_roles"},
		{"rack_role", "rack_roles"},
		{"rir", "rirs"},
		{"ipam_role", "ipam_roles"},
		{"contact_role", "contact_roles"},
		{"contact_group", "contact_groups"},
		{"cluster_type", "cluster_types"},
		{"cluster_group", "cluster_groups"},
		{"circuit_provider", "circuit_providers"},
		{"circuit_type", "circuit_types"},
	}
	for _, testCase := range cases {
		t.Run(testCase.resource, func(t *testing.T) {
			testName := testAccGetTestName("cat_" + testCase.resource)
			resourceName := fmt.Sprintf("netbox_%s.test", testCase.resource)
			resource.ParallelTest(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				PreCheck:                 func() { testAccPreCheck(t) },
				Steps: []resource.TestStep{
					{
						Config: fmt.Sprintf(`
resource "netbox_%s" "test" {
  name        = "%s"
  description = "%s"
}
data "netbox_%s" "by_name" {
  depends_on = [netbox_%s.test]
  name       = "%s"
}
data "netbox_%s" "all" {
  depends_on = [netbox_%s.test]
  filters = [
    { name = "name__ic", value = "%s" },
  ]
}`, testCase.resource, testName, testName, testCase.resource, testCase.resource, testName, testCase.plural, testCase.resource, testName),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(resourceName, "name", testName),
							resource.TestCheckResourceAttr(resourceName, "slug", getSlug(testName)),
							resource.TestCheckResourceAttr(resourceName, "description", testName),
							resource.TestCheckResourceAttrPair("data.netbox_"+testCase.resource+".by_name", "id", resourceName, "id"),
							resource.TestCheckResourceAttr("data.netbox_"+testCase.plural+".all", testCase.plural+".#", "1"),
						),
					},
					{
						ResourceName:      resourceName,
						ImportState:       true,
						ImportStateVerify: true,
					},
					{
						// Shrink: description clears. go-netbox's omitempty would drop the empty
						// string, so this is what proves the explicit clear in the PUT body works.
						Config: fmt.Sprintf(`
resource "netbox_%s" "test" {
  name = "%s"
}`, testCase.resource, testName),
						Check: resource.ComposeTestCheckFunc(
							resource.TestCheckResourceAttr(resourceName, "name", testName),
							resource.TestCheckNoResourceAttr(resourceName, "description"),
							resource.TestCheckResourceAttr(resourceName, "slug", getSlug(testName)),
						),
					},
				},
			})
		})
	}
}

func TestAccNetboxContact_basic(t *testing.T) {
	testName := testAccGetTestName("cat_contact")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_contact_group" "parent" {
  name = "%[1]s-cgp"
}
resource "netbox_contact_group" "cg" {
  name      = "%[1]s-cg"
  slug      = "%[2]s-cg"
  parent_id = netbox_contact_group.parent.id
}
resource "netbox_contact_role" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}
resource "netbox_contact" "test" {
  name        = "%[1]s"
  title       = "Chief Gurk Officer"
  phone       = "+49 123 456"
  email       = "%[1]s@example.com"
  address     = "Example Street 1"
  link        = "https://example.com/%[1]s"
  description = "%[1]s"
  group_ids   = [netbox_contact_group.cg.id]
}
data "netbox_contact" "test" {
  depends_on = [netbox_contact.test]
  name       = "%[1]s"
}
data "netbox_contacts" "test" {
  depends_on = [netbox_contact.test]
  filters = [
    { name = "name", value = "%[1]s" },
  ]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_contact.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_contact.test", "group_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_contact.test", "email", testName+"@example.com"),
					resource.TestCheckResourceAttr("netbox_contact.test", "link", "https://example.com/"+testName),
					resource.TestCheckResourceAttr("netbox_contact.test", "description", testName),
					resource.TestCheckResourceAttrPair("netbox_contact_group.cg", "parent_id", "netbox_contact_group.parent", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_contact.test", "id", "netbox_contact.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_contacts.test", "contacts.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_contact.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: every optional scalar clears, the group memberships and the group's
				// parent too. The group blocks stay so the delete ordering cannot bite.
				Config: fmt.Sprintf(`
resource "netbox_contact_group" "parent" {
  name = "%[1]s-cgp"
}
resource "netbox_contact_group" "cg" {
  name = "%[1]s-cg"
}
resource "netbox_contact" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_contact.test", "title"),
					resource.TestCheckNoResourceAttr("netbox_contact.test", "phone"),
					resource.TestCheckNoResourceAttr("netbox_contact.test", "email"),
					resource.TestCheckNoResourceAttr("netbox_contact.test", "address"),
					resource.TestCheckNoResourceAttr("netbox_contact.test", "link"),
					resource.TestCheckNoResourceAttr("netbox_contact.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_contact.test", "group_ids"),
					resource.TestCheckNoResourceAttr("netbox_contact_group.cg", "parent_id"),
				),
			},
		},
	})
}

func TestAccNetboxVLAN_basic(t *testing.T) {
	testName := testAccGetTestName("cat_vlan")
	vid := acctest.RandIntRange(2, 4000)
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_vlan_group" "test" {
  name = "%[1]s"
  vid_ranges = [
    { start = %[2]d, end = %[2]d },
  ]
}`, testName, vid)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan" "test" {
  name        = "%[1]s"
  vid         = %[2]d
  status      = "active"
  description = "%[1]s"
  comments    = "Created by acceptance test."
  tenant_id   = netbox_tenant.test.id
  site_id     = netbox_site.test.id
  group_id    = netbox_vlan_group.test.id
  role_id     = netbox_ipam_role.test.id
}
data "netbox_vlan" "test" {
  depends_on = [netbox_vlan.test]
  vid        = %[2]d
}
data "netbox_vlans" "test" {
  depends_on = [netbox_vlan.test]
  filters = [
    { name = "group_id", value = netbox_vlan_group.test.id },
  ]
}`, testName, vid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan.test", "vid", fmt.Sprintf("%d", vid)),
					resource.TestCheckResourceAttr("netbox_vlan.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_vlan.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_vlan.test", "comments", "Created by acceptance test."),
					resource.TestCheckResourceAttrPair("netbox_vlan.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test", "group_id", "netbox_vlan_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan.test", "id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_vlan.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: every reference and the description clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan" "test" {
  name = "%[1]s"
  vid  = %[2]d
}`, testName, vid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test", "status", "active"),
					resource.TestCheckNoResourceAttr("netbox_vlan.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_vlan.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_vlan.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_vlan.test", "site_id"),
					resource.TestCheckNoResourceAttr("netbox_vlan.test", "group_id"),
					resource.TestCheckNoResourceAttr("netbox_vlan.test", "role_id"),
				),
			},
		},
	})
}

func TestAccNetboxPrefix_basic(t *testing.T) {
	testName := testAccGetTestName("cat_prefix")
	prefix := fmt.Sprintf("10.%d.%d.0/24", acctest.RandIntRange(2, 250), acctest.RandIntRange(0, 250))
	vid := acctest.RandIntRange(2, 4000)
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_vlan" "test" {
  name = "%[1]s"
  vid  = %[2]d
}`, testName, vid)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix        = "%[2]s"
  status        = "active"
  is_pool       = true
  mark_utilized = true
  description   = "%[1]s"
  tenant_id     = netbox_tenant.test.id
  vrf_id        = netbox_vrf.test.id
  role_id       = netbox_ipam_role.test.id
  vlan_id       = netbox_vlan.test.id
  site_id       = netbox_site.test.id
}
data "netbox_prefix" "test" {
  depends_on = [netbox_prefix.test]
  prefix     = "%[2]s"
  vrf_id     = netbox_vrf.test.id
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", prefix),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "is_pool", "true"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "mark_utilized", "true"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "vlan_id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttr("netbox_prefix.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "scope_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "region_id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_prefix.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Cycle the scope through the remaining aliases and the explicit pair.
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix    = "%[2]s"
  region_id = netbox_region.test.id
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "scope_type", "dcim.region"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "region_id", "netbox_region.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix        = "%[2]s"
  site_group_id = netbox_site_group.test.id
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "scope_type", "dcim.sitegroup"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "site_group_id", "netbox_site_group.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix      = "%[2]s"
  location_id = netbox_location.test.id
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "scope_type", "dcim.location"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix     = "%[2]s"
  scope_type = "dcim.site"
  scope_id   = netbox_site.test.id
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "scope_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_prefix.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				// Shrink: description and every reference clear; status, is_pool and
				// mark_utilized are optional+computed and keep NetBox's value.
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[2]s"
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_prefix.test", "mark_utilized", "true"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "scope_type"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "scope_id"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "site_id"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "vrf_id"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_prefix.test", "vlan_id"),
				),
			},
		},
	})
}

func TestAccNetboxPrefixes_filters(t *testing.T) {
	testName := testAccGetTestName("cat_prefixes")
	prefix := fmt.Sprintf("10.%d.%d.0/24", acctest.RandIntRange(2, 250), acctest.RandIntRange(0, 250))
	vid := acctest.RandIntRange(2, 4000)
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_vlan" "test" {
  name = "%[1]s"
  vid  = %[2]d
}`, testName, vid)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// Every filter the pre-6.0 provider's netbox_prefixes block accepts, ANDed down
				// to the one matching prefix; the tenant scopes the query to this test's objects.
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix      = "%[2]s"
  status      = "active"
  description = "%[1]s"
  tenant_id   = netbox_tenant.test.id
  role_id     = netbox_ipam_role.test.id
  vlan_id     = netbox_vlan.test.id
  site_id     = netbox_site.test.id
}
data "netbox_prefixes" "test" {
  depends_on = [netbox_prefix.test]
  filters = [
    { name = "tenant_id", value = netbox_tenant.test.id },
    { name = "status", value = "active" },
    { name = "role_id", value = netbox_ipam_role.test.id },
    { name = "vlan_id", value = netbox_vlan.test.id },
    { name = "vlan_vid", value = %[3]d },
    { name = "site_id", value = netbox_site.test.id },
    { name = "description", value = "%[1]s" },
  ]
}`, testName, prefix, vid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_prefixes.test", "prefixes.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_prefixes.test", "prefixes.0.prefix", prefix),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.test", "prefixes.0.id", "netbox_prefix.test", "id"),
				),
			},
			{
				// region_id matches through NetBox's cached scope tree, here a prefix scoped to
				// the region itself.
				Config: deps + fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix    = "%[2]s"
  tenant_id = netbox_tenant.test.id
  region_id = netbox_region.test.id
}
data "netbox_prefixes" "test" {
  depends_on = [netbox_prefix.test]
  filters = [
    { name = "tenant_id", value = netbox_tenant.test.id },
    { name = "region_id", value = netbox_region.test.id },
  ]
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_prefixes.test", "prefixes.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_prefixes.test", "prefixes.0.region_id", "netbox_region.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxCluster_basic(t *testing.T) {
	testName := testAccGetTestName("cat_cluster")
	deps := fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}
resource "netbox_cluster_group" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName, getSlug(testName))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_cluster" "test" {
  name             = "%[1]s"
  cluster_type_id  = netbox_cluster_type.test.id
  cluster_group_id = netbox_cluster_group.test.id
  tenant_id        = netbox_tenant.test.id
  status           = "active"
  description      = "%[1]s"
  comments         = "Created by acceptance test."
  scope_type       = "dcim.region"
  scope_id         = netbox_region.test.id
}
data "netbox_cluster" "test" {
  depends_on = [netbox_cluster.test]
  name       = "%[1]s"
}
data "netbox_clusters" "test" {
  depends_on = [netbox_cluster.test]
  filters = [
    { name = "group_id", value = netbox_cluster_group.test.id },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_cluster.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "cluster_type_id", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "scope_id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "region_id", "netbox_region.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "site_id"),
					resource.TestCheckResourceAttrPair("data.netbox_cluster.test", "id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_clusters.test", "clusters.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_cluster.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Cycle the scope through each alias.
				Config: deps + fmt.Sprintf(`
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id         = netbox_site.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_group_id   = netbox_site_group.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "scope_type", "dcim.sitegroup"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "site_group_id", "netbox_site_group.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  location_id     = netbox_location.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "scope_type", "dcim.location"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  region_id       = netbox_region.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "scope_type", "dcim.region"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "region_id", "netbox_region.test", "id"),
				),
			},
			{
				// Shrink: description and every reference clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "scope_type"),
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "scope_id"),
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "region_id"),
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "cluster_group_id"),
					resource.TestCheckNoResourceAttr("netbox_cluster.test", "tenant_id"),
				),
			},
		},
	})
}

func TestAccNetboxRouteTarget_basic(t *testing.T) {
	testName := testAccGetTestName("cat_rt")
	rtNumber := acctest.RandIntRange(1, 65000)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_route_target" "test" {
  name        = "65000:%[2]d"
  description = "%[1]s"
  tenant_id   = netbox_tenant.test.id
}
data "netbox_route_target" "test" {
  depends_on = [netbox_route_target.test]
  name       = "65000:%[2]d"
}
data "netbox_route_targets" "test" {
  depends_on = [netbox_route_target.test]
  filters = [
    { name = "name", value = "65000:%[2]d" },
  ]
}`, testName, rtNumber),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_route_target.test", "description", testName),
					resource.TestCheckResourceAttrPair("netbox_route_target.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_route_target.test", "id", "netbox_route_target.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_route_targets.test", "route_targets.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_route_target.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: description and the tenant reference clear.
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_route_target" "test" {
  name = "65000:%[2]d"
}`, testName, rtNumber),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_route_target.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_route_target.test", "tenant_id"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_manufacturer",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimManufacturersList(dcim.NewDcimManufacturersListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimManufacturersDestroy(dcim.NewDcimManufacturersDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_platform",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimPlatformsList(dcim.NewDcimPlatformsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimPlatformsDestroy(dcim.NewDcimPlatformsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_device_role",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimDeviceRolesList(dcim.NewDcimDeviceRolesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimDeviceRolesDestroy(dcim.NewDcimDeviceRolesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_rack_role",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimRackRolesList(dcim.NewDcimRackRolesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimRackRolesDestroy(dcim.NewDcimRackRolesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_rir",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamRirsList(ipam.NewIpamRirsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Ipam.IpamRirsDestroy(ipam.NewIpamRirsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_ipam_role",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamRolesList(ipam.NewIpamRolesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Ipam.IpamRolesDestroy(ipam.NewIpamRolesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_vlan",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamVlansList(ipam.NewIpamVlansListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Ipam.IpamVlansDestroy(ipam.NewIpamVlansDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_contact_role",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Tenancy.TenancyContactRolesList(tenancy.NewTenancyContactRolesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Tenancy.TenancyContactRolesDestroy(tenancy.NewTenancyContactRolesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_contact_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Tenancy.TenancyContactGroupsList(tenancy.NewTenancyContactGroupsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Tenancy.TenancyContactGroupsDestroy(tenancy.NewTenancyContactGroupsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_contact",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Tenancy.TenancyContactsList(tenancy.NewTenancyContactsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Tenancy.TenancyContactsDestroy(tenancy.NewTenancyContactsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_cluster",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Virtualization.VirtualizationClustersList(virtualization.NewVirtualizationClustersListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Virtualization.VirtualizationClustersDestroy(virtualization.NewVirtualizationClustersDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_cluster_type",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Virtualization.VirtualizationClusterTypesList(virtualization.NewVirtualizationClusterTypesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Virtualization.VirtualizationClusterTypesDestroy(virtualization.NewVirtualizationClusterTypesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_cluster_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Virtualization.VirtualizationClusterGroupsList(virtualization.NewVirtualizationClusterGroupsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Virtualization.VirtualizationClusterGroupsDestroy(virtualization.NewVirtualizationClusterGroupsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_circuit_provider",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsProvidersList(circuits.NewCircuitsProvidersListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Circuits.CircuitsProvidersDestroy(circuits.NewCircuitsProvidersDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_circuit_type",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsCircuitTypesList(circuits.NewCircuitsCircuitTypesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Circuits.CircuitsCircuitTypesDestroy(circuits.NewCircuitsCircuitTypesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_prefix",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamPrefixesList(ipam.NewIpamPrefixesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, result.Description})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Ipam.IpamPrefixesDestroy(ipam.NewIpamPrefixesDestroyParams().WithID(id), nil)
			return err
		})
}

// prefix is normalize cidr: an IPv6 prefix written uppercase and uncompressed keeps that spelling
// in state although NetBox echoes it lowercase and compressed, a prefix with host bits (which
// NetBox rejects) fails at plan, and the next plan is empty.
func TestAccNetboxPrefix_normalizedSpelling(t *testing.T) {
	spelled := fmt.Sprintf("2001:DB8:0:%X:0:0:0:0/64", acctest.RandIntRange(0x100, 0xfff))
	config := fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%s"
  status = "active"
}`, spelled)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config,
				Check:  resource.TestCheckResourceAttr("netbox_prefix.test", "prefix", spelled),
			},
			{
				// Not the last step: the post-test destroy runs the last step's config.
				Config: `
resource "netbox_prefix" "test" {
  prefix = "10.0.0.5/24"
  status = "active"
}`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("(?i)host bits set"),
			},
			{
				Config: config,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
		},
	})
}
