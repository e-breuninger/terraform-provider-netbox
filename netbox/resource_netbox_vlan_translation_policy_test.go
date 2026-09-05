package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVlanTranslationPolicy_basic(t *testing.T) {
	testSlug := "vlan_translation_policy"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_vlan_translation_policy" "test" {
  name        = "%[1]s"
  description = "vlan translation policy for acctest"
  comments    = "some comments"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_translation_policy.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan_translation_policy.test", "description", "vlan translation policy for acctest"),
					resource.TestCheckResourceAttr("netbox_vlan_translation_policy.test", "comments", "some comments"),
				),
			},
			{
				ResourceName:      "netbox_vlan_translation_policy.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_vlan_translation_policy", &resource.Sweeper{
		Name:         "netbox_vlan_translation_policy",
		Dependencies: []string{"netbox_vlan_translation_rule"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamVlanTranslationPoliciesListParams()
			res, err := api.Ipam.IpamVlanTranslationPoliciesList(params, nil)
			if err != nil {
				return err
			}
			for _, policy := range res.GetPayload().Results {
				if strings.HasPrefix(*policy.Name, testPrefix) {
					deleteParams := ipam.NewIpamVlanTranslationPoliciesDeleteParams().WithID(policy.ID)
					_, err := api.Ipam.IpamVlanTranslationPoliciesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a vlan translation policy")
				}
			}
			return nil
		},
	})
}
