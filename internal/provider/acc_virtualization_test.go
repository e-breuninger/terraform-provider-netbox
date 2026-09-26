//go:build acctest

package provider_test

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxVirtualMachine_basic(t *testing.T) {
	testName := testAccGetTestName("vm")
	deps := fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}
resource "netbox_cluster" "test" {
  name            = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}
resource "netbox_device_role" "test" {
  name      = "%[1]s"
  color_hex = "112233"
  vm_role   = true
}
resource "netbox_tenant" "test" {
  name = "%[1]s"
}
resource "netbox_platform" "test" {
  name = "%[1]s"
}
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
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name         = "%[1]s"
  cluster_id   = netbox_cluster.test.id
  status       = "staged"
  local_context_data = jsonencode({ env = "test" })
  role_id      = netbox_device_role.test.id
  tenant_id    = netbox_tenant.test.id
  platform_id  = netbox_platform.test.id
  vcpus        = 2.5
  memory_mb    = 2048
  disk_size_mb = 10240
  description  = "Acceptance test VM."
  comments     = "Created by acceptance test."
}
data "netbox_virtual_machine" "test" {
  name = netbox_virtual_machine.test.name
}
data "netbox_virtual_machines" "test" {
  filters = [
    { name = "cluster_id", value = netbox_cluster.test.id },
  ]
  depends_on = [netbox_virtual_machine.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "status", "staged"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "local_context_data", `{"env":"test"}`),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "2.5"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "memory_mb", "2048"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "disk_size_mb", "10240"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machine.test", "id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "virtual_machines.#", "1"),
				),
			},
			{
				// Shrink to the minimum: every optional attribute clears (status and disk_size_mb are computed
				// and stay).
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name       = "%s"
  cluster_id = netbox_cluster.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "role_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "tenant_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "platform_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "vcpus"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "memory_mb"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "description"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "comments"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "local_context_data"),
				),
			},
			{
				// Move the VM off the cluster onto a site with a host device.
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name      = "%s"
  site_id   = netbox_site.test.id
  device_id = netbox_device.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "device_id", "netbox_device.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "cluster_id"),
				),
			},
			{
				// And back onto the cluster: site and device clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name       = "%s"
  cluster_id = netbox_cluster.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "site_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine.test", "device_id"),
				),
			},
			{
				ResourceName:      "netbox_virtual_machine.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVMInterface_basic(t *testing.T) {
	testName := testAccGetTestName("vmif")
	deps := fmt.Sprintf(`
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
resource "netbox_vlan" "untagged" {
  name = "%[1]s-untagged"
  vid  = 811
}
resource "netbox_vlan" "tagged1" {
  name = "%[1]s-tagged1"
  vid  = 812
}
resource "netbox_vlan" "tagged2" {
  name = "%[1]s-tagged2"
  vid  = 813
}
resource "netbox_vrf" "test" {
  name = "%[1]s"
}
resource "netbox_virtual_machine_interface" "parent" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%[1]s-parent"
}
resource "netbox_virtual_machine_interface" "bridge" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%[1]s-bridge"
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%[1]s"
  parent_id          = netbox_virtual_machine_interface.parent.id
  bridge_id          = netbox_virtual_machine_interface.bridge.id
  mtu                = 1500
  mode               = "tagged"
  untagged_vlan_id   = netbox_vlan.untagged.id
  tagged_vlan_ids    = [netbox_vlan.tagged1.id, netbox_vlan.tagged2.id]
  vrf_id             = netbox_vrf.test.id
  description        = "Acceptance test interface."
}
data "netbox_virtual_machine_interface" "test" {
  id = netbox_virtual_machine_interface.test.id
}
data "netbox_virtual_machine_interfaces" "test" {
  filters = [
    { name = "virtual_machine_id", value = netbox_virtual_machine.test.id },
    { name = "virtual_machine", value = netbox_virtual_machine.test.name },
    { name = "enabled", value = "true" },
  ]
  depends_on         = [netbox_virtual_machine_interface.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine_interface.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_machine_interface.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_interface.test", "mtu", "1500"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_interface.test", "mode", "tagged"),
					resource.TestCheckResourceAttr("netbox_virtual_machine_interface.test", "tagged_vlan_ids.#", "2"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface.test", "untagged_vlan_id", "netbox_vlan.untagged", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface.test", "parent_id", "netbox_virtual_machine_interface.parent", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine_interface.test", "bridge_id", "netbox_virtual_machine_interface.bridge", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machine_interface.test", "id", "netbox_virtual_machine_interface.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machine_interfaces.test", "virtual_machine_interfaces.#", "3"),
				),
			},
			{
				// Disable and shrink: mode, MTU, VLANs, VRF and the description clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%s"
  enabled            = false
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine_interface.test", "enabled", "false"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "mtu"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "parent_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "bridge_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "mode"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "untagged_vlan_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "tagged_vlan_ids.#"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "vrf_id"),
					resource.TestCheckNoResourceAttr("netbox_virtual_machine_interface.test", "description"),
				),
			},
			{
				ResourceName:      "netbox_virtual_machine_interface.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVirtualDisk_basic(t *testing.T) {
	testName := testAccGetTestName("vdisk")
	deps := fmt.Sprintf(`
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
resource "netbox_virtual_disk" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%[1]s"
  size_mb            = 20480
  description        = "Acceptance test disk."
}
data "netbox_virtual_disk" "test" {
  id = netbox_virtual_disk.test.id
}
data "netbox_virtual_disks" "test" {
  filters = [
    { name = "virtual_machine_id", value = netbox_virtual_machine.test.id },
  ]
  depends_on         = [netbox_virtual_disk.test]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_disk.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_virtual_disk.test", "size_mb", "20480"),
					resource.TestCheckResourceAttrPair("netbox_virtual_disk.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_disk.test", "id", "netbox_virtual_disk.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_virtual_disks.test", "virtual_disks.#", "1"),
				),
			},
			{
				// Resize and drop the description: it must clear.
				Config: deps + fmt.Sprintf(`
resource "netbox_virtual_disk" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "%s"
  size_mb            = 40960
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_disk.test", "size_mb", "40960"),
					resource.TestCheckNoResourceAttr("netbox_virtual_disk.test", "description"),
				),
			},
			{
				ResourceName:      "netbox_virtual_disk.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	sweep("netbox_virtual_disk",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Virtualization.VirtualizationVirtualDisksList(virtualization.NewVirtualizationVirtualDisksListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Virtualization.VirtualizationVirtualDisksDestroy(virtualization.NewVirtualizationVirtualDisksDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_virtual_machine_interface",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Virtualization.VirtualizationInterfacesList(virtualization.NewVirtualizationInterfacesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Virtualization.VirtualizationInterfacesDestroy(virtualization.NewVirtualizationInterfacesDestroyParams().WithID(id), nil)
			return err
		})
	sweep("netbox_virtual_machine",
		func(client *client.NetBoxAPI) ([]sweepItem, error) {
			res, err := client.Virtualization.VirtualizationVirtualMachinesList(virtualization.NewVirtualizationVirtualMachinesListParams(), nil)
			if err != nil {
				return nil, err
			}
			var items []sweepItem
			for _, result := range res.GetPayload().Results {
				items = append(items, sweepItem{result.ID, deref(result.Name)})
			}
			return items, nil
		},
		func(client *client.NetBoxAPI, id int64) error {
			_, err := client.Virtualization.VirtualizationVirtualMachinesDestroy(virtualization.NewVirtualizationVirtualMachinesDestroyParams().WithID(id), nil)
			return err
		})
}
