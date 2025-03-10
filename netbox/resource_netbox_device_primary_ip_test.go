package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func testAccNetboxDevicePrimaryIPFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id = netbox_site.test.id
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
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

resource "netbox_location" "test" {
  name = "%[1]s"
  site_id =netbox_site.test.id
}

resource "netbox_rack" "test" {
  name = "%[1]s"
  site_id = netbox_site.test.id
  status = "reserved"
  width = 19
  u_height = 48
  tenant_id = netbox_tenant.test.id
  location_id = netbox_location.test.id
}

resource "netbox_device" "test" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  tenant_id = netbox_tenant.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "thisisacomment"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 10

  tags = [netbox_tag.test.name]
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_device_interface" "test" {
  device_id = netbox_device.test.id
  name = "%[1]s"
  type = "1000base-t"
}
`, testName)
}

func TestAccNetboxDevicePrimaryIP4_basic(t *testing.T) {
	testSlug := "pr_ip_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + `
resource "netbox_ip_address" "test_v4" {
  ip_address = "1.1.1.12/32"
  status = "active"
  interface_id = netbox_device_interface.test.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4" {
  device_id = netbox_device.test.id
  ip_address_id = netbox_ip_address.test_v4.id
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4", "ip_address_id", "netbox_ip_address.test_v4", "id"),

					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "planned"),
				),
			},
		},
	})
}

func TestAccNetboxDevicePrimaryIP6_basic(t *testing.T) {
	testSlug := "pr_ipv6_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + `
resource "netbox_ip_address" "test_v6" {
  ip_address = "2001::1/128"
  status = "active"
  interface_id = netbox_device_interface.test.id
  object_type = "dcim.interface"
}
resource "netbox_device_primary_ip" "test_v6" {
  device_id = netbox_device.test.id
  ip_address_id = netbox_ip_address.test_v6.id
  ip_address_version = 6
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v6", "device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v6", "ip_address_id", "netbox_ip_address.test_v6", "id"),

					resource.TestCheckResourceAttr("netbox_device.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_device.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_device.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_device.test", "status", "planned"),
				),
			},
		},
	})
}

func TestAccNetboxDevicePrimaryIP4_removePrimary(t *testing.T) {
	testSlug := "pr_ip_removePrimary"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test2" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "thisisacomment"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test2" {
  device_id = netbox_device.test2.id
  name = "%[1]s"
  type = "1000base-t"
}

resource "netbox_ip_address" "test_v4_2" {
  ip_address = "1.1.1.16/32"
  status = "active"
  interface_id = netbox_device_interface.test2.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_2" {
  device_id = netbox_device.test2.id
  ip_address_id = netbox_ip_address.test_v4_2.id
}`, testName),
			},
			// A repeated second step is required, so that the resource "netbox_device" "test2" goes through a resourceNetboxDeviceRead cycle
			// This is needed because adding a netbox_device_primary_ip updates the netbox_device
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test2" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "thisisacomment"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test2" {
  device_id = netbox_device.test2.id
  name = "%[1]s"
  type = "1000base-t"
}

resource "netbox_ip_address" "test_v4_2" {
  ip_address = "1.1.1.16/32"
  status = "active"
  interface_id = netbox_device_interface.test2.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_2" {
  device_id = netbox_device.test2.id
  ip_address_id = netbox_ip_address.test_v4_2.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_2", "device_id", "netbox_device.test2", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_2", "ip_address_id", "netbox_ip_address.test_v4_2", "id"),
				),
			},
			// Now we do 2 things: modify netbox_device.test2 (changing the comment value), AND we remove the IP and primary IP
			// This fails with:
			//        Error: [PUT /dcim/devices/{id}/][400] dcim_devices_update default {"primary_ip4":["Related object not found using the provided numeric ID: 14"]}
			// because (I think) that the device is doing 1) a read of the current state, 2) the deletion of the primary IP then modifies the device, 3) the device then tries to write its changes, but its now out of date
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test2" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "thisisacomment with changes"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test2" {
  device_id = netbox_device.test2.id
  name = "%[1]s"
  type = "1000base-t"
}
`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_device.test2", "name", testName),
					resource.TestCheckResourceAttr("netbox_device.test2", "primary_ipv4", "0"),
					resource.TestCheckResourceAttr("netbox_device.test2", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_device.test2", "tags.0", testName),
					resource.TestCheckResourceAttr("netbox_device.test2", "status", "planned"),
				),
			},
		},
	})
}

