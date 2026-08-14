package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxIpsecPolicyDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_ipsec_proposal" "test" {
  name                   = "%[1]s"
  encryption_algorithm   = "aes-256-gcm"
}
`, testName)
}

func TestAccNetboxIpsecPolicy_basic(t *testing.T) {
	testSlug := "ipsec_policy"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIpsecPolicyDependencies(testName) + fmt.Sprintf(`
resource "netbox_ipsec_policy" "test" {
  name          = "%[1]s"
  description   = "This is a description."
  proposal_ids  = [netbox_ipsec_proposal.test.id]
  pfs_group     = 14
  comments      = "some comments"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "proposal_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "pfs_group", "14"),
					resource.TestCheckResourceAttr("netbox_ipsec_policy.test", "comments", "some comments"),
				),
			},
			{
				ResourceName:      "netbox_ipsec_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ipsec_policy", &resource.Sweeper{
		Name:         "netbox_ipsec_policy",
		Dependencies: []string{"netbox_ipsec_profile"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := vpn.NewVpnIpsecPoliciesListParams()
			res, err := api.Vpn.VpnIpsecPoliciesList(params, nil)
			if err != nil {
				return err
			}
			for _, ipsecPolicy := range res.GetPayload().Results {
				if strings.HasPrefix(*ipsecPolicy.Name, testPrefix) {
					deleteParams := vpn.NewVpnIpsecPoliciesDeleteParams().WithID(ipsecPolicy.ID)
					_, err := api.Vpn.VpnIpsecPoliciesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an ipsec policy")
				}
			}
			return nil
		},
	})
}
