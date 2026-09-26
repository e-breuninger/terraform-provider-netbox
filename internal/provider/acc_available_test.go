//go:build acctest

// The allocation resources: available_ip_address, available_prefix, available_vlan and
// available_asn.
package provider_test

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNetboxAvailableIPAddress_prefix(t *testing.T) {
	testName := testAccGetTestName("avail_ip")
	net := fmt.Sprintf("10.%d.0", acctest.RandIntRange(100, 199))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_prefix" "test" {
  prefix = "%[2]s.0/24"
  status = "active"
}
resource "netbox_ip_address" "peer" {
  ip_address = "%[2]s.250/24"
  status     = "active"
}
resource "netbox_available_ip_address" "test" {
  prefix_id     = netbox_prefix.test.id
  status        = "reserved"
  role          = "vip"
  dns_name      = "%[3]s.example.com"
  description   = "%[1]s"
  comments      = "Created by acceptance test."
  tenant_id     = netbox_tenant.test.id
  vrf_id        = netbox_vrf.test.id
  nat_inside_id = netbox_ip_address.peer.id
  tags          = [netbox_tag.test.slug]
}`, testName, net, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("netbox_available_ip_address.test", "ip_address", regexp.MustCompile(`^`+regexp.QuoteMeta(net)+`\.\d+/24$`)),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "role", "vip"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "dns_name", getSlug(testName)+".example.com"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "family", "4"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "nat_inside_id", "netbox_ip_address.peer", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "prefix_id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "tags.#", "1"),
				),
			},
			{
				// Update through the ordinary ip-addresses endpoint; the address stays in the prefix.
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_prefix" "test" {
  prefix = "%[2]s.0/24"
  status = "active"
}
resource "netbox_ip_address" "peer" {
  ip_address = "%[2]s.250/24"
  status     = "active"
}
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
  # The address keeps the VRF set in the previous step (vrf_id is optional+computed), so it
  # still has to go before the VRF on destroy.
  depends_on = [netbox_vrf.test]
}`, testName, net),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("netbox_available_ip_address.test", "ip_address", regexp.MustCompile(`^`+regexp.QuoteMeta(net)+`\.\d+/24$`)),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "status", "active"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "role"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "dns_name"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "tenant_id"),
					// vrf_id is optional+computed (unset means the allocation source's VRF), so
					// unsetting it keeps NetBox's value like status and is_pool do.
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "nat_inside_id"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "tags"),
				),
			},
			{
				// prefix_id is an allocation input NetBox does not report back; the adoption tests
				// below cover what an import without it means.
				ResourceName:            "netbox_available_ip_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"prefix_id"},
			},
		},
	})
}

// An address allocated without vrf_id inherits the VRF of its prefix or IP range instead of being
// moved into the global table by the configure call.
func TestAccNetboxAvailableIPAddress_inheritsVRF(t *testing.T) {
	testName := testAccGetTestName("avail_ip_vrf")
	net := fmt.Sprintf("10.%d.0", acctest.RandIntRange(100, 199))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_prefix" "test" {
  prefix = "%[2]s.0/25"
  status = "active"
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_range" "test" {
  start_address = "%[2]s.129/24"
  end_address   = "%[2]s.200/24"
  vrf_id        = netbox_vrf.test.id
}
resource "netbox_available_ip_address" "from_prefix" {
  prefix_id = netbox_prefix.test.id
}
resource "netbox_available_ip_address" "from_range" {
  ip_range_id = netbox_ip_range.test.id
}`, testName, net),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.from_prefix", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.from_range", "vrf_id", "netbox_vrf.test", "id"),
				),
			},
		},
	})
}

// A prefix allocated without vrf_id inherits the VRF of its parent.
func TestAccNetboxAvailablePrefix_inheritsVRF(t *testing.T) {
	testName := testAccGetTestName("avail_prefix_vrf")
	net := fmt.Sprintf("10.%d", acctest.RandIntRange(100, 199))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_prefix" "parent" {
  prefix = "%[2]s.0.0/16"
  status = "container"
  vrf_id = netbox_vrf.test.id
}
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
}`, testName, net),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("netbox_available_prefix.test", "prefix", regexp.MustCompile(`^`+regexp.QuoteMeta(net)+`\.\d+\.0/24$`)),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "vrf_id", "netbox_vrf.test", "id"),
				),
			},
		},
	})
}

