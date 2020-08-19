package netbox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"testing"
)

func TestAccNetboxPlatformDataSource_basic(t *testing.T) {

	testSlug := "pltf_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_platform" "test" {
  name = "%[1]s"
}
data "netbox_platform" "test" {
  depends_on = [netbox_platform.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_platform.test", "id", "netbox_platform.test", "id"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
