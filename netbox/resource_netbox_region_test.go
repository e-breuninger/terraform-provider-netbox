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

func TestAccNetboxRegion_basic(t *testing.T) {

	testSlug := "region_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%s"
  slug = "%s"
  description = "%[1]s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_region.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_region.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_region.test", "parent_region_id", "0"),
					resource.TestCheckResourceAttr("netbox_region.test", "description", testName),
				),
			},
			{
				ResourceName:      "netbox_region.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxRegion_defaultSlug(t *testing.T) {

	testSlug := "region_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_region" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_region.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_region.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_region", &resource.Sweeper{
		Name:         "netbox_region",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimRegionsListParams()
			res, err := api.Dcim.DcimRegionsList(params, nil)
			if err != nil {
				return err
			}
			for _, region := range res.GetPayload().Results {
				if strings.HasPrefix(*region.Name, testPrefix) {
					deleteParams := dcim.NewDcimRegionsDeleteParams().WithID(region.ID)
					_, err := api.Dcim.DcimRegionsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a Region")
				}
			}
			return nil
		},
	})
}
