package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxIkePolicyDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                     = "%[1]s"
  authentication_method    = "preshared-keys"
  encryption_algorithm     = "aes-256-gcm"
  group                    = 14
}
`, testName)
}

func TestAccNetboxIkePolicy_basic(t *testing.T) {
	testSlug := "ike_policy"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIkePolicyDependencies(testName) + fmt.Sprintf(`
resource "netbox_ike_policy" "test" {
  name           = "%[1]s"
  description    = "This is a description."
  version        = 2
  mode           = "main"
  proposal_ids   = [netbox_ike_proposal.test.id]
  preshared_key  = "supersecret"
  comments       = "some comments"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "version", "2"),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "mode", "main"),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "proposal_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "preshared_key", "supersecret"),
					resource.TestCheckResourceAttr("netbox_ike_policy.test", "comments", "some comments"),
				),
			},
			{
				ResourceName:            "netbox_ike_policy.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"preshared_key"},
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ike_policy", &resource.Sweeper{
		Name: "netbox_ike_policy",
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := vpn.NewVpnIkePoliciesListParams()
			res, err := api.Vpn.VpnIkePoliciesList(params, nil)
			if err != nil {
				return err
			}
			for _, ikePolicy := range res.GetPayload().Results {
				if strings.HasPrefix(*ikePolicy.Name, testPrefix) {
					deleteParams := vpn.NewVpnIkePoliciesDeleteParams().WithID(ikePolicy.ID)
					_, err := api.Vpn.VpnIkePoliciesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an ike policy")
				}
			}
			return nil
		},
	})
}
