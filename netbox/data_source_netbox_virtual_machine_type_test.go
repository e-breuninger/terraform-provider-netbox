package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualMachineTypeDataSource_basic(t *testing.T) {
	testSlug := "vm_type_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_virtual_machine_type" "test" {
  name        = "%[1]s"
  slug        = "%[1]s"
  description = "This is my virtual machine type"
  comments    = "Some comments"
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
data "netbox_virtual_machine_type" "test" {
  depends_on = []

  name = netbox_virtual_machine_type.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machine_type.test", "id", "netbox_virtual_machine_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_type.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_type.test", "slug", testName),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_type.test", "description", "This is my virtual machine type"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_type.test", "comments", "Some comments"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_type.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_type.test", "tags.0", testName),
				),
			},
		},
	})
}
