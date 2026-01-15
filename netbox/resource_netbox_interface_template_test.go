package netbox

import (
	"fmt"
	"strings"
	"testing"

	"github.com/fbreckle/go-netbox/netbox/client/dcim"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	log "github.com/sirupsen/logrus"
)

func TestAccNetboxInterfaceTemplate_basic(t *testing.T) {
	testSlug := "interface_template"
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

resource "netbox_interface_template" "test" {
	name = "%[1]s"
	device_type_id = netbox_device_type.test.id
	type = "100base-tx"
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "type", "100base-tx"),
					resource.TestCheckResourceAttrPair("netbox_interface_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				ResourceName:      "netbox_interface_template.test",
				ImportState:       true,
				ImportStateVerify: true,
			},
		},
	})
}

func TestAccNetboxInterfaceTemplate_opts(t *testing.T) {
	testSlug := "interface_template"
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

resource "netbox_interface_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	device_type_id = netbox_device_type.test.id
	type = "100base-tx"
	mgmt_only = true
}`, testName, randomSlug),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "label", fmt.Sprintf("%s label", testName)),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "type", "100base-tx"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "mgmt_only", "true"),
					resource.TestCheckResourceAttrPair("netbox_interface_template.test", "device_type_id", "netbox_device_type.test", "id"),
				),
			},
			{
				Config: fmt.Sprintf(`
resource "netbox_manufacturer" "test" {
	name = "%[1]s"
}

resource "netbox_module_type" "test" {
	manufacturer_id = netbox_manufacturer.test.id
	model           = "%[1]s"
}

resource "netbox_interface_template" "test" {
	name = "%[1]s"
	description = "%[1]s description"
	label = "%[1]s label"
	module_type_id = netbox_module_type.test.id
	type = "100base-tx"
	mgmt_only = false
}`, testName),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr("netbox_interface_template.test", "name", testName),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "description", fmt.Sprintf("%s description", testName)),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "label", fmt.Sprintf("%s label", testName)),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "type", "100base-tx"),
					resource.TestCheckResourceAttr("netbox_interface_template.test", "mgmt_only", "false"),
					resource.TestCheckResourceAttrPair("netbox_interface_template.test", "module_type_id", "netbox_module_type.test", "id"),
				),
			},
		},
	})
}

func init() {
	resource.AddTestSweepers("netbox_interface_template", &resource.Sweeper{
		Name:         "netbox_interface_template",
		Dependencies: []string{},
		F: func(region string) error {
			m, err := sharedClientForRegion(region)
			if err != nil {
				return fmt.Errorf("Error getting client: %s", err)
			}
			api := m.(*providerState)
			params := dcim.NewDcimInterfaceTemplatesListParams()
			res, err := api.Dcim.DcimInterfaceTemplatesList(params, nil)
			if err != nil {
				return err
			}
			for _, tmpl := range res.GetPayload().Results {
				if strings.HasPrefix(*tmpl.Name, testPrefix) {
					deleteParams := dcim.NewDcimInterfaceTemplatesDeleteParams().WithID(tmpl.ID)
					_, err := api.Dcim.DcimInterfaceTemplatesDelete(deleteParams, nil)
					if err != nil {
						return err
					}
					log.Print("[DEBUG] Deleted an interface template")
				}
			}
			return nil
		},
	})
}
