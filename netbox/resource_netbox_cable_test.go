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

func testAccNetboxCableFullDependencies(testName string) string {
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

resource "netbox_device_console_port" "test1" {
  device_id = netbox_device.test.id
  name = "%[1]s1"
}

resource "netbox_device_console_port" "test2" {
  device_id = netbox_device.test.id
  name = "%[1]s2"
}

resource "netbox_device_console_server_port" "test1" {
  device_id = netbox_device.test.id
  name = "%[1]s1"
}

resource "netbox_device_console_server_port" "test2" {
  device_id = netbox_device.test.id
  name = "%[1]s2"
}
`, testName)
}

func TestAccNetboxCable_basic(t *testing.T) {
	testSlug := "cable_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckCableDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxCableFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_cable" "test" {
  a_termination {
		object_type = "dcim.consoleserverport"
		object_id = netbox_device_console_server_port.test1.id
	}
	a_termination {
		object_type = "dcim.consoleserverport"
		object_id = netbox_device_console_server_port.test2.id
	}

	b_termination {
		object_type = "dcim.consoleport"
		object_id = netbox_device_console_port.test1.id
	}
	b_termination {
		object_type = "dcim.consoleport"
		object_id = netbox_device_console_port.test2.id
	}

	status = "connected"
	label = "%[1]s_label"
	type = "cat8"
	tenant_id = netbox_tenant.test.id
	color_hex = "123456"
	length = 10
	length_unit = "m"
	description = "%[1]s_description"
	comments = "%[1]s_comments"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cable.test", "status", "connected"),
					resource.TestCheckResourceAttr("netbox_cable.test", "label", testName+"_label"),
					resource.TestCheckResourceAttr("netbox_cable.test", "type", "cat8"),
					resource.TestCheckResourceAttr("netbox_cable.test", "color_hex", "123456"),
					resource.TestCheckResourceAttr("netbox_cable.test", "length", "10"),
					resource.TestCheckResourceAttr("netbox_cable.test", "length_unit", "m"),
					resource.TestCheckResourceAttr("netbox_cable.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_cable.test", "comments", testName+"_comments"),
					resource.TestCheckResourceAttr("netbox_cable.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_cable.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttr("netbox_cable.test", "a_termination.#", "2"),
					resource.TestCheckResourceAttr("netbox_cable.test", "a_termination.0.object_type", "dcim.consoleserverport"),
					resource.TestCheckResourceAttr("netbox_cable.test", "a_termination.1.object_type", "dcim.consoleserverport"),
					resource.TestCheckTypeSetElemAttrPair("netbox_cable.test", "a_termination.*.object_id", "netbox_device_console_server_port.test1", "id"),
					resource.TestCheckTypeSetElemAttrPair("netbox_cable.test", "a_termination.*.object_id", "netbox_device_console_server_port.test2", "id"),

					resource.TestCheckResourceAttr("netbox_cable.test", "b_termination.#", "2"),
					resource.TestCheckResourceAttr("netbox_cable.test", "b_termination.0.object_type", "dcim.consoleport"),
					resource.TestCheckResourceAttr("netbox_cable.test", "b_termination.1.object_type", "dcim.consoleport"),
					resource.TestCheckTypeSetElemAttrPair("netbox_cable.test", "b_termination.*.object_id", "netbox_device_console_port.test1", "id"),
					resource.TestCheckTypeSetElemAttrPair("netbox_cable.test", "b_termination.*.object_id", "netbox_device_console_port.test2", "id"),

					resource.TestCheckResourceAttrPair("netbox_cable.test", "tenant_id", "netbox_tenant.test", "id"),
				),
			},
			{
				Config: testAccNetboxCableFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_cable" "test" {
  a_termination {
		object_type = "dcim.consoleserverport"
		object_id = netbox_device_console_server_port.test1.id
	}

	b_termination {
		object_type = "dcim.consoleport"
		object_id = netbox_device_console_port.test1.id
	}

	status = "connected"
	label = "%[1]s_label"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cable.test", "status", "connected"),
					resource.TestCheckResourceAttr("netbox_cable.test", "label", testName+"_label"),
					resource.TestCheckResourceAttr("netbox_cable.test", "type", ""),
					resource.TestCheckResourceAttr("netbox_cable.test", "color_hex", ""),
					resource.TestCheckResourceAttr("netbox_cable.test", "length", "0"),
					resource.TestCheckResourceAttr("netbox_cable.test", "length_unit", ""),
					resource.TestCheckResourceAttr("netbox_cable.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_cable.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_cable.test", "tags.#", "0"),
					resource.TestCheckResourceAttr("netbox_cable.test", "tenant_id", "0"),

					resource.TestCheckResourceAttr("netbox_cable.test", "a_termination.#", "1"),
					resource.TestCheckResourceAttr("netbox_cable.test", "a_termination.0.object_type", "dcim.consoleserverport"),
					resource.TestCheckResourceAttrPair("netbox_cable.test", "a_termination.0.object_id", "netbox_device_console_server_port.test1", "id"),

					resource.TestCheckResourceAttr("netbox_cable.test", "b_termination.#", "1"),
					resource.TestCheckResourceAttr("netbox_cable.test", "b_termination.0.object_type", "dcim.consoleport"),
					resource.TestCheckResourceAttrPair("netbox_cable.test", "b_termination.0.object_id", "netbox_device_console_port.test1", "id"),
				),
			},
			{
				ResourceName:      "netbox_cable.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckCableDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each cable
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_cable" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimCablesReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimCablesRead(params, nil)

		if err == nil {
			return fmt.Errorf("cable (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimCablesReadDefault); ok {
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
	resource.AddTestSweepers("netbox_cable", &resource.Sweeper{
		Name:         "netbox_cable",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimCablesListParams()
			res, err := api.Dcim.DcimCablesList(params, nil)
			if err != nil {
				return err
			}
			for _, cable := range res.GetPayload().Results {
				if strings.HasPrefix(cable.Label, testPrefix) {
					deleteParams := dcim.NewDcimCablesDeleteParams().WithID(cable.ID)
					_, err := api.Dcim.DcimCablesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a cable")
				}
			}
			return nil
		},
	})
}
