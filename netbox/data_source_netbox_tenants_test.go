package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxTenantsDataSource_basic(t *testing.T) {

	testSlug := "tnt_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
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
data "netbox_tenants" "test" {
  depends_on = [netbox_tenant.test_list_0, netbox_tenant.test_list_1]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.0.name", "netbox_tenant.test_list_0", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.1.name", "netbox_tenant.test_list_1", "name"),
				),
			},
		},
	})
}

func TestAccNetboxTenantsDataSource_filter(t *testing.T) {

	testSlug := "tnt_ds_filter"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
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
data "netbox_tenants" "test" {
  depends_on = [netbox_tenant.test_list_0, netbox_tenant.test_list_1]

	filter {
		name = "name"
		value = "%[1]s_0"
	}
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tenants.test", "tenants.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.0.name", "netbox_tenant.test_list_0", "name"),
				),
			},
		},
	})
}
