//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxCircuit_basic(t *testing.T) {
	testName := testAccGetTestName("circuit")
	deps := fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
}
resource "netbox_circuit_type" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_circuit_provider_account" "test" {
  circuit_provider_id = netbox_circuit_provider.test.id
  account             = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_circuit" "test" {
  cid                         = "%[1]s"
  circuit_provider_id         = netbox_circuit_provider.test.id
  circuit_provider_account_id = netbox_circuit_provider_account.test.id
  circuit_type_id             = netbox_circuit_type.test.id
  tenant_id                   = netbox_tenant.test.id
  status                      = "active"
  install_date                = "2024-05-01"
  termination_date            = "2030-05-01"
  commit_rate_kbps            = 1000000
  distance                    = 12.5
  distance_unit               = "km"
  description                 = "Acceptance test circuit."
  comments                    = "Created by acceptance test."
}
data "netbox_circuit" "test" {
  cid = netbox_circuit.test.cid
}
data "netbox_circuits" "test" {
  filters = [
    { name = "provider_id", value = netbox_circuit_provider.test.id },
  ]
  depends_on          = [netbox_circuit.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit.test", "cid", testName),
					resource.TestCheckResourceAttr("netbox_circuit.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_circuit.test", "install_date", "2024-05-01"),
					resource.TestCheckResourceAttr("netbox_circuit.test", "termination_date", "2030-05-01"),
					resource.TestCheckResourceAttr("netbox_circuit.test", "commit_rate_kbps", "1000000"),
					resource.TestCheckResourceAttr("netbox_circuit.test", "distance_unit", "km"),
					resource.TestCheckResourceAttr("netbox_circuit.test", "distance", "12.5"),
					resource.TestCheckResourceAttrPair("netbox_circuit.test", "circuit_provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_circuit.test", "circuit_provider_account_id", "netbox_circuit_provider_account.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit.test", "id", "netbox_circuit.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuits.test", "circuits.#", "1"),
				),
			},
			{
				// Shrink: the plain optionals clear. status, install_date and termination_date are
				// computed and keep what NetBox holds.
				Config: deps + fmt.Sprintf(`
resource "netbox_circuit" "test" {
  cid                 = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
  circuit_type_id     = netbox_circuit_type.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_circuit.test", "install_date", "2024-05-01"),
					resource.TestCheckNoResourceAttr("netbox_circuit.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_circuit.test", "circuit_provider_account_id"),
					resource.TestCheckNoResourceAttr("netbox_circuit.test", "commit_rate_kbps"),
					resource.TestCheckNoResourceAttr("netbox_circuit.test", "distance_unit"),
					resource.TestCheckNoResourceAttr("netbox_circuit.test", "distance"),
					resource.TestCheckNoResourceAttr("netbox_circuit.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_circuit.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_circuit.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxCircuitTermination_basic(t *testing.T) {
	testName := testAccGetTestName("circuitterm")
	deps := fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}
resource "netbox_circuit_type" "test" {
  name = "%[1]s"
  slug = "%[2]s"
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
resource "netbox_circuit_provider_network" "test" {
  name                = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
}
resource "netbox_circuit" "test" {
  cid                 = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
  circuit_type_id     = netbox_circuit_type.test.id
}`, testName, getSlug(testName))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_circuit_termination" "test" {
  circuit_id          = netbox_circuit.test.id
  term_side           = "A"
  site_id             = netbox_site.test.id
  port_speed_kbps     = 1000000
  upstream_speed_kbps = 500000
  xconnect_id         = "XC-4711"
  pp_info             = "Panel 3, port 12"
  mark_connected      = true
  description         = "Acceptance test termination."
}
data "netbox_circuit_termination" "test" {
  id = netbox_circuit_termination.test.id
}
data "netbox_circuit_terminations" "test" {
  depends_on = [netbox_circuit_termination.test]
  filters = [
    { name = "circuit_id", value = netbox_circuit.test.id },
  ]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "term_side", "A"),
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "termination_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_circuit_termination.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_circuit_termination.test", "provider_network_id"),
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "xconnect_id", "XC-4711"),
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "mark_connected", "true"),
					resource.TestCheckResourceAttrPair("netbox_circuit_termination.test", "termination_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_termination.test", "id", "netbox_circuit_termination.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_terminations.test", "circuit_terminations.#", "1"),
				),
			},
			{
				// Shrink: the optionals clear, mark_connected is computed and keeps its value.
				Config: deps + `
resource "netbox_circuit_termination" "test" {
  circuit_id       = netbox_circuit.test.id
  term_side        = "A"
  termination_type = "dcim.site"
  termination_id   = netbox_site.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_circuit_termination.test", "port_speed_kbps"),
					resource.TestCheckNoResourceAttr("netbox_circuit_termination.test", "upstream_speed_kbps"),
					resource.TestCheckNoResourceAttr("netbox_circuit_termination.test", "xconnect_id"),
					resource.TestCheckNoResourceAttr("netbox_circuit_termination.test", "pp_info"),
					resource.TestCheckNoResourceAttr("netbox_circuit_termination.test", "description"),
				),
			},
			{
				// Cycle the termination through each alias and the provider network.
				Config: deps + `
resource "netbox_circuit_termination" "test" {
  circuit_id = netbox_circuit.test.id
  term_side  = "A"
  region_id  = netbox_region.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "termination_type", "dcim.region"),
					resource.TestCheckResourceAttrPair("netbox_circuit_termination.test", "region_id", "netbox_region.test", "id"),
				),
			},
			{
				Config: deps + `
resource "netbox_circuit_termination" "test" {
  circuit_id    = netbox_circuit.test.id
  term_side     = "A"
  site_group_id = netbox_site_group.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "termination_type", "dcim.sitegroup"),
					resource.TestCheckResourceAttrPair("netbox_circuit_termination.test", "site_group_id", "netbox_site_group.test", "id"),
				),
			},
			{
				Config: deps + `
resource "netbox_circuit_termination" "test" {
  circuit_id  = netbox_circuit.test.id
  term_side   = "A"
  location_id = netbox_location.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "termination_type", "dcim.location"),
					resource.TestCheckResourceAttrPair("netbox_circuit_termination.test", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				Config: deps + `
resource "netbox_circuit_termination" "test" {
  circuit_id          = netbox_circuit.test.id
  term_side           = "A"
  provider_network_id = netbox_circuit_provider_network.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_termination.test", "termination_type", "circuits.providernetwork"),
					resource.TestCheckResourceAttrPair("netbox_circuit_termination.test", "provider_network_id", "netbox_circuit_provider_network.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_circuit_termination.test", "site_id"),
				),
			},
			{
				ResourceName:      "netbox_circuit_termination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccNetboxCircuitType_color: color_hex is set and clears when unset.
func TestAccNetboxCircuitType_color(t *testing.T) {
	testName := testAccGetTestName("circuittype_color")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_circuit_type" "test" {
  name      = "%[1]s"
  color_hex = "ff0000"
}`, testName),
				Check: resource.TestCheckResourceAttr("netbox_circuit_type.test", "color_hex", "ff0000"),
			},
			{
				ResourceName:      "netbox_circuit_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_circuit_type" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_circuit_type.test", "color_hex"),
			},
		},
	})
}

