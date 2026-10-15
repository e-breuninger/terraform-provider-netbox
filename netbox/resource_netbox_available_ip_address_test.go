package netbox

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	"github.com/hashicorp/terraform-plugin-sdk/v2/terraform"
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
  role = "loopback"
}`, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "dns_name", "test.mydomain.local"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "role", "loopback"),
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

func TestAccNetboxAvailableIPAddress_deviceByObjectType(t *testing.T) {
	startAddress := "1.2.7.1/24"
	endAddress := "1.2.7.50/24"
	testSlug := "av_ipa_dev_ot"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDeviceDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_range" "test_range" {
  start_address = "%s"
  end_address = "%s"
}
resource "netbox_available_ip_address" "test" {
  ip_range_id = netbox_ip_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
  object_type = "dcim.interface"
  interface_id = netbox_device_interface.test.id
}`, startAddress, endAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "object_type", "dcim.interface"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "interface_id", "netbox_device_interface.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_deviceByFieldName(t *testing.T) {
	startAddress := "1.3.7.1/24"
	endAddress := "1.3.7.50/24"
	testSlug := "av_ipa_dev_fn"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxIPAddressFullDeviceDependencies(testName) + fmt.Sprintf(`
resource "netbox_ip_range" "test_range" {
  start_address = "%s"
  end_address = "%s"
}
resource "netbox_available_ip_address" "test" {
  ip_range_id = netbox_ip_range.test_range.id
  status = "active"
  dns_name = "test_range.mydomain.local"
  device_interface_id = netbox_device_interface.test.id
}`, startAddress, endAddress),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "device_interface_id", "netbox_device_interface.test", "id"),
				),
			},
		},
	})
}

func TestAccNetboxAvailableIPAddress_cf(t *testing.T) {
	testPrefix := "1.1.8.0/24"
	testIP := "1.1.8.1/24"
	testSlug := "av_ipa_cf"

	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_field" "test" {
  name   = "%s"
  type   = "text"
  weight = 100
  content_types = ["ipam.ipaddress"]
}

resource "netbox_prefix" "test" {
  prefix = "%s"
  status = "active"
  is_pool = false
}

resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
  status = "active"
  custom_fields = {
    "${netbox_custom_field.test.name}" = "test-field"
  }
}`, testSlug, testPrefix),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", testIP),
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", fmt.Sprintf("custom_fields.%s", testSlug), "test-field"),
				),
			},
			{
				ResourceName:      "netbox_available_ip_address.test",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{
					"prefix_id",
				},
			},
		},
	})
}

// TestAccNetboxAvailableIPAddress_adoptExistingIP imports an address allocated
// out of band - the only way to manage one - and asserts it survives: same
// NetBox record, prefix recorded in state, empty plan afterwards. Moving it to
// another prefix from there must still re-allocate.
func TestAccNetboxAvailableIPAddress_adoptExistingIP(t *testing.T) {
	prefixes := `
resource "netbox_prefix" "test" {
  prefix  = "1.1.9.0/24"
  status  = "active"
  is_pool = false
}

resource "netbox_prefix" "other" {
  prefix  = "1.1.13.0/24"
  status  = "active"
  is_pool = false
}`

	withAdoptedIP := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.test.id
}`

	movedToOtherPrefix := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.other.id
}`

	var prefixID, adoptedIPID int64

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			// Prefixes only; remember the id to allocate from.
			{
				Config: prefixes,
				Check: func(s *terraform.State) (err error) {
					prefixID, err = testAccStateID(s, "netbox_prefix.test")
					return err
				},
			},
			// Allocate behind Terraform's back, the same way the resource
			// does, then adopt it.
			{
				PreConfig: func() {
					api := testAccProvider.Meta().(*providerState)
					params := ipam.NewIpamPrefixesAvailableIpsCreateParams().WithID(prefixID).
						WithData([]*models.AvailableIP{{}})
					res, err := api.Ipam.IpamPrefixesAvailableIpsCreate(params, nil)
					if err != nil {
						t.Fatalf("allocating an out-of-band IP in prefix %d: %s", prefixID, err)
					}
					if len(res.Payload) == 0 {
						t.Fatalf("no available IP addresses in prefix %d", prefixID)
					}
					adoptedIPID = res.Payload[0].ID
				},
				Config:             withAdoptedIP,
				ResourceName:       "netbox_available_ip_address.test",
				ImportState:        true,
				ImportStatePersist: true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strconv.FormatInt(adoptedIPID, 10), nil
				},
				ImportStateCheck: testAccCheckImportedAllocationSourceUnset("prefix_id"),
			},
			// Records the prefix in state, keeps the address. Unfixed, this
			// destroys the imported address and allocates another.
			{
				Config: withAdoptedIP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "prefix_id", "netbox_prefix.test", "id"),
					func(s *terraform.State) error {
						id, err := testAccStateID(s, "netbox_available_ip_address.test")
						if err != nil {
							return err
						}
						if id != adoptedIPID {
							return fmt.Errorf("expected the adopted address (NetBox id %d) to be kept, got id %d", adoptedIPID, id)
						}
						return nil
					},
				),
			},
			// Converged.
			{
				Config:   withAdoptedIP,
				PlanOnly: true,
			},
			// Source known now, so a change re-allocates as usual.
			{
				Config: movedToOtherPrefix,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", "1.1.13.1/24"),
					func(s *terraform.State) error {
						id, err := testAccStateID(s, "netbox_available_ip_address.test")
						if err != nil {
							return err
						}
						if id == adoptedIPID {
							return fmt.Errorf("expected the address to be replaced, but it is still NetBox id %d", adoptedIPID)
						}
						return nil
					},
				),
			},
		},
	})
}

// TestAccNetboxAvailableIPAddress_adoptExistingIPFromRange is the ip_range_id
// twin of TestAccNetboxAvailableIPAddress_adoptExistingIP.
func TestAccNetboxAvailableIPAddress_adoptExistingIPFromRange(t *testing.T) {
	rangeOnly := `
