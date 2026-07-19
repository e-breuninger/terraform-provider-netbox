package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCircuitProviderNetworkDataSource_basic(t *testing.T) {
	testSlug := "circuit_provider_network_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxCircuitProviderNetworkDataSourceDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxCircuitProviderNetworkDataSource(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_network.test", "comments", "Circuit Provider Comments"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_network.test", "description", "This is my circuit provider"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_network.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_network.test", "service_id", "This is my Service ID"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_provider_network.test", "provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_network.test", "tags.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_network.test", "tags.0", testName+"-tag-1"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_network.test", "tags.1", testName+"-tag-2"),
				),
			},
		},
	})
}

func testAccNetboxCircuitProviderNetworkDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
	name	= "%[1]s-0"
}

resource "netbox_tag" "test0" {
	name	= "%[1]s-tag-1"
}

resource "netbox_tag" "test1" {
	name	= "%[1]s-tag-2"
}

resource "netbox_circuit_provider_network" "test" {
	comments 		= "Circuit Provider Comments"
	description		= "This is my circuit provider"
	name 			= "%[1]s" 
	provider_id 	= netbox_circuit_provider.test.id
	service_id 		= "This is my Service ID"
	tags			= [ netbox_tag.test0.name, netbox_tag.test1.name ]
}
`, testName)
}

func testAccNetboxCircuitProviderNetworkDataSource(testName string) string {
	return fmt.Sprintf(`
data "netbox_circuit_provider_network" "test" {
	depends_on	=	[]

	name		= "%[1]s"
}
`, testName)
}
