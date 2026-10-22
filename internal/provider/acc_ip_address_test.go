//go:build acctest

package provider_test

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccNetboxIPAddress_basic(t *testing.T) {
	testName := testAccGetTestName("ip_basic")
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
resource "netbox_ip_address" "test" {
  ip_address  = "10.0.0.60/24"
  status      = "reserved"
  role        = "vip"
  dns_name    = "%[2]s.example.com"
  description = "%[1]s"
  comments    = "Created by acceptance test."
  tenant_id   = netbox_tenant.test.id
  vrf_id      = netbox_vrf.test.id
  tags        = [netbox_tag.test.slug]
}
resource "netbox_ip_address" "outside" {
  ip_address    = "192.0.2.60/32"
  nat_inside_id = netbox_ip_address.test.id
}
data "netbox_ip_addresses" "by_tag" {
  filters = [
    { name = "tag", value = netbox_tag.test.slug },
  ]
  depends_on = [netbox_ip_address.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", "10.0.0.60/24"),
					// The tag is this test's alone, so the tag lookup finds exactly our address.
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.by_tag", "ip_addresses.0.id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "role", "vip"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "dns_name", getSlug(testName)+".example.com"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "family", "4"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tags.#", "1"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "assigned_object"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.outside", "nat_inside_id", "netbox_ip_address.test", "id"),
				),
			},
			{
				// The inside address learns its NAT peer on the next read; the config is unchanged.
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
resource "netbox_ip_address" "test" {
  ip_address  = "10.0.0.60/24"
  status      = "reserved"
  role        = "vip"
  dns_name    = "%[2]s.example.com"
  description = "%[1]s"
  comments    = "Created by acceptance test."
  tenant_id   = netbox_tenant.test.id
  vrf_id      = netbox_vrf.test.id
  tags        = [netbox_tag.test.slug]
}
resource "netbox_ip_address" "outside" {
  ip_address    = "192.0.2.60/32"
  nat_inside_id = netbox_ip_address.test.id
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "nat_outside_ids.#", "1"),
					resource.TestCheckTypeSetElemAttrPair("netbox_ip_address.test", "nat_outside_ids.*", "netbox_ip_address.outside", "id"),
				),
			},
			{
				// Shrink: everything optional clears on both addresses (status is computed and
				// keeps its value).
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
resource "netbox_ip_address" "test" {
  ip_address = "10.0.0.60/24"
}
resource "netbox_ip_address" "outside" {
  ip_address = "192.0.2.60/32"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "reserved"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "role"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "dns_name"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "vrf_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "tags"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.outside", "nat_inside_id"),
				),
			},
			{
				// Settle: the shrink updates both addresses in parallel (nothing links them any
				// more), so whether the inside address's update response still carried the NAT
				// peer is a race. As in the step above that learns the peer, an unchanged config
				// re-reads and forgets it; without this the import verify flakes on a stale
				// nat_outside_ids.
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
resource "netbox_ip_address" "test" {
  ip_address = "10.0.0.60/24"
}
resource "netbox_ip_address" "outside" {
  ip_address = "192.0.2.60/32"
}`, testName),
				Check: resource.TestCheckResourceAttr("netbox_ip_address.test", "nat_outside_ids.#", "0"),
			},
			{
				ResourceName:      "netbox_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIPAddress_ipv6(t *testing.T) {
	testName := testAccGetTestName("ip_v6")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address  = "2001:db8::10/64"
  description = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "family", "6"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
				),
			},
		},
	})
}

// TestAccNetboxIPAddress_assignment covers the assigned_object hook: the
// virtual_machine_interface_id / device_interface_id aliases and the generic
// assigned_object_type/_id pair derive each other in both directions.
func TestAccNetboxIPAddress_assignment(t *testing.T) {
	testName := testAccGetTestName("ip_assign")
	virtualMachineConfig := fmt.Sprintf(`
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
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				// Alias set: the pair is derived, in the same apply that creates the interface (the id is
				// unknown at plan time).
				Config: virtualMachineConfig + `
resource "netbox_ip_address" "test" {
  ip_address      = "10.0.0.61/24"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}
data "netbox_ip_address" "test" {
  id = netbox_ip_address.test.id
}
data "netbox_ip_addresses" "test" {
  filters = [
    { name = "address", value = "10.0.0.61/24" },
  ]
  depends_on = [netbox_ip_address.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "assigned_object_type", "virtualization.vminterface"),
					// Data sources run the companion hooks too, so the aliases derive there as well.
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckNoResourceAttr("data.netbox_ip_address.test", "device_interface_id"),
					// 10.0.0.61 belongs to this test alone, so the first match is ours.
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test", "ip_addresses.0.virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "assigned_object_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "device_interface_id"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "assigned_object.id", "netbox_virtual_machine_interface.test", "id"),
				),
			},
			{
				// The same assignment written as the explicit pair is a no-op: the alias is derived from the
				// pair.
				Config: virtualMachineConfig + `
resource "netbox_ip_address" "test" {
  ip_address           = "10.0.0.61/24"
  assigned_object_type = "virtualization.vminterface"
  assigned_object_id   = netbox_virtual_machine_interface.test.id
}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				Check: resource.TestCheckResourceAttrPair("netbox_ip_address.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
			},
			{
				// Import: the aliases derive from what NetBox returns, so nothing needs to be ignored.
				ResourceName:      "netbox_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Move to the device interface.
				Config: virtualMachineConfig + `
resource "netbox_ip_address" "test" {
  ip_address          = "10.0.0.61/24"
  device_interface_id = netbox_device_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "assigned_object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "assigned_object_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "device_interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "virtual_machine_interface_id"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "assigned_object.name", "eth0"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "assigned_object.device.name", testName),
				),
			},
			{
				// Unassign: every attribute of the assignment clears.
				Config: virtualMachineConfig + `
resource "netbox_ip_address" "test" {
  ip_address = "10.0.0.61/24"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "assigned_object_type"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "assigned_object_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "virtual_machine_interface_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "device_interface_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_address.test", "assigned_object"),
				),
			},
			{
				// Both aliases at once is a plan-time error.
				Config: virtualMachineConfig + `
resource "netbox_ip_address" "test" {
  ip_address                   = "10.0.0.61/24"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
  device_interface_id          = netbox_device_interface.test.id
}`,
				ExpectError: regexp.MustCompile("Conflicting assignment"),
			},
			{
				// An alias next to the explicit pair as well.
				Config: virtualMachineConfig + `
resource "netbox_ip_address" "test" {
  ip_address           = "10.0.0.61/24"
  virtual_machine_interface_id      = netbox_virtual_machine_interface.test.id
  assigned_object_type = "virtualization.vminterface"
  assigned_object_id   = netbox_virtual_machine_interface.test.id
}`,
				ExpectError: regexp.MustCompile("Conflicting assignment"),
			},
		},
	})
}

func TestAccNetboxIPAddressDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("ip_ds")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address  = "10.0.0.107/24"
  status      = "active"
  dns_name    = "%[2]s.example.com"
  description = "%[1]s"
}
data "netbox_ip_address" "by_id" {
  id = netbox_ip_address.test.id
}
data "netbox_ip_address" "by_dns" {
  dns_name   = "%[2]s.example.com"
  depends_on = [netbox_ip_address.test]
}
data "netbox_ip_addresses" "list" {
  filters = [
    { name = "dns_name", value = "%[2]s.example.com" },
  ]
  depends_on = [netbox_ip_address.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.by_id", "ip_address", "netbox_ip_address.test", "ip_address"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.by_dns", "id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.list", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.list", "ip_addresses.0.id", "netbox_ip_address.test", "id"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_ip_address",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Ipam.IpamIPAddressesList(ipam.NewIpamIPAddressesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// Addresses have no name; the tests put the test name in the description.
				items = append(items, sweepItem{result.ID, result.Description})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Ipam.IpamIPAddressesDestroy(ipam.NewIpamIPAddressesDestroyParams().WithID(id), nil)
			return err
		})
}

// ip_address is normalize ip: the state keeps the configured spelling although NetBox echoes the
// address compressed and lowercased, the next plan is empty, and text that is no address fails at
// plan.
func TestAccNetboxIPAddress_normalizedSpelling(t *testing.T) {
	spelled := fmt.Sprintf("2001:DB8:0:%X:0:0::60/64", acctest.RandIntRange(0x100, 0xfff))
	config := fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
}`, spelled)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", spelled),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "family", "6"),
				),
			},
			{
				// Not the last step: the post-test destroy runs the last step's config.
				Config: `
