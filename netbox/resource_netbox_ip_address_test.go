package netbox

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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

func TestAccNetboxIPAddress_cf(t *testing.T) {
	testIP := "1.1.1.8/32"
	testSlug := "ipaddr_cf"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name   = "%s"
  type   = "text"
  weight = 100
  content_types = ["ipam.ipaddress"]
}
resource "netbox_ip_address" "test_customfield" {
  ip_address = "%s"
  status = "active"
  custom_fields = {
    "${netbox_custom_field.test.name}" = "test-field"
  }
}`, testSlug, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test_customfield", fmt.Sprintf("custom_fields.%s", testSlug), "test-field"),
				),
			},
			{
				ResourceName:      "netbox_ip_address.test_customfield",
				ImportState:       true,
				ImportStateVerify: true,
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
				ResourceName:      "netbox_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
				// import step doesn't have context about device_interface_id, thus we
				// get interface_id and object_type imported instead
				ImportStateVerifyIgnore: []string{"interface_id", "object_type", "device_interface_id"},
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
				ResourceName:      "netbox_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
				// import step doesn't have context about virtual_machine_interface_id, thus we
				// get interface_id and object_type imported instead
				ImportStateVerifyIgnore: []string{"interface_id", "object_type", "virtual_machine_interface_id"},
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
	testIP := "1.1.1.10/32"
	testIPInside := "1.1.1.11/32"
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

// TestAccNetboxIPAddress_allocateFromPrefixWithInterface covers allocating an
// address and assigning it to an interface in the same apply - the combined
// workflow netbox_available_ip_address's own deviceByObjectType /
// deviceByFieldName tests cover, now on the unified resource.
func TestAccNetboxIPAddress_allocateFromPrefixWithInterface(t *testing.T) {
	testSlug := "ipadr_alloc_dev"
	testName := testAccGetTestName(testSlug)

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDeviceDependencies(testName) + `
resource "netbox_prefix" "test" {
  prefix  = "1.1.25.0/24"
  status  = "active"
  is_pool = false
}
resource "netbox_ip_address" "test" {
  prefix_id    = netbox_prefix.test.id
  object_type  = "dcim.interface"
  interface_id = netbox_device_interface.test.id
  status       = "active"
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", "1.1.25.1/24"),
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "interface_id", "netbox_device_interface.test", "id"),
				),
			},
		},
	})
}

// TestAccNetboxIPAddress_allocationSourceExactlyOneOf covers the mutual
// exclusivity between ip_address, prefix_id and ip_range_id: exactly one must
// be given.
func TestAccNetboxIPAddress_allocationSourceExactlyOneOf(t *testing.T) {
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: `
resource "netbox_ip_address" "test" {
  status = "active"
}`,
				ExpectError: regexp.MustCompile("one of `ip_address,ip_range_id,prefix_id` must be specified"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix  = "1.1.26.0/24"
  status  = "active"
  is_pool = false
}
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  prefix_id  = netbox_prefix.test.id
  status     = "active"
}`, "1.1.26.5/24"),
				ExpectError: regexp.MustCompile("only one of `ip_address,ip_range_id,prefix_id` can be specified"),
			},
		},
	})
}

// TestAccNetboxIPAddress_allocateFromPrefix covers creating a netbox_ip_address
// with no ip_address given, letting Netbox pick the next free address from a
// prefix - the same allocation netbox_available_ip_address performs, folded
// into this resource instead of a separate one.
func TestAccNetboxIPAddress_allocateFromPrefix(t *testing.T) {
	testPrefix := "1.1.20.0/24"

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
  prefix  = "%s"
  status  = "active"
  is_pool = false
}
resource "netbox_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
  status    = "active"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", "1.1.20.1/24"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
				),
			},
		},
	})
}

// TestAccNetboxIPAddress_allocateFromIPRange is the ip_range_id twin of
// TestAccNetboxIPAddress_allocateFromPrefix.
func TestAccNetboxIPAddress_allocateFromIPRange(t *testing.T) {
	startAddress := "1.1.21.1/24"
	endAddress := "1.1.21.50/24"

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_range" "test" {
  start_address = "%s"
  end_address   = "%s"
}
resource "netbox_ip_address" "test" {
  ip_range_id = netbox_ip_range.test.id
  status      = "active"
}`, startAddress, endAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", "1.1.21.1/24"),
				),
			},
		},
	})
}

