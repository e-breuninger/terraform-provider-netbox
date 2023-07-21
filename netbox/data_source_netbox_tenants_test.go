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

func testAccNetboxTenantsDataSourceManyTenants(testName string) string {
	return fmt.Sprintf(`resource "netbox_tenant" "test" {
  count = 51
  name = "%s-${count.index}"
}
`, testName)
}

func TestAccNetboxTenantsDataSource_many(t *testing.T) {
	testSlug := "tnt_ds_many"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxTenantsDataSourceManyTenants(testName) + `data "netbox_tenants" "test" {
  depends_on = [netbox_tenant.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tenants.test", "tenants.#", "51"),
				),
			},
			{
				Config: testAccNetboxTenantsDataSourceManyTenants(testName) + `data "netbox_tenants" "test" {
  depends_on = [netbox_tenant.test]
  limit = 2
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tenants.test", "tenants.#", "2"),
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

func TestAccNetboxTenantsDataSource_tenantgroups(t *testing.T) {
	testSlug := "tnt_ds_tenant_group_filter"
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
