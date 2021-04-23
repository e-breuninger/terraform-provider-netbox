package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixDataSource_basic(t *testing.T) {

	testPrefix := "10.0.0.0/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[1]s"
  status = "active"
  is_pool = true
}
data "netbox_prefix" "test" {
  depends_on = [netbox_prefix.test]
  cidr = "%[1]s"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "id", "netbox_prefix.test", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