resource "netbox_ip_address" "test" {
  ip_address = "10.0.0.300/24"
}`,
				PlanOnly:    true,
				ExpectError: regexp.MustCompile("(?i)invalid ip address"),
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

// TestAccNetboxIPAddresses_customFieldFilter filters ip_addresses by a text custom field with the
// default loose logic: the value matches as a case-insensitive substring.
func TestAccNetboxIPAddresses_customFieldFilter(t *testing.T) {
	// Custom field names are ^[a-z0-9_]+$.
	testName := strings.ToLower(strings.ReplaceAll(testAccGetTestName("ip_cf"), "-", "_"))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name         = "%[1]s"
  type         = "text"
  object_types = ["ipam.ipaddress"]
}
resource "netbox_ip_address" "test_a" {
  ip_address = "10.0.0.201/24"
  custom_fields = {
    (netbox_custom_field.test.name) = "match"
  }
}
resource "netbox_ip_address" "test_b" {
  ip_address = "10.0.0.202/24"
  custom_fields = {
    (netbox_custom_field.test.name) = "other"
  }
}
data "netbox_ip_addresses" "match" {
  depends_on = [netbox_ip_address.test_a, netbox_ip_address.test_b]
  filters = [
    { name = "cf_%[1]s", value = "match" },
  ]
}
data "netbox_ip_addresses" "contains" {
  depends_on = [netbox_ip_address.test_a, netbox_ip_address.test_b]
  filters = [
    { name = "cf_%[1]s", value = "ATC" },
  ]
}
data "netbox_ip_addresses" "none" {
  depends_on = [netbox_ip_address.test_a, netbox_ip_address.test_b]
  filters = [
    { name = "cf_%[1]s", value = "nomatch" },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.match", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.match", "ip_addresses.0.id", "netbox_ip_address.test_a", "id"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.contains", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.contains", "ip_addresses.0.id", "netbox_ip_address.test_a", "id"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.none", "ip_addresses.#", "0"),
				),
			},
		},
	})
}

// TestAccNetboxIPAddress_customFieldLookup looks an ip_address up by a custom field value, alone and
// combined with another input; a value both addresses carry does not match exactly one.
func TestAccNetboxIPAddress_customFieldLookup(t *testing.T) {
	// Custom field names are ^[a-z0-9_]+$.
	testName := strings.ToLower(strings.ReplaceAll(testAccGetTestName("ip_cf_one"), "-", "_"))
	deps := fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name         = "%[1]s"
  type         = "text"
  object_types = ["ipam.ipaddress"]
}
resource "netbox_ip_address" "test_a" {
  ip_address = "10.0.0.231/24"
  status     = "reserved"
  custom_fields = {
    (netbox_custom_field.test.name) = "match"
  }
}
resource "netbox_ip_address" "test_b" {
  ip_address = "10.0.0.232/24"
  status     = "active"
  custom_fields = {
    (netbox_custom_field.test.name) = "other"
  }
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
data "netbox_ip_address" "by_value" {
  depends_on = [netbox_ip_address.test_a, netbox_ip_address.test_b]
  custom_field_filters = {
    (netbox_custom_field.test.name) = "match"
  }
}
data "netbox_ip_address" "by_value_and_status" {
  depends_on = [netbox_ip_address.test_a, netbox_ip_address.test_b]
  status     = "active"
  custom_field_filters = {
    (netbox_custom_field.test.name) = "other"
  }
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.by_value", "id", "netbox_ip_address.test_a", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.by_value_and_status", "id", "netbox_ip_address.test_b", "id"),
				),
			},
			{
				// The field's loose logic matches "t" in both values.
				Config: deps + `
data "netbox_ip_address" "ambiguous" {
  depends_on = [netbox_ip_address.test_a, netbox_ip_address.test_b]
  custom_field_filters = {
    (netbox_custom_field.test.name) = "t"
  }
}`,
				ExpectError: regexp.MustCompile(`did not match exactly one`),
			},
		},
	})
}
