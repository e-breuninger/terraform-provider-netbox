package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxAggregateDataSource_basic(t *testing.T) {
	testSlug := "aggregate_ds_basic"
	testName := testAccGetTestName(testSlug)
	testPrefix := "10.0.0.0/8"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}

resource "netbox_aggregate" "test" {
  prefix = "%[2]s"
  rir_id = netbox_rir.test.id
}

data "netbox_aggregate" "test" {
  depends_on = [netbox_aggregate.test]
  prefix     = "%[2]s"
  rir_id     = netbox_rir.test.id
}`, testName, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_aggregate.test", "id", "netbox_aggregate.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_aggregate.test", "prefix", testPrefix),
					resource.TestCheckResourceAttrPair("data.netbox_aggregate.test", "rir_id", "netbox_rir.test", "id"),
				),
			},
		},
	})
}
