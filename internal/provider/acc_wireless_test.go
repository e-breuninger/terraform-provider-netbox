//go:build acctest

// Wireless: LAN groups, LANs and links.
package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxWirelessLANGroup_basic(t *testing.T) {
	testName := testAccGetTestName("wlangroup")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_wireless_lan_group" "parent" {
  name = "%[1]s-parent"
}
resource "netbox_wireless_lan_group" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  parent_id   = netbox_wireless_lan_group.parent.id
  description = "Acceptance test wireless LAN group."
  comments    = "Created by acceptance test."
}
data "netbox_wireless_lan_group" "test" {
  name = netbox_wireless_lan_group.test.name
}
data "netbox_wireless_lan_groups" "test" {
  filters = [
    { name = "parent_id", value = netbox_wireless_lan_group.parent.id },
  ]
  depends_on = [netbox_wireless_lan_group.test]
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan_group.test", "slug", getSlug(testName)),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan_group.test", "parent_id", "netbox_wireless_lan_group.parent", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_wireless_lan_group.test", "id", "netbox_wireless_lan_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_wireless_lan_groups.test", "wireless_lan_groups.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_wireless_lan_group" "parent" {
  name = "%[1]s-parent"
}
resource "netbox_wireless_lan_group" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_wireless_lan_group.test", "parent_id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan_group.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan_group.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_wireless_lan_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxWirelessLAN_basic(t *testing.T) {
	testName := testAccGetTestName("wlan")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_vlan" "test" {
  name = "%[1]s"
  vid  = 1234
}
resource "netbox_wireless_lan_group" "test" {
  name = "%[1]s"
}
resource "netbox_site_group" "test" {
  name = "%[1]s"
}
resource "netbox_region" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid        = "%[1]s"
  description = "Acceptance test wireless LAN."
  group_id    = netbox_wireless_lan_group.test.id
  status      = "reserved"
  vlan_id     = netbox_vlan.test.id
  tenant_id   = netbox_tenant.test.id
  auth_type   = "wpa-personal"
  auth_cipher = "aes"
  auth_psk    = "correct horse battery staple"
  scope_type  = "dcim.site"
  scope_id    = netbox_site.test.id
  comments    = "Created by acceptance test."
}
data "netbox_wireless_lan" "test" {
  ssid = netbox_wireless_lan.test.ssid
}
data "netbox_wireless_lans" "test" {
  filters = [
    { name = "ssid__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_wireless_lan.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "ssid", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_type", "wpa-personal"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_cipher", "aes"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "scope_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "location_id"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "vlan_id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_wireless_lan.test", "id", "netbox_wireless_lan.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_wireless_lans.test", "wireless_lans.#", "1"),
				),
			},
			{
				// Cycle the scope through each alias.
				Config: deps + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid    = "%[1]s"
  site_id = netbox_site.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "scope_type", "dcim.site"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "site_id", "netbox_site.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid      = "%[1]s"
  region_id = netbox_region.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "scope_type", "dcim.region"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "region_id", "netbox_region.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid          = "%[1]s"
  site_group_id = netbox_site_group.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "scope_type", "dcim.sitegroup"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "site_group_id", "netbox_site_group.test", "id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid        = "%[1]s"
  location_id = netbox_location.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "scope_type", "dcim.location"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "location_id", "netbox_location.test", "id"),
				),
			},
			{
				// Shrink: status is optional+computed and keeps NetBox's value; the rest clears.
				Config: deps + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "group_id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "vlan_id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "auth_type"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "auth_cipher"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "auth_psk"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "scope_type"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "scope_id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "site_id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_wireless_lan.test", "comments"),
				),
			},
			{
				ResourceName:      "netbox_wireless_lan.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

// TestAccNetboxWirelessLink_basic links two wireless interfaces of one device.
func TestAccNetboxWirelessLink_basic(t *testing.T) {
	testName := testAccGetTestName("wlink")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_device_interface" "a" {
  device_id = netbox_device.test.id
  name      = "wlan0"
  type      = "ieee802.11ax"
}
resource "netbox_device_interface" "b" {
  device_id = netbox_device.test.id
  name      = "wlan1"
  type      = "ieee802.11ax"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_wireless_link" "test" {
  interface_a_id = netbox_device_interface.a.id
  interface_b_id = netbox_device_interface.b.id
  ssid           = "%[1]s"
  status         = "planned"
  tenant_id      = netbox_tenant.test.id
  auth_type      = "wpa-personal"
  auth_cipher    = "aes"
  auth_psk       = "correct horse battery staple"
  distance       = 1.5
  distance_unit  = "km"
  description    = "Acceptance test wireless link."
  comments       = "Created by acceptance test."
}
data "netbox_wireless_link" "test" {
  ssid = netbox_wireless_link.test.ssid
}
data "netbox_wireless_links" "test" {
  filters = [
    { name = "ssid__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_wireless_link.test]
}
data "netbox_wireless_links" "by_interface" {
  filters = [
    { name = "interfacea_id", value = netbox_device_interface.a.id },
  ]
  depends_on     = [netbox_wireless_link.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "auth_type", "wpa-personal"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "auth_cipher", "aes"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "distance", "1.5"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "distance_unit", "km"),
					resource.TestCheckResourceAttrPair("netbox_wireless_link.test", "interface_a_id", "netbox_device_interface.a", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_wireless_link.test", "id", "netbox_wireless_link.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_wireless_links.test", "wireless_links.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_wireless_links.by_interface", "wireless_links.0.id", "netbox_wireless_link.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_wireless_link.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: status is optional+computed and keeps NetBox's value; the rest clears.
				Config: deps + `
resource "netbox_wireless_link" "test" {
  interface_a_id = netbox_device_interface.a.id
  interface_b_id = netbox_device_interface.b.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "status", "planned"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "ssid"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "auth_type"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "auth_cipher"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "auth_psk"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "distance"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "distance_unit"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_wireless_link.test", "comments"),
				),
			},
		},
	})
}

func init() {
	sweep("netbox_wireless_link",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Wireless.WirelessWirelessLinksList(wireless.NewWirelessWirelessLinksListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, result.Ssid})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Wireless.WirelessWirelessLinksDestroy(wireless.NewWirelessWirelessLinksDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_wireless_lan",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Wireless.WirelessWirelessLansList(wireless.NewWirelessWirelessLansListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Ssid)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Wireless.WirelessWirelessLansDestroy(wireless.NewWirelessWirelessLansDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_wireless_lan_group",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Wireless.WirelessWirelessLanGroupsList(wireless.NewWirelessWirelessLanGroupsListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Wireless.WirelessWirelessLanGroupsDestroy(wireless.NewWirelessWirelessLanGroupsDestroyParams().WithID(id), nil)
			return err
		})
}
