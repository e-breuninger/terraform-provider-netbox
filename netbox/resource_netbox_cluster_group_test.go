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

func TestAccNetboxClusterGroup_basic(t *testing.T) {

	testSlug := "clstrgrp_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cluster_group" "test" {
  name = "%[1]s"
  slug = "%[1]s"
  description = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_cluster_group.test", "slug", testName),
					resource.TestCheckResourceAttr("netbox_cluster_group.test", "description", testName),
				),
			},
			{
				ResourceName:      "netbox_cluster_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxClusterGroup_defaultSlug(t *testing.T) {

	testSlug := "clstrgrp_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cluster_group" "test" {
  name = "%[1]s"
  description = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cluster_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_cluster_group.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttr("netbox_cluster_group.test", "description", testName),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_cluster_group", &resource.Sweeper{
		Name:         "netbox_cluster_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := virtualization.NewVirtualizationClusterGroupsListParams()
			res, err := api.Virtualization.VirtualizationClusterGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, cluster_group := range res.GetPayload().Results {
				if strings.HasPrefix(*cluster_group.Name, testPrefix) {
					deleteParams := virtualization.NewVirtualizationClusterGroupsDeleteParams().WithID(cluster_group.ID)
					_, err := api.Virtualization.VirtualizationClusterGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a cluster_group")
				}
			}
			return nil
		},
	})
}
