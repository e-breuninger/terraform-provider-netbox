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

func TestAccNetboxCircuitProvider_basic(t *testing.T) {

	testSlug := "circuit_prov"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_provider.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_circuit_provider.test", "slug", randomSlug),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}`, testName+"2", randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_provider.test", "name", testName+"2"),
					resource.TestCheckResourceAttr("netbox_circuit_provider.test", "slug", randomSlug),
				),
			},
			{
				ResourceName:      "netbox_circuit_provider.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_circuit_provider", &resource.Sweeper{
		Name:         "netbox_circuit_provider",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := circuits.NewCircuitsProvidersListParams()
			res, err := api.Circuits.CircuitsProvidersList(params, nil)
			if err != nil {
				return err
			}
			for _, CircuitProvider := range res.GetPayload().Results {
				if strings.HasPrefix(*CircuitProvider.Name, testPrefix) {
					deleteParams := circuits.NewCircuitsProvidersDeleteParams().WithID(CircuitProvider.ID)
					_, err := api.Circuits.CircuitsProvidersDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a circuit provider")
				}
			}
			return nil
		},
	})
}
