package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVrfsSetUp() string {
	return `
resource "netbox_vrf" "test_1" {
  name = "VRF1"
}

resource "netbox_vrf" "test_2" {
  name = "VRF2"
}

resource "netbox_vrf" "test_3" {
  name = "VRF3"
}`
}

func testAccNetboxVrfsByName() string {
	return `
data "netbox_vrfs" "test" {
  filter {
	name  = "name"
	value = "VRF1"
  }
}`
}

func TestAccNetboxVrfsDataSource_basic(t *testing.T) {
	setUp := testAccNetboxVrfsSetUp()
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vrf.test_1", "name", "VRF1"),
					resource.TestCheckResourceAttr("netbox_vrf.test_2", "name", "VRF2"),
					resource.TestCheckResourceAttr("netbox_vrf.test_3", "name", "VRF3"),
				),
			},
			{
				Config: setUp + testAccNetboxVrfsByName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vrfs.test", "vrfs.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vrfs.test", "vrfs.0.name", "netbox_vrf.test_1", "name"),
				),
			},
		},
	})
}
