package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxClusterDataSource_basic(t *testing.T) {

	testSlug := "clstr_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
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
  site_id = netbox_site.test.id
  tags = [netbox_tag.test.name]
}

data "netbox_cluster" "by_name" {
  name = netbox_cluster.test.name
}

data "netbox_cluster" "by_site_id" {
  site_id = netbox_cluster.test.site_id
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_cluster.by_name", "id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cluster.by_site_id", "id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_cluster.by_name", "name", testName),
					resource.TestCheckResourceAttrPair("data.netbox_cluster.by_name", "cluster_type_id", "netbox_cluster_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cluster.by_name", "cluster_group_id", "netbox_cluster_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cluster.by_name", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_cluster.by_name", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_cluster.by_name", "tags.0", testName),
				),
			},
		},
	})
}