resource "netbox_ip_range" "test" {
  start_address = "1.1.12.1/24"
  end_address   = "1.1.12.50/24"
}`

	withAdoptedIP := rangeOnly + `
resource "netbox_available_ip_address" "test" {
  ip_range_id = netbox_ip_range.test.id
}`

	var rangeID, adoptedIPID int64

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: rangeOnly,
				Check: func(s *terraform.State) (err error) {
					rangeID, err = testAccStateID(s, "netbox_ip_range.test")
					return err
				},
			},
			{
				PreConfig: func() {
					api := testAccProvider.Meta().(*providerState)
					params := ipam.NewIpamIPRangesAvailableIpsCreateParams().WithID(rangeID).
						WithData([]*models.AvailableIP{{}})
					res, err := api.Ipam.IpamIPRangesAvailableIpsCreate(params, nil)
					if err != nil {
						t.Fatalf("allocating an out-of-band IP in range %d: %s", rangeID, err)
					}
					if len(res.Payload) == 0 {
						t.Fatalf("no available IP addresses in range %d", rangeID)
					}
					adoptedIPID = res.Payload[0].ID
				},
				Config:             withAdoptedIP,
				ResourceName:       "netbox_available_ip_address.test",
				ImportState:        true,
				ImportStatePersist: true,
				ImportStateIdFunc: func(*terraform.State) (string, error) {
					return strconv.FormatInt(adoptedIPID, 10), nil
				},
				ImportStateCheck: testAccCheckImportedAllocationSourceUnset("ip_range_id"),
			},
			{
				Config: withAdoptedIP,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("netbox_available_ip_address.test", "ip_range_id", "netbox_ip_range.test", "id"),
					func(s *terraform.State) error {
						id, err := testAccStateID(s, "netbox_available_ip_address.test")
						if err != nil {
							return err
						}
						if id != adoptedIPID {
							return fmt.Errorf("expected the adopted address (NetBox id %d) to be kept, got id %d", adoptedIPID, id)
						}
						return nil
					},
				),
			},
			{
				Config:   withAdoptedIP,
				PlanOnly: true,
			},
		},
	})
}

// TestAccNetboxAvailableIPAddress_allocationSourceChangeReallocates guards the
// other side: a real prefix change on a provider-created address re-allocates.
func TestAccNetboxAvailableIPAddress_allocationSourceChangeReallocates(t *testing.T) {
	prefixes := `
resource "netbox_prefix" "first" {
  prefix  = "1.1.10.0/24"
  status  = "active"
  is_pool = false
}

resource "netbox_prefix" "second" {
  prefix  = "1.1.11.0/24"
  status  = "active"
  is_pool = false
}`

	fromFirst := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.first.id
}`

	fromSecond := prefixes + `
resource "netbox_available_ip_address" "test" {
  prefix_id = netbox_prefix.second.id
}`

	var firstIPID int64

	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fromFirst,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", "1.1.10.1/24"),
					func(s *terraform.State) (err error) {
						firstIPID, err = testAccStateID(s, "netbox_available_ip_address.test")
						return err
					},
				),
			},
			{
				Config: fromSecond,
				Check: resource.ComposeTestCheckFunc(
					// New address and new record; an in-place update would
					// have kept both.
					resource.TestCheckResourceAttr("netbox_available_ip_address.test", "ip_address", "1.1.11.1/24"),
					func(s *terraform.State) error {
						secondIPID, err := testAccStateID(s, "netbox_available_ip_address.test")
						if err != nil {
							return err
						}
						if secondIPID == firstIPID {
							return fmt.Errorf("expected the IP address to be replaced, but it is still NetBox id %d", firstIPID)
						}
						return nil
					},
				),
			},
		},
	})
}

// testAccCheckImportedAllocationSourceUnset asserts an import could not
// populate the given allocation source, which is the root of the problem.
func testAccCheckImportedAllocationSourceUnset(attribute string) resource.ImportStateCheckFunc {
	return func(states []*terraform.InstanceState) error {
		for _, is := range states {
			if got := is.Attributes[attribute]; got != "" && got != "0" {
				return fmt.Errorf("expected imported %s to be unset, got %q", attribute, got)
			}
		}
		return nil
	}
}

// testAccStateID returns a resource's NetBox id from state.
func testAccStateID(s *terraform.State, address string) (int64, error) {
	rs, ok := s.RootModule().Resources[address]
	if !ok {
		return 0, fmt.Errorf("%s not found in state", address)
	}
	return strconv.ParseInt(rs.Primary.ID, 10, 64)
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
			api := m.(*providerState)
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
