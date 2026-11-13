//go:build acctest

package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// TestAccNetboxPrimaryIP_dualStackImport imports the primary IPv4 and IPv6 of one device by IP
// version, expects the import without a version to be refused as ambiguous, and removes the IPv6
// resource: the primary IPv4 must stay.
func TestAccNetboxPrimaryIP_dualStackImport(t *testing.T) {
	testName := testAccGetTestName("primip_import")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "eth0"
  type      = "1000base-t"
}
resource "netbox_ip_address" "v4" {
  ip_address          = "192.0.2.74/24"
  device_interface_id = netbox_device_interface.test.id
}
resource "netbox_ip_address" "v6" {
  ip_address          = "2001:db8::74/64"
  device_interface_id = netbox_device_interface.test.id
}
resource "netbox_primary_ip" "v4" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.v4.id
}`, testName)
	v6 := `
resource "netbox_primary_ip" "v6" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.v6.id
}`
	device := func(dependsOn string) string {
		return fmt.Sprintf(`
data "netbox_device" "test" {
  id         = netbox_device.test.id
  depends_on = [%s]
}`, dependsOn)
	}
	importID := func(suffix string) resource.ImportStateIdFunc {
		return func(state *terraform.State) (string, error) {
			return "device/" + state.RootModule().Resources["netbox_device.test"].Primary.ID + suffix, nil
		}
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{Config: deps + v6 + device("netbox_primary_ip.v4, netbox_primary_ip.v6")},
			{
				ResourceName:      "netbox_primary_ip.v4",
				ImportState:       true,
				ImportStateIdFunc: importID("/4"),
				ImportStateVerify: true,
				// Both resources share the id device/<id>.
				ImportStateVerifyIdentifierAttribute: "ip_address_id",
			},
			{
				ResourceName:      "netbox_primary_ip.v6",
				ImportState:       true,
				ImportStateIdFunc: importID("/6"),
				ImportStateVerify: true,
				// Both resources share the id device/<id>.
				ImportStateVerifyIdentifierAttribute: "ip_address_id",
			},
			{
				ResourceName:      "netbox_primary_ip.v6",
				ImportState:       true,
				ImportStateIdFunc: importID(""),
				ExpectError:       regexp.MustCompile("Ambiguous import ID"),
			},
			{
				ResourceName:      "netbox_primary_ip.v6",
				ImportState:       true,
				ImportStateIdFunc: importID("/5"),
				ExpectError:       regexp.MustCompile("Invalid import ID"),
			},
			// Removing the IPv6 resource; the next step reads the device afterwards.
			{Config: deps + device("netbox_primary_ip.v4")},
			{
				Config: deps + device("netbox_primary_ip.v4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_device.test", "primary_ip4_id", "netbox_ip_address.v4", "id"),
					resource.TestCheckNoResourceAttr("data.netbox_device.test", "primary_ip6_id"),
				),
			},
		},
	})
}
