package netbox

import (
	"fmt"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/go-openapi/runtime"
	"github.com/hashicorp/terraform-plugin-sdk/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/terraform"
	"log"
	"strconv"
	"strings"
	"testing"
)

func testAccNetboxVirtualMachineFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_ip_address" "test" {
  ip_address = "1.1.1.1/32"
  status = "active"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_platform" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}`, testName)
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
  name = "%s"
  cluster_id = netbox_cluster.test.id
  comments = "thisisacomment"
  memory_mb = 1024
  tenant_id = netbox_tenant.test.id
  platform_id = netbox_platform.test.id
  vcpus = 4
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "cluster_id", "netbox_cluster.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_virtual_machine.test", "platform_id", "netbox_platform.test", "id"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "comments", "thisisacomment"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "memory_mb", "1024"),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "vcpus", "4"),
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
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_virtual_machine.test", "memory_mb", "0"),
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
	conn := testAccProvider.Meta().(*client.NetBox)

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
			errorcode := err.(*runtime.APIError).Response.(runtime.ClientResponse).Code()
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
  name = "%s"
  cluster_id = netbox_cluster.test.id
  tags = ["boo"]
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
  tags = ["boo", "foo"]
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

func init() {
	resource.AddTestSweepers("netbox_virtual_machine", &resource.Sweeper{
		Name:         "netbox_virtual_machine",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBox)
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
