package netbox

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
)

func testAccNetboxServiceFullDependencies(testName string) string {
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

func TestAccNetboxService_basic(t *testing.T) {

	testSlug := "svc_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		PreCheck:     func() { testAccPreCheck(t) },
		Providers:    testAccProviders,
		CheckDestroy: testAccCheckServiceDestroy,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxServiceFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_service" "test" {
  name = "%s"
  virtual_machine_id = netbox_virtual_machine.test.id
  ports = [666]
  protocol = "tcp"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_service.test", "virtual_machine_id", "netbox_virtual_machine.test", "id"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.#", "1"),
					resource.TestCheckResourceAttr("netbox_service.test", "ports.0", "666"),
					resource.TestCheckResourceAttr("netbox_service.test", "protocol", "tcp"),
				),
			},
			{
				ResourceName:      "netbox_service.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func testAccCheckServiceDestroy(s *terraform.State) error {
	// retrieve the connection established in Provider configuration
	conn := testAccProvider.Meta().(*client.NetBoxAPI)

	// loop through the resources in state, verifying each service
	// is destroyed
	for _, rs := range s.RootModule().Resources {
		if rs.Type != "netbox_service" {
			continue
		}

		// Retrieve our service by referencing it's state ID for API lookup
		stateID, _ := strconv.ParseInt(rs.Primary.ID, 10, 64)
		params := ipam.NewIpamServicesReadParams().WithID(stateID)
		_, err := conn.Ipam.IpamServicesRead(params, nil)

		if err == nil {
			return fmt.Errorf("service (%s) still exists", rs.Primary.ID)
		}

		if err != nil {
			if errresp, ok := err.(*ipam.IpamServicesReadDefault); ok {
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
	resource.AddTestSweepers("netbox_service", &resource.Sweeper{
		Name:         "netbox_service",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamServicesListParams()
			res, err := api.Ipam.IpamServicesList(params, nil)
			if err != nil {
				return err
			}
			for _, intrface := range res.GetPayload().Results {
				if strings.HasPrefix(*intrface.Name, testPrefix) {
					deleteParams := ipam.NewIpamServicesDeleteParams().WithID(intrface.ID)
					_, err := api.Ipam.IpamServicesDelete(deleteParams, nil)
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
