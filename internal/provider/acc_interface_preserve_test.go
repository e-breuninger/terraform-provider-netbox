//go:build acctest

// Writes must leave fields alone that the resource does not own: go-netbox's writable models carry
// primary_mac_address, wireless_lans, primary_ip4 and the like without omitempty, and sending them
// as null or [] on an unrelated update would clear what a companion resource or an operator set
// out of band (issue #961).
package provider_test

import (
	"fmt"
	"testing"

	"github.com/e-breuninger/terraform-provider-netbox/internal/netboxapi"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/fbreckle/go-netbox/netbox/client/wireless"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
	"github.com/hashicorp/terraform-plugin-testing/terraform"
)

// patchOutOfBand PATCHes a device (virtualMachine=false) or virtual machine interface with a body
// the provider never sends itself.
func patchOutOfBand(t *testing.T, virtualMachine bool, id int64, body map[string]any) {
	t.Helper()
	client, err := testAccClient()
	if err != nil {
		t.Fatal(err)
	}
	if virtualMachine {
		params := virtualization.NewVirtualizationInterfacesPartialUpdateParams().WithID(id).WithData(&models.WritableVMInterface{})
		_, err = client.Virtualization.VirtualizationInterfacesPartialUpdate(params, nil, netboxapi.WithBody(body))
	} else {
		params := dcim.NewDcimInterfacesPartialUpdateParams().WithID(id).WithData(&models.WritableInterface{})
		_, err = client.Dcim.DcimInterfacesPartialUpdate(params, nil, netboxapi.WithBody(body))
	}
	if err != nil {
		t.Fatalf("out-of-band PATCH: %v", err)
	}
}

func checkDeviceInterface(name string, check func(*models.Interface) error) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		var id int64
		if err := stateID(name, &id)(state); err != nil {
			return err
		}
		client, err := testAccClient()
		if err != nil {
			return err
		}
		res, err := client.Dcim.DcimInterfacesRetrieve(dcim.NewDcimInterfacesRetrieveParams().WithID(id), nil)
		if err != nil {
			return err
		}
		return check(res.Payload)
	}
}

func checkVMInterface(name string, check func(*models.VMInterface) error) resource.TestCheckFunc {
	return func(state *terraform.State) error {
		var id int64
		if err := stateID(name, &id)(state); err != nil {
			return err
		}
		client, err := testAccClient()
		if err != nil {
			return err
		}
		res, err := client.Virtualization.VirtualizationInterfacesRetrieve(virtualization.NewVirtualizationInterfacesRetrieveParams().WithID(id), nil)
		if err != nil {
			return err
		}
		return check(res.Payload)
	}
}

