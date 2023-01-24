package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client"
	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxTag_basic(t *testing.T) {

	testSlug := "tag_basic"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
  slug = "%s"
  color_hex = "112233"
  description = "This is a test"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tag.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tag.test", "slug", randomSlug),
					resource.TestCheckResourceAttr("netbox_tag.test", "color_hex", "112233"),
					resource.TestCheckResourceAttr("netbox_tag.test", "description", "This is a test"),
				),
			},
			{
				ResourceName:      "netbox_tag.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxTag_defaultSlug(t *testing.T) {

	testSlug := "tag_defSlug"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%s"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_tag.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_tag.test", "slug", getSlug(testName)),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_tag", &resource.Sweeper{
		Name:         "netbox_tag",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*client.NetBoxAPI)
			params := extras.NewExtrasTagsListParams()
			res, err := api.Extras.ExtrasTagsList(params, nil)
			if err != nil {
				return err
			}
			for _, tag := range res.GetPayload().Results {
				if strings.HasPrefix(*tag.Name, testPrefix) {
					deleteParams := extras.NewExtrasTagsDeleteParams().WithID(tag.ID)
					_, err := api.Extras.ExtrasTagsDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a tag")
				}
			}
			return nil
		},
	})
}