func TestAccNetboxDevicePrimaryIP4_simpleUpdateDevice(t *testing.T) {
	testSlug := "pr_ip_updateDevice"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test3" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test3" {
  device_id = netbox_device.test3.id
  name = "%[1]s"
  type = "1000base-t"
}

resource "netbox_ip_address" "test_v4_3" {
  ip_address = "1.1.1.101/32"
  status = "active"
  interface_id = netbox_device_interface.test3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_3" {
  device_id = netbox_device.test3.id
  ip_address_id = netbox_ip_address.test_v4_3.id
}`, testName),
			},
			// Update the device and check the primary ip value
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test3" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment2"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test3" {
  device_id = netbox_device.test3.id
  name = "%[1]s"
  type = "virtual"
}

resource "netbox_ip_address" "test_v4_3" {
  ip_address = "1.1.1.101/32"
  status = "active"
  interface_id = netbox_device_interface.test3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_3" {
  device_id = netbox_device.test3.id
  ip_address_id = netbox_ip_address.test_v4_3.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "device_id", "netbox_device.test3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "ip_address_id", "netbox_ip_address.test_v4_3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test3", "primary_ipv4", "netbox_ip_address.test_v4_3", "id"),
				),
			},
			// Update the device and the interface and check the primary ip value
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test3" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment3"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test3" {
  device_id = netbox_device.test3.id
  name = "%[1]s"
  type = "1000base-t"
}

resource "netbox_ip_address" "test_v4_3" {
  ip_address = "1.1.1.101/32"
  status = "active"
  interface_id = netbox_device_interface.test3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_3" {
  device_id = netbox_device.test3.id
  ip_address_id = netbox_ip_address.test_v4_3.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "device_id", "netbox_device.test3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "ip_address_id", "netbox_ip_address.test_v4_3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test3", "primary_ipv4", "netbox_ip_address.test_v4_3", "id"),
				),
			},
			// Update the device, the interface and the IP address, then check the primary ip
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test3" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment4"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test3" {
  device_id = netbox_device.test3.id
  name = "%[1]s"
  type = "virtual"
}

resource "netbox_ip_address" "test_v4_3" {
  ip_address = "1.1.1.102/32"
  status = "active"
  interface_id = netbox_device_interface.test3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_3" {
  device_id = netbox_device.test3.id
  ip_address_id = netbox_ip_address.test_v4_3.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "device_id", "netbox_device.test3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "ip_address_id", "netbox_ip_address.test_v4_3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test3", "primary_ipv4", "netbox_ip_address.test_v4_3", "id"),
				),
			},
			// Force a read at the end
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test3" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "commnent4"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test3" {
  device_id = netbox_device.test3.id
  name = "%[1]s"
  type = "virtual"
}

resource "netbox_ip_address" "test_v4_3" {
  ip_address = "1.1.1.102/32"
  status = "active"
  interface_id = netbox_device_interface.test3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_3" {
  device_id = netbox_device.test3.id
  ip_address_id = netbox_ip_address.test_v4_3.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "device_id", "netbox_device.test3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_3", "ip_address_id", "netbox_ip_address.test_v4_3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test3", "primary_ipv4", "netbox_ip_address.test_v4_3", "id"),
				),
			},
		},
	})
}

func TestAccNetboxDevicePrimaryIP4_refsUpdateDevice(t *testing.T) {
	testSlug := "pr_ip_refUpdateDevice"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:                 func() { testAccPreCheck(t) },
		ProtoV6ProviderFactories: TestAccProtoV6ProviderFactories,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test4" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test4" {
  device_id = netbox_device.test4.id
  name = "%[1]s"
  type = "1000base-t"

  lifecycle {
    create_before_destroy = true
  }
}

resource "netbox_ip_address" "test_v4_4" {
  ip_address = "1.1.2.101/32"
  status = "active"
  interface_id = netbox_device_interface.test4.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_4" {
  device_id = netbox_device.test4.id
  ip_address_id = netbox_ip_address.test_v4_4.id
}`, testName),
			},
			// Switch the interface for a new one
			// the create_before_destroy lifecycle ensures the interface exists at all times
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test4" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test4_2" {
  device_id = netbox_device.test4.id
  name = "%[1]s-2"
  type = "1000base-t"

  lifecycle {
    create_before_destroy = true
  }
}

resource "netbox_ip_address" "test_v4_4" {
  ip_address = "1.1.2.101/32"
  status = "active"
  interface_id = netbox_device_interface.test4_2.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_4" {
  device_id = netbox_device.test4.id
  ip_address_id = netbox_ip_address.test_v4_4.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "device_id", "netbox_device.test4", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "ip_address_id", "netbox_ip_address.test_v4_4", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test4", "primary_ipv4", "netbox_ip_address.test_v4_4", "id"),
				),
			},
			// switch the ip address for a new one
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test4" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test4_2" {
  device_id = netbox_device.test4.id
  name = "%[1]s-2"
  type = "1000base-t"

  lifecycle {
    create_before_destroy = true
  }
}

