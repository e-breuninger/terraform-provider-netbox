package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualCircuitTerminationDataSource_basic(t *testing.T) {
	testSlug := "virtual_circuit_term_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxVirtualCircuitTerminationDependencies(testName) + `
resource "netbox_virtual_circuit_termination" "test" {
  virtual_circuit_id = netbox_virtual_circuit.test.id
  interface_id        = netbox_device_interface.test.id
  role                = "peer"
  description         = "This is my virtual circuit termination"
  tags                = [netbox_tag.test.name]
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
data "netbox_virtual_circuit_termination" "test" {
  depends_on = []

  id = netbox_virtual_circuit_termination.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit_termination.test", "id", "netbox_virtual_circuit_termination.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit_termination.test", "virtual_circuit_id", "netbox_virtual_circuit.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit_termination.test", "interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_termination.test", "role", "peer"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_termination.test", "description", "This is my virtual circuit termination"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_termination.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit_termination.test", "tags.0", testName),
				),
			},
		},
	})
}