// TestAccNetboxIPAddress_adoptExistingAddress imports an address allocated out
// of band - typically via the available-ips endpoint, exactly like
// netbox_available_ip_address's own Create does - and asserts it survives:
// same Netbox record, prefix recorded in state, empty plan afterwards.
//
// Netbox does not report which prefix an address was allocated from, so Read
// cannot populate prefix_id: after import, state holds 0 while the
// configuration carries a real prefix id. Without the CustomizeDiff guard, that
// 0 -> id difference plans a destroy+create of a live address, and the
// re-allocation hands out a different one - silently renumbering the host. This
// is issue #861 against netbox_available_ip_address; folding the allocation
// into netbox_ip_address must not reintroduce it.
func TestAccNetboxIPAddress_adoptExistingAddress(t *testing.T) {
	prefixes := `
resource "netbox_prefix" "test" {
  prefix  = "1.1.22.0/24"
  status  = "active"
  is_pool = false
}`

	withAdoptedAddress := prefixes + `
resource "netbox_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
  status    = "active"
}`

	var prefixID, adoptedAddressID int64

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: prefixes,
				Check: func(s *terraform.State) (err error) {
					prefixID, err = testAccStateID(s, "netbox_prefix.test")
					return err
				},
			},
			{
				// Allocate behind Terraform's back, the same way the resource
				// does, then adopt it.
				PreConfig: func() {
					api := testAccProvider.Meta().(*providerState)
					params := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(prefixID).
						WithData([]*models.AvailableIP{{}})
					res, err := api.Ipam.IpamPrefixesAvailableIpsCreate(params, nil)
					if err != nil {
						t.Fatalf("allocating an out-of-band IP in prefix %d: %s", prefixID, err)
					}
					if len(res.Payload) == 0 {
						t.Fatalf("no available IP addresses in prefix %d", prefixID)
					}
					adoptedAddressID = res.Payload[0].ID
				},
				Config:             withAdoptedAddress,
				ResourceName:       "netbox_ip_address.test",
				ImportState:        true,
				ImportStatePersist: true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strconv.FormatInt(adoptedAddressID, 10), nil
				},
			},
			// Records the prefix in state, keeps the address.
			{
				Config: withAdoptedAddress,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_ip_address.test", "prefix_id", "netbox_prefix.test", "id"),
					func(s *terraform.State) error {
						id, err := testAccStateID(s, "netbox_ip_address.test")
						if err != nil {
							return err
						}
						if id != adoptedAddressID {
							return fmt.Errorf("expected the adopted address (Netbox id %d) to be kept, got id %d", adoptedAddressID, id)
						}
						return nil
					},
				),
			},
			// Converged.
			{
				Config:   withAdoptedAddress,
				PlanOnly: true,
			},
		},
	})
}

// TestAccNetboxIPAddress_allocationSourceChangeReallocates guards the other
// side: a real prefix change on an allocated address must still re-allocate.
func TestAccNetboxIPAddress_allocationSourceChangeReallocates(t *testing.T) {
	prefixes := `
resource "netbox_prefix" "first" {
  prefix  = "1.1.23.0/24"
  status  = "active"
  is_pool = false
}
resource "netbox_prefix" "second" {
  prefix  = "1.1.24.0/24"
  status  = "active"
  is_pool = false
}`

	fromFirst := prefixes + `
resource "netbox_ip_address" "test" {
  prefix_id = netbox_prefix.first.id
  status    = "active"
}`

	fromSecond := prefixes + `
resource "netbox_ip_address" "test" {
  prefix_id = netbox_prefix.second.id
  status    = "active"
}`

	var firstID int64

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fromFirst,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", "1.1.23.1/24"),
					func(s *terraform.State) (err error) {
						firstID, err = testAccStateID(s, "netbox_ip_address.test")
						return err
					},
				),
			},
			{
				Config: fromSecond,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", "1.1.24.1/24"),
					func(s *terraform.State) error {
						id, err := testAccStateID(s, "netbox_ip_address.test")
						if err != nil {
							return err
						}
						if id == firstID {
							return fmt.Errorf("expected the address to be replaced, but it is still Netbox id %d", firstID)
						}
						return nil
					},
				),
			},
		},
	})
}

// testAccStateID returns a resource's Netbox id from state.
func testAccStateID(s *terraform.State, address string) (int64, error) {
	rs, ok := s.RootModule().Resources[address]
	if !ok {
		return 0, fmt.Errorf("%s not found in state", address)
	}
	return strconv.ParseInt(rs.Primary.ID, 10, 64)
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
			api := m.(*providerState)
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

func TestAccNetboxIPAddress_dnsNameCase(t *testing.T) {
	testIP := "1.1.1.253/32"
	testSlug := "ipaddr_dns_case"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				// Mixed-case dns_name: NetBox stores it as lowercase.
				// Without DiffSuppressFunc the post-apply plan check detects
				// drift and fails with "the plan was not empty" (#828).
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status     = "active"
  dns_name   = "Mytest.Example.Com"
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "dns_name", "mytest.example.com"),
				),
			},
		},
	})
}
