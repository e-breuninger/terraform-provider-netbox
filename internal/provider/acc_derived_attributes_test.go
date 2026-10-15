//go:build acctest

// Attributes NetBox fills in when they are not set (see derived_attributes.go).
package provider_test

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

// derivedFrom: NetBox derives the value on every save from another attribute.
// Unset values converge, changing the source derives them again, and removing
// the source clears them.
func TestAccNetboxDerivedAttributes(t *testing.T) {
	testName := testAccGetTestName("derived")
	slug := getSlug(testName)
	deps := fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
resource "netbox_site" "a" {
  name = "%[1]s-a"
}
resource "netbox_site" "b" {
  name = "%[1]s-b"
}
resource "netbox_location" "one" {
  name    = "%[1]s-one"
  slug    = "%[2]s-one"
  site_id = netbox_site.a.id
}
resource "netbox_location" "two" {
  name    = "%[1]s-two"
  slug    = "%[2]s-two"
  site_id = netbox_site.a.id
}
resource "netbox_rack" "one" {
  name        = "%[1]s-one"
  site_id     = netbox_site.a.id
  location_id = netbox_location.one.id
}
resource "netbox_rack" "two" {
  name        = "%[1]s-two"
  site_id     = netbox_site.a.id
  location_id = netbox_location.two.id
}
resource "netbox_rack_type" "small" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s-small"
  form_factor     = "2-post-frame"
  width           = 19
  u_height        = 42
  outer_width     = 500
  outer_unit      = "mm"
}
resource "netbox_rack_type" "large" {
  manufacturer_id = netbox_manufacturer.test.id
  model           = "%[1]s-large"
  form_factor     = "4-post-cabinet"
  width           = 23
  u_height        = 47
  outer_width     = 600
  outer_unit      = "mm"
}
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "a" {
  name            = "%[1]s-a"
  cluster_type_id = netbox_cluster_type.test.id
  scope_type      = "dcim.site"
  scope_id        = netbox_site.a.id
}
resource "netbox_cluster" "b" {
  name            = "%[1]s-b"
  cluster_type_id = netbox_cluster_type.test.id
  scope_type      = "dcim.site"
  scope_id        = netbox_site.b.id
}`, testName, slug)
	config := func(rack, rackType, cluster string) string {
		return deps + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.a.id
  rack_id        = %[2]s
}
resource "netbox_rack" "typed" {
  name         = "%[1]s-typed"
  site_id      = netbox_site.a.id
  rack_type_id = netbox_rack_type.%[3]s.id
}
resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.%[4]s.id
}`, testName, rack, rackType, cluster)
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config("netbox_rack.one.id", "small", "a"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.one", "id"),
					resource.TestCheckResourceAttr("netbox_rack.typed", "form_factor", "2-post-frame"),
					resource.TestCheckResourceAttr("netbox_rack.typed", "width", "19"),
					resource.TestCheckResourceAttr("netbox_rack.typed", "outer_width", "500"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "site_id", "netbox_site.a", "id"),
				),
			},
			{ // change the sources: NetBox derives the values again
				Config: config("netbox_rack.two.id", "large", "b"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.two", "id"),
					resource.TestCheckResourceAttr("netbox_rack.typed", "form_factor", "4-post-cabinet"),
					resource.TestCheckResourceAttr("netbox_rack.typed", "width", "23"),
					resource.TestCheckResourceAttr("netbox_rack.typed", "outer_width", "600"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "site_id", "netbox_site.b", "id"),
				),
			},
			{ // remove the rack: the location it implied is cleared
				Config: config("null", "large", "b"),
				Check:  resource.TestCheckNoResourceAttr("netbox_device.test", "location_id"),
			},
		},
	})
}

// derivedOnCreate: NetBox derives the value only on create. Creating an object
// that inherits it no longer fails; as before, an unset value is then cleared,
// and a value set in the configuration stays.
func TestAccNetboxDerivedOnCreate(t *testing.T) {
	testName := testAccGetTestName("oncreate")
	deps := fmt.Sprintf(`
resource "netbox_platform" "test" {
  name = "%[1]s"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}
resource "netbox_device_type" "test" {
  manufacturer_id     = netbox_manufacturer.test.id
  model               = "%[1]s"
  default_platform_id = netbox_platform.test.id
  airflow             = "front-to-rear"
}
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
resource "netbox_site" "test" {
  name = "%[1]s"
}
resource "netbox_virtual_machine_type" "test" {
  name                = "%[1]s"
  default_platform_id = netbox_platform.test.id
  default_vcpus       = 4
  default_memory      = 2048
}`, testName)
	config := func(platform string) string {
		return deps + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
  platform_id    = %[2]s
}
resource "netbox_virtual_machine" "test" {
  name                    = "%[1]s"
  site_id                 = netbox_site.test.id
  virtual_machine_type_id = netbox_virtual_machine_type.test.id
  platform_id             = %[2]s
}`, testName, platform)
	}
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{ // created with inherited values; the next plan clears the unset ones
				Config: config("null"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "airflow", "front-to-rear"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "4"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "memory_mb", "2048"),
				),
				ExpectNonEmptyPlan: true,
			},
			{
				Config: config("null"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device.test", "platform_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "airflow"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "vcpus"),
				),
			},
			{
				Config: config("netbox_platform.test.id"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "platform_id", "netbox_platform.test", "id"),
				),
			},
		},
	})
}
