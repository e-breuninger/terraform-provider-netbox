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

func testAccNetboxDeviceInterfaceFullDependencies(testName string) string {
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

resource "netbox_vlan" "test1" {
  name = "%[1]s_vlan1"
  vid = 1001
  tags = []
}

resource "netbox_vlan" "test2" {
  name = "%[1]s_vlan2"
  vid = 1002
  tags = []
}`, testName)
}

func testAccNetboxDeviceInterface_basic(testName string) string {
	return fmt.Sprintf(`
resource "netbox_device_interface" "test" {
  name = "%s"
  device_id = netbox_device.test.id
  tags = ["%[1]s"]
  type = "1000base-t"
}`, testName)
}

func testAccNetboxDeviceInterface_opts(testName string, testMac string) string {
	return fmt.Sprintf(`
resource "netbox_device_interface" "test" {
  name = "%[1]s"
  description = "%[1]s"
  enabled = true
  mgmtonly = true
  mac_address = "%[2]s"
  mtu = 1440
  device_id = netbox_device.test.id
  type = "1000base-t"
}`, testName, testMac)
}

func testAccNetboxDeviceInterface_vlans(testName string) string {
	return fmt.Sprintf(`
resource "netbox_device_interface" "test1" {
  name = "%[1]s_1"
  mode = "access"
  untagged_vlan = netbox_vlan.test1.id
  device_id = netbox_device.test.id
  type = "1000base-t"
}

resource "netbox_device_interface" "test2" {
  name = "%[1]s_2"
  mode = "tagged"
  tagged_vlans = [netbox_vlan.test2.id]
  untagged_vlan = netbox_vlan.test1.id
  device_id = netbox_device.test.id
  type = "1000base-t"
}

resource "netbox_device_interface" "test3" {
  name = "%[1]s_3"
  mode = "tagged-all"
  tagged_vlans = [netbox_vlan.test1.id, netbox_vlan.test2.id]
  device_id = netbox_device.test.id
  type = "1000base-t"
}`, testName)
}

func TestAccNetboxDeviceInterface_basic(t *testing.T) {
	testSlug := "iface_basic"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxDeviceInterfaceFullDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeviceInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: setUp + testAccNetboxDeviceInterface_basic(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "type", "1000base-t"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_device_interface.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxDeviceInterface_opts(t *testing.T) {
	testSlug := "iface_mac"
	testMac := "00:01:02:03:04:05"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxDeviceInterfaceFullDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeviceInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: setUp + testAccNetboxDeviceInterface_opts(testName, testMac),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "type", "1000base-t"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mgmtonly", "true"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mac_address", "00:01:02:03:04:05"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mtu", "1440"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "device_id", "netbox_device.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_interface.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxDeviceInterface_vlans(t *testing.T) {
	testSlug := "iface_vlan"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxDeviceInterfaceFullDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeviceInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: setUp + testAccNetboxDeviceInterface_vlans(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test1", "mode", "access"),
					resource.TestCheckResourceAttr("netbox_device_interface.test2", "mode", "tagged"),
					resource.TestCheckResourceAttr("netbox_device_interface.test3", "mode", "tagged-all"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test1", "untagged_vlan", "netbox_vlan.test1", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test2", "untagged_vlan", "netbox_vlan.test1", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test2", "tagged_vlans.0", "netbox_vlan.test2", "id"),
					resource.TestCheckResourceAttr("netbox_device_interface.test3", "tagged_vlans.#", "2"),
				),
			},
			{
				ResourceName:      "netbox_device_interface.test1",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_device_interface.test2",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_device_interface.test3",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDeviceInterfaceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each interface
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_device_interface" {
			continue
		}

		// Retrieve our interface by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimInterfacesReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimInterfacesRead(params, nil)

		if err == nil {
			return fmt.Errorf("device interface (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimInterfacesReadDefault); ok {
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
	resource.AddTestSweepers("netbox_device_interface", &resource.Sweeper{
		Name:         "netbox_device_interface",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimInterfacesListParams()
			res, err := api.Dcim.DcimInterfacesList(params, nil)
			if err != nil {
				return err
			}
			for _, intrface := range res.GetPayload().Results {
				if strings.HasPrefix(*intrface.Name, testPrefix) {
					deleteParams := dcim.NewDcimInterfacesDeleteParams().WithID(intrface.ID)
					_, err := api.Dcim.DcimInterfacesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device interface")
				}
			}
			return nil
		},
	})
}
