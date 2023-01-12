// Copyright (c) 2022 Cisco Systems, Inc. and its affiliates
// All rights reserved.

package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxDevicesDataSource_basic(t *testing.T) {

	testSlug := "device_ds_basic"
	testName := testAccGetTestName(testSlug)
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
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.device_type_id", "netbox_device_type.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_devices.test", "devices.0.location_id", "netbox_location.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.serial", "ABCDEF0"),
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.0.status", "staged"),
				),
			},
			{
				Config: dependencies + testAccNetboxDeviceDataSourceFilterTenant,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_devices.test", "devices.#", "4"),
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
		},
	})
}

func testAccNetboxDeviceDataSourceDependencies(testName string) string {
	return testAccNetboxDeviceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_device" "test0" {
  name = "%[1]s_0"
  comments = "this is also a comment"
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  device_type_id = netbox_device_type.test.id
  site_id = netbox_site.test.id
  platform_id = netbox_platform.test.id
  location_id = netbox_location.test.id
  serial = "ABCDEF0"
  status = "staged"
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
