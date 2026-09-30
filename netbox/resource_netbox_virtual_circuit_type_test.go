package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualCircuitType_basic(t *testing.T) {
	testSlug := "virtual_circuit_type"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_virtual_circuit_type" "test" {
  name        = "%[1]s"
  slug        = "%[1]s"
  color_hex   = "123456"
  description = "This is my virtual circuit type"
  comments    = "This is my virtual circuit type comment"
  tags        = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "slug", testName),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "color_hex", "123456"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "description", "This is my virtual circuit type"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "comments", "This is my virtual circuit type comment"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_virtual_circuit_type.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_virtual_circuit_type.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_virtual_circuit_type", &resource.Sweeper{
		Name: "netbox_virtual_circuit_type",
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := circuits.NewCircuitsVirtualCircuitTypesListParams()
			res, err := api.Circuits.CircuitsVirtualCircuitTypesList(params, nil)
			if err != nil {
				return err
			}
			for _, virtualCircuitType := range res.GetPayload().Results {
				if strings.HasPrefix(*virtualCircuitType.Name, testPrefix) {
					deleteParams := circuits.NewCircuitsVirtualCircuitTypesDeleteParams().WithID(virtualCircuitType.ID)
					_, err := api.Circuits.CircuitsVirtualCircuitTypesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a virtual circuit type")
				}
			}
			return nil
		},
	})
}
