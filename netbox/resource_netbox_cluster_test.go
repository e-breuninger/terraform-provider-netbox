package netbox

import (
	"fmt"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	"strings"
	"testing"
)

func TestAccNetboxCluster_basic(t *testing.T) {

	testSlug := "clstr_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "name", testName),
				),
			},
			{
				ResourceName:      "netbox_cluster.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_cluster", &resource.Sweeper{
		Name:         "netbox_cluster",
		Dependencies: []string{"netbox_virtual_machine"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := virtualization.NewVirtualizationClustersListParams()
			res, err := api.Virtualization.VirtualizationClustersList(params, nil)
			if err != nil {
				return err
			}
			for _, cluster := range res.GetPayload().Results {
				if strings.HasPrefix(*cluster.Name, testPrefix) {
					deleteParams := virtualization.NewVirtualizationClustersDeleteParams().WithID(cluster.ID)
					_, err := api.Virtualization.VirtualizationClustersDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a cluster")
				}
			}
			return nil
		},
	})
}
