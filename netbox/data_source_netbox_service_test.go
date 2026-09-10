package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxServiceDataSource_basic(t *testing.T) {
	testSlug := "service_ds_basic"
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
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_service" "test" {
  name               = "%[1]s"
  virtual_machine_id = netbox_virtual_machine.test.id
  protocol           = "tcp"
  ports              = [80]
}

data "netbox_service" "test" {
  depends_on         = [netbox_service.test]
  name               = "%[1]s"
  virtual_machine_id = netbox_virtual_machine.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_service.test", "id", "netbox_service.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_service.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_service.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("data.netbox_service.test", "ports.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_service.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
				),
			},
		},
	})
}
