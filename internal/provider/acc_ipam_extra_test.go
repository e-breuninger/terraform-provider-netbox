//go:build acctest

// ASN, VLAN group and VRF (the nested_id_list / int_ranges / default feature batch), plus the
// site.asns and contact.groups id lists.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxASN_basic(t *testing.T) {
	testName := testAccGetTestName("cat_asn")
	asn := acctest.RandIntRange(64512, 65500)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rir" "test" {
  name       = "%[1]s"
  slug       = "%[3]s"
  is_private = true
}
resource "netbox_ipam_role" "test" {
  name   = "%[1]s"
  slug   = "%[3]s"
  weight = 700
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "test" {
  asn         = %[2]d
  rir_id      = netbox_rir.test.id
  role_id     = netbox_ipam_role.test.id
  tenant_id   = netbox_tenant.test.id
  description = "%[1]s"
  comments    = "Created by acceptance test."
}
data "netbox_asn" "test" {
  depends_on = [netbox_asn.test]
  asn        = %[2]d
}
data "netbox_asns" "test" {
  depends_on = [netbox_asn.test]
  filters = [
    { name = "asn", value = "%[2]d" },
  ]
}`, testName, asn, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_asn.test", "asn", fmt.Sprintf("%d", asn)),
					resource.TestCheckResourceAttr("netbox_rir.test", "is_private", "true"),
					resource.TestCheckResourceAttr("netbox_ipam_role.test", "weight", "700"),
					resource.TestCheckResourceAttrPair("netbox_asn.test", "rir_id", "netbox_rir.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_asn.test", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_asn.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_asn.test", "id", "netbox_asn.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_asns.test", "asns.#", "1"),
				),
			},
			{
				// The tenant and role blocks stay: NetBox protects a tenant that is still
				// referenced, and the delete is not ordered after the update that unreferences it.
				Config: fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "test" {
  asn    = %[2]d
  rir_id = netbox_rir.test.id
}`, testName, asn),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_asn.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_asn.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_asn.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_asn.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_asn.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVLANGroup_basic(t *testing.T) {
	testName := testAccGetTestName("cat_vlangrp")
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// Ranges deliberately not in ascending order: a server-side sort must not plan as a change (the
				// implicit post-apply plan check fails otherwise).
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan_group" "test" {
  name          = "%[1]s"
  slug          = "%[2]s"
  description   = "%[1]s"
  comments      = "Created by acceptance test."
  tenant_id     = netbox_tenant.test.id
  site_group_id = netbox_site_group.test.id
  vid_ranges = [
    { start = 1200, end = 1299 },
    { start = 1100, end = 1149 },
  ]
}
resource "netbox_vlan_group" "single" {
  name = "%[1]s-single"
  vid_ranges = [
    { start = 1300, end = 1399 },
  ]
}
data "netbox_vlan_group" "test" {
  depends_on = [netbox_vlan_group.test]
  name       = "%[1]s"
}
data "netbox_vlan_groups" "test" {
  depends_on = [netbox_vlan_group.test]
  filters = [
    { name = "name", value = "%[1]s" },
  ]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "comments", "Created by acceptance test."),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "scope_type", "dcim.sitegroup"),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test", "site_group_id", "netbox_site_group.test", "id"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "vid_ranges.#", "2"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "vid_ranges.0.start", "1200"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "vid_ranges.1.end", "1149"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_group.test", "id", "netbox_vlan_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "1"),
				),
			},
			{
				// Single range: order-independent, safe to import-verify.
				ResourceName:      "netbox_vlan_group.single",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Cycle the scope through the remaining aliases and the explicit pair.
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan_group" "test" {
  name      = "%[1]s"
  region_id = netbox_region.test.id
  vid_ranges = [
    { start = 1200, end = 1299 },
    { start = 1100, end = 1149 },
  ]
}
resource "netbox_vlan_group" "single" {
  name = "%[1]s-single"
  vid_ranges = [
    { start = 1300, end = 1399 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "scope_type", "dcim.region"),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test", "region_id", "netbox_region.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan_group" "test" {
  name        = "%[1]s"
  location_id = netbox_location.test.id
  vid_ranges = [
    { start = 1200, end = 1299 },
    { start = 1100, end = 1149 },
  ]
}
resource "netbox_vlan_group" "single" {
  name = "%[1]s-single"
  vid_ranges = [
    { start = 1300, end = 1399 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "scope_type", "dcim.location"),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan_group" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
  vid_ranges = [
    { start = 1200, end = 1299 },
    { start = 1100, end = 1149 },
  ]
}
resource "netbox_vlan_group" "single" {
  name = "%[1]s-single"
  vid_ranges = [
    { start = 1300, end = 1399 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan_group" "test" {
  name       = "%[1]s"
  scope_type = "dcim.site"
  scope_id   = netbox_site.test.id
  vid_ranges = [
    { start = 1200, end = 1299 },
    { start = 1100, end = 1149 },
  ]
}
resource "netbox_vlan_group" "single" {
  name = "%[1]s-single"
  vid_ranges = [
    { start = 1300, end = 1399 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test", "scope_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vlan_group.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				// Shrink: description, comments, the tenant and the scope clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan_group" "test" {
  name = "%[1]s"
  vid_ranges = [
    { start = 1200, end = 1299 },
    { start = 1100, end = 1149 },
  ]
}
resource "netbox_vlan_group" "single" {
  name = "%[1]s-single"
  vid_ranges = [
    { start = 1300, end = 1399 },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_vlan_group.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_vlan_group.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_vlan_group.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_vlan_group.test", "scope_type"),
					resource.TestCheckNoResourceAttr("netbox_vlan_group.test", "scope_id"),
					resource.TestCheckNoResourceAttr("netbox_vlan_group.test", "site_id"),
					resource.TestCheckNoResourceAttr("netbox_vlan_group.test", "site_group_id"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test", "vid_ranges.#", "2"),
				),
			},
		},
	})
}

// TestAccNetboxVLAN_qinq: a customer VLAN carried in a service VLAN.
func TestAccNetboxVLAN_qinq(t *testing.T) {
	testName := testAccGetTestName("cat_qinq")
	svid := acctest.RandIntRange(2, 2000)
	deps := fmt.Sprintf(`
resource "netbox_vlan" "svlan" {
  name      = "%[1]s-svlan"
  vid       = %[2]d
  qinq_role = "svlan"
}`, testName, svid)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan" "cvlan" {
  name          = "%[1]s-cvlan"
  vid           = %[2]d
  qinq_role     = "cvlan"
  qinq_svlan_id = netbox_vlan.svlan.id
}`, testName, svid+2000),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.svlan", "qinq_role", "svlan"),
					resource.TestCheckResourceAttr("netbox_vlan.cvlan", "qinq_role", "cvlan"),
					resource.TestCheckResourceAttrPair("netbox_vlan.cvlan", "qinq_svlan_id", "netbox_vlan.svlan", "id"),
				),
			},
			{
				ResourceName:      "netbox_vlan.cvlan",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_vlan" "cvlan" {
  name = "%[1]s-cvlan"
  vid  = %[2]d
}`, testName, svid+2000),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_vlan.cvlan", "qinq_role"),
					resource.TestCheckNoResourceAttr("netbox_vlan.cvlan", "qinq_svlan_id"),
				),
			},
		},
	})
}

