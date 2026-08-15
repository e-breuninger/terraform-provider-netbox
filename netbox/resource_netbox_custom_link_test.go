package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCustomLink_basic(t *testing.T) {
	testSlug := "custom_link"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_link" "test" {
  name          = "%[1]s"
  content_types = ["dcim.site"]
  link_text     = "View site"
  link_url      = "https://example.com/{{ object.name }}"
  enabled       = true
  weight        = 50
  group_name    = "example-group"
  button_class  = "blue"
  new_window    = true
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_link.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "content_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "content_types.0", "dcim.site"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "link_text", "View site"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "link_url", "https://example.com/{{ object.name }}"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "enabled", "true"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "weight", "50"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "group_name", "example-group"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "button_class", "blue"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "new_window", "true"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_custom_link" "test" {
  name          = "%[1]s"
  content_types = ["dcim.site", "dcim.device"]
  link_text     = "View site updated"
  link_url      = "https://example.com/updated/{{ object.name }}"
  enabled       = false
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_custom_link.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "content_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "link_text", "View site updated"),
					resource.TestCheckResourceAttr("netbox_custom_link.test", "enabled", "false"),
				),
			},
			{
				ResourceName:      "netbox_custom_link.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_custom_link", &resource.Sweeper{
		Name:         "netbox_custom_link",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := extras.NewExtrasCustomLinksListParams()
			res, err := api.Extras.ExtrasCustomLinksList(params, nil)
			if err != nil {
				return err
			}
			for _, customLink := range res.GetPayload().Results {
				if strings.HasPrefix(*customLink.Name, testPrefix) {
					deleteParams := extras.NewExtrasCustomLinksDeleteParams().WithID(customLink.ID)
					_, err := api.Extras.ExtrasCustomLinksDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a custom link")
				}
			}
			return nil
		},
	})
}
