package netbox

import (
	"fmt"
	"testing"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVpnTunnelSetUp(testName string) string {
	return fmt.Sprintf(`
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
}`, testName)
}


const testAccNetboxVpnTunnelNoResult = `
data "vpn_tunnel" "test" {
  name = "iDontExist"
}`

func testAccNetboxVpnTunnelByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_vpn_tunnel" "test" {
  name = "%[1]s"
}`, testName)
}

func TestAccNetboxVpnTunnelDataSource_basic(t *testing.T) {
	testSlug := "vpntun_ds_basic"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxVpnTunnelSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel.test", "name", testName),
				),
			},
			{
				Config: setUp + testAccNetboxVpnTunnelByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vpn_tunnel.test", "id", "netbox_vpn_tunnel.test", "id"),
				),
			},
			{
				Config:      setUp + testAccNetboxVpnTunnelNoResult,
				ExpectError: regexp.MustCompile("no vpn tunnel found matching filter"),
			},
		},
	})
}
