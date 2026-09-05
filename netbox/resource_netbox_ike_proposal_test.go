package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIkeProposal_basic(t *testing.T) {
	testSlug := "ike_proposal"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                      = "%[1]s"
  description                = "This is a description."
  authentication_method      = "preshared-keys"
  encryption_algorithm       = "aes-256-gcm"
  authentication_algorithm   = "hmac-sha256"
  group                       = 14
  sa_lifetime                = 3600
  comments                    = "some comments"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "authentication_method", "preshared-keys"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "encryption_algorithm", "aes-256-gcm"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "group", "14"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "sa_lifetime", "3600"),
					resource.TestCheckResourceAttr("netbox_ike_proposal.test", "comments", "some comments"),
				),
			},
			{
				ResourceName:      "netbox_ike_proposal.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ike_proposal", &resource.Sweeper{
		Name: "netbox_ike_proposal",
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := vpn.NewVpnIkeProposalsListParams()
			res, err := api.Vpn.VpnIkeProposalsList(params, nil)
			if err != nil {
				return err
			}
			for _, ikeProposal := range res.GetPayload().Results {
				if strings.HasPrefix(*ikeProposal.Name, testPrefix) {
					deleteParams := vpn.NewVpnIkeProposalsDeleteParams().WithID(ikeProposal.ID)
					_, err := api.Vpn.VpnIkeProposalsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an ike proposal")
				}
			}
			return nil
		},
	})
}
