package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDeviceType_basic(t *testing.T) {
	testSlug := "device_type"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  slug = "%[2]s"
  part_number = "%[2]s"
  u_height = "0.5"
  manufacturer_id = netbox_manufacturer.test.id
  is_full_depth = true
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_device_type.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "part_number", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "u_height", "0.5"),
					resource.TestCheckResourceAttrPair("netbox_device_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "is_full_depth", "true"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  slug = "%[2]s"
  part_number = "%[2]s"
  u_height = "0.5"
  manufacturer_id = netbox_manufacturer.test.id
  is_full_depth = false
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_type.test", "model", testName),
					resource.TestCheckResourceAttr("netbox_device_type.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "part_number", randomSlug),
					resource.TestCheckResourceAttr("netbox_device_type.test", "u_height", "0.5"),
					resource.TestCheckResourceAttrPair("netbox_device_type.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_type.test", "is_full_depth", "false"),
				),
			},
			{
				ResourceName:      "netbox_device_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_device_type", &resource.Sweeper{
		Name:         "netbox_device_type",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimDeviceTypesListParams()
			res, err := api.Dcim.DcimDeviceTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, devicetype := range res.GetPayload().Results {
				if strings.HasPrefix(*devicetype.Model, testPrefix) {
					deleteParams := dcim.NewDcimDeviceTypesDeleteParams().WithID(devicetype.ID)
					_, err := api.Dcim.DcimDeviceTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a device type")
				}
			}
			return nil
		},
	})
}
