package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxWirelessLinkDataSource_basic(t *testing.T) {
	testSlug := "wlink_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxWirelessLinkDependencies(testName) + fmt.Sprintf(`
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
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_wireless_link" "test" {
  depends_on = []

  id = tonumber(netbox_wireless_link.test.id)
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_wireless_link.test", "id", "netbox_wireless_link.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_wireless_link.test", "interface_a_id", "netbox_device_interface.test_a", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_wireless_link.test", "interface_b_id", "netbox_device_interface.test_b", "id"),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "ssid", testName),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "status", "planned"),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "auth_type", "wpa-personal"),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "auth_cipher", "aes"),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "description", "test description"),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "comments", "test comments"),
					resource.TestCheckResourceAttrPair("data.netbox_wireless_link.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_wireless_link.test", "tags.0", testName),
				),
			},
		},
	})
}