// testAccCheckImportedAllocationSourceUnset asserts that an import leaves the allocation input
// (prefix_id or ip_range_id) out of state: NetBox does not report it.
func testAccCheckImportedAllocationSourceUnset(key string) resource.ImportStateCheckFunc {
	return func(states []*terraform.InstanceState) error {
		// The persisted state also holds the prefixes/ranges of earlier steps.
		for _, state := range states {
			if state.Ephemeral.Type != "netbox_available_ip_address" {
				continue
			}
			if v, ok := state.Attributes[key]; ok {
				return fmt.Errorf("expected %s to be unset after import, got %q", key, v)
			}
			return nil
		}
		return fmt.Errorf("no netbox_available_ip_address in the imported state")
	}
}

// TestAccNetboxAvailableIPAddress_adoptExistingIP imports an address allocated out of band and
// asserts it survives: same NetBox record, prefix recorded in state, empty plan afterwards.
// Moving it to another prefix from there must still re-allocate.
func TestAccNetboxAvailableIPAddress_adoptExistingIP(t *testing.T) {
	net := fmt.Sprintf("10.%d", acctest.RandIntRange(100, 199))
	prefixes := fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[1]s.20.0/24"
  status = "active"
}
resource "netbox_prefix" "other" {
  prefix = "%[1]s.21.0/24"
  status = "active"
}`, net)
	withAdoptedIP := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
}`
	movedToOtherPrefix := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.other.id
}`

	var prefixID, adoptedIPID int64
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: prefixes,
				Check: func(state *terraform.State) (err error) {
					prefixID, err = testAccStateID(state, "netbox_prefix.test")
					return err
				},
			},
			{
				// Allocate behind Terraform's back, the same way the resource does, then adopt it.
				PreConfig: func() {
					api, err := testAccClient()
					if err != nil {
						t.Fatal(err)
					}
					res, err := api.Ipam.IpamPrefixesAvailableIpsCreate(
						ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(prefixID).WithData([]*models.AvailableIPRequest{{}}), nil)
					if err != nil {
						t.Fatalf("allocating an out-of-band IP in prefix %d: %s", prefixID, err)
					}
					if len(res.Payload) == 0 {
						t.Fatalf("no available IP addresses in prefix %d", prefixID)
					}
					adoptedIPID = res.Payload[0].ID
				},
				Config:             withAdoptedIP,
				ResourceName:       "netbox_available_ip_address.test",
				ImportState:        true,
				ImportStatePersist: true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strconv.FormatInt(adoptedIPID, 10), nil
				},
				ImportStateCheck: testAccCheckImportedAllocationSourceUnset("prefix_id"),
			},
			{
				// Records the prefix in state and keeps the address. Without the write-only
				// RequiresReplaceIf this step destroys the imported address and allocates another.
				Config: withAdoptedIP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "prefix_id", "netbox_prefix.test", "id"),
					testAccCheckStateIDIs("netbox_available_ip_address.test", &adoptedIPID, true),
				),
			},
			{
				Config:   withAdoptedIP,
				PlanOnly: true,
			},
			{
				// The source is known now, so a change re-allocates as usual.
				Config: movedToOtherPrefix,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", net+".21.1/24"),
					testAccCheckStateIDIs("netbox_available_ip_address.test", &adoptedIPID, false),
				),
			},
		},
	})
}

// TestAccNetboxAvailableIPAddress_adoptExistingIPFromRange is the ip_range_id twin of
// TestAccNetboxAvailableIPAddress_adoptExistingIP.
func TestAccNetboxAvailableIPAddress_adoptExistingIPFromRange(t *testing.T) {
	net := fmt.Sprintf("10.%d.22", acctest.RandIntRange(100, 199))
	rangeOnly := fmt.Sprintf(`
