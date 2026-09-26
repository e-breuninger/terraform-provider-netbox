//go:build acctest

// Devices, device interfaces, virtual device contexts and the netbox_primary_ip /
// netbox_device_oob_ip companions.
package provider_test

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/plancheck"
)

// deviceDeps is the catalog a device needs: site, manufacturer, device type and role, all named
// after the test.
func deviceDeps(testName string) string {
	return fmt.Sprintf(`
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
}`, testName)
}

func TestAccNetboxDevice_basic(t *testing.T) {
	testName := testAccGetTestName("device")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_platform" "test" {
  name = "%[1]s"
}
resource "netbox_location" "test" {
  name    = "%[1]s"
  site_id = netbox_site.test.id
}
resource "netbox_rack" "test" {
  name        = "%[1]s"
  site_id     = netbox_site.test.id
  location_id = netbox_location.test.id
}
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_config_template" "test" {
  name          = "%[1]s"
  template_code = "hostname {{ device.name }}"
}
resource "netbox_virtual_chassis" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device" "test" {
  name               = "%[1]s"
  device_type_id     = netbox_device_type.test.id
  role_id            = netbox_device_role.test.id
  site_id            = netbox_site.test.id
  location_id        = netbox_location.test.id
  rack_id            = netbox_rack.test.id
  rack_position      = 10.5
  rack_face          = "front"
  status             = "staged"
  airflow            = "front-to-rear"
  tenant_id          = netbox_tenant.test.id
  platform_id        = netbox_platform.test.id
  cluster_id         = netbox_cluster.test.id
  config_template_id = netbox_config_template.test.id
  virtual_chassis_id       = netbox_virtual_chassis.test.id
  virtual_chassis_position = 1
  virtual_chassis_priority = 10
  serial             = "SN-%[1]s"
  asset_tag          = "AT-%[1]s"
  local_context_data = jsonencode({ ntp = ["10.0.0.1"] })
  description        = "Acceptance test device."
  comments           = "Created by acceptance test."
}
data "netbox_device" "test" {
  name = netbox_device.test.name
}
data "netbox_devices" "test" {
  filters = [
    { name = "site_id", value = netbox_site.test.id },
  ]
  depends_on = [netbox_device.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_position", "10.5"),
					resource.TestCheckResourceAttr("netbox_device.test", "rack_face", "front"),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "staged"),
					resource.TestCheckResourceAttr("netbox_device.test", "airflow", "front-to-rear"),
					resource.TestCheckResourceAttr("netbox_device.test", "serial", "SN-"+testName),
					resource.TestCheckResourceAttr("netbox_device.test", "asset_tag", "AT-"+testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "rack_id", "netbox_rack.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "config_template_id", "netbox_config_template.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "virtual_chassis_id", "netbox_virtual_chassis.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "virtual_chassis_position", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "virtual_chassis_priority", "10"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "primary_ip4_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "oob_ip_id"),
					resource.TestCheckResourceAttrPair("data.netbox_device.test", "id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_device.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// local_context_data is normalize json: the state keeps the configured text although
				// NetBox echoes it compact, sorted and escaped (see prettyJSON).
				Config: deps + fmt.Sprintf(`
resource "netbox_device" "test" {
  device_type_id     = netbox_device_type.test.id
  role_id            = netbox_device_role.test.id
  site_id            = netbox_site.test.id
  local_context_data = <<EOT
%[1]sEOT
}`, prettyJSON),
				Check: resource.TestCheckResourceAttr("netbox_device.test", "local_context_data", prettyJSON),
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device" "test" {
  device_type_id     = netbox_device_type.test.id
  role_id            = netbox_device_role.test.id
  site_id            = netbox_site.test.id
  local_context_data = <<EOT
%[1]sEOT
}`, prettyJSON),
				ConfigPlanChecks: resource.ConfigPlanChecks{
					PreApply: []plancheck.PlanCheck{plancheck.ExpectEmptyPlan()},
				},
			},
			{
				// Shrink: everything optional clears (status is computed and stays), the name
				// included — NetBox allows nameless devices.
				Config: deps + `
resource "netbox_device" "test" {
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device.test", "name"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "rack_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "location_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "rack_position"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "rack_face"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "airflow"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "platform_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "cluster_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "config_template_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "virtual_chassis_id"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "virtual_chassis_position"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "virtual_chassis_priority"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "serial"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "asset_tag"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "local_context_data"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_device.test", "comments"),
				),
			},
		},
	})
}

