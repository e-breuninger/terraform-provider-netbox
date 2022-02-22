package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxTennantListDataSource_basic(t *testing.T) {

	testSlug := "test-tenant-list"
	testName := testAccGetTestName(testSlug)
	testResource := "data.netbox_tenant_list.test"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test_list_0" {
  name = "%[1]s_0"
}

resource "netbox_tenant" "test_list_1" {
  name = "%[1]s_1"
}

data "netbox_tenant_list" "test" {
  depends_on = [netbox_tenant.test_list_0, netbox_tenant.test_list_1]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResource, "tenants.0.name", fmt.Sprintf("%s_0", testName)),
					resource.TestCheckResourceAttr(testResource, "tenants.1.name", fmt.Sprintf("%s_1", testName)),
				),
			},
		},
	})
}
