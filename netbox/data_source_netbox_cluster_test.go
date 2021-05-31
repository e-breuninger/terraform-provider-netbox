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
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
data "netbox_cluster" "test" {
  depends_on = [netbox_cluster.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_cluster.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_cluster.test", "id", "netbox_cluster.test", "id"),
				),
			},
		},
	})
}
