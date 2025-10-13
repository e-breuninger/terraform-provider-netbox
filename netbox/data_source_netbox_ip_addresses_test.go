package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpAddressesDataSource_basic(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_basic"
	testName := testAccGetTestName(testSlug)
	testIP := "140.18.8.1/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
	ip_address = "%[1]s"
	virtual_machine_interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
	role = "anycast"
}
data "netbox_ip_addresses" "test" {
	depends_on = [netbox_ip_address.test]

	filter {
		name = "ip_address"
		value = "%[1]s"
	}
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test", "ip_addresses.0.ip_address", "netbox_ip_address.test", "ip_address"),
					resource.TestCheckResourceAttr("data.netbox_ip_addresses.test", "ip_addresses.0.role", "anycast"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test", "ip_addresses.0.tags.0.name", "netbox_tag.test", "name"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_filter(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter"
	testName := testAccGetTestName(testSlug)
	testIP0 := "140.18.8.2/24"
	testIP1 := "140.18.8.3/24"
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
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "ip_address"
		value = "%s"
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
	testIP0 := "140.18.8.4/24"
	testIP1 := "140.18.8.5/24"
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
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  role = "vrrp"
  tags = [netbox_tag.test.name]
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "role"
		value = "vip"
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
	testPrefix1 := "140.18.8.0/24"
	testIP0 := "140.18.8.6/24"
	testIP1 := "192.168.21.7/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_prefix" "testv4" {
  prefix = "%s"
  status = "active"
}
resource "netbox_ip_address" "test_list_0" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1]

	filter {
		name = "parent_prefix"
		value = "%s"
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
	testIP0 := "140.18.8.8/24"
	testIP1 := "140.18.8.9/24"
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
}
resource "netbox_ip_address" "test_list_1" {
	ip_address = "%s"
	virtual_machine_interface_id = netbox_interface.test.id
	status = "active"
	tags = [netbox_tag.test.name]
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
	testIP0 := "140.18.8.10/24"
	testIP1 := "140.18.8.11/24"
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
}
resource "netbox_ip_address" "test_list_1" {
	ip_address = "%s"
	virtual_machine_interface_id = netbox_interface.test.id
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

resource "netbox_ip_address" "test" {
  count       = 51
  ip_address  = "10.11.12.${count.index}/32"
  status      = "active"
  virtual_machine_interface_id = netbox_interface.test.id
  dns_name = "%s"
}
`, testName)
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
	testTag := "default-gw"
	testName := testAccGetTestName(testSlug)
	testIP0 := "140.18.8.12/24"
	testIP1 := "140.18.8.13/24"
	testIP2 := "140.18.8.14/24"
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
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name, netbox_tag.gw_tag.name]
}
resource "netbox_ip_address" "test_list_2" {
  ip_address = "%s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.test.name]
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1, netbox_ip_address.test_list_2]

	filter {
		name = "tag"
		value = "%s"
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

func TestAccNetboxIpAddressesDataSource_nestedVM(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter_tags"
	testTag := "default-gw"
	testName := testAccGetTestName(testSlug)
	testIP0 := "140.18.8.15/24"
	testIP1 := "140.18.8.16/24"
	testIP2 := "140.18.8.17/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_tag" "gw_tag" {
  name = "%[1]s"
}

resource "netbox_virtual_machine" "test_1" {
  name = "%[5]s_1"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_interface" "test_1" {
  name = "%[5]s_1"
  virtual_machine_id = netbox_virtual_machine.test_1.id
}

resource "netbox_virtual_machine" "test_2" {
  name = "%[5]s_2"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_interface" "test_2" {
  name = "%[5]s_2"
  virtual_machine_id = netbox_virtual_machine.test_2.id
}

resource "netbox_ip_address" "test_list_0" {
  ip_address = "%[2]s"
  virtual_machine_interface_id = netbox_interface.test.id
  status = "active"
  tags = [netbox_tag.gw_tag.name]
}
resource "netbox_ip_address" "test_list_1" {
  ip_address = "%[3]s"
  virtual_machine_interface_id = netbox_interface.test_1.id
  status = "active"
  tags = [netbox_tag.gw_tag.name]
}
resource "netbox_ip_address" "test_list_2" {
  ip_address = "%[4]s"
  virtual_machine_interface_id = netbox_interface.test_2.id
  status = "active"
  tags = [netbox_tag.gw_tag.name]
}
data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test_list_0, netbox_ip_address.test_list_1, netbox_ip_address.test_list_2]

	filter {
		name = "tag"
		value = "%[1]s"
	}
}`, testTag, testIP0, testIP1, testIP2, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.assigned_object.0.name", "netbox_interface.test", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.assigned_object.0.device.0.id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.1.assigned_object.0.name", "netbox_interface.test_1", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.1.assigned_object.0.device.0.id", "netbox_virtual_machine.test_1", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.2.assigned_object.0.name", "netbox_interface.test_2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.2.assigned_object.0.device.0.id", "netbox_virtual_machine.test_2", "id"),
				),
			},
		},
	})
}

