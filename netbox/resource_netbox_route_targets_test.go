package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func getNetboxRouteTargetsResource(rtName, tenantName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "rt_acctest_basic" {
	name = %[2]s
}
resource "netbox_ipan_route_targets" "rt_acctest_basic" {
	name = "%[1]s",
	description = "rts for acctest",
	tenant_id = netbox_tenant.test.id
}`, rtName, tenantName)
}

func resourceNetboxRouteTargets_test(t *testing.T) {
	testSlug := "rts"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{

				Config: getNetboxRouteTargetsResource(testName, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_route_targets.rt_acctest_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_route_targets.rt_acctest_basic", "description", "rts for acctest"),
				),
				ResourceName:      "netbox_route_targets.rt_acctest_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_route_targets", &resource.Sweeper{
		Name:         "netbox_route_targets",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamRouteTargetsListParams()
			res, err := api.Ipam.IpamRouteTargetsList(params, nil)
			if err != nil {
				return err
			}
			for _, role := range res.GetPayload().Results {
				if strings.HasPrefix(*role.Name, testPrefix) {
					deleteParams := ipam.NewIpamRouteTargetsDeleteParams().WithID(role.ID)
					_, err := api.Ipam.IpamRouteTargetsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a rir")
				}
			}
			return nil
		},
	})
}
