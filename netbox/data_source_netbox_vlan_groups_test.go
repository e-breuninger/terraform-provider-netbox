package netbox

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxVlanGroupsSetUp() string {
	return `
resource "netbox_vlan_group" "test_1" {
  name 			 = "VLANGroup1"
  slug 			 = "vlangroup1"
  vid_ranges = [[100, 200]]
}

resource "netbox_vlan_group" "test_2" {
  name 			 = "VLANGroup2"
  slug 			 = "vlangroup2"
  vid_ranges = [[300, 400]]
}

resource "netbox_vlan_group" "test_3" {
  name 			 = "VLANGroup3"
  slug 			 = "vlangroup3"
  vid_ranges = [[500, 600]]
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

func testAccNetboxVlanGroupsBySlug() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name = "slug"
	value = "vlangroup2"
  }
}`
}

func testAccNetboxVlanGroupsWithNameIsw() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name  = "name__isw"
	value = "VLANGroup"
  }
}`
}

func testAccNetboxVlanGroupsByID() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name  = "id"
	value = netbox_vlan_group.test_1.id
  }
}`
}

func testAccNetboxVlanGroupsWithLimit() string {
	return `
data "netbox_vlan_groups" "test" {
  filter {
	name  = "name__isw"
	value = "VLANGroup"
  }
  limit = 2
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
				Config: setUp + testAccNetboxVlanGroupsBySlug(),
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
				Config: setUp,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan_group.test_1", "name", "VLANGroup1"),
				),
			},
			{
				Config: setUp + testAccNetboxVlanGroupsWithNameIsw(),
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

func TestAccNetboxVlanGroupsDataSource_byID(t *testing.T) {
	setUp := testAccNetboxVlanGroupsSetUp()
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
			},
			{
				Config: setUp + testAccNetboxVlanGroupsByID(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_vlan_groups.test", "vlan_groups.0.name", "netbox_vlan_group.test_1", "name"),
				),
			},
		},
	})
}

func TestAccNetboxVlanGroupsDataSource_limit(t *testing.T) {
	setUp := testAccNetboxVlanGroupsSetUp()
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: setUp,
			},
			{
				Config: setUp + testAccNetboxVlanGroupsWithLimit(),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_vlan_groups.test", "vlan_groups.#", "2"),
				),
			},
		},
	})
}
