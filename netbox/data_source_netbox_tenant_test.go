package netbox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccNetboxTenantDataSource_basic(t *testing.T) {

	testSlug := "tnt_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

data "netbox_tenant" "test" {
  depends_on = [netbox_tenant.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_tenant.test", "id", "netbox_tenant.test", "id"),
				),
			},
		},
	})
}