func TestAccNetboxIpAddressesDataSource_nestedDevice(t *testing.T) {
	testSlug := "ipam_ipaddrs_ds_filter_tags"
	testTag := "default-gw"
	testName := testAccGetTestName(testSlug)
	testIP0 := "140.18.8.18/24"
	testIP1 := "140.18.8.19/24"
	testIP2 := "140.18.8.20/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_tag" "gw_tag" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[5]s"
  status = "active"
}

resource "netbox_device_role" "test" {
  name = "%[5]s"
  color_hex = "123456"
}

resource "netbox_manufacturer" "test" {
  name = "%[5]s"
}

resource "netbox_device_type" "test" {
  model = "%[5]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device" "test" {
  name = "dev_%[5]s_0"
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
}

resource "netbox_device_interface" "test" {
  name = "int_%[5]s_0"
  device_id = netbox_device.test.id
  type = "1000base-t"
}

resource "netbox_ip_address" "test" {
  ip_address = "%[2]s"
  status = "active"
  device_interface_id = netbox_device_interface.test.id
  tags = [netbox_tag.gw_tag.name]
}

resource "netbox_device" "test_1" {
  name = "dev_%[5]s_1"
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
}

resource "netbox_device_interface" "test_1" {
  name = "int_%[5]s_1"
  device_id = netbox_device.test_1.id
  type = "1000base-t"
}

resource "netbox_ip_address" "test_1" {
  ip_address = "%[3]s"
  status = "active"
  device_interface_id = netbox_device_interface.test_1.id
  tags = [netbox_tag.gw_tag.name]
}

resource "netbox_device" "test_2" {
  name = "dev_%[5]s_2"
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  role_id = netbox_device_role.test.id
}

resource "netbox_device_interface" "test_2" {
  name = "int_%[5]s_2"
  device_id = netbox_device.test_2.id
  type = "1000base-t"
}

resource "netbox_ip_address" "test_2" {
  ip_address = "%[4]s"
  status = "active"
  device_interface_id = netbox_device_interface.test_2.id
  tags = [netbox_tag.gw_tag.name]
}

data "netbox_ip_addresses" "test_list" {
	depends_on = [netbox_ip_address.test, netbox_ip_address.test_1, netbox_ip_address.test_2]

	filter {
		name = "tag"
		value = "%[1]s"
	}
}`, testTag, testIP0, testIP1, testIP2, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.assigned_object.0.name", "netbox_device_interface.test", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.0.assigned_object.0.device.0.id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.1.assigned_object.0.name", "netbox_device_interface.test_1", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.1.assigned_object.0.device.0.id", "netbox_device.test_1", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.2.assigned_object.0.name", "netbox_device_interface.test_2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_addresses.test_list", "ip_addresses.2.assigned_object.0.device.0.id", "netbox_device.test_2", "id"),
				),
			},
		},
	})
}
