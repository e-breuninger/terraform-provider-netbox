//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxVpnTunnelGroup_basic(t *testing.T) {
	testName := testAccGetTestName("tunnelgrp")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vpn_tunnel_group" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "Acceptance test tunnel group."
  comments    = "Acceptance test comments."
}
data "netbox_vpn_tunnel_group" "test" {
  name = netbox_vpn_tunnel_group.test.name
}
data "netbox_vpn_tunnel_groups" "test" {
  filters = [
    { name = "name", value = netbox_vpn_tunnel_group.test.name },
  ]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_group.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttrPair("data.netbox_vpn_tunnel_group.test", "id", "netbox_vpn_tunnel_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vpn_tunnel_groups.test", "vpn_tunnel_groups.#", "1"),
				),
			},
			{
				// Drop description: it must clear.
				Config: fmt.Sprintf(`
resource "netbox_vpn_tunnel_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel_group.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel_group.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_vpn_tunnel_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVpnTunnel_basic(t *testing.T) {
	testName := testAccGetTestName("tunnel")
	deps := fmt.Sprintf(`
resource "netbox_vpn_tunnel_group" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_ike_proposal" "test" {
  name                  = "%[1]s"
  authentication_method = "preshared-keys"
  encryption_algorithm  = "aes-256-cbc"
  group                 = 14
}
resource "netbox_ike_policy" "test" {
  name             = "%[1]s"
  version          = 2
  ike_proposal_ids = [netbox_ike_proposal.test.id]
}
resource "netbox_ipsec_proposal" "test" {
  name                 = "%[1]s"
  encryption_algorithm = "aes-256-cbc"
}
resource "netbox_ipsec_policy" "test" {
  name               = "%[1]s"
  ipsec_proposal_ids = [netbox_ipsec_proposal.test.id]
}
resource "netbox_ipsec_profile" "test" {
  name            = "%[1]s"
  mode            = "esp"
  ike_policy_id   = netbox_ike_policy.test.id
  ipsec_policy_id = netbox_ipsec_policy.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_vpn_tunnel" "test" {
  name                = "%[1]s"
  encapsulation       = "ipsec-tunnel"
  status              = "active"
  vpn_tunnel_group_id = netbox_vpn_tunnel_group.test.id
  ipsec_profile_id    = netbox_ipsec_profile.test.id
  tenant_id           = netbox_tenant.test.id
  tunnel_id           = 42
  description         = "Acceptance test tunnel."
  comments            = "Created by acceptance test."
}
data "netbox_vpn_tunnel" "test" {
  name = netbox_vpn_tunnel.test.name
}
data "netbox_vpn_tunnels" "test" {
  filters = [
    { name = "group_id", value = netbox_vpn_tunnel_group.test.id },
  ]
  depends_on          = [netbox_vpn_tunnel.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel.test", "encapsulation", "ipsec-tunnel"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel.test", "ipsec_profile_id", "netbox_ipsec_profile.test", "id"),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel.test", "tunnel_id", "42"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel.test", "vpn_tunnel_group_id", "netbox_vpn_tunnel_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vpn_tunnel.test", "id", "netbox_vpn_tunnel.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vpn_tunnels.test", "vpn_tunnels.#", "1"),
				),
			},
			{
				// Shrink: encapsulation changes, everything optional clears (status is computed and stays).
				Config: deps + fmt.Sprintf(`
resource "netbox_vpn_tunnel" "test" {
  name          = "%s"
  encapsulation = "ip-ip"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel.test", "encapsulation", "ip-ip"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel.test", "vpn_tunnel_group_id"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel.test", "ipsec_profile_id"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel.test", "tunnel_id"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_vpn_tunnel.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVpnTunnelTermination_basic(t *testing.T) {
	testName := testAccGetTestName("tunnelterm")
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
resource "netbox_ip_address" "test" {
  ip_address           = "10.112.0.7/24"
  status               = "active"
  assigned_object_type = "virtualization.vminterface"
  assigned_object_id   = netbox_virtual_machine_interface.test.id
}
resource "netbox_vpn_tunnel" "test" {
  name          = "%[1]s"
  encapsulation = "ipsec-tunnel"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_vpn_tunnel_termination" "test" {
  vpn_tunnel_id    = netbox_vpn_tunnel.test.id
  role             = "hub"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
  outside_ip_address_id        = netbox_ip_address.test.id
}
data "netbox_vpn_tunnel_termination" "test" {
  id = netbox_vpn_tunnel_termination.test.id
}
data "netbox_vpn_tunnel_terminations" "test" {
  depends_on = [netbox_vpn_tunnel_termination.test]
  filters = [
    { name = "tunnel_id", value = netbox_vpn_tunnel.test.id },
  ]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_termination.test", "role", "hub"),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_termination.test", "termination_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.test", "vpn_tunnel_id", "netbox_vpn_tunnel.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.test", "termination_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.test", "outside_ip_address_id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.test", "virtual_machine_interface_id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel_termination.test", "device_interface_id"),
					resource.TestCheckResourceAttrPair("data.netbox_vpn_tunnel_termination.test", "id", "netbox_vpn_tunnel_termination.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vpn_tunnel_terminations.test", "vpn_tunnel_terminations.#", "1"),
				),
			},
			{
				// Drop the outside IP: it must clear (role is computed and stays).
				Config: deps + `
resource "netbox_vpn_tunnel_termination" "test" {
  vpn_tunnel_id    = netbox_vpn_tunnel.test.id
  termination_type = "virtualization.vminterface"
  termination_id   = netbox_virtual_machine_interface.test.id
}`,
				// The explicit pair names the same assignment as the alias.
				Check: resource.TestCheckNoResourceAttr("netbox_vpn_tunnel_termination.test", "outside_ip_address_id"),
			},
			{
				// Move to the device interface via the alias.
				Config: deps + `
resource "netbox_vpn_tunnel_termination" "test" {
  vpn_tunnel_id       = netbox_vpn_tunnel.test.id
  device_interface_id = netbox_device_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_termination.test", "termination_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.test", "device_interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_vpn_tunnel_termination.test", "virtual_machine_interface_id"),
				),
			},
			{
				ResourceName:      "netbox_vpn_tunnel_termination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_vpn_tunnel_termination",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnTunnelTerminationsList(vpn.NewVpnTunnelTerminationsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// Terminations have no name; key on the parent tunnel's.
				name := ""
				if result.Tunnel != nil {
					name = deref(result.Tunnel.Name)
				}
				items = append(items, sweepItem{result.ID, name})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Vpn.VpnTunnelTerminationsDestroy(vpn.NewVpnTunnelTerminationsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_vpn_tunnel",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnTunnelsList(vpn.NewVpnTunnelsListParams(), nil)
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
			_, err := client.Vpn.VpnTunnelsDestroy(vpn.NewVpnTunnelsDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_vpn_tunnel_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Vpn.VpnTunnelGroupsList(vpn.NewVpnTunnelGroupsListParams(), nil)
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
			_, err := client.Vpn.VpnTunnelGroupsDestroy(vpn.NewVpnTunnelGroupsDestroyParams().WithID(id), nil)
			return err
		})
}
