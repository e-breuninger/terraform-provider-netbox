package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func getNetboxDataSourceRouteTargetConfig(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "acctest_ds_rts" {
	name = "%[1]s"
}

resource "netbox_route_targets" "acctest_ds_rts"{
	name = "%[1]s"
	tenant_id = netbox_tenant.acctest_ds_rts.id
}`, testName)
}

func TestAccNetboxRouteTargets_basic(t *testing.T) {
	testSlug := "rtsds"
	testName := testAccGetTestName(testSlug)

	testLongSlug := "rts_ds_errlongname"
	testLongName := testAccGetTestName(testLongSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: getNetboxDataSourceRouteTargetConfig(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox.route_targets.acctest_ds_rts", "id", "data.netbox_route_targets.acctest_ds_rts", "id"),
				),
			},
			{
				Config:      getNetboxDataSourceRouteTargetConfig(testLongName),
				ExpectError: regexp.MustCompile(fmt.Sprintf("expected length of name to be in the range (1 - 21), got %s", testLongName)),
			},
		},
	})
}