resource "netbox_ip_range" "test" {
  start_address = "%[1]s.1/24"
  end_address   = "%[1]s.50/24"
}`, net)
	withAdoptedIP := rangeOnly + `
resource "netbox_available_ip_address" "test" {
  ip_range_id = netbox_ip_range.test.id
}`

	var rangeID, adoptedIPID int64
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: rangeOnly,
				Check: func(state *terraform.State) (err error) {
					rangeID, err = testAccStateID(state, "netbox_ip_range.test")
					return err
				},
			},
			{
				PreConfig: func() {
					api, err := testAccClient()
					if err != nil {
						t.Fatal(err)
					}
					res, err := api.Ipam.IpamIPRangesAvailableIpsCreate(
						ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(rangeID).WithData([]*models.AvailableIPRequest{{}}), nil)
					if err != nil {
						t.Fatalf("allocating an out-of-band IP in range %d: %s", rangeID, err)
					}
					if len(res.Payload) == 0 {
						t.Fatalf("no available IP addresses in range %d", rangeID)
					}
					adoptedIPID = res.Payload[0].ID
				},
				Config:             withAdoptedIP,
				ResourceName:       "netbox_available_ip_address.test",
				ImportState:        true,
				ImportStatePersist: true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strconv.FormatInt(adoptedIPID, 10), nil
				},
				ImportStateCheck: testAccCheckImportedAllocationSourceUnset("ip_range_id"),
			},
			{
				Config: withAdoptedIP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "ip_range_id", "netbox_ip_range.test", "id"),
					testAccCheckStateIDIs("netbox_available_ip_address.test", &adoptedIPID, true),
				),
			},
			{
				Config:   withAdoptedIP,
				PlanOnly: true,
			},
		},
	})
}

// TestAccNetboxAvailableIPAddress_allocationSourceChangeReallocates guards the other side: a real
// prefix change on a provider-created address still re-allocates.
func TestAccNetboxAvailableIPAddress_allocationSourceChangeReallocates(t *testing.T) {
	net := fmt.Sprintf("10.%d", acctest.RandIntRange(100, 199))
	prefixes := fmt.Sprintf(`
resource "netbox_prefix" "first" {
  prefix = "%[1]s.23.0/24"
  status = "active"
}
resource "netbox_prefix" "second" {
  prefix = "%[1]s.24.0/24"
  status = "active"
}`, net)
	fromFirst := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.first.id
}`
	fromSecond := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.second.id
}`

	var firstIPID int64
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fromFirst,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", net+".23.1/24"),
					func(state *terraform.State) (err error) {
						firstIPID, err = testAccStateID(state, "netbox_available_ip_address.test")
						return err
					},
				),
			},
			{
				Config: fromSecond,
				Check: resource.ComposeTestCheckFunc(
					// New address and new record; an in-place update would have kept both.
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", net+".24.1/24"),
					testAccCheckStateIDIs("netbox_available_ip_address.test", &firstIPID, false),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_ipRange(t *testing.T) {
	testName := testAccGetTestName("avail_iprange")
	net := fmt.Sprintf("10.%d.1", acctest.RandIntRange(100, 199))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_ip_range" "test" {
  start_address = "%[2]s.10/24"
  end_address   = "%[2]s.20/24"
  vrf_id        = netbox_vrf.test.id
}
resource "netbox_available_ip_address" "test" {
  ip_range_id = netbox_ip_range.test.id
  vrf_id      = netbox_vrf.test.id
  description = "%[1]s"
}`, testName, net),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("netbox_available_ip_address.test", "ip_address", regexp.MustCompile(`^`+regexp.QuoteMeta(net)+`\.1\d/24$`)),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "ip_range_id", "netbox_ip_range.test", "id"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "status", "active"),
				),
			},
		},
	})
}

