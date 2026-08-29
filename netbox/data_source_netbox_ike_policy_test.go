package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIkePolicyDataSource_basic(t *testing.T) {
	testSlug := "ike_policy_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxIkePolicyDependencies(testName) + fmt.Sprintf(`
resource "netbox_ike_policy" "test" {
  name           = "%[1]s"
  description    = "This is a description."
  version        = 1
  mode           = "main"
  proposal_ids   = [netbox_ike_proposal.test.id]
  preshared_key  = "supersecret"
  comments       = "some comments"
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
data "netbox_ike_policy" "test" {
  depends_on = []

  name = netbox_ike_policy.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ike_policy.test", "id", "netbox_ike_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ike_policy.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_ike_policy.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("data.netbox_ike_policy.test", "version", "1"),
					resource.TestCheckResourceAttr("data.netbox_ike_policy.test", "mode", "main"),
					resource.TestCheckResourceAttr("data.netbox_ike_policy.test", "proposal_ids.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_ike_policy.test", "preshared_key", "supersecret"),
					resource.TestCheckResourceAttr("data.netbox_ike_policy.test", "comments", "some comments"),
				),
			},
		},
	})
}