resource "netbox_ip_address" "test_v4_4_2" {
  ip_address = "1.1.2.102/32"
  status = "active"
  interface_id = netbox_device_interface.test4_2.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_4" {
  device_id = netbox_device.test4.id
  ip_address_id = netbox_ip_address.test_v4_4_2.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "device_id", "netbox_device.test4", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "ip_address_id", "netbox_ip_address.test_v4_4_2", "id"),
					// Needs another read on the device for the field to be accurate
					//resource.TestCheckResourceAttrPair("netbox_device.test4", "primary_ipv4", "netbox_ip_address.test_v4_4_2", "id"),
				),
			},
			// Force another read to update device
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test4" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test4_2" {
  device_id = netbox_device.test4.id
  name = "%[1]s-2"
  type = "1000base-t"

  lifecycle {
    create_before_destroy = true
  }
}

resource "netbox_ip_address" "test_v4_4_2" {
  ip_address = "1.1.2.102/32"
  status = "active"
  interface_id = netbox_device_interface.test4_2.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_4" {
  device_id = netbox_device.test4.id
  ip_address_id = netbox_ip_address.test_v4_4_2.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "device_id", "netbox_device.test4", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "ip_address_id", "netbox_ip_address.test_v4_4_2", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test4", "primary_ipv4", "netbox_ip_address.test_v4_4_2", "id"),
				),
			},
			// Switch both the interface and the ip address at the same time
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test4" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test4_3" {
  device_id = netbox_device.test4.id
  name = "%[1]s-3"
  type = "virtual"

  lifecycle {
    create_before_destroy = true
  }
}

resource "netbox_ip_address" "test_v4_4_3" {
  ip_address = "1.1.2.103/32"
  status = "active"
  interface_id = netbox_device_interface.test4_3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_4" {
  device_id = netbox_device.test4.id
  ip_address_id = netbox_ip_address.test_v4_4_3.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "device_id", "netbox_device.test4", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "ip_address_id", "netbox_ip_address.test_v4_4_3", "id"),
					// Require another read
					//resource.TestCheckResourceAttrPair("netbox_device.test4", "primary_ipv4", "netbox_ip_address.test_v4_4_3", "id"),
				),
			},
			// Force another read
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test4" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test4_3" {
  device_id = netbox_device.test4.id
  name = "%[1]s-3"
  type = "virtual"

  lifecycle {
    create_before_destroy = true
  }
}

resource "netbox_ip_address" "test_v4_4_3" {
  ip_address = "1.1.2.103/32"
  status = "active"
  interface_id = netbox_device_interface.test4_3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_4" {
  device_id = netbox_device.test4.id
  ip_address_id = netbox_ip_address.test_v4_4_3.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "device_id", "netbox_device.test4", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4", "ip_address_id", "netbox_ip_address.test_v4_4_3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test4", "primary_ipv4", "netbox_ip_address.test_v4_4_3", "id"),
				),
			},
			// Make a new primary IP
			{
				Config: testAccNetboxDevicePrimaryIPFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test4" {
  name = "%[1]s"
  role_id = netbox_device_role.test.id
  site_id = netbox_site.test.id
  device_type_id = netbox_device_type.test.id
  cluster_id = netbox_cluster.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  comments = "comment1"
  status = "planned"
  rack_id = netbox_rack.test.id
  rack_face = "front"
  rack_position = 11

  tags = [netbox_tag.test.name]
}

resource "netbox_device_interface" "test4_3" {
  device_id = netbox_device.test4.id
  name = "%[1]s-3"
  type = "virtual"

  lifecycle {
    create_before_destroy = true
  }
}

resource "netbox_ip_address" "test_v4_4_3" {
  ip_address = "1.1.2.103/32"
  status = "active"
  interface_id = netbox_device_interface.test4_3.id
  object_type = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4_4_2" {
  device_id = netbox_device.test4.id
  ip_address_id = netbox_ip_address.test_v4_4_3.id
  ip_address_version = 4
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4_2", "device_id", "netbox_device.test4", "id"),
					resource.TestCheckResourceAttrPair("netbox_device_primary_ip.test_v4_4_2", "ip_address_id", "netbox_ip_address.test_v4_4_3", "id"),
					resource.TestCheckResourceAttrPair("netbox_device.test4", "primary_ipv4", "netbox_ip_address.test_v4_4_3", "id"),
				),
			},
		},
	})
}
