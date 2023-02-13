package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func getNetboxDataSourceRouteTargetConfig(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "acctest_ds_rts" {
	name = "%[1]s"
}

resource "netbox_route_targets" "acctest_ds_rts" {
	name = "%[1]s"
	tenant_id = netbox_tenant.acctest_ds_rts.id
}

data "netbox_route_targets" "acctest_ds_rts" {
	name = "%[1]s"
	depends_on = [netbox_route_targets.acctest_ds_rts]
}`, testName)
}

func TestAccNetboxRouteTargets_basic(t *testing.T) {
	testSlug := "rtsds"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getNetboxDataSourceRouteTargetConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_route_targets.acctest_ds_rts", "id", "netbox_route_targets.acctest_ds_rts", "id"),
				),
			},
		},
	})
}
