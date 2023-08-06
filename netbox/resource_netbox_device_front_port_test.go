package netbox

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxDeviceFrontPortFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_tag" "test" {
  name = "%[1]sa"
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

resource "netbox_device" "test" {
  name = "%[1]s"
  device_type_id = netbox_device_type.test.id
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device_module_bay" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
}

resource "netbox_module_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model = "%[1]s"
}

resource "netbox_module" "test" {
  device_id = netbox_device.test.id
  module_bay_id = netbox_device_module_bay.test.id
  module_type_id = netbox_module_type.test.id
  status = "active"
}

resource "netbox_device_rear_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "8p8c"
  positions = 1
  mark_connected = true
}
`, testName)
}

func TestAccNetboxDeviceFrontPort_basic(t *testing.T) {
	testSlug := "device_front_port_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDeviceFrontPortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDeviceFrontPortFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "8p8c"
  rear_port_id = netbox_device_rear_port.test.id
  rear_port_position = 1

  mark_connected = true
  module_id = netbox_module.test.id
  label = "%[1]s_label"
  color_hex = "123456"
  description = "%[1]s_description"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "type", "8p8c"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "label", testName+"_label"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "color_hex", "123456"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_port_position", "1"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttrPair("netbox_device_front_port.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_front_port.test", "rear_port_id", "netbox_device_rear_port.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_front_port.test", "module_id", "netbox_module.test", "id"),
				),
			},
			{
				Config: testAccNetboxDeviceFrontPortFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device_front_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "8p8c"
  rear_port_id = netbox_device_rear_port.test.id
  rear_port_position = 1
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "type", "8p8c"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "mark_connected", "false"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "label", ""),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "color_hex", ""),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "rear_port_position", "1"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_device_front_port.test", "module_id", "0"),

					resource.TestCheckResourceAttrPair("netbox_device_front_port.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_front_port.test", "rear_port_id", "netbox_device_rear_port.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_front_port.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDeviceFrontPortDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each front port
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_device_front_port" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimFrontPortsReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimFrontPortsRead(params, nil)

		if err == nil {
			return fmt.Errorf("device_front_port (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimFrontPortsReadDefault); ok {
				errorcode := errresp.Code()
				if errorcode == 404 {
					return nil
				}
			}
			return err
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("netbox_device_front_port", &resource.Sweeper{
		Name:         "netbox_device_front_port",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimFrontPortsListParams()
			res, err := api.Dcim.DcimFrontPortsList(params, nil)
			if err != nil {
				return err
			}
			for _, frontPort := range res.GetPayload().Results {
				if strings.HasPrefix(*frontPort.Name, testPrefix) {
					deleteParams := dcim.NewDcimFrontPortsDeleteParams().WithID(frontPort.ID)
					_, err := api.Dcim.DcimFrontPortsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device_front_port")
				}
			}
			return nil
		},
	})
}
