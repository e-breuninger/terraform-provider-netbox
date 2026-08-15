package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxServiceTemplateDataSource_basic(t *testing.T) {
	testSlug := "service_template_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := testAccNetboxServiceTemplateDependencies(testName) + fmt.Sprintf(`
resource "netbox_service_template" "test" {
  name        = "%[1]s"
  protocol    = "tcp"
  ports       = [8080, 8443]
  description = "service template for acctest"
  comments    = "some comments"
  tags        = [netbox_tag.test.name]
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
data "netbox_service_template" "test" {
  depends_on = []

  name = netbox_service_template.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_service_template.test", "id", "netbox_service_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_service_template.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_service_template.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("data.netbox_service_template.test", "ports.#", "2"),
					resource.TestCheckResourceAttr("data.netbox_service_template.test", "description", "service template for acctest"),
					resource.TestCheckResourceAttr("data.netbox_service_template.test", "comments", "some comments"),
					resource.TestCheckResourceAttr("data.netbox_service_template.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_service_template.test", "tags.0", testName),
				),
			},
		},
	})
}
