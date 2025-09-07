package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/extras"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxConfigTemplate_basic(t *testing.T) {
	testName := testAccGetTestName("config_template")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_config_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	template_code = "hostname {{ name }}"
	environment_params = jsonencode({"name" = "my-hostname"})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_config_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_config_template.test", "template_code", "hostname {{ name }}"),
					resource.TestCheckResourceAttr("netbox_config_template.test", "environment_params", "{\"name\":\"my-hostname\"}"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_config_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	template_code = "hostname {{ new_var }}"
	environment_params = jsonencode({"new_var" = "my-hostname-2"})
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_config_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_config_template.test", "template_code", "hostname {{ new_var }}"),
					resource.TestCheckResourceAttr("netbox_config_template.test", "environment_params", "{\"new_var\":\"my-hostname-2\"}"),
				),
			},
			{
				ResourceName:      "netbox_config_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxConfigTemplate_tags(t *testing.T) {
	testName := testAccGetTestName("config_template_tags")
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: testAccNetboxVrfTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_config_template" "test_tags" {
  name = "%[1]s"
  template_code = "hostname test"
  tags = ["%[1]sa"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_template.test_tags", "name", testName),
					resource.TestCheckResourceAttr("netbox_config_template.test_tags", "tags.#", "1"),
					resource.TestCheckResourceAttr("netbox_config_template.test_tags", "tags.0", testName+"a"),
				),
			},
			{
				Config: testAccNetboxVrfTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_config_template" "test_tags" {
  name = "%[1]s"
  template_code = "hostname test"
  tags = ["%[1]sa", "%[1]sb"]
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_template.test_tags", "tags.#", "2"),
					resource.TestCheckResourceAttr("netbox_config_template.test_tags", "tags.0", testName+"a"),
					resource.TestCheckResourceAttr("netbox_config_template.test_tags", "tags.1", testName+"b"),
				),
			},
			{
				Config: testAccNetboxVrfTagDependencies(testName) + fmt.Sprintf(`
resource "netbox_config_template" "test_tags" {
  name = "%s"
  template_code = "hostname test"
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_config_template.test_tags", "tags.#", "0"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_config_template", &resource.Sweeper{
		Name:         "netbox_config_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			state := m.(*providerState)
			api := state.legacyAPI
			params := extras.NewExtrasConfigTemplatesListParams()
			res, err := api.Extras.ExtrasConfigTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, tmpl := range res.GetPayload().Results {
				if strings.HasPrefix(*tmpl.Name, testPrefix) {
					deleteParams := extras.NewExtrasConfigTemplatesDeleteParams().WithID(tmpl.ID)
					_, err := api.Extras.ExtrasConfigTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a config template")
				}
			}
			return nil
		},
	})
}
