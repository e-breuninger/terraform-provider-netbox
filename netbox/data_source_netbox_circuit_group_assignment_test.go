package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCircuitGroupAssignmentDataSource_basic(t *testing.T) {
	testSlug := "circuit_group_assignment_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxCircuitGroupAssignmentDependencies(testName) + `
resource "netbox_circuit_group_assignment" "test" {
	group_id    = netbox_circuit_group.test.id
	member_type = "circuits.circuit"
	member_id   = netbox_circuit.test.id
	priority    = "primary"
	tags        = [netbox_tag.test.name]
}`
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_circuit_group_assignment" "test" {
	depends_on = []

	id = netbox_circuit_group_assignment.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_circuit_group_assignment.test", "id", "netbox_circuit_group_assignment.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_group_assignment.test", "group_id", "netbox_circuit_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group_assignment.test", "member_type", "circuits.circuit"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_group_assignment.test", "member_id", "netbox_circuit.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group_assignment.test", "priority", "primary"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group_assignment.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group_assignment.test", "tags.0", testName),
				),
			},
		},
	})
}