// TestAccNetboxDeviceInterface_preservesPrimaryMACAndWirelessLANs: a description change keeps the
// primary MAC and wireless LAN assignments made out of band (neither is in the schema).
func TestAccNetboxDeviceInterface_preservesPrimaryMACAndWirelessLANs(t *testing.T) {
	testName := testAccGetTestName("if_preserve")
	client, err := testAccClient()
	if err != nil {
		t.Fatal(err)
	}
	// The body is written directly: the typed model sends tags as null, which NetBox rejects.
	lan, err := client.Wireless.WirelessWirelessLansCreate(wireless.NewWirelessWirelessLansCreateParams().WithData(&models.WritableWirelessLAN{}), nil,
		netboxapi.WithBody(map[string]any{"ssid": testName, "tags": []any{}}))
	if err != nil {
		t.Fatal(err)
	}
	lanID := lan.Payload.ID
	t.Cleanup(func() {
		_, _ = client.Wireless.WirelessWirelessLansDestroy(wireless.NewWirelessWirelessLansDestroyParams().WithID(lanID), nil)
	})
	var ifID, macID int64
	config := func(description string) string {
		return deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
}
resource "netbox_device_interface" "test" {
  device_id   = netbox_device.test.id
  name        = "wlan0"
  type        = "ieee802.11a"
  description = "%[2]s"
}
resource "netbox_mac_address" "test" {
  mac_address         = "02:42:ac:11:00:01"
  device_interface_id = netbox_device_interface.test.id
}`, testName, description)
	}
	preserved := checkDeviceInterface("netbox_device_interface.test", func(i *models.Interface) error {
		if i.PrimaryMacAddress == nil || i.PrimaryMacAddress.ID != macID {
			return fmt.Errorf("primary_mac_address = %v, want id %d", i.PrimaryMacAddress, macID)
		}
		if len(i.WirelessLans) != 1 || i.WirelessLans[0].ID != lanID {
			return fmt.Errorf("wireless_lans = %v, want [%d]", i.WirelessLans, lanID)
		}
		return nil
	})
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config("before"),
				Check: resource.ComposeTestCheckFunc(
					stateID("netbox_device_interface.test", &ifID),
					stateID("netbox_mac_address.test", &macID),
				),
			},
			{
				PreConfig: func() {
					patchOutOfBand(t, false, ifID, map[string]any{"primary_mac_address": macID, "wireless_lans": []int64{lanID}})
				},
				Config: config("after"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "description", "after"),
					preserved,
				),
			},
			{
				// A no-op plan must not touch them either.
				Config: config("after"),
				Check:  preserved,
			},
		},
	})
}

// TestAccNetboxVirtualMachineInterface_preservesPrimaryMAC: the virtualization counterpart.
func TestAccNetboxVirtualMachineInterface_preservesPrimaryMAC(t *testing.T) {
	testName := testAccGetTestName("vmif_preserve")
	var ifID, macID int64
	config := func(description string) string {
		return fmt.Sprintf(`
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
  description        = "%[2]s"
}
resource "netbox_mac_address" "test" {
  mac_address                  = "02:42:ac:11:00:02"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}`, testName, description)
	}
	preserved := checkVMInterface("netbox_virtual_machine_interface.test", func(i *models.VMInterface) error {
		if i.PrimaryMacAddress == nil || i.PrimaryMacAddress.ID != macID {
			return fmt.Errorf("primary_mac_address = %v, want id %d", i.PrimaryMacAddress, macID)
		}
		return nil
	})
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: config("before"),
				Check: resource.ComposeTestCheckFunc(
					stateID("netbox_virtual_machine_interface.test", &ifID),
					stateID("netbox_mac_address.test", &macID),
				),
			},
			{
				PreConfig: func() {
					patchOutOfBand(t, true, ifID, map[string]any{"primary_mac_address": macID})
				},
				Config: config("after"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine_interface.test", "description", "after"),
					preserved,
				),
			},
		},
	})
}

// TestAccNetboxPrimaryIP_survivesParentUpdate: updating the device or virtual machine that
// netbox_primary_ip manages keeps the primary IP (primary_ip4 is computed-only on both).
func TestAccNetboxPrimaryIP_survivesParentUpdate(t *testing.T) {
	testName := testAccGetTestName("primip_keep")
	config := func(comments string) string {
		return deviceDeps(testName) + fmt.Sprintf(`
resource "netbox_device" "test" {
  name           = "%[1]s"
  device_type_id = netbox_device_type.test.id
  role_id        = netbox_device_role.test.id
  site_id        = netbox_site.test.id
  comments       = "%[2]s"
}
resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name      = "eth0"
  type      = "1000base-t"
}
resource "netbox_ip_address" "dev" {
  ip_address          = "10.0.0.91/24"
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
  comments   = "%[2]s"
}
resource "netbox_virtual_machine_interface" "test" {
  virtual_machine_id = netbox_virtual_machine.test.id
  name               = "eth0"
}
resource "netbox_ip_address" "vm" {
  ip_address                   = "10.0.0.92/24"
  virtual_machine_interface_id = netbox_virtual_machine_interface.test.id
}
resource "netbox_primary_ip" "dev" {
  device_id     = netbox_device.test.id
  ip_address_id = netbox_ip_address.dev.id
}
resource "netbox_primary_ip" "vm" {
  virtual_machine_id = netbox_virtual_machine.test.id
  ip_address_id      = netbox_ip_address.vm.id
}
data "netbox_device" "test" {
  id         = netbox_device.test.id
  depends_on = [netbox_primary_ip.dev]
}
data "netbox_virtual_machine" "test" {
  id         = netbox_virtual_machine.test.id
  depends_on = [netbox_primary_ip.vm]
}`, testName, comments)
	}
	preserved := resource.ComposeTestCheckFunc(
		resource.TestCheckResourceAttrPair("data.netbox_device.test", "primary_ip4_id", "netbox_ip_address.dev", "id"),
		resource.TestCheckResourceAttrPair("data.netbox_virtual_machine.test", "primary_ip4_id", "netbox_ip_address.vm", "id"),
	)
	resource.ParallelTest(t, resource.TestCase{
		ProtoV6ProviderFactories: testAccProtoV6ProviderFactories,
		PreCheck:                 func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{Config: config("before"), Check: preserved},
			{
				// The device and VM are updated after the primary IP is set; the data sources re-read
				// afterwards because the companion resources depend on their targets.
				Config: config("after"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "after"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "comments", "after"),
					preserved,
				),
			},
		},
	})
}
