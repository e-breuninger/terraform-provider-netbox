package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVlanGroupsSetUp() string {
	return `
resource "netbox_vlan_group" "test_1" {
  name = "VLANGroup1"
  slug = "vlangroup1"
  min_vid = 100
  max_vid = 200
}

resource "netbox_vlan_group" "test_2" {
  name = "VLANGroup2"
  slug = "vlangroup2"
  min_vid = 300
  max_vid = 400
}

resource "netbox_vlan_group" "test_3" {
  name = "VLANGroup3"
  slug = "vlangroup3"
  min_vid = 500
  max_vid = 600
}`
}

func testAccNetboxVlanGroupsByName() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name  = "name"
	value = "VLANGroup1"
  }
}`
}

func testAccNetboxVlanGroupsByNameN() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name = "name__n"
	value = "VLANGroup1"
  }
}`
}

func testAccNetboxVlanGroupsBySlug() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name = "slug"
	value = "vlangroup2"
  }
}`
}

func testAccNetboxVlanGroupsByMinVidRange() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name = "minvid"
	value = "300"
  }
}`
}

func testAccNetboxVlanGroupsWithQ() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name = "name"
	value = "VLANGroup"
  }
}`
}

func TestAccNetboxVlanGroupsDataSource_basic(t *testing.T) {
	setUp := testAccNetboxVlanGroupsSetUp()
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test_1", "name", "VLANGroup1"),
					resource.TestCheckResourceAttr("netbox_vlan_group.test_1", "slug", "vlangroup1"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanGroupsByName(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.0.name", "netbox_vlan_group.test_1", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.0.slug", "netbox_vlan_group.test_1", "slug"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanGroupsByNameN(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.0.name", "netbox_vlan_group.test_2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.1.name", "netbox_vlan_group.test_3", "name"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanGroupsBySlug(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.0.slug", "netbox_vlan_group.test_2", "slug"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanGroupsByMinVidRange(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.0.slug", "netbox_vlan_group.test_2", "slug"),
				),
			},
		},
	})
}

func TestAccNetboxVlanGroupsDataSource_search(t *testing.T) {
	setUp := testAccNetboxVlanGroupsSetUp()
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp + testAccNetboxVlanGroupsWithQ(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "3"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.0.slug", "netbox_vlan_group.test_1", "slug"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.1.slug", "netbox_vlan_group.test_2", "slug"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.2.slug", "netbox_vlan_group.test_3", "slug"),
				),
			},
		},
	})
}
