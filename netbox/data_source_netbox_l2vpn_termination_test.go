package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxL2VPNTerminationDataSource_basic(t *testing.T) {
	testSlug := "l2vpn_term_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxL2VPNTerminationDependencies(testName) + `
resource "netbox_l2vpn_termination" "test" {
  l2vpn_id              = netbox_l2vpn.test.id
  assigned_object_type  = "dcim.interface"
  assigned_object_id    = netbox_device_interface.test.id
  tags                  = [netbox_tag.test.name]
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
data "netbox_l2vpn_termination" "test" {
  depends_on = []

  id = tonumber(netbox_l2vpn_termination.test.id)
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_l2vpn_termination.test", "id", "netbox_l2vpn_termination.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_l2vpn_termination.test", "l2vpn_id", "netbox_l2vpn.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn_termination.test", "assigned_object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("data.netbox_l2vpn_termination.test", "assigned_object_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn_termination.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_l2vpn_termination.test", "tags.0", testName),
				),
			},
		},
	})
}
