package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/vpn"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpsecProposal_basic(t *testing.T) {
	testSlug := "ipsec_proposal"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ipsec_proposal" "test" {
  name                      = "%[1]s"
  description                = "This is a description."
  encryption_algorithm       = "aes-256-gcm"
  authentication_algorithm   = "hmac-sha256"
  sa_lifetime_seconds        = 3600
  sa_lifetime_data           = 1000000
  comments                    = "some comments"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "encryption_algorithm", "aes-256-gcm"),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "sa_lifetime_seconds", "3600"),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "sa_lifetime_data", "1000000"),
					resource.TestCheckResourceAttr("netbox_ipsec_proposal.test", "comments", "some comments"),
				),
			},
			{
				ResourceName:      "netbox_ipsec_proposal.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ipsec_proposal", &resource.Sweeper{
		Name: "netbox_ipsec_proposal",
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := vpn.NewVpnIpsecProposalsListParams()
			res, err := api.Vpn.VpnIpsecProposalsList(params, nil)
			if err != nil {
				return err
			}
			for _, ipsecProposal := range res.GetPayload().Results {
				if strings.HasPrefix(*ipsecProposal.Name, testPrefix) {
					deleteParams := vpn.NewVpnIpsecProposalsDeleteParams().WithID(ipsecProposal.ID)
					_, err := api.Vpn.VpnIpsecProposalsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an ipsec proposal")
				}
			}
			return nil
		},
	})
}
