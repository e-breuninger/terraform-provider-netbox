package netbox

import (
	"fmt"
	"testing"
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVpnTunnelGroupSetUp(testName string) string {
	return fmt.Sprintf(`
resource "netbox_vpn_tunnel_group" "test" {
  name = "%[1]s"
  description = "%[1]s"
}`, testName)
}


const testAccNetboxVpnTunnelGroupNoResult = `
data "netbox_vpn_tunnel_group" "test" {
  name = "iDontExist"
}`

func testAccNetboxVpnTunnelGroupByName(testName string) string {
	return fmt.Sprintf(`
data "netbox_vpn_tunnel_group" "test" {
  name = "%[1]s"
}`, testName)
}

func TestAccNetboxVpnTunnelGroupDataSource_basic(t *testing.T) {
	testSlug := "vpntun_ds_basic"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxVpnTunnelGroupSetUp(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_group.test", "name", testName),
				),
			},
			{
				Config: setUp + testAccNetboxVpnTunnelGroupByName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vpn_tunnel_group.test", "id", "netbox_vpn_tunnel_group.test", "id"),
				),
			},
			{
				Config:      setUp + testAccNetboxVpnTunnelGroupNoResult,
				ExpectError: regexp.MustCompile("no tunnel group found matching filter"),
			},
		},
	})
}
