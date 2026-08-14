package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxProviderAccountDependencies(testName string, testSlug string) string {
	return fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
	name = "%[1]s"
	slug = "%[2]s"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName, testSlug)
}

func TestAccNetboxProviderAccount_basic(t *testing.T) {
	testSlug := "provider_account"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxProviderAccountDependencies(testName, testSlug) + fmt.Sprintf(`
resource "netbox_provider_account" "test" {
	account 		= "%[1]s"
	comments 		= "Provider Account Comments"
	description		= "This is my provider account"
	name 			= "%[1]s"
	provider_id 	= netbox_circuit_provider.test.id
	tags			= [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_provider_account.test", "account", testName),
					resource.TestCheckResourceAttr("netbox_provider_account.test", "comments", "Provider Account Comments"),
					resource.TestCheckResourceAttr("netbox_provider_account.test", "description", "This is my provider account"),
					resource.TestCheckResourceAttr("netbox_provider_account.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_provider_account.test", "provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttr("netbox_provider_account.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_provider_account.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_provider_account.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_provider_account", &resource.Sweeper{
		Name:         "netbox_provider_account",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := circuits.NewCircuitsProviderAccountsListParams()
			res, err := api.Circuits.CircuitsProviderAccountsList(params, nil)
			if err != nil {
				return err
			}
			for _, providerAccount := range res.GetPayload().Results {
				if strings.HasPrefix(providerAccount.Name, testPrefix) {
					deleteParams := circuits.NewCircuitsProviderAccountsDeleteParams().WithID(providerAccount.ID)
					_, err := api.Circuits.CircuitsProviderAccountsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a provider account")
				}
			}
			return nil
		},
	})
}
