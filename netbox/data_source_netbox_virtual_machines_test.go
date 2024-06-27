package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxVirtualMachinesDataSource_basic(t *testing.T) {
	testSlug := "vm_ds_basic"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxVirtualMachineDataSourceDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxVirtualMachineDataSourceFilterName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.0.name", testName+"_0"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.0.vcpus", "4"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.0.memory_mb", "1024"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.0.disk_size_gb", "256"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.0.comments", "thisisacomment"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.platform_slug", "netbox_platform.test", "slug"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.device_id", "netbox_device.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.device_name", "netbox_device.test", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxVirtualMachineDataSourceFilterCluster,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.#", "4"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.name", "netbox_virtual_machine.test0", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.1.name", "netbox_virtual_machine.test1", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.2.name", "netbox_virtual_machine.test2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.3.name", "netbox_virtual_machine.test3", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxVirtualMachineDataSourceFilterTenantID,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.name", "netbox_virtual_machine.test0", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxVirtualMachineDataSourceNameRegex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.#", "2"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.name", "netbox_virtual_machine.test2", "name"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.1.name", "netbox_virtual_machine.test3", "name"),
				),
			},
			{
				Config: dependencies + testAccNetboxVirtualMachineDataSourceLimit,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test", "vms.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test", "vms.0.cluster_id", "netbox_cluster.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxVirtualMachinesDataSource_tags(t *testing.T) {
	testSlug := "vm_ds_tags"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxVirtualMachineDataSourceDependenciesWithTags(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxVirtualMachineDataSourceTagA(testName) + testAccNetboxVirtualMachineDataSourceTagB(testName) + testAccNetboxVirtualMachineDataSourceTagAB(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.tag-a", "vms.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.tag-b", "vms.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.tag-ab", "vms.#", "1"),
				),
			},
		},
	})
}

func TestAccNetboxVirtualMachinesDataSource_status(t *testing.T) {
	testSlug := "vm_ds_tags"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxVirtualMachineDataSourceDependenciesWithStatus(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxVirtualMachineDataSourceStatusActive + testAccNetboxVirtualMachineDataSourceStatusDecommissioning,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test_active", "vms.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_virtual_machines.test_decommissioning", "vms.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test_active", "vms.0.status", "netbox_virtual_machine.test0", "status"),
					resource.TestCheckResourceAttrPair("data.netbox_virtual_machines.test_decommissioning", "vms.0.status", "netbox_virtual_machine.test1", "status"),
				),
			},
		},
	})
}

func testAccNetboxVirtualMachineDataSourceDependencies(testName string) string {
	return testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test0" {
  name = "%[1]s_0"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
  device_id = netbox_device.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_gb = 256
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  platform_id = netbox_platform.test.id
  vcpus = 4
}

resource "netbox_virtual_machine" "test1" {
  name = "%[1]s_1"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
}

resource "netbox_virtual_machine" "test2" {
  name = "%[1]s_2_regex"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
}

resource "netbox_virtual_machine" "test3" {
  name = "%[1]s_3_regex"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
}
`, testName)
}

func testAccNetboxVirtualMachineDataSourceDependenciesWithTags(testName string) string {
	return testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_tag" "servicea" {
	name      = "%[1]s_service-a"
}

resource "netbox_tag" "serviceb" {
	name      = "%[1]s_service-b"
}

resource "netbox_virtual_machine" "test0" {
  name = "%[1]s_0"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_gb = 256
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  platform_id = netbox_platform.test.id
  vcpus = 4
  tags = [
		netbox_tag.servicea.name,
		netbox_tag.serviceb.name,
	]
}

resource "netbox_virtual_machine" "test1" {
  name = "%[1]s_1"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
	tags = [
		netbox_tag.servicea.name,
	]
}

resource "netbox_virtual_machine" "test2" {
  name = "%[1]s_2_regex"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
	tags = [
		netbox_tag.serviceb.name,
	]
}
`, testName)
}

const testAccNetboxVirtualMachineDataSourceFilterCluster = `
data "netbox_virtual_machines" "test" {
  filter {
    name  = "cluster_id"
    value = netbox_cluster.test.id
  }
}`

const testAccNetboxVirtualMachineDataSourceFilterTenantID = `
data "netbox_virtual_machines" "test" {
  filter {
    name  = "tenant_id"
    value = netbox_tenant.test.id
  }
}`

func testAccNetboxVirtualMachineDataSourceFilterName(testName string) string {
	return fmt.Sprintf(`
data "netbox_virtual_machines" "test" {
  filter {
    name  = "name"
    value = "%[1]s_0"
  }
}`, testName)
}

const testAccNetboxVirtualMachineDataSourceNameRegex = `
data "netbox_virtual_machines" "test" {
  name_regex = "test.*_regex"
}`

const testAccNetboxVirtualMachineDataSourceLimit = `
data "netbox_virtual_machines" "test" {
  limit = 1
  filter {
    name  = "cluster_id"
    value = netbox_cluster.test.id
  }
}`

func testAccNetboxVirtualMachineDataSourceTagA(testName string) string {
	return fmt.Sprintf(`
	data "netbox_virtual_machines" "tag-a" {
		filter {
			name  = "tag"
			value = "%[1]s_service-a"
		}
	}`, testName)
}

func testAccNetboxVirtualMachineDataSourceTagB(testName string) string {
	return fmt.Sprintf(`
data "netbox_virtual_machines" "tag-b" {
  filter {
    name  = "tag"
    value = "%[1]s_service-b"
	}
}`, testName)
}

func testAccNetboxVirtualMachineDataSourceTagAB(testName string) string {
	return fmt.Sprintf(`
data "netbox_virtual_machines" "tag-ab" {
	filter {
    name  = "tag"
    value = "%[1]s_service-a"
	}
  filter {
    name  = "tag"
    value = "%[1]s_service-b"
	}
}`, testName)
}

const testAccNetboxVirtualMachineDataSourceStatusActive = `
data "netbox_virtual_machines" "test_active" {
  filter {
    name  = "status"
    value = "active"
  }
}`

const testAccNetboxVirtualMachineDataSourceStatusDecommissioning = `
data "netbox_virtual_machines" "test_decommissioning" {
  filter {
    name  = "status"
    value = "decommissioning"
  }
}`

func testAccNetboxVirtualMachineDataSourceDependenciesWithStatus(testName string) string {
	return testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_tag" "servicea" {
	name      = "%[1]s_service-a"
}

resource "netbox_virtual_machine" "test0" {
  name = "%[1]s_0"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_gb = 256
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  platform_id = netbox_platform.test.id
  vcpus = 4
	status = "active"
  tags = [
		netbox_tag.servicea.name,
	]
}

resource "netbox_virtual_machine" "test1" {
  name = "%[1]s_1"
  cluster_id = netbox_cluster.test.id
  site_id = netbox_site.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_gb = 256
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  platform_id = netbox_platform.test.id
  vcpus = 4
	status = "decommissioning"
  tags = [
		netbox_tag.servicea.name,
	]
}`, testName)
}
