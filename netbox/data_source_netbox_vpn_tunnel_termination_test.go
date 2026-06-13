package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxTunnelTerminationsSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_interface" "test1" {
  name      = "eth0"
  device_id = netbox_device.test.id
  type      = "1000base-t"
}
resource "netbox_device_interface" "test2" {
  name      = "eth1"
  device_id = netbox_device.test.id
  type      = "1000base-t"
}
resource "netbox_vpn_tunnel_group" "test" {
  name = "%[1]s"
  description = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
resource "netbox_vpn_tunnel" "test" {
  name = "%[1]s"
  encapsulation = "ipsec-transport"
  status = "active"
  tunnel_group_id = netbox_vpn_tunnel_group.test.id

  description = "%[1]s"
  tenant_id = netbox_tenant.test.id
  tunnel_id = 123
  tags = [netbox_tag.test.name]
}
resource "netbox_vpn_tunnel_termination" "test1" {
  tunnel_id = netbox_vpn_tunnel.test.id
  role = "hub"
  device_interface_id = netbox_device_interface.test1.id
}
resource "netbox_vpn_tunnel_termination" "test2" {
  tunnel_id = netbox_vpn_tunnel.test.id
  role = "spoke"
  device_interface_id = netbox_device_interface.test2.id
}`, testName)
}

func testAccNetboxTunnelTerminations() string {
	return `
data "netbox_vpn_tunnel_terminations" "test" {
}`
}

func TestAccNetboxTunnelTerminationsDataSource_basic(t *testing.T) {
	testName := testAccGetTestName("vpntun_terminations_ds_basic")
	setUp := testAccNetboxTunnelTerminationsSetUp(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_vpn_tunnel_termination.test1", "tunnel_id", "netbox_vpn_tunnel.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxTunnelTerminations(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vpn_tunnel_terminations.test", "terminations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vpn_tunnel_terminations.test", "terminations.0.id", "netbox_vpn_tunnel_termination.test1", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vpn_tunnel_terminations.test", "terminations.1.id", "netbox_vpn_tunnel_termination.test2", "id"),
				),
			},
		},
	})
}