// TestAccNetboxAvailableIPAddress_assignment mirrors TestAccNetboxIPAddress_assignment for the
// allocated address: the interface aliases and the generic assigned_object pair work on the
// available-ips path too.
func TestAccNetboxAvailableIPAddress_assignment(t *testing.T) {
	testName := testAccGetTestName("avail_ip_assign")
	net := fmt.Sprintf("10.%d.0", acctest.RandIntRange(200, 224))
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "eth0"
  type      = "1000base-t"
}
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.test.id
}
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%[1]s"
}
resource "netbox_prefix" "test" {
  prefix = "%[2]s.0/24"
  status = "active"
}`, testName, net)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_available_ip_address" "test" {
  prefix_id                    = netbox_prefix.test.id
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "assigned_object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "assigned_object_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
				),
			},
			{
				// Move to the device interface via the alias.
				Config: deps + `
resource "netbox_available_ip_address" "test" {
  prefix_id           = netbox_prefix.test.id
  device_interface_id = netbox_device_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "assigned_object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "device_interface_id", "netbox_device_interface.test", "id"),
				),
			},
			{
				// The same assignment written as the explicit pair is a no-op.
				Config: deps + `
resource "netbox_available_ip_address" "test" {
  prefix_id            = netbox_prefix.test.id
  assigned_object_type = "dcim.interface"
  assigned_object_id   = netbox_device_interface.test.id
}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				Check: resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "assigned_object_id", "netbox_device_interface.test", "id"),
			},
			{
				// Unassign: the assignment clears in place.
				Config: deps + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "assigned_object_type"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "assigned_object_id"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "device_interface_id"),
					resource.TestCheckNoResourceAttr("netbox_available_ip_address.test", "virtual_machine_interface_id"),
				),
			},
		},
	})
}

func TestAccNetboxAvailablePrefix_basic(t *testing.T) {
	testName := testAccGetTestName("avail_prefix")
	net := fmt.Sprintf("10.%d", acctest.RandIntRange(225, 250))
	vid := acctest.RandIntRange(2, 4000)
	// Every dependency stays in every step: dropping a resource and the reference to it in one
	// apply does not guarantee the reference-clearing update runs first.
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
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
  vid  = %[3]d
}
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_prefix" "parent" {
  prefix = "%[2]s.0.0/16"
  status = "container"
}`, testName, net, vid)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
  status           = "reserved"
  description      = "%[1]s"
  tenant_id        = netbox_tenant.test.id
  is_pool          = true
  mark_utilized    = true
  role_id          = netbox_ipam_role.test.id
  vlan_id          = netbox_vlan.test.id
  vrf_id           = netbox_vrf.test.id
  site_id          = netbox_site.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("netbox_available_prefix.test", "prefix", regexp.MustCompile(`^`+regexp.QuoteMeta(net)+`\.\d+\.0/24$`)),
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "is_pool", "true"),
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "mark_utilized", "true"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "parent_prefix_id", "netbox_prefix.parent", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "vlan_id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				// Cycle the scope through each alias and the explicit pair; each write is one kind.
				Config: deps + `
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
  region_id        = netbox_region.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "scope_type", "dcim.region"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "scope_id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "region_id", "netbox_region.test", "id"),
				),
			},
			{
				Config: deps + `
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
  site_group_id    = netbox_site_group.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "scope_type", "dcim.sitegroup"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "site_group_id", "netbox_site_group.test", "id"),
				),
			},
			{
				Config: deps + `
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
  location_id      = netbox_location.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "scope_type", "dcim.location"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				Config: deps + `
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
  scope_type       = "dcim.site"
  scope_id         = netbox_site.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "scope_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				Config: deps + `
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
  # The prefix keeps the VRF set earlier (vrf_id is optional+computed), so it still has to go
  # before the VRF on destroy.
  depends_on = [netbox_vrf.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("netbox_available_prefix.test", "prefix", regexp.MustCompile(`^`+regexp.QuoteMeta(net)+`\.\d+\.0/24$`)),
					// status and is_pool are optional+computed: they keep NetBox's value.
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "is_pool", "true"),
					resource.TestCheckNoResourceAttr("netbox_available_prefix.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_available_prefix.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_available_prefix.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_available_prefix.test", "vlan_id"),
					// vrf_id is optional+computed too (unset means the parent's VRF).
					resource.TestCheckResourceAttrPair("netbox_available_prefix.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_available_prefix.test", "scope_type"),
					resource.TestCheckNoResourceAttr("netbox_available_prefix.test", "scope_id"),
					resource.TestCheckNoResourceAttr("netbox_available_prefix.test", "site_id"),
				),
			},
			{
				ResourceName:      "netbox_available_prefix.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateIdFunc: func(state *terraform.State) (string, error) {
					resource := state.RootModule().Resources["netbox_available_prefix.test"]
					if resource == nil {
						return "", fmt.Errorf("netbox_available_prefix.test not in state")
					}
					a := resource.Primary.Attributes
					return a["parent_prefix_id"] + "/" + resource.Primary.ID + "/" + a["prefix_length"], nil
				},
			},
		},
	})
}

