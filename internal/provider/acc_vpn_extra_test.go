//go:build acctest

// IKE proposals and L2VPNs (the go-netbox 2026-08-29 batch).
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

func TestAccNetboxIKEProposal_basic(t *testing.T) {
	testName := testAccGetTestName("ikeprop")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                     = "%[1]s"
  authentication_method    = "preshared-keys"
  encryption_algorithm     = "aes-256-cbc"
  authentication_algorithm = "hmac-sha256"
  group                    = 14
  sa_lifetime              = 28800
  description              = "Acceptance test IKE proposal."
  comments                 = "Created by acceptance test."
}
data "netbox_ike_proposal" "test" {
  name = netbox_ike_proposal.test.name
}
data "netbox_ike_proposals" "test" {
  # Scoped by name as well as by the algorithm: other tests create proposals with
  # aes-256-cbc too, and on an unthrottled run they overlap with this one.
  filters = [
    { name = "name", value = "%[1]s" },
    { name = "encryption_algorithm", value = "aes-256-cbc" },
  ]
  depends_on           = [netbox_ike_proposal.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "authentication_method", "preshared-keys"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "encryption_algorithm", "aes-256-cbc"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "group", "14"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "sa_lifetime", "28800"),
					resource.TestCheckResourceAttrPair("data.netbox_ike_proposal.test", "id", "netbox_ike_proposal.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ike_proposals.test", "ike_proposals.#", "1"),
				),
			},
			{
				// Everything optional clears, except authentication_algorithm which NetBox refuses to blank (it
				// stays).
				Config: fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                  = "%s"
  authentication_method = "certificates"
  encryption_algorithm  = "aes-128-gcm"
  group                 = 19
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "encryption_algorithm", "aes-128-gcm"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "group", "19"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckNoResourceAttr("netbox_ike_proposal.test", "sa_lifetime"),
					resource.TestCheckNoResourceAttr("netbox_ike_proposal.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_ike_proposal.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_ike_proposal.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxL2VPN_basic(t *testing.T) {
	testName := testAccGetTestName("l2vpn")
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_route_target" "test" {
  name = "65000:%[2]d"
}`, testName, 4000+len(testName))
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_l2vpn" "test" {
  name           = "%[1]s"
  slug           = "%[2]s"
  type           = "vxlan-evpn"
  identifier     = 10042
  tenant_id      = netbox_tenant.test.id
  import_targets = [netbox_route_target.test.id]
  export_targets = [netbox_route_target.test.id]
  description    = "Acceptance test L2VPN."
  comments       = "Created by acceptance test."
}
data "netbox_l2vpn" "test" {
  name = netbox_l2vpn.test.name
}
data "netbox_l2vpns" "test" {
  # Scoped by name as well as by the type: nothing else creates a vxlan-evpn L2VPN today, but the
  # count must not depend on that staying true while the suite runs in parallel.
  filters = [
    { name = "name", value = "%[1]s" },
    { name = "type", value = "vxlan-evpn" },
  ]
  depends_on = [netbox_l2vpn.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "type", "vxlan-evpn"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "identifier", "10042"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "import_targets.#", "1"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "export_targets.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_l2vpn.test", "id", "netbox_l2vpn.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_l2vpns.test", "l2vpns.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_l2vpn" "test" {
  name = "%s"
  type = "vpls"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "type", "vpls"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn.test", "identifier"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn.test", "import_targets"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn.test", "export_targets"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_l2vpn.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccNetboxL2VPNTermination_basic: the assigned_object hook aliases (vlan_id,
// virtual_machine_interface_id, device_interface_id), moving between them, and the explicit pair as
// a no-op.
func TestAccNetboxL2VPNTermination_basic(t *testing.T) {
	testName := testAccGetTestName("l2vpnterm")
	deps := deviceDeps(testName) + fmt.Sprintf(`
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
resource "netbox_l2vpn" "test" {
  name = "%[1]s"
  type = "vxlan"
}
resource "netbox_vlan" "test" {
  name = "%[1]s"
  vid  = 3999
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
				Config: deps + `
resource "netbox_l2vpn_termination" "test" {
  l2vpn_id = netbox_l2vpn.test.id
  vlan_id  = netbox_vlan.test.id
}
data "netbox_l2vpn_termination" "test" {
  id = netbox_l2vpn_termination.test.id
}
data "netbox_l2vpn_terminations" "test" {
  filters = [
    { name = "l2vpn_id", value = netbox_l2vpn.test.id },
  ]
  depends_on = [netbox_l2vpn_termination.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_l2vpn_termination.test", "assigned_object_type", "ipam.vlan"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "assigned_object_id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "vlan_id", "netbox_vlan.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn_termination.test", "virtual_machine_interface_id"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn_termination.test", "device_interface_id"),
					resource.TestCheckResourceAttrPair("data.netbox_l2vpn_termination.test", "id", "netbox_l2vpn_termination.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn_terminations.test", "l2vpn_terminations.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_l2vpn_termination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + `
resource "netbox_l2vpn_termination" "test" {
  l2vpn_id                     = netbox_l2vpn.test.id
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_l2vpn_termination.test", "assigned_object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "assigned_object_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn_termination.test", "vlan_id"),
				),
			},
			{
				// Move to the device interface via the alias.
				Config: deps + `
resource "netbox_l2vpn_termination" "test" {
  l2vpn_id            = netbox_l2vpn.test.id
  device_interface_id = netbox_device_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_l2vpn_termination.test", "assigned_object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "device_interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_l2vpn_termination.test", "virtual_machine_interface_id"),
				),
			},
			{
				// The same assignment written as the explicit pair is a no-op.
				Config: deps + `
resource "netbox_l2vpn_termination" "test" {
  l2vpn_id             = netbox_l2vpn.test.id
  assigned_object_type = "dcim.interface"
  assigned_object_id   = netbox_device_interface.test.id
}`,
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
				Check: resource.TestCheckResourceAttrPair("netbox_l2vpn_termination.test", "assigned_object_id", "netbox_device_interface.test", "id"),
			},
		},
	})
}

func init() {
	sweep("netbox_l2vpn_termination",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnL2vpnTerminationsList(vpn.NewVpnL2vpnTerminationsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				name := ""
				if result.L2vpn != nil {
					name = deref(result.L2vpn.Name)
				}
				items = append(items, sweepItem{id: result.ID, name: name})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Vpn.VpnL2vpnTerminationsDestroy(vpn.NewVpnL2vpnTerminationsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_l2vpn",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnL2vpnsList(vpn.NewVpnL2vpnsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{id: result.ID, name: deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Vpn.VpnL2vpnsDestroy(vpn.NewVpnL2vpnsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_ike_proposal",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnIkeProposalsList(vpn.NewVpnIkeProposalsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{id: result.ID, name: deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Vpn.VpnIkeProposalsDestroy(vpn.NewVpnIkeProposalsDestroyParams().WithID(id), nil)
			return err
		})
}
