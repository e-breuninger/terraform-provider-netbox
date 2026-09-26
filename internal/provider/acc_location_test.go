//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxLocation_basic(t *testing.T) {
	testSlug := "location_basic"
	testName := testAccGetTestName(testSlug)
	testNameSub := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	randomSlugSub := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  status      = "active"
  description = "my-description"
  facility    = "Building B"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
}

resource "netbox_location" "test-sub" {
  name        = "%[3]s"
  slug        = "%[4]s"
  description = "my-description"
  parent_id   = netbox_location.test.id
  site_id     = netbox_site.test.id
}`, testName, randomSlug, testNameSub, randomSlugSub),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_location.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_location.test", "description", "my-description"),
					resource.TestCheckResourceAttr("netbox_location.test", "facility", "Building B"),
					resource.TestCheckResourceAttr("netbox_location.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_location.test", "depth", "0"),
					resource.TestCheckResourceAttr("netbox_location.test-sub", "depth", "1"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_location.test", "id", "netbox_location.test-sub", "parent_id"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test_2" {
  name = "%[1]s_b"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name = "%[1]s"
  slug = "%[2]s"
  site_id = netbox_site.test_2.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_location.test", "slug", randomSlug),
					resource.TestCheckResourceAttrPair("netbox_location.test", "site_id", "netbox_site.test_2", "id"),
					resource.TestCheckNoResourceAttr("netbox_location.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_location.test", "facility"),
					resource.TestCheckNoResourceAttr("netbox_location.test", "tenant_id"),
				),
			},
			{
				ResourceName:      "netbox_location.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxLocation_updateParent(t *testing.T) {
	testSlug := "loc_upd_parent"
	testName := testAccGetTestName(testSlug)
	testNameSub := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	randomSlugSub := testAccGetTestName(testSlug)
	withParent := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "my-description"
  site_id     = netbox_site.test.id
}

resource "netbox_location" "test_sub" {
  name        = "%[3]s"
  slug        = "%[4]s"
  description = "my-description"
  parent_id   = netbox_location.test.id
  site_id     = netbox_site.test.id
}`, testName, randomSlug, testNameSub, randomSlugSub)
	withoutParent := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "my-description"
  site_id     = netbox_site.test.id
}

resource "netbox_location" "test_sub" {
  name        = "%[3]s"
  slug        = "%[4]s"
  description = "my-description"
  site_id     = netbox_site.test.id
}`, testName, randomSlug, testNameSub, randomSlugSub)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: withParent,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_location.test", "id", "netbox_location.test_sub", "parent_id"),
				),
			},
			{
				Config: withoutParent,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_location.test", "name", testName),
					resource.TestCheckNoResourceAttr("netbox_location.test_sub", "parent_id"),
				),
			},
		},
	})
}

func TestAccNetboxLocationDataSource_basic(t *testing.T) {
	testSlug := "location_ds_basic"
	testName := testAccGetTestName(testSlug)
	testNameSub := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  description = "my-description"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
}

resource "netbox_location" "test_sub" {
  name        = "%[2]s"
  description = "my-description"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
  parent_id   = netbox_location.test.id
}

data "netbox_location" "by_name" {
  name = netbox_location.test.name
}

data "netbox_location" "by_name_and_site" {
  name    = netbox_location.test.name
  site_id = netbox_site.test.id
}

data "netbox_location" "sub_by_name" {
  name = netbox_location.test_sub.name
}

data "netbox_location" "by_id" {
  id = netbox_location.test.id
}

data "netbox_location" "by_slug" {
  slug = netbox_location.test.slug
}`, testName, testNameSub),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_id", "id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_slug", "id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_location.by_name", "name", testName),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "description", "netbox_location.test", "description"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "site_id", "netbox_location.test", "site_id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name_and_site", "site_id", "netbox_location.test", "site_id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.by_name", "tenant_id", "netbox_location.test", "tenant_id"),
					resource.TestCheckResourceAttrPair("data.netbox_location.sub_by_name", "parent_id", "netbox_location.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxLocationsDataSource_basic(t *testing.T) {
	testSlug := "locations_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_location" "test" {
  name        = "%[1]s"
  description = "my-description"
  facility    = "Building B"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
  tags        = [netbox_tag.test.slug]
}

data "netbox_locations" "by_name" {
  filters = [
    { name = "name", value = netbox_location.test.name },
  ]
}

data "netbox_locations" "no_match" {
  filters = [
    { name = "name", value = "non-existent" },
  ]
}

data "netbox_locations" "by_site_slug" {
  filters = [
    { name = "site", value = netbox_site.test.slug },
  ]
  depends_on = [netbox_location.test]
}

data "netbox_locations" "by_site_id" {
  filters = [
    { name = "site_id", value = netbox_site.test.id },
  ]
  depends_on = [netbox_location.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_locations.by_name", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_name", "locations.0.name", "netbox_location.test", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_name", "locations.0.site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_name", "locations.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckNoResourceAttr("data.netbox_locations.by_name", "locations.0.parent_id"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_name", "locations.0.description", "my-description"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_name", "locations.0.facility", "Building B"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_name", "locations.0.tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_locations.no_match", "locations.#", "0"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_site_slug", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_site_slug", "locations.0.name", "netbox_location.test", "name"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_site_id", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_site_id", "locations.0.name", "netbox_location.test", "name"),
				),
			},
		},
	})
}

func TestAccNetboxLocationsDataSource_sublocations(t *testing.T) {
	testSlug := "sublocations_ds"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_location" "parent" {
  name    = "%[1]s_p"
  site_id = netbox_site.test.id
}

resource "netbox_location" "test1" {
  name      = "%[1]s_1"
  parent_id = netbox_location.parent.id
  site_id   = netbox_site.test.id
}

resource "netbox_location" "test2" {
  name      = "%[1]s_2"
  parent_id = netbox_location.parent.id
  site_id   = netbox_site.test.id
}

data "netbox_locations" "by_parent" {
  filters = [
    { name = "parent_id", value = netbox_location.parent.id },
  ]
  depends_on = [netbox_location.test1, netbox_location.test2]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_locations.by_parent", "locations.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_parent", "locations.0.parent_id", "netbox_location.parent", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_parent", "locations.1.parent_id", "netbox_location.parent", "id"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_location",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimLocationsList(dcim.NewDcimLocationsListParams(), nil)
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
			_, err := client.Dcim.DcimLocationsDestroy(dcim.NewDcimLocationsDestroyParams().WithID(id), nil)
			return err
		})
}
