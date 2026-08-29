package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxProviderAccountDataSource_basic(t *testing.T) {
	testSlug := "provider_account_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxProviderAccountDependencies(testName, testSlug) + fmt.Sprintf(`
resource "netbox_provider_account" "test" {
	account 		= "%[1]s"
	comments 		= "Provider Account Comments"
	description		= "This is my provider account"
	name 			= "%[1]s"
	provider_id 	= netbox_circuit_provider.test.id
	tags			= [netbox_tag.test.name]
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
data "netbox_provider_account" "test" {
	depends_on	=	[]

	name = netbox_provider_account.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_provider_account.test", "id", "netbox_provider_account.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_provider_account.test", "account", testName),
					resource.TestCheckResourceAttr("data.netbox_provider_account.test", "comments", "Provider Account Comments"),
					resource.TestCheckResourceAttr("data.netbox_provider_account.test", "description", "This is my provider account"),
					resource.TestCheckResourceAttr("data.netbox_provider_account.test", "name", testName),
					resource.TestCheckResourceAttrPair("data.netbox_provider_account.test", "provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_provider_account.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_provider_account.test", "tags.0", testName),
				),
			},
		},
	})
}
