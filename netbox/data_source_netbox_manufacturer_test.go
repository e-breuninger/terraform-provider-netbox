package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxManufacturerDataSource_basic(t *testing.T) {
	testSlug := "manufacturer_ds_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxManufacturerDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_manufacturers" "all_manufacturers" {
}


data "netbox_manufacturers" "by_name" {
 filter {
   name = "name"
   value  = netbox_manufacturer.manufacturer0.name
 }
}

data "netbox_manufacturers" "by_slug" {
 filter {
   name = "slug"
   value  = netbox_manufacturer.manufacturer1.slug
 }
}

data "netbox_manufacturers" "none" {
 filter {
   name = "slug"
   value  = "nonexisting"
 }
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestMatchResourceAttr("data.netbox_manufacturers.all_manufacturers", "manufacturers.#", regexp.MustCompile("[2-9]|\\d{2,}")), // assume there are at least 2 manufacturers, exact amount depends on pre-population
					resource.TestCheckResourceAttr("data.netbox_manufacturers.by_name", "manufacturers.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_manufacturers.by_name", "manufacturers.0.name", testName+"-0"),
					resource.TestCheckResourceAttr("data.netbox_manufacturers.by_name", "manufacturers.0.slug", testName+"-0-slug"),
					resource.TestCheckResourceAttrSet("data.netbox_manufacturers.by_name", "manufacturers.0.id"),
					resource.TestCheckResourceAttr("data.netbox_manufacturers.by_slug", "manufacturers.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_manufacturers.by_slug", "manufacturers.0.name", testName+"-1"),
					resource.TestCheckResourceAttr("data.netbox_manufacturers.by_slug", "manufacturers.0.slug", testName+"-1-slug"),
					resource.TestCheckResourceAttrSet("data.netbox_manufacturers.by_slug", "manufacturers.0.id"),
					resource.TestCheckResourceAttr("data.netbox_manufacturers.none", "manufacturers.#", "0"),
				),
			},
		},
	})
}

func testAccNetboxManufacturerDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_manufacturer" "manufacturer0" {
 name = "%[1]s-0"
 slug = "%[1]s-0-slug"
}

resource "netbox_manufacturer" "manufacturer1" {
 name = "%[1]s-1"
 slug = "%[1]s-1-slug"
}

`, testName)
}
