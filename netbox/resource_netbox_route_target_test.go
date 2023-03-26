package netbox

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func getNetboxRouteTargetResource(rtName, tenantName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "rt_acctest_basic" {
	name = %[2]s
}
resource "netbox_ipam_route_target" "rt_acctest_basic" {
	name = "%[1]s",
	description = "rt for acctest",
	tenant_id = netbox_tenant.test.id
}`, rtName, tenantName)
}

func resourceNetboxRouteTarget_test(t *testing.T) {
	testSlug := "rt"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{

				Config: getNetboxRouteTargetResource(testName, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_route_target.rt_acctest_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_route_target.rt_acctest_basic", "description", "rt for acctest"),
					resource.TestCheckResourceAttrPair("netbox_route_target.rt_acc_test_basic", "tenant_id", "netbox_tenant.rt_acc_test_basic", "id"),
				),
				ResourceName:      "netbox_route_target.rt_acctest_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "rt_acctest_basic" {
	name = %[2]s
}
resource "netbox_ipam_route_target" "rt_acctest_basic" {
	name = "%[1]s",
	description = "change description",
	tenant_id = netbox_tenant.test.id
}`, testName, fmt.Sprintf("new%s", testName)),
				Check: resource.ComposeAggregateTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_route_target.rt_acctest_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_route_target.rt_acctest_basic", "description", "change description"),
					resource.TestCheckResourceAttrPair("netbox_route_target.rt_acc_test_basic", "tenant_id", "netbox_tenant.rt_acc_test_basic", "id"),
				),
				ResourceName:      "netbox_route_target.rt_acctest_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_ipam_route_target" "rt_acctest_basic" {
	name = "%[1]s",
	description = "change description",
	tenant_id = "10001"
}`, testName, fmt.Sprintf(testName)),
				ExpectError:       regexp.MustCompile("tenant_id does not exist.*"),
				ResourceName:      "netbox_route_target.rt_acctest_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_route_target", &resource.Sweeper{
		Name:         "netbox_route_target",
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
