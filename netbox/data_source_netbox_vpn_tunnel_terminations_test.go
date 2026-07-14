package netbox

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

// testAccCheckIDInList verifies that the ID of resource resName appears somewhere
// in the "<listAttr>.N.id" attributes of data source dsName. This keeps list
// data-source tests parallel-safe: they assert their own objects are present
// rather than asserting a global count/ordering that concurrent tests change.
func testAccCheckIDInList(dsName, listAttr, resName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		ds, ok := s.RootModule().Resources[dsName]
		if !ok {
			return fmt.Errorf("data source not found: %s", dsName)
		}
		res, ok := s.RootModule().Resources[resName]
		if !ok {
			return fmt.Errorf("resource not found: %s", resName)
		}
		want := res.Primary.ID
		count, _ := strconv.Atoi(ds.Primary.Attributes[listAttr+".#"])
		for i := 0; i < count; i++ {
			if ds.Primary.Attributes[fmt.Sprintf("%s.%d.id", listAttr, i)] == want {
				return nil
			}
		}
		return fmt.Errorf("%s (id %s) not found in %s.%s", resName, want, dsName, listAttr)
	}
}

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
				// The data source has no filter, so it returns every termination
				// in Netbox. Assert our two are present (order/count-independent)
				// rather than a global count that concurrent tests would change.
				Check: resource.ComposeTestCheckFunc(
					testAccCheckIDInList("data.netbox_vpn_tunnel_terminations.test", "terminations", "netbox_vpn_tunnel_termination.test1"),
					testAccCheckIDInList("data.netbox_vpn_tunnel_terminations.test", "terminations", "netbox_vpn_tunnel_termination.test2"),
				),
			},
		},
	})
}
