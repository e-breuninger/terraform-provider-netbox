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
  name = "%[1]s"
  status = "active"
}
`, testName)
}
func TestAccNetboxVlan_basic(t *testing.T) {

	testSlug := "vlan_basic"
	testName := testAccGetTestName(testSlug)
	testVid := "777"
	testDescription := "Test Description"
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxPrefixFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan" "test" {
  name = "%s"
  vid = "%s"
}`, testName, testVid),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan.test", "vid", testVid),
				),
			},
			{
				Config: testAccNetboxPrefixFullDependencies(testName) + fmt.Sprintf(`
resource "netbox_vlan" "test" {
  name = "%s"
  vid = "%s"
  description = "%s"
  status = "active"
  tenant_id = netbox_tenant.test.id
  site_id = netbox_site.test.id
  tags = [netbox_tag.test.name]
}`, testName, testVid, testDescription),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_vlan.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_vlan.test", "vid", testVid),
					resource.TestCheckResourceAttr("netbox_vlan.test", "description", testDescription),
					resource.TestCheckResourceAttr("netbox_prefix.test", "status", "active"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test", "tenant_id", "netbox_tenant.test", "id"),
					resource.TestCheckResourceAttrPair("netbox_vlan.test", "site_id", "netbox_site.test", "id"),
					resource.TestCheckResourceAttr("netbox_vlan.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_vlan.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_vlan.test",
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
