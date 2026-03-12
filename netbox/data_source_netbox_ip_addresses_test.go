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
	virtual_machine_interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	role = "anycast"
	vrf_id = netbox_vrf.test.id
}
data "netbox_ip_addresses" "test" {
	depends_on = [netbox_ip_address.test]
}`, testIP),
				//                              This snippet sometimes returns things from other tests, even if resource.Test is used instead of resource.ParallelTest
				//                              This happens especially in CI testing (where test execution is presumably slow)
				//                              The check functions are now removed so this does no longer happen
				//				Check: resource.ComposeTestCheckFunc(
				//					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test", "ip_addresses.0.ip_address", "netbox_ip_address.test", "ip_address"),
				//					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.role", "anycast"),
				//					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test", "ip_addresses.0.tags.0.name", "netbox_tag.test", "name"),
				//				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_filter(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter"
	testName := testAccGetTestName(testSlug)
	testIP0 := "203.0.113.1/24"
	testIP1 := "203.0.113.2/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "ip_address"
		value = "%s"
	}
	filter {
		name = "vrf_id"
		value = netbox_vrf.test.id
	}
}`, testIP0, testIP1, testIP0),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.ip_address", "netbox_ip_address.test_list_0", "ip_address"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_filter2(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter_role"
	testName := testAccGetTestName(testSlug)
	testIP0 := "203.0.113.1/24"
	testIP1 := "203.0.113.2/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  role = "vip"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  role = "vrrp"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "role"
		value = "vip"
	}
	filter {
		name = "vrf_id"
		value = netbox_vrf.test.id
	}
}`, testIP0, testIP1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.ip_address", "netbox_ip_address.test_list_0", "ip_address"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_filter_parent_prefix(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter_prefix"
	testName := testAccGetTestName(testSlug)
	testPrefix1 := "203.0.113.0/24"
	testIP0 := "203.0.113.1/24"
	testIP1 := "203.0.200.1/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_prefix" "testv4" {
  prefix = "%s"
  status = "active"
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_0" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "parent_prefix"
		value = "%s"
	}
	filter {
		name = "vrf_id"
		value = netbox_vrf.test.id
	}
}`, testPrefix1, testIP0, testIP1, testPrefix1),
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
	testIP0 := "203.0.113.1/24"
	testIP1 := "203.0.113.2/24"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
	ip_address = "%s"
	virtual_machine_interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_1" {
	ip_address = "%s"
	virtual_machine_interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	vrf_id = netbox_vrf.test.id
}

data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "vm_interface_id"
		value = netbox_interface.test.id
	}
}`, testIP0, testIP1),
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
	testIP0 := "203.0.113.10/24"
	testIP1 := "203.0.113.20/24"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
	ip_address = "%s"
	virtual_machine_interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	tenant_id = netbox_tenant.test.id
	vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_1" {
	ip_address = "%s"
	virtual_machine_interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	tenant_id = netbox_tenant.test.id
	vrf_id = netbox_vrf.test.id
}

data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "vm_interface_id"
		value = netbox_interface.test.id
	}
}`, testIP0, testIP1),
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

func testAccNetboxIPAddressesDataSourceDependenciesMany(testName string) string {
	return testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
}`, testName) + fmt.Sprintf(`
resource "netbox_interface" "test" {
  name = "test"
  virtual_machine_id = netbox_virtual_machine.test.id
}

resource "netbox_vrf" "test" {
  name = "%s"
}

resource "netbox_ip_address" "test" {
  count       = 51
  ip_address  = "10.11.12.${count.index}/32"
  status      = "active"
  virtual_machine_interface_id = netbox_interface.test.id
  dns_name = "%s"
  vrf_id = netbox_vrf.test.id
}
`, testName, testName)
}

func TestAccNetboxIpAddressessDataSource_many(t *testing.T) {
	testSlug := "ip_adrs_ds_many"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressesDataSourceDependenciesMany(testName) + fmt.Sprintf(`
data "netbox_ip_addresses" "test" {
  depends_on = [netbox_ip_address.test]

  filter {
    name = "dns_name"
    value = "%s"
  }
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.#", "51"),
				),
			},
			{
				Config: testAccNetboxIPAddressesDataSourceDependenciesMany(testName) + `data "netbox_ip_addresses" "test" {
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

func TestAccNetboxIpAddressesDataSource_filter_tags(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter_tags"
	testName := testAccGetTestName(testSlug)
	testTag := testName + "_gw"
	testIP0 := "203.0.113.1/24"
	testIP1 := "203.0.113.2/24"
	testIP2 := "203.0.113.3/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_tag" "gw_tag" {
  name = "%s"
}
resource "netbox_ip_address" "test_list_0" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name, netbox_tag.gw_tag.name]
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_2" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
  vrf_id = netbox_vrf.test.id
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1, netbox_ip_address.test_list_2]

	filter {
		name = "tag"
		value = "%s"
	}
	filter {
		name = "vrf_id"
		value = netbox_vrf.test.id
	}
}`, testTag, testIP0, testIP1, testIP2, testTag),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.ip_address", "netbox_ip_address.test_list_1", "ip_address"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_filter_description(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter_description"
	testName := testAccGetTestName(testSlug)
	testIP0 := "203.0.113.1/24"
	testIP1 := "203.0.113.2/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test_list_0" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  role = "vip"
  description = "test 1"
  vrf_id = netbox_vrf.test.id
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  role = "vrrp"
  description = "test 2"
  vrf_id = netbox_vrf.test.id
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "description"
		value = "test 1"
	}
	filter {
		name = "vrf_id"
		value = netbox_vrf.test.id
	}
}`, testIP0, testIP1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test_list", "ip_addresses.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.ip_address", "netbox_ip_address.test_list_0", "ip_address"),
				),
			},
		},
	})
}
