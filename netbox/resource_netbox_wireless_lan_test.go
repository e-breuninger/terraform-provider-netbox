package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxWirelessLANDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_vlan_group" "test_group" {
  name       = "%[1]s"
  slug       = "%[1]s"
  scope_type = "dcim.site"
  scope_id   = netbox_site.test.id
  vid_ranges = [[1, 4094]]
}

resource "netbox_vlan" "test" {
  name      = "%[1]s"
  vid       = 777
  status    = "active"
  tenant_id = netbox_tenant.test.id
  site_id   = netbox_site.test.id
  group_id  = netbox_vlan_group.test_group.id
}

resource "netbox_wireless_lan_group" "test" {
  name = "%[1]s"
}`, testName)
}

func TestAccNetboxWirelessLAN_basic(t *testing.T) {
	testSlug := "wlan_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "ssid", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "tags.#", "0"),
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

func TestAccNetboxWirelessLAN_withDependencies(t *testing.T) {
	testSlug := "wlan_with_deps"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxWirelessLANDependencies(testName) + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid        = "%[1]s"
  status      = "reserved"
  group_id    = netbox_wireless_lan_group.test.id
  tenant_id   = netbox_tenant.test.id
  vlan_id     = netbox_vlan.test.id
  auth_type   = "wpa-personal"
  auth_cipher = "aes"
  auth_psk    = "supersecret123"
  description = "test description"
  comments    = "test comments"
  tags        = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "ssid", testName),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_type", "wpa-personal"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_cipher", "aes"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "description", "test description"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "comments", "test comments"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "group_id", "netbox_wireless_lan_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "vlan_id", "netbox_vlan.test", "id"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "tags.0", testName),
				),
			},
			{
				ResourceName:            "netbox_wireless_lan.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auth_psk"},
			},
		},
	})
}

func TestAccNetboxWirelessLAN_clearOptionalFields(t *testing.T) {
	testSlug := "wlan_clear_opts"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxWirelessLANDependencies(testName) + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid        = "%[1]s"
  status      = "reserved"
  group_id    = netbox_wireless_lan_group.test.id
  tenant_id   = netbox_tenant.test.id
  vlan_id     = netbox_vlan.test.id
  auth_type   = "wpa-personal"
  auth_cipher = "aes"
  auth_psk    = "supersecret123"
  description = "test description"
  comments    = "test comments"
  tags        = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_type", "wpa-personal"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_cipher", "aes"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "description", "test description"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "comments", "test comments"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "group_id", "netbox_wireless_lan_group.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_wireless_lan.test", "vlan_id", "netbox_vlan.test", "id"),
				),
			},
			{
				Config: testAccNetboxWirelessLANDependencies(testName) + fmt.Sprintf(`
resource "netbox_wireless_lan" "test" {
  ssid   = "%[1]s"
  status = "active"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_type", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_cipher", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "auth_psk", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "group_id", "0"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "tenant_id", "0"),
					resource.TestCheckResourceAttr("netbox_wireless_lan.test", "vlan_id", "0"),
				),
			},
		},
	})
}
