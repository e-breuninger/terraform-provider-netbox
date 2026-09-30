package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualCircuitTypeDataSource_basic(t *testing.T) {
	testSlug := "virtual_circuit_type_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_virtual_circuit_type" "test" {
  name        = "%[1]s"
  slug        = "%[1]s"
  color_hex   = "123456"
  description = "This is my virtual circuit type"
  comments    = "This is my virtual circuit type comment"
  tags        = [netbox_tag.test.name]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_virtual_circuit_type" "test" {
  depends_on = []

  name = netbox_virtual_circuit_type.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit_type.test", "id", "netbox_virtual_circuit_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_type.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_type.test", "slug", testName),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_type.test", "color_hex", "123456"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_type.test", "description", "This is my virtual circuit type"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_type.test", "comments", "This is my virtual circuit type comment"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_type.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_type.test", "tags.0", testName),
				),
			},
		},
	})
}
