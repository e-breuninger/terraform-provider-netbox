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

func TestAccDeviceInterface_basic(t *testing.T) {
	name := "test interface"

	resource.Test(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeviceInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccCheckDeviceInterfaceConfigBasic(name),
				Check: resource.ComposeTestCheckFunc(
					testAccCheckDeviceInterfaceExists("netbox_device_interface.test-interface"),
				),
			},
		},
	})
}

func testAccCheckDeviceInterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_device_interface" {
			continue
		}

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimInterfacesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		resp, err := c.Dcim.DcimInterfacesRead(params, nil)
		if err != nil {
			if err.(*dcim.DcimInterfacesReadDefault).Code() == 404 {
				return nil
			}

			return err
		}

		return fmt.Errorf("Interface ID still exists: %d", resp.Payload.ID)
	}

	return nil
}

func testAccCheckDeviceInterfaceExists(n string) resource.TestCheckFunc {
	return func(s *terraform.State) error {
		rs, ok := s.RootModule().Resources[n]

		if !ok {
			return fmt.Errorf("Not found: %s", n)
		}

		if rs.Primary.ID == "" {
			return fmt.Errorf("No interface ID set")
		}

		c := testAccProvider.Meta().(*client.NetBoxAPI)

		objectID, err := strconv.ParseInt(rs.Primary.ID, 10, 64)
		if err != nil {
			return err
		}

		params := &dcim.DcimInterfacesReadParams{
			Context: context.Background(),
			ID:      objectID,
		}

		_, err = c.Dcim.DcimInterfacesRead(params, nil)
		if err != nil {
			return err
		}

		return nil
	}
}

func testAccCheckDeviceInterfaceConfigBasic(name string) string {
	return fmt.Sprintf(`

resource "netbox_tag" "test-interface" {
  name = "test-interface"
}
resource "netbox_site" "test-interface" {
  name = "test-interface"
  slug = "test-interface"
  status = "active"
}
resource "netbox_manufacturer" "test-interface" {
  name = "test-interface"
}
resource "netbox_device_type" "test-interface" {
  manufacturer_id = netbox_manufacturer.test-interface.id
  model = "test-interface"
  slug = "test-interface"
}
resource "netbox_device_role" "test-interface" {
  name = "test-interface"
  color_hex = "112233"
}
resource "netbox_device" "test-interface" {
  device_type_id = netbox_device_type.test-interface.id
  device_role_id = netbox_device_role.test-interface.id
  site_id = netbox_site.test-interface.id

  tags = ["test-interface"]
}
resource "netbox_device_interface" "test-interface" {
  device_id = netbox_device.test-interface.id
  type = "virtual"
  name = "%s"
}

`, name)
}
