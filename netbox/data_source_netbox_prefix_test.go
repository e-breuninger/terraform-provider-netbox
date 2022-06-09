package netbox

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxPrefixDataSource_basic(t *testing.T) {

	testPrefix := "10.0.0.0/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[1]s"
  status = "active"
  is_pool = true
}
data "netbox_prefix" "test" {
  depends_on = [netbox_prefix.test]
  cidr = "%[1]s"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "id", "netbox_prefix.test", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccNetboxPrefixDataSource_description_single(t *testing.T) {

	testPrefix := "10.0.0.0/24"
	testDesc := "test-description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[1]s"
  status = "active"
  is_pool = true
  description = "%[2]s"
}
data "netbox_prefix" "test" {
  depends_on = [netbox_prefix.test]
  description = "%[2]s"
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "id", "netbox_prefix.test", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccNetboxPrefixDataSource_description_multiple_failure(t *testing.T) {

	testPrefix := "10.0.0.0/24"
	testPrefix2 := "10.0.64.0/26"

	testDesc := "test-description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[1]s"
  status = "active"
  is_pool = true
  description = "%[3]s"
}
resource "netbox_prefix" "test2" {
	prefix = "%[2]s"
	status = "active"
	is_pool = true
	description = "%[3]s"
  }
data "netbox_prefix" "test" {
  depends_on = [netbox_prefix.test]
  description = "%[3]s"
}`, testPrefix, testPrefix2, testDesc),
				ExpectError: regexp.MustCompile(fmt.Sprintf("Multiple matches found for %[1]s, can't continue.", testDesc)),
			},
		},
	})
}

func TestAccNetboxPrefixDataSource_description_cidr(t *testing.T) {

	testPrefix := "10.0.0.0/24"
	testDesc := "test-description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix = "%[1]s"
  status = "active"
  is_pool = true
  description = "%[2]s"
}
data "netbox_prefix" "test" {
  depends_on = [netbox_prefix.test]
  description = "%[2]s"
}`, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "id", "netbox_prefix.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "cidr", "netbox_prefix.test", "cidr"),
					resource.TestCheckResourceAttrPair("data.netbox_prefix.test", "description", "netbox_prefix.test", "description"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
