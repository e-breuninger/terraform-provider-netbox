package netbox

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxVirtualInterfaceFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_cluster_type" "test" {
  name = "%[1]s"
}

resource "netbox_cluster" "test" {
  name = "%[1]s"
  cluster_type_id = netbox_cluster_type.test.id
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
}

`, testName)
}

func TestAccNetboxVirtualInterface_basic(t *testing.T) {

	testSlug := "iface_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualInterfaceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_interface" "test" {
  name = "%s"
  virtual_machine_id = netbox_virtual_machine.test.id
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_interface.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVirtualInterface_mac(t *testing.T) {

	testSlug := "iface_mac"
	testMac := "00:01:02:03:04:05"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVirtualInterfaceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_interface" "test" {
  name = "%[1]s"
  virtual_machine_id = netbox_virtual_machine.test.id
  mac_address = "%[2]s"
}`, testName, testMac),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttr("netbox_interface.test", "mac_address", "00:01:02:03:04:05"),
				),
			},
			{
				ResourceName:      "netbox_interface.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckInterfaceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each interface
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_interface" {
			continue
		}

		// Retrieve our interface by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := virtualization.NewVirtualizationInterfacesReadParams().WithID(stateID)
		_, err := conn.Virtualization.VirtualizationInterfacesRead(params, nil)

		if err == nil {
			return fmt.Errorf("interface (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			errorcode := err.(*virtualization.VirtualizationInterfacesReadDefault).Code()
			if errorcode == 404 {
				return nil
			}
			return err
		}
	}
	return nil
}

func init() {
	resource.AddTestSweepers("netbox_interface", &resource.Sweeper{
		Name:         "netbox_interface",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := virtualization.NewVirtualizationInterfacesListParams()
			res, err := api.Virtualization.VirtualizationInterfacesList(params, nil)
			if err != nil {
				return err
			}
			for _, intrface := range res.GetPayload().Results {
				if strings.HasPrefix(*intrface.Name, testPrefix) {
					deleteParams := virtualization.NewVirtualizationInterfacesDeleteParams().WithID(intrface.ID)
					_, err := api.Virtualization.VirtualizationInterfacesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an interface")
				}
			}
			return nil
		},
	})
}
