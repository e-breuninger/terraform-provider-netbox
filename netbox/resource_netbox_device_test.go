package netbox

import (
	"context"
	"fmt"
	"strconv"
	"testing"

	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
)

func TestAccNetboxDevice_basic(t *testing.T) {
	device_type_id := "7"
	device_role_id := "4"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckNetboxDeviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckNetboxDeviceConfigBasic(device_type_id, device_role_id),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckNetboxDeviceExists("netbox_dcim_device.test"),
				),
			},
		},
	})
}

func testAccCheckNetboxDeviceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_dcim_device" {
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
			if err.(*runtime.APIError).Code == 404 {
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

func testAccCheckNetboxDeviceConfigBasic(device_type_id string, device_role_id string) string {
	return fmt.Sprintf(`
resource "netbox_extras_tag" "test-device" {
  name = "Test Device"
  slug = "test-device"
}
  
resource "netbox_dcim_site" "test-device" {
	name = "test-device"
	slug = "test-device"
	status = "active"

	custom_fields = {
		tf-test = "customFieldValue"
	}
}

resource "netbox_dcim_rack" "test-device" {
	name = "rack-test-device"
	site_id = netbox_dcim_site.test-device.id

	custom_fields = {
		rackCustomField = "rackCustomeFieldValue"
	  }

}
resource "netbox_dcim_device" "test" {
	device_type_id = "%s"
	device_role_id = "%s"
	site_id = netbox_dcim_site.test-device.id


	tags {
		name = netbox_extras_tag.test-device.name
		slug = netbox_extras_tag.test-device.slug
	}
	custom_fields = {
		deviceCsutomField = "deviceCustomFieldValue"
	}
}

`, device_type_id, device_role_id)
}
