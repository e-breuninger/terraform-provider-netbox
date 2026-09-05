package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIkeProposalDataSource_basic(t *testing.T) {
	testSlug := "ike_proposal_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_ike_proposal" "test" {
  name                      = "%[1]s"
  description                = "This is a description."
  authentication_method      = "preshared-keys"
  encryption_algorithm       = "aes-256-gcm"
  authentication_algorithm   = "hmac-sha256"
  group                       = 14
  sa_lifetime                = 3600
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
data "netbox_ike_proposal" "test" {
  depends_on = []

  name = netbox_ike_proposal.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ike_proposal.test", "id", "netbox_ike_proposal.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "authentication_method", "preshared-keys"),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "encryption_algorithm", "aes-256-gcm"),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "authentication_algorithm", "hmac-sha256"),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "group", "14"),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "sa_lifetime", "3600"),
					resource.TestCheckResourceAttr("data.netbox_ike_proposal.test", "comments", "some comments"),
				),
			},
		},
	})
}
