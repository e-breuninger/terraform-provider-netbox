package netbox

import (
	"fmt"
	"log"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/fbreckle/go-netbox/netbox/models"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxAggregate_basic(t *testing.T) {
	testPrefix := "1.1.1.0/25"
	testSlug := "aggregate"
	testDesc := "test aggregate"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rir" "test" {
  name = "%s"
  slug = "%s"
}
resource "netbox_aggregate" "test" {
  prefix = "%s"
  description = "%s"
  rir_id = netbox_rir.test.id
}`, testName, randomSlug, testPrefix, testDesc),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_aggregate.test", "prefix", testPrefix),
					resource.TestCheckResourceAttr("netbox_aggregate.test", "description", testDesc),
					resource.TestCheckResourceAttrPair("netbox_aggregate.test", "rir_id", "netbox_rir.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_aggregate.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_aggregate", &resource.Sweeper{
		Name:         "netbox_aggregate",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamAggregatesListParams()
			res, err := api.Ipam.IpamAggregatesList(params, nil)
			if err != nil {
				return err
			}
			for _, prefix := range res.GetPayload().Results {
				if len(prefix.Tags) > 0 && (prefix.Tags[0] == &models.NestedTag{Name: strToPtr("acctest"), Slug: strToPtr("acctest")}) {
					deleteParams := ipam.NewIpamAggregatesDeleteParams().WithID(prefix.ID)
					_, err := api.Ipam.IpamAggregatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a aggregate")
				}
			}
			return nil
		},
	})
}
