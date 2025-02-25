package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxConfigContextDataSource_basic(t *testing.T) {
	testSlug := "cfct_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name = "%[1]s"
  weight = 1000
  data = jsonencode({"testkey" = "testval"})
}
data "netbox_config_context" "test" {
  depends_on = [netbox_config_context.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_config_context.test", "id", "netbox_config_context.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_config_context.test", "weight", "netbox_config_context.test", "weight"),
					resource.TestCheckResourceAttrPair("data.netbox_config_context.test", "data", "netbox_config_context.test", "data"),
				),
			},
		},
	})
}
