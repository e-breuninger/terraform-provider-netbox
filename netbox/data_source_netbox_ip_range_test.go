package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpRangeDataSource_basic(t *testing.T) {
	testStartIP := "10.0.0.101/24"
	testEndIP := "10.0.0.150/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_range" "test" {
  start_address = "%[1]s"
  end_address = "%[2]s"
}
data "netbox_ip_range" "test" {
  depends_on = [netbox_ip_range.test]
  contains = "%[1]s"
}`, testStartIP, testEndIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.test", "id", "netbox_ip_range.test", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