func TestAccNetboxAvailableVLAN_basic(t *testing.T) {
	testName := testAccGetTestName("avail_vlan")
	start := acctest.RandIntRange(2000, 3900)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_vlan" "svlan" {
  name      = "%[1]s-svlan"
  vid       = %[4]d
  qinq_role = "svlan"
}
resource "netbox_vlan_group" "test" {
  name = "%[1]s"
  vid_ranges = [
    { start = %[2]d, end = %[3]d },
  ]
}
resource "netbox_available_vlan" "first" {
  name          = "%[1]s-1"
  group_id      = netbox_vlan_group.test.id
  status        = "reserved"
  description   = "%[1]s"
  comments      = "Created by acceptance test."
  tenant_id     = netbox_tenant.test.id
  site_id       = netbox_site.test.id
  role_id       = netbox_ipam_role.test.id
  qinq_role     = "cvlan"
  qinq_svlan_id = netbox_vlan.svlan.id
}
resource "netbox_available_vlan" "second" {
  depends_on = [netbox_available_vlan.first]
  name       = "%[1]s-2"
  group_id   = netbox_vlan_group.test.id
}`, testName, start, start+9, start-1000),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_vlan.first", "vid", fmt.Sprint(start)),
					resource.TestCheckResourceAttr("netbox_available_vlan.first", "status", "reserved"),
					resource.TestCheckResourceAttrPair("netbox_available_vlan.first", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_vlan.first", "group_id", "netbox_vlan_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_vlan.first", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_vlan.first", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttr("netbox_available_vlan.first", "qinq_role", "cvlan"),
					resource.TestCheckResourceAttrPair("netbox_available_vlan.first", "qinq_svlan_id", "netbox_vlan.svlan", "id"),
					resource.TestCheckResourceAttr("netbox_available_vlan.first", "comments", "Created by acceptance test."),
					resource.TestCheckResourceAttr("netbox_available_vlan.second", "vid", fmt.Sprint(start+1)),
					resource.TestCheckResourceAttr("netbox_available_vlan.second", "status", "active"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_vlan" "svlan" {
  name      = "%[1]s-svlan"
  vid       = %[4]d
  qinq_role = "svlan"
}
resource "netbox_vlan_group" "test" {
  name = "%[1]s"
  vid_ranges = [
    { start = %[2]d, end = %[3]d },
  ]
}
resource "netbox_available_vlan" "first" {
  name     = "%[1]s-1"
  group_id = netbox_vlan_group.test.id
}
resource "netbox_available_vlan" "second" {
  depends_on = [netbox_available_vlan.first]
  name       = "%[1]s-2"
  group_id   = netbox_vlan_group.test.id
}`, testName, start, start+9, start-1000),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_vlan.first", "vid", fmt.Sprint(start)),
					resource.TestCheckResourceAttr("netbox_available_vlan.first", "status", "reserved"),
					resource.TestCheckNoResourceAttr("netbox_available_vlan.first", "description"),
					resource.TestCheckNoResourceAttr("netbox_available_vlan.first", "comments"),
					resource.TestCheckNoResourceAttr("netbox_available_vlan.first", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_available_vlan.first", "site_id"),
					resource.TestCheckNoResourceAttr("netbox_available_vlan.first", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_available_vlan.first", "qinq_role"),
					resource.TestCheckNoResourceAttr("netbox_available_vlan.first", "qinq_svlan_id"),
				),
			},
			{
				ResourceName:      "netbox_available_vlan.first",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxAvailableASN_basic(t *testing.T) {
	testName := testAccGetTestName("avail_asn")
	// A private-use range (RFC 6996), picked at random so parallel runs do not collide.
	start := acctest.RandIntRange(64600, 65400)
	deps := fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_asn_range" "test" {
  name   = "%[1]s"
  slug   = "%[4]s"
  rir_id = netbox_rir.test.id
  start  = %[2]d
  end    = %[3]d
}`, testName, start, start+9, getSlug(testName))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_available_asn" "first" {
  asn_range_id = netbox_asn_range.test.id
  tenant_id    = netbox_tenant.test.id
  role_id      = netbox_ipam_role.test.id
  description  = "%[1]s"
  comments     = "Created by acceptance test."
}
resource "netbox_available_asn" "second" {
  depends_on   = [netbox_available_asn.first]
  asn_range_id = netbox_asn_range.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					// NetBox hands out the range in order, and takes the RIR from the range.
					resource.TestCheckResourceAttr("netbox_available_asn.first", "asn", fmt.Sprint(start)),
					resource.TestCheckResourceAttr("netbox_available_asn.second", "asn", fmt.Sprint(start+1)),
					resource.TestCheckResourceAttrPair("netbox_available_asn.first", "rir_id", "netbox_rir.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_asn.first", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_available_asn.first", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttr("netbox_available_asn.first", "description", testName),
					resource.TestCheckResourceAttrSet("netbox_available_asn.first", "url"),
				),
			},
			{
				// Shrink: the optional attributes clear, while the allocated asn and the rir NetBox
				// derived from the range survive the update that clears them.
				Config: deps + `
resource "netbox_available_asn" "first" {
  asn_range_id = netbox_asn_range.test.id
}
resource "netbox_available_asn" "second" {
  depends_on   = [netbox_available_asn.first]
  asn_range_id = netbox_asn_range.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_asn.first", "asn", fmt.Sprint(start)),
					resource.TestCheckResourceAttrPair("netbox_available_asn.first", "rir_id", "netbox_rir.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_available_asn.first", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_available_asn.first", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_available_asn.first", "description"),
					resource.TestCheckNoResourceAttr("netbox_available_asn.first", "comments"),
				),
			},
			{
				// asn_range_id is an allocation input NetBox does not report back.
				ResourceName:            "netbox_available_asn.first",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"asn_range_id"},
			},
		},
	})
}

// TestAccNetboxAvailablePrefixDataSource_basic reads the free blocks of a /16 container after one
// /24 is allocated from it: the complement of the first /24 in a /16 is eight aligned blocks.
func TestAccNetboxAvailablePrefixDataSource_basic(t *testing.T) {
	net := fmt.Sprintf("10.%d", acctest.RandIntRange(251, 254))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "parent" {
  prefix = "%[1]s.0.0/16"
  status = "container"
}
resource "netbox_available_prefix" "test" {
  parent_prefix_id = netbox_prefix.parent.id
  prefix_length    = 24
}
data "netbox_available_prefix" "test" {
  depends_on = [netbox_available_prefix.test]
  prefix_id  = netbox_prefix.parent.id
}`, net),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_prefix.test", "prefix", net+".0.0/24"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test", "available_prefixes.#", "8"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test", "available_prefixes.0.prefix", net+".1.0/24"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test", "available_prefixes.0.family", "4"),
					resource.TestCheckNoResourceAttr("data.netbox_available_prefix.test", "available_prefixes.0.vrf_id"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test", "available_prefixes.7.prefix", net+".128.0/17"),
				),
			},
		},
	})
}
