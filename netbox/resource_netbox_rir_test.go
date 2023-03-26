package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRir_basic(t *testing.T) {

	testSlug := "rir"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rir" "test_basic" {
  name        = "%s"
  slug        = "%s"
  description = "my-description"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rir.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_rir.test_basic", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_rir.test_basic", "description", "my-description"),
				),
			},
			{
				ResourceName:      "netbox_rir.test_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_rir", &resource.Sweeper{
		Name:         "netbox_rir",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamRirsListParams()
			res, err := api.Ipam.IpamRirsList(params, nil)
			if err != nil {
				return err
			}
			for _, role := range res.GetPayload().Results {
				if strings.HasPrefix(*role.Name, testPrefix) {
					deleteParams := ipam.NewIpamRirsDeleteParams().WithID(role.ID)
					_, err := api.Ipam.IpamRirsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a rir")
				}
			}
			return nil
		},
	})
}
