//go:build acctest

// IPAM aggregates and ranges, services and service templates, and FHRP groups with their interface
// assignments.
package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/acctest"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxAggregate_basic(t *testing.T) {
	testName := testAccGetTestName("aggregate")
	prefix := fmt.Sprintf("%d.0.0.0/8", acctest.RandIntRange(20, 126))
	deps := fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_aggregate" "test" {
  prefix      = "%[2]s"
  rir_id      = netbox_rir.test.id
  tenant_id   = netbox_tenant.test.id
  date_added  = "2026-01-15"
  description = "Acceptance test aggregate."
  comments    = "Created by acceptance test."
}
data "netbox_aggregate" "test" {
  depends_on = [netbox_aggregate.test]
  prefix     = "%[2]s"
}
data "netbox_aggregates" "test" {
  depends_on = [netbox_aggregate.test]
  filters = [
    { name = "rir_id", value = netbox_rir.test.id },
  ]
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_aggregate.test", "prefix", prefix),
					resource.TestCheckResourceAttr("netbox_aggregate.test", "date_added", "2026-01-15"),
					resource.TestCheckResourceAttr("netbox_aggregate.test", "family", "4"),
					resource.TestCheckResourceAttrPair("netbox_aggregate.test", "rir_id", "netbox_rir.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_aggregate.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_aggregate.test", "id", "netbox_aggregate.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_aggregates.test", "aggregates.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_aggregate.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: tenant, description and comments clear. date_added is optional+computed
				// (a strfmt.Date crossing as raw JSON, which has no clear value) and keeps its value.
				Config: deps + fmt.Sprintf(`
resource "netbox_aggregate" "test" {
  prefix = "%[2]s"
  rir_id = netbox_rir.test.id
}`, testName, prefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_aggregate.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_aggregate.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_aggregate.test", "comments"),
					resource.TestCheckResourceAttr("netbox_aggregate.test", "date_added", "2026-01-15"),
				),
			},
		},
	})
}

func TestAccNetboxIPRange_basic(t *testing.T) {
	testName := testAccGetTestName("iprange")
	third := acctest.RandIntRange(1, 250)
	start := fmt.Sprintf("10.90.%d.1/24", third)
	end := fmt.Sprintf("10.90.%d.50/24", third)
	deps := fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_ipam_role" "test" {
  name = "%[1]s"
}
resource "netbox_vrf" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_ip_range" "test" {
  start_address  = "%[2]s"
  end_address    = "%[3]s"
  status         = "reserved"
  tenant_id      = netbox_tenant.test.id
  role_id        = netbox_ipam_role.test.id
  vrf_id         = netbox_vrf.test.id
  mark_utilized  = true
  mark_populated = true
  description    = "Acceptance test range."
  comments       = "Created by acceptance test."
}
data "netbox_ip_range" "test" {
  depends_on    = [netbox_ip_range.test]
  start_address = "%[2]s"
}
data "netbox_ip_ranges" "test" {
  depends_on = [netbox_ip_range.test]
  filters = [
    { name = "tenant_id", value = netbox_tenant.test.id },
  ]
}`, testName, start, end),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_range.test", "start_address", start),
					resource.TestCheckResourceAttr("netbox_ip_range.test", "end_address", end),
					resource.TestCheckResourceAttr("netbox_ip_range.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_ip_range.test", "size", "50"),
					resource.TestCheckResourceAttr("netbox_ip_range.test", "family", "4"),
					resource.TestCheckResourceAttr("netbox_ip_range.test", "mark_utilized", "true"),
					resource.TestCheckResourceAttr("netbox_ip_range.test", "mark_populated", "true"),
					resource.TestCheckResourceAttrPair("netbox_ip_range.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_range.test", "id", "netbox_ip_range.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_ip_ranges.test", "ip_ranges.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_ip_range.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: every reference and free-text field clears; the mark_* booleans fall back
				// to their default and status is optional+computed, keeping NetBox's value.
				Config: deps + fmt.Sprintf(`
