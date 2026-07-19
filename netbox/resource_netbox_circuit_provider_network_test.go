package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxCircuitProviderNetworkDependencies(testName string, testSlug string) string {
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

func TestAccNetboxCircuitProviderNetwork_basic(t *testing.T) {
	testSlug := "circuit_prov_network"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxCircuitProviderNetworkDependencies(testName, testSlug) + fmt.Sprintf(`
resource "netbox_circuit_provider_network" "test" {
	comments 		= "Circuit Provider Comments"
	description		= "This is my circuit provider"
	name 			= "%[1]s" 
	provider_id 	= netbox_circuit_provider.test.id
	service_id 		= "This is my Service ID"
	tags			= [netbox_tag.test.name]
}`, testName, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "comments", "Circuit Provider Comments"),
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "description", "This is my circuit provider"),
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_circuit_provider_network.test", "provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "service_id", "This is my Service ID"),
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_circuit_provider_network.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_circuit_provider_network.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_circuit_provider_network", &resource.Sweeper{
		Name:         "netbox_circuit_provider_network",
		Dependencies: []string{"network_circuit_provider"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := circuits.NewCircuitsProviderNetworksListParams()
			res, err := api.Circuits.CircuitsProviderNetworksList(params, nil)
			if err != nil {
				return err
			}
			for _, CircuitProviderNetwork := range res.GetPayload().Results {
				if strings.HasPrefix(*CircuitProviderNetwork.Name, testPrefix) {
					deleteParams := circuits.NewCircuitsProviderNetworksDeleteParams().WithID(CircuitProviderNetwork.ID)
					_, err := api.Circuits.CircuitsProviderNetworksDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a circuit provider network")
				}
			}
			return nil
		},
	})
}
