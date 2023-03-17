package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRackRoleDataSource_basic(t *testing.T) {

	testSlug := "rack_role_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_rack_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
	description = "%[1]sdescription"
  tags = [netbox_tag.test.name]
}

data "netbox_rack_role" "test" {
  depends_on = [netbox_rack_role.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_rack_role.test", "id", "netbox_rack_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_rack_role.test", "name", "netbox_rack_role.test", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_rack_role.test", "color_hex", "netbox_rack_role.test", "color_hex"),
					resource.TestCheckResourceAttrPair("data.netbox_rack_role.test", "description", "netbox_rack_role.test", "description"),
					resource.TestCheckResourceAttr("data.netbox_rack_role.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_rack_role.test", "tags.0", testName),
				),
			},
		},
	})
}
