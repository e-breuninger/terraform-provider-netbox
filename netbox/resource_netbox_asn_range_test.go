package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxAsnRangeDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxAsnRange_basic(t *testing.T) {
	testSlug := "asn_range"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxAsnRangeDependencies(testName) + fmt.Sprintf(`
resource "netbox_asn_range" "test" {
  name        = "%[1]s"
  slug        = "%[1]s"
  rir_id      = netbox_rir.test.id
  start       = 65000
  end         = 65100
  tenant_id   = netbox_tenant.test.id
  description = "asn range for acctest"
  comments    = "some comments"
  tags        = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_asn_range.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "slug", testName),
					resource.TestCheckResourceAttrPair("netbox_asn_range.test", "rir_id", "netbox_rir.test", "id"),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "start", "65000"),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "end", "65100"),
					resource.TestCheckResourceAttrPair("netbox_asn_range.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "description", "asn range for acctest"),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "comments", "some comments"),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_asn_range.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_asn_range.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_asn_range", &resource.Sweeper{
		Name:         "netbox_asn_range",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamAsnRangesListParams()
			res, err := api.Ipam.IpamAsnRangesList(params, nil)
			if err != nil {
				return err
			}
			for _, asnRange := range res.GetPayload().Results {
				if strings.HasPrefix(*asnRange.Name, testPrefix) {
					deleteParams := ipam.NewIpamAsnRangesDeleteParams().WithID(asnRange.ID)
					_, err := api.Ipam.IpamAsnRangesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an asn range")
				}
			}
			return nil
		},
	})
}
