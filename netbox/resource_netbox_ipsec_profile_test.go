package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxIpsecProfileDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                    = "%[1]s"
  authentication_method   = "preshared-keys"
  encryption_algorithm    = "aes-256-gcm"
  group                   = 14
}

resource "netbox_ike_policy" "test" {
  name          = "%[1]s"
  version       = 2
  proposal_ids  = [netbox_ike_proposal.test.id]
}

resource "netbox_ipsec_proposal" "test" {
  name                   = "%[1]s"
  encryption_algorithm   = "aes-256-gcm"
}

resource "netbox_ipsec_policy" "test" {
  name          = "%[1]s"
  proposal_ids  = [netbox_ipsec_proposal.test.id]
}
`, testName)
}

func TestAccNetboxIpsecProfile_basic(t *testing.T) {
	testSlug := "ipsec_profile"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIpsecProfileDependencies(testName) + fmt.Sprintf(`
resource "netbox_ipsec_profile" "test" {
  name             = "%[1]s"
  description      = "This is a description."
  mode             = "esp"
  ike_policy_id    = netbox_ike_policy.test.id
  ipsec_policy_id  = netbox_ipsec_policy.test.id
  comments         = "some comments"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_profile.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipsec_profile.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("netbox_ipsec_profile.test", "mode", "esp"),
					resource.TestCheckResourceAttrPair("netbox_ipsec_profile.test", "ike_policy_id", "netbox_ike_policy.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ipsec_profile.test", "ipsec_policy_id", "netbox_ipsec_policy.test", "id"),
					resource.TestCheckResourceAttr("netbox_ipsec_profile.test", "comments", "some comments"),
				),
			},
			{
				ResourceName:      "netbox_ipsec_profile.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ipsec_profile", &resource.Sweeper{
		Name: "netbox_ipsec_profile",
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := vpn.NewVpnIpsecProfilesListParams()
			res, err := api.Vpn.VpnIpsecProfilesList(params, nil)
			if err != nil {
				return err
			}
			for _, ipsecProfile := range res.GetPayload().Results {
				if strings.HasPrefix(*ipsecProfile.Name, testPrefix) {
					deleteParams := vpn.NewVpnIpsecProfilesDeleteParams().WithID(ipsecProfile.ID)
					_, err := api.Vpn.VpnIpsecProfilesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an ipsec profile")
				}
			}
			return nil
		},
	})
}