resource "netbox_ip_range" "test" {
  start_address = "%[2]s"
  end_address   = "%[3]s"
}`, testName, start, end),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_ip_range.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_range.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_range.test", "vrf_id"),
					resource.TestCheckNoResourceAttr("netbox_ip_range.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_ip_range.test", "comments"),
					resource.TestCheckResourceAttr("netbox_ip_range.test", "mark_utilized", "false"),
				),
			},
		},
	})
}

func TestAccNetboxService_basic(t *testing.T) {
	testName := testAccGetTestName("service")
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
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
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-eth0"
  type      = "virtual"
}
resource "netbox_ip_address" "test" {
  ip_address           = "10.91.0.1/24"
  status               = "active"
  assigned_object_type = "dcim.interface"
  assigned_object_id   = netbox_device_interface.test.id
}
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_virtual_machine" "test" {
  name       = "%[1]s"
  cluster_id = netbox_cluster.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_service" "test" {
  name            = "%[1]s"
  protocol        = "tcp"
  ports          = [80, 443]
  device_id      = netbox_device.test.id
  ip_address_ids = [netbox_ip_address.test.id]
  description     = "Acceptance test service."
  comments        = "Created by acceptance test."
}
data "netbox_service" "test" {
  depends_on = [netbox_service.test]
  name       = "%[1]s"
}
data "netbox_services" "test" {
  depends_on = [netbox_service.test]
  filters = [
    { name = "name", value = "%[1]s" },
  ]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_service.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.#", "2"),
					resource.TestCheckResourceAttr("netbox_service.test", "ip_address_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_service.test", "parent_object_type", "dcim.device"),
					resource.TestCheckResourceAttrPair("netbox_service.test", "parent_object_id",
						"netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_service.test", "device_id",
						"netbox_device.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_service.test", "virtual_machine_id"),
					resource.TestCheckResourceAttrPair("data.netbox_service.test", "id", "netbox_service.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_services.test", "services.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_service.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: the address binding, description and comments clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_service" "test" {
  name      = "%[1]s"
  protocol  = "tcp"
  ports     = [80]
  device_id = netbox_device.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test", "ports.#", "1"),
					resource.TestCheckNoResourceAttr("netbox_service.test", "ip_address_ids"),
					resource.TestCheckNoResourceAttr("netbox_service.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_service.test", "comments"),
				),
			},
			{
				// Switching alias moves the service to the virtual machine and clears device_id.
				Config: deps + fmt.Sprintf(`
resource "netbox_service" "test" {
  name               = "%[1]s"
  protocol           = "tcp"
  ports              = [80]
  virtual_machine_id = netbox_virtual_machine.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test", "parent_object_type",
						"virtualization.virtualmachine"),
					resource.TestCheckResourceAttrPair("netbox_service.test", "virtual_machine_id",
						"netbox_virtual_machine.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_service.test", "device_id"),
				),
			},
			{
				// The pair can also be set directly; the matching alias is filled in from it.
				Config: deps + fmt.Sprintf(`
resource "netbox_service" "test" {
  name               = "%[1]s"
  protocol           = "tcp"
  ports              = [80]
  parent_object_type = "dcim.device"
  parent_object_id   = netbox_device.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_service.test", "device_id",
						"netbox_device.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_service.test", "virtual_machine_id"),
				),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_service" "test" {
  name               = "%[1]s"
  protocol           = "tcp"
  ports              = [80]
  device_id          = netbox_device.test.id
  virtual_machine_id = netbox_virtual_machine.test.id
}`, testName),
				ExpectError: regexp.MustCompile("Conflicting assignment"),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_service" "test" {
  name               = "%[1]s"
  protocol           = "tcp"
  ports              = [80]
  device_id          = netbox_device.test.id
  parent_object_type = "dcim.device"
  parent_object_id   = netbox_device.test.id
}`, testName),
				ExpectError: regexp.MustCompile("Conflicting assignment"),
			},
			{
				// NetBox rejects a service without a parent, so an empty assignment fails at plan time.
				Config: deps + fmt.Sprintf(`
