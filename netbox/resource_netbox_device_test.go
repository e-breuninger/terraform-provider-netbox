package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccNetboxDeviceFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id = netbox_site.test.id
}

resource "netbox_location" "test" {
	name = "%[1]s"
	site_id =netbox_site.test.id
}

resource "netbox_rack_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_rack" "test" {
  name = "%[1]s"
	site_id = netbox_site.test.id
	status = "reserved"
	width = 19
	u_height = 48
	tenant_id = netbox_tenant.test.id
	location_id = netbox_location.test.id
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}

resource "netbox_tag" "test_b" {
  name = "%[1]sb"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}`, testName)
}

func TestAccNetboxDevice_basic(t *testing.T) {

	testSlug := "device_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckDeviceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDeviceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name = "%[1]s"
  comments = "thisisacomment"
  tenant_id = netbox_tenant.test.id
  platform_id = netbox_platform.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  tags = ["%[1]sa"]
  site_id = netbox_site.test.id
  cluster_id = netbox_cluster.test.id
  location_id = netbox_location.test.id
  status = "staged"
  serial = "ABCDEF"
	rack_id = netbox_rack.test.id
	rack_face = "front"
	rack_position = 10
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "rack_id", "netbox_rack.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "staged"),
					resource.TestCheckResourceAttr("netbox_device.test", "serial", "ABCDEF"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName+"a"),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_face", "front"),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_position", "10"),
				),
			},
			{
				Config: testAccNetboxDeviceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name = "%[1]s"
  comments = "thisisacomment"
  tenant_id = netbox_tenant.test.id
  platform_id = netbox_platform.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  tags = ["%[1]sa"]
  site_id = netbox_site.test.id
  cluster_id = netbox_cluster.test.id
  location_id = netbox_location.test.id
	rack_id = netbox_rack.test.id
  status = "staged"
  serial = "ABCDEF"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "rack_id", "netbox_rack.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "staged"),
					resource.TestCheckResourceAttr("netbox_device.test", "serial", "ABCDEF"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName+"a"),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_face", ""),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_position", "0"),
				),
			},
			{
				Config: testAccNetboxDeviceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name = "%[1]s"
  comments = "thisisacomment"
  tenant_id = netbox_tenant.test.id
  platform_id = netbox_platform.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  tags = ["%[1]sa"]
  site_id = netbox_site.test.id
  cluster_id = netbox_cluster.test.id
  location_id = netbox_location.test.id
  status = "staged"
  serial = "ABCDEF"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "staged"),
					resource.TestCheckResourceAttr("netbox_device.test", "serial", "ABCDEF"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName+"a"),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_id", "0"),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_face", ""),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_position", "0"),
				),
			},
			{
				ResourceName:      "netbox_device.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDeviceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each device
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_device" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimDevicesReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimDevicesRead(params, nil)

		if err == nil {
			return fmt.Errorf("device (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			errorcode := err.(*dcim.DcimDevicesReadDefault).Code()
			if errorcode == 404 {
				return nil
			}
			return err
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("netbox_device", &resource.Sweeper{
		Name:         "netbox_device",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimDevicesListParams()
			res, err := api.Dcim.DcimDevicesList(params, nil)
			if err != nil {
				return err
			}
			for _, Device := range res.GetPayload().Results {
				if strings.HasPrefix(*Device.Name, testPrefix) {
					deleteParams := dcim.NewDcimDevicesDeleteParams().WithID(Device.ID)
					_, err := api.Dcim.DcimDevicesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device")
				}
			}
			return nil
		},
	})
}
