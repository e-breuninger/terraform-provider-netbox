package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxRackGroup_basic(t *testing.T) {
	testSlug := "rack_group_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}

resource "netbox_rack_group" "test" {
  name        = "%[1]s"
  slug        = "%[2]s"
  description = "This is my rack group"
  comments    = "Rack group comments"
  tags        = [netbox_tag.test.name]
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_rack_group.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_rack_group.test", "description", "This is my rack group"),
					resource.TestCheckResourceAttr("netbox_rack_group.test", "comments", "Rack group comments"),
					resource.TestCheckResourceAttr("netbox_rack_group.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_rack_group.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_rack_group.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxRackGroup_defaultSlug(t *testing.T) {
	testSlug := "rack_group_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_rack_group" "test" {
  name = "%[1]s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_rack_group.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_rack_group.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_rack_group", &resource.Sweeper{
		Name:         "netbox_rack_group",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimRackGroupsListParams()
			res, err := api.Dcim.DcimRackGroupsList(params, nil)
			if err != nil {
				return err
			}
			for _, rackGroup := range res.GetPayload().Results {
				if strings.HasPrefix(*rackGroup.Name, testPrefix) {
					deleteParams := dcim.NewDcimRackGroupsDeleteParams().WithID(rackGroup.ID)
					_, err := api.Dcim.DcimRackGroupsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a rack_group")
				}
			}
			return nil
		},
	})
}
