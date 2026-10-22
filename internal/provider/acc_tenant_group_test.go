//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxTenantGroup_basic(t *testing.T) {
	testSlug := "tenantgroup_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%s"
  slug = "%s"
  description = "%[1]s"
  comments    = "Acceptance test comments."
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "slug", randomSlug),
					resource.TestCheckNoResourceAttr("netbox_tenant_group.test", "parent_id"),
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "comments", "Acceptance test comments."),
				),
			},
			{
				// Shrink: the description clears.
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_tenant_group.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_tenant_group.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_tenant_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxTenantGroup_defaultSlug(t *testing.T) {
	testSlug := "tenantgroup_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant_group.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func TestAccNetboxTenantGroup_parent(t *testing.T) {
	testSlug := "tenantgroup_parent"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%[1]s"
}
resource "netbox_tenant_group" "test-child" {
  name      = "%[1]s-child"
  parent_id = netbox_tenant_group.test.id
}`, testName),
				Check: resource.TestCheckResourceAttrPair("netbox_tenant_group.test-child", "parent_id", "netbox_tenant_group.test", "id"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%[1]s"
}
resource "netbox_tenant_group" "test-child" {
  name = "%[1]s-child"
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_tenant_group.test-child", "parent_id"),
			},
		},
	})
}

func TestAccNetboxTenantGroupDataSource_basic(t *testing.T) {
	testSlug := "tenantgroup_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%[1]s"
}
resource "netbox_tenant_group" "test-child" {
  name      = "%[1]s-child"
  parent_id = netbox_tenant_group.test.id
}
data "netbox_tenant_group" "test" {
  depends_on = [netbox_tenant_group.test]
  name = "%[1]s"
}
data "netbox_tenant_group" "test-child" {
  depends_on = [netbox_tenant_group.test-child]
  slug = "%[2]s"
}
data "netbox_tenant_groups" "all" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on = [netbox_tenant_group.test, netbox_tenant_group.test-child]
}`, testName, getSlug(testName+"-child")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_tenant_group.test", "id", "netbox_tenant_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant_group.test", "slug", "netbox_tenant_group.test", "slug"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant_group.test-child", "id", "netbox_tenant_group.test-child", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant_group.test-child", "parent_id", "netbox_tenant_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_tenant_groups.all", "tenant_groups.#", "2"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_tenant_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Tenancy.TenancyTenantGroupsList(tenancy.NewTenancyTenantGroupsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Tenancy.TenancyTenantGroupsDestroy(tenancy.NewTenancyTenantGroupsDestroyParams().WithID(id), nil)
			return err
		})
}
