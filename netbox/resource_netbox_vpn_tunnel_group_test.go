package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVpnTunnelGroup_basic(t *testing.T) {
	testSlug := "vpntnlgrp_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vpn_tunnel_group" "test" {
  name = "%[1]s"
  slug = "%[1]s"
  description = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_group.test", "slug", testName),
					resource.TestCheckResourceAttr("netbox_vpn_tunnel_group.test", "description", testName),
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

func init() {
	resource.AddTestSweepers("netbox_vpn_tunnel_group", &resource.Sweeper{
		Name:         "netbox_vpn_tunnel_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := vpn.NewVpnTunnelGroupsListParams()
			res, err := api.Vpn.VpnTunnelGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, clusterGroup := range res.GetPayload().Results {
				if strings.HasPrefix(*clusterGroup.Name, testPrefix) {
					deleteParams := vpn.NewVpnTunnelGroupsDeleteParams().WithID(clusterGroup.ID)
					_, err := api.Vpn.VpnTunnelGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a vpn_tunnel_group")
				}
			}
			return nil
		},
	})
}
