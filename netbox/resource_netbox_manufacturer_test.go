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

	testSlug := "manufacturer_basic"
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
  description = "112233"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "description", "112233"),
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

func TestAccNetboxManufacturer_defaultSlug(t *testing.T) {

	testSlug := "manufacturer_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_manufacturer.test", "slug", testName),
				),
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
			for _, tag := range res.GetPayload().Results {
				if strings.HasPrefix(*tag.Name, testPrefix) {
					deleteParams := dcim.NewDcimManufacturersDeleteParams().WithID(tag.ID)
					_, err := api.Dcim.DcimManufacturersDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a tag")
				}
			}
			return nil
		},
	})
}
