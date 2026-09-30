package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVlanTranslationRuleDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_vlan_translation_policy" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxVlanTranslationRule_basic(t *testing.T) {
	testSlug := "vlan_translation_rule"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVlanTranslationRuleDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan_translation_rule" "test" {
  policy_id   = netbox_vlan_translation_policy.test.id
  local_vid   = 100
  remote_vid  = 200
  description = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_vlan_translation_rule.test", "policy_id", "netbox_vlan_translation_policy.test", "id"),
					resource.TestCheckResourceAttr("netbox_vlan_translation_rule.test", "local_vid", "100"),
					resource.TestCheckResourceAttr("netbox_vlan_translation_rule.test", "remote_vid", "200"),
					resource.TestCheckResourceAttr("netbox_vlan_translation_rule.test", "description", testName),
				),
			},
			{
				ResourceName:      "netbox_vlan_translation_rule.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_vlan_translation_rule", &resource.Sweeper{
		Name:         "netbox_vlan_translation_rule",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamVlanTranslationRulesListParams()
			res, err := api.Ipam.IpamVlanTranslationRulesList(params, nil)
			if err != nil {
				return err
			}
			for _, rule := range res.GetPayload().Results {
				if strings.HasPrefix(rule.Description, testPrefix) {
					deleteParams := ipam.NewIpamVlanTranslationRulesDeleteParams().WithID(rule.ID)
					_, err := api.Ipam.IpamVlanTranslationRulesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a vlan translation rule")
				}
			}
			return nil
		},
	})
}
