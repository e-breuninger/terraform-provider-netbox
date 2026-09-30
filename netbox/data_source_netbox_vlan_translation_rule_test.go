package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVlanTranslationRuleDataSource_basic(t *testing.T) {
	testSlug := "vlan_translation_rule_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxVlanTranslationRuleDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan_translation_rule" "test" {
  policy_id   = netbox_vlan_translation_policy.test.id
  local_vid   = 100
  remote_vid  = 200
  description = "%[1]s"
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
data "netbox_vlan_translation_rule" "test" {
  depends_on = []

  policy_id = netbox_vlan_translation_policy.test.id
  local_vid = netbox_vlan_translation_rule.test.local_vid
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_vlan_translation_rule.test", "id", "netbox_vlan_translation_rule.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_translation_rule.test", "policy_id", "netbox_vlan_translation_policy.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_rule.test", "local_vid", "100"),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_rule.test", "remote_vid", "200"),
					resource.TestCheckResourceAttr("data.netbox_vlan_translation_rule.test", "description", testName),
				),
			},
		},
	})
}
