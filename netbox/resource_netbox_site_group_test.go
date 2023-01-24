package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxSiteGroup_basic(t *testing.T) {

	testSlug := "s_grp_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "parent" {
  name = "%[1]s"
  slug = "%[2]s"
  description = "foo bar."
}

resource "netbox_site_group" "child" {
  name = "%[1]s-child"
  slug = "%[2]s-c"

  parent_id = netbox_site_group.parent.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site_group.parent", "name", testName),
					resource.TestCheckResourceAttr("netbox_site_group.parent", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_site_group.parent", "description", "foo bar."),
					resource.TestCheckResourceAttr("netbox_site_group.child", "name", fmt.Sprintf("%s-child", testName)),
					resource.TestCheckResourceAttr("netbox_site_group.child", "slug", fmt.Sprintf("%s-c", randomSlug)),
					resource.TestCheckResourceAttrPair("netbox_site_group.child", "parent_id", "netbox_site_group.parent", "id"),
				),
			},
			{
				ResourceName:      "netbox_site_group.parent",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxSiteGroup_defaultSlug(t *testing.T) {

	testSlug := "sitegrp_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_site_group" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_site_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_site_group.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_site_group", &resource.Sweeper{
		Name:         "netbox_site_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := dcim.NewDcimSiteGroupsListParams()
			res, err := api.Dcim.DcimSiteGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, siteGroup := range res.GetPayload().Results {
				if strings.HasPrefix(*siteGroup.Name, testPrefix) {
					deleteParams := dcim.NewDcimSiteGroupsDeleteParams().WithID(siteGroup.ID)
					_, err := api.Dcim.DcimSiteGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a site group")
				}
			}
			return nil
		},
	})
}
