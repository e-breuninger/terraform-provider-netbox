package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxAvailablePrefixesDataSource_basic(t *testing.T) {
	testPrefix := "10.10.10.0/24"
	testSlug := "available_prefixes_ds_basic"

	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_prefix" "test_with_vrf" {
  prefix = "%[2]s"
  status = "active"
  vrf_id = netbox_vrf.test_vrf.id

}

resource "netbox_vrf" "test_vrf" {
  name = "%[1]s_test_vrf"
}

resource "netbox_available_prefix" "test_create_available_prefix" {
  parent_prefix_id = netbox_prefix.test_with_vrf.id
  prefix_length    = 27
  vrf_id           = netbox_vrf.test_vrf.id
  status           = "active"
}


data "netbox_available_prefix" "test_available_prefix" {
  prefix_id = netbox_available_prefix.test_create_available_prefix.parent_prefix_id
}

`, testName, testPrefix),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.#", "3"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.0.prefix", "10.10.10.32/27"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.1.prefix", "10.10.10.64/26"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.2.prefix", "10.10.10.128/25"),
					resource.TestCheckResourceAttrPair("data.netbox_available_prefix.test_available_prefix", "prefixes_available.0.vrf_id", "netbox_vrf.test_vrf", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_available_prefix.test_available_prefix", "prefix_id", "netbox_prefix.test_with_vrf", "id"),
				),
			},
		},
	})
}

func TestAccNetboxAvailablePrefixesDataSource_without_vrf(t *testing.T) {
	testPrefix := "10.10.10.0/24"
	testSlug := "available_prefixes_ds_without_vrf"

	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_prefix" "test_with_vrf" {
  prefix = "%[2]s"
  status = "active"
}

resource "netbox_available_prefix" "test_create_available_prefix" {
  parent_prefix_id = netbox_prefix.test_with_vrf.id
  prefix_length    = 27
  status           = "active"
}


data "netbox_available_prefix" "test_available_prefix" {
  prefix_id = netbox_available_prefix.test_create_available_prefix.parent_prefix_id
}

`, testName, testPrefix),

				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.#", "3"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.0.prefix", "10.10.10.32/27"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.1.prefix", "10.10.10.64/26"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.2.prefix", "10.10.10.128/25"),
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.2.vrf_id", "0"),
					resource.TestCheckResourceAttrPair("data.netbox_available_prefix.test_available_prefix", "prefix_id", "netbox_prefix.test_with_vrf", "id"),
				),
			},
		},
	})
}

func TestAccNetboxAvailablePrefixesDataSource_none_available(t *testing.T) {
	testPrefix := "10.10.10.0/24"
	testSlug := "available_prefixes_ds_none_available"

	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`

resource "netbox_prefix" "test_with_vrf" {
  prefix = "%[2]s"
  status = "active"
  vrf_id = netbox_vrf.test_vrf.id

}

resource "netbox_vrf" "test_vrf" {
  name = "%[1]s_test_vrf"
}

resource "netbox_available_prefix" "test_create_available_prefix" {
  count = 2
  parent_prefix_id = netbox_prefix.test_with_vrf.id
  prefix_length    = 25
  vrf_id           = netbox_vrf.test_vrf.id
  status           = "active"
}


data "netbox_available_prefix" "test_available_prefix" {
  prefix_id = netbox_available_prefix.test_create_available_prefix[1].parent_prefix_id
}

`, testName, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_available_prefix.test_available_prefix", "prefixes_available.#", "0"),
					resource.TestCheckResourceAttrPair("data.netbox_available_prefix.test_available_prefix", "prefix_id", "netbox_prefix.test_with_vrf", "id"),
				),
			},
		},
	})
}
