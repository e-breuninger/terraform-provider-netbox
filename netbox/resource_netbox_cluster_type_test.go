package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxClusterType_basic(t *testing.T) {
	testSlug := "clstr_type_data_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_cluster_type.test", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_cluster_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxClusterType_defaultSlug(t *testing.T) {
	testSlug := "clstr_type_data_default_slug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_cluster_type.test", "slug", testName),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_cluster_type", &resource.Sweeper{
		Name:         "netbox_cluster_type",
		Dependencies: []string{"netbox_cluster"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := virtualization.NewVirtualizationClusterTypesListParams()
			res, err := api.Virtualization.VirtualizationClusterTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, clusterType := range res.GetPayload().Results {
				if strings.HasPrefix(*clusterType.Name, testPrefix) {
					deleteParams := virtualization.NewVirtualizationClusterTypesDeleteParams().WithID(clusterType.ID)
					_, err := api.Virtualization.VirtualizationClusterTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a cluster type")
				}
			}
			return nil
		},
	})
}
