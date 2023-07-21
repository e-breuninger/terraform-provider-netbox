package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceRoleDataSource_basic(t *testing.T) {
	testSlug := "dvrl_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
  tags = [netbox_tag.test.name]
}

data "netbox_device_role" "test" {
  depends_on = [netbox_device_role.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_role.test", "id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_role.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_role.test", "tags.0", testName),
				),
			},
		},
	})
}
