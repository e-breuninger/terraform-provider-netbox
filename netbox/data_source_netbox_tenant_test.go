package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
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

data "netbox_tenant" "by_name" {
  depends_on = [netbox_tenant.test]
  name = "%[1]s"
}

data "netbox_tenant" "by_slug" {
  depends_on = [netbox_tenant.test]
  slug = "%[1]s"
}

data "netbox_tenant" "by_description" {
  depends_on = [netbox_tenant.test]
  name = "%[1]s"
  description = "%[1]s"
}

data "netbox_tenant" "by_both" {
  depends_on = [netbox_tenant.test]
  name = "%[1]s"
  slug = "%[1]s"
  description = "%[1]s"
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_tenant.by_name", "id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant.by_slug", "id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant.by_description", "id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant.by_both", "id", "netbox_tenant.test", "id"),
				),
			},
		},
	})
}
