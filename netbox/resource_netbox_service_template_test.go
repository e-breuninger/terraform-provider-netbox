package netbox

import (
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/ipam"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
)

func testAccNetboxServiceTemplateDependencies(testName string) string {
	return fmt.Sprintf(`
resource "netbox_tag" "test" {
  name = "%[1]s"
}
`, testName)
}

func TestAccNetboxServiceTemplate_basic(t *testing.T) {
	testSlug := "service_template"
	testName := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxServiceTemplateDependencies(testName) + fmt.Sprintf(`
resource "netbox_service_template" "test" {
  name        = "%[1]s"
  protocol    = "tcp"
  ports       = [8080, 8443]
  description = "service template for acctest"
  comments    = "some comments"
  tags        = [netbox_tag.test.name]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_service_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_service_template.test", "protocol", "tcp"),
					resource.TestCheckResourceAttr("netbox_service_template.test", "ports.#", "2"),
					resource.TestCheckResourceAttr("netbox_service_template.test", "description", "service template for acctest"),
					resource.TestCheckResourceAttr("netbox_service_template.test", "comments", "some comments"),
					resource.TestCheckResourceAttr("netbox_service_template.test", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_service_template.test", "tags.0", testName),
				),
			},
			{
				ResourceName:      "netbox_service_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_service_template", &resource.Sweeper{
		Name:         "netbox_service_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := ipam.NewIpamServiceTemplatesListParams()
			res, err := api.Ipam.IpamServiceTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, serviceTemplate := range res.GetPayload().Results {
				if strings.HasPrefix(*serviceTemplate.Name, testPrefix) {
					deleteParams := ipam.NewIpamServiceTemplatesDeleteParams().WithID(serviceTemplate.ID)
					_, err := api.Ipam.IpamServiceTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a service template")
				}
			}
			return nil
		},
	})
}
