package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxModuleType_basic(t *testing.T) {
	testModel := testAccGetTestName("module_type_basic")
	testManufacturer := testAccGetTestName("manufacturer")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model          = "%s"
  part_number    = "MT-1000"
  weight         = 2.5
  weight_unit    = "kg"
  description    = "Test module type"
}`, testManufacturer, testModel),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type.test", "model", testModel),
					resource.TestCheckResourceAttr("netbox_module_type.test", "part_number", "MT-1000"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight", "2.5"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "weight_unit", "kg"),
					resource.TestCheckResourceAttr("netbox_module_type.test", "description", "Test module type"),
				),
			},
			{
				ResourceName:      "netbox_module_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxModuleType_minimal(t *testing.T) {
	testModel := testAccGetTestName("module_type_minimal")
	testManufacturer := testAccGetTestName("manufacturer")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model          = "%s"
}`, testManufacturer, testModel),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type.test", "model", testModel),
					resource.TestCheckResourceAttr("netbox_module_type.test", "part_number", ""),
				),
			},
		},
	})
}

func TestAccNetboxModuleType_withTags(t *testing.T) {
	testModel := testAccGetTestName("module_type_tags")
	testManufacturer := testAccGetTestName("manufacturer")
	testTag := testAccGetTestName("tag")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
}

resource "netbox_tag" "test" {
  name = "%s"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model          = "%s"
  tags = [netbox_tag.test.slug]
}`, testManufacturer, testTag, testModel),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_type.test", "model", testModel),
					resource.TestCheckResourceAttr("netbox_module_type.test", "tags.#", "1"),
				),
			},
		},
	})
}
