package netbox

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func TestAccNetboxDevice_basic(t *testing.T) {
	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetboxDeviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNetboxDeviceConfigBasic(),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetboxDeviceExists("netbox_device.test-device"),
				),
			},
		},
	})
}

func testAccCheckNetboxDeviceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_device" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimDevicesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Dcim.DcimDevicesRead(params, nil)
		if err != nil {
			if err.(*dcim.DcimDevicesReadDefault).Code() == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Device ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckNetboxDeviceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No device ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimDevicesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Dcim.DcimDevicesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckNetboxDeviceConfigBasic() string {
	return fmt.Sprintf(`
resource "netbox_tag" "test-device" {
  name = "test-device"
}
resource "netbox_site" "test-device" {
  name = "test-device"
  slug = "test-device"
  status = "active"
}
resource "netbox_manufacturer" "test-device" {
  name = "test-device"
}
resource "netbox_device_type" "test-device" {
  manufacturer_id = netbox_manufacturer.test-device.id
  model = "test-device"
  slug = "test-device"
}
resource "netbox_device_role" "test-device" {
  name = "test-device"
  color_hex = "112233"
}
resource "netbox_device" "test-device" {
  device_type_id = netbox_device_type.test-device.id
  device_role_id = netbox_device_role.test-device.id
  site_id = netbox_site.test-device.id

  tags = ["test-device"]
}

`)
}
