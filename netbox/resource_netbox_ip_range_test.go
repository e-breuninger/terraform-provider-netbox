package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxIPRangeFullDependencies(testStartAddress string, testSlug string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_vrf" "test" {
  name = "%[1]s"
}

resource "netbox_ipam_role" "test" {
  name = "%[1]s"
  slug = "%[2]s"
}

resource "netbox_site" "test" {
  name = "%[1]s"
  status = "active"
}
`, testStartAddress, testSlug)
}
func TestAccNetboxIpRange_basic(t *testing.T) {
	testSlug := "range_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	testStartAddress := "10.0.0.1/24"
	testEndAddress := "10.0.0.50/24"
	testDescription := "Test Description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPRangeFullDependencies(testName, randomSlug) + fmt.Sprintf(`
resource "netbox_ip_range" "test_basic" {
  start_address = "%s"
  end_address = "%s"
  status = "active"
  description = "%s"
  tags = []
}`, testStartAddress, testEndAddress, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "start_address", testStartAddress),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "end_address", testEndAddress),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "status", "active"),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "description", testDescription),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "tags.#", "0"),
				),
			},
			{
				Config: testAccNetboxIPRangeFullDependencies(testName, randomSlug) + fmt.Sprintf(`
resource "netbox_ip_range" "test_basic" {
  start_address = "%s"
  end_address = "%s"
  vrf_id = netbox_vrf.test.id
  tenant_id = netbox_tenant.test.id
  status = "active"
  description = "%s"
  tags = []
}`, testStartAddress, testEndAddress, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "start_address", testStartAddress),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "end_address", testEndAddress),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "description", testDescription),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_ip_range.test_basic", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_range.test_basic", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_ip_range.test_basic", "tags.#", "0"),
				),
			},
			{
				ResourceName:      "netbox_ip_range.test_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxIpRange_with_dependencies(t *testing.T) {
	testSlug := "range_with_dependencies"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	testStartAddress := "10.0.0.51/24"
	testEndAddress := "10.0.0.100/24"
	testDescription := "Test Description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPRangeFullDependencies(testName, randomSlug) + fmt.Sprintf(`
resource "netbox_ip_range" "test_with_dependencies" {
  start_address = "%s"
  end_address = "%s"
  description = "%s"
  vrf_id = netbox_vrf.test.id
  status = "active"
  tenant_id = netbox_tenant.test.id
  role_id = netbox_ipam_role.test.id
  tags = [netbox_tag.test.name]
}`, testStartAddress, testEndAddress, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_ip_range.test_with_dependencies", "start_address", testStartAddress),
					resource.TestCheckResourceAttr("netbox_ip_range.test_with_dependencies", "end_address", testEndAddress),
					resource.TestCheckResourceAttr("netbox_ip_range.test_with_dependencies", "description", testDescription),
					resource.TestCheckResourceAttr("netbox_ip_range.test_with_dependencies", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_ip_range.test_with_dependencies", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_range.test_with_dependencies", "vrf_id", "netbox_vrf.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_ip_range.test_with_dependencies", "role_id", "netbox_ipam_role.test", "id"),
					resource.TestCheckResourceAttr("netbox_ip_range.test_with_dependencies", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_ip_range.test_with_dependencies", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_ip_range.test_with_dependencies",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_ip_range", &resource.Sweeper{
		Name:         "netbox_ip_range",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamIPRangesListParams()
			res, err := api.Ipam.IpamIPRangesList(params, nil)
			if err != nil {
				return err
			}
			for _, r := range res.GetPayload().Results {
				if strings.HasPrefix(*r.StartAddress, testPrefix) {
					deleteParams := ipam.NewIpamIPRangesDeleteParams().WithID(r.ID)
					_, err := api.Ipam.IpamIPRangesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a range")
				}
			}
			return nil
		},
	})
}
