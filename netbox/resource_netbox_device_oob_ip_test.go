package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxDeviceOobIPDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name           = "%[1]s"
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
}

resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s"
  type      = "1000base-t"
}

resource "netbox_ip_address" "test" {
  ip_address   = "10.0.0.5/24"
  status       = "active"
  interface_id = netbox_device_interface.test.id
  object_type  = "dcim.interface"
}
`, testName)
}

func TestAccNetboxDeviceOobIP_basic(t *testing.T) {
	testSlug := "device_oob_ip_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxDeviceOobIPDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies + `
resource "netbox_device_oob_ip" "test" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_oob_ip.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_oob_ip.test", "ip_address_id", "netbox_ip_address.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_oob_ip.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
