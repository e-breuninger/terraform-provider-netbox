//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxSiteGroup_basic(t *testing.T) {
	testSlug := "sitegroup_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "test" {
  name = "%s"
  slug = "%s"
  description = "%[1]s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site_group.test", "slug", randomSlug),
					resource.TestCheckNoResourceAttr("netbox_site_group.test", "parent_id"),
					resource.TestCheckResourceAttr("netbox_site_group.test", "description", testName),
				),
			},
			{
				ResourceName:      "netbox_site_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxSiteGroup_defaultSlug(t *testing.T) {
	testSlug := "sitegroup_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site_group.test", "slug", getSlug(testName)),
					resource.TestCheckNoResourceAttr("netbox_site_group.test", "description"),
				),
			},
		},
	})
}

func TestAccNetboxSiteGroup_parent(t *testing.T) {
	testSlug := "sitegroup_parent"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test-child" {
  name      = "%[1]s-child"
  parent_id = netbox_site_group.test.id
}`, testName),
				Check: resource.TestCheckResourceAttrPair("netbox_site_group.test-child", "parent_id", "netbox_site_group.test", "id"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test-child" {
  name = "%[1]s-child"
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_site_group.test-child", "parent_id"),
			},
		},
	})
}

func TestAccNetboxSiteGroupDataSource_basic(t *testing.T) {
	testSlug := "sitegroup_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test-child" {
  name      = "%[1]s-child"
  parent_id = netbox_site_group.test.id
}
data "netbox_site_group" "test" {
  depends_on = [netbox_site_group.test]
  name = "%[1]s"
}
data "netbox_site_group" "test-child" {
  depends_on = [netbox_site_group.test-child]
  slug = "%[2]s"
}
data "netbox_site_groups" "all" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on = [netbox_site_group.test, netbox_site_group.test-child]
}`, testName, getSlug(testName+"-child")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_site_group.test", "id", "netbox_site_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_site_group.test", "slug", "netbox_site_group.test", "slug"),
					resource.TestCheckResourceAttrPair("data.netbox_site_group.test-child", "id", "netbox_site_group.test-child", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_site_group.test-child", "parent_id", "netbox_site_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_site_groups.all", "site_groups.#", "2"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_site_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimSiteGroupsList(dcim.NewDcimSiteGroupsListParams(), nil)
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
			_, err := client.Dcim.DcimSiteGroupsDestroy(dcim.NewDcimSiteGroupsDestroyParams().WithID(id), nil)
			return err
		})
}
