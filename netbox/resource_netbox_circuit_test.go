package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxCircuitDependencies(testName string, testSlug string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
	name = "%[1]s"
	slug = "%[2]s"
}
resource "netbox_circuit_provider" "test" {
	name = "%[1]s"
	slug = "%[2]s"
}
resource "netbox_circuit_type" "test" {
	name = "%[1]s"
	slug = "%[2]s"
}
`, testName, testSlug)
}
func TestAccNetboxCircuit_basic(t *testing.T) {
	testSlug := "circuit_prov"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxCircuitDependencies(testName, randomSlug) + fmt.Sprintf(`
resource "netbox_circuit" "test" {
  cid = "%[1]s"
  status = "active"
  provider_id = netbox_circuit_provider.test.id
  type_id = netbox_circuit_type.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit.test", "cid", testName),
					resource.TestCheckResourceAttrPair("netbox_circuit.test", "provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_circuit.test", "type_id", "netbox_circuit_type.test", "id"),
				),
			},
			{
				Config: testAccNetboxCircuitDependencies(testName, randomSlug) + fmt.Sprintf(`
resource "netbox_circuit" "test" {
  cid = "%[1]s"
  status = "active"
  provider_id = netbox_circuit_provider.test.id
  type_id = netbox_circuit_type.test.id
  tenant_id = netbox_tenant.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit.test", "cid", testName),
					resource.TestCheckResourceAttrPair("netbox_circuit.test", "provider_id", "netbox_circuit_provider.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_circuit.test", "type_id", "netbox_circuit_type.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_circuit.test", "tenant_id", "netbox_tenant.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_circuit.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_circuit", &resource.Sweeper{
		Name:         "netbox_circuit",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := circuits.NewCircuitsCircuitsListParams()
			res, err := api.Circuits.CircuitsCircuitsList(params, nil)
			if err != nil {
				return err
			}
			for _, Circuit := range res.GetPayload().Results {
				if strings.HasPrefix(*Circuit.Cid, testPrefix) {
					deleteParams := circuits.NewCircuitsCircuitsDeleteParams().WithID(Circuit.ID)
					_, err := api.Circuits.CircuitsCircuitsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a circuit")
				}
			}
			return nil
		},
	})
}
