package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualCircuitDataSource_basic(t *testing.T) {
	testSlug := "virtual_circuit_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxVirtualCircuitDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_circuit" "test" {
  cid                 = "%[1]s"
  provider_network_id = netbox_circuit_provider_network.test.id
  type_id             = netbox_virtual_circuit_type.test.id
  status              = "active"
  tenant_id           = netbox_tenant.test.id
  description         = "This is my virtual circuit"
  comments             = "This is my virtual circuit comment"
  tags                = [netbox_tag.test.name]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_virtual_circuit" "test" {
  depends_on = []

  cid = netbox_virtual_circuit.test.cid
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit.test", "id", "netbox_virtual_circuit.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit.test", "cid", testName),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit.test", "provider_network_id", "netbox_circuit_provider_network.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit.test", "type_id", "netbox_virtual_circuit_type.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit.test", "status", "active"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_circuit.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit.test", "description", "This is my virtual circuit"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit.test", "comments", "This is my virtual circuit comment"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_virtual_circuit.test", "tags.0", testName),
				),
			},
		},
	})
}