// TestAccNetboxCircuitProvider_asns: the ASN id set grows, shrinks and clears when unset.
func TestAccNetboxCircuitProvider_asns(t *testing.T) {
	testName := testAccGetTestName("provider_asns")
	asn := acctest.RandIntRange(4200000000, 4200090000)
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
resource "netbox_circuit_provider" "test" {
  name    = "%[1]s"
  asn_ids = %[4]s
}`, testName, asn, asn+1, asns)
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config("[netbox_asn.a.id, netbox_asn.b.id]"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_provider.test", "asn_ids.#", "2"),
					resource.TestCheckTypeSetElemAttrPair("netbox_circuit_provider.test", "asn_ids.*", "netbox_asn.a", "id"),
					resource.TestCheckTypeSetElemAttrPair("netbox_circuit_provider.test", "asn_ids.*", "netbox_asn.b", "id"),
				),
			},
			{
				ResourceName:      "netbox_circuit_provider.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: config("[netbox_asn.b.id]"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_provider.test", "asn_ids.#", "1"),
					resource.TestCheckTypeSetElemAttrPair("netbox_circuit_provider.test", "asn_ids.*", "netbox_asn.b", "id"),
				),
			},
			{
				// Shrink: unset (null) clears the assignment.
				Config: config("null"),
				Check:  resource.TestCheckNoResourceAttr("netbox_circuit_provider.test", "asn_ids"),
			},
		},
	})
}

func TestAccNetboxCircuitProviderNetwork_basic(t *testing.T) {
	testName := testAccGetTestName("provnet")
	deps := fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_circuit_provider_network" "test" {
  name                = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
  service_id          = "SVC-4711"
  description         = "Acceptance test provider network."
  comments            = "Created by acceptance test."
}
data "netbox_circuit_provider_network" "test" {
  name = netbox_circuit_provider_network.test.name
}
data "netbox_circuit_provider_networks" "test" {
  filters = [
    { name = "name", value = netbox_circuit_provider_network.test.name },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "service_id", "SVC-4711"),
					resource.TestCheckResourceAttrPair("netbox_circuit_provider_network.test", "circuit_provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_provider_network.test", "id", "netbox_circuit_provider_network.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "circuit_provider_networks.#", "1"),
				),
			},
			{
				// Drop the optionals: they must clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_circuit_provider_network" "test" {
  name                = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_circuit_provider_network.test", "service_id"),
					resource.TestCheckNoResourceAttr("netbox_circuit_provider_network.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_circuit_provider_network.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_circuit_provider_network.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