func TestAccNetboxVRF_basic(t *testing.T) {
	testName := testAccGetTestName("cat_vrf")
	routeTarget := acctest.RandIntRange(1, 60000)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// enforce_unique is not set: the schema default (true) must apply and match what NetBox
				// reports.
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_route_target" "a" {
  name = "65000:%[2]d"
}
resource "netbox_route_target" "b" {
  name = "65001:%[2]d"
}
resource "netbox_vrf" "test" {
  name           = "%[1]s"
  rd             = "65000:%[2]d"
  description    = "%[1]s"
  tenant_id      = netbox_tenant.test.id
  import_targets = [netbox_route_target.a.id]
  export_targets = [netbox_route_target.a.id, netbox_route_target.b.id]
}
data "netbox_vrf" "test" {
  depends_on = [netbox_vrf.test]
  name       = "%[1]s"
}
data "netbox_vrfs" "test" {
  depends_on = [netbox_vrf.test]
  filters = [
    { name = "tenant_id", value = netbox_tenant.test.id },
  ]
}`, testName, routeTarget),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test", "enforce_unique", "true"),
					resource.TestCheckResourceAttr("netbox_vrf.test", "import_targets.#", "1"),
					resource.TestCheckResourceAttr("netbox_vrf.test", "export_targets.#", "2"),
					resource.TestCheckResourceAttrPair("netbox_vrf.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vrf.test", "id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vrfs.test", "vrfs.#", "1"),
				),
			},
			{
				// Flip the default off and drop everything else.
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_route_target" "a" {
  name = "65000:%[2]d"
}
resource "netbox_route_target" "b" {
  name = "65001:%[2]d"
}
resource "netbox_vrf" "test" {
  name           = "%[1]s"
  enforce_unique = false
}`, testName, routeTarget),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test", "enforce_unique", "false"),
					resource.TestCheckNoResourceAttr("netbox_vrf.test", "import_targets.#"),
					resource.TestCheckNoResourceAttr("netbox_vrf.test", "export_targets.#"),
					resource.TestCheckNoResourceAttr("netbox_vrf.test", "rd"),
					resource.TestCheckNoResourceAttr("netbox_vrf.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_vrf.test", "tenant_id"),
				),
			},
			{
				ResourceName:      "netbox_vrf.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxSite_asns(t *testing.T) {
	testName := testAccGetTestName("site_asns")
	asn := acctest.RandIntRange(64512, 65400)
	config := func(asns string) string {
		return fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_asn" "a" {
  asn         = %[2]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
}
resource "netbox_asn" "b" {
  asn         = %[3]d
  rir_id      = netbox_rir.test.id
  description = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
  asn_ids = %[4]s
}`, testName, asn, asn+1, asns)
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config("[netbox_asn.a.id, netbox_asn.b.id]"),
				Check:  resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "2"),
			},
			{
				Config: config("[netbox_asn.b.id]"),
				Check:  resource.TestCheckResourceAttr("netbox_site.test", "asn_ids.#", "1"),
			},
		},
	})
}

func TestAccNetboxContact_groups(t *testing.T) {
	testName := testAccGetTestName("contact_groups")
	config := func(groups string) string {
		return fmt.Sprintf(`
resource "netbox_contact_group" "a" {
  name = "%[1]s-a"
}
resource "netbox_contact_group" "b" {
  name = "%[1]s-b"
}
resource "netbox_contact" "test" {
  name   = "%[1]s"
  group_ids = %[2]s
}`, testName, groups)
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config("[netbox_contact_group.a.id, netbox_contact_group.b.id]"),
				Check:  resource.TestCheckResourceAttr("netbox_contact.test", "group_ids.#", "2"),
			},
			{
				Config: config("[netbox_contact_group.a.id]"),
				Check:  resource.TestCheckResourceAttr("netbox_contact.test", "group_ids.#", "1"),
			},
		},
	})
}

func init() {
	sweep("netbox_asn",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamAsnsList(ipam.NewIpamAsnsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// ASNs have no name; tests put the test name in description.
				items = append(items, sweepItem{result.ID, result.Description})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Ipam.IpamAsnsDestroy(ipam.NewIpamAsnsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_vlan_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamVlanGroupsList(ipam.NewIpamVlanGroupsListParams(), nil)
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
			_, err := client.Ipam.IpamVlanGroupsDestroy(ipam.NewIpamVlanGroupsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_vrf",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamVrfsList(ipam.NewIpamVrfsListParams(), nil)
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
			_, err := client.Ipam.IpamVrfsDestroy(ipam.NewIpamVrfsDestroyParams().WithID(id), nil)
			return err
		})
}
