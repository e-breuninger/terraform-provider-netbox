package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxModuleTypeProfileDataSource_basic(t *testing.T) {
	testSlug := "module_type_profile_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_module_type_profile" "test" {
  name        = "%[1]s"
  description = "This is my module type profile"
  comments    = "Module type profile comments"
  schema      = jsonencode({
    type = "object"
    properties = {
      port_count = {
        type = "integer"
      }
    }
  })
  tags = [netbox_tag.test.name]
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
data "netbox_module_type_profile" "test" {
  depends_on = []

  name = netbox_module_type_profile.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_module_type_profile.test", "id", "netbox_module_type_profile.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_module_type_profile.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_module_type_profile.test", "description", "This is my module type profile"),
					resource.TestCheckResourceAttr("data.netbox_module_type_profile.test", "comments", "Module type profile comments"),
					resource.TestCheckResourceAttr("data.netbox_module_type_profile.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_module_type_profile.test", "tags.0", testName),
				),
			},
		},
	})
}
