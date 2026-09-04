package netbox

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccNetboxDeviceTagDependencies(testName string) string {
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

resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxDeviceTag_basic(t *testing.T) {
	testSlug := "device_tag_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxDeviceTagDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies + `
resource "netbox_device_tag" "test" {
  device_id = netbox_device.test.id
  tag       = netbox_tag.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_tag.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_tag.test", "tag", "netbox_tag.test", "name"),
				),
				// netbox_device.test does not declare the tag itself, so once netbox_device_tag
				// attaches it out of band, netbox_device.test's own tags/tags_all reconciliation
				// wants to strip it back off. This is the documented "don't combine with a
				// netbox_device resource managing the same device" caveat, not a bug in this test.
				ExpectNonEmptyPlan: true,
			},
			{
				ResourceName:      "netbox_device_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccNetboxDeviceTagPrimaryIPDependencies(testName string) string {
	return testAccNetboxDeviceTagDependencies(testName) + fmt.Sprintf(`
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

resource "netbox_device_primary_ip" "test" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.test.id
}
`, testName)
}

// testAccCheckNetboxDevicePrimaryIP4 reads the device straight from the NetBox API (bypassing
// Terraform state, which is not refreshed for netbox_device.test as a side effect of applying
// netbox_device_tag) to verify that attaching/detaching a tag does not clobber the primary IP.
func testAccCheckNetboxDevicePrimaryIP4(deviceResourceName, ipResourceName string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		deviceRs, ok := s.RootModule().Resources[deviceResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", deviceResourceName)
		}
		ipRs, ok := s.RootModule().Resources[ipResourceName]
		if !ok {
			return fmt.Errorf("resource not found: %s", ipResourceName)
		}

		deviceID, err := strconv.ParseInt(deviceRs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}
		expectedIPID, err := strconv.ParseInt(ipRs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		api := testAccProvider.Meta().(*providerState)
		res, err := api.Dcim.DcimDevicesRead(dcim.NewDcimDevicesReadParams().WithID(deviceID), nil)
		if err != nil {
			return err
		}

		primaryIP4 := res.GetPayload().PrimaryIp4
		if primaryIP4 == nil {
			return fmt.Errorf("device %d has no primary_ip4 set, expected %d", deviceID, expectedIPID)
		}
		if primaryIP4.ID != expectedIPID {
			return fmt.Errorf("device %d primary_ip4 is %d, expected %d", deviceID, primaryIP4.ID, expectedIPID)
		}
		return nil
	}
}

func TestAccNetboxDeviceTag_preservesPrimaryIP(t *testing.T) {
	testSlug := "device_tag_preserve_primary_ip"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxDeviceTagPrimaryIPDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// baseline: primary IP set, no netbox_device_tag involved yet
				Config: dependencies,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetboxDevicePrimaryIP4("netbox_device.test", "netbox_ip_address.test"),
				),
			},
			{
				// attaching a tag must not clobber the device's primary IP
				Config: dependencies + `
resource "netbox_device_tag" "test" {
  device_id = netbox_device.test.id
  tag       = netbox_tag.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_tag.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_tag.test", "tag", "netbox_tag.test", "name"),
					testAccCheckNetboxDevicePrimaryIP4("netbox_device.test", "netbox_ip_address.test"),
				),
				// see the comment in TestAccNetboxDeviceTag_basic: netbox_device.test wants to
				// strip the tag it doesn't itself declare. Expected, not what this step tests.
				ExpectNonEmptyPlan: true,
			},
			{
				// detaching the tag (netbox_device_tag no longer in config) must not clobber it either
				Config: dependencies,
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetboxDevicePrimaryIP4("netbox_device.test", "netbox_ip_address.test"),
				),
			},
		},
	})
}
