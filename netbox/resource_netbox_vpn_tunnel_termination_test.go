package netbox

import (
	"fmt"
	"log"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVpnTunnelTermination_basic(t *testing.T) {
	testSlug := "vpnterm_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "ff00ff"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}
resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  name = "eth0"
  device_id = netbox_device.test.id
  type = "virtual"
}
resource "netbox_ip_address" "device_1" {
  ip_address = "2.2.2.0/32"
  status = "active"
  device_interface_id = netbox_device_interface.test.id
}
resource "netbox_ip_address" "device_2" {
  ip_address = "2.2.2.1/32"
  status = "active"
  device_interface_id = netbox_device_interface.test.id
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_interface" "test" {
  name = "eth0"
  virtual_machine_id = netbox_virtual_machine.test.id
}
resource "netbox_ip_address" "vm_1" {
  ip_address = "2.2.2.2/32"
  status = "active"
  virtual_machine_interface_id = netbox_interface.test.id
}
resource "netbox_ip_address" "vm_2" {
  ip_address = "2.2.2.3/32"
  status = "active"
  virtual_machine_interface_id = netbox_interface.test.id
}

resource "netbox_vpn_tunnel_group" "test" {
  name = "%[1]s"
  description = "%[1]s"
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_vpn_tunnel" "test" {
  name = "%[1]s"
  encapsulation = "ipsec-transport"
  status = "active"
  tunnel_group_id = netbox_vpn_tunnel_group.test.id
}
resource "netbox_vpn_tunnel_termination" "device" {
  role = "peer"
  tunnel_id = netbox_vpn_tunnel.test.id
  device_interface_id = netbox_device_interface.test.id

  tags = [netbox_tag.test.name]
}
resource "netbox_vpn_tunnel_termination" "vm" {
  role = "peer"
  tunnel_id = netbox_vpn_tunnel.test.id
  virtual_machine_interface_id = netbox_interface.test.id

  outside_ip_address_id = netbox_ip_address.vm_1.id
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.device", "tunnel_id", "netbox_vpn_tunnel.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.vm", "tunnel_id", "netbox_vpn_tunnel.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.vm", "outside_ip_address_id", "netbox_ip_address.vm_1", "id"),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_termination.device", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_termination.device", "tags.0", testName),
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

func init() {
	resource.AddTestSweepers("netbox_vpn_tunnel_termination", &resource.Sweeper{
		Name:         "netbox_vpn_tunnel_termination",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := vpn.NewVpnTunnelTerminationsListParams()
			res, err := api.Vpn.VpnTunnelTerminationsList(params, nil)
			if err != nil {
				return err
			}
			for _, vpnTunnelTermination := range res.GetPayload().Results {
                                deleteParams := vpn.NewVpnTunnelTerminationsDeleteParams().WithID(vpnTunnelTermination.ID)
                                _, err := api.Vpn.VpnTunnelTerminationsDelete(deleteParams, nil)
                                if err != nil {
                                        return err
                                }
                                log.Print("[DEBUG] Deleted a vpn tunnel termination")
			}
			return nil
		},
	})
}
