package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceTypeDataSource_basic(t *testing.T) {

	testSlug := "device_type_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s"
}
data "netbox_device_type" "test" {
  depends_on = [netbox_device_type.test]
  model = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_type.test", "id", "netbox_device_type.test", "id"),
				),
			},
		},
	})
}