func TestAccNetboxDeviceInterface_basic(t *testing.T) {
	testName := testAccGetTestName("devif")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_vlan" "test" {
  name = "%[1]s"
  vid  = 3998
}
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_device_interface" "lag" {
  device_id = netbox_device.test.id
  name      = "bond0"
  type      = "lag"
}
resource "netbox_device_interface" "br" {
  device_id = netbox_device.test.id
  name      = "br0"
  type      = "bridge"
}
resource "netbox_virtual_device_context" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s-vdc"
  status    = "active"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_device_interface" "test" {
  device_id        = netbox_device.test.id
  name             = "eth0"
  type             = "1000base-t"
  label            = "Port 1"
  enabled          = false
  mgmt_only        = true
  mark_connected   = true
  mtu              = 9000
  speed            = 1000000
  duplex           = "full"
  mode             = "tagged"
  untagged_vlan_id = netbox_vlan.test.id
  tagged_vlan_ids  = [netbox_vlan.test.id]
  lag_device_interface_id = netbox_device_interface.lag.id
  bridge_id        = netbox_device_interface.br.id
  vdc_ids          = [netbox_virtual_device_context.test.id]
  wwn              = "00:11:22:33:44:55:66:77"
  vrf_id           = netbox_vrf.test.id
  description      = "Acceptance test interface."
}
resource "netbox_device_interface" "sub" {
  device_id                  = netbox_device.test.id
  name                       = "eth0.100"
  type                       = "virtual"
  parent_device_interface_id = netbox_device_interface.test.id
}
data "netbox_device_interface" "test" {
  id = netbox_device_interface.test.id
}
data "netbox_device_interfaces" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_device_interface.test, netbox_device_interface.sub]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "name", "eth0"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "type", "1000base-t"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mgmt_only", "true"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mark_connected", "true"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mtu", "9000"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "speed", "1000000"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "duplex", "full"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mode", "tagged"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "tagged_vlan_ids.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "lag_device_interface_id", "netbox_device_interface.lag", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "bridge_id", "netbox_device_interface.br", "id"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "vdc_ids.#", "1"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "wwn", "00:11:22:33:44:55:66:77"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_interface.sub", "parent_device_interface_id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_device_interface.test", "id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_interfaces.test", "device_interfaces.#", "4"),
				),
			},
			{
				ResourceName:      "netbox_device_interface.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Type change and shrink. The sub-interface stays but loses its parent.
				Config: deps + `
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "eth0"
  type      = "virtual"
}
resource "netbox_device_interface" "sub" {
  device_id = netbox_device.test.id
  name      = "eth0.100"
  type      = "virtual"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "type", "virtual"),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "enabled", "true"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "label"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "mtu"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "speed"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "duplex"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "mode"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "untagged_vlan_id"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "tagged_vlan_ids"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "lag_device_interface_id"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "bridge_id"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "vdc_ids"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "wwn"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.test", "vrf_id"),
					resource.TestCheckNoResourceAttr("netbox_device_interface.sub", "parent_device_interface_id"),
				),
			},
		},
	})
}

// TestAccNetboxPrimaryIP_basic: the companion resource sets and clears
// primary_ip4/primary_ip6 on a device and on a VM.
// netbox_primary_ip has no shrink step: it is a pure link resource with no optional
// attributes to clear, so unsetting it means destroying it — which the post-test destroy already
// exercises (it PATCHes primary_ip4/6 back to null on the device).
func TestAccNetboxPrimaryIP_basic(t *testing.T) {
	testName := testAccGetTestName("primip")
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
  ip_address          = "10.0.0.71/24"
  device_interface_id = netbox_device_interface.test.id
}
resource "netbox_ip_address" "v6" {
  ip_address          = "2001:db8::71/64"
  device_interface_id = netbox_device_interface.test.id
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
}
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "eth0"
}
resource "netbox_ip_address" "vm" {
  ip_address                   = "10.0.0.72/24"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_primary_ip" "test" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.v4.id
}
resource "netbox_primary_ip" "vm" {
  virtual_machine_id = netbox_virtual_machine.test.id
  ip_address_id      = netbox_ip_address.vm.id
}
data "netbox_device" "test" {
  id         = netbox_device.test.id
  depends_on = [netbox_primary_ip.test]
}
data "netbox_virtual_machine" "test" {
  id         = netbox_virtual_machine.test.id
  depends_on = [netbox_primary_ip.vm]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_primary_ip.test", "ip_address_version", "4"),
					resource.TestCheckResourceAttrPair("netbox_primary_ip.test", "ip_address_id", "netbox_ip_address.v4", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_device.test", "primary_ip4_id", "netbox_ip_address.v4", "id"),
					resource.TestCheckResourceAttr("netbox_primary_ip.vm", "ip_address_version", "4"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machine.test", "primary_ip4_id", "netbox_ip_address.vm", "id"),
				),
			},
			{
				ResourceName:      "netbox_primary_ip.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				// Switch to the IPv6 address: primary_ip6 is set and primary_ip4 cleared.
				Config: deps + `
resource "netbox_primary_ip" "test" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.v6.id
}
data "netbox_device" "test" {
  id         = netbox_device.test.id
  depends_on = [netbox_primary_ip.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_primary_ip.test", "ip_address_version", "6"),
					resource.TestCheckResourceAttrPair("data.netbox_device.test", "primary_ip6_id", "netbox_ip_address.v6", "id"),
					resource.TestCheckNoResourceAttr("data.netbox_device.test", "primary_ip4_id"),
				),
			},
			{
				// Both targets at once is a config error.
				Config: deps + `
resource "netbox_primary_ip" "test" {
  device_id          = netbox_device.test.id
  virtual_machine_id = netbox_virtual_machine.test.id
  ip_address_id      = netbox_ip_address.v6.id
}`,
				ExpectError: regexp.MustCompile("Invalid Attribute Combination"),
			},
		},
	})
}

