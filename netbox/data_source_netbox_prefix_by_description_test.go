package netbox

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixByDescDataSource_basic(t *testing.T) {

	testPrefix := "10.0.0.0/24"
	testDesc := fmt.Sprintf("test-prefix-%d", rand.Intn(100))
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[1]s"
  status = "active"
  is_pool = true
  description = "%[2]s"
}
data "netbox_prefix_by_description" "test" {
  depends_on = [netbox_prefix.test]
  description = "%[2]s"
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "id", "netbox_prefix.test", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
