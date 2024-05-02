package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxAvailableIPMultipleCIDRsAddress_basic(t *testing.T) {
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
resource "netbox_available_multiple_cidrs_ip_address" "test" {
  prefix_ids = [netbox_prefix.test.id]
  status = "active"
  dns_name = "test.mydomain.local"
  role = "loopback"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_multiple_cidrs_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_available_multiple_cidrs_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_available_multiple_cidrs_ip_address.test", "dns_name", "test.mydomain.local"),
					resource.TestCheckResourceAttr("netbox_available_multiple_cidrs_ip_address.test", "role", "loopback"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPMultipleCIDRsAddress_basic_range(t *testing.T) {
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
resource "netbox_available_multiple_cidrs_ip_address" "test_range" {
  ip_range_idsa;sklfjdjas;.lfkjsd;flk = netbox_prefix.test.id
  status = "active"
  dns_name = "test_range.mydomain.local"
}`, startAddress, endAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_multiple_cidrs_ip_address.test_range", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_available_multiple_cidrs_ip_address.test_range", "status", "active"),
					resource.TestCheckResourceAttr("netbox_available_multiple_cidrs_ip_address.test_range", "dns_name", "test_range.mydomain.local"),
				),
			},
		},
	})
}
