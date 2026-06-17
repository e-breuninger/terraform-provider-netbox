package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceRenderConfigDataSource_basic(t *testing.T) {
	testSlug := "device_render_cfg"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_config_template" "test" {
  name = "%[1]s"
  template_code = "hostname={{ device.name }}"
}

resource "netbox_device" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
  config_template_id = netbox_config_template.test.id
}

data "netbox_device_render_config" "test" {
  depends_on = [netbox_device.test]
  device_id = netbox_device.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device_render_config.test", "id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_render_config.test", "content", fmt.Sprintf("hostname=%s", testName)),
					resource.TestCheckResourceAttrPair("data.netbox_device_render_config.test", "config_template_id", "netbox_config_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_render_config.test", "config_template_name", testName),
				),
			},
		},
	})
}
