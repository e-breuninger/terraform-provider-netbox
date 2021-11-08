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
					testAccCheckDeviceInterfaceExists("netbox_dcim_interface.test"),
				),
			},
		},
	})
}

func testAccCheckDeviceInterfaceDestroy(s *terraform.State) error {
	c := testAccProvider.Meta().(*client.NetBoxAPI)

	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_dcim_interface" {
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
			if err.(*runtime.APIError).Code == 404 {
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

resource "netbox_extras_tag" "test-interface" {
	name = "Test Interface"
	slug = "test-interface"
  }
	
resource "netbox_dcim_site" "test-interface" {
	name = "test-interface"
	slug = "test-interface"
	status = "active"
}

resource "netbox_dcim_rack" "test-interface" {
	name = "rack-test-interface"
	site_id = netbox_dcim_site.test-interface.id

}

resource "netbox_dcim_device" "test-interface" {
	device_type_id = 7
	device_role_id = 4
	site_id = netbox_dcim_site.test-interface.id

}

resource "netbox_dcim_interface" "test" {

	device_id = netbox_dcim_device.test-interface.id
	type = "virtual"
	name = "%s"
	tagged_vlan = [64]
}

`, name)
}
