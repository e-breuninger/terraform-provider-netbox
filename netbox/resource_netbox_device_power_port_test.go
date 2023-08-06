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

func testAccNetboxDevicePowerPortFullDependencies(testName string) string {
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
`, testName)
}

func TestAccNetboxDevicePowerPort_basic(t *testing.T) {
	testSlug := "device_power_port_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDevicePowerPortDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePowerPortFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device_power_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"

	maximum_draw = 750
  allocated_draw = 500
	mark_connected = true  
	type = "iec-60320-c6"
  module_id = netbox_module.test.id
  description = "%[1]s_description"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "maximum_draw", "750"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "allocated_draw", "500"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "type", "iec-60320-c6"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttrPair("netbox_device_power_port.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_power_port.test", "module_id", "netbox_module.test", "id"),
				),
			},
			{
				Config: testAccNetboxDevicePowerPortFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device_power_port" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "maximum_draw", "0"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "allocated_draw", "0"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "type", ""),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "mark_connected", "false"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_device_power_port.test", "module_id", "0"),

					resource.TestCheckResourceAttrPair("netbox_device_power_port.test", "device_id", "netbox_device.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_power_port.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDevicePowerPortDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each power port
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_device_power_port" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimPowerPortsReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimPowerPortsRead(params, nil)

		if err == nil {
			return fmt.Errorf("device_power_port (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimPowerPortsReadDefault); ok {
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
	resource.AddTestSweepers("netbox_device_power_port", &resource.Sweeper{
		Name:         "netbox_device_power_port",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimPowerPortsListParams()
			res, err := api.Dcim.DcimPowerPortsList(params, nil)
			if err != nil {
				return err
			}
			for _, powerPort := range res.GetPayload().Results {
				if strings.HasPrefix(*powerPort.Name, testPrefix) {
					deleteParams := dcim.NewDcimPowerPortsDeleteParams().WithID(powerPort.ID)
					_, err := api.Dcim.DcimPowerPortsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device_power_port")
				}
			}
			return nil
		},
	})
}
