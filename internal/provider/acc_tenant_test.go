//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/tenancy"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccNetboxTenantTagDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}

resource "netbox_tag" "test_b" {
  name = "%[1]sb"
}
`, testName)
}

func TestAccNetboxTenant_basic(t *testing.T) {
	testSlug := "tenant_basic"
	testName := testAccGetTestName(testSlug)
	testDescription := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
  slug = "%[2]s"
  description = "%[3]s"
  comments = "Created by acceptance test."
  group_id = netbox_tenant_group.test.id
}`, testName, randomSlug, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_tenant.test", "description", testDescription),
					resource.TestCheckResourceAttrPair("netbox_tenant.test", "group_id", "netbox_tenant_group.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_tenant.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: description, comments and the group clear.
				Config: fmt.Sprintf(`
resource "netbox_tenant_group" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_tenant.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_tenant.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_tenant.test", "group_id"),
				),
			},
		},
	})
}

func TestAccNetboxTenant_defaultSlug(t *testing.T) {
	testSlug := "tenant_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

// Tags are referenced by slug and read back as a set.
func TestAccNetboxTenant_tags(t *testing.T) {
	testSlug := "tenant_tags"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxTenantTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_tenant" "test_tags" {
  name = "%[1]s"
  tags = [netbox_tag.test_a.slug]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test_tags", "name", testName),
					resource.TestCheckResourceAttr("netbox_tenant.test_tags", "tags.#", "1"),
					resource.TestCheckTypeSetElemAttr("netbox_tenant.test_tags", "tags.*", getSlug(testName+"a")),
				),
			},
			{
				Config: testAccNetboxTenantTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_tenant" "test_tags" {
  name = "%[1]s"
  tags = [netbox_tag.test_a.slug, netbox_tag.test_b.slug]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tenant.test_tags", "tags.#", "2"),
					resource.TestCheckTypeSetElemAttr("netbox_tenant.test_tags", "tags.*", getSlug(testName+"a")),
					resource.TestCheckTypeSetElemAttr("netbox_tenant.test_tags", "tags.*", getSlug(testName+"b")),
				),
			},
			{
				Config: testAccNetboxTenantTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_tenant" "test_tags" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_tenant.test_tags", "tags"),
				),
			},
		},
	})
}

func TestAccNetboxTenantDataSource_basic(t *testing.T) {
	testSlug := "tnt_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
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
  slug = "%[2]s"
}

data "netbox_tenant" "by_both" {
  depends_on = [netbox_tenant.test]
  name = "%[1]s"
  slug = "%[2]s"
}
`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_tenant.by_name", "id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant.by_slug", "id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_tenant.by_both", "id", "netbox_tenant.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxTenantsDataSource_filter(t *testing.T) {
	testSlug := "tnts_ds_filter"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
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
  filters = [
    { name = "name", value = "%[1]s_0" },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_tenants.test", "tenants.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_tenants.test", "tenants.0.name", "netbox_tenant.test_list_0", "name"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_tenant",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Tenancy.TenancyTenantsList(tenancy.NewTenancyTenantsListParams(), nil)
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
			_, err := client.Tenancy.TenancyTenantsDestroy(tenancy.NewTenancyTenantsDestroyParams().WithID(id), nil)
			return err
		})
}
