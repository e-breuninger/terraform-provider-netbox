package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxCircuitGroupDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
	name = "%[1]s"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxCircuitGroup_basic(t *testing.T) {
	testSlug := "circuit_group"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxCircuitGroupDependencies(testName) + fmt.Sprintf(`
resource "netbox_circuit_group" "test" {
	name 			= "%[1]s"
	slug 			= "%[2]s"
	comments 		= "Circuit Group Comments"
	description		= "This is my circuit group"
	tenant_id 		= netbox_tenant.test.id
	tags			= [netbox_tag.test.name]
}`, testName, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "slug", testSlug),
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "comments", "Circuit Group Comments"),
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "description", "This is my circuit group"),
					resource.TestCheckResourceAttrPair("netbox_circuit_group.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_circuit_group.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_circuit_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_circuit_group", &resource.Sweeper{
		Name:         "netbox_circuit_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := circuits.NewCircuitsCircuitGroupsListParams()
			res, err := api.Circuits.CircuitsCircuitGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, circuitGroup := range res.GetPayload().Results {
				if circuitGroup.Name != nil && strings.HasPrefix(*circuitGroup.Name, testPrefix) {
					deleteParams := circuits.NewCircuitsCircuitGroupsDeleteParams().WithID(circuitGroup.ID)
					_, err := api.Circuits.CircuitsCircuitGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a circuit group")
				}
			}
			return nil
		},
	})
}
