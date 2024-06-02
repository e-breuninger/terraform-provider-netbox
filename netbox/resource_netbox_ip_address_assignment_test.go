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

func testAccNetboxIPAddressAssignmentFullDependencies(testName string, testIP string, testIP2 string) string {
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

resource "netbox_ip_address" "outer" {
  ip_address = "%[3]s"
  status = "active"
  tags = [netbox_tag.test.name]
}

resource "netbox_ip_address" "test" {
  ip_address = "%[2]s"
  status = "active"
  tags = [netbox_tag.test.name]
  dns_name = "abc.example.com"
  description = "abc"
  role = "anycast"
  nat_inside_address_id = netbox_ip_address.outer.id
}
`, testName, testIP, testIP2)
}

func testAccNetboxIPAddressAssignmentFullDeviceDependencies(testName string, testIP string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

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
resource "netbox_ip_address" "test" {
  ip_address = "%[2]s"
  status = "active"
  tags = [netbox_tag.test.name]
}
`, testName, testIP)
}

func TestAccNetboxIPAddressAssignment_basic(t *testing.T) {
	testIP := "1.2.1.1/32"
	testIP2 := "1.2.2.1/32"
	testSlug := "ipaddress_assign"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressAssignmentFullDependencies(testName, testIP, testIP2) + fmt.Sprintf(`
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = netbox_ip_address.test.id
  object_type = "virtualization.vminterface"
  interface_id = netbox_interface.test.id
}

data "netbox_ip_addresses" "test" {
	depends_on = [netbox_ip_address_assignment.test]
	filter {
		name = "ip_address"
		value = "%[1]s"
	}
}
`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "ip_address_id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttr("netbox_ip_address_assignment.test", "object_type", "virtualization.vminterface"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.dns_name", "abc.example.com"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.status", "active"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.description", "abc"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.role", "anycast"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.tags.0.name", testName),
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address_assignment.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type"},
			},
		},
	})
}

func TestAccNetboxIPAddressAssignment_deviceByObjectType(t *testing.T) {
	testIP := "1.2.1.2/32"
	testSlug := "ipadr_dev_ot_assign"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressAssignmentFullDeviceDependencies(testName, testIP) + `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = netbox_ip_address.test.id
  object_type = "dcim.interface"
  interface_id = netbox_device_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "ip_address_id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttr("netbox_ip_address_assignment.test", "object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "interface_id", "netbox_device_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address_assignment.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type"},
			},
		},
	})
}

func TestAccNetboxIPAddressAssignment_vmSwitchStyle(t *testing.T) {
	testIP := "1.2.1.9/32"
	testIP2 := "1.2.2.9/32"
	testSlug := "ipadr_vm_sw_assign"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressAssignmentFullDependencies(testName, testIP, testIP2) + `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = netbox_ip_address.test.id
  object_type = "virtualization.vminterface"
  interface_id = netbox_interface.test.id
}`,
			},
			{
				Config: testAccNetboxIPAddressAssignmentFullDependencies(testName, testIP, testIP2) + `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = netbox_ip_address.test.id
  virtual_machine_interface_id = netbox_interface.test.id
}`,
			},
			{
				ResourceName:            "netbox_ip_address_assignment.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"interface_id", "object_type", "virtual_machine_interface_id"},
			},
		},
	})
}

// TestAccNetboxIPAddressAssignment_deviceByFieldName tests if creating an ip address and linking it to a device via the `device_interface_id` field works
func TestAccNetboxIPAddressAssignment_deviceByFieldName(t *testing.T) {
	testIP := "1.2.1.4/32"
	testSlug := "ipadr_dev_fn_assign"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressAssignmentFullDeviceDependencies(testName, testIP) + `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = netbox_ip_address.test.id
  device_interface_id = netbox_device_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "ip_address_id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "device_interface_id", "netbox_device_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address_assignment.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"device_interface_id"},
			},
		},
	})
}

func TestAccNetboxIPAddressAssignment_vmByFieldName(t *testing.T) {
	testIP := "1.2.1.5/32"
	testIP2 := "1.2.2.5/32"
	testSlug := "ipadr_vm_fn_assign"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressAssignmentFullDependencies(testName, testIP, testIP2) + `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = netbox_ip_address.test.id
  virtual_machine_interface_id = netbox_interface.test.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "ip_address_id", "netbox_ip_address.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_address_assignment.test", "virtual_machine_interface_id", "netbox_interface.test", "id"),
				),
			},
			{
				ResourceName:            "netbox_ip_address_assignment.test",
				ImportState:             true,
				ImportStateVerify:       true,
				ImportStateVerifyIgnore: []string{"virtual_machine_interface_id"},
			},
		},
	})
}

func TestAccNetboxIPAddressAssignment_invalidConfig(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{ // api.Ipam.IpamIPAddressesPartialUpdate()
			// NewPatchedWritableIPAddressRequest()

			{
				Config: `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = 1
  object_type = "dcim.interface"
}`,
				ExpectError: regexp.MustCompile(".*all of `interface_id,object_type` must be specified.*"),
			},
			{
				Config: `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = 1
  interface_id = 1
}`,
				ExpectError: regexp.MustCompile(".*all of `interface_id,object_type` must be specified.*"),
			},
			{
				Config: `
resource "netbox_ip_address_assignment" "test" {
  ip_address_id = 1
  virtual_machine_interface_id = 1
  interface_id = 1
  object_type = "dcim.interface"
}`,
				ExpectError: regexp.MustCompile(".*conflicts with interface_id.*"),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ip_address_assignment", &resource.Sweeper{
		Name:         "netbox_ip_address_assignment",
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
