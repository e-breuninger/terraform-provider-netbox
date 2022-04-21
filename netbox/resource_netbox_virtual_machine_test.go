package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccNetboxVirtualMachineFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
  site_id = netbox_site.test.id
}

resource "netbox_device_role" "test" {
  name = "%[1]s"
  color_hex = "123456"
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}

resource "netbox_tag" "test_a" {
  name = "%[1]sa"
}

resource "netbox_tag" "test_b" {
  name = "%[1]sb"
}
`, testName)
}

func TestAccNetboxVirtualMachine_basic(t *testing.T) {

	testSlug := "vm_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualMachineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
				),
			},
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  disk_size_gb = 256
  tenant_id = netbox_tenant.test.id
  role_id = netbox_device_role.test.id
  platform_id = netbox_platform.test.id
  vcpus = 4
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "role_id", "netbox_device_role.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "memory_mb", "1024"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "4"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "disk_size_gb", "256"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "tags.0", testName+"a"),
				),
			},
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
  tenant_id = netbox_tenant.test.id
  platform_id = netbox_platform.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "0"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "memory_mb", "0"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "disk_size_gb", "0"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "comments", ""),
				),
			},
			{
				ResourceName:      "netbox_virtual_machine.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVirtualMachine_fractionalVcpu(t *testing.T) {

	testSlug := "vm_fracVcpu"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckVirtualMachineDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
  vcpus = 2.50
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "2.5"),
				),
			},
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
  vcpus = 4
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "4"),
				),
			},
			{
				ResourceName:      "netbox_virtual_machine.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckVirtualMachineDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each virtual machine
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_virtual_machine" {
			continue
		}

		// Retrieve our virtual machine by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := virtualization.NewVirtualizationVirtualMachinesReadParams().WithID(stateID)
		_, err := conn.Virtualization.VirtualizationVirtualMachinesRead(params, nil)

		if err == nil {
			return fmt.Errorf("virtual machine (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			errorcode := err.(*virtualization.VirtualizationVirtualMachinesReadDefault).Code()
			if errorcode == 404 {
				return nil
			}
			return err
		}
	}
	return nil
}

func TestAccNetboxVirtualMachine_tags(t *testing.T) {

	testSlug := "vm_tags"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "tags.#", "1"),
				),
			},
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
  tags = ["%[1]sa", "%[1]sb"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "tags.#", "2"),
				),
			},
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_virtual_machine" "test" {
  name = "%s"
  cluster_id = netbox_cluster.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "tags.#", "0"),
				),
			},
		},
	})
}

func TestAccNetboxVirtualMachine_customFields(t *testing.T) {
	testSlug := "vm_cf"
	testName := testAccGetTestName(testSlug)
	resource.Test(t, resource.TestCase{
		PreCheck:  func() { testAccPreCheck(t) },
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualMachineFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_custom_field" "test" {
	name          = "custom_field"
	type          = "text"
	content_types = ["virtualization.virtualmachine"]
}
resource "netbox_virtual_machine" "test" {
  name          = "%[1]s"
  cluster_id    = netbox_cluster.test.id
  custom_fields = {"${netbox_custom_field.test.name}" = "76"}
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "custom_fields.custom_field", "76"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_virtual_machine", &resource.Sweeper{
		Name:         "netbox_virtual_machine",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := virtualization.NewVirtualizationVirtualMachinesListParams()
			res, err := api.Virtualization.VirtualizationVirtualMachinesList(params, nil)
			if err != nil {
				return err
			}
			for _, virtualMachine := range res.GetPayload().Results {
				if strings.HasPrefix(*virtualMachine.Name, testPrefix) {
					deleteParams := virtualization.NewVirtualizationVirtualMachinesDeleteParams().WithID(virtualMachine.ID)
					_, err := api.Virtualization.VirtualizationVirtualMachinesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a virtual machine")
				}
			}
			return nil
		},
	})
}
