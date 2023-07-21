package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVrfDataSource_basic(t *testing.T) {
	testSlug := "tnt_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource"netbox_vrf" "test" {
  name = "%[1]s"
}

data "netbox_vrf" "test" {
  depends_on = [netbox_vrf.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vrf.test", "id", "netbox_vrf.test", "id"),
				),
			},
		},
	})
}