// TestAccNetboxDeviceOOBIP_basic: the companion resource sets and moves oob_ip. It has no shrink
// step: it is a pure link resource, so unsetting it means destroying it, which the post-test
// destroy already exercises.
func TestAccNetboxDeviceOOBIP_basic(t *testing.T) {
	testName := testAccGetTestName("oobip")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "mgmt0"
  type      = "1000base-t"
  mgmt_only = true
}
resource "netbox_ip_address" "a" {
  ip_address          = "10.0.0.81/24"
  device_interface_id = netbox_device_interface.test.id
}
resource "netbox_ip_address" "b" {
  ip_address          = "10.0.0.82/24"
  device_interface_id = netbox_device_interface.test.id
}
data "netbox_device" "test" {
  id         = netbox_device.test.id
  depends_on = [netbox_device_oob_ip.test]
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
resource "netbox_device_oob_ip" "test" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.a.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_oob_ip.test", "ip_address_id", "netbox_ip_address.a", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_device.test", "oob_ip_id", "netbox_ip_address.a", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_oob_ip.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + `
resource "netbox_device_oob_ip" "test" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.b.id
}`,
				Check: resource.TestCheckResourceAttrPair("data.netbox_device.test", "oob_ip_id", "netbox_ip_address.b", "id"),
			},
		},
	})
}

func init() {
	sweep("netbox_device_interface",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimInterfacesList(dcim.NewDcimInterfacesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				// Interfaces are named eth0 etc.; key on the device's name.
				name := ""
				if result.Device != nil {
					name = deref(result.Device.Name)
				}
				items = append(items, sweepItem{id: result.ID, name: name})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimInterfacesDestroy(dcim.NewDcimInterfacesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_device",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Dcim.DcimDevicesList(dcim.NewDcimDevicesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{id: result.ID, name: deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Dcim.DcimDevicesDestroy(dcim.NewDcimDevicesDestroyParams().WithID(id), nil)
			return err
		})
}

// TestAccNetboxVirtualDeviceContext_basic: the primary address must be assigned to an interface of
// the device, so the context is created after the interface and its address.
func TestAccNetboxVirtualDeviceContext_basic(t *testing.T) {
	testName := testAccGetTestName("vdc")
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
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "eth0"
  type      = "1000base-t"
}
resource "netbox_ip_address" "test" {
  ip_address          = "10.0.7.11/24"
  device_interface_id = netbox_device_interface.test.id
}
resource "netbox_ip_address" "v6" {
  ip_address          = "2001:db8:7::11/64"
  device_interface_id = netbox_device_interface.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_device_context" "test" {
  device_id      = netbox_device.test.id
  name           = "%[1]s"
  identifier     = 10
  status         = "planned"
  tenant_id      = netbox_tenant.test.id
  primary_ip4_id = netbox_ip_address.test.id
  primary_ip6_id = netbox_ip_address.v6.id
  description    = "Acceptance test context."
  comments       = "Created by acceptance test."
}
data "netbox_virtual_device_context" "test" {
  name = netbox_virtual_device_context.test.name
}
data "netbox_virtual_device_contexts" "test" {
  filters = [
    { name = "device_id", value = netbox_device.test.id },
  ]
  depends_on = [netbox_virtual_device_context.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "identifier", "10"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "status", "planned"),
					resource.TestCheckResourceAttrPair("netbox_virtual_device_context.test", "primary_ip4_id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_device_context.test", "primary_ip6_id", "netbox_ip_address.v6", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_device_context.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "interface_count", "0"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_device_context.test", "id", "netbox_virtual_device_context.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_device_contexts.test", "virtual_device_contexts.#", "1"),
				),
			},
			{
				ResourceName:      "netbox_virtual_device_context.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_device_context" "test" {
  device_id = netbox_device.test.id
  name      = "%[1]s"
  status    = "active"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_device_context.test", "status", "active"),
					resource.TestCheckNoResourceAttr("netbox_virtual_device_context.test", "identifier"),
					resource.TestCheckNoResourceAttr("netbox_virtual_device_context.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_device_context.test", "primary_ip4_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_device_context.test", "primary_ip6_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_device_context.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_virtual_device_context.test", "comments"),
				),
			},
		},
	})
}

