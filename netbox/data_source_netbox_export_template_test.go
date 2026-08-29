package netbox

import (
	"fmt"
	"testing"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxExportTemplateDataSource_basic(t *testing.T) {
	testSlug := "export_template_ds"
	testName := testAccGetTestName(testSlug)
	dependencies := fmt.Sprintf(`
resource "netbox_export_template" "test" {
  name           = "%[1]s"
  content_types  = ["dcim.site"]
  description    = "test description"
  template_code  = "{%% for obj in queryset %%}{{ obj.name }}\n{%% endfor %%}"
  mime_type      = "text/plain"
  file_extension = "txt"
  as_attachment  = true
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
data "netbox_export_template" "test" {
  depends_on = []

  name = netbox_export_template.test.name
}`,
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttrPair("data.netbox_export_template.test", "id", "netbox_export_template.test", "id"),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "name", testName),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "content_types.#", "1"),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "content_types.0", "dcim.site"),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "description", "test description"),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "template_code", "{% for obj in queryset %}{{ obj.name }}\n{% endfor %}"),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "mime_type", "text/plain"),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "file_extension", "txt"),
					resource.TestCheckResourceAttr("data.netbox_export_template.test", "as_attachment", "true"),
				),
			},
		},
	})
}
