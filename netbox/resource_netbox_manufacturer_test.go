package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxManufacturer_basic(t *testing.T) {
	testSlug := "manufacturer"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_manufacturer.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_manufacturer", &resource.Sweeper{
		Name:         "netbox_manufacturer",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimManufacturersListParams()
			res, err := api.Dcim.DcimManufacturersList(params, nil)
			if err != nil {
				return err
			}
			for _, manufacturer := range res.GetPayload().Results {
				if strings.HasPrefix(*manufacturer.Name, testPrefix) {
					deleteParams := dcim.NewDcimManufacturersDeleteParams().WithID(manufacturer.ID)
					_, err := api.Dcim.DcimManufacturersDelete(deleteParams, nil)
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
