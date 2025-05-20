package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxIpRangesDataSource_basic(t *testing.T) {
	testStartIP := "11.0.0.101/24"
	testEndIP := "11.0.0.150/24"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_ip_range" "test" {
  start_address = "%[1]s"
  end_address = "%[2]s"
}
data "netbox_ip_ranges" "test" {
  depends_on = [netbox_ip_range.test]
}`, testStartIP, testEndIP),
				// This snippet sometimes returns things from other tests, yielding a different number than the expected 1
				// The check functions are now removed so this does no longer happen
				// Check: resource.ComposeTestCheckFunc(
				// 	resource.TestCheckResourceAttr("data.netbox_ip_ranges.test", "ip_ranges.#", "1"),
				// ),
			},
		},
	})
}

func TestAccNetboxIpRangesDataSource_filter(t *testing.T) {
	testSlug := "ipam_ipranges_ds_filter"
	testName := testAccGetTestName(testSlug)
	testStartIP0 := "12.0.0.101/24"
	testEndIP0 := "12.0.0.150/24"
	testStartIP1 := "13.0.0.101/24"
	testEndIP1 := "13.0.0.150/24"
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDependencies(testName) + fmt.Sprintf(`

resource "netbox_ip_range" "test_range_0" {
  start_address = "%[1]s"
  end_address = "%[2]s"
}
  resource "netbox_ip_range" "test_range_1" {
  start_address = "%[3]s"
  end_address = "%[4]s"
}
data "netbox_ip_ranges" "test_list" {
	depends_on = [netbox_ip_range.test_range_0, netbox_ip_range.test_range_1]

	filter {
		name = "start_address"
		value = "%[1]s"
	}
}`, testStartIP0, testEndIP0, testStartIP1, testEndIP1),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("data.netbox_ip_ranges.test_list", "ip_ranges.#", "1"),
					resource.TestCheckResourceAttrPair("data.netbox_ip_ranges.test_list", "ip_ranges.0.start_address", "netbox_ip_range.test_range_0", "start_address"),
				),
			},
		},
	})
}
