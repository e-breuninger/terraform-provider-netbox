package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func getNetboxDataSourceRouteTargetConfig(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "acctest_ds_rt" {
	name = "%[1]s"
}

resource "netbox_route_target" "acctest_ds_rt" {
	name = "%[1]s"
	tenant_id = netbox_tenant.acctest_ds_rt.id
}

data "netbox_route_target" "acctest_ds_rt" {
	name = "%[1]s"
	depends_on = [netbox_route_target.acctest_ds_rt]
}`, testName)
}

func TestAccNetboxRouteTarget_basic(t *testing.T) {
	testSlug := "rtds"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getNetboxDataSourceRouteTargetConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_route_target.acctest_ds_rt", "id", "netbox_route_target.acctest_ds_rt", "id"),
				),
			},
		},
	})
}
