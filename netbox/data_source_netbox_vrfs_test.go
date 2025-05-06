package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccNetboxVrfsSetUp() string {
	return `
resource "netbox_tag" "test" {
  name = "test_tag"
}

resource "netbox_vrf" "test_1" {
  name = "VRF1"
}

resource "netbox_vrf" "test_2" {
  name = "VRF2"
  tags = [netbox_tag.test.name]
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

func testAccNetboxVrfsByTag() string {
	return `
data "netbox_vrfs" "test" {
  filter {
	name  = "tag"
	value = netbox_tag.test.name
  }
}`
}

func TestAccNetboxVrfsDataSource_basic(t *testing.T) {
	setUp := testAccNetboxVrfsSetUp()
	resource.Test(t, resource.TestCase{
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
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
			{
				Config: setUp + testAccNetboxVrfsByTag(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vrfs.test", "vrfs.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vrfs.test", "vrfs.0.name", "netbox_vrf.test_2", "name"),
				),
			},
		},
	})
}
