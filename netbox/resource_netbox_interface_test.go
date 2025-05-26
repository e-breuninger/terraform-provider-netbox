package netbox

import (
	"fmt"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/virtualization"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
	log "github.com/sirupsen/logrus"
)

func testAccNetboxInterfaceFullDependencies(testName string) string {
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
}

resource "netbox_virtual_machine" "test" {
  name = "%[1]s"
  cluster_id = netbox_cluster.test.id
}

resource "netbox_vlan" "test1" {
  name = "%[1]s_vlan1"
  vid = 1001
  tags = []
}

resource "netbox_vlan" "test2" {
  name = "%[1]s_vlan2"
  vid = 1002
  tags = []
}`, testName)
}

func testAccNetboxInterfaceBasic(testName string) string {
	return fmt.Sprintf(`
resource "netbox_interface" "test" {
  name = "%s"
  virtual_machine_id = netbox_virtual_machine.test.id
  tags = ["%[1]s"]
}`, testName)
}

func testAccNetboxInterfaceOpts(testName string, testMac string, enabled string) string {
	return fmt.Sprintf(`
resource "netbox_interface" "test" {
  name = "%[1]s"
  description = "%[1]s"
  enabled = %[3]s
  mac_address = "%[2]s"
  mtu = 1440
  virtual_machine_id = netbox_virtual_machine.test.id
}`, testName, testMac, enabled)
}

func testAccNetboxInterfaceVlans(testName string) string {
	return fmt.Sprintf(`
resource "netbox_interface" "test1" {
  name = "%[1]s_1"
  mode = "access"
  untagged_vlan = netbox_vlan.test1.id
  virtual_machine_id = netbox_virtual_machine.test.id
}

resource "netbox_interface" "test2" {
  name = "%[1]s_2"
  mode = "tagged"
  tagged_vlans = [netbox_vlan.test2.id]
  untagged_vlan = netbox_vlan.test1.id
  virtual_machine_id = netbox_virtual_machine.test.id
}

resource "netbox_interface" "test3" {
  name = "%[1]s_3"
  mode = "tagged-all"
  tagged_vlans = [netbox_vlan.test1.id, netbox_vlan.test2.id]
  virtual_machine_id = netbox_virtual_machine.test.id
}`, testName)
}

func TestAccNetboxInterface_basic(t *testing.T) {
	testSlug := "iface_basic"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxInterfaceFullDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: setUp + testAccNetboxInterfaceBasic(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttr("netbox_interface.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_interface.test", "tags.0", testName),
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

func TestAccNetboxInterface_opts(t *testing.T) {
	testSlug := "iface_mac"
	testMac := "00:01:02:03:04:05"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxInterfaceFullDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: setUp + testAccNetboxInterfaceOpts(testName, testMac, "true"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_interface.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_interface.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_interface.test", "mac_address", "00:01:02:03:04:05"),
					resource.TestCheckResourceAttr("netbox_interface.test", "mtu", "1440"),
					resource.TestCheckResourceAttrPair("netbox_interface.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
				),
			},
			{
				Config: setUp + testAccNetboxInterfaceOpts(testName, testMac, "false"),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_interface.test", "description", testName),
					resource.TestCheckResourceAttr("netbox_interface.test", "enabled", "false"),
					resource.TestCheckResourceAttr("netbox_interface.test", "mac_address", "00:01:02:03:04:05"),
					resource.TestCheckResourceAttr("netbox_interface.test", "mtu", "1440"),
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

func TestAccNetboxInterface_vlans(t *testing.T) {
	testSlug := "iface_vlan"
	testName := testAccGetTestName(testSlug)
	setUp := testAccNetboxInterfaceFullDependencies(testName)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckInterfaceDestroy,
		Steps: []resource.TestStep{
			{
				Config: setUp + testAccNetboxInterfaceVlans(testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface.test1", "mode", "access"),
					resource.TestCheckResourceAttr("netbox_interface.test2", "mode", "tagged"),
					resource.TestCheckResourceAttr("netbox_interface.test3", "mode", "tagged-all"),
					resource.TestCheckResourceAttrPair("netbox_interface.test1", "untagged_vlan", "netbox_vlan.test1", "id"),
					resource.TestCheckResourceAttrPair("netbox_interface.test2", "untagged_vlan", "netbox_vlan.test1", "id"),
					resource.TestCheckResourceAttrPair("netbox_interface.test2", "tagged_vlans.0", "netbox_vlan.test2", "id"),
					resource.TestCheckResourceAttr("netbox_interface.test3", "tagged_vlans.#", "2"),
				),
			},
			{
				ResourceName:      "netbox_interface.test1",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_interface.test2",
				ImportState:       true,
				ImportStateVerify: true,
			},
			{
				ResourceName:      "netbox_interface.test3",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckInterfaceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*providerState)

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
			if errresp, ok := err.(*virtualization.VirtualizationInterfacesReadDefault); ok {
				errorcode := errresp.Code()
				if errorcode == 404 {
					return nil
				}
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
			api := m.(*providerState)
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
