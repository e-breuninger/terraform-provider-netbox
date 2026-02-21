package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVlansSetUp() string {
	return `
resource "netbox_vlan" "test_1" {
  name = "VLAN1234"
  vid  = 1234
}

resource "netbox_vlan" "test_2" {
  name = "VLAN1235"
  vid  = 1235
}

resource "netbox_vlan" "test_3" {
  name = "VLAN1236"
  vid  = 1236
}
  type = "text"
  content_types = ["ipam.vlan"]
  weight = 100
  default = "red"
  validation_regex = "^.*$"
}
resource "netbox_custom_field" "test_field2" {
  name = "field2"
  type = "text"
  content_types = ["ipam.vlan"]
  weight = 100
  default = "red"
  validation_regex = "^.*$"
}
resource "netbox_vlan" "test_with_custom_fields" {
  name = "VLAN_CUSTOM"
  vid  = 999

  custom_fields = {
    field1 = "value1"
    field2 = "value2"
  }
}`
}

func testAccNetboxVlansByVid() string {
	return `
data "netbox_vlans" "test" {
  filter {
	name  = "vid"
	value = "1234"
  }
}`
}

func testAccNetboxVlansByVidN() string {
	return `
data "netbox_vlans" "test" {
  filter {
	name = "vid__n"
	value = "1234"
  }
}`
}

func testAccNetboxVlansByVidRange() string {
	return `
data "netbox_vlans" "test" {
  filter {
	name = "vid__gte"
	value = "1235"
  }

  filter {
	name = "vid__lte"
	value = "1236"
  }
}`
}

func testAccNetboxVlansByCustomFields() string {
	return `
data "netbox_vlans" "test_custom_fields" {
  filter {
    name  = "vid"
    value = "999"
  }
}`
}

func TestAccNetboxVlansDataSource_basic(t *testing.T) {
	setUp := testAccNetboxVlansSetUp()
	// This test cannot be run in parallel with other tests, because other tests create also Vlans
	// These Vlans then interfere with the __n filter test
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test_1", "vid", "1234"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByVid(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.0.vid", "netbox_vlan.test_1", "vid"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByVidN(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.0.vid", "netbox_vlan.test_2", "vid"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByVidRange(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test", "vlans.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.0.vid", "netbox_vlan.test_2", "vid"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test", "vlans.1.vid", "netbox_vlan.test_3", "vid"),
				),
			},
			{
				Config: setUp + testAccNetboxVlansByCustomFields(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlans.test_custom_fields", "vlans.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlans.test_custom_fields", "vlans.0.vid", "netbox_vlan.test_with_custom_fields", "vid"),
					resource.TestCheckResourceAttr("data.netbox_vlans.test_custom_fields", "vlans.0.custom_fields.field1", "value1"),
					resource.TestCheckResourceAttr("data.netbox_vlans.test_custom_fields", "vlans.0.custom_fields.field2", "value2"),
				),
			},
		},
	})
}
