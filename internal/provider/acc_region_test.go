//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxRegion_basic(t *testing.T) {
	testSlug := "region_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%s"
  slug = "%s"
  description = "%[1]s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_region.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_region.test", "slug", randomSlug),
					resource.TestCheckNoResourceAttr("netbox_region.test", "parent_region_id"),
					resource.TestCheckResourceAttr("netbox_region.test", "description", testName),
				),
			},
			{
				ResourceName:      "netbox_region.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxRegion_defaultSlug(t *testing.T) {
	testSlug := "region_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_region.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_region.test", "slug", getSlug(testName)),
					resource.TestCheckNoResourceAttr("netbox_region.test", "description"),
				),
			},
		},
	})
}

func TestAccNetboxRegion_parent(t *testing.T) {
	testSlug := "region_parent"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test-child" {
  name      = "%[1]s-child"
  parent_region_id = netbox_region.test.id
}`, testName),
				Check: resource.TestCheckResourceAttrPair("netbox_region.test-child", "parent_region_id", "netbox_region.test", "id"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test-child" {
  name = "%[1]s-child"
}`, testName),
				Check: resource.TestCheckNoResourceAttr("netbox_region.test-child", "parent_region_id"),
			},
		},
	})
}

func TestAccNetboxRegionDataSource_basic(t *testing.T) {
	testSlug := "region_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test-child" {
  name      = "%[1]s-child"
  parent_region_id = netbox_region.test.id
}
data "netbox_region" "test" {
  depends_on = [netbox_region.test]
  name = "%[1]s"
}
data "netbox_region" "test-child" {
  depends_on = [netbox_region.test-child]
  slug = "%[2]s"
}
data "netbox_regions" "all" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on = [netbox_region.test, netbox_region.test-child]
}`, testName, getSlug(testName+"-child")),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_region.test", "id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_region.test", "slug", "netbox_region.test", "slug"),
					resource.TestCheckResourceAttrPair("data.netbox_region.test-child", "id", "netbox_region.test-child", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_region.test-child", "parent_region_id", "netbox_region.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_regions.all", "regions.#", "2"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_region",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimRegionsList(dcim.NewDcimRegionsListParams(), nil)
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
			_, err := client.Dcim.DcimRegionsDestroy(dcim.NewDcimRegionsDestroyParams().WithID(id), nil)
			return err
		})
}
