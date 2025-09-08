package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxModule_basic(t *testing.T) {
	testSerial := testAccGetTestName("module_basic")
	testManufacturer := testAccGetTestName("manufacturer")
	testDeviceType := testAccGetTestName("device_type")
	testDevice := testAccGetTestName("device")
	testModuleType := testAccGetTestName("module_type")
	testModuleBay := testAccGetTestName("module_bay")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
}

resource "netbox_device_type" "test" {
  model          = "%s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name          = "%s"
  device_type_id = netbox_device_type.test.id
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model          = "%s"
}

resource "netbox_device_module_bay" "test" {
  device_id      = netbox_device.test.id
  name          = "%s"
}

resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status        = "active"
  serial        = "%s"
  asset_tag     = "MT-001"
  description   = "Test module"
}`, testManufacturer, testDeviceType, testDevice, testModuleType, testModuleBay, testSerial),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_module.test", "serial", testSerial),
					resource.TestCheckResourceAttr("netbox_module.test", "asset_tag", "MT-001"),
					resource.TestCheckResourceAttr("netbox_module.test", "description", "Test module"),
				),
			},
			{
				ResourceName:      "netbox_module.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxModule_minimal(t *testing.T) {
	testManufacturer := testAccGetTestName("manufacturer")
	testDeviceType := testAccGetTestName("device_type")
	testDevice := testAccGetTestName("device")
	testModuleType := testAccGetTestName("module_type")
	testModuleBay := testAccGetTestName("module_bay")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
}

resource "netbox_device_type" "test" {
  model          = "%s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name          = "%s"
  device_type_id = netbox_device_type.test.id
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model          = "%s"
}

resource "netbox_device_module_bay" "test" {
  device_id      = netbox_device.test.id
  name          = "%s"
}

resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status        = "planned"
}`, testManufacturer, testDeviceType, testDevice, testModuleType, testModuleBay),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module.test", "status", "planned"),
				),
			},
		},
	})
}

func TestAccNetboxModule_withTags(t *testing.T) {
	testSerial := testAccGetTestName("module_tags")
	testManufacturer := testAccGetTestName("manufacturer")
	testDeviceType := testAccGetTestName("device_type")
	testDevice := testAccGetTestName("device")
	testModuleType := testAccGetTestName("module_type")
	testModuleBay := testAccGetTestName("module_bay")
	testTag := testAccGetTestName("tag")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
}

resource "netbox_device_type" "test" {
  model          = "%s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name          = "%s"
  device_type_id = netbox_device_type.test.id
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model          = "%s"
}

resource "netbox_device_module_bay" "test" {
  device_id      = netbox_device.test.id
  name          = "%s"
}

resource "netbox_tag" "test" {
  name = "%s"
}

resource "netbox_module" "test" {
  device_id      = netbox_device.test.id
  module_bay_id  = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status        = "active"
  serial        = "%s"
  tags          = [netbox_tag.test.slug]
}`, testManufacturer, testDeviceType, testDevice, testModuleType, testModuleBay, testTag, testSerial),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_module.test", "serial", testSerial),
					resource.TestCheckResourceAttr("netbox_module.test", "tags.#", "1"),
				),
			},
		},
	})
}
