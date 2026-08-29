package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/circuits"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxCircuitGroupAssignmentDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_circuit_provider" "test" {
	name = "%[1]s"
	slug = "%[1]s"
}
resource "netbox_circuit_type" "test" {
	name = "%[1]s"
	slug = "%[1]s"
}
resource "netbox_circuit" "test" {
  cid         = "%[1]s"
  status      = "active"
  provider_id = netbox_circuit_provider.test.id
  type_id     = netbox_circuit_type.test.id
}
resource "netbox_circuit_group" "test" {
	name = "%[1]s"
	slug = "%[1]s"
}
resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxCircuitGroupAssignment_basic(t *testing.T) {
	testSlug := "circuit_group_assignment"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxCircuitGroupAssignmentDependencies(testName) + `
resource "netbox_circuit_group_assignment" "test" {
	group_id    = netbox_circuit_group.test.id
	member_type = "circuits.circuit"
	member_id   = netbox_circuit.test.id
	priority    = "primary"
	tags        = [netbox_tag.test.name]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_circuit_group_assignment.test", "group_id", "netbox_circuit_group.test", "id"),
					resource.TestCheckResourceAttr("netbox_circuit_group_assignment.test", "member_type", "circuits.circuit"),
					resource.TestCheckResourceAttrPair("netbox_circuit_group_assignment.test", "member_id", "netbox_circuit.test", "id"),
					resource.TestCheckResourceAttr("netbox_circuit_group_assignment.test", "priority", "primary"),
					resource.TestCheckResourceAttr("netbox_circuit_group_assignment.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_circuit_group_assignment.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_circuit_group_assignment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_circuit_group_assignment", &resource.Sweeper{
		Name:         "netbox_circuit_group_assignment",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := circuits.NewCircuitsCircuitGroupAssignmentsListParams()
			res, err := api.Circuits.CircuitsCircuitGroupAssignmentsList(params, nil)
			if err != nil {
				return err
			}
			for _, assignment := range res.GetPayload().Results {
				if assignment.Group != nil && assignment.Group.Name != nil && strings.HasPrefix(*assignment.Group.Name, testPrefix) {
					deleteParams := circuits.NewCircuitsCircuitGroupAssignmentsDeleteParams().WithID(assignment.ID)
					_, err := api.Circuits.CircuitsCircuitGroupAssignmentsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a circuit group assignment")
				}
			}
			return nil
		},
	})
}
