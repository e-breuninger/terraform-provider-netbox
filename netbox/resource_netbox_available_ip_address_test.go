package netbox

import (
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxAvailableIPAddress_basic(t *testing.T) {
	testPrefix := "1.1.2.0/24"
	testIP := "1.1.2.1/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
	prefix = "%s"
	status = "active"
	is_pool = false
}
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "dns_name", "test.mydomain.local"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_basic_range(t *testing.T) {
	startAddress := "1.1.5.1/24"
	endAddress := "1.1.5.50/24"
	testIP := "1.1.5.1/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_range" "test" {
    start_address = "%s"
    end_address = "%s"
}
resource "netbox_available_ip_address" "test_range" {
  ip_range_id = netbox_ip_range.test.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}`, startAddress, endAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test_range", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test_range", "status", "active"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test_range", "dns_name", "test_range.mydomain.local"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_multipleIpsParallel(t *testing.T) {
	testPrefix := "1.1.3.0/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
	prefix = "%s"
	status = "active"
	is_pool = false
}
resource "netbox_available_ip_address" "test1" {
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
}
resource "netbox_available_ip_address" "test2" {
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
}
resource "netbox_available_ip_address" "test3" {
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrSet("netbox_available_ip_address.test1", "ip_address"),
					resource.TestCheckResourceAttrSet("netbox_available_ip_address.test2", "ip_address"),
					resource.TestCheckResourceAttrSet("netbox_available_ip_address.test3", "ip_address"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_multipleIpsParallel_range(t *testing.T) {
	startAddress := "1.1.6.1/24"
	endAddress := "1.1.6.50/24"
	testIP := []string{"1.1.6.1/24", "1.1.6.2/24", "1.1.6.3/24"}
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_range" "test_range" {
    start_address = "%s"
    end_address = "%s"
}
resource "netbox_available_ip_address" "test_range1" {
  ip_range_id = test_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}
resource "netbox_available_ip_address" "test_range2" {
  ip_range_id = test_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}
resource "netbox_available_ip_address" "test_range3" {
  ip_range_id = test_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}`, startAddress, endAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test1", "ip_address", testIP[0]),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test2", "ip_address", testIP[1]),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test3", "ip_address", testIP[2]),
				),
				ExpectError: regexp.MustCompile(".*"),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_multipleIpsSerial(t *testing.T) {
	testPrefix := "1.1.4.0/24"
	testIP := []string{"1.1.4.1/24", "1.1.4.2/24", "1.1.4.3/24"}
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_prefix" "test" {
	prefix = "%s"
	status = "active"
	is_pool = false
}
resource "netbox_available_ip_address" "test1" {
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
}
resource "netbox_available_ip_address" "test2" {
  depends_on = [netbox_available_ip_address.test1]
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
}
resource "netbox_available_ip_address" "test3" {
  depends_on = [netbox_available_ip_address.test2]
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test1", "ip_address", testIP[0]),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test2", "ip_address", testIP[1]),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test3", "ip_address", testIP[2]),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_multipleIpsSerial_range(t *testing.T) {
	startAddress := "1.1.7.1/24"
	endAddress := "1.1.7.50/24"
	testIP := []string{"1.1.7.1/24", "1.1.7.2/24", "1.1.7.3/24"}
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_range" "test_range" {
    start_address = "%s"
    end_address = "%s"
}
resource "netbox_available_ip_address" "test_range1" {
  ip_range_id = netbox_ip_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}
resource "netbox_available_ip_address" "test_range2" {
  depends_on = [netbox_available_ip_address.test_range1]
  ip_range_id = netbox_ip_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}
resource "netbox_available_ip_address" "test_range3" {
  depends_on = [netbox_available_ip_address.test_range2]
  ip_range_id = netbox_ip_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}`, startAddress, endAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test_range1", "ip_address", testIP[0]),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test_range2", "ip_address", testIP[1]),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test_range3", "ip_address", testIP[2]),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_customFields(t *testing.T) {
	testPrefix := "1.1.8.0/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "issue" {
	name = "issue"
	type = "url"
	content_types = ["ipam.ipaddress"]
}
resource "netbox_prefix" "test" {
	prefix = "%s"
	status = "active"
	is_pool = false
}
resource "netbox_available_ip_address" "test" {
  depends_on = [netbox_custom_field.issue]
  prefix_id = netbox_prefix.test.id
  status = "active"
  dns_name = "test.mydomain.local"
  custom_fields = {
	"${netbox_custom_field.issue.name}" = "76"
  }
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "custom_fields.issue", "76"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_available_ip_address", &resource.Sweeper{
		Name:         "netbox_available_ip_address",
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
