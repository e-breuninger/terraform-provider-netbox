package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpsecPolicyDataSource_basic(t *testing.T) {
	testSlug := "ipsec_policy_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxIpsecPolicyDependencies(testName) + fmt.Sprintf(`
resource "netbox_ipsec_policy" "test" {
  name          = "%[1]s"
  description   = "This is a description."
  proposal_ids  = [netbox_ipsec_proposal.test.id]
  pfs_group     = 14
  comments      = "some comments"
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
data "netbox_ipsec_policy" "test" {
  depends_on = []

  name = netbox_ipsec_policy.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_policy.test", "id", "netbox_ipsec_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_policy.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_ipsec_policy.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("data.netbox_ipsec_policy.test", "proposal_ids.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_policy.test", "pfs_group", "14"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_policy.test", "comments", "some comments"),
				),
			},
		},
	})
}
