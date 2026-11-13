//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// TestAccNetboxVirtualChassisMaster_basic sets the master of a two-member chassis, moves it to
// the other member, then updates only the chassis: its update must keep the master the companion
// resource set (otherwise the refresh after apply drops the master resource and the plan is not
// empty). Last, the master device itself is replaced, which NetBox only allows once the master is
// cleared.
func TestAccNetboxVirtualChassisMaster_basic(t *testing.T) {
	testName := testAccGetTestName("vc_master")
	config := func(master, description string) string {
		return deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_virtual_chassis" "test" {
  name        = "%[1]s"
  description = "%[3]s"
}
resource "netbox_device" "a" {
  name                     = "%[1]s-a"
  device_type_id           = netbox_device_type.test.id
  role_id                  = netbox_device_role.test.id
  site_id                  = netbox_site.test.id
  virtual_chassis_id       = netbox_virtual_chassis.test.id
  virtual_chassis_position = 1
}
resource "netbox_device" "b" {
  name                     = "%[1]s-b"
  device_type_id           = netbox_device_type.test.id
  role_id                  = netbox_device_role.test.id
  site_id                  = netbox_site.test.id
  virtual_chassis_id       = netbox_virtual_chassis.test.id
  virtual_chassis_position = 2
}
resource "netbox_virtual_chassis_master" "test" {
  virtual_chassis_id = netbox_virtual_chassis.test.id
  device_id          = netbox_device.%[2]s.id
}
data "netbox_virtual_chassis" "test" {
  id         = netbox_virtual_chassis.test.id
  depends_on = [netbox_virtual_chassis_master.test]
}`, testName, master, description)
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config("a", "one"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_chassis_master.test", "id", "netbox_virtual_chassis.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_chassis_master.test", "device_id", "netbox_device.a", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_chassis.test", "master_id", "netbox_device.a", "id"),
				),
			},
			{
				// The master moves to the other member: the resource is replaced.
				Config: config("b", "one"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_chassis_master.test", "device_id", "netbox_device.b", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_chassis.test", "master_id", "netbox_device.b", "id"),
				),
			},
			{
				// Only the chassis changes: its update must not clear the master.
				Config: config("b", "two"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_chassis.test", "description", "two"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_chassis.test", "master_id", "netbox_device.b", "id"),
				),
			},
			{
				// The master device is replaced. NetBox refuses to delete a device that is still
				// the master, so the master resource has to be replaced with it, cleared first.
				Taint:  []string{"netbox_device.b"},
				Config: config("b", "two"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_chassis_master.test", "device_id", "netbox_device.b", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_chassis.test", "master_id", "netbox_device.b", "id"),
				),
			},
			{
				ResourceName:      "netbox_virtual_chassis_master.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}
