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

func TestAccNetboxCircuitType_basic(t *testing.T) {

	testSlug := "circuit_type"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_circuit_type" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_circuit_type.test", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_circuit_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_circuit_type", &resource.Sweeper{
		Name:         "netbox_circuit_type",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := circuits.NewCircuitsCircuitTypesListParams()
			res, err := api.Circuits.CircuitsCircuitTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, CircuitType := range res.GetPayload().Results {
				if strings.HasPrefix(*CircuitType.Name, testPrefix) {
					deleteParams := circuits.NewCircuitsCircuitTypesDeleteParams().WithID(CircuitType.ID)
					_, err := api.Circuits.CircuitsCircuitTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a circuit type")
				}
			}
			return nil
		},
	})
}
