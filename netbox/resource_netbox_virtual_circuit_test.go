package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVirtualCircuitDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
  slug = "%[1]s"
}

resource "netbox_circuit_provider_network" "test" {
  name        = "%[1]s"
  provider_id = netbox_circuit_provider.test.id
}

resource "netbox_virtual_circuit_type" "test" {
  name = "%[1]s"
  slug = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
  slug = "%[1]s"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxVirtualCircuit_basic(t *testing.T) {
	testSlug := "virtual_circuit"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualCircuitDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_circuit" "test" {
  cid                 = "%[1]s"
  provider_network_id = netbox_circuit_provider_network.test.id
  type_id             = netbox_virtual_circuit_type.test.id
  status              = "active"
  tenant_id           = netbox_tenant.test.id
  description         = "This is my virtual circuit"
  comments             = "This is my virtual circuit comment"
  tags                = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "cid", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_circuit.test", "provider_network_id", "netbox_circuit_provider_network.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_circuit.test", "type_id", "netbox_virtual_circuit_type.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_virtual_circuit.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "description", "This is my virtual circuit"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "comments", "This is my virtual circuit comment"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_virtual_circuit.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_virtual_circuit", &resource.Sweeper{
		Name:         "netbox_virtual_circuit",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := circuits.NewCircuitsVirtualCircuitsListParams()
			res, err := api.Circuits.CircuitsVirtualCircuitsList(params, nil)
			if err != nil {
				return err
			}
			for _, virtualCircuit := range res.GetPayload().Results {
				if strings.HasPrefix(*virtualCircuit.Cid, testPrefix) {
					deleteParams := circuits.NewCircuitsVirtualCircuitsDeleteParams().WithID(virtualCircuit.ID)
					_, err := api.Circuits.CircuitsVirtualCircuitsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a virtual circuit")
				}
			}
			return nil
		},
	})
}
