package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func TestAccNetboxExportTemplate_basic(t *testing.T) {
	testSlug := "export_template"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_export_template" "test" {
  name           = "%[1]s"
  content_types  = ["dcim.site"]
  description    = "test description"
  template_code  = "{%% for obj in queryset %%}{{ obj.name }}\n{%% endfor %%}"
  mime_type      = "text/plain"
  file_extension = "txt"
  as_attachment  = true
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_export_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_export_template.test", "content_types.#", "1"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "content_types.0", "dcim.site"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "description", "test description"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "template_code", "{% for obj in queryset %}{{ obj.name }}\n{% endfor %}"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "mime_type", "text/plain"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "file_extension", "txt"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "as_attachment", "true"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_export_template" "test" {
  name          = "%[1]s"
  content_types = ["dcim.site", "dcim.device"]
  template_code = "{%% for obj in queryset %%}{{ obj.name }}-updated\n{%% endfor %%}"
  as_attachment = false
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_export_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_export_template.test", "content_types.#", "2"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "template_code", "{% for obj in queryset %}{{ obj.name }}-updated\n{% endfor %}"),
					resource.TestCheckResourceAttr("netbox_export_template.test", "as_attachment", "false"),
				),
			},
			{
				ResourceName:      "netbox_export_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_export_template", &resource.Sweeper{
		Name:         "netbox_export_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := extras.NewExtrasExportTemplatesListParams()
			res, err := api.Extras.ExtrasExportTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, exportTemplate := range res.GetPayload().Results {
				if strings.HasPrefix(*exportTemplate.Name, testPrefix) {
					deleteParams := extras.NewExtrasExportTemplatesDeleteParams().WithID(exportTemplate.ID)
					_, err := api.Extras.ExtrasExportTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an export template")
				}
			}
			return nil
		},
	})
}
