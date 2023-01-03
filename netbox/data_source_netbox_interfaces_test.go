package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxInterfacesDataSource_basic(t *testing.T) {

	testSlug := "interface_ds_basic"
	testResource := "data.netbox_interfaces.test"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxInterfacesDataSourceDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + testAccNetboxInterfacesDataSourceFilterName(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResource, "interfaces.#", "1"),
					resource.TestCheckResourceAttr(testResource, "interfaces.0.name", testName+"_0"),
					resource.TestCheckResourceAttr(testResource, "interfaces.0.enabled", "true"),
					resource.TestCheckResourceAttrPair(testResource, "interfaces.0.vm_id", "netbox_virtual_machine.test0", "id"),
					resource.TestCheckResourceAttrPair(testResource, "interfaces.0.id", "netbox_interface.vm0_1", "id"),
				),
			},
			{
				Config: dependencies + testAccNetboxInterfacesDataSourceFilterVM,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResource, "interfaces.#", "2"),
					resource.TestCheckResourceAttrPair(testResource, "interfaces.0.vm_id", "netbox_virtual_machine.test1", "id"),
					resource.TestCheckResourceAttrPair(testResource, "interfaces.1.vm_id", "netbox_virtual_machine.test1", "id"),
				),
			},
			{
				Config: dependencies + testAccNetboxInterfacesDataSourceNameRegex,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(testResource, "interfaces.#", "1"),
					resource.TestCheckResourceAttr(testResource, "interfaces.0.name", testName+"_2_regex"),
				),
			},
		},
	})
}

func testAccNetboxInterfacesDataSourceDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_virtual_machine" "test0" {
  name = "%[1]s_0"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_virtual_machine" "test1" {
  name = "%[1]s_1"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_interface" "vm0_1" {
  name = "%[1]s_0"
  virtual_machine_id = netbox_virtual_machine.test0.id
}

resource "netbox_interface" "vm1_1" {
  name = "%[1]s_1"
  virtual_machine_id = netbox_virtual_machine.test1.id
}

resource "netbox_interface" "vm1_2" {
  name = "%[1]s_2_regex"
  virtual_machine_id = netbox_virtual_machine.test1.id
}

`, testName)
}

const testAccNetboxInterfacesDataSourceFilterVM = `
data "netbox_interfaces" "test" {
  filter {
    name  = "vm_id"
    value = netbox_virtual_machine.test1.id
  }
}`

func testAccNetboxInterfacesDataSourceFilterName(testName string) string {
	return fmt.Sprintf(`
data "netbox_interfaces" "test" {
  filter {
    name  = "name"
    value = "%[1]s_0"
  }
}`, testName)
}

const testAccNetboxInterfacesDataSourceNameRegex = `
data "netbox_interfaces" "test" {
  name_regex = "test.*_regex"
}`
