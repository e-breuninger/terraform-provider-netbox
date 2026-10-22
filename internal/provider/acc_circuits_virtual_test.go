//go:build acctest

// The circuits objects NetBox 4.6 added: provider accounts, circuit groups and their assignments,
// and the virtual circuit family.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxCircuitProviderAccount_basic(t *testing.T) {
	testName := testAccGetTestName("provacct")
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
resource "netbox_circuit_provider_account" "test" {
  account             = "%[1]s"
  name                = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
  description         = "Acceptance test provider account."
  comments            = "Created by acceptance test."
}
data "netbox_circuit_provider_account" "test" {
  name = netbox_circuit_provider_account.test.name
}
data "netbox_circuit_provider_accounts" "test" {
  filters = [
    { name = "provider_id", value = netbox_circuit_provider.test.id },
  ]
  depends_on          = [netbox_circuit_provider_account.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_provider_account.test", "account", testName),
					resource.TestCheckResourceAttrPair("netbox_circuit_provider_account.test", "circuit_provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_provider_account.test", "id", "netbox_circuit_provider_account.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_accounts.test", "circuit_provider_accounts.#", "1"),
				),
			},
			{
				// Shrink: name, description and comments clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_circuit_provider_account" "test" {
  account             = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_circuit_provider_account.test", "name"),
					resource.TestCheckNoResourceAttr("netbox_circuit_provider_account.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_circuit_provider_account.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_circuit_provider_account.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxCircuitGroup_basic(t *testing.T) {
	testName := testAccGetTestName("circuitgrp")
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_circuit_group" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  tenant_id   = netbox_tenant.test.id
  description = "Acceptance test circuit group."
  comments    = "Created by acceptance test."
}
data "netbox_circuit_group" "test" {
  name = netbox_circuit_group.test.name
}
data "netbox_circuit_groups" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_circuit_group.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttrPair("netbox_circuit_group.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_group.test", "id", "netbox_circuit_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_groups.test", "circuit_groups.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_circuit_group" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_circuit_group.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_circuit_group.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_circuit_group.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_circuit_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxCircuitGroupAssignment_basic(t *testing.T) {
	testName := testAccGetTestName("cgassign")
	deps := fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
}
resource "netbox_circuit_type" "test" {
  name = "%[1]s"
}
resource "netbox_circuit" "test" {
  cid                 = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
  circuit_type_id     = netbox_circuit_type.test.id
  status              = "active"
}
resource "netbox_circuit_group" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_circuit_group_assignment" "test" {
  circuit_group_id = netbox_circuit_group.test.id
  member_type      = "circuits.circuit"
  member_id        = netbox_circuit.test.id
  priority         = "primary"
}
data "netbox_circuit_group_assignment" "test" {
  id = netbox_circuit_group_assignment.test.id
}
data "netbox_circuit_group_assignments" "test" {
  filters = [
    { name = "group_id", value = netbox_circuit_group.test.id },
  ]
  depends_on       = [netbox_circuit_group_assignment.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_group_assignment.test", "member_type", "circuits.circuit"),
					resource.TestCheckResourceAttr("netbox_circuit_group_assignment.test", "priority", "primary"),
					resource.TestCheckResourceAttrPair("netbox_circuit_group_assignment.test", "member_id", "netbox_circuit.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_group_assignment.test", "id", "netbox_circuit_group_assignment.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group_assignments.test", "circuit_group_assignments.#", "1"),
				),
			},
			{
				// Shrink: priority is the only optional attribute and must clear.
				Config: deps + `
resource "netbox_circuit_group_assignment" "test" {
  circuit_group_id = netbox_circuit_group.test.id
  member_type      = "circuits.circuit"
  member_id        = netbox_circuit.test.id
}`,
				Check: resource.TestCheckNoResourceAttr("netbox_circuit_group_assignment.test", "priority"),
			},
			{
				ResourceName:      "netbox_circuit_group_assignment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVirtualCircuitType_basic(t *testing.T) {
	testName := testAccGetTestName("vctype")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_virtual_circuit_type" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "Acceptance test virtual circuit type."
  color_hex   = "ff0000"
  comments    = "Created by acceptance test."
}
data "netbox_virtual_circuit_type" "test" {
  name = netbox_virtual_circuit_type.test.name
}
data "netbox_virtual_circuit_types" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_virtual_circuit_type.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit_type.test", "id", "netbox_virtual_circuit_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_types.test", "virtual_circuit_types.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_virtual_circuit_type" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit_type.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit_type.test", "color_hex"),
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit_type.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_virtual_circuit_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// virtualCircuitDeps is a provider, its network and account, a virtual circuit type and a tenant.
func virtualCircuitDeps(testName string) string {
	return fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
}
resource "netbox_circuit_provider_network" "test" {
  name                = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
}
resource "netbox_circuit_provider_account" "test" {
  account             = "%[1]s"
  circuit_provider_id = netbox_circuit_provider.test.id
}
resource "netbox_virtual_circuit_type" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}`, testName)
}

func TestAccNetboxVirtualCircuit_basic(t *testing.T) {
	testName := testAccGetTestName("vcircuit")
	deps := virtualCircuitDeps(testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_circuit" "test" {
  cid                         = "%[1]s"
  circuit_provider_network_id = netbox_circuit_provider_network.test.id
  virtual_circuit_type_id     = netbox_virtual_circuit_type.test.id
  circuit_provider_account_id = netbox_circuit_provider_account.test.id
  tenant_id                   = netbox_tenant.test.id
  status                      = "active"
  description                 = "Acceptance test virtual circuit."
  comments                    = "Created by acceptance test."
}
data "netbox_virtual_circuit" "test" {
  cid = netbox_virtual_circuit.test.cid
}
data "netbox_virtual_circuits" "test" {
  filters = [
    { name = "provider_network_id", value = netbox_circuit_provider_network.test.id },
  ]
  depends_on                  = [netbox_virtual_circuit.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "cid", testName),
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_virtual_circuit.test", "virtual_circuit_type_id", "netbox_virtual_circuit_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit.test", "id", "netbox_virtual_circuit.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuits.test", "virtual_circuits.#", "1"),
				),
			},
			{
				// Shrink: status is optional+computed and keeps NetBox's value.
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_circuit" "test" {
  cid                         = "%[1]s"
  circuit_provider_network_id = netbox_circuit_provider_network.test.id
  virtual_circuit_type_id     = netbox_virtual_circuit_type.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "status", "active"),
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit.test", "circuit_provider_account_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_virtual_circuit.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVirtualCircuitTermination_basic(t *testing.T) {
	testName := testAccGetTestName("vcterm")
	deps := componentDeps(testName, "") + virtualCircuitDeps(testName) + fmt.Sprintf(`
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s"
  // NetBox terminates virtual circuits only on virtual interfaces.
  type = "virtual"
}
resource "netbox_virtual_circuit" "test" {
  cid                         = "%[1]s"
  circuit_provider_network_id = netbox_circuit_provider_network.test.id
  virtual_circuit_type_id     = netbox_virtual_circuit_type.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_virtual_circuit_termination" "test" {
  virtual_circuit_id  = netbox_virtual_circuit.test.id
  device_interface_id = netbox_device_interface.test.id
  role                = "hub"
  description         = "Acceptance test virtual circuit termination."
}
data "netbox_virtual_circuit_termination" "test" {
  id = netbox_virtual_circuit_termination.test.id
}
data "netbox_virtual_circuit_terminations" "test" {
  filters = [
    { name = "virtual_circuit_id", value = netbox_virtual_circuit.test.id },
  ]
  depends_on         = [netbox_virtual_circuit_termination.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_circuit_termination.test", "role", "hub"),
					resource.TestCheckResourceAttrPair("netbox_virtual_circuit_termination.test", "device_interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit_termination.test", "id", "netbox_virtual_circuit_termination.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_terminations.test", "virtual_circuit_terminations.#", "1"),
				),
			},
			{
				// Shrink: description clears, role is optional+computed and keeps its value.
				Config: deps + `
resource "netbox_virtual_circuit_termination" "test" {
  virtual_circuit_id  = netbox_virtual_circuit.test.id
  device_interface_id = netbox_device_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_circuit_termination.test", "role", "hub"),
					resource.TestCheckNoResourceAttr("netbox_virtual_circuit_termination.test", "description"),
				),
			},
			{
				ResourceName:      "netbox_virtual_circuit_termination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_virtual_circuit_termination",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsVirtualCircuitTerminationsList(circuits.NewCircuitsVirtualCircuitTerminationsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// Terminations have no name; key on the virtual circuit's cid.
				name := ""
				if result.VirtualCircuit != nil {
					name = deref(result.VirtualCircuit.Cid)
				}
				items = append(items, sweepItem{result.ID, name})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Circuits.CircuitsVirtualCircuitTerminationsDestroy(circuits.NewCircuitsVirtualCircuitTerminationsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_virtual_circuit",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsVirtualCircuitsList(circuits.NewCircuitsVirtualCircuitsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Cid)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Circuits.CircuitsVirtualCircuitsDestroy(circuits.NewCircuitsVirtualCircuitsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_virtual_circuit_type",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsVirtualCircuitTypesList(circuits.NewCircuitsVirtualCircuitTypesListParams(), nil)
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
			_, err := client.Circuits.CircuitsVirtualCircuitTypesDestroy(circuits.NewCircuitsVirtualCircuitTypesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_circuit_group_assignment",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsCircuitGroupAssignmentsList(circuits.NewCircuitsCircuitGroupAssignmentsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// Assignments have no name; key on the group's.
				name := ""
				if result.Group != nil {
					name = deref(result.Group.Name)
				}
				items = append(items, sweepItem{result.ID, name})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Circuits.CircuitsCircuitGroupAssignmentsDestroy(circuits.NewCircuitsCircuitGroupAssignmentsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_circuit_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsCircuitGroupsList(circuits.NewCircuitsCircuitGroupsListParams(), nil)
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
			_, err := client.Circuits.CircuitsCircuitGroupsDestroy(circuits.NewCircuitsCircuitGroupsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_circuit_provider_account",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Circuits.CircuitsProviderAccountsList(circuits.NewCircuitsProviderAccountsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Account)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Circuits.CircuitsProviderAccountsDestroy(circuits.NewCircuitsProviderAccountsDestroyParams().WithID(id), nil)
			return err
		})
}
