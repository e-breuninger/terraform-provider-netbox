package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxL2VPNDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_route_target" "import" {
  name      = "65000:%[2]d"
  tenant_id = netbox_tenant.test.id
}

resource "netbox_route_target" "export" {
  name      = "65000:%[3]d"
  tenant_id = netbox_tenant.test.id
}`, testName, acctest.RandIntRange(1, 65000), acctest.RandIntRange(1, 65000))
}

func TestAccNetboxL2VPN_basic(t *testing.T) {
	testSlug := "l2vpn_basic"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_l2vpn" "test" {
  name = "%[1]s"
  slug = "%[1]s"
  type = "vxlan"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "slug", testName),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "type", "vxlan"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "description", ""),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "comments", ""),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "tags.#", "0"),
				),
			},
			{
				ResourceName:      "netbox_l2vpn.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxL2VPN_withDependencies(t *testing.T) {
	testSlug := "l2vpn_deps"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxL2VPNDependencies(testName) + fmt.Sprintf(`
resource "netbox_l2vpn" "test" {
  name           = "%[1]s"
  slug           = "%[1]s"
  type           = "vpws"
  identifier     = 1234
  tenant_id      = netbox_tenant.test.id
  description    = "test description"
  comments       = "test comments"
  import_targets = [netbox_route_target.import.id]
  export_targets = [netbox_route_target.export.id]
  tags           = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "type", "vpws"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "identifier", "1234"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "description", "test description"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "comments", "test comments"),
					resource.TestCheckResourceAttrPair("netbox_l2vpn.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "import_targets.#", "1"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "export_targets.#", "1"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_l2vpn.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_l2vpn.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_l2vpn", &resource.Sweeper{
		Name:         "netbox_l2vpn",
		Dependencies: []string{"netbox_l2vpn_termination"},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamL2vpnsListParams()
			res, err := api.Ipam.IpamL2vpnsList(params, nil)
			if err != nil {
				return err
			}
			for _, l2vpn := range res.GetPayload().Results {
				if strings.HasPrefix(*l2vpn.Name, testPrefix) {
					deleteParams := ipam.NewIpamL2vpnsDeleteParams().WithID(l2vpn.ID)
					_, err := api.Ipam.IpamL2vpnsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an l2vpn")
				}
			}
			return nil
		},
	})
}
