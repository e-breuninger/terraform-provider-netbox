package netbox

import (
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"testing"
)

func TestAccNetboxDeviceRoleDataSource_basic(t *testing.T) {

	testSlug := "dvrl_ds_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}
data "netbox_device_role" "test" {
  depends_on = [netbox_device_role.test]
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_role.test", "id", "netbox_device_role.test", "id"),
				),
				ExpectNonEmptyPlan: true,
			},
		},
	})
}
