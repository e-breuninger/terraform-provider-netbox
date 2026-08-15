package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxCustomLinkDataSource_basic(t *testing.T) {
	testSlug := "custom_link_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
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
}`, testName)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: dependencies,
			},
			{
				Config: dependencies + `
data "netbox_custom_link" "test" {
  depends_on = []

  name = netbox_custom_link.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_custom_link.test", "id", "netbox_custom_link.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "content_types.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "content_types.0", "dcim.site"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "link_text", "View site"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "link_url", "https://example.com/{{ object.name }}"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "enabled", "true"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "weight", "50"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "group_name", "example-group"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "button_class", "blue"),
					resource.TestCheckResourceAttr("data.netbox_custom_link.test", "new_window", "true"),
				),
			},
		},
	})
}
