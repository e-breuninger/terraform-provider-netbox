package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxClusterGroupDataSource_basic(t *testing.T) {
	testSlug := "clstrgrp_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_cluster_group" "test" {
  name = "%[1]s"
}

data "netbox_cluster_group" "test" {
  depends_on = [netbox_cluster_group.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_cluster_group.test", "cluster_group_id", "netbox_cluster_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cluster_group.test", "id", "netbox_cluster_group.test", "id"),
				),
			},
		},
	})
}
