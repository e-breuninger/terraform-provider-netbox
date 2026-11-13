//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

func TestAccNetboxPrimaryIP_dualStackImport(t *testing.T) {
	for _, kind := range []string{"device", "vm"} {
		t.Run(kind, func(t *testing.T) {
			name := testAccGetTestName("primary-import")
			ownerType, ownerAttr, interfaceAttr := "netbox_device", "device_id", "device_interface_id"
			deps := deviceDeps(name) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = %q
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "eth0"
  type      = "1000base-t"
}
`, name)
			interfaceRef := "netbox_device_interface.test.id"
			if kind == "vm" {
				ownerType, ownerAttr, interfaceAttr = "netbox_virtual_machine", "virtual_machine_id", "virtual_machine_interface_id"
				interfaceRef = "netbox_virtual_machine_interface.test.id"
				deps = fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = %[1]q
}
resource "netbox_cluster" "test" {
  name            = %[1]q
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_virtual_machine" "test" {
  name       = %[1]q
  cluster_id = netbox_cluster.test.id
}
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "eth0"
}
`, name)
			}
			deps += fmt.Sprintf(`
resource "netbox_ip_address" "v4" {
  ip_address = "192.0.2.74/24"
  %[1]s = %[2]s
}
resource "netbox_ip_address" "v6" {
  ip_address = "2001:db8::74/64"
  %[1]s = %[2]s
}
resource "netbox_primary_ip" "v4" {
  %[3]s = %[4]s.test.id
  ip_address_id = netbox_ip_address.v4.id
}
`, interfaceAttr, interfaceRef, ownerAttr, ownerType)
			v6 := fmt.Sprintf(`
resource "netbox_primary_ip" "v6" {
  %[1]s = %[2]s.test.id
  ip_address_id = netbox_ip_address.v6.id
  depends_on = [netbox_primary_ip.v4]
}
`, ownerAttr, ownerType)
			checkOwner := func(dependencies string) string {
				return fmt.Sprintf(`
data "%s" "test" {
  id = %s.test.id
  depends_on = [%s]
}
`, ownerType, ownerType, dependencies)
			}
			config := deps + v6 + checkOwner("netbox_primary_ip.v4, netbox_primary_ip.v6")
			steps := []resource.TestStep{{Config: config}}
			for _, family := range []int{4, 6} {
				address := fmt.Sprintf("netbox_primary_ip.v%d", family)
				steps = append(steps, resource.TestStep{
					ResourceName:                         address,
					ImportState:                          true,
					ImportStatePersist:                   true,
					ImportStateVerify:                    true,
					ImportStateVerifyIdentifierAttribute: "ip_address_id",
					ImportStateIdFunc: func(state *terraform.State) (string, error) {
						return fmt.Sprintf("%s/%s/%d", kind, state.RootModule().Resources[ownerType+".test"].Primary.ID, family), nil
					},
				}, resource.TestStep{
					Config:   config,
					PlanOnly: true,
					ConfigPlanChecks: resource.ConfigPlanChecks{
						PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
					},
				})
			}
			steps = append(steps, resource.TestStep{
				// Removing the imported IPv6 binding must preserve primary IPv4.
				Config: deps + checkOwner("netbox_primary_ip.v4"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data."+ownerType+".test", "primary_ip4_id", "netbox_ip_address.v4", "id"),
					resource.TestCheckNoResourceAttr("data."+ownerType+".test", "primary_ip6_id"),
				),
			})
			resource.ParallelTest(t, resource.TestCase{
				ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
				PreCheck:                 func() { testAccPreCheck(t) },
				Steps:                    steps,
			})
		})
	}
}
