package netbox

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxDeviceBayFullDependencies(testName string) string {
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

resource "netbox_device" "test_installed" {
  name = "%[1]s_installed"
  device_type_id = netbox_device_type.test.id
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
}
`, testName)
}

func TestAccNetboxDeviceBay_basic(t *testing.T) {
	testSlug := "device_bay_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers:    testAccProviders,
		PreCheck:     func() { testAccPreCheck(t) },
		CheckDestroy: testAccCheckDeviceBayDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDeviceBayFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device_bay" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  label = "%[1]s_label"
  description = "%[1]s_description"
  installed_device_id = netbox_device.test_installed.id
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_bay.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "label", testName+"_label"),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "description", testName+"_description"),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "tags.0", testName+"a"),

					resource.TestCheckResourceAttrPair("netbox_device_bay.test", "installed_device_id", "netbox_device.test_installed", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_bay.test", "device_id", "netbox_device.test", "id"),
				),
			},
			{
				Config: testAccNetboxDeviceBayFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device_bay" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_bay.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "label", ""),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "installed_device_id", ""),
					resource.TestCheckResourceAttr("netbox_device_bay.test", "tags.#", "0"),

					resource.TestCheckResourceAttrPair("netbox_device_bay.test", "device_id", "netbox_device.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_bay.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckDeviceBayDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*providerState)

	// loop through the resources in state, verifying each module bay
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_device_bay" {
			continue
		}

		// Retrieve our device by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := dcim.NewDcimDeviceBaysReadParams().WithID(stateID)
		_, err := conn.Dcim.DcimDeviceBaysRead(params, nil)

		if err == nil {
			return fmt.Errorf("device_bay (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*dcim.DcimDeviceBaysReadDefault); ok {
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
	resource.AddTestSweepers("netbox_device_bay", &resource.Sweeper{
		Name:         "netbox_device_bay",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimDeviceBaysListParams()
			res, err := api.Dcim.DcimDeviceBaysList(params, nil)
			if err != nil {
				return err
			}
			for _, deviceBay := range res.GetPayload().Results {
				if strings.HasPrefix(*deviceBay.Name, testPrefix) {
					deleteParams := dcim.NewDcimDeviceBaysDeleteParams().WithID(deviceBay.ID)
					_, err := api.Dcim.DcimDeviceBaysDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device_bay")
				}
			}
			return nil
		},
	})
}