// TestAccNetboxDeviceRole_hierarchy covers the device role attributes the catalog table does not
// (the parent role, the default config template, color_hex and vm_role) and the platform's
// manufacturer_id.
func TestAccNetboxDeviceRole_hierarchy(t *testing.T) {
	testName := testAccGetTestName("rolehier")
	deps := fmt.Sprintf(`
resource "netbox_device_role" "parent" {
  name = "%[1]s-parent"
}
resource "netbox_config_template" "test" {
  name          = "%[1]s"
  template_code = "hostname {{ device.name }}"
}
resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_role" "test" {
  name               = "%[1]s"
  slug               = "%[2]s"
  color_hex          = "112233"
  vm_role            = true
  parent_id          = netbox_device_role.parent.id
  config_template_id = netbox_config_template.test.id
}
resource "netbox_platform" "test" {
  name            = "%[1]s"
  slug            = "%[2]s"
  manufacturer_id = netbox_manufacturer.test.id
}`, testName, getSlug(testName)),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_role.test", "parent_id", "netbox_device_role.parent", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_role.test", "config_template_id", "netbox_config_template.test", "id"),
					resource.TestCheckResourceAttr("netbox_device_role.test", "color_hex", "112233"),
					resource.TestCheckResourceAttr("netbox_device_role.test", "vm_role", "true"),
					resource.TestCheckResourceAttrPair("netbox_platform.test", "manufacturer_id", "netbox_manufacturer.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_device_role.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_device_role" "test" {
  name = "%[1]s"
}
resource "netbox_platform" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_device_role.test", "parent_id"),
					resource.TestCheckNoResourceAttr("netbox_device_role.test", "config_template_id"),
					resource.TestCheckNoResourceAttr("netbox_platform.test", "manufacturer_id"),
				),
			},
		},
	})
}

// TestAccNetboxDeviceRenderConfig_basic renders a device through its own config template, then
// through a second template named explicitly.
func TestAccNetboxDeviceRenderConfig_basic(t *testing.T) {
	testName := testAccGetTestName("render_config")
	deps := deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_config_template" "test" {
  name          = "%[1]s"
  template_code = "hostname {{ device.name }}"
}
resource "netbox_config_template" "other" {
  name          = "%[1]s-other"
  template_code = "! {{ device.name }} other"
}
resource "netbox_device" "test" {
  name               = "%[1]s"
  device_type_id     = netbox_device_type.test.id
  role_id            = netbox_device_role.test.id
  site_id            = netbox_site.test.id
  config_template_id = netbox_config_template.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + `
data "netbox_device_render_config" "own" {
  device_id = netbox_device.test.id
}
data "netbox_device_render_config" "other" {
  device_id          = netbox_device.test.id
  config_template_id = netbox_config_template.other.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_device_render_config.own", "content", "hostname "+testName),
					resource.TestCheckResourceAttrPair("data.netbox_device_render_config.own", "config_template_id", "netbox_config_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_device_render_config.own", "config_template_name", testName),
					resource.TestCheckResourceAttr("data.netbox_device_render_config.other", "content", "! "+testName+" other"),
					resource.TestCheckResourceAttrPair("data.netbox_device_render_config.other", "config_template_id", "netbox_config_template.other", "id"),
				),
			},
		},
	})
}
