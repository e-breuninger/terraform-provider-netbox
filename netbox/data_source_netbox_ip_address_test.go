package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxIPAddressDataSourceFullDeviceDependencies(testName string) string {
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

func TestAccNetboxIpAddressDataSource_basic(t *testing.T) {
	ipAddress := "10.0.0.107/24"
	status := "active"
	testSlug := "IPAddressDataSourceBasic"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressDataSourceFullDeviceDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%[1]s"
  status = "%[2]s"
}
data "netbox_ip_address" "test" {
  depends_on = [netbox_ip_address.test]
  id = netbox_ip_address.test.id
}`, ipAddress, status, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.test", "ip_address", "netbox_ip_address.test", "ip_address"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccNetboxIpAddressDataSource_nestedDevice(t *testing.T) {
	ipAddress := "10.0.0.123/24"
	status := "active"
	testSlug := "IPAddressDataSourceNestedDevice"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressDataSourceFullDeviceDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%[1]s"
  status = "%[2]s"
  device_interface_id = netbox_device_interface.test.id
}

data "netbox_ip_address" "test" {
  depends_on = [netbox_ip_address.test]
  id = netbox_ip_address.test.id
}`, ipAddress, status, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.test", "assigned_object.0.id", "netbox_device_interface.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.test", "assigned_object.0.device.0.id", "netbox_device.test", "id"),

				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}

func TestAccNetboxIpAddressDataSource_nestedVM(t *testing.T) {
	ipAddress := "10.0.0.254/24"
	status := "active"
	testSlug := "IPAddressDataSourceNestedVM"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressDataSourceFullDeviceDependencies(testName) + fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[3]s"
}

resource "netbox_cluster_type" "test" {
  name = "%[3]s"
}

resource "netbox_cluster_group" "test" {
  name = "%[3]s"
}

resource "netbox_cluster" "test" {
  name = "%[3]s"
  cluster_type_id = netbox_cluster_type.test.id
  cluster_group_id = netbox_cluster_group.test.id
  site_id = netbox_site.test.id
  comments = "testcomments"
  description = "testdescription"
  tags = [netbox_tag.test.name]
}

resource "netbox_virtual_machine" "test" {
  name = "%[3]s_0"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_mb = 256
  vcpus = 4
}

resource "netbox_interface" "test" {
  name               = "eth0"
  virtual_machine_id = netbox_virtual_machine.test.id
}

resource "netbox_ip_address" "test" {
  ip_address = "%[1]s"
  status = "%[2]s"
  interface_id = netbox_interface.test.id
  object_type  = "virtualization.vminterface"
}

data "netbox_ip_address" "test" {
  depends_on = [netbox_ip_address.test]
  id = netbox_ip_address.test.id
}`, ipAddress, status, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.test", "ip_address", "netbox_ip_address.test", "ip_address"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.test", "assigned_object.0.name", "netbox_interface.test", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_address.test", "assigned_object.0.device.0.id", "netbox_virtual_machine.test", "id"),
				),
				ExpectNonEmptyPlan: false,
			},
		},
	})
}
