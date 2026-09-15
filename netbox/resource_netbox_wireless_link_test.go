package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxWirelessLinkDependencies(testName string) string {
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

resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model           = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test_a" {
  name           = "%[1]s_a"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device" "test_b" {
  name           = "%[1]s_b"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}

resource "netbox_device_interface" "test_a" {
  name      = "%[1]s_a"
  device_id = netbox_device.test_a.id
  type      = "ieee802.11a"
}

resource "netbox_device_interface" "test_b" {
  name      = "%[1]s_b"
  device_id = netbox_device.test_b.id
  type      = "ieee802.11a"
}`, testName)
}

func TestAccNetboxWirelessLink_basic(t *testing.T) {
	testSlug := "wlink_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxWirelessLinkDependencies(testName) + `
resource "netbox_wireless_link" "test" {
  interface_a_id = netbox_device_interface.test_a.id
  interface_b_id = netbox_device_interface.test_b.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_wireless_link.test", "interface_a_id", "netbox_device_interface.test_a", "id"),
					resource.TestCheckResourceAttrPair("netbox_wireless_link.test", "interface_b_id", "netbox_device_interface.test_b", "id"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "status", "connected"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "ssid", ""),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "tags.#", "0"),
				),
			},
			{
				ResourceName:      "netbox_wireless_link.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxWirelessLink_withDependencies(t *testing.T) {
	testSlug := "wlink_deps"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxWirelessLinkDependencies(testName) + fmt.Sprintf(`
resource "netbox_wireless_link" "test" {
  interface_a_id = netbox_device_interface.test_a.id
  interface_b_id = netbox_device_interface.test_b.id
  ssid            = "%[1]s"
  status          = "planned"
  tenant_id       = netbox_tenant.test.id
  auth_type       = "wpa-personal"
  auth_cipher     = "aes"
  auth_psk        = "supersecret123"
  description     = "test description"
  comments        = "test comments"
  tags            = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "ssid", testName),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "status", "planned"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "auth_type", "wpa-personal"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "auth_cipher", "aes"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "description", "test description"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "comments", "test comments"),
					resource.TestCheckResourceAttrPair("netbox_wireless_link.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_wireless_link.test", "tags.0", testName),
				),
			},
			{
				ResourceName:            "netbox_wireless_link.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"auth_psk"},
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_wireless_link", &resource.Sweeper{
		Name:         "netbox_wireless_link",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := wireless.NewWirelessWirelessLinksListParams()
			res, err := api.Wireless.WirelessWirelessLinksList(params, nil)
			if err != nil {
				return err
			}
			for _, link := range res.GetPayload().Results {
				if link.Ssid != "" && strings.HasPrefix(link.Ssid, testPrefix) {
					deleteParams := wireless.NewWirelessWirelessLinksDeleteParams().WithID(link.ID)
					_, err := api.Wireless.WirelessWirelessLinksDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a wireless link")
				}
			}
			return nil
		},
	})
}
