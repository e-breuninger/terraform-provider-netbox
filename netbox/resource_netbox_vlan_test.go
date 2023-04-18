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

func testAccNetboxVlanFullDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_tenant" "test" {
  name = "%[1]s"
}

resource "netbox_site" "test" {
  name   = "%[1]s"
  status = "active"
}

resource "netbox_vlan_group" "test_group" {
	name       = "%[1]s"
	slug       = "%[1]s"
	min_vid    = 1
	max_vid    = 4094
	scope_type = "dcim.site"
	scope_id   = netbox_site.test.id
}
`, testName)
}
func TestAccNetboxVlan_basic(t *testing.T) {

	testSlug := "vlan_basic"
	testName := testAccGetTestName(testSlug)
	testVid := "777"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVlanFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan" "test_basic" {
  name = "%s"
  vid  = "%s"
  tags = []
}`, testName, testVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test_basic", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan.test_basic", "vid", testVid),
					resource.TestCheckResourceAttr("netbox_vlan.test_basic", "status", "active"),
					resource.TestCheckResourceAttr("netbox_vlan.test_basic", "description", ""),
					resource.TestCheckResourceAttr("netbox_vlan.test_basic", "tags.#", "0"),
				),
			},
			{
				ResourceName:      "netbox_vlan.test_basic",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxVlan_with_dependencies(t *testing.T) {

	testSlug := "vlan_with_dependencies"
	testName := testAccGetTestName(testSlug)
	testVid := "666"
	testDescription := "Test Description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVlanFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan" "test_with_dependencies" {
  name        = "%s"
  vid         = "%s"
  description = "%s"
  status      = "active"
  tenant_id   = netbox_tenant.test.id
  site_id     = netbox_site.test.id
  group_id    = netbox_vlan_group.test_group.id
  tags        = [netbox_tag.test.name]
}`, testName, testVid, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test_with_dependencies", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan.test_with_dependencies", "vid", testVid),
					resource.TestCheckResourceAttr("netbox_vlan.test_with_dependencies", "description", testDescription),
					resource.TestCheckResourceAttr("netbox_vlan.test_with_dependencies", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test_with_dependencies", "group_id", "netbox_vlan_group.test_group", "id"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test_with_dependencies", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test_with_dependencies", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_vlan.test_with_dependencies", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_vlan.test_with_dependencies", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_vlan.test_with_dependencies",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_vlan", &resource.Sweeper{
		Name:         "netbox_vlan",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := ipam.NewIpamVlansListParams()
			res, err := api.Ipam.IpamVlansList(params, nil)
			if err != nil {
				return err
			}
			for _, vlan := range res.GetPayload().Results {
				if strings.HasPrefix(*vlan.Name, testPrefix) {
					deleteParams := ipam.NewIpamVlansDeleteParams().WithID(vlan.ID)
					_, err := api.Ipam.IpamVlansDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a vlan")
				}
			}
			return nil
		},
	})
}
