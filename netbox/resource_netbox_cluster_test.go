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

func TestAccNetboxCluster_basic(t *testing.T) {
	testSlug := "clstr_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_group" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  cluster_group_id = netbox_cluster_group.test.id
  comments = "%[1]scomments"
  description = "%[1]sdescription"
  site_id = netbox_site.test.id
  tags = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "cluster_type_id", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "cluster_group_id", "netbox_cluster_group.test", "id"),
					resource.TestCheckResourceAttr("netbox_cluster.test", "comments", testName+"comments"),
					resource.TestCheckResourceAttr("netbox_cluster.test", "description", testName+"description"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_cluster.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_cluster.test", "tags.0", testName),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test_updatetag" {
  name = "%[1]s-a"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_group" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name   = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  cluster_group_id = netbox_cluster_group.test.id
  tenant_id = netbox_tenant.test.id
  tags = [netbox_tag.test.name, netbox_tag.test_updatetag.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "cluster_type_id", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "cluster_group_id", "netbox_cluster_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_cluster.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_cluster.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("netbox_cluster.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_cluster.test", "tags.1", fmt.Sprintf("%[1]s-a", testName)),
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
		Dependencies: []string{"netbox_virtual_machine", "netbox_site"},
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
