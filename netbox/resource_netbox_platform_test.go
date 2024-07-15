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

func TestAccNetboxPlatform_basic(t *testing.T) {
	testSlug := "platform_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_platform" "test" {
  name = "%s"
  slug = "%s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_platform.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_platform.test", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_platform.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPlatform_manufacturer(t *testing.T) {
	testSlug := "platform_manufacturer"
	testName := testAccGetTestName(testSlug)
	testManufacturer := "manu_test"
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[3]s"
}

resource "netbox_platform" "test" {
  name = "%[1]s"
  slug = "%[2]s"
  manufacturer_id = netbox_manufacturer.test.id
}`, testName, randomSlug, testManufacturer),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_platform.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_platform.test", "slug", randomSlug),
					resource.TestCheckResourceAttrPair("netbox_platform.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_platform.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxPlatform_defaultSlug(t *testing.T) {
	testSlug := "platform_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_platform" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_platform.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_platform.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_platform", &resource.Sweeper{
		Name:         "netbox_platform",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimPlatformsListParams()
			res, err := api.Dcim.DcimPlatformsList(params, nil)
			if err != nil {
				return err
			}
			for _, platform := range res.GetPayload().Results {
				if strings.HasPrefix(*platform.Name, testPrefix) {
					deleteParams := dcim.NewDcimPlatformsDeleteParams().WithID(platform.ID)
					_, err := api.Dcim.DcimPlatformsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a platform")
				}
			}
			return nil
		},
	})
}
