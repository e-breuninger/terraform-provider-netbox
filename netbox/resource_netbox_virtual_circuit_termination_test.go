package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVirtualCircuitTerminationDependencies(testName string) string {
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

resource "netbox_virtual_circuit" "test" {
  cid                 = "%[1]s"
  provider_network_id = netbox_circuit_provider_network.test.id
  type_id             = netbox_virtual_circuit_type.test.id
  status              = "active"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name            = "%[1]s"
  device_type_id  = netbox_device_type.test.id
  role_id         = netbox_device_role.test.id
  site_id         = netbox_site.test.id
}

resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s"
  type      = "virtual"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxVirtualCircuitTermination_basic(t *testing.T) {
	testSlug := "virtual_circuit_term"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualCircuitTerminationDependencies(testName) + `
resource "netbox_virtual_circuit_termination" "test" {
  virtual_circuit_id = netbox_virtual_circuit.test.id
  interface_id        = netbox_device_interface.test.id
  role                = "peer"
  description         = "This is my virtual circuit termination"
  tags                = [netbox_tag.test.name]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_circuit_termination.test", "virtual_circuit_id", "netbox_virtual_circuit.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_circuit_termination.test", "interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_termination.test", "role", "peer"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_termination.test", "description", "This is my virtual circuit termination"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_termination.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_termination.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_virtual_circuit_termination.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_virtual_circuit_termination", &resource.Sweeper{
		Name:         "netbox_virtual_circuit_termination",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := circuits.NewCircuitsVirtualCircuitTerminationsListParams()
			res, err := api.Circuits.CircuitsVirtualCircuitTerminationsList(params, nil)
			if err != nil {
				return err
			}
			for _, term := range res.GetPayload().Results {
				if term.VirtualCircuit != nil && term.VirtualCircuit.Cid != nil && strings.HasPrefix(*term.VirtualCircuit.Cid, testPrefix) {
					deleteParams := circuits.NewCircuitsVirtualCircuitTerminationsDeleteParams().WithID(term.ID)
					_, err := api.Circuits.CircuitsVirtualCircuitTerminationsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a virtual circuit termination")
				}
			}
			return nil
		},
	})
}
