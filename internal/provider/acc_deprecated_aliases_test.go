//go:build acctest

// The deprecated alias attributes (spec alias_of): a configuration written
// against the old names must apply, fill the canonical attributes, warn, and replan empty when
// rewritten to the canonical names. The full alias list is locked schema-side by
// TestAliasAttributesDeclared; this file exercises the plan-time behaviour on representatives of
// each shape (bool with default, scalar id, set of ids, list aliases whose name changes
// entirely) against the real API. Clear-on-unset of aliased canonicals is covered by the
// resources' own shrink steps (e.g. TestAccNetboxVpnTunnel_basic).
package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccNetboxAliases_deviceInterface: scalar and collection aliases (mgmtonly, untagged_vlan,
// tagged_vlans), the alias -> canonical rewrite as a no-op, and the both-set conflict.
func TestAccNetboxAliases_deviceInterface(t *testing.T) {
	testName := testAccGetTestName("alias_if")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_vlan" "tagged" {
  name   = "%[1]s-tagged"
  vid    = 3101
  status = "active"
}
resource "netbox_vlan" "untagged" {
  name   = "%[1]s-untagged"
  vid    = 3102
  status = "active"
}`, testName)
	aliased := deps + `
resource "netbox_device_interface" "test" {
  device_id     = netbox_device.test.id
  name          = "eth0"
  type          = "1000base-t"
  mode          = "tagged"
  mgmtonly      = true
  untagged_vlan = netbox_vlan.untagged.id
  tagged_vlans  = [netbox_vlan.tagged.id]
}`
	canonical := deps + `
resource "netbox_device_interface" "test" {
  device_id        = netbox_device.test.id
  name             = "eth0"
  type             = "1000base-t"
  mode             = "tagged"
  mgmt_only        = true
  untagged_vlan_id = netbox_vlan.untagged.id
  tagged_vlan_ids  = [netbox_vlan.tagged.id]
}`
	resourceName := "netbox_device_interface.test"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: aliased,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "mgmt_only", "true"),
					resource.TestCheckResourceAttr(resourceName, "mgmtonly", "true"),
					resource.TestCheckResourceAttrPair(resourceName, "untagged_vlan_id", "netbox_vlan.untagged", "id"),
					resource.TestCheckResourceAttrPair(resourceName, "untagged_vlan", "netbox_vlan.untagged", "id"),
					resource.TestCheckResourceAttr(resourceName, "tagged_vlan_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "tagged_vlans.#", "1"),
				),
			},
			{
				// The same assignment through the canonical names is a no-op: the migration off the
				// aliases costs nothing.
				Config:   canonical,
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + `
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "eth0"
  type      = "1000base-t"
  mgmtonly  = true
  mgmt_only = false
}`,
				ExpectError: regexp.MustCompile("deprecated alias of mgmt_only"),
			},
		},
	})
}

// TestAccNetboxAliases_configContext: the twelve scope-list aliases share one shape; sites and
// roles cover it (roles also renames entirely, to device_role_ids).
func TestAccNetboxAliases_configContext(t *testing.T) {
	testName := testAccGetTestName("alias_cc")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}`, testName)
	resourceName := "netbox_config_context.test"
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name  = "%[1]s"
  data  = jsonencode({ alias_test = true })
  sites = [netbox_site.test.id]
  roles = [netbox_device_role.test.id]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "site_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "device_role_ids.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "sites.#", "1"),
					resource.TestCheckResourceAttr(resourceName, "roles.#", "1"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_config_context" "test" {
  name            = "%[1]s"
  data            = jsonencode({ alias_test = true })
  site_ids        = [netbox_site.test.id]
  device_role_ids = [netbox_device_role.test.id]
}`, testName),
				PlanOnly: true,
			},
			{
				ResourceName:      resourceName,
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
