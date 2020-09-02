package netbox

import (
	"fmt"
	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"log"
	"regexp"
	"testing"
)

func testAccNetboxIPAddressFullDependencies(testName string) string {
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

resource "netbox_interface" "test" {
  name = "%[1]s"
  virtual_machine_id = netbox_virtual_machine.test.id
`, testName)
}

func TestAccNetboxIPAddress_basic(t *testing.T) {

	testIP := "1.1.1.1/32"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "active"
  tags = ["acctest"]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tags.#", "1"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "reserved"
  tags = ["acctest"]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "reserved"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "dhcp"
  tags = ["acctest"]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "dhcp"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "provoke_error"
  tags = ["acctest"]
}`, testIP),
				ExpectError: regexp.MustCompile("expected status to be one of .*"),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "deprecated"
  tags = ["acctest"]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "deprecated"),
				),
			},
			{
			Config: fmt.Sprintf(`
resource "netbox_ip_address" "test" {
  ip_address = "%s"
  status = "active"
  dns_name = "mytest.example.com"
  tags = ["acctest"]
}`, testIP),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_ip_address.test", "dns_name", "mytest.example.com"),
				),
			},
			{
				ResourceName:      "netbox_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ip_address", &resource.Sweeper{
		Name:         "netbox_ip_address",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBox)
			params := ipam.NewIpamIPAddressesListParams()
			res, err := api.Ipam.IpamIPAddressesList(params, nil)
			if err != nil {
				return err
			}
			for _, ipAddress := range res.GetPayload().Results {
				if len(ipAddress.Tags) > 0 && ipAddress.Tags[0] == "acctest" {
					deleteParams := ipam.NewIpamIPAddressesDeleteParams().WithID(ipAddress.ID)
					_, err := api.Ipam.IpamIPAddressesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an ip address")
				}
			}
			return nil
		},
	})
}
