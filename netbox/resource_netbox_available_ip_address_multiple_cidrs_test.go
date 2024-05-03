package netbox

import (
	"fmt"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxAvailableIPAddressMultipleCidrs_basic(t *testing.T) {
	testPrefix1 := "1.1.2.0/24"
	testPrefix2 := "2.1.2.0/24"
	testIP := "1.1.2.1/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test1" {
  prefix = "%[1]s"
  status = "active"
  is_pool = false
}
resource "netbox_prefix" "test2" {
	prefix = "%[2]s"
	status = "active"
	is_pool = false
  }
resource "netbox_available_ip_address_multiple_cidrs" "test" {
  prefix_ids = [netbox_prefix.test1.id, netbox_prefix.test2.id]
  status = "active"
  dns_name = "test.mydomain.local"
  role = "loopback"
}`, testPrefix1, testPrefix2),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address_multiple_cidrs.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_available_ip_address_multiple_cidrs.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_available_ip_address_multiple_cidrs.test", "dns_name", "test.mydomain.local"),
					resource.TestCheckResourceAttr("netbox_available_ip_address_multiple_cidrs.test", "role", "loopback"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_available_ip_address_multiple_cidrs", &resource.Sweeper{
		Name:         "netbox_available_ip_address_multiple_cidrs",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamIPAddressesListParams()
			res, err := api.Ipam.IpamIPAddressesList(params, nil)
			if err != nil {
				return err
			}
			for _, ipAddress := range res.GetPayload().Results {
				if len(ipAddress.Tags) > 0 && (ipAddress.Tags[0] == &models.NestedTag{Name: strToPtr("acctest"), Slug: strToPtr("acctest")}) {
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
