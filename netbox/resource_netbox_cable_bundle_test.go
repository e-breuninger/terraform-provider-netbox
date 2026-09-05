package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCableBundle_basic(t *testing.T) {
	testSlug := "cable_bundle"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_cable_bundle" "test" {
	name        = "%[1]s"
	description = "This is my cable bundle"
	comments    = "Cable Bundle Comments"
	tags        = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_cable_bundle.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_cable_bundle.test", "description", "This is my cable bundle"),
					resource.TestCheckResourceAttr("netbox_cable_bundle.test", "comments", "Cable Bundle Comments"),
					resource.TestCheckResourceAttr("netbox_cable_bundle.test", "cable_count", "0"),
					resource.TestCheckResourceAttr("netbox_cable_bundle.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_cable_bundle.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_cable_bundle.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_cable_bundle", &resource.Sweeper{
		Name:         "netbox_cable_bundle",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimCableBundlesListParams()
			res, err := api.Dcim.DcimCableBundlesList(params, nil)
			if err != nil {
				return err
			}
			for _, bundle := range res.GetPayload().Results {
				if strings.HasPrefix(*bundle.Name, testPrefix) {
					deleteParams := dcim.NewDcimCableBundlesDeleteParams().WithID(bundle.ID)
					_, err := api.Dcim.DcimCableBundlesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a cable bundle")
				}
			}
			return nil
		},
	})
}
