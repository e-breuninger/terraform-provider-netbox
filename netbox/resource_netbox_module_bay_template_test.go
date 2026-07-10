package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxModuleBayTemplate_basic(t *testing.T) {
	testSlug := "module_bay_template"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_module_bay_template" "test" {
	name = "%[1]s"
	device_type_id = netbox_device_type.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "name", testName),
					resource.TestCheckResourceAttrPair("netbox_module_bay_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_module_bay_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxModuleBayTemplate_opts(t *testing.T) {
	testSlug := "module_bay_template"
	testName := testAccGetTestName(testSlug)
	randomSlug := testAccGetTestName(testSlug)
	resource.ParallelTest(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck(t) },
		Steps: []resource.TestStep{
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_module_bay_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	position = "1"
	device_type_id = netbox_device_type.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "label", fmt.Sprintf("%s label", testName)),
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "position", "1"),
					resource.TestCheckResourceAttrPair("netbox_module_bay_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_device_type" "test" {
	model = "%[1]s"
	slug = "%[2]s"
	part_number = "%[2]s"
	manufacturer_id = netbox_manufacturer.test.id
}

resource "netbox_module_bay_template" "test" {
	name = "%[1]s"
	description = "%[1]s new description"
	label = "%[1]s new label"
	position = "2"
	device_type_id = netbox_device_type.test.id
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "description", fmt.Sprintf("%s new description", testName)),
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "label", fmt.Sprintf("%s new label", testName)),
					resource.TestCheckResourceAttr("netbox_module_bay_template.test", "position", "2"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_module_bay_template", &resource.Sweeper{
		Name:         "netbox_module_bay_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimModuleBayTemplatesListParams()
			res, err := api.Dcim.DcimModuleBayTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, tmpl := range res.GetPayload().Results {
				if strings.HasPrefix(*tmpl.Name, testPrefix) {
					deleteParams := dcim.NewDcimModuleBayTemplatesDeleteParams().WithID(tmpl.ID)
					_, err := api.Dcim.DcimModuleBayTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted a module bay template")
				}
			}
			return nil
		},
	})
}
