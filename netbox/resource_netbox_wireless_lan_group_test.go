package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxWirelessLANGroup_basic(t *testing.T) {
	testSlug := "wlangrp_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_wireless_lan_group" "parent" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "foo bar"
  tags        = [netbox_tag.test.name]
}

resource "netbox_wireless_lan_group" "child" {
  name      = "%[1]s-child"
  slug      = "%[2]s-c"
  parent_id = netbox_wireless_lan_group.parent.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.parent", "name", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.parent", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.parent", "description", "foo bar"),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.parent", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.parent", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.child", "name", fmt.Sprintf("%s-child", testName)),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.child", "slug", fmt.Sprintf("%s-c", randomSlug)),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan_group.child", "parent_id", "netbox_wireless_lan_group.parent", "id"),
				),
			},
			{
				ResourceName:      "netbox_wireless_lan_group.parent",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxWirelessLANGroup_defaultSlug(t *testing.T) {
	testSlug := "wlangrp_defslug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_wireless_lan_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func TestAccNetboxWirelessLANGroup_updateParentAndDescription(t *testing.T) {
	testSlug := "wlangrp_update"
	testName := testAccGetTestName(testSlug)
	parentName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_wireless_lan_group" "parent" {
  name = "%[1]s"
}

resource "netbox_wireless_lan_group" "test" {
  name        = "%[2]s"
  description = "foo bar"
  parent_id   = netbox_wireless_lan_group.parent.id
}`, parentName, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "description", "foo bar"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan_group.test", "parent_id", "netbox_wireless_lan_group.parent", "id"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_wireless_lan_group" "parent" {
  name = "%[1]s"
}

resource "netbox_wireless_lan_group" "test" {
  name = "%[2]s"
}`, parentName, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "parent_id", "0"),
				),
			},
		},
	})
}
