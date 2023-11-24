package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxLocationsDataSource_basic(t *testing.T) {
	testSlug := "location_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
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
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
	tags        = [netbox_tag.test.slug]
}

data "netbox_locations" "by_name" {
	filter {
		name  = "name"
		value = netbox_location.test.name
	}
}

data "netbox_locations" "no_match" {
	filter {
		name  = "name"
		value = "non-existent"
	}
}

data "netbox_locations" "by_site_slug" {
	filter {
		name  = "site"
		value = netbox_site.test.slug
	}
	depends_on = [netbox_location.test]
}

data "netbox_locations" "by_site_id" {
	filter {
		name  = "site_id"
		value = netbox_site.test.id
	}
	depends_on = [netbox_location.test]
}

data "netbox_locations" "by_tenant_slug" {
	filter {
		name  = "tenant"
		value = netbox_tenant.test.slug
	}
	depends_on = [netbox_location.test]
}

data "netbox_locations" "by_tenant_id" {
	filter {
		name  = "tenant_id"
		value = netbox_tenant.test.id
	}
	depends_on = [netbox_location.test]
}

data "netbox_locations" "by_tags" {
	tags 			 = [netbox_tag.test.slug]
	depends_on = [netbox_location.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_locations.by_name", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_name", "locations.0.name", "netbox_location.test", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_name", "locations.0.site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_name", "locations.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_name", "locations.0.description", "my-description"),
					resource.TestCheckResourceAttr("data.netbox_locations.no_match", "locations.#", "0"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_site_slug", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_site_slug", "locations.0.name", "netbox_location.test", "name"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_site_id", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_site_id", "locations.0.name", "netbox_location.test", "name"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_tenant_slug", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_tenant_slug", "locations.0.name", "netbox_location.test", "name"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_tenant_id", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_tenant_id", "locations.0.name", "netbox_location.test", "name"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_tags", "locations.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_locations.by_tags", "locations.0.name", "netbox_location.test", "name"),
				),
			},
		},
	})
}

func TestAccNetboxLocationsDataSource_multiple(t *testing.T) {
	testSlug := "location_ds_multiple"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test1" {
	name = "%[1]s_1"
}

resource "netbox_tag" "test2" {
	name = "%[1]s_2"
}

resource "netbox_location" "test1" {
  name        = "%[1]s_1"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
	tags        = [netbox_tag.test1.slug]
}

resource "netbox_location" "test2" {
  name        = "%[1]s_2"
  site_id     = netbox_site.test.id
  tenant_id   = netbox_tenant.test.id
	tags        = [netbox_tag.test2.slug]
}

data "netbox_locations" "by_site" {
	filter {
		name  = "site"
		value = netbox_site.test.name
	}
	depends_on = [netbox_location.test1, netbox_location.test2]
}

data "netbox_locations" "by_tag" {
	tags = [netbox_tag.test1.slug]
	depends_on = [netbox_location.test1, netbox_location.test2]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_locations.by_site", "locations.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_locations.by_tag", "locations.#", "1"),
				),
			},
		},
	})
}
