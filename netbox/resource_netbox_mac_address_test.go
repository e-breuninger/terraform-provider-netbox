package netbox

import (
	"fmt"
	"log"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxMACAddressFullVmDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_vrf" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_interface" "test" {
  name = "%[1]s"
  virtual_machine_id = netbox_virtual_machine.test.id
}`, testName)
}

func testAccNetboxMACAddressFullDeviceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
}

resource "netbox_device_interface" "test" {
  name = "%[1]s"
  device_id = netbox_device.test.id
  type = "1000base-t"
}`, testName)
}

func TestAccNetboxMACAddress_standalone(t *testing.T) {
	testSlug := "mac-addr"
	macAddress := "00:1A:2B:3C:4D:5E"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%[1]s"
  description = "%[2]s"
  comments    = "%[2]s"
}
`, macAddress, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", macAddress),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "description", testSlug),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "comments", testSlug),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
}`, macAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", macAddress),
				),
			},
			{
				ResourceName:      "netbox_mac_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxMACAddress_customFields(t *testing.T) {
	testSlug := "mac-addr_cf"
	macAddress := "00:1A:2B:3F:4D:5E"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name          = "mac_custom_field"
  type          = "text"
  content_types = ["dcim.macaddress"]
}

resource "netbox_mac_address" "test" {
  mac_address   = "%[1]s"
  description   = "%[2]s"
  comments      = "%[2]s"
  custom_fields = {"${netbox_custom_field.test.name}" = "foomac"}
}
`, macAddress, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "custom_fields.mac_custom_field", "foomac"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
}`, macAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", macAddress),
				),
			},
			{
				ResourceName:      "netbox_mac_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxMACAddress_deviceByFieldName(t *testing.T) {
	testSlug := "mac-addr-dev-fn"
	macAddress := "01:1A:2B:3C:4D:5E"
	macAddress2 := "01:1A:2B:7C:4D:5E"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxMACAddressFullDeviceDependencies(testSlug) + fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%[1]s"
  device_interface_id = netbox_device_interface.test.id
  description = "%[3]s"
}

resource "netbox_mac_address" "test2" {
  mac_address = "%[2]s"
  device_interface_id = netbox_device_interface.test.id
  description = "%[3]s"
  comments    = "%[3]s"
}`, macAddress, macAddress2, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", macAddress),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "description", testSlug),
					resource.TestCheckResourceAttrPair("netbox_mac_address.test", "device_interface_id", "netbox_device_interface.test", "id"),
				),
			},
			{
				RefreshState: true,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mac_address", macAddress),
					resource.TestCheckResourceAttr("netbox_device_interface.test", "mac_addresses.#", "2"),
					resource.TestCheckResourceAttrSet("netbox_device_interface.test", "mac_addresses.0.id"),
					resource.TestCheckResourceAttrSet("netbox_device_interface.test", "mac_addresses.0.mac_address"),
					resource.TestCheckResourceAttrSet("netbox_device_interface.test", "mac_addresses.0.description"),
				),
			},
			{
				ResourceName:            "netbox_mac_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"device_interface_id"},
			},
		},
	})
}

func TestAccNetboxMACAddress_vmByFieldName(t *testing.T) {
	testSlug := "mac-addr-vm-fn"
	macAddress := "02:1A:2B:3C:4D:5E"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxMACAddressFullVmDependencies(testSlug) + fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  description = "%s"
}`, macAddress, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", macAddress),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "description", testSlug),
					resource.TestCheckResourceAttrPair("netbox_mac_address.test", "virtual_machine_interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_mac_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"virtual_machine_interface_id"},
			},
		},
	})
}

func TestAccNetboxMACAddress_deviceByObjectType(t *testing.T) {
	testSlug := "mac-addr-dev-ot"
	macAddress := "03:1A:2B:3C:4D:5E"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxMACAddressFullDeviceDependencies(testSlug) + fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
  object_type = "dcim.interface"
  interface_id = netbox_device_interface.test.id
  description = "%s"
}`, macAddress, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", macAddress),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "description", testSlug),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_mac_address.test", "interface_id", "netbox_device_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_mac_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type"},
			},
		},
	})
}

func TestAccNetboxMACAddress_vmByObjectType(t *testing.T) {
	testSlug := "mac-addr-vm-ot"
	macAddress := "04:1A:2B:3C:4D:5E"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxMACAddressFullVmDependencies(testSlug) + fmt.Sprintf(`
resource "netbox_mac_address" "test" {
  mac_address = "%s"
  object_type = "virtualization.vminterface"
  interface_id = netbox_interface.test.id
  description = "%s"
}`, macAddress, testSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_mac_address.test", "mac_address", macAddress),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "description", testSlug),
					resource.TestCheckResourceAttr("netbox_mac_address.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_mac_address.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_mac_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type"},
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_mac_address", &resource.Sweeper{
		Name:         "netbox_mac_address",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimMacAddressesListParams()
			res, err := api.Dcim.DcimMacAddressesList(params, nil)
			if err != nil {
				return err
			}
			for _, macAddress := range res.GetPayload().Results {
				if len(macAddress.Tags) > 0 && (macAddress.Tags[0] == &models.NestedTag{Name: strToPtr("acctest"), Slug: strToPtr("acctest")}) {
					deleteParams := dcim.NewDcimMacAddressesDeleteParams().WithID(macAddress.ID)
					_, err := api.Dcim.DcimMacAddressesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a mac address")
				}
			}
			return nil
		},
	})
}
