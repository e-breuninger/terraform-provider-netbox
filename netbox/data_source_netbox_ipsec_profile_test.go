package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpsecProfileDataSource_basic(t *testing.T) {
	testSlug := "ipsec_profile_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxIpsecProfileDependencies(testName) + fmt.Sprintf(`
resource "netbox_ipsec_profile" "test" {
  name             = "%[1]s"
  description      = "This is a description."
  mode             = "esp"
  ike_policy_id    = netbox_ike_policy.test.id
  ipsec_policy_id  = netbox_ipsec_policy.test.id
  comments         = "some comments"
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
data "netbox_ipsec_profile" "test" {
  depends_on = []

  name = netbox_ipsec_profile.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_profile.test", "id", "netbox_ipsec_profile.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_profile.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_ipsec_profile.test", "description", "This is a description."),
					resource.TestCheckResourceAttr("data.netbox_ipsec_profile.test", "mode", "esp"),
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_profile.test", "ike_policy_id", "netbox_ike_policy.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ipsec_profile.test", "ipsec_policy_id", "netbox_ipsec_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ipsec_profile.test", "comments", "some comments"),
				),
			},
		},
	})
}
