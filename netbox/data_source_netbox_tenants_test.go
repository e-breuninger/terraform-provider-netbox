package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxTenantsDataSource_basic(t *testing.T) {
	testSlug := "tnts_ds_basic"
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
				//                              This snippet sometimes returns things from other tests, even if resource.Test is used instead of resource.ParallelTest
				//                              This happens especially in CI testing (where test execution is presumably slow)
				//                              The check functions are now removed so this does no longer happen
				//				Check: resource.ComposeTestCheckFunc(
				//					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.0.name", "netbox_tenant.test_list_0", "name"),
				//					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.1.name", "netbox_tenant.test_list_1", "name"),
				//				),
			},
		},
	})
}

func TestAccNetboxTenantsDataSource_filter(t *testing.T) {
	testSlug := "tnts_ds_filter"
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

func TestAccNetboxTenantsDataSource_tenantgroups(t *testing.T) {
	testSlug := "tnts_ds_tenant_group_filter"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "group_0" {
  name = "group_%[1]s_1"
}

resource "netbox_tenant" "tenant_0" {
  name = "tenant_%[1]s_0"
  group_id = netbox_tenant_group.group_0.id
}

data "netbox_tenants" "test" {
  depends_on = [netbox_tenant.tenant_0, netbox_tenant_group.group_0]

  filter {
    name = "name"
    value = "tenant_%[1]s_0"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tenants.test", "tenants.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.0.tenant_group.0.name", "netbox_tenant_group.group_0", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.0.tenant_group.0.slug", "netbox_tenant_group.group_0", "slug"),
				),
			},
		},
	})
}
