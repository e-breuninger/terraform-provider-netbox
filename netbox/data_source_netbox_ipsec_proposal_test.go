package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpsecProposalDataSource_basic(t *testing.T) {
	testSlug := "ipsec_proposal_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_ipsec_proposal" "test" {
  name                      = "%[1]s"
  description                = "This is a description."
  encryption_algorithm       = "aes-256-gcm"
  authentication_algorithm   = "hmac-sha256"
  sa_lifetime_seconds        = 3600
  sa_lifetime_data           = 1000000
  comments                    = "some comments"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_ipsec_proposal" "test" {
  depends_on = []

  name = netbox_ipsec_proposal.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_proposal.test", "id", "netbox_ipsec_proposal.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposal.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposal.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposal.test", "encryption_algorithm", "aes-256-gcm"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposal.test", "sa_lifetime_seconds", "3600"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposal.test", "sa_lifetime_data", "1000000"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_proposal.test", "comments", "some comments"),
				),
			},
		},
	})
}
