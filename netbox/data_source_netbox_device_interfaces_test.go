package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceInterfacesDataSource_basic(t *testing.T) {
	testSlug := "dev_ifaces_ds_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxDeviceInterfacesDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_device_interfaces" "by_name" {
  filter {
    name = "name"
    value  = netbox_device_interface.test.name
  }
}

data "netbox_device_interfaces" "by_device_id" {
  filter {
    name = "device_id"
    value  = netbox_device.test.id
  }
}

data "netbox_device_interfaces" "by_mac_address" {
  filter {
    name = "mac_address"
    value  = netbox_mac_address.test.mac_address
  }
}

data "netbox_device_interfaces" "by_tag" {
  filter {
    name = "tag"
    value  = netbox_tag.test.name
  }
}
`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.by_name", "interfaces.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.by_name", "interfaces.0.type", "1000base-t"),
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.by_name", "interfaces.0.name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.by_name", "interfaces.0.enabled", "true"),
					resource.TestCheckResourceAttrPair("data.netbox_device_interfaces.by_name", "interfaces.0.device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.by_device_id", "interfaces.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.by_mac_address", "interfaces.#", "1"),
					resource.TestCheckResourceAttrSet("data.netbox_device_interfaces.by_mac_address", "interfaces.0.mac_addresses.0.id"),
					resource.TestCheckResourceAttrSet("data.netbox_device_interfaces.by_mac_address", "interfaces.0.mac_addresses.0.mac_address"),
					resource.TestCheckResourceAttrSet("data.netbox_device_interfaces.by_mac_address", "interfaces.0.mac_addresses.0.description"),
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.by_tag", "interfaces.#", "2"),
				),
			},
		},
	})
}

func testAccNetboxDeviceInterfacesDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device_interface" "test" {
  name = "%[1]s"
  device_id = netbox_device.test.id
  tags = [netbox_tag.test.name]
  type = "1000base-t"
}

resource "netbox_device_interface" "test2" {
  name = "%[1]s_two"
  device_id = netbox_device.test.id
  tags = [netbox_tag.test.name]
  type = "1000base-t"
}

resource "netbox_mac_address" "test" {
  mac_address = "F4:02:BA:7F:FD:F8"
  device_interface_id = netbox_device_interface.test.id
  description = "%[1]s"
}
`, testName)
}
