// Copyright (c) 2022 Cisco Systems, Inc. and its affiliates
// All rights reserved.

package netbox

import (
	"encoding/json"
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"
)

func TestAccNetboxDevicesDataSource_basic(t *testing.T) {
	testSlug := "devices_ds_basic"
	testName := testAccGetTestName(testSlug)
	testLocalContextData, _ := json.Marshal(map[string]string{"testkey0": "testvalue0"})
	dependencies := testAccNetboxDeviceDataSourceDependencies(testName)
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxDeviceDataSourceFilterName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.name", testName+"_0"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.comments", "this is also a comment"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.description", "this is also a description"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.serial", "ABCDEF0"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.status", "staged"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.local_context_data", string(testLocalContextData)),
				),
			},
			{
				Config: dependencies + testAccNetboxDeviceDataSourceFilterTenant,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "4"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.primary_ipv4", "10.0.0.60"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.name", "netbox_device.test0", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.1.name", "netbox_device.test1", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.2.name", "netbox_device.test2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.3.name", "netbox_device.test3", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxDeviceDataSourceFilterRole,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "4"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.1.role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.name", "netbox_device.test0", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.1.name", "netbox_device.test1", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.2.name", "netbox_device.test2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.3.name", "netbox_device.test3", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxDeviceDataSourceNameRegex(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.name", "netbox_device.test2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.1.name", "netbox_device.test3", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxDeviceDataSourceLimit,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.tenant_id", "netbox_tenant.test", "id"),
				),
			},
			{
				Config: dependencies + testAccNetBoxDeviceDataSourceFilterTagsAndStatus,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.tag_devices", "devices.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_devices.tag_devices", "devices.0.tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_devices.tag_devices", "devices.0.status", "staged"),
				),
			},
			{
				Config: dependencies + testAccNetBoxDeviceDataSourceMultipleTagsFilter,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.multiple_filter_devices", "devices.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_devices.multiple_filter_devices", "devices.0.tags.#", "2"),
				),
			},
		},
	})
}

func testAccNetboxDeviceDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id = netbox_site.test.id
}

resource "netbox_location" "test" {
  name = "%[1]s"
  site_id =netbox_site.test.id
}

resource "netbox_rack_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
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

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}

resource "netbox_tag" "test_b" {
  name = "%[1]sb"
}

resource "netbox_tag" "test_c" {
  name = "%[1]sc"
}

resource "netbox_manufacturer" "test" {
  name = "%[1]s"
}

resource "netbox_device_type" "test" {
  model = "%[1]s"
  manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_device_interface" "test" {
  name      = "eth0"
  device_id = netbox_device.test0.id
  type      = "1000base-t"
}

resource "netbox_ip_address" "test" {
  ip_address   = "10.0.0.60/24"
  status       = "active"
  interface_id = netbox_device_interface.test.id
  object_type  = "dcim.interface"
}

resource "netbox_device_primary_ip" "test_v4" {
  device_id     = netbox_device.test0.id
  ip_address_id = netbox_ip_address.test.id
}

resource "netbox_device" "test0" {
  name = "%[1]s_0"
  comments = "this is also a comment"
  description = "this is also a description"
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  site_id = netbox_site.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  serial = "ABCDEF0"
  status = "staged"
  tags = [netbox_tag.test_a.name]
  local_context_data = jsonencode({"testkey0"="testvalue0"})
}

resource "netbox_device" "test1" {
  name = "%[1]s_1"
  comments = "this is also first comment"
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  site_id = netbox_site.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  serial = "ABCDEF1"
  local_context_data = jsonencode({"testkey1"="testvalue1"})
}

resource "netbox_device" "test2" {
  name = "%[1]s_2_regex"
  comments = "this is also second comment"
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  site_id = netbox_site.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  serial = "ABCDEF2"
  tags = [netbox_tag.test_b.name, netbox_tag.test_c.name]
  local_context_data = jsonencode({"testkey2"="testvalue2"})
}

resource "netbox_device" "test3" {
  name = "%[1]s_3_regex"
  comments = "this is also third comment"
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  site_id = netbox_site.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  serial = "ABCDEF3"
  local_context_data = jsonencode({"testkey3"="testvalue3"})
}
`, testName)
}

func testAccNetboxDeviceDataSourceFilterName(testName string) string {
	return fmt.Sprintf(`
data "netbox_devices" "test" {
  filter {
    name  = "name"
    value = "%[1]s_0"
  }
}`, testName)
}

const testAccNetboxDeviceDataSourceFilterTenant = `
data "netbox_devices" "test" {
  filter {
    name  = "tenant_id"
    value = netbox_tenant.test.id
  }
}`

const testAccNetBoxDeviceDataSourceFilterTagsAndStatus = `
data "netbox_devices" "tag_devices" {
  filter {
    name  = "tags"
    value = netbox_tag.test_a.name
  }
  filter {
	name  = "status"
    value = "staged"
  }
}`

const testAccNetBoxDeviceDataSourceMultipleTagsFilter = `
data "netbox_devices" "multiple_filter_devices" {
  filter {
    name  = "tags"
    value = join(",", [netbox_tag.test_b.name, netbox_tag.test_c.name])
  }
}`

const testAccNetboxDeviceDataSourceFilterRole = `
data "netbox_devices" "test" {
  filter {
    name  = "role_id"
    value = netbox_device_role.test.id
  }
}`

func testAccNetboxDeviceDataSourceNameRegex(testName string) string {
	return fmt.Sprintf(`
data "netbox_devices" "test" {
  name_regex = "%[1]s.*_regex"
}`, testName)
}

const testAccNetboxDeviceDataSourceLimit = `
data "netbox_devices" "test" {
  limit = 1
  filter {
    name  = "tenant_id"
    value = netbox_tenant.test.id
  }
}`

func TestAccNetboxDevicesDataSource_CustomFields(t *testing.T) {
	testSlug := "device_ds_customfields"
	testName := testAccGetTestName(testSlug)
	testField := strings.ReplaceAll(testAccGetTestName(testSlug), "-", "_")
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxDeviceFullDependencies(testName) + fmt.Sprintf(`
data "netbox_devices" "test" {
  depends_on = [
    netbox_device.test,
    netbox_custom_field.test,
  ]

  filter {
    name  = "name"
    value = "%[2]s"
  }
}

resource "netbox_custom_field" "test" {
  name          = "%[1]s"
  type          = "text"
  content_types = ["dcim.device"]
}

resource "netbox_device" "test" {
  name = "%[2]s"
  comments = "thisisacomment"
  description = "thisisadescription"
  tenant_id = netbox_tenant.test.id
  platform_id = netbox_platform.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  tags = ["%[2]sa"]
  site_id = netbox_site.test.id
  cluster_id = netbox_cluster.test.id
  location_id = netbox_location.test.id
  status = "staged"
  serial = "ABCDEF"
  custom_fields = {"${netbox_custom_field.test.name}" = "81"}
}
`, testField, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.name", testName),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.comments", "thisisacomment"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.description", "thisisadescription"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.serial", "ABCDEF"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.status", "staged"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.custom_fields."+testField, "81"),
				),
			},
		},
	})
}