resource "netbox_service" "test" {
  name     = "%[1]s"
  protocol = "tcp"
  ports    = [80]
}`, testName),
				ExpectError: regexp.MustCompile("Missing assignment"),
			},
		},
	})
}

// TestAccNetboxFHRPGroup_basic also covers netbox_fhrp_group_assignment, which needs a group and an
// interface to link and has no attributes of its own to clear.
func TestAccNetboxFHRPGroup_basic(t *testing.T) {
	testName := testAccGetTestName("fhrp")
	groupID := acctest.RandIntRange(100, 4000)
	deps := fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
}
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
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-eth0"
  type      = "virtual"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_fhrp_group" "test" {
  name        = "%[1]s"
  protocol    = "vrrp3"
  group_id    = %[2]d
  auth_type   = "plaintext"
  auth_key    = "%[1]s-key"
  description = "Acceptance test FHRP group."
  comments    = "Created by acceptance test."
}
resource "netbox_fhrp_group_assignment" "test" {
  group_id       = netbox_fhrp_group.test.id
  interface_type = "dcim.interface"
  interface_id   = netbox_device_interface.test.id
  priority       = 100
}
data "netbox_fhrp_group" "test" {
  depends_on = [netbox_fhrp_group.test]
  name       = "%[1]s"
}
data "netbox_fhrp_groups" "test" {
  # group_id alone is a random number, so two parallel tests can draw the same one: scope the count
  # by name too.
  depends_on = [netbox_fhrp_group.test]
  filters = [
    { name = "name", value = "%[1]s" },
    { name = "group_id", value = %[2]d },
  ]
}
data "netbox_fhrp_group_assignment" "test" {
  id = netbox_fhrp_group_assignment.test.id
}
data "netbox_fhrp_group_assignments" "test" {
  depends_on = [netbox_fhrp_group_assignment.test]
  filters = [
    { name = "group_id", value = netbox_fhrp_group.test.id },
  ]
}`, testName, groupID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "protocol", "vrrp3"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "group_id", fmt.Sprint(groupID)),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "auth_type", "plaintext"),
					resource.TestCheckResourceAttr("netbox_fhrp_group.test", "auth_key", testName+"-key"),
					resource.TestCheckResourceAttr("netbox_fhrp_group_assignment.test", "priority", "100"),
					resource.TestCheckResourceAttr("netbox_fhrp_group_assignment.test", "interface_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_fhrp_group_assignment.test", "group_id",
						"netbox_fhrp_group.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_fhrp_group.test", "id", "netbox_fhrp_group.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_fhrp_groups.test", "fhrp_groups.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_fhrp_group_assignment.test", "id",
						"netbox_fhrp_group_assignment.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_fhrp_group_assignments.test", "fhrp_group_assignments.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_fhrp_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_fhrp_group_assignment.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Shrink: name, auth and free-text fields clear on the group.
				Config: deps + fmt.Sprintf(`
resource "netbox_fhrp_group" "test" {
  protocol = "vrrp3"
  group_id = %[2]d
}`, testName, groupID),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_fhrp_group.test", "name"),
					resource.TestCheckNoResourceAttr("netbox_fhrp_group.test", "auth_type"),
					resource.TestCheckNoResourceAttr("netbox_fhrp_group.test", "auth_key"),
					resource.TestCheckNoResourceAttr("netbox_fhrp_group.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_fhrp_group.test", "comments"),
				),
			},
		},
	})
}

func TestAccNetboxServiceTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("svctmpl")
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_service_template" "test" {
  name        = "%[1]s"
  protocol    = "udp"
  ports       = [53, 5353]
  description = "Acceptance test service template."
  comments    = "Created by acceptance test."
}
data "netbox_service_template" "test" {
  name = netbox_service_template.test.name
}
data "netbox_service_templates" "test" {
  filters = [
    { name = "name__ic", value = "%[1]s" },
  ]
  depends_on    = [netbox_service_template.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service_template.test", "protocol", "udp"),
					resource.TestCheckResourceAttr("netbox_service_template.test", "ports.#", "2"),
					resource.TestCheckTypeSetElemAttr("netbox_service_template.test", "ports.*", "5353"),
					resource.TestCheckResourceAttrPair("data.netbox_service_template.test", "id", "netbox_service_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_service_templates.test", "service_templates.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_service_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_service_template" "test" {
  name     = "%[1]s"
  protocol = "tcp"
  ports    = [53]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service_template.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("netbox_service_template.test", "ports.#", "1"),
					resource.TestCheckNoResourceAttr("netbox_service_template.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_service_template.test", "comments"),
				),
			},
		},
	})
}
