package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCircuitGroupDataSource_basic(t *testing.T) {
	testSlug := "circuit_group_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxCircuitGroupDependencies(testName) + fmt.Sprintf(`
resource "netbox_circuit_group" "test" {
	name 			= "%[1]s"
	slug 			= "%[1]s"
	comments 		= "Circuit Group Comments"
	description		= "This is my circuit group"
	tenant_id 		= netbox_tenant.test.id
	tags			= [netbox_tag.test.name]
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
data "netbox_circuit_group" "test" {
	depends_on	=	[]

	name = netbox_circuit_group.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_circuit_group.test", "id", "netbox_circuit_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_circuit_group.test", "slug", testName),
					resource.TestCheckResourceAttr("data.netbox_circuit_group.test", "comments", "Circuit Group Comments"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group.test", "description", "This is my circuit group"),
					resource.TestCheckResourceAttrPair("data.netbox_circuit_group.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_circuit_group.test", "tags.0", testName),
				),
			},
		},
	})
}
