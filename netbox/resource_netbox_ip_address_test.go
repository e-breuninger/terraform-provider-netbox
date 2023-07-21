package netbox

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxIPAddressFullDependencies(testName string) string {
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
}
`, testName)
}

func testAccNetboxIPAddressFullDeviceDependencies(testName string) string {
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
}
`, testName)
}

func TestAccNetboxIPAddress_basic(t *testing.T) {
	testIP := "1.1.1.1/32"
	testSlug := "ipaddress"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  object_type = "virtualization.vminterface"
  status = "active"
  tags = [netbox_tag.test.name]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tenant_id", "0"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "vrf_id", "0"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  object_type = "virtualization.vminterface"
  status = "reserved"
  tenant_id = netbox_tenant.test.id
  vrf_id = netbox_vrf.test.id
  tags = [netbox_tag.test.name]
  description = "description for %[1]s"
  role = "loopback"
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "reserved"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "description", fmt.Sprintf("description for %[1]s", testIP)),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "role", "loopback"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  object_type = "virtualization.vminterface"
  status = "dhcp"
  tags = [netbox_tag.test.name]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "dhcp"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tenant_id", "0"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "vrf_id", "0"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "role", ""),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  object_type = "virtualization.vminterface"
  status = "provoke_error"
  tags = [netbox_tag.test.name]
}`, testIP),
				ExpectError: regexp.MustCompile("expected status to be one of .*"),
			},
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  object_type = "virtualization.vminterface"
  status = "deprecated"
  tags = [netbox_tag.test.name]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "deprecated"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  object_type = "virtualization.vminterface"
  status = "active"
  dns_name = "mytest.example.com"
  tags = [netbox_tag.test.name]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "dns_name", "mytest.example.com"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type"},
			},
		},
	})
}

func TestAccNetboxIPAddress_deviceByObjectType(t *testing.T) {
	testIP := "1.1.1.2/32"
	testSlug := "ipadr_dev_ot"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDeviceDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  object_type = "dcim.interface"
  interface_id = netbox_device_interface.test.id
  status = "active"
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_device_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type"},
			},
		},
	})
}

func TestAccNetboxIPAddress_vmByObjectType(t *testing.T) {
	testIP := "1.1.1.3/32"
	testSlug := "ipadr_vm_ot"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  object_type = "virtualization.vminterface"
  interface_id = netbox_interface.test.id
  status = "active"
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type"},
			},
		},
	})
}

func TestAccNetboxIPAddress_vmSwitchStyle(t *testing.T) {
	testIP := "1.1.1.9/32"
	testSlug := "ipadr_vm_sw"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  object_type = "virtualization.vminterface"
  interface_id = netbox_interface.test.id
  status = "active"
}`, testIP),
			},
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
}`, testIP),
			},
			{
				ResourceName:            "netbox_ip_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type", "virtual_machine_interface_id"},
			},
		},
	})
}

// TestAccNetboxIPAddress_deviceByFieldName tests if creating an ip address and linking it to a device via the `device_interface_id` field works
func TestAccNetboxIPAddress_deviceByFieldName(t *testing.T) {
	testIP := "1.1.1.4/32"
	testSlug := "ipadr_dev_fn"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDeviceDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  device_interface_id = netbox_device_interface.test.id
  status = "active"
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "device_interface_id", "netbox_device_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"device_interface_id"},
			},
		},
	})
}

func TestAccNetboxIPAddress_vmByFieldName(t *testing.T) {
	testIP := "1.1.1.5/32"
	testSlug := "ipadr_vm_fn"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "virtual_machine_interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"virtual_machine_interface_id"},
			},
		},
	})
}

// TestAccNetboxIPAddress_standalone tests the case where an ip address is not linked to a vm or device
func TestAccNetboxIPAddress_standalone(t *testing.T) {
	testIP := "1.1.1.6/32"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "active"
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
				),
			},
			{
				ResourceName:      "netbox_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIPAddress_nat(t *testing.T) {
	testIP := "1.1.1.8/32"
	testIPInside := "1.1.1.9/32"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "active"
}

resource "netbox_ip_address" "inside" {
  ip_address = "%s"
  status = "active"
  nat_inside_address_id = netbox_ip_address.test.id
}
`, testIP, testIPInside),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.inside", "nat_inside_address_id", "netbox_ip_address.test", "id"),
				),
			},
			// we have to make another step because netbox_ip_address.test.nat_outside_addresses needs a refresh
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "active"
}

resource "netbox_ip_address" "inside" {
  ip_address = "%s"
  status = "active"
  nat_inside_address_id = netbox_ip_address.test.id
}
`, testIP, testIPInside),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "nat_outside_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "nat_outside_addresses.0.id", "netbox_ip_address.inside", "id"),
				),
			},
			{
				ResourceName:      "netbox_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIPAddress_invalidConfig(t *testing.T) {
	testIP := "1.1.1.7/32"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  object_type = "dcim.interface"
  status = "active"
}`, testIP),
				ExpectError: regexp.MustCompile(".*all of `interface_id,object_type` must be specified.*"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  interface_id = 1
  status = "active"
}`, testIP),
				ExpectError: regexp.MustCompile(".*all of `interface_id,object_type` must be specified.*"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  virtual_machine_interface_id = 1
  interface_id = 1
  object_type = "dcim.interface"
  status = "active"
}`, testIP),
				ExpectError: regexp.MustCompile(".*conflicts with interface_id.*"),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ip_address", &resource.Sweeper{
		Name:         "netbox_ip_address",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamIPAddressesListParams()
			res, err := api.Ipam.IpamIPAddressesList(params, nil)
			if err != nil {
				return err
			}
			for _, ipAddress := range res.GetPayload().Results {
				if len(ipAddress.Tags) > 0 && (ipAddress.Tags[0] == &models.NestedTag{Name: strToPtr("acctest"), Slug: strToPtr("acctest")}) {
					deleteParams := ipam.NewIpamIPAddressesDeleteParams().WithID(ipAddress.ID)
					_, err := api.Ipam.IpamIPAddressesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an ip address")
				}
			}
			return nil
		},
	})
}
