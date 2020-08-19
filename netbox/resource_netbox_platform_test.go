package netbox

import (
	"fmt"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"log"
	"strings"
	"testing"
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
					resource.TestCheckResourceAttr("netbox_platform.test", "slug", testName),
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
			api := m.(*client.NetBox)
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
