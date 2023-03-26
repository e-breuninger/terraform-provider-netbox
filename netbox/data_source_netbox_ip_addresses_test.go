package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpAddressesDataSource_basic(t *testing.T) {

	testSlug := "ipam_ipaddrs_ds_basic"
	testName := testAccGetTestName(testSlug)
	testIP := "203.0.113.1/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
	ip_address = "%s"
	interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	role = "anycast"
}
data "netbox_ip_addresses" "test" {
	depends_on = [netbox_ip_address.test]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test", "ip_addresses.0.ip_address", "netbox_ip_address.test", "ip_address"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.role", "anycast"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_filter(t *testing.T) {

	testSlug := "ipam_ipaddrs_ds_filter"
	testName := testAccGetTestName(testSlug)
	testIP_0 := "203.0.113.1/24"
	testIP_1 := "203.0.113.2/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "ip_address"
		value = "%s"
	}
}`, testIP_0, testIP_1, testIP_0),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.ip_address", "netbox_ip_address.test_list_0", "ip_address"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_multiple(t *testing.T) {

	testSlug := "ipam_ipaddrs_ds_multiple"
	testIP_0 := "203.0.113.1/24"
	testIP_1 := "203.0.113.2/24"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
	ip_address = "%s"
	interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
}
resource "netbox_ip_address" "test_list_1" {
	ip_address = "%s"
	interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
}

data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "vm_interface_id"
		value = netbox_interface.test.id
	}
}`, testIP_0, testIP_1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.ip_address", "netbox_ip_address.test_list_0", "ip_address"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.1.ip_address", "netbox_ip_address.test_list_1", "ip_address"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_flattenTenant(t *testing.T) {

	testSlug := "ipam_ipaddrs_ds_flattenTenant"
	testIP_0 := "203.0.113.10/24"
	testIP_1 := "203.0.113.20/24"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
	ip_address = "%s"
	interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	tenant_id = netbox_tenant.test.id
}
resource "netbox_ip_address" "test_list_1" {
	ip_address = "%s"
	interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	tenant_id = netbox_tenant.test.id
}

data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "vm_interface_id"
		value = netbox_interface.test.id
	}
}`, testIP_0, testIP_1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.0.tenant.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.1.tenant.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.tenant.0.name", "netbox_tenant.test", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.1.tenant.0.name", "netbox_tenant.test", "name"),
				),
			},
		},
	})
}

func testAccNetboxIpAddressesDataSourceDependencies_many(testName string) string {
	return testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
}`, testName) + `
resource "netbox_interface" "test" {
  name = "test"
  virtual_machine_id = netbox_virtual_machine.test.id
}

resource "netbox_ip_address" "test" {
  count       = 51
  ip_address  = "10.11.12.${count.index}/32"
  status      = "active"
  interface_id = netbox_interface.test.id
}
`
}

func TestAccNetboxIpAddressessDataSource_many(t *testing.T) {

	testSlug := "ip_adrs_ds_many"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIpAddressesDataSourceDependencies_many(testName) + `data "netbox_ip_addresses" "test" {
  depends_on = [netbox_ip_address.test]
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.#", "51"),
				),
			},
			{
				Config: testAccNetboxIpAddressesDataSourceDependencies_many(testName) + `data "netbox_ip_addresses" "test" {
  depends_on = [netbox_ip_address.test]
  limit = 2
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.#", "2"),
				),
			},
		},
	})
}
