package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCircuitProviderNetworksDataSource_basic(t *testing.T) {
	testSlug := "circuit_provider_networks_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxCircuitProviderNetworksDataSourceDependencies(testName, testSlug)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxCircuitProviderNetworksDataSourceFilterName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "provider_networks.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "provider_networks.0.comments", testName+"_comments-a"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "provider_networks.0.description", testName+"_description-a"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "provider_networks.0.name", testName+"_0"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_provider_networks.test", "provider_networks.0.provider_id", "netbox_circuit_provider.providera", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "provider_networks.0.service_id", testName+"_service-id-a"),
				),
			},
			{
				Config: dependencies + testAccNetboxCircuitProviderNetworksDataSourceLimit,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "provider_networks.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_provider_networks.test", "provider_networks.0.provider_id", "netbox_circuit_provider.providera", "id"),
				),
			},
			{
				Config: dependencies + testAccNetboxCircuitProviderNetworksDataSourceNameRegex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.test", "provider_networks.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_provider_networks.test", "provider_networks.0.name", "netbox_circuit_provider_network.providernetworkc", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_provider_networks.test", "provider_networks.1.name", "netbox_circuit_provider_network.providernetworkd", "name"),
				),
			},
		},
	})
}

func TestAccNetboxCircuitProviderNetworksDataSource_tags(t *testing.T) {
	testSlug := "circuit_provider_networks_tags"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxCircuitProviderNetworksDataSourceDependenciesWithTags(testName, testSlug)

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxCircuitProviderNetworksDataSourceTagA(testName) + testAccNetboxCircuitProviderNetworksDataSourceTagB(testName) + testAccNetboxCircuitProviderNetworksDataSourceTagAB(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.tag-a", "provider_networks.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.tag-b", "provider_networks.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_circuit_provider_networks.tag-ab", "provider_networks.#", "1"),
				),
			},
		},
	})
}

func testAccNetboxCircuitProviderNetworksDataSourceDependencies(testName string, testSlug string) string {
	return fmt.Sprintf(`
resource "netbox_circuit_provider" "providera" {
	name		= "%[1]s_provider-a"
	slug		= "%[2]s_provider-a"
}

resource "netbox_circuit_provider" "providerb" {
	name		= "%[1]s_provider-b"
	slug		= "%[2]s_provider-b"
}

resource "netbox_circuit_provider_network" "providernetworka" {
	comments	= "%[1]s_comments-a"
	description	= "%[1]s_description-a"
	name		= "%[1]s_0"
	provider_id	= netbox_circuit_provider.providera.id
	service_id	= "%[1]s_service-id-a"
}

resource "netbox_circuit_provider_network" "providernetworkb" {
	comments	= "%[1]s_comments-b"
	description	= "%[1]s_description-b"
	name		= "%[1]s_1"
	provider_id	= netbox_circuit_provider.providerb.id
	service_id	= "%[1]s_service-id-b"
}

resource "netbox_circuit_provider_network" "providernetworkc" {
	comments	= "%[1]s_comments-c"
	description	= "%[1]s_description-c"
	name		= "%[1]s_2_regex"
	provider_id	= netbox_circuit_provider.providera.id
	service_id	= "%[1]s_service-id-c"
}

resource "netbox_circuit_provider_network" "providernetworkd" {
	comments	= "%[1]s_comment-d"
	description	= "%[1]s_description-d"
	name		= "%[1]s_3_regex"
	provider_id	= netbox_circuit_provider.providerb.id
	service_id	= "%[1]s_service-id-d"
}
`, testName, testSlug)
}

func testAccNetboxCircuitProviderNetworksDataSourceDependenciesWithTags(testName string, testSlug string) string {
	return fmt.Sprintf(`
resource "netbox_circuit_provider" "providera" {
	name		= "%[1]s_provider-a"
	slug		= "%[2]s_provider-a"
}

resource "netbox_circuit_provider" "providerb" {
	name		= "%[1]s_provider-b"
	slug		= "%[2]s_provider-b"
}

resource "netbox_tag" "tag-a" {
	name		= "%[1]s_service-a"
}

resource "netbox_tag" "tag-b" {
	name		= "%[1]s_service-b"
}

resource "netbox_circuit_provider_network" "providernetworka" {
	comments	= "%[1]s_comments-a"
	description	= "%[1]s_description-a"
	name		= "%[1]s_0"
	provider_id	= netbox_circuit_provider.providera.id
	service_id	= "%[1]s_service-id-a"
	tags		= [
		netbox_tag.tag-a.name
	]
}

resource "netbox_circuit_provider_network" "providernetworkb" {
	comments	= "%[1]s_comments-b"
	description	= "%[1]s_description-b"
	name		= "%[1]s_1"
	provider_id	= netbox_circuit_provider.providerb.id
	service_id	= "%[1]s_service-id-b"
	tags		= [
		netbox_tag.tag-b.name
	]
}

resource "netbox_circuit_provider_network" "providernetworkc" {
	comments	= "%[1]s_comments-c"
	description	= "%[1]s_description-c"
	name		= "%[1]s_2"
	provider_id	= netbox_circuit_provider.providera.id
	service_id	= "%[1]s_service-id-c"
	tags		= [
		netbox_tag.tag-a.name,
		netbox_tag.tag-b.name
	]
}
`, testName, testSlug)
}

func testAccNetboxCircuitProviderNetworksDataSourceFilterName(testName string) string {
	return fmt.Sprintf(`
data "netbox_circuit_provider_networks" "test" {
	filter {
		name 	= "name"
		value	= "%[1]s_0"
	}
}
`, testName)
}

const testAccNetboxCircuitProviderNetworksDataSourceNameRegex = `
data "netbox_circuit_provider_networks" "test" {
	name_regex = "test.*_regex"
}
`

const testAccNetboxCircuitProviderNetworksDataSourceLimit = `
data "netbox_circuit_provider_networks" "test" {
	limit 		= 1
	filter {
		name	= "provider_id"
		value	= netbox_circuit_provider.providera.id
	}
}
`

func testAccNetboxCircuitProviderNetworksDataSourceTagA(testName string) string {
	return fmt.Sprintf(`
data "netbox_circuit_provider_networks" "tag-a" {
	filter {
		name	= "tag"
		value	= "%[1]s_service-a"
	}
}
`, testName)
}

func testAccNetboxCircuitProviderNetworksDataSourceTagB(testName string) string {
	return fmt.Sprintf(`
data "netbox_circuit_provider_networks" "tag-b" {
	filter {
		name	= "tag"
		value	= "%[1]s_service-b"
	}
}
`, testName)
}

func testAccNetboxCircuitProviderNetworksDataSourceTagAB(testName string) string {
	return fmt.Sprintf(`
data "netbox_circuit_provider_networks" "tag-ab" {
	filter {
		name	= "tag"
		value	= "%[1]s_service-a"
	}
	filter {
		name	= "tag"
		value	= "%[1]s_service-b"
	}
}
`, testName)
}
