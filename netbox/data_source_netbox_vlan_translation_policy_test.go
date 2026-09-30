package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVlanTranslationPolicyDataSource_basic(t *testing.T) {
	testSlug := "vlan_translation_policy_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_vlan_translation_policy" "test" {
  name        = "%[1]s"
  description = "vlan translation policy for acctest"
  comments    = "some comments"
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
data "netbox_vlan_translation_policy" "test" {
  depends_on = []

  name = netbox_vlan_translation_policy.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan_translation_policy.test", "id", "netbox_vlan_translation_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_policy.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_policy.test", "description", "vlan translation policy for acctest"),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_policy.test", "comments", "some comments"),
				),
			},
		},
	})
}
